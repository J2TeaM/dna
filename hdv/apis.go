package hdv

import (
	"dna"
	"dna/http"
	"encoding/json"
)

type APIMovieJSON struct {
	Error dna.Int  `json:"e"`
	Movie APIMovie `json:"r"`
}

type APIMovie struct {
	MovieId        dna.Int
	EpId           dna.Int
	Title          dna.String      `json:"MovieName"`
	LinkPlayBackup dna.String      `json:"LinkPlayBackup"`
	Link           dna.String      `json:"LinkPlay"`
	LinkPlayOther  dna.String      `json:"LinkPlayOther"`
	SubtitleExt    APISubtitleList `json:"SubtitleExt"`
	SubtitleExtSe  APISubtitleList `json:"SubtitleExtSe"`
	Subtitle       APISubtitleList `json:"Subtitle"`
	Episode        dna.String      `json:"Episode"`
	Audiodub       dna.String      `json:"AudioDub"`
	Audio          dna.String      `json:"Audio"`
	Season         []Season        `json:"Season"`
	Adver          dna.StringArray `json:"Adver"`
}

func (apiMovie *APIMovie) ToEpisode() *Episode {
	episode := NewEpisode()
	episode.MovieId = apiMovie.MovieId
	episode.EpId = apiMovie.EpId
	episode.Title = apiMovie.Title
	episode.LinkPlayBackup = apiMovie.LinkPlayBackup
	episode.Link = apiMovie.Link
	episode.LinkPlayOther = apiMovie.LinkPlayOther
	episode.SubtitleExt = dna.StringArray{apiMovie.SubtitleExt.Vietnamese.Source, apiMovie.SubtitleExt.English.Source}
	episode.SubtitleExtSe = dna.StringArray{apiMovie.SubtitleExtSe.Vietnamese.Source, apiMovie.SubtitleExtSe.English.Source}
	episode.Subtitle = dna.StringArray{apiMovie.Subtitle.Vietnamese.Source, apiMovie.Subtitle.English.Source}
	episode.EpisodeId = apiMovie.Episode.ToInt()
	episode.Audiodub = apiMovie.Audiodub.ToInt()
	episode.Audio = apiMovie.Audio.ToInt()
	season, err := json.Marshal(apiMovie.Season)
	if err == nil {
		episode.Season = dna.String(string(season))
	}
	return episode
}

func GetAPIMovie(movieid, ep dna.Int) (*APIMovie, error) {
	urlb := NewURLBuilder()
	var link dna.String = ""
	if ep == 0 {
		link = urlb.GetMovie(movieid)
	} else {
		link = urlb.GetEpisole(movieid, ep)
	}
	result, err := http.Get(link)
	if err == nil {
		var apiMoveiJSON = &APIMovieJSON{}
		errd := json.Unmarshal(result.Data.ToBytes(), apiMoveiJSON)
		if errd == nil {
			apiMoveiJSON.Movie.MovieId = movieid
			apiMoveiJSON.Movie.EpId = ep
			return &apiMoveiJSON.Movie, nil
		} else {
			return nil, errd
		}
	} else {
		return nil, err
	}
}

type APISubtitle struct {
	Label  dna.String `json:"Label"`
	Source dna.String `json:"Source"`
}

type APISubtitleList struct {
	Vietnamese APISubtitle `json:"VIE"`
	English    APISubtitle `json:"ENG"`
}

type APIChannelJSON struct {
	Error   dna.Int    `json:"e"`
	Channel APIChannel `json:"r"`
}

type APIChannel struct {
	Name          dna.String `json:"ChannelName"`
	Image         dna.String `json:"ChannelImage"`
	Link          dna.String `json:"LinkPlay"`
	LinkPlayOther dna.String `json:"LinkPlayOther"`
}

type Season struct {
	Id    dna.String `json:"MovieID"`
	Title dna.String `json:"Name"`
}
