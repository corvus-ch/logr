package log_test

import (
	"bytes"
	"fmt"
	stdlog "log"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/corvus-ch/logr/buffered"
	test "github.com/corvus-ch/logr/internal"
	"github.com/corvus-ch/logr/log"
	"github.com/corvus-ch/logr/std"
	"github.com/stretchr/testify/assert"
)

func Example() {
	l := buffered.New(1)
	log.SetLogger(l)
	log.Info("Info level log message")
	log.Infof("%X", "Info level log message printed in hex values")
	log.Error("Error level log message")
	log.Errorf("%X", "Error level log message printed in hex values")
	log.NewWithPrefix("adipiscing").Info("This message is prefixed")
	log.V(1).Info("Debug level message")
	log.V(1).Infof("%X", "Debug level message in hex values")
	log.V(2).Info("This message will not be printed as its verbosity exceeds the maximum")
	log.Print("Info level log message using log.Logger compatibility")
	log.Printf("%X", "Info level log message using log.Logger compatibility in hex values")
	log.Println("Info level log line using log.Logger compatibility")
	l.Buf().WriteTo(os.Stdout)
	// Output:
	// INFO Info level log message
	// INFO 496E666F206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// ERROR Error level log message
	// ERROR 4572726F72206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// INFO adipiscingThis message is prefixed
	// V[1] Debug level message
	// V[1] 4465627567206C6576656C206D65737361676520696E206865782076616C756573
	// INFO Info level log message using log.Logger compatibility
	// INFO 496E666F206C6576656C206C6F67206D657373616765207573696E67206C6F672E4C6F6767657220636F6D7061746962696C69747920696E206865782076616C756573
	// INFO Info level log line using log.Logger compatibility
}

func Example_callDepth() {
	l := std.New(1, stdlog.New(os.Stdout, "", stdlog.Lshortfile))
	l.SetCallDepth(3)
	log.SetLogger(l)
	log.Error("Cras justo odio, dapibus ac facilisis.")
	log.NewWithPrefix("adipiscing").Info("Cras justo odio, dapibus ac facilisis.")
	log.V(1).Info("Cras justo odio, dapibus ac facilisis.")
	// Output:
	// logger_test.go:50: Cras justo odio, dapibus ac facilisis.
	// adipiscinglogger_test.go:51: Cras justo odio, dapibus ac facilisis.
	// logger_test.go:52: Cras justo odio, dapibus ac facilisis.
}

func setup() *bytes.Buffer {
	l := buffered.New(1)
	log.SetLogger(l)
	return l.Buf()
}

func testPanic(t *testing.T, name, panic, out string, f func()) {
	t.Run(name, func(t *testing.T) {
		buf := setup()
		assert.PanicsWithValue(t, panic, f)
		assert.Equal(t, fmt.Sprintf("ERROR %s\n", out), buf.String())
	})
}

func TestFatal(t *testing.T) {
	defer monkey.Patch(os.Exit, func(code int) { panic(fmt.Sprintf("exit status %d", code)) }).Unpatch()
	testPanic(t, "fatal", "exit status 1", test.Msg, func() {
		log.Fatal(test.Msg)
	})
	testPanic(t, "fatalf", "exit status 1", test.Formatted, func() {
		log.Fatalf("%X", test.Msg)
	})
	testPanic(t, "fatalln", "exit status 1", test.Msg, func() {
		log.Fatalln(test.Msg)
	})
}

func TestPanic(t *testing.T) {
	testPanic(t, "panic", test.Msg, test.Msg, func() {
		log.Panic(test.Msg)
	})
	testPanic(t, "panicf", test.Formatted, test.Formatted, func() {
		log.Panicf("%X", test.Msg)
	})
	testPanic(t, "panicln", test.Msg, test.Msg, func() {
		log.Panicln(test.Msg)
	})
}
