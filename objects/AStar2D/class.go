package AStar2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
An implementation of the A* algorithm, used to find the shortest path between two vertices on a connected graph in 2D space.
See [AStar3D] for a more thorough explanation on how to use this class. [AStar2D] is a wrapper for [AStar3D] that enforces 2D coordinates.

	// AStar2D methods that can be overridden by a [Class] that extends it.
	type AStar2D interface {
		//Called when estimating the cost between a point and the path's ending point.
		//Note that this function is hidden in the default [AStar2D] class.
		EstimateCost(from_id int, to_id int) Float.X
		//Called when computing the cost between two connected points.
		//Note that this function is hidden in the default [AStar2D] class.
		ComputeCost(from_id int, to_id int) Float.X
	}
*/
type Instance [1]classdb.AStar2D

/*
Called when estimating the cost between a point and the path's ending point.
Note that this function is hidden in the default [AStar2D] class.
*/
func (Instance) _estimate_cost(impl func(ptr unsafe.Pointer, from_id int, to_id int) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var to_id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(from_id), int(to_id))
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Called when computing the cost between two connected points.
Note that this function is hidden in the default [AStar2D] class.
*/
func (Instance) _compute_cost(impl func(ptr unsafe.Pointer, from_id int, to_id int) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var to_id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(from_id), int(to_id))
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Returns the next available point ID with no point associated to it.
*/
func (self Instance) GetAvailablePointId() int {
	return int(int(class(self).GetAvailablePointId()))
}

/*
Adds a new point at the given position with the given identifier. The [param id] must be 0 or larger, and the [param weight_scale] must be 0.0 or greater.
The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point. Thus, all else being equal, the algorithm prefers points with lower [param weight_scale]s to form a path.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(1, 0), 4) # Adds the point (1, 0) with weight_scale 4 and id 1
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(1, 0), 4); // Adds the point (1, 0) with weight_scale 4 and id 1
[/csharp]
[/codeblocks]
If there already exists a point for the given [param id], its position and weight scale are updated to the given values.
*/
func (self Instance) AddPoint(id int, position Vector2.XY) {
	class(self).AddPoint(gd.Int(id), gd.Vector2(position), gd.Float(1.0))
}

/*
Returns the position of the point associated with the given [param id].
*/
func (self Instance) GetPointPosition(id int) Vector2.XY {
	return Vector2.XY(class(self).GetPointPosition(gd.Int(id)))
}

/*
Sets the [param position] for the point with the given [param id].
*/
func (self Instance) SetPointPosition(id int, position Vector2.XY) {
	class(self).SetPointPosition(gd.Int(id), gd.Vector2(position))
}

/*
Returns the weight scale of the point associated with the given [param id].
*/
func (self Instance) GetPointWeightScale(id int) Float.X {
	return Float.X(Float.X(class(self).GetPointWeightScale(gd.Int(id))))
}

/*
Sets the [param weight_scale] for the point with the given [param id]. The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point.
*/
func (self Instance) SetPointWeightScale(id int, weight_scale Float.X) {
	class(self).SetPointWeightScale(gd.Int(id), gd.Float(weight_scale))
}

/*
Removes the point associated with the given [param id] from the points pool.
*/
func (self Instance) RemovePoint(id int) {
	class(self).RemovePoint(gd.Int(id))
}

/*
Returns whether a point associated with the given [param id] exists.
*/
func (self Instance) HasPoint(id int) bool {
	return bool(class(self).HasPoint(gd.Int(id)))
}

/*
Returns an array with the IDs of the points that form the connection with the given point.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 1))
astar.add_point(3, Vector2(1, 1))
astar.add_point(4, Vector2(2, 0))

astar.connect_points(1, 2, true)
astar.connect_points(1, 3, true)

var neighbors = astar.get_point_connections(1) # Returns [2, 3]
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 1));
astar.AddPoint(3, new Vector2(1, 1));
astar.AddPoint(4, new Vector2(2, 0));

astar.ConnectPoints(1, 2, true);
astar.ConnectPoints(1, 3, true);

int[] neighbors = astar.GetPointConnections(1); // Returns [2, 3]
[/csharp]
[/codeblocks]
*/
func (self Instance) GetPointConnections(id int) []int64 {
	return []int64(class(self).GetPointConnections(gd.Int(id)).AsSlice())
}

/*
Returns an array of all point IDs.
*/
func (self Instance) GetPointIds() []int64 {
	return []int64(class(self).GetPointIds().AsSlice())
}

/*
Disables or enables the specified point for pathfinding. Useful for making a temporary obstacle.
*/
func (self Instance) SetPointDisabled(id int) {
	class(self).SetPointDisabled(gd.Int(id), true)
}

/*
Returns whether a point is disabled or not for pathfinding. By default, all points are enabled.
*/
func (self Instance) IsPointDisabled(id int) bool {
	return bool(class(self).IsPointDisabled(gd.Int(id)))
}

/*
Creates a segment between the given points. If [param bidirectional] is [code]false[/code], only movement from [param id] to [param to_id] is allowed, not the reverse direction.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(1, 1))
astar.add_point(2, Vector2(0, 5))
astar.connect_points(1, 2, false)
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(1, 1));
astar.AddPoint(2, new Vector2(0, 5));
astar.ConnectPoints(1, 2, false);
[/csharp]
[/codeblocks]
*/
func (self Instance) ConnectPoints(id int, to_id int) {
	class(self).ConnectPoints(gd.Int(id), gd.Int(to_id), true)
}

/*
Deletes the segment between the given points. If [param bidirectional] is [code]false[/code], only movement from [param id] to [param to_id] is prevented, and a unidirectional segment possibly remains.
*/
func (self Instance) DisconnectPoints(id int, to_id int) {
	class(self).DisconnectPoints(gd.Int(id), gd.Int(to_id), true)
}

/*
Returns whether there is a connection/segment between the given points. If [param bidirectional] is [code]false[/code], returns whether movement from [param id] to [param to_id] is possible through this segment.
*/
func (self Instance) ArePointsConnected(id int, to_id int) bool {
	return bool(class(self).ArePointsConnected(gd.Int(id), gd.Int(to_id), true))
}

/*
Returns the number of points currently in the points pool.
*/
func (self Instance) GetPointCount() int {
	return int(int(class(self).GetPointCount()))
}

/*
Returns the capacity of the structure backing the points, useful in conjunction with [method reserve_space].
*/
func (self Instance) GetPointCapacity() int {
	return int(int(class(self).GetPointCapacity()))
}

/*
Reserves space internally for [param num_nodes] points, useful if you're adding a known large number of points at once, such as points on a grid. New capacity must be greater or equals to old capacity.
*/
func (self Instance) ReserveSpace(num_nodes int) {
	class(self).ReserveSpace(gd.Int(num_nodes))
}

/*
Clears all the points and segments.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Returns the ID of the closest point to [param to_position], optionally taking disabled points into account. Returns [code]-1[/code] if there are no points in the points pool.
[b]Note:[/b] If several points are the closest to [param to_position], the one with the smallest ID will be returned, ensuring a deterministic result.
*/
func (self Instance) GetClosestPoint(to_position Vector2.XY) int {
	return int(int(class(self).GetClosestPoint(gd.Vector2(to_position), false)))
}

/*
Returns the closest position to [param to_position] that resides inside a segment between two connected points.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 5))
astar.connect_points(1, 2)
var res = astar.get_closest_position_in_segment(Vector2(3, 3)) # Returns (0, 3)
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 5));
astar.ConnectPoints(1, 2);
Vector2 res = astar.GetClosestPositionInSegment(new Vector2(3, 3)); // Returns (0, 3)
[/csharp]
[/codeblocks]
The result is in the segment that goes from [code]y = 0[/code] to [code]y = 5[/code]. It's the closest position in the segment to the given point.
*/
func (self Instance) GetClosestPositionInSegment(to_position Vector2.XY) Vector2.XY {
	return Vector2.XY(class(self).GetClosestPositionInSegment(gd.Vector2(to_position)))
}

/*
Returns an array with the points that are in the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[b]Note:[/b] This method is not thread-safe. If called from a [Thread], it will return an empty array and will print an error message.
*/
func (self Instance) GetPointPath(from_id int, to_id int) []Vector2.XY {
	return []Vector2.XY(class(self).GetPointPath(gd.Int(from_id), gd.Int(to_id), false).AsSlice())
}

/*
Returns an array with the IDs of the points that form the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 1), 1) # Default weight is 1
astar.add_point(3, Vector2(1, 1))
astar.add_point(4, Vector2(2, 0))

astar.connect_points(1, 2, false)
astar.connect_points(2, 3, false)
astar.connect_points(4, 3, false)
astar.connect_points(1, 4, false)

var res = astar.get_id_path(1, 3) # Returns [1, 2, 3]
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 1), 1); // Default weight is 1
astar.AddPoint(3, new Vector2(1, 1));
astar.AddPoint(4, new Vector2(2, 0));

astar.ConnectPoints(1, 2, false);
astar.ConnectPoints(2, 3, false);
astar.ConnectPoints(4, 3, false);
astar.ConnectPoints(1, 4, false);
int[] res = astar.GetIdPath(1, 3); // Returns [1, 2, 3]
[/csharp]
[/codeblocks]
If you change the 2nd point's weight to 3, then the result will be [code][1, 4, 3][/code] instead, because now even though the distance is longer, it's "easier" to get through point 4 than through point 2.
*/
func (self Instance) GetIdPath(from_id int, to_id int) []int64 {
	return []int64(class(self).GetIdPath(gd.Int(from_id), gd.Int(to_id), false).AsSlice())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AStar2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AStar2D"))
	return Instance{classdb.AStar2D(object)}
}

/*
Called when estimating the cost between a point and the path's ending point.
Note that this function is hidden in the default [AStar2D] class.
*/
func (class) _estimate_cost(impl func(ptr unsafe.Pointer, from_id gd.Int, to_id gd.Int) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var to_id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when computing the cost between two connected points.
Note that this function is hidden in the default [AStar2D] class.
*/
func (class) _compute_cost(impl func(ptr unsafe.Pointer, from_id gd.Int, to_id gd.Int) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var to_id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the next available point ID with no point associated to it.
*/
//go:nosplit
func (self class) GetAvailablePointId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_available_point_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new point at the given position with the given identifier. The [param id] must be 0 or larger, and the [param weight_scale] must be 0.0 or greater.
The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point. Thus, all else being equal, the algorithm prefers points with lower [param weight_scale]s to form a path.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(1, 0), 4) # Adds the point (1, 0) with weight_scale 4 and id 1
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(1, 0), 4); // Adds the point (1, 0) with weight_scale 4 and id 1
[/csharp]
[/codeblocks]
If there already exists a point for the given [param id], its position and weight scale are updated to the given values.
*/
//go:nosplit
func (self class) AddPoint(id gd.Int, position gd.Vector2, weight_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, position)
	callframe.Arg(frame, weight_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the position of the point associated with the given [param id].
*/
//go:nosplit
func (self class) GetPointPosition(id gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param position] for the point with the given [param id].
*/
//go:nosplit
func (self class) SetPointPosition(id gd.Int, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_set_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the weight scale of the point associated with the given [param id].
*/
//go:nosplit
func (self class) GetPointWeightScale(id gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_point_weight_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param weight_scale] for the point with the given [param id]. The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point.
*/
//go:nosplit
func (self class) SetPointWeightScale(id gd.Int, weight_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, weight_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_set_point_weight_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the point associated with the given [param id] from the points pool.
*/
//go:nosplit
func (self class) RemovePoint(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether a point associated with the given [param id] exists.
*/
//go:nosplit
func (self class) HasPoint(id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_has_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array with the IDs of the points that form the connection with the given point.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 1))
astar.add_point(3, Vector2(1, 1))
astar.add_point(4, Vector2(2, 0))

astar.connect_points(1, 2, true)
astar.connect_points(1, 3, true)

var neighbors = astar.get_point_connections(1) # Returns [2, 3]
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 1));
astar.AddPoint(3, new Vector2(1, 1));
astar.AddPoint(4, new Vector2(2, 0));

astar.ConnectPoints(1, 2, true);
astar.ConnectPoints(1, 3, true);

int[] neighbors = astar.GetPointConnections(1); // Returns [2, 3]
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetPointConnections(id gd.Int) gd.PackedInt64Array {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_point_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt64Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of all point IDs.
*/
//go:nosplit
func (self class) GetPointIds() gd.PackedInt64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_point_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt64Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Disables or enables the specified point for pathfinding. Useful for making a temporary obstacle.
*/
//go:nosplit
func (self class) SetPointDisabled(id gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_set_point_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether a point is disabled or not for pathfinding. By default, all points are enabled.
*/
//go:nosplit
func (self class) IsPointDisabled(id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_is_point_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a segment between the given points. If [param bidirectional] is [code]false[/code], only movement from [param id] to [param to_id] is allowed, not the reverse direction.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(1, 1))
astar.add_point(2, Vector2(0, 5))
astar.connect_points(1, 2, false)
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(1, 1));
astar.AddPoint(2, new Vector2(0, 5));
astar.ConnectPoints(1, 2, false);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) ConnectPoints(id gd.Int, to_id gd.Int, bidirectional bool) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, bidirectional)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_connect_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Deletes the segment between the given points. If [param bidirectional] is [code]false[/code], only movement from [param id] to [param to_id] is prevented, and a unidirectional segment possibly remains.
*/
//go:nosplit
func (self class) DisconnectPoints(id gd.Int, to_id gd.Int, bidirectional bool) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, bidirectional)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_disconnect_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether there is a connection/segment between the given points. If [param bidirectional] is [code]false[/code], returns whether movement from [param id] to [param to_id] is possible through this segment.
*/
//go:nosplit
func (self class) ArePointsConnected(id gd.Int, to_id gd.Int, bidirectional bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, bidirectional)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_are_points_connected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of points currently in the points pool.
*/
//go:nosplit
func (self class) GetPointCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the capacity of the structure backing the points, useful in conjunction with [method reserve_space].
*/
//go:nosplit
func (self class) GetPointCapacity() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_point_capacity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Reserves space internally for [param num_nodes] points, useful if you're adding a known large number of points at once, such as points on a grid. New capacity must be greater or equals to old capacity.
*/
//go:nosplit
func (self class) ReserveSpace(num_nodes gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, num_nodes)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_reserve_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears all the points and segments.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the ID of the closest point to [param to_position], optionally taking disabled points into account. Returns [code]-1[/code] if there are no points in the points pool.
[b]Note:[/b] If several points are the closest to [param to_position], the one with the smallest ID will be returned, ensuring a deterministic result.
*/
//go:nosplit
func (self class) GetClosestPoint(to_position gd.Vector2, include_disabled bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	callframe.Arg(frame, include_disabled)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the closest position to [param to_position] that resides inside a segment between two connected points.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 5))
astar.connect_points(1, 2)
var res = astar.get_closest_position_in_segment(Vector2(3, 3)) # Returns (0, 3)
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 5));
astar.ConnectPoints(1, 2);
Vector2 res = astar.GetClosestPositionInSegment(new Vector2(3, 3)); // Returns (0, 3)
[/csharp]
[/codeblocks]
The result is in the segment that goes from [code]y = 0[/code] to [code]y = 5[/code]. It's the closest position in the segment to the given point.
*/
//go:nosplit
func (self class) GetClosestPositionInSegment(to_position gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_closest_position_in_segment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array with the points that are in the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[b]Note:[/b] This method is not thread-safe. If called from a [Thread], it will return an empty array and will print an error message.
*/
//go:nosplit
func (self class) GetPointPath(from_id gd.Int, to_id gd.Int, allow_partial_path bool) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, from_id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, allow_partial_path)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_point_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array with the IDs of the points that form the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[codeblocks]
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 1), 1) # Default weight is 1
astar.add_point(3, Vector2(1, 1))
astar.add_point(4, Vector2(2, 0))

astar.connect_points(1, 2, false)
astar.connect_points(2, 3, false)
astar.connect_points(4, 3, false)
astar.connect_points(1, 4, false)

var res = astar.get_id_path(1, 3) # Returns [1, 2, 3]
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 1), 1); // Default weight is 1
astar.AddPoint(3, new Vector2(1, 1));
astar.AddPoint(4, new Vector2(2, 0));

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
func (self class) GetIdPath(from_id gd.Int, to_id gd.Int, allow_partial_path bool) gd.PackedInt64Array {
	var frame = callframe.New()
	callframe.Arg(frame, from_id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, allow_partial_path)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AStar2D.Bind_get_id_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt64Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsAStar2D() Advanced            { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAStar2D() Instance         { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_estimate_cost":
		return reflect.ValueOf(self._estimate_cost)
	case "_compute_cost":
		return reflect.ValueOf(self._compute_cost)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_estimate_cost":
		return reflect.ValueOf(self._estimate_cost)
	case "_compute_cost":
		return reflect.ValueOf(self._compute_cost)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() { classdb.Register("AStar2D", func(ptr gd.Object) any { return classdb.AStar2D(ptr) }) }
