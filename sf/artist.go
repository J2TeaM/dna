package sf

import (
	"dna"
	"time"
)

type Artist struct {
	// Auto increased id
	Id        dna.Int
	AMG       dna.Int
	Name      dna.String
	Genres    dna.StringArray
	UrlSlug   dna.String
	Image     dna.String
	Rating    dna.IntArray
	Bio       dna.String
	Checktime time.Time
}

// NewArtist return default new artist
func NewArtist() *Artist {
	artist := new(Artist)
	artist.Id = 0
	artist.AMG = 0
	artist.UrlSlug = ""
	artist.Image = ""
	artist.Genres = dna.StringArray{}
	artist.Name = ""
	artist.Rating = dna.IntArray{0, 0, 0}
	artist.Bio = "{}"
	artist.Checktime = time.Time{}
	return artist
}

//CSVRecord returns a record to write csv format.
//
//psql -c "COPY sfartists (id,amg,name,genres,url_slug,image,rating,bio,checktime) FROM '/Users/daonguyenanbinh/Box Documents/Sites/golang/sfartists.csv' DELIMITER ',' CSV"
func (artist *Artist) CSVRecord() []string {
	return []string{
		artist.Id.ToString().String(),
		artist.AMG.ToString().String(),
		artist.Name.String(),
		dna.Sprintf("%#v", artist.Genres).Replace("dna.StringArray", "").String(),
		artist.UrlSlug.String(),
		artist.Image.String(),
		dna.Sprintf("%#v", artist.Rating).Replace("dna.IntArray", "").String(),
		artist.Bio.String(),
		artist.Checktime.Format("2006-01-02 15:04:05"),
	}
}
