package ml

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"os"
	"sync"
	"time"
)

type Song struct {
	Id             dna.Int
	Title          dna.String
	Artists        dna.StringArray
	Url            dna.String
	Lyric          dna.String
	Status         dna.Int
	Gracenote      dna.Int
	Publishers     dna.String
	Writers        dna.String
	LineCount      dna.Int
	LineTimestamps dna.IntArray
	Checktime      time.Time
}

func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Artists = dna.StringArray{}
	song.Url = ""
	song.Lyric = ""
	song.Status = 0
	song.Gracenote = 0
	song.Publishers = ""
	song.Writers = ""
	song.LineCount = 0
	song.LineTimestamps = dna.IntArray{}
	song.Checktime = time.Now()
	return song
}

func GetSong(id dna.Int) (*Song, error) {
	apisong, err := GetAPISong(id)
	if err != nil {
		return nil, err
	} else {
		song := NewSong()
		apisong.FillSong(song)
		return song, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (song *Song) Fetch() error {
	_song, err := GetSong(song.Id)
	if err != nil {
		return err
	} else {
		*song = *_song
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (song *Song) GetId() dna.Int {
	return song.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (song *Song) New() item.Item {
	return item.Item(NewSong())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (song *Song) Init(v interface{}) {
	switch v.(type) {
	case int:
		song.Id = dna.Int(v.(int))
	case dna.Int:
		song.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

var (
	mutex      = &sync.Mutex{}
	SongFile   *os.File
	TotalBytes dna.Int
)

func (song *Song) Save(db *sqlpg.DB) error {
	insertStmt := sqlpg.GetInsertStatement("mlsongs", song, false) + "\n"
	mutex.Lock()
	n, err := SongFile.WriteString(insertStmt.String())
	TotalBytes += dna.Int(n)
	mutex.Unlock()
	if err == nil {
		return nil
	} else {
		return err
	}

}
