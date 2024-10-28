package SurfaceTool

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The [SurfaceTool] is used to construct a [Mesh] by specifying vertex attributes individually. It can be used to construct a [Mesh] from a script. All properties except indices need to be added before calling [method add_vertex]. For example, to add vertex colors and UVs:
[codeblocks]
[gdscript]
var st = SurfaceTool.new()
st.begin(Mesh.PRIMITIVE_TRIANGLES)
st.set_color(Color(1, 0, 0))
st.set_uv(Vector2(0, 0))
st.add_vertex(Vector3(0, 0, 0))
[/gdscript]
[csharp]
var st = new SurfaceTool();
st.Begin(Mesh.PrimitiveType.Triangles);
st.SetColor(new Color(1, 0, 0));
st.SetUV(new Vector2(0, 0));
st.AddVertex(new Vector3(0, 0, 0));
[/csharp]
[/codeblocks]
The above [SurfaceTool] now contains one vertex of a triangle which has a UV coordinate and a specified [Color]. If another vertex were added without calling [method set_uv] or [method set_color], then the last values would be used.
Vertex attributes must be passed [b]before[/b] calling [method add_vertex]. Failure to do so will result in an error when committing the vertex information to a mesh.
Additionally, the attributes used before the first vertex is added determine the format of the mesh. For example, if you only add UVs to the first vertex, you cannot add color to any of the subsequent vertices.
See also [ArrayMesh], [ImmediateMesh] and [MeshDataTool] for procedural geometry generation.
[b]Note:[/b] Godot uses clockwise [url=https://learnopengl.com/Advanced-OpenGL/Face-culling]winding order[/url] for front faces of triangle primitive modes.

*/
type Go [1]classdb.SurfaceTool

/*
Set to [constant SKIN_8_WEIGHTS] to indicate that up to 8 bone influences per vertex may be used.
By default, only 4 bone influences are used ([constant SKIN_4_WEIGHTS])
[b]Note:[/b] This function takes an enum, not the exact number of weights.
*/
func (self Go) SetSkinWeightCount(count classdb.SurfaceToolSkinWeightCount) {
	class(self).SetSkinWeightCount(count)
}

/*
By default, returns [constant SKIN_4_WEIGHTS] to indicate only 4 bone influences per vertex are used.
Returns [constant SKIN_8_WEIGHTS] if up to 8 influences are used.
[b]Note:[/b] This function returns an enum, not the exact number of weights.
*/
func (self Go) GetSkinWeightCount() classdb.SurfaceToolSkinWeightCount {
	return classdb.SurfaceToolSkinWeightCount(class(self).GetSkinWeightCount())
}

/*
Sets the color format for this custom [param channel_index]. Use [constant CUSTOM_MAX] to disable.
Must be invoked after [method begin] and should be set before [method commit] or [method commit_to_arrays].
*/
func (self Go) SetCustomFormat(channel_index int, format classdb.SurfaceToolCustomFormat) {
	class(self).SetCustomFormat(gd.Int(channel_index), format)
}

/*
Returns the format for custom [param channel_index] (currently up to 4). Returns [constant CUSTOM_MAX] if this custom channel is unused.
*/
func (self Go) GetCustomFormat(channel_index int) classdb.SurfaceToolCustomFormat {
	return classdb.SurfaceToolCustomFormat(class(self).GetCustomFormat(gd.Int(channel_index)))
}

/*
Called before adding any vertices. Takes the primitive type as an argument (e.g. [constant Mesh.PRIMITIVE_TRIANGLES]).
*/
func (self Go) Begin(primitive classdb.MeshPrimitiveType) {
	class(self).Begin(primitive)
}

/*
Specifies the position of current vertex. Should be called after specifying other vertex properties (e.g. Color, UV).
*/
func (self Go) AddVertex(vertex gd.Vector3) {
	class(self).AddVertex(vertex)
}

/*
Specifies a [Color] to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
[b]Note:[/b] The material must have [member BaseMaterial3D.vertex_color_use_as_albedo] enabled for the vertex color to be visible.
*/
func (self Go) SetColor(color gd.Color) {
	class(self).SetColor(color)
}

/*
Specifies a normal to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
func (self Go) SetNormal(normal gd.Vector3) {
	class(self).SetNormal(normal)
}

/*
Specifies a tangent to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
func (self Go) SetTangent(tangent gd.Plane) {
	class(self).SetTangent(tangent)
}

/*
Specifies a set of UV coordinates to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
func (self Go) SetUv(uv gd.Vector2) {
	class(self).SetUv(uv)
}

/*
Specifies an optional second set of UV coordinates to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
func (self Go) SetUv2(uv2 gd.Vector2) {
	class(self).SetUv2(uv2)
}

/*
Specifies an array of bones to use for the [i]next[/i] vertex. [param bones] must contain 4 integers.
*/
func (self Go) SetBones(bones []int32) {
	class(self).SetBones(gd.NewPackedInt32Slice(bones))
}

/*
Specifies weight values to use for the [i]next[/i] vertex. [param weights] must contain 4 values. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
func (self Go) SetWeights(weights []float32) {
	class(self).SetWeights(gd.NewPackedFloat32Slice(weights))
}

/*
Sets the custom value on this vertex for [param channel_index].
[method set_custom_format] must be called first for this [param channel_index]. Formats which are not RGBA will ignore other color channels.
*/
func (self Go) SetCustom(channel_index int, custom_color gd.Color) {
	class(self).SetCustom(gd.Int(channel_index), custom_color)
}

/*
Specifies the smooth group to use for the [i]next[/i] vertex. If this is never called, all vertices will have the default smooth group of [code]0[/code] and will be smoothed with adjacent vertices of the same group. To produce a mesh with flat normals, set the smooth group to [code]-1[/code].
[b]Note:[/b] This function actually takes a [code]uint32_t[/code], so C# users should use [code]uint32.MaxValue[/code] instead of [code]-1[/code] to produce a mesh with flat normals.
*/
func (self Go) SetSmoothGroup(index int) {
	class(self).SetSmoothGroup(gd.Int(index))
}

/*
Inserts a triangle fan made of array data into [Mesh] being constructed.
Requires the primitive type be set to [constant Mesh.PRIMITIVE_TRIANGLES].
*/
func (self Go) AddTriangleFan(vertices []gd.Vector3) {
	class(self).AddTriangleFan(gd.NewPackedVector3Slice(vertices), gd.NewPackedVector2Slice(([1][]gd.Vector2{}[0])), gd.NewPackedColorSlice(([1][]gd.Color{}[0])), gd.NewPackedVector2Slice(([1][]gd.Vector2{}[0])), gd.NewPackedVector3Slice(([1][]gd.Vector3{}[0])), ([1]gd.Array{}[0]))
}

/*
Adds a vertex to index array if you are using indexed vertices. Does not need to be called before adding vertices.
*/
func (self Go) AddIndex(index int) {
	class(self).AddIndex(gd.Int(index))
}

/*
Shrinks the vertex array by creating an index array. This can improve performance by avoiding vertex reuse.
*/
func (self Go) Index() {
	class(self).Index()
}

/*
Removes the index array by expanding the vertex array.
*/
func (self Go) Deindex() {
	class(self).Deindex()
}

/*
Generates normals from vertices so you do not have to do it manually. If [param flip] is [code]true[/code], the resulting normals will be inverted. [method generate_normals] should be called [i]after[/i] generating geometry and [i]before[/i] committing the mesh using [method commit] or [method commit_to_arrays]. For correct display of normal-mapped surfaces, you will also have to generate tangents using [method generate_tangents].
[b]Note:[/b] [method generate_normals] only works if the primitive type to be set to [constant Mesh.PRIMITIVE_TRIANGLES].
[b]Note:[/b] [method generate_normals] takes smooth groups into account. To generate smooth normals, set the smooth group to a value greater than or equal to [code]0[/code] using [method set_smooth_group] or leave the smooth group at the default of [code]0[/code]. To generate flat normals, set the smooth group to [code]-1[/code] using [method set_smooth_group] prior to adding vertices.
*/
func (self Go) GenerateNormals() {
	class(self).GenerateNormals(false)
}

/*
Generates a tangent vector for each vertex. Requires that each vertex have UVs and normals set already (see [method generate_normals]).
*/
func (self Go) GenerateTangents() {
	class(self).GenerateTangents()
}

/*
Optimizes triangle sorting for performance. Requires that [method get_primitive_type] is [constant Mesh.PRIMITIVE_TRIANGLES].
*/
func (self Go) OptimizeIndicesForCache() {
	class(self).OptimizeIndicesForCache()
}

/*
Returns the axis-aligned bounding box of the vertex positions.
*/
func (self Go) GetAabb() gd.AABB {
	return gd.AABB(class(self).GetAabb())
}

/*
Generates an LOD for a given [param nd_threshold] in linear units (square root of quadric error metric), using at most [param target_index_count] indices.
*/
func (self Go) GenerateLod(nd_threshold float64) []int32 {
	return []int32(class(self).GenerateLod(gd.Float(nd_threshold), gd.Int(3)).AsSlice())
}

/*
Sets [Material] to be used by the [Mesh] you are constructing.
*/
func (self Go) SetMaterial(material gdclass.Material) {
	class(self).SetMaterial(material)
}

/*
Returns the type of mesh geometry, such as [constant Mesh.PRIMITIVE_TRIANGLES].
*/
func (self Go) GetPrimitiveType() classdb.MeshPrimitiveType {
	return classdb.MeshPrimitiveType(class(self).GetPrimitiveType())
}

/*
Clear all information passed into the surface tool so far.
*/
func (self Go) Clear() {
	class(self).Clear()
}

/*
Creates a vertex array from an existing [Mesh].
*/
func (self Go) CreateFrom(existing gdclass.Mesh, surface int) {
	class(self).CreateFrom(existing, gd.Int(surface))
}

/*
Creates this SurfaceTool from existing vertex arrays such as returned by [method commit_to_arrays], [method Mesh.surface_get_arrays], [method Mesh.surface_get_blend_shape_arrays], [method ImporterMesh.get_surface_arrays], and [method ImporterMesh.get_surface_blend_shape_arrays]. [param primitive_type] controls the type of mesh data, defaulting to [constant Mesh.PRIMITIVE_TRIANGLES].
*/
func (self Go) CreateFromArrays(arrays gd.Array) {
	class(self).CreateFromArrays(arrays, 3)
}

/*
Creates a vertex array from the specified blend shape of an existing [Mesh]. This can be used to extract a specific pose from a blend shape.
*/
func (self Go) CreateFromBlendShape(existing gdclass.Mesh, surface int, blend_shape string) {
	class(self).CreateFromBlendShape(existing, gd.Int(surface), gd.NewString(blend_shape))
}

/*
Append vertices from a given [Mesh] surface onto the current vertex array with specified [Transform3D].
*/
func (self Go) AppendFrom(existing gdclass.Mesh, surface int, transform gd.Transform3D) {
	class(self).AppendFrom(existing, gd.Int(surface), transform)
}

/*
Returns a constructed [ArrayMesh] from current information passed in. If an existing [ArrayMesh] is passed in as an argument, will add an extra surface to the existing [ArrayMesh].
[b]FIXME:[/b] Document possible values for [param flags], it changed in 4.0. Likely some combinations of [enum Mesh.ArrayFormat].
*/
func (self Go) Commit() gdclass.ArrayMesh {
	return gdclass.ArrayMesh(class(self).Commit(([1]gdclass.ArrayMesh{}[0]), gd.Int(0)))
}

/*
Commits the data to the same format used by [method ArrayMesh.add_surface_from_arrays], [method ImporterMesh.add_surface], and [method create_from_arrays]. This way you can further process the mesh data using the [ArrayMesh] or [ImporterMesh] APIs.
*/
func (self Go) CommitToArrays() gd.Array {
	return gd.Array(class(self).CommitToArrays())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SurfaceTool
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SurfaceTool"))
	return Go{classdb.SurfaceTool(object)}
}

/*
Set to [constant SKIN_8_WEIGHTS] to indicate that up to 8 bone influences per vertex may be used.
By default, only 4 bone influences are used ([constant SKIN_4_WEIGHTS])
[b]Note:[/b] This function takes an enum, not the exact number of weights.
*/
//go:nosplit
func (self class) SetSkinWeightCount(count classdb.SurfaceToolSkinWeightCount)  {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_skin_weight_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
By default, returns [constant SKIN_4_WEIGHTS] to indicate only 4 bone influences per vertex are used.
Returns [constant SKIN_8_WEIGHTS] if up to 8 influences are used.
[b]Note:[/b] This function returns an enum, not the exact number of weights.
*/
//go:nosplit
func (self class) GetSkinWeightCount() classdb.SurfaceToolSkinWeightCount {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SurfaceToolSkinWeightCount](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_get_skin_weight_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the color format for this custom [param channel_index]. Use [constant CUSTOM_MAX] to disable.
Must be invoked after [method begin] and should be set before [method commit] or [method commit_to_arrays].
*/
//go:nosplit
func (self class) SetCustomFormat(channel_index gd.Int, format classdb.SurfaceToolCustomFormat)  {
	var frame = callframe.New()
	callframe.Arg(frame, channel_index)
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_custom_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the format for custom [param channel_index] (currently up to 4). Returns [constant CUSTOM_MAX] if this custom channel is unused.
*/
//go:nosplit
func (self class) GetCustomFormat(channel_index gd.Int) classdb.SurfaceToolCustomFormat {
	var frame = callframe.New()
	callframe.Arg(frame, channel_index)
	var r_ret = callframe.Ret[classdb.SurfaceToolCustomFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_get_custom_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Called before adding any vertices. Takes the primitive type as an argument (e.g. [constant Mesh.PRIMITIVE_TRIANGLES]).
*/
//go:nosplit
func (self class) Begin(primitive classdb.MeshPrimitiveType)  {
	var frame = callframe.New()
	callframe.Arg(frame, primitive)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies the position of current vertex. Should be called after specifying other vertex properties (e.g. Color, UV).
*/
//go:nosplit
func (self class) AddVertex(vertex gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_add_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies a [Color] to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
[b]Note:[/b] The material must have [member BaseMaterial3D.vertex_color_use_as_albedo] enabled for the vertex color to be visible.
*/
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies a normal to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
//go:nosplit
func (self class) SetNormal(normal gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies a tangent to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
//go:nosplit
func (self class) SetTangent(tangent gd.Plane)  {
	var frame = callframe.New()
	callframe.Arg(frame, tangent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies a set of UV coordinates to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
//go:nosplit
func (self class) SetUv(uv gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, uv)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies an optional second set of UV coordinates to use for the [i]next[/i] vertex. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
//go:nosplit
func (self class) SetUv2(uv2 gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, uv2)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies an array of bones to use for the [i]next[/i] vertex. [param bones] must contain 4 integers.
*/
//go:nosplit
func (self class) SetBones(bones gd.PackedInt32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(bones))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies weight values to use for the [i]next[/i] vertex. [param weights] must contain 4 values. If every vertex needs to have this information set and you fail to submit it for the first vertex, this information may not be used at all.
*/
//go:nosplit
func (self class) SetWeights(weights gd.PackedFloat32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(weights))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the custom value on this vertex for [param channel_index].
[method set_custom_format] must be called first for this [param channel_index]. Formats which are not RGBA will ignore other color channels.
*/
//go:nosplit
func (self class) SetCustom(channel_index gd.Int, custom_color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, channel_index)
	callframe.Arg(frame, custom_color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_custom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Specifies the smooth group to use for the [i]next[/i] vertex. If this is never called, all vertices will have the default smooth group of [code]0[/code] and will be smoothed with adjacent vertices of the same group. To produce a mesh with flat normals, set the smooth group to [code]-1[/code].
[b]Note:[/b] This function actually takes a [code]uint32_t[/code], so C# users should use [code]uint32.MaxValue[/code] instead of [code]-1[/code] to produce a mesh with flat normals.
*/
//go:nosplit
func (self class) SetSmoothGroup(index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_smooth_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Inserts a triangle fan made of array data into [Mesh] being constructed.
Requires the primitive type be set to [constant Mesh.PRIMITIVE_TRIANGLES].
*/
//go:nosplit
func (self class) AddTriangleFan(vertices gd.PackedVector3Array, uvs gd.PackedVector2Array, colors gd.PackedColorArray, uv2s gd.PackedVector2Array, normals gd.PackedVector3Array, tangents gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(vertices))
	callframe.Arg(frame, discreet.Get(uvs))
	callframe.Arg(frame, discreet.Get(colors))
	callframe.Arg(frame, discreet.Get(uv2s))
	callframe.Arg(frame, discreet.Get(normals))
	callframe.Arg(frame, discreet.Get(tangents))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_add_triangle_fan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a vertex to index array if you are using indexed vertices. Does not need to be called before adding vertices.
*/
//go:nosplit
func (self class) AddIndex(index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_add_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Shrinks the vertex array by creating an index array. This can improve performance by avoiding vertex reuse.
*/
//go:nosplit
func (self class) Index()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the index array by expanding the vertex array.
*/
//go:nosplit
func (self class) Deindex()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_deindex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates normals from vertices so you do not have to do it manually. If [param flip] is [code]true[/code], the resulting normals will be inverted. [method generate_normals] should be called [i]after[/i] generating geometry and [i]before[/i] committing the mesh using [method commit] or [method commit_to_arrays]. For correct display of normal-mapped surfaces, you will also have to generate tangents using [method generate_tangents].
[b]Note:[/b] [method generate_normals] only works if the primitive type to be set to [constant Mesh.PRIMITIVE_TRIANGLES].
[b]Note:[/b] [method generate_normals] takes smooth groups into account. To generate smooth normals, set the smooth group to a value greater than or equal to [code]0[/code] using [method set_smooth_group] or leave the smooth group at the default of [code]0[/code]. To generate flat normals, set the smooth group to [code]-1[/code] using [method set_smooth_group] prior to adding vertices.
*/
//go:nosplit
func (self class) GenerateNormals(flip bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, flip)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_generate_normals, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates a tangent vector for each vertex. Requires that each vertex have UVs and normals set already (see [method generate_normals]).
*/
//go:nosplit
func (self class) GenerateTangents()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_generate_tangents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Optimizes triangle sorting for performance. Requires that [method get_primitive_type] is [constant Mesh.PRIMITIVE_TRIANGLES].
*/
//go:nosplit
func (self class) OptimizeIndicesForCache()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_optimize_indices_for_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the axis-aligned bounding box of the vertex positions.
*/
//go:nosplit
func (self class) GetAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Generates an LOD for a given [param nd_threshold] in linear units (square root of quadric error metric), using at most [param target_index_count] indices.
*/
//go:nosplit
func (self class) GenerateLod(nd_threshold gd.Float, target_index_count gd.Int) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, nd_threshold)
	callframe.Arg(frame, target_index_count)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_generate_lod, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets [Material] to be used by the [Mesh] you are constructing.
*/
//go:nosplit
func (self class) SetMaterial(material gdclass.Material)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the type of mesh geometry, such as [constant Mesh.PRIMITIVE_TRIANGLES].
*/
//go:nosplit
func (self class) GetPrimitiveType() classdb.MeshPrimitiveType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MeshPrimitiveType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_get_primitive_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clear all information passed into the surface tool so far.
*/
//go:nosplit
func (self class) Clear()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a vertex array from an existing [Mesh].
*/
//go:nosplit
func (self class) CreateFrom(existing gdclass.Mesh, surface gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(existing[0])[0])
	callframe.Arg(frame, surface)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_create_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates this SurfaceTool from existing vertex arrays such as returned by [method commit_to_arrays], [method Mesh.surface_get_arrays], [method Mesh.surface_get_blend_shape_arrays], [method ImporterMesh.get_surface_arrays], and [method ImporterMesh.get_surface_blend_shape_arrays]. [param primitive_type] controls the type of mesh data, defaulting to [constant Mesh.PRIMITIVE_TRIANGLES].
*/
//go:nosplit
func (self class) CreateFromArrays(arrays gd.Array, primitive_type classdb.MeshPrimitiveType)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(arrays))
	callframe.Arg(frame, primitive_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_create_from_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a vertex array from the specified blend shape of an existing [Mesh]. This can be used to extract a specific pose from a blend shape.
*/
//go:nosplit
func (self class) CreateFromBlendShape(existing gdclass.Mesh, surface gd.Int, blend_shape gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(existing[0])[0])
	callframe.Arg(frame, surface)
	callframe.Arg(frame, discreet.Get(blend_shape))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_create_from_blend_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Append vertices from a given [Mesh] surface onto the current vertex array with specified [Transform3D].
*/
//go:nosplit
func (self class) AppendFrom(existing gdclass.Mesh, surface gd.Int, transform gd.Transform3D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(existing[0])[0])
	callframe.Arg(frame, surface)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_append_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a constructed [ArrayMesh] from current information passed in. If an existing [ArrayMesh] is passed in as an argument, will add an extra surface to the existing [ArrayMesh].
[b]FIXME:[/b] Document possible values for [param flags], it changed in 4.0. Likely some combinations of [enum Mesh.ArrayFormat].
*/
//go:nosplit
func (self class) Commit(existing gdclass.ArrayMesh, flags gd.Int) gdclass.ArrayMesh {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(existing[0])[0])
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_commit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.ArrayMesh{classdb.ArrayMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Commits the data to the same format used by [method ArrayMesh.add_surface_from_arrays], [method ImporterMesh.add_surface], and [method create_from_arrays]. This way you can further process the mesh data using the [ArrayMesh] or [ImporterMesh] APIs.
*/
//go:nosplit
func (self class) CommitToArrays() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SurfaceTool.Bind_commit_to_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsSurfaceTool() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSurfaceTool() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("SurfaceTool", func(ptr gd.Object) any { return classdb.SurfaceTool(ptr) })}
type CustomFormat = classdb.SurfaceToolCustomFormat

const (
/*Limits range of data passed to [method set_custom] to unsigned normalized 0 to 1 stored in 8 bits per channel. See [constant Mesh.ARRAY_CUSTOM_RGBA8_UNORM].*/
	CustomRgba8Unorm CustomFormat = 0
/*Limits range of data passed to [method set_custom] to signed normalized -1 to 1 stored in 8 bits per channel. See [constant Mesh.ARRAY_CUSTOM_RGBA8_SNORM].*/
	CustomRgba8Snorm CustomFormat = 1
/*Stores data passed to [method set_custom] as half precision floats, and uses only red and green color channels. See [constant Mesh.ARRAY_CUSTOM_RG_HALF].*/
	CustomRgHalf CustomFormat = 2
/*Stores data passed to [method set_custom] as half precision floats and uses all color channels. See [constant Mesh.ARRAY_CUSTOM_RGBA_HALF].*/
	CustomRgbaHalf CustomFormat = 3
/*Stores data passed to [method set_custom] as full precision floats, and uses only red color channel. See [constant Mesh.ARRAY_CUSTOM_R_FLOAT].*/
	CustomRFloat CustomFormat = 4
/*Stores data passed to [method set_custom] as full precision floats, and uses only red and green color channels. See [constant Mesh.ARRAY_CUSTOM_RG_FLOAT].*/
	CustomRgFloat CustomFormat = 5
/*Stores data passed to [method set_custom] as full precision floats, and uses only red, green and blue color channels. See [constant Mesh.ARRAY_CUSTOM_RGB_FLOAT].*/
	CustomRgbFloat CustomFormat = 6
/*Stores data passed to [method set_custom] as full precision floats, and uses all color channels. See [constant Mesh.ARRAY_CUSTOM_RGBA_FLOAT].*/
	CustomRgbaFloat CustomFormat = 7
/*Used to indicate a disabled custom channel.*/
	CustomMax CustomFormat = 8
)
type SkinWeightCount = classdb.SurfaceToolSkinWeightCount

const (
/*Each individual vertex can be influenced by only 4 bone weights.*/
	Skin4Weights SkinWeightCount = 0
/*Each individual vertex can be influenced by up to 8 bone weights.*/
	Skin8Weights SkinWeightCount = 1
)
