// Package SkeletonIK3D provides methods for working with SkeletonIK3D object instances.
package SkeletonIK3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/SkeletonModifier3D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"

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
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
SkeletonIK3D is used to rotate all bones of a [Skeleton3D] bone chain a way that places the end bone at a desired 3D position. A typical scenario for IK in games is to place a character's feet on the ground or a character's hands on a currently held object. SkeletonIK uses FabrikInverseKinematic internally to solve the bone chain and applies the results to the [Skeleton3D] [code]bones_global_pose_override[/code] property for all affected bones in the chain. If fully applied, this overwrites any bone transform from [Animation]s or bone custom poses set by users. The applied amount can be controlled with the [member SkeletonModifier3D.influence] property.
[codeblock]
# Apply IK effect automatically on every new frame (not the current)
skeleton_ik_node.start()

# Apply IK effect only on the current frame
skeleton_ik_node.start(true)

# Stop IK effect and reset bones_global_pose_override on Skeleton
skeleton_ik_node.stop()

# Apply full IK effect
skeleton_ik_node.set_influence(1.0)

# Apply half IK effect
skeleton_ik_node.set_influence(0.5)

# Apply zero IK effect (a value at or below 0.01 also removes bones_global_pose_override on Skeleton)
skeleton_ik_node.set_influence(0.0)
[/codeblock]
*/
type Instance [1]gdclass.SkeletonIK3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonIK3D() Instance
}

/*
Returns the parent [Skeleton3D] Node that was present when SkeletonIK entered the [SceneTree]. Returns null if the parent node was not a [Skeleton3D] Node when SkeletonIK3D entered the [SceneTree].
*/
func (self Instance) GetParentSkeleton() [1]gdclass.Skeleton3D { //gd:SkeletonIK3D.get_parent_skeleton
	return [1]gdclass.Skeleton3D(class(self).GetParentSkeleton())
}

/*
Returns [code]true[/code] if SkeletonIK is applying IK effects on continues frames to the [Skeleton3D] bones. Returns [code]false[/code] if SkeletonIK is stopped or [method start] was used with the [code]one_time[/code] parameter set to [code]true[/code].
*/
func (self Instance) IsRunning() bool { //gd:SkeletonIK3D.is_running
	return bool(class(self).IsRunning())
}

/*
Starts applying IK effects on each frame to the [Skeleton3D] bones but will only take effect starting on the next frame. If [param one_time] is [code]true[/code], this will take effect immediately but also reset on the next frame.
*/
func (self Instance) Start() { //gd:SkeletonIK3D.start
	class(self).Start(false)
}

/*
Stops applying IK effects on each frame to the [Skeleton3D] bones and also calls [method Skeleton3D.clear_bones_global_pose_override] to remove existing overrides on all bones.
*/
func (self Instance) Stop() { //gd:SkeletonIK3D.stop
	class(self).Stop()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonIK3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonIK3D"))
	casted := Instance{*(*gdclass.SkeletonIK3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) RootBone() string {
	return string(class(self).GetRootBone().String())
}

func (self Instance) SetRootBone(value string) {
	class(self).SetRootBone(String.Name(String.New(value)))
}

func (self Instance) TipBone() string {
	return string(class(self).GetTipBone().String())
}

func (self Instance) SetTipBone(value string) {
	class(self).SetTipBone(String.Name(String.New(value)))
}

func (self Instance) Target() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetTargetTransform())
}

func (self Instance) SetTarget(value Transform3D.BasisOrigin) {
	class(self).SetTargetTransform(Transform3D.BasisOrigin(value))
}

func (self Instance) OverrideTipBasis() bool {
	return bool(class(self).IsOverrideTipBasis())
}

func (self Instance) SetOverrideTipBasis(value bool) {
	class(self).SetOverrideTipBasis(value)
}

func (self Instance) UseMagnet() bool {
	return bool(class(self).IsUsingMagnet())
}

func (self Instance) SetUseMagnet(value bool) {
	class(self).SetUseMagnet(value)
}

func (self Instance) Magnet() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetMagnetPosition())
}

func (self Instance) SetMagnet(value Vector3.XYZ) {
	class(self).SetMagnetPosition(Vector3.XYZ(value))
}

func (self Instance) TargetNode() string {
	return string(class(self).GetTargetNode().String())
}

func (self Instance) SetTargetNode(value string) {
	class(self).SetTargetNode(Path.ToNode(String.New(value)))
}

func (self Instance) MinDistance() Float.X {
	return Float.X(Float.X(class(self).GetMinDistance()))
}

func (self Instance) SetMinDistance(value Float.X) {
	class(self).SetMinDistance(float64(value))
}

func (self Instance) MaxIterations() int {
	return int(int(class(self).GetMaxIterations()))
}

func (self Instance) SetMaxIterations(value int) {
	class(self).SetMaxIterations(int64(value))
}

func (self Instance) Interpolation() Float.X {
	return Float.X(Float.X(class(self).GetInterpolation()))
}

func (self Instance) SetInterpolation(value Float.X) {
	class(self).SetInterpolation(float64(value))
}

//go:nosplit
func (self class) SetRootBone(root_bone String.Name) { //gd:SkeletonIK3D.set_root_bone
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(root_bone)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_root_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootBone() String.Name { //gd:SkeletonIK3D.get_root_bone
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_root_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTipBone(tip_bone String.Name) { //gd:SkeletonIK3D.set_tip_bone
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(tip_bone)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_tip_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTipBone() String.Name { //gd:SkeletonIK3D.get_tip_bone
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_tip_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetTransform(target Transform3D.BasisOrigin) { //gd:SkeletonIK3D.set_target_transform
	var frame = callframe.New()
	callframe.Arg(frame, target)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_target_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetTransform() Transform3D.BasisOrigin { //gd:SkeletonIK3D.get_target_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_target_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetNode(node Path.ToNode) { //gd:SkeletonIK3D.set_target_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetNode() Path.ToNode { //gd:SkeletonIK3D.get_target_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOverrideTipBasis(override bool) { //gd:SkeletonIK3D.set_override_tip_basis
	var frame = callframe.New()
	callframe.Arg(frame, override)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_override_tip_basis, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOverrideTipBasis() bool { //gd:SkeletonIK3D.is_override_tip_basis
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_is_override_tip_basis, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseMagnet(use bool) { //gd:SkeletonIK3D.set_use_magnet
	var frame = callframe.New()
	callframe.Arg(frame, use)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_use_magnet, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingMagnet() bool { //gd:SkeletonIK3D.is_using_magnet
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_is_using_magnet, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMagnetPosition(local_position Vector3.XYZ) { //gd:SkeletonIK3D.set_magnet_position
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_magnet_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMagnetPosition() Vector3.XYZ { //gd:SkeletonIK3D.get_magnet_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_magnet_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the parent [Skeleton3D] Node that was present when SkeletonIK entered the [SceneTree]. Returns null if the parent node was not a [Skeleton3D] Node when SkeletonIK3D entered the [SceneTree].
*/
//go:nosplit
func (self class) GetParentSkeleton() [1]gdclass.Skeleton3D { //gd:SkeletonIK3D.get_parent_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_parent_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skeleton3D{gd.PointerMustAssertInstanceID[gdclass.Skeleton3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if SkeletonIK is applying IK effects on continues frames to the [Skeleton3D] bones. Returns [code]false[/code] if SkeletonIK is stopped or [method start] was used with the [code]one_time[/code] parameter set to [code]true[/code].
*/
//go:nosplit
func (self class) IsRunning() bool { //gd:SkeletonIK3D.is_running
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_is_running, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinDistance(min_distance float64) { //gd:SkeletonIK3D.set_min_distance
	var frame = callframe.New()
	callframe.Arg(frame, min_distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_min_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinDistance() float64 { //gd:SkeletonIK3D.get_min_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_min_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxIterations(iterations int64) { //gd:SkeletonIK3D.set_max_iterations
	var frame = callframe.New()
	callframe.Arg(frame, iterations)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_max_iterations, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxIterations() int64 { //gd:SkeletonIK3D.get_max_iterations
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_max_iterations, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts applying IK effects on each frame to the [Skeleton3D] bones but will only take effect starting on the next frame. If [param one_time] is [code]true[/code], this will take effect immediately but also reset on the next frame.
*/
//go:nosplit
func (self class) Start(one_time bool) { //gd:SkeletonIK3D.start
	var frame = callframe.New()
	callframe.Arg(frame, one_time)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops applying IK effects on each frame to the [Skeleton3D] bones and also calls [method Skeleton3D.clear_bones_global_pose_override] to remove existing overrides on all bones.
*/
//go:nosplit
func (self class) Stop() { //gd:SkeletonIK3D.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetInterpolation(interpolation float64) { //gd:SkeletonIK3D.set_interpolation
	var frame = callframe.New()
	callframe.Arg(frame, interpolation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_set_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInterpolation() float64 { //gd:SkeletonIK3D.get_interpolation
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonIK3D.Bind_get_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonIK3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkeletonIK3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.Advanced {
	return *((*SkeletonModifier3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModifier3D() SkeletonModifier3D.Instance {
	return *((*SkeletonModifier3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModifier3D.Advanced(self.AsSkeletonModifier3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModifier3D.Instance(self.AsSkeletonModifier3D()), name)
	}
}
func init() {
	gdclass.Register("SkeletonIK3D", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonIK3D{*(*gdclass.SkeletonIK3D)(unsafe.Pointer(&ptr))}
	})
}
