package TileMapLayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Node for 2D tile-based maps. A [TileMapLayer] uses a [TileSet] which contain a list of tiles which are used to create grid-based maps. Unlike the [TileMap] node, which is deprecated, [TileMapLayer] has only one layer of tiles. You can use several [TileMapLayer] to achieve the same result as a [TileMap] node.
For performance reasons, all TileMap updates are batched at the end of a frame. Notably, this means that scene tiles from a [TileSetScenesCollectionSource] may be initialized after their parent. This is only queued when inside the scene tree.
To force an update earlier on, call [method update_internals].
	// TileMapLayer methods that can be overridden by a [Class] that extends it.
	type TileMapLayer interface {
		//Should return [code]true[/code] if the tile at coordinates [param coords] requires a runtime update.
		//[b]Warning:[/b] Make sure this function only returns [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
		//[b]Note:[/b] If the result of this function should change, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
		UseTileDataRuntimeUpdate(coords gd.Vector2i) bool
		//Called with a [TileData] object about to be used internally by the [TileMapLayer], allowing its modification at runtime.
		//This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords].
		//[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
		//[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
		TileDataRuntimeUpdate(coords gd.Vector2i, tile_data object.TileData) 
	}

*/
type Simple [1]classdb.TileMapLayer
func (Simple) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords gd.Vector2i) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, coords)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords gd.Vector2i, tile_data [1]classdb.TileData) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,0)
		var tile_data [1]classdb.TileData
		tile_data[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, coords, tile_data)
		gc.End()
	}
}
func (self Simple) SetCell(coords gd.Vector2i, source_id int, atlas_coords gd.Vector2i, alternative_tile int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCell(coords, gd.Int(source_id), atlas_coords, gd.Int(alternative_tile))
}
func (self Simple) EraseCell(coords gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseCell(coords)
}
func (self Simple) FixInvalidTiles() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FixInvalidTiles()
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) GetCellSourceId(coords gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCellSourceId(coords)))
}
func (self Simple) GetCellAtlasCoords(coords gd.Vector2i) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetCellAtlasCoords(coords))
}
func (self Simple) GetCellAlternativeTile(coords gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCellAlternativeTile(coords)))
}
func (self Simple) GetCellTileData(coords gd.Vector2i) [1]classdb.TileData {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TileData(Expert(self).GetCellTileData(gc, coords))
}
func (self Simple) GetUsedCells() gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetUsedCells(gc))
}
func (self Simple) GetUsedCellsById(source_id int, atlas_coords gd.Vector2i, alternative_tile int) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetUsedCellsById(gc, gd.Int(source_id), atlas_coords, gd.Int(alternative_tile)))
}
func (self Simple) GetUsedRect() gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).GetUsedRect())
}
func (self Simple) GetPattern(coords_array gd.ArrayOf[gd.Vector2i]) [1]classdb.TileMapPattern {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TileMapPattern(Expert(self).GetPattern(gc, coords_array))
}
func (self Simple) SetPattern(position gd.Vector2i, pattern [1]classdb.TileMapPattern) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPattern(position, pattern)
}
func (self Simple) SetCellsTerrainConnect(cells gd.ArrayOf[gd.Vector2i], terrain_set int, terrain int, ignore_empty_terrains bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellsTerrainConnect(cells, gd.Int(terrain_set), gd.Int(terrain), ignore_empty_terrains)
}
func (self Simple) SetCellsTerrainPath(path gd.ArrayOf[gd.Vector2i], terrain_set int, terrain int, ignore_empty_terrains bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellsTerrainPath(path, gd.Int(terrain_set), gd.Int(terrain), ignore_empty_terrains)
}
func (self Simple) HasBodyRid(body gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasBodyRid(body))
}
func (self Simple) GetCoordsForBodyRid(body gd.RID) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetCoordsForBodyRid(body))
}
func (self Simple) UpdateInternals() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateInternals()
}
func (self Simple) NotifyRuntimeTileDataUpdate() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).NotifyRuntimeTileDataUpdate()
}
func (self Simple) MapPattern(position_in_tilemap gd.Vector2i, coords_in_pattern gd.Vector2i, pattern [1]classdb.TileMapPattern) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).MapPattern(position_in_tilemap, coords_in_pattern, pattern))
}
func (self Simple) GetSurroundingCells(coords gd.Vector2i) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetSurroundingCells(gc, coords))
}
func (self Simple) GetNeighborCell(coords gd.Vector2i, neighbor classdb.TileSetCellNeighbor) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetNeighborCell(coords, neighbor))
}
func (self Simple) MapToLocal(map_position gd.Vector2i) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).MapToLocal(map_position))
}
func (self Simple) LocalToMap(local_position gd.Vector2) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).LocalToMap(local_position))
}
func (self Simple) SetTileMapDataFromArray(tile_map_layer_data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileMapDataFromArray(gc.PackedByteSlice(tile_map_layer_data))
}
func (self Simple) GetTileMapDataAsArray() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetTileMapDataAsArray(gc).Bytes())
}
func (self Simple) SetEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnabled(enabled)
}
func (self Simple) IsEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEnabled())
}
func (self Simple) SetTileSet(tile_set [1]classdb.TileSet) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileSet(tile_set)
}
func (self Simple) GetTileSet() [1]classdb.TileSet {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TileSet(Expert(self).GetTileSet(gc))
}
func (self Simple) SetYSortOrigin(y_sort_origin int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetYSortOrigin(gd.Int(y_sort_origin))
}
func (self Simple) GetYSortOrigin() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetYSortOrigin()))
}
func (self Simple) SetXDrawOrderReversed(x_draw_order_reversed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetXDrawOrderReversed(x_draw_order_reversed)
}
func (self Simple) IsXDrawOrderReversed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsXDrawOrderReversed())
}
func (self Simple) SetRenderingQuadrantSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRenderingQuadrantSize(gd.Int(size))
}
func (self Simple) GetRenderingQuadrantSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetRenderingQuadrantSize()))
}
func (self Simple) SetCollisionEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionEnabled(enabled)
}
func (self Simple) IsCollisionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCollisionEnabled())
}
func (self Simple) SetUseKinematicBodies(use_kinematic_bodies bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseKinematicBodies(use_kinematic_bodies)
}
func (self Simple) IsUsingKinematicBodies() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingKinematicBodies())
}
func (self Simple) SetCollisionVisibilityMode(visibility_mode classdb.TileMapLayerDebugVisibilityMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionVisibilityMode(visibility_mode)
}
func (self Simple) GetCollisionVisibilityMode() classdb.TileMapLayerDebugVisibilityMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TileMapLayerDebugVisibilityMode(Expert(self).GetCollisionVisibilityMode())
}
func (self Simple) SetNavigationEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationEnabled(enabled)
}
func (self Simple) IsNavigationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNavigationEnabled())
}
func (self Simple) SetNavigationMap(mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationMap(mapping)
}
func (self Simple) GetNavigationMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetNavigationMap())
}
func (self Simple) SetNavigationVisibilityMode(show_navigation classdb.TileMapLayerDebugVisibilityMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationVisibilityMode(show_navigation)
}
func (self Simple) GetNavigationVisibilityMode() classdb.TileMapLayerDebugVisibilityMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TileMapLayerDebugVisibilityMode(Expert(self).GetNavigationVisibilityMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TileMapLayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Should return [code]true[/code] if the tile at coordinates [param coords] requires a runtime update.
[b]Warning:[/b] Make sure this function only returns [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
[b]Note:[/b] If the result of this function should change, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
*/
func (class) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords gd.Vector2i) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, coords)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called with a [TileData] object about to be used internally by the [TileMapLayer], allowing its modification at runtime.
This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords].
[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the [TileMapLayer] it needs an update.
*/
func (class) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, coords gd.Vector2i, tile_data object.TileData) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,0)
		var tile_data object.TileData
		tile_data[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, coords, tile_data)
		ctx.End()
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
func (self class) SetCell(coords gd.Vector2i, source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Erases the cell at coordinates [param coords].
*/
//go:nosplit
func (self class) EraseCell(coords gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_erase_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears cells containing tiles that do not exist in the [member tile_set].
*/
//go:nosplit
func (self class) FixInvalidTiles()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_fix_invalid_tiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all cells.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tile source ID of the cell at coordinates [param coords]. Returns [code]-1[/code] if the cell does not exist.
*/
//go:nosplit
func (self class) GetCellSourceId(coords gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_cell_source_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile atlas coordinates ID of the cell at coordinates [param coords]. Returns [code]Vector2i(-1, -1)[/code] if the cell does not exist.
*/
//go:nosplit
func (self class) GetCellAtlasCoords(coords gd.Vector2i) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_cell_atlas_coords, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile alternative ID of the cell at coordinates [param coords].
*/
//go:nosplit
func (self class) GetCellAlternativeTile(coords gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_cell_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetCellTileData(ctx gd.Lifetime, coords gd.Vector2i) object.TileData {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_cell_tile_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TileData
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a [Vector2i] array with the positions of all cells containing a tile. A cell is considered empty if its source identifier equals [code]-1[/code], its atlas coordinate identifier is [code]Vector2(-1, -1)[/code] and its alternative identifier is [code]-1[/code].
*/
//go:nosplit
func (self class) GetUsedCells(ctx gd.Lifetime) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Returns a [Vector2i] array with the positions of all cells containing a tile. Tiles may be filtered according to their source ([param source_id]), their atlas coordinates ([param atlas_coords]), or alternative id ([param alternative_tile]).
If a parameter has its value set to the default one, this parameter is not used to filter a cell. Thus, if all parameters have their respective default values, this method returns the same result as [method get_used_cells].
A cell is considered empty if its source identifier equals [code]-1[/code], its atlas coordinate identifier is [code]Vector2(-1, -1)[/code] and its alternative identifier is [code]-1[/code].
*/
//go:nosplit
func (self class) GetUsedCellsById(ctx gd.Lifetime, source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_used_cells_by_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Returns a rectangle enclosing the used (non-empty) tiles of the map.
*/
//go:nosplit
func (self class) GetUsedRect() gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_used_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates and returns a new [TileMapPattern] from the given array of cells. See also [method set_pattern].
*/
//go:nosplit
func (self class) GetPattern(ctx gd.Lifetime, coords_array gd.ArrayOf[gd.Vector2i]) object.TileMapPattern {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(coords_array))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TileMapPattern
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Pastes the [TileMapPattern] at the given [param position] in the tile map. See also [method get_pattern].
*/
//go:nosplit
func (self class) SetPattern(position gd.Vector2i, pattern object.TileMapPattern)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, mmm.Get(pattern[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Update all the cells in the [param cells] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. If an updated cell has the same terrain as one of its neighboring cells, this function tries to join the two. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
[b]Note:[/b] To work correctly, this method requires the [TileMapLayer]'s TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainConnect(cells gd.ArrayOf[gd.Vector2i], terrain_set gd.Int, terrain gd.Int, ignore_empty_terrains bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(cells))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_cells_terrain_connect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Update all the cells in the [param path] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. The function will also connect two successive cell in the path with the same terrain. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
[b]Note:[/b] To work correctly, this method requires the [TileMapLayer]'s TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainPath(path gd.ArrayOf[gd.Vector2i], terrain_set gd.Int, terrain gd.Int, ignore_empty_terrains bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_cells_terrain_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the provided [param body] [RID] belongs to one of this [TileMapLayer]'s cells.
*/
//go:nosplit
func (self class) HasBodyRid(body gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_has_body_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the coordinates of the tile for given physics body [RID]. Such an [RID] can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
//go:nosplit
func (self class) GetCoordsForBodyRid(body gd.RID) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_coords_for_body_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) UpdateInternals()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_update_internals, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Notifies the [TileMapLayer] node that calls to [method _use_tile_data_runtime_update] or [method _tile_data_runtime_update] will lead to different results. This will thus trigger a [TileMapLayer] update.
[b]Warning:[/b] Updating the [TileMapLayer] is computationally expensive and may impact performance. Try to limit the number of calls to this function to avoid unnecessary update.
[b]Note:[/b] This does not trigger a direct update of the [TileMapLayer], the update will be done at the end of the frame as usual (unless you call [method update_internals]).
*/
//go:nosplit
func (self class) NotifyRuntimeTileDataUpdate()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_notify_runtime_tile_data_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns for the given coordinates [param coords_in_pattern] in a [TileMapPattern] the corresponding cell coordinates if the pattern was pasted at the [param position_in_tilemap] coordinates (see [method set_pattern]). This mapping is required as in half-offset tile shapes, the mapping might not work by calculating [code]position_in_tile_map + coords_in_pattern[/code].
*/
//go:nosplit
func (self class) MapPattern(position_in_tilemap gd.Vector2i, coords_in_pattern gd.Vector2i, pattern object.TileMapPattern) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position_in_tilemap)
	callframe.Arg(frame, coords_in_pattern)
	callframe.Arg(frame, mmm.Get(pattern[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_map_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of all neighboring cells to the one at [param coords].
*/
//go:nosplit
func (self class) GetSurroundingCells(ctx gd.Lifetime, coords gd.Vector2i) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_surrounding_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Returns the neighboring cell to the one at coordinates [param coords], identified by the [param neighbor] direction. This method takes into account the different layouts a TileMap can take.
*/
//go:nosplit
func (self class) GetNeighborCell(coords gd.Vector2i, neighbor classdb.TileSetCellNeighbor) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, neighbor)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_neighbor_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the centered position of a cell in the [TileMapLayer]'s local coordinate space. To convert the returned value into global coordinates, use [method Node2D.to_global]. See also [method local_to_map].
[b]Note:[/b] This may not correspond to the visual position of the tile, i.e. it ignores the [member TileData.texture_origin] property of individual tiles.
*/
//go:nosplit
func (self class) MapToLocal(map_position gd.Vector2i) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, map_position)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_map_to_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node2D.to_local] before passing it to this method. See also [method map_to_local].
*/
//go:nosplit
func (self class) LocalToMap(local_position gd.Vector2) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_local_to_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTileMapDataFromArray(tile_map_layer_data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tile_map_layer_data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_tile_map_data_from_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTileMapDataAsArray(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_tile_map_data_as_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTileSet(tile_set object.TileSet)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tile_set[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_tile_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTileSet(ctx gd.Lifetime) object.TileSet {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_tile_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TileSet
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetYSortOrigin(y_sort_origin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, y_sort_origin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetYSortOrigin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetXDrawOrderReversed(x_draw_order_reversed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x_draw_order_reversed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_x_draw_order_reversed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsXDrawOrderReversed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_is_x_draw_order_reversed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderingQuadrantSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRenderingQuadrantSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollisionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_is_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseKinematicBodies(use_kinematic_bodies bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_kinematic_bodies)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_use_kinematic_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingKinematicBodies() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_is_using_kinematic_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionVisibilityMode(visibility_mode classdb.TileMapLayerDebugVisibilityMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visibility_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionVisibilityMode() classdb.TileMapLayerDebugVisibilityMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TileMapLayerDebugVisibilityMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsNavigationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_is_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a custom [param map] as a [NavigationServer2D] navigation map. If not set, uses the default [World2D] navigation map instead.
*/
//go:nosplit
func (self class) SetNavigationMap(mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the [NavigationServer2D] navigation used by this [TileMapLayer].
By default this returns the default [World2D] navigation map, unless a custom map was provided using [method set_navigation_map].
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationVisibilityMode(show_navigation classdb.TileMapLayerDebugVisibilityMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, show_navigation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_set_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationVisibilityMode() classdb.TileMapLayerDebugVisibilityMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TileMapLayerDebugVisibilityMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMapLayer.Bind_get_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTileMapLayer() Expert { return self[0].AsTileMapLayer() }


//go:nosplit
func (self Simple) AsTileMapLayer() Simple { return self[0].AsTileMapLayer() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_use_tile_data_runtime_update": return reflect.ValueOf(self._use_tile_data_runtime_update);
	case "_tile_data_runtime_update": return reflect.ValueOf(self._tile_data_runtime_update);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_use_tile_data_runtime_update": return reflect.ValueOf(self._use_tile_data_runtime_update);
	case "_tile_data_runtime_update": return reflect.ValueOf(self._tile_data_runtime_update);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TileMapLayer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type DebugVisibilityMode = classdb.TileMapLayerDebugVisibilityMode

const (
/*Hide the collisions or navigation debug shapes in the editor, and use the debug settings to determine their visibility in game (i.e. [member SceneTree.debug_collisions_hint] or [member SceneTree.debug_navigation_hint]).*/
	DebugVisibilityModeDefault DebugVisibilityMode = 0
/*Always hide the collisions or navigation debug shapes.*/
	DebugVisibilityModeForceHide DebugVisibilityMode = 2
/*Always show the collisions or navigation debug shapes.*/
	DebugVisibilityModeForceShow DebugVisibilityMode = 1
)
