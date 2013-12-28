package site

import (
	"dna"
	"dna/terminal"
	"io/ioutil"
	"os"
)

var (
	QuittingMessage dna.String = "Quitting in"
	EndingMessage   dna.String = "PROCESS COMPLETED!"
	SqlConfigPath   dna.String = "./config/app.ini"
	SiteConfigPath  dna.String = "./config/app.ini"
	SqlErrorLogPath dna.String = "./log/sql_error.log"
)

var (

	// SQLERROR is a default logger to print warning message
	HTTPERROR = terminal.NewLogger(terminal.Magenta, ioutil.Discard, "", "./log/http_error.log", 0)
	// SQLERROR is a default logger to print warning message
	SQLERROR = terminal.NewLogger(terminal.Magenta, ioutil.Discard, "", "./log/sql_error.log", 0)
	// INFO is a default logger to print info message
	INFO = terminal.NewLogger(terminal.White, os.Stdout, "INFO:", "./log/std.log", terminal.Ldate|terminal.Ltime)
	// // WARNING is a default logger to print warning message
	// WARNING = terminal.NewLogger(terminal.Magenta, os.Stdout, "WARNING:", "./log/std.log", terminal.Ldate|terminal.Ltime|terminal.Lshortfile)
	// // ERROR is a default logger to print error message
	// ERROR = terminal.NewLogger(terminal.Red, os.Stderr, "ERROR:", "./log/std.log", terminal.Ldate|terminal.Ltime|terminal.Lshortfile)
)
