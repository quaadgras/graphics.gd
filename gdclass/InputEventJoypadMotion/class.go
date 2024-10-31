package InputEventJoypadMotion

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Stores information about joystick motions. One [InputEventJoypadMotion] represents one axis at a time. For gamepad buttons, see [InputEventJoypadButton].
*/
type Instance [1]classdb.InputEventJoypadMotion

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventJoypadMotion

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventJoypadMotion"))
	return Instance{classdb.InputEventJoypadMotion(object)}
}

func (self Instance) Axis() gdconst.JoyAxis {
	return gdconst.JoyAxis(class(self).GetAxis())
}

func (self Instance) SetAxis(value gdconst.JoyAxis) {
	class(self).SetAxis(value)
}

func (self Instance) AxisValue() float64 {
	return float64(float64(class(self).GetAxisValue()))
}

func (self Instance) SetAxisValue(value float64) {
	class(self).SetAxisValue(gd.Float(value))
}

//go:nosplit
func (self class) SetAxis(axis gdconst.JoyAxis) {
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_set_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAxis() gdconst.JoyAxis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdconst.JoyAxis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_get_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAxisValue(axis_value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, axis_value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_set_axis_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAxisValue() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadMotion.Bind_get_axis_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEvent(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEvent(), name)
	}
}
func init() {
	classdb.Register("InputEventJoypadMotion", func(ptr gd.Object) any { return classdb.InputEventJoypadMotion(ptr) })
}
