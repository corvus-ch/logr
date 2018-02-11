package zerolog_test

import (
	"io/ioutil"
	"os"
	"testing"

	test "github.com/corvus-ch/logr/internal"
	log "github.com/corvus-ch/logr/zerolog"
	"github.com/rs/zerolog"
)

func Example() {
	l := log.New(1, zerolog.New(os.Stdout))
	l.Info("Info level log message")
	l.Infof("%X", "Info level log message printed in hex values")
	l.Error("Error level log message")
	l.Errorf("%X", "Error level log message printed in hex values")
	l.NewWithPrefix("adipiscing").Info("This message has a prefix field")
	l.V(1).Info("This message will be printed with debug level")
	l.V(1).Infof("%X", "This message will be printed with debug level as hex values")
	l.V(2).Info("This message will not be printed as its verbosity exceeds the maximum")
	// Output:
	// {"level":"info","message":"Info level log message"}
	// {"level":"info","message":"496E666F206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573"}
	// {"level":"error","message":"Error level log message"}
	// {"level":"error","message":"4572726F72206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573"}
	// {"level":"info","prefix":"adipiscing","message":"This message has a prefix field"}
	// {"level":"debug","message":"This message will be printed with debug level"}
	// {"level":"debug","message":"54686973206D6573736167652077696C6C206265207072696E7465642077697468206465627567206C6576656C206173206865782076616C756573"}
}

func Benchmark(b *testing.B) {
	l := log.New(1, zerolog.New(ioutil.Discard))
	test.Benchmark(b, "error", l.Error)
	test.Benchmarkf(b, "errorf", l.Errorf)
	test.Benchmark(b, "info", l.Info)
	test.Benchmarkf(b, "infof", l.Infof)
	test.Benchmark(b, "v", l.V(1).Info)
	test.Benchmarkf(b, "vf", l.V(1).Infof)
	test.Benchmark(b, "disabled", l.V(2).Info)
	test.Benchmarkf(b, "disabledf", l.V(2).Infof)
}
