package GLTFSkin

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
var _ = pointers.Root

type Instance [1]classdb.GLTFSkin
type Any interface {
	gd.IsClass
	AsGLTFSkin() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFSkin

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSkin"))
	return Instance{classdb.GLTFSkin(object)}
}

func (self Instance) SkinRoot() int {
	return int(int(class(self).GetSkinRoot()))
}

func (self Instance) SetSkinRoot(value int) {
	class(self).SetSkinRoot(gd.Int(value))
}

func (self Instance) JointsOriginal() []int32 {
	return []int32(class(self).GetJointsOriginal().AsSlice())
}

func (self Instance) SetJointsOriginal(value []int32) {
	class(self).SetJointsOriginal(gd.NewPackedInt32Slice(value))
}

func (self Instance) InverseBinds() gd.Array {
	return gd.Array(class(self).GetInverseBinds())
}

func (self Instance) SetInverseBinds(value gd.Array) {
	class(self).SetInverseBinds(value)
}

func (self Instance) Joints() []int32 {
	return []int32(class(self).GetJoints().AsSlice())
}

func (self Instance) SetJoints(value []int32) {
	class(self).SetJoints(gd.NewPackedInt32Slice(value))
}

func (self Instance) NonJoints() []int32 {
	return []int32(class(self).GetNonJoints().AsSlice())
}

func (self Instance) SetNonJoints(value []int32) {
	class(self).SetNonJoints(gd.NewPackedInt32Slice(value))
}

func (self Instance) Roots() []int32 {
	return []int32(class(self).GetRoots().AsSlice())
}

func (self Instance) SetRoots(value []int32) {
	class(self).SetRoots(gd.NewPackedInt32Slice(value))
}

func (self Instance) Skeleton() int {
	return int(int(class(self).GetSkeleton()))
}

func (self Instance) SetSkeleton(value int) {
	class(self).SetSkeleton(gd.Int(value))
}

func (self Instance) JointIToBoneI() Dictionary.Any {
	return Dictionary.Any(class(self).GetJointIToBoneI())
}

func (self Instance) SetJointIToBoneI(value Dictionary.Any) {
	class(self).SetJointIToBoneI(value)
}

func (self Instance) JointIToName() Dictionary.Any {
	return Dictionary.Any(class(self).GetJointIToName())
}

func (self Instance) SetJointIToName(value Dictionary.Any) {
	class(self).SetJointIToName(value)
}

func (self Instance) GodotSkin() objects.Skin {
	return objects.Skin(class(self).GetGodotSkin())
}

func (self Instance) SetGodotSkin(value objects.Skin) {
	class(self).SetGodotSkin(value)
}

//go:nosplit
func (self class) GetSkinRoot() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_skin_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkinRoot(skin_root gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, skin_root)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_skin_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointsOriginal() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joints_original, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointsOriginal(joints_original gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joints_original))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joints_original, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInverseBinds() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_inverse_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInverseBinds(inverse_binds gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(inverse_binds))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_inverse_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJoints() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJoints(joints gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joints))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNonJoints() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_non_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNonJoints(non_joints gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(non_joints))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_non_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoots() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoots(roots gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(roots))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeleton() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkeleton(skeleton gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointIToBoneI() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joint_i_to_bone_i, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointIToBoneI(joint_i_to_bone_i gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joint_i_to_bone_i))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joint_i_to_bone_i, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointIToName() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joint_i_to_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointIToName(joint_i_to_name gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joint_i_to_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joint_i_to_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGodotSkin() objects.Skin {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_godot_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Skin{classdb.Skin(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGodotSkin(godot_skin objects.Skin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(godot_skin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_godot_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFSkin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFSkin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() { classdb.Register("GLTFSkin", func(ptr gd.Object) any { return classdb.GLTFSkin(ptr) }) }
