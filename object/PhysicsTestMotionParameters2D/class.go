package PhysicsTestMotionParameters2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
By changing various properties of this object, such as the motion, you can configure the parameters for [method PhysicsServer2D.body_test_motion].

*/
type Simple [1]classdb.PhysicsTestMotionParameters2D
func (self Simple) GetFrom() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetFrom())
}
func (self Simple) SetFrom(from gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrom(from)
}
func (self Simple) GetMotion() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMotion())
}
func (self Simple) SetMotion(motion gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMotion(motion)
}
func (self Simple) GetMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMargin()))
}
func (self Simple) SetMargin(margin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMargin(gd.Float(margin))
}
func (self Simple) IsCollideSeparationRayEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCollideSeparationRayEnabled())
}
func (self Simple) SetCollideSeparationRayEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollideSeparationRayEnabled(enabled)
}
func (self Simple) GetExcludeBodies() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).GetExcludeBodies(gc))
}
func (self Simple) SetExcludeBodies(exclude_list gd.ArrayOf[gd.RID]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExcludeBodies(exclude_list)
}
func (self Simple) GetExcludeObjects() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Int](Expert(self).GetExcludeObjects(gc))
}
func (self Simple) SetExcludeObjects(exclude_list gd.ArrayOf[gd.Int]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExcludeObjects(exclude_list)
}
func (self Simple) IsRecoveryAsCollisionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRecoveryAsCollisionEnabled())
}
func (self Simple) SetRecoveryAsCollisionEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRecoveryAsCollisionEnabled(enabled)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsTestMotionParameters2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsPhysicsTestMotionParameters2D() Expert { return self[0].AsPhysicsTestMotionParameters2D() }


//go:nosplit
func (self Simple) AsPhysicsTestMotionParameters2D() Simple { return self[0].AsPhysicsTestMotionParameters2D() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
func init() {classdb.Register("PhysicsTestMotionParameters2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
