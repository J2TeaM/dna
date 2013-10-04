package terminal

import (
	. "dna"
	"time"
)

func ExampleIndicator_SetFrames() {
	indicator := NewIndicator()
	indicator.ActiveColor = Blue
	indicator.SetFrames(StringArray{".", "o", "0", "o"})
	for i := 1; i <= 10; i++ {
		indicator.Show("  $indicator Getting items")
		time.Sleep(100 * time.Millisecond)
	}
	// Output is repeated frames of blue color
}

func ExampleNewIndicator() {
	indicator := NewIndicator()
	indicator.ActiveColor = Blue
	indicator.InactiveColor = Black
	indicator.Inversed = false
	indicator.ActiveChar = "★"
	indicator.InactiveChar = "☆"
	indicator.Length = 6
	for i := 1; i <= 3000; i++ {
		indicator.Show("  $indicator Getting items")
		time.Sleep(100 * time.Millisecond)
	}
	// The output is:
	//   ★★★★★☆ Getting items
	//   ★ is blue
	//   ☆ is black
}

func ExampleNewIndicatorWithTheme() {
	idc := NewIndicatorWithTheme(ThemeDefault)
	Logv(idc.ActiveColor)
	Logv(idc.InactiveColor)
	Logv(idc.ActiveChar)
	Logv(idc.InactiveChar)
	Logv(idc.Length)
	Logv(idc.Inversed)
	Logv(idc.HasFlash)
	Logv(idc.FramesPerFlash)
	// Output:
	// 2
	// 7
	// "•"
	// "•"
	// 3
	// true
	// false
	// 2
}
