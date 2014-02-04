package nct

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
)

// SongFinder defines a new song page
type SongFinder struct {
	PathIndex    dna.Int
	Page         dna.Int
	SongPortions dna.StringArray
}

func NewSongFinder() *SongFinder {
	sf := new(SongFinder)
	sf.PathIndex = 0
	sf.Page = 0
	sf.SongPortions = dna.StringArray{}
	return sf
}

func GetNewestSongPortions(pathIndex, page dna.Int) (dna.StringArray, error) {
	var baseUrl = dna.String("http://www.nhaccuatui.com")
	var officialArr dna.StringArray
	var keyArr []dna.StringArray
	var ret = dna.StringArray{}
	link := baseUrl + NewestSongPaths[pathIndex].Replace(`.html`, "."+page.ToString()+".html")
	// dna.Log(link)
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		similarKeysArr := data.FindAllString(`(?mis)<div class="info_song">.+?</div>`, -1)
		for _, val := range similarKeysArr {
			officialArr = val.FindAllString(`<a.+href.+">.+</a>`, -1)
			if officialArr.Length() > 0 {
				keyArr = officialArr[0].GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
				if len(keyArr) > 0 {
					ret.Push(keyArr[0][1])
				}
			}
		}

		return ret, nil
	} else {
		return nil, err
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (sf *SongFinder) Fetch() error {
	ret, err := GetNewestSongPortions(sf.PathIndex, sf.Page)
	if err != nil {
		return err
	} else {
		sf.SongPortions = ret
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (sf *SongFinder) GetId() dna.Int {
	return sf.PathIndex
}

// New implements item.Item interface
// Returns new item.Item interface
func (sf *SongFinder) New() item.Item {
	return item.Item(NewSongFinder())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (sf *SongFinder) Init(v interface{}) {
	var n dna.Int
	var NSongPaths dna.Int = NewestSongPaths.Length() // The total of song paths
	switch v.(type) {
	case int:
		n = dna.Int(v.(int))
	case dna.Int:
		n = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
	pathIndex := dna.Int(n / TotalSongPages)
	if pathIndex >= NSongPaths {
		pathIndex = NSongPaths - 1
	}
	sf.PathIndex = pathIndex
	sf.Page = n%TotalSongPages + 1
}

func (sf *SongFinder) Save(db *sqlpg.DB) error {

	// dna.Log(sf.SongPortions)
	sf.SongPortions = sf.SongPortions.Unique()
	err := FilterKeys(&sf.SongPortions, "nctsongs", db)
	// err := sf.SongPortions.FilterByIds("nctsongs", db)
	if err != nil {
		return err
	} else {
		mutex.Lock()
		NewestSongPortions = NewestSongPortions.Concat(sf.SongPortions)
		// *NewestSongPortions = append(*NewestSongPortions, *sf.SongPortions...)
		mutex.Unlock()
		return nil
	}

}
