package soumu

import (
	"log"
	"os"
)

type Logger interface {
	Warnf(string, ...interface{})
	Debugf(string, ...interface{})
}

var logger Logger = &BasicLogger{
	Logger: log.New(os.Stderr, "", log.LstdFlags),
}

func SetLogger(l Logger) {
	logger = l
}

type BasicLogger struct {
	Logger *log.Logger
}

func (bl *BasicLogger) Warnf(format string, v ...interface{}) {
	bl.Logger.Printf(format, v...)
}

func (bl *BasicLogger) Debugf(format string, v ...interface{}) {
	// suppress debug logs
}
