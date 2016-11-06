package newrelic

import (
	"github.com/apex/log"
	"github.com/newrelic/go-agent"
)

// NewRelicLogger wraps apex/log for use with New Relic.
type NewRelicLogger struct {
	l     log.Interface
	debug bool
}

// NewNewRelicLogger returns a new New Relic logger.
func NewNewRelicLogger(l log.Interface) *NewRelicLogger {
	return &NewRelicLogger{l: l, debug: false}
}

// SetDebug sets the debug state.
func (n *NewRelicLogger) SetDebug(debug bool) {
	n.debug = debug
}

// Error logs an error message.
func (n *NewRelicLogger) Error(msg string, context map[string]interface{}) {
	n.l.WithField("component", "newrelic").WithFields(log.Fields(context)).Error(msg)
}

// Warn logs an warn message.
func (n *NewRelicLogger) Warn(msg string, context map[string]interface{}) {
	n.l.WithField("component", "newrelic").WithFields(log.Fields(context)).Warn(msg)
}

// Info logs an info message.
func (n *NewRelicLogger) Info(msg string, context map[string]interface{}) {
	n.l.WithField("component", "newrelic").WithFields(log.Fields(context)).Info(msg)
}

// Debug logs an debug message, if debug is enabled.
func (n *NewRelicLogger) Debug(msg string, context map[string]interface{}) {
	if n.debug {
		n.l.WithField("component", "newrelic").WithFields(log.Fields(context)).Debug(msg)
	}
}

// DebugEnabled returns the debug state.
func (n *NewRelicLogger) DebugEnabled() bool {
	return n.debug
}

