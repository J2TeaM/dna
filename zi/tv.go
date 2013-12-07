package zi

import (
	. "dna"
	"dna/site"
	"dna/sqlpg"
	"errors"
	"time"
)

// Defines API key and session key TV
const (
	TV_API_KEY     = String("d04210a70026ad9323076716781c223f")
	TV_SESSION_KEY = String("91618dfec493ed7dc9d61ac088dff36b")
)

// TV defines basic TV type
//
// NOTICE: SubTitle and Tracking fields are not properly decoded.
type TV struct {
	Id              Int
	Key             String
	Title           String
	Fullname        String
	Episode         Int
	DateReleased    time.Time
	Duration        Int
	Thumbnail       String
	FileUrl         String
	ResolutionFlags Int
	// LinkUrl          String
	ProgramId        Int
	ProgramName      String
	ProgramThumbnail String
	ProgramGenreIds  IntArray
	ProgramGenres    StringArray
	Plays            Int
	Comments         Int
	Likes            Int
	Rating           Float
	Subtitle         String
	Tracking         String
	Signature        String
	Checktime        time.Time
}

// NewTV return a pointer to new TV
func NewTV() *TV {
	tv := new(TV)
	tv.Key = ""
	tv.Id = 0
	tv.Title = ""
	tv.Fullname = ""
	tv.Episode = 0
	tv.DateReleased = time.Time{}
	tv.Duration = 0
	tv.Thumbnail = ""
	tv.FileUrl = ""
	tv.ResolutionFlags = 0
	// tv.LinkUrl = ""
	tv.ProgramId = 0
	tv.ProgramName = ""
	tv.ProgramThumbnail = ""
	tv.ProgramGenreIds = IntArray{}
	tv.ProgramGenres = StringArray{}
	tv.Plays = 0
	tv.Comments = 0
	tv.Likes = 0
	tv.Rating = 0
	tv.Subtitle = ""
	tv.Tracking = ""
	tv.Signature = ""
	tv.Checktime = time.Time{}
	return tv
}

// GetEncodedKey gets an encoded key of a video
func (tv *TV) GetEncodedKey() String {
	return getCipherText(GetId(tv.Key), IntArray{10, 2, 0, 1, 0})
}

// GetDirectLink gets a direct url for specific episode
func (tv *TV) GetDirectLink() String {
	return TV_BASE_URL.Concat(tv.GetEncodedKey(), "/")
}

// GetTV returns a tv or an error
func GetTV(id Int) (*TV, error) {
	var tv *TV = NewTV()
	apiTV, err := GetAPITV(id)
	if err == nil {
		tv.Id = apiTV.Id + ID_DIFFERENCE

		if tv.Id != id {
			return nil, errors.New(string(Sprintf("Item id: %v - key:%v does not match", tv.Id, tv.Key)))
		}
		tv.Key = GetKey(tv.Id)
		tv.Title = apiTV.Title
		tv.Fullname = apiTV.Fullname
		tv.Episode = apiTV.Episode
		if apiTV.DateReleased != "" {
			timeFlds := apiTV.DateReleased.Split(`/`)
			if timeFlds.Length() != 3 {
				return nil, errors.New(string(Sprintf("Date released of item id: %v - key:%v cannot be decoded", tv.Id, tv.Key)))
			}
			tv.DateReleased = time.Date(int(timeFlds[2].ToInt()), time.Month(timeFlds[1].ToInt()), int(timeFlds[0].ToInt()), 0, 0, 0, 0, time.UTC)
		}
		tv.Duration = apiTV.Duration
		tv.Thumbnail = apiTV.Thumbnail
		tv.FileUrl = apiTV.FileUrl
		flags := Int(0)
		tmp := apiTV.FileUrl.FindAllStringSubmatch(`format=(.+)&`, -1)
		if len(tmp) > 0 {
			switch tmp[0][1] {
			case "f3gp":
				flags = flags | LRe240p
			case "f360":
				flags = flags | LRe360p
			case "f480":
				flags = flags | LRe480p
			case "f720":
				flags = flags | LRe720p
			case "f1080":
				flags = flags | LRe1080p
			}
		} else {
			return nil, errors.New(string(Sprintf("File url  of item id: %v - key:%v is not properly formated: No resolution found", tv.Id, tv.Key)))
		}

		for key, val := range apiTV.OtherUrl {
			switch {
			case key == "Video3GP" && val != "":
				flags = flags | LRe240p
			case key == "Video360" && val != "":
				flags = flags | LRe360p
			case key == "Video480" && val != "":
				flags = flags | LRe480p
			case key == "Video720" && val != "":
				flags = flags | LRe720p
			case key == "Video1080" && val != "":
				flags = flags | LRe1080p
			}
		}
		tv.ResolutionFlags = flags
		// tv.LinkUrl = apiTV.LinkUrl
		tv.ProgramId = apiTV.ProgramId
		tv.ProgramName = apiTV.ProgramName
		tv.ProgramThumbnail = apiTV.ProgramThumbnail
		tv.ProgramGenreIds = IntArray{}
		tv.ProgramGenres = StringArray{}
		for _, genre := range apiTV.ProgramGenres {
			tv.ProgramGenreIds.Push(genre.Id)
			tv.ProgramGenres.Push(genre.Name)
		}
		tv.Plays = apiTV.Plays
		tv.Comments = apiTV.Comments
		tv.Likes = apiTV.Likes
		tv.Rating = apiTV.Rating
		tv.Subtitle = apiTV.SubTitle
		tv.Tracking = apiTV.Tracking
		tv.Signature = apiTV.Signature
		tv.Checktime = time.Now()
		return tv, nil
	} else {
		return nil, err
	}
}

// Fetch implements site.Item interface.
// Returns error if can not get item
func (tv *TV) Fetch() error {
	_tv, err := GetTV(tv.Id)
	if err != nil {
		return err
	} else {
		*tv = *_tv
		return nil
	}
}

// New implements site.Item interface
// Returns new site.Item interface
func (tv *TV) New() site.Item {
	return site.Item(NewTV())
}

// Init implements site.Item interface.
// It sets Id or key.
// Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (tv *TV) Init(v interface{}) {
	switch v.(type) {
	case int:
		tv.Id = Int(v.(int))
	case Int:
		tv.Id = v.(Int)
	default:
		panic("Interface v has to be int")
	}
}

func (tv *TV) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(tv)
}
