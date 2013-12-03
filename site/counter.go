package site

import (
	. "dna"
	"sync"
	"time"
)

// Counter defines counter struct
//
// Notice: CFail uses only with 1st pattern of StateHandler. It's reset to 0 once next item passed.
type Counter struct {
	mu           sync.RWMutex
	Total        Int           // Total items if specified
	Count        Int           // Total items at a point of running
	Fail         Int           // Total items failed
	CFail        Int           // Current continuous items failed.
	Pass         Int           // Total items done
	Speed        Int           // How many items counter processes per second
	ElapsedTime  time.Duration // ElapsedTime time.
	startingTime time.Time
}

func NewCounter(state *StateHandler) *Counter {
	var total Int = 0
	switch state.GetPattern() {
	case 2:
		total = state.GetRange().Total
	case 3:
		total = state.GetExtSlice().Length()
	}
	return &Counter{
		Total:        total,
		Count:        0,
		Fail:         0,
		CFail:        0,
		Pass:         0,
		Speed:        0,
		ElapsedTime:  0,
		startingTime: time.Now(),
	}
}

// Tick changes its values when an item is being processed.
//
// 	* hasError: determines whether a processed item is successful or not.
func (c *Counter) Tick(hasError Bool) {
	c.mu.Lock()
	if hasError == true {
		c.Fail += 1
		c.CFail += 1
	} else {
		c.Pass += 1
		c.CFail = 0
	}
	c.Count += 1
	c.ElapsedTime = time.Since(c.startingTime)
	if c.ElapsedTime/time.Second > 0 {
		c.Speed = Int(int64(c.Count) / int64(c.ElapsedTime/time.Second))
	}

	c.mu.Unlock()
}

// getTimeFmt returns default format of time.Duration into 1.12s or 1m23.12s
// Only take 2 digits after second unit
func getTimeFmt(duration time.Duration) String {
	return Sprintf("%v", duration).ReplaceWithRegexp(`(^.+\.[0-9]{2})[0-9]+(.+)`, `$1$2`)
}

func (c *Counter) GetCount() Int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Count
}

func (c *Counter) GetCFail() Int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.CFail
}

func (c Counter) String() string {
	c.mu.RLock()
	if c.Total > 0 {
		format := String("%v✘ | %v✔")
		return string(Sprintf(format, c.Fail, c.Pass))
	} else {
		format := String("t:%v | n:%v | fail:%v | pass:%v | speed:%v")
		return string(Sprintf(format, getTimeFmt(c.ElapsedTime), c.Count, c.Fail, c.Pass, c.Speed))
	}
	c.mu.RUnlock()
	return ""
}

// FinalString prints the last result of a counter
func (c *Counter) FinalString() string {
	format := String("N:%v | t=%v | (%v✘ - %v✔) | ν=%vitems/s")
	return string(Sprintf(format, c.Count, getTimeFmt(c.ElapsedTime), c.Fail, c.Pass, c.Speed))
}
