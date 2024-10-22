package FramebufferCacheRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Framebuffer cache manager for Rendering Device based renderers. Provides a way to create a framebuffer and reuse it in subsequent calls for as long as the used textures exists. Framebuffers will automatically be cleaned up when dependent objects are freed.

*/
type Go [1]classdb.FramebufferCacheRD

/*
Creates, or obtains a cached, framebuffer. [param textures] lists textures accessed. [param passes] defines the subpasses and texture allocation, if left empty a single pass is created and textures are allocated depending on their usage flags. [param views] defines the number of views used when rendering.
*/
func (self Go) GetCacheMultipass(textures gd.ArrayOf[gd.RID], passes gd.ArrayOf[gdclass.RDFramebufferPass], views int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetCacheMultipass(gc, textures, passes, gd.Int(views)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FramebufferCacheRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("FramebufferCacheRD"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Creates, or obtains a cached, framebuffer. [param textures] lists textures accessed. [param passes] defines the subpasses and texture allocation, if left empty a single pass is created and textures are allocated depending on their usage flags. [param views] defines the number of views used when rendering.
*/
//go:nosplit
func (self class) GetCacheMultipass(ctx gd.Lifetime, textures gd.ArrayOf[gd.RID], passes gd.ArrayOf[gdclass.RDFramebufferPass], views gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(textures))
	callframe.Arg(frame, mmm.Get(passes))
	callframe.Arg(frame, views)
	var r_ret = callframe.Ret[gd.RID](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FramebufferCacheRD.Bind_get_cache_multipass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFramebufferCacheRD() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFramebufferCacheRD() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("FramebufferCacheRD", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
