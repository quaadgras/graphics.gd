// Package InputEventMouseButton provides methods for working with InputEventMouseButton object instances.
package InputEventMouseButton

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
import "graphics.gd/classdb/InputEventMouse"
import "graphics.gd/classdb/InputEventWithModifiers"
import "graphics.gd/classdb/InputEventFromWindow"
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

/*
Stores information about mouse click events. See [method Node._input].
[b]Note:[/b] On Wear OS devices, rotary input is mapped to [constant MOUSE_BUTTON_WHEEL_UP] and [constant MOUSE_BUTTON_WHEEL_DOWN]. This can be changed to [constant MOUSE_BUTTON_WHEEL_LEFT] and [constant MOUSE_BUTTON_WHEEL_RIGHT] with the [member ProjectSettings.input_devices/pointing/android/rotary_input_scroll_axis] setting.
*/
type Instance [1]gdclass.InputEventMouseButton

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEventMouseButton() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventMouseButton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventMouseButton"))
	casted := Instance{*(*gdclass.InputEventMouseButton)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Factor() Float.X {
	return Float.X(Float.X(class(self).GetFactor()))
}

func (self Instance) SetFactor(value Float.X) {
	class(self).SetFactor(gd.Float(value))
}

func (self Instance) ButtonIndex() MouseButton {
	return MouseButton(class(self).GetButtonIndex())
}

func (self Instance) SetButtonIndex(value MouseButton) {
	class(self).SetButtonIndex(value)
}

func (self Instance) SetCanceled(value bool) {
	class(self).SetCanceled(value)
}

func (self Instance) SetPressed(value bool) {
	class(self).SetPressed(value)
}

func (self Instance) DoubleClick() bool {
	return bool(class(self).IsDoubleClick())
}

func (self Instance) SetDoubleClick(value bool) {
	class(self).SetDoubleClick(value)
}

//go:nosplit
func (self class) SetFactor(factor gd.Float) { //gd:InputEventMouseButton.set_factor
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFactor() gd.Float { //gd:InputEventMouseButton.get_factor
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_get_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetButtonIndex(button_index MouseButton) { //gd:InputEventMouseButton.set_button_index
	var frame = callframe.New()
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_button_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetButtonIndex() MouseButton { //gd:InputEventMouseButton.get_button_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[MouseButton](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_get_button_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressed(pressed bool) { //gd:InputEventMouseButton.set_pressed
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCanceled(canceled bool) { //gd:InputEventMouseButton.set_canceled
	var frame = callframe.New()
	callframe.Arg(frame, canceled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_canceled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDoubleClick(double_click bool) { //gd:InputEventMouseButton.set_double_click
	var frame = callframe.New()
	callframe.Arg(frame, double_click)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_set_double_click, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDoubleClick() bool { //gd:InputEventMouseButton.is_double_click
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseButton.Bind_is_double_click, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEventMouse.Advanced(self.AsInputEventMouse()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEventMouse.Instance(self.AsInputEventMouse()), name)
	}
}
func init() {
	gdclass.Register("InputEventMouseButton", func(ptr gd.Object) any {
		return [1]gdclass.InputEventMouseButton{*(*gdclass.InputEventMouseButton)(unsafe.Pointer(&ptr))}
	})
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
