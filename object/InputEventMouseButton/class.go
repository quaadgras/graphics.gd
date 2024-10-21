package InputEventMouseButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/InputEventMouse"
import "grow.graphics/gd/object/InputEventWithModifiers"
import "grow.graphics/gd/object/InputEventFromWindow"
import "grow.graphics/gd/object/InputEvent"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Stores information about mouse click events. See [method Node._input].
[b]Note:[/b] On Wear OS devices, rotary input is mapped to [constant MOUSE_BUTTON_WHEEL_UP] and [constant MOUSE_BUTTON_WHEEL_DOWN]. This can be changed to [constant MOUSE_BUTTON_WHEEL_LEFT] and [constant MOUSE_BUTTON_WHEEL_RIGHT] with the [member ProjectSettings.input_devices/pointing/android/rotary_input_scroll_axis] setting.

*/
type Simple [1]classdb.InputEventMouseButton
func (self Simple) SetFactor(factor float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFactor(gd.Float(factor))
}
func (self Simple) GetFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFactor()))
}
func (self Simple) SetButtonIndex(button_index gd.MouseButton) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetButtonIndex(button_index)
}
func (self Simple) GetButtonIndex() gd.MouseButton {
	gc := gd.GarbageCollector(); _ = gc
	return gd.MouseButton(Expert(self).GetButtonIndex())
}
func (self Simple) SetPressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPressed(pressed)
}
func (self Simple) SetCanceled(canceled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanceled(canceled)
}
func (self Simple) SetDoubleClick(double_click bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDoubleClick(double_click)
}
func (self Simple) IsDoubleClick() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDoubleClick())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.InputEventMouseButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetFactor(factor gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_get_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetButtonIndex(button_index gd.MouseButton)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, button_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonIndex() gd.MouseButton {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.MouseButton](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_get_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCanceled(canceled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canceled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_canceled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDoubleClick(double_click bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, double_click)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_double_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDoubleClick() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_is_double_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsInputEventMouseButton() Expert { return self[0].AsInputEventMouseButton() }


//go:nosplit
func (self Simple) AsInputEventMouseButton() Simple { return self[0].AsInputEventMouseButton() }


//go:nosplit
func (self class) AsInputEventMouse() InputEventMouse.Expert { return self[0].AsInputEventMouse() }


//go:nosplit
func (self Simple) AsInputEventMouse() InputEventMouse.Simple { return self[0].AsInputEventMouse() }


//go:nosplit
func (self class) AsInputEventWithModifiers() InputEventWithModifiers.Expert { return self[0].AsInputEventWithModifiers() }


//go:nosplit
func (self Simple) AsInputEventWithModifiers() InputEventWithModifiers.Simple { return self[0].AsInputEventWithModifiers() }


//go:nosplit
func (self class) AsInputEventFromWindow() InputEventFromWindow.Expert { return self[0].AsInputEventFromWindow() }


//go:nosplit
func (self Simple) AsInputEventFromWindow() InputEventFromWindow.Simple { return self[0].AsInputEventFromWindow() }


//go:nosplit
func (self class) AsInputEvent() InputEvent.Expert { return self[0].AsInputEvent() }


//go:nosplit
func (self Simple) AsInputEvent() InputEvent.Simple { return self[0].AsInputEvent() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("InputEventMouseButton", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
