package CameraAttributesPhysical

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/CameraAttributes"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[CameraAttributesPhysical] is used to set rendering settings based on a physically-based camera's settings. It is responsible for exposure, auto-exposure, and depth of field.
When used in a [WorldEnvironment] it provides default settings for exposure, auto-exposure, and depth of field that will be used by all cameras without their own [CameraAttributes], including the editor camera. When used in a [Camera3D] it will override any [CameraAttributes] set in the [WorldEnvironment] and will override the [Camera3D]s [member Camera3D.far], [member Camera3D.near], [member Camera3D.fov], and [member Camera3D.keep_aspect] properties. When used in [VoxelGI] or [LightmapGI], only the exposure settings will be used.
The default settings are intended for use in an outdoor environment, tips for settings for use in an indoor environment can be found in each setting's documentation.
[b]Note:[/b] Depth of field blur is only supported in the Forward+ and Mobile rendering methods, not Compatibility.

*/
type Simple [1]classdb.CameraAttributesPhysical
func (self Simple) SetAperture(aperture float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAperture(gd.Float(aperture))
}
func (self Simple) GetAperture() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAperture()))
}
func (self Simple) SetShutterSpeed(shutter_speed float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShutterSpeed(gd.Float(shutter_speed))
}
func (self Simple) GetShutterSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetShutterSpeed()))
}
func (self Simple) SetFocalLength(focal_length float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFocalLength(gd.Float(focal_length))
}
func (self Simple) GetFocalLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFocalLength()))
}
func (self Simple) SetFocusDistance(focus_distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFocusDistance(gd.Float(focus_distance))
}
func (self Simple) GetFocusDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFocusDistance()))
}
func (self Simple) SetNear(near float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNear(gd.Float(near))
}
func (self Simple) GetNear() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetNear()))
}
func (self Simple) SetFar(far float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFar(gd.Float(far))
}
func (self Simple) GetFar() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFar()))
}
func (self Simple) GetFov() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFov()))
}
func (self Simple) SetAutoExposureMaxExposureValue(exposure_value_max float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoExposureMaxExposureValue(gd.Float(exposure_value_max))
}
func (self Simple) GetAutoExposureMaxExposureValue() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAutoExposureMaxExposureValue()))
}
func (self Simple) SetAutoExposureMinExposureValue(exposure_value_min float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoExposureMinExposureValue(gd.Float(exposure_value_min))
}
func (self Simple) GetAutoExposureMinExposureValue() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAutoExposureMinExposureValue()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CameraAttributesPhysical
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetAperture(aperture gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aperture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_aperture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAperture() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_aperture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShutterSpeed(shutter_speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shutter_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_shutter_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShutterSpeed() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_shutter_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFocalLength(focal_length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, focal_length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_focal_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFocalLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_focal_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFocusDistance(focus_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, focus_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_focus_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFocusDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_focus_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNear(near gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, near)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNear() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFar(far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFar() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the vertical field of view that corresponds to the [member frustum_focal_length]. This value is calculated internally whenever [member frustum_focal_length] is changed.
*/
//go:nosplit
func (self class) GetFov() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoExposureMaxExposureValue(exposure_value_max gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exposure_value_max)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_auto_exposure_max_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoExposureMaxExposureValue() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_auto_exposure_max_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoExposureMinExposureValue(exposure_value_min gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exposure_value_min)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_set_auto_exposure_min_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoExposureMinExposureValue() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributesPhysical.Bind_get_auto_exposure_min_exposure_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCameraAttributesPhysical() Expert { return self[0].AsCameraAttributesPhysical() }


//go:nosplit
func (self Simple) AsCameraAttributesPhysical() Simple { return self[0].AsCameraAttributesPhysical() }


//go:nosplit
func (self class) AsCameraAttributes() CameraAttributes.Expert { return self[0].AsCameraAttributes() }


//go:nosplit
func (self Simple) AsCameraAttributes() CameraAttributes.Simple { return self[0].AsCameraAttributes() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


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
func init() {classdb.Register("CameraAttributesPhysical", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
