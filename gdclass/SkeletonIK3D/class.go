package SkeletonIK3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SkeletonModifier3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.SkeletonIK3D

/*
Returns the parent [Skeleton3D] Node that was present when SkeletonIK entered the [SceneTree]. Returns null if the parent node was not a [Skeleton3D] Node when SkeletonIK3D entered the [SceneTree].
*/
func (self Go) GetParentSkeleton() gdclass.Skeleton3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Skeleton3D(class(self).GetParentSkeleton(gc))
}

/*
Returns [code]true[/code] if SkeletonIK is applying IK effects on continues frames to the [Skeleton3D] bones. Returns [code]false[/code] if SkeletonIK is stopped or [method start] was used with the [code]one_time[/code] parameter set to [code]true[/code].
*/
func (self Go) IsRunning() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsRunning())
}

/*
Starts applying IK effects on each frame to the [Skeleton3D] bones but will only take effect starting on the next frame. If [param one_time] is [code]true[/code], this will take effect immediately but also reset on the next frame.
*/
func (self Go) Start() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Start(false)
}

/*
Stops applying IK effects on each frame to the [Skeleton3D] bones and also calls [method Skeleton3D.clear_bones_global_pose_override] to remove existing overrides on all bones.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonIK3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonIK3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) RootBone() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetRootBone(gc).String())
}

func (self Go) SetRootBone(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRootBone(gc.StringName(value))
}

func (self Go) TipBone() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTipBone(gc).String())
}

func (self Go) SetTipBone(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTipBone(gc.StringName(value))
}

func (self Go) Target() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform3D(class(self).GetTargetTransform())
}

func (self Go) SetTarget(value gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetTransform(value)
}

func (self Go) OverrideTipBasis() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsOverrideTipBasis())
}

func (self Go) SetOverrideTipBasis(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOverrideTipBasis(value)
}

func (self Go) UseMagnet() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingMagnet())
}

func (self Go) SetUseMagnet(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseMagnet(value)
}

func (self Go) Magnet() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetMagnetPosition())
}

func (self Go) SetMagnet(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMagnetPosition(value)
}

func (self Go) TargetNode() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTargetNode(gc).String())
}

func (self Go) SetTargetNode(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetNode(gc.String(value).NodePath(gc))
}

func (self Go) MinDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMinDistance()))
}

func (self Go) SetMinDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinDistance(gd.Float(value))
}

func (self Go) MaxIterations() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxIterations()))
}

func (self Go) SetMaxIterations(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxIterations(gd.Int(value))
}

func (self Go) Interpolation() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetInterpolation()))
}

func (self Go) SetInterpolation(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInterpolation(gd.Float(value))
}

//go:nosplit
func (self class) SetRootBone(root_bone gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(root_bone))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_root_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootBone(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_root_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTipBone(tip_bone gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tip_bone))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_tip_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTipBone(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_tip_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetTransform(target gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, target)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_target_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_target_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetNode(node gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOverrideTipBasis(override bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, override)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_override_tip_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsOverrideTipBasis() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_is_override_tip_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseMagnet(use bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_use_magnet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingMagnet() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_is_using_magnet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMagnetPosition(local_position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_magnet_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMagnetPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_magnet_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the parent [Skeleton3D] Node that was present when SkeletonIK entered the [SceneTree]. Returns null if the parent node was not a [Skeleton3D] Node when SkeletonIK3D entered the [SceneTree].
*/
//go:nosplit
func (self class) GetParentSkeleton(ctx gd.Lifetime) gdclass.Skeleton3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_parent_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Skeleton3D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if SkeletonIK is applying IK effects on continues frames to the [Skeleton3D] bones. Returns [code]false[/code] if SkeletonIK is stopped or [method start] was used with the [code]one_time[/code] parameter set to [code]true[/code].
*/
//go:nosplit
func (self class) IsRunning() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_is_running, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinDistance(min_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_min_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_min_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxIterations(iterations gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, iterations)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_max_iterations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxIterations() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_max_iterations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Starts applying IK effects on each frame to the [Skeleton3D] bones but will only take effect starting on the next frame. If [param one_time] is [code]true[/code], this will take effect immediately but also reset on the next frame.
*/
//go:nosplit
func (self class) Start(one_time bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, one_time)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops applying IK effects on each frame to the [Skeleton3D] bones and also calls [method Skeleton3D.clear_bones_global_pose_override] to remove existing overrides on all bones.
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetInterpolation(interpolation gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interpolation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_set_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInterpolation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonIK3D.Bind_get_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonIK3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonIK3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.GD { return *((*SkeletonModifier3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModifier3D() SkeletonModifier3D.Go { return *((*SkeletonModifier3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}
func init() {classdb.Register("SkeletonIK3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
