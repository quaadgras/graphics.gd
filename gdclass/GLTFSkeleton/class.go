package GLTFSkeleton

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

type Go [1]classdb.GLTFSkeleton
func (self Go) GetGodotSkeleton() gdclass.Skeleton3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Skeleton3D(class(self).GetGodotSkeleton(gc))
}
func (self Go) GetBoneAttachmentCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBoneAttachmentCount()))
}
func (self Go) GetBoneAttachment(idx int) gdclass.BoneAttachment3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.BoneAttachment3D(class(self).GetBoneAttachment(gc, gd.Int(idx)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFSkeleton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFSkeleton"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Joints() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetJoints(gc).AsSlice())
}

func (self Go) SetJoints(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJoints(gc.PackedInt32Slice(value))
}

func (self Go) Roots() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetRoots(gc).AsSlice())
}

func (self Go) SetRoots(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRoots(gc.PackedInt32Slice(value))
}

func (self Go) UniqueNames() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.String](class(self).GetUniqueNames(gc))
}

func (self Go) SetUniqueNames(value gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUniqueNames(value)
}

func (self Go) GodotBoneNode() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetGodotBoneNode(gc))
}

func (self Go) SetGodotBoneNode(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGodotBoneNode(value)
}

//go:nosplit
func (self class) GetJoints(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRoots(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGodotSkeleton(ctx gd.Lifetime) gdclass.Skeleton3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_get_godot_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Skeleton3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetUniqueNames(ctx gd.Lifetime) gd.ArrayOf[gd.String] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.String](ret)
}
//go:nosplit
func (self class) SetUniqueNames(unique_names gd.ArrayOf[gd.String])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(unique_names))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) GetGodotBoneNode(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_get_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) SetGodotBoneNode(godot_bone_node gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(godot_bone_node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_set_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneAttachmentCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_get_bone_attachment_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetBoneAttachment(ctx gd.Lifetime, idx gd.Int) gdclass.BoneAttachment3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSkeleton.Bind_get_bone_attachment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.BoneAttachment3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsGLTFSkeleton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFSkeleton() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("GLTFSkeleton", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
