// Package CameraAttributesPhysical provides methods for working with CameraAttributesPhysical object instances.
package CameraAttributesPhysical

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

/*
[CameraAttributesPhysical] is used to set rendering settings based on a physically-based camera's settings. It is responsible for exposure, auto-exposure, and depth of field.
When used in a [WorldEnvironment] it provides default settings for exposure, auto-exposure, and depth of field that will be used by all cameras without their own [CameraAttributes], including the editor camera. When used in a [Camera3D] it will override any [CameraAttributes] set in the [WorldEnvironment] and will override the [Camera3D]s [member Camera3D.far], [member Camera3D.near], [member Camera3D.fov], and [member Camera3D.keep_aspect] properties. When used in [VoxelGI] or [LightmapGI], only the exposure settings will be used.
The default settings are intended for use in an outdoor environment, tips for settings for use in an indoor environment can be found in each setting's documentation.
[b]Note:[/b] Depth of field blur is only supported in the Forward+ and Mobile rendering methods, not Compatibility.
*/
type Instance [1]gdclass.CameraAttributesPhysical

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCameraAttributesPhysical() Instance
}

/*
Returns the vertical field of view that corresponds to the [member frustum_focal_length]. This value is calculated internally whenever [member frustum_focal_length] is changed.
*/
func (self Instance) GetFov() Float.X { //gd:CameraAttributesPhysical.get_fov
	return Float.X(Float.X(class(self).GetFov()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CameraAttributesPhysical

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraAttributesPhysical"))
	casted := Instance{*(*gdclass.CameraAttributesPhysical)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) FrustumFocusDistance() Float.X {
	return Float.X(Float.X(class(self).GetFocusDistance()))
}

func (self Instance) SetFrustumFocusDistance(value Float.X) {
	class(self).SetFocusDistance(gd.Float(value))
}

func (self Instance) FrustumFocalLength() Float.X {
	return Float.X(Float.X(class(self).GetFocalLength()))
}

func (self Instance) SetFrustumFocalLength(value Float.X) {
	class(self).SetFocalLength(gd.Float(value))
}

func (self Instance) FrustumNear() Float.X {
	return Float.X(Float.X(class(self).GetNear()))
}

func (self Instance) SetFrustumNear(value Float.X) {
	class(self).SetNear(gd.Float(value))
}

func (self Instance) FrustumFar() Float.X {
	return Float.X(Float.X(class(self).GetFar()))
}

func (self Instance) SetFrustumFar(value Float.X) {
	class(self).SetFar(gd.Float(value))
}

func (self Instance) ExposureAperture() Float.X {
	return Float.X(Float.X(class(self).GetAperture()))
}

func (self Instance) SetExposureAperture(value Float.X) {
	class(self).SetAperture(gd.Float(value))
}

func (self Instance) ExposureShutterSpeed() Float.X {
	return Float.X(Float.X(class(self).GetShutterSpeed()))
}

func (self Instance) SetExposureShutterSpeed(value Float.X) {
	class(self).SetShutterSpeed(gd.Float(value))
}

func (self Instance) AutoExposureMinExposureValue() Float.X {
	return Float.X(Float.X(class(self).GetAutoExposureMinExposureValue()))
}

func (self Instance) SetAutoExposureMinExposureValue(value Float.X) {
	class(self).SetAutoExposureMinExposureValue(gd.Float(value))
}

func (self Instance) AutoExposureMaxExposureValue() Float.X {
	return Float.X(Float.X(class(self).GetAutoExposureMaxExposureValue()))
}

func (self Instance) SetAutoExposureMaxExposureValue(value Float.X) {
	class(self).SetAutoExposureMaxExposureValue(gd.Float(value))
}

//go:nosplit
func (self class) SetAperture(aperture gd.Float) { //gd:CameraAttributesPhysical.set_aperture
	var frame = callframe.New()
	callframe.Arg(frame, aperture)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_aperture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAperture() gd.Float { //gd:CameraAttributesPhysical.get_aperture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_aperture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShutterSpeed(shutter_speed gd.Float) { //gd:CameraAttributesPhysical.set_shutter_speed
	var frame = callframe.New()
	callframe.Arg(frame, shutter_speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_shutter_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShutterSpeed() gd.Float { //gd:CameraAttributesPhysical.get_shutter_speed
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_shutter_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFocalLength(focal_length gd.Float) { //gd:CameraAttributesPhysical.set_focal_length
	var frame = callframe.New()
	callframe.Arg(frame, focal_length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_focal_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFocalLength() gd.Float { //gd:CameraAttributesPhysical.get_focal_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_focal_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFocusDistance(focus_distance gd.Float) { //gd:CameraAttributesPhysical.set_focus_distance
	var frame = callframe.New()
	callframe.Arg(frame, focus_distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_focus_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFocusDistance() gd.Float { //gd:CameraAttributesPhysical.get_focus_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_focus_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNear(near gd.Float) { //gd:CameraAttributesPhysical.set_near
	var frame = callframe.New()
	callframe.Arg(frame, near)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_near, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNear() gd.Float { //gd:CameraAttributesPhysical.get_near
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_near, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFar(far gd.Float) { //gd:CameraAttributesPhysical.set_far
	var frame = callframe.New()
	callframe.Arg(frame, far)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_far, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFar() gd.Float { //gd:CameraAttributesPhysical.get_far
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_far, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the vertical field of view that corresponds to the [member frustum_focal_length]. This value is calculated internally whenever [member frustum_focal_length] is changed.
*/
//go:nosplit
func (self class) GetFov() gd.Float { //gd:CameraAttributesPhysical.get_fov
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMaxExposureValue(exposure_value_max gd.Float) { //gd:CameraAttributesPhysical.set_auto_exposure_max_exposure_value
	var frame = callframe.New()
	callframe.Arg(frame, exposure_value_max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_auto_exposure_max_exposure_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMaxExposureValue() gd.Float { //gd:CameraAttributesPhysical.get_auto_exposure_max_exposure_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_auto_exposure_max_exposure_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureMinExposureValue(exposure_value_min gd.Float) { //gd:CameraAttributesPhysical.set_auto_exposure_min_exposure_value
	var frame = callframe.New()
	callframe.Arg(frame, exposure_value_min)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_set_auto_exposure_min_exposure_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureMinExposureValue() gd.Float { //gd:CameraAttributesPhysical.get_auto_exposure_min_exposure_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributesPhysical.Bind_get_auto_exposure_min_exposure_value, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("CameraAttributesPhysical", func(ptr gd.Object) any {
		return [1]gdclass.CameraAttributesPhysical{*(*gdclass.CameraAttributesPhysical)(unsafe.Pointer(&ptr))}
	})
}
