package ArrayMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Mesh"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Array"
import "grow.graphics/gd/variant/Dictionary"
import "grow.graphics/gd/variant/Transform3D"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/AABB"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The [ArrayMesh] is used to construct a [Mesh] by specifying the attributes as arrays.
The most basic example is the creation of a single triangle:
[codeblocks]
[gdscript]
var vertices = PackedVector3Array()
vertices.push_back(Vector3(0, 1, 0))
vertices.push_back(Vector3(1, 0, 0))
vertices.push_back(Vector3(0, 0, 1))

# Initialize the ArrayMesh.
var arr_mesh = ArrayMesh.new()
var arrays = []
arrays.resize(Mesh.ARRAY_MAX)
arrays[Mesh.ARRAY_VERTEX] = vertices

# Create the Mesh.
arr_mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, arrays)
var m = MeshInstance3D.new()
m.mesh = arr_mesh
[/gdscript]
[csharp]
var vertices = new Vector3[]

	{
	    new Vector3(0, 1, 0),
	    new Vector3(1, 0, 0),
	    new Vector3(0, 0, 1),
	};

// Initialize the ArrayMesh.
var arrMesh = new ArrayMesh();
var arrays = new Godot.Collections.Array();
arrays.Resize((int)Mesh.ArrayType.Max);
arrays[(int)Mesh.ArrayType.Vertex] = vertices;

// Create the Mesh.
arrMesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, arrays);
var m = new MeshInstance3D();
m.Mesh = arrMesh;
[/csharp]
[/codeblocks]
The [MeshInstance3D] is ready to be added to the [SceneTree] to be shown.
See also [ImmediateMesh], [MeshDataTool] and [SurfaceTool] for procedural geometry generation.
[b]Note:[/b] Godot uses clockwise [url=https://learnopengl.com/Advanced-OpenGL/Face-culling]winding order[/url] for front faces of triangle primitive modes.
*/
type Instance [1]classdb.ArrayMesh

/*
Adds name for a blend shape that will be added with [method add_surface_from_arrays]. Must be called before surface is added.
*/
func (self Instance) AddBlendShape(name string) {
	class(self).AddBlendShape(gd.NewStringName(name))
}

/*
Returns the number of blend shapes that the [ArrayMesh] holds.
*/
func (self Instance) GetBlendShapeCount() int {
	return int(int(class(self).GetBlendShapeCount()))
}

/*
Returns the name of the blend shape at this index.
*/
func (self Instance) GetBlendShapeName(index int) string {
	return string(class(self).GetBlendShapeName(gd.Int(index)).String())
}

/*
Sets the name of the blend shape at this index.
*/
func (self Instance) SetBlendShapeName(index int, name string) {
	class(self).SetBlendShapeName(gd.Int(index), gd.NewStringName(name))
}

/*
Removes all blend shapes from this [ArrayMesh].
*/
func (self Instance) ClearBlendShapes() {
	class(self).ClearBlendShapes()
}

/*
Creates a new surface. [method Mesh.get_surface_count] will become the [code]surf_idx[/code] for this new surface.
Surfaces are created to be rendered using a [param primitive], which may be any of the values defined in [enum Mesh.PrimitiveType].
The [param arrays] argument is an array of arrays. Each of the [constant Mesh.ARRAY_MAX] elements contains an array with some of the mesh data for this surface as described by the corresponding member of [enum Mesh.ArrayType] or [code]null[/code] if it is not used by the surface. For example, [code]arrays[0][/code] is the array of vertices. That first vertex sub-array is always required; the others are optional. Adding an index array puts this surface into "index mode" where the vertex and other arrays become the sources of data and the index array defines the vertex order. All sub-arrays must have the same length as the vertex array (or be an exact multiple of the vertex array's length, when multiple elements of a sub-array correspond to a single vertex) or be empty, except for [constant Mesh.ARRAY_INDEX] if it is used.
The [param blend_shapes] argument is an array of vertex data for each blend shape. Each element is an array of the same structure as [param arrays], but [constant Mesh.ARRAY_VERTEX], [constant Mesh.ARRAY_NORMAL], and [constant Mesh.ARRAY_TANGENT] are set if and only if they are set in [param arrays] and all other entries are [code]null[/code].
The [param lods] argument is a dictionary with [float] keys and [PackedInt32Array] values. Each entry in the dictionary represents an LOD level of the surface, where the value is the [constant Mesh.ARRAY_INDEX] array to use for the LOD level and the key is roughly proportional to the distance at which the LOD stats being used. I.e., increasing the key of an LOD also increases the distance that the objects has to be from the camera before the LOD is used.
The [param flags] argument is the bitwise or of, as required: One value of [enum Mesh.ArrayCustomFormat] left shifted by [code]ARRAY_FORMAT_CUSTOMn_SHIFT[/code] for each custom channel in use, [constant Mesh.ARRAY_FLAG_USE_DYNAMIC_UPDATE], [constant Mesh.ARRAY_FLAG_USE_8_BONE_WEIGHTS], or [constant Mesh.ARRAY_FLAG_USES_EMPTY_VERTEX_ARRAY].
[b]Note:[/b] When using indices, it is recommended to only use points, lines, or triangles.
*/
func (self Instance) AddSurfaceFromArrays(primitive classdb.MeshPrimitiveType, arrays Array.Any) {
	class(self).AddSurfaceFromArrays(primitive, arrays, [1]gd.Array{}[0], [1]Dictionary.Any{}[0], 0)
}

/*
Removes all surfaces from this [ArrayMesh].
*/
func (self Instance) ClearSurfaces() {
	class(self).ClearSurfaces()
}
func (self Instance) SurfaceUpdateVertexRegion(surf_idx int, offset int, data []byte) {
	class(self).SurfaceUpdateVertexRegion(gd.Int(surf_idx), gd.Int(offset), gd.NewPackedByteSlice(data))
}
func (self Instance) SurfaceUpdateAttributeRegion(surf_idx int, offset int, data []byte) {
	class(self).SurfaceUpdateAttributeRegion(gd.Int(surf_idx), gd.Int(offset), gd.NewPackedByteSlice(data))
}
func (self Instance) SurfaceUpdateSkinRegion(surf_idx int, offset int, data []byte) {
	class(self).SurfaceUpdateSkinRegion(gd.Int(surf_idx), gd.Int(offset), gd.NewPackedByteSlice(data))
}

/*
Returns the length in vertices of the vertex array in the requested surface (see [method add_surface_from_arrays]).
*/
func (self Instance) SurfaceGetArrayLen(surf_idx int) int {
	return int(int(class(self).SurfaceGetArrayLen(gd.Int(surf_idx))))
}

/*
Returns the length in indices of the index array in the requested surface (see [method add_surface_from_arrays]).
*/
func (self Instance) SurfaceGetArrayIndexLen(surf_idx int) int {
	return int(int(class(self).SurfaceGetArrayIndexLen(gd.Int(surf_idx))))
}

/*
Returns the format mask of the requested surface (see [method add_surface_from_arrays]).
*/
func (self Instance) SurfaceGetFormat(surf_idx int) classdb.MeshArrayFormat {
	return classdb.MeshArrayFormat(class(self).SurfaceGetFormat(gd.Int(surf_idx)))
}

/*
Returns the primitive type of the requested surface (see [method add_surface_from_arrays]).
*/
func (self Instance) SurfaceGetPrimitiveType(surf_idx int) classdb.MeshPrimitiveType {
	return classdb.MeshPrimitiveType(class(self).SurfaceGetPrimitiveType(gd.Int(surf_idx)))
}

/*
Returns the index of the first surface with this name held within this [ArrayMesh]. If none are found, -1 is returned.
*/
func (self Instance) SurfaceFindByName(name string) int {
	return int(int(class(self).SurfaceFindByName(gd.NewString(name))))
}

/*
Sets a name for a given surface.
*/
func (self Instance) SurfaceSetName(surf_idx int, name string) {
	class(self).SurfaceSetName(gd.Int(surf_idx), gd.NewString(name))
}

/*
Gets the name assigned to this surface.
*/
func (self Instance) SurfaceGetName(surf_idx int) string {
	return string(class(self).SurfaceGetName(gd.Int(surf_idx)).String())
}

/*
Regenerates tangents for each of the [ArrayMesh]'s surfaces.
*/
func (self Instance) RegenNormalMaps() {
	class(self).RegenNormalMaps()
}

/*
Performs a UV unwrap on the [ArrayMesh] to prepare the mesh for lightmapping.
*/
func (self Instance) LightmapUnwrap(transform Transform3D.BasisOrigin, texel_size Float.X) error {
	return error(class(self).LightmapUnwrap(gd.Transform3D(transform), gd.Float(texel_size)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ArrayMesh

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ArrayMesh"))
	return Instance{classdb.ArrayMesh(object)}
}

func (self Instance) BlendShapeMode() classdb.MeshBlendShapeMode {
	return classdb.MeshBlendShapeMode(class(self).GetBlendShapeMode())
}

func (self Instance) SetBlendShapeMode(value classdb.MeshBlendShapeMode) {
	class(self).SetBlendShapeMode(value)
}

func (self Instance) CustomAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetCustomAabb())
}

func (self Instance) SetCustomAabb(value AABB.PositionSize) {
	class(self).SetCustomAabb(gd.AABB(value))
}

func (self Instance) ShadowMesh() objects.ArrayMesh {
	return objects.ArrayMesh(class(self).GetShadowMesh())
}

func (self Instance) SetShadowMesh(value objects.ArrayMesh) {
	class(self).SetShadowMesh(value)
}

/*
Adds name for a blend shape that will be added with [method add_surface_from_arrays]. Must be called before surface is added.
*/
//go:nosplit
func (self class) AddBlendShape(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_add_blend_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of blend shapes that the [ArrayMesh] holds.
*/
//go:nosplit
func (self class) GetBlendShapeCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_get_blend_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the blend shape at this index.
*/
//go:nosplit
func (self class) GetBlendShapeName(index gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_get_blend_shape_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the name of the blend shape at this index.
*/
//go:nosplit
func (self class) SetBlendShapeName(index gd.Int, name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_set_blend_shape_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all blend shapes from this [ArrayMesh].
*/
//go:nosplit
func (self class) ClearBlendShapes() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_clear_blend_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetBlendShapeMode(mode classdb.MeshBlendShapeMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_set_blend_shape_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendShapeMode() classdb.MeshBlendShapeMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MeshBlendShapeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_get_blend_shape_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new surface. [method Mesh.get_surface_count] will become the [code]surf_idx[/code] for this new surface.
Surfaces are created to be rendered using a [param primitive], which may be any of the values defined in [enum Mesh.PrimitiveType].
The [param arrays] argument is an array of arrays. Each of the [constant Mesh.ARRAY_MAX] elements contains an array with some of the mesh data for this surface as described by the corresponding member of [enum Mesh.ArrayType] or [code]null[/code] if it is not used by the surface. For example, [code]arrays[0][/code] is the array of vertices. That first vertex sub-array is always required; the others are optional. Adding an index array puts this surface into "index mode" where the vertex and other arrays become the sources of data and the index array defines the vertex order. All sub-arrays must have the same length as the vertex array (or be an exact multiple of the vertex array's length, when multiple elements of a sub-array correspond to a single vertex) or be empty, except for [constant Mesh.ARRAY_INDEX] if it is used.
The [param blend_shapes] argument is an array of vertex data for each blend shape. Each element is an array of the same structure as [param arrays], but [constant Mesh.ARRAY_VERTEX], [constant Mesh.ARRAY_NORMAL], and [constant Mesh.ARRAY_TANGENT] are set if and only if they are set in [param arrays] and all other entries are [code]null[/code].
The [param lods] argument is a dictionary with [float] keys and [PackedInt32Array] values. Each entry in the dictionary represents an LOD level of the surface, where the value is the [constant Mesh.ARRAY_INDEX] array to use for the LOD level and the key is roughly proportional to the distance at which the LOD stats being used. I.e., increasing the key of an LOD also increases the distance that the objects has to be from the camera before the LOD is used.
The [param flags] argument is the bitwise or of, as required: One value of [enum Mesh.ArrayCustomFormat] left shifted by [code]ARRAY_FORMAT_CUSTOMn_SHIFT[/code] for each custom channel in use, [constant Mesh.ARRAY_FLAG_USE_DYNAMIC_UPDATE], [constant Mesh.ARRAY_FLAG_USE_8_BONE_WEIGHTS], or [constant Mesh.ARRAY_FLAG_USES_EMPTY_VERTEX_ARRAY].
[b]Note:[/b] When using indices, it is recommended to only use points, lines, or triangles.
*/
//go:nosplit
func (self class) AddSurfaceFromArrays(primitive classdb.MeshPrimitiveType, arrays gd.Array, blend_shapes gd.Array, lods gd.Dictionary, flags classdb.MeshArrayFormat) {
	var frame = callframe.New()
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, pointers.Get(arrays))
	callframe.Arg(frame, pointers.Get(blend_shapes))
	callframe.Arg(frame, pointers.Get(lods))
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_add_surface_from_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all surfaces from this [ArrayMesh].
*/
//go:nosplit
func (self class) ClearSurfaces() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_clear_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SurfaceUpdateVertexRegion(surf_idx gd.Int, offset gd.Int, data gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, pointers.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_update_vertex_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SurfaceUpdateAttributeRegion(surf_idx gd.Int, offset gd.Int, data gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, pointers.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_update_attribute_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SurfaceUpdateSkinRegion(surf_idx gd.Int, offset gd.Int, data gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, pointers.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_update_skin_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the length in vertices of the vertex array in the requested surface (see [method add_surface_from_arrays]).
*/
//go:nosplit
func (self class) SurfaceGetArrayLen(surf_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_get_array_len, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the length in indices of the index array in the requested surface (see [method add_surface_from_arrays]).
*/
//go:nosplit
func (self class) SurfaceGetArrayIndexLen(surf_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_get_array_index_len, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the format mask of the requested surface (see [method add_surface_from_arrays]).
*/
//go:nosplit
func (self class) SurfaceGetFormat(surf_idx gd.Int) classdb.MeshArrayFormat {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[classdb.MeshArrayFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the primitive type of the requested surface (see [method add_surface_from_arrays]).
*/
//go:nosplit
func (self class) SurfaceGetPrimitiveType(surf_idx gd.Int) classdb.MeshPrimitiveType {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[classdb.MeshPrimitiveType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_get_primitive_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the first surface with this name held within this [ArrayMesh]. If none are found, -1 is returned.
*/
//go:nosplit
func (self class) SurfaceFindByName(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_find_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a name for a given surface.
*/
//go:nosplit
func (self class) SurfaceSetName(surf_idx gd.Int, name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_set_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets the name assigned to this surface.
*/
//go:nosplit
func (self class) SurfaceGetName(surf_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Regenerates tangents for each of the [ArrayMesh]'s surfaces.
*/
//go:nosplit
func (self class) RegenNormalMaps() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_regen_normal_maps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Performs a UV unwrap on the [ArrayMesh] to prepare the mesh for lightmapping.
*/
//go:nosplit
func (self class) LightmapUnwrap(transform gd.Transform3D, texel_size gd.Float) error {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	callframe.Arg(frame, texel_size)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_lightmap_unwrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowMesh(mesh objects.ArrayMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_set_shadow_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowMesh() objects.ArrayMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_get_shadow_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ArrayMesh{classdb.ArrayMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsArrayMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsArrayMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.Advanced    { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsMesh(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMesh(), name)
	}
}
func init() { classdb.Register("ArrayMesh", func(ptr gd.Object) any { return classdb.ArrayMesh(ptr) }) }

type Error int

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
