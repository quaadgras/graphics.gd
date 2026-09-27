package Basis_test

import (
	"fmt"
	"testing"

	"graphics.gd/internal/gdtests"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Quaternion"
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

// rotations is a spread of unit quaternions to check conversions with,
// including half turns (where a transposed matrix is its own inverse and
// hides the mistake) and rotations about every axis.
func rotations() []Quaternion.IJKX {
	var qs []Quaternion.IJKX
	for _, axis := range []Vector3.XYZ{Vector3.Right, Vector3.Up, Vector3.Back, Vector3.Normalized(Vector3.New(1, 2, -3)), Vector3.Normalized(Vector3.New(-0.4, 0.1, 0.9))} {
		for _, angle := range []Angle.Radians{0.3, -1.2, Angle.Pi / 2, -Angle.Pi / 2, 2.8, Angle.Pi} {
			s, c := Angle.Sin(angle/2), Angle.Cos(angle/2)
			qs = append(qs, Quaternion.IJKX{I: axis.X * s, J: axis.Y * s, K: axis.Z * s, X: c})
		}
	}
	return qs
}

func near(a, b Vector3.XYZ) bool { return Vector3.Distance(a, b) < 1e-4 }

// TestQuaternionConversions checks that a quaternion and its basis turn
// vectors the same way in both directions: AsQuaternion and AsBasis used
// to read and write the matrix transposed, giving the inverse rotation.
func TestQuaternionConversions(t *testing.T) {
	probes := []Vector3.XYZ{Vector3.Right, Vector3.Up, Vector3.Back, Vector3.New(0.3, -2, 1.5)}
	for _, q := range rotations() {
		b := Quaternion.AsBasis(q)
		scaled := Basis.RotatesScales(q, Vector3.New(1, 1, 1))
		back := Basis.AsQuaternion(b)
		axisAngle := Basis.RotatesAxisAngle(Vector3.Normalized(Vector3.New(q.I, q.J, q.K)), Quaternion.AngleInRadians(q))
		for _, v := range probes {
			want := Quaternion.Rotate(v, q)
			if got := Basis.Transform(v, b); !near(got, want) {
				t.Fatalf("Quaternion.AsBasis(%v) turns %v to %v, the quaternion to %v", q, v, got, want)
			}
			if got := Basis.Transform(v, scaled); !near(got, want) {
				t.Fatalf("Basis.RotatesScales(%v) turns %v to %v, the quaternion to %v", q, v, got, want)
			}
			if got := Quaternion.Rotate(v, back); !near(got, want) {
				t.Fatalf("Basis.AsQuaternion(%v) turns %v to %v, the basis to %v", b, v, got, want)
			}
			if got := Basis.Transform(v, axisAngle); !near(got, want) {
				t.Fatalf("Basis.RotatesAxisAngle for %v turns %v to %v, the quaternion to %v", q, v, got, want)
			}
		}
		// Scale must not leak into the rotation.
		if got := Quaternion.Rotate(Vector3.Up, Basis.AsQuaternion(Basis.ScaledLocal(b, Vector3.New(2, 3, 0.5)))); !near(got, Quaternion.Rotate(Vector3.Up, q)) {
			t.Fatalf("Basis.AsQuaternion of a scaled %v turns up to %v", q, got)
		}
		// And the Euler angles of the quaternion are those of its basis.
		e := Quaternion.EulerRadians(Angle.OrderYXZ, q)
		if got, want := Basis.Transform(Vector3.Back, Basis.FromEuler(e, Angle.OrderYXZ)), Quaternion.Rotate(Vector3.Back, q); !near(got, want) {
			t.Fatalf("Quaternion.EulerRadians(%v) = %v turns back to %v, want %v", q, e, got, want)
		}
	}
}

// TestSlerp checks Basis.Slerp lands on the rotation part way between.
func TestSlerp(t *testing.T) {
	from := Basis.RotatesAxisAngle(Vector3.Up, 0.2)
	to := Basis.RotatesAxisAngle(Vector3.Up, 1.4)
	for _, w := range []float64{0, 0.25, 0.5, 1} {
		want := Basis.Transform(Vector3.Right, Basis.RotatesAxisAngle(Vector3.Up, Angle.Radians(0.2+1.2*w)))
		if got := Basis.Transform(Vector3.Right, Basis.Slerp(from, to, w)); !near(got, want) {
			t.Fatalf("Slerp weight %v turns right to %v, want %v", w, got, want)
		}
	}
}

// TestOuter checks the outer product is v times with transposed: column
// j is v scaled by with's j-th component.
func TestOuter(t *testing.T) {
	v, with := Vector3.New(1, 2, 3), Vector3.New(4, 5, 6)
	got := Basis.Transform(Vector3.New(0.5, -1, 2), Basis.Outer(v, with))
	want := Vector3.MulX(v, Vector3.Dot(with, Vector3.New(0.5, -1, 2)))
	if !near(got, want) {
		t.Fatalf("Outer(%v, %v) maps to %v, want %v", v, with, got, want)
	}
}

// TestTransposedDot checks tdotx/y/z are the dot products with the
// basis's columns, i.e. the transpose applied to the vector.
func TestTransposedDot(t *testing.T) {
	b := Basis.Mul(Basis.RotatesAxisAngle(Vector3.Normalized(Vector3.New(1, 2, 3)), 0.7), Basis.Scales(Vector3.New(1, 2, 3)))
	v := Vector3.New(0.3, -1, 2)
	want := Basis.Transform(v, Basis.Transposed(b))
	got := Vector3.New(Basis.TransposedDotX(b, v), Basis.TransposedDotY(b, v), Basis.TransposedDotZ(b, v))
	if !near(got, want) {
		t.Fatalf("TransposedDot of %v = %v, want %v", v, got, want)
	}
}

// TestLookingAt checks the forward axis (-Z) points at the target and
// up stays up.
func TestLookingAt(t *testing.T) {
	target := Vector3.New(3, 1, -2)
	b := Basis.LookingAt(target, Vector3.Up)
	if got := Basis.Transform(Vector3.Forward, b); !near(got, Vector3.Normalized(target)) {
		t.Fatalf("LookingAt(%v) points forward along %v", target, got)
	}
	if up := Basis.Transform(Vector3.Up, b); up.Y <= 0 || Vector3.Dot(up, target) > 1e-4 {
		t.Fatalf("LookingAt(%v) has up %v", target, up)
	}
	if !Basis.IsOrthonormal(b) || Basis.Determinant(b) < 0.999 {
		t.Fatalf("LookingAt(%v) is not a rotation: %v", target, b)
	}
}
