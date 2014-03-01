package ke

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

// Album defines main Video type.
type Album struct {
	Id          dna.Int
	Key         dna.String
	Title       dna.String
	Artists     dna.StringArray
	Nsongs      dna.Int
	Plays       dna.Int
	Coverart    dna.String
	Description dna.String
	Songids     dna.IntArray
	DateCreated time.Time
	Checktime   time.Time
}

// NewAlbum returns new album.
func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Key = ""
	album.Title = ""
	album.Artists = dna.StringArray{}
	album.Plays = 0
	album.Songids = dna.IntArray{}
	album.Nsongs = 0
	album.Description = ""
	album.Coverart = ""
	album.DateCreated = time.Time{}
	album.Checktime = time.Time{}
	return album
}

// GetAlbum returns a album.
func GetAlbum(id dna.Int) (*Album, error) {
	apiAlbum, err := GetAPIAlbum(id)
	if err != nil {
		return nil, err
	} else {
		album := apiAlbum.ToAlbum()
		if album.Id == 0 {
			return nil, errors.New(dna.Sprintf("Keeng - Album ID: %v not found", id).String())
		} else {
			return album, nil
		}
	}
}

// Fetch implements item.Item interface
func (album *Album) Fetch() error {
	_album, err := GetAlbum(album.Id)
	if err != nil {
		return err
	} else {
		*album = *_album
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (album *Album) GetId() dna.Int {
	return album.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (album *Album) New() item.Item {
	return item.Item(NewAlbum())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (album *Album) Init(v interface{}) {
	switch v.(type) {
	case int:
		album.Id = dna.Int(v.(int))
	case dna.Int:
		album.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (album *Album) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(album)
}
