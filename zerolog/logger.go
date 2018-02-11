// Package std implements logr.Logger by Tim Hockin using Zerolog.
package zerolog

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/thockin/logr"
)

// New creates a new instance of logr.Logger.
func New(verbosity int, l zerolog.Logger) *logger {
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
	logger    zerolog.Logger
}

// Info implements logr.Logger.Info() by writing an event with info level or debug level in case a sub logger was
// created using V() with level greater than zero.
func (l logger) Info(args ...interface{}) {
	if l.Enabled() {
		l.event().Msg(fmt.Sprint(args...))
	}
}

// Infof implements logr.Logger.Infof() by writing an event with info level or debug level in case a sub logger was
// created using V() with level greater than zero.
func (l logger) Infof(format string, args ...interface{}) {
	if l.Enabled() {
		l.event().Msgf(format, args...)
	}
}

// Enabled implements logr.Logger.Enabled() by checking if the current verbosity level is less or equal than the loggers
// maximum verbosity.
func (l logger) Enabled() bool {
	return l.level <= l.verbosity
}

// Error implements logr.Logger.Error() by writing an event with error level.
func (l logger) Error(args ...interface{}) {
	l.logger.Error().Msg(fmt.Sprint(args...))
}

// Errorf implements logr.Logger.Errorf() by writing an event with error level.
func (l logger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
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
		logger:    l.logger.With().Str("prefix", prefix).Logger(),
	}
}

func (l logger) event() *zerolog.Event {
	if l.level > 0 {
		return l.logger.Debug()
	}

	return l.logger.Info()
}
