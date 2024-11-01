package MobileVRInterface

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/XRInterface"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This is a generic mobile VR implementation where you need to provide details about the phone and HMD used. It does not rely on any existing framework. This is the most basic interface we have. For the best effect, you need a mobile phone with a gyroscope and accelerometer.
Note that even though there is no positional tracking, the camera will assume the headset is at a height of 1.85 meters. You can change this by setting [member eye_height].
You can initialize this interface as follows:
[codeblock]
var interface = XRServer.find_interface("Native mobile")
if interface and interface.initialize():

	get_viewport().use_xr = true

[/codeblock]
*/
type Instance [1]classdb.MobileVRInterface

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MobileVRInterface

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MobileVRInterface"))
	return Instance{classdb.MobileVRInterface(object)}
}

func (self Instance) EyeHeight() float64 {
	return float64(float64(class(self).GetEyeHeight()))
}

func (self Instance) SetEyeHeight(value float64) {
	class(self).SetEyeHeight(gd.Float(value))
}

func (self Instance) Iod() float64 {
	return float64(float64(class(self).GetIod()))
}

func (self Instance) SetIod(value float64) {
	class(self).SetIod(gd.Float(value))
}

func (self Instance) DisplayWidth() float64 {
	return float64(float64(class(self).GetDisplayWidth()))
}

func (self Instance) SetDisplayWidth(value float64) {
	class(self).SetDisplayWidth(gd.Float(value))
}

func (self Instance) DisplayToLens() float64 {
	return float64(float64(class(self).GetDisplayToLens()))
}

func (self Instance) SetDisplayToLens(value float64) {
	class(self).SetDisplayToLens(gd.Float(value))
}

func (self Instance) OffsetRect() gd.Rect2 {
	return gd.Rect2(class(self).GetOffsetRect())
}

func (self Instance) SetOffsetRect(value gd.Rect2) {
	class(self).SetOffsetRect(value)
}

func (self Instance) Oversample() float64 {
	return float64(float64(class(self).GetOversample()))
}

func (self Instance) SetOversample(value float64) {
	class(self).SetOversample(gd.Float(value))
}

func (self Instance) K1() float64 {
	return float64(float64(class(self).GetK1()))
}

func (self Instance) SetK1(value float64) {
	class(self).SetK1(gd.Float(value))
}

func (self Instance) K2() float64 {
	return float64(float64(class(self).GetK2()))
}

func (self Instance) SetK2(value float64) {
	class(self).SetK2(gd.Float(value))
}

func (self Instance) VrsMinRadius() float64 {
	return float64(float64(class(self).GetVrsMinRadius()))
}

func (self Instance) SetVrsMinRadius(value float64) {
	class(self).SetVrsMinRadius(gd.Float(value))
}

func (self Instance) VrsStrength() float64 {
	return float64(float64(class(self).GetVrsStrength()))
}

func (self Instance) SetVrsStrength(value float64) {
	class(self).SetVrsStrength(gd.Float(value))
}

//go:nosplit
func (self class) SetEyeHeight(eye_height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, eye_height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_eye_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEyeHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_eye_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIod(iod gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, iod)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_iod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIod() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_iod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisplayWidth(display_width gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, display_width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_display_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisplayWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_display_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisplayToLens(display_to_lens gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, display_to_lens)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_display_to_lens, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisplayToLens() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_display_to_lens, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffsetRect(offset_rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset_rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_offset_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffsetRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_offset_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOversample(oversample gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, oversample)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_oversample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOversample() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_oversample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetK1(k gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, k)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_k1, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetK1() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_k1, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetK2(k gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, k)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_k2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetK2() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_k2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVrsMinRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsMinRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVrsStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsStrength(strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsMobileVRInterface() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMobileVRInterface() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsXRInterface() XRInterface.Advanced {
	return *((*XRInterface.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRInterface() XRInterface.Instance {
	return *((*XRInterface.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRInterface(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRInterface(), name)
	}
}
func init() {
	classdb.Register("MobileVRInterface", func(ptr gd.Object) any { return classdb.MobileVRInterface(ptr) })
}
