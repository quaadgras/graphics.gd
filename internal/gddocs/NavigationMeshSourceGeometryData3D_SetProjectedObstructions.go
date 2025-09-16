/*
[gdscript]
"vertices" : PackedFloat32Array
"elevation" : float
"height" : float
"carve" : bool
[/gdscript]
*/

package main

import "graphics.gd/classdb/NavigationMeshSourceGeometryData3D"

func NavigationMeshSourceGeometryData3D_SetProjectedObstructions() {
	type ProjectedObstruction3D struct {
		Vertices  []float32 `gd:"vertices"`
		Elevation float32   `gd:"elevation"`
		Height    float32   `gd:"height"`
		Carve     bool      `gd:"carve"`
	}
	_ = 0
	_ = NavigationMeshSourceGeometryData3D.ProjectedObstruction3D(ProjectedObstruction3D{})
}
