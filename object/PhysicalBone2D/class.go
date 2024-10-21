package PhysicalBone2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/RigidBody2D"
import "grow.graphics/gd/object/PhysicsBody2D"
import "grow.graphics/gd/object/CollisionObject2D"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [PhysicalBone2D] node is a [RigidBody2D]-based node that can be used to make [Bone2D]s in a [Skeleton2D] react to physics.
[b]Note:[/b] To make the [Bone2D]s visually follow the [PhysicalBone2D] node, use a [SkeletonModification2DPhysicalBones] modification on the [Skeleton2D] parent.
[b]Note:[/b] The [PhysicalBone2D] node does not automatically create a [Joint2D] node to keep [PhysicalBone2D] nodes together. They must be created manually. For most cases, you want to use a [PinJoint2D] node. The [PhysicalBone2D] node will automatically configure the [Joint2D] node once it's been added as a child node.

*/
type Simple [1]classdb.PhysicalBone2D
func (self Simple) GetJoint() [1]classdb.Joint2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Joint2D(Expert(self).GetJoint(gc))
}
func (self Simple) GetAutoConfigureJoint() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAutoConfigureJoint())
}
func (self Simple) SetAutoConfigureJoint(auto_configure_joint bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoConfigureJoint(auto_configure_joint)
}
func (self Simple) SetSimulatePhysics(simulate_physics bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSimulatePhysics(simulate_physics)
}
func (self Simple) GetSimulatePhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSimulatePhysics())
}
func (self Simple) IsSimulatingPhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSimulatingPhysics())
}
func (self Simple) SetBone2dNodepath(nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBone2dNodepath(nodepath)
}
func (self Simple) GetBone2dNodepath() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetBone2dNodepath(gc))
}
func (self Simple) SetBone2dIndex(bone_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBone2dIndex(gd.Int(bone_index))
}
func (self Simple) GetBone2dIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBone2dIndex()))
}
func (self Simple) SetFollowBoneWhenSimulating(follow_bone bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFollowBoneWhenSimulating(follow_bone)
}
func (self Simple) GetFollowBoneWhenSimulating() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFollowBoneWhenSimulating())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicalBone2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the first [Joint2D] child node, if one exists. This is mainly a helper function to make it easier to get the [Joint2D] that the [PhysicalBone2D] is autoconfiguring.
*/
//go:nosplit
func (self class) GetJoint(ctx gd.Lifetime) object.Joint2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_get_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Joint2D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetAutoConfigureJoint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_get_auto_configure_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoConfigureJoint(auto_configure_joint bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, auto_configure_joint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_set_auto_configure_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSimulatePhysics(simulate_physics bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, simulate_physics)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_set_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimulatePhysics() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_get_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a boolean that indicates whether the [PhysicalBone2D] is running and simulating using the Godot 2D physics engine. When [code]true[/code], the PhysicalBone2D node is using physics.
*/
//go:nosplit
func (self class) IsSimulatingPhysics() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBone2dNodepath(nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_set_bone2d_nodepath, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBone2dNodepath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_get_bone2d_nodepath, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBone2dIndex(bone_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_set_bone2d_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBone2dIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_get_bone2d_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFollowBoneWhenSimulating(follow_bone bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, follow_bone)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_set_follow_bone_when_simulating, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFollowBoneWhenSimulating() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone2D.Bind_get_follow_bone_when_simulating, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPhysicalBone2D() Expert { return self[0].AsPhysicalBone2D() }


//go:nosplit
func (self Simple) AsPhysicalBone2D() Simple { return self[0].AsPhysicalBone2D() }


//go:nosplit
func (self class) AsRigidBody2D() RigidBody2D.Expert { return self[0].AsRigidBody2D() }


//go:nosplit
func (self Simple) AsRigidBody2D() RigidBody2D.Simple { return self[0].AsRigidBody2D() }


//go:nosplit
func (self class) AsPhysicsBody2D() PhysicsBody2D.Expert { return self[0].AsPhysicsBody2D() }


//go:nosplit
func (self Simple) AsPhysicsBody2D() PhysicsBody2D.Simple { return self[0].AsPhysicsBody2D() }


//go:nosplit
func (self class) AsCollisionObject2D() CollisionObject2D.Expert { return self[0].AsCollisionObject2D() }


//go:nosplit
func (self Simple) AsCollisionObject2D() CollisionObject2D.Simple { return self[0].AsCollisionObject2D() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


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
func init() {classdb.Register("PhysicalBone2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
