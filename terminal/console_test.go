package terminal

func ExampleConsole() {
	var console *Console = NewConsole()
	console.Right(3).Write("Default String").Down(1).Foreground(Green).Write("Green text").Display(ResetCode)
	console.Down(1).Left(100).Write("Back to normal string").Background(Blue).Foreground(Yellow).Write("Yellow text with blue background")
	console.Display(ResetCode).Down(1).Left(100).Display(ReverseCode).Write("Swap foreground and background")
	console.Display(ResetCode).Down(1).Left(100).Write("Before line reset").Erase(Line).Left(100).Write("After line reset")
	console.Display(ResetCode).Erase(Line).Column(1).Write("Final line.")
	//Output is:"   Default String"
	//"                 Green text"
	//"Back to normal stringYellow text with blue background"
	//"Swap forground and background"
	//"Final line."
}

func ExampleConsole_Display() {
	var console *Console = NewConsole()
	console.Display(ResetCode)
	console.Display(ReverseCode)
}

func ExampleConsole_Foreground() {
	var console *Console = NewConsole()
	console.Foreground(Red)
	console.Foreground(Blue)
}

func ExampleConsole_Background() {
	var console *Console = NewConsole()
	console.Background(Red)
	console.Background(Blue)
}

func ExampleConsole_Erase() {
	var console *Console = NewConsole()
	console.Erase(Line)
	console.Erase(Screen)
}
