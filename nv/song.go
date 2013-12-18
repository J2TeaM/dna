package nv

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
	Authors   dna.StringArray
	Topics    dna.StringArray
	Plays     dna.Int
	Lyric     dna.String
	Link      dna.String
	Link320   dna.String
	Type      dna.String
	Checktime time.Time
}

// NewSong returns new song whose id is 0
func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Artists = dna.StringArray{}
	song.Authors = dna.StringArray{}
	song.Plays = 0
	song.Lyric = ""
	song.Link = ""
	song.Link320 = ""
	song.Topics = dna.StringArray{}
	song.Type = ""
	song.Checktime = time.Time{}
	return song
}

func refineAuthorsOrArtists(str dna.String) dna.StringArray {
	tmp := str.ToStringArray().SplitWithRegexp(` / `).SplitWithRegexp(` - `).SplitWithRegexp(` – `)
	tmp = tmp.SplitWithRegexp(`, `).SplitWithRegexp(` ft `).SplitWithRegexp(` feat `).SplitWithRegexp(` ft. `)
	tmp = tmp.SplitWithRegexp(` feat. `).SplitWithRegexp(` Feat. `).SplitWithRegexp(` Ft. `)
	tmp = tmp.SplitWithRegexp(` & `).SplitWithRegexp(` vs. `).SplitWithRegexp(`- `).SplitWithRegexp(` & `)
	tmp = dna.StringArray(tmp.Map(func(val dna.String, idx dna.Int) dna.String {
		rv := val.Replace(`Đang Cập Nhật...`, ``).Replace(`Đang Cập Nhật (QT)`, ``)
		rv = rv.Replace(`Đang Cập Nhật (VN)`, ``).Replace(`Nhạc Phim QT`, `Nhạc Phim Quốc Tế`)
		rv = rv.Replace(`Đang cập nhật`, ``).Replace(`Nhiều Ca Sỹ`, `Various Artists`)
		return rv
	}).([]dna.String)).Filter(func(val dna.String, idx dna.Int) dna.Bool {
		if val != "" {
			return true
		} else {
			return false
		}
	})
	return tmp
}

// getSongXML returns song from main page
func getSongXML(song *Song) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://hcm.nhac.vui.vn/asx2.php?type=1&id=" + song.Id.ToString()
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil && !result.Data.Match(`Bài hát không tồn tại`) {
			data := &result.Data
			linkArr := data.FindAllStringSubmatch(`<jwplayer:file><\!\[CDATA\[(.+)\]\]></jwplayer:file>`, 1)
			if len(linkArr) > 0 {
				song.Link = linkArr[0][1].Trim()
				switch {
				case song.Link.Match("mp3") == true:
					song.Type = "song"
				case song.Link.Match("(?mis)mp4") == true:
					song.Type = "video"
				case song.Link.Match("flv") == true:
					song.Type = "video"
				}
			}

		}
		channel <- true

	}()
	return channel
}

// getSongFromMainPage returns song from main page
func getSongFromMainPage(song *Song) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://hcm.nhac.vui.vn/google-bot-m" + song.Id.ToString() + "c2p1a1.html"
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil && !result.Data.Match(`Bài hát không tồn tại`) {
			data := &result.Data
			authorArr := data.FindAllStringSubmatch(`Nhạc sĩ: <span>(.+)</span>`, 1)
			if len(authorArr) > 0 {
				song.Authors = refineAuthorsOrArtists(authorArr[0][1])
			}
			tmpArr := data.FindAllStringSubmatch(`<div class="nghenhac-baihat"><h1>(.+)</h1></div>`, 1)
			if len(tmpArr) > 0 {
				tmp := tmpArr[0][1].Trim().Split(` - `)
				song.Title = tmp[0]
				if tmp.Length() > 1 {
					tmp.Shift()
					song.Artists = refineAuthorsOrArtists(tmp.Join(" - "))
				}
			}

			topicsArr := data.FindAllStringSubmatch(`Thể loại: (<a.+?)\|`, 1)
			if len(topicsArr) > 0 {
				song.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().ToStringArray().SplitWithRegexp(` / `).SplitWithRegexp(`/`)
			}

			playsArr := data.FindAllStringSubmatch(`Lượt nghe: (.+?)</p></div>`, 1)
			if len(playsArr) > 0 {
				song.Plays = playsArr[0][1].Trim().Replace(",", "").ToInt()
			}
			lyricArr := data.FindAllString(`(?mis)<div class="nghenhac-loibaihat-cnt.+<div class="loi-bai-hat-footer">`, 1)
			if lyricArr.Length() > 0 {
				lyric := lyricArr[0].Trim().Replace(`<br/>`, "\n").RemoveHtmlTags("").Trim()
				if !lyric.Match(`Hiện bài hát.+chưa có lời.`) {
					song.Lyric = lyric
				}
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
	if song.Type == "video" {
		song.Link = ""
		FoundVideos.Push(song.Id)
	}
	if song.Link == "" {
		return nil, errors.New(dna.Sprintf("Nhacvui - Song %v: Mp3 link not found", song.Id).String())
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
