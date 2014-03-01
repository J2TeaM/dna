package hdv

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"encoding/binary"
	"errors"
	// "io/ioutil"
	"encoding/base64"
	"time"
	"unicode/utf16"
	"unicode/utf8"
)

type Episode struct {
	MovieId        dna.Int
	EpId           dna.Int
	Title          dna.String
	LinkPlayBackup dna.String
	Link           dna.String
	LinkPlayOther  dna.String
	SubtitleExt    dna.StringArray
	SubtitleExtSe  dna.StringArray
	Subtitle       dna.StringArray
	EpisodeId      dna.Int
	Audiodub       dna.Int
	Audio          dna.Int
	Season         dna.String
	PlaylistM3u8   dna.String
	ViSrt          dna.String // Vietnamese subtitle - encoded in Base64
	EnSrt          dna.String // English subtitle - encoded in Base64
	EpisodeM3u8    dna.String
	Checktime      time.Time
}

func NewEpisode() *Episode {
	episode := new(Episode)
	episode.MovieId = 0
	episode.EpId = 0
	episode.Title = ""
	episode.LinkPlayBackup = ""
	episode.Link = ""
	episode.LinkPlayOther = ""
	episode.SubtitleExt = dna.StringArray{}
	episode.SubtitleExtSe = dna.StringArray{}
	episode.Subtitle = dna.StringArray{}
	episode.EpisodeId = 0
	episode.Audiodub = 0
	episode.Audio = 0
	episode.Season = "{}"
	episode.PlaylistM3u8 = ""
	episode.ViSrt = ""
	episode.EnSrt = ""
	episode.EpisodeM3u8 = ""
	episode.Checktime = time.Now()
	return episode
}
func UTF16ToUTF8String(b []byte, o binary.ByteOrder) dna.String {
	utf := make([]uint16, (len(b)+(2-1))/2)
	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}
	return dna.String(string(utf16.Decode(utf)))
}
func ISO8859_1ToUTF8String(iso8859_1_buf []byte) dna.String {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return dna.String(string(buf))
}
func getSrtContent(episode *Episode, isEn dna.Bool) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		var result *http.Result
		var err error
		if isEn == true {
			result, err = http.Get(episode.SubtitleExt[1])
		} else {
			result, err = http.Get(episode.SubtitleExt[0])
		}

		if err == nil {
			if isEn == true {

				// It is hard to detect an encoding of a string.
				// Therefore we convert them to BASE64
				episode.EnSrt = dna.String(base64.StdEncoding.EncodeToString(result.Data.ToBytes()))
				// episode.EnSrt = ISO8859_1ToUTF8String(result.Data.ToBytes())
				// ioutil.WriteFile("./dump/test_en_srt.srt", result.Data.ToBytes(), 0644)
			} else {
				// Vietnamese Subtitle encoded in UTF-16 Little Ending
				// It has to be converted to UTF-8
				if result.Data.Match(`^[0-9a-fA-F]+$`) == false {
					// episode.ViSrt = UTF16ToUTF8String(result.Data.ToBytes(), binary.LittleEndian)
					episode.ViSrt = dna.String(base64.StdEncoding.EncodeToString(result.Data.ToBytes()))
				}

				// dna.Log(result.Data.Substring(0, 100))
			}
		}
		channel <- true

	}()
	return channel
}

func getPlaylistM3u8(episode *Episode) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		result, err := http.Get(episode.Link)
		if err == nil {
			episode.PlaylistM3u8 = result.Data
			m3u8Filenames := episode.PlaylistM3u8.FindAllString(`.+m3u8`, -1)
			// Get file with highest resolution
			if m3u8Filenames.Length() > 0 {
				lastM3u8Filename := m3u8Filenames[m3u8Filenames.Length()-1]
				baseM3u8Url := episode.Link.ReplaceWithRegexp(`[^/]+m3u8$`, "")
				lastM3u8FilePath := baseM3u8Url + lastM3u8Filename
				resultm3u8, errm3u8 := http.Get(lastM3u8FilePath)
				if errm3u8 == nil {
					episode.EpisodeM3u8 = resultm3u8.Data.ReplaceWithRegexp(`(.+\.ts)`, baseM3u8Url+"${1}")
					// ioutil.WriteFile("./dump/test.m3u8", episode.EpisodeM3u8.ToBytes(), 0644)
				}
			}

		}
		channel <- true

	}()
	return channel
}

func GetEpisode(movieid, ep dna.Int) (*Episode, error) {
	apimovie, err := GetAPIMovie(movieid, ep)
	if err == nil {
		episode := apimovie.ToEpisode()
		c := make(chan bool, 3)
		go func() {
			c <- <-getPlaylistM3u8(episode)
		}()
		go func() {
			c <- <-getSrtContent(episode, false)
		}()
		go func() {
			c <- <-getSrtContent(episode, true)
		}()
		for i := 0; i < 3; i++ {
			<-c
		}
		return episode, nil
	} else {
		return nil, err
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (episode *Episode) Fetch() error {
	_episode, err := GetEpisode(episode.MovieId, episode.EpId)
	if err != nil {
		return err
	} else {
		*episode = *_episode
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (episode *Episode) GetId() dna.Int {
	return ToEpisodeKey(episode.MovieId, episode.EpId)
}

// New implements item.Item interface
// Returns new item.Item interface
func (episode *Episode) New() item.Item {
	return item.Item(NewEpisode())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (episode *Episode) Init(v interface{}) {
	switch v.(type) {
	case int:
		episode.MovieId, episode.EpId = ToMovieIdAndEpisodeId(dna.Int(v.(int)))
	case dna.Int:
		episode.MovieId, episode.EpId = ToMovieIdAndEpisodeId(v.(dna.Int))
	default:
		panic("Interface v has to be int")
	}
}

func (episode *Episode) Save(db *sqlpg.DB) error {
	insertStmt := getInsertStmt(episode, dna.Sprintf("WHERE NOT EXISTS (SELECT 1 FROM %v WHERE movie_id=%v and ep_id=%v)", getTableName(episode), episode.MovieId, episode.EpId))
	_, err := db.Exec(insertStmt.String())
	if err != nil {
		err = errors.New(err.Error() + " $$$error$$$" + insertStmt.String() + "$$$error$$$")
	}
	return err
}
