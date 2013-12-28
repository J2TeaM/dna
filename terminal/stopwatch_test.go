package terminal

import (
	"time"
)

func ExampleStopWatch() {
	stopWatch := NewStopWatch()
	stopWatch.Start()
	for stopWatch.Elapsed < time.Second*5 {
		stopWatch.Show("Watch running - 🕒:")
	}
	stopWatch.Stop()
}
