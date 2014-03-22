package ml

import (
	"dna"
	"dna/http"
	"encoding/json"
	"errors"
)

type APISong struct {
	Id         dna.Int    `json:"lyricid"`
	Title      dna.String `json:"title"`
	Artist     dna.String `json:"artist"`
	Url        dna.String `json:"url"`
	Lyric      dna.String `json:"song"`
	Status     dna.Int    `json:"content_status"`
	Gracenote  dna.Int    `json:"gracenote"`
	Publishers dna.String `json:"publishers"`
	Writers    dna.String `json:"songwriters"`
	LineCount  dna.Int    `json:"line_count"`
	// SongmeaningLines   []undefined `json:"songmeaning_lines"`
	// Songmeanings       []undefined `json:"songmeanings"`
	LineTimestamps []struct {
		Line      dna.Int `json:"line"`
		Timestamp dna.Int `json:"timestamp"`
	} `json:"songLinetimestamps"`
}

func (apiSong *APISong) FillSong(song *Song) {
	song.Id = apiSong.Id
	song.Title = apiSong.Title
	song.Artists = apiSong.Artist.ToStringArray()
	song.Url = apiSong.Url
	song.Lyric = apiSong.Lyric
	song.Status = apiSong.Status
	song.Gracenote = apiSong.Gracenote
	song.Publishers = apiSong.Publishers
	song.Writers = apiSong.Writers
	song.LineCount = apiSong.LineCount
	lt := dna.IntArray{}
	for _, ts := range apiSong.LineTimestamps {
		lt.Push(ts.Line)
		lt.Push(ts.Timestamp)

	}
	song.LineTimestamps = lt

}

func GetAPISong(id dna.Int) (*APISong, error) {
	var apisong = new(APISong)
	apisong.Id = id
	link := "http://api.metrolyrics.com/v1/get/fullbody/?lyricid=" + apisong.Id.ToString() + "&X-API-KEY=" + dna.String(X_API_KEY) + "&format=json"
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apisong)
		if apisong.Title == "" {
			return nil, errors.New(dna.Sprintf("MetroLyrics - Song ID: %v not found", apisong.Id).String())
		} else {
			return apisong, nil
		}
	} else {
		return nil, err
	}
}
