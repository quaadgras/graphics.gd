package CameraAttributes

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Controls camera-specific attributes such as depth of field and exposure override.
When used in a [WorldEnvironment] it provides default settings for exposure, auto-exposure, and depth of field that will be used by all cameras without their own [CameraAttributes], including the editor camera. When used in a [Camera3D] it will override any [CameraAttributes] set in the [WorldEnvironment]. When used in [VoxelGI] or [LightmapGI], only the exposure settings will be used.
See also [Environment] for general 3D environment settings.
This is a pure virtual class that is inherited by [CameraAttributesPhysical] and [CameraAttributesPractical].

*/
type Go [1]classdb.CameraAttributes
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CameraAttributes
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CameraAttributes"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ExposureSensitivity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExposureSensitivity()))
}

func (self Go) SetExposureSensitivity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExposureSensitivity(gd.Float(value))
}

func (self Go) ExposureMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExposureMultiplier()))
}

func (self Go) SetExposureMultiplier(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExposureMultiplier(gd.Float(value))
}

func (self Go) AutoExposureEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoExposureEnabled())
}

func (self Go) SetAutoExposureEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoExposureEnabled(value)
}

func (self Go) AutoExposureScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAutoExposureScale()))
}

func (self Go) SetAutoExposureScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoExposureScale(gd.Float(value))
}

func (self Go) AutoExposureSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAutoExposureSpeed()))
}

func (self Go) SetAutoExposureSpeed(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoExposureSpeed(gd.Float(value))
}

//go:nosplit
func (self class) SetExposureMultiplier(multiplier gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_set_exposure_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExposureMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_get_exposure_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExposureSensitivity(sensitivity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sensitivity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_set_exposure_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExposureSensitivity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_get_exposure_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoExposureEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_set_auto_exposure_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoExposureEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_is_auto_exposure_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoExposureSpeed(exposure_speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exposure_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_set_auto_exposure_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoExposureSpeed() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_get_auto_exposure_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoExposureScale(exposure_grey gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exposure_grey)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_set_auto_exposure_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoExposureScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraAttributes.Bind_get_auto_exposure_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraAttributes() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCameraAttributes() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("CameraAttributes", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
