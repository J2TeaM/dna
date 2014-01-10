package sf

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"time"
)

type Video struct {
	// Auto increased id
	Songid      dna.Int
	YoutubeId   dna.String
	Title       dna.String
	Description dna.String
	Duration    dna.Int
	Thumbnail   dna.String
	Checktime   time.Time
}

// NewVideo return default new video
func NewVideo() *Video {
	video := new(Video)
	video.Songid = 0
	video.YoutubeId = ""
	video.Duration = 0
	video.Thumbnail = ""
	video.Title = ""
	video.Description = ""
	video.Checktime = time.Now()
	return video
}

//
//
//psql -c "COPY sfvideos (songid,youtube_id,title,description,duration,thumbnail,checktime) FROM '/Users/daonguyenanbinh/Box Documents/Sites/golang/sfvideos.csv' DELIMITER ',' CSV"
func (video *Video) CSVRecord() []string {
	return []string{
		video.Songid.ToString().String(),
		video.YoutubeId.String(),
		video.Title.String(),
		video.Description.String(),
		video.Duration.ToString().String(),
		video.Thumbnail.String(),
		video.Checktime.Format("2006-02-01 15:04:05"),
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (video *Video) Fetch() error {
	return nil
}

// GetId implements GetId methods of item.Item interface
func (video *Video) GetId() dna.Int {
	return 0
}

// New implements item.Item interface
// Returns new item.Item interface
func (video *Video) New() item.Item {
	return item.Item(NewVideo())
}

func (video *Video) Init(v interface{}) {
	// do nothing
}

func (video *Video) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(video)
}
