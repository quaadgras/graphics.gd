package NavigationMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A navigation mesh is a collection of polygons that define which areas of an environment are traversable to aid agents in pathfinding through complicated spaces.

*/
type Go [1]classdb.NavigationMesh

/*
Based on [param value], enables or disables the specified layer in the [member geometry_collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member geometry_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Adds a polygon using the indices of the vertices you get when calling [method get_vertices].
*/
func (self Go) AddPolygon(polygon []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddPolygon(gc.PackedInt32Slice(polygon))
}

/*
Returns the number of polygons in the navigation mesh.
*/
func (self Go) GetPolygonCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPolygonCount()))
}

/*
Returns a [PackedInt32Array] containing the indices of the vertices of a created polygon.
*/
func (self Go) GetPolygon(idx int) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetPolygon(gc, gd.Int(idx)).AsSlice())
}

/*
Clears the array of polygons, but it doesn't clear the array of vertices.
*/
func (self Go) ClearPolygons() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearPolygons()
}

/*
Initializes the navigation mesh by setting the vertices and indices according to a [Mesh].
[b]Note:[/b] The given [param mesh] must be of type [constant Mesh.PRIMITIVE_TRIANGLES] and have an index array.
*/
func (self Go) CreateFromMesh(mesh gdclass.Mesh) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CreateFromMesh(mesh)
}

/*
Clears the internal arrays for vertices and polygon indices.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NavigationMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("NavigationMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Vertices() []gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector3(class(self).GetVertices(gc).AsSlice())
}

func (self Go) SetVertices(value []gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertices(gc.PackedVector3Slice(value))
}

func (self Go) SamplePartitionType() classdb.NavigationMeshSamplePartitionType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationMeshSamplePartitionType(class(self).GetSamplePartitionType())
}

func (self Go) SetSamplePartitionType(value classdb.NavigationMeshSamplePartitionType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSamplePartitionType(value)
}

func (self Go) GeometryParsedGeometryType() classdb.NavigationMeshParsedGeometryType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationMeshParsedGeometryType(class(self).GetParsedGeometryType())
}

func (self Go) SetGeometryParsedGeometryType(value classdb.NavigationMeshParsedGeometryType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParsedGeometryType(value)
}

func (self Go) GeometryCollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetGeometryCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) GeometrySourceGeometryMode() classdb.NavigationMeshSourceGeometryMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationMeshSourceGeometryMode(class(self).GetSourceGeometryMode())
}

func (self Go) SetGeometrySourceGeometryMode(value classdb.NavigationMeshSourceGeometryMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSourceGeometryMode(value)
}

func (self Go) GeometrySourceGroupName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSourceGroupName(gc).String())
}

func (self Go) SetGeometrySourceGroupName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSourceGroupName(gc.StringName(value))
}

func (self Go) CellSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCellSize()))
}

func (self Go) SetCellSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellSize(gd.Float(value))
}

func (self Go) CellHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCellHeight()))
}

func (self Go) SetCellHeight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellHeight(gd.Float(value))
}

func (self Go) BorderSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBorderSize()))
}

func (self Go) SetBorderSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderSize(gd.Float(value))
}

func (self Go) AgentHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAgentHeight()))
}

func (self Go) SetAgentHeight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAgentHeight(gd.Float(value))
}

func (self Go) AgentRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAgentRadius()))
}

func (self Go) SetAgentRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAgentRadius(gd.Float(value))
}

func (self Go) AgentMaxClimb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAgentMaxClimb()))
}

func (self Go) SetAgentMaxClimb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAgentMaxClimb(gd.Float(value))
}

func (self Go) AgentMaxSlope() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAgentMaxSlope()))
}

func (self Go) SetAgentMaxSlope(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAgentMaxSlope(gd.Float(value))
}

func (self Go) RegionMinSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRegionMinSize()))
}

func (self Go) SetRegionMinSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRegionMinSize(gd.Float(value))
}

func (self Go) RegionMergeSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRegionMergeSize()))
}

func (self Go) SetRegionMergeSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRegionMergeSize(gd.Float(value))
}

func (self Go) EdgeMaxLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEdgeMaxLength()))
}

func (self Go) SetEdgeMaxLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEdgeMaxLength(gd.Float(value))
}

func (self Go) EdgeMaxError() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEdgeMaxError()))
}

func (self Go) SetEdgeMaxError(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEdgeMaxError(gd.Float(value))
}

func (self Go) VerticesPerPolygon() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVerticesPerPolygon()))
}

func (self Go) SetVerticesPerPolygon(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVerticesPerPolygon(gd.Float(value))
}

func (self Go) DetailSampleDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDetailSampleDistance()))
}

func (self Go) SetDetailSampleDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDetailSampleDistance(gd.Float(value))
}

func (self Go) DetailSampleMaxError() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDetailSampleMaxError()))
}

func (self Go) SetDetailSampleMaxError(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDetailSampleMaxError(gd.Float(value))
}

func (self Go) FilterLowHangingObstacles() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFilterLowHangingObstacles())
}

func (self Go) SetFilterLowHangingObstacles(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilterLowHangingObstacles(value)
}

func (self Go) FilterLedgeSpans() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFilterLedgeSpans())
}

func (self Go) SetFilterLedgeSpans(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilterLedgeSpans(value)
}

func (self Go) FilterWalkableLowHeightSpans() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFilterWalkableLowHeightSpans())
}

func (self Go) SetFilterWalkableLowHeightSpans(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilterWalkableLowHeightSpans(value)
}

func (self Go) FilterBakingAabb() gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
		return gd.AABB(class(self).GetFilterBakingAabb())
}

func (self Go) SetFilterBakingAabb(value gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilterBakingAabb(value)
}

func (self Go) FilterBakingAabbOffset() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetFilterBakingAabbOffset())
}

func (self Go) SetFilterBakingAabbOffset(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilterBakingAabbOffset(value)
}

//go:nosplit
func (self class) SetSamplePartitionType(sample_partition_type classdb.NavigationMeshSamplePartitionType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sample_partition_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_sample_partition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSamplePartitionType() classdb.NavigationMeshSamplePartitionType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationMeshSamplePartitionType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_sample_partition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParsedGeometryType(geometry_type classdb.NavigationMeshParsedGeometryType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, geometry_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParsedGeometryType() classdb.NavigationMeshParsedGeometryType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationMeshParsedGeometryType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member geometry_collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member geometry_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSourceGeometryMode(mask classdb.NavigationMeshSourceGeometryMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSourceGeometryMode() classdb.NavigationMeshSourceGeometryMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationMeshSourceGeometryMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSourceGroupName(mask gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mask))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_source_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSourceGroupName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_source_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellSize(cell_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cell_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellHeight(cell_height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cell_height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_cell_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_cell_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBorderSize(border_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, border_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBorderSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAgentHeight(agent_height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent_height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_agent_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAgentHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_agent_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAgentRadius(agent_radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent_radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAgentRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAgentMaxClimb(agent_max_climb gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent_max_climb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_agent_max_climb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAgentMaxClimb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_agent_max_climb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAgentMaxSlope(agent_max_slope gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent_max_slope)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_agent_max_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAgentMaxSlope() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_agent_max_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRegionMinSize(region_min_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region_min_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_region_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRegionMinSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_region_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRegionMergeSize(region_merge_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region_merge_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_region_merge_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRegionMergeSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_region_merge_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEdgeMaxLength(edge_max_length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edge_max_length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_edge_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEdgeMaxLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_edge_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEdgeMaxError(edge_max_error gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edge_max_error)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_edge_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEdgeMaxError() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_edge_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVerticesPerPolygon(vertices_per_polygon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vertices_per_polygon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_vertices_per_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVerticesPerPolygon() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_vertices_per_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDetailSampleDistance(detail_sample_dist gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail_sample_dist)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_detail_sample_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDetailSampleDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_detail_sample_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDetailSampleMaxError(detail_sample_max_error gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail_sample_max_error)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_detail_sample_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDetailSampleMaxError() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_detail_sample_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilterLowHangingObstacles(filter_low_hanging_obstacles bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, filter_low_hanging_obstacles)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_filter_low_hanging_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilterLowHangingObstacles() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_filter_low_hanging_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilterLedgeSpans(filter_ledge_spans bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, filter_ledge_spans)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_filter_ledge_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilterLedgeSpans() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_filter_ledge_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilterWalkableLowHeightSpans(filter_walkable_low_height_spans bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, filter_walkable_low_height_spans)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_filter_walkable_low_height_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilterWalkableLowHeightSpans() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_filter_walkable_low_height_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilterBakingAabb(baking_aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, baking_aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_filter_baking_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilterBakingAabb() gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_filter_baking_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilterBakingAabbOffset(baking_aabb_offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, baking_aabb_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_filter_baking_aabb_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilterBakingAabbOffset() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_filter_baking_aabb_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the vertices that can be then indexed to create polygons with the [method add_polygon] method.
*/
//go:nosplit
func (self class) SetVertices(vertices gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [PackedVector3Array] containing all the vertices being used to create the polygons.
*/
//go:nosplit
func (self class) GetVertices(ctx gd.Lifetime) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a polygon using the indices of the vertices you get when calling [method get_vertices].
*/
//go:nosplit
func (self class) AddPolygon(polygon gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(polygon))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_add_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of polygons in the navigation mesh.
*/
//go:nosplit
func (self class) GetPolygonCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_polygon_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [PackedInt32Array] containing the indices of the vertices of a created polygon.
*/
//go:nosplit
func (self class) GetPolygon(ctx gd.Lifetime, idx gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears the array of polygons, but it doesn't clear the array of vertices.
*/
//go:nosplit
func (self class) ClearPolygons()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_clear_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Initializes the navigation mesh by setting the vertices and indices according to a [Mesh].
[b]Note:[/b] The given [param mesh] must be of type [constant Mesh.PRIMITIVE_TRIANGLES] and have an index array.
*/
//go:nosplit
func (self class) CreateFromMesh(mesh gdclass.Mesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_create_from_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the internal arrays for vertices and polygon indices.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMesh.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsNavigationMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNavigationMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("NavigationMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type SamplePartitionType = classdb.NavigationMeshSamplePartitionType

const (
/*Watershed partitioning. Generally the best choice if you precompute the navigation mesh, use this if you have large open areas.*/
	SamplePartitionWatershed SamplePartitionType = 0
/*Monotone partitioning. Use this if you want fast navigation mesh generation.*/
	SamplePartitionMonotone SamplePartitionType = 1
/*Layer partitioning. Good choice to use for tiled navigation mesh with medium and small sized tiles.*/
	SamplePartitionLayers SamplePartitionType = 2
/*Represents the size of the [enum SamplePartitionType] enum.*/
	SamplePartitionMax SamplePartitionType = 3
)
type ParsedGeometryType = classdb.NavigationMeshParsedGeometryType

const (
/*Parses mesh instances as geometry. This includes [MeshInstance3D], [CSGShape3D], and [GridMap] nodes.*/
	ParsedGeometryMeshInstances ParsedGeometryType = 0
/*Parses [StaticBody3D] colliders as geometry. The collider should be in any of the layers specified by [member geometry_collision_mask].*/
	ParsedGeometryStaticColliders ParsedGeometryType = 1
/*Both [constant PARSED_GEOMETRY_MESH_INSTANCES] and [constant PARSED_GEOMETRY_STATIC_COLLIDERS].*/
	ParsedGeometryBoth ParsedGeometryType = 2
/*Represents the size of the [enum ParsedGeometryType] enum.*/
	ParsedGeometryMax ParsedGeometryType = 3
)
type SourceGeometryMode = classdb.NavigationMeshSourceGeometryMode

const (
/*Scans the child nodes of the root node recursively for geometry.*/
	SourceGeometryRootNodeChildren SourceGeometryMode = 0
/*Scans nodes in a group and their child nodes recursively for geometry. The group is specified by [member geometry_source_group_name].*/
	SourceGeometryGroupsWithChildren SourceGeometryMode = 1
/*Uses nodes in a group for geometry. The group is specified by [member geometry_source_group_name].*/
	SourceGeometryGroupsExplicit SourceGeometryMode = 2
/*Represents the size of the [enum SourceGeometryMode] enum.*/
	SourceGeometryMax SourceGeometryMode = 3
)
