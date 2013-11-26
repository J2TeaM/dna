package terminal

import (
	. "dna"
)

// Using with token:
// 	$indicator
// Example: The format look like: " $indicator Getting items".
//
// Normal order is only from left to right. The inversed option sets from left to right and from right to left.
//
// Some special chars: ● ○ • ・ ◎ ◉ ⦿ ■ □ ▪ ▫ - ⁃ ✦ ✷ ✔ ⁍ ➤ ❯ ✤ ✧ ★ ☆ ✯ ✡ ✩ ▸ ▹ ▶ ▷ might be used as active or inactive chars.
type Indicator struct {
	Frames         StringArray // String array detemines all the frames looped through
	ActiveColor    Int         // Color of characters at active status. Default -1.
	InactiveColor  Int         // Color of characters at inactive status. Default -1.
	ActiveChar     String      // A single character represents indicator when active
	InactiveChar   String      // A single character represents indicator when inactive
	Length         Int         // Length of repeated characters
	Inversed       Bool        // Set inversed order options
	HasFlash       Bool        // Enable flash options
	FramesPerFlash Int         // The number of frames per flash
	firstTime      Bool
	counter        Int
	console        *Console
}

func NewIndicator() *Indicator {
	idc := new(Indicator)
	idc.counter = 0
	idc.console = NewConsole()
	idc.ActiveColor = Green
	idc.InactiveColor = White
	idc.ActiveChar = "•"
	idc.InactiveChar = "•"
	idc.Length = 4
	idc.Inversed = true
	idc.HasFlash = true
	idc.FramesPerFlash = idc.Length / 2
	idc.Frames = StringArray{}

	idc.firstTime = true
	return idc
}

// Some default indicators:
// 	ThemeDefault: [•••] green & white dots with no flash & inversed order
// 	ThemeSimple: \|-/ rotating
// 	ThemeSimpleBar: [••••] white & black dots with flash & no inversed order
// 	ThemeDot: .... Blue and white dots
// 	ThemeStar: [★★★★] green & black stars
func NewIndicatorWithTheme(theme Theme) *Indicator {
	idc := NewIndicator()
	switch theme {
	case ThemeDefault:
		idc.Length = 3
		idc.HasFlash = false
		idc.makeFrames()
		idc.Frames = StringArray(idc.Frames.Map(func(v String, idx Int) String {
			return String("[") + v + String("]")
		}).([]String))
	case ThemeSimple:
		idc.ActiveColor = White
		idc.InactiveColor = Black
		idc.Frames = StringArray{"/", "-", "\\", "|"}
	case ThemeSimpleBar:
		idc.ActiveColor = White
		idc.InactiveColor = Black
		idc.makeFrames()
		idc.Frames = StringArray(idc.Frames.Map(func(v String, idx Int) String {
			return String("[") + v + String("]")
		}).([]String))
	case ThemeDot:
		idc.ActiveColor = Blue
		idc.InactiveColor = White
		idc.ActiveChar = "▪"
		idc.InactiveChar = "▪"
	case ThemeStar:
		idc.ActiveColor = Green
		idc.InactiveColor = Black
		idc.Inversed = false
		idc.ActiveChar = "★"
		idc.InactiveChar = "☆"
		idc.HasFlash = false
		idc.makeFrames()
		idc.Frames = StringArray(idc.Frames.Map(func(v String, idx Int) String {
			return String("[") + v + String("]")
		}).([]String))
	}
	return idc
}

func (idc *Indicator) makeFrames() {
	var str String
	var length Int = idc.Length
	var frames, frames2 StringArray
	for i := Int(0); i < length+1; i++ {
		str = NewColorString(idc.ActiveChar.Repeat(i)).Color(idc.ActiveColor).Value()
		str += NewColorString(idc.InactiveChar.Repeat(length - i)).Color(idc.InactiveColor).Bold().Value()
		frames.Push(str)
		if idc.HasFlash {
			if i.IsDivisibleBy(idc.FramesPerFlash) {
				// adding flash
				frames.Push(String(" ").Repeat(length * idc.ActiveChar.Length()))
			}
		}

		if idc.Inversed {
			str = NewColorString(idc.ActiveChar.Repeat(i)).Color(idc.InactiveColor).Bold().Value()
			str += NewColorString(idc.InactiveChar.Repeat(length - i)).Color(idc.ActiveColor).Value()
			frames2.Push(str)
			if idc.HasFlash {
				if i.IsDivisibleBy(idc.FramesPerFlash) {
					// adding flash
					frames2.Push(String(" ").Repeat(length * idc.ActiveChar.Length()))
				}
			}
		}
	}
	idc.Frames = idc.Frames.Concat(frames)
	idc.Frames.Push(String(" ").Repeat(length * idc.ActiveChar.Length()))
	if idc.Inversed {
		idc.Frames = idc.Frames.Concat(frames2.Reverse())
		idc.Frames.Push(String(" ").Repeat(length * idc.ActiveChar.Length()))
	}
}

// Build your own custom frames.
// The InactiveColor, Char , Length  Inversed properties will be disable.
// Only AcitiveColor will be used and color will be set for each frame
func (idc *Indicator) SetFrames(frames StringArray) {
	idc.Frames = StringArray(frames.Map(func(v String, idx Int) String {
		return NewColorString(v).Color(idc.ActiveColor).Value()
	}).([]String))
	idc.Length = frames.Length()
}

func (idc *Indicator) isFirstTimeRunning() Bool {
	if idc.firstTime == true {
		idc.firstTime = false
		return true
	} else {
		return false
	}
}

// Show prints indicator to console. Theme is specified by token $indicator
func (idc *Indicator) Show(formatString String) {
	if idc.isFirstTimeRunning() == false {
		idc.console.Erase(Line).Column(1)
	} else {
		if idc.Frames.Length() == 0 {
			idc.makeFrames()
		}
	}
	renderedText := formatString.Replace("$indicator", idc.Frames[idc.counter%idc.Frames.Length()])
	idc.console.Write(renderedText)
	idc.counter += 1

}

//Close prints the final result.
func (idc *Indicator) Close(formatString String) {
	idc.console.Erase(Line).Column(1)
	str := NewColorString("[" + idc.ActiveChar.Repeat(idc.Length) + "]").Color(idc.ActiveColor).Value()
	renderedText := formatString.Replace("$indicator", str)
	idc.console.Log(renderedText)
}
