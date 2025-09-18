/*
[gdscript]
var target = get_collider() # A CollisionObject3D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject3D)GetCollider(); // A CollisionObject3D.
var shapeId = GetColliderShape(); // The shape index in the collider.
var ownerId = target.ShapeFindOwner(shapeId); // The owner ID in the collider.
var shape = target.ShapeOwnerGetOwner(ownerId);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/CollisionObject3D"
	"graphics.gd/classdb/RayCast3D"
	"graphics.gd/variant/Object"
)

var rayCast3D RayCast3D.Instance

func RayCast3D_GetColliderShape() {
	var target = Object.To[CollisionObject3D.Instance](rayCast3D.GetCollider()) // A CollisionObject3D.
	var shapeID int = rayCast3D.GetColliderShape()
	var ownerID int = target.ShapeFindOwner(shapeID)
	var shape = target.ShapeOwnerGetOwner(ownerID)
	_ = shape
}
