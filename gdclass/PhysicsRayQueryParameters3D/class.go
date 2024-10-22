package PhysicsRayQueryParameters3D

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
By changing various properties of this object, such as the ray position, you can configure the parameters for [method PhysicsDirectSpaceState3D.intersect_ray].

*/
type Go [1]classdb.PhysicsRayQueryParameters3D

/*
Returns a new, pre-configured [PhysicsRayQueryParameters3D] object. Use it to quickly create query parameters using the most common options.
[codeblock]
var query = PhysicsRayQueryParameters3D.create(position, position + Vector3(0, -10, 0))
var collision = get_world_3d().direct_space_state.intersect_ray(query)
[/codeblock]
*/
func (self Go) Create(from gd.Vector3, to gd.Vector3) gdclass.PhysicsRayQueryParameters3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.PhysicsRayQueryParameters3D(class(self).Create(gc, from, to, gd.Int(4294967295), ([1]gd.ArrayOf[gd.RID]{}[0])))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsRayQueryParameters3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicsRayQueryParameters3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) From() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetFrom())
}

func (self Go) SetFrom(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrom(value)
}

func (self Go) To() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetTo())
}

func (self Go) SetTo(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTo(value)
}

func (self Go) CollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) Exclude() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.RID](class(self).GetExclude(gc))
}

func (self Go) SetExclude(value gd.ArrayOf[gd.RID]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExclude(value)
}

func (self Go) CollideWithBodies() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCollideWithBodiesEnabled())
}

func (self Go) SetCollideWithBodies(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollideWithBodies(value)
}

func (self Go) CollideWithAreas() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCollideWithAreasEnabled())
}

func (self Go) SetCollideWithAreas(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollideWithAreas(value)
}

func (self Go) HitFromInside() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHitFromInsideEnabled())
}

func (self Go) SetHitFromInside(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHitFromInside(value)
}

func (self Go) HitBackFaces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHitBackFacesEnabled())
}

func (self Go) SetHitBackFaces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHitBackFaces(value)
}

/*
Returns a new, pre-configured [PhysicsRayQueryParameters3D] object. Use it to quickly create query parameters using the most common options.
[codeblock]
var query = PhysicsRayQueryParameters3D.create(position, position + Vector3(0, -10, 0))
var collision = get_world_3d().direct_space_state.intersect_ray(query)
[/codeblock]
*/
//go:nosplit
func (self class) Create(ctx gd.Lifetime, from gd.Vector3, to gd.Vector3, collision_mask gd.Int, exclude gd.ArrayOf[gd.RID]) gdclass.PhysicsRayQueryParameters3D {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, collision_mask)
	callframe.Arg(frame, mmm.Get(exclude))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.PhysicsRayQueryParameters3D.Bind_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PhysicsRayQueryParameters3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrom(from gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrom() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_get_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTo(to gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTo() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_get_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExclude(exclude gd.ArrayOf[gd.RID])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(exclude))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_exclude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExclude(ctx gd.Lifetime) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_get_exclude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
//go:nosplit
func (self class) SetCollideWithBodies(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollideWithAreas(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHitFromInside(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_hit_from_inside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHitFromInsideEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_is_hit_from_inside_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHitBackFaces(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_set_hit_back_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHitBackFacesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsRayQueryParameters3D.Bind_is_hit_back_faces_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsRayQueryParameters3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsRayQueryParameters3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PhysicsRayQueryParameters3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
