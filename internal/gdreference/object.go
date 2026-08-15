package gdreference

import (
	"runtime"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadcheck"
)

var now uint64 = 2

// Barrier needs to be called whenever Let references are
// invalidated.
func Barrier() {
	if threadcheck.Main() {
		now++
	}
}

func init() {
	if unsafe.Sizeof(runtime.Cleanup{}) != unsafe.Sizeof(object{}) {
		panic("gdreference: size of runtime.Cleanup does not match size of object")
	}
}

// Object reference that's safe to use from a single goroutine.
type Object struct {
	_ [0]*Object

	assigned object
	sentinel *object
	revision uint64
}

// RawObject returns an unsafe [Object] reference from a raw
// [gdextension.Object] pointer, no memory safety protections
// will apply to the result.
func RawObject(obj gdextension.Object) Object {
	if obj == 0 {
		return Object{}
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(obj, gdextension.CallReturns[gdextension.ObjectID](&id))
	return Object{assigned: object{inEngine: obj, objectID: id}}
}

// LetObject creates an engine-owned [Object] reference.
func LetObject(obj gdextension.Object) Object {
	if obj == 0 {
		return Object{}
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(obj, gdextension.CallReturns[gdextension.ObjectID](&id))
	var revision uint64
	if threadcheck.Main() {
		revision = now
	}
	return Object{
		assigned: object{objectID: id, inEngine: obj},
		sentinel: &borrowSentinel,
		revision: revision,
	}
}

// PinObject writes a [gdextension.Object] into an existing [Object] pointer on the heap.
// Useful for extension classes, object is not automatically freed.
func PinObject(obj *Object, raw gdextension.Object) {
	if raw == 0 {
		*obj = Object{}
		return
	}
	if obj.assigned.inEngine == raw {
		obj.revision = 0
		return
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(raw, gdextension.CallReturns[gdextension.ObjectID](&id))
	obj.sentinel = &obj.assigned
	obj.assigned = object{objectID: id, inEngine: raw}
}

// OwnObject creates a Go-owned [Object] reference.
//
// The choice between the two lifetime models is [threadcheck.FrameTemporaries],
// not [threadcheck.Main]: Main governs dispatch (on wasm it is unconditionally
// true, because a single-threaded runtime must never queue to a ring nobody
// else drains), while this is a question about lifetimes, and pooled objects
// are swept by the per-frame [GC]. Frame-alignment is what makes that sweep
// safe, and on wasm only goroutine identity can establish it — see
// threadcheck_wasm.go. Without this, a goroutine on web that held an object
// across a frame had it freed underneath it; this is the object-side match of
// the same gate on gd.Wrap*, which anchors reference types.
func OwnObject(obj gdextension.Object, free func(gdextension.Object)) Object {
	if obj == 0 {
		return Object{}
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(obj, gdextension.CallReturns[gdextension.ObjectID](&id))
	var sentinel *object
	var revision uint64
	var result Object
	if threadcheck.FrameTemporaries() {
		if len(pool_free) > 0 {
			sentinel = pool_free[len(pool_free)-1]
			pool_free = pool_free[: len(pool_free)-1 : cap(pool_free)]
		} else {
			var bucket, i = tail / 128, tail % 128
			if bucket >= len(pool) {
				pool = append(pool, new([128]object))
			}
			sentinel = &pool[bucket][i]
			tail++
		}
		sentinel.inEngine = obj
		sentinel.objectID = id
		revision = now
		result.assigned.objectID = id
	} else {
		sentinel = new(object)
		cleanup := runtime.AddCleanup(sentinel, func(obj gdextension.Object) {
			// Queue the free behind any still-buffered cross-thread calls
			// that reference the object: the wrapper was necessarily alive
			// when those were recorded, so FIFO order runs the free after
			// every queued use. A deferred callable would instead run at
			// the next Callable.Cycle, which can precede the frame drain.
			ring.Threads.Defer(func() {
				free(obj)
			})
		}, obj)
		*sentinel = *(*object)(unsafe.Pointer(&cleanup))
	}
	result.assigned.inEngine = obj
	result.sentinel = sentinel
	result.revision = revision
	return result
}

// NewObject returns a new static [Object] with a pointer value not known
// in advance, can be set with [SetObject].
func NewObject() Object {
	return Object{
		sentinel: new(object),
	}
}

// GetObject returns the underlying engine pointer for an [Object].
func GetObject(obj Object) gdextension.Object {
	if obj.sentinel == nil {
		return obj.assigned.inEngine
	}
	// Fast path for pinned/pooled objects (e.g. an extension's own self, the
	// receiver of every outbound call a virtual makes): this replicates the
	// TypePinned/TypePooled branch of [AskObject] inline so the hot path
	// avoids the call. It runs BEFORE the revision check because it is pure
	// loads — the revision path costs a threadcheck.Main (an assembly call)
	// that pinned self-references, whose revision is 0, would pay for
	// nothing. Both guards return assigned.inEngine, so the order does not
	// change any result. The guards match AskObject exactly — a non-zero
	// assigned id whose sentinel refers to the same object — so
	// borrow/thread/static references still fall through. The revision tag
	// check excludes deferred references, whose assigned.objectID is a
	// context (not the target) and could otherwise match a cache slot that
	// resolved to the context object itself.
	if obj.revision&deferredBit == 0 && obj.sentinel != &borrowSentinel && obj.assigned.objectID != 0 && obj.sentinel.objectID == obj.assigned.objectID {
		return obj.assigned.inEngine
	}
	if threadcheck.Main() && obj.revision == now {
		return obj.assigned.inEngine
	}
	raw, _ := AskObject(obj)
	return raw
}

// Anchor returns the garbage-collected allocation that holds obj's engine-side
// lifetime open, for passing to [runtime.KeepAlive]. For references created off
// the main thread that is the sentinel carrying the [runtime.AddCleanup] which
// queues the free, so keeping it alive keeps the cleanup from running.
//
// Every caller that takes a raw pointer out of an [Object] with [GetObject] and
// records it somewhere the collector cannot see — the cross-thread ring, an
// argument pack — must keep the anchor alive until the record has been made.
// The wrapper is dead the moment its pointer has been extracted, so without
// this the collector is free to run the cleanup, and the queued free then
// overtakes the very call that was about to be enqueued behind it.
func (obj Object) Anchor() unsafe.Pointer { return unsafe.Pointer(obj.sentinel) }

// SetObject sets the underlying engine pointer for a [TypeStatic]
// [Object] created with [NewObject].
func SetObject(obj Object, val gdextension.Object) {
	if obj.assigned != (object{}) {
		panic("SetObject can only be used with objects created by NewObject")
	}
	if val == 0 {
		*obj.sentinel = object{}
		return
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(val, gdextension.CallReturns[gdextension.ObjectID](&id))
	obj.sentinel.inEngine = val
	obj.sentinel.objectID = id
}

var borrowSentinel object

// AskObject returns lifetime information for the object.
func AskObject(obj Object) (gdextension.Object, Type) {
	if obj.revision&deferredBit != 0 {
		return resolveDeferred(obj), TypeBorrow
	}
	switch obj.sentinel {
	case nil:
		return obj.assigned.inEngine, TypeUnsafe
	case &borrowSentinel:
		// The epoch fast path is main-thread only: `now` is written by the
		// frame Barrier on the main thread without synchronization, and a
		// borrow's cached pointer is only guaranteed until that Barrier — an
		// off-main reader could validate against a stale epoch mid-frame and
		// call through a pointer the frame GC just invalidated (silent
		// wrong-object calls). Off the main thread, and for main borrows
		// whose epoch has passed, resolve by object id instead.
		if threadcheck.Main() && obj.revision == now {
			return obj.assigned.inEngine, TypeBorrow
		}
		return gdextension.Host.Objects.Lookup(obj.assigned.objectID), TypeBorrow
	}
	if obj.assigned.objectID == 0 {
		if obj.assigned.inEngine == 0 {
			if obj.sentinel.inEngine == 0 {
				return gdextension.Host.Objects.Lookup(obj.sentinel.objectID), TypeStatic
			}
			return obj.sentinel.inEngine, TypeStatic
		}
		if *obj.sentinel == obj.assigned {
			return 0, TypeThread
		}
		if obj.sentinel.inEngine == 0 && obj.sentinel.objectID != 0 {
			// Ownership was transferred to the engine (see EndObject):
			// borrow the object back through the object database, which
			// returns 0 if the engine has since freed it.
			return gdextension.Host.Objects.Lookup(obj.sentinel.objectID), TypeBorrow
		}
		return obj.assigned.inEngine, TypeThread
	}
	if obj.sentinel.objectID == obj.assigned.objectID {
		if obj.revision <= 1 {
			return obj.assigned.inEngine, TypePinned
		}
		return obj.assigned.inEngine, TypePooled
	}
	return gdextension.Host.Objects.Lookup(obj.assigned.objectID), TypeBorrow
}

// EndObject leaks the object, releasing ownership to the engine.
func EndObject(obj Object) (gdextension.Object, bool) {
	raw, t := AskObject(obj)
	switch t {
	case TypePooled:
		*obj.sentinel = object{}
		pool_free = append(pool_free, obj.sentinel)
	case TypeThread:
		cleanup := (*runtime.Cleanup)(unsafe.Pointer(obj.sentinel))
		cleanup.Stop()
		// The engine owns the object now: leave the wrapper usable as a
		// looked-up borrow (parity with the pooled main-thread case) so a
		// goroutine can keep calling methods on a node it has handed over,
		// e.g. after adding it to the scene tree. The zero inEngine field
		// distinguishes this state from live cleanup bits, whose first word
		// is a non-zero cleanup id.
		var id gdextension.ObjectID
		gdextension.Host.Objects.ID.Get(raw, gdextension.CallReturns[gdextension.ObjectID](&id))
		*obj.sentinel = object{inEngine: 0, objectID: id}
	case TypeUnsafe, TypePinned:
	case TypeStatic:
		obj.sentinel.inEngine = 0
	case TypeBorrow:
		return raw, false
	}
	return raw, true
}

// CutObject either ends the object (true) or gets it (false)
func CutObject(obj Object, end bool) gdextension.Object {
	if end {
		raw, _ := EndObject(obj)
		return raw
	}
	return GetObject(obj)
}

// UseObject marks the object as used, preventing it from being
// freed for one frame.
func UseObject(obj *Object) {
	if obj.revision&deferredBit != 0 {
		return // deferred references do not participate in frame pooling.
	}
	if obj.sentinel == &obj.assigned {
		obj.revision = 0
		return
	}
	if !BadObject(*obj) && obj.sentinel != nil && obj.sentinel != &borrowSentinel && obj.assigned.objectID != 0 && obj.sentinel.objectID == obj.assigned.objectID {
		obj.sentinel.inEngine = obj.assigned.inEngine
	}
}

// BadObject returns true if the reference has been invalidated.
func BadObject(obj Object) bool {
	return obj == Object{} || obj == Object{revision: 1} || GetObject(obj) == 0
}

type object struct {
	objectID gdextension.ObjectID
	inEngine gdextension.Object
}

var tail int
var pool = []*[128]object{{}}
var pool_free []*object
