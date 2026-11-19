//go:build !generate

package gd_test

import (
	"runtime"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/ClassDB"
	"graphics.gd/classdb/Node"
	"graphics.gd/internal/gdreference"
	"graphics.gd/variant/Object"
)

// SelfInstantiating is registered with a custom constructor that calls a
// method on the object before returning it — which instantiates the engine
// side of the value inside the constructor. The engine's create callback
// used to register such values a second time, orphaning the first
// registration while its dispatch word was still pinned: the orphan's
// collection then aborted the whole program with "runtime.Pinner: found
// leaking pinned pointer" (seen at editor shutdown in projects registering
// constructors like this, e.g. graphics.eg/3d/third-person-controller).
type SelfInstantiating struct {
	Node.Extension[SelfInstantiating]
}

func NewSelfInstantiating() *SelfInstantiating {
	self := new(SelfInstantiating)
	self.AsNode().SetProcessMode(Node.ProcessModeAlways) // instantiates the engine side.
	return self
}

func init() {
	classdb.Register[SelfInstantiating](NewSelfInstantiating)
}

func TestCustomConstructorSelfInstantiates(t *testing.T) {
	raw, ok := rawObject(ClassDB.Instantiate("SelfInstantiating"))
	if !ok {
		t.Fatal("expected ClassDB.Instantiate to return an object")
	}
	obj := Object.Instance([1]gdreference.Object{raw})
	self, ok := Object.As[*SelfInstantiating](obj)
	if !ok {
		t.Fatal("expected the instance to resolve to its Go struct")
	}
	// The engine-side object must be the one the constructor instantiated,
	// not a second one registered over the top of it.
	if gdreference.GetObject(self.AsObject()[0]) != gdreference.GetObject(raw) {
		t.Fatal("expected the constructor's object to be reused, got a second instance")
	}
	if mode := self.AsNode().ProcessMode(); mode != Node.ProcessModeAlways {
		t.Fatalf("expected constructor's state to survive, process mode is %v", mode)
	}
	// The orphaned registration used to only abort once its pinner was
	// collected: force the collection that would trigger it.
	runtime.GC()
	runtime.GC()
	self.AsNode().QueueFree()
}

// rawObject is the engine object an any from the engine holds: the base
// Object.Instance (for a class without a Go package, like Go-defined ones).
func rawObject(v any) (gdreference.Object, bool) {
	switch o := v.(type) {
	case gdreference.Object:
		return o, true
	case interface{ AsObject() [1]gdreference.Object }:
		return o.AsObject()[0], true
	}
	return gdreference.Object{}, false
}

// TestObjectAsBaseWrapper: an engine object whose class has no Go package
// (here a Go-defined class) comes back from the engine as an Object.Any, so
// callers can use it, not as an internal reference type.
func TestObjectAsBaseWrapper(t *testing.T) {
	v := ClassDB.Instantiate("SelfInstantiating")
	if _, ok := v.(Object.Any); !ok {
		t.Fatalf("ClassDB.Instantiate returned %T, want an Object.Any", v)
	}
}
