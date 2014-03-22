package lf

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"dna/utils"
	"encoding/json"
	"errors"
	"time"
)

// http://api.lyricfind.com/metadata.do?apikey=4b59d60b5b74512a662b89dfb1b28680&reqtype=availablelyrics&artistid=amg:7362&offset=0&limit=10&displaykey=ccabb2c8bf7302e1d8c9b87be793bfb0&output=json
// METADATA

type APIArtist struct {
	Id         dna.Int    `json:"amg"`
	Image      dna.String `json:"image"`
	Genre      dna.String `json:"genre"`
	Popularity dna.Int    `json:"popularity"`
	Name       dna.String `json:"name"`
}

type APIAlbum struct {
	Id         dna.Int    `json:"amg"`
	Year       dna.String `json:"year"`
	Image      dna.String `json:"image"`
	Largeimage dna.String `json:"largeimage"`
	Importance dna.Int    `json:"importance"`
	Title      dna.String `json:"title"`
	Artist     APIArtist  `json:"artist"`
}

type APIResponse struct {
	Code        dna.Int    `json:"code"`
	Description dna.String `json:"description"`
}

type APIFullSong struct {
	Response APIResponse `json:"response"`
	Track    struct {
		Id           dna.Int    `json:"amg"`
		Title        dna.String `json:"title"`
		Instrumental dna.Bool   `json:"instrumental"`
		Viewable     dna.Bool   `json:"viewable"`
		HasLrc       dna.Bool   `json:"has_lrc"`
		LrcVerified  dna.Bool   `json:"lrc_verified"`

		Duration   dna.String  `json:"duration"`
		Lyricid    dna.Int     `json:"lyric"`
		Album      APIAlbum    `json:"album"`
		Artists    []APIArtist `json:"artists"`
		LastUpdate dna.String  `json:"last_update"`
		Lyrics     dna.String  `json:"lyrics"`
		Copyright  dna.String  `json:"copyright"`
		Writer     dna.String  `json:"writer"`
	} `json:"track"`
}

func NewAPIFullSong() *APIFullSong {
	return new(APIFullSong)
}

func (apiSong *APIFullSong) ToArtists() []*Artist {
	artists := []*Artist{}

	// Getting artist from album
	albumArtist := NewArtist()
	albumArtist.Id = apiSong.Track.Album.Artist.Id
	albumArtist.Name = apiSong.Track.Album.Artist.Name
	albumArtist.Genres = apiSong.Track.Album.Artist.Genre.ToStringArray()
	albumArtist.Popularity = apiSong.Track.Album.Artist.Popularity
	albumArtist.Image = apiSong.Track.Album.Artist.Image
	artists = append(artists, albumArtist)

	// Getting artists from song
	for _, artist := range apiSong.Track.Artists {
		// dna.Log(artist)
		// Only get new artists
		if artist.Id != albumArtist.Id {
			trackArtist := NewArtist()
			trackArtist.Id = artist.Id
			trackArtist.Name = artist.Name
			trackArtist.Genres = artist.Genre.ToStringArray()
			trackArtist.Popularity = artist.Popularity
			trackArtist.Image = artist.Image
			artists = append(artists, trackArtist)
		}
	}
	return artists
}

func (apiSong *APIFullSong) ToSong() *Song {
	song := NewSong()
	apiSong.FillSong(song)
	return song
}

func (apiSong *APIFullSong) ToAlbum() *Album {
	// dna.Log(apiSong.Track.Album)
	album := NewAlbum()
	album.Id = apiSong.Track.Album.Id
	album.Title = apiSong.Track.Album.Title
	album.Artistid = apiSong.Track.Album.Artist.Id
	album.Artists = apiSong.Track.Album.Artist.Name.ToStringArray()
	album.Year = apiSong.Track.Album.Year
	album.Importance = apiSong.Track.Album.Importance
	album.Image = apiSong.Track.Album.Image
	album.LargeImage = apiSong.Track.Album.Largeimage
	return album
}

func (apiSong *APIFullSong) FillSong(song *Song) {
	song.Id = apiSong.Track.Id
	song.Title = apiSong.Track.Title

	artists := dna.StringArray{}
	for _, artist := range apiSong.Track.Artists {
		artists.Push(artist.Name)
	}
	song.Artists = artists
	artistids := dna.IntArray{}
	for _, artist := range apiSong.Track.Artists {
		artistids.Push(artist.Id)
	}

	song.Artistids = artistids
	song.Albumid = apiSong.Track.Album.Id
	song.AlbumTitle = apiSong.Track.Album.Title
	song.Duration = utils.ToSeconds(apiSong.Track.Duration)
	song.Instrumental = apiSong.Track.Instrumental
	song.Viewable = apiSong.Track.Viewable
	song.HasLrc = apiSong.Track.HasLrc
	song.LrcVerified = apiSong.Track.LrcVerified
	song.Lyricid = apiSong.Track.Lyricid
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

func getAPIFullSongWithLyric(id dna.Int) (*APIFullSong, error) {
	var apisong = new(APIFullSong)
	var link string
	link = "https://api.lyricfind.com/lyric.do?apikey=" + API_KEY + "&reqtype=default&trackid=amg:" + id.ToString().String() + "&output=json"
	result, err := http.Get(dna.String(link))
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apisong)
		switch apisong.Response.Code {
		case 100, 101, 111, 102:
			return apisong, nil
		case 203, 204, 205: // "description": "LYRIC NOT AVAILABLE"
			return getAPIFullSongWithMetadata(id)
		default:
			return nil, errors.New(dna.Sprintf("LyricFind - Song ID: %v - %v %v", id, apisong.Response.Code, apisong.Response.Description).String())

		}
		return apisong, nil
	} else {
		return nil, err
	}
}

func getAPIFullSongWithMetadata(id dna.Int) (*APIFullSong, error) {
	var apisong = new(APIFullSong)
	var link string

	link = "http://api.lyricfind.com/metadata.do?apikey=" + METADATA_KEY + "&reqtype=metadata&trackid=amg:" + id.ToString().String() + "&displaykey=" + API_KEY + "&output=json"
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

// GetAPIFullSong runs with 2 steps the following procedure:
// 	It checks whether a given songid has lyric and metadata.
// 	If the lyric is not available but the metadata is available,
// 	then it fetches metadata.
// 	An error is returned when no info is found.
func GetAPIFullSong(id dna.Int) (*APIFullSong, error) {
	return getAPIFullSongWithLyric(id)
}

func (apiSong *APIFullSong) Fetch() error {
	_apiSong, err := GetAPIFullSong(apiSong.Track.Id)
	if err != nil {
		return err
	} else {
		*apiSong = *_apiSong
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (apiSong *APIFullSong) GetId() dna.Int {
	return apiSong.Track.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (apiSong *APIFullSong) New() item.Item {
	return item.Item(NewAPIFullSong())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (apiSong *APIFullSong) Init(v interface{}) {
	switch v.(type) {
	case int:
		apiSong.Track.Id = dna.Int(v.(int))
	case dna.Int:
		apiSong.Track.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (apiSong *APIFullSong) Save(db *sqlpg.DB) error {

	var queries = dna.StringArray{}
	var err error

	// Getting artist queries
	artists := apiSong.ToArtists()
	for _, artist := range artists {
		queries.Push(sqlpg.GetInsertIgnoreStatement(sqlpg.GetTableName(artist), artist, "id", artist.Id, false))
	}

	// Getting album query
	album := apiSong.ToAlbum()
	queries.Push(sqlpg.GetInsertIgnoreStatement(sqlpg.GetTableName(album), album, "id", album.Id, false))

	// Getting song query
	song := apiSong.ToSong()
	queries.Push(sqlpg.GetInsertStatement(sqlpg.GetTableName(song), song, false))

	for _, query := range queries {
		_, err = db.Exec(query.String())
	}

	if err != nil {
		errQueries := dna.StringArray(queries.Map(func(val dna.String, idx dna.Int) dna.String {
			return "$$$error$$$" + val + "$$$error$$$"
		}).([]dna.String))
		return errors.New(err.Error() + errQueries.Join("\n").String())
	} else {
		return nil
	}

}
