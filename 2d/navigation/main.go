package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/NavigationAgent2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

type Character struct {
	CharacterBody2D.Extension[Character]

	MovementSpeed Float.X

	NavigationAgent NavigationAgent2D.Instance
}

func NewCharacter() *Character {
	return &Character{
		MovementSpeed: 200,
	}
}

func (c *Character) Ready() {
	// These values need to be adjusted for the actor's speed
	// and the navigation layout.
	c.NavigationAgent.SetPathDesiredDistance(2)
	c.NavigationAgent.SetTargetDesiredDistance(2)
	c.NavigationAgent.SetDebugEnabled(true)
}

// The "click" event is a custom input action defined in
// Project > Project Settings > Input Map tab.
func (c *Character) UnhandledInput(event InputEvent.Instance) {
	if !event.IsActionPressed("click") {
		return
	}
	c.SetMovementTarget(c.AsCanvasItem().GetGlobalMousePosition())
}

func (c *Character) SetMovementTarget(movement_target Vector2.XY) {
	c.NavigationAgent.SetTargetPosition(movement_target)
}

func (c *Character) PhysicsProcess(delta Float.X) {
	if c.NavigationAgent.IsNavigationFinished() {
		return
	}
	var current_agent_position = c.AsNode2D().GlobalPosition()
	var next_path_position = c.NavigationAgent.GetNextPathPosition()
	c.AsCharacterBody2D().SetVelocity(Vector2.MulX(Vector2.Direction(current_agent_position, next_path_position), c.MovementSpeed))
	c.AsCharacterBody2D().MoveAndSlide()
}

func main() {
	classdb.Register[Character](NewCharacter)
	startup.Scene()
}
