package lifetime

import (
	"runtime"
	"sync"
)

type Group[API any] struct {
	API *API
	rev uint32 // revision, incremented when recycled.
	buf []uintptr
}

var pool sync.Pool

func New[API any]() *Group[API] {
	lifetime, ok := pool.Get().(*Group[API])
	if !ok {
		lifetime = new(Group[API])
		runtime.SetFinalizer(lifetime, lifetime.Free)
	}
	return lifetime
}

func (g Group[API]) Free() {
	pool.Put(g)
}

type PointerSize interface {
	uintptr | [2]uintptr | [3]uintptr
}

type Bound[API any, T any, S PointerSize] struct {
	rev uint32
	idx uint32
	pin *Group[API]
}
