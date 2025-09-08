package main

import (
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/RenderingServer"
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

type Racing struct {
	Node2D.Extension[Racing]
	//Police *Police
	Road *Road

	Player *Player
}

func (r *Racing) UnhandledInput(event InputEvent.Instance) {
	if event, ok := Object.As[InputEventKey.Instance](event); ok {
		if event.Keycode() == Input.KeyTab && event.AsInputEvent().IsPressed() {
			Debug = !Debug
		}
	}
}

func (game *Racing) Ready() {
	RenderingServer.SetDefaultClearColor(Color.W3C.DarkGreen)
}

func (game *Racing) Process(dt Float.X) {
	if game.Road.Travel(dt, Vector2.New(real(game.Player.pos), imag(game.Player.pos)), real(game.Player.speed), Angle.Radians(real(game.Player.angle))) {
		game.Player.speed = 0
	}
}

func main() {
	classdb.Register[Racing]()
	classdb.Register[Player]()
	classdb.Register[Road]()
	startup.Scene()
}
