package am

import (
	"dna"
)

type Discography struct {
	Id       dna.Int
	Title    dna.String
	Coverart dna.String
}

type Category struct {
	Id   dna.Int
	Name dna.String
}

type Person struct {
	Id   dna.Int
	Name dna.String
}

type Song struct {
	Id        dna.Int
	Title     dna.String
	Artists   []Person
	Composers []Person
	Duration  dna.Int
}

type Credit struct {
	Id     dna.Int
	Artist dna.String
	Job    dna.String
}

//[{"average":81.428571428571,"count":7,"itemId":"MW0002585207"}]
type AverageRating struct {
	Average dna.Float  `json:"average"`
	Count   dna.Int    `json:"count"`
	ItemId  dna.String `json:"itemId"`
}
