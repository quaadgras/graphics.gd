// Package PhysicsTestMotionParameters3D provides methods for working with PhysicsTestMotionParameters3D object instances.
package PhysicsTestMotionParameters3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/RID"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

/*
By changing various properties of this object, such as the motion, you can configure the parameters for [method PhysicsServer3D.body_test_motion].
*/
type Instance [1]gdclass.PhysicsTestMotionParameters3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsTestMotionParameters3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsTestMotionParameters3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsTestMotionParameters3D"))
	casted := Instance{*(*gdclass.PhysicsTestMotionParameters3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) From() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetFrom())
}

func (self Instance) SetFrom(value Transform3D.BasisOrigin) {
	class(self).SetFrom(gd.Transform3D(value))
}

func (self Instance) Motion() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetMotion())
}

func (self Instance) SetMotion(value Vector3.XYZ) {
	class(self).SetMotion(gd.Vector3(value))
}

func (self Instance) Margin() Float.X {
	return Float.X(Float.X(class(self).GetMargin()))
}

func (self Instance) SetMargin(value Float.X) {
	class(self).SetMargin(gd.Float(value))
}

func (self Instance) MaxCollisions() int {
	return int(int(class(self).GetMaxCollisions()))
}

func (self Instance) SetMaxCollisions(value int) {
	class(self).SetMaxCollisions(gd.Int(value))
}

func (self Instance) CollideSeparationRay() bool {
	return bool(class(self).IsCollideSeparationRayEnabled())
}

func (self Instance) SetCollideSeparationRay(value bool) {
	class(self).SetCollideSeparationRayEnabled(value)
}

func (self Instance) ExcludeBodies() []RID.Any {
	return []RID.Any(gd.ArrayAs[[]RID.Any](gd.InternalArray(class(self).GetExcludeBodies())))
}

func (self Instance) SetExcludeBodies(value []RID.Any) {
	class(self).SetExcludeBodies(gd.ArrayFromSlice[Array.Contains[gd.RID]](value))
}

func (self Instance) ExcludeObjects() []int {
	return []int(gd.ArrayAs[[]int](gd.InternalArray(class(self).GetExcludeObjects())))
}

func (self Instance) SetExcludeObjects(value []int) {
	class(self).SetExcludeObjects(gd.ArrayFromSlice[Array.Contains[gd.Int]](value))
}

func (self Instance) RecoveryAsCollision() bool {
	return bool(class(self).IsRecoveryAsCollisionEnabled())
}

func (self Instance) SetRecoveryAsCollision(value bool) {
	class(self).SetRecoveryAsCollisionEnabled(value)
}

//go:nosplit
func (self class) GetFrom() gd.Transform3D { //gd:PhysicsTestMotionParameters3D.get_from
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_get_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrom(from gd.Transform3D) { //gd:PhysicsTestMotionParameters3D.set_from
	var frame = callframe.New()
	callframe.Arg(frame, from)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotion() gd.Vector3 { //gd:PhysicsTestMotionParameters3D.get_motion
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_get_motion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMotion(motion gd.Vector3) { //gd:PhysicsTestMotionParameters3D.set_motion
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_motion, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Float { //gd:PhysicsTestMotionParameters3D.get_margin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Float) { //gd:PhysicsTestMotionParameters3D.set_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxCollisions() gd.Int { //gd:PhysicsTestMotionParameters3D.get_max_collisions
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_get_max_collisions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxCollisions(max_collisions gd.Int) { //gd:PhysicsTestMotionParameters3D.set_max_collisions
	var frame = callframe.New()
	callframe.Arg(frame, max_collisions)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_max_collisions, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideSeparationRayEnabled() bool { //gd:PhysicsTestMotionParameters3D.is_collide_separation_ray_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_is_collide_separation_ray_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideSeparationRayEnabled(enabled bool) { //gd:PhysicsTestMotionParameters3D.set_collide_separation_ray_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_collide_separation_ray_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeBodies() Array.Contains[gd.RID] { //gd:PhysicsTestMotionParameters3D.get_exclude_bodies
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_get_exclude_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeBodies(exclude_list Array.Contains[gd.RID]) { //gd:PhysicsTestMotionParameters3D.set_exclude_bodies
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(exclude_list)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_exclude_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeObjects() Array.Contains[gd.Int] { //gd:PhysicsTestMotionParameters3D.get_exclude_objects
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_get_exclude_objects, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Int]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeObjects(exclude_list Array.Contains[gd.Int]) { //gd:PhysicsTestMotionParameters3D.set_exclude_objects
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(exclude_list)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_exclude_objects, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRecoveryAsCollisionEnabled() bool { //gd:PhysicsTestMotionParameters3D.is_recovery_as_collision_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_is_recovery_as_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRecoveryAsCollisionEnabled(enabled bool) { //gd:PhysicsTestMotionParameters3D.set_recovery_as_collision_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionParameters3D.Bind_set_recovery_as_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsPhysicsTestMotionParameters3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsTestMotionParameters3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("PhysicsTestMotionParameters3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsTestMotionParameters3D{*(*gdclass.PhysicsTestMotionParameters3D)(unsafe.Pointer(&ptr))}
	})
}
