package gdreference

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/ring"
)

// An object created off the main thread is freed by a cleanup once Go no
// longer refers to it (see [OwnObject]), which only happens when the garbage
// collector gets around to it, so whatever is still waiting on one when the
// engine shuts down would be reported as leaked. They are kept track of here
// for [FreeThreadObjects] to free.
var threads struct {
	sync.Mutex
	owned map[uint64]*threadObject // by the id of the cleanup that frees it
}

type threadObject struct {
	id   uint64
	obj  gdextension.Object
	free func(gdextension.Object)
}

func init() {
	// the id of a cleanup is how its object is found again.
	if field := reflect.TypeFor[runtime.Cleanup]().Field(0); field.Name != "id" || (field.Type.Kind() != reflect.Uint64 && !cleanupIsID) || field.Offset != 0 {
		panic("gdreference: runtime.Cleanup no longer starts with its id")
	}
}

func cleanupID(cleanup *runtime.Cleanup) uint64 { return *(*uint64)(unsafe.Pointer(cleanup)) }

// ownThreadObject has the object freed when the sentinel is garbage collected,
// returning the cleanup to store in the sentinel.
func ownThreadObject(sentinel *object, obj gdextension.Object, free func(gdextension.Object)) runtime.Cleanup {
	owned := &threadObject{obj: obj, free: free}
	cleanup := runtime.AddCleanup(sentinel, func(owned *threadObject) {
		// Queue the free behind any still-buffered cross-thread calls
		// that reference the object: the wrapper was necessarily alive
		// when those were recorded, so FIFO order runs the free after
		// every queued use. A deferred callable would instead run at
		// the next Callable.Cycle, which can precede the frame drain.
		ring.Threads.Defer(owned.release)
	}, owned)
	owned.id = cleanupID(&cleanup)
	threads.Lock()
	if threads.owned == nil {
		threads.owned = make(map[uint64]*threadObject)
	}
	threads.owned[owned.id] = owned
	threads.Unlock()
	return cleanup
}

// release frees the object, unless the engine was given it in the meantime
// or has shut down and it was freed then.
func (owned *threadObject) release() {
	threads.Lock()
	_, ok := threads.owned[owned.id]
	delete(threads.owned, owned.id)
	threads.Unlock()
	if ok {
		owned.free(owned.obj)
	}
}

// disownThreadObject stops the cleanup in the sentinel from freeing its
// object, as the engine owns it now.
func disownThreadObject(cleanup *runtime.Cleanup) {
	threads.Lock()
	delete(threads.owned, cleanupID(cleanup))
	threads.Unlock()
	cleanup.Stop()
}

// FreeThreadObjects frees every object created off the main thread that Go
// still owns, for when the engine shuts down. Main thread only; any reference
// to them that Go still holds is invalid afterwards.
func FreeThreadObjects() {
	threads.Lock()
	owned := threads.owned
	threads.owned = nil
	threads.Unlock()
	for _, owned := range owned {
		owned.free(owned.obj)
	}
}
