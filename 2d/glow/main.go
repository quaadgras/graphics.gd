package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

const (
	CaveLimit = 1000
)

var GlowMap = Resource.Load[Texture.Instance]("res://glow_map.webp")

type BeachCave struct {
	Node2D.Extension[BeachCave]

	Cave             Node2D.Instance
	WorldEnvironment WorldEnvironment.Instance
}

func (b *BeachCave) UnhandledInput(event InputEvent.Instance) {
	if event, ok := Object.As[InputEventMouseMotion.Instance](event); ok && event.AsInputEventMouse().ButtonMask() > 0 {
		position := b.Cave.Position()
		position.X = Float.Clamp(position.X+event.Relative().X*2, -CaveLimit, 0)
		b.Cave.SetPosition(position)
	}
	if event.IsActionPressed("toggle_glow_map") {
		env := b.WorldEnvironment.Environment()
		if env.GlowMap() != Texture.Nil {
			env.SetGlowMap(Texture.Nil)
			env.SetGlowIntensity(0.8)
		} else {
			env.SetGlowMap(GlowMap)
			env.SetGlowIntensity(1.6)
		}
	}
}

func main() {
	classdb.Register[BeachCave]()
	startup.Scene()
}
