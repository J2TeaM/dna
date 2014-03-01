package nct

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"time"
)

// Artist defines an artist structure.
type Artist struct {
	Id        dna.Int
	Name      dna.String
	Avatar    dna.String
	NSongs    dna.Int
	NAlbums   dna.Int
	NVideos   dna.Int
	Checktime time.Time
}

// NewArtist returns a new artist.
func NewArtist() *Artist {
	artist := new(Artist)
	artist.Id = 0
	artist.Name = ""
	artist.Avatar = ""
	artist.NSongs = 0
	artist.NAlbums = 0
	artist.NVideos = 0
	artist.Checktime = time.Now()
	return artist
}

// GetArtist returns an artist given an id or an error if available-
func GetArtist(id dna.Int) (*Artist, error) {
	apiArtist, err := GetAPIArtist(id)
	if err == nil {
		return apiArtist.ToArtist(), nil
	} else {
		return nil, err
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (artist *Artist) Fetch() error {
	_artist, err := GetArtist(artist.Id)
	if err != nil {
		return err
	} else {
		*artist = *_artist
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (artist *Artist) GetId() dna.Int {
	return artist.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (artist *Artist) New() item.Item {
	return item.Item(NewArtist())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (artist *Artist) Init(v interface{}) {
	switch v.(type) {
	case int:
		artist.Id = dna.Int(v.(int))
	case dna.Int:
		artist.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (artist *Artist) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(artist)
}
