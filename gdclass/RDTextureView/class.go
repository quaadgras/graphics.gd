package RDTextureView

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
This object is used by [RenderingDevice].

*/
type Go [1]classdb.RDTextureView
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDTextureView
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDTextureView"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) FormatOverride() classdb.RenderingDeviceDataFormat {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceDataFormat(class(self).GetFormatOverride())
}

func (self Go) SetFormatOverride(value classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFormatOverride(value)
}

func (self Go) SwizzleR() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleR())
}

func (self Go) SetSwizzleR(value classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSwizzleR(value)
}

func (self Go) SwizzleG() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleG())
}

func (self Go) SetSwizzleG(value classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSwizzleG(value)
}

func (self Go) SwizzleB() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleB())
}

func (self Go) SetSwizzleB(value classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSwizzleB(value)
}

func (self Go) SwizzleA() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureSwizzle(class(self).GetSwizzleA())
}

func (self Go) SetSwizzleA(value classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSwizzleA(value)
}

//go:nosplit
func (self class) SetFormatOverride(p_member classdb.RenderingDeviceDataFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_set_format_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFormatOverride() classdb.RenderingDeviceDataFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceDataFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_get_format_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSwizzleR(p_member classdb.RenderingDeviceTextureSwizzle)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_set_swizzle_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSwizzleR() classdb.RenderingDeviceTextureSwizzle {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_get_swizzle_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSwizzleG(p_member classdb.RenderingDeviceTextureSwizzle)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_set_swizzle_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSwizzleG() classdb.RenderingDeviceTextureSwizzle {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_get_swizzle_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSwizzleB(p_member classdb.RenderingDeviceTextureSwizzle)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_set_swizzle_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSwizzleB() classdb.RenderingDeviceTextureSwizzle {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_get_swizzle_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSwizzleA(p_member classdb.RenderingDeviceTextureSwizzle)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_set_swizzle_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSwizzleA() classdb.RenderingDeviceTextureSwizzle {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSwizzle](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureView.Bind_get_swizzle_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDTextureView() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDTextureView() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("RDTextureView", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
