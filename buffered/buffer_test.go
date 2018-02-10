package buffered_test

import (
	"os"

	log "github.com/corvus-ch/logr/buffered"
)

func Example() {
	l := log.New(1)
	l.Info("Info level log message")
	l.Infof("%X", "Info level log message printed in hex values")
	l.Error("Error level log message")
	l.Errorf("%X", "Error level log message printed in hex values")
	l.NewWithPrefix("adipiscing").Info("This message has a prefix field")
	l.V(1).Info("This message will be printed with verbose level")
	l.V(1).Infof("%X", "This message will be printed with verbose level as hex values")
	l.V(2).Info("This message will not be printed as its verbosity exceeds the maximum")
	l.Buf().WriteTo(os.Stdout)
	// Output:
	// INFO Info level log message
	// INFO 496E666F206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// ERROR Error level log message
	// ERROR 4572726F72206C6576656C206C6F67206D657373616765207072696E74656420696E206865782076616C756573
	// INFO adipiscingThis message has a prefix field
	// V[1] This message will be printed with verbose level
	// V[1] 54686973206D6573736167652077696C6C206265207072696E746564207769746820766572626F7365206C6576656C206173206865782076616C756573
}

func Example_mutex() {
	l := log.New(1)
	l.Mutex().Lock()
	defer l.Mutex().Unlock()
	go func() {
		// Due to the active lock, the following log call has to wait.
		l.Info("Cras mattis consectetur purus sit amet fermentum.")
	}()
	l.Buf().WriteString("Duis mollis, ")
	l.Buf().WriteString("est non commodo luctus, ")
	l.Buf().WriteString("nisi erat porttitor ligula, ")
	l.Buf().WriteString("eget lacinia odio sem nec elit.")
	l.Buf().WriteRune('\n')
	// The buffers content is written to the output before releasing the lock. Thus the info log call does not show up
	// in the output.
	l.Buf().WriteTo(os.Stdout)
	// Output: Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit.
}
