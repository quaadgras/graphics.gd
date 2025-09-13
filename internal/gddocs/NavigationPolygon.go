/*
[gdscript]
var new_navigation_mesh = NavigationPolygon.new()
var bounding_outline = PackedVector2Array([Vector2(0, 0), Vector2(0, 50), Vector2(50, 50), Vector2(50, 0)])
new_navigation_mesh.add_outline(bounding_outline)
NavigationServer2D.bake_from_source_geometry_data(new_navigation_mesh, NavigationMeshSourceGeometryData2D.new());
$NavigationRegion2D.navigation_polygon = new_navigation_mesh
[/gdscript]
[csharp]
var newNavigationMesh = new NavigationPolygon();
Vector2[] boundingOutline = [new Vector2(0, 0), new Vector2(0, 50), new Vector2(50, 50), new Vector2(50, 0)];
newNavigationMesh.AddOutline(boundingOutline);
NavigationServer2D.BakeFromSourceGeometryData(newNavigationMesh, new NavigationMeshSourceGeometryData2D());
GetNode<NavigationRegion2D>("NavigationRegion2D").NavigationPolygon = newNavigationMesh;
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/NavigationMeshSourceGeometryData2D"
	"graphics.gd/classdb/NavigationPolygon"
	"graphics.gd/classdb/NavigationRegion2D"
	"graphics.gd/classdb/NavigationServer2D"
	"graphics.gd/variant/Vector2"
)

func ExampleNavigationPolygon(region NavigationRegion2D.Instance) {
	var new_navigation_mesh = NavigationPolygon.New()
	var bounding_outline = []Vector2.XY{{0, 0}, {0, 50}, {50, 50}, {50, 0}}
	new_navigation_mesh.AddOutline(bounding_outline)
	NavigationServer2D.BakeFromSourceGeometryData(new_navigation_mesh, NavigationMeshSourceGeometryData2D.New(), nil)
	region.SetNavigationPolygon(new_navigation_mesh)
}
