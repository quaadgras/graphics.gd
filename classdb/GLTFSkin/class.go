// Package GLTFSkin provides methods for working with GLTFSkin object instances.
package GLTFSkin

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
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"

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

type Instance [1]gdclass.GLTFSkin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFSkin() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFSkin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSkin"))
	casted := Instance{*(*gdclass.GLTFSkin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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

func (self Instance) InverseBinds() []Transform3D.BasisOrigin {
	return []Transform3D.BasisOrigin(gd.ArrayAs[[]Transform3D.BasisOrigin](gd.InternalArray(class(self).GetInverseBinds())))
}

func (self Instance) SetInverseBinds(value []Transform3D.BasisOrigin) {
	class(self).SetInverseBinds(gd.ArrayFromSlice[Array.Contains[gd.Transform3D]](value))
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

func (self Instance) JointIToBoneI() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetJointIToBoneI()))
}

func (self Instance) SetJointIToBoneI(value map[any]any) {
	class(self).SetJointIToBoneI(gd.DictionaryFromMap(value))
}

func (self Instance) JointIToName() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetJointIToName()))
}

func (self Instance) SetJointIToName(value map[any]any) {
	class(self).SetJointIToName(gd.DictionaryFromMap(value))
}

func (self Instance) GodotSkin() [1]gdclass.Skin {
	return [1]gdclass.Skin(class(self).GetGodotSkin())
}

func (self Instance) SetGodotSkin(value [1]gdclass.Skin) {
	class(self).SetGodotSkin(value)
}

//go:nosplit
func (self class) GetSkinRoot() gd.Int { //gd:GLTFSkin.get_skin_root
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_skin_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkinRoot(skin_root gd.Int) { //gd:GLTFSkin.set_skin_root
	var frame = callframe.New()
	callframe.Arg(frame, skin_root)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_skin_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointsOriginal() gd.PackedInt32Array { //gd:GLTFSkin.get_joints_original
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joints_original, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointsOriginal(joints_original gd.PackedInt32Array) { //gd:GLTFSkin.set_joints_original
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joints_original))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joints_original, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInverseBinds() Array.Contains[gd.Transform3D] { //gd:GLTFSkin.get_inverse_binds
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_inverse_binds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Transform3D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInverseBinds(inverse_binds Array.Contains[gd.Transform3D]) { //gd:GLTFSkin.set_inverse_binds
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(inverse_binds)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_inverse_binds, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJoints() gd.PackedInt32Array { //gd:GLTFSkin.get_joints
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJoints(joints gd.PackedInt32Array) { //gd:GLTFSkin.set_joints
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joints))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNonJoints() gd.PackedInt32Array { //gd:GLTFSkin.get_non_joints
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_non_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNonJoints(non_joints gd.PackedInt32Array) { //gd:GLTFSkin.set_non_joints
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(non_joints))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_non_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoots() gd.PackedInt32Array { //gd:GLTFSkin.get_roots
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoots(roots gd.PackedInt32Array) { //gd:GLTFSkin.set_roots
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(roots))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeleton() gd.Int { //gd:GLTFSkin.get_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkeleton(skeleton gd.Int) { //gd:GLTFSkin.set_skeleton
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointIToBoneI() Dictionary.Any { //gd:GLTFSkin.get_joint_i_to_bone_i
	var frame = callframe.New()
	var r_ret = callframe.Ret[Dictionary.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joint_i_to_bone_i, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointIToBoneI(joint_i_to_bone_i Dictionary.Any) { //gd:GLTFSkin.set_joint_i_to_bone_i
	var frame = callframe.New()
	callframe.Arg(frame, joint_i_to_bone_i)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joint_i_to_bone_i, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointIToName() Dictionary.Any { //gd:GLTFSkin.get_joint_i_to_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[Dictionary.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_joint_i_to_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointIToName(joint_i_to_name Dictionary.Any) { //gd:GLTFSkin.set_joint_i_to_name
	var frame = callframe.New()
	callframe.Arg(frame, joint_i_to_name)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_joint_i_to_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGodotSkin() [1]gdclass.Skin { //gd:GLTFSkin.get_godot_skin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_get_godot_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skin{gd.PointerWithOwnershipTransferredToGo[gdclass.Skin](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGodotSkin(godot_skin [1]gdclass.Skin) { //gd:GLTFSkin.set_godot_skin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(godot_skin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkin.Bind_set_godot_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("GLTFSkin", func(ptr gd.Object) any { return [1]gdclass.GLTFSkin{*(*gdclass.GLTFSkin)(unsafe.Pointer(&ptr))} })
}
