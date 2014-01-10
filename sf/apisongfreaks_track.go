package sf

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"dna/terminal"
	"dna/utils"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"time"
)

var (
	VideosEnable   = false
	CommentsEnable = false
	SQLERROR       = terminal.NewLogger(terminal.Magenta, ioutil.Discard, "", "./log/sql_error.log", 0)
)

type SongFreak struct {
	Id        dna.Int
	Track     dna.String
	Videos    dna.String
	Checktime time.Time
}

func NewSongFreak() *SongFreak {
	sf := new(SongFreak)
	sf.Id = 0
	sf.Track = "{}"
	sf.Videos = "[]"
	return sf
}

type APISongFreaksTrack struct {
	Id       dna.Int
	XMLName  xml.Name    `xml:"songfreaks"`
	Response APIResponse `xml:"response"`
	Track    APITrack    `xml:"track"`
	Videos   []APIVideo  `xml:"videos>video"`
	Comments APIComments `xml:"comments"`
}

func NewAPISongFreaksTrack() *APISongFreaksTrack {
	sf := new(APISongFreaksTrack)
	sf.Id = 0
	sf.Response = APIResponse{202, 0, ""}
	sf.Track = APITrack{}
	sf.Videos = []APIVideo{}
	sf.Comments = APIComments{}
	return sf
}

func GetSongFreaksTrack(id dna.Int) (*APISongFreaksTrack, error) {
	var link dna.String = "http://apiv2.songfreaks.com//lyric.do?"
	// Log(link)
	PostData.SetIdKey(id)
	if VideosEnable == true {
		PostData.VideosEnable = true
	}
	if CommentsEnable == true {
		PostData.CommentsEnable = true
	}
	result, err := Post(link, PostData.Encode())
	mutex.Lock()
	Cookie = result.Header.Get("Set-Cookie")
	mutex.Unlock()
	if err == nil {
		songfreaks := &APISongFreaksTrack{}
		songfreaks.Id = id
		merr := xml.Unmarshal([]byte(result.Data.String()), songfreaks)
		if merr == nil {
			// dna.Log(result.Data)
			if sferr := songfreaks.HasError(); sferr != nil {
				return nil, sferr
			} else {
				return songfreaks, nil
			}
		} else {
			return nil, merr
		}
	}
	return nil, err
}

func (sf *APISongFreaksTrack) HasError() error {
	switch sf.Response.Code {
	case 101, 102, 103, 111:
		// do nothing, request successful
		// code 103 has SUCCESS: LICENSE, NO LYRICS
	case 202, 201:
		return errors.New("No id found at index:" + sf.Id.ToString().String())
	default:
		mes := dna.Sprintf("%v", sf.Response).String()
		return errors.New("Unknow response code" + mes)
	}
	if sf.Track.Id == 0 {
		return errors.New("No track id found at index" + sf.Id.ToString().String())
	}
	return nil
}

func (sf *APISongFreaksTrack) ToSongFreak() (*SongFreak, error) {
	if sferr := sf.HasError(); sferr != nil {
		return nil, sferr
	}
	sofre := NewSongFreak()
	sofre.Id = sf.Id
	track, err := xml.MarshalIndent(sf.Track, "", "\t")
	if err == nil {
		sofre.Track = dna.String(string(track))
	}
	videos, err := xml.MarshalIndent(sf.Videos, "", "\t")
	if err == nil {
		sofre.Videos = dna.String(string(videos))
	}
	sofre.Checktime = time.Now()
	return sofre, nil
}

func (sf *APISongFreaksTrack) ToSong() (*Song, error) {
	if sferr := sf.HasError(); sferr != nil {
		return nil, sferr
	}
	song := NewSong()
	song.Id = sf.Track.Id
	song.TrackGroupId = sf.Track.TrackGroupId
	song.AMG = sf.Track.AMG
	song.UrlSlug = sf.Track.UrlSlug
	song.IsInstrumental = sf.Track.IsInstrumental
	song.Viewable = sf.Track.Viewable
	song.Duration = utils.ToSeconds(sf.Track.Duration)
	song.Lyricid = sf.Track.LyricId
	song.HasLrc = sf.Track.HasLrc
	song.TrackNumber = sf.Track.TrackNumber
	song.DiscNumber = sf.Track.DiscNumber
	song.Title = sf.Track.Title
	song.Rating = dna.IntArray{sf.Track.Rating.AverageRating.ToInt(), sf.Track.Rating.UserRating, sf.Track.Rating.TotalRatings}
	song.Albumid = sf.Track.Album.Id

	artistIds := dna.IntArray{}
	artists := dna.StringArray{}
	for _, artist := range sf.Track.Artists {
		artistIds.Push(artist.Id)
		artists.Push(artist.Name)
	}
	song.Artistids = artistIds
	song.Artists = artists

	if sf.Track.Lrc.Lines != nil && len(sf.Track.Lrc.Lines) > 0 {
		lines, err := json.Marshal(sf.Track.Lrc.Lines)
		if err == nil {
			song.Lrc = dna.String(string(lines))

		}
	}

	song.Link = sf.Track.Link
	song.Lyric = sf.Track.Lyrics
	if song.Lyric != "" {
		song.HasLyric = true
	}
	song.Copyright = sf.Track.Copyright
	song.Writer = sf.Track.Writer
	song.SubmittedLyric = sf.Track.SubmittedLyric
	song.Checktime = time.Now()
	return song, nil
}
func (sf *APISongFreaksTrack) ToAlbum() (*Album, error) {
	if sferr := sf.HasError(); sferr != nil {
		return nil, sferr
	}
	album := NewAlbum()
	album.Id = sf.Track.Album.Id
	album.AMG = sf.Track.Album.AMG
	album.UrlSlug = sf.Track.Album.UrlSlug
	album.Year = sf.Track.Album.Year
	album.Coverart = sf.Track.Album.Coverart
	album.CoverartLarge = sf.Track.Album.CoverartLarge
	album.Title = sf.Track.Album.Title
	// album.Ratings = dna.IntArray{0, 0, 0}
	// NOT IMPLEMENTING
	album.Artistid = sf.Track.Album.Artist.Id
	album.Artists = sf.Track.Album.Artist.Name.ToStringArray()
	album.Link = sf.Track.Album.Link
	album.Songids = dna.IntArray{}
	// album.ReviewAuthor = ""
	// album.Review = ""
	// NOT IMPLEMENTING
	album.Checktime = time.Now()
	return album, nil
}

func (sf *APISongFreaksTrack) ToAlbumArtist() (*Artist, error) {
	if sferr := sf.HasError(); sferr != nil {
		return nil, sferr
	}
	artist := NewArtist()
	albumArtist := sf.Track.Album.Artist
	artist.Id = albumArtist.Id
	artist.AMG = albumArtist.AMG
	artist.UrlSlug = albumArtist.UrlSlug
	artist.Image = albumArtist.Image
	artist.Genres = albumArtist.Genres.ToStringArray()
	artist.Name = albumArtist.Name
	artist.Rating = dna.IntArray{0, 0, 0}
	// No Rating found in an artist of an album of a track
	artist.Bio = "{}"
	// No BIO found in an artist of an album of a track
	artist.Checktime = time.Now()
	if albumArtist.Id > 0 {
		return artist, nil
	} else {
		return nil, errors.New("Cannot convert artist")
	}

}

func (sf *APISongFreaksTrack) ToVideos() ([]*Video, error) {
	if sferr := sf.HasError(); sferr != nil {
		return nil, sferr
	}
	videos := []*Video{}
	for _, tvideo := range sf.Videos {
		video := NewVideo()
		video.Songid = sf.Track.Id
		video.YoutubeId = tvideo.YoutubeId
		video.Duration = tvideo.Duration
		video.Thumbnail = tvideo.Thumbnail
		video.Title = tvideo.Title
		video.Description = tvideo.Description
		video.Checktime = time.Now()
		videos = append(videos, video)
	}
	return videos, nil
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (sf *APISongFreaksTrack) Fetch() error {
	_sf, err := GetSongFreaksTrack(sf.Id)
	if err != nil {
		return err
	} else {
		*sf = *_sf
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (sf *APISongFreaksTrack) GetId() dna.Int {
	return sf.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (sf *APISongFreaksTrack) New() item.Item {
	return item.Item(NewAPISongFreaksTrack())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (sf *APISongFreaksTrack) Init(v interface{}) {
	switch v.(type) {
	case int:
		sf.Id = dna.Int(v.(int))
	case dna.Int:
		sf.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (sf *APISongFreaksTrack) Save(db *sqlpg.DB) error {

	sofre, err := sf.ToSongFreak()
	if err == nil {
		return db.InsertIgnore(sofre)
	} else {
		return err
	}

	// var rerr error = nil

	// err := db.Ping()

	// artist, err := sf.ToAlbumArtist()
	// if err == nil {
	// 	rerr = artist.Save(db)
	// 	if rerr != nil {
	// 		SQLERROR.Println("SF-ERROR: " + rerr.Error())
	// 	}
	// }

	// album, err := sf.ToAlbum()
	// if err == nil {
	// 	rerr = album.Save(db)
	// 	if rerr != nil {
	// 		SQLERROR.Println("SF-ERROR: " + rerr.Error())
	// 	}
	// }

	// song, err := sf.ToSong()
	// if err == nil {
	// 	rerr = song.Save(db)
	// 	if rerr != nil {
	// 		SQLERROR.Println("SF-ERROR: " + rerr.Error())
	// 	}
	// }

	// videos, err := sf.ToVideos()
	// if err == nil {
	// 	for _, video := range videos {
	// 		rerr = video.Save(db)
	// 		if rerr != nil {
	// 			SQLERROR.Println("SF-ERROR: " + rerr.Error())
	// 		}
	// 	}
	// }
	// return nil
	// return db.InsertIgnore(sf)
}
