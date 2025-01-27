// Package TileSetAtlasSource provides methods for working with TileSetAtlasSource object instances.
package TileSetAtlasSource

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/TileSetSource"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
An atlas is a grid of tiles laid out on a texture. Each tile in the grid must be exposed using [method create_tile]. Those tiles are then indexed using their coordinates in the grid.
Each tile can also have a size in the grid coordinates, making it more or less cells in the atlas.
Alternatives version of a tile can be created using [method create_alternative_tile], which are then indexed using an alternative ID. The main tile (the one in the grid), is accessed with an alternative ID equal to 0.
Each tile alternate has a set of properties that is defined by the source's [TileSet] layers. Those properties are stored in a TileData object that can be accessed and modified using [method get_tile_data].
As TileData properties are stored directly in the TileSetAtlasSource resource, their properties might also be set using [code]TileSetAtlasSource.set("<coords_x>:<coords_y>/<alternative_id>/<tile_data_property>")[/code].
*/
type Instance [1]gdclass.TileSetAtlasSource

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTileSetAtlasSource() Instance
}

/*
Creates a new tile at coordinates [param atlas_coords] with the given [param size].
*/
func (self Instance) CreateTile(atlas_coords Vector2i.XY) { //gd:TileSetAtlasSource.create_tile
	class(self).CreateTile(gd.Vector2i(atlas_coords), gd.Vector2i(gd.Vector2i{1, 1}))
}

/*
Remove a tile and its alternative at coordinates [param atlas_coords].
*/
func (self Instance) RemoveTile(atlas_coords Vector2i.XY) { //gd:TileSetAtlasSource.remove_tile
	class(self).RemoveTile(gd.Vector2i(atlas_coords))
}

/*
Move the tile and its alternatives at the [param atlas_coords] coordinates to the [param new_atlas_coords] coordinates with the [param new_size] size. This functions will fail if a tile is already present in the given area.
If [param new_atlas_coords] is [code]Vector2i(-1, -1)[/code], keeps the tile's coordinates. If [param new_size] is [code]Vector2i(-1, -1)[/code], keeps the tile's size.
To avoid an error, first check if a move is possible using [method has_room_for_tile].
*/
func (self Instance) MoveTileInAtlas(atlas_coords Vector2i.XY) { //gd:TileSetAtlasSource.move_tile_in_atlas
	class(self).MoveTileInAtlas(gd.Vector2i(atlas_coords), gd.Vector2i(gd.Vector2i{-1, -1}), gd.Vector2i(gd.Vector2i{-1, -1}))
}

/*
Returns the size of the tile (in the grid coordinates system) at coordinates [param atlas_coords].
*/
func (self Instance) GetTileSizeInAtlas(atlas_coords Vector2i.XY) Vector2i.XY { //gd:TileSetAtlasSource.get_tile_size_in_atlas
	return Vector2i.XY(class(self).GetTileSizeInAtlas(gd.Vector2i(atlas_coords)))
}

/*
Returns whether there is enough room in an atlas to create/modify a tile with the given properties. If [param ignored_tile] is provided, act as is the given tile was not present in the atlas. This may be used when you want to modify a tile's properties.
*/
func (self Instance) HasRoomForTile(atlas_coords Vector2i.XY, size Vector2i.XY, animation_columns int, animation_separation Vector2i.XY, frames_count int) bool { //gd:TileSetAtlasSource.has_room_for_tile
	return bool(class(self).HasRoomForTile(gd.Vector2i(atlas_coords), gd.Vector2i(size), gd.Int(animation_columns), gd.Vector2i(animation_separation), gd.Int(frames_count), gd.Vector2i(gd.Vector2i{-1, -1})))
}

/*
Returns an array of tiles coordinates ID that will be automatically removed when modifying one or several of those properties: [param texture], [param margins], [param separation] or [param texture_region_size]. This can be used to undo changes that would have caused tiles data loss.
*/
func (self Instance) GetTilesToBeRemovedOnChange(texture [1]gdclass.Texture2D, margins Vector2i.XY, separation Vector2i.XY, texture_region_size Vector2i.XY) []Vector2.XY { //gd:TileSetAtlasSource.get_tiles_to_be_removed_on_change
	return []Vector2.XY(slices.Collect(class(self).GetTilesToBeRemovedOnChange(texture, gd.Vector2i(margins), gd.Vector2i(separation), gd.Vector2i(texture_region_size)).Values()))
}

/*
If there is a tile covering the [param atlas_coords] coordinates, returns the top-left coordinates of the tile (thus its coordinate ID). Returns [code]Vector2i(-1, -1)[/code] otherwise.
*/
func (self Instance) GetTileAtCoords(atlas_coords Vector2i.XY) Vector2i.XY { //gd:TileSetAtlasSource.get_tile_at_coords
	return Vector2i.XY(class(self).GetTileAtCoords(gd.Vector2i(atlas_coords)))
}

/*
Checks if the source has any tiles that don't fit the texture area (either partially or completely).
*/
func (self Instance) HasTilesOutsideTexture() bool { //gd:TileSetAtlasSource.has_tiles_outside_texture
	return bool(class(self).HasTilesOutsideTexture())
}

/*
Removes all tiles that don't fit the available texture area. This method iterates over all the source's tiles, so it's advised to use [method has_tiles_outside_texture] beforehand.
*/
func (self Instance) ClearTilesOutsideTexture() { //gd:TileSetAtlasSource.clear_tiles_outside_texture
	class(self).ClearTilesOutsideTexture()
}

/*
Sets the number of columns in the animation layout of the tile at coordinates [param atlas_coords]. If set to 0, then the different frames of the animation are laid out as a single horizontal line in the atlas.
*/
func (self Instance) SetTileAnimationColumns(atlas_coords Vector2i.XY, frame_columns int) { //gd:TileSetAtlasSource.set_tile_animation_columns
	class(self).SetTileAnimationColumns(gd.Vector2i(atlas_coords), gd.Int(frame_columns))
}

/*
Returns how many columns the tile at [param atlas_coords] has in its animation layout.
*/
func (self Instance) GetTileAnimationColumns(atlas_coords Vector2i.XY) int { //gd:TileSetAtlasSource.get_tile_animation_columns
	return int(int(class(self).GetTileAnimationColumns(gd.Vector2i(atlas_coords))))
}

/*
Sets the margin (in grid tiles) between each tile in the animation layout of the tile at coordinates [param atlas_coords] has.
*/
func (self Instance) SetTileAnimationSeparation(atlas_coords Vector2i.XY, separation Vector2i.XY) { //gd:TileSetAtlasSource.set_tile_animation_separation
	class(self).SetTileAnimationSeparation(gd.Vector2i(atlas_coords), gd.Vector2i(separation))
}

/*
Returns the separation (as in the atlas grid) between each frame of an animated tile at coordinates [param atlas_coords].
*/
func (self Instance) GetTileAnimationSeparation(atlas_coords Vector2i.XY) Vector2i.XY { //gd:TileSetAtlasSource.get_tile_animation_separation
	return Vector2i.XY(class(self).GetTileAnimationSeparation(gd.Vector2i(atlas_coords)))
}

/*
Sets the animation speed of the tile at coordinates [param atlas_coords] has.
*/
func (self Instance) SetTileAnimationSpeed(atlas_coords Vector2i.XY, speed Float.X) { //gd:TileSetAtlasSource.set_tile_animation_speed
	class(self).SetTileAnimationSpeed(gd.Vector2i(atlas_coords), gd.Float(speed))
}

/*
Returns the animation speed of the tile at coordinates [param atlas_coords].
*/
func (self Instance) GetTileAnimationSpeed(atlas_coords Vector2i.XY) Float.X { //gd:TileSetAtlasSource.get_tile_animation_speed
	return Float.X(Float.X(class(self).GetTileAnimationSpeed(gd.Vector2i(atlas_coords))))
}

/*
Sets the tile animation mode of the tile at [param atlas_coords] to [param mode]. See also [method get_tile_animation_mode].
*/
func (self Instance) SetTileAnimationMode(atlas_coords Vector2i.XY, mode gdclass.TileSetAtlasSourceTileAnimationMode) { //gd:TileSetAtlasSource.set_tile_animation_mode
	class(self).SetTileAnimationMode(gd.Vector2i(atlas_coords), mode)
}

/*
Returns the tile animation mode of the tile at [param atlas_coords]. See also [method set_tile_animation_mode].
*/
func (self Instance) GetTileAnimationMode(atlas_coords Vector2i.XY) gdclass.TileSetAtlasSourceTileAnimationMode { //gd:TileSetAtlasSource.get_tile_animation_mode
	return gdclass.TileSetAtlasSourceTileAnimationMode(class(self).GetTileAnimationMode(gd.Vector2i(atlas_coords)))
}

/*
Sets how many animation frames the tile at coordinates [param atlas_coords] has.
*/
func (self Instance) SetTileAnimationFramesCount(atlas_coords Vector2i.XY, frames_count int) { //gd:TileSetAtlasSource.set_tile_animation_frames_count
	class(self).SetTileAnimationFramesCount(gd.Vector2i(atlas_coords), gd.Int(frames_count))
}

/*
Returns how many animation frames has the tile at coordinates [param atlas_coords].
*/
func (self Instance) GetTileAnimationFramesCount(atlas_coords Vector2i.XY) int { //gd:TileSetAtlasSource.get_tile_animation_frames_count
	return int(int(class(self).GetTileAnimationFramesCount(gd.Vector2i(atlas_coords))))
}

/*
Sets the animation frame [param duration] of frame [param frame_index] for the tile at coordinates [param atlas_coords].
*/
func (self Instance) SetTileAnimationFrameDuration(atlas_coords Vector2i.XY, frame_index int, duration Float.X) { //gd:TileSetAtlasSource.set_tile_animation_frame_duration
	class(self).SetTileAnimationFrameDuration(gd.Vector2i(atlas_coords), gd.Int(frame_index), gd.Float(duration))
}

/*
Returns the animation frame duration of frame [param frame_index] for the tile at coordinates [param atlas_coords].
*/
func (self Instance) GetTileAnimationFrameDuration(atlas_coords Vector2i.XY, frame_index int) Float.X { //gd:TileSetAtlasSource.get_tile_animation_frame_duration
	return Float.X(Float.X(class(self).GetTileAnimationFrameDuration(gd.Vector2i(atlas_coords), gd.Int(frame_index))))
}

/*
Returns the sum of the sum of the frame durations of the tile at coordinates [param atlas_coords]. This value needs to be divided by the animation speed to get the actual animation loop duration.
*/
func (self Instance) GetTileAnimationTotalDuration(atlas_coords Vector2i.XY) Float.X { //gd:TileSetAtlasSource.get_tile_animation_total_duration
	return Float.X(Float.X(class(self).GetTileAnimationTotalDuration(gd.Vector2i(atlas_coords))))
}

/*
Creates an alternative tile for the tile at coordinates [param atlas_coords]. If [param alternative_id_override] is -1, give it an automatically generated unique ID, or assigns it the given ID otherwise.
Returns the new alternative identifier, or -1 if the alternative could not be created with a provided [param alternative_id_override].
*/
func (self Instance) CreateAlternativeTile(atlas_coords Vector2i.XY) int { //gd:TileSetAtlasSource.create_alternative_tile
	return int(int(class(self).CreateAlternativeTile(gd.Vector2i(atlas_coords), gd.Int(-1))))
}

/*
Remove a tile's alternative with alternative ID [param alternative_tile].
Calling this function with [param alternative_tile] equals to 0 will fail, as the base tile alternative cannot be removed.
*/
func (self Instance) RemoveAlternativeTile(atlas_coords Vector2i.XY, alternative_tile int) { //gd:TileSetAtlasSource.remove_alternative_tile
	class(self).RemoveAlternativeTile(gd.Vector2i(atlas_coords), gd.Int(alternative_tile))
}

/*
Change a tile's alternative ID from [param alternative_tile] to [param new_id].
Calling this function with [param new_id] of 0 will fail, as the base tile alternative cannot be moved.
*/
func (self Instance) SetAlternativeTileId(atlas_coords Vector2i.XY, alternative_tile int, new_id int) { //gd:TileSetAtlasSource.set_alternative_tile_id
	class(self).SetAlternativeTileId(gd.Vector2i(atlas_coords), gd.Int(alternative_tile), gd.Int(new_id))
}

/*
Returns the alternative ID a following call to [method create_alternative_tile] would return.
*/
func (self Instance) GetNextAlternativeTileId(atlas_coords Vector2i.XY) int { //gd:TileSetAtlasSource.get_next_alternative_tile_id
	return int(int(class(self).GetNextAlternativeTileId(gd.Vector2i(atlas_coords))))
}

/*
Returns the [TileData] object for the given atlas coordinates and alternative ID.
*/
func (self Instance) GetTileData(atlas_coords Vector2i.XY, alternative_tile int) [1]gdclass.TileData { //gd:TileSetAtlasSource.get_tile_data
	return [1]gdclass.TileData(class(self).GetTileData(gd.Vector2i(atlas_coords), gd.Int(alternative_tile)))
}

/*
Returns the atlas grid size, which depends on how many tiles can fit in the texture. It thus depends on the [member texture]'s size, the atlas [member margins], and the tiles' [member texture_region_size].
*/
func (self Instance) GetAtlasGridSize() Vector2i.XY { //gd:TileSetAtlasSource.get_atlas_grid_size
	return Vector2i.XY(class(self).GetAtlasGridSize())
}

/*
Returns a tile's texture region in the atlas texture. For animated tiles, a [param frame] argument might be provided for the different frames of the animation.
*/
func (self Instance) GetTileTextureRegion(atlas_coords Vector2i.XY) Rect2i.PositionSize { //gd:TileSetAtlasSource.get_tile_texture_region
	return Rect2i.PositionSize(class(self).GetTileTextureRegion(gd.Vector2i(atlas_coords), gd.Int(0)))
}

/*
If [member use_texture_padding] is [code]false[/code], returns [member texture]. Otherwise, returns and internal [ImageTexture] created that includes the padding.
*/
func (self Instance) GetRuntimeTexture() [1]gdclass.Texture2D { //gd:TileSetAtlasSource.get_runtime_texture
	return [1]gdclass.Texture2D(class(self).GetRuntimeTexture())
}

/*
Returns the region of the tile at coordinates [param atlas_coords] for the given [param frame] inside the texture returned by [method get_runtime_texture].
[b]Note:[/b] If [member use_texture_padding] is [code]false[/code], returns the same as [method get_tile_texture_region].
*/
func (self Instance) GetRuntimeTileTextureRegion(atlas_coords Vector2i.XY, frame_ int) Rect2i.PositionSize { //gd:TileSetAtlasSource.get_runtime_tile_texture_region
	return Rect2i.PositionSize(class(self).GetRuntimeTileTextureRegion(gd.Vector2i(atlas_coords), gd.Int(frame_)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TileSetAtlasSource

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileSetAtlasSource"))
	casted := Instance{*(*gdclass.TileSetAtlasSource)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Texture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) Margins() Vector2i.XY {
	return Vector2i.XY(class(self).GetMargins())
}

func (self Instance) SetMargins(value Vector2i.XY) {
	class(self).SetMargins(gd.Vector2i(value))
}

func (self Instance) Separation() Vector2i.XY {
	return Vector2i.XY(class(self).GetSeparation())
}

func (self Instance) SetSeparation(value Vector2i.XY) {
	class(self).SetSeparation(gd.Vector2i(value))
}

func (self Instance) TextureRegionSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetTextureRegionSize())
}

func (self Instance) SetTextureRegionSize(value Vector2i.XY) {
	class(self).SetTextureRegionSize(gd.Vector2i(value))
}

func (self Instance) UseTexturePadding() bool {
	return bool(class(self).GetUseTexturePadding())
}

func (self Instance) SetUseTexturePadding(value bool) {
	class(self).SetUseTexturePadding(value)
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) { //gd:TileSetAtlasSource.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D { //gd:TileSetAtlasSource.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargins(margins gd.Vector2i) { //gd:TileSetAtlasSource.set_margins
	var frame = callframe.New()
	callframe.Arg(frame, margins)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_margins, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargins() gd.Vector2i { //gd:TileSetAtlasSource.get_margins
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_margins, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSeparation(separation gd.Vector2i) { //gd:TileSetAtlasSource.set_separation
	var frame = callframe.New()
	callframe.Arg(frame, separation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_separation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeparation() gd.Vector2i { //gd:TileSetAtlasSource.get_separation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_separation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRegionSize(texture_region_size gd.Vector2i) { //gd:TileSetAtlasSource.set_texture_region_size
	var frame = callframe.New()
	callframe.Arg(frame, texture_region_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_texture_region_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRegionSize() gd.Vector2i { //gd:TileSetAtlasSource.get_texture_region_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_texture_region_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseTexturePadding(use_texture_padding bool) { //gd:TileSetAtlasSource.set_use_texture_padding
	var frame = callframe.New()
	callframe.Arg(frame, use_texture_padding)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_use_texture_padding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseTexturePadding() bool { //gd:TileSetAtlasSource.get_use_texture_padding
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_use_texture_padding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new tile at coordinates [param atlas_coords] with the given [param size].
*/
//go:nosplit
func (self class) CreateTile(atlas_coords gd.Vector2i, size gd.Vector2i) { //gd:TileSetAtlasSource.create_tile
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_create_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Remove a tile and its alternative at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) RemoveTile(atlas_coords gd.Vector2i) { //gd:TileSetAtlasSource.remove_tile
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_remove_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Move the tile and its alternatives at the [param atlas_coords] coordinates to the [param new_atlas_coords] coordinates with the [param new_size] size. This functions will fail if a tile is already present in the given area.
If [param new_atlas_coords] is [code]Vector2i(-1, -1)[/code], keeps the tile's coordinates. If [param new_size] is [code]Vector2i(-1, -1)[/code], keeps the tile's size.
To avoid an error, first check if a move is possible using [method has_room_for_tile].
*/
//go:nosplit
func (self class) MoveTileInAtlas(atlas_coords gd.Vector2i, new_atlas_coords gd.Vector2i, new_size gd.Vector2i) { //gd:TileSetAtlasSource.move_tile_in_atlas
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, new_atlas_coords)
	callframe.Arg(frame, new_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_move_tile_in_atlas, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the size of the tile (in the grid coordinates system) at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileSizeInAtlas(atlas_coords gd.Vector2i) gd.Vector2i { //gd:TileSetAtlasSource.get_tile_size_in_atlas
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_size_in_atlas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether there is enough room in an atlas to create/modify a tile with the given properties. If [param ignored_tile] is provided, act as is the given tile was not present in the atlas. This may be used when you want to modify a tile's properties.
*/
//go:nosplit
func (self class) HasRoomForTile(atlas_coords gd.Vector2i, size gd.Vector2i, animation_columns gd.Int, animation_separation gd.Vector2i, frames_count gd.Int, ignored_tile gd.Vector2i) bool { //gd:TileSetAtlasSource.has_room_for_tile
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, size)
	callframe.Arg(frame, animation_columns)
	callframe.Arg(frame, animation_separation)
	callframe.Arg(frame, frames_count)
	callframe.Arg(frame, ignored_tile)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_has_room_for_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of tiles coordinates ID that will be automatically removed when modifying one or several of those properties: [param texture], [param margins], [param separation] or [param texture_region_size]. This can be used to undo changes that would have caused tiles data loss.
*/
//go:nosplit
func (self class) GetTilesToBeRemovedOnChange(texture [1]gdclass.Texture2D, margins gd.Vector2i, separation gd.Vector2i, texture_region_size gd.Vector2i) Packed.Array[Vector2.XY] { //gd:TileSetAtlasSource.get_tiles_to_be_removed_on_change
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, margins)
	callframe.Arg(frame, separation)
	callframe.Arg(frame, texture_region_size)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tiles_to_be_removed_on_change, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
If there is a tile covering the [param atlas_coords] coordinates, returns the top-left coordinates of the tile (thus its coordinate ID). Returns [code]Vector2i(-1, -1)[/code] otherwise.
*/
//go:nosplit
func (self class) GetTileAtCoords(atlas_coords gd.Vector2i) gd.Vector2i { //gd:TileSetAtlasSource.get_tile_at_coords
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_at_coords, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Checks if the source has any tiles that don't fit the texture area (either partially or completely).
*/
//go:nosplit
func (self class) HasTilesOutsideTexture() bool { //gd:TileSetAtlasSource.has_tiles_outside_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_has_tiles_outside_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes all tiles that don't fit the available texture area. This method iterates over all the source's tiles, so it's advised to use [method has_tiles_outside_texture] beforehand.
*/
//go:nosplit
func (self class) ClearTilesOutsideTexture() { //gd:TileSetAtlasSource.clear_tiles_outside_texture
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_clear_tiles_outside_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the number of columns in the animation layout of the tile at coordinates [param atlas_coords]. If set to 0, then the different frames of the animation are laid out as a single horizontal line in the atlas.
*/
//go:nosplit
func (self class) SetTileAnimationColumns(atlas_coords gd.Vector2i, frame_columns gd.Int) { //gd:TileSetAtlasSource.set_tile_animation_columns
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_columns)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_tile_animation_columns, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns how many columns the tile at [param atlas_coords] has in its animation layout.
*/
//go:nosplit
func (self class) GetTileAnimationColumns(atlas_coords gd.Vector2i) gd.Int { //gd:TileSetAtlasSource.get_tile_animation_columns
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_animation_columns, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the margin (in grid tiles) between each tile in the animation layout of the tile at coordinates [param atlas_coords] has.
*/
//go:nosplit
func (self class) SetTileAnimationSeparation(atlas_coords gd.Vector2i, separation gd.Vector2i) { //gd:TileSetAtlasSource.set_tile_animation_separation
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, separation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_tile_animation_separation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the separation (as in the atlas grid) between each frame of an animated tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationSeparation(atlas_coords gd.Vector2i) gd.Vector2i { //gd:TileSetAtlasSource.get_tile_animation_separation
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_animation_separation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the animation speed of the tile at coordinates [param atlas_coords] has.
*/
//go:nosplit
func (self class) SetTileAnimationSpeed(atlas_coords gd.Vector2i, speed gd.Float) { //gd:TileSetAtlasSource.set_tile_animation_speed
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_tile_animation_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the animation speed of the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationSpeed(atlas_coords gd.Vector2i) gd.Float { //gd:TileSetAtlasSource.get_tile_animation_speed
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_animation_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tile animation mode of the tile at [param atlas_coords] to [param mode]. See also [method get_tile_animation_mode].
*/
//go:nosplit
func (self class) SetTileAnimationMode(atlas_coords gd.Vector2i, mode gdclass.TileSetAtlasSourceTileAnimationMode) { //gd:TileSetAtlasSource.set_tile_animation_mode
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_tile_animation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tile animation mode of the tile at [param atlas_coords]. See also [method set_tile_animation_mode].
*/
//go:nosplit
func (self class) GetTileAnimationMode(atlas_coords gd.Vector2i) gdclass.TileSetAtlasSourceTileAnimationMode { //gd:TileSetAtlasSource.get_tile_animation_mode
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gdclass.TileSetAtlasSourceTileAnimationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_animation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets how many animation frames the tile at coordinates [param atlas_coords] has.
*/
//go:nosplit
func (self class) SetTileAnimationFramesCount(atlas_coords gd.Vector2i, frames_count gd.Int) { //gd:TileSetAtlasSource.set_tile_animation_frames_count
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frames_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_tile_animation_frames_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns how many animation frames has the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationFramesCount(atlas_coords gd.Vector2i) gd.Int { //gd:TileSetAtlasSource.get_tile_animation_frames_count
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_animation_frames_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the animation frame [param duration] of frame [param frame_index] for the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) SetTileAnimationFrameDuration(atlas_coords gd.Vector2i, frame_index gd.Int, duration gd.Float) { //gd:TileSetAtlasSource.set_tile_animation_frame_duration
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_index)
	callframe.Arg(frame, duration)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_tile_animation_frame_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the animation frame duration of frame [param frame_index] for the tile at coordinates [param atlas_coords].
*/
//go:nosplit
func (self class) GetTileAnimationFrameDuration(atlas_coords gd.Vector2i, frame_index gd.Int) gd.Float { //gd:TileSetAtlasSource.get_tile_animation_frame_duration
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_animation_frame_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the sum of the sum of the frame durations of the tile at coordinates [param atlas_coords]. This value needs to be divided by the animation speed to get the actual animation loop duration.
*/
//go:nosplit
func (self class) GetTileAnimationTotalDuration(atlas_coords gd.Vector2i) gd.Float { //gd:TileSetAtlasSource.get_tile_animation_total_duration
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_animation_total_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates an alternative tile for the tile at coordinates [param atlas_coords]. If [param alternative_id_override] is -1, give it an automatically generated unique ID, or assigns it the given ID otherwise.
Returns the new alternative identifier, or -1 if the alternative could not be created with a provided [param alternative_id_override].
*/
//go:nosplit
func (self class) CreateAlternativeTile(atlas_coords gd.Vector2i, alternative_id_override gd.Int) gd.Int { //gd:TileSetAtlasSource.create_alternative_tile
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_id_override)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_create_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Remove a tile's alternative with alternative ID [param alternative_tile].
Calling this function with [param alternative_tile] equals to 0 will fail, as the base tile alternative cannot be removed.
*/
//go:nosplit
func (self class) RemoveAlternativeTile(atlas_coords gd.Vector2i, alternative_tile gd.Int) { //gd:TileSetAtlasSource.remove_alternative_tile
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_remove_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Change a tile's alternative ID from [param alternative_tile] to [param new_id].
Calling this function with [param new_id] of 0 will fail, as the base tile alternative cannot be moved.
*/
//go:nosplit
func (self class) SetAlternativeTileId(atlas_coords gd.Vector2i, alternative_tile gd.Int, new_id gd.Int) { //gd:TileSetAtlasSource.set_alternative_tile_id
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	callframe.Arg(frame, new_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_set_alternative_tile_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the alternative ID a following call to [method create_alternative_tile] would return.
*/
//go:nosplit
func (self class) GetNextAlternativeTileId(atlas_coords gd.Vector2i) gd.Int { //gd:TileSetAtlasSource.get_next_alternative_tile_id
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_next_alternative_tile_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [TileData] object for the given atlas coordinates and alternative ID.
*/
//go:nosplit
func (self class) GetTileData(atlas_coords gd.Vector2i, alternative_tile gd.Int) [1]gdclass.TileData { //gd:TileSetAtlasSource.get_tile_data
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TileData{gd.PointerMustAssertInstanceID[gdclass.TileData](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the atlas grid size, which depends on how many tiles can fit in the texture. It thus depends on the [member texture]'s size, the atlas [member margins], and the tiles' [member texture_region_size].
*/
//go:nosplit
func (self class) GetAtlasGridSize() gd.Vector2i { //gd:TileSetAtlasSource.get_atlas_grid_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_atlas_grid_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a tile's texture region in the atlas texture. For animated tiles, a [param frame] argument might be provided for the different frames of the animation.
*/
//go:nosplit
func (self class) GetTileTextureRegion(atlas_coords gd.Vector2i, frame_ gd.Int) gd.Rect2i { //gd:TileSetAtlasSource.get_tile_texture_region
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_tile_texture_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [member use_texture_padding] is [code]false[/code], returns [member texture]. Otherwise, returns and internal [ImageTexture] created that includes the padding.
*/
//go:nosplit
func (self class) GetRuntimeTexture() [1]gdclass.Texture2D { //gd:TileSetAtlasSource.get_runtime_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_runtime_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the region of the tile at coordinates [param atlas_coords] for the given [param frame] inside the texture returned by [method get_runtime_texture].
[b]Note:[/b] If [member use_texture_padding] is [code]false[/code], returns the same as [method get_tile_texture_region].
*/
//go:nosplit
func (self class) GetRuntimeTileTextureRegion(atlas_coords gd.Vector2i, frame_ gd.Int) gd.Rect2i { //gd:TileSetAtlasSource.get_runtime_tile_texture_region
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetAtlasSource.Bind_get_runtime_tile_texture_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTileSetAtlasSource() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTileSetAtlasSource() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTileSetSource() TileSetSource.Advanced {
	return *((*TileSetSource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTileSetSource() TileSetSource.Instance {
	return *((*TileSetSource.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(TileSetSource.Advanced(self.AsTileSetSource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(TileSetSource.Instance(self.AsTileSetSource()), name)
	}
}
func init() {
	gdclass.Register("TileSetAtlasSource", func(ptr gd.Object) any {
		return [1]gdclass.TileSetAtlasSource{*(*gdclass.TileSetAtlasSource)(unsafe.Pointer(&ptr))}
	})
}

type TileAnimationMode = gdclass.TileSetAtlasSourceTileAnimationMode //gd:TileSetAtlasSource.TileAnimationMode

const (
	/*Tile animations start at same time, looking identical.*/
	TileAnimationModeDefault TileAnimationMode = 0
	/*Tile animations start at random times, looking varied.*/
	TileAnimationModeRandomStartTimes TileAnimationMode = 1
	/*Represents the size of the [enum TileAnimationMode] enum.*/
	TileAnimationModeMax TileAnimationMode = 2
)
