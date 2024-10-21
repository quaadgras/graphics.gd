package FramebufferCacheRD

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Framebuffer cache manager for Rendering Device based renderers. Provides a way to create a framebuffer and reuse it in subsequent calls for as long as the used textures exists. Framebuffers will automatically be cleaned up when dependent objects are freed.

*/
type Simple [1]classdb.FramebufferCacheRD
func (self Simple) GetCacheMultipass(textures gd.ArrayOf[gd.RID], passes gd.ArrayOf[[1]classdb.RDFramebufferPass], views int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetCacheMultipass(gc, textures, passes, gd.Int(views)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.FramebufferCacheRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates, or obtains a cached, framebuffer. [param textures] lists textures accessed. [param passes] defines the subpasses and texture allocation, if left empty a single pass is created and textures are allocated depending on their usage flags. [param views] defines the number of views used when rendering.
*/
//go:nosplit
func (self class) GetCacheMultipass(ctx gd.Lifetime, textures gd.ArrayOf[gd.RID], passes gd.ArrayOf[object.RDFramebufferPass], views gd.Int) gd.RID {
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

//go:nosplit
func (self class) AsFramebufferCacheRD() Expert { return self[0].AsFramebufferCacheRD() }


//go:nosplit
func (self Simple) AsFramebufferCacheRD() Simple { return self[0].AsFramebufferCacheRD() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("FramebufferCacheRD", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
