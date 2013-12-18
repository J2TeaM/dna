package cc

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

const (
	LRe240p  = 1 << iota // Flag of  240p resolution
	LRe360p              // Flag of 360p resolution
	LRe480p              // Flag of 480p resolution
	LRe720p              // Flag of 720p resolution
	LRe1080p             // Flag of 1080p resolutions
)

type Video struct {
	Id              dna.Int
	Title           dna.String
	Artists         dna.StringArray
	Topics          dna.StringArray
	Plays           dna.Int
	ResolutionFlags dna.Int
	Thumbnail       dna.String
	Lyric           dna.String
	Links           dna.StringArray
	YearReleased    dna.Int
	Checktime       time.Time
}

// NewVideo returns new video whose id is 0
func NewVideo() *Video {
	video := new(Video)
	video.Id = 0
	video.Title = ""
	video.Artists = dna.StringArray{}
	video.Topics = dna.StringArray{}
	video.Plays = 0
	video.ResolutionFlags = 0
	video.Thumbnail = ""
	video.Lyric = ""
	video.Links = dna.StringArray{}
	video.YearReleased = 0
	video.Checktime = time.Time{}
	return video
}

// getVideoXML returns video from main page
func getVideoXML(video *Video) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://www.chacha.vn/player/videoXml/" + video.Id.ToString()
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil && !result.Data.Match(`Bài hát không tồn tại`) {
			data := &result.Data
			tagArr := data.FindAllString(`<enclosure.+label=".+?"`, -1)
			if tagArr.Length() > 0 {
				for _, tag := range tagArr {
					tmp := tag.FindAllStringSubmatch(`label="(.+?)"`, -1)
					if len(tmp) > 0 {
						switch tmp[0][1] {
						case "240p":
							video.ResolutionFlags |= LRe240p
						case "360p":
							video.ResolutionFlags |= LRe360p
						case "480p":
							video.ResolutionFlags |= LRe480p
						case "720p":
							video.ResolutionFlags |= LRe720p
						case "1080p":
							video.ResolutionFlags |= LRe1080p
						}
					}
					video.Links.Push(tag.GetTagAttributes("url"))
				}
			}
		}
		channel <- true

	}()
	return channel
}

// getVideoFromMainPage returns video from main page
func getVideoFromMainPage(video *Video) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://www.chacha.vn/video/google-bot," + video.Id.ToString() + ".html"
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil && !result.Data.Match(`Không tìm thấy trang`) {
			data := &result.Data
			titleArr := data.FindAllStringSubmatch(`<h1 class="name">(.+)</h1>`, 1)
			if len(titleArr) > 0 {
				video.Title = titleArr[0][1].Trim()
			}

			artistArr := data.FindAllStringSubmatch(`(?mis)<li>Nghệ sĩ:(.+?)</li>`, 1)
			if len(artistArr) > 0 {
				video.Artists = refineAuthorsOrArtists(artistArr[0][1].RemoveHtmlTags("").Trim())
			}

			topicArr := data.FindAllStringSubmatch(`(?mis)<li>Thể loại:(.+?)</li>`, 1)
			if len(topicArr) > 0 {
				video.Topics = topicArr[0][1].RemoveHtmlTags("").Trim().ToStringArray()
			}

			yearReleasedArr := data.FindAllStringSubmatch(`Năm phát hành: <span>(.+?)</span>`, 1)
			if len(yearReleasedArr) > 0 {
				video.YearReleased = yearReleasedArr[0][1].Trim().ToInt()
			}

			playsArr := data.FindAllStringSubmatch(`([0-9]+) lượt xem`, 1)
			if len(playsArr) > 0 {
				video.Plays = playsArr[0][1].ToInt()
			}

			coverartArr := data.FindAllString(`<meta property="og:image".+`, 1)
			if coverartArr.Length() > 0 {
				video.Thumbnail = coverartArr[0].GetTagAttributes("content")
			}

			lyricArr := data.FindAllStringSubmatch(`(?mis)<p class="lyric" id="lyric_box">(.+?)<a class="fs11 more" id="lyric_more".+?</a>`, 1)
			if len(lyricArr) > 0 {
				video.Lyric = lyricArr[0][1].Replace(`<br /> `, "\n").Replace(`<br />`, "\n").RemoveHtmlTags("").Trim()
			}

		}
		channel <- true

	}()
	return channel
}

// GetVideo returns a video or an error
//
// Direct link:
// 	curl 'http://hcm.nhac.vui.vn/ajax/nghe_bai_hat/download_320k/472092' -H 'Cookie: pageCookie=13; ACCOUNT_ID=965257; token=3f363de2c081a3a3a685b1033e6f03b1%7C52ab4c37;' -v
func GetVideo(id dna.Int) (*Video, error) {
	var video *Video = NewVideo()
	video.Id = id
	c := make(chan bool, 2)

	go func() {
		c <- <-getVideoXML(video)
	}()
	go func() {
		c <- <-getVideoFromMainPage(video)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}

	// Check how many bits in ResolutionFlags have value 1
	var count uint32 = 0
	for i := uint32(0); i < 5; i++ {
		if (int(video.ResolutionFlags)>>i)&1 == 1 {
			count += 1
		}
	}

	if video.Links.Length() == 0 {
		return nil, errors.New(dna.Sprintf("Chacha - Video %v: Link not found", video.Id).String())
	} else if dna.Int(count) != video.Links.Length() {
		return nil, errors.New(dna.Sprintf("Chacha - Video %v: Video Resolution flags and links do not match", video.Id).String())
	} else {
		video.Checktime = time.Now()
		return video, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (video *Video) Fetch() error {
	_video, err := GetVideo(video.Id)
	if err != nil {
		return err
	} else {
		*video = *_video
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (video *Video) GetId() dna.Int {
	return video.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (video *Video) New() item.Item {
	return item.Item(NewVideo())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (video *Video) Init(v interface{}) {
	switch v.(type) {
	case int:
		video.Id = dna.Int(v.(int))
	case dna.Int:
		video.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (video *Video) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(video)
}
