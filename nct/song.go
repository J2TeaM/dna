package nct

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

// Song defines a song type
type Song struct {
	Id        dna.Int
	Key       dna.String
	Title     dna.String
	Artists   dna.StringArray
	Topics    dna.StringArray
	Plays     dna.Int
	Type      dna.String
	Bitrate   dna.Int
	Official  dna.Int
	LinkKey   dna.String
	Lyric     dna.String
	HasLrc    dna.Bool
	LrcUrl    dna.String
	Lrc       dna.String
	Checktime time.Time
}

// NewSong returns new song with default settings.
func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Key = ""
	song.Title = ""
	song.Plays = 0
	song.Artists = dna.StringArray{}
	song.Topics = dna.StringArray{}
	song.Type = ""
	song.Bitrate = 0
	song.Official = 0
	song.LinkKey = ""
	song.Lyric = ""
	song.HasLrc = false
	song.LrcUrl = ""
	song.Lrc = ""
	song.Checktime = time.Time{}
	return song
}

func getSongLrc(song *Song) {
	result, err := http.Get(song.LrcUrl)
	if err == nil {
		lrc, derr := DecryptLRC(result.Data)
		if derr == nil {
			song.Lrc = lrc
		} else {
			dna.Log("ERR WHILE DECRYPT SONG ", song.Id)
			dna.Log("-----\n")
		}
	}
}

func getSongLrcUrl(song *Song) {

	link := "http://www.nhaccuatui.com/flash/xml?key1=" + song.LinkKey
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		lrcArr := data.FindAllStringSubmatch(`<lyric><\!\[CDATA\[(.+)\]\]></lyric>`, 1)
		if len(lrcArr) > 0 {
			song.LrcUrl = lrcArr[0][1].Trim()
			getSongLrc(song)
		}
	}

}

// getSongPlays returns song plays
func getSongPlays(song *Song, body dna.String) {
	// POST METHOD
	// link := "http://www.nhaccuatui.com/interaction/api/hit-counter?jsoncallback=nct"
	// http.DefaulHeader.Set("Content-Type", "application/x-www-form-urlencoded ")
	// result, err := http.Post(dna.String(link), body)
	// // Log(link)
	// if err == nil {
	// 	data := &result.Data
	// 	tpl := dna.String(`{"counter":([0-9]+)}`)
	// 	playsArr := data.FindAllStringSubmatch(tpl, -1)
	// 	if len(playsArr) > 0 {
	// 		song.Plays = playsArr[0][1].ToInt()
	// 	}
	// }
	// GET METHOD
	link := "http://www.nhaccuatui.com/interaction/api/counter?jsoncallback=nct&listSongIds=" + song.Id.ToString()
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		tpl := dna.Sprintf(`{"%v":([0-9]+)}`, song.Id)
		// dna.Log(data)
		playsArr := data.FindAllStringSubmatch(tpl, -1)
		if len(playsArr) > 0 {
			song.Plays = playsArr[0][1].ToInt()
		}
	}
}

// getSongFromMainPage returns song from main page
func getSongFromMainPage(song *Song) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.nhaccuatui.com/bai-hat/google-bot." + song.Key + ".html"
		result, err := http.Get(link)
		// Log(link)
		// Log(result.Data)
		if err == nil {
			data := &result.Data

			topicsArr := data.FindAllStringSubmatch(`<strong>Thể loại</strong></p>[\n\t\r]+(.+)`, 1)
			if len(topicsArr) > 0 {
				song.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().Split(", ")
			}

			if data.Match(`<span class="tag.+official`) == true {
				song.Official = 1
			}

			titleArr := data.FindAllStringSubmatch(`<h1 itemprop="name">(.+?)</h1>`, 1)
			if len(titleArr) > 0 {
				song.Title = titleArr[0][1].Trim().SplitWithRegexp(" - ", 2)[0].Trim()
			}

			artistsArr := data.FindAllStringSubmatch(`<h1 itemprop="name">(.+?)</h1>`, 1)
			if len(artistsArr) > 0 {
				artists := artistsArr[0][1].RemoveHtmlTags("").SplitWithRegexp(" - ", 2)
				if artists.Length() == 2 {
					song.Artists = artists[1].Split(", ").Filter(func(v dna.String, i dna.Int) dna.Bool {
						if v != "" {
							return true
						} else {
							return false
						}
					})
				}
			}

			linkKeyArr := data.FindAllStringSubmatch(`file=http://www.nhaccuatui.com/flash/xml\?key1=(.+?)"`, -1)
			if len(linkKeyArr) > 0 {
				song.LinkKey = linkKeyArr[0][1].Trim()
			}

			lyricArr := data.FindAllStringSubmatch(`(?mis)<div id="divLyric".+?>(.+?)<div class="more_add".+?/>`, -1)
			if len(lyricArr) > 0 {
				song.Lyric = lyricArr[0][1].DecodeHTML().Trim().ReplaceWithRegexp(`</div>$`, "")
			}

			bitrate := data.FindAllStringSubmatch(`<span class="tag orange">(.+?)k</span>`, -1)
			if len(bitrate) > 0 {
				song.Bitrate = bitrate[0][1].ToInt()
			}

			// Find params for the number of song plays
			// POST METHOD
			// itemIdArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('(.+?)'.+`, 1)
			// timeArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('.+?', '(.+?)'.+\);`, 1)
			// signArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('.+?', '.+?', '(.+?)'.+;`, 1)
			// typeArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('.+?', '.+?', '.+?', "(.+?)"\);`, 1)
			// if len(itemIdArr) > 0 && len(timeArr) > 0 && len(signArr) > 0 && len(typeArr) > 0 {
			// 	// boday has post form:
			// 	// item_id=2870710&time=1389009424631&sign=2499ab08f6662842a02b06aad603d8ab&type=song
			// 	body := dna.Sprintf(`item_id=%v&time=%v&sign=%v&type=%v`, itemIdArr[0][1], timeArr[0][1], signArr[0][1], typeArr[0][1])
			// 	song.Type = typeArr[0][1].Trim()
			// 	song.Id = itemIdArr[0][1].ToInt()
			// 	getSongPlays(song, body)
			// }
			// GET METHOD
			itemIdArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('(.+?)'.+`, 1)
			typeArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('.+?', '.+?', '.+?', "(.+?)"\);`, 1)
			if len(itemIdArr) > 0 && len(typeArr) > 0 {
				// boday has post form:
				// item_id=2870710&time=1389009424631&sign=2499ab08f6662842a02b06aad603d8ab&type=song
				song.Type = typeArr[0][1].Trim()
				song.Id = itemIdArr[0][1].ToInt()
				getSongPlays(song, "")
			}

			GetRelevantPortions(&result.Data)

			// Getting LRC
			if data.Match(`<span class="tag.+KARAOKE`) == true {
				song.HasLrc = true
				getSongLrcUrl(song)
			}

		}
		channel <- true
	}()
	return channel
}

// GetSong returns a song or an error
// 	* key: A unique key of a song
// 	* Official : 0 or 1, if its value is unknown, set to 0
// 	* Returns a found song or an error
func GetSong(key dna.String) (*Song, error) {
	var song *Song = NewSong()
	song.Key = key
	c := make(chan bool, 1)
	go func() {
		c <- <-getSongFromMainPage(song)
	}()
	for i := 0; i < 1; i++ {
		<-c
	}
	// getSongPlays(song)
	if song.LinkKey == "" {
		return nil, errors.New(dna.Sprintf("Nhaccuatui - Song id %v:  - Key: %v  Link key not found", song.Id, song.Key).String())
	} else {
		song.Checktime = time.Now()
		return song, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (song *Song) Fetch() error {
	_song, err := GetSong(song.Key)
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
		idx := dna.Int(v.(int))
		length := NewestSongPortions.Length()
		if idx >= length {
			idx = length - 1
		}
		if length > 0 {
			song.Key = NewestSongPortions[idx]
		}

	case dna.Int:
		idx := v.(dna.Int)
		length := NewestSongPortions.Length()
		if idx >= length {
			idx = length - 1
		}
		if length > 0 {
			song.Key = NewestSongPortions[idx]
		}

	default:
		panic("Interface v has to be int")
	}
}

func (song *Song) Save(db *sqlpg.DB) error {
	FilterRelevants(db)
	return db.InsertIgnore(song)
}
