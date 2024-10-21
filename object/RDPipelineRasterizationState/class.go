package RDPipelineRasterizationState

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
This object is used by [RenderingDevice].

*/
type Simple [1]classdb.RDPipelineRasterizationState
func (self Simple) SetEnableDepthClamp(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableDepthClamp(p_member)
}
func (self Simple) GetEnableDepthClamp() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableDepthClamp())
}
func (self Simple) SetDiscardPrimitives(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDiscardPrimitives(p_member)
}
func (self Simple) GetDiscardPrimitives() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDiscardPrimitives())
}
func (self Simple) SetWireframe(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWireframe(p_member)
}
func (self Simple) GetWireframe() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetWireframe())
}
func (self Simple) SetCullMode(p_member classdb.RenderingDevicePolygonCullMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCullMode(p_member)
}
func (self Simple) GetCullMode() classdb.RenderingDevicePolygonCullMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDevicePolygonCullMode(Expert(self).GetCullMode())
}
func (self Simple) SetFrontFace(p_member classdb.RenderingDevicePolygonFrontFace) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontFace(p_member)
}
func (self Simple) GetFrontFace() classdb.RenderingDevicePolygonFrontFace {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDevicePolygonFrontFace(Expert(self).GetFrontFace())
}
func (self Simple) SetDepthBiasEnabled(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthBiasEnabled(p_member)
}
func (self Simple) GetDepthBiasEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDepthBiasEnabled())
}
func (self Simple) SetDepthBiasConstantFactor(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthBiasConstantFactor(gd.Float(p_member))
}
func (self Simple) GetDepthBiasConstantFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepthBiasConstantFactor()))
}
func (self Simple) SetDepthBiasClamp(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthBiasClamp(gd.Float(p_member))
}
func (self Simple) GetDepthBiasClamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepthBiasClamp()))
}
func (self Simple) SetDepthBiasSlopeFactor(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthBiasSlopeFactor(gd.Float(p_member))
}
func (self Simple) GetDepthBiasSlopeFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepthBiasSlopeFactor()))
}
func (self Simple) SetLineWidth(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLineWidth(gd.Float(p_member))
}
func (self Simple) GetLineWidth() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLineWidth()))
}
func (self Simple) SetPatchControlPoints(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPatchControlPoints(gd.Int(p_member))
}
func (self Simple) GetPatchControlPoints() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPatchControlPoints()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDPipelineRasterizationState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsRDPipelineRasterizationState() Expert { return self[0].AsRDPipelineRasterizationState() }


//go:nosplit
func (self Simple) AsRDPipelineRasterizationState() Simple { return self[0].AsRDPipelineRasterizationState() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
func init() {classdb.Register("RDPipelineRasterizationState", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
