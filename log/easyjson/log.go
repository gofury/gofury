package logfactory

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
)

func NewTestLogger(level log.Level) (*log.Logger, *memory.Handler) {
	h := memory.New()
	l := &log.Logger{
		Handler: h,
		Level:   level,
	}
	return l, h
}