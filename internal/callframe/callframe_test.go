package callframe_test

import (
	"testing"
	"unsafe"

	"grow.graphics/gd/internal/callframe"
)

func TestFrame(t *testing.T) {
	var frame = callframe.New()

	ptr := callframe.Arg(frame, 100)

	if *(*int)(unsafe.Pointer(ptr.Uintptr())) != 100 {
		t.Fatal("ptr != 100")
	}
}

var TestEscape = func(a, b, c *int, ret *int) { *ret = 22 }

func BenchmarkFrame(b *testing.B) {
	var frame = callframe.New()
	b.ResetTimer()
	frame.Free()

	for i := 0; i < b.N; i++ {
		frame = callframe.New()
		ptr1 := callframe.Arg(frame, 100)
		ret := callframe.Ret[int](frame)
		ptr2 := callframe.Arg(frame, 200)
		ptr3 := callframe.Arg(frame, 300)
		TestEscape((*int)(unsafe.Pointer(ptr1.Uintptr())), (*int)(unsafe.Pointer(ptr2.Uintptr())), (*int)(unsafe.Pointer(ptr3.Uintptr())), (*int)(unsafe.Pointer(ret.Uintptr())))
		if ret.Get() != 22 {
			b.Fatal("ret != 22")
		}
		frame.Free()
	}
}

func BenchmarkEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arg = 100
		var arg2 = 200
		var arg3 = 200
		var ret int
		TestEscape(&arg, &arg2, &arg3, &ret)
		if ret != 22 {
			b.Fatal("ret != 22")
		}
	}
}
