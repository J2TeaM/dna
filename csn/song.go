package csn

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

var LastSongId dna.Int

type Song struct {
	Id            dna.Int
	Title         dna.String
	Artists       dna.StringArray
	Authors       dna.StringArray
	Topics        dna.StringArray
	AlbumTitle    dna.String
	AlbumHref     dna.String
	AlbumCoverart dna.String
	Producer      dna.String
	Downloads     dna.Int
	Plays         dna.Int
	Formats       dna.String
	Href          dna.String
	IsLyric       dna.Int
	Lyric         dna.String
	DateReleased  dna.String
	DateCreated   time.Time
	// Type decides it is a song or a video.
	// If it is a video, method ToVideo will be called
	Type      dna.Bool
	Checktime time.Time
}

// NewSong returns new song whose id is 0
func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Artists = dna.StringArray{}
	song.Authors = dna.StringArray{}
	song.Topics = dna.StringArray{}
	song.AlbumTitle = ""
	song.AlbumHref = ""
	song.AlbumCoverart = ""
	song.Producer = ""
	song.Downloads = 0
	song.Plays = 0
	song.Formats = ""
	song.Href = ""
	song.IsLyric = 0
	song.Lyric = ""
	song.DateReleased = ""
	song.Type = false
	song.DateCreated = time.Time{}
	song.Checktime = time.Time{}
	return song
}

func refineAuthorsOrArtists(str dna.String) dna.StringArray {

	tmp := str.ToStringArray().SplitWithRegexp(`; `).SplitWithRegexp(` / `).SplitWithRegexp(` - `).SplitWithRegexp(` – `)
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

func getSongFormats(song *Song) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://download.chiasenhac.com/google-bot~" + song.Id.ToString() + "_download.html"
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil {
			data := &result.Data
			urlsArr := data.FindAllString(`<a href.+(Link|Mobile) Download.+`, -1)
			if urlsArr.Length() > 0 {
				song.Formats, song.Type = GetFormats(urlsArr)
				// dna.Log(song.Formats)
			}

			topicsArr := data.FindAllStringSubmatch(`<div class="plt-text">(.+)`, 1)
			if len(topicsArr) > 0 {
				topics := topicsArr[0][1].RemoveHtmlTags("").Trim().Split(`-&gt;`)
				topics.Pop()
				song.Topics = dna.StringArray(topics.Map(func(val dna.String, idx dna.Int) dna.String {
					return val.Replace("...", "").Replace("Nước khác", "Nhạc Các Nước Khác").Replace("Âu, Mỹ", "Âu Mỹ").Trim().Title()
				}).([]dna.String)).SplitWithRegexp(`, `).Unique()
			}

			hrefArr := data.FindAllString(`Download: <a.+`, 1)
			if hrefArr.Length() > 0 {
				song.Href = "http://chiasenhac.com/" + hrefArr[0].GetTagAttributes("href").Trim()
				titleArtists := hrefArr[0].Replace("Download:", "").Trim().RemoveHtmlTags("").Split(" - ")
				song.Artists = refineAuthorsOrArtists(titleArtists[titleArtists.Length()-1])
				titleArtists.Pop()
				song.Title = titleArtists.Join(" - ")

				<-getSongFromMainPage(song)

			}

			dlArr := data.FindAllStringSubmatch(`([0-9]+) downloads`, 1)
			if len(dlArr) > 0 {
				song.Downloads = dlArr[0][1].Trim().Replace(".", "").ToInt()
			}

			dateCreatedArr := data.FindAllStringSubmatch(`<img src="images/tain5.gif" title="Upload at (.+?)" />`, 1)
			if len(dateCreatedArr) > 0 {
				tmp := dateCreatedArr[0][1].Trim()
				switch {
				case tmp.Match(`Hôm qua`) == true:
					song.DateCreated = time.Now().AddDate(0, 0, -1)
				case tmp.Match(`Hôm nay`) == true:
					song.DateCreated = time.Now()
				default:
					val := tmp.ReplaceWithRegexp(`(\d+)/(\d+)/(\d+) (\d+):(\d+)`, "${3}-${2}-${1}T${4}:${5}:00Z")
					parsedDate, err := time.Parse(time.RFC3339, val.String())
					if err == nil {
						song.DateCreated = parsedDate
					}
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
		// link := "http://playlist.chiasenhac.com/google-bot~" + song.Id.ToString() + ".html"
		result, err := http.Get(song.Href)
		// dna.Log(song.Href)
		if err == nil {
			data := &result.Data
			// titleArr := data.FindAllStringSubmatch(`<span class="maintitle">(.+)`, 1)
			// if len(titleArr) > 0 {
			// 	song.Title = titleArr[0][1].RemoveHtmlTags("").Trim()
			// }

			// artistArr := data.FindAllStringSubmatch(`Trình bày:(.+)`, 1)
			// if len(artistArr) > 0 {
			// 	song.Artists = refineAuthorsOrArtists(artistArr[0][1].RemoveHtmlTags("").Trim())
			// }

			// artistArr = data.FindAllStringSubmatch(`Ca sĩ:(.+)`, 1)
			// if len(artistArr) > 0 {
			// 	song.Artists = refineAuthorsOrArtists(artistArr[0][1].RemoveHtmlTags("").Trim())
			// }

			albumTitleArr := data.FindAllStringSubmatch(`Album:</u>(.+?)</a>`, 1)
			if len(albumTitleArr) > 0 {
				song.AlbumTitle = albumTitleArr[0][1].RemoveHtmlTags("").Trim().DecodeHTML()
			}

			albumHrefArr := data.FindAllStringSubmatch(`Album:</u>.+?</a>(.+)`, 1)
			if len(albumHrefArr) > 0 {
				song.AlbumHref = albumHrefArr[0][1].GetTagAttributes("href").Trim()
			}

			albumCoverArr := data.FindAllStringSubmatch(`(<link rel="image_src".+)`, 1)
			if len(albumCoverArr) > 0 {
				song.AlbumCoverart = albumCoverArr[0][1].GetTagAttributes("href").Trim()
			}

			producerArr := data.FindAllStringSubmatch(`Sản xuất:(.+)`, 1)
			if len(producerArr) > 0 {
				song.Producer = producerArr[0][1].RemoveHtmlTags("").Trim()
			}

			authorArr := data.FindAllStringSubmatch(`<u>Sáng tác:(.+)`, 1)
			if len(authorArr) > 0 {
				song.Authors = refineAuthorsOrArtists(authorArr[0][1].RemoveHtmlTags("").Trim())
			}

			playsArr := data.FindAllStringSubmatch(`(?mis)<img src="images/bh1.gif".+?<span>(.+?)</span>`, 1)
			if len(playsArr) > 0 {
				song.Plays = playsArr[0][1].Trim().Replace(".", "").ToInt()
			}

			// dlArr := data.FindAllStringSubmatch(`(?mis)images/bh3.gif.+?<span>(.+?)</span>`, 1)
			// if len(dlArr) > 0 {
			// 	song.Downloads = dlArr[0][1].Trim().Replace(".", "").ToInt()
			// }

			dateReleased := data.FindAllStringSubmatch(`Năm phát hành.+<b>(.+)<\/b>`, 1)
			if len(dateReleased) > 0 {
				song.DateReleased = dateReleased[0][1].Trim()
			}

			lyricArr := data.FindAllString(`(?mis)<p class="genmed".+?<div id="morelyric".+?</div>`, 1)
			if lyricArr.Length() > 0 {
				song.Lyric = lyricArr[0].RemoveHtmlTags("").Replace("<br /> ", "\n").Replace("<br />", "").Trim()
				song.IsLyric = 1
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
	c := make(chan bool)

	// go func() {
	// 	c <- <-getSongFromMainPage(song)
	// }()

	go func() {
		c <- <-getSongFormats(song)
	}()

	for i := 0; i < 1; i++ {
		<-c
	}
	song.Checktime = time.Now()
	if song.Formats == "" {
		return nil, errors.New(dna.Sprintf("Chiasenhac - Song %v: Mp3 link not found", song.Id).String())
	} else {
		return song, nil
	}
}

func (song *Song) ToVideo() *Video {
	video := NewVideo()
	video.Id = song.Id
	video.Title = song.Title
	video.Artists = song.Artists
	video.Authors = song.Authors
	video.Topics = song.Topics
	video.Thumbnail = song.AlbumCoverart
	video.Producer = song.Producer
	video.Downloads = song.Downloads
	video.Plays = song.Plays
	video.Formats = song.Formats
	video.Href = song.Href
	video.IsLyric = song.IsLyric
	video.Lyric = song.Lyric
	video.DateReleased = song.DateReleased
	video.DateCreated = song.DateCreated
	video.Type = song.Type
	video.Checktime = song.Checktime
	return video
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
	return item.Item(NewSongVideo())
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
