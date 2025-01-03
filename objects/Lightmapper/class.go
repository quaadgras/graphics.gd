package Lightmapper

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class should be extended by custom lightmapper classes. Lightmappers can then be used with [LightmapGI] to provide fast baked global illumination in 3D.
Godot contains a built-in GPU-based lightmapper [LightmapperRD] that uses compute shaders, but custom lightmappers can be implemented by C++ modules.
*/
type Instance [1]classdb.Lightmapper
type Any interface {
	gd.IsClass
	AsLightmapper() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Lightmapper

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Lightmapper"))
	return Instance{*(*classdb.Lightmapper)(unsafe.Pointer(&object))}
}

func (self class) AsLightmapper() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLightmapper() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("Lightmapper", func(ptr gd.Object) any { return [1]classdb.Lightmapper{*(*classdb.Lightmapper)(unsafe.Pointer(&ptr))} })
}
