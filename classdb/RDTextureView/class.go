// Package RDTextureView provides methods for working with RDTextureView object instances.
package RDTextureView

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

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDTextureView

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDTextureView() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDTextureView

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDTextureView"))
	casted := Instance{*(*gdclass.RDTextureView)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) FormatOverride() gdclass.RenderingDeviceDataFormat {
	return gdclass.RenderingDeviceDataFormat(class(self).GetFormatOverride())
}

func (self Instance) SetFormatOverride(value gdclass.RenderingDeviceDataFormat) {
	class(self).SetFormatOverride(value)
}

func (self Instance) SwizzleR() gdclass.RenderingDeviceTextureSwizzle {
	return gdclass.RenderingDeviceTextureSwizzle(class(self).GetSwizzleR())
}

func (self Instance) SetSwizzleR(value gdclass.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleR(value)
}

func (self Instance) SwizzleG() gdclass.RenderingDeviceTextureSwizzle {
	return gdclass.RenderingDeviceTextureSwizzle(class(self).GetSwizzleG())
}

func (self Instance) SetSwizzleG(value gdclass.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleG(value)
}

func (self Instance) SwizzleB() gdclass.RenderingDeviceTextureSwizzle {
	return gdclass.RenderingDeviceTextureSwizzle(class(self).GetSwizzleB())
}

func (self Instance) SetSwizzleB(value gdclass.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleB(value)
}

func (self Instance) SwizzleA() gdclass.RenderingDeviceTextureSwizzle {
	return gdclass.RenderingDeviceTextureSwizzle(class(self).GetSwizzleA())
}

func (self Instance) SetSwizzleA(value gdclass.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleA(value)
}

//go:nosplit
func (self class) SetFormatOverride(p_member gdclass.RenderingDeviceDataFormat) { //gd:RDTextureView.set_format_override
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_format_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormatOverride() gdclass.RenderingDeviceDataFormat { //gd:RDTextureView.get_format_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceDataFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_format_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleR(p_member gdclass.RenderingDeviceTextureSwizzle) { //gd:RDTextureView.set_swizzle_r
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_r, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleR() gdclass.RenderingDeviceTextureSwizzle { //gd:RDTextureView.get_swizzle_r
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_r, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleG(p_member gdclass.RenderingDeviceTextureSwizzle) { //gd:RDTextureView.set_swizzle_g
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_g, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleG() gdclass.RenderingDeviceTextureSwizzle { //gd:RDTextureView.get_swizzle_g
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_g, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleB(p_member gdclass.RenderingDeviceTextureSwizzle) { //gd:RDTextureView.set_swizzle_b
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_b, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleB() gdclass.RenderingDeviceTextureSwizzle { //gd:RDTextureView.get_swizzle_b
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_b, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleA(p_member gdclass.RenderingDeviceTextureSwizzle) { //gd:RDTextureView.set_swizzle_a
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_a, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleA() gdclass.RenderingDeviceTextureSwizzle { //gd:RDTextureView.get_swizzle_a
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_a, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDTextureView() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDTextureView() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDTextureView", func(ptr gd.Object) any {
		return [1]gdclass.RDTextureView{*(*gdclass.RDTextureView)(unsafe.Pointer(&ptr))}
	})
}
