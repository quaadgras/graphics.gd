package AnimatableBody2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/StaticBody2D"
import "graphics.gd/objects/PhysicsBody2D"
import "graphics.gd/objects/CollisionObject2D"
import "graphics.gd/objects/Node2D"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
An animatable 2D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform2D].
When [AnimatableBody2D] is moved, its linear and angular velocity are estimated and used to affect other physics bodies in its path. This makes it useful for moving platforms, doors, and other moving objects.
*/
type Instance [1]classdb.AnimatableBody2D
type Any interface {
	gd.IsClass
	AsAnimatableBody2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimatableBody2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimatableBody2D"))
	return Instance{classdb.AnimatableBody2D(object)}
}

func (self Instance) SyncToPhysics() bool {
	return bool(class(self).IsSyncToPhysicsEnabled())
}

func (self Instance) SetSyncToPhysics(value bool) {
	class(self).SetSyncToPhysics(value)
}

//go:nosplit
func (self class) SetSyncToPhysics(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatableBody2D.Bind_set_sync_to_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSyncToPhysicsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatableBody2D.Bind_is_sync_to_physics_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimatableBody2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimatableBody2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStaticBody2D() StaticBody2D.Advanced {
	return *((*StaticBody2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStaticBody2D() StaticBody2D.Instance {
	return *((*StaticBody2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPhysicsBody2D() PhysicsBody2D.Advanced {
	return *((*PhysicsBody2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsBody2D() PhysicsBody2D.Instance {
	return *((*PhysicsBody2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCollisionObject2D() CollisionObject2D.Advanced {
	return *((*CollisionObject2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject2D() CollisionObject2D.Instance {
	return *((*CollisionObject2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStaticBody2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStaticBody2D(), name)
	}
}
func init() {
	classdb.Register("AnimatableBody2D", func(ptr gd.Object) any { return [1]classdb.AnimatableBody2D{classdb.AnimatableBody2D(ptr)} })
}
