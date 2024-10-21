package TileData

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[TileData] object represents a single tile in a [TileSet]. It is usually edited using the tileset editor, but it can be modified at runtime using [method TileMap._tile_data_runtime_update].

*/
type Simple [1]classdb.TileData
func (self Simple) SetFlipH(flip_h bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlipH(flip_h)
}
func (self Simple) GetFlipH() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlipH())
}
func (self Simple) SetFlipV(flip_v bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlipV(flip_v)
}
func (self Simple) GetFlipV() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlipV())
}
func (self Simple) SetTranspose(transpose bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTranspose(transpose)
}
func (self Simple) GetTranspose() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetTranspose())
}
func (self Simple) SetMaterial(material [1]classdb.Material) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaterial(material)
}
func (self Simple) GetMaterial() [1]classdb.Material {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Material(Expert(self).GetMaterial(gc))
}
func (self Simple) SetTextureOrigin(texture_origin gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureOrigin(texture_origin)
}
func (self Simple) GetTextureOrigin() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetTextureOrigin())
}
func (self Simple) SetModulate(modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetModulate(modulate)
}
func (self Simple) GetModulate() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetModulate())
}
func (self Simple) SetZIndex(z_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetZIndex(gd.Int(z_index))
}
func (self Simple) GetZIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetZIndex()))
}
func (self Simple) SetYSortOrigin(y_sort_origin int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetYSortOrigin(gd.Int(y_sort_origin))
}
func (self Simple) GetYSortOrigin() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetYSortOrigin()))
}
func (self Simple) SetOccluder(layer_id int, occluder_polygon [1]classdb.OccluderPolygon2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOccluder(gd.Int(layer_id), occluder_polygon)
}
func (self Simple) GetOccluder(layer_id int, flip_h bool, flip_v bool, transpose bool) [1]classdb.OccluderPolygon2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OccluderPolygon2D(Expert(self).GetOccluder(gc, gd.Int(layer_id), flip_h, flip_v, transpose))
}
func (self Simple) SetConstantLinearVelocity(layer_id int, velocity gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstantLinearVelocity(gd.Int(layer_id), velocity)
}
func (self Simple) GetConstantLinearVelocity(layer_id int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetConstantLinearVelocity(gd.Int(layer_id)))
}
func (self Simple) SetConstantAngularVelocity(layer_id int, velocity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstantAngularVelocity(gd.Int(layer_id), gd.Float(velocity))
}
func (self Simple) GetConstantAngularVelocity(layer_id int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetConstantAngularVelocity(gd.Int(layer_id))))
}
func (self Simple) SetCollisionPolygonsCount(layer_id int, polygons_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionPolygonsCount(gd.Int(layer_id), gd.Int(polygons_count))
}
func (self Simple) GetCollisionPolygonsCount(layer_id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCollisionPolygonsCount(gd.Int(layer_id))))
}
func (self Simple) AddCollisionPolygon(layer_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCollisionPolygon(gd.Int(layer_id))
}
func (self Simple) RemoveCollisionPolygon(layer_id int, polygon_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveCollisionPolygon(gd.Int(layer_id), gd.Int(polygon_index))
}
func (self Simple) SetCollisionPolygonPoints(layer_id int, polygon_index int, polygon gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionPolygonPoints(gd.Int(layer_id), gd.Int(polygon_index), polygon)
}
func (self Simple) GetCollisionPolygonPoints(layer_id int, polygon_index int) gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetCollisionPolygonPoints(gc, gd.Int(layer_id), gd.Int(polygon_index)))
}
func (self Simple) SetCollisionPolygonOneWay(layer_id int, polygon_index int, one_way bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionPolygonOneWay(gd.Int(layer_id), gd.Int(polygon_index), one_way)
}
func (self Simple) IsCollisionPolygonOneWay(layer_id int, polygon_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCollisionPolygonOneWay(gd.Int(layer_id), gd.Int(polygon_index)))
}
func (self Simple) SetCollisionPolygonOneWayMargin(layer_id int, polygon_index int, one_way_margin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionPolygonOneWayMargin(gd.Int(layer_id), gd.Int(polygon_index), gd.Float(one_way_margin))
}
func (self Simple) GetCollisionPolygonOneWayMargin(layer_id int, polygon_index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCollisionPolygonOneWayMargin(gd.Int(layer_id), gd.Int(polygon_index))))
}
func (self Simple) SetTerrainSet(terrain_set int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTerrainSet(gd.Int(terrain_set))
}
func (self Simple) GetTerrainSet() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTerrainSet()))
}
func (self Simple) SetTerrain(terrain int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTerrain(gd.Int(terrain))
}
func (self Simple) GetTerrain() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTerrain()))
}
func (self Simple) SetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor, terrain int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTerrainPeeringBit(peering_bit, gd.Int(terrain))
}
func (self Simple) GetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTerrainPeeringBit(peering_bit)))
}
func (self Simple) IsValidTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsValidTerrainPeeringBit(peering_bit))
}
func (self Simple) SetNavigationPolygon(layer_id int, navigation_polygon [1]classdb.NavigationPolygon) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationPolygon(gd.Int(layer_id), navigation_polygon)
}
func (self Simple) GetNavigationPolygon(layer_id int, flip_h bool, flip_v bool, transpose bool) [1]classdb.NavigationPolygon {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.NavigationPolygon(Expert(self).GetNavigationPolygon(gc, gd.Int(layer_id), flip_h, flip_v, transpose))
}
func (self Simple) SetProbability(probability float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProbability(gd.Float(probability))
}
func (self Simple) GetProbability() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetProbability()))
}
func (self Simple) SetCustomData(layer_name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomData(gc.String(layer_name), value)
}
func (self Simple) GetCustomData(layer_name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetCustomData(gc, gc.String(layer_name)))
}
func (self Simple) SetCustomDataByLayerId(layer_id int, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomDataByLayerId(gd.Int(layer_id), value)
}
func (self Simple) GetCustomDataByLayerId(layer_id int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetCustomDataByLayerId(gc, gd.Int(layer_id)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TileData
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetFlipH(flip_h bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip_h)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlipH() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlipV(flip_v bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip_v)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlipV() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTranspose(transpose bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, transpose)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_transpose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTranspose() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_transpose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaterial(material object.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) object.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureOrigin(texture_origin gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture_origin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_texture_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureOrigin() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_texture_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModulate(modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZIndex(z_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, z_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetYSortOrigin(y_sort_origin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, y_sort_origin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetYSortOrigin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the occluder for the TileSet occlusion layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetOccluder(layer_id gd.Int, occluder_polygon object.OccluderPolygon2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, mmm.Get(occluder_polygon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_occluder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the occluder polygon of the tile for the TileSet occlusion layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
//go:nosplit
func (self class) GetOccluder(ctx gd.Lifetime, layer_id gd.Int, flip_h bool, flip_v bool, transpose bool) object.OccluderPolygon2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, flip_h)
	callframe.Arg(frame, flip_v)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_occluder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OccluderPolygon2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the constant linear velocity. This does not move the tile. This linear velocity is applied to objects colliding with this tile. This is useful to create conveyor belts.
*/
//go:nosplit
func (self class) SetConstantLinearVelocity(layer_id gd.Int, velocity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the constant linear velocity applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) GetConstantLinearVelocity(layer_id gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the constant angular velocity. This does not rotate the tile. This angular velocity is applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) SetConstantAngularVelocity(layer_id gd.Int, velocity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the constant angular velocity applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) GetConstantAngularVelocity(layer_id gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the polygons count for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonsCount(layer_id gd.Int, polygons_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygons_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_collision_polygons_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns how many polygons the tile has for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonsCount(layer_id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_collision_polygons_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a collision polygon to the tile on the given TileSet physics layer.
*/
//go:nosplit
func (self class) AddCollisionPolygon(layer_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_add_collision_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) RemoveCollisionPolygon(layer_id gd.Int, polygon_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_remove_collision_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonPoints(layer_id gd.Int, polygon_index gd.Int, polygon gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, mmm.Get(polygon))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_collision_polygon_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonPoints(ctx gd.Lifetime, layer_id gd.Int, polygon_index gd.Int) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_collision_polygon_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonOneWay(layer_id gd.Int, polygon_index gd.Int, one_way bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, one_way)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_collision_polygon_one_way, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether one-way collisions are enabled for the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) IsCollisionPolygonOneWay(layer_id gd.Int, polygon_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_is_collision_polygon_one_way, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonOneWayMargin(layer_id gd.Int, polygon_index gd.Int, one_way_margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, one_way_margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_collision_polygon_one_way_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the one-way margin (for one-way platforms) of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonOneWayMargin(layer_id gd.Int, polygon_index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_collision_polygon_one_way_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTerrainSet(terrain_set gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_terrain_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTerrainSet() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_terrain_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTerrain(terrain gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, terrain)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_terrain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTerrain() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_terrain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
//go:nosplit
func (self class) SetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor, terrain gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	callframe.Arg(frame, terrain)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
//go:nosplit
func (self class) GetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether the given [param peering_bit] direction is valid for this tile.
*/
//go:nosplit
func (self class) IsValidTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_is_valid_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the navigation polygon for the TileSet navigation layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetNavigationPolygon(layer_id gd.Int, navigation_polygon object.NavigationPolygon)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, mmm.Get(navigation_polygon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the navigation polygon of the tile for the TileSet navigation layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
//go:nosplit
func (self class) GetNavigationPolygon(ctx gd.Lifetime, layer_id gd.Int, flip_h bool, flip_v bool, transpose bool) object.NavigationPolygon {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, flip_h)
	callframe.Arg(frame, flip_v)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.NavigationPolygon
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProbability(probability gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probability)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_probability, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProbability() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_probability, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tile's custom data value for the TileSet custom data layer with name [param layer_name].
*/
//go:nosplit
func (self class) SetCustomData(layer_name gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(layer_name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom data value for custom data layer named [param layer_name].
*/
//go:nosplit
func (self class) GetCustomData(ctx gd.Lifetime, layer_name gd.String) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(layer_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the tile's custom data value for the TileSet custom data layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCustomDataByLayerId(layer_id gd.Int, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_set_custom_data_by_layer_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom data value for custom data layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCustomDataByLayerId(ctx gd.Lifetime, layer_id gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileData.Bind_get_custom_data_by_layer_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTileData() Expert { return self[0].AsTileData() }


//go:nosplit
func (self Simple) AsTileData() Simple { return self[0].AsTileData() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TileData", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
