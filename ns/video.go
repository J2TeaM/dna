package ns

import (
	. "dna"
	"dna/http"
	"errors"
	"fmt"
	"math"
	"time"
)

// Define new Album type.
// Notice: Artistid should be Artistids , but this field is not important, then it will be ignored.
type Video struct {
	Id           Int
	Title        String
	Artists      StringArray
	Topics       StringArray
	Plays        Int
	Duration     Int
	Official     Int
	Proceducerid Int
	Link         String
	Sublink      String
	Thumbnail    String
	DateCreated  time.Time
	Checktime    time.Time
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
	video.Proceducerid = 0
	video.Link = ""
	video.Sublink = ""
	video.Thumbnail = ""
	video.DateCreated = time.Time{}
	video.Checktime = time.Time{}

	return video
}

func getVideoDurationAndSublink(id Int) <-chan *Video {
	channel := make(chan *Video, 1)
	go func() {
		video := NewVideo()
		video.Id = id
		link := "http://nhacso.net/flash/video/xnl/1/id/" + GetKey(id)
		// Log(link)
		result, err := http.Get(link)
		if err != nil {
			// do nothing
		} else {
			data := &result.Data
			// getValueXML is in song.go
			// Log(getValueXML(data, "duration", 0))
			video.Duration = getValueXML(data, "duration", 0).RemoveHtmlTags("").ToInt()
			video.Sublink = getValueXML(data, "subUrl", 0)
		}
		channel <- video
	}()
	return channel
}

func getVideoFromMainPage(id Int) <-chan *Video {
	channel := make(chan *Video, 1)
	go func() {
		video := NewVideo()
		link := "http://nhacso.net/xem-video/joke-link." + GetKey(id) + "=.html"
		// Log(link)
		result, err := http.Get(link)
		if err != nil || result.Data.Match("Rất tiếc, chúng tôi không tìm thấy thông tin bạn yêu cầu!") {
			// do nothing cause it has an error
		} else {
			data := &result.Data
			video.Id = id
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
					ts := video.Link.FindAllStringSubmatch(`([0-9]+)_`, 1)[0][1]
					secs := float64(ts.ToInt()) * math.Pow10(13-len(ts))
					// Log(secs)
					video.DateCreated = Float(secs / 1000).ToInt().ToTime()

				}
			}
			proceducerid := data.FindAllStringSubmatch(`getProducerByListIds\('(\d+)', '#producer_'\);`, 1)
			if len(proceducerid) > 0 {
				video.Proceducerid = proceducerid[0][1].ToInt()
			}

		}
		channel <- video
	}()
	return channel

}

// GetVideo returns a video and an error (if available)
func GetVideo(id Int) (*Video, error) {
	var video *Video = NewVideo()
	c := make(chan *Video)

	go func() {
		c <- <-getVideoFromMainPage(id)
	}()

	go func() {
		c <- <-getVideoDurationAndSublink(id)
	}()

	for i := 0; i < 2; i++ {
		tempVideo := <-c
		if tempVideo.Id > 0 {
			video.Id = tempVideo.Id
		}
		if tempVideo.Title != "" {
			video.Title = tempVideo.Title
		}
		if tempVideo.Artists.Length() > 0 {
			video.Artists = tempVideo.Artists
		}
		if tempVideo.Topics.Length() > 0 {
			video.Topics = tempVideo.Topics
		}
		if tempVideo.Plays > 0 {
			video.Plays = tempVideo.Plays
		}
		if tempVideo.Duration > 0 {
			video.Duration = tempVideo.Duration
		}
		if tempVideo.Official > 0 {
			video.Official = tempVideo.Official
		}
		if tempVideo.Proceducerid > 0 {
			video.Proceducerid = tempVideo.Proceducerid
		}
		if tempVideo.Link != "" {
			video.Link = tempVideo.Link
		}
		if tempVideo.Sublink != "" {
			video.Sublink = tempVideo.Sublink
		}
		if tempVideo.Thumbnail != "" {
			video.Thumbnail = tempVideo.Thumbnail
		}
		if !tempVideo.DateCreated.IsZero() {
			video.DateCreated = tempVideo.DateCreated
		}

	}
	if video.Link == "" {
		return nil, errors.New(fmt.Sprintf("Video %v : Link not found", video.Id))
	} else {
		video.Checktime = time.Now()
		return video, nil
	}

}
