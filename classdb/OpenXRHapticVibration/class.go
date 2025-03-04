// Package OpenXRHapticVibration provides methods for working with OpenXRHapticVibration object instances.
package OpenXRHapticVibration

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/OpenXRHapticBase"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
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
This haptic feedback resource makes it possible to define a vibration based haptic feedback pulse that can be triggered through actions in the OpenXR action map.
*/
type Instance [1]gdclass.OpenXRHapticVibration

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRHapticVibration() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRHapticVibration

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRHapticVibration"))
	casted := Instance{*(*gdclass.OpenXRHapticVibration)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Duration() int {
	return int(int(class(self).GetDuration()))
}

func (self Instance) SetDuration(value int) {
	class(self).SetDuration(int64(value))
}

func (self Instance) Frequency() Float.X {
	return Float.X(Float.X(class(self).GetFrequency()))
}

func (self Instance) SetFrequency(value Float.X) {
	class(self).SetFrequency(float64(value))
}

func (self Instance) Amplitude() Float.X {
	return Float.X(Float.X(class(self).GetAmplitude()))
}

func (self Instance) SetAmplitude(value Float.X) {
	class(self).SetAmplitude(float64(value))
}

//go:nosplit
func (self class) SetDuration(duration int64) { //gd:OpenXRHapticVibration.set_duration
	var frame = callframe.New()
	callframe.Arg(frame, duration)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHapticVibration.Bind_set_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDuration() int64 { //gd:OpenXRHapticVibration.get_duration
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHapticVibration.Bind_get_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrequency(frequency float64) { //gd:OpenXRHapticVibration.set_frequency
	var frame = callframe.New()
	callframe.Arg(frame, frequency)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHapticVibration.Bind_set_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrequency() float64 { //gd:OpenXRHapticVibration.get_frequency
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHapticVibration.Bind_get_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmplitude(amplitude float64) { //gd:OpenXRHapticVibration.set_amplitude
	var frame = callframe.New()
	callframe.Arg(frame, amplitude)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHapticVibration.Bind_set_amplitude, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmplitude() float64 { //gd:OpenXRHapticVibration.get_amplitude
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHapticVibration.Bind_get_amplitude, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRHapticVibration() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRHapticVibration() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsOpenXRHapticBase() OpenXRHapticBase.Advanced {
	return *((*OpenXRHapticBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRHapticBase() OpenXRHapticBase.Instance {
	return *((*OpenXRHapticBase.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(OpenXRHapticBase.Advanced(self.AsOpenXRHapticBase()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(OpenXRHapticBase.Instance(self.AsOpenXRHapticBase()), name)
	}
}
func init() {
	gdclass.Register("OpenXRHapticVibration", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRHapticVibration{*(*gdclass.OpenXRHapticVibration)(unsafe.Pointer(&ptr))}
	})
}
