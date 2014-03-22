package lf

import (
	"dna"
	"dna/http"
	"encoding/json"
	"errors"
	"time"
)

type APISong struct {
	Response struct {
		Code        dna.Int    `json:"code"`
		Description dna.String `json:"description"`
	} `json:"response"`
	Track struct {
		Amg          dna.Int    `json:"amg"`
		Instrumental dna.Bool   `json:"instrumental"`
		Viewable     dna.Bool   `json:"viewable"`
		HasLrc       dna.Bool   `json:"has_lrc"`
		LrcVerified  dna.Bool   `json:"lrc_verified"`
		Title        dna.String `json:"title"`
		Artist       struct {
			Name dna.String `json:"name"`
		} `json:"artist"`
		LastUpdate dna.String `json:"last_update"`
		Lyrics     dna.String `json:"lyrics"`
		Copyright  dna.String `json:"copyright"`
		Writer     dna.String `json:"writer"`
	} `json:"track"`
}

func (apiSong *APISong) FillSong(song *Song) {
	song.Id = apiSong.Track.Amg
	song.Title = apiSong.Track.Title
	song.Artists = apiSong.Track.Artist.Name.Split(" and ").SplitWithRegexp(", ")
	song.Instrumental = apiSong.Track.Instrumental
	song.Viewable = apiSong.Track.Viewable
	song.HasLrc = apiSong.Track.HasLrc
	song.LrcVerified = apiSong.Track.LrcVerified
	song.Lyric = apiSong.Track.Lyrics
	song.Copyright = apiSong.Track.Copyright
	song.Writers = apiSong.Track.Writer
	// Mon Jan 2 15:04:05 MST 2006
	if apiSong.Track.LastUpdate != "" {
		lastUpdate, err := time.Parse("2006-01-02 15:04:05", apiSong.Track.LastUpdate.String())
		if err == nil {
			song.DateUpdated = lastUpdate
		} else {
			dna.Log(err.Error(), " Song ID:", song.Id, " GOT:", apiSong.Track.LastUpdate, "\n\n")
		}
	}
}

func GetAPISong(id dna.Int) (*APISong, error) {
	var apisong = new(APISong)
	link := "https://api.lyricfind.com/lyric.do?apikey=" + API_KEY + "&reqtype=default&trackid=amg:" + id.ToString().String() + "&output=json"
	// result, err := GetProxy(dna.String(link), 0)
	result, err := http.Get(dna.String(link))
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apisong)
		switch apisong.Response.Code {
		case 100, 101, 111, 102:
			return apisong, nil
		default:
			return nil, errors.New(dna.Sprintf("LyricFind - Song ID: %v - %v %v", id, apisong.Response.Code, apisong.Response.Description).String())
		}
		return apisong, nil
	} else {
		return nil, err
	}
}
