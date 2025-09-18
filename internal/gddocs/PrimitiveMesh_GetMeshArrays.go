/*
[gdscript]
var c = CylinderMesh.new()
var arr_mesh = ArrayMesh.new()
arr_mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, c.get_mesh_arrays())
[/gdscript]
[csharp]
var c = new CylinderMesh();
var arrMesh = new ArrayMesh();
arrMesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, c.GetMeshArrays());
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/ArrayMesh"
	"graphics.gd/classdb/CylinderMesh"
	"graphics.gd/classdb/Mesh"
)

func PrimitiveMesh_GetMeshArrays() {
	var c = CylinderMesh.New()
	var arrMesh = ArrayMesh.New()
	arrMesh.AddSurfaceFromArrays(Mesh.PrimitiveTriangles, c.AsPrimitiveMesh().GetMeshArrays())
}
