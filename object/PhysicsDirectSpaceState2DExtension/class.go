package PhysicsDirectSpaceState2DExtension

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PhysicsDirectSpaceState2D"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class extends [PhysicsDirectSpaceState2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectSpaceState2D].
	// PhysicsDirectSpaceState2DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsDirectSpaceState2DExtension interface {
		IntersectRay(from gd.Vector2, to gd.Vector2, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, result *classdb.PhysicsServer2DExtensionRayResult) bool
		IntersectPoint(position gd.Vector2, canvas_instance_id gd.Int, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results *classdb.PhysicsServer2DExtensionShapeResult, max_results gd.Int) gd.Int
		IntersectShape(shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, result *classdb.PhysicsServer2DExtensionShapeResult, max_results gd.Int) gd.Int
		CastMotion(shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64) bool
		CollideShape(shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results gd.Int, result_count *int32) bool
		RestInfo(shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, rest_info *classdb.PhysicsServer2DExtensionShapeRestInfo) bool
	}

*/
type Simple [1]classdb.PhysicsDirectSpaceState2DExtension
func (self Simple) IsBodyExcludedFromQuery(body gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBodyExcludedFromQuery(body))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsDirectSpaceState2DExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func (class) _intersect_ray(impl func(ptr unsafe.Pointer, from gd.Vector2, to gd.Vector2, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, result *classdb.PhysicsServer2DExtensionRayResult) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from = gd.UnsafeGet[gd.Vector2](p_args,0)
		var to = gd.UnsafeGet[gd.Vector2](p_args,1)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args,2)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args,3)
		var collide_with_areas = gd.UnsafeGet[bool](p_args,4)
		var hit_from_inside = gd.UnsafeGet[bool](p_args,5)
		var result = gd.UnsafeGet[*classdb.PhysicsServer2DExtensionRayResult](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from, to, collision_mask, collide_with_bodies, collide_with_areas, hit_from_inside, result)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _intersect_point(impl func(ptr unsafe.Pointer, position gd.Vector2, canvas_instance_id gd.Int, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results *classdb.PhysicsServer2DExtensionShapeResult, max_results gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var canvas_instance_id = gd.UnsafeGet[gd.Int](p_args,1)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args,2)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args,3)
		var collide_with_areas = gd.UnsafeGet[bool](p_args,4)
		var results = gd.UnsafeGet[*classdb.PhysicsServer2DExtensionShapeResult](p_args,5)
		var max_results = gd.UnsafeGet[gd.Int](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, position, canvas_instance_id, collision_mask, collide_with_bodies, collide_with_areas, results, max_results)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _intersect_shape(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, result *classdb.PhysicsServer2DExtensionShapeResult, max_results gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args,4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args,5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args,6)
		var result = gd.UnsafeGet[*classdb.PhysicsServer2DExtensionShapeResult](p_args,7)
		var max_results = gd.UnsafeGet[gd.Int](p_args,8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, result, max_results)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _cast_motion(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args,4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args,5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args,6)
		var closest_safe = gd.UnsafeGet[*float64](p_args,7)
		var closest_unsafe = gd.UnsafeGet[*float64](p_args,8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, closest_safe, closest_unsafe)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _collide_shape(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results gd.Int, result_count *int32) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args,4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args,5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args,6)
		var results = gd.UnsafeGet[unsafe.Pointer](p_args,7)
		var max_results = gd.UnsafeGet[gd.Int](p_args,8)
		var result_count = gd.UnsafeGet[*int32](p_args,9)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, results, max_results, result_count)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _rest_info(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, rest_info *classdb.PhysicsServer2DExtensionShapeRestInfo) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args,4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args,5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args,6)
		var rest_info = gd.UnsafeGet[*classdb.PhysicsServer2DExtensionShapeRestInfo](p_args,7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, rest_info)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) IsBodyExcludedFromQuery(body gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsDirectSpaceState2DExtension.Bind_is_body_excluded_from_query, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPhysicsDirectSpaceState2DExtension() Expert { return self[0].AsPhysicsDirectSpaceState2DExtension() }


//go:nosplit
func (self Simple) AsPhysicsDirectSpaceState2DExtension() Simple { return self[0].AsPhysicsDirectSpaceState2DExtension() }


//go:nosplit
func (self class) AsPhysicsDirectSpaceState2D() PhysicsDirectSpaceState2D.Expert { return self[0].AsPhysicsDirectSpaceState2D() }


//go:nosplit
func (self Simple) AsPhysicsDirectSpaceState2D() PhysicsDirectSpaceState2D.Simple { return self[0].AsPhysicsDirectSpaceState2D() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_intersect_ray": return reflect.ValueOf(self._intersect_ray);
	case "_intersect_point": return reflect.ValueOf(self._intersect_point);
	case "_intersect_shape": return reflect.ValueOf(self._intersect_shape);
	case "_cast_motion": return reflect.ValueOf(self._cast_motion);
	case "_collide_shape": return reflect.ValueOf(self._collide_shape);
	case "_rest_info": return reflect.ValueOf(self._rest_info);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PhysicsDirectSpaceState2DExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
