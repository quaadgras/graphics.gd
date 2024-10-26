package Angle

import "math"

type Radians float
type Degrees float
type XYZ = struct {
	X Radians
	Y Radians
	Z Radians
}

type vector2 = struct {
	X float
	Y float
}

// AsVector2 creates a unit Vector2 rotated to the given angle in radians. This is equivalent
// to doing:
//
//	Vector2.New(math.Cos(angle), math.Sin(angle))
//	Vector2.Rotated(Vector2.Right, angle).
func (angle Radians) AsVector2() vector2 { //gd:Vector2.from_angle
	return vector2{float(math.Cos(float64(angle))), float(math.Sin(float64(angle)))}
}
