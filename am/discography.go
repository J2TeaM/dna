package am

import (
	"dna"
)

type APIDiscography struct {
	Id       dna.Int
	Title    dna.String
	Coverart dna.String
}

func (apidisco *APIDiscography) ToDiscography() *Discography {
	disco := NewDiscography()
	disco.Id = apidisco.Id
	disco.Title = apidisco.Title
	disco.Coverart = apidisco.Coverart
	return disco
}

type Discography struct {
	Id       dna.Int
	Title    dna.String
	Artistid dna.Int
	Coverart dna.String
}

func NewDiscography() *Discography {
	disco := new(Discography)
	disco.Id = 0
	disco.Artistid = 0
	disco.Title = ""
	disco.Coverart = ""
	return disco
}

func (disco *Discography) CSVRecord() []string {
	return []string{
		disco.Id.ToString().String(),
		disco.Title.String(),
		disco.Artistid.ToString().String(),
		disco.Coverart.String(),
	}
}
