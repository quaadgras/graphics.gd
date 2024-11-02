package InputEventMouseButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/InputEventMouse"
import "grow.graphics/gd/objects/InputEventWithModifiers"
import "grow.graphics/gd/objects/InputEventFromWindow"
import "grow.graphics/gd/objects/InputEvent"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Stores information about mouse click events. See [method Node._input].
[b]Note:[/b] On Wear OS devices, rotary input is mapped to [constant MOUSE_BUTTON_WHEEL_UP] and [constant MOUSE_BUTTON_WHEEL_DOWN]. This can be changed to [constant MOUSE_BUTTON_WHEEL_LEFT] and [constant MOUSE_BUTTON_WHEEL_RIGHT] with the [member ProjectSettings.input_devices/pointing/android/rotary_input_scroll_axis] setting.
*/
type Instance [1]classdb.InputEventMouseButton

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventMouseButton

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventMouseButton"))
	return Instance{classdb.InputEventMouseButton(object)}
}

func (self Instance) Factor() float64 {
	return float64(float64(class(self).GetFactor()))
}

func (self Instance) SetFactor(value float64) {
	class(self).SetFactor(gd.Float(value))
}

func (self Instance) ButtonIndex() MouseButton {
	return MouseButton(class(self).GetButtonIndex())
}

func (self Instance) SetButtonIndex(value MouseButton) {
	class(self).SetButtonIndex(value)
}

func (self Instance) DoubleClick() bool {
	return bool(class(self).IsDoubleClick())
}

func (self Instance) SetDoubleClick(value bool) {
	class(self).SetDoubleClick(value)
}

//go:nosplit
func (self class) SetFactor(factor gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFactor() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_get_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetButtonIndex(button_index MouseButton) {
	var frame = callframe.New()
	callframe.Arg(frame, button_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetButtonIndex() MouseButton {
	var frame = callframe.New()
	var r_ret = callframe.Ret[MouseButton](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_get_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressed(pressed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCanceled(canceled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, canceled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_canceled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDoubleClick(double_click bool) {
	var frame = callframe.New()
	callframe.Arg(frame, double_click)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_double_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDoubleClick() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_is_double_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMouseButton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventMouseButton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsInputEventMouse() InputEventMouse.Advanced {
	return *((*InputEventMouse.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventMouse() InputEventMouse.Instance {
	return *((*InputEventMouse.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEventWithModifiers() InputEventWithModifiers.Advanced {
	return *((*InputEventWithModifiers.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventWithModifiers() InputEventWithModifiers.Instance {
	return *((*InputEventWithModifiers.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEventFromWindow() InputEventFromWindow.Advanced {
	return *((*InputEventFromWindow.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventFromWindow() InputEventFromWindow.Instance {
	return *((*InputEventFromWindow.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsInputEventMouse(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEventMouse(), name)
	}
}
func init() {
	classdb.Register("InputEventMouseButton", func(ptr gd.Object) any { return classdb.InputEventMouseButton(ptr) })
}

type MouseButton int

const (
	/*Enum value which doesn't correspond to any mouse button. This is used to initialize [enum MouseButton] properties with a generic state.*/
	MouseButtonNone MouseButton = 0
	/*Primary mouse button, usually assigned to the left button.*/
	MouseButtonLeft MouseButton = 1
	/*Secondary mouse button, usually assigned to the right button.*/
	MouseButtonRight MouseButton = 2
	/*Middle mouse button.*/
	MouseButtonMiddle MouseButton = 3
	/*Mouse wheel scrolling up.*/
	MouseButtonWheelUp MouseButton = 4
	/*Mouse wheel scrolling down.*/
	MouseButtonWheelDown MouseButton = 5
	/*Mouse wheel left button (only present on some mice).*/
	MouseButtonWheelLeft MouseButton = 6
	/*Mouse wheel right button (only present on some mice).*/
	MouseButtonWheelRight MouseButton = 7
	/*Extra mouse button 1. This is sometimes present, usually to the sides of the mouse.*/
	MouseButtonXbutton1 MouseButton = 8
	/*Extra mouse button 2. This is sometimes present, usually to the sides of the mouse.*/
	MouseButtonXbutton2 MouseButton = 9
)
