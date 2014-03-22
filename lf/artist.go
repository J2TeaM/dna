package lf

import (
	"dna"
	"time"
)

type Artist struct {
	Id         dna.Int
	Name       dna.String
	Genres     dna.StringArray
	Popularity dna.Int
	Image      dna.String
	Checktime  time.Time
}

func NewArtist() *Artist {
	artist := new(Artist)
	artist.Id = 0
	artist.Name = ""
	artist.Genres = dna.StringArray{}
	artist.Popularity = 0
	artist.Image = ""
	artist.Checktime = time.Now()
	return artist
}

func (artist *Artist) Equal(cprArtist *Artist) dna.Bool {
	if artist.Id == cprArtist.Id {
		return true
	} else {
		return false
	}
}
