/*
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
Vector3[] vertices =
[
    new Vector3(0, 1, 0),
    new Vector3(1, 0, 0),
    new Vector3(0, 0, 1),
];

// Initialize the ArrayMesh.
var arrMesh = new ArrayMesh();
Godot.Collections.Array arrays = [];
arrays.Resize((int)Mesh.ArrayType.Max);
arrays[(int)Mesh.ArrayType.Vertex] = vertices;

// Create the Mesh.
arrMesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, arrays);
var m = new MeshInstance3D();
m.Mesh = arrMesh;
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/ArrayMesh"
	"graphics.gd/classdb/Mesh"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/variant/Vector3"
)

func ExampleArrayMesh() {
	var vertices = []Vector3.XYZ{
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, 1},
	}
	// Initialize the ArrayMesh.
	var arrMesh = ArrayMesh.New()
	var arrays = make([]any, Mesh.ArrayMax)
	arrays[Mesh.ArrayVertex] = vertices

	// Create the Mesh.
	arrMesh.AddSurfaceFromArrays(Mesh.PrimitiveTriangles, arrays)
	var m = MeshInstance3D.New()
	m.SetMesh(arrMesh.AsMesh())
}
