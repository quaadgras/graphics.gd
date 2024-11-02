package CameraAttributesPractical

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/CameraAttributes"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Controls camera-specific attributes such as auto-exposure, depth of field, and exposure override.
When used in a [WorldEnvironment] it provides default settings for exposure, auto-exposure, and depth of field that will be used by all cameras without their own [CameraAttributes], including the editor camera. When used in a [Camera3D] it will override any [CameraAttributes] set in the [WorldEnvironment]. When used in [VoxelGI] or [LightmapGI], only the exposure settings will be used.
*/
type Instance [1]classdb.CameraAttributesPractical

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CameraAttributesPractical

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraAttributesPractical"))
	return Instance{classdb.CameraAttributesPractical(object)}
}

func (self Instance) DofBlurFarEnabled() bool {
	return bool(class(self).IsDofBlurFarEnabled())
}

func (self Instance) SetDofBlurFarEnabled(value bool) {
	class(self).SetDofBlurFarEnabled(value)
}

func (self Instance) DofBlurFarDistance() float64 {
	return float64(float64(class(self).GetDofBlurFarDistance()))
}

func (self Instance) SetDofBlurFarDistance(value float64) {
	class(self).SetDofBlurFarDistance(gd.Float(value))
}

func (self Instance) DofBlurFarTransition() float64 {
	return float64(float64(class(self).GetDofBlurFarTransition()))
}

func (self Instance) SetDofBlurFarTransition(value float64) {
	class(self).SetDofBlurFarTransition(gd.Float(value))
}

func (self Instance) DofBlurNearEnabled() bool {
	return bool(class(self).IsDofBlurNearEnabled())
}

func (self Instance) SetDofBlurNearEnabled(value bool) {
	class(self).SetDofBlurNearEnabled(value)
}

func (self Instance) DofBlurNearDistance() float64 {
	return float64(float64(class(self).GetDofBlurNearDistance()))
}

func (self Instance) SetDofBlurNearDistance(value float64) {
	class(self).SetDofBlurNearDistance(gd.Float(value))
}

func (self Instance) DofBlurNearTransition() float64 {
	return float64(float64(class(self).GetDofBlurNearTransition()))
}

func (self Instance) SetDofBlurNearTransition(value float64) {
	class(self).SetDofBlurNearTransition(gd.Float(value))
}

func (self Instance) DofBlurAmount() float64 {
	return float64(float64(class(self).GetDofBlurAmount()))
}

func (self Instance) SetDofBlurAmount(value float64) {
	class(self).SetDofBlurAmount(gd.Float(value))
}

func (self Instance) AutoExposureMinSensitivity() float64 {
	return float64(float64(class(self).GetAutoExposureMinSensitivity()))
}

func (self Instance) SetAutoExposureMinSensitivity(value float64) {
	class(self).SetAutoExposureMinSensitivity(gd.Float(value))
}

func (self Instance) AutoExposureMaxSensitivity() float64 {
	return float64(float64(class(self).GetAutoExposureMaxSensitivity()))
}

func (self Instance) SetAutoExposureMaxSensitivity(value float64) {
	class(self).SetAutoExposureMaxSensitivity(gd.Float(value))
}

//go:nosplit
func (self class) SetDofBlurFarEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_far_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDofBlurFarEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_is_dof_blur_far_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurFarDistance(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_far_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurFarDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_far_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurFarTransition(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_far_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurFarTransition() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_far_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurNearEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_near_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDofBlurNearEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_is_dof_blur_near_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurNearDistance(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_near_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurNearDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_near_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurNearTransition(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_near_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurNearTransition() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_near_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDofBlurAmount(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_dof_blur_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDofBlurAmount() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_dof_blur_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMaxSensitivity(max_sensitivity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, max_sensitivity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_auto_exposure_max_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMaxSensitivity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_auto_exposure_max_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMinSensitivity(min_sensitivity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, min_sensitivity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_set_auto_exposure_min_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMinSensitivity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPractical.Bind_get_auto_exposure_min_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsCameraAttributes(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsCameraAttributes(), name)
	}
}
func init() {
	classdb.Register("CameraAttributesPractical", func(ptr gd.Object) any { return classdb.CameraAttributesPractical(ptr) })
}
