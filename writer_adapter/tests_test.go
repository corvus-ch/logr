package writer_adapter_test

import (
	"os"

	"github.com/corvus-ch/logr/buffered"
	"github.com/corvus-ch/logr/writer_adapter"
)

func Example() {
	l := buffered.New(0)
	w1 := writer_adapter.NewInfoWriter(l)
	w2 := writer_adapter.NewErrorWriter(l)
	w1.Write([]byte("I am written with info level\n"))
	w1.Write([]byte("and so am I."))
	w2.Write([]byte("I am written with error level\n"))
	w2.Write([]byte("and I am an error too."))
	l.Buf().WriteTo(os.Stdout)
	// Output:
	// INFO I am written with info level
	// INFO and so am I.
	// ERROR I am written with error level
	// ERROR and I am an error too.
}

func Example_buffered() {
	l := buffered.New(0)
	w1 := writer_adapter.NewBufferedInfoWriter(l)
	w2 := writer_adapter.NewBufferedErrorWriter(l)
	w1.Write([]byte("I am written with info level\n"))
	w1.Write([]byte("and I am need to wait until the close"))
	w2.Write([]byte("I am written with error level\n"))
	w2.Write([]byte("and I am an error which needs to wait too."))
	// This will print all complete lines.
	os.Stdout.Write(l.Buf().Bytes())
	w1.Close()
	w2.Close()
	// This will print all complete lines plus the fragments kept in the buffer.
	os.Stdout.Write(l.Buf().Bytes())
	// Output:
	// INFO I am written with info level
	// ERROR I am written with error level
	// INFO I am written with info level
	// ERROR I am written with error level
	// INFO and I am need to wait until the close
	// ERROR and I am an error which needs to wait too.
}
