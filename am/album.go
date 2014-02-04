package am

import (
	"dna"
	"encoding/json"
	"time"
)

func convertCategoryToStringArray(catStr dna.String) dna.StringArray {
	var cats []Category
	var ret = dna.StringArray{}
	err := json.Unmarshal([]byte(string(catStr)), &cats)
	if err != nil {
		dna.Log(catStr)
		dna.Log(err.Error())
		panic("Invalid category string input")
	} else {
		for _, cat := range cats {
			ret.Push(cat.Name)
		}
	}
	return ret
}

func convertSongToIntArray(songStr dna.String) dna.IntArray {
	var apiSongs []APISong
	var ret = dna.IntArray{}
	err := json.Unmarshal([]byte(string(songStr)), &apiSongs)
	if err != nil {
		panic("Invalid song string input")
	} else {
		for _, song := range apiSongs {
			ret.Push(song.Id)
		}
	}
	return ret
}

type Album struct {
	Id           dna.Int
	Title        dna.String
	Artistids    dna.IntArray
	Artists      dna.StringArray
	Review       dna.String
	Coverart     dna.String
	Duration     dna.Int
	Ratings      dna.IntArray
	Similars     dna.IntArray
	Genres       dna.StringArray
	Styles       dna.StringArray
	Moods        dna.StringArray
	Themes       dna.StringArray
	Songids      dna.IntArray
	DateReleased time.Time
	Checktime    time.Time
}

func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Title = ""
	album.Artistids = dna.IntArray{}
	album.Artists = dna.StringArray{}
	album.Review = ""
	album.Coverart = ""
	album.Duration = 0
	album.Ratings = dna.IntArray{0, 0, 0}
	album.Similars = dna.IntArray{}
	album.Genres = dna.StringArray{}
	album.Styles = dna.StringArray{}
	album.Moods = dna.StringArray{}
	album.Themes = dna.StringArray{}
	album.Songids = dna.IntArray{}
	album.DateReleased = time.Time{}
	album.Checktime = time.Time{}
	return album
}

func (album *Album) CSVRecord() []string {
	return []string{
		album.Id.ToString().String(),
		album.Title.String(),
		dna.Sprintf("%#v", album.Artistids).Replace("dna.IntArray", "").String(),
		dna.Sprintf("%#v", album.Artists).Replace("dna.StringArray", "").String(),
		album.Review.String(),
		album.Coverart.String(),
		album.Duration.ToString().String(),
		dna.Sprintf("%#v", album.Ratings).Replace("dna.IntArray", "").String(),
		dna.Sprintf("%#v", album.Similars).Replace("dna.IntArray", "").String(),
		dna.Sprintf("%#v", album.Genres).Replace("dna.StringArray", "").String(),
		dna.Sprintf("%#v", album.Styles).Replace("dna.StringArray", "").String(),
		dna.Sprintf("%#v", album.Moods).Replace("dna.StringArray", "").String(),
		dna.Sprintf("%#v", album.Themes).Replace("dna.StringArray", "").String(),
		dna.Sprintf("%#v", album.Songids).Replace("dna.IntArray", "").String(),
		album.DateReleased.Format("2006-01-02 15:04:05"),
		album.Checktime.Format("2006-01-02 15:04:05"),
	}
}
