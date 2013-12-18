package cc

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

type Album struct {
	Id           dna.Int
	Title        dna.String
	Artists      dna.StringArray
	Topics       dna.StringArray
	Nsongs       dna.Int
	Plays        dna.Int
	Coverart     dna.String
	Description  dna.String
	Songids      dna.IntArray
	YearReleased dna.Int
	Checktime    time.Time
}

func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Title = ""
	album.Artists = dna.StringArray{}
	album.Topics = dna.StringArray{}
	album.Plays = 0
	album.Songids = dna.IntArray{}
	album.Nsongs = 0
	album.Description = ""
	album.Coverart = ""
	album.YearReleased = 0
	album.Checktime = time.Time{}
	return album
}

// getAlbumFromMainPage returns album from main page
func getAlbumFromMainPage(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.chacha.vn/album/google," + album.Id.ToString() + ".html"
		// dna.Log(link)
		result, err := http.Get(link)
		if err == nil && !result.Data.Match(`Không tìm thấy trang`) {
			data := &result.Data

			titleArr := data.FindAllStringSubmatch(`<h1 class="name">(.+)</h1>`, 1)
			if len(titleArr) > 0 {
				album.Title = titleArr[0][1].Trim()
			}

			artistArr := data.FindAllStringSubmatch(`(?mis)<li>Nghệ sĩ:(.+?)</li>`, 1)
			if len(artistArr) > 0 {
				album.Artists = refineAuthorsOrArtists(artistArr[0][1].RemoveHtmlTags("").Trim())
			}

			topicArr := data.FindAllStringSubmatch(`(?mis)<li>Thể loại:(.+?)</li>`, 1)
			if len(topicArr) > 0 {
				album.Topics = topicArr[0][1].RemoveHtmlTags("").Trim().ToStringArray()
			}

			yearReleasedArr := data.FindAllStringSubmatch(`Năm phát hành: <span>(.+?)</span>`, 1)
			if len(yearReleasedArr) > 0 {
				album.YearReleased = yearReleasedArr[0][1].Trim().ToInt()
			}

			nsongsArr := data.FindAllStringSubmatch(`Số bài hát: <span>(.+)</span><`, 1)
			if len(nsongsArr) > 0 {
				album.Nsongs = nsongsArr[0][1].Trim().ToInt()
			}

			playsArr := data.FindAllStringSubmatch(`([0-9]+) Lượt nghe`, 1)
			if len(playsArr) > 0 {
				album.Plays = playsArr[0][1].ToInt()
			}

			coverartArr := data.FindAllString(`<img class="detail".+`, 1)
			if coverartArr.Length() > 0 {
				album.Coverart = coverartArr[0].GetTagAttributes("src")
			}

			descArr := data.FindAllStringSubmatch(`(?mis)<p class="clb desc".+?">(.+?)<a class="fs11 more" id="desc_more".+Xem thêm`, 1)
			if len(descArr) > 0 {
				album.Description = descArr[0][1].Replace(`<br /> `, "\n").Replace(`<br />`, "\n").RemoveHtmlTags("").DecodeHTML().Trim()
			}

			songidsArr := data.FindAllString(`<li id="song_[0-9]+" value="[0-9]+">`, -1)
			songidsArr.ForEach(func(val dna.String, idx dna.Int) {
				tmp := val.FindAllStringSubmatch(`<li id="song_[0-9]+" value="([0-9]+)">`, 1)
				if len(tmp) > 0 {
					album.Songids.Push(tmp[0][1].ToInt())
				}
			})

		}
		channel <- true
	}()
	return channel
}

// GetAlbum returns a album or an error
// 	* key: A unique key of a album
// 	* Official : 0 or 1, if its value is unknown, set to 0
// 	* Returns a found album or an error
func GetAlbum(id dna.Int) (*Album, error) {
	var album *Album = NewAlbum()
	album.Id = id
	c := make(chan bool, 1)
	go func() {
		c <- <-getAlbumFromMainPage(album)
	}()
	for i := 0; i < 1; i++ {
		<-c
	}
	if album.Nsongs != album.Songids.Length() {
		return nil, errors.New(dna.Sprintf("Nhacvui - Album %v: Songids and Nsongs do not match", album.Id).String())
	} else if album.Nsongs == 0 && album.Songids.Length() == 0 {
		return nil, errors.New(dna.Sprintf("Nhacvui - Album %v: No song found", album.Id).String())
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
	default:
		panic("Interface v has to be int")
	}
}

func (album *Album) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(album)
}
