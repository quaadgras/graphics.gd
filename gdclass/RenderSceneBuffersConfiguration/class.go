package RenderSceneBuffersConfiguration

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
This configuration object is created and populated by the render engine on a viewport change and used to (re)configure a [RenderSceneBuffers] object.

*/
type Go [1]classdb.RenderSceneBuffersConfiguration
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RenderSceneBuffersConfiguration
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RenderSceneBuffersConfiguration"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) RenderTarget() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
		return gd.RID(class(self).GetRenderTarget())
}

func (self Go) SetRenderTarget(value gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRenderTarget(value)
}

func (self Go) InternalSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2i(class(self).GetInternalSize())
}

func (self Go) SetInternalSize(value gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInternalSize(value)
}

func (self Go) TargetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2i(class(self).GetTargetSize())
}

func (self Go) SetTargetSize(value gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetSize(value)
}

func (self Go) ViewCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetViewCount()))
}

func (self Go) SetViewCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetViewCount(gd.Int(value))
}

func (self Go) Scaling3dMode() classdb.RenderingServerViewportScaling3DMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingServerViewportScaling3DMode(class(self).GetScaling3dMode())
}

func (self Go) SetScaling3dMode(value classdb.RenderingServerViewportScaling3DMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScaling3dMode(value)
}

func (self Go) Msaa3d() classdb.RenderingServerViewportMSAA {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingServerViewportMSAA(class(self).GetMsaa3d())
}

func (self Go) SetMsaa3d(value classdb.RenderingServerViewportMSAA) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMsaa3d(value)
}

func (self Go) ScreenSpaceAa() classdb.RenderingServerViewportScreenSpaceAA {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingServerViewportScreenSpaceAA(class(self).GetScreenSpaceAa())
}

func (self Go) SetScreenSpaceAa(value classdb.RenderingServerViewportScreenSpaceAA) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScreenSpaceAa(value)
}

func (self Go) FsrSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFsrSharpness()))
}

func (self Go) SetFsrSharpness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFsrSharpness(gd.Float(value))
}

func (self Go) TextureMipmapBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTextureMipmapBias()))
}

func (self Go) SetTextureMipmapBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureMipmapBias(gd.Float(value))
}

//go:nosplit
func (self class) GetRenderTarget() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_render_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderTarget(render_target gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, render_target)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_render_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInternalSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_internal_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInternalSize(internal_size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, internal_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_internal_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_target_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetSize(target_size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, target_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_target_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetViewCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetViewCount(view_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, view_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScaling3dMode() classdb.RenderingServerViewportScaling3DMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportScaling3DMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScaling3dMode(scaling_3d_mode classdb.RenderingServerViewportScaling3DMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scaling_3d_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsaa3d() classdb.RenderingServerViewportMSAA {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportMSAA](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsaa3d(msaa_3d classdb.RenderingServerViewportMSAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msaa_3d)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScreenSpaceAa() classdb.RenderingServerViewportScreenSpaceAA {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportScreenSpaceAA](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScreenSpaceAa(screen_space_aa classdb.RenderingServerViewportScreenSpaceAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_space_aa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFsrSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFsrSharpness(fsr_sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fsr_sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureMipmapBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_get_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureMipmapBias(texture_mipmap_bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture_mipmap_bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersConfiguration.Bind_set_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsRenderSceneBuffersConfiguration() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderSceneBuffersConfiguration() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RenderSceneBuffersConfiguration", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
