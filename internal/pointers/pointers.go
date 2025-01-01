// Package pointers provides managed pointers that are invisible to the Go runtime.
//
// Usually, the Go runtime attempts to scan all pointers within a program and as such,
// the more pointers in use, the more overhead there will be to scan them.
//
// When working with fixed pointers allocated outside of the runtime, the naive approach
// would be to use a uintptr instead of a typed pointer but this can easily lead right
// into unsafe and undefined behavior as well as memory leaks. [runtime.SetFinalizer] is
// another option, except, this requires additional allocations and overhead that makes it
// unsuitable for scenarios where thousands of pointers are in use.
//
// Pointers exclusively managed through this package are protected against double frees,
// use-after-free and leaks. Runtime panics will be raised if any unsafe use is detected.
//
// The following constructors are available:
//
//	New(ptr) // panics-on-use after two garbage collection cycles.
//	Let(ptr) // like [New], but [End] operations on this pointer always return false.
//	Raw(ptr) // unsafe.Pointer equivalent, zero overhead, but provides no protection.
//	Add(ptr) // allocates and returns a static pointer that can be mutated with Set.
//
// The following methods are available:
//
//	Set(X, ptr) // sets and returns the mutable pointer at slot N.
//	End(X) (ptr, bool) // returns true on the first call to [End], false on all subsequent calls, GC-related metadata is marked for reuse.
//	Use(ptr) // refreshes a [New] pointer, as if it were just allocated, similar (in some sense) to [runtime.KeepAlive].
//	Pin(ptr) // pin a [New] pointer, so that it will be available for use until the next call to [Free]
//	Bad(ptr) // returns true if the pointer is nil, or freed.
//
// [Cycle] should be called periodically to invalidate any pointer-related resources for re-use. Invalidated pointers
// will eventually have their [Free] method called as long as the program continues to construct new pointers of the same types.
//
// The [For] function returns an iterator for all pointers of a given type, this can be used to free all pointers of a given type.
// Up to 16 pointer types for each shape are currently supported.
package pointers

import (
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"
)

const pageSize = 3 * 4 * 5 * 20
const shapesMax = 4
const cyclesMax = 4
const uniqueMax = 16

// tables stores the pointers that are managed by the pointers package.
// there is one table for each pointer shape, followed by the pointers
// themselves which have the following structure:
//
//	type entry[T Shape] struct {
//		rev uintptr // when 1, locks the entry, when 0, the entry is free, when 2, pinned.
//		age uintptr // when 1, locks the entry, when 2, queues mark, 3 and 4 are used to track marking.
//		ptr [unsafe.Sizeof([1]T{})/unsafe.Sizeof(uintptr(0))]uintptr
//	 }
var tables [uniqueMax][shapesMax]atomicSlice[[pageSize]atomic.Uintptr]

var static [shapesMax][][pageSize]uintptr
var counts [shapesMax]atomic.Uintptr // static counts

const (
	offRev = iota
	offAge
	offPtr
)

// writes stores a small ring of write indicies for each global, the garbage
// collection will cycle through these indicies to determine the write head
// for each table. Any entries -2 or less will be freed.
var writes [uniqueMax][shapesMax]atomic.Uintptr

// current cycle number, used as the first index into [writes].
var cyclen atomic.Uintptr

func init() {
	cyclen.Store(5)
}

// Cycle triggers an deadline garbage collection cycle, to clean up temporary
// objects, only pointers allocated in the current or last cycle will
// be preserved.
func Cycle() {
	cyclen.Add(1)
}

// Root is the root of the object graph. Pointers must exist within this
// value to be considered reachable. Roots must either be thread-safe to
// read or implement the [sync.Locker] interface.
var Root any

// MarkAndSweep performs a mark and sweep garbage collection cycle.
// The root parameter is the [Root] of the object graph, all reachable
// objects will be marked and all unmarked objects will be freed.
func MarkAndSweep() {
	now := cyclen.Load()
	mark(now, false, reflect.ValueOf(Root))
	var count int

	for u := range uniqueMax {
		for s := range shapesMax {
			tab := &tables[u][s]
			for j := range tab.len.Load() {
				page := tab.Index(j)
				for i := uintptr(0); i < pageSize; i += uintptr(s + 2) {
					age := page[i+offAge].Load()
					if age == 0 {
						break
					}
					if age > 2 && age-2 < now-2 {
						end(page[i+offRev].Load(), u, s, j*pageSize+i)
						count++
					}
				}
			}
		}
	}
}

var typeLocker = reflect.TypeOf([0]sync.Locker{}).Elem()
var typePointer = reflect.TypeOf([0]PointerAny{}).Elem()

func mark(now uintptr, locked bool, rvalue reflect.Value) {
	if !rvalue.IsValid() || rvalue.IsNil() {
		return
	}
	if !locked && rvalue.Type().Implements(typeLocker) {
		rvalue.Interface().(sync.Locker).Lock()
		defer rvalue.Interface().(sync.Locker).Unlock()
		locked = true
	}
	switch rvalue.Kind() {
	case reflect.Pointer:
		if !rvalue.IsNil() {
			mark(now, locked, rvalue.Elem())
			return
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rvalue.Len(); i++ {
			mark(now, locked, rvalue.Index(i))
		}
	case reflect.Map:
		for _, key := range rvalue.MapKeys() {
			mark(now, locked, rvalue.MapIndex(key))
		}
		for _, key := range rvalue.MapKeys() {
			mark(now, locked, key)
		}
	case reflect.Struct:
		if rvalue.Type().Implements(typePointer) {
			ptr := rvalue.Interface().(PointerAny)
			addr, shape, ptype := ptr.pointer()
			tab := &tables[ptype][shape]
			page, off := addr/pageSize, addr%pageSize
			arr := tab.Index(page)
			arr[off+offAge].Store(now)
		}
		for i := 0; i < rvalue.NumField(); i++ {
			mark(now, locked, rvalue.Field(i))
		}
	}
}

// Shape of up to [3]uintptr's, suitable for supporting fat pointers.
type Shape interface {
	[1]uintptr | [2]uintptr | [3]uintptr
}

// PointerNamed that is invisible to the Go runtime. The pointers package manages it instead.
// Supports pointer shapes up to [3]uintptr's in size. The first type parameter should be the
// named type being defined as a pointers pointer type.
type PointerNamed[T Pointer[T, S, N], S Shape, N PointerType] struct {
	_ [0]N // prevents converting between different pointer types.

	// if both the address and the revision are zero, then this is an
	// unsafe pointer and the ptr value is exposed directly, in this
	// case, no protections are provided.
	//
	// if the address is non-zero and the revision is zero, then this
	// is a static pointer, it cannot be freed and will the ptr field
	// is ignored.

	address[S, N]

	rev uintptr // revision counter.
	ptr S       // pointer checksum.
}

type samplePointer PointerNamed[samplePointer, [1]uintptr, [1]Type]

var _ PointerAny = samplePointer{}

func (p samplePointer) Free() { End(p) }

type address[S Shape, N PointerType] uintptr

func (a address[S, N]) pointer() (addr, shape, ptype uintptr) {
	return uintptr(a), uintptr(len([1]S{}[0])), uintptr(len([1]N{}[0]))
}

type PointerAny interface {
	pointer() (addr, shape, ptype uintptr)
}

// Pointer is the underlying type for a [PointerNamed].
type Pointer[T any, S Shape, N PointerType] interface {
	~struct {
		_ [0]N

		address[S, N]

		rev uintptr
		ptr S
	}
	Free()
}

// New manages the given pointer value discretely.
func New[T Pointer[T, P, N], P Shape, N PointerType](ptr P) T {
	now := cyclen.Load()
	tab := &tables[len([1]N{}[0])][len(ptr)]
	wat := &writes[len([1]N{}[0])][len(ptr)]
	for {
		idx := wat.Load()
		page, addr := idx/pageSize, idx%pageSize
		arr := tab.Index(page)
		age := arr[addr+offAge].Load()
		rev := arr[addr+offRev].Load()
		nxt := age
		if rev == 0 || rev != 2 {
			nxt = idx + offPtr + uintptr(len(ptr))
		}
		if wat.CompareAndSwap(idx, nxt) && age != 1 && arr[addr+offAge].CompareAndSwap(age, 1) {
			var current struct {
				_ [0]N

				address[P, N]

				rev uintptr
				ptr P
			}
			current.address = address[P, N](idx)
			if rev > 2 {
				current.rev = arr[addr+offRev].Load()
				for i := range uintptr(len(ptr)) {
					current.ptr[i] = arr[addr+offPtr+i].Load()
				}
				T(current).Free()
				if arr[addr+offRev].Load() != 1 {
					panic("bad free")
				}
			}
			arr[addr+offAge].Store(now)
			arr[addr+offRev].Store(now)
			for i := range uintptr(len(ptr)) {
				arr[addr+offPtr+i].Store(uintptr(ptr[i]))
			}
			current.rev = now
			current.ptr = ptr
			return T(current)
		}
	}
}

// End the lifetime of the pointer, returning the underlying pointer value
// and ok=true if this is the first time the pointer has been freed.
func End[T Pointer[T, P, N], P Shape, N PointerType](ptr T) (P, bool) {
	p := (struct {
		_ [0]N

		address[P, N]

		rev uintptr
		ptr P
	})(ptr)
	if p.ptr == [1]P{}[0] {
		return [1]P{}[0], false
	}
	if end(p.rev, len([1]N{}[0]), len(p.ptr), uintptr(p.address)) {
		return p.ptr, true
	}
	return [1]P{}[0], false
}

func end(rev uintptr, u, s int, p uintptr) bool {
	page, addr := uintptr(p/pageSize), uintptr(p%pageSize)
	arr := tables[u][s].Index(page)
	if arr[addr+offRev].Load() > 1 && arr[addr+offRev].CompareAndSwap(rev, 1) {
		age := arr[addr+offAge].Load()
		if age == 1 {
			return true
		}
		arr[addr+offRev].Store(2) // next free.
		for {
			end := writes[u][s].Load()
			arr[addr+offAge].Store(end)
			if writes[u][s].CompareAndSwap(end, p) {
				return true
			}
		}
	}
	return false
}

// Pin the pointer, preventing it from being freed until the next mark
// and sweep cycle. The pointer must be placed in a root in order for
// it to remain reachable.

// Get returns the underlying pointer value, or panics if the pointer
// has been freed.
func Get[T Pointer[T, P, N], P Shape, N PointerType](ptr T) P {
	p := (struct {
		_ [0]N

		address[P, N]

		rev uintptr
		ptr P
	})(ptr)
	if p.rev == 0 && p.address == 0 {
		return p.ptr
	}
	if p.rev == 0 {
		var result P
		for i := 0; i < len(p.ptr); i++ {
			result[i] = static[len(p.ptr)][p.address/pageSize][uintptr(p.address)%pageSize+uintptr(i)]
		}
		return result
	}
	page, addr := uintptr(p.address/pageSize), uintptr(p.address%pageSize)
	arr := tables[len([1]N{}[0])][len(p.ptr)].Index(page)
	rev := arr[addr+offRev].Load()
	if rev != 2 && rev != uintptr(p.rev) {
		panic("expired pointer")
	}
	for i := range uintptr(len(p.ptr)) {
		if arr[addr+offPtr+i].Load() != p.ptr[i] {
			panic("expired pointer")
		}
	}
	return p.ptr
}

// Add allocates a new pointer that can be mutated with [Set].
func Add[T Pointer[T, P, N], P Shape, N PointerType](val P) T {
	addr := counts[len(val)].Add(uintptr(len(val)))
	var result struct {
		_ [0]N

		address[P, N]

		rev uintptr
		ptr P
	}
	if len(static[len(val)]) <= int(addr/pageSize) {
		static[len(val)] = append(static[len(val)], [pageSize]uintptr{})
	}
	for i := 0; i < len(val); i++ {
		static[len(val)][addr/pageSize][addr%pageSize+uintptr(i)] = uintptr(val[i])
	}
	result.address = address[P, N](addr)
	return T(result)
}

// Set overwrites the underlying added pointer value.
func Set[T Pointer[T, P, N], P Shape, N PointerType](ptr T, val P) {
	p := (struct {
		_ [0]N

		address[P, N]

		rev uintptr
		ptr P
	})(ptr)
	for i := 0; i < len(val); i++ {
		static[len(val)][p.address/pageSize][uintptr(p.address)%pageSize+uintptr(i)] = uintptr(val[i])
	}
}

// PointerType restricts the unique pointer types that can be used with the pointers package.
// A unique pointer type is a pointer
type PointerType interface {
	[0]Type | [1]Type | [2]Type | [3]Type |
		[4]Type | [5]Type | [6]Type | [7]Type |
		[8]Type | [9]Type | [10]Type | [11]Type |
		[12]Type | [13]Type | [14]Type | [15]Type
}

// Type ID, should be prefixed by an array size.
type Type struct{}

type atomicSlice[T any] struct {
	mut sync.Mutex // only locked for appends.
	ptr atomic.Pointer[*T]
	len atomic.Uintptr
}

// Index returns the element at the given index, extending the slice if necessary.
func (s *atomicSlice[T]) Index(i uintptr) *T {
	l := s.len.Load()
	old := unsafe.Slice(s.ptr.Load(), l)
	if uintptr(i) >= l {
		s.mut.Lock()
		defer s.mut.Unlock()
		l = s.len.Load()
		if !(uintptr(i) >= l) {
			old := unsafe.Slice(s.ptr.Load(), l)
			return old[i]
		}
		end := new(T)
		new := make([]*T, len(old)+1)
		copy(new, old)
		new[len(old)] = end
		s.ptr.Store(&new[0])
		s.len.Store(uintptr(len(new)))
		return end
	}
	return old[i]
}

// Raw returns an [unsafe.Pointer] equivalent to [Pointer], zero overhead, but it provides no protections on use.
func Raw[T Pointer[T, P, N], P Shape, N PointerType](ptr P) T {
	var result struct {
		_ [0]N

		address[P, N]

		rev uintptr
		ptr P
	}
	result.ptr = ptr
	return T(result)
}

// Pin the pointer, preventing it from being freed until [Free] is called on it.
func Pin[T Pointer[T, P, N], P Shape, N PointerType](ptr T) T {
	p := (struct {
		_ [0]N

		address[P, N]

		rev uintptr
		ptr P
	})(ptr)
	if p.rev == 0 {
		panic("cannot pin a nil pointer")
	}
	page, addr := uintptr(p.address/pageSize), uintptr(p.address%pageSize)
	arr := tables[len([1]N{}[0])][len(p.ptr)].Index(page)
	rev := arr[addr+offRev].Load()
	if rev != uintptr(p.rev) {
		panic("expired pointer")
	}
	for i := range uintptr(len(p.ptr)) {
		if arr[addr+offPtr+i].Load() != p.ptr[i] {
			panic("expired pointer")
		}
	}
	p.rev = 2
	arr[addr+offRev].Store(2)
	return T(p)
}
