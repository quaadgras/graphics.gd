package startup

import (
	"sync"
	"sync/atomic"
)

type cgoHandle uintptr

func cgoNewHandle(v any) cgoHandle {
	h := handleIdx.Add(1)
	if h == 0 {
		panic("runtime/cgo: ran out of handle space")
	}
	handles.Store(h, v)
	return cgoHandle(h)
}

// Value returns the associated Go value for a valid handle.
//
// The method panics if the handle is invalid.
func (h cgoHandle) Value() any {
	v, ok := handles.Load(uintptr(h))
	if !ok {
		panic("runtime/cgo: misuse of an invalid Handle")
	}
	return v
}

// Delete invalidates a handle. This method should only be called once
// the program no longer needs to pass the handle to C and the C code
// no longer has a copy of the handle value.
//
// The method panics if the handle is invalid.
func (h cgoHandle) Delete() {
	_, ok := handles.LoadAndDelete(uintptr(h))
	if !ok {
		panic("runtime/cgo: misuse of an invalid Handle")
	}
}

var (
	handles   = sync.Map{} // map[Handle]interface{}
	handleIdx atomic.Uintptr
)
