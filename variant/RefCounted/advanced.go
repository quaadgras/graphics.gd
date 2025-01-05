package RefCounted

import (
	"reflect"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
)

type Advanced [1]gdclass.RefCounted

func (obj Advanced) AsObject() [1]gd.Object          { return obj[0].AsObject() }
func (self *Advanced) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

// Virtual method lookup.
func (obj Advanced) Virtual(name string) reflect.Value { return obj[0].Virtual(name) }
