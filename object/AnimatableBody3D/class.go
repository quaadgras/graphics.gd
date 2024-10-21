package AnimatableBody3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/StaticBody3D"
import "grow.graphics/gd/object/PhysicsBody3D"
import "grow.graphics/gd/object/CollisionObject3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An animatable 3D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform3D].
When [AnimatableBody3D] is moved, its linear and angular velocity are estimated and used to affect other physics bodies in its path. This makes it useful for moving platforms, doors, and other moving objects.

*/
type Simple [1]classdb.AnimatableBody3D
func (self Simple) SetSyncToPhysics(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSyncToPhysics(enable)
}
func (self Simple) IsSyncToPhysicsEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSyncToPhysicsEnabled())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimatableBody3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSyncToPhysics(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatableBody3D.Bind_set_sync_to_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSyncToPhysicsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatableBody3D.Bind_is_sync_to_physics_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimatableBody3D() Expert { return self[0].AsAnimatableBody3D() }


//go:nosplit
func (self Simple) AsAnimatableBody3D() Simple { return self[0].AsAnimatableBody3D() }


//go:nosplit
func (self class) AsStaticBody3D() StaticBody3D.Expert { return self[0].AsStaticBody3D() }


//go:nosplit
func (self Simple) AsStaticBody3D() StaticBody3D.Simple { return self[0].AsStaticBody3D() }


//go:nosplit
func (self class) AsPhysicsBody3D() PhysicsBody3D.Expert { return self[0].AsPhysicsBody3D() }


//go:nosplit
func (self Simple) AsPhysicsBody3D() PhysicsBody3D.Simple { return self[0].AsPhysicsBody3D() }


//go:nosplit
func (self class) AsCollisionObject3D() CollisionObject3D.Expert { return self[0].AsCollisionObject3D() }


//go:nosplit
func (self Simple) AsCollisionObject3D() CollisionObject3D.Simple { return self[0].AsCollisionObject3D() }


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
func init() {classdb.Register("AnimatableBody3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
