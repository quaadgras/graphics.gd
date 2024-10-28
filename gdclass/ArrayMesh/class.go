package ArrayMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Mesh"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

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
type Go [1]classdb.ArrayMesh

/*
Adds name for a blend shape that will be added with [method add_surface_from_arrays]. Must be called before surface is added.
*/
func (self Go) AddBlendShape(name string) {
	class(self).AddBlendShape(gd.NewStringName(name))
}

/*
Returns the number of blend shapes that the [ArrayMesh] holds.
*/
func (self Go) GetBlendShapeCount() int {
	return int(int(class(self).GetBlendShapeCount()))
}

/*
Returns the name of the blend shape at this index.
*/
func (self Go) GetBlendShapeName(index int) string {
	return string(class(self).GetBlendShapeName(gd.Int(index)).String())
}

/*
Sets the name of the blend shape at this index.
*/
func (self Go) SetBlendShapeName(index int, name string) {
	class(self).SetBlendShapeName(gd.Int(index), gd.NewStringName(name))
}

/*
Removes all blend shapes from this [ArrayMesh].
*/
func (self Go) ClearBlendShapes() {
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
func (self Go) AddSurfaceFromArrays(primitive classdb.MeshPrimitiveType, arrays gd.Array) {
	class(self).AddSurfaceFromArrays(primitive, arrays, ([1]gd.Array{}[0]), ([1]gd.Dictionary{}[0]), 0)
}

/*
Removes all surfaces from this [ArrayMesh].
*/
func (self Go) ClearSurfaces() {
	class(self).ClearSurfaces()
}
func (self Go) SurfaceUpdateVertexRegion(surf_idx int, offset int, data []byte) {
	class(self).SurfaceUpdateVertexRegion(gd.Int(surf_idx), gd.Int(offset), gd.NewPackedByteSlice(data))
}
func (self Go) SurfaceUpdateAttributeRegion(surf_idx int, offset int, data []byte) {
	class(self).SurfaceUpdateAttributeRegion(gd.Int(surf_idx), gd.Int(offset), gd.NewPackedByteSlice(data))
}
func (self Go) SurfaceUpdateSkinRegion(surf_idx int, offset int, data []byte) {
	class(self).SurfaceUpdateSkinRegion(gd.Int(surf_idx), gd.Int(offset), gd.NewPackedByteSlice(data))
}

/*
Returns the length in vertices of the vertex array in the requested surface (see [method add_surface_from_arrays]).
*/
func (self Go) SurfaceGetArrayLen(surf_idx int) int {
	return int(int(class(self).SurfaceGetArrayLen(gd.Int(surf_idx))))
}

/*
Returns the length in indices of the index array in the requested surface (see [method add_surface_from_arrays]).
*/
func (self Go) SurfaceGetArrayIndexLen(surf_idx int) int {
	return int(int(class(self).SurfaceGetArrayIndexLen(gd.Int(surf_idx))))
}

/*
Returns the format mask of the requested surface (see [method add_surface_from_arrays]).
*/
func (self Go) SurfaceGetFormat(surf_idx int) classdb.MeshArrayFormat {
	return classdb.MeshArrayFormat(class(self).SurfaceGetFormat(gd.Int(surf_idx)))
}

/*
Returns the primitive type of the requested surface (see [method add_surface_from_arrays]).
*/
func (self Go) SurfaceGetPrimitiveType(surf_idx int) classdb.MeshPrimitiveType {
	return classdb.MeshPrimitiveType(class(self).SurfaceGetPrimitiveType(gd.Int(surf_idx)))
}

/*
Returns the index of the first surface with this name held within this [ArrayMesh]. If none are found, -1 is returned.
*/
func (self Go) SurfaceFindByName(name string) int {
	return int(int(class(self).SurfaceFindByName(gd.NewString(name))))
}

/*
Sets a name for a given surface.
*/
func (self Go) SurfaceSetName(surf_idx int, name string) {
	class(self).SurfaceSetName(gd.Int(surf_idx), gd.NewString(name))
}

/*
Gets the name assigned to this surface.
*/
func (self Go) SurfaceGetName(surf_idx int) string {
	return string(class(self).SurfaceGetName(gd.Int(surf_idx)).String())
}

/*
Regenerates tangents for each of the [ArrayMesh]'s surfaces.
*/
func (self Go) RegenNormalMaps() {
	class(self).RegenNormalMaps()
}

/*
Performs a UV unwrap on the [ArrayMesh] to prepare the mesh for lightmapping.
*/
func (self Go) LightmapUnwrap(transform gd.Transform3D, texel_size float64) gd.Error {
	return gd.Error(class(self).LightmapUnwrap(transform, gd.Float(texel_size)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ArrayMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ArrayMesh"))
	return Go{classdb.ArrayMesh(object)}
}

func (self Go) BlendShapeMode() classdb.MeshBlendShapeMode {
		return classdb.MeshBlendShapeMode(class(self).GetBlendShapeMode())
}

func (self Go) SetBlendShapeMode(value classdb.MeshBlendShapeMode) {
	class(self).SetBlendShapeMode(value)
}

func (self Go) CustomAabb() gd.AABB {
		return gd.AABB(class(self).GetCustomAabb())
}

func (self Go) SetCustomAabb(value gd.AABB) {
	class(self).SetCustomAabb(value)
}

func (self Go) ShadowMesh() gdclass.ArrayMesh {
		return gdclass.ArrayMesh(class(self).GetShadowMesh())
}

func (self Go) SetShadowMesh(value gdclass.ArrayMesh) {
	class(self).SetShadowMesh(value)
}

/*
Adds name for a blend shape that will be added with [method add_surface_from_arrays]. Must be called before surface is added.
*/
//go:nosplit
func (self class) AddBlendShape(name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
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
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the name of the blend shape at this index.
*/
//go:nosplit
func (self class) SetBlendShapeName(index gd.Int, name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_set_blend_shape_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all blend shapes from this [ArrayMesh].
*/
//go:nosplit
func (self class) ClearBlendShapes()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_clear_blend_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBlendShapeMode(mode classdb.MeshBlendShapeMode)  {
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
func (self class) AddSurfaceFromArrays(primitive classdb.MeshPrimitiveType, arrays gd.Array, blend_shapes gd.Array, lods gd.Dictionary, flags classdb.MeshArrayFormat)  {
	var frame = callframe.New()
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, discreet.Get(arrays))
	callframe.Arg(frame, discreet.Get(blend_shapes))
	callframe.Arg(frame, discreet.Get(lods))
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_add_surface_from_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all surfaces from this [ArrayMesh].
*/
//go:nosplit
func (self class) ClearSurfaces()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_clear_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SurfaceUpdateVertexRegion(surf_idx gd.Int, offset gd.Int, data gd.PackedByteArray)  {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, discreet.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_update_vertex_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SurfaceUpdateAttributeRegion(surf_idx gd.Int, offset gd.Int, data gd.PackedByteArray)  {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, discreet.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_surface_update_attribute_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SurfaceUpdateSkinRegion(surf_idx gd.Int, offset gd.Int, data gd.PackedByteArray)  {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, discreet.Get(data))
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
	callframe.Arg(frame, discreet.Get(name))
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
func (self class) SurfaceSetName(surf_idx gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, discreet.Get(name))
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
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Regenerates tangents for each of the [ArrayMesh]'s surfaces.
*/
//go:nosplit
func (self class) RegenNormalMaps()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_regen_normal_maps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Performs a UV unwrap on the [ArrayMesh] to prepare the mesh for lightmapping.
*/
//go:nosplit
func (self class) LightmapUnwrap(transform gd.Transform3D, texel_size gd.Float) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	callframe.Arg(frame, texel_size)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_lightmap_unwrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB)  {
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
func (self class) SetShadowMesh(mesh gdclass.ArrayMesh)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_set_shadow_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowMesh() gdclass.ArrayMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayMesh.Bind_get_shadow_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.ArrayMesh{classdb.ArrayMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsArrayMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsArrayMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.GD { return *((*Mesh.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMesh() Mesh.Go { return *((*Mesh.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMesh(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMesh(), name)
	}
}
func init() {classdb.Register("ArrayMesh", func(ptr gd.Object) any { return classdb.ArrayMesh(ptr) })}
