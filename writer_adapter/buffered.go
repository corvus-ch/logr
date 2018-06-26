package writer_adapter

import (
	"bytes"
	"io"

	"github.com/bketelsen/logr"
)

// NewBuffered creates a line buffered bufferedWriter which writes lines to the given function.
func NewBuffered(out func(args ...interface{})) io.WriteCloser {
	return &bufferedWriter{out, &bytes.Buffer{}}
}

// NewBufferedInfoWriter creates a new bufferedWriter which writes lines to the given logger using info level.
func NewBufferedInfoWriter(l logr.InfoLogger) io.WriteCloser {
	return NewBuffered(l.Info)
}

// NewBufferedErrorWriter creates a new bufferedWriter which writes lines to the given logger using error level.
func NewBufferedErrorWriter(l logr.Logger) io.WriteCloser {
	return NewBuffered(l.Error)
}

type bufferedWriter struct {
	out func(args ...interface{})
	buf *bytes.Buffer
}

func (w bufferedWriter) Write(p []byte) (n int, err error) {
	if n, err = w.buf.Write(p); err != nil {
		return
	}

	err = w.writeLines()

	return
}

func (w bufferedWriter) Close() error {
	if line := w.buf.String(); len(line) > 0 {
		w.out(line)
	}
	w.buf.Reset()
	return nil
}

func (w bufferedWriter) writeLines() error {
	for {
		line, err := w.buf.ReadString('\n')
		if len(line) > 0 {
			if err := w.writeLine(line); err != nil {
				return err
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (w bufferedWriter) writeLine(line string) error {
	if line[len(line)-1] == '\n' {
		w.out(line[:len(line)-1])
	} else {
		if _, err := w.buf.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}
