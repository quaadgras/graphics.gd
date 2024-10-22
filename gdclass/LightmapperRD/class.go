package LightmapperRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Lightmapper"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
LightmapperRD ("RD" stands for [RenderingDevice]) is the built-in GPU-based lightmapper for use with [LightmapGI]. On most dedicated GPUs, it can bake lightmaps much faster than most CPU-based lightmappers. LightmapperRD uses compute shaders to bake lightmaps, so it does not require CUDA or OpenCL libraries to be installed to be usable.
[b]Note:[/b] Only usable when using the Vulkan backend (Forward+ or Mobile), not OpenGL.

*/
type Go [1]classdb.LightmapperRD
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.LightmapperRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("LightmapperRD"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsLightmapperRD() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLightmapperRD() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsLightmapper() Lightmapper.GD { return *((*Lightmapper.GD)(unsafe.Pointer(&self))) }
func (self Go) AsLightmapper() Lightmapper.Go { return *((*Lightmapper.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsLightmapper(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsLightmapper(), name)
	}
}
func init() {classdb.Register("LightmapperRD", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
