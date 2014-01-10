package sf

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"time"
)

type Album struct {
	Id            dna.Int
	AMG           dna.Int
	Title         dna.String
	Artistid      dna.Int
	Artists       dna.StringArray
	UrlSlug       dna.String
	Year          dna.Int
	Coverart      dna.String
	CoverartLarge dna.String
	Ratings       dna.IntArray
	Link          dna.String
	Songids       dna.IntArray
	ReviewAuthor  dna.String
	Review        dna.String
	Checktime     time.Time
}

// NewAlbum return default new album
func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.AMG = 0
	album.UrlSlug = ""
	album.Year = 0
	album.Coverart = ""
	album.CoverartLarge = ""
	album.Title = ""
	album.Ratings = dna.IntArray{0, 0, 0}
	album.Artistid = 0
	album.Artists = dna.StringArray{}
	album.Link = ""
	album.Songids = dna.IntArray{}
	album.ReviewAuthor = ""
	album.Review = ""
	album.Checktime = time.Time{}
	return album
}

//CSVRecord returns a record to write csv format.
//
//psql -c "COPY sfalbums (id,amg,title,artistid,artists,url_slug,year,coverart,coverart_large,ratings,link,songids,review_author,review,checktime) FROM '/Users/daonguyenanbinh/Box Documents/Sites/golang/sfalbums.csv' DELIMITER ',' CSV"
func (album *Album) CSVRecord() []string {
	return []string{
		album.Id.ToString().String(),
		album.AMG.ToString().String(),
		album.Title.String(),
		album.Artistid.ToString().String(),
		dna.Sprintf("%#v", album.Artists).Replace("dna.StringArray", "").String(),
		album.UrlSlug.String(),
		album.Year.ToString().String(),
		album.Coverart.String(),
		album.CoverartLarge.String(),
		dna.Sprintf("%#v", album.Ratings).Replace("dna.IntArray", "").String(),
		album.Link.String(),
		dna.Sprintf("%#v", album.Songids).Replace("dna.IntArray", "").String(),
		album.ReviewAuthor.String(),
		album.Review.String(),
		album.Checktime.Format("2006-02-01 15:04:05"),
	}
}

func getAlbumFromMainAPI(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		var link dna.String = "http://apiv2.songfreaks.com//lyric.do?"
		// Log(link)
		PostData.SetIdKey(album.Id)
		dna.Log(PostData.Encode())
		result, err := Post(link, PostData.Encode())
		mutex.Lock()
		Cookie = result.Header.Get("Set-Cookie")
		mutex.Unlock()
		if err == nil {
			dna.Log(result.Data)
		}

		channel <- true
	}()
	return channel
}

// GetAlbum returns an album and an error (if available)
func GetAlbum(id dna.Int) (*Album, error) {
	var album *Album = NewAlbum()
	album.Id = id
	c := make(chan bool)

	go func() {
		c <- <-getAlbumFromMainAPI(album)
	}()
	for i := 0; i < 1; i++ {
		<-c
	}

	return album, nil
}

// Fetch implements item.Item interface.
// Returns error if can not get item
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
