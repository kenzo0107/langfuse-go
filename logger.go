package langfuse

import (
	"fmt"
	"log"
)

var logFatal = log.Fatal

type logger interface {
	Output(int, string) error
}

type ilogger interface {
	logger
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

type debug interface {
	Debug() bool
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
}

type internalLog struct {
	logger
}

func (t internalLog) Println(v ...interface{}) {
	if err := t.Output(2, fmt.Sprintln(v...)); err != nil {
		logFatal(err)
	}
}

func (t internalLog) Printf(format string, v ...interface{}) {
	if err := t.Output(2, fmt.Sprintf(format, v...)); err != nil {
		logFatal(err)
	}
}

func (t internalLog) Print(v ...interface{}) {
	if err := t.Output(2, fmt.Sprint(v...)); err != nil {
		logFatal(err)
	}
}
