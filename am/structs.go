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

//AverageRating defines average rating
//correspondent to 2nd and 3rd elements of Ratings field.
//[{"average":81.428571428571,"count":7,"itemId":"MW0002585207"}]
type AverageRating struct {
	Average dna.Float  `json:"average"`
	Count   dna.Int    `json:"count"`
	ItemId  dna.String `json:"itemId"`
}

type Release struct {
	Id     dna.Int
	Title  dna.String
	Format dna.String
	Year   dna.Int
	Label  dna.String
}

type Award struct {
	Id    dna.Int
	Title dna.String
	Year  dna.Int
	Chart dna.String
	Peak  dna.Int
	Type  dna.Int // 0: undefined, 1 : album, 2 song, 3 song & album
	Award dna.String
	// If Chart is empty, then Award is the name if the award
	// artists reveive
	Winners []Person
}

// AwardSection defines section for group of awards
// such as Billboard Albums
type AwardSection struct {
	Name   dna.String
	Type   dna.String
	Awards []Award
}
