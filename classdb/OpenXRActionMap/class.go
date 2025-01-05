// Package OpenXRActionMap provides methods for working with OpenXRActionMap object instances.
package OpenXRActionMap

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
OpenXR uses an action system similar to Godots Input map system to bind inputs and outputs on various types of XR controllers to named actions. OpenXR specifies more detail on these inputs and outputs than Godot supports.
Another important distinction is that OpenXR offers no control over these bindings. The bindings we register are suggestions, it is up to the XR runtime to offer users the ability to change these bindings. This allows the XR runtime to fill in the gaps if new hardware becomes available.
The action map therefore needs to be loaded at startup and can't be changed afterwards. This resource is a container for the entire action map.
*/
type Instance [1]gdclass.OpenXRActionMap
type Any interface {
	gd.IsClass
	AsOpenXRActionMap() Instance
}

/*
Retrieve the number of actions sets in our action map.
*/
func (self Instance) GetActionSetCount() int {
	return int(int(class(self).GetActionSetCount()))
}

/*
Retrieve an action set by name.
*/
func (self Instance) FindActionSet(name string) [1]gdclass.OpenXRActionSet {
	return [1]gdclass.OpenXRActionSet(class(self).FindActionSet(gd.NewString(name)))
}

/*
Retrieve the action set at this index.
*/
func (self Instance) GetActionSet(idx int) [1]gdclass.OpenXRActionSet {
	return [1]gdclass.OpenXRActionSet(class(self).GetActionSet(gd.Int(idx)))
}

/*
Add an action set.
*/
func (self Instance) AddActionSet(action_set [1]gdclass.OpenXRActionSet) {
	class(self).AddActionSet(action_set)
}

/*
Remove an action set.
*/
func (self Instance) RemoveActionSet(action_set [1]gdclass.OpenXRActionSet) {
	class(self).RemoveActionSet(action_set)
}

/*
Retrieve the number of interaction profiles in our action map.
*/
func (self Instance) GetInteractionProfileCount() int {
	return int(int(class(self).GetInteractionProfileCount()))
}

/*
Find an interaction profile by its name (path).
*/
func (self Instance) FindInteractionProfile(name string) [1]gdclass.OpenXRInteractionProfile {
	return [1]gdclass.OpenXRInteractionProfile(class(self).FindInteractionProfile(gd.NewString(name)))
}

/*
Get the interaction profile at this index.
*/
func (self Instance) GetInteractionProfile(idx int) [1]gdclass.OpenXRInteractionProfile {
	return [1]gdclass.OpenXRInteractionProfile(class(self).GetInteractionProfile(gd.Int(idx)))
}

/*
Add an interaction profile.
*/
func (self Instance) AddInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) {
	class(self).AddInteractionProfile(interaction_profile)
}

/*
Remove an interaction profile.
*/
func (self Instance) RemoveInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) {
	class(self).RemoveInteractionProfile(interaction_profile)
}

/*
Setup this action set with our default actions.
*/
func (self Instance) CreateDefaultActionSets() {
	class(self).CreateDefaultActionSets()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRActionMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRActionMap"))
	return Instance{*(*gdclass.OpenXRActionMap)(unsafe.Pointer(&object))}
}

func (self Instance) ActionSets() Array.Any {
	return Array.Any(class(self).GetActionSets())
}

func (self Instance) SetActionSets(value Array.Any) {
	class(self).SetActionSets(value)
}

func (self Instance) InteractionProfiles() Array.Any {
	return Array.Any(class(self).GetInteractionProfiles())
}

func (self Instance) SetInteractionProfiles(value Array.Any) {
	class(self).SetInteractionProfiles(value)
}

//go:nosplit
func (self class) SetActionSets(action_sets gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action_sets))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_set_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetActionSets() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Retrieve the number of actions sets in our action map.
*/
//go:nosplit
func (self class) GetActionSetCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_action_set_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Retrieve an action set by name.
*/
//go:nosplit
func (self class) FindActionSet(name gd.String) [1]gdclass.OpenXRActionSet {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_find_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.OpenXRActionSet{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRActionSet](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Retrieve the action set at this index.
*/
//go:nosplit
func (self class) GetActionSet(idx gd.Int) [1]gdclass.OpenXRActionSet {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.OpenXRActionSet{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRActionSet](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Add an action set.
*/
//go:nosplit
func (self class) AddActionSet(action_set [1]gdclass.OpenXRActionSet) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action_set[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_add_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Remove an action set.
*/
//go:nosplit
func (self class) RemoveActionSet(action_set [1]gdclass.OpenXRActionSet) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action_set[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_remove_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetInteractionProfiles(interaction_profiles gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interaction_profiles))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_set_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInteractionProfiles() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Retrieve the number of interaction profiles in our action map.
*/
//go:nosplit
func (self class) GetInteractionProfileCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_interaction_profile_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Find an interaction profile by its name (path).
*/
//go:nosplit
func (self class) FindInteractionProfile(name gd.String) [1]gdclass.OpenXRInteractionProfile {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_find_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.OpenXRInteractionProfile{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRInteractionProfile](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Get the interaction profile at this index.
*/
//go:nosplit
func (self class) GetInteractionProfile(idx gd.Int) [1]gdclass.OpenXRInteractionProfile {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.OpenXRInteractionProfile{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRInteractionProfile](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Add an interaction profile.
*/
//go:nosplit
func (self class) AddInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interaction_profile[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_add_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Remove an interaction profile.
*/
//go:nosplit
func (self class) RemoveInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interaction_profile[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_remove_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Setup this action set with our default actions.
*/
//go:nosplit
func (self class) CreateDefaultActionSets() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_create_default_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRActionMap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRActionMap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("OpenXRActionMap", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRActionMap{*(*gdclass.OpenXRActionMap)(unsafe.Pointer(&ptr))}
	})
}
