package RenderDataRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/RenderData"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This object manages all render data for the rendering device based renderers.
[b]Note:[/b] This is an internal rendering server object only exposed for GDExtension plugins.
*/
type Instance [1]classdb.RenderDataRD
type Any interface {
	gd.IsClass
	AsRenderDataRD() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderDataRD

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderDataRD"))
	return Instance{classdb.RenderDataRD(object)}
}

func (self class) AsRenderDataRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderDataRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRenderData() RenderData.Advanced {
	return *((*RenderData.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderData() RenderData.Instance {
	return *((*RenderData.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderData(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderData(), name)
	}
}
func init() {
	classdb.Register("RenderDataRD", func(ptr gd.Object) any { return [1]classdb.RenderDataRD{classdb.RenderDataRD(ptr)} })
}
