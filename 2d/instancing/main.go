package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

var BallScene = Resource.Load[PackedScene.Is[Node2D.Instance]]("res://ball.tscn")

type BallFactory struct {
	Node2D.Extension[BallFactory]
}

func (b *BallFactory) UnhandledInput(event InputEvent.Instance) {
	if event.IsEcho() {
		return
	}
	if event, ok := Object.As[InputEventMouseButton.Instance](event); ok && event.AsInputEvent().IsPressed() {
		if event.ButtonIndex() == Input.MouseButtonLeft {
			b.Spawn(event.AsInputEventMouse().GlobalPosition())
		}
	}
}

func (b *BallFactory) Spawn(position Vector2.XY) {
	var instance = BallScene.Instantiate()
	instance.SetGlobalPosition(position)
	b.AsNode().AddChild(instance.AsNode())
}

func main() {
	classdb.Register[BallFactory]()
	startup.Scene()
}
