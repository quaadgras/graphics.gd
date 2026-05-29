//go:build !generate

package gd

import (
	basis "graphics.gd/variant/Basis"
	"graphics.gd/variant/Transform3D"
)

func Transposed(t Transform3D.BasisOrigin) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin{
		Basis:  basis.Transposed(t.Basis),
		Origin: t.Origin,
	}
}
