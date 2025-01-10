// Package GridMap provides methods for working with GridMap object instances.
package GridMap

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Vector3i"
import "graphics.gd/variant/Basis"
import "graphics.gd/variant/Array"

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
GridMap lets you place meshes on a grid interactively. It works both from the editor and from scripts, which can help you create in-game level editors.
GridMaps use a [MeshLibrary] which contains a list of tiles. Each tile is a mesh with materials plus optional collision and navigation shapes.
A GridMap contains a collection of cells. Each grid cell refers to a tile in the [MeshLibrary]. All cells in the map have the same dimensions.
Internally, a GridMap is split into a sparse collection of octants for efficient rendering and physics processing. Every octant has the same dimensions and can contain several cells.
[b]Note:[/b] GridMap doesn't extend [VisualInstance3D] and therefore can't be hidden or cull masked based on [member VisualInstance3D.layers]. If you make a light not affect the first layer, the whole GridMap won't be lit by the light in question.
*/
type Instance [1]gdclass.GridMap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGridMap() Instance
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionLayerValue(layer_number int, value bool) {
	class(self).SetCollisionLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionLayerValue(layer_number int) bool {
	return bool(class(self).GetCollisionLayerValue(gd.Int(layer_number)))
}

/*
Sets the [RID] of the navigation map this GridMap node should use for its cell baked navigation meshes.
*/
func (self Instance) SetNavigationMap(navigation_map Resource.ID) {
	class(self).SetNavigationMap(navigation_map)
}

/*
Returns the [RID] of the navigation map this GridMap node uses for its cell baked navigation meshes.
This function returns always the map set on the GridMap node and not the map on the NavigationServer. If the map is changed directly with the NavigationServer API the GridMap node will not be aware of the map change.
*/
func (self Instance) GetNavigationMap() Resource.ID {
	return Resource.ID(class(self).GetNavigationMap())
}

/*
Sets the mesh index for the cell referenced by its grid coordinates.
A negative item index such as [constant INVALID_CELL_ITEM] will clear the cell.
Optionally, the item's orientation can be passed. For valid orientation values, see [method get_orthogonal_index_from_basis].
*/
func (self Instance) SetCellItem(position Vector3i.XYZ, item int) {
	class(self).SetCellItem(gd.Vector3i(position), gd.Int(item), gd.Int(0))
}

/*
The [MeshLibrary] item index located at the given grid coordinates. If the cell is empty, [constant INVALID_CELL_ITEM] will be returned.
*/
func (self Instance) GetCellItem(position Vector3i.XYZ) int {
	return int(int(class(self).GetCellItem(gd.Vector3i(position))))
}

/*
The orientation of the cell at the given grid coordinates. [code]-1[/code] is returned if the cell is empty.
*/
func (self Instance) GetCellItemOrientation(position Vector3i.XYZ) int {
	return int(int(class(self).GetCellItemOrientation(gd.Vector3i(position))))
}

/*
Returns the basis that gives the specified cell its orientation.
*/
func (self Instance) GetCellItemBasis(position Vector3i.XYZ) Basis.XYZ {
	return Basis.XYZ(class(self).GetCellItemBasis(gd.Vector3i(position)))
}

/*
Returns one of 24 possible rotations that lie along the vectors (x,y,z) with each component being either -1, 0, or 1. For further details, refer to the Godot source code.
*/
func (self Instance) GetBasisWithOrthogonalIndex(index int) Basis.XYZ {
	return Basis.XYZ(class(self).GetBasisWithOrthogonalIndex(gd.Int(index)))
}

/*
This function considers a discretization of rotations into 24 points on unit sphere, lying along the vectors (x,y,z) with each component being either -1, 0, or 1, and returns the index (in the range from 0 to 23) of the point best representing the orientation of the object. For further details, refer to the Godot source code.
*/
func (self Instance) GetOrthogonalIndexFromBasis(basis Basis.XYZ) int {
	return int(int(class(self).GetOrthogonalIndexFromBasis(gd.Basis(basis))))
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node3D.to_local] before passing it to this method. See also [method map_to_local].
*/
func (self Instance) LocalToMap(local_position Vector3.XYZ) Vector3i.XYZ {
	return Vector3i.XYZ(class(self).LocalToMap(gd.Vector3(local_position)))
}

/*
Returns the position of a grid cell in the GridMap's local coordinate space. To convert the returned value into global coordinates, use [method Node3D.to_global]. See also [method local_to_map].
*/
func (self Instance) MapToLocal(map_position Vector3i.XYZ) Vector3.XYZ {
	return Vector3.XYZ(class(self).MapToLocal(gd.Vector3i(map_position)))
}

/*
This method does nothing.
*/
func (self Instance) ResourceChanged(resource [1]gdclass.Resource) {
	class(self).ResourceChanged(resource)
}

/*
Clear all cells.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Returns an array of [Vector3] with the non-empty cell coordinates in the grid map.
*/
func (self Instance) GetUsedCells() gd.Array {
	return gd.Array(class(self).GetUsedCells())
}

/*
Returns an array of all cells with the given item index specified in [param item].
*/
func (self Instance) GetUsedCellsByItem(item int) gd.Array {
	return gd.Array(class(self).GetUsedCellsByItem(gd.Int(item)))
}

/*
Returns an array of [Transform3D] and [Mesh] references corresponding to the non-empty cells in the grid. The transforms are specified in local space.
*/
func (self Instance) GetMeshes() Array.Any {
	return Array.Any(class(self).GetMeshes())
}

/*
Returns an array of [ArrayMesh]es and [Transform3D] references of all bake meshes that exist within the current GridMap.
*/
func (self Instance) GetBakeMeshes() Array.Any {
	return Array.Any(class(self).GetBakeMeshes())
}

/*
Returns [RID] of a baked mesh with the given [param idx].
*/
func (self Instance) GetBakeMeshInstance(idx int) Resource.ID {
	return Resource.ID(class(self).GetBakeMeshInstance(gd.Int(idx)))
}

/*
Clears all baked meshes. See [method make_baked_meshes].
*/
func (self Instance) ClearBakedMeshes() {
	class(self).ClearBakedMeshes()
}

/*
Bakes lightmap data for all meshes in the assigned [MeshLibrary].
*/
func (self Instance) MakeBakedMeshes() {
	class(self).MakeBakedMeshes(false, gd.Float(0.1))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GridMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GridMap"))
	casted := Instance{*(*gdclass.GridMap)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) MeshLibrary() [1]gdclass.MeshLibrary {
	return [1]gdclass.MeshLibrary(class(self).GetMeshLibrary())
}

func (self Instance) SetMeshLibrary(value [1]gdclass.MeshLibrary) {
	class(self).SetMeshLibrary(value)
}

func (self Instance) PhysicsMaterial() [1]gdclass.PhysicsMaterial {
	return [1]gdclass.PhysicsMaterial(class(self).GetPhysicsMaterial())
}

func (self Instance) SetPhysicsMaterial(value [1]gdclass.PhysicsMaterial) {
	class(self).SetPhysicsMaterial(value)
}

func (self Instance) CellSize() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetCellSize())
}

func (self Instance) SetCellSize(value Vector3.XYZ) {
	class(self).SetCellSize(gd.Vector3(value))
}

func (self Instance) CellOctantSize() int {
	return int(int(class(self).GetOctantSize()))
}

func (self Instance) SetCellOctantSize(value int) {
	class(self).SetOctantSize(gd.Int(value))
}

func (self Instance) CellCenterX() bool {
	return bool(class(self).GetCenterX())
}

func (self Instance) SetCellCenterX(value bool) {
	class(self).SetCenterX(value)
}

func (self Instance) CellCenterY() bool {
	return bool(class(self).GetCenterY())
}

func (self Instance) SetCellCenterY(value bool) {
	class(self).SetCenterY(value)
}

func (self Instance) CellCenterZ() bool {
	return bool(class(self).GetCenterZ())
}

func (self Instance) SetCellCenterZ(value bool) {
	class(self).SetCenterZ(value)
}

func (self Instance) CellScale() Float.X {
	return Float.X(Float.X(class(self).GetCellScale()))
}

func (self Instance) SetCellScale(value Float.X) {
	class(self).SetCellScale(gd.Float(value))
}

func (self Instance) CollisionLayer() int {
	return int(int(class(self).GetCollisionLayer()))
}

func (self Instance) SetCollisionLayer(value int) {
	class(self).SetCollisionLayer(gd.Int(value))
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) CollisionPriority() Float.X {
	return Float.X(Float.X(class(self).GetCollisionPriority()))
}

func (self Instance) SetCollisionPriority(value Float.X) {
	class(self).SetCollisionPriority(gd.Float(value))
}

func (self Instance) BakeNavigation() bool {
	return bool(class(self).IsBakingNavigation())
}

func (self Instance) SetBakeNavigation(value bool) {
	class(self).SetBakeNavigation(value)
}

//go:nosplit
func (self class) SetCollisionLayer(layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionPriority(priority gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionPriority() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsMaterial(material [1]gdclass.PhysicsMaterial) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_physics_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsMaterial() [1]gdclass.PhysicsMaterial {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_physics_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.PhysicsMaterial{gd.PointerWithOwnershipTransferredToGo[gdclass.PhysicsMaterial](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeNavigation(bake_navigation bool) {
	var frame = callframe.New()
	callframe.Arg(frame, bake_navigation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_bake_navigation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsBakingNavigation() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_is_baking_navigation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [RID] of the navigation map this GridMap node should use for its cell baked navigation meshes.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [RID] of the navigation map this GridMap node uses for its cell baked navigation meshes.
This function returns always the map set on the GridMap node and not the map on the NavigationServer. If the map is changed directly with the NavigationServer API the GridMap node will not be aware of the map change.
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMeshLibrary(mesh_library [1]gdclass.MeshLibrary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh_library[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_mesh_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMeshLibrary() [1]gdclass.MeshLibrary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_mesh_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.MeshLibrary{gd.PointerWithOwnershipTransferredToGo[gdclass.MeshLibrary](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellSize(size gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_cell_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOctantSize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_octant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOctantSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_octant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetCellItem(position gd.Vector3i, item gd.Int, orientation gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, item)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_cell_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
The [MeshLibrary] item index located at the given grid coordinates. If the cell is empty, [constant INVALID_CELL_ITEM] will be returned.
*/
//go:nosplit
func (self class) GetCellItem(position gd.Vector3i) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The orientation of the cell at the given grid coordinates. [code]-1[/code] is returned if the cell is empty.
*/
//go:nosplit
func (self class) GetCellItemOrientation(position gd.Vector3i) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_item_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the basis that gives the specified cell its orientation.
*/
//go:nosplit
func (self class) GetCellItemBasis(position gd.Vector3i) gd.Basis {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_item_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns one of 24 possible rotations that lie along the vectors (x,y,z) with each component being either -1, 0, or 1. For further details, refer to the Godot source code.
*/
//go:nosplit
func (self class) GetBasisWithOrthogonalIndex(index gd.Int) gd.Basis {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_basis_with_orthogonal_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
This function considers a discretization of rotations into 24 points on unit sphere, lying along the vectors (x,y,z) with each component being either -1, 0, or 1, and returns the index (in the range from 0 to 23) of the point best representing the orientation of the object. For further details, refer to the Godot source code.
*/
//go:nosplit
func (self class) GetOrthogonalIndexFromBasis(basis gd.Basis) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, basis)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_orthogonal_index_from_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node3D.to_local] before passing it to this method. See also [method map_to_local].
*/
//go:nosplit
func (self class) LocalToMap(local_position gd.Vector3) gd.Vector3i {
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[gd.Vector3i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_local_to_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of a grid cell in the GridMap's local coordinate space. To convert the returned value into global coordinates, use [method Node3D.to_global]. See also [method local_to_map].
*/
//go:nosplit
func (self class) MapToLocal(map_position gd.Vector3i) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, map_position)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_map_to_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) ResourceChanged(resource [1]gdclass.Resource) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_resource_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCenterX(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_center_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterX() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_center_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterY(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_center_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterY() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_center_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterZ(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_center_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterZ() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_center_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clear all cells.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of [Vector3] with the non-empty cell coordinates in the grid map.
*/
//go:nosplit
func (self class) GetUsedCells() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of all cells with the given item index specified in [param item].
*/
//go:nosplit
func (self class) GetUsedCellsByItem(item gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, item)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_used_cells_by_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of [Transform3D] and [Mesh] references corresponding to the non-empty cells in the grid. The transforms are specified in local space.
*/
//go:nosplit
func (self class) GetMeshes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of [ArrayMesh]es and [Transform3D] references of all bake meshes that exist within the current GridMap.
*/
//go:nosplit
func (self class) GetBakeMeshes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_bake_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [RID] of a baked mesh with the given [param idx].
*/
//go:nosplit
func (self class) GetBakeMeshInstance(idx gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_bake_mesh_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears all baked meshes. See [method make_baked_meshes].
*/
//go:nosplit
func (self class) ClearBakedMeshes() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_clear_baked_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Bakes lightmap data for all meshes in the assigned [MeshLibrary].
*/
//go:nosplit
func (self class) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, gen_lightmap_uv)
	callframe.Arg(frame, lightmap_uv_texel_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_make_baked_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnCellSizeChanged(cb func(cell_size Vector3.XYZ)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("cell_size_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self class) AsGridMap() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGridMap() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("GridMap", func(ptr gd.Object) any { return [1]gdclass.GridMap{*(*gdclass.GridMap)(unsafe.Pointer(&ptr))} })
}
