package site

import (
	"dna"
	"dna/terminal"
	"time"
)

func CountDown(duration time.Duration, message, endingMess dna.String) {
	hourGlass := terminal.NewHourGlass(duration)
	hourGlass.Start()
	for hourGlass.Done == false {
		time.Sleep(time.Millisecond * 500)
		hourGlass.Show(message)
	}
	dna.Log("\n" + endingMess)
}
