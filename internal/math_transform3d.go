//go:build !generate

package gd

type Transform3D struct {
	Basis  Basis
	Origin Vector3
}
