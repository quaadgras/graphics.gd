//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"sync/atomic"
)

type Object struct {
	_Object struct{}
	obj     safeObject
}

type RefCounted struct {
	_RefCounted struct{}
	obj         safeObject
}

func NewRefCounted(ctx InstanceOwner, at *RefCounted) *RefCounted {
	if at == nil {
		at = new(RefCounted)
	}
	at.obj.new(ctx, at.class(), true)
	return at
}
func (gdClass RefCounted) Free()          { gdClass.obj.free() }
func (gdClass RefCounted) owner() cObject { return gdClass.obj.get() }
func (gdClass RefCounted) Object() Object { return Object{obj: gdClass.obj} }
func (RefCounted) class() string          { return "RefCounted\000" }

var methodRefCounted [3]cMethodBind

func (gdClass RefCounted) virtual(rtype reflect.Type, name string) (method reflect.Method, ok bool) {
	return
}
func (gdClass RefCounted) InitRef() bool {
	return methodCall[bool](gdClass.obj.get(), methodRefCounted[0])
}
func (gdClass RefCounted) Reference() bool {
	return methodCall[bool](gdClass.obj.get(), methodRefCounted[1])
}
func (gdClass RefCounted) Unreference() bool {
	return methodCall[bool](gdClass.obj.get(), methodRefCounted[2])
}

// safeObject wraps an engine object and provides a memory-safe interface
// that tracks ownership, references and panics on double-free.
type safeObject struct {
	self *safeObject // pointer to the 'owner'
	cPtr cObject
	refC atomic.Int64 // Go reference counter.

	owner InstanceOwner
}

// new creates a new instance of the given class and takes
// ownership of it.
func (obj *safeObject) new(ctx InstanceOwner, class string, ref bool) {
	obj.take(ctx, classDB.construct_object(class))
	if ref {
		gdCounter := RefCounted{obj: *obj}
		if !gdCounter.InitRef() || !gdCounter.Reference() {
			panic("failed to reference object")
		}
		obj.refC.Add(1)
	}
}

// hold a singleton reference that cannot be freed.
func (obj *safeObject) hold(singleton cObject) {
	if obj.self != nil || obj.cPtr != 0 {
		panic("object already initialized")
	}
	obj.self = obj // take 'ownership'.
	obj.cPtr = singleton
	obj.refC.Store(-1) // not freeable.
}

// get returns the engine pointer to the object
// will panic if the object is not initialised.
// result will be zero if the object has been
// previously freed or given back to the engine.
func (obj *safeObject) get() cObject {
	if obj.self == obj {
		return obj.cPtr
	}
	return obj.self.cPtr
}

// take marks that ownership of the given pointer has been taken
// from the engine. It may no longer be freed by the engine.
// This function panics if the object is already initialised.
func (obj *safeObject) take(ctx InstanceOwner, ptr cObject) {
	if obj.self != nil || obj.cPtr != 0 {
		panic("object already initialized")
	}
	obj.self = obj // take ownership.
	obj.cPtr = ptr
	obj.owner = ctx
	ctx.children.Add(1)
}

// give marks that the object has been given back to the engine. It will
// no longer be available for use in Go. This function panics if the object
// is not owned by Go.
func (obj *safeObject) give(ctx InstanceOwner) cObject {
	obj = obj.self // ignore possible copy.
	if obj == nil || obj.self == nil {
		panic("object does not have Go ownership")
	}
	obj.self = nil
	ctx.children.Add(-1)
	return obj.cPtr
}

// destroy the object, called from the engine.
func (obj *safeObject) destroy() {
	obj = obj.self // ignore possible copy.
	if obj == nil || obj.self == nil {
		panic("object does not have Go ownership")
	}
	obj.cPtr = 0
}

// free frees the underlying Godot object, free panics if
// the object was already freed, or if it is not owned by
// Go.
func (obj *safeObject) free() {
	obj = obj.self // ignore possible copy.
	if obj == nil || obj.self == nil {
		panic("object does not have Go ownership")
	}

	ctx := obj.owner

	if obj.cPtr == 0 || ctx == belongsToGodot {
		//panic("object already freed")
		return
	}

	if obj.refC.Load() == 0 {
		obj.cPtr.destroy()
	} else {
		result := obj.refC.Add(-1)
		if result < 0 {
			panic("object already freed")
		}
		if result == 0 {
			gdCounter := RefCounted{obj: *obj}
			if !gdCounter.Unreference() {
				panic("failed to unreference object")
			}
		}
	}
	ctx.children.Add(-1)
	fmt.Println("free", obj)
}
