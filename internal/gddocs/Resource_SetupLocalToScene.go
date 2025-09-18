/*
extends Resource

var damage = 0

func _setup_local_to_scene():
	damage = randi_range(10, 40)
*/

package main

import (
	"graphics.gd/classdb/Resource"
	"graphics.gd/variant/Int"
)

type MyResource struct {
	Resource.Extension[MyResource]

	damage int
}

func (r *MyResource) SetupLocalToScene() {
	r.damage = Int.RandomBetween(10, 40)
}
