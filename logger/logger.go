package logger

import "fmt"

// Logger struct
type Logger struct {
	quiet   bool
	verbose bool
}

// New logger
func NewLogger(quiet, verbose bool) Logger {
	return Logger{
		quiet:   quiet,
		verbose: verbose,
	}
}

// Print prints its argument if the --quiet flag is not passed
func (l Logger) Print(args ...interface{}) {
	if l.quiet == false {
		fmt.Print(args...)
	}
}

// Print prints its argument if the --quiet flag is not passed
func (l Logger) Println(args ...interface{}) {
	if l.quiet == false {
		fmt.Println(args...)
	}
}

func (l Logger) Printf(format string, args ...interface{}) {
	if l.quiet == false {
		fmt.Printf(format, args...)
	}
}

func (l Logger) Verbose(args ...interface{}) {
	if l.verbose == true {
		l.Println(args...)
	}
}
