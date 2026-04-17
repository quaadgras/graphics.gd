package swizzle

import (
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

func RGB[T vec3.XYZ | vec4.XYZW | vec3.RGB | vec4.RGBA](v T) vec3.RGB {
	switch v := any(v).(type) {
	case vec4.RGBA:
		return vec3.RGB{R: v.R, G: v.G, B: v.B}
	case vec3.XYZ:
		return vec3.RGB{R: v.X, G: v.Y, B: v.Z}
	case vec4.XYZW:
		return vec3.RGB{R: v.X, G: v.Y, B: v.Z}
	case vec3.RGB:
		return v
	default:
		panic("unreachable")
	}
}
