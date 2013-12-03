package zi

import (
	. "dna"
	"dna/http"
	"dna/site"
	"dna/sqlpg"
	"errors"
	"time"
)

// The basic song type
type Album struct {
	Id           Int
	Key          String
	EncodedKey   String
	Title        String
	Artists      StringArray
	Coverart     String
	Topics       StringArray
	Plays        Int
	Songids      IntArray
	YearReleased String
	Nsongs       Int
	Description  String
	DateCreated  time.Time
	Checktime    time.Time
}

// NewAlbum returns a new pointer to Album
func NewAlbum() *Album {
	album := new(Album)
	album.Key = ""
	album.Id = 0
	album.EncodedKey = ""
	album.Title = ""
	album.Artists = StringArray{}
	album.Coverart = ""
	album.Topics = StringArray{}
	album.Plays = 0
	album.Songids = IntArray{}
	album.YearReleased = ""
	album.Nsongs = 0
	album.Description = ""
	album.DateCreated = time.Time{}
	album.Checktime = time.Time{}
	return album
}

// getSongFromMainPage returns album from main page
func getAlbumFromMainPage(album *Album) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://mp3.zing.vn/album/google-bot/" + album.Key + ".html"
		result, err := http.Get(link)
		// Log(link)
		// Log(result.Data)
		if err == nil {
			data := &result.Data

			encodedKeyArr := data.FindAllStringSubmatch(`xmlURL=http://mp3.zing.vn/xml/album-xml/(.+)&amp;`, -1)
			if len(encodedKeyArr) > 0 {
				album.EncodedKey = encodedKeyArr[0][1]
			}

			playsArr := data.FindAllStringSubmatch(`Lượt nghe:</span>(.+)</p>`, -1)
			if len(playsArr) > 0 {
				album.Plays = playsArr[0][1].Trim().Replace(".", "").ToInt()
			}

			yearsArr := data.FindAllStringSubmatch(`Năm phát hành:</span>(.+)</p>`, -1)
			if len(yearsArr) > 0 {
				album.YearReleased = yearsArr[0][1].Trim()
			}

			nsongsArr := data.FindAllStringSubmatch(`Số bài hát:</span>(.+)</p>`, -1)
			if len(nsongsArr) > 0 {
				album.Nsongs = nsongsArr[0][1].Trim().ToInt()
			}

			topicsArr := data.FindAllStringSubmatch(`Thể loại:(.+)`, -1)
			if len(topicsArr) > 0 {
				album.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().Split(", ").SplitWithRegexp(` / `).Unique()
			}

			descArr := data.FindAllStringSubmatch(`(?mis)(<p id="_albumIntro" class="rows2".+#_albumIntro">)Xem toàn bộ</a>`, -1)
			if len(descArr) > 0 {
				album.Description = descArr[0][1].RemoveHtmlTags("").Trim()
			}

			titleArr := data.FindAllStringSubmatch(`<h1 class="detail-title">(.+) - <a.+`, -1)
			if len(titleArr) > 0 {
				album.Title = titleArr[0][1].RemoveHtmlTags("").Trim()
			}

			artistsArr := data.FindAllStringSubmatch(`<h1 class="detail-title">.+(<a.+)`, -1)
			if len(artistsArr) > 0 {
				album.Artists = StringArray(artistsArr[0][1].RemoveHtmlTags("").Trim().Split(" ft. ").Unique().Map(func(val String, idx Int) String {
					return val.Trim()
				}).([]String))
			}

			covertArr := data.FindAllStringSubmatch(`<span class="album-detail-img">(.+)`, -1)
			if len(covertArr) > 0 {
				album.Coverart = covertArr[0][1].GetTagAttributes("src")
				datecreatedArr := album.Coverart.FindAllStringSubmatch(`_([0-9]+)\..+$`, -1)
				if len(datecreatedArr) > 0 {
					Log(int64(datecreatedArr[0][1].ToInt()))
					album.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()), 0)
				}
			}

			songidsArr := data.FindAllString(`id="_divPlsLite.+?"`, -1)
			if songidsArr.Length() > 0 {
				album.Songids = IntArray(songidsArr.Map(func(val String, idx Int) Int {
					return GetId(val.FindAllStringSubmatch(`id="_divPlsLite(.+)"`, -1)[0][1])
				}).([]Int))
			}

		}
		channel <- true

	}()
	return channel
}

// GetAlbum returns a pointer to Album
func GetAlbum(id Int) (*Album, error) {
	var album *Album = NewAlbum()
	album.Id = id
	album.Key = GetKey(id)
	c := make(chan bool, 1)

	go func() {
		c <- <-getAlbumFromMainPage(album)
	}()

	for i := 0; i < 1; i++ {
		<-c
	}

	if album.Nsongs != album.Songids.Length() {
		return nil, errors.New(Sprintf("Zing - Album %v: Songids and Nsongs do not match", album.Id).String())
	} else if album.Nsongs == 0 {
		return nil, errors.New(Sprintf("Zing - Album %v: No song found", album.Id).String())
	} else {
		album.Checktime = time.Now()
		return album, nil
	}
}

// Fetch implements site.Item interface.
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

// New implements site.Item interface
// Returns new site.Item interface
func (album *Album) New() site.Item {
	return site.Item(NewAlbum())
}

// Init implements site.Item interface.
// It sets Id or key.
// Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (album *Album) Init(v interface{}) {
	switch v.(type) {
	case int:
		album.Id = Int(v.(int))
	case Int:
		album.Id = v.(Int)
	// case string:
	// 	album.Key = String(v.(string))
	// case String:
	// 	album.Key = v.(String)
	default:
		panic("Interface v has to be int")
	}
}

func (album *Album) Save(db *sqlpg.DB) error {
	return db.Update(album, "id", "artist_id", "video_id", "album_id", "is_hit", "is_official", "download_status", "copyright", "bitrate_flags", "likes", "comments")
}
