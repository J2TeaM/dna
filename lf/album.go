package lf

import (
	"dna"
	"time"
)

type Album struct {
	Id         dna.Int
	Title      dna.String
	Artistid   dna.Int
	Artists    dna.StringArray
	Year       dna.String
	Importance dna.Int
	Image      dna.String
	LargeImage dna.String
	Checktime  time.Time
}

func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Title = ""
	album.Artistid = 0
	album.Artists = dna.StringArray{}
	album.Year = ""
	album.Importance = 0
	album.Image = ""
	album.LargeImage = ""
	album.Checktime = time.Now()
	return album
}
