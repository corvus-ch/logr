// Package logrus implements logr.Logger by Tim Hockin using Logrus.
package logrus

import (
	"github.com/bketelsen/logr"
	"github.com/sirupsen/logrus"
)

// New creates a new instance of logr.Logger.
func New(verbosity int, l *logrus.Logger) *logger {
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
	logger    *logrus.Logger
}

// Info implements logr.Logger.Info() by writing an event with info level or debug level in case a sub logger was
// created using V() with level greater than zero.
func (l logger) Info(args ...interface{}) {
	if l.Enabled() {
		if l.level > 0 {
			l.entry().Debug(args...)
		} else {
			l.entry().Info(args...)
		}
	}
}

// Infof implements logr.Logger.Infof() by writing an event with info level or debug level in case a sub logger was
// created using V() with level greater than zero.
func (l logger) Infof(format string, args ...interface{}) {
	if l.Enabled() {
		if l.level > 0 {
			l.entry().Debugf(format, args...)
		} else {
			l.entry().Infof(format, args...)
		}
	}
}

// Enabled implements logr.Logger.Enabled() by checking if the current verbosity level is less or equal than the loggers
// maximum verbosity.
func (l logger) Enabled() bool {
	return l.level <= l.verbosity
}

// Error implements logr.Logger.Error() by writing an event with error level.
func (l logger) Error(args ...interface{}) {
	l.entry().Error(args...)
}

// Errorf implements logr.Logger.Errorf() by writing an event with error level.
func (l logger) Errorf(format string, args ...interface{}) {
	l.entry().Errorf(format, args...)
}

// V implements logr.Logger.V.
func (l logger) V(level int) logr.InfoLogger {
	return &logger{
		level:     level,
		verbosity: l.verbosity,
		prefix:    l.prefix,
		logger:    l.logger,
	}
}

// NewWithPrefix implements logr.Logger.NewWithPrefix by adding a field named "prefix".
func (l logger) NewWithPrefix(prefix string) logr.Logger {
	return &logger{
		level:     l.level,
		verbosity: l.verbosity,
		prefix:    prefix,
		logger:    l.logger,
	}
}

func (l logger) entry() *logrus.Entry {
	if len(l.prefix) > 0 {
		return l.logger.WithField("prefix", l.prefix)
	}

	return logrus.NewEntry(l.logger)
}
