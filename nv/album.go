package nv

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

type Album struct {
	Id          dna.Int
	Title       dna.String
	Artists     dna.StringArray
	Topics      dna.StringArray
	Nsongs      dna.Int
	Plays       dna.Int
	Coverart    dna.String
	Description dna.String
	DateCreated time.Time
	Songids     dna.IntArray
	Checktime   time.Time
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
	album.Coverart = ""
	album.Description = ""
	album.DateCreated = time.Time{}
	album.Checktime = time.Time{}
	return album
}

// getAlbumFromMainPage returns album from main page
func getAlbumFromMainPage(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://hcm.nhac.vui.vn/google-bot-a" + album.Id.ToString() + "p1.html"
		result, err := http.Get(link)
		if err == nil && !result.Data.Match(`Album không tồn tại.`) {
			data := &result.Data

			tmpArr := data.FindAllStringSubmatch(`<div class="nghenhac-baihat">(.+)</div>`, 1)
			if len(tmpArr) > 0 {
				tmp := tmpArr[0][1].RemoveHtmlTags("").Split(" - ")
				album.Title = tmp[0]
				tmp.Shift()
				album.Artists = refineAuthorsOrArtists(tmp.Join(" - "))
			}
			topicsArr := data.FindAllStringSubmatch(`Thể loại:(.+)`, 1)
			if len(topicsArr) > 0 {
				album.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().ToStringArray()
			}
			playsArr := data.FindAllStringSubmatch(`Lượt nghe:(.+)`, 1)
			if len(playsArr) > 0 {
				album.Plays = playsArr[0][1].RemoveHtmlTags("").Trim().Replace(",", "").ToInt()
			}
			nsongArr := data.FindAllStringSubmatch(`Số bài hát:(.+)`, 1)
			if len(nsongArr) > 0 {
				album.Nsongs = nsongArr[0][1].RemoveHtmlTags("").Trim().Replace(",", "").ToInt()
			}

			descArr := data.FindAllStringSubmatch(`(?mis)<p class="albumInfo-rutgon" id="album_info">(.+)<a href="javascript:;".+Xem toàn bộ</a>`, 1)
			if len(descArr) > 0 {
				album.Description = descArr[0][1].RemoveHtmlTags("").Trim().Replace("<br />", "\n")
			}

			coverartArr := data.FindAllString(`(?mis)<span class="albumInfo-img">.+?</span>`, 1)
			if coverartArr.Length() > 0 {
				album.Coverart = coverartArr[0].GetTagAttributes("src")
				datecreatedArr := album.Coverart.FindAllStringSubmatch(`/([0-9]+)_.+$`, -1)
				if len(datecreatedArr) > 0 {
					// Log(int64(datecreatedArr[0][1].ToInt()))
					album.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()), 0)
				} else {
					dateCreatedArr := album.Coverart.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
					if len(dateCreatedArr) > 0 {
						year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
						month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
						day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
						album.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

					}
				}
			}

			songidsArr := data.FindAllString(`javascript:liked_onclick\('[0-9]+'\)`, -1)
			songidsArr.ForEach(func(val dna.String, idx dna.Int) {
				idArr := val.FindAllStringSubmatch(`javascript:liked_onclick\('([0-9]+)'\)`, 1)
				if len(idArr) > 0 {
					album.Songids.Push(idArr[0][1].ToInt())
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
	// return db.Update(album, "id", "description")
	return db.InsertIgnore(album)
}
