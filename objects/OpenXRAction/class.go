package OpenXRAction

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This resource defines an OpenXR action. Actions can be used both for inputs (buttons, joysticks, triggers, etc.) and outputs (haptics).
OpenXR performs automatic conversion between action type and input type whenever possible. An analog trigger bound to a boolean action will thus return [code]false[/code] if the trigger is depressed and [code]true[/code] if pressed fully.
Actions are not directly bound to specific devices, instead OpenXR recognizes a limited number of top level paths that identify devices by usage. We can restrict which devices an action can be bound to by these top level paths. For instance an action that should only be used for hand held controllers can have the top level paths "/user/hand/left" and "/user/hand/right" associated with them. See the [url=https://www.khronos.org/registry/OpenXR/specs/1.0/html/xrspec.html#semantic-path-reserved]reserved path section in the OpenXR specification[/url] for more info on the top level paths.
Note that the name of the resource is used to register the action with.
*/
type Instance [1]classdb.OpenXRAction
type Any interface {
	gd.IsClass
	AsOpenXRAction() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OpenXRAction

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRAction"))
	return Instance{*(*classdb.OpenXRAction)(unsafe.Pointer(&object))}
}

func (self Instance) LocalizedName() string {
	return string(class(self).GetLocalizedName().String())
}

func (self Instance) SetLocalizedName(value string) {
	class(self).SetLocalizedName(gd.NewString(value))
}

func (self Instance) ActionType() classdb.OpenXRActionActionType {
	return classdb.OpenXRActionActionType(class(self).GetActionType())
}

func (self Instance) SetActionType(value classdb.OpenXRActionActionType) {
	class(self).SetActionType(value)
}

func (self Instance) ToplevelPaths() []string {
	return []string(class(self).GetToplevelPaths().Strings())
}

func (self Instance) SetToplevelPaths(value []string) {
	class(self).SetToplevelPaths(gd.NewPackedStringSlice(value))
}

//go:nosplit
func (self class) SetLocalizedName(localized_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(localized_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAction.Bind_set_localized_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLocalizedName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAction.Bind_get_localized_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActionType(action_type classdb.OpenXRActionActionType) {
	var frame = callframe.New()
	callframe.Arg(frame, action_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAction.Bind_set_action_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetActionType() classdb.OpenXRActionActionType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OpenXRActionActionType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAction.Bind_get_action_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetToplevelPaths(toplevel_paths gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(toplevel_paths))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAction.Bind_set_toplevel_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetToplevelPaths() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAction.Bind_get_toplevel_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsOpenXRAction() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRAction() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() {
	classdb.Register("OpenXRAction", func(ptr gd.Object) any {
		return [1]classdb.OpenXRAction{*(*classdb.OpenXRAction)(unsafe.Pointer(&ptr))}
	})
}

type ActionType = classdb.OpenXRActionActionType

const (
	/*This action provides a boolean value.*/
	OpenxrActionBool ActionType = 0
	/*This action provides a float value between [code]0.0[/code] and [code]1.0[/code] for any analog input such as triggers.*/
	OpenxrActionFloat ActionType = 1
	/*This action provides a [Vector2] value and can be bound to embedded trackpads and joysticks.*/
	OpenxrActionVector2 ActionType = 2
	OpenxrActionPose    ActionType = 3
)
