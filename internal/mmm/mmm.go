// Package mmm provides a way to manually manage memory and resource lifetimes with protections against unsafe double-free and use-after-free errors.
package mmm

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type revision uintptr

const maxRevision = 4611686018427387904 - 1 // 2**62 - 1

func (rev revision) isPinned() bool {
	return rev&(1<<63) != 0
}

func (rev *revision) pin() {
	*rev |= 1 << 63
}

func (rev revision) isRefCounted() bool {
	return rev&(1<<62) != 0
}

func (rev *revision) setRefCounted() {
	*rev |= 1 << 62
}

func (rev revision) matches(other revision) bool {
	return rev&^(1<<62|1<<63) == other&^(1<<62|1<<63)
}

type pointer[API any, Size PointerSize] struct {
	rev revision
	ref *freeableWith[Size]
}

type genericPointer struct {
	rev revision
	pin unsafe.Pointer
}

// Lifetime group, keeps track of a sequence of [Pointer] values and calls their
// [Free] methods when their lifetime has ended. Not safe for use by multiple
// goroutines. Beware of using [Lifetime] types to manage the lifetime of other
// [Lifetime] values, as this may can lead to reference cycles (a form of manual
// memory management deadlock) that will prevent the lifetime from ever ending.
//
// Usually you'll want to create a [Lifetime] and immediately defer its [End] and
// pass it down the call stack as if it were a [context.Context] (or inside of one).
// If the lifetime is never ended, the underlying [Pointer] values will never
// be freed and will leak.
type Lifetime struct {
	rev    revision // revision number, incremented when freed, enables recycling the root.
	root   *freeable
	allocs *atomic.Int64
	lock   bool
}

var Mutex sync.Mutex

// NewLifetime returns a new lifetime, call [End] to free any [Pointers] associated
// with this lifetime. The Lifetime may no longer be used after calling [End].
func NewLifetime() Lifetime {
	root, ok := roots.Get().(*freeable)
	if !ok {
		root = new(freeable)
		root.prv = root
		root.nxt = root
	}
	root.api = nil
	return Lifetime{
		rev:  root.rev,
		root: root,
	}
}

func (lifetime Lifetime) WithAllocationCounter(counter *atomic.Int64) Lifetime {
	lifetime.allocs = counter
	return lifetime
}

// NewThreadSafeLifetime returns a new lifetime, call [End] to free any [Pointers] associated
// with this lifetime. The Lifetime may no longer be used after calling [End].
func NewThreadSafeLifetime() Lifetime {
	root, ok := threadSafeRoots.Get().(*freeable)
	if !ok {
		root = new(freeable)
		root.prv = root
		root.nxt = root
	}
	root.api = nil
	return Lifetime{
		rev:  root.rev,
		root: root,
		lock: true,
	}
}

// pool for each PointerSize type.
var pools [3 + 1]sync.Pool
var roots sync.Pool
var threadSafeRoots sync.Pool
var refcounters sync.Pool

type rc struct {
	int int
	api unsafe.Pointer
}

func (obj *freeable) free() {
	obj.record("free")
	obj.prv.nxt = obj.nxt
	obj.nxt.prv = obj.prv
	if obj.rev.isRefCounted() {
		rc := (*rc)(obj.api)
		rc.int--
		if rc.int == 0 {
			rc.api = nil
			refcounters.Put(rc)
		}
		if rc.int > 0 {
			obj.end = nil
			return
		}
	}
	obj.end(genericPointer{
		rev: obj.rev,
		pin: unsafe.Pointer(obj),
	})
	if obj.end != nil {
		crash(obj, "grow.graphics/gd/internal/mmm error: pointer escaped from free")
	}
}

type freeableWith[Size PointerSize] struct {
	freeable
	ptr Size
}

func newObjectWith[Size PointerSize]() *freeableWith[Size] {
	var zero Size
	block, ok := pools[unsafe.Sizeof(zero)/unsafe.Sizeof(uintptr(0))].Get().(*freeableWith[Size])
	if !ok {
		return new(freeableWith[Size])
	}
	return block
}

// End the lifetime, freeing any unfreed [Pointers] associated with it. This is
// an idempotent operation, calling it multiple times is safe.
func (lifetime Lifetime) End() {
	if lifetime.lock {
		Mutex.Lock()
		defer Mutex.Unlock()
	}
	if lifetime.root == nil || lifetime.root.rev != lifetime.rev {
		return
	}
	if lifetime.root.api != nil { // not a root.
		lifetime.root.free()
		return
	}
	root := lifetime.root
	for {
		next := root.nxt
		if next.api == nil {
			break
		}
		root.nxt.free()
	}
	root.rev++
	if root.rev < maxRevision {
		if lifetime.lock {
			threadSafeRoots.Put(root)
		} else {
			roots.Put(root)
		}
	}
}

// PointerSize is a valid pointer size.
type PointerSize interface {
	~uintptr | ~[0]uintptr | ~[1]uintptr | ~[2]uintptr | ~[3]uintptr
}

// Pointer of unique type T belonging to the given API, using the given RAM allocator
// that is reponsible for freeing the pointer.
type Pointer[API any, T PointerWithFree[API, T, Size], Size PointerSize] struct {
	_ [0]*T // prevents unsafe type conversions from one pointer type to another.
	pointer[API, Size]
}

type leased[API any, T any, Size PointerSize] Pointer[API, leased[API, T, Size], Size]

func (ptr leased[API, T, Size]) Free() { End(ptr) }

type access[API any, T PointerWithFree[API, T, Size], Size PointerSize] struct {
	_ [0]*T
	pointer[API, Size]
}

// ManagedPointer is any [Pointer] value with a PointerSize of 'Size'.
type ManagedPointer[Size PointerSize] interface {
	readPointer() Size
}

// Get the underlying pointer value from a manually managed pointer. If the pointer
// has already been freed, this function will (safely) panic.
func Get[T ManagedPointer[Size], Size PointerSize](ptr T) Size {
	return ptr.readPointer()
}

func (ptr pointer[API, Size]) readPointer() Size {
	if ptr == (pointer[API, Size]{}) {
		var zero Size
		return zero
	}
	if !ptr.ref.rev.matches(ptr.rev) {
		crash(&ptr.ref.freeable, "grow.graphics/gd/internal/mmm error: use after free")
	}
	return ptr.ref.ptr
}

// Move the given pointer from its previously bound lifetime to a new lifetime, it will no
// longer be freed when its previous lifetime ends and will instead be freed when the new
// one ends. The pointer must have been created with [New] otherwise this function will
// panic. The pointer may not have been copied.
func Move[API any, T PointerWithFree[API, T, Size], Size PointerSize](ptr T, lifetime Lifetime) {
	if lifetime.lock {
		Mutex.Lock()
		defer Mutex.Unlock()
	}
	val := access[API, T, Size](ptr)
	if val.ref.rev.isPinned() {
		crash(&val.ref.freeable, "grow.graphics/gd/internal/mmm error: move after pin")
	}
	if val.ref.rev.isRefCounted() {
		crash(&val.ref.freeable, "grow.graphics/gd/internal/mmm error: move after copy")
	}
	val.ref.record("move")
	internalNew[T](lifetime, getAPI(ptr), End(ptr))
}

// Life returns the lifetime of the given 'pinned' pointer, panics if the pointer is not
// pinned.
func Life[API any, T PointerWithFree[API, T, Size], Size PointerSize](ptr T) Lifetime {
	Mutex.Lock()
	defer Mutex.Unlock()

	val := access[API, T, Size](ptr).ref
	if !val.rev.isPinned() {
		crash(&val.freeable, "grow.graphics/gd/internal/mmm error: unpinned pointer used as a lifetime")
	}
	return Lifetime{
		rev:  val.rev,
		root: &val.freeable,
	}
}

// Copy returns a copy of the given pointer that is bound to a new lifetime. T.Free will
// be called once all copies of the pointer have been freed. The pointer must have been
// created with [New] otherwise this function will panic.
func Copy[API any, T PointerWithFree[API, T, Size], Size PointerSize](ptr T, lifetime Lifetime) T {
	if lifetime.lock {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	val := access[API, T, Size](ptr)
	if val.ref.rev.isPinned() {
		crash(&val.ref.freeable, "grow.graphics/gd/internal/mmm error: copy after pin")
	}
	val.ref.record("copy")
	if !val.ref.rev.isRefCounted() {
		api := val.ref.api
		count, ok := refcounters.Get().(*rc)
		if !ok {
			count = new(rc)
		}
		count.int = 1
		count.api = api
		val.ref.api = unsafe.Pointer(api)
		val.ref.rev.setRefCounted()
	}
	var result = access[API, T, Size](internalNew[T](lifetime, getAPI(ptr), Get(ptr)))
	rc := (*rc)(val.ref.api)
	rc.int++
	result.ref.api = unsafe.Pointer(rc)
	return T(result)
}

// Set the pointer value to the given 'ptr' value, the pointer must have been created
// with [New] otherwise this function will panic.
func Set[API any, T PointerWithFree[API, T, Size], Size PointerSize](ptr *T, value Size) {
	Mutex.Lock()
	defer Mutex.Unlock()

	val := access[API, T, Size](*ptr)
	if !val.ref.rev.matches(val.rev) {
		crash(&val.ref.freeable, "grow.graphics/gd/internal/mmm error: use after free")
	}
	val.ref.record("set")
	val.ref.ptr = value
}

// End the lifetime of a pointer, this function must be called from within the Free
// method of a [PointerWithFree]. Does not free any underlying resources associated
// with the pointer, use [Free] for that.
func End[API any, T PointerWithFree[API, T, Size], Size PointerSize](ptr T) Size {
	var zero Size
	var empty T
	if ptr == empty {
		return zero
	}
	val := access[API, T, Size](ptr)
	if !val.ref.rev.matches(val.rev) {
		crash(&val.ref.freeable, "grow.graphics/gd/internal/mmm error: use after free")
	}
	val.ref.record("end")
	tmp := val.ref.ptr
	val.ref.end = nil
	if val.ref.ptr != zero {
		val.ref.ptr = zero
		val.ref.rev++
	}
	prv := val.ref.prv
	nxt := val.ref.nxt
	prv.nxt = nxt
	nxt.prv = prv
	if val.ref.rev < maxRevision {
		pools[unsafe.Sizeof(zero)/unsafe.Sizeof(uintptr(0))].Put(val.ref)
	}
	return tmp
}

// API returns the API associated with the given pointer.
func API[API any, T PointerWithFree[API, T, Size], Size PointerSize](ptr T) *API {
	var zero T
	if ptr == zero {
		panic("nil pointer dereference")
	}
	val := access[API, T, Size](ptr)
	if !val.ref.rev.matches(val.rev) {
		crash(&val.ref.freeable, "grow.graphics/gd/internal/mmm error: use after free")
	}
	if val.rev.isRefCounted() {
		return (*API)((*rc)(val.ref.api).api)
	}
	return (*API)(val.ref.api)
}

func getAPI[API any, T PointerWithFree[API, T, Size], Size PointerSize](ptr T) *API {
	val := access[API, T, Size](ptr)
	if !val.ref.rev.matches(val.rev) {
		crash(&val.ref.freeable, "grow.graphics/gd/internal/mmm error: use after free")
	}
	if val.rev.isRefCounted() {
		return (*API)((*rc)(val.ref.api).api)
	}
	return (*API)(val.ref.api)
}

// PointerWithFree is a value with an underlying [Pointer] type and a [Free] method
// that calls [End] on the pointer.
type PointerWithFree[API any, T any, Size PointerSize] interface {
	~struct {
		_ [0]*T
		pointer[API, Size]
	}

	ManagedPointer[Size]

	// Free should release any underlying resources associated with the pointer
	// and call [End] on the pointer. Calling this multiple times is safe but will
	// raise a recoverable runtime panic.
	Free()
}

// Let takes a raw 'Size' pointer value that points at foreign memory or an external resource that is permitted to
// be borrowed until the end of the specified lifetime. The pointer returned by this function may not be moved to
// another lifetime, nor may it be freed.
//
// Pointers passed to this function retain ownership of the original 'ptr' value, however such code may not free
// the pointer until [Lifetime.End] has been called.
func Let[T PointerWithFree[API, T, Size], API any, Size PointerSize](lifetime Lifetime, api *API, ptr Size) T {
	if lifetime.lock {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	var result access[API, T, Size]
	var leased = internalNew[leased[API, T, Size]](lifetime, api, ptr)
	result.ref = leased.ref
	result.rev = leased.rev
	leased.ref.api = unsafe.Pointer(api)
	result.ref.rev.pin()
	result.ref.record("Let")
	return T(result)
}

// New takes a raw 'Size' pointer value that points at memory or a resource that needs to be manually managed.
// The pointer will be assoicated with the specified lifetime, when the lifetime ends the pointer's Free
// method will be called "exactly once". Invalid usage of the pointer remains safe and will raise a runtime
// panic that can be recovered from.
//
// Pointers passed to this function transfer ownership of [T.Free] to T, copies of 'ptr' must behave as if
// T.Free was called immediately on the pointer after being passed to this function.
func New[T PointerWithFree[API, T, Size], API any, Size PointerSize](lifetime Lifetime, api *API, ptr Size) T {
	if lifetime.lock {
		Mutex.Lock()
		defer Mutex.Unlock()
	}
	return internalNew[T, API, Size](lifetime, api, ptr)
}

func internalNew[T PointerWithFree[API, T, Size], API any, Size PointerSize](lifetime Lifetime, api *API, ptr Size) T {
	if lifetime.allocs != nil {
		lifetime.allocs.Add(1)
	}
	block := newObjectWith[Size]()
	block.api = unsafe.Pointer(api)
	block.ptr = ptr
	free := T.Free
	block.end = *(*func(genericPointer))(unsafe.Pointer(&free))

	last := lifetime.root.prv
	next := last.nxt
	last.nxt = &block.freeable
	block.prv = last
	block.nxt = next
	block.api = unsafe.Pointer(api)
	block.record("new")

	return T(access[API, T, Size]{
		pointer: pointer[API, Size]{
			rev: block.rev,
			ref: block,
		},
	})
}

// Pin behaves like [New] except that the pointer is pinned to the lifetime and may not be moved to another one.
// [T.Free] will be called when the lifetime ends.
func Pin[T PointerWithFree[API, T, Size], API any, Size PointerSize](lifetime Lifetime, api *API, ptr Size) T {
	if lifetime.lock {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	var pinned = internalNew[T, API, Size](lifetime, api, ptr)
	var result = access[API, T, Size](pinned)
	result.ref.rev.pin()
	return T(result)
}
