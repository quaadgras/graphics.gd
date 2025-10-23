package gdreference

import (
	"runtime"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/threadsafe"
)

// Object reference that's safe to use from a single goroutine.
type Object struct {
	_ [0]*Object

	// Cases
	//
	// sentinel == nil
	// 	the object is an unsafe/raw reference
	//
	// assigned.objectID == 0 && sentinel.objectID == 0
	// 	the object is owned by Go and there is a cleanup associated with it.
	//
	// assigned.objectID == 0 && sentinel.objectID != 0
	// 	the values in the sentinel object override the values in the assigned object
	//
	// assigned.objectID != 0 && sentinel.objectID == 0
	// 	the object is owned by the mainthread and has been used within the last frame.
	//
	// assigned.objectID != sentinel.objectID
	// 	the object may have been invalidated, ignore the sentinel and only lookup the assigned objectID

	assigned object
	sentinel *object
}

func NewObject(obj gdextension.Object, id gdextension.ObjectID, on_main_thread bool) Object {
	var sentinel *object
	if on_main_thread {
		if pool_free != nil {
			sentinel = pool_free[len(pool_free)-1]
			pool_free = pool_free[: len(pool_free)-1 : cap(pool_free)]
		} else {
			var bucket, i = tail / 128, tail % 128
			if bucket >= len(pool) {
				pool = append(pool, [128]object{})
			}
			sentinel = &pool[bucket][i]
		}
	} else {
		select {
		case sentinel = <-free:
		default:
			sentinel = new(object)
		}
		cleanups.Insert(id, runtime.AddCleanup(sentinel, func(id gdextension.ObjectID) {
			if obj := gdextension.Host.Objects.Lookup(id); obj != 0 {
				gdextension.Host.Objects.Unsafe.Free(obj)
			}
			cleanups.Remove(id)
		}, id))
	}
	sentinel.objectID = id
	sentinel.inEngine = obj
	return Object{
		sentinel: sentinel,
		assigned: object{
			objectID: id,
			inEngine: obj,
		},
	}
}

func RawObject(obj gdextension.Object, id gdextension.ObjectID) Object {
	return Object{assigned: object{objectID: id, inEngine: obj}}
}

func (obj Object) lookup() gdextension.Object {
	if obj.sentinel == nil {
		return obj.assigned.inEngine
	}
	objectID := obj.sentinel.objectID
	inEngine := obj.sentinel.inEngine
	if objectID != obj.assigned.objectID && obj.assigned.objectID != 0 {
		return gdextension.Host.Objects.Lookup(obj.assigned.objectID)
	}
	return inEngine
}

func (obj Object) end() gdextension.Object {
	if obj.sentinel == nil {
		return obj.assigned.inEngine
	}
	objectID := obj.sentinel.objectID
	inEngine := obj.sentinel.inEngine
	if obj.assigned.objectID == 0 {
		obj.sentinel.inEngine = 0
		return inEngine
	}
	if objectID != obj.assigned.objectID {
		return 0
	}
	obj.sentinel.inEngine = 0
	if obj.assigned.objectID == 0 && objectID == 0 {
		if cleanup, ok := cleanups.Lookup(objectID); ok {
			cleanup.Stop()
			cleanups.Remove(objectID)
			select {
			case free <- obj.sentinel:
			default:
			}
		}
	}
	return inEngine
}

func (obj Object) Free() {
	if raw := obj.end(); raw != 0 {
		gdextension.Host.Objects.Unsafe.Free(raw)
	}
}

type object struct {
	objectID gdextension.ObjectID
	inEngine gdextension.Object
	lifetime *lifetime
}

type lifetime struct {
	cleanup runtime.Cleanup
	created uint64
}

var tail int
var pool = [][128]object{{}}
var pool_free []*object
var free = make(chan *object, 128)
var cleanups threadsafe.Map[gdextension.ObjectID, runtime.Cleanup]
