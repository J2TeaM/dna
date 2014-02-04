package am

import (
	"dna"
	"time"
)

type APISong struct {
	Id        dna.Int
	Title     dna.String
	Artists   []Person
	Composers []Person
	Duration  dna.Int
}

func (apisong *APISong) ToSong() *Song {
	song := NewSong()
	song.Id = apisong.Id
	song.Title = apisong.Title.DecodeHTML()
	song.Duration = apisong.Duration

	artistids := dna.IntArray{}
	artists := dna.StringArray{}
	for _, artist := range apisong.Artists {
		artistids.Push(artist.Id)
		artists.Push(artist.Name)
	}
	song.Artistids = artistids
	song.Artists = artists

	composerids := dna.IntArray{}
	composers := dna.StringArray{}
	for _, composer := range apisong.Composers {
		composerids.Push(composer.Id)
		composers.Push(composer.Name)
	}
	song.Composerids = composerids
	song.Composers = composers
	song.Checktime = time.Now()

	return song
}

type Song struct {
	Id          dna.Int
	Title       dna.String
	Artistids   dna.IntArray
	Artists     dna.StringArray
	Albumid     dna.Int
	Composerids dna.IntArray
	Composers   dna.StringArray
	Duration    dna.Int
	Checktime   time.Time
}

func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.Title = ""
	song.Duration = 0
	song.Artistids = dna.IntArray{}
	song.Artists = dna.StringArray{}
	song.Albumid = 0
	song.Composerids = dna.IntArray{}
	song.Composers = dna.StringArray{}
	song.Checktime = time.Time{}
	return song
}

func (song *Song) CSVRecord() []string {
	return []string{
		song.Id.ToString().String(),
		song.Title.String(),
		dna.Sprintf("%#v", song.Artistids).Replace("dna.IntArray", "").String(),
		dna.Sprintf("%#v", song.Artists).Replace("dna.StringArray", "").String(),
		song.Albumid.ToString().String(),
		dna.Sprintf("%#v", song.Composerids).Replace("dna.IntArray", "").String(),
		dna.Sprintf("%#v", song.Composers).Replace("dna.StringArray", "").String(),
		song.Duration.ToString().String(),
	}
}
