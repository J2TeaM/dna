package ke

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

type Artist struct {
	Id        dna.Int
	Title     dna.String
	Coverart  dna.String
	Nsongs    dna.Int
	Nalbums   dna.Int
	Nvideos   dna.Int
	Checktime time.Time
}

func NewArtist() *Artist {
	artist := new(Artist)
	artist.Id = 0
	artist.Title = ""
	artist.Coverart = ""
	artist.Nsongs = 0
	artist.Nalbums = 0
	artist.Nvideos = 0
	artist.Checktime = time.Time{}
	return artist
}

// GetArtist returns a album.
func GetArtist(id dna.Int) (*Artist, error) {
	apiArtist, err := GetAPIArtistEntry(id)
	if err != nil {
		return nil, err
	} else {
		artist := apiArtist.ToArtist()
		if artist.Id == 0 {
			return nil, errors.New("Invalid artistid: Zero value found")
		} else {
			return artist, nil
		}
	}
}

// GetSongs returns a list of song from an artist or an error
func (artist *Artist) GetSongs() ([]Song, error) {
	var songs []Song
	apiArtistSongs, err := GetAPIArtistSongs(artist.Id, 1, artist.Nsongs+100)
	if err != nil {
		return nil, err
	} else {
		for _, apiSong := range apiArtistSongs.Data {
			songs = append(songs, *apiSong.ToSong())
		}
		return songs, nil
	}
}

// GetAlbums returns a list of albums from an artist or an error
func (artist *Artist) GetAlbums() ([]Album, error) {
	var albums []Album
	apiArtistAlbums, err := GetAPIArtistAlbums(artist.Id, 1, artist.Nalbums+100)
	if err != nil {
		return nil, err
	} else {
		for _, apiAlbum := range apiArtistAlbums.Data {
			albums = append(albums, *apiAlbum.ToAlbum())
		}
		return albums, nil
	}
}

// GetVideos returns a list of videos from an artist or an error
func (artist *Artist) GetVideos() ([]Video, error) {
	var videos []Video
	apiArtistVideos, err := GetAPIArtistVideos(artist.Id, 1, artist.Nvideos+100)
	if err != nil {
		return nil, err
	} else {
		for _, apiVideo := range apiArtistVideos.Data {
			videos = append(videos, *apiVideo.ToVideo())
		}
		return videos, nil
	}
}

// Fetch implements item.Item interface
func (artist *Artist) Fetch() error {
	_artist, err := GetArtist(artist.Id)
	if err != nil {
		return err
	} else {
		*artist = *_artist
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (artist *Artist) GetId() dna.Int {
	return artist.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (artist *Artist) New() item.Item {
	return item.Item(NewArtist())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (artist *Artist) Init(v interface{}) {
	switch v.(type) {
	case int:
		artist.Id = dna.Int(v.(int))
	case dna.Int:
		artist.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (artist *Artist) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(artist)
}
