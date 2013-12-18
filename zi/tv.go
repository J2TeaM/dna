package zi

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

// Defines API key and session key TV
const (
	TV_API_KEY     = dna.String("d04210a70026ad9323076716781c223f")
	TV_SESSION_KEY = dna.String("91618dfec493ed7dc9d61ac088dff36b")
)

// TV defines basic TV type
//
// NOTICE: SubTitle and Tracking fields are not properly decoded.
type TV struct {
	Id              dna.Int
	Key             dna.String
	Title           dna.String
	Fullname        dna.String
	Episode         dna.Int
	DateReleased    time.Time
	Duration        dna.Int
	Thumbnail       dna.String
	FileUrl         dna.String
	ResolutionFlags dna.Int
	// LinkUrl          dna.String
	ProgramId        dna.Int
	ProgramName      dna.String
	ProgramThumbnail dna.String
	ProgramGenreIds  dna.IntArray
	ProgramGenres    dna.StringArray
	Plays            dna.Int
	Comments         dna.Int
	Likes            dna.Int
	Rating           dna.Float
	Subtitle         dna.String
	Tracking         dna.String
	Signature        dna.String
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
	tv.ProgramGenreIds = dna.IntArray{}
	tv.ProgramGenres = dna.StringArray{}
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
func (tv *TV) GetEncodedKey() dna.String {
	return getCipherText(GetId(tv.Key), dna.IntArray{10, 2, 0, 1, 0})
}

// GetDirectLink gets a direct url for specific episode
func (tv *TV) GetDirectLink() dna.String {
	return TV_BASE_URL.Concat(tv.GetEncodedKey(), "/")
}

// GetTV returns a tv or an error
func GetTV(id dna.Int) (*TV, error) {
	var tv *TV = NewTV()
	apiTV, err := GetAPITV(id)
	if err == nil {
		tv.Id = apiTV.Id + ID_DIFFERENCE

		if tv.Id != id {
			return nil, errors.New(string(dna.Sprintf("Item id: %v - key:%v does not match", tv.Id, tv.Key)))
		}
		tv.Key = GetKey(tv.Id)
		tv.Title = apiTV.Title
		tv.Fullname = apiTV.Fullname
		tv.Episode = apiTV.Episode
		if apiTV.DateReleased != "" {
			timeFlds := apiTV.DateReleased.Trim().Split(`/`)
			if timeFlds.Length() != 3 {
				return nil, errors.New(string(dna.Sprintf("Date released of item id: %v - key:%v cannot be decoded", tv.Id, tv.Key)))
			}
			tv.DateReleased = time.Date(int(timeFlds[2].ToInt()), time.Month(timeFlds[1].ToInt()), int(timeFlds[0].ToInt()), 0, 0, 0, 0, time.UTC)
		}
		tv.Duration = apiTV.Duration
		tv.Thumbnail = apiTV.Thumbnail
		tv.FileUrl = apiTV.FileUrl
		flags := dna.Int(0)
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
			return nil, errors.New(string(dna.Sprintf("File url  of item id: %v - key:%v is not properly formated: No resolution found", tv.Id, tv.Key)))
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
		tv.ProgramGenreIds = dna.IntArray{}
		tv.ProgramGenres = dna.StringArray{}
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

// Fetch implements item.Item interface.
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

// GetId implements GetId methods of item.Item interface
func (tv *TV) GetId() dna.Int {
	return tv.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (tv *TV) New() item.Item {
	return item.Item(NewTV())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (tv *TV) Init(v interface{}) {
	switch v.(type) {
	case int:
		tv.Id = dna.Int(v.(int))
	case dna.Int:
		tv.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (tv *TV) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(tv)
}
