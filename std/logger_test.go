package std_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	stdlog "log"
	"os"
	"testing"

	test "github.com/corvus-ch/logr/internal"
	log "github.com/corvus-ch/logr/std"
)

func Example() {
	l := log.New(1, stdlog.New(os.Stdout, "", stdlog.Lshortfile))
	l.Info("Info level log message")
	l.Infof("%X", "Info level log message printed in hex values")
	l.Error("Error level log message")
	l.Errorf("%X", "Error level log message printed in hex values")
	l.NewWithPrefix("adipiscing").Info("This message is prefixed")
	l.V(1).Info("Debug level message")
	l.V(1).Infof("%X", "Debug level message in hex values")
	l.V(2).Info("This message will not be printed as its verbosity exceeds the maximum")
	// Output:
	// logger_test.go:17: Info level log message
	// logger_test.go:18: 496E666F206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// logger_test.go:19: Error level log message
	// logger_test.go:20: 4572726F72206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// adipiscinglogger_test.go:21: This message is prefixed
	// logger_test.go:22: Debug level message
	// logger_test.go:23: 4465627567206C6576656C206D65737361676520696E206865782076616C756573
}

func Example_twoLoggers() {
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	log1 := stdlog.New(buf1, "", 0)
	log2 := stdlog.New(buf2, "", 0)
	l := log.New(1, log1, log2)
	l.Info("Info level log message")
	l.Infof("%X", "Info level log message printed in hex values")
	l.Error("Error level log message")
	l.Errorf("%X", "Error level log message printed in hex values")
	l.NewWithPrefix("adipiscing").Info("This message is prefixed")
	l.V(1).Info("This message will be printed with debug level")
	l.V(1).Infof("%X", "This message will be printed with debug level as hex values")
	l.V(2).Info("This message will not be printed as its verbosity exceeds the maximum")
	fmt.Fprintln(os.Stdout, "Output of first logger:")
	buf1.WriteTo(os.Stdout)
	fmt.Fprintln(os.Stdout, "Output of second logger:")
	buf2.WriteTo(os.Stdout)
	// Output:
	// Output of first logger:
	// Error level log message
	// 4572726F72206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// Output of second logger:
	// Info level log message
	// 496E666F206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// adipiscingThis message is prefixed
	// This message will be printed with debug level
	// 54686973206D6573736167652077696C6C206265207072696E7465642077697468206465627567206C6576656C206173206865782076616C756573
}

func Example_threeLoggers() {
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	buf3 := &bytes.Buffer{}
	log1 := stdlog.New(buf1, "", 0)
	log2 := stdlog.New(buf2, "", 0)
	log3 := stdlog.New(buf3, "", 0)
	l := log.New(1, log1, log2, log3)
	l.Info("Info level log message")
	l.Infof("%X", "Info level log message printed in hex values")
	l.Error("Error level log message")
	l.Errorf("%X", "Error level log message printed in hex values")
	l.NewWithPrefix("adipiscing").Info("This message is prefixed")
	l.V(1).Info("This message will be printed with debug level")
	l.V(1).Infof("%X", "This message will be printed with debug level as hex values")
	l.V(2).Info("This message will not be printed as its verbosity exceeds the maximum")
	fmt.Fprintln(os.Stdout, "Output of first logger:")
	buf1.WriteTo(os.Stdout)
	fmt.Fprintln(os.Stdout, "Output of second logger:")
	buf2.WriteTo(os.Stdout)
	fmt.Fprintln(os.Stdout, "Output of third logger:")
	buf3.WriteTo(os.Stdout)
	// Output:
	// Output of first logger:
	// Error level log message
	// 4572726F72206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// Output of second logger:
	// Info level log message
	// 496E666F206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// adipiscingThis message is prefixed
	// Output of third logger:
	// This message will be printed with debug level
	// 54686973206D6573736167652077696C6C206265207072696E7465642077697468206465627567206C6576656C206173206865782076616C756573
}

func Benchmark(b *testing.B) {
	l := log.New(
		1,
		stdlog.New(ioutil.Discard, "", 0),
		stdlog.New(ioutil.Discard, "", 0),
		stdlog.New(ioutil.Discard, "", 0),
	)
	test.Benchmark(b, "error", l.Error)
	test.Benchmarkf(b, "errorf", l.Errorf)
	test.Benchmark(b, "info", l.Info)
	test.Benchmarkf(b, "infof", l.Infof)
	test.Benchmark(b, "v", l.V(1).Info)
	test.Benchmarkf(b, "vf", l.V(1).Infof)
	test.Benchmark(b, "disabled", l.V(2).Info)
	test.Benchmarkf(b, "disabledf", l.V(2).Infof)
}
