package nct

import (
	"dna"
	"dna/sqlpg"
	"dna/utils"
	"errors"
)

// Portion defines only basic fields : Id, Key, IsOfficial of Song, Album, Video.
// IsOfficial only applied to Song and correspondent to Official field in Song type.
// It is a subset of basic typse but it is a result when getting relevant or similar songs, albums or videos
// from a page.
type Portion struct {
	Id  int32
	Key string
}

//NewPortion returns a new portion
func NewPortion() *Portion {
	portion := new(Portion)
	portion.Id = 0
	portion.Key = ""
	return portion
}

// Portions defines a list of portions
type Portions []Portion

func (pl *Portions) Reset() {
	tmp := &Portions{}
	*pl = *tmp
	// for i := 0; i < len(*pl); i++ {
	// 	(*pl) = append((*pl)[:i], (*pl)[i+1:]...)
	// }
}

// GetIds returns a slice of ids from a list of portions
func (pl *Portions) GetIds() *dna.IntArray {
	ret := &dna.IntArray{}
	for _, portion := range *pl {
		ret.Push(dna.Int(portion.Id))
	}
	return ret
}

// GetIds returns a slice of ids from a list of portions
func (pl *Portions) GetKeys() *dna.StringArray {
	ret := &dna.StringArray{}
	for _, portion := range *pl {
		ret.Push(dna.String(portion.Key))
	}
	return ret
}

// GetKeysFromIds returns a list of correspondent keys from an id slice
// func (pl *Portions) GetKeysFromIds(ids *dna.IntArray) *dna.StringArray {
// 	ret := &dna.StringArray{}
// 	for _, id := range *ids {
// 		for _, portion := range *pl {
// 			if id == portion.Id {
// 				ret.Push(portion.Key)
// 			}
// 		}
// 	}
// 	return ret
// }

// GetPortion returns a portion from an id.
// Ok determines whether a portion is found or not.
func (pl *Portions) GetPortion(id dna.Int) (*Portion, error) {
	for _, portion := range *pl {
		if portion.Id == int32(id) {
			return &portion, nil
		}
	}
	return nil, errors.New("No portion found")
}

// Push adds a new potion to a portion list
func (pl *Portions) Push(portion *Portion) {
	// mutex.Lock()
	*pl = append(*pl, *portion)
	// mutex.Unlock()
}

func (pl *Portions) Delete(i dna.Int) {
	*pl = append((*pl)[:i], (*pl)[i+1:]...)
}

// FilterByIds gets a new portion list that ids are not in a specified table.
func (pl *Portions) FilterByIds(tblName dna.String, db *sqlpg.DB) error {
	// mutex.Lock()
	// defer mutex.Unlock()
	if pl.Length() > 0 {
		ids, err := utils.SelectMissingIds(tblName, pl.GetIds(), db)
		if err != nil {
			return err
		}
		ret := &Portions{}
		for _, id := range *ids {
			portion, err := pl.GetPortion(id)
			if err == nil {
				// only push found portion
				ret.Push(portion)
			}
		}
		// pl.Reset()
		*pl = *ret
		return nil
	} else {
		return nil
	}

}

// FilterByKeys gets a new portion list that keys are not in a specified table.
func (pl *Portions) FilterByKeys(tblName dna.String, db *sqlpg.DB) error {
	// mutex.Lock()
	// defer mutex.Unlock()
	if pl.Length() > 0 {
		keys, err := utils.SelectMissingKeys(tblName, pl.GetKeys(), db)
		if err != nil {
			return err
		}
		ret := &Portions{}
		if keys != nil {
			for _, key := range *keys {
				portion := NewPortion()
				portion.Key = string(key)
				ret.Push(portion)
			}
		}
		// pl.Reset()
		*pl = *ret
		return nil
	} else {
		return nil
	}
}

// UniqueByKeys only gets unique value from Portions by keys
func (pl *Portions) UniqueByKeys() {
	// mutex.Lock()
	// defer mutex.Unlock()
	ret := &Portions{}
	tmp := (*pl.GetKeys()).Unique()
	for _, key := range tmp {
		pt := NewPortion()
		pt.Key = string(key)
		ret.Push(pt)
	}
	// pl.Reset()
	*pl = *ret
}

// UniqueByIds only gets unique value from Portions by ids
func (pl *Portions) UniqueByIds() {
	// mutex.Lock()
	// defer mutex.Unlock()
	ret := &Portions{}
	tmp := (*pl.GetIds()).Unique()
	for _, id := range tmp {
		portion, err := pl.GetPortion(id)
		if err == nil {
			ret.Push(portion)
		}
	}
	// pl.Reset()
	*pl = *ret
}

func (pl *Portions) Length() dna.Int {
	return dna.Int(len(*pl))
}

// GetRelevantPortions returns relevant songs, albums or videos
func GetRelevantPortions(data *dna.String) {
	if EnableRelevantPortionsMode {
		var body *dna.String
		// This is is supposed to get relevant or similar songs'
		// var body *dna.String
		songPart := data.FindAllString(`(?mis)^.+BXH Bài Hát`, 1)
		if songPart.Length() > 0 {
			body = &songPart[0]
		} else {
			body = data
		}
		similarSongs := body.FindAllString(`(?mis)<div class="info_data">.+?0</span>`, -1)
		similarSongs.ForEach(func(val dna.String, idx dna.Int) {
			portion := NewPortion()
			keyArr := val.FindAllStringSubmatch(`\.([0-9a-zA-Z]+)\.html" class="name_song"`, 1)
			idArr := val.FindAllStringSubmatch(`"NCTCounter_sg_([0-9]+)`, 1)
			if len(keyArr) > 0 {
				portion.Key = keyArr[0][1].String()
			}
			if len(idArr) > 0 {
				portion.Id = int32(idArr[0][1].ToInt())
			}
			RelevantSongs.Push(portion)
		})

		similarSongs = data.FindAllString(`(?mis)<div class="info_song">.+?0</span>`, -1)
		similarSongs.ForEach(func(val dna.String, idx dna.Int) {
			portion := NewPortion()
			keyArr := val.FindAllStringSubmatch(`\.([0-9a-zA-Z]+)\.html" class="name_song"`, 1)
			idArr := val.FindAllStringSubmatch(`"NCTCounter_sg_([0-9]+)`, 1)
			if len(keyArr) > 0 {
				portion.Key = keyArr[0][1].String()
			}
			if len(idArr) > 0 {
				portion.Id = int32(idArr[0][1].ToInt())
			}
			RelevantSongs.Push(portion)
		})

		RelevantSongs.UniqueByIds()
		// }

		// This part to find similar albums

		albumPart := data.FindAllString(`(?mis)^.+BXH Playlist`, 1)
		if albumPart.Length() > 0 {
			body = &albumPart[0]
		} else {
			body = data
		}
		albumKeyArr := body.FindAllString(`<a href="http://www.nhaccuatui.com/playlist/.+html.+`, -1)
		albumKeys := dna.StringArray(albumKeyArr.Map(func(val dna.String, idx dna.Int) dna.String {
			keyArr := val.GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
			if len(keyArr) > 0 {
				return keyArr[0][1]
			} else {
				return ""
			}
		}).([]dna.String)).Unique()
		for _, key := range albumKeys {
			portion := NewPortion()
			portion.Key = key.String()
			RelevantAlbums.Push(portion)
		}

		// This part to find similar videos
		videoPart := data.FindAllString(`(?mis)^.+BXH Video`, 1)
		if videoPart.Length() > 0 {
			body = &videoPart[0]
		} else {
			body = data
		}
		videoKeyArr := body.FindAllString(`<a href="http://www.nhaccuatui.com/video/.+html.+`, -1)
		videoKeys := dna.StringArray(videoKeyArr.Map(func(val dna.String, idx dna.Int) dna.String {
			keyArr := val.GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
			if len(keyArr) > 0 {
				return keyArr[0][1]
			} else {
				return ""
			}
		}).([]dna.String)).Unique().Filter(func(val dna.String, idx dna.Int) dna.Bool {
			if val != "com/video/top-20" {
				return true
			} else {
				return false
			}
		})
		for _, key := range videoKeys {
			portion := NewPortion()
			portion.Key = key.String()
			RelevantVideos.Push(portion)
		}

	}

}
