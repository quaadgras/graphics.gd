// Package InputMap provides methods for working with InputMap object instances.
package InputMap

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Manages all [InputEventAction] which can be created/modified from the project settings menu [b]Project > Project Settings > Input Map[/b] or in code with [method add_action] and [method action_add_event]. See [method Node._input].
*/
var self [1]gdclass.InputMap
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.InputMap)
	self = *(*[1]gdclass.InputMap)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if the [InputMap] has a registered action with the given name.
*/
func HasAction(action string) bool {
	once.Do(singleton)
	return bool(class(self).HasAction(gd.NewStringName(action)))
}

/*
Returns an array of all actions in the [InputMap].
*/
func GetActions() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetActions())
}

/*
Adds an empty action to the [InputMap] with a configurable [param deadzone].
An [InputEvent] can then be added to this action with [method action_add_event].
*/
func AddAction(action string) {
	once.Do(singleton)
	class(self).AddAction(gd.NewStringName(action), gd.Float(0.5))
}

/*
Removes an action from the [InputMap].
*/
func EraseAction(action string) {
	once.Do(singleton)
	class(self).EraseAction(gd.NewStringName(action))
}

/*
Sets a deadzone value for the action.
*/
func ActionSetDeadzone(action string, deadzone Float.X) {
	once.Do(singleton)
	class(self).ActionSetDeadzone(gd.NewStringName(action), gd.Float(deadzone))
}

/*
Returns a deadzone value for the action.
*/
func ActionGetDeadzone(action string) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).ActionGetDeadzone(gd.NewStringName(action))))
}

/*
Adds an [InputEvent] to an action. This [InputEvent] will trigger the action.
*/
func ActionAddEvent(action string, event [1]gdclass.InputEvent) {
	once.Do(singleton)
	class(self).ActionAddEvent(gd.NewStringName(action), event)
}

/*
Returns [code]true[/code] if the action has the given [InputEvent] associated with it.
*/
func ActionHasEvent(action string, event [1]gdclass.InputEvent) bool {
	once.Do(singleton)
	return bool(class(self).ActionHasEvent(gd.NewStringName(action), event))
}

/*
Removes an [InputEvent] from an action.
*/
func ActionEraseEvent(action string, event [1]gdclass.InputEvent) {
	once.Do(singleton)
	class(self).ActionEraseEvent(gd.NewStringName(action), event)
}

/*
Removes all events from an action.
*/
func ActionEraseEvents(action string) {
	once.Do(singleton)
	class(self).ActionEraseEvents(gd.NewStringName(action))
}

/*
Returns an array of [InputEvent]s associated with a given action.
[b]Note:[/b] When used in the editor (e.g. a tool script or [EditorPlugin]), this method will return events for the editor action. If you want to access your project's input binds from the editor, read the [code]input/*[/code] settings from [ProjectSettings].
*/
func ActionGetEvents(action string) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).ActionGetEvents(gd.NewStringName(action)))
}

/*
Returns [code]true[/code] if the given event is part of an existing action. This method ignores keyboard modifiers if the given [InputEvent] is not pressed (for proper release detection). See [method action_has_event] if you don't want this behavior.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func EventIsAction(event [1]gdclass.InputEvent, action string) bool {
	once.Do(singleton)
	return bool(class(self).EventIsAction(event, gd.NewStringName(action), false))
}

/*
Clears all [InputEventAction] in the [InputMap] and load it anew from [ProjectSettings].
*/
func LoadFromProjectSettings() {
	once.Do(singleton)
	class(self).LoadFromProjectSettings()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.InputMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns [code]true[/code] if the [InputMap] has a registered action with the given name.
*/
//go:nosplit
func (self class) HasAction(action gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_has_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of all actions in the [InputMap].
*/
//go:nosplit
func (self class) GetActions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_get_actions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds an empty action to the [InputMap] with a configurable [param deadzone].
An [InputEvent] can then be added to this action with [method action_add_event].
*/
//go:nosplit
func (self class) AddAction(action gd.StringName, deadzone gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, deadzone)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_add_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an action from the [InputMap].
*/
//go:nosplit
func (self class) EraseAction(action gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_erase_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a deadzone value for the action.
*/
//go:nosplit
func (self class) ActionSetDeadzone(action gd.StringName, deadzone gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, deadzone)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_set_deadzone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a deadzone value for the action.
*/
//go:nosplit
func (self class) ActionGetDeadzone(action gd.StringName) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_get_deadzone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds an [InputEvent] to an action. This [InputEvent] will trigger the action.
*/
//go:nosplit
func (self class) ActionAddEvent(action gd.StringName, event [1]gdclass.InputEvent) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_add_event, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the action has the given [InputEvent] associated with it.
*/
//go:nosplit
func (self class) ActionHasEvent(action gd.StringName, event [1]gdclass.InputEvent) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_has_event, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes an [InputEvent] from an action.
*/
//go:nosplit
func (self class) ActionEraseEvent(action gd.StringName, event [1]gdclass.InputEvent) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_erase_event, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all events from an action.
*/
//go:nosplit
func (self class) ActionEraseEvents(action gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_erase_events, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of [InputEvent]s associated with a given action.
[b]Note:[/b] When used in the editor (e.g. a tool script or [EditorPlugin]), this method will return events for the editor action. If you want to access your project's input binds from the editor, read the [code]input/*[/code] settings from [ProjectSettings].
*/
//go:nosplit
func (self class) ActionGetEvents(action gd.StringName) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_get_events, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given event is part of an existing action. This method ignores keyboard modifiers if the given [InputEvent] is not pressed (for proper release detection). See [method action_has_event] if you don't want this behavior.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) EventIsAction(event [1]gdclass.InputEvent, action gd.StringName, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_event_is_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears all [InputEventAction] in the [InputMap] and load it anew from [ProjectSettings].
*/
//go:nosplit
func (self class) LoadFromProjectSettings() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_load_from_project_settings, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("InputMap", func(ptr gd.Object) any { return [1]gdclass.InputMap{*(*gdclass.InputMap)(unsafe.Pointer(&ptr))} })
}
