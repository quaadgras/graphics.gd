package VoxelGIData

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
[VoxelGIData] contains baked voxel global illumination for use in a [VoxelGI] node. [VoxelGIData] also offers several properties to adjust the final appearance of the global illumination. These properties can be adjusted at run-time without having to bake the [VoxelGI] node again.
[b]Note:[/b] To prevent text-based scene files ([code].tscn[/code]) from growing too much and becoming slow to load and save, always save [VoxelGIData] to an external binary resource file ([code].res[/code]) instead of embedding it within the scene. This can be done by clicking the dropdown arrow next to the [VoxelGIData] resource, choosing [b]Edit[/b], clicking the floppy disk icon at the top of the Inspector then choosing [b]Save As...[/b].

*/
type Go [1]classdb.VoxelGIData
func (self Go) Allocate(to_cell_xform gd.Transform3D, aabb gd.AABB, octree_size gd.Vector3, octree_cells []byte, data_cells []byte, distance_field []byte, level_counts []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Allocate(to_cell_xform, aabb, octree_size, gc.PackedByteSlice(octree_cells), gc.PackedByteSlice(data_cells), gc.PackedByteSlice(distance_field), gc.PackedInt32Slice(level_counts))
}

/*
Returns the bounds of the baked voxel data as an [AABB], which should match [member VoxelGI.size] after being baked (which only contains the size as a [Vector3]).
[b]Note:[/b] If the size was modified without baking the VoxelGI data, then the value of [method get_bounds] and [member VoxelGI.size] will not match.
*/
func (self Go) GetBounds() gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(class(self).GetBounds())
}
func (self Go) GetOctreeSize() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetOctreeSize())
}
func (self Go) GetToCellXform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(class(self).GetToCellXform())
}
func (self Go) GetOctreeCells() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).GetOctreeCells(gc).Bytes())
}
func (self Go) GetDataCells() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).GetDataCells(gc).Bytes())
}
func (self Go) GetLevelCounts() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetLevelCounts(gc).AsSlice())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VoxelGIData
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VoxelGIData"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) DynamicRange() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDynamicRange()))
}

func (self Go) SetDynamicRange(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDynamicRange(gd.Float(value))
}

func (self Go) Energy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEnergy()))
}

func (self Go) SetEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnergy(gd.Float(value))
}

func (self Go) Bias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBias()))
}

func (self Go) SetBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBias(gd.Float(value))
}

func (self Go) NormalBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetNormalBias()))
}

func (self Go) SetNormalBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNormalBias(gd.Float(value))
}

func (self Go) Propagation() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPropagation()))
}

func (self Go) SetPropagation(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPropagation(gd.Float(value))
}

func (self Go) UseTwoBounces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingTwoBounces())
}

func (self Go) SetUseTwoBounces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseTwoBounces(value)
}

func (self Go) Interior() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsInterior())
}

func (self Go) SetInterior(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInterior(value)
}

//go:nosplit
func (self class) Allocate(to_cell_xform gd.Transform3D, aabb gd.AABB, octree_size gd.Vector3, octree_cells gd.PackedByteArray, data_cells gd.PackedByteArray, distance_field gd.PackedByteArray, level_counts gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, to_cell_xform)
	callframe.Arg(frame, aabb)
	callframe.Arg(frame, octree_size)
	callframe.Arg(frame, mmm.Get(octree_cells))
	callframe.Arg(frame, mmm.Get(data_cells))
	callframe.Arg(frame, mmm.Get(distance_field))
	callframe.Arg(frame, mmm.Get(level_counts))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_allocate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the bounds of the baked voxel data as an [AABB], which should match [member VoxelGI.size] after being baked (which only contains the size as a [Vector3]).
[b]Note:[/b] If the size was modified without baking the VoxelGI data, then the value of [method get_bounds] and [member VoxelGI.size] will not match.
*/
//go:nosplit
func (self class) GetBounds() gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetOctreeSize() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_octree_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetToCellXform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_to_cell_xform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetOctreeCells(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_octree_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetDataCells(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_data_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetLevelCounts(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_level_counts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDynamicRange(dynamic_range gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, dynamic_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_set_dynamic_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDynamicRange() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_dynamic_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnergy(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_set_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_set_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNormalBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_set_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNormalBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPropagation(propagation gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, propagation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_set_propagation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPropagation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_get_propagation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInterior(interior bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interior)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_set_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsInterior() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_is_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseTwoBounces(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_set_use_two_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingTwoBounces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VoxelGIData.Bind_is_using_two_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVoxelGIData() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVoxelGIData() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("VoxelGIData", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
