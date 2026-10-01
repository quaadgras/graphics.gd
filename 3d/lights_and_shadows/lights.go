package main

import (
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OmniLight3D"
	"graphics.gd/classdb/SpotLight3D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
)

// DayNightCycle slowly rotates the sun.
type DayNightCycle struct {
	DirectionalLight3D.Extension[DayNightCycle]
}

func (d *DayNightCycle) Process(delta Float.X) {
	d.AsNode3D().RotateObjectLocal(Vector3.Right, Angle.Radians(0.025*delta))
}

// spin moves a light around in a circle while turning it.
func spin(node Node3D.Instance, increment *Float.X, delta Float.X) {
	var position = node.Position()
	position.X = Angle.Sin(Angle.Radians(*increment))
	position.Z = Angle.Cos(Angle.Radians(*increment))
	node.SetPosition(position)
	// Avoid precision issues over time by rolling over every full turn.
	var rotation = node.Rotation()
	rotation.Y = Angle.Radians(Float.Mod(*increment, Float.X(Angle.Tau)))
	node.SetRotation(rotation)

	*increment += delta
}

type SpinOmniLight struct {
	OmniLight3D.Extension[SpinOmniLight]

	increment Float.X
}

func (s *SpinOmniLight) Process(delta Float.X) { spin(s.AsNode3D(), &s.increment, delta) }

type SpinSpotLight struct {
	SpotLight3D.Extension[SpinSpotLight]

	increment Float.X
}

func (s *SpinSpotLight) Process(delta Float.X) { spin(s.AsNode3D(), &s.increment, delta) }
