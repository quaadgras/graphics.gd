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
	raw, ok := ClassDB.Instantiate("SelfInstantiating").(gdreference.Object)
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
