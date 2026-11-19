package gdclass

import (
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdreference"
)

type object gdreference.Object

func (obj object) AsObject() [1]gdreference.Object {
	return [1]gdreference.Object{gdreference.Object(obj)}
}

// Anchor returns the value to pass to [runtime.KeepAlive] in order to hold the
// object's engine-side lifetime open across an engine call, see
// [gdreference.Anchor]. It is a single word, so that the generated bindings can
// keep their receiver reachable across the enqueue without spilling the whole
// wrapper.
func (obj object) Anchor() unsafe.Pointer {
	return gdreference.Object(obj).Anchor()
}

type Any interface {
	AsObject() [1]gdreference.Object
}

type Object = gdreference.Object
type RefCounted = gd.RefCounted

var classDB = make(map[string]func(gdreference.Object) any)

func Register(name string, constructor func(gdreference.Object) any) {
	classDB[name] = constructor
}

func init() {
	gd.ObjectAs = func(name string, ptr gdreference.Object) any {
		if constructor, ok := classDB[name]; ok {
			return constructor(ptr)
		}
		// A class whose package isn't imported (or a Go-defined class):
		// the base Object wrapper, so that it still satisfies Object.Any.
		if constructor, ok := classDB["Object"]; ok {
			return constructor(ptr)
		}
		return ptr
	}
}
