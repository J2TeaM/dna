package ns

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"fmt"
	"math"
	"time"
)

type Song struct {
	Id          dna.Int
	Title       dna.String
	Artists     dna.StringArray
	Artistid    dna.Int
	Authors     dna.StringArray
	Authorid    dna.Int
	Plays       dna.Int
	Duration    dna.Int
	Link        dna.String
	Topics      dna.StringArray
	Category    dna.StringArray
	Bitrate     dna.Int
	Official    dna.Int
	Islyric     dna.Int
	DateCreated time.Time
	DateUpdated time.Time
	Lyric       dna.String
	SameArtist  dna.Int
	Checktime   time.Time
}

// NewSong returns new song whose id is 0
func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Artists = dna.StringArray{}
	song.Artistid = 0
	song.Authors = dna.StringArray{}
	song.Authorid = 0
	song.Plays = 0
	song.Duration = 0
	song.Link = ""
	song.Topics = dna.StringArray{}
	song.Category = dna.StringArray{}
	song.Bitrate = 0
	song.Official = 0
	song.Islyric = 0
	song.Lyric = ""
	song.DateCreated = time.Time{}
	song.DateUpdated = time.Time{}
	song.SameArtist = 0
	song.Checktime = time.Time{}
	return song
}

func getValueXML(data *dna.String, tag dna.String, position dna.Int) dna.String {
	v := (*data).FindAllString("<"+tag+">.+<\\/"+tag+">", -1)
	if v.Length() > position {
		return v[position].ReplaceWithRegexp(`\]\].+$`, "").ReplaceWithRegexp(`^.+CDATA\[`, "")
	} else {
		return ""
	}
}

// getSongFromMainPage returns song from main page
func getSongFromMainPage(song *Song) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/nghe-nhac/google-bot." + GetKey(song.Id) + "==.html"
		result, err := http.Get(link)
		// Log(link)
		// Log(result.Data)
		if err == nil && !result.Data.Match("Rất tiếc, chúng tôi không tìm thấy thông tin bạn yêu cầu!") {
			data := &result.Data
			if data.Match("official") {
				song.Official = 1
			}

			bitrate := data.FindAllString(`\d+kb\/s`, 1)[0]
			if !bitrate.IsBlank() {
				song.Bitrate = bitrate.FindAllString(`\d+`, 1)[0].ToInt()
			}

			plays := data.FindAllString("total_listen_song_detail_\\d+.+", 1)[0]
			if !plays.IsBlank() {
				song.Plays = plays.ReplaceWithRegexp("<\\/span>.+$", "").ReplaceWithRegexp("^.+>", "").ReplaceWithRegexp("\\.", "").ToInt()
			}

			topics := data.FindAllString("<li><a\\shref\\=\\\"http\\:\\/\\/nhacso\\.net\\/the-loai.+", 1)[0]
			if !topics.IsBlank() {
				topics = topics.ReplaceWithRegexp("^.+\\\">|<\\/a><\\/li>", "")
				song.Topics = topics.ToStringArray().SplitWithRegexp(" - ").SplitWithRegexp("/")
				temp := data.FindAllString("<a.+class=\"casi\".+>(.+?)<\\/a>", 1)
				var singer dna.String
				if temp.Length() > 0 {
					singer = temp[0]
				} else {
					singer = ""
				}
				if topics.Match("Nhạc Hoa") && singer.Match(` / `) {
					song.SameArtist = 1
				}
			}

			lyric := data.FindAllString(`(?mis)txtlyric.+Bạn chưa nhập nội bài hát`, 1)[0]
			if !lyric.IsBlank() {
				song.Islyric = 1
				song.Lyric = lyric.ReplaceWithRegexp("(?mis)<\\/textarea>.+$", "").ReplaceWithRegexp("^.+>", "")
				if song.Lyric.Match("Hãy đóng góp lời bài hát chính xác cho Nhacso nhé") {
					song.Lyric = ``
					song.Islyric = 0
				}
			}
		}
		channel <- true

	}()
	return channel
}

// getSongFromXML returns values from url: http://nhacso.net/flash/song/xnl/1/id/
func getSongFromXML(song *Song) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/flash/song/xnl/1/id/" + GetKey(song.Id)
		result, err := http.Get(link)
		if err == nil {
			song.Title = getValueXML(&result.Data, "name", 1).Trim()
			song.Artists = getValueXML(&result.Data, "artist", 0).ToStringArray().SplitWithRegexp("\\|\\|").SplitWithRegexp(" / ").SplitWithRegexp(" - ")
			song.Artistid = getValueXML(&result.Data, "artistlink", 0).ReplaceWithRegexp("\\.html", "").ReplaceWithRegexp(`^.+-`, "").ToInt()
			authors := getValueXML(&result.Data, "author", 0)
			if !authors.IsBlank() {
				song.Authors = authors.ToStringArray().SplitWithRegexp("\\|\\|").SplitWithRegexp(" / ").SplitWithRegexp(" - ")
				song.Authorid = getValueXML(&result.Data, "authorlink", 0).ReplaceWithRegexp(`\.html`, "").ReplaceWithRegexp(`^.+-`, "").ToInt()

			}
			duration := result.Data.FindAllString("<totalTime.+totalTime>", 1)
			if duration.Length() > 0 {
				song.Duration = duration[0].RemoveHtmlTags("").Trim().ToInt()
			}

			song.Link = getValueXML(&result.Data, "mp3link", 0)

			if song.Title != "" && song.Link != "/" {
				ts := song.Link.FindAllString(`\/[0-9]+_`, 1)[0].ReplaceWithRegexp(`\/`, "").ReplaceWithRegexp(`_`, "")
				unix := ts.ToInt().ToFloat() * dna.Float(math.Pow10(13-len(ts)))
				song.DateCreated = dna.Int(int64(unix) / 1000).ToTime()
				song.DateUpdated = time.Now()
			}
		}
		channel <- true

	}()
	return channel
}

// GetSong returns a song or an error
func GetSong(id dna.Int) (*Song, error) {
	var song *Song = NewSong()
	song.Id = id
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

	if song.Link == "" || song.Link == "/" {
		return nil, errors.New(fmt.Sprintf("Nhacso - Song %v: Mp3 link not found", song.Id))
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
	// case string:
	// 	song.Key = dna.String(v.(string))
	// case dna.String:
	// 	song.Key = v.(String)
	default:
		panic("Interface v has to be int")
	}
}

func (song *Song) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(song)
}
