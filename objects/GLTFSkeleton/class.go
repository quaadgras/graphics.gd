package GLTFSkeleton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Dictionary"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]classdb.GLTFSkeleton
type Any interface {
	gd.IsClass
	AsGLTFSkeleton() Instance
}

func (self Instance) GetGodotSkeleton() objects.Skeleton3D {
	return objects.Skeleton3D(class(self).GetGodotSkeleton())
}
func (self Instance) GetBoneAttachmentCount() int {
	return int(int(class(self).GetBoneAttachmentCount()))
}
func (self Instance) GetBoneAttachment(idx int) objects.BoneAttachment3D {
	return objects.BoneAttachment3D(class(self).GetBoneAttachment(gd.Int(idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFSkeleton

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSkeleton"))
	return Instance{*(*classdb.GLTFSkeleton)(unsafe.Pointer(&object))}
}

func (self Instance) Joints() []int32 {
	return []int32(class(self).GetJoints().AsSlice())
}

func (self Instance) SetJoints(value []int32) {
	class(self).SetJoints(gd.NewPackedInt32Slice(value))
}

func (self Instance) Roots() []int32 {
	return []int32(class(self).GetRoots().AsSlice())
}

func (self Instance) SetRoots(value []int32) {
	class(self).SetRoots(gd.NewPackedInt32Slice(value))
}

func (self Instance) UniqueNames() gd.Array {
	return gd.Array(class(self).GetUniqueNames())
}

func (self Instance) SetUniqueNames(value gd.Array) {
	class(self).SetUniqueNames(value)
}

func (self Instance) GodotBoneNode() Dictionary.Any {
	return Dictionary.Any(class(self).GetGodotBoneNode())
}

func (self Instance) SetGodotBoneNode(value Dictionary.Any) {
	class(self).SetGodotBoneNode(value)
}

//go:nosplit
func (self class) GetJoints() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJoints(joints gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joints))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoots() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoots(roots gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(roots))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGodotSkeleton() objects.Skeleton3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Skeleton3D{gd.PointerWithOwnershipTransferredToGo[classdb.Skeleton3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUniqueNames() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUniqueNames(unique_names gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(unique_names))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) GetGodotBoneNode() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) SetGodotBoneNode(godot_bone_node gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(godot_bone_node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneAttachmentCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBoneAttachment(idx gd.Int) objects.BoneAttachment3D {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.BoneAttachment3D{gd.PointerWithOwnershipTransferredToGo[classdb.BoneAttachment3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsGLTFSkeleton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFSkeleton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("GLTFSkeleton", func(ptr gd.Object) any {
		return [1]classdb.GLTFSkeleton{*(*classdb.GLTFSkeleton)(unsafe.Pointer(&ptr))}
	})
}
