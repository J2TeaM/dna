package zi

import (
	. "dna"
	"dna/http"
	"dna/site"
	"dna/sqlpg"
	"errors"
)

// TV defines basic TV type
type TV struct {
	Key     String
	Id      Int
	Title   String
	Artists StringArray
	Authors StringArray
	Link    String
}

// NewTV return a pointer to new TV
func NewTV() *TV {
	tv := new(TV)
	tv.Key = ""
	tv.Id = 0
	tv.Title = ""
	tv.Artists = StringArray{}
	tv.Authors = StringArray{}
	tv.Link = ""
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

// getTVFromMainPage returns song from main page
func getTVFromMainPage(tv *TV) <-chan bool {

	channel := make(chan bool, 1)
	go func() {
		link := "xxxxxxxxxxxxxx" + GetKey(tv.Id) + ".html"
		result, err := http.Get(link)
		// Log(link)
		// Log(result.Data)
		data := &result.Data
		Log(data.Match("<title>Thông báo</title>"))
		if err == nil && !data.Match("<title>Thông báo</title>") {

			//////// CHANGE!!!!!
			topicsArr := data.FindAllStringSubmatch(`Thể loại:(.+)\|`, -1)
			Log(topicsArr)

		}
		channel <- true

	}()
	return channel
}

// GetTV returns a tv or an error
func GetTV(id Int) (*TV, error) {
	var tv *TV = NewTV()
	tv.Id = id
	c := make(chan bool, 2)

	// go func() {
	// 	c <- <-getSongFromXML(tv)
	// }()
	go func() {
		c <- <-getTVFromMainPage(tv)
	}()

	for i := 0; i < 2; i++ {
		<-c
	}

	if tv.Key == "" {
		return nil, errors.New(Sprintf("Zing - Song %v: Mp3 link not found", tv.Id).String())
	} else {
		return tv, nil
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
	return site.Item(NewSong())
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
