package nct

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

type Video struct {
	Id          dna.Int
	Key         dna.String
	Title       dna.String
	Artists     dna.StringArray
	Topics      dna.StringArray
	Plays       dna.Int
	Duration    dna.Int
	Thumbnail   dna.String
	Type        dna.String
	LinkKey     dna.String
	Lyric       dna.String
	DateCreated time.Time
	Checktime   time.Time
}

func NewVideo() *Video {
	video := new(Video)
	video.Id = 0
	video.Key = ""
	video.Title = ""
	video.Artists = dna.StringArray{}
	video.Topics = dna.StringArray{}
	video.Plays = 0
	video.Duration = 0
	video.Thumbnail = ""
	video.Type = ""
	video.LinkKey = ""
	video.Lyric = ""
	video.DateCreated = time.Time{}
	video.Checktime = time.Time{}
	return video
}

// getVideoPlays returns video plays
func getVideoPlays(video *Video) {
	link := "http://www.nhaccuatui.com/wg/get-counter?listVideoIds=" + video.Id.ToString()
	result, err := http.Get(link)
	// dna.Log(link)
	if err == nil {
		data := &result.Data
		tpl := dna.Sprintf(`{"%v":([0-9]+)}`, video.Id)
		playsArr := data.FindAllStringSubmatch(tpl, -1)
		if len(playsArr) > 0 {
			video.Plays = playsArr[0][1].ToInt()
		}
	}
}

// getVideoFromMainPage returns video from main page
func getVideoFromMainPage(video *Video) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.nhaccuatui.com/video/google-bot." + video.Key + ".html"
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			// new version in June 2013 does not support lyric of videos

			idArr := data.FindAllStringSubmatch(`value="(.+)" id="inpHiddenId"`, 1)
			if len(idArr) > 0 {
				video.Id = idArr[0][1].ToInt()
			}

			titleArr := data.FindAllStringSubmatch(`(?mis)<h1 class="songname">(.+?)- <a`, 1)
			if len(titleArr) > 0 {
				video.Title = titleArr[0][1].Trim()
			}

			artistArr := data.FindAllStringSubmatch(`(?mis)<h1 class="songname">.+(<a.+?)</h1>`, 1)
			if len(artistArr) > 0 {
				video.Artists = artistArr[0][1].RemoveHtmlTags("").Trim().Split(", ")
			}

			topicArr := data.FindAllStringSubmatch(`Thể loại: (<a.+?</a></p>)`, 1)
			if len(topicArr) > 0 {
				video.Topics = topicArr[0][1].RemoveHtmlTags("").Trim().Split(",&nbsp;")
			}

			thumbArr := data.FindAllString(`<link rel="image_src".+`, 1)
			if thumbArr.Length() > 0 {
				video.Thumbnail = thumbArr[0].GetTagAttributes("href").Trim()
				datecreatedArr := video.Thumbnail.FindAllStringSubmatch(`/([0-9]+)_[0-9]+\..+$`, -1)
				if len(datecreatedArr) > 0 {
					// Log(int64(datecreatedArr[0][1].ToInt()))
					video.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()/1000), 0)
				} else {
					dateCreatedArr := video.Thumbnail.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
					if len(dateCreatedArr) > 0 {
						year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
						month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
						day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
						video.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
					}
				}
			}

			typeArr := data.FindAllStringSubmatch(`value="(.+)" id="inpHiddenType"`, -1)
			if len(typeArr) > 0 {
				video.Type = typeArr[0][1].Trim()
			}

			linkkeyArr := data.FindAllStringSubmatch(`"flashPlayer", ".+", "(.+?)"`, 1)
			if len(linkkeyArr) > 0 {
				video.LinkKey = linkkeyArr[0][1].Trim()
			}

			GetRelevantPortions(&result.Data)
		}
		channel <- true
	}()
	return channel
}

// GetVideo returns a video or an error
// 	* key: A unique key of a video
// 	* Official : 0 or 1, if its value is unknown, set to 0
// 	* Returns a found video or an error
func GetVideo(key dna.String) (*Video, error) {
	var video *Video = NewVideo()
	video.Key = key
	c := make(chan bool, 1)
	go func() {
		c <- <-getVideoFromMainPage(video)
	}()
	for i := 0; i < 1; i++ {
		<-c
	}
	getVideoPlays(video)

	if video.LinkKey == "" {
		return nil, errors.New(dna.Sprintf("Nhaccuatui - Video id:%v, key:%v Link key not found", video.Id, video.Key).String())
	} else {
		video.Checktime = time.Now()
		return video, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (video *Video) Fetch() error {
	_video, err := GetVideo(video.Key)
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
		idx := dna.Int(v.(int))
		length := (*NewestVideoPortions).Length()
		if idx >= length {
			idx = length - 1
		}
		if length > 0 {
			video.Key = dna.String((*NewestVideoPortions)[idx].Key)
		}
	case dna.Int:
		idx := v.(dna.Int)
		length := (*NewestVideoPortions).Length()
		if idx >= length {
			idx = length - 1
		}
		if length > 0 {
			video.Key = dna.String((*NewestVideoPortions)[idx].Key)
		}

	default:
		panic("Interface v has to be int")
	}
}

func (video *Video) Save(db *sqlpg.DB) error {
	filterRelevants(db)
	return db.InsertIgnore(video)
}
