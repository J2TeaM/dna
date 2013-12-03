package terminal

import (
	. "dna"
	"fmt"
	"time"
)

// Define new theme type
type Theme Int

// Set default themes.
// Preview: https://raw.github.com/anbinh/images/master/theme.jpg
const (
	ThemeDefault Theme = iota
	ThemeSimple
	ThemeSimpleBar
	ThemeDot
	ThemeStar
)

/* ProgressBar implements general way to print progress bar to console.

The format of the theme string includes the following tokens:
	$total : show total items
	$current : show total items completed
	$elapsed : show time elapsed
	$eta : show estimated completion time in seconds
	$speed : show speed
	$percent : show completion percentage
	$custom : insert your custom data
	$[...] : progress bar itself. "..." represents custom tokens. Overlay tokens info on the bar

Ex: "Running... $[$current/$total - Elapsed Time: $elapsed]"

Notice: If bar is not overlayed (only $[]), properties such as CompleteSymbol, IncompleteSymbol will be used.
Otherwise, they will be dismissed and the bar will use whitespace + overlayed string instead.

*/
type ProgressBar struct {
	TotalItems          Int           // The number of items of progress bar. Must specific
	CurrentItemsCount   Int           // The current number of items
	Width               Int           // The progress bar width. Default 70
	CompleteSymbol      String        // The symbol represents completion character
	IncompleteSymbol    String        // The symbol represents incompletion character
	CompleteBGColor     Int           // The BG color presents complete chars. Default -1.
	IncompleteBGColor   Int           // The BG color presents Running chars. Default -1.
	CompleteTextColor   Int           // The BG color presents complete chars. Default -1.
	IncompleteTextColor Int           // The BG color presents Running chars. Default -1.
	CompleteTheme       String        // The theme string  represents when complete. Default "".
	RunningTheme        String        // The theme string  represents when running. Default "".
	elapsedTime         time.Duration // The time in seconds goes by from the beginning
	remainingTime       time.Duration // Estimated completion time in seconds
	speed               Int           // How many items passed per second
	percent             Float         // Complete percentage
	totalLines          Int           // The number of lines counted
	startingTime        time.Time     // Store the time when progress bar is constructed
	console             *Console
	firstTime           Bool   // check whether it is the first time the progress bar is running
	completeCustomData  String // data will be shown when complete. Set empty string to disable
	runningCustomData   String // data will be shown when running. Set empty string to disable
}

// Setup new progress bar with default properties
func NewProgressBar(totalItems Int, runningTheme, completeTheme String) *ProgressBar {
	progressBar := new(ProgressBar)
	progressBar.TotalItems = totalItems
	progressBar.Width = 70
	progressBar.CompleteSymbol = " "
	progressBar.IncompleteSymbol = " "
	progressBar.CompleteBGColor = Green
	progressBar.IncompleteBGColor = Black
	progressBar.CompleteTextColor = Black
	progressBar.IncompleteTextColor = -1
	progressBar.elapsedTime = 0
	progressBar.remainingTime = 0
	progressBar.speed = 0
	progressBar.startingTime = time.Now()
	if runningTheme == "" {
		progressBar.RunningTheme =
			`AB $[Running...  $percent% - t:$elapsed - t_left:$eta]
(current items:$current/$total)
Speed:$speed`
	} else {
		progressBar.RunningTheme = runningTheme
	}
	if completeTheme == "" {
		progressBar.CompleteTheme = `$[Done! - t:$elapsed - ($current/$total) - speed:$speed]`
	} else {
		progressBar.CompleteTheme = completeTheme
	}
	progressBar.totalLines = progressBar.RunningTheme.ToLines().Length()
	progressBar.console = NewConsole()
	progressBar.firstTime = true
	return progressBar
}

// Define progress bar with input theme
func NewProgressBarWithTheme(totalItems Int, theme Theme) *ProgressBar {
	progressBar := new(ProgressBar)
	progressBar.TotalItems = totalItems
	switch theme {
	case ThemeDefault:
		{
			progressBar.Width = 70
			progressBar.CompleteSymbol = " "
			progressBar.IncompleteSymbol = " "
			progressBar.CompleteBGColor = Green
			progressBar.IncompleteBGColor = White
			progressBar.CompleteTextColor = Black
			progressBar.IncompleteTextColor = Black
			var rt String = "$[Running...   $percent%   $current/$total]"
			rt += "\nElapsed $elapsed    Remaining $eta    Speed $speeditems/s"
			var ct String = `$[Done!    t:$elapsed    Total:$total    Speed:$speeditems/s]`
			progressBar.RunningTheme = rt
			progressBar.CompleteTheme = ct
		}
	case ThemeSimple:
		{
			progressBar.Width = 40
			progressBar.CompleteSymbol = "▬"
			progressBar.IncompleteSymbol = "․"
			progressBar.CompleteBGColor = -1
			progressBar.IncompleteBGColor = -1
			progressBar.CompleteTextColor = -1
			progressBar.IncompleteTextColor = -1
			progressBar.RunningTheme = `[$[]]   $percent%   $current/$total`
			var rt String = `Complete!    t:$elapsed    Total:$total    Speed:$speeditems/s`
			progressBar.CompleteTheme = rt
		}
	case ThemeSimpleBar:
		{
			progressBar.Width = 50
			progressBar.CompleteSymbol = " "
			progressBar.IncompleteSymbol = " "
			progressBar.CompleteBGColor = Blue
			progressBar.IncompleteBGColor = White
			progressBar.CompleteTextColor = Black
			progressBar.IncompleteTextColor = Black
			progressBar.RunningTheme = `$[$percent% $current/$total $elapsed]`
			progressBar.CompleteTheme = `$[Done! t:$elapsed total:$total v:$speeditems/s]`
		}
	case ThemeStar:
		{
			progressBar.Width = 40
			progressBar.CompleteSymbol = "★"
			progressBar.IncompleteSymbol = "☆"
			progressBar.CompleteBGColor = -1
			progressBar.IncompleteBGColor = -1
			progressBar.CompleteTextColor = Yellow
			progressBar.IncompleteTextColor = -1
			// getting yellow left bracket
			ylb := NewColorString("[").Yellow().Value()
			// getting right bracket yellow color
			yrb := NewColorString("]").Yellow().Value()
			rt := fmt.Sprintf("%v$[]%v   $percent%%   $current/$total", ylb, yrb)
			ct := fmt.Sprintf("%v$[]%v\n$percent%% t:$elapsed total:$total", ylb, yrb)
			progressBar.RunningTheme = String(rt)
			progressBar.CompleteTheme = String(ct)
		}
	case ThemeDot:
		{
			progressBar.Width = 40
			progressBar.CompleteSymbol = "."
			progressBar.IncompleteSymbol = "․"
			progressBar.CompleteBGColor = -1
			progressBar.IncompleteBGColor = -1
			progressBar.CompleteTextColor = Green
			progressBar.IncompleteTextColor = -1
			progressBar.RunningTheme = "$[]   $percent%   $current/$total"
			progressBar.CompleteTheme = "$[]\n$percent% t:$elapsed total:$total"
		}
	}
	progressBar.elapsedTime = 0
	progressBar.remainingTime = 0
	progressBar.speed = 0
	progressBar.startingTime = time.Now()
	progressBar.totalLines = progressBar.RunningTheme.ToLines().Length()
	progressBar.console = NewConsole()
	progressBar.firstTime = true
	return progressBar
}

func (b *ProgressBar) getTime(seconds Int) String {
	return seconds.ToString()
}

func (b *ProgressBar) formatDuration(d time.Duration) String {
	totalSecs := int(d.Seconds())
	hours := totalSecs / 3600
	totalSecs = totalSecs % 3600
	minutes := totalSecs / 60
	seconds := totalSecs % 60
	if hours > 0 {
		return String(fmt.Sprintf("%02dh:%02dm:%02ds", hours, minutes, seconds))
	} else {
		if minutes > 0 {
			return String(fmt.Sprintf("%02dm:%02ds", minutes, seconds))
		} else {
			return String(fmt.Sprintf("%02ds", seconds))
		}
	}
}

func (b *ProgressBar) calculateElapsedTime() {
	b.elapsedTime = time.Now().Sub(b.startingTime)
}

func (b *ProgressBar) calculatePercent() {
	b.percent = b.CurrentItemsCount.ToFloat() / b.TotalItems.ToFloat() * 100
}

func (b *ProgressBar) calculateSpeed() {
	if b.elapsedTime.Nanoseconds() > 0 {
		x := b.CurrentItemsCount.ToFloat() / Float(b.elapsedTime.Seconds())
		b.speed = x.Ceil()
	} else {
		b.speed = 0
	}
}

func (b *ProgressBar) calculateRemainingTime() {
	if b.speed > 0 {
		x := (b.TotalItems.ToFloat() - b.CurrentItemsCount.ToFloat()) / b.speed.ToFloat()
		b.remainingTime = time.Duration(x * Float(time.Second))
	} else {
		b.remainingTime = 10000000
	}

}

func (b *ProgressBar) isComplete() Bool {
	if b.TotalItems == b.CurrentItemsCount {
		return true
	} else {
		return false
	}
}

func (b *ProgressBar) isFirstTimeRunning() Bool {
	if b.firstTime == true {
		b.firstTime = false
		return true
	} else {
		return false
	}
}

func (b *ProgressBar) hasBar(line String) Bool {
	return line.Match("\\$\\[.*?\\]")
}

func (b *ProgressBar) getParsedTokens(tokens String) String {
	var result String
	result = tokens
	result = result.ReplaceWithRegexp("\\$total", b.TotalItems.ToCommaFormat())
	result = result.ReplaceWithRegexp("\\$current", b.CurrentItemsCount.ToCommaFormat())
	result = result.ReplaceWithRegexp("\\$elapsed", b.formatDuration(b.elapsedTime))
	result = result.ReplaceWithRegexp("\\$eta", b.formatDuration(b.remainingTime))
	result = result.ReplaceWithRegexp("\\$speed", b.speed.ToCommaFormat())
	result = result.ReplaceWithRegexp("\\$percent", b.percent.ToFormattedString(1, 2))
	if b.isComplete() {
		result = result.ReplaceWithRegexp("\\$custom", b.completeCustomData)
	} else {
		result = result.ReplaceWithRegexp("\\$custom", b.runningCustomData)
	}
	return result
}

func (b *ProgressBar) getDrawingBar(overlayedString String) String {
	var result String
	var incompletePart, completePart *ColorString
	if overlayedString.IsBlank() {
		bgBarWidth := Float(b.percent.ToPrimitiveValue() / 100 * float64(b.Width.ToPrimitiveType())).Floor()
		completePart = NewColorString(b.CompleteSymbol.Repeat(bgBarWidth))
		incompletePart = NewColorString(b.IncompleteSymbol.Repeat(b.Width - bgBarWidth))
	} else {
		barLeft := b.Width - overlayedString.Length()
		var left Int
		if barLeft.IsEven() {
			left = barLeft / 2
		} else {
			left = barLeft/2 + 1
		}
		right := barLeft / 2
		finalString := String(" ").Repeat(left) + overlayedString + String(" ").Repeat(right)
		bgBarWidth := Float(b.percent.ToPrimitiveValue() / 100 * float64(b.Width.ToPrimitiveType())).Floor()
		completePart = NewColorString(finalString.Substring(0, bgBarWidth))
		incompletePart = NewColorString(finalString.Substring(bgBarWidth, b.Width))
	}
	result += completePart.Background(b.CompleteBGColor).Color(b.CompleteTextColor).Value()
	result += incompletePart.Background(b.IncompleteBGColor).Color(b.IncompleteTextColor).Value()

	return result
}

func (b *ProgressBar) getParsedCustomData(data String) String {
	if b.isComplete() {
		return data.ReplaceWithRegexp("\\$custom", b.completeCustomData)
	} else {
		return data.ReplaceWithRegexp("\\$custom", b.runningCustomData)
	}
}

func (b *ProgressBar) parseTheme() {
	// Broken theme string into lines
	// Loop through line to find tokens
	var lines StringArray
	if b.isComplete() {
		lines = b.CompleteTheme.ToLines()
	} else {
		lines = b.RunningTheme.ToLines()
	}

	for _, line := range lines {
		if b.hasBar(line) {
			// getting all tokens from bar
			barTokens := line.FindAllStringSubmatch("\\$\\[(.*?)\\]", 1)[0][1]
			parsedTokensString := b.getParsedTokens(barTokens)
			// render all bar to color string
			drawingBar := b.getDrawingBar(parsedTokensString)
			// replace the bar token with rendered bar
			content := line.ReplaceWithRegexp("\\$\\[.*?\\]", drawingBar)
			// parse the rest of line (outside bar token)
			// add custom data and write to console
			b.console.Erase(Line).Write(b.getParsedCustomData(b.getParsedTokens(content)))
		} else {
			parsedTokensLine := b.getParsedTokens(line)
			b.console.Erase(Line).Write(b.getParsedCustomData(parsedTokensLine))
		}
		b.console.Down(1).Column(1).HideCursor()
	}

	if b.isComplete() {
		b.console.ShowCursor()
	}

}

func (b *ProgressBar) calculateStats() {
	b.calculateElapsedTime()
	b.calculatePercent()
	b.calculateSpeed()
	b.calculateRemainingTime()
}

// Show runs the progress bar.
// The first param is the number of current items in comparison with total items.
// The second and the last is custom data passed at runtime. Set empty strong ("") to neglect
// The custom data is used with token $custom in theme string.
func (b *ProgressBar) Show(currentItemsCount Int, runningCData, completeCData String) {
	b.CurrentItemsCount = currentItemsCount
	b.runningCustomData = runningCData
	b.completeCustomData = completeCData
	b.calculateStats()
	if b.isFirstTimeRunning() == false {
		b.console.Up(b.totalLines).Erase(Line).Column(1)
	}
	b.parseTheme()

}
