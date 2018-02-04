package pointers

import (
	"errors"
	"unsafe"
)

// Worker is the test worker object
type Worker struct {
}

// Do simulates working routine
func (np *Worker) Do() error {
	return errors.New("Not implemented")
}

// WorkerPool contains worker pointers for testing
type WorkerPool struct {
	CurWorker       *Worker
	CurUnsafeWorker unsafe.Pointer
}
