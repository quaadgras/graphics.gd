//go:build !cgo && !js

package ring

import (
	"unsafe"

	gdunsafe "graphics.gd"
)

func flush(entries unsafe.Pointer, tail, head uint32) {
	ring := (*[Size]Entry)(entries)
	for i := tail; i != head; i++ {
		CrashIndex = i & Mask
		e := &ring[i&Mask]
		gdunsafe.Object(e.Object).UnsafeCall(
			gdunsafe.MethodForClass(e.Method),
			unsafe.Pointer(&e.Result[0]),
			uint64(e.Shape),
			unsafe.Pointer(&e.Args[0]),
		)
	}
	CrashIndex = 0xFFFFFFFF
}
