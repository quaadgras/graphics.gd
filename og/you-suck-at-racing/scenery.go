package main

import (
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

var Scene []Scenery

type Scenery struct {
	Pos   Vector2.XY
	Image Texture2D.Instance
}

func AddToScene(resize_offset, Pos Vector2.XY, Image Texture2D.Instance) {
	if Pos.Y+resize_offset.Y > 0 {
		return
	}
	for i, scenery := range Scene {
		if scenery.Image == Texture2D.Nil {
			Scene[i] = Scenery{Pos: Pos, Image: Image}
			return
		}
	}
	Scene = append(Scene, Scenery{Pos: Pos, Image: Image})
}

func UpdateScene(dt, speed Float.X, angle Angle.Radians) {
	var game_size = DisplayServer.WindowGetSize(0)
	speed = Angle.Cos(Angle.Radians(angle)) * speed * dt
	for i, scenery := range Scene {
		if scenery.Image == Texture2D.Nil {
			continue
		}
		Object.Use(scenery.Image)
		if i < len(Scene) {
			Scene[i].Pos.Y += speed
			if Scene[i].Pos.Y-scenery.Image.GetSize().Y > Float.X(game_size.Y) {
				a := Scene
				a[i].Image = Texture2D.Nil
				Scene = a
			}
		}
	}
}

func DrawScene(canvas CanvasItem.Instance, offset Vector2.XY) {
	var count int
	for i := len(Scene) - 1; i >= 0; i-- {
		Scene[i].Draw(canvas, offset)
		count++
	}
}

func (scenery *Scenery) Draw(canvas CanvasItem.Instance, offset Vector2.XY) {
	if scenery.Image == Texture2D.Nil {
		return
	}
	size := scenery.Image.GetSize()
	canvas.DrawTexture(scenery.Image, Vector2.Add(Vector2.Add(scenery.Pos, Vector2.New(-size.X/2, -size.Y)), offset))
}
