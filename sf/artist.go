package sf

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
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
		artist.Checktime.Format("2006-02-01 15:04:05"),
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (artist *Artist) Fetch() error {
	return nil
}

// GetId implements GetId methods of item.Item interface
func (artist *Artist) GetId() dna.Int {
	return 0
}

// New implements item.Item interface
// Returns new item.Item interface
func (artist *Artist) New() item.Item {
	return item.Item(NewArtist())
}

func (artist *Artist) Init(v interface{}) {
	// do nothing
}

func (artist *Artist) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(artist)
}
