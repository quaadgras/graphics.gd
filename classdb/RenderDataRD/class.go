// Package RenderDataRD provides methods for working with RenderDataRD object instances.
package RenderDataRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/RenderData"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This object manages all render data for the rendering device based renderers.
[b]Note:[/b] This is an internal rendering server object only exposed for GDExtension plugins.
*/
type Instance [1]gdclass.RenderDataRD

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRenderDataRD() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderDataRD

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderDataRD"))
	casted := Instance{*(*gdclass.RenderDataRD)(unsafe.Pointer(&object))}
	return casted
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
		return gd.VirtualByName(RenderData.Advanced(self.AsRenderData()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RenderData.Instance(self.AsRenderData()), name)
	}
}
func init() {
	gdclass.Register("RenderDataRD", func(ptr gd.Object) any {
		return [1]gdclass.RenderDataRD{*(*gdclass.RenderDataRD)(unsafe.Pointer(&ptr))}
	})
}
