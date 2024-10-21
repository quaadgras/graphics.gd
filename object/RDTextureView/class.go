package RDTextureView

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
This object is used by [RenderingDevice].

*/
type Simple [1]classdb.RDTextureView
func (self Simple) SetFormatOverride(p_member classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFormatOverride(p_member)
}
func (self Simple) GetFormatOverride() classdb.RenderingDeviceDataFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceDataFormat(Expert(self).GetFormatOverride())
}
func (self Simple) SetSwizzleR(p_member classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSwizzleR(p_member)
}
func (self Simple) GetSwizzleR() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSwizzle(Expert(self).GetSwizzleR())
}
func (self Simple) SetSwizzleG(p_member classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSwizzleG(p_member)
}
func (self Simple) GetSwizzleG() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSwizzle(Expert(self).GetSwizzleG())
}
func (self Simple) SetSwizzleB(p_member classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSwizzleB(p_member)
}
func (self Simple) GetSwizzleB() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSwizzle(Expert(self).GetSwizzleB())
}
func (self Simple) SetSwizzleA(p_member classdb.RenderingDeviceTextureSwizzle) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSwizzleA(p_member)
}
func (self Simple) GetSwizzleA() classdb.RenderingDeviceTextureSwizzle {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSwizzle(Expert(self).GetSwizzleA())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDTextureView
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsRDTextureView() Expert { return self[0].AsRDTextureView() }


//go:nosplit
func (self Simple) AsRDTextureView() Simple { return self[0].AsRDTextureView() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RDTextureView", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
