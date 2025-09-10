package main

import (
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

type Player struct {
	Node2D.Extension[Player]

	texture Texture2D.Instance

	pos, angle, speed complex64

	deathVector complex64
}

func (p *Player) Ready() {
	p.texture = Resource.Load[Texture2D.Instance]("res://car.png")
	p.AsCanvasItem().SetZIndex(-1)
}

func (p *Player) Process(delta Float.X) {
	Object.Use(p.texture)

	dt := complex(delta, 0)

	p.AsCanvasItem().QueueRedraw()

	if GameOver {
		if real(p.speed) > 100*6 {
			p.pos += p.deathVector * p.speed / 2 * dt
			p.angle += (2 * π) * p.speed * dt / 500
		}
		return
	}

	if Input.IsKeyPressed(Input.KeyW) || Input.IsKeyPressed(Input.KeyUp) {
		p.speed += 200 * dt
	}
	if Input.IsKeyPressed(Input.KeyS) || Input.IsKeyPressed(Input.KeyDown) || Input.IsKeyPressed(Input.KeySpace) {
		p.speed -= 500 * dt
		if real(p.speed) < 0 {
			p.speed = 0
		}
	}
	if p.speed == 0 {
		return
	}
	p.pos += complex(Angle.Sin(Angle.Radians(real(p.angle))), 0) * p.speed * dt

	if Input.IsKeyPressed(Input.KeyA) || Input.IsKeyPressed(Input.KeyLeft) {
		p.angle -= (π / 2 * (1 + p.speed/400)) * dt
		if real(p.angle) < -π/2 {
			p.angle = -π / 2
		}
	} else if Input.IsKeyPressed(Input.KeyD) || Input.IsKeyPressed(Input.KeyRight) {
		p.angle += (π / 2 * (1 + p.speed/400)) * dt
		if real(p.angle) > π/2 {
			p.angle = π / 2
		}
	}
}

func (p *Player) Draw() {
	game_size := DisplayServer.WindowGetSize(0)
	center := Vector2.New(game_size.X/2, game_size.Y/2)
	if GameOver {
		center = Vector2.Add(center, Vector2.New(real(p.pos), imag(p.pos)))
	}
	CanvasItem.Expanded(p.AsCanvasItem()).DrawSetTransform(center, Angle.Radians(real(p.angle)), Vector2.One)
	p.AsCanvasItem().DrawTexture(p.texture, Vector2.New(-p.texture.GetWidth()/2, -p.texture.GetHeight()/2))
	if Debug {
		p.AsCanvasItem().DrawCircle(Vector2.Zero, 32, Color.Bytes(0, 0, 100, 50))
	}
}
