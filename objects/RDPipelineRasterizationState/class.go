package RDPipelineRasterizationState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This object is used by [RenderingDevice].
*/
type Instance [1]classdb.RDPipelineRasterizationState

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDPipelineRasterizationState

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineRasterizationState"))
	return Instance{classdb.RDPipelineRasterizationState(object)}
}

func (self Instance) EnableDepthClamp() bool {
	return bool(class(self).GetEnableDepthClamp())
}

func (self Instance) SetEnableDepthClamp(value bool) {
	class(self).SetEnableDepthClamp(value)
}

func (self Instance) DiscardPrimitives() bool {
	return bool(class(self).GetDiscardPrimitives())
}

func (self Instance) SetDiscardPrimitives(value bool) {
	class(self).SetDiscardPrimitives(value)
}

func (self Instance) Wireframe() bool {
	return bool(class(self).GetWireframe())
}

func (self Instance) SetWireframe(value bool) {
	class(self).SetWireframe(value)
}

func (self Instance) CullMode() classdb.RenderingDevicePolygonCullMode {
	return classdb.RenderingDevicePolygonCullMode(class(self).GetCullMode())
}

func (self Instance) SetCullMode(value classdb.RenderingDevicePolygonCullMode) {
	class(self).SetCullMode(value)
}

func (self Instance) FrontFace() classdb.RenderingDevicePolygonFrontFace {
	return classdb.RenderingDevicePolygonFrontFace(class(self).GetFrontFace())
}

func (self Instance) SetFrontFace(value classdb.RenderingDevicePolygonFrontFace) {
	class(self).SetFrontFace(value)
}

func (self Instance) DepthBiasEnabled() bool {
	return bool(class(self).GetDepthBiasEnabled())
}

func (self Instance) SetDepthBiasEnabled(value bool) {
	class(self).SetDepthBiasEnabled(value)
}

func (self Instance) DepthBiasConstantFactor() Float.X {
	return Float.X(Float.X(class(self).GetDepthBiasConstantFactor()))
}

func (self Instance) SetDepthBiasConstantFactor(value Float.X) {
	class(self).SetDepthBiasConstantFactor(gd.Float(value))
}

func (self Instance) DepthBiasClamp() Float.X {
	return Float.X(Float.X(class(self).GetDepthBiasClamp()))
}

func (self Instance) SetDepthBiasClamp(value Float.X) {
	class(self).SetDepthBiasClamp(gd.Float(value))
}

func (self Instance) DepthBiasSlopeFactor() Float.X {
	return Float.X(Float.X(class(self).GetDepthBiasSlopeFactor()))
}

func (self Instance) SetDepthBiasSlopeFactor(value Float.X) {
	class(self).SetDepthBiasSlopeFactor(gd.Float(value))
}

func (self Instance) LineWidth() Float.X {
	return Float.X(Float.X(class(self).GetLineWidth()))
}

func (self Instance) SetLineWidth(value Float.X) {
	class(self).SetLineWidth(gd.Float(value))
}

func (self Instance) PatchControlPoints() int {
	return int(int(class(self).GetPatchControlPoints()))
}

func (self Instance) SetPatchControlPoints(value int) {
	class(self).SetPatchControlPoints(gd.Int(value))
}

//go:nosplit
func (self class) SetEnableDepthClamp(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_enable_depth_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthClamp() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_enable_depth_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDiscardPrimitives(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_discard_primitives, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDiscardPrimitives() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_discard_primitives, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWireframe(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_wireframe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWireframe() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_wireframe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMode(p_member classdb.RenderingDevicePolygonCullMode) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMode() classdb.RenderingDevicePolygonCullMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDevicePolygonCullMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontFace(p_member classdb.RenderingDevicePolygonFrontFace) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_front_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontFace() classdb.RenderingDevicePolygonFrontFace {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDevicePolygonFrontFace](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_front_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasEnabled(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasConstantFactor(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_constant_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasConstantFactor() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_constant_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasClamp(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasClamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_clamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasSlopeFactor(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_slope_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasSlopeFactor() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_slope_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLineWidth(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLineWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPatchControlPoints(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_patch_control_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPatchControlPoints() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_patch_control_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineRasterizationState() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineRasterizationState() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
	classdb.Register("RDPipelineRasterizationState", func(ptr gd.Object) any { return classdb.RDPipelineRasterizationState(ptr) })
}
