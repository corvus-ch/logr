// Package std implements Tim Hockin logr.Logger using Golangs standard log.Logger.
package std

import (
	"log"

	"fmt"
	"github.com/thockin/logr"
)

// New creates a new instance of logr.Logger.
//
// The function takes one or several *log.Logger instances. If only one *log.Logger is provided, this logger will be
// used for all log levels. If two *log.Logger instances are provided, the first one will be used for logs with level
// error and the second for all info levels. If tree or more *log.Logger instances are provided, the third logger and
// any consecutive loggers are used for the verbose levels created with logr.Logger.V().
//
// The verbosity defines the upper limit. When creating a new sub logger using logr.Logger.V(), if the level passed to
// V is greater than verbosity, the sub logger will be silenced.
//
// If verbosity is smaller than the number of *log.Logger instances, any additional logger will be ignored. If verbosity
// is greater than the number of *log.Logger instances, the last logger will be used to fill in for any missing levels.
//
// This implementation takes control over the prefix of *log.Logger. Any prefix set on instantiation, will be ignored.
// Flags are preserved.
//
// Example:
//
//     l1 := log.New(os.Stderr, "", 0)
//     l2 := log.New(os.Stdout, "", 0)
//     l3 := log.New(&bytes.Buffer{}, "", 0)
//
//     // All output written to the same logger.
//     lgr1 := New(1, l2)
//     lgr1.Error("I will be written to STDERR")
//     lgr1.Info("I will be written to STDERR too")
//     lgr1.V(1).Info("And mee too")
//     lgr1.V(2).Info("I will be ignored")
//
//     // Output written to tree different loggers depending on the severity and verbosity.
//     lgr2 := New(1, l1, l2, l3)
//     lgr2.Error("I will be written to STDERR")
//     lgr2.Info("I will be written to STDOUT")
//     lgr2.V(1).Info("I will be buffered")
//     lgr2.V(2).Info("I will be ignored")
//
//     // Output written only to the first two loggers. The third one is ignored.
//     lgr3 := New(0, l1, l2, l3)
//     lgr3.Error("I will be written to STDERR")
//     lgr3.Info("I will be written to STDOUT")
//     lgr3.V(1).Info("I will be ignored")
//     lgr3.V(2).Info("And I will be ignored too")
//
func New(verbosity int, ll ...*log.Logger) *logger {
	loggers := make([]*log.Logger, verbosity+2)
	for i := range loggers {
		if i < len(ll) {
			loggers[i] = ll[i]
		} else {
			loggers[i] = ll[len(ll)-1]
		}
	}
	return &logger{
		level:     0,
		verbosity: verbosity,
		prefix:    "",
		loggers:   loggers,
		callDepth: 2,
	}
}

type logger struct {
	logr.Logger
	level     int
	verbosity int
	prefix    string
	loggers   []*log.Logger
	callDepth int
}

// Info implements logr.Logger.Info by writing to log.Logger of with the matching level.
func (l logger) Info(args ...interface{}) {
	if l.Enabled() {
		l.loggers[l.index()].SetPrefix(l.prefix)
		l.loggers[l.index()].Output(l.callDepth, fmt.Sprint(args...))
	}
}

// Infof implements logr.Logger.Infof by writing to log.Logger of with the matching level.
func (l logger) Infof(format string, args ...interface{}) {
	if l.Enabled() {
		l.loggers[l.index()].SetPrefix(l.prefix)
		l.loggers[l.index()].Output(l.callDepth, fmt.Sprintf(format, args...))
	}
}

// Enabled implements logr.Logger.Enabled by checking if the current verbosity level is less or equal than the loggers
// maximum verbosity.
func (l logger) Enabled() bool {
	return l.level <= l.verbosity
}

// Error implements logr.Logger.Error by writing to the first log.Logger.
func (l logger) Error(args ...interface{}) {
	l.loggers[0].SetPrefix(l.prefix)
	l.loggers[0].Output(l.callDepth, fmt.Sprint(args...))
}

// Errorf implements logr.Logger.Errorf by writing to the first log.Logger.
func (l logger) Errorf(format string, args ...interface{}) {
	l.loggers[0].SetPrefix(l.prefix)
	l.loggers[0].Output(l.callDepth, fmt.Sprintf(format, args...))
}

// V implements logr.Logger.V.
func (l logger) V(level int) logr.InfoLogger {
	return logger{
		level:     level,
		verbosity: l.verbosity,
		prefix:    l.prefix,
		loggers:   l.loggers,
		callDepth: l.callDepth + 1,
	}
}

// NewWithPrefix implements logr.Logger.NewWithPrefix.
func (l logger) NewWithPrefix(prefix string) logr.Logger {
	return logger{
		level:     l.level,
		verbosity: l.verbosity,
		prefix:    prefix,
		loggers:   l.loggers,
		callDepth: l.callDepth + 1,
	}
}

// SetCallDepth sets the call depth passed to log.Logger.Output.
func (l *logger) SetCallDepth(depth int) {
	l.callDepth = depth
}

func (l logger) index() int {
	return l.level + 1
}
