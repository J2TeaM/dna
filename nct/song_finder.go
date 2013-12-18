package nct

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
)

// SongFinder defines a new song page
type SongFinder struct {
	PathIndex    dna.Int
	Page         dna.Int
	SongPortions *Portions
}

func NewSongFinder() *SongFinder {
	sf := new(SongFinder)
	sf.PathIndex = 0
	sf.Page = 0
	sf.SongPortions = &Portions{}
	return sf
}

func GetNewestSongPortions(pathIndex, page dna.Int) (*Portions, error) {
	var baseUrl = dna.String("http://www.nhaccuatui.com")
	var similarIdsArr dna.StringArray
	var portion *Portion = NewPortion()
	var officialArr dna.StringArray
	var keyArr []dna.StringArray
	var key string
	var ret = &Portions{}
	link := baseUrl + NewestSongPaths[pathIndex].Replace(`.html`, "."+page.ToString()+".html")
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		similarKeysArr := data.FindAllString(`(?mis)<span class="rel">.+?</span>`, -1)
		similarIdsArr = data.FindAllString(`<div class="num">.+?</div>`, -1)
		if similarIdsArr.Length() != similarKeysArr.Length() {
			return nil, errors.New("CRITICAL ERROR: Lengths of id and key list do not match")
		} else {
			for idx, val := range similarKeysArr {
				officialArr = val.FindAllString(`<a.+href.+">.+</a>`, -1)
				if officialArr.Length() > 0 {
					keyArr = officialArr[0].GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
					if len(keyArr) > 0 {
						key = string(keyArr[0][1])
					}
					if officialArr[0].GetTagAttributes("class") == "mof" {
						portion.IsOfficial = true
					}
				}
				idArr := similarIdsArr[idx].FindAllStringSubmatch(`NCTCounter_sg_([0-9]+)`, -1)
				if len(idArr) > 0 {
					portion.Id = int32(idArr[0][1].ToInt())
				}
				ret.Push(&Portion{int32(idArr[0][1].ToInt()), key, false})
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
	var NSongPaths dna.Int = dna.Int(NewestSongPaths.Length()) // The total of song paths
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
	sf.SongPortions.UniqueByIds()
	err := sf.SongPortions.FilterByIds("nctsongs", db)
	if err != nil {
		return err
	} else {
		mutex.Lock()
		*NewestSongPortions = append(*NewestSongPortions, *sf.SongPortions...)
		mutex.Unlock()
		return nil
	}

}
