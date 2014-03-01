package ke

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
)

type Lyric struct {
	Id      dna.Int
	Content dna.String
}

func NewLyric() *Lyric {
	lyric := new(Lyric)
	lyric.Id = 0
	lyric.Content = ""
	return lyric
}

func GetLyric(id dna.Int) (*Lyric, error) {
	apiLyric, err := GetAPILyric(id)
	if err != nil {
		return nil, err
	} else {
		if apiLyric.Data != "" {
			lrc := NewLyric()
			lrc.Id = id
			lrc.Content = apiLyric.Data
			return lrc, nil
		} else {
			return nil, errors.New(dna.Sprintf("Keeng - Lyric ID: %v not found", id).String())
		}
	}
}

func (lyric *Lyric) Fetch() error {
	_lyric, err := GetLyric(lyric.Id)
	if err != nil {
		return err
	} else {
		*lyric = *_lyric
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (lyric *Lyric) GetId() dna.Int {
	return lyric.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (lyric *Lyric) New() item.Item {
	return item.Item(NewLyric())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (lyric *Lyric) Init(v interface{}) {
	switch v.(type) {
	case int:
		lyric.Id = dna.Int(v.(int))
	case dna.Int:
		lyric.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (lyric *Lyric) Save(db *sqlpg.DB) error {
	song := NewSong()
	song.Id = lyric.Id
	song.Lyric = lyric.Content
	song.HasLyric = true
	return db.Update(song, "id", "lyric", "has_lyric")
}
