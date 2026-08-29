package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/RigidBody3D"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
)

var CubeRigidBody = Resource.Load[PackedScene.Is[RigidBody3D.Instance]]("res://cube_rigidbody.tscn")

type Level struct {
	Node3D.Extension[Level]
}

// OnSpawnTimerTimeout randomly spawns Rigidbody cubes.
func (l *Level) OnSpawnTimerTimeout() {
	var new_rb = CubeRigidBody.Instantiate()
	var position = new_rb.AsNode3D().Position()
	position.Y = 15
	position.X = Float.RandomBetween(-5, 5)
	position.Z = Float.RandomBetween(-5, 5)
	new_rb.AsNode3D().SetPosition(position)
	l.AsNode().AddChild(new_rb.AsNode())
}

func main() {
	classdb.Register[Level]()
	classdb.Register[Cubio]()
	classdb.Register[FollowCamera](NewFollowCamera)
	startup.Scene()
}
