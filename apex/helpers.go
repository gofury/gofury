package apex

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
)

func ParseApexLogLevel(logLevel string) log.Level {
	if level, err := log.ParseLevel(logLevel); err != nil {
		return log.InfoLevel
	} else {
		return level
	}
}

func NewTestLogger(level log.Level) (*log.Logger, *memory.Handler) {
	h := memory.New()
	l := &log.Logger{
		Handler: h,
		Level:   level,
	}
	return l, h
}

