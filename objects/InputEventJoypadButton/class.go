package InputEventJoypadButton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/InputEvent"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Input event type for gamepad buttons. For gamepad analog sticks and joysticks, see [InputEventJoypadMotion].
*/
type Instance [1]classdb.InputEventJoypadButton
type Any interface {
	gd.IsClass
	AsInputEventJoypadButton() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventJoypadButton

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventJoypadButton"))
	return Instance{*(*classdb.InputEventJoypadButton)(unsafe.Pointer(&object))}
}

func (self Instance) ButtonIndex() JoyButton {
	return JoyButton(class(self).GetButtonIndex())
}

func (self Instance) SetButtonIndex(value JoyButton) {
	class(self).SetButtonIndex(value)
}

func (self Instance) Pressure() Float.X {
	return Float.X(Float.X(class(self).GetPressure()))
}

func (self Instance) SetPressure(value Float.X) {
	class(self).SetPressure(gd.Float(value))
}

func (self Instance) SetPressed(value bool) {
	class(self).SetPressed(value)
}

//go:nosplit
func (self class) SetButtonIndex(button_index JoyButton) {
	var frame = callframe.New()
	callframe.Arg(frame, button_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadButton.Bind_set_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetButtonIndex() JoyButton {
	var frame = callframe.New()
	var r_ret = callframe.Ret[JoyButton](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadButton.Bind_get_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressure(pressure gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pressure)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadButton.Bind_set_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPressure() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadButton.Bind_get_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressed(pressed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventJoypadButton.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsInputEventJoypadButton() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventJoypadButton() Instance {
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
	classdb.Register("InputEventJoypadButton", func(ptr gd.Object) any {
		return [1]classdb.InputEventJoypadButton{*(*classdb.InputEventJoypadButton)(unsafe.Pointer(&ptr))}
	})
}

type JoyButton int

const (
	/*An invalid game controller button.*/
	JoyButtonInvalid JoyButton = -1
	/*Game controller SDL button A. Corresponds to the bottom action button: Sony Cross, Xbox A, Nintendo B.*/
	JoyButtonA JoyButton = 0
	/*Game controller SDL button B. Corresponds to the right action button: Sony Circle, Xbox B, Nintendo A.*/
	JoyButtonB JoyButton = 1
	/*Game controller SDL button X. Corresponds to the left action button: Sony Square, Xbox X, Nintendo Y.*/
	JoyButtonX JoyButton = 2
	/*Game controller SDL button Y. Corresponds to the top action button: Sony Triangle, Xbox Y, Nintendo X.*/
	JoyButtonY JoyButton = 3
	/*Game controller SDL back button. Corresponds to the Sony Select, Xbox Back, Nintendo - button.*/
	JoyButtonBack JoyButton = 4
	/*Game controller SDL guide button. Corresponds to the Sony PS, Xbox Home button.*/
	JoyButtonGuide JoyButton = 5
	/*Game controller SDL start button. Corresponds to the Sony Options, Xbox Menu, Nintendo + button.*/
	JoyButtonStart JoyButton = 6
	/*Game controller SDL left stick button. Corresponds to the Sony L3, Xbox L/LS button.*/
	JoyButtonLeftStick JoyButton = 7
	/*Game controller SDL right stick button. Corresponds to the Sony R3, Xbox R/RS button.*/
	JoyButtonRightStick JoyButton = 8
	/*Game controller SDL left shoulder button. Corresponds to the Sony L1, Xbox LB button.*/
	JoyButtonLeftShoulder JoyButton = 9
	/*Game controller SDL right shoulder button. Corresponds to the Sony R1, Xbox RB button.*/
	JoyButtonRightShoulder JoyButton = 10
	/*Game controller D-pad up button.*/
	JoyButtonDpadUp JoyButton = 11
	/*Game controller D-pad down button.*/
	JoyButtonDpadDown JoyButton = 12
	/*Game controller D-pad left button.*/
	JoyButtonDpadLeft JoyButton = 13
	/*Game controller D-pad right button.*/
	JoyButtonDpadRight JoyButton = 14
	/*Game controller SDL miscellaneous button. Corresponds to Xbox share button, PS5 microphone button, Nintendo Switch capture button.*/
	JoyButtonMisc1 JoyButton = 15
	/*Game controller SDL paddle 1 button.*/
	JoyButtonPaddle1 JoyButton = 16
	/*Game controller SDL paddle 2 button.*/
	JoyButtonPaddle2 JoyButton = 17
	/*Game controller SDL paddle 3 button.*/
	JoyButtonPaddle3 JoyButton = 18
	/*Game controller SDL paddle 4 button.*/
	JoyButtonPaddle4 JoyButton = 19
	/*Game controller SDL touchpad button.*/
	JoyButtonTouchpad JoyButton = 20
	/*The number of SDL game controller buttons.*/
	JoyButtonSdlMax JoyButton = 21
	/*The maximum number of game controller buttons supported by the engine. The actual limit may be lower on specific platforms:
	  - [b]Android:[/b] Up to 36 buttons.
	  - [b]Linux:[/b] Up to 80 buttons.
	  - [b]Windows[/b] and [b]macOS:[/b] Up to 128 buttons.*/
	JoyButtonMax JoyButton = 128
)
