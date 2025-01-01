package PhysicalBoneSimulator3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/SkeletonModifier3D"
import "graphics.gd/objects/Node3D"
import "graphics.gd/objects/Node"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Node that can be the parent of [PhysicalBone3D] and can apply the simulation results to [Skeleton3D].
*/
type Instance [1]classdb.PhysicalBoneSimulator3D
type Any interface {
	gd.IsClass
	AsPhysicalBoneSimulator3D() Instance
}

/*
Returns a boolean that indicates whether the [PhysicalBoneSimulator3D] is running and simulating.
*/
func (self Instance) IsSimulatingPhysics() bool {
	return bool(class(self).IsSimulatingPhysics())
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
func (self Instance) PhysicalBonesStopSimulation() {
	class(self).PhysicalBonesStopSimulation()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
func (self Instance) PhysicalBonesStartSimulation() {
	class(self).PhysicalBonesStartSimulation([1]gd.Array{}[0])
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesAddCollisionException(exception Resource.ID) {
	class(self).PhysicalBonesAddCollisionException(exception)
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesRemoveCollisionException(exception Resource.ID) {
	class(self).PhysicalBonesRemoveCollisionException(exception)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PhysicalBoneSimulator3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalBoneSimulator3D"))
	return Instance{classdb.PhysicalBoneSimulator3D(object)}
}

/*
Returns a boolean that indicates whether the [PhysicalBoneSimulator3D] is running and simulating.
*/
//go:nosplit
func (self class) IsSimulatingPhysics() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
//go:nosplit
func (self class) PhysicalBonesStopSimulation() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
//go:nosplit
func (self class) PhysicalBonesStartSimulation(bones gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bones))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_start_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesAddCollisionException(exception gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesRemoveCollisionException(exception gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicalBoneSimulator3D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicalBoneSimulator3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}
func init() {
	classdb.Register("PhysicalBoneSimulator3D", func(ptr gd.Object) any {
		return [1]classdb.PhysicalBoneSimulator3D{classdb.PhysicalBoneSimulator3D(ptr)}
	})
}
