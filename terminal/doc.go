/*
Package terminal implements the library relating to console, color string. It prints colors, support position of cursor.
It utilises VT100-compatible escape sequences.

It can be chained when initalize variable typed Console. Example:
	console  := NewConsole()
	console.Right(3).Write("Default String")
	// Take cursor to the right 3 spaces and write "Default String" to console
	console.Down(1).Foreground(Green).Write("The second text")
	// Go down 1 line and write "The second text" in green color
	console.Display(ResetCode)
	// Reset all attributes so next text printed on screen does not have its green color
	console.Write("Before line reset").Erase(Line).Write("After line reset")
	// Only "After line reset" string is printed

Usage of ColoString. Example:
	var console *Console = NewConsole()
	// Define new console
	console.Write(NewColorString("Hello world").Green().Inverse().Value()).Down(1)
	// Print "Hello world" with green background and black text
	console.Write(NewColorString("Hello world").SetTextColor(Red).YellowBackground().Value())
	// Print "Hello world" with yellow background and red text
	console.Write(NewColorString("Hello world").SetTextColor(Cyan).SetAttribute(ReverseCode).Value())
	// Print "Hello world" with cyan background and black text

Writing custom progress bar:
	var runningTheme String = `Downloading... 	[$[$percent% - t:$elapsed - t_left:$eta]]
		Current items:$current/$total - $custom
		Speed:$speed`
	var completeTheme String = "[$[]]\n$percent% - $custom"
	bar := NewProgressBar(245, runningTheme, completeTheme)
	bar.CompleteSymbol = "★"
	bar.IncompleteSymbol = "☆"
	bar.CompleteBGColor = -1
	bar.IncompleteBGColor = -1
	bar.CompleteTextColor = Green
	bar.IncompleteTextColor = -1
	bar.Width = 40
	for i := 1; i <= 245; i++ {
		bar.Show(Int(i), Int(100).Rand().ToString()+"random data", "Completion custom data")
		time.Sleep(time.Duration(Int(100).Rand().ToPrimitiveValue()) * time.Millisecond)
	}

Or using one of defaul theme:
	bar := NewProgressBarWithTheme(245, ThemeDot)
	for i := 1; i <= 245; i++ {
		bar.Show(Int(i), "", "")
		time.Sleep(time.Duration(Int(100).Rand().ToPrimitiveValue()) * time.Millisecond)
	}

*/
package terminal
