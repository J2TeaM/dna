package lw

import (
	"dna"
	"dna/http"
	"time"
)

type Song struct {
	Id                     dna.Int
	Title                  dna.String
	ArtistId               dna.Int
	Artists                dna.StringArray
	AlbumId                dna.Int
	AlbumTitle             dna.String
	Year                   dna.String
	Lyric                  dna.String
	DownloadDone           dna.Int
	IsGracenote            dna.String
	DownloadGracenoteDone  dna.Int
	GracenoteLyric         dna.String
	GracenoteSongwriters   dna.String
	GracenotePublishers    dna.String
	Metrolyric             dna.String
	Tags                   dna.String
	DownloadMetrolyricDone dna.Int
	AllmusicLink           dna.String
	MusicbrainzLink        dna.String
	YoutubeLink            dna.String
	Language               dna.String
	DateCreated            time.Time
}

func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.ArtistId = 0
	song.Artists = dna.StringArray{}
	song.AlbumId = 0
	song.AlbumTitle = ""
	song.Year = ""
	song.Lyric = ""
	song.DownloadDone = 0
	song.IsGracenote = ""
	song.DownloadGracenoteDone = 0
	song.GracenoteLyric = ""
	song.GracenoteSongwriters = ""
	song.GracenotePublishers = ""
	song.Metrolyric = ""
	song.Tags = ""
	song.DownloadMetrolyricDone = 0
	song.AllmusicLink = ""
	song.MusicbrainzLink = ""
	song.YoutubeLink = ""
	song.Language = ""
	song.DateCreated = time.Now()
	return song
}

func getGracenoteSongLyric(artist, title dna.String, song *Song) {
	link := "http://lyrics.wikia.com/Gracenote:" + artist.Replace(" ", "_") + ":" + title.Replace(" ", "_")
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data

		writersArr := data.FindAllString(`Songwriters.+`, 1)
		if writersArr.Length() > 0 {
			song.GracenoteSongwriters = writersArr[0].GetTags("em")[0].RemoveHtmlTags("").DecodeHTML()
		}

		publisheraArr := data.FindAllString(`Publishers.+`, 1)
		if publisheraArr.Length() > 0 {
			song.GracenotePublishers = publisheraArr[0].GetTags("em")[0].RemoveHtmlTags("").DecodeHTML()
		}

		lyricArr := data.FindAllStringSubmatch(`(?mis)<div class='lyricbox'>(.+?)<\!--`, 1)
		if len(lyricArr) > 0 {
			song.GracenoteLyric = lyricArr[0][1].Trim().DecodeHTML().ReplaceWithRegexp(`(?mis)^<div.+?</span></div>`, "").Trim().Replace("<br />", "\n")
		}

		if song.GracenoteLyric != "" {
			song.DownloadGracenoteDone = 1
		} else {
			song.DownloadGracenoteDone = 0
		}

	}
}

func GetSong(artist, title dna.String) (*Song, error) {
	link := "http://lyrics.wikia.com/" + artist.Replace(" ", "_") + ":" + title.Replace(" ", "_")
	result, err := http.Get(link)
	// Log(link)
	// Log(result.Data)
	song := NewSong()
	song.Title = title
	song.Artists = artist.ToStringArray()
	if err == nil {
		data := &result.Data

		if data.Match(`class='lyricbox'.+Instrumental.+TrebleClef`) == true {
			song.Lyric = "Instrumental"
		}

		lyricArr := data.FindAllStringSubmatch(`(?mis)<div class='lyricbox'>(.+?)<!--`, 1)
		if len(lyricArr) > 0 {
			song.Lyric = lyricArr[0][1].Trim().DecodeHTML().ReplaceWithRegexp(`(?mis)^<div.+?</span></div>`, "").Trim().Replace("<br />", "\n")
		}

		if song.Lyric == "" {
			song.DownloadDone = -1
		} else {
			song.DownloadDone = 1
		}

		musicBrainzArr := data.FindAllString(`http://musicbrainz.+html`, 1)
		if musicBrainzArr.Length() > 0 {
			song.MusicbrainzLink = musicBrainzArr[0]
		}

		allmusicArr := data.FindAllString(`http://www\.allmusic\.com.+`, 1)
		if allmusicArr.Length() > 0 {
			song.AllmusicLink = allmusicArr[0].ReplaceWithRegexp(`".+$`, "")
		}

		youtubeArr := data.FindAllString(`http://youtube\.com.+`, 1)
		if youtubeArr.Length() > 0 {
			song.YoutubeLink = youtubeArr[0].ReplaceWithRegexp(`".+$`, "")
		}

		if data.Match("View the Gracenote") {
			song.IsGracenote = "1"
		}

		languageArr := data.FindAllStringSubmatch(`category normal" data-name="Language/(.+)" data-namespace`, 1)
		if len(languageArr) > 0 {
			song.Language = languageArr[0][1]
		}

		albumArr := data.FindAllStringSubmatch(`appears on the album.+">(.+)</a></i>`, 1)
		if len(albumArr) > 0 {
			song.AlbumTitle = albumArr[0][1]
		}

		if song.AlbumTitle.Match(`[0-9]+`) {
			song.Year = song.AlbumTitle.FindAllStringSubmatch(`([0-9]+)`, 1)[0][1]
		}

		if data.Match(`It has been suggested that Gracenote`) {
			song.IsGracenote = "1"
			getGracenoteSongLyric(artist, title, song)
		}

		return song, nil
	} else {
		return nil, err
	}

}
