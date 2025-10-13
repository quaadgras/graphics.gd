package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Marker2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

type State Enum.Int[struct {
	Idle,
	Follow State
}]

var States = Enum.Values[State]()

const (
	Mass           = 10
	ArriveDistance = 10
)

type Character struct {
	Marker2D.Extension[Character]

	speed          Float.X `range:"10,500,0.1,or_greater"`
	state          State
	velocity       Vector2.XY
	click_position Vector2.XY
	path           []Vector2.XY
	next_point     Vector2.XY

	TileMap *PathFindAStar `gd:"../TileMapLayer"`
}

func NewCharacter() *Character {
	return &Character{
		speed: 200,
	}
}

func (c *Character) Ready() { c.change_state(States.Idle) }
func (c *Character) Process(delta Float.X) {
	if c.state != States.Follow {
		return
	}
	if arrived_to_next_point := c.move_to(c.next_point); arrived_to_next_point {
		c.path = c.path[1:]
		if len(c.path) == 0 {
			c.change_state(States.Idle)
			return
		}
		c.next_point = c.path[0]
	}
}
func (c *Character) UnhandledInput(event InputEvent.Instance) {
	c.click_position = c.AsCanvasItem().GetGlobalMousePosition()
	if c.TileMap.IsPointWalkable(c.click_position) {
		if event.MoreArgs().IsActionPressed("teleport_to", false, true) {
			c.change_state(States.Idle)
			c.AsNode2D().SetGlobalPosition(Vector2.From(c.TileMap.RoundLocalPosition(Vector2i.From(c.click_position))))
		} else if event.IsActionPressed("move_to") {
			c.change_state(States.Follow)
		}
	}
}
func (c *Character) move_to(local_position Vector2.XY) bool {
	var desired_velocity = Vector2.MulX(Vector2.Normalized(Vector2.Sub(local_position, c.AsNode2D().Position())), c.speed)
	var steering = Vector2.Sub(desired_velocity, c.velocity)
	c.velocity = Vector2.Add(c.velocity, Vector2.DivX(steering, Mass))
	c.AsNode2D().SetPosition(Vector2.Add(c.AsNode2D().Position(), Vector2.MulX(c.velocity, c.AsNode().GetProcessDeltaTime())))
	c.AsNode2D().SetRotation(Vector2.AngleRadians(c.velocity))
	return Vector2.Distance(c.AsNode2D().Position(), local_position) < ArriveDistance
}
func (c *Character) change_state(new_state State) {
	switch new_state {
	case States.Idle:
		c.TileMap.ClearPath()
	case States.Follow:
		c.path = c.TileMap.FindPath(c.AsNode2D().Position(), c.click_position)
		if len(c.path) < 2 {
			new_state = States.Idle
			break
		}
		c.next_point = c.path[1]
	}
	c.state = new_state
}

func main() {
	classdb.Register[Character](NewCharacter)
	classdb.Register[PathFindAStar](NewPathFindAStar)
	startup.Scene()
}
