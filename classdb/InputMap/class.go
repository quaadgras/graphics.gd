// Package InputMap provides methods for working with InputMap object instances.
package InputMap

import "unsafe"
import "sync"
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
func HasAction(action string) bool { //gd:InputMap.has_action
	once.Do(singleton)
	return bool(class(self).HasAction(String.Name(String.New(action))))
}

/*
Returns an array of all actions in the [InputMap].
*/
func GetActions() []string { //gd:InputMap.get_actions
	once.Do(singleton)
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetActions())))
}

/*
Adds an empty action to the [InputMap] with a configurable [param deadzone].
An [InputEvent] can then be added to this action with [method action_add_event].
*/
func AddAction(action string) { //gd:InputMap.add_action
	once.Do(singleton)
	class(self).AddAction(String.Name(String.New(action)), gd.Float(0.5))
}

/*
Removes an action from the [InputMap].
*/
func EraseAction(action string) { //gd:InputMap.erase_action
	once.Do(singleton)
	class(self).EraseAction(String.Name(String.New(action)))
}

/*
Sets a deadzone value for the action.
*/
func ActionSetDeadzone(action string, deadzone Float.X) { //gd:InputMap.action_set_deadzone
	once.Do(singleton)
	class(self).ActionSetDeadzone(String.Name(String.New(action)), gd.Float(deadzone))
}

/*
Returns a deadzone value for the action.
*/
func ActionGetDeadzone(action string) Float.X { //gd:InputMap.action_get_deadzone
	once.Do(singleton)
	return Float.X(Float.X(class(self).ActionGetDeadzone(String.Name(String.New(action)))))
}

/*
Adds an [InputEvent] to an action. This [InputEvent] will trigger the action.
*/
func ActionAddEvent(action string, event [1]gdclass.InputEvent) { //gd:InputMap.action_add_event
	once.Do(singleton)
	class(self).ActionAddEvent(String.Name(String.New(action)), event)
}

/*
Returns [code]true[/code] if the action has the given [InputEvent] associated with it.
*/
func ActionHasEvent(action string, event [1]gdclass.InputEvent) bool { //gd:InputMap.action_has_event
	once.Do(singleton)
	return bool(class(self).ActionHasEvent(String.Name(String.New(action)), event))
}

/*
Removes an [InputEvent] from an action.
*/
func ActionEraseEvent(action string, event [1]gdclass.InputEvent) { //gd:InputMap.action_erase_event
	once.Do(singleton)
	class(self).ActionEraseEvent(String.Name(String.New(action)), event)
}

/*
Removes all events from an action.
*/
func ActionEraseEvents(action string) { //gd:InputMap.action_erase_events
	once.Do(singleton)
	class(self).ActionEraseEvents(String.Name(String.New(action)))
}

/*
Returns an array of [InputEvent]s associated with a given action.
[b]Note:[/b] When used in the editor (e.g. a tool script or [EditorPlugin]), this method will return events for the editor action. If you want to access your project's input binds from the editor, read the [code]input/*[/code] settings from [ProjectSettings].
*/
func ActionGetEvents(action string) [][1]gdclass.InputEvent { //gd:InputMap.action_get_events
	once.Do(singleton)
	return [][1]gdclass.InputEvent(gd.ArrayAs[[][1]gdclass.InputEvent](gd.InternalArray(class(self).ActionGetEvents(String.Name(String.New(action))))))
}

/*
Returns [code]true[/code] if the given event is part of an existing action. This method ignores keyboard modifiers if the given [InputEvent] is not pressed (for proper release detection). See [method action_has_event] if you don't want this behavior.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func EventIsAction(event [1]gdclass.InputEvent, action string) bool { //gd:InputMap.event_is_action
	once.Do(singleton)
	return bool(class(self).EventIsAction(event, String.Name(String.New(action)), false))
}

/*
Clears all [InputEventAction] in the [InputMap] and load it anew from [ProjectSettings].
*/
func LoadFromProjectSettings() { //gd:InputMap.load_from_project_settings
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
func (self class) HasAction(action String.Name) bool { //gd:InputMap.has_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
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
func (self class) GetActions() Array.Contains[String.Name] { //gd:InputMap.get_actions
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_get_actions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Name]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Adds an empty action to the [InputMap] with a configurable [param deadzone].
An [InputEvent] can then be added to this action with [method action_add_event].
*/
//go:nosplit
func (self class) AddAction(action String.Name, deadzone gd.Float) { //gd:InputMap.add_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
	callframe.Arg(frame, deadzone)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_add_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an action from the [InputMap].
*/
//go:nosplit
func (self class) EraseAction(action String.Name) { //gd:InputMap.erase_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_erase_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a deadzone value for the action.
*/
//go:nosplit
func (self class) ActionSetDeadzone(action String.Name, deadzone gd.Float) { //gd:InputMap.action_set_deadzone
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
	callframe.Arg(frame, deadzone)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_set_deadzone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a deadzone value for the action.
*/
//go:nosplit
func (self class) ActionGetDeadzone(action String.Name) gd.Float { //gd:InputMap.action_get_deadzone
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
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
func (self class) ActionAddEvent(action String.Name, event [1]gdclass.InputEvent) { //gd:InputMap.action_add_event
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_add_event, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the action has the given [InputEvent] associated with it.
*/
//go:nosplit
func (self class) ActionHasEvent(action String.Name, event [1]gdclass.InputEvent) bool { //gd:InputMap.action_has_event
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
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
func (self class) ActionEraseEvent(action String.Name, event [1]gdclass.InputEvent) { //gd:InputMap.action_erase_event
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_erase_event, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all events from an action.
*/
//go:nosplit
func (self class) ActionEraseEvents(action String.Name) { //gd:InputMap.action_erase_events
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_erase_events, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of [InputEvent]s associated with a given action.
[b]Note:[/b] When used in the editor (e.g. a tool script or [EditorPlugin]), this method will return events for the editor action. If you want to access your project's input binds from the editor, read the [code]input/*[/code] settings from [ProjectSettings].
*/
//go:nosplit
func (self class) ActionGetEvents(action String.Name) Array.Contains[[1]gdclass.InputEvent] { //gd:InputMap.action_get_events
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputMap.Bind_action_get_events, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.InputEvent]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given event is part of an existing action. This method ignores keyboard modifiers if the given [InputEvent] is not pressed (for proper release detection). See [method action_has_event] if you don't want this behavior.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) EventIsAction(event [1]gdclass.InputEvent, action String.Name, exact_match bool) bool { //gd:InputMap.event_is_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(action)))
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
func (self class) LoadFromProjectSettings() { //gd:InputMap.load_from_project_settings
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
