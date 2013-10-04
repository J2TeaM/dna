package terminal

import (
	. "dna"
)

const (
	_ESC  = "\033"
	_CSI  = _ESC + "["
	clear = _CSI + "2J"
)

type Console struct{}

func NewConsole() *Console {
	return &Console{}
}

// Write prints str to terminal
func (c *Console) Write(str String) *Console {
	Print(str)
	return c
}

// Log prints str to terminal with newline
func (c *Console) Log(str String) *Console {
	Print(str)
	Print("\n")
	return c
}

// Up moves cursor up by n columns
func (c *Console) Up(n Int) *Console {
	Print(_CSI + n.ToString() + "A")
	return c
}

// Down moves cursor down by n columns
func (c *Console) Down(n Int) *Console {
	Print(_CSI + n.ToString() + "B")
	return c
}

// Right moves cursor right by n columns
func (c *Console) Right(n Int) *Console {
	Print(_CSI + n.ToString() + "C")
	return c
}

// Left moves cursor left by n columns
func (c *Console) Left(n Int) *Console {
	Print(_CSI + n.ToString() + "D")
	return c
}

// Move moves cursor to line x, column y
func (c *Console) Move(x, y Int) *Console {
	Print(_CSI + x.ToString() + ";" + y.ToString() + "H")
	return c
}

// Column moves cursor to column n
func (c *Console) Column(n Int) *Console {
	Print(_CSI + n.ToString() + "G")
	return c
}

// Region constants. It is used as parameter of method Erase
const (
	EndLine   = iota // Erase from the cursor to the end of the line
	StartLine        // Erase from the cursor to the start of the line
	Line             // Erase the current line
	DownLine         // Erase everything below the current line
	UpLine           // Erase everything above the current line
	Screen           // Erase the entire screen
)

func (c *Console) Erase(region Int) *Console {
	switch region {
	case 0:
		Print(_CSI + "K")
	case 1:
		Print(_CSI + "1K")
	case 2:
		Print(_CSI + "2K")
	case 3:
		Print(_CSI + "J")
	case 4:
		Print(_CSI + "1J")
	case 5:
		Print(_CSI + "1J")
	default:
		Print(_CSI + "2K")
	}
	return c
}

// Colour attributes. It is used with method Display
const (
	ResetCode  = iota // All attributes off
	BrightCode        // Bold or increase intensity
	DimCode           // Faint (decreased intensity)
	_SKIPPED3
	UnderscoreCode // Underline
	BlinkCode      // Blink
	_SKIPPED6
	ReverseCode // Inverse or reverse; swap foreground and background
	HiddenCode  // Not widely supported
)

// Display sets mode to print from color attribute constants
func (c *Console) Display(attr Int) *Console {
	Print(_CSI + attr.ToString() + "m")
	return c
}

// Colours. It is used with methods Foreground and Background
const (
	Black   = iota // Black color
	Red            // Red color
	Green          // Green color
	Yellow         // Yellow color
	Blue           // Blue color
	Magenta        // Magenta color
	Cyan           // Cyan color
	White          // White color
)

// Foreground sets foreground color
func (c *Console) Foreground(color Int) *Console {
	var code Int
	if color >= 0 && color <= 7 {
		code = color + 30
	} else {
		panic("Color foreground input invalid")
		code = 30
	}
	Print(_CSI + code.ToString() + "m")
	return c
}

// Background set background color
func (c *Console) Background(color Int) *Console {
	var code Int
	if color >= 0 && color <= 7 {
		code = color + 40
	} else {
		panic("Color background input invalid")
		code = 40
	}
	Print(_CSI + code.ToString() + "m")
	return c
}

// HideCursor hides the cursor
func (c *Console) HideCursor() *Console {
	Print(_CSI + "?25l")
	return c
}

// ShowCursor shows the cursor
func (c *Console) ShowCursor() *Console {
	Print(_CSI + "?25h")
	return c
}
