package zi

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

// The basic song type
type Album struct {
	Id           dna.Int
	Key          dna.String
	EncodedKey   dna.String
	Title        dna.String
	Artists      dna.StringArray
	Coverart     dna.String
	Topics       dna.StringArray
	Plays        dna.Int
	Songids      dna.IntArray
	YearReleased dna.String
	Nsongs       dna.Int
	Description  dna.String
	DateCreated  time.Time
	Checktime    time.Time
	// add more 6 fields
	IsAlbum    dna.Int
	IsHit      dna.Int
	IsOfficial dna.Int
	Likes      dna.Int
	Comments   dna.Int
	StatusId   dna.Int
	ArtistIds  dna.IntArray
}

// NewAlbum returns a new pointer to Album
func NewAlbum() *Album {
	album := new(Album)
	album.Key = ""
	album.Id = 0
	album.EncodedKey = ""
	album.Title = ""
	album.Artists = dna.StringArray{}
	album.Coverart = ""
	album.Topics = dna.StringArray{}
	album.Plays = 0
	album.Songids = dna.IntArray{}
	album.YearReleased = ""
	album.Nsongs = 0
	album.Description = ""
	album.DateCreated = time.Time{}
	album.Checktime = time.Time{}
	// add more 6 fields
	album.IsAlbum = 0
	album.IsHit = 0
	album.IsOfficial = 0
	album.Likes = 0
	album.StatusId = 0
	album.Comments = 0
	album.ArtistIds = dna.IntArray{}
	return album
}

//GetAlbumFromAPI gets a album from API. It does not get content from main site.
func GetAlbumFromAPI(id dna.Int) (*Album, error) {
	var album *Album = NewAlbum()
	album.Id = id
	apialbum, err := GetAPIAlbum(id)
	if err != nil {
		return nil, err
	} else {
		if apialbum.Response.MsgCode == 1 {
			if GetKey(apialbum.Id+307843200) != GetKey(album.Id) {
				panic("Resulted key and computed key are not match.")
			}

			album.Title = apialbum.Title
			album.Artists = dna.StringArray(apialbum.Artists.Split(" , ").Map(func(val dna.String, idx dna.Int) dna.String {
				return val.Trim()
			}).([]dna.String)).SplitWithRegexp(",").Filter(func(v dna.String, i dna.Int) dna.Bool {
				if v != "" {
					return true
				} else {
					return false
				}
			})

			album.Topics = dna.StringArray(apialbum.Topics.Split(", ").Map(func(val dna.String, idx dna.Int) dna.String {
				return val.Trim()
			}).([]dna.String)).SplitWithRegexp(" / ").Unique().Filter(func(v dna.String, i dna.Int) dna.Bool {
				if v != "" {
					return true
				} else {
					return false
				}
			})
			album.Plays = apialbum.Plays
			// album.Songids
			// album.Nsongs
			// album.EncodedKey
			// album.Coverart
			// album.DateCreated
			album.YearReleased = apialbum.YearReleased
			album.Description = apialbum.Description.RemoveHtmlTags("")

			album.ArtistIds = apialbum.ArtistIds.Split(",").ToIntArray()
			album.IsAlbum = apialbum.IsAlbum
			album.IsHit = apialbum.IsHit
			album.IsOfficial = apialbum.IsOfficial
			album.Likes = apialbum.Likes
			album.StatusId = apialbum.StatusId
			album.Comments = apialbum.Comments
			album.Checktime = time.Now()
			return album, nil
		} else {
			return nil, errors.New("Message code invalid " + apialbum.Response.MsgCode.ToString().String())
		}
	}
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

			// playsArr := data.FindAllStringSubmatch(`Lượt nghe:</span>(.+)</p>`, -1)
			// if len(playsArr) > 0 {
			// 	album.Plays = playsArr[0][1].Trim().Replace(".", "").ToInt()
			// }

			// yearsArr := data.FindAllStringSubmatch(`Năm phát hành:</span>(.+)</p>`, -1)
			// if len(yearsArr) > 0 {
			// 	album.YearReleased = yearsArr[0][1].Trim()
			// }

			nsongsArr := data.FindAllStringSubmatch(`Số bài hát:</span>(.+)</p>`, -1)
			if len(nsongsArr) > 0 {
				album.Nsongs = nsongsArr[0][1].Trim().ToInt()
			}

			// topicsArr := data.FindAllStringSubmatch(`Thể loại:(.+)`, -1)
			// if len(topicsArr) > 0 {
			// 	album.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().Split(", ").SplitWithRegexp(` / `).Unique()
			// }

			// descArr := data.FindAllStringSubmatch(`(?mis)(<p id="_albumIntro" class="rows2".+#_albumIntro">)Xem toàn bộ</a>`, -1)
			// if len(descArr) > 0 {
			// 	album.Description = descArr[0][1].RemoveHtmlTags("").Trim()
			// }

			// titleArr := data.FindAllStringSubmatch(`<h1 class="detail-title">(.+) - <a.+`, -1)
			// if len(titleArr) > 0 {
			// 	album.Title = titleArr[0][1].RemoveHtmlTags("").Trim()
			// }

			// artistsArr := data.FindAllStringSubmatch(`<h1 class="detail-title">.+(<a.+)`, -1)
			// if len(artistsArr) > 0 {
			// 	album.Artists = dna.StringArray(artistsArr[0][1].RemoveHtmlTags("").Trim().Split(" ft. ").Unique().Map(func(val dna.String, idx dna.Int) dna.String {
			// 		return val.Trim()
			// 	}).([]String))
			// }

			covertArr := data.FindAllStringSubmatch(`<span class="album-detail-img">(.+)`, -1)
			if len(covertArr) > 0 {
				album.Coverart = covertArr[0][1].GetTagAttributes("src")
				datecreatedArr := album.Coverart.FindAllStringSubmatch(`_([0-9]+)\..+$`, -1)
				if len(datecreatedArr) > 0 {
					// Log(int64(datecreatedArr[0][1].ToInt()))
					album.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()), 0)
				}
			}

			songidsArr := data.FindAllString(`id="_divPlsLite.+?"`, -1)
			if songidsArr.Length() > 0 {
				album.Songids = dna.IntArray(songidsArr.Map(func(val dna.String, idx dna.Int) dna.Int {
					return GetId(val.FindAllStringSubmatch(`id="_divPlsLite(.+)"`, -1)[0][1])
				}).([]dna.Int))
			}

		}
		channel <- true

	}()
	return channel
}

// getAlbumFromAPI returns album from API
func getAlbumFromAPI(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		apialbum, err := GetAlbumFromAPI(album.Id)
		if err == nil {
			album.Title = apialbum.Title
			album.Artists = apialbum.Artists
			album.Topics = apialbum.Topics
			album.Plays = apialbum.Plays
			album.YearReleased = apialbum.YearReleased
			album.Description = apialbum.Description
			album.ArtistIds = apialbum.ArtistIds
			album.IsAlbum = apialbum.IsAlbum
			album.IsHit = apialbum.IsHit
			album.IsOfficial = apialbum.IsOfficial
			album.Likes = apialbum.Likes
			album.StatusId = apialbum.StatusId
			album.Comments = apialbum.Comments
			album.Checktime = time.Now()
		}
		channel <- true

	}()
	return channel
}

// GetAlbum returns a pointer to Album
//
// Notice: Once getting special albums' titles such as "Bảng Xếp Hạng Bài Hát Hàn Quốc ...",
// the albums will be discarded because album.Nsongs and album.Songids.Length() do not match.
func GetAlbum(id dna.Int) (*Album, error) {
	var album *Album = NewAlbum()
	album.Id = id
	album.Key = GetKey(id)
	c := make(chan bool, 2)

	go func() {
		c <- <-getAlbumFromMainPage(album)
	}()
	go func() {
		c <- <-getAlbumFromAPI(album)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}

	if album.Nsongs != album.Songids.Length() {
		return nil, errors.New(dna.Sprintf("Zing - Album %v: Songids and Nsongs do not match", album.Id).String())
	} else if album.Nsongs == 0 && album.Songids.Length() == 0 {
		return nil, errors.New(dna.Sprintf("Zing - Album %v: No song found", album.Id).String())
	} else {
		album.Checktime = time.Now()
		return album, nil
	}
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
	// case string:
	// 	album.Key = dna.String(v.(string))
	// case dna.String:
	// 	album.Key = v.(String)
	default:
		panic("Interface v has to be int")
	}
}

func (album *Album) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(album)
}
