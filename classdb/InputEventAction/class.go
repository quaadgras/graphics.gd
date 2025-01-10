// Package InputEventAction provides methods for working with InputEventAction object instances.
package InputEventAction

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/InputEvent"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Contains a generic action which can be targeted from several types of inputs. Actions and their events can be set in the [b]Input Map[/b] tab in [b]Project > Project Settings[/b], or with the [InputMap] class.
[b]Note:[/b] Unlike the other [InputEvent] subclasses which map to unique physical events, this virtual one is not emitted by the engine. This class is useful to emit actions manually with [method Input.parse_input_event], which are then received in [method Node._input]. To check if a physical event matches an action from the Input Map, use [method InputEvent.is_action] and [method InputEvent.is_action_pressed].
*/
type Instance [1]gdclass.InputEventAction

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEventAction() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventAction

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventAction"))
	casted := Instance{*(*gdclass.InputEventAction)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Action() string {
	return string(class(self).GetAction().String())
}

func (self Instance) SetAction(value string) {
	class(self).SetAction(gd.NewStringName(value))
}

func (self Instance) SetPressed(value bool) {
	class(self).SetPressed(value)
}

func (self Instance) Strength() Float.X {
	return Float.X(Float.X(class(self).GetStrength()))
}

func (self Instance) SetStrength(value Float.X) {
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEvent.Advanced(self.AsInputEvent()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEvent.Instance(self.AsInputEvent()), name)
	}
}
func init() {
	gdclass.Register("InputEventAction", func(ptr gd.Object) any {
		return [1]gdclass.InputEventAction{*(*gdclass.InputEventAction)(unsafe.Pointer(&ptr))}
	})
}
