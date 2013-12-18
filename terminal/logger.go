package terminal

import (
	. "dna"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// Define constant flags for Logger
const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date: 2009/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// Logger reimplements log.Logger from standard library. It adds color and path fields to Logger{}.
// It supports multiwriter. By default it writes error, warning and info logger to standard output
// and std.log file.Logger prints color outputs to os.Stdout or os.Stderr.
// And it removes color then writing to a file
//
//	// Example
// 	ERROR.Println("This is error message")
// 	INFO.Println("This is info message")
// 	WARNING.Println("This is a warning message")
//
//***************************************************************************
//
//***NOTICE***: This Logger is not complete, path field should be global var.
type Logger struct {
	mu         sync.Mutex // ensures atomic writes; protects the following fields
	prefix     string     // prefix to write at beginning of each line
	flag       int        // properties
	out        io.Writer  // destination for output
	buf        []byte     // for accumulating text to write
	color      Int
	console    *Console
	path       String
	defaultOut io.Writer // default writer stdout or stderr
}

// New creates a new Logger.   The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func NewLogger(color Int, out io.Writer, prefix, path String, flag Int) *Logger {
	var multi io.Writer
	if path != "" {
		file, err := os.OpenFile(path.String(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			Log(fmt.Sprintf("Failed to open log file %v : %v - It uses standard output", path, err))
			multi = io.MultiWriter(out)
		} else {
			multi = io.MultiWriter(file, out)
		}

	} else {
		multi = io.MultiWriter(out)
	}
	return &Logger{out: multi, prefix: prefix.String(), path: path, flag: int(flag), color: color, console: NewConsole(), defaultOut: out}
}

// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
// Knows the buffer has capacity.
func itoa(buf *[]byte, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		*buf = append(*buf, '0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	*buf = append(*buf, b[bp:]...)
}

func (l *Logger) formatHeader(buf *[]byte, t time.Time, file string, line int) {
	*buf = append(*buf, l.prefix...)
	if l.flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if l.flag&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if l.flag&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if l.flag&(Lshortfile|Llongfile) != 0 {
		if l.flag&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}
}

// Output writes the output for a logging event.  The string s contains
// the text to print after the prefix specified by the flags of the
// Logger.  A newline is appended if the last character of s is not
// already a newline.  Calldepth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l *Logger) Output(calldepth int, s string) error {
	now := time.Now() // get this early.
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(Lshortfile|Llongfile) != 0 {
		// release lock while getting caller info - it's expensive.
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, file, line)
	l.buf = append(l.buf, s...)
	if len(s) > 0 && s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)
	return err
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
//
// 	* Notice: It adds '\n' to the end of the value if newline is not specified based on standard library
func (l *Logger) Printf(format String, v ...interface{}) {
	l.console.Foreground(l.color)
	l.Output(2, fmt.Sprintf(format.String(), v...))
	l.console.Display(ResetCode)
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
//
// 	* Notice: It adds '\n' to the end of the value if newline is not specified based on standard library
func (l *Logger) Print(v ...interface{}) {
	l.console.Foreground(l.color)
	l.Output(2, fmt.Sprint(v...))
	l.console.Display(ResetCode)
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Println(v ...interface{}) {
	l.console.Foreground(l.color)
	l.Output(2, fmt.Sprintln(v...))
	l.console.Display(ResetCode)
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *Logger) Fatal(v ...interface{}) {
	l.console.Foreground(l.color)
	l.Output(2, fmt.Sprint(v...))
	l.console.Display(ResetCode)
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format String, v ...interface{}) {
	l.console.Foreground(l.color)
	l.Output(2, fmt.Sprintf(format.String(), v...))
	l.console.Display(ResetCode)
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (l *Logger) Fatalln(v ...interface{}) {
	l.console.Foreground(l.color)
	l.Output(2, fmt.Sprintln(v...))
	l.console.Display(ResetCode)
	os.Exit(1)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	l.console.Foreground(l.color)
	s := fmt.Sprint(v...)
	l.Output(2, s)
	l.console.Display(ResetCode)
	panic(s)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (l *Logger) Panicf(format String, v ...interface{}) {
	l.console.Foreground(l.color)
	s := fmt.Sprintf(format.String(), v...)
	l.Output(2, s)
	l.console.Display(ResetCode)
	panic(s)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (l *Logger) Panicln(v ...interface{}) {
	l.console.Foreground(l.color)
	s := fmt.Sprintln(v...)
	l.Output(2, s)
	l.console.Display(ResetCode)
	panic(s)
}

// Flags returns the output flags for the logger.
func (l *Logger) Flags() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.flag
}

// SetFlags sets the output flags for the logger.
func (l *Logger) SetFlags(flag Int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flag = int(flag)
}

// Prefix returns the output prefix for the logger.
func (l *Logger) Prefix() String {
	l.mu.Lock()
	defer l.mu.Unlock()
	return String(l.prefix)
}

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(prefix String) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix.String()
}

// Color returns the output Color for the logger.
func (l *Logger) Color() Int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.color
}

// SetColor sets the output color for the logger.
func (l *Logger) SetColor(color Int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = color
}

// Path returns the output Path for the logger.
func (l *Logger) Path() String {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.path
}

// SetPath sets the output color for the logger.
func (l *Logger) SetPath(path String) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.path = path
	var multi io.Writer
	if path != "" {
		file, err := os.OpenFile(path.String(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Failed to open log file %v : %v", path, err))
		}
		multi = io.MultiWriter(file, l.defaultOut)
	} else {
		multi = io.MultiWriter(l.defaultOut)
	}
	l.out = multi
}
