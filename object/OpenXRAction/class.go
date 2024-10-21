package OpenXRAction

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource defines an OpenXR action. Actions can be used both for inputs (buttons, joysticks, triggers, etc.) and outputs (haptics).
OpenXR performs automatic conversion between action type and input type whenever possible. An analog trigger bound to a boolean action will thus return [code]false[/code] if the trigger is depressed and [code]true[/code] if pressed fully.
Actions are not directly bound to specific devices, instead OpenXR recognizes a limited number of top level paths that identify devices by usage. We can restrict which devices an action can be bound to by these top level paths. For instance an action that should only be used for hand held controllers can have the top level paths "/user/hand/left" and "/user/hand/right" associated with them. See the [url=https://www.khronos.org/registry/OpenXR/specs/1.0/html/xrspec.html#semantic-path-reserved]reserved path section in the OpenXR specification[/url] for more info on the top level paths.
Note that the name of the resource is used to register the action with.

*/
type Simple [1]classdb.OpenXRAction
func (self Simple) SetLocalizedName(localized_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLocalizedName(gc.String(localized_name))
}
func (self Simple) GetLocalizedName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLocalizedName(gc).String())
}
func (self Simple) SetActionType(action_type classdb.OpenXRActionActionType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActionType(action_type)
}
func (self Simple) GetActionType() classdb.OpenXRActionActionType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRActionActionType(Expert(self).GetActionType())
}
func (self Simple) SetToplevelPaths(toplevel_paths gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetToplevelPaths(toplevel_paths)
}
func (self Simple) GetToplevelPaths() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetToplevelPaths(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OpenXRAction
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsOpenXRAction() Expert { return self[0].AsOpenXRAction() }


//go:nosplit
func (self Simple) AsOpenXRAction() Simple { return self[0].AsOpenXRAction() }


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
func init() {classdb.Register("OpenXRAction", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
