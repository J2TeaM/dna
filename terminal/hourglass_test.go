package terminal

import (
	"time"
)

func ExampleHourGlass() {
	hourGlass := NewHourGlass(5 * time.Second)
	hourGlass.Start()
	for hourGlass.Done == false {
		hourGlass.Show("Watch countdown - 🕒:")
	}
}
