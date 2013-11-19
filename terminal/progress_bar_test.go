package terminal

import (
	. "dna"
	"time"
)

func ExampleNewProgressBar() {
	var runningTheme String = `Downloading... 	[$[$percent% - t:$elapsed - t_left:$eta]]`
	var completeTheme String = "[$[]]\n$percent% - $custom"
	bar := NewProgressBar(245, runningTheme, completeTheme)
	Logv(bar.TotalItems)          // 245
	Logv(bar.CurrentItemsCount)   // Not run yet, value 0
	Logv(bar.Width)               // Default 70
	Logv(bar.CompleteSymbol)      // Whitespace
	Logv(bar.IncompleteSymbol)    // Whitespace
	Logv(bar.CompleteBGColor)     // Green color, value 2
	Logv(bar.IncompleteBGColor)   // Black color, value 0
	Logv(bar.CompleteTextColor)   // Black color, value 0
	Logv(bar.IncompleteTextColor) // Use default console setting, value -1
	Logv(bar.CompleteTheme)
	Logv(bar.RunningTheme)
	// Output:
	// 245
	// 0
	// 70
	// " "
	// " "
	// 2
	// 0
	// 0
	// -1
	// "[$[]]\n$percent% - $custom"
	// "Downloading... \t[$[$percent% - t:$elapsed - t_left:$eta]]"
}

func ExampleNewProgressBarWithTheme() {
	bar := NewProgressBarWithTheme(245, ThemeDot)
	Logv(bar.TotalItems)          // 245
	Logv(bar.CurrentItemsCount)   // Not run yet, value 0
	Logv(bar.Width)               // Default 40
	Logv(bar.CompleteSymbol)      // dot (.)
	Logv(bar.IncompleteSymbol)    // dot (.)
	Logv(bar.CompleteBGColor)     // Use default console setting, value -1
	Logv(bar.IncompleteBGColor)   // Use default console setting, value -1
	Logv(bar.CompleteTextColor)   // Green color, value 2
	Logv(bar.IncompleteTextColor) // Use default console setting, value -1
	Logv(bar.CompleteTheme)
	Logv(bar.RunningTheme)
	// Output:
	// 245
	// 0
	// 40
	// "."
	// "․"
	// -1
	// -1
	// 2
	// -1
	// "$[]\n$percent% t:$elapsed total:$total"
	// "$[]   $percent%   $current/$total"
}

func ExampleProgressBar_Show() {
	var runningTheme String = `Downloading... 	[$[$percent% - t:$elapsed - t_left:$eta]]
		Current items:$current/$total - Speed:$speed - $custom`
	var completeTheme String = "[$[]]\n$percent% - $custom"
	bar := NewProgressBar(434, runningTheme, completeTheme)
	bar.CompleteSymbol = "★"
	bar.IncompleteSymbol = "☆"
	bar.CompleteBGColor = -1
	bar.IncompleteBGColor = -1
	bar.CompleteTextColor = -1
	bar.IncompleteTextColor = -1
	bar.Width = 40
	for i := 1; i <= 34; i++ {
		bar.Show(Int(i), Int(100).Rand().ToString()+"random data", "Completion custom data")
		time.Sleep(time.Duration(Int(100).Rand().ToPrimitiveType()) * time.Millisecond)
	}
	// Progress bar appearance while running:
	// Downloading... 	[       55.07% - t:02s - t_left:02s      ]
	// 			Current items:239/434 - Speed:97 - 94random data
	// Complete bar is:
	// [★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★]
	// 100.00% - Completion custom data

}
