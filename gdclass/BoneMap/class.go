package BoneMap

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
This class contains a dictionary that uses a list of bone names in [SkeletonProfile] as key names.
By assigning the actual [Skeleton3D] bone name as the key value, it maps the [Skeleton3D] to the [SkeletonProfile].

*/
type Go [1]classdb.BoneMap

/*
Returns a skeleton bone name is mapped to [param profile_bone_name].
In the retargeting process, the returned bone name is the bone name of the source skeleton.
*/
func (self Go) GetSkeletonBoneName(profile_bone_name string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSkeletonBoneName(gc, gc.StringName(profile_bone_name)).String())
}

/*
Maps a skeleton bone name to [param profile_bone_name].
In the retargeting process, the setting bone name is the bone name of the source skeleton.
*/
func (self Go) SetSkeletonBoneName(profile_bone_name string, skeleton_bone_name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkeletonBoneName(gc.StringName(profile_bone_name), gc.StringName(skeleton_bone_name))
}

/*
Returns a profile bone name having [param skeleton_bone_name]. If not found, an empty [StringName] will be returned.
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
func (self Go) FindProfileBoneName(skeleton_bone_name string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).FindProfileBoneName(gc, gc.StringName(skeleton_bone_name)).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BoneMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("BoneMap"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Profile() gdclass.SkeletonProfile {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.SkeletonProfile(class(self).GetProfile(gc))
}

func (self Go) SetProfile(value gdclass.SkeletonProfile) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProfile(value)
}

//go:nosplit
func (self class) GetProfile(ctx gd.Lifetime) gdclass.SkeletonProfile {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneMap.Bind_get_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SkeletonProfile
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProfile(profile gdclass.SkeletonProfile)  {
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
func (self Go) OnBoneMapUpdated(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("bone_map_updated"), gc.Callable(cb), 0)
}


func (self Go) OnProfileUpdated(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("profile_updated"), gc.Callable(cb), 0)
}


func (self class) AsBoneMap() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoneMap() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("BoneMap", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
