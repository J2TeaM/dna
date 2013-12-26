package ke

import (
	"dna"
	"net/http"
	"time"
)

// Header defines a struc header posted to keeng service.
var Header = http.Header{
	"Accept":          []string{"*/*"},
	"Accept-Encoding": []string{"gzip,deflate"},
	"Content-Type":    []string{"text/xml; charset=utf-8"},
	"Accept-Language": []string{"en-US,en"},
	// "SOAPAction":      []string{"http://tempuri.org/getAlbum_v2"},
	"Connection": []string{"keep-alive"},
	"User-Agent": []string{"Keeng/2.0 CFNetwork/609.1.4 Darwin/13.0.0"},

	// req.Header.Add("Content-Length", "336")
}

// APILyric defines a struct of lyric from API.
type APILyric struct {
	Data   dna.String `json:"data"`
	Status dna.Int    `json:"status"`
}

// APIStatusAlbum defines a struct of status album from API.
type APIStatusAlbum struct {
	Status dna.Int  `json:"status"`
	Data   APIAlbum `json:"data"`
}

// APIStatusSong defines a struct of status song from API.
type APIStatusSong struct {
	Status dna.Int      `json:"status"`
	Data   APISongEntry `json:"data"`
}

// APISongEntry defines a struct of sing from API.
// When a song is fetched, all of its similar songs are fetched as well.
type APISongEntry struct {
	MainSong      APISong   `json:"song_detail"`
	RelevantSongs []APISong `json:"song_lienquan"`
}

// APIAlbum defines a struct of album from API.
type APIAlbum struct {
	Id       dna.Int    `json:"id"`
	Title    dna.String `json:"name"`
	Artists  dna.String `json:"singer"`
	Coverart dna.String `json:"image"`
	Url      dna.String `json:"url"`
	Plays    dna.Int    `json:"listen_no"`
	SongList []APISong  `json:"song_list"`
}

// ToAlbum converts APIAlbum to Album type.
func (apiAlbum *APIAlbum) ToAlbum() *Album {
	album := NewAlbum()
	album.Id = apiAlbum.Id
	keyArr := apiAlbum.Url.FindAllStringSubmatch(`.+/(.+?)\.html`, 1)
	if len(keyArr) > 0 {
		album.Key = keyArr[0][1]
	}
	album.Title = apiAlbum.Title
	album.Artists = apiAlbum.Artists.Split(" ft ")
	album.Plays = apiAlbum.Plays
	songids := dna.IntArray{}
	for _, song := range apiAlbum.SongList {
		songids.Push(song.Id)
	}
	album.Songids = songids
	album.Nsongs = dna.Int(len(apiAlbum.SongList))
	album.Description = ""
	album.Coverart = apiAlbum.Coverart
	dateCreatedArr := album.Coverart.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
	if len(dateCreatedArr) > 0 {
		year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
		month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
		day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
		album.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

	}
	album.Checktime = time.Now()
	return album
}

// APIAlbum defines a struct of song from API.
type APISong struct {
	Id                dna.Int    `json:"id"`
	Title             dna.String `json:"name"`
	Artists           dna.String `json:"singer"`
	Plays             dna.Int    `json:"listen_no"`
	ListenType        dna.Int    `json:"listen_type"`
	Lyric             dna.String `json:"lyric"`
	Link              dna.String `json:"media_url"`
	MediaUrlMono      dna.String `json:"media_url_mono"`
	MediaUrlPre       dna.String `json:"media_url_pre"`
	DownloadUrl       dna.String `json:"download_url"`
	IsDownload        dna.Int    `json:"is_download"`
	RingbacktoneCode  dna.String `json:"ringbacktone_code"`
	RingbacktonePrice dna.Int    `json:"ringbacktone_price"`
	Url               dna.String `json:"url"`
	Price             dna.Int    `json:"price"`
	Copyright         dna.Int    `json:"copyright"`
	CrbtId            dna.Int    `json:"crbt_id"`
	Coverart          dna.String `json:"image"`    // return "null" (string) if not available
	Coverart310       dna.String `json:"image310"` // return "null" (string) if not available
}

// ToSong converts APISong to Song type.
func (apiSong *APISong) ToSong() *Song {
	song := NewSong()
	song.Id = apiSong.Id
	keyArr := apiSong.Url.FindAllStringSubmatch(`.+/(.+?)\.html`, 1)
	if len(keyArr) > 0 {
		song.Key = keyArr[0][1]
	}
	song.Title = apiSong.Title
	song.Artists = apiSong.Artists.Split(" ft ")
	song.Plays = apiSong.Plays
	song.ListenType = apiSong.ListenType
	song.Lyric = apiSong.Lyric
	song.Link = apiSong.Link
	song.MediaUrlMono = apiSong.MediaUrlMono
	song.MediaUrlPre = apiSong.MediaUrlPre
	song.DownloadUrl = apiSong.DownloadUrl
	song.IsDownload = apiSong.IsDownload
	song.RingbacktoneCode = apiSong.RingbacktoneCode
	song.RingbacktonePrice = apiSong.RingbacktonePrice
	// song.Url = apiSong.Url
	song.Price = apiSong.Price
	song.Copyright = apiSong.Copyright
	song.CrbtId = apiSong.CrbtId
	song.Coverart = apiSong.Coverart
	song.Coverart310 = apiSong.Coverart310
	dateCreatedArr := song.Link.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
	if len(dateCreatedArr) > 0 {
		year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
		month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
		day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
		song.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

	}
	song.Checktime = time.Now()
	return song
}

// APIStatusArtist defines a struct of status artist from API.
type APIStatusArtist struct {
	Status dna.Int   `json:"status"`
	Data   APIArtist `json:"data"`
}

// APIArtist defines a struct of artist from API.
type APIArtist struct {
	Artist  APIArtistProfile `json:"singer"`
	Nsongs  dna.Int          `json:"num_song"`
	Nalbums dna.Int          `json:"num_album"`
	Nvideos dna.Int          `json:"num_video"`
}

// ToArtist converts APIArtist to Artist type
func (apiArtistEntry *APIArtist) ToArtist() *Artist {
	artist := NewArtist()
	artist.Id = apiArtistEntry.Artist.Id
	artist.Title = apiArtistEntry.Artist.Title
	artist.Coverart = apiArtistEntry.Artist.Coverart
	artist.Nsongs = apiArtistEntry.Nsongs
	artist.Nalbums = apiArtistEntry.Nalbums
	artist.Nvideos = apiArtistEntry.Nvideos
	artist.Checktime = time.Now()
	return artist
}

// APiArtistProfile defines a struct of artist profile from API.
type APIArtistProfile struct {
	Id       dna.Int    `json:"id"`
	Title    dna.String `json:"name"`
	Coverart dna.String `json:"image"`
}

// APIArtistSongs defines a struct representing some songs from an artist.
type APIArtistSongs struct {
	Status dna.Int   `json:"status"`
	Data   []APISong `json:"data"`
}

// APIArtistAlbums defines a struct representing some albums from an artist.
type APIArtistAlbums struct {
	Status dna.Int    `json:"status"`
	Data   []APIAlbum `json:"data"`
}

// APIArtistAlbums defines a struct representing some videos from an artist.
type APIArtistVideos struct {
	Status dna.Int    `json:"status"`
	Data   []APIVideo `json:"data"`
}

// APIVideo defines a struct of video from API.
type APIVideo struct {
	Id                dna.Int    `json:"id"`
	Title             dna.String `json:"name"`
	Artists           dna.String `json:"singer"`
	Plays             dna.Int    `json:"listen_no"`
	ListenType        dna.Int    `json:"listen_type"`
	Link              dna.String `json:"media_url"`
	IsDownload        dna.Int    `json:"is_download"`
	DownloadUrl       dna.String `json:"download_url"`
	RingbacktoneCode  dna.String `json:"ringbacktone_code"`
	RingbacktonePrice dna.Int    `json:"ringbacktone_price"`
	Url               dna.String `json:"url"`
	Price             dna.Int    `json:"price"`
	Copyright         dna.Int    `json:"copyright"`
	CrbtId            dna.Int    `json:"crbt_id"`
	Thumbnail         dna.String `json:"image"` // return "null" (string) if not available
}

// ToVideo converts APIVideo to Video type
func (apiVideo *APIVideo) ToVideo() *Video {
	video := NewVideo()
	video.Id = apiVideo.Id
	keyArr := apiVideo.Url.FindAllStringSubmatch(`.+/(.+?)\.html`, 1)
	if len(keyArr) > 0 {
		video.Key = keyArr[0][1]
	}
	video.Title = apiVideo.Title
	video.Artists = apiVideo.Artists.Split(" ft ")
	video.Plays = apiVideo.Plays
	video.ListenType = apiVideo.ListenType
	video.Link = apiVideo.Link
	video.IsDownload = apiVideo.IsDownload
	video.DownloadUrl = apiVideo.DownloadUrl
	video.RingbacktoneCode = apiVideo.RingbacktoneCode
	video.RingbacktonePrice = apiVideo.RingbacktonePrice
	video.Price = apiVideo.Price
	video.Copyright = apiVideo.Copyright
	video.CrbtId = apiVideo.CrbtId
	video.Thumbnail = apiVideo.Thumbnail
	dateCreatedArr := video.Link.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
	if len(dateCreatedArr) > 0 {
		year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
		month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
		day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
		video.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
	}
	video.Checktime = time.Now()
	return video
}
