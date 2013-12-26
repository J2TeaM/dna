package nv

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

var FoundVideos = &dna.IntArray{}

// Define new Album type.
// Notice: Artistid should be Artistids , but this field is not important, then it will be ignored.
type Video struct {
	Id        dna.Int
	Title     dna.String
	Artists   dna.StringArray
	Authors   dna.StringArray
	Topics    dna.StringArray
	Plays     dna.Int
	Lyric     dna.String
	Link      dna.String
	Link320   dna.String
	Thumbnail dna.String
	Type      dna.String
	Checktime time.Time
}

// NewVideo return default new video
func NewVideo() *Video {
	video := new(Video)
	video.Id = 0
	video.Title = ""
	video.Artists = dna.StringArray{}
	video.Authors = dna.StringArray{}
	video.Topics = dna.StringArray{}
	video.Plays = 0
	video.Lyric = ""
	video.Link = ""
	video.Link320 = ""
	video.Thumbnail = ""
	video.Type = ""
	video.Checktime = time.Time{}

	return video
}

// getVideoFromXML returns video from main page
func getVideoFromXML(video *Video) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://hcm.nhac.vui.vn/asx2.php?type=1&id=" + video.Id.ToString()
		result, err := http.Get(link)
		// dna.Log(link)
		if err == nil {
			data := &result.Data
			linkArr := data.FindAllStringSubmatch(`<jwplayer:file><\!\[CDATA\[(.+)\]\]></jwplayer:file>`, 1)
			if len(linkArr) > 0 {
				video.Link = linkArr[0][1].Trim()
				switch {
				case video.Link.Match("mp3") == true:
					video.Type = "song"
				case video.Link.Match("(?mis)mp4") == true:
					video.Type = "video"
				case video.Link.Match("flv") == true:
					video.Type = "video"
				}
			}

		}
		channel <- true

	}()
	return channel
}

func getVideoFromMainPage(video *Video) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://hcm.nhac.vui.vn/google-bot-clip" + video.Id.ToString() + "c1.html"
		// Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data

			tmpArr := data.FindAllStringSubmatch(`<div class="nghenhac-baihat"><h1>(.+)</h1></div>`, 1)
			if len(tmpArr) > 0 {
				tmp := tmpArr[0][1].Trim().Split(` - `)
				video.Title = tmp[0]
				if tmp.Length() > 1 {
					tmp.Shift()
					video.Artists = refineAuthorsOrArtists(tmp.Join(" - "))
				}
			}

			topicsArr := data.FindAllStringSubmatch(`Thể loại: (<a.+?)\|`, 1)
			if len(topicsArr) > 0 {
				video.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().ToStringArray().SplitWithRegexp(` / `).SplitWithRegexp(`/`)
			}

			playsArr := data.FindAllStringSubmatch(`Lượt nghe: (.+?)</div>`, 1)
			if len(playsArr) > 0 {
				video.Plays = playsArr[0][1].Trim().Replace(",", "").ToInt()
			}

			thumbArr := data.FindAllStringSubmatch(`(<meta property="og:image".+)`, 1)
			if len(thumbArr) > 0 {
				video.Thumbnail = thumbArr[0][1].GetTagAttributes("content").Trim()
			}

			lyricArr := data.FindAllString(`(?mis)<div class="nghenhac-loibaihat-cnt.+<div class="loi-bai-hat-footer">`, 1)
			if lyricArr.Length() > 0 {
				lyric := lyricArr[0].Trim().Replace(`<br/>`, "\n").RemoveHtmlTags("").Trim()
				if !lyric.Match(`Hiện bài hát.+chưa có lời.`) {
					video.Lyric = lyric
				}
			}

		}
		channel <- true
	}()
	return channel

}

// GetVideo returns a video and an error (if available)
func GetVideo(id dna.Int) (*Video, error) {
	var video *Video = NewVideo()
	video.Id = id
	c := make(chan bool)

	go func() {
		c <- <-getVideoFromMainPage(video)
	}()
	go func() {
		c <- <-getVideoFromXML(video)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}
	if video.Link == "" {
		return nil, errors.New(dna.Sprintf("Nhacvui - Video %v : Link not found", video.Id).String())
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
	// return db.Update(video, "id", "lyric")
	return db.InsertIgnore(video)
}
