package soumu

import (
	"log"
	"os"
)

var logger *log.Logger

func DebugEnable() {
	logger = log.New(os.Stdout, "[go-soumuradio]", log.LstdFlags)
}

func SetLogger(l *log.Logger) {
	logger = l
}

func LogDisable() {
	logger = nil
}

func logf(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.Printf(format, v...)
}
