package ibix

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"encoding/csv"
	"os"
	"sync"
)

var out = csv.NewWriter(os.Stdout)
var mutex = &sync.Mutex{}
var SubjectList = dna.StringArray{"Aptitude", "Logical Reasoning", "Verbal Ability", "General Knowledge"}
var Engineering = dna.StringArray{"Engineering"}
var MainList = dna.StringArray{}

type Category struct {
	Subject dna.String
	Name    dna.String
	Url     dna.String
}

func NewCategory() *Category {
	cat := new(Category)
	cat.Subject = ""
	cat.Name = ""
	cat.Url = ""
	return cat
}

func (cat Category) ToString() string {
	return string(cat.Subject + " - " + cat.Name + " - " + cat.Url)
}

type Subject struct {
	Id   dna.Int
	Name dna.String
	Cats []*Category
}

func NewSubject(id dna.Int, name dna.String) *Subject {
	sub := new(Subject)
	sub.Id = id
	sub.Name = name
	sub.Cats = []*Category{}
	return sub
}

func getSongFormats(subject *Subject) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.indiabix.com/" + subject.Name.ToDashCase().ToLowerCase()
		// dna.Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			tables := data.FindAllString(`(?mis)<table width="100%" id="ib-tbl-topics">.+?</table>`, -1)
			for _, table := range tables {
				cells := table.FindAllString(`<td>.+?</td>`, -1)
				if cells.Length() > 0 {
					for _, cell := range cells {
						cat := &Category{subject.Name, cell.RemoveHtmlTags(""), cell.GetTagAttributes("href")}
						subject.Cats = append(subject.Cats, cat)
						// dna.Log(cat)
					}
				}
			}
		}
		channel <- true

	}()
	return channel
}

func GetSubject(idx dna.Int) (*Subject, error) {

	subject := NewSubject(idx, MainList[idx])
	c := make(chan bool)
	go func() {
		c <- <-getSongFormats(subject)
	}()

	for i := 0; i < 1; i++ {
		<-c
	}
	return subject, nil
}

func GetSubjectByName(name dna.String) (*Subject, error) {

	subject := NewSubject(0, name)
	c := make(chan bool)

	go func() {
		c <- <-getSongFormats(subject)
	}()

	for i := 0; i < 1; i++ {
		<-c
	}
	return subject, nil
}

func (subject *Subject) Fetch() error {
	_subject, err := GetSubject(subject.Id)
	if err != nil {
		return err
	} else {
		*subject = *_subject
		return nil
	}
}

func (subject *Subject) GetId() dna.Int {
	return subject.Id
}

func (subject *Subject) New() item.Item {
	return item.Item(NewSubject(0, ""))
}

func (subject *Subject) Init(v interface{}) {
	switch v.(type) {
	case int:
		subject.Id = dna.Int(v.(int))
	case dna.Int:
		subject.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (subject *Subject) Save(db *sqlpg.DB) error {
	mutex.Lock()
	for _, cat := range subject.Cats {
		record := make([]string, 3)
		record[0] = cat.Subject.String()
		record[1] = cat.Name.String()
		record[2] = cat.Url.String()
		out.Write(record)
		out.Flush()
	}
	mutex.Unlock()
	return nil
}
