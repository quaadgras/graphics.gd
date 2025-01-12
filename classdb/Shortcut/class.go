// Package Shortcut provides methods for working with Shortcut object instances.
package Shortcut

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Shortcuts are commonly used for interacting with a [Control] element from an [InputEvent] (also known as hotkeys).
One shortcut can contain multiple [InputEvent]'s, allowing the possibility of triggering one action with multiple different inputs.
*/
type Instance [1]gdclass.Shortcut

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsShortcut() Instance
}

/*
Returns whether [member events] contains an [InputEvent] which is valid.
*/
func (self Instance) HasValidEvent() bool {
	return bool(class(self).HasValidEvent())
}

/*
Returns whether any [InputEvent] in [member events] equals [param event].
*/
func (self Instance) MatchesEvent(event [1]gdclass.InputEvent) bool {
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
type class [1]gdclass.Shortcut

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Shortcut"))
	casted := Instance{*(*gdclass.Shortcut)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_set_events, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEvents() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_get_events, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_has_valid_event, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether any [InputEvent] in [member events] equals [param event].
*/
//go:nosplit
func (self class) MatchesEvent(event [1]gdclass.InputEvent) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_matches_event, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shortcut.Bind_get_as_text, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("Shortcut", func(ptr gd.Object) any { return [1]gdclass.Shortcut{*(*gdclass.Shortcut)(unsafe.Pointer(&ptr))} })
}
