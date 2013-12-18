package cc

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

type Song struct {
	Id        dna.Int
	Title     dna.String
	Artists   dna.StringArray
	Artistid  dna.Int
	Topics    dna.StringArray
	Plays     dna.Int
	Duration  dna.Int
	Bitrate   dna.Int
	Coverart  dna.String
	Lyrics    dna.String
	Link      dna.String
	Checktime time.Time
}

// NewSong returns new song whose id is 0
func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Artists = dna.StringArray{}
	song.Artistid = 0
	song.Topics = dna.StringArray{}
	song.Plays = 0
	song.Duration = 0
	song.Bitrate = 0
	song.Coverart = ""
	song.Lyrics = ""
	song.Link = ""
	song.Checktime = time.Time{}
	return song
}

// getSongXML returns song from main page
func getSongXML(song *Song) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://www.chacha.vn/player/songXml/" + song.Id.ToString()
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil && !result.Data.Match(`Không tìm thấy bài hát`) {
			data := &result.Data
			if data.Match(`<enclosure.+label="320K"`) {
				song.Bitrate = 320
				link := data.FindAllString(`<enclosure.+label="320K".+/>`, -1)
				if link.Length() > 0 {
					song.Link = link[0].GetTagAttributes("url")
				}
			} else {
				if data.Match(`<enclosure.+label="128K"`) {
					song.Bitrate = 128
					link := data.FindAllString(`<enclosure.+label="128K".+?`, -1)
					if link.Length() > 0 {
						song.Link = link[0].GetTagAttributes("url")
					}
				}
			}

		}
		channel <- true

	}()
	return channel
}

func refineAuthorsOrArtists(str dna.String) dna.StringArray {

	tmp := str.ToStringArray().SplitWithRegexp(`-&nbsp;`).SplitWithRegexp(` / `).SplitWithRegexp(` - `).SplitWithRegexp(` – `)
	tmp = tmp.SplitWithRegexp(`, `).SplitWithRegexp(` ft `).SplitWithRegexp(` feat `).SplitWithRegexp(` ft. `)
	tmp = tmp.SplitWithRegexp(` feat. `).SplitWithRegexp(` Feat. `).SplitWithRegexp(` Ft. `)
	tmp = tmp.SplitWithRegexp(` & `).SplitWithRegexp(` vs. `).SplitWithRegexp(`- `).SplitWithRegexp(` _ `)
	tmp = dna.StringArray(tmp.Map(func(val dna.String, idx dna.Int) dna.String {
		rv := val.Replace(`Đang Cập Nhật...`, ``).Replace(`Đang Cập Nhật (QT)`, ``)
		rv = rv.Replace(`Đang Cập Nhật (VN)`, ``).Replace(`Nhạc Phim QT`, `Nhạc Phim Quốc Tế`)
		rv = rv.Replace(`Đang cập nhật`, ``).Replace(`Nhiều Ca Sỹ`, `Various Artists`)
		rv = rv.Replace("Nhiều ca sĩ", "Various Artists").Replace("V.A", "Various Artists")
		return rv.Trim()
	}).([]dna.String)).Filter(func(val dna.String, idx dna.Int) dna.Bool {
		if val != "" {
			return true
		} else {
			return false
		}
	})
	return tmp
}

// getSongFromMainPage returns song from main page
func getSongFromMainPage(song *Song) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://www.chacha.vn/song/google-bot," + song.Id.ToString() + ".html"
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil && !result.Data.Match(`Không tìm thấy bài hát`) {
			data := &result.Data
			titleArr := data.FindAllStringSubmatch(`<h1 class="name">(.+)</h1>`, 1)
			if len(titleArr) > 0 {
				song.Title = titleArr[0][1].Trim()
			}

			artistArr := data.FindAllStringSubmatch(`(?mis)<li>Nghệ sĩ:(.+?)</li>`, 1)
			if len(artistArr) > 0 {
				song.Artists = refineAuthorsOrArtists(artistArr[0][1].RemoveHtmlTags("").Trim())
				artistid := artistArr[0][1].GetTagAttributes("href").FindAllString(`[0-9]+`, 1)
				if artistid.Length() > 0 {
					song.Artistid = artistid[0].ToInt()
				}
			}

			topicArr := data.FindAllStringSubmatch(`(?mis)<li>Thể loại:(.+?)</li>`, 1)
			if len(topicArr) > 0 {
				song.Topics = topicArr[0][1].RemoveHtmlTags("").Trim().ToStringArray()
			}

			playsArr := data.FindAllStringSubmatch(`([0-9]+) lượt nghe`, 1)
			if len(playsArr) > 0 {
				song.Plays = playsArr[0][1].ToInt()
			}

			lyricArr := data.FindAllStringSubmatch(`(?mis)<p class="lyric" id="lyric_box">(.+?)<a class="fs11 more" id="lyric_more".+?</a>`, 1)
			if len(lyricArr) > 0 {
				song.Lyrics = lyricArr[0][1].Replace(`<br /> `, "\n").Replace(`<br />`, "\n").RemoveHtmlTags("").Trim()
			}

			coverartArr := data.FindAllString(`<meta property="og:image".+`, 1)
			if coverartArr.Length() > 0 {
				song.Coverart = coverartArr[0].GetTagAttributes("content")
			}

		}
		channel <- true

	}()
	return channel
}

// GetSong returns a song or an error
//
// Direct link:
// 	curl 'http://hcm.nhac.vui.vn/ajax/nghe_bai_hat/download_320k/472092' -H 'Cookie: pageCookie=13; ACCOUNT_ID=965257; token=3f363de2c081a3a3a685b1033e6f03b1%7C52ab4c37;' -v
func GetSong(id dna.Int) (*Song, error) {
	var song *Song = NewSong()
	song.Id = id
	c := make(chan bool, 2)

	go func() {
		c <- <-getSongXML(song)
	}()
	go func() {
		c <- <-getSongFromMainPage(song)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}
	if song.Link == "" {
		return nil, errors.New(dna.Sprintf("Chacha - Song %v: Mp3 link not found", song.Id).String())
	} else {
		song.Checktime = time.Now()
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

func (song *Song) Save(db *sqlpg.DB) error {

	return db.InsertIgnore(song)
}
