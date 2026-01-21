---
title: Canvas
slug: guide/rendering/canvas
sidebar:
  order: 24
---

A high-level immediate-mode 2D rendering option is available by implementing the `Draw()` method on a
[Node2D.Extension](https://pkg.go.dev/graphics.gd/classdb/Node2D#Extension).

You can structure the entire project with a single [Node2D](https://pkg.go.dev/graphics.gd/classdb/Node2D) root
added to the default scene, and pass down the [CanvasItem](https://pkg.go.dev/graphics.gd/classdb/CanvasItem)
as needed for rendering everything, along with implementing `Ready()` and `Process(Float.X)` this will give you
a familiar development experience to something like [Love2D](https://love2d.org/) or [Ebiten](https://ebiten.org/).

```go
package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
	"graphics.gd/startup"
)

type MyProject struct {
	Node2D.Extension[MyProject]
}

func (m *MyProject) Ready() {
	// Initialization code here
}

func (m *MyProject) Process(delta Float.X) {
	// Update code here
}

func (m *MyProject) Draw() {
	canvas := m.AsCanvasItem() // pass this to all your Draw functions
	canvas.DrawCircle(Vector2.New(100, 100), 50, Color.X11.Red)
}

func main() {
	startup.LoadingScene()
	SceneTree.Add(new(MyProject))
	startup.Scene()
}

```
