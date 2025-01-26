// Package TileMapLayer provides methods for working with TileMapLayer object instances.
package TileMapLayer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Rect2i"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Node for 2D tile-based maps. A [TileMapLayer] uses a [TileSet] which contain a list of tiles which are used to create grid-based maps. Unlike the [TileMap] node, which is deprecated, [TileMapLayer] has only one layer of tiles. You can use several [TileMapLayer] to achieve the same result as a [TileMap] node.
For performance reasons, all TileMap updates are batched at the end of a frame. Notably, this means that scene tiles from a [TileSetScenesCollectionSource] may be initialized after their parent. This is only queued when inside the scene tree.
To force an update earlier on, call [method update_internals].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=TileMapLayer)
*/
type Instance [1]gdclass.TileMapLayer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTileMapLayer() Instance
}
type Interface interface {
	//Should return [code]true[/code] if the tile at coordinates [param coords] requires a runtime update.
	//[b]Warning:[/b] Make sure this function only returns [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
	//[b]Note:[/b] If the result of this function should change, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
	UseTileDataRuntimeUpdate(coords Vector2i.XY) bool
	//Called with a [TileData] object about to be used internally by the [TileMapLayer], allowing its modification at runtime.
	//This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords].
	//[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
	//[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
	TileDataRuntimeUpdate(coords Vector2i.XY, tile_data [1]gdclass.TileData)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) UseTileDataRuntimeUpdate(coords Vector2i.XY) (_ bool) { return }
func (self implementation) TileDataRuntimeUpdate(coords Vector2i.XY, tile_data [1]gdclass.TileData) {
	return
}

/*
Should return [code]true[/code] if the tile at coordinates [param coords] requires a runtime update.
[b]Warning:[/b] Make sure this function only returns [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
[b]Note:[/b] If the result of this function should change, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
*/
func (Instance) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords Vector2i.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var coords = gd.UnsafeGet[gd.Vector2i](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, coords)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called with a [TileData] object about to be used internally by the [TileMapLayer], allowing its modification at runtime.
This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords].
[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
*/
func (Instance) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords Vector2i.XY, tile_data [1]gdclass.TileData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var coords = gd.UnsafeGet[gd.Vector2i](p_args, 0)

		var tile_data = [1]gdclass.TileData{pointers.New[gdclass.TileData]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(tile_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, coords, tile_data)
	}
}

/*
Sets the tile identifiers for the cell at coordinates [param coords]. Each tile of the [TileSet] is identified using three parts:
- The source identifier [param source_id] identifies a [TileSetSource] identifier. See [method TileSet.set_source_id],
- The atlas coordinate identifier [param atlas_coords] identifies a tile coordinates in the atlas (if the source is a [TileSetAtlasSource]). For [TileSetScenesCollectionSource] it should always be [code]Vector2i(0, 0)[/code],
- The alternative tile identifier [param alternative_tile] identifies a tile alternative in the atlas (if the source is a [TileSetAtlasSource]), and the scene for a [TileSetScenesCollectionSource].
If [param source_id] is set to [code]-1[/code], [param atlas_coords] to [code]Vector2i(-1, -1)[/code], or [param alternative_tile] to [code]-1[/code], the cell will be erased. An erased cell gets [b]all[/b] its identifiers automatically set to their respective invalid values, namely [code]-1[/code], [code]Vector2i(-1, -1)[/code] and [code]-1[/code].
*/
func (self Instance) SetCell(coords Vector2i.XY) { //gd:TileMapLayer.set_cell
	class(self).SetCell(gd.Vector2i(coords), gd.Int(-1), gd.Vector2i(gd.Vector2i{-1, -1}), gd.Int(0))
}

/*
Erases the cell at coordinates [param coords].
*/
func (self Instance) EraseCell(coords Vector2i.XY) { //gd:TileMapLayer.erase_cell
	class(self).EraseCell(gd.Vector2i(coords))
}

/*
Clears cells containing tiles that do not exist in the [member tile_set].
*/
func (self Instance) FixInvalidTiles() { //gd:TileMapLayer.fix_invalid_tiles
	class(self).FixInvalidTiles()
}

/*
Clears all cells.
*/
func (self Instance) Clear() { //gd:TileMapLayer.clear
	class(self).Clear()
}

/*
Returns the tile source ID of the cell at coordinates [param coords]. Returns [code]-1[/code] if the cell does not exist.
*/
func (self Instance) GetCellSourceId(coords Vector2i.XY) int { //gd:TileMapLayer.get_cell_source_id
	return int(int(class(self).GetCellSourceId(gd.Vector2i(coords))))
}

/*
Returns the tile atlas coordinates ID of the cell at coordinates [param coords]. Returns [code]Vector2i(-1, -1)[/code] if the cell does not exist.
*/
func (self Instance) GetCellAtlasCoords(coords Vector2i.XY) Vector2i.XY { //gd:TileMapLayer.get_cell_atlas_coords
	return Vector2i.XY(class(self).GetCellAtlasCoords(gd.Vector2i(coords)))
}

/*
Returns the tile alternative ID of the cell at coordinates [param coords].
*/
func (self Instance) GetCellAlternativeTile(coords Vector2i.XY) int { //gd:TileMapLayer.get_cell_alternative_tile
	return int(int(class(self).GetCellAlternativeTile(gd.Vector2i(coords))))
}

/*
Returns the [TileData] object associated with the given cell, or [code]null[/code] if the cell does not exist or is not a [TileSetAtlasSource].
[codeblock]
func get_clicked_tile_power():

	var clicked_cell = tile_map_layer.local_to_map(tile_map_layer.get_local_mouse_position())
	var data = tile_map_layer.get_cell_tile_data(clicked_cell)
	if data:
	    return data.get_custom_data("power")
	else:
	    return 0

[/codeblock]
*/
func (self Instance) GetCellTileData(coords Vector2i.XY) [1]gdclass.TileData { //gd:TileMapLayer.get_cell_tile_data
	return [1]gdclass.TileData(class(self).GetCellTileData(gd.Vector2i(coords)))
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile. A cell is considered empty if its source identifier equals [code]-1[/code], its atlas coordinate identifier is [code]Vector2(-1, -1)[/code] and its alternative identifier is [code]-1[/code].
*/
func (self Instance) GetUsedCells() []Vector2i.XY { //gd:TileMapLayer.get_used_cells
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetUsedCells())))
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile. Tiles may be filtered according to their source ([param source_id]), their atlas coordinates ([param atlas_coords]), or alternative id ([param alternative_tile]).
If a parameter has its value set to the default one, this parameter is not used to filter a cell. Thus, if all parameters have their respective default values, this method returns the same result as [method get_used_cells].
A cell is considered empty if its source identifier equals [code]-1[/code], its atlas coordinate identifier is [code]Vector2(-1, -1)[/code] and its alternative identifier is [code]-1[/code].
*/
func (self Instance) GetUsedCellsById() []Vector2i.XY { //gd:TileMapLayer.get_used_cells_by_id
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetUsedCellsById(gd.Int(-1), gd.Vector2i(gd.Vector2i{-1, -1}), gd.Int(-1)))))
}

/*
Returns a rectangle enclosing the used (non-empty) tiles of the map.
*/
func (self Instance) GetUsedRect() Rect2i.PositionSize { //gd:TileMapLayer.get_used_rect
	return Rect2i.PositionSize(class(self).GetUsedRect())
}

/*
Creates and returns a new [TileMapPattern] from the given array of cells. See also [method set_pattern].
*/
func (self Instance) GetPattern(coords_array []Vector2i.XY) [1]gdclass.TileMapPattern { //gd:TileMapLayer.get_pattern
	return [1]gdclass.TileMapPattern(class(self).GetPattern(gd.ArrayFromSlice[Array.Contains[gd.Vector2i]](coords_array)))
}

/*
Pastes the [TileMapPattern] at the given [param position] in the tile map. See also [method get_pattern].
*/
func (self Instance) SetPattern(position Vector2i.XY, pattern [1]gdclass.TileMapPattern) { //gd:TileMapLayer.set_pattern
	class(self).SetPattern(gd.Vector2i(position), pattern)
}

/*
Update all the cells in the [param cells] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. If an updated cell has the same terrain as one of its neighboring cells, this function tries to join the two. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
[b]Note:[/b] To work correctly, this method requires the [TileMapLayer]'s TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
func (self Instance) SetCellsTerrainConnect(cells []Vector2i.XY, terrain_set int, terrain int) { //gd:TileMapLayer.set_cells_terrain_connect
	class(self).SetCellsTerrainConnect(gd.ArrayFromSlice[Array.Contains[gd.Vector2i]](cells), gd.Int(terrain_set), gd.Int(terrain), true)
}

/*
Update all the cells in the [param path] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. The function will also connect two successive cell in the path with the same terrain. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
[b]Note:[/b] To work correctly, this method requires the [TileMapLayer]'s TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
func (self Instance) SetCellsTerrainPath(path []Vector2i.XY, terrain_set int, terrain int) { //gd:TileMapLayer.set_cells_terrain_path
	class(self).SetCellsTerrainPath(gd.ArrayFromSlice[Array.Contains[gd.Vector2i]](path), gd.Int(terrain_set), gd.Int(terrain), true)
}

/*
Returns whether the provided [param body] [RID] belongs to one of this [TileMapLayer]'s cells.
*/
func (self Instance) HasBodyRid(body RID.Body2D) bool { //gd:TileMapLayer.has_body_rid
	return bool(class(self).HasBodyRid(gd.RID(body)))
}

/*
Returns the coordinates of the tile for given physics body [RID]. Such an [RID] can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
func (self Instance) GetCoordsForBodyRid(body RID.Body2D) Vector2i.XY { //gd:TileMapLayer.get_coords_for_body_rid
	return Vector2i.XY(class(self).GetCoordsForBodyRid(gd.RID(body)))
}

/*
Triggers a direct update of the [TileMapLayer]. Usually, calling this function is not needed, as [TileMapLayer] node updates automatically when one of its properties or cells is modified.
However, for performance reasons, those updates are batched and delayed to the end of the frame. Calling this function will force the [TileMapLayer] to update right away instead.
[b]Warning:[/b] Updating the [TileMapLayer] is computationally expensive and may impact performance. Try to limit the number of updates and how many tiles they impact.
*/
func (self Instance) UpdateInternals() { //gd:TileMapLayer.update_internals
	class(self).UpdateInternals()
}

/*
Notifies the [TileMapLayer] node that calls to [method _use_tile_data_runtime_update] or [method _tile_data_runtime_update] will lead to different results. This will thus trigger a [TileMapLayer] update.
[b]Warning:[/b] Updating the [TileMapLayer] is computationally expensive and may impact performance. Try to limit the number of calls to this function to avoid unnecessary update.
[b]Note:[/b] This does not trigger a direct update of the [TileMapLayer], the update will be done at the end of the frame as usual (unless you call [method update_internals]).
*/
func (self Instance) NotifyRuntimeTileDataUpdate() { //gd:TileMapLayer.notify_runtime_tile_data_update
	class(self).NotifyRuntimeTileDataUpdate()
}

/*
Returns for the given coordinates [param coords_in_pattern] in a [TileMapPattern] the corresponding cell coordinates if the pattern was pasted at the [param position_in_tilemap] coordinates (see [method set_pattern]). This mapping is required as in half-offset tile shapes, the mapping might not work by calculating [code]position_in_tile_map + coords_in_pattern[/code].
*/
func (self Instance) MapPattern(position_in_tilemap Vector2i.XY, coords_in_pattern Vector2i.XY, pattern [1]gdclass.TileMapPattern) Vector2i.XY { //gd:TileMapLayer.map_pattern
	return Vector2i.XY(class(self).MapPattern(gd.Vector2i(position_in_tilemap), gd.Vector2i(coords_in_pattern), pattern))
}

/*
Returns the list of all neighboring cells to the one at [param coords].
*/
func (self Instance) GetSurroundingCells(coords Vector2i.XY) []Vector2i.XY { //gd:TileMapLayer.get_surrounding_cells
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetSurroundingCells(gd.Vector2i(coords)))))
}

/*
Returns the neighboring cell to the one at coordinates [param coords], identified by the [param neighbor] direction. This method takes into account the different layouts a TileMap can take.
*/
func (self Instance) GetNeighborCell(coords Vector2i.XY, neighbor gdclass.TileSetCellNeighbor) Vector2i.XY { //gd:TileMapLayer.get_neighbor_cell
	return Vector2i.XY(class(self).GetNeighborCell(gd.Vector2i(coords), neighbor))
}

/*
Returns the centered position of a cell in the [TileMapLayer]'s local coordinate space. To convert the returned value into global coordinates, use [method Node2D.to_global]. See also [method local_to_map].
[b]Note:[/b] This may not correspond to the visual position of the tile, i.e. it ignores the [member TileData.texture_origin] property of individual tiles.
*/
func (self Instance) MapToLocal(map_position Vector2i.XY) Vector2.XY { //gd:TileMapLayer.map_to_local
	return Vector2.XY(class(self).MapToLocal(gd.Vector2i(map_position)))
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node2D.to_local] before passing it to this method. See also [method map_to_local].
*/
func (self Instance) LocalToMap(local_position Vector2.XY) Vector2i.XY { //gd:TileMapLayer.local_to_map
	return Vector2i.XY(class(self).LocalToMap(gd.Vector2(local_position)))
}

/*
Sets a custom [param map] as a [NavigationServer2D] navigation map. If not set, uses the default [World2D] navigation map instead.
*/
func (self Instance) SetNavigationMap(mapping RID.NavigationMap2D) { //gd:TileMapLayer.set_navigation_map
	class(self).SetNavigationMap(gd.RID(mapping))
}

/*
Returns the [RID] of the [NavigationServer2D] navigation used by this [TileMapLayer].
By default this returns the default [World2D] navigation map, unless a custom map was provided using [method set_navigation_map].
*/
func (self Instance) GetNavigationMap() RID.NavigationMap2D { //gd:TileMapLayer.get_navigation_map
	return RID.NavigationMap2D(class(self).GetNavigationMap())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TileMapLayer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileMapLayer"))
	casted := Instance{*(*gdclass.TileMapLayer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TileMapData() []byte {
	return []byte(class(self).GetTileMapDataAsArray().Bytes())
}

func (self Instance) SetTileMapData(value []byte) {
	class(self).SetTileMapDataFromArray(gd.NewPackedByteSlice(value))
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) TileSet() [1]gdclass.TileSet {
	return [1]gdclass.TileSet(class(self).GetTileSet())
}

func (self Instance) SetTileSet(value [1]gdclass.TileSet) {
	class(self).SetTileSet(value)
}

func (self Instance) YSortOrigin() int {
	return int(int(class(self).GetYSortOrigin()))
}

func (self Instance) SetYSortOrigin(value int) {
	class(self).SetYSortOrigin(gd.Int(value))
}

func (self Instance) XDrawOrderReversed() bool {
	return bool(class(self).IsXDrawOrderReversed())
}

func (self Instance) SetXDrawOrderReversed(value bool) {
	class(self).SetXDrawOrderReversed(value)
}

func (self Instance) RenderingQuadrantSize() int {
	return int(int(class(self).GetRenderingQuadrantSize()))
}

func (self Instance) SetRenderingQuadrantSize(value int) {
	class(self).SetRenderingQuadrantSize(gd.Int(value))
}

func (self Instance) CollisionEnabled() bool {
	return bool(class(self).IsCollisionEnabled())
}

func (self Instance) SetCollisionEnabled(value bool) {
	class(self).SetCollisionEnabled(value)
}

func (self Instance) UseKinematicBodies() bool {
	return bool(class(self).IsUsingKinematicBodies())
}

func (self Instance) SetUseKinematicBodies(value bool) {
	class(self).SetUseKinematicBodies(value)
}

func (self Instance) CollisionVisibilityMode() gdclass.TileMapLayerDebugVisibilityMode {
	return gdclass.TileMapLayerDebugVisibilityMode(class(self).GetCollisionVisibilityMode())
}

func (self Instance) SetCollisionVisibilityMode(value gdclass.TileMapLayerDebugVisibilityMode) {
	class(self).SetCollisionVisibilityMode(value)
}

func (self Instance) NavigationEnabled() bool {
	return bool(class(self).IsNavigationEnabled())
}

func (self Instance) SetNavigationEnabled(value bool) {
	class(self).SetNavigationEnabled(value)
}

func (self Instance) NavigationVisibilityMode() gdclass.TileMapLayerDebugVisibilityMode {
	return gdclass.TileMapLayerDebugVisibilityMode(class(self).GetNavigationVisibilityMode())
}

func (self Instance) SetNavigationVisibilityMode(value gdclass.TileMapLayerDebugVisibilityMode) {
	class(self).SetNavigationVisibilityMode(value)
}

/*
Should return [code]true[/code] if the tile at coordinates [param coords] requires a runtime update.
[b]Warning:[/b] Make sure this function only returns [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
[b]Note:[/b] If the result of this function should change, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
*/
func (class) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords gd.Vector2i) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var coords = gd.UnsafeGet[gd.Vector2i](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, coords)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called with a [TileData] object about to be used internally by the [TileMapLayer], allowing its modification at runtime.
This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords].
[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
*/
func (class) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords gd.Vector2i, tile_data [1]gdclass.TileData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var coords = gd.UnsafeGet[gd.Vector2i](p_args, 0)

		var tile_data = [1]gdclass.TileData{pointers.New[gdclass.TileData]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(tile_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, coords, tile_data)
	}
}

/*
Sets the tile identifiers for the cell at coordinates [param coords]. Each tile of the [TileSet] is identified using three parts:
- The source identifier [param source_id] identifies a [TileSetSource] identifier. See [method TileSet.set_source_id],
- The atlas coordinate identifier [param atlas_coords] identifies a tile coordinates in the atlas (if the source is a [TileSetAtlasSource]). For [TileSetScenesCollectionSource] it should always be [code]Vector2i(0, 0)[/code],
- The alternative tile identifier [param alternative_tile] identifies a tile alternative in the atlas (if the source is a [TileSetAtlasSource]), and the scene for a [TileSetScenesCollectionSource].
If [param source_id] is set to [code]-1[/code], [param atlas_coords] to [code]Vector2i(-1, -1)[/code], or [param alternative_tile] to [code]-1[/code], the cell will be erased. An erased cell gets [b]all[/b] its identifiers automatically set to their respective invalid values, namely [code]-1[/code], [code]Vector2i(-1, -1)[/code] and [code]-1[/code].
*/
//go:nosplit
func (self class) SetCell(coords gd.Vector2i, source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int) { //gd:TileMapLayer.set_cell
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Erases the cell at coordinates [param coords].
*/
//go:nosplit
func (self class) EraseCell(coords gd.Vector2i) { //gd:TileMapLayer.erase_cell
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_erase_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears cells containing tiles that do not exist in the [member tile_set].
*/
//go:nosplit
func (self class) FixInvalidTiles() { //gd:TileMapLayer.fix_invalid_tiles
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_fix_invalid_tiles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears all cells.
*/
//go:nosplit
func (self class) Clear() { //gd:TileMapLayer.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tile source ID of the cell at coordinates [param coords]. Returns [code]-1[/code] if the cell does not exist.
*/
//go:nosplit
func (self class) GetCellSourceId(coords gd.Vector2i) gd.Int { //gd:TileMapLayer.get_cell_source_id
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_cell_source_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tile atlas coordinates ID of the cell at coordinates [param coords]. Returns [code]Vector2i(-1, -1)[/code] if the cell does not exist.
*/
//go:nosplit
func (self class) GetCellAtlasCoords(coords gd.Vector2i) gd.Vector2i { //gd:TileMapLayer.get_cell_atlas_coords
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_cell_atlas_coords, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tile alternative ID of the cell at coordinates [param coords].
*/
//go:nosplit
func (self class) GetCellAlternativeTile(coords gd.Vector2i) gd.Int { //gd:TileMapLayer.get_cell_alternative_tile
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_cell_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [TileData] object associated with the given cell, or [code]null[/code] if the cell does not exist or is not a [TileSetAtlasSource].
[codeblock]
func get_clicked_tile_power():
    var clicked_cell = tile_map_layer.local_to_map(tile_map_layer.get_local_mouse_position())
    var data = tile_map_layer.get_cell_tile_data(clicked_cell)
    if data:
        return data.get_custom_data("power")
    else:
        return 0
[/codeblock]
*/
//go:nosplit
func (self class) GetCellTileData(coords gd.Vector2i) [1]gdclass.TileData { //gd:TileMapLayer.get_cell_tile_data
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_cell_tile_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TileData{gd.PointerMustAssertInstanceID[gdclass.TileData](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile. A cell is considered empty if its source identifier equals [code]-1[/code], its atlas coordinate identifier is [code]Vector2(-1, -1)[/code] and its alternative identifier is [code]-1[/code].
*/
//go:nosplit
func (self class) GetUsedCells() Array.Contains[gd.Vector2i] { //gd:TileMapLayer.get_used_cells
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector2i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile. Tiles may be filtered according to their source ([param source_id]), their atlas coordinates ([param atlas_coords]), or alternative id ([param alternative_tile]).
If a parameter has its value set to the default one, this parameter is not used to filter a cell. Thus, if all parameters have their respective default values, this method returns the same result as [method get_used_cells].
A cell is considered empty if its source identifier equals [code]-1[/code], its atlas coordinate identifier is [code]Vector2(-1, -1)[/code] and its alternative identifier is [code]-1[/code].
*/
//go:nosplit
func (self class) GetUsedCellsById(source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int) Array.Contains[gd.Vector2i] { //gd:TileMapLayer.get_used_cells_by_id
	var frame = callframe.New()
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_used_cells_by_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector2i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a rectangle enclosing the used (non-empty) tiles of the map.
*/
//go:nosplit
func (self class) GetUsedRect() gd.Rect2i { //gd:TileMapLayer.get_used_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_used_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates and returns a new [TileMapPattern] from the given array of cells. See also [method set_pattern].
*/
//go:nosplit
func (self class) GetPattern(coords_array Array.Contains[gd.Vector2i]) [1]gdclass.TileMapPattern { //gd:TileMapLayer.get_pattern
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(coords_array)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TileMapPattern{gd.PointerWithOwnershipTransferredToGo[gdclass.TileMapPattern](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Pastes the [TileMapPattern] at the given [param position] in the tile map. See also [method get_pattern].
*/
//go:nosplit
func (self class) SetPattern(position gd.Vector2i, pattern [1]gdclass.TileMapPattern) { //gd:TileMapLayer.set_pattern
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, pointers.Get(pattern[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Update all the cells in the [param cells] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. If an updated cell has the same terrain as one of its neighboring cells, this function tries to join the two. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
[b]Note:[/b] To work correctly, this method requires the [TileMapLayer]'s TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainConnect(cells Array.Contains[gd.Vector2i], terrain_set gd.Int, terrain gd.Int, ignore_empty_terrains bool) { //gd:TileMapLayer.set_cells_terrain_connect
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(cells)))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_cells_terrain_connect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Update all the cells in the [param path] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. The function will also connect two successive cell in the path with the same terrain. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
[b]Note:[/b] To work correctly, this method requires the [TileMapLayer]'s TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainPath(path Array.Contains[gd.Vector2i], terrain_set gd.Int, terrain gd.Int, ignore_empty_terrains bool) { //gd:TileMapLayer.set_cells_terrain_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(path)))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_cells_terrain_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the provided [param body] [RID] belongs to one of this [TileMapLayer]'s cells.
*/
//go:nosplit
func (self class) HasBodyRid(body gd.RID) bool { //gd:TileMapLayer.has_body_rid
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_has_body_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the coordinates of the tile for given physics body [RID]. Such an [RID] can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
//go:nosplit
func (self class) GetCoordsForBodyRid(body gd.RID) gd.Vector2i { //gd:TileMapLayer.get_coords_for_body_rid
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_coords_for_body_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Triggers a direct update of the [TileMapLayer]. Usually, calling this function is not needed, as [TileMapLayer] node updates automatically when one of its properties or cells is modified.
However, for performance reasons, those updates are batched and delayed to the end of the frame. Calling this function will force the [TileMapLayer] to update right away instead.
[b]Warning:[/b] Updating the [TileMapLayer] is computationally expensive and may impact performance. Try to limit the number of updates and how many tiles they impact.
*/
//go:nosplit
func (self class) UpdateInternals() { //gd:TileMapLayer.update_internals
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_update_internals, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Notifies the [TileMapLayer] node that calls to [method _use_tile_data_runtime_update] or [method _tile_data_runtime_update] will lead to different results. This will thus trigger a [TileMapLayer] update.
[b]Warning:[/b] Updating the [TileMapLayer] is computationally expensive and may impact performance. Try to limit the number of calls to this function to avoid unnecessary update.
[b]Note:[/b] This does not trigger a direct update of the [TileMapLayer], the update will be done at the end of the frame as usual (unless you call [method update_internals]).
*/
//go:nosplit
func (self class) NotifyRuntimeTileDataUpdate() { //gd:TileMapLayer.notify_runtime_tile_data_update
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_notify_runtime_tile_data_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns for the given coordinates [param coords_in_pattern] in a [TileMapPattern] the corresponding cell coordinates if the pattern was pasted at the [param position_in_tilemap] coordinates (see [method set_pattern]). This mapping is required as in half-offset tile shapes, the mapping might not work by calculating [code]position_in_tile_map + coords_in_pattern[/code].
*/
//go:nosplit
func (self class) MapPattern(position_in_tilemap gd.Vector2i, coords_in_pattern gd.Vector2i, pattern [1]gdclass.TileMapPattern) gd.Vector2i { //gd:TileMapLayer.map_pattern
	var frame = callframe.New()
	callframe.Arg(frame, position_in_tilemap)
	callframe.Arg(frame, coords_in_pattern)
	callframe.Arg(frame, pointers.Get(pattern[0])[0])
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_map_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of all neighboring cells to the one at [param coords].
*/
//go:nosplit
func (self class) GetSurroundingCells(coords gd.Vector2i) Array.Contains[gd.Vector2i] { //gd:TileMapLayer.get_surrounding_cells
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_surrounding_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector2i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the neighboring cell to the one at coordinates [param coords], identified by the [param neighbor] direction. This method takes into account the different layouts a TileMap can take.
*/
//go:nosplit
func (self class) GetNeighborCell(coords gd.Vector2i, neighbor gdclass.TileSetCellNeighbor) gd.Vector2i { //gd:TileMapLayer.get_neighbor_cell
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, neighbor)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_neighbor_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the centered position of a cell in the [TileMapLayer]'s local coordinate space. To convert the returned value into global coordinates, use [method Node2D.to_global]. See also [method local_to_map].
[b]Note:[/b] This may not correspond to the visual position of the tile, i.e. it ignores the [member TileData.texture_origin] property of individual tiles.
*/
//go:nosplit
func (self class) MapToLocal(map_position gd.Vector2i) gd.Vector2 { //gd:TileMapLayer.map_to_local
	var frame = callframe.New()
	callframe.Arg(frame, map_position)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_map_to_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node2D.to_local] before passing it to this method. See also [method map_to_local].
*/
//go:nosplit
func (self class) LocalToMap(local_position gd.Vector2) gd.Vector2i { //gd:TileMapLayer.local_to_map
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_local_to_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTileMapDataFromArray(tile_map_layer_data gd.PackedByteArray) { //gd:TileMapLayer.set_tile_map_data_from_array
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tile_map_layer_data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_tile_map_data_from_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTileMapDataAsArray() gd.PackedByteArray { //gd:TileMapLayer.get_tile_map_data_as_array
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_tile_map_data_as_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnabled(enabled bool) { //gd:TileMapLayer.set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool { //gd:TileMapLayer.is_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTileSet(tile_set [1]gdclass.TileSet) { //gd:TileMapLayer.set_tile_set
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tile_set[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_tile_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTileSet() [1]gdclass.TileSet { //gd:TileMapLayer.get_tile_set
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_tile_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TileSet{gd.PointerWithOwnershipTransferredToGo[gdclass.TileSet](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetYSortOrigin(y_sort_origin gd.Int) { //gd:TileMapLayer.set_y_sort_origin
	var frame = callframe.New()
	callframe.Arg(frame, y_sort_origin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetYSortOrigin() gd.Int { //gd:TileMapLayer.get_y_sort_origin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetXDrawOrderReversed(x_draw_order_reversed bool) { //gd:TileMapLayer.set_x_draw_order_reversed
	var frame = callframe.New()
	callframe.Arg(frame, x_draw_order_reversed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_x_draw_order_reversed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsXDrawOrderReversed() bool { //gd:TileMapLayer.is_x_draw_order_reversed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_is_x_draw_order_reversed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRenderingQuadrantSize(size gd.Int) { //gd:TileMapLayer.set_rendering_quadrant_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRenderingQuadrantSize() gd.Int { //gd:TileMapLayer.get_rendering_quadrant_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionEnabled(enabled bool) { //gd:TileMapLayer.set_collision_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollisionEnabled() bool { //gd:TileMapLayer.is_collision_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_is_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseKinematicBodies(use_kinematic_bodies bool) { //gd:TileMapLayer.set_use_kinematic_bodies
	var frame = callframe.New()
	callframe.Arg(frame, use_kinematic_bodies)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_use_kinematic_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingKinematicBodies() bool { //gd:TileMapLayer.is_using_kinematic_bodies
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_is_using_kinematic_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionVisibilityMode(visibility_mode gdclass.TileMapLayerDebugVisibilityMode) { //gd:TileMapLayer.set_collision_visibility_mode
	var frame = callframe.New()
	callframe.Arg(frame, visibility_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionVisibilityMode() gdclass.TileMapLayerDebugVisibilityMode { //gd:TileMapLayer.get_collision_visibility_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TileMapLayerDebugVisibilityMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationEnabled(enabled bool) { //gd:TileMapLayer.set_navigation_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsNavigationEnabled() bool { //gd:TileMapLayer.is_navigation_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_is_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a custom [param map] as a [NavigationServer2D] navigation map. If not set, uses the default [World2D] navigation map instead.
*/
//go:nosplit
func (self class) SetNavigationMap(mapping gd.RID) { //gd:TileMapLayer.set_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [RID] of the [NavigationServer2D] navigation used by this [TileMapLayer].
By default this returns the default [World2D] navigation map, unless a custom map was provided using [method set_navigation_map].
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID { //gd:TileMapLayer.get_navigation_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationVisibilityMode(show_navigation gdclass.TileMapLayerDebugVisibilityMode) { //gd:TileMapLayer.set_navigation_visibility_mode
	var frame = callframe.New()
	callframe.Arg(frame, show_navigation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_set_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationVisibilityMode() gdclass.TileMapLayerDebugVisibilityMode { //gd:TileMapLayer.get_navigation_visibility_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TileMapLayerDebugVisibilityMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapLayer.Bind_get_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self class) AsTileMapLayer() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTileMapLayer() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_use_tile_data_runtime_update":
		return reflect.ValueOf(self._use_tile_data_runtime_update)
	case "_tile_data_runtime_update":
		return reflect.ValueOf(self._tile_data_runtime_update)
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_use_tile_data_runtime_update":
		return reflect.ValueOf(self._use_tile_data_runtime_update)
	case "_tile_data_runtime_update":
		return reflect.ValueOf(self._tile_data_runtime_update)
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("TileMapLayer", func(ptr gd.Object) any {
		return [1]gdclass.TileMapLayer{*(*gdclass.TileMapLayer)(unsafe.Pointer(&ptr))}
	})
}

type DebugVisibilityMode = gdclass.TileMapLayerDebugVisibilityMode //gd:TileMapLayer.DebugVisibilityMode

const (
	/*Hide the collisions or navigation debug shapes in the editor, and use the debug settings to determine their visibility in game (i.e. [member SceneTree.debug_collisions_hint] or [member SceneTree.debug_navigation_hint]).*/
	DebugVisibilityModeDefault DebugVisibilityMode = 0
	/*Always hide the collisions or navigation debug shapes.*/
	DebugVisibilityModeForceHide DebugVisibilityMode = 2
	/*Always show the collisions or navigation debug shapes.*/
	DebugVisibilityModeForceShow DebugVisibilityMode = 1
)
