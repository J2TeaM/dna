package nct

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
)

// AlbumFinder defines a new song page
type AlbumFinder struct {
	PathIndex     dna.Int
	Page          dna.Int
	AlbumPortions dna.StringArray
}

func NewAlbumFinder() *AlbumFinder {
	af := new(AlbumFinder)
	af.PathIndex = 0
	af.Page = 0
	af.AlbumPortions = dna.StringArray{}
	return af
}

func GetNewestAlbumPortions(pathIndex, page dna.Int) (dna.StringArray, error) {
	var baseUrl = dna.String("http://www.nhaccuatui.com")
	var ret = dna.StringArray{}
	link := baseUrl + NewestAlbumPaths[pathIndex].Replace(`.html`, "."+page.ToString()+".html")
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		albumKeyArr := data.FindAllString(`(?mis)<div class="box_left">.+?</div>`, -1)
		if albumKeyArr.Length() > 0 {
			albumKeyArr.ForEach(func(val dna.String, idx dna.Int) {
				keyArr := val.GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
				if len(keyArr) > 0 {
					if !keyArr[0][1].Contains("-") {
						ret.Push(keyArr[0][1])
					}

				}
			})
		}
		return ret, nil
	} else {
		return nil, err
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (af *AlbumFinder) Fetch() error {
	ret, err := GetNewestAlbumPortions(af.PathIndex, af.Page)
	if err != nil {
		return err
	} else {
		af.AlbumPortions = ret
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (af *AlbumFinder) GetId() dna.Int {
	return af.PathIndex
}

// New implements item.Item interface
// Returns new item.Item interface
func (af *AlbumFinder) New() item.Item {
	return item.Item(NewAlbumFinder())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (af *AlbumFinder) Init(v interface{}) {
	var n dna.Int
	var NAlbumPaths dna.Int = dna.Int(NewestAlbumPaths.Length()) // The total of song paths
	switch v.(type) {
	case int:
		n = dna.Int(v.(int))
	case dna.Int:
		n = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
	pathIndex := dna.Int(n / TotalAlbumPages)
	if pathIndex >= NAlbumPaths {
		pathIndex = NAlbumPaths - 1
	}
	af.PathIndex = pathIndex
	af.Page = n%TotalAlbumPages + 1
}

func (af *AlbumFinder) Save(db *sqlpg.DB) error {
	af.AlbumPortions = af.AlbumPortions.Unique()
	err := FilterKeys(&af.AlbumPortions, "nctalbums", db)
	// err := af.AlbumPortions.FilterByKeys("nctalbums", db)
	if err != nil {
		return err
	} else {
		mutex.Lock()
		NewestAlbumPortions = NewestAlbumPortions.Concat(af.AlbumPortions)
		mutex.Unlock()
		return nil
	}
}
