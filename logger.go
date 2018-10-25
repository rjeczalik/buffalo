package buffalo

import (
	"github.com/gobuffalo/logger"
)

// Logger interface is used throughout Buffalo
// apps to log a whole manner of things.
type Logger interface {
	WithField(string, interface{}) Logger
	WithFields(map[string]interface{}) Logger
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Printf(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
}

// NewLogger based on the specified log level.
// This logger will log to the STDOUT in a human readable,
// but parseable form.
/*
	Example: time="2016-12-01T21:02:07-05:00" level=info duration=225.283µs human_size="106 B" method=GET path="/" render=199.79µs request_id=2265736089 size=106 status=200
*/
func NewLogger(level string) Logger {
	return lWrap{logger.NewLogger(level)}
}

type lWrap struct {
	logger.FieldLogger
}

func (l lWrap) WithField(s string, i interface{}) Logger {
	return lWrap{l.FieldLogger.WithField(s, i)}
}

func (l lWrap) WithFields(m map[string]interface{}) Logger {
	return lWrap{l.FieldLogger.WithFields(m)}
}
