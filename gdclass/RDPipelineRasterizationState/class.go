package RDPipelineRasterizationState

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
type Go [1]classdb.RDPipelineRasterizationState
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDPipelineRasterizationState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDPipelineRasterizationState"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EnableDepthClamp() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableDepthClamp())
}

func (self Go) SetEnableDepthClamp(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableDepthClamp(value)
}

func (self Go) DiscardPrimitives() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDiscardPrimitives())
}

func (self Go) SetDiscardPrimitives(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDiscardPrimitives(value)
}

func (self Go) Wireframe() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetWireframe())
}

func (self Go) SetWireframe(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWireframe(value)
}

func (self Go) CullMode() classdb.RenderingDevicePolygonCullMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDevicePolygonCullMode(class(self).GetCullMode())
}

func (self Go) SetCullMode(value classdb.RenderingDevicePolygonCullMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCullMode(value)
}

func (self Go) FrontFace() classdb.RenderingDevicePolygonFrontFace {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDevicePolygonFrontFace(class(self).GetFrontFace())
}

func (self Go) SetFrontFace(value classdb.RenderingDevicePolygonFrontFace) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontFace(value)
}

func (self Go) DepthBiasEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDepthBiasEnabled())
}

func (self Go) SetDepthBiasEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepthBiasEnabled(value)
}

func (self Go) DepthBiasConstantFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDepthBiasConstantFactor()))
}

func (self Go) SetDepthBiasConstantFactor(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepthBiasConstantFactor(gd.Float(value))
}

func (self Go) DepthBiasClamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDepthBiasClamp()))
}

func (self Go) SetDepthBiasClamp(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepthBiasClamp(gd.Float(value))
}

func (self Go) DepthBiasSlopeFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDepthBiasSlopeFactor()))
}

func (self Go) SetDepthBiasSlopeFactor(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepthBiasSlopeFactor(gd.Float(value))
}

func (self Go) LineWidth() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLineWidth()))
}

func (self Go) SetLineWidth(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineWidth(gd.Float(value))
}

func (self Go) PatchControlPoints() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPatchControlPoints()))
}

func (self Go) SetPatchControlPoints(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPatchControlPoints(gd.Int(value))
}

//go:nosplit
func (self class) SetEnableDepthClamp(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_enable_depth_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableDepthClamp() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_enable_depth_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiscardPrimitives(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_discard_primitives, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDiscardPrimitives() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_discard_primitives, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWireframe(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_wireframe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWireframe() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_wireframe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCullMode(p_member classdb.RenderingDevicePolygonCullMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCullMode() classdb.RenderingDevicePolygonCullMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDevicePolygonCullMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontFace(p_member classdb.RenderingDevicePolygonFrontFace)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_front_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontFace() classdb.RenderingDevicePolygonFrontFace {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDevicePolygonFrontFace](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_front_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthBiasEnabled(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_depth_bias_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthBiasEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_depth_bias_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthBiasConstantFactor(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_depth_bias_constant_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthBiasConstantFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_depth_bias_constant_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthBiasClamp(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_depth_bias_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthBiasClamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_depth_bias_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthBiasSlopeFactor(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_depth_bias_slope_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthBiasSlopeFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_depth_bias_slope_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLineWidth(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLineWidth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPatchControlPoints(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_set_patch_control_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPatchControlPoints() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineRasterizationState.Bind_get_patch_control_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineRasterizationState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDPipelineRasterizationState() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDPipelineRasterizationState", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
