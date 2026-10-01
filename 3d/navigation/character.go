package main

import (
	"graphics.gd/classdb/Marker3D"
	"graphics.gd/classdb/NavigationAgent3D"
	"graphics.gd/classdb/NavigationServer3D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
)

type Character struct {
	Marker3D.Extension[Character]

	CharacterSpeed Float.X
	ShowPath       bool

	NavAgent NavigationAgent3D.Instance `gd:"NavigationAgent3D"`

	navPathLine *Line3D
}

func NewCharacter() *Character {
	return &Character{CharacterSpeed: 10, ShowPath: true}
}

func (c *Character) Ready() {
	c.navPathLine = new(Line3D)
	c.AsNode().AddChild(c.navPathLine.AsNode())
	c.navPathLine.AsNode3D().SetTopLevel(true)
}

func (c *Character) PhysicsProcess(delta Float.X) {
	if c.NavAgent.IsNavigationFinished() {
		return
	}
	node := c.AsNode3D()
	var nextPosition = c.NavAgent.GetNextPathPosition()
	var offset = Vector3.Sub(nextPosition, node.GlobalPosition())
	node.SetGlobalPosition(Vector3.Move(node.GlobalPosition(), nextPosition, delta*c.CharacterSpeed))

	// Make the robot look at the direction we're traveling.
	// Clamp Y to 0 so the robot only looks left and right, not up/down.
	offset.Y = 0
	if !Vector3.IsApproximatelyZero(offset) {
		node.LookAt(Vector3.Add(node.GlobalPosition(), offset))
	}
}

func (c *Character) SetTargetPosition(targetPosition Vector3.XYZ) {
	c.NavAgent.SetTargetPosition(targetPosition)
	// Get a full navigation path with the NavigationServer API.
	if c.ShowPath {
		var startPosition = c.AsNode3D().GlobalTransform().Origin
		const optimize = true
		var navigationMap = c.AsNode3D().GetWorld3d().NavigationMap()
		var path = NavigationServer3D.MapGetPath(
			navigationMap,
			startPosition,
			targetPosition,
			optimize,
		)
		c.navPathLine.DrawPath(path)
	}
}
