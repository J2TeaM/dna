package csn

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"time"
)

type Video struct {
	Id           dna.Int
	Title        dna.String
	Artists      dna.StringArray
	Authors      dna.StringArray
	Topics       dna.StringArray
	Thumbnail    dna.String
	Producer     dna.String
	Downloads    dna.Int
	Plays        dna.Int
	Formats      dna.String
	Href         dna.String
	IsLyric      dna.Int
	Lyric        dna.String
	DateReleased dna.String
	DateCreated  time.Time
	Type         dna.Bool
	Checktime    time.Time
}

// NewVideo returns new video whose id is 0
func NewVideo() *Video {
	video := new(Video)
	video.Id = 0
	video.Title = ""
	video.Artists = dna.StringArray{}
	video.Authors = dna.StringArray{}
	video.Topics = dna.StringArray{}
	video.Thumbnail = ""
	video.Producer = ""
	video.Downloads = 0
	video.Plays = 0
	video.Formats = ""
	video.Href = ""
	video.IsLyric = 0
	video.Lyric = ""
	video.DateReleased = ""
	video.DateCreated = time.Time{}
	video.Type = true
	video.Checktime = time.Time{}
	return video
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (video *Video) Fetch() error {
	return nil
	// _video, err := GetVideo(video.Id)
	// if err != nil {
	// 	return err
	// } else {
	// 	*video = *_video
	// 	return nil
	// }
}

// GetId implements GetId methods of item.Item interface
func (video *Video) GetId() dna.Int {
	return video.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (video *Video) New() item.Item {
	return item.Item(NewVideo())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (video *Video) Init(v interface{}) {
	switch v.(type) {
	case int:
		video.Id = dna.Int(v.(int))
	case dna.Int:
		video.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (video *Video) Save(db *sqlpg.DB) error {

	return db.InsertIgnore(video)
}
