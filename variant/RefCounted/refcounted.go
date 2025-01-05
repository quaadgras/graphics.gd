package RefCounted

import (
	"reflect"
	"unsafe"

	gd "graphics.gd/internal"
)

type Instance [1]gd.RefCounted

func (obj Instance) AsObject() [1]gd.Object          { return obj[0].AsObject() }
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

// Virtual method lookup.
func (obj Instance) Virtual(name string) reflect.Value { return obj[0].Virtual(name) }
