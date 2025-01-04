// Package LightmapperRD provides methods for working with LightmapperRD object instances.
package LightmapperRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Lightmapper"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
LightmapperRD ("RD" stands for [RenderingDevice]) is the built-in GPU-based lightmapper for use with [LightmapGI]. On most dedicated GPUs, it can bake lightmaps much faster than most CPU-based lightmappers. LightmapperRD uses compute shaders to bake lightmaps, so it does not require CUDA or OpenCL libraries to be installed to be usable.
[b]Note:[/b] Only usable when using the Vulkan backend (Forward+ or Mobile), not OpenGL.
*/
type Instance [1]gdclass.LightmapperRD
type Any interface {
	gd.IsClass
	AsLightmapperRD() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.LightmapperRD

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("LightmapperRD"))
	return Instance{*(*gdclass.LightmapperRD)(unsafe.Pointer(&object))}
}

func (self class) AsLightmapperRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLightmapperRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsLightmapper() Lightmapper.Advanced {
	return *((*Lightmapper.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsLightmapper() Lightmapper.Instance {
	return *((*Lightmapper.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsLightmapper(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsLightmapper(), name)
	}
}
func init() {
	gdclass.Register("LightmapperRD", func(ptr gd.Object) any {
		return [1]gdclass.LightmapperRD{*(*gdclass.LightmapperRD)(unsafe.Pointer(&ptr))}
	})
}
