// Package callframe provides a way to create Godot-compatible callframes.
package callframe

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

var pin runtime.Pinner

var frames sync.Pool

// Frame can be used to reduce the number of allocations required to pass
// pointers across API boundaries. Since dynamic function calls and CGO
// always cause pointers to escape, a frame enables a block of memory to be
// used to copy in const value arguments. The frame is cached using [sync.Pool]
// and recycled across calls. Suitable for use when the function being
// called is only borrowing the pointer for the lifetime of the call. A
// frame has a fixed size, if it runs out of space, arguments will be
// allocated onto the Go heap as usual.
type Frame struct {
	len uint8
	idx uint8
	//oom bool

	pin runtime.Pinner
	ptr [16]*uintptr
	typ [16]reflect.Type
	buf [16 * 6]uintptr
}

// Nil implements [Any].
type Nil struct{}

// Uintptr always returns 0.
func (Nil) Uintptr() uintptr { return 0 }

// New returns either a new [Frame], or a recycled frame.
func New() *Frame {
	frame, ok := frames.Get().(*Frame)
	if !ok {
		frame = new(Frame)
		pin.Pin(frame)
	}
	return frame
}

type array struct{}

// Array of arguments.
type Args struct {
	_    [0]array
	void unsafe.Pointer
}

func (array Args) Uintptr() uintptr {
	return uintptr(array.void)
}

// Array returns a pointer to array of arguments offset by i.
func (frame *Frame) Array(i int) Args {
	return Args{
		void: unsafe.Pointer(&frame.ptr[i]),
	}
}

// Free recycles the [Frame]. Do not reuse the frame after
// calling this method.
func (frame *Frame) Free() {
	frame.len = 0
	frame.idx = 0
	clear(frame.buf[:frame.len])
	frame.pin.Unpin()
	frames.Put(frame)
}

// Ptr that does not escape as an allocation
// to the heap, T must be a value type that
// does not contain any pointers. The value
// must not be used after the frame has been
// freed.
type Ptr[T comparable] struct {
	_    [0]T
	void unsafe.Pointer
}

// Uintptr returns the uintptr value of the pointer. Useful for passing to C code.
func (ptr Ptr[T]) Uintptr() uintptr {
	return uintptr(ptr.void)
}

// UnsafePoiner returns the unsafe.Pointer value of the pointer.
func (ptr Ptr[T]) UnsafePointer() unsafe.Pointer {
	return ptr.void
}

// Get returns the value of the pointer.
func (ptr Ptr[T]) Get() T {
	return *(*T)(unsafe.Pointer(ptr.void))
}

// Arg adds a new argument to the call frame.
func Arg[T comparable](frame *Frame, arg T) Ptr[T] {
	size := unsafe.Sizeof(arg) / unsafe.Sizeof(uintptr(0))

	if uintptr(frame.idx)+size >= uintptr(len(frame.buf)) {
		copy := arg
		ptr := Ptr[T]{void: unsafe.Pointer(&copy)}
		frame.pin.Pin(ptr)
		return ptr
	}
	if frame.len >= uint8(len(frame.ptr)) {
		// FIXME when frame.len is too big, we still need a continous set of pointers for the Array method.
		panic("too many arguments for call.Frame FIXME")
	}
	copy(frame.buf[frame.idx:], unsafe.Slice((*uintptr)(unsafe.Pointer(&arg)), size))
	frame.typ[frame.len] = reflect.TypeOf(arg)
	frame.ptr[frame.len] = &frame.buf[frame.idx]
	frame.len++
	frame.idx += uint8(size)
	return Ptr[T]{void: unsafe.Pointer(frame.ptr[frame.len-1])}
}

// Ret prepares an expected return value that will be available after the call has been
// made.
func Ret[T comparable](frame *Frame) Ptr[T] {
	var zero T
	size := unsafe.Sizeof(zero) / unsafe.Sizeof(uintptr(0))
	if uintptr(frame.idx)+size >= uintptr(len(frame.buf)) || frame.len >= uint8(len(frame.ptr)) {
		ptr := Ptr[T]{void: unsafe.Pointer(new(T))}
		frame.pin.Pin(ptr)
		return ptr
	}
	*(*T)(unsafe.Pointer(&frame.buf[frame.idx])) = zero
	frame.idx += uint8(size)
	return Ptr[T]{void: unsafe.Pointer(&frame.buf[frame.idx-uint8(size)])}
}
