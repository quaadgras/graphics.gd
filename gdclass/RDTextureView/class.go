package RDTextureView

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This object is used by [RenderingDevice].
*/
type Instance [1]classdb.RDTextureView

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDTextureView

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDTextureView"))
	return Instance{classdb.RDTextureView(object)}
}

func (self Instance) FormatOverride() classdb.RenderingDeviceDataFormat {
	return classdb.RenderingDeviceDataFormat(class(self).GetFormatOverride())
}

func (self Instance) SetFormatOverride(value classdb.RenderingDeviceDataFormat) {
	class(self).SetFormatOverride(value)
}

func (self Instance) SwizzleR() classdb.RenderingDeviceTextureSwizzle {
	return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleR())
}

func (self Instance) SetSwizzleR(value classdb.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleR(value)
}

func (self Instance) SwizzleG() classdb.RenderingDeviceTextureSwizzle {
	return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleG())
}

func (self Instance) SetSwizzleG(value classdb.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleG(value)
}

func (self Instance) SwizzleB() classdb.RenderingDeviceTextureSwizzle {
	return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleB())
}

func (self Instance) SetSwizzleB(value classdb.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleB(value)
}

func (self Instance) SwizzleA() classdb.RenderingDeviceTextureSwizzle {
	return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleA())
}

func (self Instance) SetSwizzleA(value classdb.RenderingDeviceTextureSwizzle) {
	class(self).SetSwizzleA(value)
}

//go:nosplit
func (self class) SetFormatOverride(p_member classdb.RenderingDeviceDataFormat) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_format_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormatOverride() classdb.RenderingDeviceDataFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceDataFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_format_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleR(p_member classdb.RenderingDeviceTextureSwizzle) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleR() classdb.RenderingDeviceTextureSwizzle {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleG(p_member classdb.RenderingDeviceTextureSwizzle) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleG() classdb.RenderingDeviceTextureSwizzle {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleB(p_member classdb.RenderingDeviceTextureSwizzle) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleB() classdb.RenderingDeviceTextureSwizzle {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSwizzleA(p_member classdb.RenderingDeviceTextureSwizzle) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_set_swizzle_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwizzleA() classdb.RenderingDeviceTextureSwizzle {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureView.Bind_get_swizzle_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDTextureView() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDTextureView() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("RDTextureView", func(ptr gd.Object) any { return classdb.RDTextureView(ptr) })
}
