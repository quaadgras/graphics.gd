package PhysicalBone2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/RigidBody2D"
import "grow.graphics/gd/gdclass/PhysicsBody2D"
import "grow.graphics/gd/gdclass/CollisionObject2D"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The [PhysicalBone2D] node is a [RigidBody2D]-based node that can be used to make [Bone2D]s in a [Skeleton2D] react to physics.
[b]Note:[/b] To make the [Bone2D]s visually follow the [PhysicalBone2D] node, use a [SkeletonModification2DPhysicalBones] modification on the [Skeleton2D] parent.
[b]Note:[/b] The [PhysicalBone2D] node does not automatically create a [Joint2D] node to keep [PhysicalBone2D] nodes together. They must be created manually. For most cases, you want to use a [PinJoint2D] node. The [PhysicalBone2D] node will automatically configure the [Joint2D] node once it's been added as a child node.

*/
type Go [1]classdb.PhysicalBone2D

/*
Returns the first [Joint2D] child node, if one exists. This is mainly a helper function to make it easier to get the [Joint2D] that the [PhysicalBone2D] is autoconfiguring.
*/
func (self Go) GetJoint() gdclass.Joint2D {
	return gdclass.Joint2D(class(self).GetJoint())
}

/*
Returns a boolean that indicates whether the [PhysicalBone2D] is running and simulating using the Godot 2D physics engine. When [code]true[/code], the PhysicalBone2D node is using physics.
*/
func (self Go) IsSimulatingPhysics() bool {
	return bool(class(self).IsSimulatingPhysics())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicalBone2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalBone2D"))
	return Go{classdb.PhysicalBone2D(object)}
}

func (self Go) Bone2dNodepath() string {
		return string(class(self).GetBone2dNodepath().String())
}

func (self Go) SetBone2dNodepath(value string) {
	class(self).SetBone2dNodepath(gd.NewString(value).NodePath())
}

func (self Go) Bone2dIndex() int {
		return int(int(class(self).GetBone2dIndex()))
}

func (self Go) SetBone2dIndex(value int) {
	class(self).SetBone2dIndex(gd.Int(value))
}

func (self Go) AutoConfigureJoint() bool {
		return bool(class(self).GetAutoConfigureJoint())
}

func (self Go) SetAutoConfigureJoint(value bool) {
	class(self).SetAutoConfigureJoint(value)
}

func (self Go) SimulatePhysics() bool {
		return bool(class(self).GetSimulatePhysics())
}

func (self Go) SetSimulatePhysics(value bool) {
	class(self).SetSimulatePhysics(value)
}

func (self Go) FollowBoneWhenSimulating() bool {
		return bool(class(self).GetFollowBoneWhenSimulating())
}

func (self Go) SetFollowBoneWhenSimulating(value bool) {
	class(self).SetFollowBoneWhenSimulating(value)
}

/*
Returns the first [Joint2D] child node, if one exists. This is mainly a helper function to make it easier to get the [Joint2D] that the [PhysicalBone2D] is autoconfiguring.
*/
//go:nosplit
func (self class) GetJoint() gdclass.Joint2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Joint2D{classdb.Joint2D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetAutoConfigureJoint() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_auto_configure_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoConfigureJoint(auto_configure_joint bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, auto_configure_joint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_auto_configure_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSimulatePhysics(simulate_physics bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, simulate_physics)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimulatePhysics() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a boolean that indicates whether the [PhysicalBone2D] is running and simulating using the Godot 2D physics engine. When [code]true[/code], the PhysicalBone2D node is using physics.
*/
//go:nosplit
func (self class) IsSimulatingPhysics() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBone2dNodepath(nodepath gd.NodePath)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_bone2d_nodepath, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBone2dNodepath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_bone2d_nodepath, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBone2dIndex(bone_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_bone2d_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBone2dIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_bone2d_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFollowBoneWhenSimulating(follow_bone bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, follow_bone)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_follow_bone_when_simulating, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFollowBoneWhenSimulating() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_follow_bone_when_simulating, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicalBone2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicalBone2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRigidBody2D() RigidBody2D.GD { return *((*RigidBody2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsRigidBody2D() RigidBody2D.Go { return *((*RigidBody2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody2D() PhysicsBody2D.GD { return *((*PhysicsBody2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsBody2D() PhysicsBody2D.Go { return *((*PhysicsBody2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject2D() CollisionObject2D.GD { return *((*CollisionObject2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionObject2D() CollisionObject2D.Go { return *((*CollisionObject2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRigidBody2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRigidBody2D(), name)
	}
}
func init() {classdb.Register("PhysicalBone2D", func(ptr gd.Object) any { return classdb.PhysicalBone2D(ptr) })}
