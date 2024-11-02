package CameraAttributesPhysical

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
[CameraAttributesPhysical] is used to set rendering settings based on a physically-based camera's settings. It is responsible for exposure, auto-exposure, and depth of field.
When used in a [WorldEnvironment] it provides default settings for exposure, auto-exposure, and depth of field that will be used by all cameras without their own [CameraAttributes], including the editor camera. When used in a [Camera3D] it will override any [CameraAttributes] set in the [WorldEnvironment] and will override the [Camera3D]s [member Camera3D.far], [member Camera3D.near], [member Camera3D.fov], and [member Camera3D.keep_aspect] properties. When used in [VoxelGI] or [LightmapGI], only the exposure settings will be used.
The default settings are intended for use in an outdoor environment, tips for settings for use in an indoor environment can be found in each setting's documentation.
[b]Note:[/b] Depth of field blur is only supported in the Forward+ and Mobile rendering methods, not Compatibility.
*/
type Instance [1]classdb.CameraAttributesPhysical

/*
Returns the vertical field of view that corresponds to the [member frustum_focal_length]. This value is calculated internally whenever [member frustum_focal_length] is changed.
*/
func (self Instance) GetFov() float64 {
	return float64(float64(class(self).GetFov()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CameraAttributesPhysical

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraAttributesPhysical"))
	return Instance{classdb.CameraAttributesPhysical(object)}
}

func (self Instance) FrustumFocusDistance() float64 {
	return float64(float64(class(self).GetFocusDistance()))
}

func (self Instance) SetFrustumFocusDistance(value float64) {
	class(self).SetFocusDistance(gd.Float(value))
}

func (self Instance) FrustumFocalLength() float64 {
	return float64(float64(class(self).GetFocalLength()))
}

func (self Instance) SetFrustumFocalLength(value float64) {
	class(self).SetFocalLength(gd.Float(value))
}

func (self Instance) FrustumNear() float64 {
	return float64(float64(class(self).GetNear()))
}

func (self Instance) SetFrustumNear(value float64) {
	class(self).SetNear(gd.Float(value))
}

func (self Instance) FrustumFar() float64 {
	return float64(float64(class(self).GetFar()))
}

func (self Instance) SetFrustumFar(value float64) {
	class(self).SetFar(gd.Float(value))
}

func (self Instance) ExposureAperture() float64 {
	return float64(float64(class(self).GetAperture()))
}

func (self Instance) SetExposureAperture(value float64) {
	class(self).SetAperture(gd.Float(value))
}

func (self Instance) ExposureShutterSpeed() float64 {
	return float64(float64(class(self).GetShutterSpeed()))
}

func (self Instance) SetExposureShutterSpeed(value float64) {
	class(self).SetShutterSpeed(gd.Float(value))
}

func (self Instance) AutoExposureMinExposureValue() float64 {
	return float64(float64(class(self).GetAutoExposureMinExposureValue()))
}

func (self Instance) SetAutoExposureMinExposureValue(value float64) {
	class(self).SetAutoExposureMinExposureValue(gd.Float(value))
}

func (self Instance) AutoExposureMaxExposureValue() float64 {
	return float64(float64(class(self).GetAutoExposureMaxExposureValue()))
}

func (self Instance) SetAutoExposureMaxExposureValue(value float64) {
	class(self).SetAutoExposureMaxExposureValue(gd.Float(value))
}

//go:nosplit
func (self class) SetAperture(aperture gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, aperture)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_aperture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAperture() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_aperture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShutterSpeed(shutter_speed gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, shutter_speed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_shutter_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShutterSpeed() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_shutter_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFocalLength(focal_length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, focal_length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_focal_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFocalLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_focal_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFocusDistance(focus_distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, focus_distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_focus_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFocusDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_focus_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNear(near gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, near)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNear() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFar(far gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, far)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFar() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the vertical field of view that corresponds to the [member frustum_focal_length]. This value is calculated internally whenever [member frustum_focal_length] is changed.
*/
//go:nosplit
func (self class) GetFov() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMaxExposureValue(exposure_value_max gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, exposure_value_max)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_auto_exposure_max_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMaxExposureValue() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_auto_exposure_max_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMinExposureValue(exposure_value_min gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, exposure_value_min)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_auto_exposure_min_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMinExposureValue() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_auto_exposure_min_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraAttributesPhysical() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCameraAttributesPhysical() Instance {
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
	classdb.Register("CameraAttributesPhysical", func(ptr gd.Object) any { return classdb.CameraAttributesPhysical(ptr) })
}
