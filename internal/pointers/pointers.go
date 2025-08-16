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
//	Lay(ptr) // converts a [New] pointer to a [Let] pointer.
//
// [Cycle] should be called periodically to invalidate any pointer-related resources for re-use. Invalidated pointers
// will eventually have their [Free] method called as long as the program continues to construct new pointers of the same types.
//
// The [For] function returns an iterator for all pointers of a given type, this can be used to free all pointers of a given type.
// Up to 16 pointer types for each shape are currently supported.
package pointers

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

const debug = false

const pageSize = 3 * 4 * 5 * 20
const shapesMax = 7

// tables stores the pointers that are managed by the pointers package.
// there is one table for each pointer shape, followed by the pointers
// themselves which have the following structure:
//
//	type entry[T Size] struct {
//		revision revision  // the current revision number.
//		freefunc func(any) // unsafe pointer to the free function method
//		pointers T		   // the actual pointer values.
//	 }
var tables [shapesMax]atomicSlice[[pageSize]atomic.Uint64]

var static [shapesMax][][pageSize]uint64
var counts [shapesMax]atomic.Uint64 // static counts

const (
	offsetRevision = iota
	offsetFreeFunc
	offsetPointers
)

// revision number, two most significant bits (in order of most to least) are used
// to indicate whether the pointer is pinned and whether or not it has been used in
// the last cycle.
type revision uint64

const (
	revisionEOF    = 0
	revisionLocked = 1
)

// matches ignores the most significant 2 bits.
func (r revision) matches(other revision) bool {
	return r&0b0011111111111111111111111111111111111111111111111111111111111111 == other&0b001111111111111111111111111111111111111111111111111111111111111
}

// isPinned returns true if the pointer is pinned.
// (the most significant bit is set).
func (r revision) isPinned() bool {
	return r&(1<<63) != 0
}

// isActive returns true if the pointer was used in the last cycle.
func (r revision) isActive() bool {
	return r&(1<<62) != 0 || r.isPinned()
}

// active returns an active revision.
func (r revision) active() revision {
	return r | 1<<62
}

func (r revision) isClosed() bool  { return r == 0 || r&(1<<61) != 0 }
func (r revision) close() revision { return r | 1<<61 }
func (r revision) reset() revision { return r &^ (1 << 61) }

// expire an active pointer, so that
// it will need to be marked as active
// again to prevent it from being
// freed in the next cycle.
func (r revision) expire() revision {
	return r &^ (1 << 62)
}

// pinned returns a pinned revision.
func (r revision) pinned() revision {
	return r | 1<<63
}

// unpinned returns an unpinned revision.
func (r revision) unpinned() revision {
	return r &^ 1 << 63
}

// writes stores a small ring of write indicies for each global, the garbage
// collection will cycle through these indicies to determine the write head
// for each table. Any entries -2 or less will be freed.
var writes [shapesMax]atomic.Uint64

// Cycle triggers an deadline garbage collection cycle, to clean up temporary
// objects, only pointers allocated in the current or last cycle will
// be preserved.
func Cycle() {
	for s := range shapesMax {
		tab := &tables[s]
		for j := range tab.len.Load() {
			page := tab.Index(j)
			for i := uint64(0); i < pageSize; i += uint64(s + 2) {
				rev := revision(page[i+offsetRevision].Load())
				if rev == revisionEOF {
					break // end of the table.
				}
				if rev.isClosed() {
					continue
				}
				if rev.isActive() {
					if !rev.isPinned() {
						page[i+offsetRevision].CompareAndSwap(uint64(rev), uint64(rev.expire()))
					}
				} else {
					jump := uintptr(page[i+offsetFreeFunc].Load())
					if jump == 0 {
						end(rev, s, uint64(j*pageSize+i))
					} else if page[i+offsetFreeFunc].CompareAndSwap(uint64(jump), 0) {
						switch s {
						case 1:
							type Solo struct {
								sentinal uint64
								revision revision
								checksum [1]uint64
							}
							free := *(*func(Solo))(unsafe.Pointer(&jump))
							if debug {
								pc := reflect.ValueOf(free).Pointer()
								fmt.Println(runtime.FuncForPC(pc).FileLine(pc))
							}
							free(Solo{
								sentinal: j*pageSize + i,
								revision: rev,
								checksum: [1]uint64{page[i+offsetPointers].Load()},
							})
						case 2:
							type Pair struct {
								sentinal uint64
								revision revision
								checksum [2]uint64
							}
							free := *(*func(Pair))(unsafe.Pointer(&jump))
							if debug {
								pc := reflect.ValueOf(free).Pointer()
								fmt.Println(runtime.FuncForPC(pc).FileLine(pc))
							}
							free(Pair{
								sentinal: j*pageSize + i,
								revision: revision(page[i+offsetRevision].Load()),
								checksum: [2]uint64{
									page[i+offsetPointers].Load(),
									page[i+offsetPointers+1].Load(),
								},
							})
						case 3:
							type Trio struct {
								sentinal uint64
								revision revision
								checksum [3]uint64
							}
							free := *(*func(Trio))(unsafe.Pointer(&jump))
							if debug {
								pc := reflect.ValueOf(free).Pointer()
								fmt.Println(runtime.FuncForPC(pc).FileLine(pc))
							}
							free(Trio{
								sentinal: j*pageSize + i,
								revision: revision(page[i+offsetRevision].Load()),
								checksum: [3]uint64{
									page[i+offsetPointers].Load(),
									page[i+offsetPointers+1].Load(),
									page[i+offsetPointers+2].Load(),
								},
							})
						}
						// TODO confirm that [End] was called?
					}
				}
			}
		}
	}
}

// New manages the given pointer value discretely.
func New[T Generic[T, P], P Size](ptr P) T {
	return malloc(ptr, T.Free)
}

// Let returns a borrowed pointer, which we will not be freed when [End] is called.
func Let[T Generic[T, P], P Size](ptr P) T {
	return malloc(ptr, (func(T))(nil))
}

func malloc[T Generic[T, P], P Size](ptr P, free func(T)) T {
	tab := &tables[len(ptr)]
	wat := &writes[len(ptr)]
	for {
		idx := wat.Load()
		page, addr := idx/pageSize, idx%pageSize
		arr := tab.Index(page)
		rev := revision(arr[addr+offsetRevision].Load())
		nxt := arr[addr+offsetPointers].Load()
		if rev == 0 {
			nxt = idx + offsetPointers + uint64(len(ptr))
		}
		if wat.CompareAndSwap(idx, nxt) && rev != revisionLocked && rev.isClosed() && arr[addr+offsetRevision].CompareAndSwap(uint64(rev), revisionLocked) {
			var current struct {
				_ [0]*T

				sentinal uint64
				revision revision
				checksum P
			}
			current.sentinal = idx
			rev := (max(rev, 2) + 1).active().reset()
			arr[addr+offsetRevision].Store(uint64(rev))
			//
			// NOTE the below function extraction is somewhat unsafe and
			// relies on specific assumptions on how static function pointers
			// are represented by the Go compiler. Tests should catch when
			// this assumption breaks. We do this to avoid allocations and
			// to avoid the Go runtime from pointer-scanning our [tables].
			//
			var (
				free = free
				jump = uint64(uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&free))))
			)
			var (
				local [3]uint64
			)
			*(*P)(unsafe.Pointer(&local)) = ptr
			arr[addr+offsetFreeFunc].Store(jump)
			for i := range uint64(len(ptr)) {
				arr[addr+offsetPointers+i].Store(local[i])
			}
			current.revision = rev
			current.checksum = ptr
			return T(current)
		}
	}
}

func end(rev revision, s int, p uint64) bool {
	if rev < 2 {
		return false
	}
	page, addr := uint64(p/pageSize), uint64(p%pageSize)
	arr := tables[s].Index(page)
	for {
		existing := revision(arr[addr+offsetRevision].Load())
		if !rev.matches(existing) {
			return false
		}
		if existing != revisionLocked && arr[addr+offsetRevision].CompareAndSwap(uint64(existing), revisionLocked) {
			for {
				end := writes[s].Load()
				arr[addr+offsetPointers].Store(end)
				if writes[s].CompareAndSwap(end, p) {
					arr[addr+offsetRevision].Store(uint64(existing.close())) // next free.
					return true
				}
			}
		}
	}
}

// Cut either gets the pointer, or ends it, depending on the parameter passed.
func Cut[T Generic[T, P], P Size](ptr T, end bool) P {
	if end {
		raw, ok := End(ptr)
		if !ok {
			panic("pointer already freed")
		}
		return raw
	}
	return Get[T, P](ptr)
}

// Get returns the underlying pointer value, or panics if the pointer
// has been freed.
func Get[T Generic[T, P], P Size](ptr T) P {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	})(ptr)
	if p.revision == 0 && p.sentinal == 0 {
		return p.checksum
	}
	if p.revision == 0 {
		var result [3]uint64
		for i := 0; i < len(p.checksum); i++ {
			result[i] = static[len(p.checksum)][p.sentinal/pageSize][uintptr(p.sentinal)%pageSize+uintptr(i)]
		}
		return *(*P)(unsafe.Pointer(&result))
	}
	page, addr := uint64(p.sentinal/pageSize), uint64(p.sentinal%pageSize)
	arr := tables[len(p.checksum)].Index(page)
	var ptrs [3]uint64
	for i := 0; i < len(p.checksum); i++ {
		ptrs[i] = arr[addr+offsetPointers+uint64(i)].Load()
	}
	for {
		rev := revision(arr[addr+offsetRevision].Load())
		if rev == revisionLocked {
			continue
		}
		if !rev.matches(p.revision) {
			fmt.Printf("%b != %b\b", rev&0b00111111111111111111111111111111111111111111111111111111111111, p.revision&0b00111111111111111111111111111111111111111111111111111111111111)
			panic("expired pointer")
		}
		if !rev.isActive() {
			if live, ok := any(T(p)).(Liveness[P]); ok && !live.IsAlive(*(*P)(unsafe.Pointer(&ptrs))) {
				panic("dead pointer")
			}
			arr[addr+offsetRevision].CompareAndSwap(uint64(rev), uint64(rev.active()))
		}
		return *(*P)(unsafe.Pointer(&ptrs))
	}
}

// Bad returns true if the pointer is nil, or freed.
func Bad[T Generic[T, P], P Size](ptr T) bool {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	})(ptr)
	if p.revision == 0 && p.sentinal == 0 {
		return p.checksum != [1]P{}[0]
	}
	if p.revision == 0 {
		return true
	}
	page, addr := uint64(p.sentinal/pageSize), uint64(p.sentinal%pageSize)
	arr := tables[len(p.checksum)].Index(page)
	rev := revision(arr[addr+offsetRevision].Load())
	if rev.matches(p.revision) {
		if !rev.isActive() {
			arr[addr+offsetRevision].CompareAndSwap(uint64(rev), uint64(rev.active()))
		}
		return true
	}
	return false
}

// Add allocates a new pointer that can be mutated with [Set].
func Add[T Generic[T, P], P Size](val P) T {
	addr := counts[len(val)].Add(uint64(len(val)))
	var result struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	}
	if len(static[len(val)]) <= int(addr/pageSize) {
		static[len(val)] = append(static[len(val)], [pageSize]uint64{})
	}
	var local [3]uint64
	*(*P)(unsafe.Pointer(&local)) = val
	for i := 0; i < len(val); i++ {
		static[len(val)][addr/pageSize][addr%pageSize+uint64(i)] = uint64(local[i])
	}
	result.sentinal = addr
	return T(result)
}

// Set overwrites the underlying added pointer value.
func Set[T Generic[T, P], P Size](ptr T, val P) {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	})(ptr)
	if p.revision == 0 && p.sentinal == 0 {
		return
	}
	if p.revision == 0 {
		var local [3]uint64
		*(*P)(unsafe.Pointer(&local)) = val
		for i := 0; i < len(val); i++ {
			static[len(val)][p.sentinal/pageSize][uint64(p.sentinal)%pageSize+uint64(i)] = uint64(local[i])
		}
		return
	}
	page, addr := uint64(p.sentinal/pageSize), uint64(p.sentinal%pageSize)
	arr := tables[len(p.checksum)].Index(page)
	for {
		rev := revision(arr[addr+offsetRevision].Load())
		if !rev.matches(p.revision) {
			panic("expired pointer")
		}
		if rev != revisionLocked && arr[addr+offsetRevision].CompareAndSwap(uint64(rev), revisionLocked) {
			var local [3]uint64
			*(*P)(unsafe.Pointer(&local)) = val
			for i := 0; i < len(val); i++ {
				arr[addr+offsetPointers+uint64(i)].Store(uint64(local[i]))
			}
			arr[addr+offsetRevision].Store(uint64(rev.active()))
			return
		}
	}
}

type atomicSlice[T any] struct {
	mut sync.Mutex // only locked for appends.
	ptr atomic.Pointer[*T]
	len atomic.Uint64
}

// Index returns the element at the given index, extending the slice if necessary.
func (s *atomicSlice[T]) Index(i uint64) *T {
	l := s.len.Load()
	old := unsafe.Slice(s.ptr.Load(), l)
	if uint64(i) >= l {
		s.mut.Lock()
		defer s.mut.Unlock()
		l = s.len.Load()
		if !(uint64(i) >= l) {
			old := unsafe.Slice(s.ptr.Load(), l)
			return old[i]
		}
		end := new(T)
		new := make([]*T, len(old)+1)
		copy(new, old)
		new[len(old)] = end
		s.ptr.Store(&new[0])
		s.len.Store(uint64(len(new)))
		return end
	}
	return old[i]
}

// Raw returns an [unsafe.Pointer] equivalent to [Pointer], zero overhead, but it provides no protections on use.
func Raw[T Generic[T, P], P Size](ptr P) T {
	var result struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	}
	result.checksum = ptr
	return T(result)
}

// Pin the pointer, preventing it from being freed until [Free] is called on it.
func Pin[T Generic[T, P], P Size](ptr T) T {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	})(ptr)
	if (p.revision == 0 && p.sentinal == 0) || p.revision.isPinned() {
		return ptr
	}
	if p.revision == 0 {
		panic("cannot pin a nil pointer")
	}
	page, addr := uint64(p.sentinal/pageSize), uint64(p.sentinal%pageSize)
	arr := tables[len(p.checksum)].Index(page)
	rev := revision(arr[addr+offsetRevision].Load())
	if !rev.matches(p.revision) {
		panic("expired pointer")
	}
	arr[addr+offsetRevision].CompareAndSwap(uint64(rev), uint64(rev.pinned()))
	return ptr
}

// Debug prints the current state of the pointer.
func Debug[T Generic[T, P], P Size](ptr T) {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	})(ptr)
	if p.revision == 0 && p.sentinal == 0 {
		fmt.Printf("pointer raw %d\n", p.sentinal)
	}
	if p.revision == 0 {
		fmt.Printf("pointer static %d\n", p.sentinal)
	}
	page, addr := uint64(p.sentinal/pageSize), uint64(p.sentinal%pageSize)
	arr := tables[len(p.checksum)].Index(page)
	rev := revision(arr[addr+offsetRevision].Load())
	free := arr[addr+offsetFreeFunc].Load()
	fmt.Printf("pointer %d active %v closed %v freefunc %v\n", p.sentinal, rev.isActive(), rev.isClosed(), free != 0)
}

// Lay the pointer, preventing the free operation from being called on it (it can still expire).
func Lay[T Generic[T, P], P Size](ptr T) T {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	})(ptr)
	if p.revision == 0 && p.sentinal == 0 {
		return ptr
	}
	if p.revision == 0 {
		panic("cannot let a nil pointer")
	}
	page, addr := uint64(p.sentinal/pageSize), uint64(p.sentinal%pageSize)
	arr := tables[len(p.checksum)].Index(page)
	for {
		rev := revision(arr[addr+offsetRevision].Load())
		if !rev.matches(p.revision) {
			panic("expired pointer")
		}
		if rev != revisionLocked && arr[addr+offsetRevision].CompareAndSwap(uint64(rev), revisionLocked) {
			arr[addr+offsetFreeFunc].Store(0)
			arr[addr+offsetRevision].Store(uint64(rev.active()))
			return ptr
		}
	}
}

func Pack[T Generic[T, P], P Size](ptr T) complex128 {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	})(ptr)
	return *(*complex128)(unsafe.Pointer(&p.sentinal))
}

func Load[T Generic[T, P], P Size](data complex128) T {
	var result struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum P
	}
	*(*complex128)(unsafe.Pointer(&result.sentinal)) = data
	result.checksum = Get[T](result)
	return T(result)
}

// Size of a pointer up to [3]uintptr's, suitable for supporting fat pointers.
type Size interface {
	[1]uint32 | [1]uint64 | [2]uint64 | [3]uint64
}

// Generic pointer.
type Generic[T any, S Size] interface {
	~struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum S
	}

	Free()
}

type Liveness[S Size] interface {
	IsAlive(raw S) bool
}

// End the lifetime of the pointer, returning the underlying pointer value
// and ok=true if this is the first time the pointer has been freed.
func End[T Generic[T, Raw], Raw Size](ptr T) (Raw, bool) {
	p := (struct {
		_ [0]*T

		sentinal uint64
		revision revision
		checksum Raw
	})(ptr)
	if p.checksum == [1]Raw{}[0] {
		return [1]Raw{}[0], false
	}
	if end(p.revision, len(p.checksum), uint64(p.sentinal)) {
		return p.checksum, true
	}
	return [1]Raw{}[0], false
}

// Half pointer value that safely wraps a single uintptr-sized value.
type Half[T Generic[T, [1]uint32]] struct {
	_ [0]*T // prevents converting between different pointer types.

	// if both 'sentinal' and 'revision' fields are zero, then this is
	// an unsafe pointer and the checksum value is used directly, in
	// this case, no memory-safety protections are provided.
	//
	// if the 'sentinal' is non-zero but the 'revision' is zero, then
	// this is a static pointer, it cannot be freed and the checksum
	// value is ignored.
	sentinal uint64
	revision revision
	checksum [1]uint32
}

// Solo pointer value that safely wraps a single uintptr-sized value.
type Solo[T Generic[T, [1]uint64]] struct {
	_ [0]*T // prevents converting between different pointer types.

	// if both 'sentinal' and 'revision' fields are zero, then this is
	// an unsafe pointer and the checksum value is used directly, in
	// this case, no memory-safety protections are provided.
	//
	// if the 'sentinal' is non-zero but the 'revision' is zero, then
	// this is a static pointer, it cannot be freed and the checksum
	// value is ignored.
	sentinal uint64
	revision revision
	checksum [1]uint64
}

// Pair pointer value that safely wraps a pair of uintptr-sized values.
type Pair[T Generic[T, [2]uint64]] struct {
	_ [0]*T // prevents converting between different pointer types.

	// if both 'sentinal' and 'revision' fields are zero, then this is
	// an unsafe pointer and the checksum value is used directly, in
	// this case, no memory-safety protections are provided.
	//
	// if the 'sentinal' is non-zero but the 'revision' is zero, then
	// this is a static pointer, it cannot be freed and the checksum
	// value is ignored.
	sentinal uint64
	revision revision
	checksum [2]uint64
}

// Trio pointer value that safely wraps three uintptr-sized values.
type Trio[T Generic[T, [3]uint64]] struct {
	_ [0]*T // prevents converting between different pointer types.

	// if both 'sentinal' and 'revision' fields are zero, then this is
	// an unsafe pointer and the checksum value is used directly, in
	// this case, no memory-safety protections are provided.
	//
	// if the 'sentinal' is non-zero but the 'revision' is zero, then
	// this is a static pointer, it cannot be freed and the checksum
	// value is ignored.
	sentinal uint64
	revision revision
	checksum [3]uint64
}
