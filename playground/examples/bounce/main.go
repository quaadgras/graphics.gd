// A class written in Go: a ball that the engine moves and draws every frame.
package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

type Ball struct {
	Node2D.Extension[Ball]

	Radius   Float.X
	Velocity Vector2.XY
	Bounces  int
}

// Process is called by the engine every frame.
func (b *Ball) Process(delta Float.X) {
	size := DisplayServer.WindowGetSize(0)
	pos := Vector2.Add(b.AsNode2D().Position(), Vector2.MulX(b.Velocity, delta))
	if pos.X < b.Radius || pos.X > Float.X(size.X)-b.Radius {
		b.Velocity.X = -b.Velocity.X
		b.bounced()
	}
	if pos.Y < b.Radius || pos.Y > Float.X(size.Y)-b.Radius {
		b.Velocity.Y = -b.Velocity.Y
		b.bounced()
	}
	pos.X = min(max(pos.X, b.Radius), Float.X(size.X)-b.Radius)
	pos.Y = min(max(pos.Y, b.Radius), Float.X(size.Y)-b.Radius)
	b.AsNode2D().SetPosition(pos)
}

func (b *Ball) bounced() {
	b.Bounces++
	fmt.Println("bounce", b.Bounces)
	b.AsCanvasItem().QueueRedraw()
}

// Draw is called by the engine when the ball has to be drawn again.
func (b *Ball) Draw() {
	hue := Float.X(b.Bounces%12) / 12
	b.AsCanvasItem().DrawCircle(Vector2.Zero, b.Radius, Color.HSV(hue, 0.7, 1))
}

func main() {
	classdb.Register[Ball]()
	startup.LoadingScene()
	ball := new(Ball)
	ball.Radius = 40
	ball.Velocity = Vector2.New(260, 190)
	ball.AsNode2D().SetPosition(Vector2.New(120, 90))
	SceneTree.Add(ball)
	startup.Scene()
}
