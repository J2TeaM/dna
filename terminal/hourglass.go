package terminal

import (
	"dna"
	"time"
)

type HourGlass struct {
	Duration        time.Duration
	Done            dna.Bool
	startTime       time.Time
	stopTime        time.Time
	lastLap         time.Time
	lastLapDuration time.Duration
	console         *Console
}

func NewHourGlass(duration time.Duration) *HourGlass {
	hg := new(HourGlass)
	hg.Duration = duration
	hg.lastLapDuration = 0
	hg.Done = false
	hg.console = NewConsole()
	return hg
}

func (hg *HourGlass) Start() {
	hg.startTime = time.Now()
	hg.lastLap = hg.startTime
}

// lap calculates interval between two consecutive laps.
func (hg *HourGlass) lap() time.Duration {
	now := time.Now()
	duration := now.Sub(hg.lastLap)
	hg.lastLap = now
	return duration
}

func (hg *HourGlass) GetRemainingTime() time.Duration {
	sub := hg.Duration - time.Since(hg.startTime)
	if sub > 0 {
		return sub
	} else {
		hg.Done = true
		return 0
	}
}

func (hg *HourGlass) Cancel() {
	*hg = *NewHourGlass(hg.Duration)
}

func (hg *HourGlass) Show(message dna.String) {
	hg.lastLapDuration += hg.lap()
	if hg.lastLapDuration >= time.Second/2 {
		fmtClock := dna.Sprintf("%v", hg.GetRemainingTime()/time.Millisecond*time.Millisecond)
		format := dna.Sprintf("%v %v", message, fmtClock)
		hg.console.Erase(Line).Column(0)
		hg.console.Write(format)
		hg.lastLapDuration = 0
	}
}
