// Package RDSamplerState provides methods for working with RDSamplerState object instances.
package RDSamplerState

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDSamplerState

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDSamplerState() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDSamplerState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDSamplerState"))
	casted := Instance{*(*gdclass.RDSamplerState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MagFilter() gdclass.RenderingDeviceSamplerFilter {
	return gdclass.RenderingDeviceSamplerFilter(class(self).GetMagFilter())
}

func (self Instance) SetMagFilter(value gdclass.RenderingDeviceSamplerFilter) {
	class(self).SetMagFilter(value)
}

func (self Instance) MinFilter() gdclass.RenderingDeviceSamplerFilter {
	return gdclass.RenderingDeviceSamplerFilter(class(self).GetMinFilter())
}

func (self Instance) SetMinFilter(value gdclass.RenderingDeviceSamplerFilter) {
	class(self).SetMinFilter(value)
}

func (self Instance) MipFilter() gdclass.RenderingDeviceSamplerFilter {
	return gdclass.RenderingDeviceSamplerFilter(class(self).GetMipFilter())
}

func (self Instance) SetMipFilter(value gdclass.RenderingDeviceSamplerFilter) {
	class(self).SetMipFilter(value)
}

func (self Instance) RepeatU() gdclass.RenderingDeviceSamplerRepeatMode {
	return gdclass.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatU())
}

func (self Instance) SetRepeatU(value gdclass.RenderingDeviceSamplerRepeatMode) {
	class(self).SetRepeatU(value)
}

func (self Instance) RepeatV() gdclass.RenderingDeviceSamplerRepeatMode {
	return gdclass.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatV())
}

func (self Instance) SetRepeatV(value gdclass.RenderingDeviceSamplerRepeatMode) {
	class(self).SetRepeatV(value)
}

func (self Instance) RepeatW() gdclass.RenderingDeviceSamplerRepeatMode {
	return gdclass.RenderingDeviceSamplerRepeatMode(class(self).GetRepeatW())
}

func (self Instance) SetRepeatW(value gdclass.RenderingDeviceSamplerRepeatMode) {
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

func (self Instance) CompareOp() gdclass.RenderingDeviceCompareOperator {
	return gdclass.RenderingDeviceCompareOperator(class(self).GetCompareOp())
}

func (self Instance) SetCompareOp(value gdclass.RenderingDeviceCompareOperator) {
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

func (self Instance) BorderColor() gdclass.RenderingDeviceSamplerBorderColor {
	return gdclass.RenderingDeviceSamplerBorderColor(class(self).GetBorderColor())
}

func (self Instance) SetBorderColor(value gdclass.RenderingDeviceSamplerBorderColor) {
	class(self).SetBorderColor(value)
}

func (self Instance) UnnormalizedUvw() bool {
	return bool(class(self).GetUnnormalizedUvw())
}

func (self Instance) SetUnnormalizedUvw(value bool) {
	class(self).SetUnnormalizedUvw(value)
}

//go:nosplit
func (self class) SetMagFilter(p_member gdclass.RenderingDeviceSamplerFilter) { //gd:RDSamplerState.set_mag_filter
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_mag_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMagFilter() gdclass.RenderingDeviceSamplerFilter { //gd:RDSamplerState.get_mag_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceSamplerFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_mag_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinFilter(p_member gdclass.RenderingDeviceSamplerFilter) { //gd:RDSamplerState.set_min_filter
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_min_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinFilter() gdclass.RenderingDeviceSamplerFilter { //gd:RDSamplerState.get_min_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceSamplerFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_min_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMipFilter(p_member gdclass.RenderingDeviceSamplerFilter) { //gd:RDSamplerState.set_mip_filter
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_mip_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMipFilter() gdclass.RenderingDeviceSamplerFilter { //gd:RDSamplerState.get_mip_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceSamplerFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_mip_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatU(p_member gdclass.RenderingDeviceSamplerRepeatMode) { //gd:RDSamplerState.set_repeat_u
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_repeat_u, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatU() gdclass.RenderingDeviceSamplerRepeatMode { //gd:RDSamplerState.get_repeat_u
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceSamplerRepeatMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_repeat_u, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatV(p_member gdclass.RenderingDeviceSamplerRepeatMode) { //gd:RDSamplerState.set_repeat_v
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_repeat_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatV() gdclass.RenderingDeviceSamplerRepeatMode { //gd:RDSamplerState.get_repeat_v
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceSamplerRepeatMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_repeat_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatW(p_member gdclass.RenderingDeviceSamplerRepeatMode) { //gd:RDSamplerState.set_repeat_w
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_repeat_w, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatW() gdclass.RenderingDeviceSamplerRepeatMode { //gd:RDSamplerState.get_repeat_w
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceSamplerRepeatMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_repeat_w, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLodBias(p_member gd.Float) { //gd:RDSamplerState.set_lod_bias
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_lod_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLodBias() gd.Float { //gd:RDSamplerState.get_lod_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_lod_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAnisotropy(p_member bool) { //gd:RDSamplerState.set_use_anisotropy
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_use_anisotropy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseAnisotropy() bool { //gd:RDSamplerState.get_use_anisotropy
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_use_anisotropy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnisotropyMax(p_member gd.Float) { //gd:RDSamplerState.set_anisotropy_max
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_anisotropy_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnisotropyMax() gd.Float { //gd:RDSamplerState.get_anisotropy_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_anisotropy_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableCompare(p_member bool) { //gd:RDSamplerState.set_enable_compare
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_enable_compare, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableCompare() bool { //gd:RDSamplerState.get_enable_compare
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_enable_compare, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCompareOp(p_member gdclass.RenderingDeviceCompareOperator) { //gd:RDSamplerState.set_compare_op
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_compare_op, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompareOp() gdclass.RenderingDeviceCompareOperator { //gd:RDSamplerState.get_compare_op
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_compare_op, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinLod(p_member gd.Float) { //gd:RDSamplerState.set_min_lod
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_min_lod, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinLod() gd.Float { //gd:RDSamplerState.get_min_lod
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_min_lod, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxLod(p_member gd.Float) { //gd:RDSamplerState.set_max_lod
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_max_lod, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxLod() gd.Float { //gd:RDSamplerState.get_max_lod
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_max_lod, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderColor(p_member gdclass.RenderingDeviceSamplerBorderColor) { //gd:RDSamplerState.set_border_color
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_border_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderColor() gdclass.RenderingDeviceSamplerBorderColor { //gd:RDSamplerState.get_border_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceSamplerBorderColor](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_border_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUnnormalizedUvw(p_member bool) { //gd:RDSamplerState.set_unnormalized_uvw
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_set_unnormalized_uvw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUnnormalizedUvw() bool { //gd:RDSamplerState.get_unnormalized_uvw
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDSamplerState.Bind_get_unnormalized_uvw, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDSamplerState() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDSamplerState() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDSamplerState", func(ptr gd.Object) any {
		return [1]gdclass.RDSamplerState{*(*gdclass.RDSamplerState)(unsafe.Pointer(&ptr))}
	})
}
