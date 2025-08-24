package RefCounted

import (
	"reflect"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
)

type Instance [1]gd.RefCounted

func (obj Instance) AsObject() [1]gd.Object          { return obj[0].AsObject() }
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

// Virtual method lookup.
func (obj Instance) Virtual(name string) reflect.Value { return obj[0].Virtual(name) }

func (obj *Instance) SetObject(object [1]gd.Object) bool {
	return obj[0].SetObject(object)
}

// Extension can be embedded in a struct to create a new class. T should be the type of the struct
// that embeds this Extension.
type Extension[T gdclass.Interface] struct {
	gdclass.Extension[T, Instance]
}

func (class *Extension[T]) AsObject() [1]gdclass.Object    { return class.Super().AsObject() }
func (class *Extension[T]) AsRefCounted() [1]gd.RefCounted { return class.Super() }
