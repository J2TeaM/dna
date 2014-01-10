package sf

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"encoding/xml"
	"errors"
	"time"
)

type APISongFreaksAlbum struct {
	Id       dna.Int
	XMLName  xml.Name    `xml:"songfreaks"`
	Response APIResponse `xml:"response"`
	Album    APIAlbum    `xml:"album"`
}

func NewAPISongFreaksAlbum() *APISongFreaksAlbum {
	sf := new(APISongFreaksAlbum)
	sf.Id = 0
	sf.Response = APIResponse{202, 0, ""}
	sf.Album = APIAlbum{}
	return sf
}

func GetSongFreaksAlbum(id dna.Int) (*APISongFreaksAlbum, error) {
	var link dna.String = "http://apiv2.songfreaks.com//album.do?"
	// Log(link)
	PostData.SetIdKey(id)
	PostData.VideosEnable = false
	PostData.CommentsEnable = false
	result, err := Post(link, PostData.Encode())
	mutex.Lock()
	Cookie = result.Header.Get("Set-Cookie")
	mutex.Unlock()
	if err == nil {
		songfreaks := &APISongFreaksAlbum{}
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

func (sf *APISongFreaksAlbum) HasError() error {
	switch sf.Response.Code {
	case 100, 101, 102, 103, 111:
		// do nothing, request successful
		// code 103 has SUCCESS: LICENSE, NO LYRICS
	case 202, 201:
		return errors.New("No id found at index:" + sf.Id.ToString().String())
	default:
		mes := dna.Sprintf("%v", sf.Response).String()
		return errors.New("Unknow response code" + mes)
	}

	return nil
}

func (sf *APISongFreaksAlbum) ToAlbum() (*Album, error) {
	if sferr := sf.HasError(); sferr != nil {
		return nil, sferr
	}
	album := NewAlbum()
	album.Id = sf.Album.Id
	album.AMG = sf.Album.AMG
	album.UrlSlug = sf.Album.UrlSlug
	album.Year = sf.Album.Year
	album.Coverart = sf.Album.Coverart
	album.CoverartLarge = sf.Album.CoverartLarge
	album.Title = sf.Album.Title
	album.Ratings = dna.IntArray{sf.Album.Rating.AverageRating.ToInt(), sf.Album.Rating.UserRating, sf.Album.Rating.TotalRatings}
	album.Artistid = sf.Album.Artist.Id
	album.Artists = sf.Album.Artist.Name.ToStringArray()
	album.Link = sf.Album.Link
	songids := dna.IntArray{}
	for _, track := range sf.Album.Tracks {
		songids.Push(track.Id)
	}
	album.Songids = songids
	album.ReviewAuthor = sf.Album.Review.Author
	album.Review = sf.Album.Review.Value.DecodeHTML()
	album.Checktime = time.Now()
	return album, nil
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (sf *APISongFreaksAlbum) Fetch() error {
	_sf, err := GetSongFreaksAlbum(sf.Id)
	if err != nil {
		return err
	} else {
		*sf = *_sf
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (sf *APISongFreaksAlbum) GetId() dna.Int {
	return sf.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (sf *APISongFreaksAlbum) New() item.Item {
	return item.Item(NewAPISongFreaksAlbum())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (sf *APISongFreaksAlbum) Init(v interface{}) {
	switch v.(type) {
	case int:
		sf.Id = dna.Int(v.(int))
	case dna.Int:
		sf.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (sf *APISongFreaksAlbum) Save(db *sqlpg.DB) error {
	album, err := sf.ToAlbum()

	if err != nil {
		return err
	} else {
		return db.Update(album, "id", "ratings", "songids", "review_author", "review")
	}

}
