package csn

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
)

//SongVideo returns a song or video type
type SongVideo struct {
	Id   dna.Int
	Item interface{}
}

func NewSongVideo() *SongVideo {
	return &SongVideo{0, nil}
}

// GetSongVideo returns an interface of Song or Video instance
func GetSongVideo(id dna.Int) (interface{}, error) {
	song, err := GetSong(id)
	if err != nil {
		return nil, err
	} else {
		if song.Type == IS_VIDEO {
			return song.ToVideo(), nil
		} else {
			return song, nil
		}
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (sv *SongVideo) Fetch() error {
	item, err := GetSongVideo(sv.Id)
	if err != nil {
		return err
	} else {
		sv.Item = item
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (sv *SongVideo) GetId() dna.Int {
	return sv.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (sv *SongVideo) New() item.Item {
	return item.Item(NewSongVideo())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (sv *SongVideo) Init(v interface{}) {
	switch v.(type) {
	case int:
		sv.Id = dna.Int(v.(int))
	case dna.Int:
		sv.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (sv *SongVideo) Save(db *sqlpg.DB) error {
	if sv.Item != nil {
		return db.InsertIgnore(sv.Item)
	} else {
		return errors.New("Item field of struct is nil")
	}
}
