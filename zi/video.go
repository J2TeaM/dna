package zi

import (
	. "dna"
	"dna/http"
	"dna/site"
	"dna/sqlpg"
	"errors"
	"time"
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
	return video
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
				datecreatedArr := video.Thumbnail.FindAllStringSubmatch(`_([0-9]+)\..+$`, -1)
				if len(datecreatedArr) > 0 {
					Log(int64(datecreatedArr[0][1].ToInt()))
					video.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()), 0)
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

// GetVideo returns a video or an error
func GetVideo(id Int) (*Video, error) {
	var video *Video = NewVideo()
	video.Id = id
	c := make(chan bool, 1)

	// go func() {
	// 	c <- <-getSongFromXML(video)
	// }()
	go func() {
		c <- <-getVideoFromMainPage(video)
	}()

	for i := 0; i < 1; i++ {
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
