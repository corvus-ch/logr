// Package zap implements logr.Logger by Tim Hockin using the Zap logger by Uber.
package zap

import (
	"fmt"

	"github.com/thockin/logr"
	"go.uber.org/zap"
)

// New creates a new instance of logr.Logger.
func New(verbosity int, l *zap.Logger) *logger {
	return &logger{
		level:     0,
		verbosity: verbosity,
		prefix:    "",
		logger:    l,
	}
}

type logger struct {
	logr.Logger
	level     int
	verbosity int
	prefix    string
	logger    *zap.Logger
}

// Info implements logr.Logger.Info() by writing an event with info level or debug level in case a sub logger was
// created using V() with level greater than zero.
func (l logger) Info(args ...interface{}) {
	if l.Enabled() {
		l.infoDebug(fmt.Sprint(args...))
	}
}

// Infof implements logr.Logger.Infof() by writing an event with info level or debug level in case a sub logger was
// created using V() with level greater than zero.
func (l logger) Infof(format string, args ...interface{}) {
	if l.Enabled() {
		l.infoDebug(fmt.Sprintf(format, args...))
	}
}

// Enabled implements logr.Logger.Enabled() by checking if the current verbosity level is less or equal than the loggers
// maximum verbosity.
func (l logger) Enabled() bool {
	return l.level <= l.verbosity
}

// Error implements logr.Logger.Error() by writing an event with error level.
func (l logger) Error(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

// Errorf implements logr.Logger.Errorf() by writing an event with error level.
func (l logger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

// V implements logr.Logger.V.
func (l logger) V(level int) logr.InfoLogger {
	return logger{
		level:     level,
		verbosity: l.verbosity,
		prefix:    l.prefix,
		logger:    l.logger,
	}
}

// NewWithPrefix implements logr.Logger.NewWithPrefix by adding a field named "prefix".
func (l logger) NewWithPrefix(prefix string) logr.Logger {
	return logger{
		level:     l.level,
		verbosity: l.verbosity,
		prefix:    prefix,
		logger:    l.logger.With(zap.String("prefix", prefix)),
	}
}

func (l logger) infoDebug(msg string) {
	if l.level > 0 {
		l.logger.Debug(msg)
	} else {
		l.logger.Info(msg)
	}
}
