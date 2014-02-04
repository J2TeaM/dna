package am

import (
	"dna"
)

type Category struct {
	Id   dna.Int
	Name dna.String
}

func NewCategory() *Category {
	category := new(Category)
	category.Id = 0
	category.Name = ""
	return category
}
