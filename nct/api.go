package nct

import (
	"dna"
	"dna/http"
	"encoding/json"
	"errors"
	"time"
)

type apiVideoInfoType struct {
	Result dna.Bool `json:"Result"`
	Data   APIVideo `json:"Data"`
}

// APIVideo defines video structure from API url.
type APIVideo struct {
	Id        dna.Int    `json:"VideoId"`
	Key       dna.String `json:"VideoKey"`
	Title     dna.String `json:"VideoTitle"`
	Thumbnail dna.String `json:"VideoThumb"`
	Image     dna.String `json:"VideoImage"`
	Artist    dna.String `json:"Singername"`
	Time      dna.String `json:"Time"`
	Artistid  dna.Int    `json:"ArtistId"`
	Likes     dna.Int    `json:"Liked"`
	Plays     dna.Int    `json:"Listened"`
	Linkshare dna.String `json:"LinkShare"`
	StreamUrl dna.String `json:"StreamURL"`
	ObjType   dna.String `json:"ObjType"`
}

func (apiVideo *APIVideo) FillVideo(video *Video) {
	video.Id = apiVideo.Id
	video.Key = apiVideo.Key
	video.Title = apiVideo.Title
	video.Image = apiVideo.Thumbnail // Image is for small image
	video.Thumbnail = apiVideo.Image // Thumbnail is for the large, kinda opposite
	// Getting dateCreated
	datecreatedArr := video.Thumbnail.FindAllStringSubmatch(`/([0-9]+)_[0-9]+\..+$`, -1)
	if len(datecreatedArr) > 0 {
		// Log(int64(datecreatedArr[0][1].ToInt()))
		video.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()/1000), 0)
	} else {
		dateCreatedArr := video.Thumbnail.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
		if len(dateCreatedArr) > 0 {
			year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
			month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
			day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
			video.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
		}
	}
	video.Artists = apiVideo.Artist.Split(",").SplitWithRegexp(" & ")
	durationArr := apiVideo.Time.Split(":")
	switch durationArr.Length() {
	case 2:
		duration, err := time.ParseDuration((durationArr[0] + "m" + durationArr[1] + "s").String())
		if err == nil {
			video.Duration = dna.Int(int(duration.Seconds()))
		} else {
			dna.Log("Critical error: cannot parse duration of video id:", video.Id)
			dna.Log("\n\n\n")
			panic("Cannot parse duration")
		}
	case 3:
		duration, err := time.ParseDuration((durationArr[0] + "h" + durationArr[1] + "m" + durationArr[2] + "s").String())
		if err == nil {
			video.Duration = dna.Int(int(duration.Seconds()))
		} else {
			dna.Log("Critical error: cannot parse duration of video id:", video.Id)
			dna.Log("\n\n\n")
			panic("Cannot parse duration")
		}
	default:
		dna.Log("Critical error: Unknown duration format of video id:", video.Id)
		dna.Log("\n\n\n")
		panic("Cannot parse duration")
	}

	video.Artistid = apiVideo.Artistid
	video.Likes = apiVideo.Likes
	video.Plays = apiVideo.Plays
	video.LinkShare = apiVideo.Linkshare
	video.StreamUrl = apiVideo.StreamUrl
	video.Type = apiVideo.ObjType
}

// GetAPIVideo returns APIVideo from video id.
// The returned APIVideo will have nil value if an error occurs.
func GetAPIVideo(id dna.Int) (*APIVideo, error) {
	urlb := NewURLBuilder()
	link := urlb.GetVideoInfo(id)
	result, err := http.Get(link)
	if err == nil {
		var apiVideoInfo = &apiVideoInfoType{}
		errd := json.Unmarshal(result.Data.ToBytes(), apiVideoInfo)
		if errd == nil {
			if apiVideoInfo.Data.Id == 0 && apiVideoInfo.Data.Key == "" {
				return nil, errors.New(dna.Sprintf("NCT - Video ID:%v not found", id).String())
			} else {
				return &apiVideoInfo.Data, nil
			}
		} else {
			if errd.Error() == "json: cannot unmarshal string into Go value of type nct.APIVideo" && result.Data.Trim() == `{"Result":false,"Data":""}` {
				return nil, errors.New(dna.Sprintf("NCT - Video ID:%v not found", id).String())
			} else {
				return nil, errd
			}
		}
	} else {
		return nil, err
	}
}

type apiSongInfoType struct {
	Result dna.Bool `json:"Result"`
	Data   APISong  `json:"Data"`
}

// APISong defines song structure from API url.
type APISong struct {
	Id         dna.Int    `json:"SongId"`
	Key        dna.String `json:"SongKey"`
	Title      dna.String `json:"SongTitle"`
	Artist     dna.String `json:"Singername"`
	Likes      dna.Int    `json:"Liked"`
	Plays      dna.Int    `json:"Listened"`
	LinkShare  dna.String `json:"LinkShare"`
	StreamUrl  dna.String `json:"StreamURL"`
	Image      dna.String `json:"Image"`
	Coverart   dna.String `json:"PlaylistThumb"`
	ObjType    dna.String `json:"ObjType"`
	Duration   dna.Int    `json:"Duration"`
	Linkdown   dna.String `json:"Linkdown"`
	LinkdownHQ dna.String `json:"LinkdownHQ"`
}

// FillSong overwrites info to Song fields which are
// similar between APISong & Song structs.
func (apiSong *APISong) FillSong(song *Song) {
	song.Id = apiSong.Id
	song.Key = apiSong.Key
	song.Title = apiSong.Title
	song.Artists = apiSong.Artist.Split(", ")
	song.Likes = apiSong.Likes
	song.Plays = apiSong.Plays
	song.LinkShare = apiSong.LinkShare
	song.StreamUrl = apiSong.StreamUrl
	song.Image = apiSong.Image
	song.Coverart = apiSong.Coverart
	song.Type = apiSong.ObjType
	if song.StreamUrl.EndsWith("mp4") == true {
		song.Type = "VIDEO"
	}
	song.Duration = apiSong.Duration
	song.Linkdown = apiSong.Linkdown
	song.LinkdownHQ = apiSong.LinkdownHQ
}

// GetAPISong returns APISong from video id.
// The returned APISong will have nil value if an error occurs.
func GetAPISong(id dna.Int) (*APISong, error) {
	urlb := NewURLBuilder()
	link := urlb.GetSongInfo(id)
	result, err := http.Get(link)
	if err == nil {
		var apiSongInfo = &apiSongInfoType{}
		// dna.Log(result.Data)
		errd := json.Unmarshal(result.Data.ToBytes(), apiSongInfo)
		if errd == nil {
			return &apiSongInfo.Data, nil
		} else {
			if errd.Error() == "json: cannot unmarshal string into Go value of type nct.APISong" && result.Data.Trim() == `{"Result":false,"Data":""}` {
				return nil, errors.New(dna.Sprintf("NCT - Song ID:%v not found", id).String())
			} else {
				return nil, errd
			}
		}
	} else {
		return nil, err
	}
}

type apiAlbumInfoType struct {
	Result dna.Bool `json:"Result"`
	Data   APIAlbum `json:"Data"`
}

// APIAlbum defines album structure from API url.
type APIAlbum struct {
	Id          dna.Int    `json:"PlaylistId"`
	Key         dna.String `json:"PlaylistKey"`
	Title       dna.String `json:"PlaylistTitle"`
	Thumbnail   dna.String `json:"PlaylistThumb"`
	Coverart    dna.String `json:"PlaylistCover"`
	Image       dna.String `json:"PlaylistImage"`
	Artist      dna.String `json:"Singername"`
	Likes       dna.Int    `json:"Liked"`
	Plays       dna.Int    `json:"Listened"`
	Linkshare   dna.String `json:"LinkShare"`
	Listsong    []APISong  `json:"ListSong"`
	Description dna.String `json:"Description"`
	Genre       dna.String `json:"Genre"`
	ObjType     dna.String `json:"ObjType"`
}

func (apiAlbum *APIAlbum) FillAlbum(album *Album) {
	var (
		nSongs, nVideos dna.Int = 0, 0
	)
	album.Id = apiAlbum.Id
	album.Key = apiAlbum.Key
	album.Title = apiAlbum.Title
	album.Artists = apiAlbum.Artist.Split(", ")
	album.Topics = apiAlbum.Genre.ToStringArray()
	album.Likes = apiAlbum.Likes
	album.Plays = apiAlbum.Plays
	ids := dna.IntArray{}
	for _, song := range apiAlbum.Listsong {
		if isSongFormat(song.StreamUrl) == true {
			nSongs += 1
		}
		if isVideoFormat(song.StreamUrl) == true {
			nVideos += 1
		}
		ids.Push(song.Id)
	}
	// dna.Log("NSONGS AND VIDEOS:", nSongs, nVideos)
	album.Songids = ids
	album.Nsongs = dna.Int(len(apiAlbum.Listsong))
	album.Description = apiAlbum.Description
	album.Coverart = apiAlbum.Coverart
	album.LinkShare = apiAlbum.Linkshare
	switch {
	case nSongs == album.Nsongs:
		album.Type = "PLAYLIST_SONG"
	case nVideos == album.Nsongs:
		album.Type = "PLAYLIST_VIDEO"
	case nVideos+nSongs == album.Nsongs:
		album.Type = "PLAYLIST_MIXED"
	default:
		album.Type = "PLAYLIST"
	}
	datecreatedArr := album.Coverart.FindAllStringSubmatch(`/([0-9]+)[_500]*\..+$`, -1)
	if len(datecreatedArr) > 0 {
		// Log(int64(datecreatedArr[0][1].ToInt()))
		album.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()/1000), 0)
	} else {
		dateCreatedArr := album.Coverart.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
		if len(dateCreatedArr) > 0 {
			year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
			month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
			day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
			album.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

		}
	}
	album.Checktime = time.Now()
}

// GetAPIAlbum returns APIAlbum from video id.
// The returned APIAlbum will have nil value if an error occurs.
func GetAPIAlbum(id dna.Int) (*APIAlbum, error) {
	urlb := NewURLBuilder()
	link := urlb.GetPlaylistInfo(id)
	result, err := http.Get(link)
	if err == nil {
		var apiAlbumInfo = &apiAlbumInfoType{}
		errd := json.Unmarshal(result.Data.ToBytes(), apiAlbumInfo)
		if errd == nil {
			return &apiAlbumInfo.Data, nil
		} else {
			if errd.Error() == "json: cannot unmarshal string into Go value of type nct.APIAlbum" && result.Data.Trim() == `{"Result":false,"Data":""}` {
				return nil, errors.New(dna.Sprintf("NCT - Album ID:%v not found", id).String())
			} else {
				return nil, errd
			}
		}
	} else {
		return nil, err
	}
}

type apiLyricInfo struct {
	Result dna.Bool `json:"Result"`
	Data   APILyric `json:"Data"`
}

// APILyric defines lyric structure from API url.
type APILyric struct {
	Lyricid         dna.Int    `json:"LyricId"`
	Lyric           dna.String `json:"Lyric"`
	Songid          dna.Int    `json:"RefId"`
	Status          dna.Int    `json:"Status"`
	TimedLyric      dna.String `json:"TimedLyric"`
	TimedLyricFile  dna.String `json:"TimedLyricFile"`
	UsernameCreated dna.String `json:"UsernameCreated"`
}

// FillSong overwrites info to Song fields which are
// similar between APILyric & Song structs.
func (apiLyric *APILyric) FillSong(song *Song) {
	song.Lyricid = apiLyric.Lyricid
	song.Lyric = apiLyric.Lyric
	song.LyricStatus = apiLyric.Status
	song.Lrc = apiLyric.TimedLyric.ReplaceWithRegexp("^\ufeff", "").Trim()
	if song.Lrc != "" {
		song.HasLrc = true
	}
	if song.Lyric != "" {
		song.HasLyric = true
	}
	song.LrcUrl = apiLyric.TimedLyricFile
	song.UsernameCreated = apiLyric.UsernameCreated
}

// GetAPILyric returns APILyric from video id.
// The returned APILyric will have nil value if an error occurs.
func GetAPILyric(songid dna.Int) (*APILyric, error) {
	urlb := NewURLBuilder()
	link := urlb.GetSongLyric(songid)
	result, err := http.Get(link)
	if err == nil {
		var apiSongLrcInfo = &apiLyricInfo{}
		errd := json.Unmarshal(result.Data.ToBytes(), apiSongLrcInfo)
		if errd == nil {
			if apiSongLrcInfo.Result == false {
				return nil, errors.New(dna.Sprintf("NCT - Song ID:%v Lyric not found", songid).String())
			} else {
				return &apiSongLrcInfo.Data, nil
			}
		} else {
			return nil, errd
		}
	} else {
		return nil, err
	}
}

type apiArtistInfo struct {
	Result dna.Bool  `json:"Result"`
	Data   APIArtist `json:"Data"`
}

// APIArtist defines an artist structure from API url.
type APIArtist struct {
	Id      dna.Int    `json:"ArtistId"`
	Name    dna.String `json:"ArtistName"`
	Avatar  dna.String `json:"ArtistAvatar"`
	NSongs  dna.Int    `json:"SongCount"`
	NAlbums dna.Int    `json:"PlaylistCount"`
	NVideos dna.Int    `json:"VideoCount"`
	ObjType dna.String `json:"ObjType"`
}

func (apiArtist *APIArtist) CSVRecord() []string {
	return []string{
		apiArtist.Id.ToString().String(),
		apiArtist.Name.String(),
		apiArtist.Avatar.String(),
		apiArtist.NSongs.ToString().String(),
		apiArtist.NAlbums.ToString().String(),
		apiArtist.NVideos.ToString().String(),
		apiArtist.ObjType.String(),
	}
}

func (apiArtist *APIArtist) ToArtist() *Artist {
	artist := NewArtist()
	artist.Id = apiArtist.Id
	artist.Name = apiArtist.Name
	artist.Avatar = apiArtist.Avatar
	artist.NSongs = apiArtist.NSongs
	artist.NAlbums = apiArtist.NAlbums
	artist.NVideos = apiArtist.NVideos
	return artist
}

// GetAPIArtist returns APIArtsit from video id.
// The returned APIArtsit will have nil value if an error occurs.
func GetAPIArtist(id dna.Int) (*APIArtist, error) {
	urlb := NewURLBuilder()
	link := urlb.GetArtistInfo(id)
	result, err := http.Get(link)
	if err == nil {
		var aArtistInfo = &apiArtistInfo{}
		errd := json.Unmarshal(result.Data.ToBytes(), aArtistInfo)
		if errd == nil {
			if aArtistInfo.Result == false || aArtistInfo.Data.Id == 0 {
				return nil, errors.New(dna.Sprintf("NCT - Artist ID:%v not found", id).String())
			} else {
				return &aArtistInfo.Data, nil
			}
		} else {
			return nil, errd
		}
	} else {
		return nil, err
	}

}

type GenreList struct {
	Result dna.Bool `json:"Result"`
	Ismore dna.Bool `json:"IsMore"`
	Data   []struct {
		Id      dna.String `json:"GenreId"`
		Name    dna.String `json:"GenreName"`
		ObjType dna.String `json:"ObjType"`
	} `json:"Data"`
}
