package pointers

import (
	"sync/atomic"
	"unsafe"
)

func (wp *WorkerPool) switchCurWorker(newWorker *Worker) {
	storePointer(unsafe.Pointer(&wp.CurWorker), unsafe.Pointer(newWorker))
}

func (wp *WorkerPool) switchCurUnsafeWorker(newWorker *Worker) {
	atomic.StorePointer(&wp.CurUnsafeWorker, unsafe.Pointer(newWorker))
}

func storePointer(destPtr unsafe.Pointer, newPtr unsafe.Pointer)
