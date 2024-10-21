package LightmapperRD

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Lightmapper"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
LightmapperRD ("RD" stands for [RenderingDevice]) is the built-in GPU-based lightmapper for use with [LightmapGI]. On most dedicated GPUs, it can bake lightmaps much faster than most CPU-based lightmappers. LightmapperRD uses compute shaders to bake lightmaps, so it does not require CUDA or OpenCL libraries to be installed to be usable.
[b]Note:[/b] Only usable when using the Vulkan backend (Forward+ or Mobile), not OpenGL.

*/
type Simple [1]classdb.LightmapperRD
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.LightmapperRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsLightmapperRD() Expert { return self[0].AsLightmapperRD() }


//go:nosplit
func (self Simple) AsLightmapperRD() Simple { return self[0].AsLightmapperRD() }


//go:nosplit
func (self class) AsLightmapper() Lightmapper.Expert { return self[0].AsLightmapper() }


//go:nosplit
func (self Simple) AsLightmapper() Lightmapper.Simple { return self[0].AsLightmapper() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("LightmapperRD", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
