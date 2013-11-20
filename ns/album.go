package ns

import (
	. "dna"
	"dna/http"
	"errors"
	"fmt"
	"time"
)

// Define new Album type.
// Notice: Artistid should be Artistids , but this field is not important, then it will be ignored.
type Album struct {
	Id           Int
	Title        String
	Artists      StringArray
	Artistid     Int
	Topics       StringArray
	Genres       StringArray
	Category     StringArray
	Coverart     String
	Nsongs       Int
	Plays        Int
	Songids      IntArray
	Description  String
	Label        String
	DateReleased String
	Checktime    time.Time
}

// NewAlbum return default new album
func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Title = ""
	album.Artists = StringArray{}
	album.Artistid = 0
	album.Topics = StringArray{}
	album.Genres = StringArray{}
	album.Category = StringArray{}
	album.Coverart = ""
	album.Nsongs = 0
	album.Plays = 0
	album.Songids = IntArray{}
	album.Description = ""
	album.Label = ""
	album.DateReleased = ""
	album.Checktime = time.Time{}
	return album
}

func getLabelFromDesc(desc String) String {
	var ret String
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
func getGenresFromDesc(desc String) StringArray {
	var ret StringArray
	genres := desc.FindAllString(`(?i)genres?(\s+)?:?.+`, 1)
	if genres.Length() > 0 {
		ret = StringArray(genres[0].ReplaceWithRegexp(`(?mis)genres?(\s+)?:?`, "").Trim().Split(",").Map(func(val String, idx Int) String {
			return val.ReplaceWithRegexp(":", "").Trim()
		}).([]String))
		if ret.Length() == 1 {
			arr := StringArray{}
			if ret[0].FindAllString(`(?mis)K-Pop`, 1).Length() > 0 {
				arr.Push("Korean Pop")
				arr.Push(ret[0].ReplaceWithRegexp(`(?mis)\(?K-Pop\)?`, "").Trim())
				ret = arr
			}
		}
	}
	return ret.SplitWithRegexp(` > `)
}
func getAlbumTotalPlays(id Int) <-chan *Album {
	channel := make(chan *Album, 1)
	go func() {
		album := NewAlbum()
		album.Id = id
		link := "http://nhacso.net/statistic/albumtotallisten?listIds=" + id.ToString()
		result, err := http.Get(link)
		if err != nil {
			channel <- album
		} else {
			plays := result.Data.FindAllStringSubmatch(`"totalListen":"([0-9]+)"`, 1)
			if len(plays) > 0 && plays[0].Length() > 1 {
				album.Plays = plays[0][1].ToInt()
			}

			channel <- album
		}
	}()
	return channel
}
func getAlbumIssuedTime(id Int) <-chan *Album {
	channel := make(chan *Album, 1)
	go func() {
		album := NewAlbum()
		album.Id = id
		link := "http://nhacso.net/album/getdescandissuetime?listIds=" + id.ToString()
		// Log(link)
		result, err := http.Get(link)
		if err != nil {
			channel <- album
		} else {
			data := &result.Data
			dateReleased := data.FindAllStringSubmatch(`"IssueTime":"(.+?)"`, 1)
			if len(dateReleased) > 0 && dateReleased[0].Length() > 1 {
				album.DateReleased = dateReleased[0][1]
			}
			channel <- album
		}
	}()
	return channel
}
func getAlbumTotalSongs(id Int) <-chan *Album {
	channel := make(chan *Album, 1)
	go func() {
		album := NewAlbum()
		album.Id = id
		link := "http://nhacso.net/album/gettotalsong?listIds=" + id.ToString()
		result, err := http.Get(link)
		if err != nil {
			channel <- album
		} else {
			data := &result.Data
			nsongs := data.FindAllString(`"TotalSong":"[0-9]+"`, 1)
			if nsongs.Length() > 0 {
				album.Nsongs = nsongs[0].FindAllString(`[0-9]+`, 1)[0].ToInt()
			}
			channel <- album
		}
	}()
	return channel
}
func getAlbumFromMainPage(id Int) <-chan *Album {
	channel := make(chan *Album, 1)
	go func() {
		album := NewAlbum()
		album.Id = id
		link := "http://nhacso.net/nghe-album/ab." + GetKey(id) + ".html"
		// Log(link)
		result, err := http.Get(link)
		if err != nil || result.Data.Match("Rất tiếc, chúng tôi không tìm thấy thông tin bạn yêu cầu!") {
			channel <- album
		} else {
			data := &result.Data
			temp := data.FindAllString(`(?mis)class="intro_album_detail.+id="divPlayer`, 1)[0]
			if !temp.IsBlank() {
				title := temp.GetTags("strong")[0]
				if !title.IsBlank() {
					album.Title = title.RemoveHtmlTags("")
				}
				artists := temp.FindAllString(`strong.+`, 1)[0]
				if !artists.IsBlank() {
					album.Artists = artists.ReplaceWithRegexp(`^.+>`, "").ToStringArray().SplitWithRegexp(`\|\|`).SplitWithRegexp(" / ").SplitWithRegexp(" - ")
					artistid := artists.FindAllString(`\d+\.html`, 1)
					if artistid.Length() > 0 {
						album.Artistid = artistid[0].ReplaceWithRegexp(`\.html`, "").ToInt()
					}
				}

				// get multiple artists, overwrite the artists var above
				newArs := temp.FindAllString(`<p><span>.+?</span></p>`, -1)
				if newArs.Length() > 0 {
					album.Artists = StringArray(newArs.Map(func(val String, idx Int) String {
						return val.RemoveHtmlTags("").Trim()
					}).([]String))
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
				songids.ForEach(func(value String, index Int) {
					album.Songids.Push(value.ReplaceWithRegexp(`songid_`, "").ToInt())
				})
			}
			channel <- album
		}
	}()
	return channel
}

// GetAlbum returns an album and an error (if available)
func GetAlbum(id Int) (*Album, error) {
	var album *Album = NewAlbum()
	c := make(chan *Album)

	go func() {
		c <- <-getAlbumFromMainPage(id)
	}()
	go func() {
		c <- <-getAlbumTotalSongs(id)
	}()
	go func() {
		c <- <-getAlbumIssuedTime(id)
	}()
	go func() {
		c <- <-getAlbumTotalPlays(id)
	}()
	for i := 0; i < 4; i++ {
		tempAlbum := <-c
		if tempAlbum.Id > 0 {
			album.Id = tempAlbum.Id
		}
		if tempAlbum.Title != "" {
			album.Title = tempAlbum.Title
		}
		if tempAlbum.Artists.Length() > 0 {
			album.Artists = tempAlbum.Artists
		}
		if tempAlbum.Artistid > 0 {
			album.Artistid = tempAlbum.Artistid
		}
		if tempAlbum.Topics.Length() > 0 {
			album.Topics = tempAlbum.Topics
		}
		if tempAlbum.Genres.Length() > 0 {
			album.Genres = tempAlbum.Genres
		}
		if tempAlbum.Coverart != "" {
			album.Coverart = tempAlbum.Coverart
		}
		if tempAlbum.Category.Length() > 0 {
			album.Category = tempAlbum.Category
		}
		if tempAlbum.Nsongs > 0 {
			album.Nsongs = tempAlbum.Nsongs
		}
		if tempAlbum.Plays > 0 {
			album.Plays = tempAlbum.Plays
		}
		if tempAlbum.Songids.Length() > 0 {
			album.Songids = tempAlbum.Songids
		}
		if tempAlbum.Description != "" {
			album.Description = tempAlbum.Description
		}
		if tempAlbum.Label != "" {
			album.Label = tempAlbum.Label
		}
		if tempAlbum.DateReleased != "" {
			album.DateReleased = tempAlbum.DateReleased
		}

	}
	if album.Nsongs != album.Songids.Length() {
		return nil, errors.New(fmt.Sprintf("Album %v: Songids and Nsongs do not match", album.Id))
	} else if album.Nsongs == 0 {
		return nil, errors.New(fmt.Sprintf("Album %v: No song found", album.Id))
	} else {
		album.Checktime = time.Now()
		return album, nil
	}
}
