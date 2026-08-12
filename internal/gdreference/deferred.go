package gdreference

import (
	"sync"
	"sync/atomic"

	"graphics.gd/internal/gdextension"
)

// Deferred references occupy otherwise-unused capacity in [Object]: bit 63 of
// the revision field (which the frame counter will never reach) marks the
// state, the remaining revision bits index a global table of resolution
// functions, assigned.objectID carries an optional context object ID and the
// sentinel points at a cache slot shared by all copies of the wrapper.
//
// This enables references that name an object indirectly — a resource by its
// path, a node by its path relative to a scene root — and only resolve it on
// first use. [Resolver] functions are registered once per resolution rule
// (per resource path, or per struct field), while [DeferObject] wrappers are
// cheap and carry the per-instance context.
const deferredBit uint64 = 1 << 63

var deferredFuncs atomic.Pointer[[]func(gdextension.ObjectID) gdextension.Object]
var deferredMutex sync.Mutex

// Resolver indexes the deferred-resolution function table.
type Resolver uint64

// NewResolver registers fn in the deferred-resolution function table. fn
// receives the context object ID the wrapper was created with (zero if none)
// and returns the resolved object, or zero if it cannot be resolved. Safe to
// call from any goroutine, at any time, including before startup.
func NewResolver(fn func(context gdextension.ObjectID) gdextension.Object) Resolver {
	deferredMutex.Lock()
	defer deferredMutex.Unlock()
	var table []func(gdextension.ObjectID) gdextension.Object
	if existing := deferredFuncs.Load(); existing != nil {
		table = append(table, *existing...)
	}
	table = append(table, fn)
	deferredFuncs.Store(&table)
	return Resolver(len(table) - 1)
}

// DeferObject returns an [Object] that resolves through the given [Resolver]
// on first use. The result of the first successful resolution is cached by
// object ID and revalidated through the object database on each subsequent
// access, so the reference reads as null once the resolved object is freed.
// The engine (or the resolver itself) is expected to own the resolved object.
func DeferObject(context gdextension.ObjectID, resolver Resolver) Object {
	return Object{
		assigned: object{objectID: context},
		sentinel: new(object),
		revision: deferredBit | uint64(resolver),
	}
}

func resolveDeferred(obj Object) gdextension.Object {
	if id := obj.sentinel.objectID; id != 0 {
		return gdextension.Host.Objects.Lookup(id)
	}
	table := deferredFuncs.Load()
	index := obj.revision &^ deferredBit
	if table == nil || index >= uint64(len(*table)) {
		return 0
	}
	raw := (*table)[index](obj.assigned.objectID)
	if raw != 0 {
		var id gdextension.ObjectID
		gdextension.Host.Objects.ID.Get(raw, gdextension.CallReturns[gdextension.ObjectID](&id))
		obj.sentinel.objectID = id
	}
	return raw
}
