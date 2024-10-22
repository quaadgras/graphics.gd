package PhysicsTestMotionParameters2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
By changing various properties of this object, such as the motion, you can configure the parameters for [method PhysicsServer2D.body_test_motion].

*/
type Go [1]classdb.PhysicsTestMotionParameters2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsTestMotionParameters2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicsTestMotionParameters2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) From() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform2D(class(self).GetFrom())
}

func (self Go) SetFrom(value gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrom(value)
}

func (self Go) Motion() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetMotion())
}

func (self Go) SetMotion(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMotion(value)
}

func (self Go) Margin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMargin()))
}

func (self Go) SetMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMargin(gd.Float(value))
}

func (self Go) CollideSeparationRay() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCollideSeparationRayEnabled())
}

func (self Go) SetCollideSeparationRay(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollideSeparationRayEnabled(value)
}

func (self Go) ExcludeBodies() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.RID](class(self).GetExcludeBodies(gc))
}

func (self Go) SetExcludeBodies(value gd.ArrayOf[gd.RID]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExcludeBodies(value)
}

func (self Go) ExcludeObjects() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.Int](class(self).GetExcludeObjects(gc))
}

func (self Go) SetExcludeObjects(value gd.ArrayOf[gd.Int]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExcludeObjects(value)
}

func (self Go) RecoveryAsCollision() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsRecoveryAsCollisionEnabled())
}

func (self Go) SetRecoveryAsCollision(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRecoveryAsCollisionEnabled(value)
}

//go:nosplit
func (self class) GetFrom() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_get_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrom(from gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_set_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMotion() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_get_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMotion(motion gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_set_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMargin(margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideSeparationRayEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_is_collide_separation_ray_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollideSeparationRayEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_set_collide_separation_ray_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExcludeBodies(ctx gd.Lifetime) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_get_exclude_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
//go:nosplit
func (self class) SetExcludeBodies(exclude_list gd.ArrayOf[gd.RID])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(exclude_list))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_set_exclude_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExcludeObjects(ctx gd.Lifetime) gd.ArrayOf[gd.Int] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_get_exclude_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Int](ret)
}
//go:nosplit
func (self class) SetExcludeObjects(exclude_list gd.ArrayOf[gd.Int])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(exclude_list))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_set_exclude_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRecoveryAsCollisionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_is_recovery_as_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRecoveryAsCollisionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionParameters2D.Bind_set_recovery_as_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicsTestMotionParameters2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsTestMotionParameters2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("PhysicsTestMotionParameters2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
