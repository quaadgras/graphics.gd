package RDSamplerState

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
type Instance [1]classdb.RDSamplerState

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDSamplerState

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDSamplerState"))
	return Instance{classdb.RDSamplerState(object)}
}

func (self Instance) MagFilter() classdb.RenderingDeviceSamplerFilter {
	return classdb.RenderingDeviceSamplerFilter(class(self).GetMagFilter())
}

func (self Instance) SetMagFilter(value classdb.RenderingDeviceSamplerFilter) {
	class(self).SetMagFilter(value)
}

func (self Instance) MinFilter() classdb.RenderingDeviceSamplerFilter {
	return classdb.RenderingDeviceSamplerFilter(class(self).GetMinFilter())
}

func (self Instance) SetMinFilter(value classdb.RenderingDeviceSamplerFilter) {
	class(self).SetMinFilter(value)
}

func (self Instance) MipFilter() classdb.RenderingDeviceSamplerFilter {
	return classdb.RenderingDeviceSamplerFilter(class(self).GetMipFilter())
}

func (self Instance) SetMipFilter(value classdb.RenderingDeviceSamplerFilter) {
	class(self).SetMipFilter(value)
}

func (self Instance) RepeatU() classdb.RenderingDeviceSamplerRepeatMode {
	return classdb.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatU())
}

func (self Instance) SetRepeatU(value classdb.RenderingDeviceSamplerRepeatMode) {
	class(self).SetRepeatU(value)
}

func (self Instance) RepeatV() classdb.RenderingDeviceSamplerRepeatMode {
	return classdb.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatV())
}

func (self Instance) SetRepeatV(value classdb.RenderingDeviceSamplerRepeatMode) {
	class(self).SetRepeatV(value)
}

func (self Instance) RepeatW() classdb.RenderingDeviceSamplerRepeatMode {
	return classdb.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatW())
}

func (self Instance) SetRepeatW(value classdb.RenderingDeviceSamplerRepeatMode) {
	class(self).SetRepeatW(value)
}

func (self Instance) LodBias() Float.X {
	return Float.X(Float.X(class(self).GetLodBias()))
}

func (self Instance) SetLodBias(value Float.X) {
	class(self).SetLodBias(gd.Float(value))
}

func (self Instance) UseAnisotropy() bool {
	return bool(class(self).GetUseAnisotropy())
}

func (self Instance) SetUseAnisotropy(value bool) {
	class(self).SetUseAnisotropy(value)
}

func (self Instance) AnisotropyMax() Float.X {
	return Float.X(Float.X(class(self).GetAnisotropyMax()))
}

func (self Instance) SetAnisotropyMax(value Float.X) {
	class(self).SetAnisotropyMax(gd.Float(value))
}

func (self Instance) EnableCompare() bool {
	return bool(class(self).GetEnableCompare())
}

func (self Instance) SetEnableCompare(value bool) {
	class(self).SetEnableCompare(value)
}

func (self Instance) CompareOp() classdb.RenderingDeviceCompareOperator {
	return classdb.RenderingDeviceCompareOperator(class(self).GetCompareOp())
}

func (self Instance) SetCompareOp(value classdb.RenderingDeviceCompareOperator) {
	class(self).SetCompareOp(value)
}

func (self Instance) MinLod() Float.X {
	return Float.X(Float.X(class(self).GetMinLod()))
}

func (self Instance) SetMinLod(value Float.X) {
	class(self).SetMinLod(gd.Float(value))
}

func (self Instance) MaxLod() Float.X {
	return Float.X(Float.X(class(self).GetMaxLod()))
}

func (self Instance) SetMaxLod(value Float.X) {
	class(self).SetMaxLod(gd.Float(value))
}

func (self Instance) BorderColor() classdb.RenderingDeviceSamplerBorderColor {
	return classdb.RenderingDeviceSamplerBorderColor(class(self).GetBorderColor())
}

func (self Instance) SetBorderColor(value classdb.RenderingDeviceSamplerBorderColor) {
	class(self).SetBorderColor(value)
}

func (self Instance) UnnormalizedUvw() bool {
	return bool(class(self).GetUnnormalizedUvw())
}

func (self Instance) SetUnnormalizedUvw(value bool) {
	class(self).SetUnnormalizedUvw(value)
}

//go:nosplit
func (self class) SetMagFilter(p_member classdb.RenderingDeviceSamplerFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_mag_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMagFilter() classdb.RenderingDeviceSamplerFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_mag_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinFilter(p_member classdb.RenderingDeviceSamplerFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_min_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinFilter() classdb.RenderingDeviceSamplerFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_min_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMipFilter(p_member classdb.RenderingDeviceSamplerFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_mip_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMipFilter() classdb.RenderingDeviceSamplerFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_mip_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatU(p_member classdb.RenderingDeviceSamplerRepeatMode) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_repeat_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatU() classdb.RenderingDeviceSamplerRepeatMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerRepeatMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_repeat_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatV(p_member classdb.RenderingDeviceSamplerRepeatMode) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_repeat_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatV() classdb.RenderingDeviceSamplerRepeatMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerRepeatMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_repeat_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatW(p_member classdb.RenderingDeviceSamplerRepeatMode) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_repeat_w, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatW() classdb.RenderingDeviceSamplerRepeatMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerRepeatMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_repeat_w, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLodBias(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLodBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAnisotropy(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_use_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseAnisotropy() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_use_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnisotropyMax(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_anisotropy_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnisotropyMax() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_anisotropy_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableCompare(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_enable_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableCompare() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_enable_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCompareOp(p_member classdb.RenderingDeviceCompareOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_compare_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompareOp() classdb.RenderingDeviceCompareOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_compare_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinLod(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_min_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinLod() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_min_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxLod(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_max_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxLod() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_max_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderColor(p_member classdb.RenderingDeviceSamplerBorderColor) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderColor() classdb.RenderingDeviceSamplerBorderColor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerBorderColor](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUnnormalizedUvw(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_unnormalized_uvw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUnnormalizedUvw() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_unnormalized_uvw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDSamplerState() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDSamplerState() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("RDSamplerState", func(ptr gd.Object) any { return classdb.RDSamplerState(ptr) })
}
