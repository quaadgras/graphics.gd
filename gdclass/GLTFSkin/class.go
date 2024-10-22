package GLTFSkin

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

type Go [1]classdb.GLTFSkin
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFSkin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFSkin"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SkinRoot() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSkinRoot()))
}

func (self Go) SetSkinRoot(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkinRoot(gd.Int(value))
}

func (self Go) JointsOriginal() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetJointsOriginal(gc).AsSlice())
}

func (self Go) SetJointsOriginal(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointsOriginal(gc.PackedInt32Slice(value))
}

func (self Go) InverseBinds() gd.ArrayOf[gd.Transform3D] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.Transform3D](class(self).GetInverseBinds(gc))
}

func (self Go) SetInverseBinds(value gd.ArrayOf[gd.Transform3D]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInverseBinds(value)
}

func (self Go) Joints() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetJoints(gc).AsSlice())
}

func (self Go) SetJoints(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJoints(gc.PackedInt32Slice(value))
}

func (self Go) NonJoints() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetNonJoints(gc).AsSlice())
}

func (self Go) SetNonJoints(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNonJoints(gc.PackedInt32Slice(value))
}

func (self Go) Roots() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetRoots(gc).AsSlice())
}

func (self Go) SetRoots(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRoots(gc.PackedInt32Slice(value))
}

func (self Go) Skeleton() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSkeleton()))
}

func (self Go) SetSkeleton(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkeleton(gd.Int(value))
}

func (self Go) JointIToBoneI() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetJointIToBoneI(gc))
}

func (self Go) SetJointIToBoneI(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointIToBoneI(value)
}

func (self Go) JointIToName() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetJointIToName(gc))
}

func (self Go) SetJointIToName(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointIToName(value)
}

func (self Go) GodotSkin() gdclass.Skin {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Skin(class(self).GetGodotSkin(gc))
}

func (self Go) SetGodotSkin(value gdclass.Skin) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGodotSkin(value)
}

//go:nosplit
func (self class) GetSkinRoot() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_skin_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkinRoot(skin_root gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skin_root)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_skin_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJointsOriginal(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_joints_original, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJointsOriginal(joints_original gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(joints_original))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_joints_original, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInverseBinds(ctx gd.Lifetime) gd.ArrayOf[gd.Transform3D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_inverse_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Transform3D](ret)
}
//go:nosplit
func (self class) SetInverseBinds(inverse_binds gd.ArrayOf[gd.Transform3D])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(inverse_binds))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_inverse_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJoints(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJoints(joints gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(joints))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNonJoints(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_non_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNonJoints(non_joints gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(non_joints))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_non_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRoots(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRoots(roots gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(roots))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkeleton() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkeleton(skeleton gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJointIToBoneI(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_joint_i_to_bone_i, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJointIToBoneI(joint_i_to_bone_i gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(joint_i_to_bone_i))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_joint_i_to_bone_i, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJointIToName(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_joint_i_to_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJointIToName(joint_i_to_name gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(joint_i_to_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_joint_i_to_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGodotSkin(ctx gd.Lifetime) gdclass.Skin {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_get_godot_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Skin
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGodotSkin(godot_skin gdclass.Skin)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(godot_skin[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkin.Bind_set_godot_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFSkin() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFSkin() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("GLTFSkin", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
