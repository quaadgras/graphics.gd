package RefCounted

import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdreference"
)

type Advanced [1]gdclass.RefCounted

func (obj Advanced) AsObject() [1]gdreference.Object { return obj[0].AsObject() }
func (self *Advanced) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

// Virtual method lookup.
func (obj Advanced) Virtual(name string) reflect.Value { return obj[0].Virtual(name) }
