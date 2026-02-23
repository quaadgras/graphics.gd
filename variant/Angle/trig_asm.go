//go:build !precision_double && amd64

package Angle

func sin32(x float32) float32
func cos32(x float32) float32
func sincos32(x float32) (float32, float32)
