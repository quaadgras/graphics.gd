package PhysicsTestMotionParameters2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
By changing various properties of this object, such as the motion, you can configure the parameters for [method PhysicsServer2D.body_test_motion].
*/
type Instance [1]classdb.PhysicsTestMotionParameters2D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PhysicsTestMotionParameters2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsTestMotionParameters2D"))
	return Instance{classdb.PhysicsTestMotionParameters2D(object)}
}

func (self Instance) From() gd.Transform2D {
	return gd.Transform2D(class(self).GetFrom())
}

func (self Instance) SetFrom(value gd.Transform2D) {
	class(self).SetFrom(value)
}

func (self Instance) Motion() gd.Vector2 {
	return gd.Vector2(class(self).GetMotion())
}

func (self Instance) SetMotion(value gd.Vector2) {
	class(self).SetMotion(value)
}

func (self Instance) Margin() float64 {
	return float64(float64(class(self).GetMargin()))
}

func (self Instance) SetMargin(value float64) {
	class(self).SetMargin(gd.Float(value))
}

func (self Instance) CollideSeparationRay() bool {
	return bool(class(self).IsCollideSeparationRayEnabled())
}

func (self Instance) SetCollideSeparationRay(value bool) {
	class(self).SetCollideSeparationRayEnabled(value)
}

func (self Instance) ExcludeBodies() gd.Array {
	return gd.Array(class(self).GetExcludeBodies())
}

func (self Instance) SetExcludeBodies(value gd.Array) {
	class(self).SetExcludeBodies(value)
}

func (self Instance) ExcludeObjects() gd.Array {
	return gd.Array(class(self).GetExcludeObjects())
}

func (self Instance) SetExcludeObjects(value gd.Array) {
	class(self).SetExcludeObjects(value)
}

func (self Instance) RecoveryAsCollision() bool {
	return bool(class(self).IsRecoveryAsCollisionEnabled())
}

func (self Instance) SetRecoveryAsCollision(value bool) {
	class(self).SetRecoveryAsCollisionEnabled(value)
}

//go:nosplit
func (self class) GetFrom() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_get_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrom(from gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_set_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotion() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_get_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMotion(motion gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_set_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideSeparationRayEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_is_collide_separation_ray_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideSeparationRayEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_set_collide_separation_ray_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeBodies() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_get_exclude_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeBodies(exclude_list gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(exclude_list))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_set_exclude_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeObjects() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_get_exclude_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeObjects(exclude_list gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(exclude_list))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_set_exclude_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsRecoveryAsCollisionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_is_recovery_as_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRecoveryAsCollisionEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters2D.Bind_set_recovery_as_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicsTestMotionParameters2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsTestMotionParameters2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("PhysicsTestMotionParameters2D", func(ptr gd.Object) any { return classdb.PhysicsTestMotionParameters2D(ptr) })
}
