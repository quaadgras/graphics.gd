/*
[gdscript]
"vertices" : PackedFloat32Array
"carve" : bool
[/gdscript]
*/

package main

func NavigationMeshSourceGeometryData2D_SetProjectedObstructions() {
	type ProjectedObstruction struct {
		Vertices []float32 `gd:"vertices"`
		Carve    bool      `gd:"carve"`
	}
}
