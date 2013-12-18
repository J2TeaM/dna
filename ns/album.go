package ns

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"fmt"
	"time"
)

// Define new Album type.
// Notice: Artistid should be Artistids , but this field is not important, then it will be ignored.
type Album struct {
	Id           dna.Int
	Title        dna.String
	Artists      dna.StringArray
	Artistid     dna.Int
	Topics       dna.StringArray
	Genres       dna.StringArray
	Category     dna.StringArray
	Coverart     dna.String
	Nsongs       dna.Int
	Plays        dna.Int
	Songids      dna.IntArray
	Description  dna.String
	Label        dna.String
	DateReleased dna.String
	Checktime    time.Time
}

// NewAlbum return default new album
func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Title = ""
	album.Artists = dna.StringArray{}
	album.Artistid = 0
	album.Topics = dna.StringArray{}
	album.Genres = dna.StringArray{}
	album.Category = dna.StringArray{}
	album.Coverart = ""
	album.Nsongs = 0
	album.Plays = 0
	album.Songids = dna.IntArray{}
	album.Description = ""
	album.Label = ""
	album.DateReleased = ""
	album.Checktime = time.Time{}
	return album
}

func getLabelFromDesc(desc dna.String) dna.String {
	var ret dna.String
	label := desc.FindAllString(`(?i)label:?.+`, 1)
	if label.Length() > 0 {
		ret = label[0].ReplaceWithRegexp(`(?i)label:?`, "").Trim()
		if ret.FindAllString(`(?mis)(.)?Publisher(\s)+:?.+`, 1).Length() > 0 {
			ret = ret.ReplaceWithRegexp(`(?mis)(.)?Publisher(\s)+:?`, "").Trim()
		}
		return ret
	}
	if desc.FindAllString(`℗.+`, 1).Length() > 0 {
		ret = desc.FindAllString(`℗.+`, 1)[0].ReplaceWithRegexp(`℗`, "").ReplaceWithRegexp(`[0-9]{4}`, "").Trim()
		return ret
	}
	if label.Length() > 0 {
		ret = label[0].ReplaceWithRegexp(`(?mis)/?PUBLISHER(\s+)?:?`, "").Trim()
		return ret
	}

	label = desc.FindAllString(`(?mis).?Publisher(\s)+:?.+`, 1)
	if label.Length() > 0 {
		ret = label[0].ReplaceWithRegexp(`(.+)?Publisher(\s)+:?`, "").Trim()
	}
	return ret
}
func getGenresFromDesc(desc dna.String) dna.StringArray {
	var ret dna.StringArray
	genres := desc.FindAllString(`(?i)genres?(\s+)?:?.+`, 1)
	// "Released:" found in album id: 836258
	if genres.Length() > 0 {
		ret = dna.StringArray(genres[0].ReplaceWithRegexp(`(?mis)genres?(\s+)?:?`, "").ReplaceWithRegexp(`\.?\s*Released:.+`, "").Trim().Split(",").Map(func(val dna.String, idx dna.Int) dna.String {
			return val.ReplaceWithRegexp(":", "").Trim()
		}).([]dna.String))
		if ret.Length() == 1 {
			arr := dna.StringArray{}
			if ret[0].FindAllString(`(?mis)K-Pop`, 1).Length() > 0 {
				arr.Push("Korean Pop")
				arr.Push(ret[0].ReplaceWithRegexp(`(?mis)\(?K-Pop\)?`, "").Trim())
				ret = arr
			}
		}
	}
	return ret.SplitWithRegexp(` > `).SplitWithRegexp(`/`)
}
func getAlbumTotalPlays(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/statistic/albumtotallisten?listIds=" + album.Id.ToString()
		result, err := http.Get(link)
		if err == nil {
			plays := result.Data.FindAllStringSubmatch(`"totalListen":"([0-9]+)"`, 1)
			if len(plays) > 0 && plays[0].Length() > 1 {
				album.Plays = plays[0][1].ToInt()
			}
		}
		channel <- true
	}()
	return channel
}
func getAlbumIssuedTime(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/album/getdescandissuetime?listIds=" + album.Id.ToString()
		// Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			dateReleased := data.FindAllStringSubmatch(`"IssueTime":"(.+?)"`, 1)
			if len(dateReleased) > 0 && dateReleased[0].Length() > 1 {
				album.DateReleased = dateReleased[0][1]
			}
		}
		channel <- true
	}()
	return channel
}
func getAlbumTotalSongs(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/album/gettotalsong?listIds=" + album.Id.ToString()
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			nsongs := data.FindAllString(`"TotalSong":"[0-9]+"`, 1)
			if nsongs.Length() > 0 {
				album.Nsongs = nsongs[0].FindAllString(`[0-9]+`, 1)[0].ToInt()
			}
		}
		channel <- true
	}()
	return channel
}
func getAlbumFromMainPage(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/nghe-album/ab." + GetKey(album.Id) + ".html"
		// Log(link)
		result, err := http.Get(link)
		if err == nil && !result.Data.Match("Rất tiếc, chúng tôi không tìm thấy thông tin bạn yêu cầu!") {
			data := &result.Data
			temp := data.FindAllString(`(?mis)class="intro_album_detail.+id="divPlayer`, 1)[0]
			if !temp.IsBlank() {
				title := temp.GetTags("strong")[0]
				if !title.IsBlank() {
					album.Title = title.RemoveHtmlTags("")
				}
				artists := temp.FindAllString(`strong.+`, 1)[0]
				if !artists.IsBlank() {
					album.Artists = artists.ReplaceWithRegexp(`^.+>`, "").ToStringArray().SplitWithRegexp(`\|\|`).SplitWithRegexp(` / `).SplitWithRegexp(" - ")
					artistid := artists.FindAllString(`\d+\.html`, 1)
					if artistid.Length() > 0 {
						album.Artistid = artistid[0].ReplaceWithRegexp(`\.html`, "").ToInt()
					}
				}

				// get multiple artists, overwrite the artists var above
				newArs := temp.FindAllString(`<p><span>.+?</span></p>`, -1)
				if newArs.Length() > 0 {
					album.Artists = dna.StringArray(newArs.Map(func(val dna.String, idx dna.Int) dna.String {
						return val.RemoveHtmlTags("").Trim()
					}).([]dna.String)).SplitWithRegexp(`\|\|`).SplitWithRegexp(` / `).SplitWithRegexp(" - ")
				}

				coverart := temp.GetTags(`img`)[0]
				if !coverart.IsBlank() {
					album.Coverart = coverart.GetTagAttributes("src")
				}
			}

			description := data.FindAllString(`<p class="desc".+?</p>`, 1)
			if description.Length() > 0 {
				album.Description = description[0].Trim().Replace("<br>", "\n").RemoveHtmlTags("")
				if album.Description.Match(`thưởng thức nhạc chất lượng cao và chia sẻ cảm xúc với bạn bè tại Nhacso.net`) {
					album.Description = ""
				}
				album.Genres = getGenresFromDesc(album.Description)
				album.Label = getLabelFromDesc(album.Description)
			}

			topics := data.FindAllString(`<li class="bg">.+</li>`, 1)[0]
			if !topics.IsBlank() {
				album.Topics = topics.RemoveHtmlTags("").ToStringArray().SplitWithRegexp(" / ").SplitWithRegexp(" - ")
			}

			songids := data.FindAllString(`songid_\d+`, -1)
			if songids.Length() > 0 {
				songids.ForEach(func(value dna.String, index dna.Int) {
					album.Songids.Push(value.ReplaceWithRegexp(`songid_`, "").ToInt())
				})
			}
		}
		channel <- true
	}()
	return channel
}

// GetAlbum returns an album and an error (if available)
func GetAlbum(id dna.Int) (*Album, error) {
	var album *Album = NewAlbum()
	album.Id = id
	c := make(chan bool)

	go func() {
		c <- <-getAlbumFromMainPage(album)
	}()
	go func() {
		c <- <-getAlbumTotalSongs(album)
	}()
	go func() {
		c <- <-getAlbumIssuedTime(album)
	}()
	go func() {
		c <- <-getAlbumTotalPlays(album)
	}()
	for i := 0; i < 4; i++ {
		<-c
	}
	if album.Nsongs != album.Songids.Length() {
		return nil, errors.New(fmt.Sprintf("Nhacso - Album %v: Songids and Nsongs do not match", album.Id))
	} else if album.Nsongs == 0 && album.Songids.Length() == 0 {
		return nil, errors.New(fmt.Sprintf("Nhacso - Album %v: No song found", album.Id))
	} else {
		album.Checktime = time.Now()
		return album, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (album *Album) Fetch() error {
	_album, err := GetAlbum(album.Id)
	if err != nil {
		return err
	} else {
		*album = *_album
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (album *Album) GetId() dna.Int {
	return album.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (album *Album) New() item.Item {
	return item.Item(NewAlbum())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (album *Album) Init(v interface{}) {
	switch v.(type) {
	case int:
		album.Id = dna.Int(v.(int))
	case dna.Int:
		album.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (album *Album) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(album)
}
