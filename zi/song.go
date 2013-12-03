package zi

import (
	. "dna"
	"dna/http"
	"dna/site"
	"dna/sqlpg"
	"errors"
	"time"
)

const (
	LBr128      = 1 << iota // Flag of  128kbps bitrate
	LBr320                  // Flag of 320kbps bitrate
	LBrLossless             // Flag of lossless bitrate
)

// The basic song type
type Song struct {
	Id          Int
	Key         String
	Title       String
	Artists     StringArray
	Authors     StringArray
	Plays       Int
	Topics      StringArray
	Link        String
	Path        String
	Lyric       String
	DateCreated time.Time
	Checktime   time.Time
	// Add more 11 fields from official API
	ArtistId       Int
	VideoId        Int
	AlbumId        Int
	IsHit          Int
	IsOfficial     Int
	DownloadStatus Int
	Copyright      String
	BitrateFlags   Int
	Likes          Int
	Comments       Int
	Thumbnail      String
}

// Song constructor with key
func NewSong() *Song {
	song := new(Song)
	song.Key = ""
	song.Id = 0
	song.Title = ""
	song.Artists = StringArray{}
	song.Authors = StringArray{}
	song.Plays = 0
	song.Topics = StringArray{}
	song.Link = ""
	song.Path = ""
	song.Lyric = ""
	song.DateCreated = time.Time{}
	song.Checktime = time.Time{}
	// Add more 10 fields
	song.ArtistId = 0
	song.VideoId = 0
	song.AlbumId = 0
	song.IsHit = 0
	song.IsOfficial = 0
	song.DownloadStatus = 0
	song.Copyright = ""
	song.BitrateFlags = 0
	song.Likes = 0
	song.Comments = 0
	song.Thumbnail = ""
	return song
}

func getSongFromXML(song *Song) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://m.mp3.zing.vn/xml/song/" + song.GetEncodedKey(Bitrate128)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			linkArr := data.FindAllStringSubmatch(`"source":"(.+?)"`, -1)
			if len(linkArr) > 0 {
				song.Link = linkArr[0][1].Replace(`\/`, `/`)
				pathArr := song.Link.FindAllStringSubmatch(`song-load/(.+)`, -1)
				if len(linkArr) > 0 {
					song.Path = DecodePath(pathArr[0][1])
					dateCreatedArr := song.Path.FindAllStringSubmatch(`^/(\d{4}/\d{2}/\d{2})`, -1)
					if len(dateCreatedArr) > 0 {
						year := dateCreatedArr[0][1].FindAllStringSubmatch(`^(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
						month := dateCreatedArr[0][1].FindAllStringSubmatch(`^\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
						day := dateCreatedArr[0][1].FindAllStringSubmatch(`^\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
						song.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

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
		link := "http://mp3.zing.vn/bai-hat/joke-link/" + song.Key + ".html"
		result, err := http.Get(link)
		// Log(link)
		// Log(result.Data)
		if err == nil {
			data := &result.Data

			titleArr := data.FindAllStringSubmatch(`<h1 class="detail-title">(.+?)</h1>`, -1)
			if len(titleArr) > 0 {
				song.Title = titleArr[0][1].Trim()
			}

			artistsArr := data.FindAllStringSubmatch(`<h1 class="detail-title">.+</h1><span>-</span>(.+)`, -1)
			if len(artistsArr) > 0 {
				artists := artistsArr[0][1].RemoveHtmlTags("").Trim().Split(", ").SplitWithRegexp("ft. ")
				song.Artists = StringArray(artists.Map(func(val String, i Int) String {
					return val.Trim()
				}).([]String))
			}

			playsArr := data.FindAllStringSubmatch(`Lượt nghe: (.+)</p>`, -1)
			if len(playsArr) > 0 {
				song.Plays = playsArr[0][1].Replace(".", "").ToInt()
			}

			authorsArr := data.FindAllStringSubmatch(`Sáng tác:(.+?)\|`, -1)
			if len(authorsArr) > 0 {
				authors := authorsArr[0][1].RemoveHtmlTags("").Trim().Split(", ").SplitWithRegexp(" / ").SplitWithRegexp(" & ")
				song.Authors = StringArray(authors.Map(func(val String, idx Int) String {
					switch {
					case val == "Đang Cập Nhật":
						return ""
					case val == "Đang cập nhật":
						return ""
					default:
						return val
					}
				}).([]String)).Filter(func(v String, i Int) Bool {
					if v != "" {
						return true
					} else {
						return false
					}
				})
			}

			topicsArr := data.FindAllStringSubmatch(`Thể loại:(.+?)\|`, -1)
			if len(topicsArr) > 0 {
				song.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().Split(", ").SplitWithRegexp(" / ").Unique()
			}

			lyricsArr := data.FindAllStringSubmatch(`(?mis)<p class="_lyricContent.+</span></span>(.+?)</p>`, -1)
			if len(lyricsArr) > 0 {
				song.Lyric = lyricsArr[0][1].Trim()
			}

		}
		channel <- true

	}()
	return channel
}

// GetSong returns a song or an error
func GetSong(id Int) (*Song, error) {
	var song *Song = NewSong()
	song.Id = id
	song.Key = GetKey(id)
	c := make(chan bool, 2)

	go func() {
		c <- <-getSongFromXML(song)
	}()
	go func() {
		c <- <-getSongFromMainPage(song)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}

	if song.Link == "" {
		return nil, errors.New(Sprintf("Zing - Song %v: Mp3 link not found", song.Id).String())
	} else {
		song.Checktime = time.Now()
		return song, nil
	}
}

// Getting encoded key used for XML file or direct link
func (song *Song) GetEncodedKey(bitrate Bitrate) String {
	var temp IntArray
	if bitrate == Lossless {
		temp = IntArray{11, 12, 13, 13, 11, 14, 13, 13}
	} else {
		temp = Int(bitrate).ToString().Split("").ToIntArray()
	}
	tailArray := IntArray{10}.Concat(temp).Concat(IntArray{10, 2, 0, 1, 0})
	return getCipherText(GetId(song.Key), tailArray)

}

// Notice: The interface of getting mp3 direct link with high quality( 320kbps or lossless) has been deprecated. Need to check it out later!!!
func (song *Song) GetDirectLink(bitrate Bitrate) String {
	return SONG_BASE_URL.Concat(song.GetEncodedKey(bitrate), "/")
}

// Fetch implements site.Item interface.
// Returns error if can not get item
func (song *Song) Fetch() error {
	_song, err := GetSongFromAPI(song.Id)
	if err != nil {
		return err
	} else {
		*song = *_song
		return nil
	}
}

// New implements site.Item interface
// Returns new site.Item interface
func (song *Song) New() site.Item {
	return site.Item(NewSong())
}

// Init implements site.Item interface.
// It sets Id or key.
// Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (song *Song) Init(v interface{}) {
	switch v.(type) {
	case int:
		song.Id = Int(v.(int))
	case Int:
		song.Id = v.(Int)
	// case string:
	// 	song.Key = String(v.(string))
	// case String:
	// 	song.Key = v.(String)
	default:
		panic("Interface v has to be int")
	}
}

func (song *Song) Save(db *sqlpg.DB) error {
	return db.Update(song, "id", "artist_id", "video_id", "album_id", "is_hit", "is_official", "download_status", "copyright", "bitrate_flags", "likes", "comments", "thumbnail")
}
