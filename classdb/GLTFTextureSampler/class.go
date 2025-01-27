// Package GLTFTextureSampler provides methods for working with GLTFTextureSampler object instances.
package GLTFTextureSampler

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
Represents a texture sampler as defined by the base GLTF spec. Texture samplers in GLTF specify how to sample data from the texture's base image, when rendering the texture on an object.
*/
type Instance [1]gdclass.GLTFTextureSampler

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFTextureSampler() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFTextureSampler

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFTextureSampler"))
	casted := Instance{*(*gdclass.GLTFTextureSampler)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MagFilter() int {
	return int(int(class(self).GetMagFilter()))
}

func (self Instance) SetMagFilter(value int) {
	class(self).SetMagFilter(gd.Int(value))
}

func (self Instance) MinFilter() int {
	return int(int(class(self).GetMinFilter()))
}

func (self Instance) SetMinFilter(value int) {
	class(self).SetMinFilter(gd.Int(value))
}

func (self Instance) WrapS() int {
	return int(int(class(self).GetWrapS()))
}

func (self Instance) SetWrapS(value int) {
	class(self).SetWrapS(gd.Int(value))
}

func (self Instance) WrapT() int {
	return int(int(class(self).GetWrapT()))
}

func (self Instance) SetWrapT(value int) {
	class(self).SetWrapT(gd.Int(value))
}

//go:nosplit
func (self class) GetMagFilter() gd.Int { //gd:GLTFTextureSampler.get_mag_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_mag_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMagFilter(filter_mode gd.Int) { //gd:GLTFTextureSampler.set_mag_filter
	var frame = callframe.New()
	callframe.Arg(frame, filter_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_mag_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinFilter() gd.Int { //gd:GLTFTextureSampler.get_min_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_min_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinFilter(filter_mode gd.Int) { //gd:GLTFTextureSampler.set_min_filter
	var frame = callframe.New()
	callframe.Arg(frame, filter_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_min_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWrapS() gd.Int { //gd:GLTFTextureSampler.get_wrap_s
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_wrap_s, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWrapS(wrap_mode gd.Int) { //gd:GLTFTextureSampler.set_wrap_s
	var frame = callframe.New()
	callframe.Arg(frame, wrap_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_wrap_s, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWrapT() gd.Int { //gd:GLTFTextureSampler.get_wrap_t
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_wrap_t, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWrapT(wrap_mode gd.Int) { //gd:GLTFTextureSampler.set_wrap_t
	var frame = callframe.New()
	callframe.Arg(frame, wrap_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_wrap_t, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsGLTFTextureSampler() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFTextureSampler() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("GLTFTextureSampler", func(ptr gd.Object) any {
		return [1]gdclass.GLTFTextureSampler{*(*gdclass.GLTFTextureSampler)(unsafe.Pointer(&ptr))}
	})
}
