package logger

import (
	"os"

	"github.com/withmandala/go-log"
)

var logger *log.Logger

func Logger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stderr)
	}
	return logger
}
