// Package log provides a global logger using logr.Logger.
//
// By default, it uses github.com/corvus-ch/std which is configured to write to STDERR.
package log

import (
	"fmt"
	"log"
	"os"

	"github.com/corvus-ch/logr/std"
	"github.com/thockin/logr"
)

var logger logr.Logger

func init() {
	l := std.New(0, log.New(os.Stderr, "", log.LstdFlags))
	l.SetCallDepth(4)
	SetLogger(l)
}

// SetLogger sets a new default logger.
func SetLogger(l logr.Logger) {
	logger = l
}

// Info calls Info() of the default logger.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof calls Infof() of the default logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Error calls Error() of the default logger.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf calls Errorf() the default logger.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// V calls V() of the default logger.
func V(level int) logr.InfoLogger {
	l := logger.V(level)
	sl, ok := l.(std.Logger)
	if ok {
		sl.SetCallDepth(2)
		return sl
	}

	return l
}

// NewWithPrefix calls NewWithPrefix() of the default logger.
func NewWithPrefix(prefix string) logr.Logger {
	l := logger.NewWithPrefix(prefix)
	sl, ok := l.(std.Logger)
	if ok {
		sl.SetCallDepth(2)
		return sl
	}

	return l
}

// Print is equivalent to Info()
func Print(args ...interface{}) {
	logger.Info(args...)
}

// Printf is equivalent to Infof()
func Printf(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Println is equivalent to Info()
func Println(args ...interface{}) {
	logger.Info(args...)
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	logger.Error(args...)
	os.Exit(1)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
	os.Exit(1)
}

// Fatalln is equivalent to Error() followed by a call to os.Exit(1).
func Fatalln(args ...interface{}) {
	logger.Error(args...)
	os.Exit(1)
}

// Panic is equivalent to Error() followed by a call to panic().
func Panic(args ...interface{}) {
	logger.Error(args...)
	panic(fmt.Sprint(args...))
}

// Panicf is equivalent to Errorf() followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
	panic(fmt.Sprintf(format, args...))
}

// Panicln is equivalent to Error() followed by a call to panic().
func Panicln(args ...interface{}) {
	logger.Error(args...)
	panic(fmt.Sprint(args...))
}
