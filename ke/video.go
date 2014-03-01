package ke

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

// Video defines main Video type.
type Video struct {
	Id                dna.Int
	Key               dna.String
	Title             dna.String
	Artists           dna.StringArray
	Plays             dna.Int
	ListenType        dna.Int
	Link              dna.String
	IsDownload        dna.Int
	DownloadUrl       dna.String
	RingbacktoneCode  dna.String
	RingbacktonePrice dna.Int
	// Url               dna.String
	Price       dna.Int
	Copyright   dna.Int
	CrbtId      dna.Int
	Thumbnail   dna.String
	DateCreated time.Time
	Checktime   time.Time
}

// NewVideo returns new video.
func NewVideo() *Video {
	video := new(Video)
	video.Id = 0
	video.Key = ""
	video.Title = ""
	video.Artists = dna.StringArray{}
	video.Plays = 0
	video.ListenType = 0
	video.Link = ""
	video.IsDownload = 0
	video.DownloadUrl = ""
	video.RingbacktoneCode = ""
	video.RingbacktonePrice = 0
	// video.Url = ""
	video.Price = 0
	video.Copyright = 0
	video.CrbtId = 0
	video.Thumbnail = ""
	video.DateCreated = time.Time{}
	video.Checktime = time.Time{}
	return video
}

// GetVideo returns a video.
func GetVideo(id dna.Int) (*Video, error) {
	apiVideo, err := GetAPIVideo(id)
	if err != nil {
		return nil, err
	} else {
		video := apiVideo.ToVideo()
		if video.Id == 0 {
			return nil, errors.New(dna.Sprintf("Keeng - Video ID: %v not found", id).String())
		} else {
			return video, nil
		}
	}
}

// Do not implement
func (video *Video) Fetch() error {
	_video, err := GetVideo(video.Id)
	if err != nil {
		return err
	} else {
		*video = *_video
		return nil
	}
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
