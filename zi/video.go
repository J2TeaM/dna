package zi

import (
	. "dna"
	"dna/http"
	"dna/site"
	"dna/sqlpg"
	"errors"
	"time"
)

// Defines resolution constant flags with correspondent values: 1, 2, 4, 8, 16
const (
	LRe240p  = 1 << iota // Flag of  240p resolution
	LRe360p              // Flag of 360p resolution
	LRe480p              // Flag of 480p resolution
	LRe720p              // Flag of 720p resolution
	LRe1080p             // Flag of 1080p resolutions
)

// Video defines a basic video type
type Video struct {
	Id          Int
	Title       String
	Artists     StringArray
	Topics      StringArray
	Plays       Int
	Thumbnail   String
	Link        String
	Lyric       String
	DateCreated time.Time
	Checktime   time.Time
	// add new 6 fields
	ArtistIds       IntArray
	Duration        Int
	StatusId        Int
	ResolutionFlags Int
	Likes           Int
	Comments        Int
}

// NewVideo returns a pointer to a new video
func NewVideo() *Video {
	video := new(Video)
	video.Id = 0
	video.Title = ""
	video.Artists = StringArray{}
	video.Topics = StringArray{}
	video.Plays = 0
	video.Thumbnail = ""
	video.Lyric = ""
	video.Link = ""
	video.DateCreated = time.Time{}
	video.Checktime = time.Time{}
	// add new 6 fields
	video.ArtistIds = IntArray{}
	video.Duration = 0
	video.StatusId = 0
	video.Likes = 0
	video.Comments = 0
	video.ResolutionFlags = 0
	return video
}

//GetVideoFromAPI gets a video from API. It does not get content from main site.
func GetVideoFromAPI(id Int) (*Video, error) {
	var video *Video = NewVideo()
	video.Id = id
	apivideo, err := GetAPIVideo(id)
	if err != nil {
		return nil, err
	} else {
		if apivideo.Response.MsgCode == 1 {
			video.Title = apivideo.Title
			video.Artists = StringArray(apivideo.Artists.Split(" , ").Map(func(val String, idx Int) String {
				return val.Trim()
			}).([]String)).Filter(func(v String, i Int) Bool {
				if v != "" {
					return true
				} else {
					return false
				}
			})
			video.Topics = StringArray(apivideo.Topics.Split(", ").Map(func(val String, idx Int) String {
				return val.Trim()
			}).([]String)).SplitWithRegexp(" / ").Unique().Filter(func(v String, i Int) Bool {
				if v != "" {
					return true
				} else {
					return false
				}
			})
			video.Plays = apivideo.Plays
			video.Thumbnail = "http://image.mp3.zdn.vn/" + apivideo.Thumbnail
			// video.Lyric =

			datecreatedArr := video.Thumbnail.FindAllStringSubmatch(`_([0-9]+)\..+$`, -1)
			if len(datecreatedArr) > 0 && datecreatedArr[0][1].Length() > 6 {
				video.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()), 0)
			} else {
				dateCreatedArr := video.Thumbnail.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
				if len(dateCreatedArr) > 0 {
					year := dateCreatedArr[0][1].FindAllStringSubmatch(`^(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
					month := dateCreatedArr[0][1].FindAllStringSubmatch(`^\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
					day := dateCreatedArr[0][1].FindAllStringSubmatch(`^\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
					video.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

				}
			}
			video.Plays = apivideo.Plays
			video.ArtistIds = apivideo.ArtistIds.Split(",").ToIntArray()
			flags := 0
			for key, val := range apivideo.Source {
				video.Link = val
				switch {
				case key == "240" && val != "":
					flags = flags | LRe240p
				case key == "360" && val != "":
					flags = flags | LRe360p
				case key == "480" && val != "":
					flags = flags | LRe480p
				case key == "720" && val != "":
					flags = flags | LRe720p
				case key == "1080" && val != "":
					flags = flags | LRe1080p
				}
			}
			video.ResolutionFlags = Int(flags)
			video.Duration = apivideo.Duration
			video.Likes = apivideo.Likes
			video.StatusId = apivideo.StatusId
			video.Comments = apivideo.Comments
			video.Checktime = time.Now()
			return video, nil
		} else {
			return nil, errors.New("Message code invalid " + apivideo.Response.MsgCode.ToString().String())
		}
	}
}

// getVideoFromMainPage returns song from main page
func getVideoFromMainPage(video *Video) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://mp3.zing.vn/video-clip/google-bot/" + GetKey(video.Id) + ".html"
		result, err := http.Get(link)
		// Log(link)
		// Log(result.Data)
		data := &result.Data
		Log(data.Match("<title>Thông báo</title>"))
		if err == nil && !data.Match("<title>Thông báo</title>") {

			topicsArr := data.FindAllStringSubmatch(`Thể loại:(.+)\|`, -1)
			if len(topicsArr) > 0 {
				video.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().Split(", ").SplitWithRegexp(` / `).Unique()
			}

			playsArr := data.FindAllStringSubmatch(`Lượt xem:(.+)</p>`, -1)
			if len(playsArr) > 0 {
				video.Plays = playsArr[0][1].Trim().Replace(".", "").ToInt()
			}

			titleArr := data.FindAllStringSubmatch(`<h1 class="detail-title">(.+?)</h1>`, -1)
			if len(titleArr) > 0 {
				video.Title = titleArr[0][1].RemoveHtmlTags("").Trim()
			}

			artistsArr := data.FindAllStringSubmatch(`<h1 class="detail-title">.+(<a.+)`, -1)
			if len(artistsArr) > 0 {
				video.Artists = StringArray(artistsArr[0][1].RemoveHtmlTags("").Trim().Split(" ft. ").Unique().Map(func(val String, idx Int) String {
					return val.Trim()
				}).([]String))
			}

			thumbnailArr := data.FindAllString(`<meta property="og:image".+`, -1)
			if thumbnailArr.Length() > 0 {
				video.Thumbnail = thumbnailArr[0].GetTagAttributes("content")
				// datecreatedArr := video.Thumbnail.FindAllStringSubmatch(`_([0-9]+)\..+$`, -1)
				// if len(datecreatedArr) > 0 {
				// 	// Log(int64(datecreatedArr[0][1].ToInt()))
				// 	video.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()), 0)
				// }
				datecreatedArr := video.Thumbnail.FindAllStringSubmatch(`_([0-9]+)\..+$`, -1)
				if len(datecreatedArr) > 0 && datecreatedArr[0][1].Length() > 6 {
					video.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()), 0)
				} else {
					dateCreatedArr := video.Thumbnail.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
					if len(dateCreatedArr) > 0 {
						year := dateCreatedArr[0][1].FindAllStringSubmatch(`^(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
						month := dateCreatedArr[0][1].FindAllStringSubmatch(`^\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
						day := dateCreatedArr[0][1].FindAllStringSubmatch(`^\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
						video.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

					}
				}
			}

			video.Link = video.GetDirectLink(Resolution360p)

			lyricArr := data.FindAllStringSubmatch(`(?mis)<p class="_lyricContent.+</span></span>(.+?)<p class="seo">`, -1)
			if len(lyricArr) > 0 {
				video.Lyric = lyricArr[0][1].ReplaceWithRegexp(`(?mis)<p class="seo">.+`, "").Trim().Replace("</p>\r\n\t</div>\r\n\t\t</div>", "").Trim()
			}

		}
		channel <- true

	}()
	return channel
}

func getVideoLyricFromApi(video *Video) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		videoLyric, err := GetAPIVideoLyric(video.Id)
		if err == nil {
			video.Lyric = videoLyric.Content
		}
		channel <- true

	}()
	return channel
}

// getVideoFromAPI returns video from API
func getVideoFromAPI(video *Video) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		apivideo, err := GetVideoFromAPI(video.Id)
		if err == nil {
			video.Title = apivideo.Title
			video.Artists = apivideo.Artists
			video.Topics = apivideo.Topics
			video.Plays = apivideo.Plays
			video.Thumbnail = apivideo.Thumbnail
			video.DateCreated = apivideo.DateCreated
			video.Plays = apivideo.Plays
			video.ArtistIds = apivideo.ArtistIds
			video.Link = apivideo.Link
			video.ResolutionFlags = apivideo.ResolutionFlags
			video.Duration = apivideo.Duration
			video.Likes = apivideo.Likes
			video.StatusId = apivideo.StatusId
			video.Comments = apivideo.Comments
			video.Checktime = time.Now()
		}
		channel <- true

	}()
	return channel
}

// GetVideo returns a video or an error
func GetVideo(id Int) (*Video, error) {
	var video *Video = NewVideo()
	video.Id = id
	c := make(chan bool, 2)

	go func() {
		c <- <-getVideoLyricFromApi(video)
	}()
	go func() {
		c <- <-getVideoFromAPI(video)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}

	if video.Link == "" {
		return nil, errors.New(Sprintf("Zing - Video %v: Link not found", video.Id).String())
	} else {
		video.Checktime = time.Now()
		return video, nil
	}

}

// GetEncodedKey gets a encoded key used for XML link or getting direct video url
func (video *Video) GetEncodedKey(resolution Resolution) String {
	tailArray := IntArray{10}.Concat(Int(resolution).ToString().Split("").ToIntArray()).Concat(IntArray{10, 2, 0, 1, 0})
	return getCipherText(video.Id, tailArray)
}

// GetDirectLink gets a direct video link from the site with various qualities
func (video *Video) GetDirectLink(resolution Resolution) String {
	return VIDEO_BASE_URL.Concat(video.GetEncodedKey(resolution), "/")
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
