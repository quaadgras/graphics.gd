package OpenXRActionMap

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
OpenXR uses an action system similar to Godots Input map system to bind inputs and outputs on various types of XR controllers to named actions. OpenXR specifies more detail on these inputs and outputs than Godot supports.
Another important distinction is that OpenXR offers no control over these bindings. The bindings we register are suggestions, it is up to the XR runtime to offer users the ability to change these bindings. This allows the XR runtime to fill in the gaps if new hardware becomes available.
The action map therefore needs to be loaded at startup and can't be changed afterwards. This resource is a container for the entire action map.

*/
type Simple [1]classdb.OpenXRActionMap
func (self Simple) SetActionSets(action_sets gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActionSets(action_sets)
}
func (self Simple) GetActionSets() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetActionSets(gc))
}
func (self Simple) GetActionSetCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetActionSetCount()))
}
func (self Simple) FindActionSet(name string) [1]classdb.OpenXRActionSet {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OpenXRActionSet(Expert(self).FindActionSet(gc, gc.String(name)))
}
func (self Simple) GetActionSet(idx int) [1]classdb.OpenXRActionSet {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OpenXRActionSet(Expert(self).GetActionSet(gc, gd.Int(idx)))
}
func (self Simple) AddActionSet(action_set [1]classdb.OpenXRActionSet) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddActionSet(action_set)
}
func (self Simple) RemoveActionSet(action_set [1]classdb.OpenXRActionSet) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveActionSet(action_set)
}
func (self Simple) SetInteractionProfiles(interaction_profiles gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInteractionProfiles(interaction_profiles)
}
func (self Simple) GetInteractionProfiles() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetInteractionProfiles(gc))
}
func (self Simple) GetInteractionProfileCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInteractionProfileCount()))
}
func (self Simple) FindInteractionProfile(name string) [1]classdb.OpenXRInteractionProfile {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OpenXRInteractionProfile(Expert(self).FindInteractionProfile(gc, gc.String(name)))
}
func (self Simple) GetInteractionProfile(idx int) [1]classdb.OpenXRInteractionProfile {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OpenXRInteractionProfile(Expert(self).GetInteractionProfile(gc, gd.Int(idx)))
}
func (self Simple) AddInteractionProfile(interaction_profile [1]classdb.OpenXRInteractionProfile) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddInteractionProfile(interaction_profile)
}
func (self Simple) RemoveInteractionProfile(interaction_profile [1]classdb.OpenXRInteractionProfile) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveInteractionProfile(interaction_profile)
}
func (self Simple) CreateDefaultActionSets() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateDefaultActionSets()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OpenXRActionMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetActionSets(action_sets gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_sets))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_set_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetActionSets(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Retrieve the number of actions sets in our action map.
*/
//go:nosplit
func (self class) GetActionSetCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_action_set_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve an action set by name.
*/
//go:nosplit
func (self class) FindActionSet(ctx gd.Lifetime, name gd.String) object.OpenXRActionSet {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_find_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OpenXRActionSet
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Retrieve the action set at this index.
*/
//go:nosplit
func (self class) GetActionSet(ctx gd.Lifetime, idx gd.Int) object.OpenXRActionSet {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OpenXRActionSet
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Add an action set.
*/
//go:nosplit
func (self class) AddActionSet(action_set object.OpenXRActionSet)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_set[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_add_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove an action set.
*/
//go:nosplit
func (self class) RemoveActionSet(action_set object.OpenXRActionSet)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_set[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_remove_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetInteractionProfiles(interaction_profiles gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interaction_profiles))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_set_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInteractionProfiles(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Retrieve the number of interaction profiles in our action map.
*/
//go:nosplit
func (self class) GetInteractionProfileCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_interaction_profile_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Find an interaction profile by its name (path).
*/
//go:nosplit
func (self class) FindInteractionProfile(ctx gd.Lifetime, name gd.String) object.OpenXRInteractionProfile {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_find_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OpenXRInteractionProfile
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Get the interaction profile at this index.
*/
//go:nosplit
func (self class) GetInteractionProfile(ctx gd.Lifetime, idx gd.Int) object.OpenXRInteractionProfile {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OpenXRInteractionProfile
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Add an interaction profile.
*/
//go:nosplit
func (self class) AddInteractionProfile(interaction_profile object.OpenXRInteractionProfile)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interaction_profile[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_add_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove an interaction profile.
*/
//go:nosplit
func (self class) RemoveInteractionProfile(interaction_profile object.OpenXRInteractionProfile)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interaction_profile[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_remove_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Setup this action set with our default actions.
*/
//go:nosplit
func (self class) CreateDefaultActionSets()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_create_default_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsOpenXRActionMap() Expert { return self[0].AsOpenXRActionMap() }


//go:nosplit
func (self Simple) AsOpenXRActionMap() Simple { return self[0].AsOpenXRActionMap() }


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
func init() {classdb.Register("OpenXRActionMap", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
