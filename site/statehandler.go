package site

import (
	. "dna"
	"dna/sqlpg"
	"sync"
)

// Range defines a range from first element to last one
type Range struct {
	First Int
	Last  Int
	Total Int
}

func NewRange(first, last Int) *Range {
	total := last - first + 1
	return &Range{first, last, total}
}

// StateHandler defines the state of Update(). It ensures its fields are only called once.
//
// StateHandler resolves 3 common patterns to update new item from a site.
//
// 	* 1st general pattern: Update items from last ids of (song, album...) of a site to the newest ones.
// 	It will stop after N continuous failures which comes from SiteConfig. It means the newest ones are found.
// 	The Cid is the lasted id of item in a table.
// 	Fields used: IsOver
//
// 	* 2nd pattern: Update through range. Usually to fetch items from m to n.
// 	Ex: getting all songs from X with range from 1000 to 2000.
// 	Fields used: Cid, IsOver, Range
//
// 	* 3rd pattern: Update through an external slice.
// 	Ex: Re-getting all failed ids from log file whose ids are not in order.
// 	Or in the case of nhaccuatui, a key is encrypted and is used to display a page, so songid is hidden.
// 	Therefore no way to loop through a range. Only list of keys is found, which is translated into ids.
// 	The ids is an integer slice. So 3rd pattern will be applied.
// 	Fields used: Cid, IsOver, ExtSlice
type StateHandler struct {
	mu         sync.RWMutex
	Cid        Int         // Current ID of a page Update() is getting.
	SiteConfig *SiteConfig // Site config containing n continuous failures - NCFail (1st pattern)
	Range      *Range      // Looping through range if available (2nd pattern)
	ExtSlice   *IntArray   // Looping through an slice (3rd pattern)
	Db         *sqlpg.DB   // Connection-to-db state
	Pattern    Int         // pattern number
	IsOver     Bool        // IsOver is true when nothing to be updated
	Item       Item
	TableName  String
}

// CheckStateHandler panics if StateHandler is not in proper format.
func CheckStateHandler(state *StateHandler) {
	switch state.GetPattern() {
	case 1:
		if state.GetRange() != nil || state.GetExtSlice() != nil {
			panic("Wrong 1st pattern!")
		}
	case 2:
		if state.GetRange() == nil || state.GetExtSlice() != nil {
			panic("Wrong 2nd pattern")
		}
	case 3:
		if state.GetExtSlice() == nil || state.GetRange() != nil {
			panic("Wrong 3rd pattern")
		}
	default:
		panic("Wrong pattern number")
	}
}

// NewStateHandler returns default updates (1st pattern).
func NewStateHandler(item Item, config *SiteConfig, db *sqlpg.DB) *StateHandler {
	tableName := sqlpg.GetTableName(item)
	return &StateHandler{
		Cid:        0,
		SiteConfig: config,
		Range:      nil,
		ExtSlice:   nil,
		Db:         db,
		Pattern:    1,
		IsOver:     false,
		Item:       item,
		TableName:  tableName,
	}
}

// NewStateHandlerWithRange returns new StateHandler with 2nd pattern.
func NewStateHandlerWithRange(item Item, r *Range, config *SiteConfig, db *sqlpg.DB) *StateHandler {
	tableName := sqlpg.GetTableName(item)
	return &StateHandler{
		Cid:        r.First - 1,
		SiteConfig: config,
		Range:      r,
		ExtSlice:   nil,
		Db:         db,
		Pattern:    2,
		IsOver:     false,
		Item:       item,
		TableName:  tableName,
	}
}

// NewStateHandlerWithExtSlice returns  new StateHandler with 3rd pattern.
func NewStateHandlerWithExtSlice(item Item, extSlice *IntArray, config *SiteConfig, db *sqlpg.DB) *StateHandler {
	tableName := sqlpg.GetTableName(item)
	return &StateHandler{
		Cid:        -1, // In this case, Cid means index of current element in external slice
		SiteConfig: config,
		Range:      nil,
		ExtSlice:   extSlice,
		Db:         db,
		Pattern:    3,
		IsOver:     false,
		Item:       item,
		TableName:  tableName,
	}
}

// IncreaseCid increases Cid by a unit and returns the increased id.
func (sh *StateHandler) IncreaseCid() {
	sh.mu.Lock()
	sh.Cid += 1
	sh.mu.Unlock()
}

func (sh *StateHandler) SetCid(n Int) {
	sh.mu.Lock()
	sh.Cid = n
	sh.mu.Unlock()
}

// GetCid returns Cid.
func (sh *StateHandler) GetCid() Int {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	if sh.Pattern == 1 || sh.Pattern == 2 {
		return sh.Cid
	} else {
		idx := sh.Cid
		length := sh.ExtSlice.Length()
		if idx >= length {
			idx = length - 1
		}
		return (*sh.ExtSlice)[idx]
	}
}

// IsComplete returns the value of IsOver
func (sh *StateHandler) IsComplete() Bool {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	return sh.IsOver
}

// SetCompletion sets IsOver to be true
func (sh *StateHandler) SetCompletion() {
	sh.mu.Lock()
	sh.IsOver = true
	sh.mu.Unlock()
}

func (sh *StateHandler) GetRange() *Range {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	return sh.Range
}

func (sh *StateHandler) GetPattern() Int {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	return sh.Pattern
}

func (sh *StateHandler) GetExtSlice() *IntArray {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	return sh.ExtSlice
}

func (sh *StateHandler) GetNCFail() Int {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	switch {
	case sh.TableName.Match(`song`) == true:
		return sh.SiteConfig.NCSongFail
	case sh.TableName.Match(`album`) == true:
		return sh.SiteConfig.NCAlbumFail
	case sh.TableName.Match(`video`) == true:
		return sh.SiteConfig.NCVideoFail
	default:
		panic("Cannot find type of NCFail: it has to be song, album or video")
		return 0
	}
}

func (sh *StateHandler) GetItem() Item {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	return sh.Item
}

func (sh *StateHandler) GetTableName() String {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	return sh.TableName
}

func (sh *StateHandler) GetDb() *sqlpg.DB {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	return sh.Db
}

// InsertIgnore implements StateHandler.Db.InsertIgnore() method
func (sh *StateHandler) InsertIgnore(itm Item) {
	go func() {
		err := sh.Db.InsertIgnore(itm)
		if err != nil {
			Log("Cannot insert new item. ", err.Error())
		}
	}()
}
