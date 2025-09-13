/*
[gdscript]
var mesh = ImmediateMesh.new()
mesh.surface_begin(Mesh.PRIMITIVE_TRIANGLES)
mesh.surface_add_vertex(Vector3.LEFT)
mesh.surface_add_vertex(Vector3.FORWARD)
mesh.surface_add_vertex(Vector3.ZERO)
mesh.surface_end()
[/gdscript]
[csharp]
var mesh = new ImmediateMesh();
mesh.SurfaceBegin(Mesh.PrimitiveType.Triangles);
mesh.SurfaceAddVertex(Vector3.Left);
mesh.SurfaceAddVertex(Vector3.Forward);
mesh.SurfaceAddVertex(Vector3.Zero);
mesh.SurfaceEnd();
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/ImmediateMesh"
	"graphics.gd/classdb/Mesh"
	"graphics.gd/variant/Vector3"
)

func ExampleImmediateMesh() {
	var mesh = ImmediateMesh.New()
	mesh.SurfaceBegin(Mesh.PrimitiveTriangles)
	mesh.SurfaceAddVertex(Vector3.Left)
	mesh.SurfaceAddVertex(Vector3.Forward)
	mesh.SurfaceAddVertex(Vector3.Zero)
	mesh.SurfaceEnd()
}
