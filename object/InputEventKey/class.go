package InputEventKey

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
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
An input event for keys on a keyboard. Supports key presses, key releases and [member echo] events. It can also be received in [method Node._unhandled_key_input].
[b]Note:[/b] Events received from the keyboard usually have all properties set. Event mappings should have only one of the [member keycode], [member physical_keycode] or [member unicode] set.
When events are compared, properties are checked in the following priority - [member keycode], [member physical_keycode] and [member unicode]. Events with the first matching value will be considered equal.

*/
type Simple [1]classdb.InputEventKey
func (self Simple) SetPressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPressed(pressed)
}
func (self Simple) SetKeycode(keycode gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeycode(keycode)
}
func (self Simple) GetKeycode() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetKeycode())
}
func (self Simple) SetPhysicalKeycode(physical_keycode gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicalKeycode(physical_keycode)
}
func (self Simple) GetPhysicalKeycode() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetPhysicalKeycode())
}
func (self Simple) SetKeyLabel(key_label gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeyLabel(key_label)
}
func (self Simple) GetKeyLabel() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetKeyLabel())
}
func (self Simple) SetUnicode(unicode int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUnicode(gd.Int(unicode))
}
func (self Simple) GetUnicode() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetUnicode()))
}
func (self Simple) SetLocation(location gd.KeyLocation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLocation(location)
}
func (self Simple) GetLocation() gd.KeyLocation {
	gc := gd.GarbageCollector(); _ = gc
	return gd.KeyLocation(Expert(self).GetLocation())
}
func (self Simple) SetEcho(echo bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEcho(echo)
}
func (self Simple) GetKeycodeWithModifiers() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetKeycodeWithModifiers())
}
func (self Simple) GetPhysicalKeycodeWithModifiers() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetPhysicalKeycodeWithModifiers())
}
func (self Simple) GetKeyLabelWithModifiers() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetKeyLabelWithModifiers())
}
func (self Simple) AsTextKeycode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).AsTextKeycode(gc).String())
}
func (self Simple) AsTextPhysicalKeycode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).AsTextPhysicalKeycode(gc).String())
}
func (self Simple) AsTextKeyLabel() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).AsTextKeyLabel(gc).String())
}
func (self Simple) AsTextLocation() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).AsTextLocation(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.InputEventKey
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetKeycode(keycode gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeycode() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicalKeycode(physical_keycode gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, physical_keycode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_physical_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicalKeycode() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_physical_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeyLabel(key_label gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, key_label)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_key_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeyLabel() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_key_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUnicode(unicode gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, unicode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_unicode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUnicode() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_unicode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLocation(location gd.KeyLocation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, location)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLocation() gd.KeyLocation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.KeyLocation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEcho(echo bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, echo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_echo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the Latin keycode combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_keycode_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
//go:nosplit
func (self class) GetKeycodeWithModifiers() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_keycode_with_modifiers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the physical keycode combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_physical_keycode_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
//go:nosplit
func (self class) GetPhysicalKeycodeWithModifiers() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_physical_keycode_with_modifiers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the localized key label combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_key_label_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
//go:nosplit
func (self class) GetKeyLabelWithModifiers() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_key_label_with_modifiers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member keycode] and modifiers.
*/
//go:nosplit
func (self class) AsTextKeycode(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member physical_keycode] and modifiers.
*/
//go:nosplit
func (self class) AsTextPhysicalKeycode(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_physical_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member key_label] and modifiers.
*/
//go:nosplit
func (self class) AsTextKeyLabel(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_key_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member location]. This will be a blank string if the event is not specific to a location.
*/
//go:nosplit
func (self class) AsTextLocation(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsInputEventKey() Expert { return self[0].AsInputEventKey() }


//go:nosplit
func (self Simple) AsInputEventKey() Simple { return self[0].AsInputEventKey() }


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
func init() {classdb.Register("InputEventKey", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
