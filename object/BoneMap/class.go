package BoneMap

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
This class contains a dictionary that uses a list of bone names in [SkeletonProfile] as key names.
By assigning the actual [Skeleton3D] bone name as the key value, it maps the [Skeleton3D] to the [SkeletonProfile].

*/
type Simple [1]classdb.BoneMap
func (self Simple) GetProfile() [1]classdb.SkeletonProfile {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SkeletonProfile(Expert(self).GetProfile(gc))
}
func (self Simple) SetProfile(profile [1]classdb.SkeletonProfile) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProfile(profile)
}
func (self Simple) GetSkeletonBoneName(profile_bone_name string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSkeletonBoneName(gc, gc.StringName(profile_bone_name)).String())
}
func (self Simple) SetSkeletonBoneName(profile_bone_name string, skeleton_bone_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkeletonBoneName(gc.StringName(profile_bone_name), gc.StringName(skeleton_bone_name))
}
func (self Simple) FindProfileBoneName(skeleton_bone_name string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).FindProfileBoneName(gc, gc.StringName(skeleton_bone_name)).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.BoneMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) GetProfile(ctx gd.Lifetime) object.SkeletonProfile {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneMap.Bind_get_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SkeletonProfile
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProfile(profile object.SkeletonProfile)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(profile[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneMap.Bind_set_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a skeleton bone name is mapped to [param profile_bone_name].
In the retargeting process, the returned bone name is the bone name of the source skeleton.
*/
//go:nosplit
func (self class) GetSkeletonBoneName(ctx gd.Lifetime, profile_bone_name gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(profile_bone_name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneMap.Bind_get_skeleton_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Maps a skeleton bone name to [param profile_bone_name].
In the retargeting process, the setting bone name is the bone name of the source skeleton.
*/
//go:nosplit
func (self class) SetSkeletonBoneName(profile_bone_name gd.StringName, skeleton_bone_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(profile_bone_name))
	callframe.Arg(frame, mmm.Get(skeleton_bone_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneMap.Bind_set_skeleton_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a profile bone name having [param skeleton_bone_name]. If not found, an empty [StringName] will be returned.
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
//go:nosplit
func (self class) FindProfileBoneName(ctx gd.Lifetime, skeleton_bone_name gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skeleton_bone_name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneMap.Bind_find_profile_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsBoneMap() Expert { return self[0].AsBoneMap() }


//go:nosplit
func (self Simple) AsBoneMap() Simple { return self[0].AsBoneMap() }


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
func init() {classdb.Register("BoneMap", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
