package ring

import (
	"structs"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

type Entry struct {
	_      structs.HostLayout
	Object uintptr    // gdextension.Object
	Method uintptr    // gdextension.MethodForClass
	Shape  uint64     // gdextension.Shape
	Args   [256]byte  // copied packed args
	Result [64]byte   // result slot (for future phases)
	Refs   [16]uint16 // intra-buffer references (for future phases)
	Owner  uintptr    // back-pointer for future result offloading
}

const Size = 256 // power of 2
const Mask = Size - 1

type Ring struct {
	_       structs.HostLayout
	head    uint32
	tail    uint32
	Entries [Size]Entry
}

var Main Ring

func (r *Ring) Pending() bool {
	return r.head != r.tail
}

func (r *Ring) Buffer(object, method uintptr, shape uint64, args unsafe.Pointer) {
	if r.head-r.tail >= Size {
		r.Flush()
	}
	e := &r.Entries[r.head&Mask]
	e.Object = object
	e.Method = method
	e.Shape = shape
	n := gdextension.Shape(shape).SizeArguments()
	if n > 0 && args != nil {
		copy(e.Args[:n], unsafe.Slice((*byte)(args), n))
	}
	r.head++
}

func (r *Ring) Flush() {
	if r.head == r.tail {
		return
	}
	flush(unsafe.Pointer(&r.Entries[0]), r.tail, r.head)
	r.tail = r.head
}
