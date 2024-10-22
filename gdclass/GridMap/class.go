package GridMap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
GridMap lets you place meshes on a grid interactively. It works both from the editor and from scripts, which can help you create in-game level editors.
GridMaps use a [MeshLibrary] which contains a list of tiles. Each tile is a mesh with materials plus optional collision and navigation shapes.
A GridMap contains a collection of cells. Each grid cell refers to a tile in the [MeshLibrary]. All cells in the map have the same dimensions.
Internally, a GridMap is split into a sparse collection of octants for efficient rendering and physics processing. Every octant has the same dimensions and can contain several cells.
[b]Note:[/b] GridMap doesn't extend [VisualInstance3D] and therefore can't be hidden or cull masked based on [member VisualInstance3D.layers]. If you make a light not affect the first layer, the whole GridMap won't be lit by the light in question.

*/
type Go [1]classdb.GridMap

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionLayerValue(gd.Int(layer_number)))
}

/*
Sets the [RID] of the navigation map this GridMap node should use for its cell baked navigation meshes.
*/
func (self Go) SetNavigationMap(navigation_map gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationMap(navigation_map)
}

/*
Returns the [RID] of the navigation map this GridMap node uses for its cell baked navigation meshes.
This function returns always the map set on the GridMap node and not the map on the NavigationServer. If the map is changed directly with the NavigationServer API the GridMap node will not be aware of the map change.
*/
func (self Go) GetNavigationMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetNavigationMap())
}

/*
Sets the mesh index for the cell referenced by its grid coordinates.
A negative item index such as [constant INVALID_CELL_ITEM] will clear the cell.
Optionally, the item's orientation can be passed. For valid orientation values, see [method get_orthogonal_index_from_basis].
*/
func (self Go) SetCellItem(position gd.Vector3i, item int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellItem(position, gd.Int(item), gd.Int(0))
}

/*
The [MeshLibrary] item index located at the given grid coordinates. If the cell is empty, [constant INVALID_CELL_ITEM] will be returned.
*/
func (self Go) GetCellItem(position gd.Vector3i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCellItem(position)))
}

/*
The orientation of the cell at the given grid coordinates. [code]-1[/code] is returned if the cell is empty.
*/
func (self Go) GetCellItemOrientation(position gd.Vector3i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCellItemOrientation(position)))
}

/*
Returns the basis that gives the specified cell its orientation.
*/
func (self Go) GetCellItemBasis(position gd.Vector3i) gd.Basis {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Basis(class(self).GetCellItemBasis(position))
}

/*
Returns one of 24 possible rotations that lie along the vectors (x,y,z) with each component being either -1, 0, or 1. For further details, refer to the Godot source code.
*/
func (self Go) GetBasisWithOrthogonalIndex(index int) gd.Basis {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Basis(class(self).GetBasisWithOrthogonalIndex(gd.Int(index)))
}

/*
This function considers a discretization of rotations into 24 points on unit sphere, lying along the vectors (x,y,z) with each component being either -1, 0, or 1, and returns the index (in the range from 0 to 23) of the point best representing the orientation of the object. For further details, refer to the Godot source code.
*/
func (self Go) GetOrthogonalIndexFromBasis(basis gd.Basis) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetOrthogonalIndexFromBasis(basis)))
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node3D.to_local] before passing it to this method. See also [method map_to_local].
*/
func (self Go) LocalToMap(local_position gd.Vector3) gd.Vector3i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3i(class(self).LocalToMap(local_position))
}

/*
Returns the position of a grid cell in the GridMap's local coordinate space. To convert the returned value into global coordinates, use [method Node3D.to_global]. See also [method local_to_map].
*/
func (self Go) MapToLocal(map_position gd.Vector3i) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).MapToLocal(map_position))
}

/*
This method does nothing.
*/
func (self Go) ResourceChanged(resource gdclass.Resource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ResourceChanged(resource)
}

/*
Clear all cells.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Returns an array of [Vector3] with the non-empty cell coordinates in the grid map.
*/
func (self Go) GetUsedCells() gd.ArrayOf[gd.Vector3i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector3i](class(self).GetUsedCells(gc))
}

/*
Returns an array of all cells with the given item index specified in [param item].
*/
func (self Go) GetUsedCellsByItem(item int) gd.ArrayOf[gd.Vector3i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector3i](class(self).GetUsedCellsByItem(gc, gd.Int(item)))
}

/*
Returns an array of [Transform3D] and [Mesh] references corresponding to the non-empty cells in the grid. The transforms are specified in local space.
*/
func (self Go) GetMeshes() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetMeshes(gc))
}

/*
Returns an array of [ArrayMesh]es and [Transform3D] references of all bake meshes that exist within the current GridMap.
*/
func (self Go) GetBakeMeshes() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetBakeMeshes(gc))
}

/*
Returns [RID] of a baked mesh with the given [param idx].
*/
func (self Go) GetBakeMeshInstance(idx int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetBakeMeshInstance(gd.Int(idx)))
}

/*
Clears all baked meshes. See [method make_baked_meshes].
*/
func (self Go) ClearBakedMeshes() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearBakedMeshes()
}

/*
Bakes lightmap data for all meshes in the assigned [MeshLibrary].
*/
func (self Go) MakeBakedMeshes() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MakeBakedMeshes(false, gd.Float(0.1))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GridMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GridMap"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MeshLibrary() gdclass.MeshLibrary {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.MeshLibrary(class(self).GetMeshLibrary(gc))
}

func (self Go) SetMeshLibrary(value gdclass.MeshLibrary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMeshLibrary(value)
}

func (self Go) PhysicsMaterial() gdclass.PhysicsMaterial {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.PhysicsMaterial(class(self).GetPhysicsMaterial(gc))
}

func (self Go) SetPhysicsMaterial(value gdclass.PhysicsMaterial) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicsMaterial(value)
}

func (self Go) CellSize() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetCellSize())
}

func (self Go) SetCellSize(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellSize(value)
}

func (self Go) CellOctantSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetOctantSize()))
}

func (self Go) SetCellOctantSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOctantSize(gd.Int(value))
}

func (self Go) CellCenterX() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetCenterX())
}

func (self Go) SetCellCenterX(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCenterX(value)
}

func (self Go) CellCenterY() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetCenterY())
}

func (self Go) SetCellCenterY(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCenterY(value)
}

func (self Go) CellCenterZ() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetCenterZ())
}

func (self Go) SetCellCenterZ(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCenterZ(value)
}

func (self Go) CellScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCellScale()))
}

func (self Go) SetCellScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellScale(gd.Float(value))
}

func (self Go) CollisionLayer() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionLayer()))
}

func (self Go) SetCollisionLayer(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionLayer(gd.Int(value))
}

func (self Go) CollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) CollisionPriority() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCollisionPriority()))
}

func (self Go) SetCollisionPriority(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionPriority(gd.Float(value))
}

func (self Go) BakeNavigation() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsBakingNavigation())
}

func (self Go) SetBakeNavigation(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakeNavigation(value)
}

//go:nosplit
func (self class) SetCollisionLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionPriority(priority gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionPriority() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicsMaterial(material gdclass.PhysicsMaterial)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_physics_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsMaterial(ctx gd.Lifetime) gdclass.PhysicsMaterial {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_physics_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PhysicsMaterial
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakeNavigation(bake_navigation bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bake_navigation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_bake_navigation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsBakingNavigation() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_is_baking_navigation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [RID] of the navigation map this GridMap node should use for its cell baked navigation meshes.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the navigation map this GridMap node uses for its cell baked navigation meshes.
This function returns always the map set on the GridMap node and not the map on the NavigationServer. If the map is changed directly with the NavigationServer API the GridMap node will not be aware of the map change.
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMeshLibrary(mesh_library gdclass.MeshLibrary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh_library[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_mesh_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMeshLibrary(ctx gd.Lifetime) gdclass.MeshLibrary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_mesh_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.MeshLibrary
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellSize(size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellSize() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_cell_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_cell_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOctantSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_octant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOctantSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_octant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the mesh index for the cell referenced by its grid coordinates.
A negative item index such as [constant INVALID_CELL_ITEM] will clear the cell.
Optionally, the item's orientation can be passed. For valid orientation values, see [method get_orthogonal_index_from_basis].
*/
//go:nosplit
func (self class) SetCellItem(position gd.Vector3i, item gd.Int, orientation gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, item)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_cell_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The [MeshLibrary] item index located at the given grid coordinates. If the cell is empty, [constant INVALID_CELL_ITEM] will be returned.
*/
//go:nosplit
func (self class) GetCellItem(position gd.Vector3i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_cell_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The orientation of the cell at the given grid coordinates. [code]-1[/code] is returned if the cell is empty.
*/
//go:nosplit
func (self class) GetCellItemOrientation(position gd.Vector3i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_cell_item_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the basis that gives the specified cell its orientation.
*/
//go:nosplit
func (self class) GetCellItemBasis(position gd.Vector3i) gd.Basis {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Basis](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_cell_item_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns one of 24 possible rotations that lie along the vectors (x,y,z) with each component being either -1, 0, or 1. For further details, refer to the Godot source code.
*/
//go:nosplit
func (self class) GetBasisWithOrthogonalIndex(index gd.Int) gd.Basis {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Basis](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_basis_with_orthogonal_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
This function considers a discretization of rotations into 24 points on unit sphere, lying along the vectors (x,y,z) with each component being either -1, 0, or 1, and returns the index (in the range from 0 to 23) of the point best representing the orientation of the object. For further details, refer to the Godot source code.
*/
//go:nosplit
func (self class) GetOrthogonalIndexFromBasis(basis gd.Basis) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, basis)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_orthogonal_index_from_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node3D.to_local] before passing it to this method. See also [method map_to_local].
*/
//go:nosplit
func (self class) LocalToMap(local_position gd.Vector3) gd.Vector3i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[gd.Vector3i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_local_to_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position of a grid cell in the GridMap's local coordinate space. To convert the returned value into global coordinates, use [method Node3D.to_global]. See also [method local_to_map].
*/
//go:nosplit
func (self class) MapToLocal(map_position gd.Vector3i) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, map_position)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_map_to_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
This method does nothing.
*/
//go:nosplit
func (self class) ResourceChanged(resource gdclass.Resource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_resource_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCenterX(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_center_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterX() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_center_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterY(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_center_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterY() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_center_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterZ(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_set_center_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterZ() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_center_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clear all cells.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of [Vector3] with the non-empty cell coordinates in the grid map.
*/
//go:nosplit
func (self class) GetUsedCells(ctx gd.Lifetime) gd.ArrayOf[gd.Vector3i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector3i](ret)
}
/*
Returns an array of all cells with the given item index specified in [param item].
*/
//go:nosplit
func (self class) GetUsedCellsByItem(ctx gd.Lifetime, item gd.Int) gd.ArrayOf[gd.Vector3i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_used_cells_by_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector3i](ret)
}
/*
Returns an array of [Transform3D] and [Mesh] references corresponding to the non-empty cells in the grid. The transforms are specified in local space.
*/
//go:nosplit
func (self class) GetMeshes(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of [ArrayMesh]es and [Transform3D] references of all bake meshes that exist within the current GridMap.
*/
//go:nosplit
func (self class) GetBakeMeshes(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_bake_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [RID] of a baked mesh with the given [param idx].
*/
//go:nosplit
func (self class) GetBakeMeshInstance(idx gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_get_bake_mesh_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears all baked meshes. See [method make_baked_meshes].
*/
//go:nosplit
func (self class) ClearBakedMeshes()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_clear_baked_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Bakes lightmap data for all meshes in the assigned [MeshLibrary].
*/
//go:nosplit
func (self class) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gen_lightmap_uv)
	callframe.Arg(frame, lightmap_uv_texel_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridMap.Bind_make_baked_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnCellSizeChanged(cb func(cell_size gd.Vector3)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("cell_size_changed"), gc.Callable(cb), 0)
}


func (self Go) OnChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("changed"), gc.Callable(cb), 0)
}


func (self class) AsGridMap() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGridMap() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("GridMap", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
