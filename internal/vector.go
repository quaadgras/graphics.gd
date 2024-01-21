//go:build !generate

package gd

import "math"

type Vector2 struct {
	X, Y float32
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector2) ScaledBy(f Float) Vector2 {
	return Vector2{X: float32(Float(v.X) * f), Y: float32(Float(v.Y) * f)}
}

func (v Vector2) Length() Float {
	return Float(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) Normalized() Vector2 {
	var length = v.Length()
	return Vector2{X: float32(Float(v.X) / length), Y: float32(Float(v.Y) / length)}
}

type Vector2i struct {
	X, Y int32
}

type Vector3 struct {
	X, Y, Z float32
}

type Vector3i struct {
	X, Y, Z int32
}
