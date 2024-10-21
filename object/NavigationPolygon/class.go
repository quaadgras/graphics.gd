package NavigationPolygon

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A navigation mesh can be created either by baking it with the help of the [NavigationServer2D], or by adding vertices and convex polygon indices arrays manually.
To bake a navigation mesh at least one outline needs to be added that defines the outer bounds of the baked area.
[codeblocks]
[gdscript]
var new_navigation_mesh = NavigationPolygon.new()
var bounding_outline = PackedVector2Array([Vector2(0, 0), Vector2(0, 50), Vector2(50, 50), Vector2(50, 0)])
new_navigation_mesh.add_outline(bounding_outline)
NavigationServer2D.bake_from_source_geometry_data(new_navigation_mesh, NavigationMeshSourceGeometryData2D.new());
$NavigationRegion2D.navigation_polygon = new_navigation_mesh
[/gdscript]
[csharp]
var newNavigationMesh = new NavigationPolygon();
var boundingOutline = new Vector2[] { new Vector2(0, 0), new Vector2(0, 50), new Vector2(50, 50), new Vector2(50, 0) };
newNavigationMesh.AddOutline(boundingOutline);
NavigationServer2D.BakeFromSourceGeometryData(newNavigationMesh, new NavigationMeshSourceGeometryData2D());
GetNode<NavigationRegion2D>("NavigationRegion2D").NavigationPolygon = newNavigationMesh;
[/csharp]
[/codeblocks]
Adding vertices and polygon indices manually.
[codeblocks]
[gdscript]
var new_navigation_mesh = NavigationPolygon.new()
var new_vertices = PackedVector2Array([Vector2(0, 0), Vector2(0, 50), Vector2(50, 50), Vector2(50, 0)])
new_navigation_mesh.vertices = new_vertices
var new_polygon_indices = PackedInt32Array([0, 1, 2, 3])
new_navigation_mesh.add_polygon(new_polygon_indices)
$NavigationRegion2D.navigation_polygon = new_navigation_mesh
[/gdscript]
[csharp]
var newNavigationMesh = new NavigationPolygon();
var newVertices = new Vector2[] { new Vector2(0, 0), new Vector2(0, 50), new Vector2(50, 50), new Vector2(50, 0) };
newNavigationMesh.Vertices = newVertices;
var newPolygonIndices = new int[] { 0, 1, 2, 3 };
newNavigationMesh.AddPolygon(newPolygonIndices);
GetNode<NavigationRegion2D>("NavigationRegion2D").NavigationPolygon = newNavigationMesh;
[/csharp]
[/codeblocks]

*/
type Simple [1]classdb.NavigationPolygon
func (self Simple) SetVertices(vertices gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVertices(vertices)
}
func (self Simple) GetVertices() gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetVertices(gc))
}
func (self Simple) AddPolygon(polygon gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPolygon(polygon)
}
func (self Simple) GetPolygonCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPolygonCount()))
}
func (self Simple) GetPolygon(idx int) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetPolygon(gc, gd.Int(idx)))
}
func (self Simple) ClearPolygons() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearPolygons()
}
func (self Simple) GetNavigationMesh() [1]classdb.NavigationMesh {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.NavigationMesh(Expert(self).GetNavigationMesh(gc))
}
func (self Simple) AddOutline(outline gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddOutline(outline)
}
func (self Simple) AddOutlineAtIndex(outline gd.PackedVector2Array, index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddOutlineAtIndex(outline, gd.Int(index))
}
func (self Simple) GetOutlineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutlineCount()))
}
func (self Simple) SetOutline(idx int, outline gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutline(gd.Int(idx), outline)
}
func (self Simple) GetOutline(idx int) gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetOutline(gc, gd.Int(idx)))
}
func (self Simple) RemoveOutline(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveOutline(gd.Int(idx))
}
func (self Simple) ClearOutlines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearOutlines()
}
func (self Simple) MakePolygonsFromOutlines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MakePolygonsFromOutlines()
}
func (self Simple) SetCellSize(cell_size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellSize(gd.Float(cell_size))
}
func (self Simple) GetCellSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCellSize()))
}
func (self Simple) SetBorderSize(border_size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBorderSize(gd.Float(border_size))
}
func (self Simple) GetBorderSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBorderSize()))
}
func (self Simple) SetParsedGeometryType(geometry_type classdb.NavigationPolygonParsedGeometryType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParsedGeometryType(geometry_type)
}
func (self Simple) GetParsedGeometryType() classdb.NavigationPolygonParsedGeometryType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.NavigationPolygonParsedGeometryType(Expert(self).GetParsedGeometryType())
}
func (self Simple) SetParsedCollisionMask(mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParsedCollisionMask(gd.Int(mask))
}
func (self Simple) GetParsedCollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetParsedCollisionMask()))
}
func (self Simple) SetParsedCollisionMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParsedCollisionMaskValue(gd.Int(layer_number), value)
}
func (self Simple) GetParsedCollisionMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetParsedCollisionMaskValue(gd.Int(layer_number)))
}
func (self Simple) SetSourceGeometryMode(geometry_mode classdb.NavigationPolygonSourceGeometryMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSourceGeometryMode(geometry_mode)
}
func (self Simple) GetSourceGeometryMode() classdb.NavigationPolygonSourceGeometryMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.NavigationPolygonSourceGeometryMode(Expert(self).GetSourceGeometryMode())
}
func (self Simple) SetSourceGeometryGroupName(group_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSourceGeometryGroupName(gc.StringName(group_name))
}
func (self Simple) GetSourceGeometryGroupName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSourceGeometryGroupName(gc).String())
}
func (self Simple) SetAgentRadius(agent_radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAgentRadius(gd.Float(agent_radius))
}
func (self Simple) GetAgentRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAgentRadius()))
}
func (self Simple) SetBakingRect(rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBakingRect(rect)
}
func (self Simple) GetBakingRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetBakingRect())
}
func (self Simple) SetBakingRectOffset(rect_offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBakingRectOffset(rect_offset)
}
func (self Simple) GetBakingRectOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetBakingRectOffset())
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationPolygon
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the vertices that can be then indexed to create polygons with the [method add_polygon] method.
*/
//go:nosplit
func (self class) SetVertices(vertices gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [PackedVector2Array] containing all the vertices being used to create the polygons.
*/
//go:nosplit
func (self class) GetVertices(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_add_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the count of all polygons.
*/
//go:nosplit
func (self class) GetPolygonCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_polygon_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears the array of polygons, but it doesn't clear the array of outlines and vertices.
*/
//go:nosplit
func (self class) ClearPolygons()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_clear_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [NavigationMesh] resulting from this navigation polygon. This navigation mesh can be used to update the navigation mesh of a region with the [method NavigationServer3D.region_set_navigation_mesh] API directly (as 2D uses the 3D server behind the scene).
*/
//go:nosplit
func (self class) GetNavigationMesh(ctx gd.Lifetime) object.NavigationMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.NavigationMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Appends a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines.
*/
//go:nosplit
func (self class) AddOutline(outline gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(outline))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_add_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines at a fixed position.
*/
//go:nosplit
func (self class) AddOutlineAtIndex(outline gd.PackedVector2Array, index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(outline))
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_add_outline_at_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of outlines that were created in the editor or by script.
*/
//go:nosplit
func (self class) GetOutlineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_outline_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Changes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
//go:nosplit
func (self class) SetOutline(idx gd.Int, outline gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(outline))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [PackedVector2Array] containing the vertices of an outline that was created in the editor or by script.
*/
//go:nosplit
func (self class) GetOutline(ctx gd.Lifetime, idx gd.Int) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
//go:nosplit
func (self class) RemoveOutline(idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_remove_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the array of the outlines, but it doesn't clear the vertices and the polygons that were created by them.
*/
//go:nosplit
func (self class) ClearOutlines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_clear_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates polygons from the outlines added in the editor or by script.
*/
//go:nosplit
func (self class) MakePolygonsFromOutlines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_make_polygons_from_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCellSize(cell_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cell_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBorderSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParsedGeometryType(geometry_type classdb.NavigationPolygonParsedGeometryType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, geometry_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParsedGeometryType() classdb.NavigationPolygonParsedGeometryType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPolygonParsedGeometryType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParsedCollisionMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_parsed_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParsedCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_parsed_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member parsed_collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetParsedCollisionMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_parsed_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member parsed_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetParsedCollisionMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_parsed_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSourceGeometryMode(geometry_mode classdb.NavigationPolygonSourceGeometryMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, geometry_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSourceGeometryMode() classdb.NavigationPolygonSourceGeometryMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPolygonSourceGeometryMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSourceGeometryGroupName(group_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(group_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_source_geometry_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSourceGeometryGroupName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_source_geometry_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAgentRadius(agent_radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent_radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAgentRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakingRect(rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_baking_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakingRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_baking_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakingRectOffset(rect_offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_set_baking_rect_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakingRectOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_baking_rect_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears the internal arrays for vertices and polygon indices.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsNavigationPolygon() Expert { return self[0].AsNavigationPolygon() }


//go:nosplit
func (self Simple) AsNavigationPolygon() Simple { return self[0].AsNavigationPolygon() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


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
func init() {classdb.Register("NavigationPolygon", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ParsedGeometryType = classdb.NavigationPolygonParsedGeometryType

const (
/*Parses mesh instances as obstruction geometry. This includes [Polygon2D], [MeshInstance2D], [MultiMeshInstance2D], and [TileMap] nodes.
Meshes are only parsed when they use a 2D vertices surface format.*/
	ParsedGeometryMeshInstances ParsedGeometryType = 0
/*Parses [StaticBody2D] and [TileMap] colliders as obstruction geometry. The collider should be in any of the layers specified by [member parsed_collision_mask].*/
	ParsedGeometryStaticColliders ParsedGeometryType = 1
/*Both [constant PARSED_GEOMETRY_MESH_INSTANCES] and [constant PARSED_GEOMETRY_STATIC_COLLIDERS].*/
	ParsedGeometryBoth ParsedGeometryType = 2
/*Represents the size of the [enum ParsedGeometryType] enum.*/
	ParsedGeometryMax ParsedGeometryType = 3
)
type SourceGeometryMode = classdb.NavigationPolygonSourceGeometryMode

const (
/*Scans the child nodes of the root node recursively for geometry.*/
	SourceGeometryRootNodeChildren SourceGeometryMode = 0
/*Scans nodes in a group and their child nodes recursively for geometry. The group is specified by [member source_geometry_group_name].*/
	SourceGeometryGroupsWithChildren SourceGeometryMode = 1
/*Uses nodes in a group for geometry. The group is specified by [member source_geometry_group_name].*/
	SourceGeometryGroupsExplicit SourceGeometryMode = 2
/*Represents the size of the [enum SourceGeometryMode] enum.*/
	SourceGeometryMax SourceGeometryMode = 3
)
