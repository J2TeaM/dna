package site

import (
	. "dna"
	"dna/sqlpg"
	"sync"
)

// StateHandler defines the state of Update(). It ensures its fields are only called once.
type StateHandler struct {
	mu           sync.Mutex
	CurrentId    Int       // Current ID of a page Update() is getting.
	Done         Bool      // Done is true when there is nothing to update.
	FailureCount Int       // Total continous failures during update. It's reset to 0 once next item passed.
	Db           *sqlpg.DB // Connection state
}

func NewStateHandler(id Int, db *sqlpg.DB) *StateHandler {
	return &StateHandler{
		CurrentId:    id,
		Done:         false,
		FailureCount: 0,
		Db:           db,
	}
}

// IncreaseId increases CurrentId by an amount specified and returns the increased id.
func (sh *StateHandler) IncreaseId(amount Int) Int {
	sh.mu.Lock()
	sh.CurrentId += amount
	sh.mu.Unlock()
	return sh.CurrentId
}

// UpdateFailure return one-unit incremental Failure and return new increased value.
func (sh *StateHandler) UpdateFailure(hasError Bool) Int {
	sh.mu.Lock()
	if hasError == true {
		sh.FailureCount += 1
	} else {
		sh.FailureCount = 0
	}

	sh.mu.Unlock()
	return sh.FailureCount
}

// SetDone sets Done field by bool value
func (sh *StateHandler) SetDone(v Bool) {
	sh.mu.Lock()
	sh.Done = v
	sh.mu.Unlock()
}

// InsertIgnore implements StateHandler.Db.InsertIgnore() method
func (sh *StateHandler) InsertIgnore(itm Item) {
	go func() {
		err := sh.Db.InsertIgnore(itm)
		if err != nil {
			Log(err.Error())
		}
	}()
}
