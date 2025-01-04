// Package AnimatableBody3D provides methods for working with AnimatableBody3D object instances.
package AnimatableBody3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/StaticBody3D"
import "graphics.gd/classdb/PhysicsBody3D"
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
An animatable 3D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform3D].
When [AnimatableBody3D] is moved, its linear and angular velocity are estimated and used to affect other physics bodies in its path. This makes it useful for moving platforms, doors, and other moving objects.
*/
type Instance [1]gdclass.AnimatableBody3D
type Any interface {
	gd.IsClass
	AsAnimatableBody3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimatableBody3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimatableBody3D"))
	return Instance{*(*gdclass.AnimatableBody3D)(unsafe.Pointer(&object))}
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatableBody3D.Bind_set_sync_to_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSyncToPhysicsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatableBody3D.Bind_is_sync_to_physics_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimatableBody3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimatableBody3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStaticBody3D() StaticBody3D.Advanced {
	return *((*StaticBody3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStaticBody3D() StaticBody3D.Instance {
	return *((*StaticBody3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPhysicsBody3D() PhysicsBody3D.Advanced {
	return *((*PhysicsBody3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsBody3D() PhysicsBody3D.Instance {
	return *((*PhysicsBody3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCollisionObject3D() CollisionObject3D.Advanced {
	return *((*CollisionObject3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject3D() CollisionObject3D.Instance {
	return *((*CollisionObject3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStaticBody3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStaticBody3D(), name)
	}
}
func init() {
	gdclass.Register("AnimatableBody3D", func(ptr gd.Object) any {
		return [1]gdclass.AnimatableBody3D{*(*gdclass.AnimatableBody3D)(unsafe.Pointer(&ptr))}
	})
}
