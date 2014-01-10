package sf

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"time"
)

type Song struct {
	Id             dna.Int
	TrackGroupId   dna.Int
	AMG            dna.Int
	Title          dna.String
	Albumid        dna.Int
	Artistids      dna.IntArray
	Artists        dna.StringArray
	UrlSlug        dna.String
	IsInstrumental dna.Bool
	Viewable       dna.Bool
	Duration       dna.Int // in seconds
	HasLrc         dna.Bool
	HasLyric       dna.Bool
	SubmittedLyric dna.Bool
	Lyricid        dna.Int
	TrackNumber    dna.Int
	DiscNumber     dna.Int
	Rating         dna.IntArray
	Link           dna.String
	Lrc            dna.String
	Lyric          dna.String
	Copyright      dna.String
	Writer         dna.String
	Checktime      time.Time
}

// NewSong return default new song
func NewSong() *Song {
	song := new(Song)
	song.Id = 0
	song.TrackGroupId = 0
	song.AMG = 0
	song.UrlSlug = ""
	song.IsInstrumental = false
	song.Viewable = false
	song.Duration = 0
	song.Lyricid = 0
	song.HasLrc = false
	song.HasLyric = false
	song.TrackNumber = 0
	song.DiscNumber = 0
	song.Title = ""
	song.Rating = dna.IntArray{0, 0, 0}
	song.Albumid = 0
	song.Artistids = dna.IntArray{}
	song.Artists = dna.StringArray{}
	song.Lrc = "{}"
	song.Link = ""
	song.Lyric = ""
	song.Copyright = ""
	song.Writer = ""
	song.SubmittedLyric = false
	song.Checktime = time.Time{}
	return song
}

//CSVRecord returns a record to write csv format.
//
//psql -c "COPY sfsongs (id,track_group_id,amg,title,albumid,artistids,artists,url_slug,is_instrumental,viewable,duration,lyricid,has_lrc,has_lyric,track_number,disc_number,rating,lrc,link,lyric,copyright,writer,submitted_lyric,checktime) FROM '/Users/daonguyenanbinh/Box Documents/Sites/golang/sfsongs.csv' DELIMITER ',' CSV"
func (song *Song) CSVRecord() []string {
	return []string{
		song.Id.ToString().String(),
		song.TrackGroupId.ToString().String(),
		song.AMG.ToString().String(),
		song.Title.String(),
		song.Albumid.ToString().String(),
		dna.Sprintf("%#v", song.Artistids).Replace("dna.IntArray", "").String(),
		dna.Sprintf("%#v", song.Artists).Replace("dna.StringArray", "").String(),
		song.UrlSlug.String(),
		song.IsInstrumental.String(),
		song.Viewable.String(),
		song.Duration.ToString().String(),
		song.Lyricid.ToString().String(),
		song.HasLrc.String(),
		song.HasLyric.String(),
		song.TrackNumber.ToString().String(),
		song.DiscNumber.ToString().String(),
		dna.Sprintf("%#v", song.Rating).Replace("dna.IntArray", "").String(),
		song.Lrc.String(),
		song.Link.String(),
		song.Lyric.String(),
		song.Copyright.String(),
		song.Writer.String(),
		song.SubmittedLyric.String(),
		song.Checktime.Format("2006-02-01 15:04:05"),
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (song *Song) Fetch() error {
	return nil
}

// GetId implements GetId methods of item.Item interface
func (song *Song) GetId() dna.Int {
	return song.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (song *Song) New() item.Item {
	return item.Item(NewSong())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (song *Song) Init(v interface{}) {
	switch v.(type) {
	case int:
		song.Id = dna.Int(v.(int))
	case dna.Int:
		song.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (song *Song) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(song)
}
