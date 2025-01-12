// Package PhysicsDirectSpaceState3D provides methods for working with PhysicsDirectSpaceState3D object instances.
package PhysicsDirectSpaceState3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Provides direct access to a physics space in the [PhysicsServer3D]. It's used mainly to do queries against objects and areas residing in a given space.
*/
type Instance [1]gdclass.PhysicsDirectSpaceState3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsDirectSpaceState3D() Instance
}

/*
Checks whether a point is inside any solid shape. Position and other parameters are defined through [PhysicsPointQueryParameters3D]. The shapes the point is inside of are returned in an array containing dictionaries with the following fields:
[code]collider[/code]: The colliding object.
[code]collider_id[/code]: The colliding object's ID.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
The number of intersections can be limited with the [param max_results] parameter, to reduce the processing time.
*/
func (self Instance) IntersectPoint(parameters [1]gdclass.PhysicsPointQueryParameters3D) gd.Array {
	return gd.Array(class(self).IntersectPoint(parameters, gd.Int(32)))
}

/*
Intersects a ray in a given space. Ray position and other parameters are defined through [PhysicsRayQueryParameters3D]. The returned object is a dictionary with the following fields:
[code]collider[/code]: The colliding object.
[code]collider_id[/code]: The colliding object's ID.
[code]normal[/code]: The object's surface normal at the intersection point, or [code]Vector3(0, 0, 0)[/code] if the ray starts inside the shape and [member PhysicsRayQueryParameters3D.hit_from_inside] is [code]true[/code].
[code]position[/code]: The intersection point.
[code]face_index[/code]: The face index at the intersection point.
[b]Note:[/b] Returns a valid number only if the intersected shape is a [ConcavePolygonShape3D]. Otherwise, [code]-1[/code] is returned.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
If the ray did not intersect anything, then an empty dictionary is returned instead.
*/
func (self Instance) IntersectRay(parameters [1]gdclass.PhysicsRayQueryParameters3D) Dictionary.Any {
	return Dictionary.Any(class(self).IntersectRay(parameters))
}

/*
Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters3D] object, against the space. The intersected shapes are returned in an array containing dictionaries with the following fields:
[code]collider[/code]: The colliding object.
[code]collider_id[/code]: The colliding object's ID.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
The number of intersections can be limited with the [param max_results] parameter, to reduce the processing time.
[b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object.
*/
func (self Instance) IntersectShape(parameters [1]gdclass.PhysicsShapeQueryParameters3D) gd.Array {
	return gd.Array(class(self).IntersectShape(parameters, gd.Int(32)))
}

/*
Checks how far a [Shape3D] can move without colliding. All the parameters for the query, including the shape, are supplied through a [PhysicsShapeQueryParameters3D] object.
Returns an array with the safe and unsafe proportions (between 0 and 1) of the motion. The safe proportion is the maximum fraction of the motion that can be made without a collision. The unsafe proportion is the minimum fraction of the distance that must be moved for a collision. If no collision is detected a result of [code][1.0, 1.0][/code] will be returned.
[b]Note:[/b] Any [Shape3D]s that the shape is already colliding with e.g. inside of, will be ignored. Use [method collide_shape] to determine the [Shape3D]s that the shape is already colliding with.
*/
func (self Instance) CastMotion(parameters [1]gdclass.PhysicsShapeQueryParameters3D) []float32 {
	return []float32(class(self).CastMotion(parameters).AsSlice())
}

/*
Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters3D] object, against the space. The resulting array contains a list of points where the shape intersects another. Like with [method intersect_shape], the number of returned results can be limited to save processing time.
Returned points are a list of pairs of contact points. For each pair the first one is in the shape passed in [PhysicsShapeQueryParameters3D] object, second one is in the collided shape from the physics space.
[b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object.
*/
func (self Instance) CollideShape(parameters [1]gdclass.PhysicsShapeQueryParameters3D) gd.Array {
	return gd.Array(class(self).CollideShape(parameters, gd.Int(32)))
}

/*
Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters3D] object, against the space. If it collides with more than one shape, the nearest one is selected. The returned object is a dictionary containing the following fields:
[code]collider_id[/code]: The colliding object's ID.
[code]linear_velocity[/code]: The colliding object's velocity [Vector3]. If the object is an [Area3D], the result is [code](0, 0, 0)[/code].
[code]normal[/code]: The object's surface normal at the intersection point.
[code]point[/code]: The intersection point.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
If the shape did not intersect anything, then an empty dictionary is returned instead.
[b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object.
*/
func (self Instance) GetRestInfo(parameters [1]gdclass.PhysicsShapeQueryParameters3D) Dictionary.Any {
	return Dictionary.Any(class(self).GetRestInfo(parameters))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsDirectSpaceState3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsDirectSpaceState3D"))
	casted := Instance{*(*gdclass.PhysicsDirectSpaceState3D)(unsafe.Pointer(&object))}
	return casted
}

/*
Checks whether a point is inside any solid shape. Position and other parameters are defined through [PhysicsPointQueryParameters3D]. The shapes the point is inside of are returned in an array containing dictionaries with the following fields:
[code]collider[/code]: The colliding object.
[code]collider_id[/code]: The colliding object's ID.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
The number of intersections can be limited with the [param max_results] parameter, to reduce the processing time.
*/
//go:nosplit
func (self class) IntersectPoint(parameters [1]gdclass.PhysicsPointQueryParameters3D, max_results gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	callframe.Arg(frame, max_results)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState3D.Bind_intersect_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Intersects a ray in a given space. Ray position and other parameters are defined through [PhysicsRayQueryParameters3D]. The returned object is a dictionary with the following fields:
[code]collider[/code]: The colliding object.
[code]collider_id[/code]: The colliding object's ID.
[code]normal[/code]: The object's surface normal at the intersection point, or [code]Vector3(0, 0, 0)[/code] if the ray starts inside the shape and [member PhysicsRayQueryParameters3D.hit_from_inside] is [code]true[/code].
[code]position[/code]: The intersection point.
[code]face_index[/code]: The face index at the intersection point.
[b]Note:[/b] Returns a valid number only if the intersected shape is a [ConcavePolygonShape3D]. Otherwise, [code]-1[/code] is returned.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
If the ray did not intersect anything, then an empty dictionary is returned instead.
*/
//go:nosplit
func (self class) IntersectRay(parameters [1]gdclass.PhysicsRayQueryParameters3D) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState3D.Bind_intersect_ray, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters3D] object, against the space. The intersected shapes are returned in an array containing dictionaries with the following fields:
[code]collider[/code]: The colliding object.
[code]collider_id[/code]: The colliding object's ID.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
The number of intersections can be limited with the [param max_results] parameter, to reduce the processing time.
[b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object.
*/
//go:nosplit
func (self class) IntersectShape(parameters [1]gdclass.PhysicsShapeQueryParameters3D, max_results gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	callframe.Arg(frame, max_results)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState3D.Bind_intersect_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Checks how far a [Shape3D] can move without colliding. All the parameters for the query, including the shape, are supplied through a [PhysicsShapeQueryParameters3D] object.
Returns an array with the safe and unsafe proportions (between 0 and 1) of the motion. The safe proportion is the maximum fraction of the motion that can be made without a collision. The unsafe proportion is the minimum fraction of the distance that must be moved for a collision. If no collision is detected a result of [code][1.0, 1.0][/code] will be returned.
[b]Note:[/b] Any [Shape3D]s that the shape is already colliding with e.g. inside of, will be ignored. Use [method collide_shape] to determine the [Shape3D]s that the shape is already colliding with.
*/
//go:nosplit
func (self class) CastMotion(parameters [1]gdclass.PhysicsShapeQueryParameters3D) gd.PackedFloat32Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState3D.Bind_cast_motion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters3D] object, against the space. The resulting array contains a list of points where the shape intersects another. Like with [method intersect_shape], the number of returned results can be limited to save processing time.
Returned points are a list of pairs of contact points. For each pair the first one is in the shape passed in [PhysicsShapeQueryParameters3D] object, second one is in the collided shape from the physics space.
[b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object.
*/
//go:nosplit
func (self class) CollideShape(parameters [1]gdclass.PhysicsShapeQueryParameters3D, max_results gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	callframe.Arg(frame, max_results)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState3D.Bind_collide_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters3D] object, against the space. If it collides with more than one shape, the nearest one is selected. The returned object is a dictionary containing the following fields:
[code]collider_id[/code]: The colliding object's ID.
[code]linear_velocity[/code]: The colliding object's velocity [Vector3]. If the object is an [Area3D], the result is [code](0, 0, 0)[/code].
[code]normal[/code]: The object's surface normal at the intersection point.
[code]point[/code]: The intersection point.
[code]rid[/code]: The intersecting object's [RID].
[code]shape[/code]: The shape index of the colliding shape.
If the shape did not intersect anything, then an empty dictionary is returned instead.
[b]Note:[/b] This method does not take into account the [code]motion[/code] property of the object.
*/
//go:nosplit
func (self class) GetRestInfo(parameters [1]gdclass.PhysicsShapeQueryParameters3D) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState3D.Bind_get_rest_info, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsPhysicsDirectSpaceState3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectSpaceState3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("PhysicsDirectSpaceState3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsDirectSpaceState3D{*(*gdclass.PhysicsDirectSpaceState3D)(unsafe.Pointer(&ptr))}
	})
}
