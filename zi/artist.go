package zi

import (
	. "dna"
	"dna/http"
	"dna/site"
	"dna/sqlpg"
	"errors"
	"time"
)

type Artist struct {
	Id       Int
	Name     String
	Birthday time.Time
}

func NewArtist() *Artist {
	artist := new(Artist)
	artist.Id = 0
	artist.Name = ""
	artist.Birthday = time.Time{}
	return artist
}

// getArtistFromMainPage returns song from main page
func getArtistFromMainPage(artist *Artist) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "xxxxxxxxxxxxxx" + GetKey(artist.Id) + ".html"
		result, err := http.Get(link)
		// Log(link)
		// Log(result.Data)
		data := &result.Data
		Log(data.Match("<title>Thông báo</title>"))
		if err == nil && !data.Match("<title>Thông báo</title>") {

			//////// CHANGE!!!!!
			topicsArr := data.FindAllStringSubmatch(`Thể loại:(.+)\|`, -1)
			Log(topicsArr)

		}
		channel <- true

	}()
	return channel
}

// GetArtist returns a artist or an error
func GetArtist(id Int) (*Artist, error) {
	var artist *Artist = NewArtist()
	artist.Id = id
	c := make(chan bool, 2)

	// go func() {
	// 	c <- <-getSongFromXML(artist)
	// }()
	go func() {
		c <- <-getArtistFromMainPage(artist)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}

	if artist.Name == "" {
		return nil, errors.New(Sprintf("Zing - Song %v: Mp3 link not found", artist.Id).String())
	} else {
		return artist, nil
	}
}

// Fetch implements site.Item interface.
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

// New implements site.Item interface
// Returns new site.Item interface
func (artist *Artist) New() site.Item {
	return site.Item(NewSong())
}

// Init implements site.Item interface.
// It sets Id or key.
// Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (artist *Artist) Init(v interface{}) {
	switch v.(type) {
	case int:
		artist.Id = Int(v.(int))
	case Int:
		artist.Id = v.(Int)
	default:
		panic("Interface v has to be int")
	}
}

func (artist *Artist) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(artist)
}
