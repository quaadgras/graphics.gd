package NavigationMeshSourceGeometryData3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Container for parsed source geometry data used in navigation mesh baking.
*/
type Instance [1]classdb.NavigationMeshSourceGeometryData3D

/*
Appends arrays of [param vertices] and [param indices] at the end of the existing arrays. Adds the existing index as an offset to the appended indices.
*/
func (self Instance) AppendArrays(vertices []float32, indices []int32) {
	class(self).AppendArrays(gd.NewPackedFloat32Slice(vertices), gd.NewPackedInt32Slice(indices))
}

/*
Clears the internal data.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
func (self Instance) HasData() bool {
	return bool(class(self).HasData())
}

/*
Adds the geometry data of a [Mesh] resource to the navigation mesh baking data. The mesh must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
func (self Instance) AddMesh(mesh objects.Mesh, xform gd.Transform3D) {
	class(self).AddMesh(mesh, xform)
}

/*
Adds an [Array] the size of [constant Mesh.ARRAY_MAX] and with vertices at index [constant Mesh.ARRAY_VERTEX] and indices at index [constant Mesh.ARRAY_INDEX] to the navigation mesh baking data. The array must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
func (self Instance) AddMeshArray(mesh_array gd.Array, xform gd.Transform3D) {
	class(self).AddMeshArray(mesh_array, xform)
}

/*
Adds an array of vertex positions to the geometry data for navigation mesh baking to form triangulated faces. For each face the array must have three vertex positions in clockwise winding order. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
func (self Instance) AddFaces(faces []gd.Vector3, xform gd.Transform3D) {
	class(self).AddFaces(gd.NewPackedVector3Slice(faces), xform)
}

/*
Adds the geometry data of another [NavigationMeshSourceGeometryData3D] to the navigation mesh baking data.
*/
func (self Instance) Merge(other_geometry objects.NavigationMeshSourceGeometryData3D) {
	class(self).Merge(other_geometry)
}

/*
Adds a projected obstruction shape to the source geometry. The [param vertices] are considered projected on a xz-axes plane, placed at the global y-axis [param elevation] and extruded by [param height]. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
func (self Instance) AddProjectedObstruction(vertices []gd.Vector3, elevation float64, height float64, carve bool) {
	class(self).AddProjectedObstruction(gd.NewPackedVector3Slice(vertices), gd.Float(elevation), gd.Float(height), carve)
}

/*
Clears all projected obstructions.
*/
func (self Instance) ClearProjectedObstructions() {
	class(self).ClearProjectedObstructions()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.NavigationMeshSourceGeometryData3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationMeshSourceGeometryData3D"))
	return Instance{classdb.NavigationMeshSourceGeometryData3D(object)}
}

func (self Instance) Vertices() []float32 {
	return []float32(class(self).GetVertices().AsSlice())
}

func (self Instance) SetVertices(value []float32) {
	class(self).SetVertices(gd.NewPackedFloat32Slice(value))
}

func (self Instance) Indices() []int32 {
	return []int32(class(self).GetIndices().AsSlice())
}

func (self Instance) SetIndices(value []int32) {
	class(self).SetIndices(gd.NewPackedInt32Slice(value))
}

func (self Instance) ProjectedObstructions() gd.Array {
	return gd.Array(class(self).GetProjectedObstructions())
}

func (self Instance) SetProjectedObstructions(value gd.Array) {
	class(self).SetProjectedObstructions(value)
}

/*
Sets the parsed source geometry data vertices. The vertices need to be matched with appropriated indices.
[b]Warning:[/b] Inappropriate data can crash the baking process of the involved third-party libraries.
*/
//go:nosplit
func (self class) SetVertices(vertices gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the parsed source geometry data vertices array.
*/
//go:nosplit
func (self class) GetVertices() gd.PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_get_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the parsed source geometry data indices. The indices need to be matched with appropriated vertices.
[b]Warning:[/b] Inappropriate data can crash the baking process of the involved third-party libraries.
*/
//go:nosplit
func (self class) SetIndices(indices gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(indices))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_set_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the parsed source geometry data indices array.
*/
//go:nosplit
func (self class) GetIndices() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_get_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Appends arrays of [param vertices] and [param indices] at the end of the existing arrays. Adds the existing index as an offset to the appended indices.
*/
//go:nosplit
func (self class) AppendArrays(vertices gd.PackedFloat32Array, indices gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	callframe.Arg(frame, pointers.Get(indices))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_append_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears the internal data.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
//go:nosplit
func (self class) HasData() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_has_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds the geometry data of a [Mesh] resource to the navigation mesh baking data. The mesh must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddMesh(mesh objects.Mesh, xform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds an [Array] the size of [constant Mesh.ARRAY_MAX] and with vertices at index [constant Mesh.ARRAY_VERTEX] and indices at index [constant Mesh.ARRAY_INDEX] to the navigation mesh baking data. The array must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddMeshArray(mesh_array gd.Array, xform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh_array))
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_mesh_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds an array of vertex positions to the geometry data for navigation mesh baking to form triangulated faces. For each face the array must have three vertex positions in clockwise winding order. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddFaces(faces gd.PackedVector3Array, xform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(faces))
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds the geometry data of another [NavigationMeshSourceGeometryData3D] to the navigation mesh baking data.
*/
//go:nosplit
func (self class) Merge(other_geometry objects.NavigationMeshSourceGeometryData3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(other_geometry[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_merge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a projected obstruction shape to the source geometry. The [param vertices] are considered projected on a xz-axes plane, placed at the global y-axis [param elevation] and extruded by [param height]. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
//go:nosplit
func (self class) AddProjectedObstruction(vertices gd.PackedVector3Array, elevation gd.Float, height gd.Float, carve bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	callframe.Arg(frame, elevation)
	callframe.Arg(frame, height)
	callframe.Arg(frame, carve)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_projected_obstruction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears all projected obstructions.
*/
//go:nosplit
func (self class) ClearProjectedObstructions() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_clear_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetProjectedObstructions(projected_obstructions gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(projected_obstructions))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_set_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetProjectedObstructions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_get_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsNavigationMeshSourceGeometryData3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNavigationMeshSourceGeometryData3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("NavigationMeshSourceGeometryData3D", func(ptr gd.Object) any { return classdb.NavigationMeshSourceGeometryData3D(ptr) })
}
