package BoneMap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This class contains a dictionary that uses a list of bone names in [SkeletonProfile] as key names.
By assigning the actual [Skeleton3D] bone name as the key value, it maps the [Skeleton3D] to the [SkeletonProfile].
*/
type Instance [1]classdb.BoneMap

/*
Returns a skeleton bone name is mapped to [param profile_bone_name].
In the retargeting process, the returned bone name is the bone name of the source skeleton.
*/
func (self Instance) GetSkeletonBoneName(profile_bone_name string) string {
	return string(class(self).GetSkeletonBoneName(gd.NewStringName(profile_bone_name)).String())
}

/*
Maps a skeleton bone name to [param profile_bone_name].
In the retargeting process, the setting bone name is the bone name of the source skeleton.
*/
func (self Instance) SetSkeletonBoneName(profile_bone_name string, skeleton_bone_name string) {
	class(self).SetSkeletonBoneName(gd.NewStringName(profile_bone_name), gd.NewStringName(skeleton_bone_name))
}

/*
Returns a profile bone name having [param skeleton_bone_name]. If not found, an empty [StringName] will be returned.
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
func (self Instance) FindProfileBoneName(skeleton_bone_name string) string {
	return string(class(self).FindProfileBoneName(gd.NewStringName(skeleton_bone_name)).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.BoneMap

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoneMap"))
	return Instance{classdb.BoneMap(object)}
}

func (self Instance) Profile() objects.SkeletonProfile {
	return objects.SkeletonProfile(class(self).GetProfile())
}

func (self Instance) SetProfile(value objects.SkeletonProfile) {
	class(self).SetProfile(value)
}

//go:nosplit
func (self class) GetProfile() objects.SkeletonProfile {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_get_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.SkeletonProfile{classdb.SkeletonProfile(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProfile(profile objects.SkeletonProfile) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_set_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a skeleton bone name is mapped to [param profile_bone_name].
In the retargeting process, the returned bone name is the bone name of the source skeleton.
*/
//go:nosplit
func (self class) GetSkeletonBoneName(profile_bone_name gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile_bone_name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_get_skeleton_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Maps a skeleton bone name to [param profile_bone_name].
In the retargeting process, the setting bone name is the bone name of the source skeleton.
*/
//go:nosplit
func (self class) SetSkeletonBoneName(profile_bone_name gd.StringName, skeleton_bone_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile_bone_name))
	callframe.Arg(frame, pointers.Get(skeleton_bone_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_set_skeleton_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a profile bone name having [param skeleton_bone_name]. If not found, an empty [StringName] will be returned.
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
//go:nosplit
func (self class) FindProfileBoneName(skeleton_bone_name gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skeleton_bone_name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_find_profile_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnBoneMapUpdated(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("bone_map_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProfileUpdated(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("profile_updated"), gd.NewCallable(cb), 0)
}

func (self class) AsBoneMap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBoneMap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() { classdb.Register("BoneMap", func(ptr gd.Object) any { return classdb.BoneMap(ptr) }) }
