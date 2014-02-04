package sf

import (
	"dna"
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
		video.Checktime.Format("2006-01-02 15:04:05"),
	}
}
