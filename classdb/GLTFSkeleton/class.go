// Package GLTFSkeleton provides methods for working with GLTFSkeleton object instances.
package GLTFSkeleton

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
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
var _ Dictionary.Any
var _ RID.Any

type Instance [1]gdclass.GLTFSkeleton

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFSkeleton() Instance
}

func (self Instance) GetGodotSkeleton() [1]gdclass.Skeleton3D { //gd:GLTFSkeleton.get_godot_skeleton
	return [1]gdclass.Skeleton3D(class(self).GetGodotSkeleton())
}
func (self Instance) GetBoneAttachmentCount() int { //gd:GLTFSkeleton.get_bone_attachment_count
	return int(int(class(self).GetBoneAttachmentCount()))
}
func (self Instance) GetBoneAttachment(idx int) [1]gdclass.BoneAttachment3D { //gd:GLTFSkeleton.get_bone_attachment
	return [1]gdclass.BoneAttachment3D(class(self).GetBoneAttachment(gd.Int(idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFSkeleton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSkeleton"))
	casted := Instance{*(*gdclass.GLTFSkeleton)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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

func (self Instance) UniqueNames() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetUniqueNames())))
}

func (self Instance) SetUniqueNames(value []string) {
	class(self).SetUniqueNames(gd.ArrayFromSlice[Array.Contains[gd.String]](value))
}

func (self Instance) GodotBoneNode() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetGodotBoneNode()))
}

func (self Instance) SetGodotBoneNode(value map[any]any) {
	class(self).SetGodotBoneNode(gd.DictionaryFromMap(value))
}

//go:nosplit
func (self class) GetJoints() gd.PackedInt32Array { //gd:GLTFSkeleton.get_joints
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJoints(joints gd.PackedInt32Array) { //gd:GLTFSkeleton.set_joints
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joints))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoots() gd.PackedInt32Array { //gd:GLTFSkeleton.get_roots
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoots(roots gd.PackedInt32Array) { //gd:GLTFSkeleton.set_roots
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(roots))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGodotSkeleton() [1]gdclass.Skeleton3D { //gd:GLTFSkeleton.get_godot_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skeleton3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Skeleton3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUniqueNames() Array.Contains[gd.String] { //gd:GLTFSkeleton.get_unique_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.String]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUniqueNames(unique_names Array.Contains[gd.String]) { //gd:GLTFSkeleton.set_unique_names
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(unique_names)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) GetGodotBoneNode() Dictionary.Any { //gd:GLTFSkeleton.get_godot_bone_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) SetGodotBoneNode(godot_bone_node Dictionary.Any) { //gd:GLTFSkeleton.set_godot_bone_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(godot_bone_node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneAttachmentCount() gd.Int { //gd:GLTFSkeleton.get_bone_attachment_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBoneAttachment(idx gd.Int) [1]gdclass.BoneAttachment3D { //gd:GLTFSkeleton.get_bone_attachment
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.BoneAttachment3D{gd.PointerWithOwnershipTransferredToGo[gdclass.BoneAttachment3D](r_ret.Get())}
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
	gdclass.Register("GLTFSkeleton", func(ptr gd.Object) any {
		return [1]gdclass.GLTFSkeleton{*(*gdclass.GLTFSkeleton)(unsafe.Pointer(&ptr))}
	})
}
