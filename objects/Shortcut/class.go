package Shortcut

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Array"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Shortcuts are commonly used for interacting with a [Control] element from an [InputEvent] (also known as hotkeys).
One shortcut can contain multiple [InputEvent]'s, allowing the possibility of triggering one action with multiple different inputs.
*/
type Instance [1]classdb.Shortcut

/*
Returns whether [member events] contains an [InputEvent] which is valid.
*/
func (self Instance) HasValidEvent() bool {
	return bool(class(self).HasValidEvent())
}

/*
Returns whether any [InputEvent] in [member events] equals [param event].
*/
func (self Instance) MatchesEvent(event objects.InputEvent) bool {
	return bool(class(self).MatchesEvent(event))
}

/*
Returns the shortcut's first valid [InputEvent] as a [String].
*/
func (self Instance) GetAsText() string {
	return string(class(self).GetAsText().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Shortcut

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Shortcut"))
	return Instance{classdb.Shortcut(object)}
}

func (self Instance) Events() Array.Any {
	return Array.Any(class(self).GetEvents())
}

func (self Instance) SetEvents(value Array.Any) {
	class(self).SetEvents(value)
}

//go:nosplit
func (self class) SetEvents(events gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(events))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_set_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEvents() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_get_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether [member events] contains an [InputEvent] which is valid.
*/
//go:nosplit
func (self class) HasValidEvent() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_has_valid_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether any [InputEvent] in [member events] equals [param event].
*/
//go:nosplit
func (self class) MatchesEvent(event objects.InputEvent) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_matches_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the shortcut's first valid [InputEvent] as a [String].
*/
//go:nosplit
func (self class) GetAsText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_get_as_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsShortcut() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShortcut() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("Shortcut", func(ptr gd.Object) any { return classdb.Shortcut(ptr) }) }
