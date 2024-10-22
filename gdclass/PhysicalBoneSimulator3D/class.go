package PhysicalBoneSimulator3D

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
Node that can be the parent of [PhysicalBone3D] and can apply the simulation results to [Skeleton3D].

*/
type Go [1]classdb.PhysicalBoneSimulator3D

/*
Returns a boolean that indicates whether the [PhysicalBoneSimulator3D] is running and simulating.
*/
func (self Go) IsSimulatingPhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsSimulatingPhysics())
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
func (self Go) PhysicalBonesStopSimulation() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PhysicalBonesStopSimulation()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
func (self Go) PhysicalBonesStartSimulation() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PhysicalBonesStartSimulation(([1]gd.ArrayOf[gd.StringName]{}[0]))
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Go) PhysicalBonesAddCollisionException(exception gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PhysicalBonesAddCollisionException(exception)
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Go) PhysicalBonesRemoveCollisionException(exception gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PhysicalBonesRemoveCollisionException(exception)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicalBoneSimulator3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicalBoneSimulator3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

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
func (self class) AsPhysicalBoneSimulator3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicalBoneSimulator3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PhysicalBoneSimulator3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
