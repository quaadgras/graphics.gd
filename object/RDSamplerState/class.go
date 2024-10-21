package RDSamplerState

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
type Simple [1]classdb.RDSamplerState
func (self Simple) SetMagFilter(p_member classdb.RenderingDeviceSamplerFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMagFilter(p_member)
}
func (self Simple) GetMagFilter() classdb.RenderingDeviceSamplerFilter {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceSamplerFilter(Expert(self).GetMagFilter())
}
func (self Simple) SetMinFilter(p_member classdb.RenderingDeviceSamplerFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinFilter(p_member)
}
func (self Simple) GetMinFilter() classdb.RenderingDeviceSamplerFilter {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceSamplerFilter(Expert(self).GetMinFilter())
}
func (self Simple) SetMipFilter(p_member classdb.RenderingDeviceSamplerFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMipFilter(p_member)
}
func (self Simple) GetMipFilter() classdb.RenderingDeviceSamplerFilter {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceSamplerFilter(Expert(self).GetMipFilter())
}
func (self Simple) SetRepeatU(p_member classdb.RenderingDeviceSamplerRepeatMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRepeatU(p_member)
}
func (self Simple) GetRepeatU() classdb.RenderingDeviceSamplerRepeatMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceSamplerRepeatMode(Expert(self).GetRepeatU())
}
func (self Simple) SetRepeatV(p_member classdb.RenderingDeviceSamplerRepeatMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRepeatV(p_member)
}
func (self Simple) GetRepeatV() classdb.RenderingDeviceSamplerRepeatMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceSamplerRepeatMode(Expert(self).GetRepeatV())
}
func (self Simple) SetRepeatW(p_member classdb.RenderingDeviceSamplerRepeatMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRepeatW(p_member)
}
func (self Simple) GetRepeatW() classdb.RenderingDeviceSamplerRepeatMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceSamplerRepeatMode(Expert(self).GetRepeatW())
}
func (self Simple) SetLodBias(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLodBias(gd.Float(p_member))
}
func (self Simple) GetLodBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLodBias()))
}
func (self Simple) SetUseAnisotropy(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseAnisotropy(p_member)
}
func (self Simple) GetUseAnisotropy() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseAnisotropy())
}
func (self Simple) SetAnisotropyMax(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnisotropyMax(gd.Float(p_member))
}
func (self Simple) GetAnisotropyMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAnisotropyMax()))
}
func (self Simple) SetEnableCompare(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableCompare(p_member)
}
func (self Simple) GetEnableCompare() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableCompare())
}
func (self Simple) SetCompareOp(p_member classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCompareOp(p_member)
}
func (self Simple) GetCompareOp() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceCompareOperator(Expert(self).GetCompareOp())
}
func (self Simple) SetMinLod(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinLod(gd.Float(p_member))
}
func (self Simple) GetMinLod() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMinLod()))
}
func (self Simple) SetMaxLod(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxLod(gd.Float(p_member))
}
func (self Simple) GetMaxLod() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMaxLod()))
}
func (self Simple) SetBorderColor(p_member classdb.RenderingDeviceSamplerBorderColor) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBorderColor(p_member)
}
func (self Simple) GetBorderColor() classdb.RenderingDeviceSamplerBorderColor {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceSamplerBorderColor(Expert(self).GetBorderColor())
}
func (self Simple) SetUnnormalizedUvw(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUnnormalizedUvw(p_member)
}
func (self Simple) GetUnnormalizedUvw() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUnnormalizedUvw())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDSamplerState
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsRDSamplerState() Expert { return self[0].AsRDSamplerState() }


//go:nosplit
func (self Simple) AsRDSamplerState() Simple { return self[0].AsRDSamplerState() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RDSamplerState", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
