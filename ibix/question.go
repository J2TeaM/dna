package ibix

import (
	"dna"
	"dna/http"
)

type Question struct {
	Cat             Category
	Description     dna.String
	QuestionContent dna.String
	OptionA         dna.String
	OptionB         dna.String
	OptionC         dna.String
	OptionD         dna.String
	OptionE         dna.String
	Answer          dna.String
	Explaination    dna.String
	ExerciseNo      dna.Int
}

func NewQuestion() *Question {
	qt := new(Question)
	qt.Cat = Category{}
	qt.Description = ""
	qt.QuestionContent = ""
	qt.OptionA = ""
	qt.OptionB = ""
	qt.OptionC = ""
	qt.OptionD = ""
	qt.OptionE = ""
	qt.Answer = ""
	qt.Explaination = ""
	qt.ExerciseNo = 0
	return qt
}

func (q Question) ToRecord() []string {

	if q.Cat.Subject == "non-verbal-reasoning/questions-and-answers" {
		imagesArr := q.QuestionContent.FindAllString(`<img.+?src=.+?>`, -1)
		images := dna.StringArray{}
		for _, image := range imagesArr {
			images.Push("http://indiabix.com" + image.GetTagAttributes("src"))
		}
		return []string{"Non Verbal Reasoning", q.Cat.Name.String(), q.Description.String(), images.Join(",").String(), q.QuestionContent.String(), q.OptionA.String(), q.OptionB.String(), q.OptionC.String(), q.OptionD.String(), q.OptionE.String(), q.Answer.String(), q.Explaination.String(), q.ExerciseNo.ToString().String()}
	} else {
		return []string{q.Cat.Subject.String(), q.Cat.Name.String(), q.Description.String(), q.QuestionContent.String(), q.OptionA.String(), q.OptionB.String(), q.OptionC.String(), q.OptionD.String(), q.OptionE.String(), q.Answer.String(), q.Explaination.String(), q.ExerciseNo.ToString().String()}
	}
}

type Questions []Question

func NewQuestions() *Questions {
	return new(Questions)
}

func getOptionsForQuestion(tableOptions dna.String, question *Question) {
	// dna.Log(tableOptions)

	options := tableOptions.Replace(`<div class="bix-div-answer"`, "").Trim().Split("tdOptionDt")
	if options.Length() > 6 {
		panic("Answer options greater than 5 at: " + question.Cat.ToString())
	}
	for idx, option := range options {
		if idx > 0 {
			option := option.ReplaceWithRegexp(`^.+?>`, ``).ReplaceWithRegexp(`(?mis)<td class="bix-td-option".+$`, "").ReplaceWithRegexp(`</td>$`, "").ReplaceWithRegexp(`</td></tr><tr>$`, "").ReplaceWithRegexp(`</td></tr></table><input type="hidden".+$`, "").Trim()
			switch idx {
			case 1:
				question.OptionA = option
			case 2:
				question.OptionB = option
			case 3:
				question.OptionC = option
			case 4:
				question.OptionD = option
			case 5:
				question.OptionE = option
			}
		}
	}

}

func getQuestions(cat Category, isFirstPage dna.Bool, pageUrl dna.String, pageLinks *dna.StringArray) <-chan Questions {
	channel := make(chan Questions)
	questions := Questions{}
	go func() {
		var url dna.String
		if isFirstPage == true {
			url = cat.Url
		} else {
			url = pageUrl
		}
		link := "http://www.indiabix.com" + url
		// dna.Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data

			if isFirstPage == true {
				// getting page links if the first page found
				pageLinkArr := data.FindAllString(`<p class="ib-pager">.+</p>`, 1)
				if pageLinkArr.Length() > 0 {
					links := pageLinkArr[0].FindAllString(`<a href.+?>`, -1)
					for _, link := range links {
						pageLinks.Push(link.GetTagAttributes("href"))
					}
				}
			}

			// Getting direction for the questions
			descriptionArr := data.FindAllString(`(?mis)id="divDirectionText".+?<div class="bix-div-container">`, 1)
			description := dna.String("")
			if descriptionArr.Length() > 0 {
				description = descriptionArr[0].ReplaceWithRegexp(`<div class="bix-div-container">$`, "").Trim().ReplaceWithRegexp(`<td style="padding-left:5px" valign="top">$`, "").ReplaceWithRegexp(`id="divDirectionText">?`, "").Trim().ReplaceWithRegexp(`<tr>$`, "").Trim().ReplaceWithRegexp(`</div></div></td></tr>$`, "").Trim().ReplaceWithRegexp(`^<p>`, "").Trim().ReplaceWithRegexp(`</p>$`, "").Trim()
			}

			questionTables := data.FindAllString(`(?mis)<table class="bix-tbl-container".+?<hr />`, -1)
			// if questionTables.Length() > 0 {
			// 	dna.Log("# of questions: ", questionTables.Length())
			// }
			for _, questionTable := range questionTables {
				question := NewQuestion()
				question.Cat = cat
				question.Description = description

				// Getting question content
				questionContentArr := questionTable.FindAllString(`(?mis)<td class="bix-td-qtxt".+?<td class="bix-td-miscell"`, 1)
				if questionContentArr.Length() > 0 {
					question.QuestionContent = questionContentArr[0].ReplaceWithRegexp(`^<td class="bix-td-qtxt.+?>`, "").ReplaceWithRegexp(`<td class="bix-td-miscell"`, "").Trim().ReplaceWithRegexp(`^<p>`, "").Trim().ReplaceWithRegexp(`<tr>$`, "").Trim().ReplaceWithRegexp(`</tr>$`, "").Trim().ReplaceWithRegexp(`</p></td>$`, "").Trim()
				}
				// Getting question answers
				// answersArr is table html contains a list of options for the question
				optionArr := questionTable.FindAllString(`(?mis)<td class="bix-td-option.+tdOptionDt.+?<div class="bix-div-answer"`, -1)
				getOptionsForQuestion(optionArr[0], question)

				// Getting explaination for the question
				explantionArr := questionTable.FindAllString(`(?mis)Explanation:</b></span></p>.+?<div class="bix-div-workspace"`, 1)
				if explantionArr.Length() > 0 {
					question.Explaination = explantionArr[0].ReplaceWithRegexp(`^Explanation:</b></span></p>`, "").Trim().ReplaceWithRegexp(`^<p>`, "").Trim().ReplaceWithRegexp(`<div class="bix-div-workspace"$`, "").Trim().ReplaceWithRegexp(`</div>$`, "").Trim().ReplaceWithRegexp(`</div>$`, "").Trim().ReplaceWithRegexp(`</p>$`, "").Trim()
				}

				questions = append(questions, *question)

			}

			// Getting answers for all question
			// If length of answers != length of questions, raise an error
			answersArr := data.FindAllString(`<input id="hdnAjaxImageCacheKey".+`, -1)
			var answers = dna.StringArray{}
			if answersArr.Length() > 0 {
				answer := answersArr[0].GetTagAttributes("value")
				// $('input' + '#' + 'hdn' + 'Ajax' + 'Image' + 'Cache' + 'Key').val().substr(18).split('').reverse().join('').substr(17).toUpperCase().split('');
				tmp := answer.Substring(18, answer.Length()).Split("").Reverse().Join("")
				answers = tmp.Substring(17, tmp.Length()).ToUpperCase().Split("")
			}

			if answers.Length().ToPrimitiveValue() != len(questions) {
				panic("Answer length differs from question length - " + cat.ToString())
			}
			for idx, ans := range answers {
				questions[idx].Answer = ans
			}

		}
		channel <- questions

	}()
	return channel
}

func GetQuestions(cat Category, execiseNo dna.Int) (Questions, error) {

	finalQuestions := []Question{}

	c := make(chan Questions)
	var pageLinks = &dna.StringArray{}
	go func() {
		c <- <-getQuestions(cat, true, "", pageLinks)
	}()

	for i := 0; i < 1; i++ {
		qts := <-c
		for _, qt := range qts {
			finalQuestions = append(finalQuestions, qt)
		}

	}

	// dna.Log(pageLinks.Unique())

	for _, pageLink := range pageLinks.Unique() {
		go func(pageLink dna.String) {
			c <- <-getQuestions(cat, false, pageLink, pageLinks)
		}(pageLink)
	}

	for i := 0; i < pageLinks.Unique().Length().ToPrimitiveValue(); i++ {
		qts := <-c
		for _, qt := range qts {
			finalQuestions = append(finalQuestions, qt)
		}
	}

	for idx, _ := range finalQuestions {
		finalQuestions[idx].ExerciseNo = execiseNo
	}
	return finalQuestions, nil
}
