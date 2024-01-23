//go:build !generate

package gd

import "math"

const (
	x = iota
	y
	z
	w
)

type Vector2 [2]float32

func (v Vector2) Add(other Vector2) Vector2 { return Vector2{v[x] + other[x], v[y] + other[y]} }

func (v Vector2) ScaledBy(f Float) Vector2 {
	return Vector2{float32(Float(v[x]) * f), float32(Float(v[y]) * f)}
}

func (v Vector2) Length() Float {
	return Float(math.Sqrt(float64(Float(v[x])*Float(v[x]) + Float(v[y])*Float(v[y]))))
}

func (v Vector2) Normalized() Vector2 {
	length := v.Length()
	return Vector2{float32(Float(v[x]) / length), float32(Float(v[y]) / length)}
}

type Vector2i [2]int32

type Vector3 [3]float32

func (v Vector3) ScaledBy(f Float) Vector3 {
	return Vector3{float32(Float(v[x]) * f), float32(Float(v[y]) * f), float32(Float(v[z]) * f)}
}

func (v Vector3) Length() Float {
	return Float(math.Sqrt(float64(Float(v[x])*Float(v[x]) + Float(v[y])*Float(v[y]) + Float(v[z])*Float(v[z]))))
}

type Vector3i [3]int32

type Vector4 [4]float32

type Vector4i [4]int32
