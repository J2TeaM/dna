package ibix

import (
	"dna"
	"dna/http"
)

type Exercise struct {
	Cat Category
	No  dna.Int
}

func (e Exercise) ToRecords() []string {
	ret := []string{}
	ret = append(ret, e.Cat.Subject.String())
	ret = append(ret, e.Cat.Name.String())
	ret = append(ret, e.Cat.Url.String())
	ret = append(ret, e.No.ToString().String())
	return ret
}

func NewExercise() *Exercise {
	exercise := new(Exercise)
	exercise.Cat = Category{}
	exercise.No = 0
	return exercise
}

type Exercises []Exercise

func getExercises(cat Category) <-chan Exercises {
	channel := make(chan Exercises)
	exercises := Exercises{}
	go func() {
		link := "http://www.indiabix.com" + cat.Url
		// dna.Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			exercisesArr := data.FindAllStringSubmatch(`(?mis)ib-lefttbar-container.+Exercise(.+?)id="ib-main-bar"`, 1)
			// dna.Log(exercisesArr)
			if len(exercisesArr) > 0 {
				sectionArr := exercisesArr[0][1].FindAllString(`<a href=.+?</a>`, -1)
				for _, section := range sectionArr {
					exc := NewExercise()
					exc.Cat = cat
					exc.Cat.Url = section.GetTagAttributes("href")
					exc.No = exc.Cat.Url.ReplaceWithRegexp(`/$`, "").ReplaceWithRegexp(`^.+/`, "").ToInt()
					exercises = append(exercises, *exc)
					// dna.Log(exc.Cat.Url.ReplaceWithRegexp(`/$`, "").Replace(`^.+/`, ""))
				}
			}
		}
		channel <- exercises

	}()
	return channel
}

func GetExercises(cat Category) (Exercises, error) {

	finalExercises := Exercises{Exercise{Cat: cat, No: 1}}

	c := make(chan Exercises)

	go func() {
		c <- <-getExercises(cat)
	}()

	for i := 0; i < 1; i++ {
		qts := <-c
		for _, qt := range qts {
			finalExercises = append(finalExercises, qt)
		}

	}

	return finalExercises, nil
}
