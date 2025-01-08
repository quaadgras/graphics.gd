// Package InputEvent provides methods for working with InputEvent object instances.
package InputEvent

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Abstract base class of all types of input events. See [method Node._input].
*/
type Instance [1]gdclass.InputEvent

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEvent() Instance
}

/*
Returns [code]true[/code] if this input event matches a pre-defined action of any type.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Instance) IsAction(action string) bool {
	return bool(class(self).IsAction(gd.NewStringName(action), false))
}

/*
Returns [code]true[/code] if the given action is being pressed (and is not an echo event for [InputEventKey] events, unless [param allow_echo] is [code]true[/code]). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
func (self Instance) IsActionPressed(action string) bool {
	return bool(class(self).IsActionPressed(gd.NewStringName(action), false, false))
}

/*
Returns [code]true[/code] if the given action is released (i.e. not pressed). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Instance) IsActionReleased(action string) bool {
	return bool(class(self).IsActionReleased(gd.NewStringName(action), false))
}

/*
Returns a value between 0.0 and 1.0 depending on the given actions' state. Useful for getting the value of events of type [InputEventJoypadMotion].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Instance) GetActionStrength(action string) Float.X {
	return Float.X(Float.X(class(self).GetActionStrength(gd.NewStringName(action), false)))
}

/*
Returns [code]true[/code] if this input event has been canceled.
*/
func (self Instance) IsCanceled() bool {
	return bool(class(self).IsCanceled())
}

/*
Returns [code]true[/code] if this input event is pressed. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
[b]Note:[/b] Due to keyboard ghosting, [method is_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
func (self Instance) IsPressed() bool {
	return bool(class(self).IsPressed())
}

/*
Returns [code]true[/code] if this input event is released. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
*/
func (self Instance) IsReleased() bool {
	return bool(class(self).IsReleased())
}

/*
Returns [code]true[/code] if this input event is an echo event (only for events of type [InputEventKey]). An echo event is a repeated key event sent when the user is holding down the key. Any other event type returns [code]false[/code].
[b]Note:[/b] The rate at which echo events are sent is typically around 20 events per second (after holding down the key for roughly half a second). However, the key repeat delay/speed can be changed by the user or disabled entirely in the operating system settings. To ensure your project works correctly on all configurations, do not assume the user has a specific key repeat configuration in your project's behavior.
*/
func (self Instance) IsEcho() bool {
	return bool(class(self).IsEcho())
}

/*
Returns a [String] representation of the event.
*/
func (self Instance) AsText() string {
	return string(class(self).AsText().String())
}

/*
Returns [code]true[/code] if the specified [param event] matches this event. Only valid for action events i.e key ([InputEventKey]), button ([InputEventMouseButton] or [InputEventJoypadButton]), axis [InputEventJoypadMotion] or action ([InputEventAction]) events.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Instance) IsMatch(event [1]gdclass.InputEvent) bool {
	return bool(class(self).IsMatch(event, true))
}

/*
Returns [code]true[/code] if this input event's type is one that can be assigned to an input action.
*/
func (self Instance) IsActionType() bool {
	return bool(class(self).IsActionType())
}

/*
Returns [code]true[/code] if the given input event and this input event can be added together (only for events of type [InputEventMouseMotion]).
The given input event's position, global position and speed will be copied. The resulting [code]relative[/code] is a sum of both events. Both events' modifiers have to be identical.
*/
func (self Instance) Accumulate(with_event [1]gdclass.InputEvent) bool {
	return bool(class(self).Accumulate(with_event))
}

/*
Returns a copy of the given input event which has been offset by [param local_ofs] and transformed by [param xform]. Relevant for events of type [InputEventMouseButton], [InputEventMouseMotion], [InputEventScreenTouch], [InputEventScreenDrag], [InputEventMagnifyGesture] and [InputEventPanGesture].
*/
func (self Instance) XformedBy(xform Transform2D.OriginXY) [1]gdclass.InputEvent {
	return [1]gdclass.InputEvent(class(self).XformedBy(gd.Transform2D(xform), gd.Vector2(gd.Vector2{0, 0})))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEvent

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEvent"))
	casted := Instance{*(*gdclass.InputEvent)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Device() int {
	return int(int(class(self).GetDevice()))
}

func (self Instance) SetDevice(value int) {
	class(self).SetDevice(gd.Int(value))
}

//go:nosplit
func (self class) SetDevice(device gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_set_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDevice() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_get_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this input event matches a pre-defined action of any type.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) IsAction(action gd.StringName, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given action is being pressed (and is not an echo event for [InputEventKey] events, unless [param allow_echo] is [code]true[/code]). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
//go:nosplit
func (self class) IsActionPressed(action gd.StringName, allow_echo bool, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, allow_echo)
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_action_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given action is released (i.e. not pressed). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) IsActionReleased(action gd.StringName, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_action_released, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a value between 0.0 and 1.0 depending on the given actions' state. Useful for getting the value of events of type [InputEventJoypadMotion].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) GetActionStrength(action gd.StringName, exact_match bool) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_get_action_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this input event has been canceled.
*/
//go:nosplit
func (self class) IsCanceled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_canceled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this input event is pressed. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
[b]Note:[/b] Due to keyboard ghosting, [method is_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
//go:nosplit
func (self class) IsPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this input event is released. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
*/
//go:nosplit
func (self class) IsReleased() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_released, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this input event is an echo event (only for events of type [InputEventKey]). An echo event is a repeated key event sent when the user is holding down the key. Any other event type returns [code]false[/code].
[b]Note:[/b] The rate at which echo events are sent is typically around 20 events per second (after holding down the key for roughly half a second). However, the key repeat delay/speed can be changed by the user or disabled entirely in the operating system settings. To ensure your project works correctly on all configurations, do not assume the user has a specific key repeat configuration in your project's behavior.
*/
//go:nosplit
func (self class) IsEcho() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_echo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [String] representation of the event.
*/
//go:nosplit
func (self class) AsText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_as_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified [param event] matches this event. Only valid for action events i.e key ([InputEventKey]), button ([InputEventMouseButton] or [InputEventJoypadButton]), axis [InputEventJoypadMotion] or action ([InputEventAction]) events.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) IsMatch(event [1]gdclass.InputEvent, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_match, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this input event's type is one that can be assigned to an input action.
*/
//go:nosplit
func (self class) IsActionType() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_is_action_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given input event and this input event can be added together (only for events of type [InputEventMouseMotion]).
The given input event's position, global position and speed will be copied. The resulting [code]relative[/code] is a sum of both events. Both events' modifiers have to be identical.
*/
//go:nosplit
func (self class) Accumulate(with_event [1]gdclass.InputEvent) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(with_event[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_accumulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a copy of the given input event which has been offset by [param local_ofs] and transformed by [param xform]. Relevant for events of type [InputEventMouseButton], [InputEventMouseMotion], [InputEventScreenTouch], [InputEventScreenDrag], [InputEventMagnifyGesture] and [InputEventPanGesture].
*/
//go:nosplit
func (self class) XformedBy(xform gd.Transform2D, local_ofs gd.Vector2) [1]gdclass.InputEvent {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	callframe.Arg(frame, local_ofs)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEvent.Bind_xformed_by, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.InputEvent{gd.PointerWithOwnershipTransferredToGo[gdclass.InputEvent](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsInputEvent() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEvent() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("InputEvent", func(ptr gd.Object) any { return [1]gdclass.InputEvent{*(*gdclass.InputEvent)(unsafe.Pointer(&ptr))} })
}
