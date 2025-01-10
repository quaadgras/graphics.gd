// Package PhysicsRayQueryParameters2D provides methods for working with PhysicsRayQueryParameters2D object instances.
package PhysicsRayQueryParameters2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
By changing various properties of this object, such as the ray position, you can configure the parameters for [method PhysicsDirectSpaceState2D.intersect_ray].
*/
type Instance [1]gdclass.PhysicsRayQueryParameters2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsRayQueryParameters2D() Instance
}

/*
Returns a new, pre-configured [PhysicsRayQueryParameters2D] object. Use it to quickly create query parameters using the most common options.
[codeblock]
var query = PhysicsRayQueryParameters2D.create(global_position, global_position + Vector2(0, 100))
var collision = get_world_2d().direct_space_state.intersect_ray(query)
[/codeblock]
*/
func Create(from Vector2.XY, to Vector2.XY) [1]gdclass.PhysicsRayQueryParameters2D {
	self := Instance{}
	return [1]gdclass.PhysicsRayQueryParameters2D(class(self).Create(gd.Vector2(from), gd.Vector2(to), gd.Int(4294967295), [1]gd.Array{}[0]))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsRayQueryParameters2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsRayQueryParameters2D"))
	casted := Instance{*(*gdclass.PhysicsRayQueryParameters2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) From() Vector2.XY {
	return Vector2.XY(class(self).GetFrom())
}

func (self Instance) SetFrom(value Vector2.XY) {
	class(self).SetFrom(gd.Vector2(value))
}

func (self Instance) To() Vector2.XY {
	return Vector2.XY(class(self).GetTo())
}

func (self Instance) SetTo(value Vector2.XY) {
	class(self).SetTo(gd.Vector2(value))
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) Exclude() gd.Array {
	return gd.Array(class(self).GetExclude())
}

func (self Instance) SetExclude(value gd.Array) {
	class(self).SetExclude(value)
}

func (self Instance) CollideWithBodies() bool {
	return bool(class(self).IsCollideWithBodiesEnabled())
}

func (self Instance) SetCollideWithBodies(value bool) {
	class(self).SetCollideWithBodies(value)
}

func (self Instance) CollideWithAreas() bool {
	return bool(class(self).IsCollideWithAreasEnabled())
}

func (self Instance) SetCollideWithAreas(value bool) {
	class(self).SetCollideWithAreas(value)
}

func (self Instance) HitFromInside() bool {
	return bool(class(self).IsHitFromInsideEnabled())
}

func (self Instance) SetHitFromInside(value bool) {
	class(self).SetHitFromInside(value)
}

/*
Returns a new, pre-configured [PhysicsRayQueryParameters2D] object. Use it to quickly create query parameters using the most common options.
[codeblock]
var query = PhysicsRayQueryParameters2D.create(global_position, global_position + Vector2(0, 100))
var collision = get_world_2d().direct_space_state.intersect_ray(query)
[/codeblock]
*/
//go:nosplit
func (self class) Create(from gd.Vector2, to gd.Vector2, collision_mask gd.Int, exclude gd.Array) [1]gdclass.PhysicsRayQueryParameters2D {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, collision_mask)
	callframe.Arg(frame, pointers.Get(exclude))
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.PhysicsRayQueryParameters2D{gd.PointerWithOwnershipTransferredToGo[gdclass.PhysicsRayQueryParameters2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrom(from gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_set_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrom() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_get_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTo(to gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_set_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTo() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_get_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExclude(exclude gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(exclude))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_set_exclude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExclude() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_get_exclude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithBodies(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithAreas(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHitFromInside(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_set_hit_from_inside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHitFromInsideEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters2D.Bind_is_hit_from_inside_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsRayQueryParameters2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsRayQueryParameters2D() Instance {
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
	gdclass.Register("PhysicsRayQueryParameters2D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsRayQueryParameters2D{*(*gdclass.PhysicsRayQueryParameters2D)(unsafe.Pointer(&ptr))}
	})
}
