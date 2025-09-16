/*
[gdscript]
"vertices" : PackedFloat32Array
"carve" : bool
[/gdscript]
*/

package main

import "graphics.gd/classdb/NavigationMeshSourceGeometryData2D"

func NavigationMeshSourceGeometryData2D_SetProjectedObstructions() {
	type ProjectedObstruction2D struct {
		Vertices []float32 `gd:"vertices"`
		Carve    bool      `gd:"carve"`
	}
	_ = 0
	_ = NavigationMeshSourceGeometryData2D.ProjectedObstruction2D(ProjectedObstruction2D{})
}
