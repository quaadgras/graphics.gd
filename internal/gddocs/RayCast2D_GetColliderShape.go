/*
[gdscript]
var target = get_collider() # A CollisionObject2D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject2D)GetCollider(); // A CollisionObject2D.
var shapeId = GetColliderShape(); // The shape index in the collider.
var ownerId = target.ShapeFindOwner(shapeId); // The owner ID in the collider.
var shape = target.ShapeOwnerGetOwner(ownerId);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/CollisionObject2D"
	"graphics.gd/classdb/RayCast2D"
	"graphics.gd/variant/Object"
)

var rayCast2D RayCast2D.Instance

func RayCast2D_GetColliderShape() {
	var target = Object.To[CollisionObject2D.Instance](rayCast2D.GetCollider()) // A CollisionObject2D.
	var shapeID int = rayCast2D.GetColliderShape()
	var ownerID int = target.ShapeFindOwner(shapeID) // The owner ID in the collider.
	var shape = target.ShapeOwnerGetOwner(ownerID)
	_ = shape
}
