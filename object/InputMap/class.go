package InputMap

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Manages all [InputEventAction] which can be created/modified from the project settings menu [b]Project > Project Settings > Input Map[/b] or in code with [method add_action] and [method action_add_event]. See [method Node._input].

*/
type Simple [1]classdb.InputMap
func (self Simple) HasAction(action string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAction(gc.StringName(action)))
}
func (self Simple) GetActions() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](Expert(self).GetActions(gc))
}
func (self Simple) AddAction(action string, deadzone float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddAction(gc.StringName(action), gd.Float(deadzone))
}
func (self Simple) EraseAction(action string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseAction(gc.StringName(action))
}
func (self Simple) ActionSetDeadzone(action string, deadzone float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ActionSetDeadzone(gc.StringName(action), gd.Float(deadzone))
}
func (self Simple) ActionGetDeadzone(action string) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ActionGetDeadzone(gc.StringName(action))))
}
func (self Simple) ActionAddEvent(action string, event [1]classdb.InputEvent) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ActionAddEvent(gc.StringName(action), event)
}
func (self Simple) ActionHasEvent(action string, event [1]classdb.InputEvent) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ActionHasEvent(gc.StringName(action), event))
}
func (self Simple) ActionEraseEvent(action string, event [1]classdb.InputEvent) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ActionEraseEvent(gc.StringName(action), event)
}
func (self Simple) ActionEraseEvents(action string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ActionEraseEvents(gc.StringName(action))
}
func (self Simple) ActionGetEvents(action string) gd.ArrayOf[[1]classdb.InputEvent] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.InputEvent](Expert(self).ActionGetEvents(gc, gc.StringName(action)))
}
func (self Simple) EventIsAction(event [1]classdb.InputEvent, action string, exact_match bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).EventIsAction(event, gc.StringName(action), exact_match))
}
func (self Simple) LoadFromProjectSettings() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LoadFromProjectSettings()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.InputMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if the [InputMap] has a registered action with the given name.
*/
//go:nosplit
func (self class) HasAction(action gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_has_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of all actions in the [InputMap].
*/
//go:nosplit
func (self class) GetActions(ctx gd.Lifetime) gd.ArrayOf[gd.StringName] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_get_actions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.StringName](ret)
}
/*
Adds an empty action to the [InputMap] with a configurable [param deadzone].
An [InputEvent] can then be added to this action with [method action_add_event].
*/
//go:nosplit
func (self class) AddAction(action gd.StringName, deadzone gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, deadzone)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_add_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes an action from the [InputMap].
*/
//go:nosplit
func (self class) EraseAction(action gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_erase_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a deadzone value for the action.
*/
//go:nosplit
func (self class) ActionSetDeadzone(action gd.StringName, deadzone gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, deadzone)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_action_set_deadzone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a deadzone value for the action.
*/
//go:nosplit
func (self class) ActionGetDeadzone(action gd.StringName) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_action_get_deadzone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an [InputEvent] to an action. This [InputEvent] will trigger the action.
*/
//go:nosplit
func (self class) ActionAddEvent(action gd.StringName, event object.InputEvent)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_action_add_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the action has the given [InputEvent] associated with it.
*/
//go:nosplit
func (self class) ActionHasEvent(action gd.StringName, event object.InputEvent) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_action_has_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes an [InputEvent] from an action.
*/
//go:nosplit
func (self class) ActionEraseEvent(action gd.StringName, event object.InputEvent)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_action_erase_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all events from an action.
*/
//go:nosplit
func (self class) ActionEraseEvents(action gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_action_erase_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of [InputEvent]s associated with a given action.
[b]Note:[/b] When used in the editor (e.g. a tool script or [EditorPlugin]), this method will return events for the editor action. If you want to access your project's input binds from the editor, read the [code]input/*[/code] settings from [ProjectSettings].
*/
//go:nosplit
func (self class) ActionGetEvents(ctx gd.Lifetime, action gd.StringName) gd.ArrayOf[object.InputEvent] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_action_get_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.InputEvent](ret)
}
/*
Returns [code]true[/code] if the given event is part of an existing action. This method ignores keyboard modifiers if the given [InputEvent] is not pressed (for proper release detection). See [method action_has_event] if you don't want this behavior.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) EventIsAction(event object.InputEvent, action gd.StringName, exact_match bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_event_is_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears all [InputEventAction] in the [InputMap] and load it anew from [ProjectSettings].
*/
//go:nosplit
func (self class) LoadFromProjectSettings()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputMap.Bind_load_from_project_settings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsInputMap() Expert { return self[0].AsInputMap() }


//go:nosplit
func (self Simple) AsInputMap() Simple { return self[0].AsInputMap() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("InputMap", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
