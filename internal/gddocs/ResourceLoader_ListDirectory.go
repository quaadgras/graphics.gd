/*
# Prints ["extra_data/", "model.gltf", "model.tscn", "model_slime.png"]
print(ResourceLoader.list_directory("res://assets/enemies/slime"))
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/ResourceLoader"
)

func ResourceLoader_ListDirectory() {
	// Prints ["extra_data/", "model.gltf", "model.tscn", "model_slime.png"]
	fmt.Println(ResourceLoader.ListDirectory("res://assets/enemies/slime"))
}
