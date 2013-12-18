package nct

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
)

// VideoFinder defines a new song page
type VideoFinder struct {
	PathIndex     dna.Int
	Page          dna.Int
	VideoPortions *Portions
}

func NewVideoFinder() *VideoFinder {
	vf := new(VideoFinder)
	vf.PathIndex = 0
	vf.Page = 0
	vf.VideoPortions = &Portions{}
	return vf
}

func GetNewestVideoPortions(pathIndex, page dna.Int) (*Portions, error) {
	var baseUrl = dna.String("http://www.nhaccuatui.com")
	var ret = &Portions{}
	link := baseUrl + NewestVideoPaths[pathIndex].Replace(`.html`, "."+page.ToString()+".html")
	// dna.Log(link)
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		albumKeyArr := data.FindAllString(`<a rel="nofollow" href="http://www\.nhaccuatui.com/video.+?">`, -1)
		if albumKeyArr.Length() > 0 {
			albumKeyArr.ForEach(func(val dna.String, idx dna.Int) {
				portion := NewPortion()
				keyArr := val.GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
				if len(keyArr) > 0 {
					portion.Key = string(keyArr[0][1])
					ret.Push(portion)
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
func (vf *VideoFinder) Fetch() error {
	ret, err := GetNewestVideoPortions(vf.PathIndex, vf.Page)
	if err != nil {
		return err
	} else {
		vf.VideoPortions = ret
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (vf *VideoFinder) GetId() dna.Int {
	return vf.PathIndex
}

// New implements item.Item interface
// Returns new item.Item interface
func (vf *VideoFinder) New() item.Item {
	return item.Item(NewVideoFinder())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (vf *VideoFinder) Init(v interface{}) {
	var n dna.Int
	var NVideoPaths dna.Int = dna.Int(NewestVideoPaths.Length()) // The total of song paths
	switch v.(type) {
	case int:
		n = dna.Int(v.(int))
	case dna.Int:
		n = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
	pathIndex := dna.Int(n / TotalVideoPages)
	if pathIndex >= NVideoPaths {
		pathIndex = NVideoPaths - 1
	}
	vf.PathIndex = pathIndex
	vf.Page = n%TotalVideoPages + 1
}

func (vf *VideoFinder) Save(db *sqlpg.DB) error {
	vf.VideoPortions.UniqueByKeys()
	err := vf.VideoPortions.FilterByKeys("nctvideos", db)
	if err != nil {
		return err
	} else {
		mutex.Lock()
		*NewestVideoPortions = append(*NewestVideoPortions, *vf.VideoPortions...)
		mutex.Unlock()
		return nil
	}
}
