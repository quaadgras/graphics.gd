package UniformSetCacheRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
Uniform set cache manager for Rendering Device based renderers. Provides a way to create a uniform set and reuse it in subsequent calls for as long as the uniform set exists. Uniform set will automatically be cleaned up when dependent objects are freed.

*/
type Simple [1]classdb.UniformSetCacheRD
func (self Simple) GetCache(shader gd.RID, set int, uniforms gd.ArrayOf[[1]classdb.RDUniform]) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetCache(gc, shader, gd.Int(set), uniforms))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.UniformSetCacheRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates/returns a cached uniform set based on the provided uniforms for a given shader.
*/
//go:nosplit
func (self class) GetCache(ctx gd.Lifetime, shader gd.RID, set gd.Int, uniforms gd.ArrayOf[object.RDUniform]) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, set)
	callframe.Arg(frame, mmm.Get(uniforms))
	var r_ret = callframe.Ret[gd.RID](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.UniformSetCacheRD.Bind_get_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsUniformSetCacheRD() Expert { return self[0].AsUniformSetCacheRD() }


//go:nosplit
func (self Simple) AsUniformSetCacheRD() Simple { return self[0].AsUniformSetCacheRD() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("UniformSetCacheRD", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
