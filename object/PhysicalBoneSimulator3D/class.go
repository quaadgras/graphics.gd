package PhysicalBoneSimulator3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/SkeletonModifier3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Node that can be the parent of [PhysicalBone3D] and can apply the simulation results to [Skeleton3D].

*/
type Simple [1]classdb.PhysicalBoneSimulator3D
func (self Simple) IsSimulatingPhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSimulatingPhysics())
}
func (self Simple) PhysicalBonesStopSimulation() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesStopSimulation()
}
func (self Simple) PhysicalBonesStartSimulation(bones gd.ArrayOf[gd.StringName]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesStartSimulation(bones)
}
func (self Simple) PhysicalBonesAddCollisionException(exception gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesAddCollisionException(exception)
}
func (self Simple) PhysicalBonesRemoveCollisionException(exception gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesRemoveCollisionException(exception)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicalBoneSimulator3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns a boolean that indicates whether the [PhysicalBoneSimulator3D] is running and simulating.
*/
//go:nosplit
func (self class) IsSimulatingPhysics() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBoneSimulator3D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
//go:nosplit
func (self class) PhysicalBonesStopSimulation()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBoneSimulator3D.Bind_physical_bones_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
//go:nosplit
func (self class) PhysicalBonesStartSimulation(bones gd.ArrayOf[gd.StringName])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bones))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBoneSimulator3D.Bind_physical_bones_start_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesAddCollisionException(exception gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBoneSimulator3D.Bind_physical_bones_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesRemoveCollisionException(exception gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBoneSimulator3D.Bind_physical_bones_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsPhysicalBoneSimulator3D() Expert { return self[0].AsPhysicalBoneSimulator3D() }


//go:nosplit
func (self Simple) AsPhysicalBoneSimulator3D() Simple { return self[0].AsPhysicalBoneSimulator3D() }


//go:nosplit
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.Expert { return self[0].AsSkeletonModifier3D() }


//go:nosplit
func (self Simple) AsSkeletonModifier3D() SkeletonModifier3D.Simple { return self[0].AsSkeletonModifier3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PhysicalBoneSimulator3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
