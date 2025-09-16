/*
func _enter_tree():
	# Depending on when the node is added to the tree,
	# prints either "true" or "false".
	print(Engine.is_in_physics_frame())

func _process(delta):
	print(Engine.is_in_physics_frame()) # Prints false

func _physics_process(delta):
	print(Engine.is_in_physics_frame()) # Prints true
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
)

type ExampleInPhysicsFrame struct {
	Node.Extension[ExampleInPhysicsFrame]
}

func (e *ExampleInPhysicsFrame) EnterTree() {
	// Depending on when the node is added to the tree,
	// prints either "true" or "false".
	fmt.Println(Engine.IsInPhysicsFrame())
}

func (e *ExampleInPhysicsFrame) Process(delta float64) {
	fmt.Println(Engine.IsInPhysicsFrame()) // Prints false
}

func (e *ExampleInPhysicsFrame) PhysicsProcess(delta float64) {
	fmt.Println(Engine.IsInPhysicsFrame()) // Prints true
}
