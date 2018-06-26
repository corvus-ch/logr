// Package buffered implements Tim Hockins logr.Logger using an internal bytes.Buffer.
//
// This implementation is meant to be used in testing. Feel free to come up with other uses.
package buffered

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/bketelsen/logr"
)

const (
	levelError = "ERROR "
	levelInfo  = "INFO "
	levelV     = "V[%d] "
)

// New creates a new logr.Logger instance.
//
// The verbosity defines the upper limit. When creating a new sub logger using logr.Logger.V(), if the level passed to
// V is greater than verbosity, the sub logger will be silenced.
//
// Each written line will be prefixed depending of the loggers current level. When using logr.Logger.Error or
// logr.Logger.Errorf, the prefix will be "ERROR". When using logr.Logger.Info or logr.Logger.Infof and if the level is
// zero, the prefix will be "INFO".  For levels above zero, the prefix will be V[<level>] where <level> will be the
// current logger level.
//
// When creating a new sub logger using logr.Logger.NewWithPrefix(), the prefix will be written after the logging level
// but before the message. No whitespace is added between the prefix and the message.
//
//
func New(verbosity int) *logger {
	return &logger{
		level:     0,
		verbosity: verbosity,
		prefix:    "",
		buf:       &bytes.Buffer{},
	}
}

type logger struct {
	logr.Logger
	level     int
	verbosity int
	prefix    string
	buf       *bytes.Buffer
	mu        sync.Mutex
}

// Info implements logr.Logger.Info by writing the line to the internal buffer.
func (l logger) Info(args ...interface{}) {
	if l.Enabled() {
		l.writeLine(l.levelString(), fmt.Sprint(args...))
	}
}

// Infof implements logr.Logger.Infof by writing the line to the internal buffer.
//
// Each line will be prefixed depending of the loggers current level. If the level is zero, the prefix will be "INFO".
// For levels above zero, the prefix will be V[<level>] where <level> will be the current logger level.
func (l logger) Infof(format string, args ...interface{}) {
	if l.Enabled() {
		l.writeLine(l.levelString(), fmt.Sprintf(format, args...))
	}
}

// Enabled implements logr.Logger.Enabled by checking if the current verbosity level is less or equal than the loggers
// maximum verbosity.
func (l logger) Enabled() bool {
	return l.level <= l.verbosity
}

// Error implements logr.Logger.Error by prefixing the line with "ERROR" and write it to the internal buffer.
func (l logger) Error(args ...interface{}) {
	l.writeLine(levelError, fmt.Sprint(args...))
}

// Error implements logr.Logger.Errorf by prefixing the line with "ERROR" and write it to the internal buffer.
func (l logger) Errorf(format string, args ...interface{}) {
	l.writeLine(levelError, fmt.Sprintf(format, args...))
}

// V implements logr.Logger.V.
func (l logger) V(level int) logr.InfoLogger {
	return logger{
		level:     level,
		verbosity: l.verbosity,
		prefix:    l.prefix,
		buf:       l.buf,
		mu:        l.mu,
	}
}

// NewWithPrefix implements logr.Logger.NewWithPrefix.
//
//
func (l logger) NewWithPrefix(prefix string) logr.Logger {
	return logger{
		level:     l.level,
		verbosity: l.verbosity,
		prefix:    prefix,
		buf:       l.buf,
		mu:        l.mu,
	}
}

// Buf returns the internal buffer.
//
// Wrap with Mutex().Lock() and Mutex().Unlock() when doing write calls to preserve the write order.
func (l logger) Buf() *bytes.Buffer {
	return l.buf
}

// Mutex returns the sync.Mutex used to preserve the order of writes to the buffer.
func (l *logger) Mutex() *sync.Mutex {
	return &l.mu
}

func (l logger) writeLine(level, line string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf.WriteString(level)
	l.buf.WriteString(l.prefix)
	l.buf.WriteString(line)
	if !strings.HasSuffix(line, "\n") {
		l.buf.WriteRune('\n')
	}
}

func (l logger) levelString() string {
	if l.level > 0 {
		return fmt.Sprintf(levelV, l.level)
	}

	return levelInfo
}
