//go:build !generate

package internal

import (
	"testing"
	"unsafe"
)

func TestCallFrame(t *testing.T) {
	var Godot = new(API)
	Godot.Allocate = func(size uintptr) unsafe.Pointer {
		return unsafe.Pointer(unsafe.SliceData(make([]byte, size)))
	}
	var frame = Godot.newFrame()
	for i := 0; i < maxMethodArgs; i++ {
		frameSet[uintptr](i, frame, 0x12345678+uintptr(i))
		if *(*uintptr)(unsafe.Pointer(frame.Get(i))) != 0x12345678+uintptr(i) {
			t.Fail()
		}
	}
	if **(**uintptr)(unsafe.Pointer(frame.Args())) != 0x12345678 {
		t.Fail()
	}
}
