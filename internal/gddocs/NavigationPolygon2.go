/*
[gdscript]
var new_navigation_mesh = NavigationPolygon.new()
var new_vertices = PackedVector2Array([Vector2(0, 0), Vector2(0, 50), Vector2(50, 50), Vector2(50, 0)])
new_navigation_mesh.vertices = new_vertices
var new_polygon_indices = PackedInt32Array([0, 1, 2, 3])
new_navigation_mesh.add_polygon(new_polygon_indices)
$NavigationRegion2D.navigation_polygon = new_navigation_mesh
[/gdscript]
[csharp]
var newNavigationMesh = new NavigationPolygon();
Vector2[] newVertices = [new Vector2(0, 0), new Vector2(0, 50), new Vector2(50, 50), new Vector2(50, 0)];
newNavigationMesh.Vertices = newVertices;
int[] newPolygonIndices = [0, 1, 2, 3];
newNavigationMesh.AddPolygon(newPolygonIndices);
GetNode<NavigationRegion2D>("NavigationRegion2D").NavigationPolygon = newNavigationMesh;
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/NavigationPolygon"
	"graphics.gd/classdb/NavigationRegion2D"
	"graphics.gd/variant/Vector2"
)

func ExampleNavigationPolygon2(region NavigationRegion2D.Instance) {
	var new_navigation_mesh = NavigationPolygon.New()
	var new_vertices = []Vector2.XY{{0, 0}, {0, 50}, {50, 50}, {50, 0}}
	new_navigation_mesh.SetVertices(new_vertices)
	var new_polygon_indices = []int32{0, 1, 2, 3}
	new_navigation_mesh.AddPolygon(new_polygon_indices)
	region.SetNavigationPolygon(new_navigation_mesh.AsNavigationPolygon())
}
