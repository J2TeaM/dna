package ke

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

// Song defines main Video type.
type Song struct {
	Id                dna.Int
	Key               dna.String
	Title             dna.String
	Artists           dna.StringArray
	Plays             dna.Int
	ListenType        dna.Int
	Lyric             dna.String
	Link              dna.String
	MediaUrlMono      dna.String
	MediaUrlPre       dna.String
	DownloadUrl       dna.String
	IsDownload        dna.Int
	RingbacktoneCode  dna.String
	RingbacktonePrice dna.Int
	// Url               dna.String
	Price       dna.Int
	Copyright   dna.Int
	CrbtId      dna.Int
	Coverart    dna.String
	Coverart310 dna.String
	DateCreated time.Time
	Checktime   time.Time
}

// Newsong returns new song.
func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Artists = dna.StringArray{}
	song.Plays = 0
	song.ListenType = 0
	song.Lyric = ""
	song.Link = ""
	song.MediaUrlMono = ""
	song.MediaUrlPre = ""
	song.DownloadUrl = ""
	song.IsDownload = 0
	song.RingbacktoneCode = ""
	song.RingbacktonePrice = 0
	// song.Url = ""
	song.Price = 0
	song.Copyright = 0
	song.CrbtId = 0
	song.Coverart = ""
	song.Coverart310 = ""
	song.DateCreated = time.Time{}
	song.Checktime = time.Time{}
	return song
}

// GetSong returns a album.
func GetSong(id dna.Int) (*Song, error) {
	apiSong, err := GetAPISongEntry(id)
	if err != nil {
		return nil, err
	} else {
		song := apiSong.MainSong.ToSong()
		if song.Id == 0 {
			return nil, errors.New("Invalid songid: Zero value found")
		} else {
			return song, nil
		}
	}
}

// Fetch implements item.Item interface
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

func (song *Song) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(song)
}
