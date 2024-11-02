package InputEventAction

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/InputEvent"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Contains a generic action which can be targeted from several types of inputs. Actions and their events can be set in the [b]Input Map[/b] tab in [b]Project > Project Settings[/b], or with the [InputMap] class.
[b]Note:[/b] Unlike the other [InputEvent] subclasses which map to unique physical events, this virtual one is not emitted by the engine. This class is useful to emit actions manually with [method Input.parse_input_event], which are then received in [method Node._input]. To check if a physical event matches an action from the Input Map, use [method InputEvent.is_action] and [method InputEvent.is_action_pressed].
*/
type Instance [1]classdb.InputEventAction

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventAction

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventAction"))
	return Instance{classdb.InputEventAction(object)}
}

func (self Instance) Action() string {
	return string(class(self).GetAction().String())
}

func (self Instance) SetAction(value string) {
	class(self).SetAction(gd.NewStringName(value))
}

func (self Instance) Strength() float64 {
	return float64(float64(class(self).GetStrength()))
}

func (self Instance) SetStrength(value float64) {
	class(self).SetStrength(gd.Float(value))
}

func (self Instance) EventIndex() int {
	return int(int(class(self).GetEventIndex()))
}

func (self Instance) SetEventIndex(value int) {
	class(self).SetEventIndex(gd.Int(value))
}

//go:nosplit
func (self class) SetAction(action gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventAction.Bind_set_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAction() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventAction.Bind_get_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressed(pressed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventAction.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetStrength(strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventAction.Bind_set_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventAction.Bind_get_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEventIndex(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventAction.Bind_set_event_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEventIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventAction.Bind_get_event_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventAction() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventAction() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("InputEventAction", func(ptr gd.Object) any { return classdb.InputEventAction(ptr) })
}
