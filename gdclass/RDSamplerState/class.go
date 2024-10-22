package RDSamplerState

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
type Go [1]classdb.RDSamplerState
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDSamplerState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDSamplerState"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MagFilter() classdb.RenderingDeviceSamplerFilter {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceSamplerFilter(class(self).GetMagFilter())
}

func (self Go) SetMagFilter(value classdb.RenderingDeviceSamplerFilter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMagFilter(value)
}

func (self Go) MinFilter() classdb.RenderingDeviceSamplerFilter {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceSamplerFilter(class(self).GetMinFilter())
}

func (self Go) SetMinFilter(value classdb.RenderingDeviceSamplerFilter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinFilter(value)
}

func (self Go) MipFilter() classdb.RenderingDeviceSamplerFilter {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceSamplerFilter(class(self).GetMipFilter())
}

func (self Go) SetMipFilter(value classdb.RenderingDeviceSamplerFilter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMipFilter(value)
}

func (self Go) RepeatU() classdb.RenderingDeviceSamplerRepeatMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatU())
}

func (self Go) SetRepeatU(value classdb.RenderingDeviceSamplerRepeatMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRepeatU(value)
}

func (self Go) RepeatV() classdb.RenderingDeviceSamplerRepeatMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatV())
}

func (self Go) SetRepeatV(value classdb.RenderingDeviceSamplerRepeatMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRepeatV(value)
}

func (self Go) RepeatW() classdb.RenderingDeviceSamplerRepeatMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatW())
}

func (self Go) SetRepeatW(value classdb.RenderingDeviceSamplerRepeatMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRepeatW(value)
}

func (self Go) LodBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLodBias()))
}

func (self Go) SetLodBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLodBias(gd.Float(value))
}

func (self Go) UseAnisotropy() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseAnisotropy())
}

func (self Go) SetUseAnisotropy(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseAnisotropy(value)
}

func (self Go) AnisotropyMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAnisotropyMax()))
}

func (self Go) SetAnisotropyMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAnisotropyMax(gd.Float(value))
}

func (self Go) EnableCompare() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableCompare())
}

func (self Go) SetEnableCompare(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableCompare(value)
}

func (self Go) CompareOp() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceCompareOperator(class(self).GetCompareOp())
}

func (self Go) SetCompareOp(value classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCompareOp(value)
}

func (self Go) MinLod() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMinLod()))
}

func (self Go) SetMinLod(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinLod(gd.Float(value))
}

func (self Go) MaxLod() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMaxLod()))
}

func (self Go) SetMaxLod(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxLod(gd.Float(value))
}

func (self Go) BorderColor() classdb.RenderingDeviceSamplerBorderColor {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceSamplerBorderColor(class(self).GetBorderColor())
}

func (self Go) SetBorderColor(value classdb.RenderingDeviceSamplerBorderColor) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderColor(value)
}

func (self Go) UnnormalizedUvw() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUnnormalizedUvw())
}

func (self Go) SetUnnormalizedUvw(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUnnormalizedUvw(value)
}

//go:nosplit
func (self class) SetMagFilter(p_member classdb.RenderingDeviceSamplerFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_mag_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMagFilter() classdb.RenderingDeviceSamplerFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_mag_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinFilter(p_member classdb.RenderingDeviceSamplerFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_min_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinFilter() classdb.RenderingDeviceSamplerFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_min_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMipFilter(p_member classdb.RenderingDeviceSamplerFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_mip_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMipFilter() classdb.RenderingDeviceSamplerFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_mip_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRepeatU(p_member classdb.RenderingDeviceSamplerRepeatMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_repeat_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRepeatU() classdb.RenderingDeviceSamplerRepeatMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerRepeatMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_repeat_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRepeatV(p_member classdb.RenderingDeviceSamplerRepeatMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_repeat_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRepeatV() classdb.RenderingDeviceSamplerRepeatMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerRepeatMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_repeat_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRepeatW(p_member classdb.RenderingDeviceSamplerRepeatMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_repeat_w, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRepeatW() classdb.RenderingDeviceSamplerRepeatMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerRepeatMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_repeat_w, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLodBias(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLodBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseAnisotropy(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_use_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseAnisotropy() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_use_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnisotropyMax(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_anisotropy_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnisotropyMax() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_anisotropy_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableCompare(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_enable_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableCompare() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_enable_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCompareOp(p_member classdb.RenderingDeviceCompareOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_compare_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCompareOp() classdb.RenderingDeviceCompareOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_compare_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinLod(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_min_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinLod() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_min_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxLod(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_max_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxLod() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_max_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBorderColor(p_member classdb.RenderingDeviceSamplerBorderColor)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBorderColor() classdb.RenderingDeviceSamplerBorderColor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceSamplerBorderColor](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUnnormalizedUvw(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_set_unnormalized_uvw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUnnormalizedUvw() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDSamplerState.Bind_get_unnormalized_uvw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDSamplerState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDSamplerState() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDSamplerState", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
