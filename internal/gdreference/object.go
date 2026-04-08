package gdreference

import (
	"runtime"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/variant/Callable"
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
	id := gdextension.ObjectID(gdunsafe.Object(obj).ID())
	return Object{assigned: object{inEngine: obj, objectID: id}}
}

// LetObject creates an engine-owned [Object] reference.
func LetObject(obj gdextension.Object) Object {
	if obj == 0 {
		return Object{}
	}
	id := gdextension.ObjectID(gdunsafe.Object(obj).ID())
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
	id := gdextension.ObjectID(gdunsafe.Object(raw).ID())
	obj.sentinel = &obj.assigned
	obj.assigned = object{objectID: id, inEngine: raw}
}

// OwnObject creates a Go-owned [Object] reference.
func OwnObject(obj gdextension.Object, free func(gdextension.Object)) Object {
	if obj == 0 {
		return Object{}
	}
	id := gdextension.ObjectID(gdunsafe.Object(obj).ID())
	var sentinel *object
	var revision uint64
	var result Object
	if threadcheck.Main() {
		if len(pool_free) > 0 {
			sentinel = pool_free[len(pool_free)-1]
			pool_free = pool_free[: len(pool_free)-1 : cap(pool_free)]
		} else {
			var bucket, i = tail / 128, tail % 128
			if bucket >= len(pool) {
				pool = append(pool, [128]object{})
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
			Callable.Defer(Callable.New(func() {
				free(obj)
			}))
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
	if obj.sentinel == nil || (threadcheck.Main() && obj.revision == now) {
		return obj.assigned.inEngine
	}
	raw, _ := AskObject(obj)
	return raw
}

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
	id := gdextension.ObjectID(gdunsafe.Object(val).ID())
	obj.sentinel.inEngine = val
	obj.sentinel.objectID = id
}

var borrowSentinel object

// AskObject returns lifetime information for the object.
func AskObject(obj Object) (gdextension.Object, Type) {
	switch obj.sentinel {
	case nil:
		return obj.assigned.inEngine, TypeUnsafe
	case &borrowSentinel:
		if obj.revision == now {
			return obj.assigned.inEngine, TypeBorrow
		}
		return gdextension.Object(gdunsafe.ObjectID(obj.assigned.objectID).Object()), TypeBorrow
	}
	if obj.assigned.objectID == 0 {
		if obj.assigned.inEngine == 0 {
			if obj.sentinel.inEngine == 0 {
				return gdextension.Object(gdunsafe.ObjectID(obj.sentinel.objectID).Object()), TypeStatic
			}
			return obj.sentinel.inEngine, TypeStatic
		}
		if *obj.sentinel == obj.assigned {
			return 0, TypeThread
		}
		return obj.assigned.inEngine, TypeThread
	}
	if obj.sentinel.objectID == obj.assigned.objectID {
		if obj.revision <= 1 {
			return obj.assigned.inEngine, TypePinned
		}
		return obj.assigned.inEngine, TypePooled
	}
	return gdextension.Object(gdunsafe.ObjectID(obj.assigned.objectID).Object()), TypeBorrow
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
		//fmt.Println("Stopped cleanup for thread", raw, cleanup)
		*obj.sentinel = obj.assigned
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
var pool = [][128]object{{}}
var pool_free []*object
