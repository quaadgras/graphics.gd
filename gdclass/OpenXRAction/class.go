package OpenXRAction

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource defines an OpenXR action. Actions can be used both for inputs (buttons, joysticks, triggers, etc.) and outputs (haptics).
OpenXR performs automatic conversion between action type and input type whenever possible. An analog trigger bound to a boolean action will thus return [code]false[/code] if the trigger is depressed and [code]true[/code] if pressed fully.
Actions are not directly bound to specific devices, instead OpenXR recognizes a limited number of top level paths that identify devices by usage. We can restrict which devices an action can be bound to by these top level paths. For instance an action that should only be used for hand held controllers can have the top level paths "/user/hand/left" and "/user/hand/right" associated with them. See the [url=https://www.khronos.org/registry/OpenXR/specs/1.0/html/xrspec.html#semantic-path-reserved]reserved path section in the OpenXR specification[/url] for more info on the top level paths.
Note that the name of the resource is used to register the action with.

*/
type Go [1]classdb.OpenXRAction
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRAction
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRAction"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) LocalizedName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLocalizedName(gc).String())
}

func (self Go) SetLocalizedName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLocalizedName(gc.String(value))
}

func (self Go) ActionType() classdb.OpenXRActionActionType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.OpenXRActionActionType(class(self).GetActionType())
}

func (self Go) SetActionType(value classdb.OpenXRActionActionType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetActionType(value)
}

func (self Go) ToplevelPaths() []string {
	gc := gd.GarbageCollector(); _ = gc
		return []string(class(self).GetToplevelPaths(gc).Strings(gc))
}

func (self Go) SetToplevelPaths(value []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetToplevelPaths(gc.PackedStringSlice(value))
}

//go:nosplit
func (self class) SetLocalizedName(localized_name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(localized_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAction.Bind_set_localized_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLocalizedName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAction.Bind_get_localized_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetActionType(action_type classdb.OpenXRActionActionType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, action_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAction.Bind_set_action_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetActionType() classdb.OpenXRActionActionType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OpenXRActionActionType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAction.Bind_get_action_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetToplevelPaths(toplevel_paths gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(toplevel_paths))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAction.Bind_set_toplevel_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetToplevelPaths(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAction.Bind_get_toplevel_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsOpenXRAction() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRAction() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("OpenXRAction", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ActionType = classdb.OpenXRActionActionType

const (
/*This action provides a boolean value.*/
	OpenxrActionBool ActionType = 0
/*This action provides a float value between [code]0.0[/code] and [code]1.0[/code] for any analog input such as triggers.*/
	OpenxrActionFloat ActionType = 1
/*This action provides a [Vector2] value and can be bound to embedded trackpads and joysticks.*/
	OpenxrActionVector2 ActionType = 2
	OpenxrActionPose ActionType = 3
)
