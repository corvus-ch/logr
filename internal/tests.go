package internal

import (
	"testing"
)

var (
	Msg       = "Cras mattis consectetur purus sit amet fermentum."
	Formatted = "43726173206D617474697320636F6E73656374657475722070757275732073697420616D6574206665726D656E74756D2E"
)

func Benchmark(b *testing.B, name string, f func(...interface{})) {
	b.Run(name, func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				f(Msg)
			}
		})
	})
}

func Benchmarkf(b *testing.B, name string, f func(string, ...interface{})) {
	b.Run(name, func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				f("%X", Msg)
			}
		})
	})
}
