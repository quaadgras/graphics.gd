// Package OpenXRActionSet provides methods for working with OpenXRActionSet object instances.
package OpenXRActionSet

import "unsafe"
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
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Action sets in OpenXR define a collection of actions that can be activated in unison. This allows games to easily change between different states that require different inputs or need to reinterpret inputs. For instance we could have an action set that is active when a menu is open, an action set that is active when the player is freely walking around and an action set that is active when the player is controlling a vehicle.
Action sets can contain the same action with the same name, if such action sets are active at the same time the action set with the highest priority defines which binding is active.
*/
type Instance [1]gdclass.OpenXRActionSet

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRActionSet() Instance
}

/*
Retrieve the number of actions in our action set.
*/
func (self Instance) GetActionCount() int { //gd:OpenXRActionSet.get_action_count
	return int(int(class(self).GetActionCount()))
}

/*
Add an action to this action set.
*/
func (self Instance) AddAction(action [1]gdclass.OpenXRAction) { //gd:OpenXRActionSet.add_action
	class(self).AddAction(action)
}

/*
Remove an action from this action set.
*/
func (self Instance) RemoveAction(action [1]gdclass.OpenXRAction) { //gd:OpenXRActionSet.remove_action
	class(self).RemoveAction(action)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRActionSet

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRActionSet"))
	casted := Instance{*(*gdclass.OpenXRActionSet)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) LocalizedName() string {
	return string(class(self).GetLocalizedName().String())
}

func (self Instance) SetLocalizedName(value string) {
	class(self).SetLocalizedName(gd.NewString(value))
}

func (self Instance) Priority() int {
	return int(int(class(self).GetPriority()))
}

func (self Instance) SetPriority(value int) {
	class(self).SetPriority(gd.Int(value))
}

func (self Instance) Actions() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetActions())))
}

func (self Instance) SetActions(value []any) {
	class(self).SetActions(gd.EngineArrayFromSlice(value))
}

//go:nosplit
func (self class) SetLocalizedName(localized_name gd.String) { //gd:OpenXRActionSet.set_localized_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(localized_name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_set_localized_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLocalizedName() gd.String { //gd:OpenXRActionSet.get_localized_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_localized_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPriority(priority gd.Int) { //gd:OpenXRActionSet.set_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPriority() gd.Int { //gd:OpenXRActionSet.get_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Retrieve the number of actions in our action set.
*/
//go:nosplit
func (self class) GetActionCount() gd.Int { //gd:OpenXRActionSet.get_action_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_action_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActions(actions Array.Any) { //gd:OpenXRActionSet.set_actions
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(actions)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_set_actions, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetActions() Array.Any { //gd:OpenXRActionSet.get_actions
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_actions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Add an action to this action set.
*/
//go:nosplit
func (self class) AddAction(action [1]gdclass.OpenXRAction) { //gd:OpenXRActionSet.add_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_add_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Remove an action from this action set.
*/
//go:nosplit
func (self class) RemoveAction(action [1]gdclass.OpenXRAction) { //gd:OpenXRActionSet.remove_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_remove_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsOpenXRActionSet() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRActionSet() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("OpenXRActionSet", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRActionSet{*(*gdclass.OpenXRActionSet)(unsafe.Pointer(&ptr))}
	})
}
