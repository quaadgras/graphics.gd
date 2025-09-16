/*
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
*/

package main

import (
	"graphics.gd/classdb/ArrayMesh"
	"graphics.gd/classdb/BoxMesh"
	"graphics.gd/classdb/Mesh"
	"graphics.gd/classdb/MeshDataTool"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Vector3"
)

func ExampleMeshDataTool(parent Node.Instance) {
	var mesh = ArrayMesh.New()
	mesh.AddSurfaceFromArrays(Mesh.PrimitiveTriangles, BoxMesh.New().AsPrimitiveMesh().GetMeshArrays())
	var mdt = MeshDataTool.New()
	mdt.CreateFromSurface(mesh, 0)
	for i := 0; i < mdt.GetVertexCount(); i++ {
		var vertex = mdt.GetVertex(i)
		// In this example we extend the mesh by one unit, which results in separated faces as it is flat shaded.
		vertex = Vector3.Add(vertex, mdt.GetVertexNormal(i))
		// Save your change.
		mdt.SetVertex(i, vertex)
	}
	mesh.ClearSurfaces()
	mdt.CommitToSurface(mesh)
	var mi = MeshInstance3D.New()
	mi.SetMesh(mesh.AsMesh())
	parent.AddChild(mi.AsNode())
}
