package AStar3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
A* (A star) is a computer algorithm used in pathfinding and graph traversal, the process of plotting short paths among vertices (points), passing through a given set of edges (segments). It enjoys widespread use due to its performance and accuracy. Godot's A* implementation uses points in 3D space and Euclidean distances by default.
You must add points manually with [method add_point] and create segments manually with [method connect_points]. Once done, you can test if there is a path between two points with the [method are_points_connected] function, get a path containing indices by [method get_id_path], or one containing actual coordinates with [method get_point_path].
It is also possible to use non-Euclidean distances. To do so, create a class that extends [AStar3D] and override methods [method _compute_cost] and [method _estimate_cost]. Both take two indices and return a length, as is shown in the following example.
[codeblocks]
[gdscript]
class MyAStar:
    extends AStar3D

    func _compute_cost(u, v):
        return abs(u - v)

    func _estimate_cost(u, v):
        return min(0, abs(u - v) - 1)
[/gdscript]
[csharp]
public partial class MyAStar : AStar3D
{
    public override float _ComputeCost(long fromId, long toId)
    {
        return Mathf.Abs((int)(fromId - toId));
    }

    public override float _EstimateCost(long fromId, long toId)
    {
        return Mathf.Min(0, Mathf.Abs((int)(fromId - toId)) - 1);
    }
}
[/csharp]
[/codeblocks]
[method _estimate_cost] should return a lower bound of the distance, i.e. [code]_estimate_cost(u, v) <= _compute_cost(u, v)[/code]. This serves as a hint to the algorithm because the custom [method _compute_cost] might be computation-heavy. If this is not the case, make [method _estimate_cost] return the same value as [method _compute_cost] to provide the algorithm with the most accurate information.
If the default [method _estimate_cost] and [method _compute_cost] methods are used, or if the supplied [method _estimate_cost] method returns a lower bound of the cost, then the paths returned by A* will be the lowest-cost paths. Here, the cost of a path equals the sum of the [method _compute_cost] results of all segments in the path multiplied by the [code]weight_scale[/code]s of the endpoints of the respective segments. If the default methods are used and the [code]weight_scale[/code]s of all points are set to [code]1.0[/code], then this equals the sum of Euclidean distances of all segments in the path.
	// AStar3D methods that can be overridden by a [Class] that extends it.
	type AStar3D interface {
		//Called when estimating the cost between a point and the path's ending point.
		//Note that this function is hidden in the default [AStar3D] class.
		EstimateCost(from_id gd.Int, to_id gd.Int) gd.Float
		//Called when computing the cost between two connected points.
		//Note that this function is hidden in the default [AStar3D] class.
		ComputeCost(from_id gd.Int, to_id gd.Int) gd.Float
	}

*/
type Simple [1]classdb.AStar3D
func (self Simple) GetAvailablePointId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAvailablePointId()))
}
func (self Simple) AddPoint(id int, position gd.Vector3, weight_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPoint(gd.Int(id), position, gd.Float(weight_scale))
}
func (self Simple) GetPointPosition(id int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetPointPosition(gd.Int(id)))
}
func (self Simple) SetPointPosition(id int, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointPosition(gd.Int(id), position)
}
func (self Simple) GetPointWeightScale(id int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPointWeightScale(gd.Int(id))))
}
func (self Simple) SetPointWeightScale(id int, weight_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointWeightScale(gd.Int(id), gd.Float(weight_scale))
}
func (self Simple) RemovePoint(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemovePoint(gd.Int(id))
}
func (self Simple) HasPoint(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasPoint(gd.Int(id)))
}
func (self Simple) GetPointConnections(id int) gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).GetPointConnections(gc, gd.Int(id)))
}
func (self Simple) GetPointIds() gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).GetPointIds(gc))
}
func (self Simple) SetPointDisabled(id int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointDisabled(gd.Int(id), disabled)
}
func (self Simple) IsPointDisabled(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPointDisabled(gd.Int(id)))
}
func (self Simple) ConnectPoints(id int, to_id int, bidirectional bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ConnectPoints(gd.Int(id), gd.Int(to_id), bidirectional)
}
func (self Simple) DisconnectPoints(id int, to_id int, bidirectional bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectPoints(gd.Int(id), gd.Int(to_id), bidirectional)
}
func (self Simple) ArePointsConnected(id int, to_id int, bidirectional bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ArePointsConnected(gd.Int(id), gd.Int(to_id), bidirectional))
}
func (self Simple) GetPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPointCount()))
}
func (self Simple) GetPointCapacity() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPointCapacity()))
}
func (self Simple) ReserveSpace(num_nodes int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReserveSpace(gd.Int(num_nodes))
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) GetClosestPoint(to_position gd.Vector3, include_disabled bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetClosestPoint(to_position, include_disabled)))
}
func (self Simple) GetClosestPositionInSegment(to_position gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetClosestPositionInSegment(to_position))
}
func (self Simple) GetPointPath(from_id int, to_id int, allow_partial_path bool) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).GetPointPath(gc, gd.Int(from_id), gd.Int(to_id), allow_partial_path))
}
func (self Simple) GetIdPath(from_id int, to_id int, allow_partial_path bool) gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).GetIdPath(gc, gd.Int(from_id), gd.Int(to_id), allow_partial_path))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AStar3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Called when estimating the cost between a point and the path's ending point.
Note that this function is hidden in the default [AStar3D] class.
*/
func (class) _estimate_cost(impl func(ptr unsafe.Pointer, from_id gd.Int, to_id gd.Int) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from_id = gd.UnsafeGet[gd.Int](p_args,0)
		var to_id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when computing the cost between two connected points.
Note that this function is hidden in the default [AStar3D] class.
*/
func (class) _compute_cost(impl func(ptr unsafe.Pointer, from_id gd.Int, to_id gd.Int) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from_id = gd.UnsafeGet[gd.Int](p_args,0)
		var to_id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the next available point ID with no point associated to it.
*/
//go:nosplit
func (self class) GetAvailablePointId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_available_point_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new point at the given position with the given identifier. The [param id] must be 0 or larger, and the [param weight_scale] must be 0.0 or greater.
The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point. Thus, all else being equal, the algorithm prefers points with lower [param weight_scale]s to form a path.
[codeblocks]
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(1, 0, 0), 4) # Adds the point (1, 0, 0) with weight_scale 4 and id 1
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(1, 0, 0), 4); // Adds the point (1, 0, 0) with weight_scale 4 and id 1
[/csharp]
[/codeblocks]
If there already exists a point for the given [param id], its position and weight scale are updated to the given values.
*/
//go:nosplit
func (self class) AddPoint(id gd.Int, position gd.Vector3, weight_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, position)
	callframe.Arg(frame, weight_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the point associated with the given [param id].
*/
//go:nosplit
func (self class) GetPointPosition(id gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param position] for the point with the given [param id].
*/
//go:nosplit
func (self class) SetPointPosition(id gd.Int, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_set_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the weight scale of the point associated with the given [param id].
*/
//go:nosplit
func (self class) GetPointWeightScale(id gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_point_weight_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param weight_scale] for the point with the given [param id]. The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point.
*/
//go:nosplit
func (self class) SetPointWeightScale(id gd.Int, weight_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, weight_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_set_point_weight_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the point associated with the given [param id] from the points pool.
*/
//go:nosplit
func (self class) RemovePoint(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether a point associated with the given [param id] exists.
*/
//go:nosplit
func (self class) HasPoint(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_has_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with the IDs of the points that form the connection with the given point.
[codeblocks]
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(0, 0, 0))
astar.add_point(2, Vector3(0, 1, 0))
astar.add_point(3, Vector3(1, 1, 0))
astar.add_point(4, Vector3(2, 0, 0))

astar.connect_points(1, 2, true)
astar.connect_points(1, 3, true)

var neighbors = astar.get_point_connections(1) # Returns [2, 3]
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(0, 0, 0));
astar.AddPoint(2, new Vector3(0, 1, 0));
astar.AddPoint(3, new Vector3(1, 1, 0));
astar.AddPoint(4, new Vector3(2, 0, 0));
astar.ConnectPoints(1, 2, true);
astar.ConnectPoints(1, 3, true);

int[] neighbors = astar.GetPointConnections(1); // Returns [2, 3]
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetPointConnections(ctx gd.Lifetime, id gd.Int) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_point_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of all point IDs.
*/
//go:nosplit
func (self class) GetPointIds(ctx gd.Lifetime) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_point_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Disables or enables the specified point for pathfinding. Useful for making a temporary obstacle.
*/
//go:nosplit
func (self class) SetPointDisabled(id gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_set_point_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether a point is disabled or not for pathfinding. By default, all points are enabled.
*/
//go:nosplit
func (self class) IsPointDisabled(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_is_point_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a segment between the given points. If [param bidirectional] is [code]false[/code], only movement from [param id] to [param to_id] is allowed, not the reverse direction.
[codeblocks]
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(1, 1, 0))
astar.add_point(2, Vector3(0, 5, 0))
astar.connect_points(1, 2, false)
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(1, 1, 0));
astar.AddPoint(2, new Vector3(0, 5, 0));
astar.ConnectPoints(1, 2, false);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) ConnectPoints(id gd.Int, to_id gd.Int, bidirectional bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, bidirectional)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_connect_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes the segment between the given points. If [param bidirectional] is [code]false[/code], only movement from [param id] to [param to_id] is prevented, and a unidirectional segment possibly remains.
*/
//go:nosplit
func (self class) DisconnectPoints(id gd.Int, to_id gd.Int, bidirectional bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, bidirectional)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_disconnect_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the two given points are directly connected by a segment. If [param bidirectional] is [code]false[/code], returns whether movement from [param id] to [param to_id] is possible through this segment.
*/
//go:nosplit
func (self class) ArePointsConnected(id gd.Int, to_id gd.Int, bidirectional bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, bidirectional)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_are_points_connected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of points currently in the points pool.
*/
//go:nosplit
func (self class) GetPointCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the capacity of the structure backing the points, useful in conjunction with [method reserve_space].
*/
//go:nosplit
func (self class) GetPointCapacity() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_point_capacity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Reserves space internally for [param num_nodes] points. Useful if you're adding a known large number of points at once, such as points on a grid. New capacity must be greater or equals to old capacity.
*/
//go:nosplit
func (self class) ReserveSpace(num_nodes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, num_nodes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_reserve_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all the points and segments.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the ID of the closest point to [param to_position], optionally taking disabled points into account. Returns [code]-1[/code] if there are no points in the points pool.
[b]Note:[/b] If several points are the closest to [param to_position], the one with the smallest ID will be returned, ensuring a deterministic result.
*/
//go:nosplit
func (self class) GetClosestPoint(to_position gd.Vector3, include_disabled bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	callframe.Arg(frame, include_disabled)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the closest position to [param to_position] that resides inside a segment between two connected points.
[codeblocks]
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(0, 0, 0))
astar.add_point(2, Vector3(0, 5, 0))
astar.connect_points(1, 2)
var res = astar.get_closest_position_in_segment(Vector3(3, 3, 0)) # Returns (0, 3, 0)
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(0, 0, 0));
astar.AddPoint(2, new Vector3(0, 5, 0));
astar.ConnectPoints(1, 2);
Vector3 res = astar.GetClosestPositionInSegment(new Vector3(3, 3, 0)); // Returns (0, 3, 0)
[/csharp]
[/codeblocks]
The result is in the segment that goes from [code]y = 0[/code] to [code]y = 5[/code]. It's the closest position in the segment to the given point.
*/
//go:nosplit
func (self class) GetClosestPositionInSegment(to_position gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_closest_position_in_segment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with the points that are in the path found by AStar3D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[b]Note:[/b] This method is not thread-safe. If called from a [Thread], it will return an empty array and will print an error message.
*/
//go:nosplit
func (self class) GetPointPath(ctx gd.Lifetime, from_id gd.Int, to_id gd.Int, allow_partial_path bool) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, allow_partial_path)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_point_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with the IDs of the points that form the path found by AStar3D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[codeblocks]
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(0, 0, 0))
astar.add_point(2, Vector3(0, 1, 0), 1) # Default weight is 1
astar.add_point(3, Vector3(1, 1, 0))
astar.add_point(4, Vector3(2, 0, 0))

astar.connect_points(1, 2, false)
astar.connect_points(2, 3, false)
astar.connect_points(4, 3, false)
astar.connect_points(1, 4, false)

var res = astar.get_id_path(1, 3) # Returns [1, 2, 3]
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(0, 0, 0));
astar.AddPoint(2, new Vector3(0, 1, 0), 1); // Default weight is 1
astar.AddPoint(3, new Vector3(1, 1, 0));
astar.AddPoint(4, new Vector3(2, 0, 0));
astar.ConnectPoints(1, 2, false);
astar.ConnectPoints(2, 3, false);
astar.ConnectPoints(4, 3, false);
astar.ConnectPoints(1, 4, false);
int[] res = astar.GetIdPath(1, 3); // Returns [1, 2, 3]
[/csharp]
[/codeblocks]
If you change the 2nd point's weight to 3, then the result will be [code][1, 4, 3][/code] instead, because now even though the distance is longer, it's "easier" to get through point 4 than through point 2.
*/
//go:nosplit
func (self class) GetIdPath(ctx gd.Lifetime, from_id gd.Int, to_id gd.Int, allow_partial_path bool) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, allow_partial_path)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStar3D.Bind_get_id_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAStar3D() Expert { return self[0].AsAStar3D() }


//go:nosplit
func (self Simple) AsAStar3D() Simple { return self[0].AsAStar3D() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_estimate_cost": return reflect.ValueOf(self._estimate_cost);
	case "_compute_cost": return reflect.ValueOf(self._compute_cost);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AStar3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
