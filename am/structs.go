package am

import (
	"dna"
)

//AverageRating defines average rating
//correspondent to 2nd and 3rd elements of Ratings field.
//[{"average":81.428571428571,"count":7,"itemId":"MW0002585207"}]
type AverageRating struct {
	Average dna.Float  `json:"average"`
	Count   dna.Int    `json:"count"`
	ItemId  dna.String `json:"itemId"`
}

// APIAwardSection defines section for group of awards
// such as Billboard Albums
type APIAwardSection struct {
	Name   dna.String
	Type   dna.String
	Awards []APIAward
}
