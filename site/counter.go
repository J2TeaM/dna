package site

import (
	. "dna"
	"time"
)

type Counter struct {
	Total        Int           // Total items if specified
	Count        Int           // Total items at a point of running
	Failure      Int           // Total items failed
	Pass         Int           // Total items done
	Speed        Int           // How many items counter processes per second
	ElapsedTime  time.Duration // ElapsedTime time.
	RemainedTime time.Duration // The time left (if total is specified)
	startingTime time.Time
}

func NewCounter() *Counter {
	return &Counter{
		Total:        0,
		Count:        0,
		Failure:      0,
		Pass:         0,
		Speed:        0,
		ElapsedTime:  0,
		RemainedTime: 0,
		startingTime: time.Now(),
	}
}

func (c *Counter) Tick(hasError Bool) {
	if hasError == true {
		c.Failure += 1
	} else {
		c.Pass += 1
	}
	c.Count += 1
	c.ElapsedTime = time.Since(c.startingTime)
	if c.ElapsedTime/time.Second > 0 {
		c.Speed = Int(int64(c.Count) / int64(c.ElapsedTime/time.Second))
	}

	if c.Total > 0 {
		c.RemainedTime = time.Duration(int64((c.Total-c.Count)/c.Speed) * int64(time.Second))
	}
}

// getTimeFmt returns default format of time.Duration into 1.12s or 1m23.12s
// Only take 2 digits after second unit
func getTimeFmt(duration time.Duration) String {
	return Sprintf("%v", duration).ReplaceWithRegexp(`(^.+\.[0-9]{2})[0-9]+(.+)`, `$1$2`)
}

func (c Counter) String() string {
	if c.Total > 0 {
		format := String("t:%d n:%v fail:%v pass:%v speed:%v total:%v remained:%v")
		return string(Sprintf(format, getTimeFmt(c.ElapsedTime), c.Count, c.Failure, c.Pass, c.Speed, c.Total, getTimeFmt(c.RemainedTime)))
	} else {
		format := String("t:%v n:%v fail:%v pass:%v speed:%v")
		return string(Sprintf(format, getTimeFmt(c.ElapsedTime), c.Count, c.Failure, c.Pass, c.Speed))
	}
	return ""
}
