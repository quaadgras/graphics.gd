/*
[gdscript]
for i in get_slide_collision_count():
	var collision = get_slide_collision(i)
	print("Collided with: ", collision.get_collider().name)
[/gdscript]
[csharp]
for (int i = 0; i < GetSlideCollisionCount(); i++)
{
	KinematicCollision2D collision = GetSlideCollision(i);
	GD.Print("Collided with: ", (collision.GetCollider() as Node).Name);
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Object"
)

var characterBody2D CharacterBody2D.Instance

func CharacterBody2D_GetSlideCollision() {
	for i := range characterBody2D.GetSlideCollisionCount() {
		var collision = characterBody2D.GetSlideCollision(i)
		fmt.Println("Collided with: ", Object.To[Node.Instance](collision.GetCollider()).Name())
	}
}
