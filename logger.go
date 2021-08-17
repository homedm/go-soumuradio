package soumuradio

import (
	"log"
	"os"
)

var logger *log.Logger

// DebugEnable enable debugging
func DebugEnable() {
	logger = log.New(os.Stdout, "[go-soumuradio]", log.LstdFlags)
}

// SetLogger set original logger
func SetLogger(l *log.Logger) {
	logger = l
}

// LogDisable disable debugging
func LogDisable() {
	logger = nil
}

func logf(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.Printf(format, v...)
}
