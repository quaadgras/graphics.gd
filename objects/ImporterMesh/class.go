package ImporterMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2i"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
ImporterMesh is a type of [Resource] analogous to [ArrayMesh]. It contains vertex array-based geometry, divided in [i]surfaces[/i]. Each surface contains a completely separate array and a material used to draw it. Design wise, a mesh with multiple surfaces is preferred to a single surface, because objects created in 3D editing software commonly contain multiple materials.
Unlike its runtime counterpart, [ImporterMesh] contains mesh data before various import steps, such as lod and shadow mesh generation, have taken place. Modify surface data by calling [method clear], followed by [method add_surface] for each surface.
*/
type Instance [1]classdb.ImporterMesh
type Any interface {
	gd.IsClass
	AsImporterMesh() Instance
}

/*
Adds name for a blend shape that will be added with [method add_surface]. Must be called before surface is added.
*/
func (self Instance) AddBlendShape(name string) {
	class(self).AddBlendShape(gd.NewString(name))
}

/*
Returns the number of blend shapes that the mesh holds.
*/
func (self Instance) GetBlendShapeCount() int {
	return int(int(class(self).GetBlendShapeCount()))
}

/*
Returns the name of the blend shape at this index.
*/
func (self Instance) GetBlendShapeName(blend_shape_idx int) string {
	return string(class(self).GetBlendShapeName(gd.Int(blend_shape_idx)).String())
}

/*
Sets the blend shape mode to one of [enum Mesh.BlendShapeMode].
*/
func (self Instance) SetBlendShapeMode(mode classdb.MeshBlendShapeMode) {
	class(self).SetBlendShapeMode(mode)
}

/*
Returns the blend shape mode for this Mesh.
*/
func (self Instance) GetBlendShapeMode() classdb.MeshBlendShapeMode {
	return classdb.MeshBlendShapeMode(class(self).GetBlendShapeMode())
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
func (self Instance) AddSurface(primitive classdb.MeshPrimitiveType, arrays Array.Any) {
	class(self).AddSurface(primitive, arrays, [1]gd.Array{}[0], [1]Dictionary.Any{}[0], [1]objects.Material{}[0], gd.NewString(""), gd.Int(0))
}

/*
Returns the number of surfaces that the mesh holds.
*/
func (self Instance) GetSurfaceCount() int {
	return int(int(class(self).GetSurfaceCount()))
}

/*
Returns the primitive type of the requested surface (see [method add_surface]).
*/
func (self Instance) GetSurfacePrimitiveType(surface_idx int) classdb.MeshPrimitiveType {
	return classdb.MeshPrimitiveType(class(self).GetSurfacePrimitiveType(gd.Int(surface_idx)))
}

/*
Gets the name assigned to this surface.
*/
func (self Instance) GetSurfaceName(surface_idx int) string {
	return string(class(self).GetSurfaceName(gd.Int(surface_idx)).String())
}

/*
Returns the arrays for the vertices, normals, UVs, etc. that make up the requested surface. See [method add_surface].
*/
func (self Instance) GetSurfaceArrays(surface_idx int) Array.Any {
	return Array.Any(class(self).GetSurfaceArrays(gd.Int(surface_idx)))
}

/*
Returns a single set of blend shape arrays for the requested blend shape index for a surface.
*/
func (self Instance) GetSurfaceBlendShapeArrays(surface_idx int, blend_shape_idx int) Array.Any {
	return Array.Any(class(self).GetSurfaceBlendShapeArrays(gd.Int(surface_idx), gd.Int(blend_shape_idx)))
}

/*
Returns the number of lods that the mesh holds on a given surface.
*/
func (self Instance) GetSurfaceLodCount(surface_idx int) int {
	return int(int(class(self).GetSurfaceLodCount(gd.Int(surface_idx))))
}

/*
Returns the screen ratio which activates a lod for a surface.
*/
func (self Instance) GetSurfaceLodSize(surface_idx int, lod_idx int) Float.X {
	return Float.X(Float.X(class(self).GetSurfaceLodSize(gd.Int(surface_idx), gd.Int(lod_idx))))
}

/*
Returns the index buffer of a lod for a surface.
*/
func (self Instance) GetSurfaceLodIndices(surface_idx int, lod_idx int) []int32 {
	return []int32(class(self).GetSurfaceLodIndices(gd.Int(surface_idx), gd.Int(lod_idx)).AsSlice())
}

/*
Returns a [Material] in a given surface. Surface is rendered using this material.
*/
func (self Instance) GetSurfaceMaterial(surface_idx int) objects.Material {
	return objects.Material(class(self).GetSurfaceMaterial(gd.Int(surface_idx)))
}

/*
Returns the format of the surface that the mesh holds.
*/
func (self Instance) GetSurfaceFormat(surface_idx int) int {
	return int(int(class(self).GetSurfaceFormat(gd.Int(surface_idx))))
}

/*
Sets a name for a given surface.
*/
func (self Instance) SetSurfaceName(surface_idx int, name string) {
	class(self).SetSurfaceName(gd.Int(surface_idx), gd.NewString(name))
}

/*
Sets a [Material] for a given surface. Surface will be rendered using this material.
*/
func (self Instance) SetSurfaceMaterial(surface_idx int, material objects.Material) {
	class(self).SetSurfaceMaterial(gd.Int(surface_idx), material)
}

/*
Generates all lods for this ImporterMesh.
[param normal_merge_angle] and [param normal_split_angle] are in degrees and used in the same way as the importer settings in [code]lods[/code]. As a good default, use 25 and 60 respectively.
The number of generated lods can be accessed using [method get_surface_lod_count], and each LOD is available in [method get_surface_lod_size] and [method get_surface_lod_indices].
[param bone_transform_array] is an [Array] which can be either empty or contain [Transform3D]s which, for each of the mesh's bone IDs, will apply mesh skinning when generating the LOD mesh variations. This is usually used to account for discrepancies in scale between the mesh itself and its skinning data.
*/
func (self Instance) GenerateLods(normal_merge_angle Float.X, normal_split_angle Float.X, bone_transform_array Array.Any) {
	class(self).GenerateLods(gd.Float(normal_merge_angle), gd.Float(normal_split_angle), bone_transform_array)
}

/*
Returns the mesh data represented by this [ImporterMesh] as a usable [ArrayMesh].
This method caches the returned mesh, and subsequent calls will return the cached data until [method clear] is called.
If not yet cached and [param base_mesh] is provided, [param base_mesh] will be used and mutated.
*/
func (self Instance) GetMesh() objects.ArrayMesh {
	return objects.ArrayMesh(class(self).GetMesh([1]objects.ArrayMesh{}[0]))
}

/*
Removes all surfaces and blend shapes from this [ImporterMesh].
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Sets the size hint of this mesh for lightmap-unwrapping in UV-space.
*/
func (self Instance) SetLightmapSizeHint(size Vector2i.XY) {
	class(self).SetLightmapSizeHint(gd.Vector2i(size))
}

/*
Returns the size hint of this mesh for lightmap-unwrapping in UV-space.
*/
func (self Instance) GetLightmapSizeHint() Vector2i.XY {
	return Vector2i.XY(class(self).GetLightmapSizeHint())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ImporterMesh

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImporterMesh"))
	return Instance{*(*classdb.ImporterMesh)(unsafe.Pointer(&object))}
}

/*
Adds name for a blend shape that will be added with [method add_surface]. Must be called before surface is added.
*/
//go:nosplit
func (self class) AddBlendShape(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_add_blend_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of blend shapes that the mesh holds.
*/
//go:nosplit
func (self class) GetBlendShapeCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_blend_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the blend shape at this index.
*/
//go:nosplit
func (self class) GetBlendShapeName(blend_shape_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_blend_shape_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the blend shape mode to one of [enum Mesh.BlendShapeMode].
*/
//go:nosplit
func (self class) SetBlendShapeMode(mode classdb.MeshBlendShapeMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_set_blend_shape_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the blend shape mode for this Mesh.
*/
//go:nosplit
func (self class) GetBlendShapeMode() classdb.MeshBlendShapeMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MeshBlendShapeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_blend_shape_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddSurface(primitive classdb.MeshPrimitiveType, arrays gd.Array, blend_shapes gd.Array, lods gd.Dictionary, material objects.Material, name gd.String, flags gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, pointers.Get(arrays))
	callframe.Arg(frame, pointers.Get(blend_shapes))
	callframe.Arg(frame, pointers.Get(lods))
	callframe.Arg(frame, pointers.Get(material[0])[0])
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_add_surface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of surfaces that the mesh holds.
*/
//go:nosplit
func (self class) GetSurfaceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the primitive type of the requested surface (see [method add_surface]).
*/
//go:nosplit
func (self class) GetSurfacePrimitiveType(surface_idx gd.Int) classdb.MeshPrimitiveType {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	var r_ret = callframe.Ret[classdb.MeshPrimitiveType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_primitive_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the name assigned to this surface.
*/
//go:nosplit
func (self class) GetSurfaceName(surface_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the arrays for the vertices, normals, UVs, etc. that make up the requested surface. See [method add_surface].
*/
//go:nosplit
func (self class) GetSurfaceArrays(surface_idx gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a single set of blend shape arrays for the requested blend shape index for a surface.
*/
//go:nosplit
func (self class) GetSurfaceBlendShapeArrays(surface_idx gd.Int, blend_shape_idx gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	callframe.Arg(frame, blend_shape_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_blend_shape_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of lods that the mesh holds on a given surface.
*/
//go:nosplit
func (self class) GetSurfaceLodCount(surface_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_lod_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the screen ratio which activates a lod for a surface.
*/
//go:nosplit
func (self class) GetSurfaceLodSize(surface_idx gd.Int, lod_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	callframe.Arg(frame, lod_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_lod_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index buffer of a lod for a surface.
*/
//go:nosplit
func (self class) GetSurfaceLodIndices(surface_idx gd.Int, lod_idx gd.Int) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	callframe.Arg(frame, lod_idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_lod_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a [Material] in a given surface. Surface is rendered using this material.
*/
//go:nosplit
func (self class) GetSurfaceMaterial(surface_idx gd.Int) objects.Material {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Material{gd.PointerWithOwnershipTransferredToGo[classdb.Material](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the format of the surface that the mesh holds.
*/
//go:nosplit
func (self class) GetSurfaceFormat(surface_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_surface_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a name for a given surface.
*/
//go:nosplit
func (self class) SetSurfaceName(surface_idx gd.Int, name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_set_surface_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets a [Material] for a given surface. Surface will be rendered using this material.
*/
//go:nosplit
func (self class) SetSurfaceMaterial(surface_idx gd.Int, material objects.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, surface_idx)
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_set_surface_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Generates all lods for this ImporterMesh.
[param normal_merge_angle] and [param normal_split_angle] are in degrees and used in the same way as the importer settings in [code]lods[/code]. As a good default, use 25 and 60 respectively.
The number of generated lods can be accessed using [method get_surface_lod_count], and each LOD is available in [method get_surface_lod_size] and [method get_surface_lod_indices].
[param bone_transform_array] is an [Array] which can be either empty or contain [Transform3D]s which, for each of the mesh's bone IDs, will apply mesh skinning when generating the LOD mesh variations. This is usually used to account for discrepancies in scale between the mesh itself and its skinning data.
*/
//go:nosplit
func (self class) GenerateLods(normal_merge_angle gd.Float, normal_split_angle gd.Float, bone_transform_array gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, normal_merge_angle)
	callframe.Arg(frame, normal_split_angle)
	callframe.Arg(frame, pointers.Get(bone_transform_array))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_generate_lods, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the mesh data represented by this [ImporterMesh] as a usable [ArrayMesh].
This method caches the returned mesh, and subsequent calls will return the cached data until [method clear] is called.
If not yet cached and [param base_mesh] is provided, [param base_mesh] will be used and mutated.
*/
//go:nosplit
func (self class) GetMesh(base_mesh objects.ArrayMesh) objects.ArrayMesh {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base_mesh[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ArrayMesh{gd.PointerWithOwnershipTransferredToGo[classdb.ArrayMesh](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Removes all surfaces and blend shapes from this [ImporterMesh].
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the size hint of this mesh for lightmap-unwrapping in UV-space.
*/
//go:nosplit
func (self class) SetLightmapSizeHint(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_set_lightmap_size_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the size hint of this mesh for lightmap-unwrapping in UV-space.
*/
//go:nosplit
func (self class) GetLightmapSizeHint() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMesh.Bind_get_lightmap_size_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsImporterMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImporterMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("ImporterMesh", func(ptr gd.Object) any {
		return [1]classdb.ImporterMesh{*(*classdb.ImporterMesh)(unsafe.Pointer(&ptr))}
	})
}
