package TileSetAtlasSource

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TileSetSource"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An atlas is a grid of tiles laid out on a texture. Each tile in the grid must be exposed using [method create_tile]. Those tiles are then indexed using their coordinates in the grid.
Each tile can also have a size in the grid coordinates, making it more or less cells in the atlas.
Alternatives version of a tile can be created using [method create_alternative_tile], which are then indexed using an alternative ID. The main tile (the one in the grid), is accessed with an alternative ID equal to 0.
Each tile alternate has a set of properties that is defined by the source's [TileSet] layers. Those properties are stored in a TileData object that can be accessed and modified using [method get_tile_data].
As TileData properties are stored directly in the TileSetAtlasSource resource, their properties might also be set using [code]TileSetAtlasSource.set("<coords_x>:<coords_y>/<alternative_id>/<tile_data_property>")[/code].

*/
type Simple [1]classdb.TileSetAtlasSource
func (self Simple) SetTexture(texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTexture(texture)
}
func (self Simple) GetTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTexture(gc))
}
func (self Simple) SetMargins(margins gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMargins(margins)
}
func (self Simple) GetMargins() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetMargins())
}
func (self Simple) SetSeparation(separation gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSeparation(separation)
}
func (self Simple) GetSeparation() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSeparation())
}
func (self Simple) SetTextureRegionSize(texture_region_size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureRegionSize(texture_region_size)
}
func (self Simple) GetTextureRegionSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetTextureRegionSize())
}
func (self Simple) SetUseTexturePadding(use_texture_padding bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseTexturePadding(use_texture_padding)
}
func (self Simple) GetUseTexturePadding() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseTexturePadding())
}
func (self Simple) CreateTile(atlas_coords gd.Vector2i, size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateTile(atlas_coords, size)
}
func (self Simple) RemoveTile(atlas_coords gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTile(atlas_coords)
}
func (self Simple) MoveTileInAtlas(atlas_coords gd.Vector2i, new_atlas_coords gd.Vector2i, new_size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveTileInAtlas(atlas_coords, new_atlas_coords, new_size)
}
func (self Simple) GetTileSizeInAtlas(atlas_coords gd.Vector2i) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetTileSizeInAtlas(atlas_coords))
}
func (self Simple) HasRoomForTile(atlas_coords gd.Vector2i, size gd.Vector2i, animation_columns int, animation_separation gd.Vector2i, frames_count int, ignored_tile gd.Vector2i) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasRoomForTile(atlas_coords, size, gd.Int(animation_columns), animation_separation, gd.Int(frames_count), ignored_tile))
}
func (self Simple) GetTilesToBeRemovedOnChange(texture [1]classdb.Texture2D, margins gd.Vector2i, separation gd.Vector2i, texture_region_size gd.Vector2i) gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetTilesToBeRemovedOnChange(gc, texture, margins, separation, texture_region_size))
}
func (self Simple) GetTileAtCoords(atlas_coords gd.Vector2i) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetTileAtCoords(atlas_coords))
}
func (self Simple) HasTilesOutsideTexture() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasTilesOutsideTexture())
}
func (self Simple) ClearTilesOutsideTexture() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearTilesOutsideTexture()
}
func (self Simple) SetTileAnimationColumns(atlas_coords gd.Vector2i, frame_columns int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileAnimationColumns(atlas_coords, gd.Int(frame_columns))
}
func (self Simple) GetTileAnimationColumns(atlas_coords gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTileAnimationColumns(atlas_coords)))
}
func (self Simple) SetTileAnimationSeparation(atlas_coords gd.Vector2i, separation gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileAnimationSeparation(atlas_coords, separation)
}
func (self Simple) GetTileAnimationSeparation(atlas_coords gd.Vector2i) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetTileAnimationSeparation(atlas_coords))
}
func (self Simple) SetTileAnimationSpeed(atlas_coords gd.Vector2i, speed float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileAnimationSpeed(atlas_coords, gd.Float(speed))
}
func (self Simple) GetTileAnimationSpeed(atlas_coords gd.Vector2i) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTileAnimationSpeed(atlas_coords)))
}
func (self Simple) SetTileAnimationMode(atlas_coords gd.Vector2i, mode classdb.TileSetAtlasSourceTileAnimationMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileAnimationMode(atlas_coords, mode)
}
func (self Simple) GetTileAnimationMode(atlas_coords gd.Vector2i) classdb.TileSetAtlasSourceTileAnimationMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TileSetAtlasSourceTileAnimationMode(Expert(self).GetTileAnimationMode(atlas_coords))
}
func (self Simple) SetTileAnimationFramesCount(atlas_coords gd.Vector2i, frames_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileAnimationFramesCount(atlas_coords, gd.Int(frames_count))
}
func (self Simple) GetTileAnimationFramesCount(atlas_coords gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTileAnimationFramesCount(atlas_coords)))
}
func (self Simple) SetTileAnimationFrameDuration(atlas_coords gd.Vector2i, frame_index int, duration float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileAnimationFrameDuration(atlas_coords, gd.Int(frame_index), gd.Float(duration))
}
func (self Simple) GetTileAnimationFrameDuration(atlas_coords gd.Vector2i, frame_index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTileAnimationFrameDuration(atlas_coords, gd.Int(frame_index))))
}
func (self Simple) GetTileAnimationTotalDuration(atlas_coords gd.Vector2i) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTileAnimationTotalDuration(atlas_coords)))
}
func (self Simple) CreateAlternativeTile(atlas_coords gd.Vector2i, alternative_id_override int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).CreateAlternativeTile(atlas_coords, gd.Int(alternative_id_override))))
}
func (self Simple) RemoveAlternativeTile(atlas_coords gd.Vector2i, alternative_tile int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveAlternativeTile(atlas_coords, gd.Int(alternative_tile))
}
func (self Simple) SetAlternativeTileId(atlas_coords gd.Vector2i, alternative_tile int, new_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlternativeTileId(atlas_coords, gd.Int(alternative_tile), gd.Int(new_id))
}
func (self Simple) GetNextAlternativeTileId(atlas_coords gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetNextAlternativeTileId(atlas_coords)))
}
func (self Simple) GetTileData(atlas_coords gd.Vector2i, alternative_tile int) [1]classdb.TileData {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TileData(Expert(self).GetTileData(gc, atlas_coords, gd.Int(alternative_tile)))
}
func (self Simple) GetAtlasGridSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetAtlasGridSize())
}
func (self Simple) GetTileTextureRegion(atlas_coords gd.Vector2i, frame_ int) gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).GetTileTextureRegion(atlas_coords, gd.Int(frame_)))
}
func (self Simple) GetRuntimeTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetRuntimeTexture(gc))
}
func (self Simple) GetRuntimeTileTextureRegion(atlas_coords gd.Vector2i, frame_ int) gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).GetRuntimeTileTextureRegion(atlas_coords, gd.Int(frame_)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TileSetAtlasSource
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTexture(texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMargins(margins gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margins)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_margins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMargins() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_margins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSeparation(separation gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, separation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_separation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSeparation() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_separation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureRegionSize(texture_region_size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture_region_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_texture_region_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureRegionSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_texture_region_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseTexturePadding(use_texture_padding bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_texture_padding)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_use_texture_padding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseTexturePadding() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_use_texture_padding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new tile at coordinates [param atlas_coords] with the given [param size].
*/
//go:nosplit
func (self class) CreateTile(atlas_coords gd.Vector2i, size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_create_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove a tile and its alternative at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) RemoveTile(atlas_coords gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_remove_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Move the tile and its alternatives at the [param atlas_coords] coordinates to the [param new_atlas_coords] coordinates with the [param new_size] size. This functions will fail if a tile is already present in the given area.
If [param new_atlas_coords] is [code]Vector2i(-1, -1)[/code], keeps the tile's coordinates. If [param new_size] is [code]Vector2i(-1, -1)[/code], keeps the tile's size.
To avoid an error, first check if a move is possible using [method has_room_for_tile].
*/
//go:nosplit
func (self class) MoveTileInAtlas(atlas_coords gd.Vector2i, new_atlas_coords gd.Vector2i, new_size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, new_atlas_coords)
	callframe.Arg(frame, new_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_move_tile_in_atlas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of the tile (in the grid coordinates system) at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileSizeInAtlas(atlas_coords gd.Vector2i) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_size_in_atlas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether there is enough room in an atlas to create/modify a tile with the given properties. If [param ignored_tile] is provided, act as is the given tile was not present in the atlas. This may be used when you want to modify a tile's properties.
*/
//go:nosplit
func (self class) HasRoomForTile(atlas_coords gd.Vector2i, size gd.Vector2i, animation_columns gd.Int, animation_separation gd.Vector2i, frames_count gd.Int, ignored_tile gd.Vector2i) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, size)
	callframe.Arg(frame, animation_columns)
	callframe.Arg(frame, animation_separation)
	callframe.Arg(frame, frames_count)
	callframe.Arg(frame, ignored_tile)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_has_room_for_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of tiles coordinates ID that will be automatically removed when modifying one or several of those properties: [param texture], [param margins], [param separation] or [param texture_region_size]. This can be used to undo changes that would have caused tiles data loss.
*/
//go:nosplit
func (self class) GetTilesToBeRemovedOnChange(ctx gd.Lifetime, texture object.Texture2D, margins gd.Vector2i, separation gd.Vector2i, texture_region_size gd.Vector2i) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, margins)
	callframe.Arg(frame, separation)
	callframe.Arg(frame, texture_region_size)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tiles_to_be_removed_on_change, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If there is a tile covering the [param atlas_coords] coordinates, returns the top-left coordinates of the tile (thus its coordinate ID). Returns [code]Vector2i(-1, -1)[/code] otherwise.
*/
//go:nosplit
func (self class) GetTileAtCoords(atlas_coords gd.Vector2i) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_at_coords, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Checks if the source has any tiles that don't fit the texture area (either partially or completely).
*/
//go:nosplit
func (self class) HasTilesOutsideTexture() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_has_tiles_outside_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all tiles that don't fit the available texture area. This method iterates over all the source's tiles, so it's advised to use [method has_tiles_outside_texture] beforehand.
*/
//go:nosplit
func (self class) ClearTilesOutsideTexture()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_clear_tiles_outside_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the number of columns in the animation layout of the tile at coordinates [param atlas_coords]. If set to 0, then the different frames of the animation are laid out as a single horizontal line in the atlas.
*/
//go:nosplit
func (self class) SetTileAnimationColumns(atlas_coords gd.Vector2i, frame_columns gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_columns)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_tile_animation_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns how many columns the tile at [param atlas_coords] has in its animation layout.
*/
//go:nosplit
func (self class) GetTileAnimationColumns(atlas_coords gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_animation_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the margin (in grid tiles) between each tile in the animation layout of the tile at coordinates [param atlas_coords] has.
*/
//go:nosplit
func (self class) SetTileAnimationSeparation(atlas_coords gd.Vector2i, separation gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, separation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_tile_animation_separation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the separation (as in the atlas grid) between each frame of an animated tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationSeparation(atlas_coords gd.Vector2i) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_animation_separation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the animation speed of the tile at coordinates [param atlas_coords] has.
*/
//go:nosplit
func (self class) SetTileAnimationSpeed(atlas_coords gd.Vector2i, speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_tile_animation_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the animation speed of the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationSpeed(atlas_coords gd.Vector2i) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_animation_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tile animation mode of the tile at [param atlas_coords] to [param mode]. See also [method get_tile_animation_mode].
*/
//go:nosplit
func (self class) SetTileAnimationMode(atlas_coords gd.Vector2i, mode classdb.TileSetAtlasSourceTileAnimationMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_tile_animation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tile animation mode of the tile at [param atlas_coords]. See also [method set_tile_animation_mode].
*/
//go:nosplit
func (self class) GetTileAnimationMode(atlas_coords gd.Vector2i) classdb.TileSetAtlasSourceTileAnimationMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[classdb.TileSetAtlasSourceTileAnimationMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_animation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets how many animation frames the tile at coordinates [param atlas_coords] has.
*/
//go:nosplit
func (self class) SetTileAnimationFramesCount(atlas_coords gd.Vector2i, frames_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frames_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_tile_animation_frames_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns how many animation frames has the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationFramesCount(atlas_coords gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_animation_frames_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the animation frame [param duration] of frame [param frame_index] for the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) SetTileAnimationFrameDuration(atlas_coords gd.Vector2i, frame_index gd.Int, duration gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_index)
	callframe.Arg(frame, duration)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_tile_animation_frame_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the animation frame duration of frame [param frame_index] for the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationFrameDuration(atlas_coords gd.Vector2i, frame_index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_animation_frame_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the sum of the sum of the frame durations of the tile at coordinates [param atlas_coords]. This value needs to be divided by the animation speed to get the actual animation loop duration.
*/
//go:nosplit
func (self class) GetTileAnimationTotalDuration(atlas_coords gd.Vector2i) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_animation_total_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an alternative tile for the tile at coordinates [param atlas_coords]. If [param alternative_id_override] is -1, give it an automatically generated unique ID, or assigns it the given ID otherwise.
Returns the new alternative identifier, or -1 if the alternative could not be created with a provided [param alternative_id_override].
*/
//go:nosplit
func (self class) CreateAlternativeTile(atlas_coords gd.Vector2i, alternative_id_override gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_id_override)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_create_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Remove a tile's alternative with alternative ID [param alternative_tile].
Calling this function with [param alternative_tile] equals to 0 will fail, as the base tile alternative cannot be removed.
*/
//go:nosplit
func (self class) RemoveAlternativeTile(atlas_coords gd.Vector2i, alternative_tile gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_remove_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Change a tile's alternative ID from [param alternative_tile] to [param new_id].
Calling this function with [param new_id] of 0 will fail, as the base tile alternative cannot be moved.
*/
//go:nosplit
func (self class) SetAlternativeTileId(atlas_coords gd.Vector2i, alternative_tile gd.Int, new_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	callframe.Arg(frame, new_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_set_alternative_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the alternative ID a following call to [method create_alternative_tile] would return.
*/
//go:nosplit
func (self class) GetNextAlternativeTileId(atlas_coords gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_next_alternative_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [TileData] object for the given atlas coordinates and alternative ID.
*/
//go:nosplit
func (self class) GetTileData(ctx gd.Lifetime, atlas_coords gd.Vector2i, alternative_tile gd.Int) object.TileData {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TileData
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the atlas grid size, which depends on how many tiles can fit in the texture. It thus depends on the [member texture]'s size, the atlas [member margins], and the tiles' [member texture_region_size].
*/
//go:nosplit
func (self class) GetAtlasGridSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_atlas_grid_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a tile's texture region in the atlas texture. For animated tiles, a [param frame] argument might be provided for the different frames of the animation.
*/
//go:nosplit
func (self class) GetTileTextureRegion(atlas_coords gd.Vector2i, frame_ gd.Int) gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_tile_texture_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [member use_texture_padding] is [code]false[/code], returns [member texture]. Otherwise, returns and internal [ImageTexture] created that includes the padding.
*/
//go:nosplit
func (self class) GetRuntimeTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_runtime_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the region of the tile at coordinates [param atlas_coords] for the given [param frame] inside the texture returned by [method get_runtime_texture].
[b]Note:[/b] If [member use_texture_padding] is [code]false[/code], returns the same as [method get_tile_texture_region].
*/
//go:nosplit
func (self class) GetRuntimeTileTextureRegion(atlas_coords gd.Vector2i, frame_ gd.Int) gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetAtlasSource.Bind_get_runtime_tile_texture_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTileSetAtlasSource() Expert { return self[0].AsTileSetAtlasSource() }


//go:nosplit
func (self Simple) AsTileSetAtlasSource() Simple { return self[0].AsTileSetAtlasSource() }


//go:nosplit
func (self class) AsTileSetSource() TileSetSource.Expert { return self[0].AsTileSetSource() }


//go:nosplit
func (self Simple) AsTileSetSource() TileSetSource.Simple { return self[0].AsTileSetSource() }


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
func init() {classdb.Register("TileSetAtlasSource", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TileAnimationMode = classdb.TileSetAtlasSourceTileAnimationMode

const (
/*Tile animations start at same time, looking identical.*/
	TileAnimationModeDefault TileAnimationMode = 0
/*Tile animations start at random times, looking varied.*/
	TileAnimationModeRandomStartTimes TileAnimationMode = 1
/*Represents the size of the [enum TileAnimationMode] enum.*/
	TileAnimationModeMax TileAnimationMode = 2
)
