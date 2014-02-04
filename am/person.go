package am

import (
	"dna"
)

type Person struct {
	Id   dna.Int
	Name dna.String
}

func NewPerson() *Person {
	person := new(Person)
	person.Id = 0
	person.Name = ""
	return person
}
