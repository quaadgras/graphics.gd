package gdreference

import (
	"runtime"
	"sync/atomic"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/threadsafe"
)

type Object struct {
	_ [0]*Object

	assigned object
	sentinel *object
}

func New(obj gdextension.Object, id gdextension.ObjectID, on_main_thread bool) Object {
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
	atomic.StoreUint64(&sentinel.objectID, uint64(id))
	atomic.StoreUint64(&sentinel.inEngine, uint64(obj))
	atomic.StoreUint64(&sentinel.checksum, uint64(id)^uint64(obj))
	return Object{
		sentinel: sentinel,
		assigned: object{
			objectID: uint64(id),
			inEngine: uint64(obj),
			checksum: uint64(id) ^ uint64(obj),
		},
	}
}

func (obj Object) lookup() (gdextension.Object, bool) {
	if obj.sentinel == nil {
		return gdextension.Object(obj.assigned.inEngine), true
	}
	objectID := atomic.LoadUint64(&obj.sentinel.objectID)
	inEngine := atomic.LoadUint64(&obj.sentinel.inEngine)
	checksum := atomic.LoadUint64(&obj.sentinel.checksum)
	if checksum != objectID^uint64(inEngine) || objectID != obj.assigned.objectID {
		raw := gdextension.Host.Objects.Lookup(gdextension.ObjectID(obj.assigned.objectID))
		return raw, raw != 0
	}
	return gdextension.Object(inEngine), true
}

func (obj Object) end() (gdextension.Object, bool) {
	if obj.sentinel == nil {
		return gdextension.Object(obj.assigned.inEngine), true
	}
	objectID := atomic.LoadUint64(&obj.sentinel.objectID)
	inEngine := atomic.LoadUint64(&obj.sentinel.inEngine)
	checksum := atomic.LoadUint64(&obj.sentinel.checksum)
	if checksum != objectID^uint64(inEngine) || objectID != obj.assigned.objectID {
		return 0, false
	}
	if atomic.CompareAndSwapUint64(&obj.sentinel.inEngine, inEngine, 0) {
		atomic.StoreUint64(&obj.sentinel.checksum, 0)
		atomic.StoreUint64(&obj.sentinel.objectID, 0)
		if cleanup, ok := cleanups.Lookup(gdextension.ObjectID(objectID)); ok {
			cleanup.Stop()
			cleanups.Remove(gdextension.ObjectID(objectID))
			select {
			case free <- obj.sentinel:
			default:
			}
		}
		return gdextension.Object(inEngine), true
	}
	return 0, false
}

func (obj Object) Free() {
	if raw, ok := obj.end(); ok {
		gdextension.Host.Objects.Unsafe.Free(raw)
	}
}

type object struct {
	objectID uint64
	inEngine uint64
	checksum uint64
}

var tail int
var pool = [][128]object{{}}
var pool_free []*object
var free = make(chan *object, 128)
var cleanups threadsafe.Map[gdextension.ObjectID, runtime.Cleanup]
