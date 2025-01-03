package TextServerAdvanced

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/TextServerExtension"
import "graphics.gd/objects/TextServer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
An implementation of [TextServer] that uses HarfBuzz, ICU and SIL Graphite to support BiDi, complex text layouts and contextual OpenType features. This is Godot's default primary [TextServer] interface.
*/
type Instance [1]classdb.TextServerAdvanced
type Any interface {
	gd.IsClass
	AsTextServerAdvanced() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TextServerAdvanced

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextServerAdvanced"))
	return Instance{*(*classdb.TextServerAdvanced)(unsafe.Pointer(&object))}
}

func (self class) AsTextServerAdvanced() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextServerAdvanced() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTextServerExtension() TextServerExtension.Advanced {
	return *((*TextServerExtension.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextServerExtension() TextServerExtension.Instance {
	return *((*TextServerExtension.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTextServer() TextServer.Advanced {
	return *((*TextServer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextServer() TextServer.Instance {
	return *((*TextServer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}
func init() {
	classdb.Register("TextServerAdvanced", func(ptr gd.Object) any {
		return [1]classdb.TextServerAdvanced{*(*classdb.TextServerAdvanced)(unsafe.Pointer(&ptr))}
	})
}
