//go:build !cgo

package ring

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

func flush(entries unsafe.Pointer, tail, head uint32) {
	ring := (*[Size]Entry)(entries)
	for i := tail; i != head; i++ {
		CrashIndex = i & Mask
		e := &ring[i&Mask]
		gdextension.Host.Objects.Unsafe.Call(
			gdextension.Object(e.Object),
			gdextension.MethodForClass(e.Method),
			gdextension.CallReturns[any](unsafe.Pointer(&e.Result[0])),
			gdextension.Shape(e.Shape),
			gdextension.CallAccepts[any](unsafe.Pointer(&e.Args[0])),
		)
	}
	CrashIndex = 0xFFFFFFFF
}
