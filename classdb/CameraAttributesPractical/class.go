// Package CameraAttributesPractical provides methods for working with CameraAttributesPractical object instances.
package CameraAttributesPractical

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
import "graphics.gd/classdb/CameraAttributes"
import "graphics.gd/classdb/Resource"
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
Controls camera-specific attributes such as auto-exposure, depth of field, and exposure override.
When used in a [WorldEnvironment] it provides default settings for exposure, auto-exposure, and depth of field that will be used by all cameras without their own [CameraAttributes], including the editor camera. When used in a [Camera3D] it will override any [CameraAttributes] set in the [WorldEnvironment]. When used in [VoxelGI] or [LightmapGI], only the exposure settings will be used.
*/
type Instance [1]gdclass.CameraAttributesPractical

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCameraAttributesPractical() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CameraAttributesPractical

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraAttributesPractical"))
	casted := Instance{*(*gdclass.CameraAttributesPractical)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) DofBlurFarEnabled() bool {
	return bool(class(self).IsDofBlurFarEnabled())
}

func (self Instance) SetDofBlurFarEnabled(value bool) {
	class(self).SetDofBlurFarEnabled(value)
}

func (self Instance) DofBlurFarDistance() Float.X {
	return Float.X(Float.X(class(self).GetDofBlurFarDistance()))
}

func (self Instance) SetDofBlurFarDistance(value Float.X) {
	class(self).SetDofBlurFarDistance(gd.Float(value))
}

func (self Instance) DofBlurFarTransition() Float.X {
	return Float.X(Float.X(class(self).GetDofBlurFarTransition()))
}

func (self Instance) SetDofBlurFarTransition(value Float.X) {
	class(self).SetDofBlurFarTransition(gd.Float(value))
}

func (self Instance) DofBlurNearEnabled() bool {
	return bool(class(self).IsDofBlurNearEnabled())
}

func (self Instance) SetDofBlurNearEnabled(value bool) {
	class(self).SetDofBlurNearEnabled(value)
}

func (self Instance) DofBlurNearDistance() Float.X {
	return Float.X(Float.X(class(self).GetDofBlurNearDistance()))
}

func (self Instance) SetDofBlurNearDistance(value Float.X) {
	class(self).SetDofBlurNearDistance(gd.Float(value))
}

func (self Instance) DofBlurNearTransition() Float.X {
	return Float.X(Float.X(class(self).GetDofBlurNearTransition()))
}

func (self Instance) SetDofBlurNearTransition(value Float.X) {
	class(self).SetDofBlurNearTransition(gd.Float(value))
}

func (self Instance) DofBlurAmount() Float.X {
	return Float.X(Float.X(class(self).GetDofBlurAmount()))
}

func (self Instance) SetDofBlurAmount(value Float.X) {
	class(self).SetDofBlurAmount(gd.Float(value))
}

func (self Instance) AutoExposureMinSensitivity() Float.X {
	return Float.X(Float.X(class(self).GetAutoExposureMinSensitivity()))
}

func (self Instance) SetAutoExposureMinSensitivity(value Float.X) {
	class(self).SetAutoExposureMinSensitivity(gd.Float(value))
}

func (self Instance) AutoExposureMaxSensitivity() Float.X {
	return Float.X(Float.X(class(self).GetAutoExposureMaxSensitivity()))
}

func (self Instance) SetAutoExposureMaxSensitivity(value Float.X) {
	class(self).SetAutoExposureMaxSensitivity(gd.Float(value))
}

//go:nosplit
func (self class) SetDofBlurFarEnabled(enabled bool) { //gd:CameraAttributesPractical.set_dof_blur_far_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_far_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDofBlurFarEnabled() bool { //gd:CameraAttributesPractical.is_dof_blur_far_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_is_dof_blur_far_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurFarDistance(distance gd.Float) { //gd:CameraAttributesPractical.set_dof_blur_far_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_far_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurFarDistance() gd.Float { //gd:CameraAttributesPractical.get_dof_blur_far_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_far_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurFarTransition(distance gd.Float) { //gd:CameraAttributesPractical.set_dof_blur_far_transition
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_far_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurFarTransition() gd.Float { //gd:CameraAttributesPractical.get_dof_blur_far_transition
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_far_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurNearEnabled(enabled bool) { //gd:CameraAttributesPractical.set_dof_blur_near_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_near_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDofBlurNearEnabled() bool { //gd:CameraAttributesPractical.is_dof_blur_near_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_is_dof_blur_near_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurNearDistance(distance gd.Float) { //gd:CameraAttributesPractical.set_dof_blur_near_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_near_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurNearDistance() gd.Float { //gd:CameraAttributesPractical.get_dof_blur_near_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_near_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurNearTransition(distance gd.Float) { //gd:CameraAttributesPractical.set_dof_blur_near_transition
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_near_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurNearTransition() gd.Float { //gd:CameraAttributesPractical.get_dof_blur_near_transition
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_near_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurAmount(amount gd.Float) { //gd:CameraAttributesPractical.set_dof_blur_amount
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurAmount() gd.Float { //gd:CameraAttributesPractical.get_dof_blur_amount
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMaxSensitivity(max_sensitivity gd.Float) { //gd:CameraAttributesPractical.set_auto_exposure_max_sensitivity
	var frame = callframe.New()
	callframe.Arg(frame, max_sensitivity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_auto_exposure_max_sensitivity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMaxSensitivity() gd.Float { //gd:CameraAttributesPractical.get_auto_exposure_max_sensitivity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_auto_exposure_max_sensitivity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMinSensitivity(min_sensitivity gd.Float) { //gd:CameraAttributesPractical.set_auto_exposure_min_sensitivity
	var frame = callframe.New()
	callframe.Arg(frame, min_sensitivity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_auto_exposure_min_sensitivity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMinSensitivity() gd.Float { //gd:CameraAttributesPractical.get_auto_exposure_min_sensitivity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_auto_exposure_min_sensitivity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraAttributesPractical() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCameraAttributesPractical() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCameraAttributes() CameraAttributes.Advanced {
	return *((*CameraAttributes.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCameraAttributes() CameraAttributes.Instance {
	return *((*CameraAttributes.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CameraAttributes.Advanced(self.AsCameraAttributes()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CameraAttributes.Instance(self.AsCameraAttributes()), name)
	}
}
func init() {
	gdclass.Register("CameraAttributesPractical", func(ptr gd.Object) any {
		return [1]gdclass.CameraAttributesPractical{*(*gdclass.CameraAttributesPractical)(unsafe.Pointer(&ptr))}
	})
}
