package logger

import "fmt"

// Logger struct
type Logger struct {
	quiet bool
}

// New logger
func NewLogger(quiet bool) Logger {
	return Logger{quiet: quiet}
}

// Print prints its argument if the --quiet flag is not passed
func (l Logger) Print(args ...interface{}) {
	if l.quiet == false {
		fmt.Println(args...)
	}
}
func (l Logger) Printf(format string, args ...interface{}) {
	if l.quiet == false {
		fmt.Printf(format, args...)
	}
}
func (l Logger) Println(args ...interface{}) {
	l.Print(args...)
}
