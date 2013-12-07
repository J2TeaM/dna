package ns

import (
	. "dna"
	"dna/http"
	"dna/site"
	"dna/sqlpg"
	"errors"
	"fmt"
	"math"
	"time"
)

// Define new Album type.
// Notice: Artistid should be Artistids , but this field is not important, then it will be ignored.
type Video struct {
	Id          Int
	Title       String
	Artists     StringArray
	Topics      StringArray
	Plays       Int
	Duration    Int
	Official    Int
	Producerid  Int
	Link        String
	Sublink     String
	Thumbnail   String
	DateCreated time.Time
	Checktime   time.Time
}

// NewVideo return default new video
func NewVideo() *Video {
	video := new(Video)
	video.Id = 0
	video.Title = ""
	video.Artists = StringArray{}
	video.Topics = StringArray{}
	video.Plays = 0
	video.Duration = 0
	video.Official = 0
	video.Producerid = 0
	video.Link = ""
	video.Sublink = ""
	video.Thumbnail = ""
	video.DateCreated = time.Time{}
	video.Checktime = time.Time{}

	return video
}

func getVideoDurationAndSublink(video *Video) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/flash/video/xnl/1/id/" + GetKey(video.Id)
		// Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			// getValueXML is in song.go
			// Log(getValueXML(data, "duration", 0))
			video.Duration = getValueXML(data, "duration", 0).RemoveHtmlTags("").ToInt()
			video.Sublink = getValueXML(data, "subUrl", 0)
		}
		channel <- true
	}()
	return channel
}

func getVideoFromMainPage(video *Video) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/xem-video/google-bot." + GetKey(video.Id) + "=.html"
		// Log(link)
		result, err := http.Get(link)
		if err == nil && !result.Data.Match("Rất tiếc, chúng tôi không tìm thấy thông tin bạn yêu cầu!") {
			data := &result.Data
			temp := data.FindAllString(`(?mis)<p class="title_video.+Đăng bởi`, 1)
			if temp.Length() > 0 {
				title := temp[0].FindAllString(`<h1 class="title">.+</h1>`, 1)
				if title.Length() > 0 {
					video.Title = title[0].RemoveHtmlTags("").Trim()
				}

				if temp[0].Match(`official`) {
					video.Official = 1
				}

				artists := temp[0].FindAllString(`<h2>.+</h2>`, -1)
				if artists.Length() > 0 {
					video.Artists = StringArray(artists.Map(func(val String, idx Int) String {
						return val.RemoveHtmlTags("").Trim()
					}).([]String)).Unique()
				}

			}
			topics := data.FindAllString(`<li><a href="http://nhacso.net/the-loai-video/.+</a></li>`, 1)
			if topics.Length() > 0 {
				video.Topics = topics[0].RemoveHtmlTags("").ToStringArray().SplitWithRegexp(` - `).SplitWithRegexp(`/`)
			}

			plays := data.FindAllString(`<span>.+</span><ins>&nbsp;lượt xem</ins>`, 1)
			if plays.Length() > 0 {
				video.Plays = plays[0].GetTags("span")[0].RemoveHtmlTags("").Trim().Replace(".", "").ToInt()
			}

			thumbLink := data.FindAllString(`poster="(.+)" src="(.+)" data`, 1)
			if thumbLink.Length() > 0 {
				video.Thumbnail = thumbLink[0].FindAllStringSubmatch(`poster="(.+?)" `, 1)[0][1]
				video.Link = thumbLink[0].FindAllStringSubmatch(`src="(.+?)" `, 1)[0][1]
				if video.Link != "" {
					ts := video.Link.FindAllStringSubmatch(`([0-9]+)_`, 1)
					if len(ts) > 0 {
						secs := float64(ts[0][1].ToInt()) * math.Pow10(13-len(ts[0][1]))
						// Log(secs)
						video.DateCreated = Float(secs / 1000).ToInt().ToTime()
					}

				}
			}
			producerid := data.FindAllStringSubmatch(`getProducerByListIds\('(\d+)', '#producer_'\);`, 1)
			if len(producerid) > 0 {
				video.Producerid = producerid[0][1].ToInt()
			}
		}
		channel <- true
	}()
	return channel

}

// GetVideo returns a video and an error (if available)
func GetVideo(id Int) (*Video, error) {
	var video *Video = NewVideo()
	video.Id = id
	c := make(chan bool)

	go func() {
		c <- <-getVideoFromMainPage(video)
	}()

	go func() {
		c <- <-getVideoDurationAndSublink(video)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}
	if video.Link == "" {
		return nil, errors.New(fmt.Sprintf("Nhacso - Video %v : Link not found", video.Id))
	} else {
		video.Checktime = time.Now()
		return video, nil
	}

}

// Fetch implements site.Item interface.
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

// New implements site.Item interface
// Returns new site.Item interface
func (video *Video) New() site.Item {
	return site.Item(NewVideo())
}

// Init implements site.Item interface.
// It sets Id or key.
// Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (video *Video) Init(v interface{}) {
	switch v.(type) {
	case int:
		video.Id = Int(v.(int))
	case Int:
		video.Id = v.(Int)
	default:
		panic("Interface v has to be int")
	}
}

func (video *Video) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(video)
}
