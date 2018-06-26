package writer_adapter

import (
	"io"

	"github.com/bketelsen/logr"
)

// New creates a writer which directly writes to the given logger function.
func New(out func(args ...interface{})) io.Writer {
	return &writer{out}
}

// NewInfoWriter creates a writer which directly writes to the given logger using info level.
func NewInfoWriter(l logr.InfoLogger) io.Writer {
	return New(l.Info)
}

// NewInfoWriter creates a writer which directly writes to the given logger using error level.
func NewErrorWriter(l logr.Logger) io.Writer {
	return New(l.Error)
}

type writer struct {
	out func(args ...interface{})
}

func (w writer) Write(p []byte) (int, error) {
	w.out(string(p))
	return len(p), nil
}
