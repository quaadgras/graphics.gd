package UniformSetCacheRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Uniform set cache manager for Rendering Device based renderers. Provides a way to create a uniform set and reuse it in subsequent calls for as long as the uniform set exists. Uniform set will automatically be cleaned up when dependent objects are freed.
*/
type Instance [1]classdb.UniformSetCacheRD
type Any interface {
	gd.IsClass
	AsUniformSetCacheRD() Instance
}

/*
Creates/returns a cached uniform set based on the provided uniforms for a given shader.
*/
func GetCache(shader Resource.ID, set int, uniforms gd.Array) Resource.ID {
	self := Instance{}
	return Resource.ID(class(self).GetCache(shader, gd.Int(set), uniforms))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.UniformSetCacheRD

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("UniformSetCacheRD"))
	return Instance{classdb.UniformSetCacheRD(object)}
}

/*
Creates/returns a cached uniform set based on the provided uniforms for a given shader.
*/
//go:nosplit
func (self class) GetCache(shader gd.RID, set gd.Int, uniforms gd.Array) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, set)
	callframe.Arg(frame, pointers.Get(uniforms))
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UniformSetCacheRD.Bind_get_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsUniformSetCacheRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsUniformSetCacheRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("UniformSetCacheRD", func(ptr gd.Object) any { return [1]classdb.UniformSetCacheRD{classdb.UniformSetCacheRD(ptr)} })
}
