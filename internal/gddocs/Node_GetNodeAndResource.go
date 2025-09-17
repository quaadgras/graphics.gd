/*
[gdscript]
var a = get_node_and_resource("Area2D/Sprite2D")
print(a[0].name) # Prints Sprite2D
print(a[1])      # Prints <null>
print(a[2])      # Prints ^""

var b = get_node_and_resource("Area2D/Sprite2D:texture:atlas")
print(b[0].name)        # Prints Sprite2D
print(b[1].get_class()) # Prints AtlasTexture
print(b[2])             # Prints ^""

var c = get_node_and_resource("Area2D/Sprite2D:texture:atlas:region")
print(c[0].name)        # Prints Sprite2D
print(c[1].get_class()) # Prints AtlasTexture
print(c[2])             # Prints ^":region"
[/gdscript]
[csharp]
var a = GetNodeAndResource(NodePath("Area2D/Sprite2D"));
GD.Print(a[0].Name); // Prints Sprite2D
GD.Print(a[1]);      // Prints <null>
GD.Print(a[2]);      // Prints ^"

var b = GetNodeAndResource(NodePath("Area2D/Sprite2D:texture:atlas"));
GD.Print(b[0].name);        // Prints Sprite2D
GD.Print(b[1].get_class()); // Prints AtlasTexture
GD.Print(b[2]);             // Prints ^""

var c = GetNodeAndResource(NodePath("Area2D/Sprite2D:texture:atlas:region"));
GD.Print(c[0].name);        // Prints Sprite2D
GD.Print(c[1].get_class()); // Prints AtlasTexture
GD.Print(c[2]);             // Prints ^":region"
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/variant/Object"
)

func Node_GetNodeAndResource() {
	node, res, path := node.GetNodeAndResource("Area2D/Sprite2D")
	fmt.Println(node.Name()) // Prints Sprite2D
	fmt.Println(res)         // Prints <null>
	fmt.Println(path)        // Prints ^""

	node, res, path = node.GetNodeAndResource("Area2D/Sprite2D:texture:atlas")
	fmt.Println(node.Name())                                 // Prints Sprite2D
	fmt.Println(Object.Instance(res.AsObject()).ClassName()) // Prints AtlasTexture
	fmt.Println(path)                                        // Prints ^""

	node, res, path = node.GetNodeAndResource("Area2D/Sprite2D:texture:atlas:region")
	fmt.Println(node.Name())                                 // Prints Sprite2D
	fmt.Println(Object.Instance(res.AsObject()).ClassName()) // Prints AtlasTexture
	fmt.Println(path)                                        // Prints ^":region"
}
