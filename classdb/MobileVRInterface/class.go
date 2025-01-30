// Package MobileVRInterface provides methods for working with MobileVRInterface object instances.
package MobileVRInterface

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/XRInterface"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

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
type Instance [1]gdclass.MobileVRInterface

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMobileVRInterface() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MobileVRInterface

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MobileVRInterface"))
	casted := Instance{*(*gdclass.MobileVRInterface)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) EyeHeight() Float.X {
	return Float.X(Float.X(class(self).GetEyeHeight()))
}

func (self Instance) SetEyeHeight(value Float.X) {
	class(self).SetEyeHeight(float64(value))
}

func (self Instance) Iod() Float.X {
	return Float.X(Float.X(class(self).GetIod()))
}

func (self Instance) SetIod(value Float.X) {
	class(self).SetIod(float64(value))
}

func (self Instance) DisplayWidth() Float.X {
	return Float.X(Float.X(class(self).GetDisplayWidth()))
}

func (self Instance) SetDisplayWidth(value Float.X) {
	class(self).SetDisplayWidth(float64(value))
}

func (self Instance) DisplayToLens() Float.X {
	return Float.X(Float.X(class(self).GetDisplayToLens()))
}

func (self Instance) SetDisplayToLens(value Float.X) {
	class(self).SetDisplayToLens(float64(value))
}

func (self Instance) OffsetRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetOffsetRect())
}

func (self Instance) SetOffsetRect(value Rect2.PositionSize) {
	class(self).SetOffsetRect(Rect2.PositionSize(value))
}

func (self Instance) Oversample() Float.X {
	return Float.X(Float.X(class(self).GetOversample()))
}

func (self Instance) SetOversample(value Float.X) {
	class(self).SetOversample(float64(value))
}

func (self Instance) K1() Float.X {
	return Float.X(Float.X(class(self).GetK1()))
}

func (self Instance) SetK1(value Float.X) {
	class(self).SetK1(float64(value))
}

func (self Instance) K2() Float.X {
	return Float.X(Float.X(class(self).GetK2()))
}

func (self Instance) SetK2(value Float.X) {
	class(self).SetK2(float64(value))
}

func (self Instance) VrsMinRadius() Float.X {
	return Float.X(Float.X(class(self).GetVrsMinRadius()))
}

func (self Instance) SetVrsMinRadius(value Float.X) {
	class(self).SetVrsMinRadius(float64(value))
}

func (self Instance) VrsStrength() Float.X {
	return Float.X(Float.X(class(self).GetVrsStrength()))
}

func (self Instance) SetVrsStrength(value Float.X) {
	class(self).SetVrsStrength(float64(value))
}

//go:nosplit
func (self class) SetEyeHeight(eye_height float64) { //gd:MobileVRInterface.set_eye_height
	var frame = callframe.New()
	callframe.Arg(frame, eye_height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_eye_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEyeHeight() float64 { //gd:MobileVRInterface.get_eye_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_eye_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIod(iod float64) { //gd:MobileVRInterface.set_iod
	var frame = callframe.New()
	callframe.Arg(frame, iod)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_iod, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIod() float64 { //gd:MobileVRInterface.get_iod
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_iod, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisplayWidth(display_width float64) { //gd:MobileVRInterface.set_display_width
	var frame = callframe.New()
	callframe.Arg(frame, display_width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_display_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisplayWidth() float64 { //gd:MobileVRInterface.get_display_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_display_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisplayToLens(display_to_lens float64) { //gd:MobileVRInterface.set_display_to_lens
	var frame = callframe.New()
	callframe.Arg(frame, display_to_lens)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_display_to_lens, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisplayToLens() float64 { //gd:MobileVRInterface.get_display_to_lens
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_display_to_lens, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffsetRect(offset_rect Rect2.PositionSize) { //gd:MobileVRInterface.set_offset_rect
	var frame = callframe.New()
	callframe.Arg(frame, offset_rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_offset_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffsetRect() Rect2.PositionSize { //gd:MobileVRInterface.get_offset_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[Rect2.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_offset_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOversample(oversample float64) { //gd:MobileVRInterface.set_oversample
	var frame = callframe.New()
	callframe.Arg(frame, oversample)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_oversample, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOversample() float64 { //gd:MobileVRInterface.get_oversample
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_oversample, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetK1(k float64) { //gd:MobileVRInterface.set_k1
	var frame = callframe.New()
	callframe.Arg(frame, k)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_k1, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetK1() float64 { //gd:MobileVRInterface.get_k1
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_k1, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetK2(k float64) { //gd:MobileVRInterface.set_k2
	var frame = callframe.New()
	callframe.Arg(frame, k)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_k2, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetK2() float64 { //gd:MobileVRInterface.get_k2
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_k2, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVrsMinRadius() float64 { //gd:MobileVRInterface.get_vrs_min_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsMinRadius(radius float64) { //gd:MobileVRInterface.set_vrs_min_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVrsStrength() float64 { //gd:MobileVRInterface.get_vrs_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_get_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsStrength(strength float64) { //gd:MobileVRInterface.set_vrs_strength
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MobileVRInterface.Bind_set_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(XRInterface.Advanced(self.AsXRInterface()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(XRInterface.Instance(self.AsXRInterface()), name)
	}
}
func init() {
	gdclass.Register("MobileVRInterface", func(ptr gd.Object) any {
		return [1]gdclass.MobileVRInterface{*(*gdclass.MobileVRInterface)(unsafe.Pointer(&ptr))}
	})
}
