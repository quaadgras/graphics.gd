package NavigationMeshSourceGeometryData3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Container for parsed source geometry data used in navigation mesh baking.

*/
type Simple [1]classdb.NavigationMeshSourceGeometryData3D
func (self Simple) SetVertices(vertices gd.PackedFloat32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVertices(vertices)
}
func (self Simple) GetVertices() gd.PackedFloat32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat32Array(Expert(self).GetVertices(gc))
}
func (self Simple) SetIndices(indices gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIndices(indices)
}
func (self Simple) GetIndices() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetIndices(gc))
}
func (self Simple) AppendArrays(vertices gd.PackedFloat32Array, indices gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AppendArrays(vertices, indices)
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) HasData() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasData())
}
func (self Simple) AddMesh(mesh [1]classdb.Mesh, xform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddMesh(mesh, xform)
}
func (self Simple) AddMeshArray(mesh_array gd.Array, xform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddMeshArray(mesh_array, xform)
}
func (self Simple) AddFaces(faces gd.PackedVector3Array, xform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddFaces(faces, xform)
}
func (self Simple) Merge(other_geometry [1]classdb.NavigationMeshSourceGeometryData3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Merge(other_geometry)
}
func (self Simple) AddProjectedObstruction(vertices gd.PackedVector3Array, elevation float64, height float64, carve bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddProjectedObstruction(vertices, gd.Float(elevation), gd.Float(height), carve)
}
func (self Simple) ClearProjectedObstructions() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearProjectedObstructions()
}
func (self Simple) SetProjectedObstructions(projected_obstructions gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProjectedObstructions(projected_obstructions)
}
func (self Simple) GetProjectedObstructions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetProjectedObstructions(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationMeshSourceGeometryData3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the parsed source geometry data vertices. The vertices need to be matched with appropriated indices.
[b]Warning:[/b] Inappropriate data can crash the baking process of the involved third-party libraries.
*/
//go:nosplit
func (self class) SetVertices(vertices gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the parsed source geometry data vertices array.
*/
//go:nosplit
func (self class) GetVertices(ctx gd.Lifetime) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_get_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the parsed source geometry data indices. The indices need to be matched with appropriated vertices.
[b]Warning:[/b] Inappropriate data can crash the baking process of the involved third-party libraries.
*/
//go:nosplit
func (self class) SetIndices(indices gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(indices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_set_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the parsed source geometry data indices array.
*/
//go:nosplit
func (self class) GetIndices(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_get_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Appends arrays of [param vertices] and [param indices] at the end of the existing arrays. Adds the existing index as an offset to the appended indices.
*/
//go:nosplit
func (self class) AppendArrays(vertices gd.PackedFloat32Array, indices gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	callframe.Arg(frame, mmm.Get(indices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_append_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the internal data.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
//go:nosplit
func (self class) HasData() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_has_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the geometry data of a [Mesh] resource to the navigation mesh baking data. The mesh must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddMesh(mesh object.Mesh, xform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_add_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an [Array] the size of [constant Mesh.ARRAY_MAX] and with vertices at index [constant Mesh.ARRAY_VERTEX] and indices at index [constant Mesh.ARRAY_INDEX] to the navigation mesh baking data. The array must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddMeshArray(mesh_array gd.Array, xform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh_array))
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_add_mesh_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an array of vertex positions to the geometry data for navigation mesh baking to form triangulated faces. For each face the array must have three vertex positions in clockwise winding order. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddFaces(faces gd.PackedVector3Array, xform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(faces))
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_add_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds the geometry data of another [NavigationMeshSourceGeometryData3D] to the navigation mesh baking data.
*/
//go:nosplit
func (self class) Merge(other_geometry object.NavigationMeshSourceGeometryData3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(other_geometry[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_merge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a projected obstruction shape to the source geometry. The [param vertices] are considered projected on a xz-axes plane, placed at the global y-axis [param elevation] and extruded by [param height]. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
//go:nosplit
func (self class) AddProjectedObstruction(vertices gd.PackedVector3Array, elevation gd.Float, height gd.Float, carve bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	callframe.Arg(frame, elevation)
	callframe.Arg(frame, height)
	callframe.Arg(frame, carve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_add_projected_obstruction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all projected obstructions.
*/
//go:nosplit
func (self class) ClearProjectedObstructions()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_clear_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the projected obstructions with an Array of Dictionaries with the following key value pairs:
[codeblocks]
[gdscript]
"vertices" : PackedFloat32Array
"elevation" : float
"height" : float
"carve" : bool
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) SetProjectedObstructions(projected_obstructions gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(projected_obstructions))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_set_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the projected obstructions as an [Array] of dictionaries. Each [Dictionary] contains the following entries:
- [code]vertices[/code] - A [PackedFloat32Array] that defines the outline points of the projected shape.
- [code]elevation[/code] - A [float] that defines the projected shape placement on the y-axis.
- [code]height[/code] - A [float] that defines how much the projected shape is extruded along the y-axis.
- [code]carve[/code] - A [bool] that defines how the obstacle affects the navigation mesh baking. If [code]true[/code] the projected shape will not be affected by addition offsets, e.g. agent radius.
*/
//go:nosplit
func (self class) GetProjectedObstructions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData3D.Bind_get_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsNavigationMeshSourceGeometryData3D() Expert { return self[0].AsNavigationMeshSourceGeometryData3D() }


//go:nosplit
func (self Simple) AsNavigationMeshSourceGeometryData3D() Simple { return self[0].AsNavigationMeshSourceGeometryData3D() }


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
func init() {classdb.Register("NavigationMeshSourceGeometryData3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
