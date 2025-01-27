// Package MeshDataTool provides methods for working with MeshDataTool object instances.
package MeshDataTool

import "unsafe"
import "reflect"
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
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Plane"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Color"

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

/*
MeshDataTool provides access to individual vertices in a [Mesh]. It allows users to read and edit vertex data of meshes. It also creates an array of faces and edges.
To use MeshDataTool, load a mesh with [method create_from_surface]. When you are finished editing the data commit the data to a mesh with [method commit_to_surface].
Below is an example of how MeshDataTool may be used.
[codeblocks]
[gdscript]
var mesh = ArrayMesh.new()
mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, BoxMesh.new().get_mesh_arrays())
var mdt = MeshDataTool.new()
mdt.create_from_surface(mesh, 0)
for i in range(mdt.get_vertex_count()):

	var vertex = mdt.get_vertex(i)
	# In this example we extend the mesh by one unit, which results in separated faces as it is flat shaded.
	vertex += mdt.get_vertex_normal(i)
	# Save your change.
	mdt.set_vertex(i, vertex)

mesh.clear_surfaces()
mdt.commit_to_surface(mesh)
var mi = MeshInstance.new()
mi.mesh = mesh
add_child(mi)
[/gdscript]
[csharp]
var mesh = new ArrayMesh();
mesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, new BoxMesh().GetMeshArrays());
var mdt = new MeshDataTool();
mdt.CreateFromSurface(mesh, 0);
for (var i = 0; i < mdt.GetVertexCount(); i++)

	{
	    Vector3 vertex = mdt.GetVertex(i);
	    // In this example we extend the mesh by one unit, which results in separated faces as it is flat shaded.
	    vertex += mdt.GetVertexNormal(i);
	    // Save your change.
	    mdt.SetVertex(i, vertex);
	}

mesh.ClearSurfaces();
mdt.CommitToSurface(mesh);
var mi = new MeshInstance();
mi.Mesh = mesh;
AddChild(mi);
[/csharp]
[/codeblocks]
See also [ArrayMesh], [ImmediateMesh] and [SurfaceTool] for procedural geometry generation.
[b]Note:[/b] Godot uses clockwise [url=https://learnopengl.com/Advanced-OpenGL/Face-culling]winding order[/url] for front faces of triangle primitive modes.
*/
type Instance [1]gdclass.MeshDataTool

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMeshDataTool() Instance
}

/*
Clears all data currently in MeshDataTool.
*/
func (self Instance) Clear() { //gd:MeshDataTool.clear
	class(self).Clear()
}

/*
Uses specified surface of given [Mesh] to populate data for MeshDataTool.
Requires [Mesh] with primitive type [constant Mesh.PRIMITIVE_TRIANGLES].
*/
func (self Instance) CreateFromSurface(mesh [1]gdclass.ArrayMesh, surface int) error { //gd:MeshDataTool.create_from_surface
	return error(gd.ToError(class(self).CreateFromSurface(mesh, gd.Int(surface))))
}

/*
Adds a new surface to specified [Mesh] with edited data.
*/
func (self Instance) CommitToSurface(mesh [1]gdclass.ArrayMesh) error { //gd:MeshDataTool.commit_to_surface
	return error(gd.ToError(class(self).CommitToSurface(mesh, gd.Int(0))))
}

/*
Returns the [Mesh]'s format as a combination of the [enum Mesh.ArrayFormat] flags. For example, a mesh containing both vertices and normals would return a format of [code]3[/code] because [constant Mesh.ARRAY_FORMAT_VERTEX] is [code]1[/code] and [constant Mesh.ARRAY_FORMAT_NORMAL] is [code]2[/code].
*/
func (self Instance) GetFormat() int { //gd:MeshDataTool.get_format
	return int(int(class(self).GetFormat()))
}

/*
Returns the total number of vertices in [Mesh].
*/
func (self Instance) GetVertexCount() int { //gd:MeshDataTool.get_vertex_count
	return int(int(class(self).GetVertexCount()))
}

/*
Returns the number of edges in this [Mesh].
*/
func (self Instance) GetEdgeCount() int { //gd:MeshDataTool.get_edge_count
	return int(int(class(self).GetEdgeCount()))
}

/*
Returns the number of faces in this [Mesh].
*/
func (self Instance) GetFaceCount() int { //gd:MeshDataTool.get_face_count
	return int(int(class(self).GetFaceCount()))
}

/*
Sets the position of the given vertex.
*/
func (self Instance) SetVertex(idx int, vertex Vector3.XYZ) { //gd:MeshDataTool.set_vertex
	class(self).SetVertex(gd.Int(idx), gd.Vector3(vertex))
}

/*
Returns the position of the given vertex.
*/
func (self Instance) GetVertex(idx int) Vector3.XYZ { //gd:MeshDataTool.get_vertex
	return Vector3.XYZ(class(self).GetVertex(gd.Int(idx)))
}

/*
Sets the normal of the given vertex.
*/
func (self Instance) SetVertexNormal(idx int, normal Vector3.XYZ) { //gd:MeshDataTool.set_vertex_normal
	class(self).SetVertexNormal(gd.Int(idx), gd.Vector3(normal))
}

/*
Returns the normal of the given vertex.
*/
func (self Instance) GetVertexNormal(idx int) Vector3.XYZ { //gd:MeshDataTool.get_vertex_normal
	return Vector3.XYZ(class(self).GetVertexNormal(gd.Int(idx)))
}

/*
Sets the tangent of the given vertex.
*/
func (self Instance) SetVertexTangent(idx int, tangent Plane.NormalD) { //gd:MeshDataTool.set_vertex_tangent
	class(self).SetVertexTangent(gd.Int(idx), gd.Plane(tangent))
}

/*
Returns the tangent of the given vertex.
*/
func (self Instance) GetVertexTangent(idx int) Plane.NormalD { //gd:MeshDataTool.get_vertex_tangent
	return Plane.NormalD(class(self).GetVertexTangent(gd.Int(idx)))
}

/*
Sets the UV of the given vertex.
*/
func (self Instance) SetVertexUv(idx int, uv Vector2.XY) { //gd:MeshDataTool.set_vertex_uv
	class(self).SetVertexUv(gd.Int(idx), gd.Vector2(uv))
}

/*
Returns the UV of the given vertex.
*/
func (self Instance) GetVertexUv(idx int) Vector2.XY { //gd:MeshDataTool.get_vertex_uv
	return Vector2.XY(class(self).GetVertexUv(gd.Int(idx)))
}

/*
Sets the UV2 of the given vertex.
*/
func (self Instance) SetVertexUv2(idx int, uv2 Vector2.XY) { //gd:MeshDataTool.set_vertex_uv2
	class(self).SetVertexUv2(gd.Int(idx), gd.Vector2(uv2))
}

/*
Returns the UV2 of the given vertex.
*/
func (self Instance) GetVertexUv2(idx int) Vector2.XY { //gd:MeshDataTool.get_vertex_uv2
	return Vector2.XY(class(self).GetVertexUv2(gd.Int(idx)))
}

/*
Sets the color of the given vertex.
*/
func (self Instance) SetVertexColor(idx int, color Color.RGBA) { //gd:MeshDataTool.set_vertex_color
	class(self).SetVertexColor(gd.Int(idx), gd.Color(color))
}

/*
Returns the color of the given vertex.
*/
func (self Instance) GetVertexColor(idx int) Color.RGBA { //gd:MeshDataTool.get_vertex_color
	return Color.RGBA(class(self).GetVertexColor(gd.Int(idx)))
}

/*
Sets the bones of the given vertex.
*/
func (self Instance) SetVertexBones(idx int, bones []int32) { //gd:MeshDataTool.set_vertex_bones
	class(self).SetVertexBones(gd.Int(idx), gd.NewPackedInt32Slice(bones))
}

/*
Returns the bones of the given vertex.
*/
func (self Instance) GetVertexBones(idx int) []int32 { //gd:MeshDataTool.get_vertex_bones
	return []int32(class(self).GetVertexBones(gd.Int(idx)).AsSlice())
}

/*
Sets the bone weights of the given vertex.
*/
func (self Instance) SetVertexWeights(idx int, weights []float32) { //gd:MeshDataTool.set_vertex_weights
	class(self).SetVertexWeights(gd.Int(idx), gd.NewPackedFloat32Slice(weights))
}

/*
Returns bone weights of the given vertex.
*/
func (self Instance) GetVertexWeights(idx int) []float32 { //gd:MeshDataTool.get_vertex_weights
	return []float32(class(self).GetVertexWeights(gd.Int(idx)).AsSlice())
}

/*
Sets the metadata associated with the given vertex.
*/
func (self Instance) SetVertexMeta(idx int, meta any) { //gd:MeshDataTool.set_vertex_meta
	class(self).SetVertexMeta(gd.Int(idx), gd.NewVariant(meta))
}

/*
Returns the metadata associated with the given vertex.
*/
func (self Instance) GetVertexMeta(idx int) any { //gd:MeshDataTool.get_vertex_meta
	return any(class(self).GetVertexMeta(gd.Int(idx)).Interface())
}

/*
Returns an array of edges that share the given vertex.
*/
func (self Instance) GetVertexEdges(idx int) []int32 { //gd:MeshDataTool.get_vertex_edges
	return []int32(class(self).GetVertexEdges(gd.Int(idx)).AsSlice())
}

/*
Returns an array of faces that share the given vertex.
*/
func (self Instance) GetVertexFaces(idx int) []int32 { //gd:MeshDataTool.get_vertex_faces
	return []int32(class(self).GetVertexFaces(gd.Int(idx)).AsSlice())
}

/*
Returns index of specified vertex connected to given edge.
Vertex argument can only be 0 or 1 because edges are comprised of two vertices.
*/
func (self Instance) GetEdgeVertex(idx int, vertex int) int { //gd:MeshDataTool.get_edge_vertex
	return int(int(class(self).GetEdgeVertex(gd.Int(idx), gd.Int(vertex))))
}

/*
Returns array of faces that touch given edge.
*/
func (self Instance) GetEdgeFaces(idx int) []int32 { //gd:MeshDataTool.get_edge_faces
	return []int32(class(self).GetEdgeFaces(gd.Int(idx)).AsSlice())
}

/*
Sets the metadata of the given edge.
*/
func (self Instance) SetEdgeMeta(idx int, meta any) { //gd:MeshDataTool.set_edge_meta
	class(self).SetEdgeMeta(gd.Int(idx), gd.NewVariant(meta))
}

/*
Returns meta information assigned to given edge.
*/
func (self Instance) GetEdgeMeta(idx int) any { //gd:MeshDataTool.get_edge_meta
	return any(class(self).GetEdgeMeta(gd.Int(idx)).Interface())
}

/*
Returns the specified vertex index of the given face.
Vertex argument must be either 0, 1, or 2 because faces contain three vertices.
[b]Example:[/b]
[codeblocks]
[gdscript]
var index = mesh_data_tool.get_face_vertex(0, 1) # Gets the index of the second vertex of the first face.
var position = mesh_data_tool.get_vertex(index)
var normal = mesh_data_tool.get_vertex_normal(index)
[/gdscript]
[csharp]
int index = meshDataTool.GetFaceVertex(0, 1); // Gets the index of the second vertex of the first face.
Vector3 position = meshDataTool.GetVertex(index);
Vector3 normal = meshDataTool.GetVertexNormal(index);
[/csharp]
[/codeblocks]
*/
func (self Instance) GetFaceVertex(idx int, vertex int) int { //gd:MeshDataTool.get_face_vertex
	return int(int(class(self).GetFaceVertex(gd.Int(idx), gd.Int(vertex))))
}

/*
Returns specified edge associated with given face.
Edge argument must be either 0, 1, or 2 because a face only has three edges.
*/
func (self Instance) GetFaceEdge(idx int, edge int) int { //gd:MeshDataTool.get_face_edge
	return int(int(class(self).GetFaceEdge(gd.Int(idx), gd.Int(edge))))
}

/*
Sets the metadata of the given face.
*/
func (self Instance) SetFaceMeta(idx int, meta any) { //gd:MeshDataTool.set_face_meta
	class(self).SetFaceMeta(gd.Int(idx), gd.NewVariant(meta))
}

/*
Returns the metadata associated with the given face.
*/
func (self Instance) GetFaceMeta(idx int) any { //gd:MeshDataTool.get_face_meta
	return any(class(self).GetFaceMeta(gd.Int(idx)).Interface())
}

/*
Calculates and returns the face normal of the given face.
*/
func (self Instance) GetFaceNormal(idx int) Vector3.XYZ { //gd:MeshDataTool.get_face_normal
	return Vector3.XYZ(class(self).GetFaceNormal(gd.Int(idx)))
}

/*
Sets the material to be used by newly-constructed [Mesh].
*/
func (self Instance) SetMaterial(material [1]gdclass.Material) { //gd:MeshDataTool.set_material
	class(self).SetMaterial(material)
}

/*
Returns the material assigned to the [Mesh].
*/
func (self Instance) GetMaterial() [1]gdclass.Material { //gd:MeshDataTool.get_material
	return [1]gdclass.Material(class(self).GetMaterial())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MeshDataTool

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshDataTool"))
	casted := Instance{*(*gdclass.MeshDataTool)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Clears all data currently in MeshDataTool.
*/
//go:nosplit
func (self class) Clear() { //gd:MeshDataTool.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Uses specified surface of given [Mesh] to populate data for MeshDataTool.
Requires [Mesh] with primitive type [constant Mesh.PRIMITIVE_TRIANGLES].
*/
//go:nosplit
func (self class) CreateFromSurface(mesh [1]gdclass.ArrayMesh, surface gd.Int) gd.Error { //gd:MeshDataTool.create_from_surface
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_create_from_surface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new surface to specified [Mesh] with edited data.
*/
//go:nosplit
func (self class) CommitToSurface(mesh [1]gdclass.ArrayMesh, compression_flags gd.Int) gd.Error { //gd:MeshDataTool.commit_to_surface
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, compression_flags)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_commit_to_surface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Mesh]'s format as a combination of the [enum Mesh.ArrayFormat] flags. For example, a mesh containing both vertices and normals would return a format of [code]3[/code] because [constant Mesh.ARRAY_FORMAT_VERTEX] is [code]1[/code] and [constant Mesh.ARRAY_FORMAT_NORMAL] is [code]2[/code].
*/
//go:nosplit
func (self class) GetFormat() gd.Int { //gd:MeshDataTool.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of vertices in [Mesh].
*/
//go:nosplit
func (self class) GetVertexCount() gd.Int { //gd:MeshDataTool.get_vertex_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of edges in this [Mesh].
*/
//go:nosplit
func (self class) GetEdgeCount() gd.Int { //gd:MeshDataTool.get_edge_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of faces in this [Mesh].
*/
//go:nosplit
func (self class) GetFaceCount() gd.Int { //gd:MeshDataTool.get_face_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the given vertex.
*/
//go:nosplit
func (self class) SetVertex(idx gd.Int, vertex gd.Vector3) { //gd:MeshDataTool.set_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the given vertex.
*/
//go:nosplit
func (self class) GetVertex(idx gd.Int) gd.Vector3 { //gd:MeshDataTool.get_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the normal of the given vertex.
*/
//go:nosplit
func (self class) SetVertexNormal(idx gd.Int, normal gd.Vector3) { //gd:MeshDataTool.set_vertex_normal
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, normal)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the normal of the given vertex.
*/
//go:nosplit
func (self class) GetVertexNormal(idx gd.Int) gd.Vector3 { //gd:MeshDataTool.get_vertex_normal
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tangent of the given vertex.
*/
//go:nosplit
func (self class) SetVertexTangent(idx gd.Int, tangent gd.Plane) { //gd:MeshDataTool.set_vertex_tangent
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, tangent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tangent of the given vertex.
*/
//go:nosplit
func (self class) GetVertexTangent(idx gd.Int) gd.Plane { //gd:MeshDataTool.get_vertex_tangent
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Plane](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the UV of the given vertex.
*/
//go:nosplit
func (self class) SetVertexUv(idx gd.Int, uv gd.Vector2) { //gd:MeshDataTool.set_vertex_uv
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, uv)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the UV of the given vertex.
*/
//go:nosplit
func (self class) GetVertexUv(idx gd.Int) gd.Vector2 { //gd:MeshDataTool.get_vertex_uv
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the UV2 of the given vertex.
*/
//go:nosplit
func (self class) SetVertexUv2(idx gd.Int, uv2 gd.Vector2) { //gd:MeshDataTool.set_vertex_uv2
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, uv2)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_uv2, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the UV2 of the given vertex.
*/
//go:nosplit
func (self class) GetVertexUv2(idx gd.Int) gd.Vector2 { //gd:MeshDataTool.get_vertex_uv2
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_uv2, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the color of the given vertex.
*/
//go:nosplit
func (self class) SetVertexColor(idx gd.Int, color gd.Color) { //gd:MeshDataTool.set_vertex_color
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the color of the given vertex.
*/
//go:nosplit
func (self class) GetVertexColor(idx gd.Int) gd.Color { //gd:MeshDataTool.get_vertex_color
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bones of the given vertex.
*/
//go:nosplit
func (self class) SetVertexBones(idx gd.Int, bones gd.PackedInt32Array) { //gd:MeshDataTool.set_vertex_bones
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(bones))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the bones of the given vertex.
*/
//go:nosplit
func (self class) GetVertexBones(idx gd.Int) gd.PackedInt32Array { //gd:MeshDataTool.get_vertex_bones
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the bone weights of the given vertex.
*/
//go:nosplit
func (self class) SetVertexWeights(idx gd.Int, weights gd.PackedFloat32Array) { //gd:MeshDataTool.set_vertex_weights
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(weights))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns bone weights of the given vertex.
*/
//go:nosplit
func (self class) GetVertexWeights(idx gd.Int) gd.PackedFloat32Array { //gd:MeshDataTool.get_vertex_weights
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the metadata associated with the given vertex.
*/
//go:nosplit
func (self class) SetVertexMeta(idx gd.Int, meta gd.Variant) { //gd:MeshDataTool.set_vertex_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the metadata associated with the given vertex.
*/
//go:nosplit
func (self class) GetVertexMeta(idx gd.Int) gd.Variant { //gd:MeshDataTool.get_vertex_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of edges that share the given vertex.
*/
//go:nosplit
func (self class) GetVertexEdges(idx gd.Int) gd.PackedInt32Array { //gd:MeshDataTool.get_vertex_edges
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_edges, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of faces that share the given vertex.
*/
//go:nosplit
func (self class) GetVertexFaces(idx gd.Int) gd.PackedInt32Array { //gd:MeshDataTool.get_vertex_faces
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns index of specified vertex connected to given edge.
Vertex argument can only be 0 or 1 because edges are comprised of two vertices.
*/
//go:nosplit
func (self class) GetEdgeVertex(idx gd.Int, vertex gd.Int) gd.Int { //gd:MeshDataTool.get_edge_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns array of faces that touch given edge.
*/
//go:nosplit
func (self class) GetEdgeFaces(idx gd.Int) gd.PackedInt32Array { //gd:MeshDataTool.get_edge_faces
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the metadata of the given edge.
*/
//go:nosplit
func (self class) SetEdgeMeta(idx gd.Int, meta gd.Variant) { //gd:MeshDataTool.set_edge_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_edge_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns meta information assigned to given edge.
*/
//go:nosplit
func (self class) GetEdgeMeta(idx gd.Int) gd.Variant { //gd:MeshDataTool.get_edge_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the specified vertex index of the given face.
Vertex argument must be either 0, 1, or 2 because faces contain three vertices.
[b]Example:[/b]
[codeblocks]
[gdscript]
var index = mesh_data_tool.get_face_vertex(0, 1) # Gets the index of the second vertex of the first face.
var position = mesh_data_tool.get_vertex(index)
var normal = mesh_data_tool.get_vertex_normal(index)
[/gdscript]
[csharp]
int index = meshDataTool.GetFaceVertex(0, 1); // Gets the index of the second vertex of the first face.
Vector3 position = meshDataTool.GetVertex(index);
Vector3 normal = meshDataTool.GetVertexNormal(index);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetFaceVertex(idx gd.Int, vertex gd.Int) gd.Int { //gd:MeshDataTool.get_face_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns specified edge associated with given face.
Edge argument must be either 0, 1, or 2 because a face only has three edges.
*/
//go:nosplit
func (self class) GetFaceEdge(idx gd.Int, edge gd.Int) gd.Int { //gd:MeshDataTool.get_face_edge
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, edge)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the metadata of the given face.
*/
//go:nosplit
func (self class) SetFaceMeta(idx gd.Int, meta gd.Variant) { //gd:MeshDataTool.set_face_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_face_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the metadata associated with the given face.
*/
//go:nosplit
func (self class) GetFaceMeta(idx gd.Int) gd.Variant { //gd:MeshDataTool.get_face_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Calculates and returns the face normal of the given face.
*/
//go:nosplit
func (self class) GetFaceNormal(idx gd.Int) gd.Vector3 { //gd:MeshDataTool.get_face_normal
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the material to be used by newly-constructed [Mesh].
*/
//go:nosplit
func (self class) SetMaterial(material [1]gdclass.Material) { //gd:MeshDataTool.set_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the material assigned to the [Mesh].
*/
//go:nosplit
func (self class) GetMaterial() [1]gdclass.Material { //gd:MeshDataTool.get_material
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsMeshDataTool() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMeshDataTool() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("MeshDataTool", func(ptr gd.Object) any {
		return [1]gdclass.MeshDataTool{*(*gdclass.MeshDataTool)(unsafe.Pointer(&ptr))}
	})
}

type Error = gd.Error //gd:Error

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
