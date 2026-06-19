/*
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
*/

package main

import "graphics.gd/classdb/MeshDataTool"

func ExampleMeshDataToolGetFaceVertex(meshDataTool MeshDataTool.Instance) {
	var index = meshDataTool.GetFaceVertex(0, 1) // Gets the index of the second vertex of the first face.
	var position = meshDataTool.GetVertex(index)
	var normal = meshDataTool.GetVertexNormal(index)
	_, _ = position, normal
}
