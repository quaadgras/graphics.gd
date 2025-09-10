package main

import (
	"fmt"
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/ThemeDB"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

const (
	π = math.Pi
)

var Debug bool
var GameOver bool

type Racing struct {
	Node2D.Extension[Racing]
	//Police *Police
	Road *Road

	Player *Player

	GameOver bool

	last_pos Vector2.XY

	resize_offset Vector2.XY
}

func (r *Racing) UnhandledInput(event InputEvent.Instance) {
	if event, ok := Object.As[InputEventKey.Instance](event); ok {
		if event.Keycode() == Input.KeyTab && event.AsInputEvent().IsPressed() {
			Debug = !Debug
		}
		if event.Keycode() == Input.KeyR && event.AsInputEvent().IsPressed() && GameOver {
			r.AsNode().QueueFree()
			GameOver = false
			Scene = Scene[:0:cap(Scene)]
			SceneTree.Add(new(Racing))
		}
	}
}

func (game *Racing) Ready() {
	RenderingServer.SetDefaultClearColor(Color.W3C.DarkGreen)
	first_size := DisplayServer.WindowGetSize(0)
	Viewport.Get(game.AsNode()).OnSizeChanged(func() {
		new_size := DisplayServer.WindowGetSize(0)
		game.resize_offset.X = -Float.X(first_size.X-new_size.X) / 2
		game.resize_offset.Y = -Float.X(first_size.Y-new_size.Y) / 2
	})
}

func (game *Racing) Process(dt Float.X) {
	game.AsCanvasItem().QueueRedraw()
	if game.Road.Travel(dt, Vector2.New(real(game.Player.pos), imag(game.Player.pos)), real(game.Player.speed), Angle.Radians(real(game.Player.angle)), game.resize_offset) {
		spin := Angle.Radians(real(game.Player.angle)) - π/2
		game.Player.deathVector = complex(Angle.Cos(spin), Angle.Sin(spin))
		game.last_pos = Vector2.New(real(game.Player.pos), imag(game.Player.pos))
		game.Player.pos = 0
		GameOver = true
	}
	if !GameOver {
		UpdateScene(dt, real(game.Player.speed), Angle.Radians(real(game.Player.angle)))
	}
}

func (game *Racing) Draw() {
	game_size := DisplayServer.WindowGetSize(0)
	center := Vector2.New(game_size.X/2, game_size.Y/2)

	var offset = Vector2.New(-real(game.Player.pos), 0)
	if GameOver {
		offset = Vector2.New(-game.last_pos.X, 0)
	}
	DrawScene(game.AsCanvasItem(), Vector2.Add(offset, game.resize_offset))

	font := ThemeDB.FallbackFont()
	game.AsCanvasItem().DrawString(font, Vector2.New(5, font.GetHeight()/2+5), fmt.Sprintf("Kph: %d", int(Float.Round(real(game.Player.speed)/6))))

	if GameOver {
		game.AsCanvasItem().DrawString(font, center, "YOU SUCK AT RACING")
		game.AsCanvasItem().DrawString(font, Vector2.New(Float.X(game_size.X/2)-font.GetStringSize("(Press R to restart!)").X/2, font.GetHeight()/2+5), "(Press R to restart!)")
	} else {
		game.AsCanvasItem().DrawString(font, Vector2.New(Float.X(game_size.X/2)-font.GetStringSize("(Press Space to brake and Tab to debug!)").X/2, font.GetHeight()/2+5), "(Press Space to brake and Tab to debug!)")
	}
}

func main() {
	classdb.Register[Racing]()
	classdb.Register[Player]()
	classdb.Register[Road]()
	startup.Scene()
}
