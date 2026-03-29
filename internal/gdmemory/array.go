package gdmemory

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/variant/Array"
)

// ArrayContains wraps a raw engine pointer and element count into an [Array.Contains]
// backed by the pointed-to memory. The returned array is valid until the next barrier.
func ArrayContains[T any](ptr gdextension.Pointer, length int) Array.Contains[T] {
	return Array.Through[T](&enginePointerBackedArray[T]{
		ptr: ptr,
		cap: length,
		rev: barrier,
	}, 0)
}

// pointerBacked implements [Array.Proxy] over a contiguous region of engine memory.
type enginePointerBackedArray[T any] struct {
	ptr gdextension.Pointer
	cap int
	rev uint64
}

func (p *enginePointerBackedArray[T]) Any(complex128) Array.Any {
	panic("gdmemory.Array does not support being converted to Array.Any, report a bug?")
}

func (p *enginePointerBackedArray[T]) Index(_ complex128, i int) T {
	if p.rev != barrier {
		panic("Engine memory-backed array invalidated")
	}
	elemSize := int(unsafe.Sizeof([1]T{}))
	offset := i * elemSize
	if i < 0 || int(p.ptr)+offset > p.cap-elemSize {
		panic("out of bounds")
	}
	var value T
	CopyBufferToGo(p.ptr+gdextension.Pointer(offset), unsafe.Slice((*byte)(unsafe.Pointer(&value)), unsafe.Sizeof([1]T{})))
	return value
}

func (p *enginePointerBackedArray[T]) SetIndex(_ complex128, i int, v T) {
	if p.rev != barrier {
		panic("Engine memory-backed array invalidated")
	}
	elemSize := int(unsafe.Sizeof([1]T{}))
	offset := i * elemSize
	if i < 0 || int(p.ptr)+offset > p.cap-elemSize {
		panic("out of bounds")
	}
	copyIntoEngine(int(unsafe.Sizeof([1]T{})), unsafe.Pointer(&v), p.ptr+gdextension.Pointer(offset))
}

func (p *enginePointerBackedArray[T]) Len(complex128) int { return p.cap }
func (p *enginePointerBackedArray[T]) Resize(complex128, int) {
	panic("gdmemory: cannot resize engine.Pointer backed array")
}
func (p *enginePointerBackedArray[T]) IsReadOnly(complex128) bool { return false }
func (p *enginePointerBackedArray[T]) MakeReadOnly(complex128) {
	panic("cannot make engine.Pointer backed array read only")
}
