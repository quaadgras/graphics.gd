package Callable

import (
	"reflect"
	"testing"
)

type fakeRef struct{ id int }

// fakeInstance stands in for Object.Instance: a one-element array of the reference.
type fakeInstance [1]fakeRef

// fakeClass stands in for a class instance that takes the reference through SetObject.
type fakeClass struct{ ref fakeRef }

func (c *fakeClass) SetObject(obj [1]fakeRef) bool { c.ref = obj[0]; return true }

func TestWrapObject(t *testing.T) {
	value := reflect.ValueOf(fakeRef{id: 7})

	wrapped, ok := wrapObject(value, reflect.TypeFor[fakeInstance]())
	if !ok || wrapped.Interface().(fakeInstance)[0].id != 7 {
		t.Fatalf("one-element array parameter not wrapped: ok=%v %v", ok, wrapped)
	}

	wrapped, ok = wrapObject(value, reflect.TypeFor[fakeClass]())
	if !ok || wrapped.Interface().(fakeClass).ref.id != 7 {
		t.Fatalf("SetObject parameter not wrapped: ok=%v %v", ok, wrapped)
	}

	if _, ok := wrapObject(value, reflect.TypeFor[int]()); ok {
		t.Fatal("unrelated parameter type reported as wrapped")
	}
	if _, ok := wrapObject(value, reflect.TypeFor[[2]fakeRef]()); ok {
		t.Fatal("two-element array reported as wrapped")
	}
}
