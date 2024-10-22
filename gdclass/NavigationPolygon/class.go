package NavigationPolygon

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
type Go [1]classdb.NavigationPolygon

/*
Adds a polygon using the indices of the vertices you get when calling [method get_vertices].
*/
func (self Go) AddPolygon(polygon []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddPolygon(gc.PackedInt32Slice(polygon))
}

/*
Returns the count of all polygons.
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
Clears the array of polygons, but it doesn't clear the array of outlines and vertices.
*/
func (self Go) ClearPolygons() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearPolygons()
}

/*
Returns the [NavigationMesh] resulting from this navigation polygon. This navigation mesh can be used to update the navigation mesh of a region with the [method NavigationServer3D.region_set_navigation_mesh] API directly (as 2D uses the 3D server behind the scene).
*/
func (self Go) GetNavigationMesh() gdclass.NavigationMesh {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.NavigationMesh(class(self).GetNavigationMesh(gc))
}

/*
Appends a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines.
*/
func (self Go) AddOutline(outline []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddOutline(gc.PackedVector2Slice(outline))
}

/*
Adds a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines at a fixed position.
*/
func (self Go) AddOutlineAtIndex(outline []gd.Vector2, index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddOutlineAtIndex(gc.PackedVector2Slice(outline), gd.Int(index))
}

/*
Returns the number of outlines that were created in the editor or by script.
*/
func (self Go) GetOutlineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetOutlineCount()))
}

/*
Changes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
func (self Go) SetOutline(idx int, outline []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOutline(gd.Int(idx), gc.PackedVector2Slice(outline))
}

/*
Returns a [PackedVector2Array] containing the vertices of an outline that was created in the editor or by script.
*/
func (self Go) GetOutline(idx int) []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector2(class(self).GetOutline(gc, gd.Int(idx)).AsSlice())
}

/*
Removes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
func (self Go) RemoveOutline(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveOutline(gd.Int(idx))
}

/*
Clears the array of the outlines, but it doesn't clear the vertices and the polygons that were created by them.
*/
func (self Go) ClearOutlines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearOutlines()
}

/*
Creates polygons from the outlines added in the editor or by script.
*/
func (self Go) MakePolygonsFromOutlines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MakePolygonsFromOutlines()
}

/*
Based on [param value], enables or disables the specified layer in the [member parsed_collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetParsedCollisionMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParsedCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member parsed_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetParsedCollisionMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetParsedCollisionMaskValue(gd.Int(layer_number)))
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
type class [1]classdb.NavigationPolygon
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("NavigationPolygon"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Vertices() []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector2(class(self).GetVertices(gc).AsSlice())
}

func (self Go) SetVertices(value []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertices(gc.PackedVector2Slice(value))
}

func (self Go) ParsedGeometryType() classdb.NavigationPolygonParsedGeometryType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPolygonParsedGeometryType(class(self).GetParsedGeometryType())
}

func (self Go) SetParsedGeometryType(value classdb.NavigationPolygonParsedGeometryType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParsedGeometryType(value)
}

func (self Go) ParsedCollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetParsedCollisionMask()))
}

func (self Go) SetParsedCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParsedCollisionMask(gd.Int(value))
}

func (self Go) SourceGeometryMode() classdb.NavigationPolygonSourceGeometryMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPolygonSourceGeometryMode(class(self).GetSourceGeometryMode())
}

func (self Go) SetSourceGeometryMode(value classdb.NavigationPolygonSourceGeometryMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSourceGeometryMode(value)
}

func (self Go) SourceGeometryGroupName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSourceGeometryGroupName(gc).String())
}

func (self Go) SetSourceGeometryGroupName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSourceGeometryGroupName(gc.StringName(value))
}

func (self Go) CellSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCellSize()))
}

func (self Go) SetCellSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellSize(gd.Float(value))
}

func (self Go) BorderSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBorderSize()))
}

func (self Go) SetBorderSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderSize(gd.Float(value))
}

func (self Go) AgentRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAgentRadius()))
}

func (self Go) SetAgentRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAgentRadius(gd.Float(value))
}

func (self Go) BakingRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Rect2(class(self).GetBakingRect())
}

func (self Go) SetBakingRect(value gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakingRect(value)
}

func (self Go) BakingRectOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetBakingRectOffset())
}

func (self Go) SetBakingRectOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakingRectOffset(value)
}

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
func (self class) GetNavigationMesh(ctx gd.Lifetime) gdclass.NavigationMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPolygon.Bind_get_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.NavigationMesh
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
func (self class) AsNavigationPolygon() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNavigationPolygon() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("NavigationPolygon", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
