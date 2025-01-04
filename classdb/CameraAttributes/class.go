// Package CameraAttributes provides methods for working with CameraAttributes object instances.
package CameraAttributes

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Controls camera-specific attributes such as depth of field and exposure override.
When used in a [WorldEnvironment] it provides default settings for exposure, auto-exposure, and depth of field that will be used by all cameras without their own [CameraAttributes], including the editor camera. When used in a [Camera3D] it will override any [CameraAttributes] set in the [WorldEnvironment]. When used in [VoxelGI] or [LightmapGI], only the exposure settings will be used.
See also [Environment] for general 3D environment settings.
This is a pure virtual class that is inherited by [CameraAttributesPhysical] and [CameraAttributesPractical].
*/
type Instance [1]gdclass.CameraAttributes
type Any interface {
	gd.IsClass
	AsCameraAttributes() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CameraAttributes

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraAttributes"))
	return Instance{*(*gdclass.CameraAttributes)(unsafe.Pointer(&object))}
}

func (self Instance) ExposureSensitivity() Float.X {
	return Float.X(Float.X(class(self).GetExposureSensitivity()))
}

func (self Instance) SetExposureSensitivity(value Float.X) {
	class(self).SetExposureSensitivity(gd.Float(value))
}

func (self Instance) ExposureMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetExposureMultiplier()))
}

func (self Instance) SetExposureMultiplier(value Float.X) {
	class(self).SetExposureMultiplier(gd.Float(value))
}

func (self Instance) AutoExposureEnabled() bool {
	return bool(class(self).IsAutoExposureEnabled())
}

func (self Instance) SetAutoExposureEnabled(value bool) {
	class(self).SetAutoExposureEnabled(value)
}

func (self Instance) AutoExposureScale() Float.X {
	return Float.X(Float.X(class(self).GetAutoExposureScale()))
}

func (self Instance) SetAutoExposureScale(value Float.X) {
	class(self).SetAutoExposureScale(gd.Float(value))
}

func (self Instance) AutoExposureSpeed() Float.X {
	return Float.X(Float.X(class(self).GetAutoExposureSpeed()))
}

func (self Instance) SetAutoExposureSpeed(value Float.X) {
	class(self).SetAutoExposureSpeed(gd.Float(value))
}

//go:nosplit
func (self class) SetExposureMultiplier(multiplier gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_set_exposure_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExposureMultiplier() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_get_exposure_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExposureSensitivity(sensitivity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sensitivity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_set_exposure_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExposureSensitivity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_get_exposure_sensitivity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_set_auto_exposure_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoExposureEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_is_auto_exposure_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureSpeed(exposure_speed gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, exposure_speed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_set_auto_exposure_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureSpeed() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_get_auto_exposure_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoExposureScale(exposure_grey gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, exposure_grey)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_set_auto_exposure_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoExposureScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraAttributes.Bind_get_auto_exposure_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraAttributes() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCameraAttributes() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	gdclass.Register("CameraAttributes", func(ptr gd.Object) any {
		return [1]gdclass.CameraAttributes{*(*gdclass.CameraAttributes)(unsafe.Pointer(&ptr))}
	})
}
