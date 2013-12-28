package terminal

import (
	"dna"
	"time"
)

// StopWatch defines a stopwatch type.
type StopWatch struct {
	Elapsed         time.Duration
	StartTime       time.Time
	StopTime        time.Time
	lastLap         time.Time
	lastTick        time.Time
	lastLapDuration time.Duration
	console         *Console
}

// NewStopWatch returns a new StopWatch.
func NewStopWatch() *StopWatch {
	return new(StopWatch)
}

// Start begins the running watch.
func (sw *StopWatch) Start() {
	sw.StartTime = time.Now()
	sw.lastLap = sw.StartTime
	sw.lastTick = sw.StartTime
	sw.lastLapDuration = 0
	sw.console = NewConsole()
}

// Stop terminates the running watch.
func (sw *StopWatch) Stop() {
	sw.StopTime = time.Now()
	sw.Elapsed = sw.StopTime.Sub(sw.StartTime)
	sw.console.ShowCursor()
}

// Tick calculates time elapsed while running
func (sw *StopWatch) Tick() time.Duration {
	sw.Elapsed = time.Since(sw.StartTime)
	return sw.Elapsed
}

// Lap calculates interval between two consecutive laps.
func (sw *StopWatch) Lap() time.Duration {
	now := time.Now()
	duration := now.Sub(sw.lastLap)
	sw.lastLap = now
	return duration
}

// Show displays a message if lap is greater than 500 milliseconds
func (sw *StopWatch) Show(message dna.String) {
	sw.lastLapDuration += sw.Lap()
	if sw.lastLapDuration >= time.Second/2 {
		fmtClock := dna.Sprintf("%v", sw.Tick()/time.Millisecond*time.Millisecond)
		format := dna.Sprintf("%v %v", message, fmtClock)
		sw.console.Erase(Line).Column(0)
		sw.console.Write(format).HideCursor()
		sw.lastLapDuration = 0
	}

}
