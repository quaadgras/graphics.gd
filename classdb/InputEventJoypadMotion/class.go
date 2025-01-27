// Package InputEventJoypadMotion provides methods for working with InputEventJoypadMotion object instances.
package InputEventJoypadMotion

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
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/InputEvent"
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
var _ Path.ToNode

/*
Stores information about joystick motions. One [InputEventJoypadMotion] represents one axis at a time. For gamepad buttons, see [InputEventJoypadButton].
*/
type Instance [1]gdclass.InputEventJoypadMotion

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEventJoypadMotion() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventJoypadMotion

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventJoypadMotion"))
	casted := Instance{*(*gdclass.InputEventJoypadMotion)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Axis() JoyAxis {
	return JoyAxis(class(self).GetAxis())
}

func (self Instance) SetAxis(value JoyAxis) {
	class(self).SetAxis(value)
}

func (self Instance) AxisValue() Float.X {
	return Float.X(Float.X(class(self).GetAxisValue()))
}

func (self Instance) SetAxisValue(value Float.X) {
	class(self).SetAxisValue(gd.Float(value))
}

//go:nosplit
func (self class) SetAxis(axis JoyAxis) { //gd:InputEventJoypadMotion.set_axis
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_set_axis, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAxis() JoyAxis { //gd:InputEventJoypadMotion.get_axis
	var frame = callframe.New()
	var r_ret = callframe.Ret[JoyAxis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_get_axis, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAxisValue(axis_value gd.Float) { //gd:InputEventJoypadMotion.set_axis_value
	var frame = callframe.New()
	callframe.Arg(frame, axis_value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_set_axis_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAxisValue() gd.Float { //gd:InputEventJoypadMotion.get_axis_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_get_axis_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventJoypadMotion() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventJoypadMotion() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEvent() InputEvent.Advanced {
	return *((*InputEvent.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEvent() InputEvent.Instance {
	return *((*InputEvent.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(InputEvent.Advanced(self.AsInputEvent()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEvent.Instance(self.AsInputEvent()), name)
	}
}
func init() {
	gdclass.Register("InputEventJoypadMotion", func(ptr gd.Object) any {
		return [1]gdclass.InputEventJoypadMotion{*(*gdclass.InputEventJoypadMotion)(unsafe.Pointer(&ptr))}
	})
}

type JoyAxis int

const (
	/*An invalid game controller axis.*/
	JoyAxisInvalid JoyAxis = -1
	/*Game controller left joystick x-axis.*/
	JoyAxisLeftX JoyAxis = 0
	/*Game controller left joystick y-axis.*/
	JoyAxisLeftY JoyAxis = 1
	/*Game controller right joystick x-axis.*/
	JoyAxisRightX JoyAxis = 2
	/*Game controller right joystick y-axis.*/
	JoyAxisRightY JoyAxis = 3
	/*Game controller left trigger axis.*/
	JoyAxisTriggerLeft JoyAxis = 4
	/*Game controller right trigger axis.*/
	JoyAxisTriggerRight JoyAxis = 5
	/*The number of SDL game controller axes.*/
	JoyAxisSdlMax JoyAxis = 6
	/*The maximum number of game controller axes: OpenVR supports up to 5 Joysticks making a total of 10 axes.*/
	JoyAxisMax JoyAxis = 10
)
