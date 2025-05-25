// Package Euler provides euler angle types for representing 3D rotations.
package Euler

import (
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
)

type Radians struct {
	X Angle.Radians
	Y Angle.Radians
	Z Angle.Radians
}

type vector3 = struct {
	X Float.X // The vector's X component.
	Y Float.X // The vector's Y component.
	Z Float.X // The vector's Z component.
}

func (r Radians) Vector3() vector3 {
	return vector3{
		X: Float.X(r.X),
		Y: Float.X(r.Y),
		Z: Float.X(r.Z),
	}
}

func (r Radians) Degrees() Degrees {
	return Degrees{
		X: Angle.InDegrees(r.X),
		Y: Angle.InDegrees(r.Y),
		Z: Angle.InDegrees(r.Z),
	}
}

type Degrees struct {
	X Angle.Degrees
	Y Angle.Degrees
	Z Angle.Degrees
}

func (d Degrees) Vector3() vector3 {
	return vector3{
		X: Float.X(d.X),
		Y: Float.X(d.Y),
		Z: Float.X(d.Z),
	}
}

func (r Degrees) Radians() Radians {
	return Radians{
		X: Angle.InRadians(r.X),
		Y: Angle.InRadians(r.Y),
		Z: Angle.InRadians(r.Z),
	}
}
