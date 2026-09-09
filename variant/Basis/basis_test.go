package Basis_test

import (
	"fmt"
	"testing"

	"graphics.gd/internal/gdtests"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Vector3"
)

func TestIdentity(t *testing.T) {
	var basis = Basis.Identity
	gdtests.Print(t, "| X | Y | Z", "| X | Y | Z")
	gdtests.Print(t, "| 1 | 0 | 0", fmt.Sprintf("| %v | %v | %v", basis.X.X, basis.Y.X, basis.Z.X))
	gdtests.Print(t, "| 0 | 1 | 0", fmt.Sprintf("| %v | %v | %v", basis.X.Y, basis.Y.Y, basis.Z.Y))
	gdtests.Print(t, "| 0 | 0 | 1", fmt.Sprintf("| %v | %v | %v", basis.X.Z, basis.Y.Z, basis.Z.Z))
	// Prints:
	// | X | Y | Z
	// | 1 | 0 | 0
	// | 0 | 1 | 0
	// | 0 | 0 | 1
}

func TestScaled(t *testing.T) {
	var my_basis = Basis.XYZ{
		X: Vector3.XYZ{X: 1, Y: 1, Z: 1},
		Y: Vector3.XYZ{X: 2, Y: 2, Z: 2},
		Z: Vector3.XYZ{X: 3, Y: 3, Z: 3},
	}
	my_basis = Basis.Scaled(my_basis, Vector3.New(0, 2, -2))
	gdtests.Print(t, "(0.0, 2.0, -2.0)", fmt.Sprintf("(%.1f, %.1f, %.1f)", my_basis.X.X, my_basis.X.Y, my_basis.X.Z)) // Prints (0.0, 0.0, 0.0)
	gdtests.Print(t, "(0.0, 4.0, -4.0)", fmt.Sprintf("(%.1f, %.1f, %.1f)", my_basis.Y.X, my_basis.Y.Y, my_basis.Y.Z)) // Prints (4.0, 4.0, 4.0)
	gdtests.Print(t, "(0.0, 6.0, -6.0)", fmt.Sprintf("(%.1f, %.1f, %.1f)", my_basis.Z.X, my_basis.Z.Y, my_basis.Z.Z)) // Prints (-6.0, -6.0, -6.0)
}

func TestScaledLocal(t *testing.T) {
	var my_basis = Basis.XYZ{
		X: Vector3.XYZ{X: 1, Y: 1, Z: 1},
		Y: Vector3.XYZ{X: 2, Y: 2, Z: 2},
		Z: Vector3.XYZ{X: 3, Y: 3, Z: 3},
	}
	my_basis = Basis.ScaledLocal(my_basis, Vector3.New(0, 2, -2))
	gdtests.Print(t, "(0.0, 0.0, 0.0)", fmt.Sprintf("(%.1f, %.1f, %.1f)", my_basis.X.X, my_basis.X.Y, my_basis.X.Z))    // Prints (0.0, 0.0, 0.0)
	gdtests.Print(t, "(4.0, 4.0, 4.0)", fmt.Sprintf("(%.1f, %.1f, %.1f)", my_basis.Y.X, my_basis.Y.Y, my_basis.Y.Z))    // Prints (4.0, 4.0, 4.0)
	gdtests.Print(t, "(-6.0, -6.0, -6.0)", fmt.Sprintf("(%.1f, %.1f, %.1f)", my_basis.Z.X, my_basis.Z.Y, my_basis.Z.Z)) // Prints (-6.0, -6.0, -6.0)
}

// TestEulerRoundTrip checks that AsEulerAngles inverts FromEuler for
// every Euler order: the decomposition used to read the matrix
// transposed, so a +90° yaw came back as -90°.
func TestEulerRoundTrip(t *testing.T) {
	orders := []Angle.Order{Angle.OrderXYZ, Angle.OrderXZY, Angle.OrderYXZ, Angle.OrderYZX, Angle.OrderZXY, Angle.OrderZYX}
	angles := []Angle.Radians{0, 0.3, -0.7, Angle.Pi / 2, -Angle.Pi / 2, Angle.Pi, 2.5}
	for _, order := range orders {
		for _, x := range angles {
			for _, y := range angles {
				for _, z := range angles {
					want := Basis.FromEuler(Euler.Radians{X: x, Y: y, Z: z}, order)
					got := Basis.FromEuler(Basis.AsEulerAngles(want, order), order)
					for _, axis := range []Vector3.XYZ{Vector3.Right, Vector3.Up, Vector3.Back} {
						a, b := Basis.Transform(axis, want), Basis.Transform(axis, got)
						if Vector3.Distance(a, b) > 1e-4 {
							t.Fatalf("order %v euler (%v, %v, %v): axis %v maps to %v, round-tripped to %v", order, x, y, z, axis, a, b)
						}
					}
				}
			}
		}
	}
}
