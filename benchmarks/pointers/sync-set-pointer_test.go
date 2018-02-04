package pointers

import "testing"

// Case: set pointer synchronously
// atomic.SetPointer() vs Golang assembly

func BenchmarkSwitchPtrWithAssembly(b *testing.B) {
	wp := WorkerPool{}
	for n := 0; n < b.N; n++ {
		wp.switchCurWorker(&Worker{})
		wp.CurWorker.Do()
	}
}

func BenchmarkSwitchPtrWithAtomic(b *testing.B) {
	wp := WorkerPool{}
	for n := 0; n < b.N; n++ {
		wp.switchCurUnsafeWorker(&Worker{})
		(*Worker)(wp.CurUnsafeWorker).Do()
	}
}

/* 2018-01-11 go 1.9.1 linux/amd64

BenchmarkSwitchPtrWithAssembly-8        100000000               12.1 ns/op
BenchmarkSwitchPtrWithAtomic-8          100000000               14.5 ns/op

*/
