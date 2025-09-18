/*
var query = PhysicsRayQueryParameters2D.create(global_position, global_position + Vector2(0, 100))
var collision = get_world_2d().direct_space_state.intersect_ray(query)
*/

package main

import (
	"graphics.gd/classdb/PhysicsRayQueryParameters2D"
	"graphics.gd/variant/Vector2"
)

func PhysicsRayQueryParameters2D_Create() {
	var query = PhysicsRayQueryParameters2D.Create(node2d.GlobalPosition(), Vector2.Add(node2d.GlobalPosition(), Vector2.XY{0, 100}), nil)
	var collision = node2d.AsCanvasItem().GetWorld2d().DirectSpaceState().IntersectRay(query)
	_ = collision
}
