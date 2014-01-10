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
	console := terminal.NewConsole()
	console.Erase(terminal.Line).Column(0)
	console.Write(endingMess)
	console.Write("\n")
}
