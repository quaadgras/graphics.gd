// Package NavigationPolygon provides methods for working with NavigationPolygon object instances.
package NavigationPolygon

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
type Instance [1]gdclass.NavigationPolygon
type Any interface {
	gd.IsClass
	AsNavigationPolygon() Instance
}

/*
Adds a polygon using the indices of the vertices you get when calling [method get_vertices].
*/
func (self Instance) AddPolygon(polygon []int32) {
	class(self).AddPolygon(gd.NewPackedInt32Slice(polygon))
}

/*
Returns the count of all polygons.
*/
func (self Instance) GetPolygonCount() int {
	return int(int(class(self).GetPolygonCount()))
}

/*
Returns a [PackedInt32Array] containing the indices of the vertices of a created polygon.
*/
func (self Instance) GetPolygon(idx int) []int32 {
	return []int32(class(self).GetPolygon(gd.Int(idx)).AsSlice())
}

/*
Clears the array of polygons, but it doesn't clear the array of outlines and vertices.
*/
func (self Instance) ClearPolygons() {
	class(self).ClearPolygons()
}

/*
Returns the [NavigationMesh] resulting from this navigation polygon. This navigation mesh can be used to update the navigation mesh of a region with the [method NavigationServer3D.region_set_navigation_mesh] API directly (as 2D uses the 3D server behind the scene).
*/
func (self Instance) GetNavigationMesh() [1]gdclass.NavigationMesh {
	return [1]gdclass.NavigationMesh(class(self).GetNavigationMesh())
}

/*
Appends a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines.
*/
func (self Instance) AddOutline(outline []Vector2.XY) {
	class(self).AddOutline(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&outline))))
}

/*
Adds a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines at a fixed position.
*/
func (self Instance) AddOutlineAtIndex(outline []Vector2.XY, index int) {
	class(self).AddOutlineAtIndex(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&outline))), gd.Int(index))
}

/*
Returns the number of outlines that were created in the editor or by script.
*/
func (self Instance) GetOutlineCount() int {
	return int(int(class(self).GetOutlineCount()))
}

/*
Changes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
func (self Instance) SetOutline(idx int, outline []Vector2.XY) {
	class(self).SetOutline(gd.Int(idx), gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&outline))))
}

/*
Returns a [PackedVector2Array] containing the vertices of an outline that was created in the editor or by script.
*/
func (self Instance) GetOutline(idx int) []Vector2.XY {
	return []Vector2.XY(class(self).GetOutline(gd.Int(idx)).AsSlice())
}

/*
Removes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
func (self Instance) RemoveOutline(idx int) {
	class(self).RemoveOutline(gd.Int(idx))
}

/*
Clears the array of the outlines, but it doesn't clear the vertices and the polygons that were created by them.
*/
func (self Instance) ClearOutlines() {
	class(self).ClearOutlines()
}

/*
Creates polygons from the outlines added in the editor or by script.
*/
func (self Instance) MakePolygonsFromOutlines() {
	class(self).MakePolygonsFromOutlines()
}

/*
Based on [param value], enables or disables the specified layer in the [member parsed_collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetParsedCollisionMaskValue(layer_number int, value bool) {
	class(self).SetParsedCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member parsed_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetParsedCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetParsedCollisionMaskValue(gd.Int(layer_number)))
}

/*
Clears the internal arrays for vertices and polygon indices.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationPolygon

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationPolygon"))
	casted := Instance{*(*gdclass.NavigationPolygon)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Vertices() []Vector2.XY {
	return []Vector2.XY(class(self).GetVertices().AsSlice())
}

func (self Instance) SetVertices(value []Vector2.XY) {
	class(self).SetVertices(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

func (self Instance) ParsedGeometryType() gdclass.NavigationPolygonParsedGeometryType {
	return gdclass.NavigationPolygonParsedGeometryType(class(self).GetParsedGeometryType())
}

func (self Instance) SetParsedGeometryType(value gdclass.NavigationPolygonParsedGeometryType) {
	class(self).SetParsedGeometryType(value)
}

func (self Instance) ParsedCollisionMask() int {
	return int(int(class(self).GetParsedCollisionMask()))
}

func (self Instance) SetParsedCollisionMask(value int) {
	class(self).SetParsedCollisionMask(gd.Int(value))
}

func (self Instance) SourceGeometryMode() gdclass.NavigationPolygonSourceGeometryMode {
	return gdclass.NavigationPolygonSourceGeometryMode(class(self).GetSourceGeometryMode())
}

func (self Instance) SetSourceGeometryMode(value gdclass.NavigationPolygonSourceGeometryMode) {
	class(self).SetSourceGeometryMode(value)
}

func (self Instance) SourceGeometryGroupName() string {
	return string(class(self).GetSourceGeometryGroupName().String())
}

func (self Instance) SetSourceGeometryGroupName(value string) {
	class(self).SetSourceGeometryGroupName(gd.NewStringName(value))
}

func (self Instance) CellSize() Float.X {
	return Float.X(Float.X(class(self).GetCellSize()))
}

func (self Instance) SetCellSize(value Float.X) {
	class(self).SetCellSize(gd.Float(value))
}

func (self Instance) BorderSize() Float.X {
	return Float.X(Float.X(class(self).GetBorderSize()))
}

func (self Instance) SetBorderSize(value Float.X) {
	class(self).SetBorderSize(gd.Float(value))
}

func (self Instance) AgentRadius() Float.X {
	return Float.X(Float.X(class(self).GetAgentRadius()))
}

func (self Instance) SetAgentRadius(value Float.X) {
	class(self).SetAgentRadius(gd.Float(value))
}

func (self Instance) BakingRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetBakingRect())
}

func (self Instance) SetBakingRect(value Rect2.PositionSize) {
	class(self).SetBakingRect(gd.Rect2(value))
}

func (self Instance) BakingRectOffset() Vector2.XY {
	return Vector2.XY(class(self).GetBakingRectOffset())
}

func (self Instance) SetBakingRectOffset(value Vector2.XY) {
	class(self).SetBakingRectOffset(gd.Vector2(value))
}

/*
Sets the vertices that can be then indexed to create polygons with the [method add_polygon] method.
*/
//go:nosplit
func (self class) SetVertices(vertices gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [PackedVector2Array] containing all the vertices being used to create the polygons.
*/
//go:nosplit
func (self class) GetVertices() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a polygon using the indices of the vertices you get when calling [method get_vertices].
*/
//go:nosplit
func (self class) AddPolygon(polygon gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_add_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the count of all polygons.
*/
//go:nosplit
func (self class) GetPolygonCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_polygon_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [PackedInt32Array] containing the indices of the vertices of a created polygon.
*/
//go:nosplit
func (self class) GetPolygon(idx gd.Int) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the array of polygons, but it doesn't clear the array of outlines and vertices.
*/
//go:nosplit
func (self class) ClearPolygons() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_clear_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [NavigationMesh] resulting from this navigation polygon. This navigation mesh can be used to update the navigation mesh of a region with the [method NavigationServer3D.region_set_navigation_mesh] API directly (as 2D uses the 3D server behind the scene).
*/
//go:nosplit
func (self class) GetNavigationMesh() [1]gdclass.NavigationMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.NavigationMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.NavigationMesh](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Appends a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines.
*/
//go:nosplit
func (self class) AddOutline(outline gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(outline))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_add_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a [PackedVector2Array] that contains the vertices of an outline to the internal array that contains all the outlines at a fixed position.
*/
//go:nosplit
func (self class) AddOutlineAtIndex(outline gd.PackedVector2Array, index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(outline))
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_add_outline_at_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of outlines that were created in the editor or by script.
*/
//go:nosplit
func (self class) GetOutlineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_outline_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Changes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
//go:nosplit
func (self class) SetOutline(idx gd.Int, outline gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(outline))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [PackedVector2Array] containing the vertices of an outline that was created in the editor or by script.
*/
//go:nosplit
func (self class) GetOutline(idx gd.Int) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes an outline created in the editor or by script. You have to call [method make_polygons_from_outlines] for the polygons to update.
*/
//go:nosplit
func (self class) RemoveOutline(idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_remove_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears the array of the outlines, but it doesn't clear the vertices and the polygons that were created by them.
*/
//go:nosplit
func (self class) ClearOutlines() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_clear_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Creates polygons from the outlines added in the editor or by script.
*/
//go:nosplit
func (self class) MakePolygonsFromOutlines() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_make_polygons_from_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCellSize(cell_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, cell_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderSize(border_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, border_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParsedGeometryType(geometry_type gdclass.NavigationPolygonParsedGeometryType) {
	var frame = callframe.New()
	callframe.Arg(frame, geometry_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetParsedGeometryType() gdclass.NavigationPolygonParsedGeometryType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPolygonParsedGeometryType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParsedCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_parsed_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetParsedCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_parsed_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member parsed_collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetParsedCollisionMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_parsed_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member parsed_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetParsedCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_parsed_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSourceGeometryMode(geometry_mode gdclass.NavigationPolygonSourceGeometryMode) {
	var frame = callframe.New()
	callframe.Arg(frame, geometry_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSourceGeometryMode() gdclass.NavigationPolygonSourceGeometryMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPolygonSourceGeometryMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSourceGeometryGroupName(group_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(group_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_source_geometry_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSourceGeometryGroupName() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_source_geometry_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAgentRadius(agent_radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, agent_radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAgentRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakingRect(rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_baking_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakingRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_baking_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakingRectOffset(rect_offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, rect_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_set_baking_rect_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakingRectOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_get_baking_rect_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears the internal arrays for vertices and polygon indices.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPolygon.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsNavigationPolygon() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNavigationPolygon() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("NavigationPolygon", func(ptr gd.Object) any {
		return [1]gdclass.NavigationPolygon{*(*gdclass.NavigationPolygon)(unsafe.Pointer(&ptr))}
	})
}

type ParsedGeometryType = gdclass.NavigationPolygonParsedGeometryType

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

type SourceGeometryMode = gdclass.NavigationPolygonSourceGeometryMode

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
