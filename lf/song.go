package lf

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"time"
)

// Song defines a song struct whose ID is form AMG.
type Song struct {
	Id           dna.Int
	Title        dna.String
	Artistids    dna.IntArray // New
	Artists      dna.StringArray
	Albumid      dna.Int    // New
	AlbumTitle   dna.String // New
	Instrumental dna.Bool
	Viewable     dna.Bool
	HasLrc       dna.Bool
	LrcVerified  dna.Bool
	Lyricid      dna.Int // New
	Lyric        dna.String
	Duration     dna.Int // New
	Copyright    dna.String
	Writers      dna.String
	DateUpdated  time.Time
	Checktime    time.Time
}

func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Artistids = dna.IntArray{}
	song.Artists = dna.StringArray{}
	song.Albumid = 0
	song.AlbumTitle = ""
	song.Instrumental = false
	song.Viewable = false
	song.HasLrc = false
	song.LrcVerified = false
	song.Lyricid = 0
	song.Lyric = ""
	song.Duration = 0
	song.Copyright = ""
	song.Writers = ""
	song.DateUpdated = time.Time{}
	song.Checktime = time.Now()
	return song
}

// func GetSong(id dna.Int) (*Song, error) {
// 	apisong, err := GetAPISong(id)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		song := NewSong()
// 		apisong.FillSong(song)
// 		return song, nil
// 	}
// }

// Fetch implements item.Item interface.
// Returns error if can not get item
func (song *Song) Fetch() error {
	return nil
}

// GetId implements GetId methods of item.Item interface
func (song *Song) GetId() dna.Int {
	return song.Id
}

func (song *Song) New() item.Item {
	return item.Item(NewAPIFullSong())
}

func (song *Song) Init(v interface{}) {
}
func GetSong(id dna.Int) (*Song, error) {
	return nil, nil
}
func (song *Song) Save(db *sqlpg.DB) error {
	return nil
}
