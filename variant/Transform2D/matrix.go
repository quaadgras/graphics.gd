// Package Transform2D provides a 2×3 matrix representing a 2D transformation.
package Transform2D

import (
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

// Transform2D is a 2×3 matrix representing a transformation in 2D space.
// It contains three Vector2 values: x, y, and origin. Together, they can represent translation,
// rotation, scale, and skew.
//
// The x and y axes form a 2×2 matrix, known as the transform's basis. The length of each axis
// (Vector2.Length) influences the transform's scale, while the direction of all axes influence
// the rotation. Usually, both axes are perpendicular to one another. However, when you rotate
// one axis individually, the transform becomes skewed. Applying a skewed transform to a 2D sprite
// will make the sprite appear distorted.
//
// For a general introduction, see the Matrices and transforms tutorial.
//
// Note: Unlike Transform3D, there is no 2D equivalent to the Basis type. All mentions of "basis"
// refer to the x and y components of Transform2D.
type OriginXY = struct {
	// The transform basis's X axis, and the column 0 of the matrix. Combined with y, this represents
	// the transform's rotation, scale, and skew.
	//
	// 	On the identity transform, this vector points right [Vector2.Right].
	X Vector2.XY

	// The transform basis's Y axis, and the column 1 of the matrix. Combined with x, this represents
	// the transform's rotation, scale, and skew.
	//
	// On the identity transform, this vector points up (Vector2.Up).
	Y Vector2.XY

	// Origin of this transform, and the column 2 of the matrix. In 2D space, this can be seen
	// as the position.
	Origin Vector2.XY
}

var (
	// Identity Transform2D. A transform with no translation, no rotation, and its scale being 1.
	// When multiplied by another value such as Rect2 or another Transform2D, no transformation
	// occurs. This means that:
	//
	//  - The x points right [Vector2.Right]
	//  - The y points up [Vector2.Up]
	//
	Identity = OriginXY{
		X: Vector2.Right,
		Y: Vector2.XY{0, 1},
	}

	// When any transform is multiplied by FlipX, it negates all components of the x axis (the X column).
	// When FlipX is multiplied by any basis, it negates the Vector2.X component of all axes (the X row).
	FlipX = OriginXY{
		X: Vector2.New(-1, 0),
		Y: Vector2.New(0, 1),
	}

	// When any transform is multiplied by FlipY, it negates all components of the y axis (the Y column).
	//
	// When FlipY is multiplied by any basis, it negates the Vector2.Y component of all axes (the Y row).
	FlipY = OriginXY{
		X: Vector2.New(1, 0),
		Y: Vector2.New(0, -1),
	}
)

// New returns the identity transform.
func New() OriginXY {
	return OriginXY{
		X: Vector2.Right,
		Y: Vector2.Up,
	}
}

// RotationScaleSkewPosition constructs a new Transform2D from the given rotation and position.
func RotationScaleSkewPosition(rotation Angle.Radians, scale Vector2.XY, skew Angle.Radians, position Vector2.XY) OriginXY { //gd:Transform2D(rotation:float,scale:Vector2,skew:float,position:Vector2)
	r := rotation
	s := scale
	k := skew
	p := position
	return OriginXY{
		Vector2.New(s.X*Float.X(Angle.Cos(r)), s.Y*Angle.Sin(r)),
		Vector2.New(s.X*-Angle.Sin(r+k), s.Y*Angle.Cos(r+k)),
		p,
	}
}

func tdotx(t OriginXY, v Vector2.XY) Float.X { return t.X.X*v.X + t.Y.X*v.Y }
func tdoty(t OriginXY, v Vector2.XY) Float.X { return t.X.Y*v.X + t.Y.Y*v.Y }

// AffineInverse returns the inverse of the transform, under the assumption that the basis is
// invertible (must have non-zero determinant).
func AffineInverse(t OriginXY) OriginXY { //gd:Transform2D.affine_inverse
	det := Determinant(t)
	idet := 1 / det
	t.X.X, t.Y.Y = t.Y.Y, t.X.X
	t.X = Vector2.Mul(t.X, Vector2.New(idet, -idet))
	t.Y = Vector2.Mul(t.Y, Vector2.New(-idet, idet))
	t.Origin = BasisTransform(t, Vector2.Neg(t.Origin))
	return t
}

// BasisTransform returns a vector transformed (multiplied) by the basis matrix.
//
// This method does not account for translation (the origin vector).
func BasisTransform(t OriginXY, v Vector2.XY) Vector2.XY { //gd:Transform2D.basis_xform
	return Vector2.New(tdotx(t, v), tdoty(t, v))
}

// InverseBasisTransform returns a vector transformed (multiplied) by the inverse basis matrix,
// under the assumption that the basis is orthonormal (i.e. rotation/reflection is fine, scaling/skew is not).
//
// This method does not account for translation (the origin vector).
//
// transform.InverseBasisTransform(vector) is equivalent to transform.Inverse().BasisTransform(vector).
// See [Transform2D.Inverse].
//
// For non-orthonormal transforms (e.g. with scaling) transform.AffineInverse().BasisTransform(vector)
// can be used instead. See [Transform2D.AffineInverse].
func InverseBasisTransform(t OriginXY, v Vector2.XY) Vector2.XY { //gd:Transform2D.basis_xform_inv
	return BasisTransform(Inverse(t), v)
}

// Determinant returns the determinant of the basis matrix. If the basis is uniformly scaled, then its
// determinant equals the square of the scale factor.
//
// A negative determinant means the basis was flipped, so one part of the scale is negative. A zero
// determinant means the basis isn't invertible, and is usually considered invalid.
func Determinant(t OriginXY) Float.X { //gd:Transform2D.determinant
	return Float.X(t.X.X*t.Y.Y - t.X.Y*t.X.X)
}

// Origin returns the transform's origin (translation).
func Origin(t OriginXY) Vector2.XY { return t.Origin } //gd:Transform2D.get_origin

// Rotation returns the transform's rotation.
func Rotation(t OriginXY) Angle.Radians { return Angle.Atan2(t.X.Y, t.X.X) } //gd:Transform2D.get_rotation

// Scale returns the transform's scale.
func Scale(t OriginXY) Vector2.XY { //gd:Transform2D.get_scale
	var det_sign = Float.Sign(Determinant(t))
	return Vector2.New(Vector2.Length(t.X), det_sign*(Vector2.Length(t.Y)))
}

// Skew returns the transform's skew (in radians)
func Skew(t OriginXY) Angle.Radians { //gd:Transform2D.get_skew
	return Angle.Acos(Vector2.Dot(Vector2.Normalized(t.X), Vector2.MulX(Vector2.Normalized(t.Y), Float.Sign(Determinant(t))))) - Angle.Pi*0.5
}

// Lerp returns a transform interpolated between this transform and
// another by a given weight (on the range of 0.0 to 1.0).
func Lerp[X Float.Any](a, b OriginXY, weight X) OriginXY { //gd:Transform2D.interpolate_with
	return RotationScaleSkewPosition(
		Angle.Lerp(Rotation(a), Rotation(b), weight),
		Vector2.Lerp(Scale(a), Scale(b), weight),
		Angle.Lerp(Skew(a), Skew(b), weight),
		Vector2.Lerp(Origin(a), Origin(b), weight),
	)
}

// Inverse returns the inverse of the transform, under the assumption that the transformation basis is
// orthonormal (i.e. rotation/reflection is fine, scaling/skew is not). Use affine_inverse for
// non-orthonormal transforms (e.g. with scaling).
func Inverse(t OriginXY) OriginXY { //gd:Transform2D.inverse
	t.X.Y, t.Y.X = t.Y.X, t.X.Y
	t.Origin = InverseBasisTransform(t, Vector2.Neg(t.Origin))
	return t
}

// IsConformal returns true if the transform's basis is conformal, meaning it preserves angles and distance
// ratios, and may only be composed of rotation and uniform scale. Returns false if the transform's basis
// has non-uniform scale or shear/skew. This can be used to validate if the transform is non-distorted,
// which is important for physics and other use cases.
func IsConformal(t OriginXY) bool { //gd:Transform2D.is_conformal
	if Float.IsApproximatelyEqual(t.X.X, t.Y.Y) && Float.IsApproximatelyEqual(t.X.Y, -t.Y.X) {
		return true // Non-flipped case.
	}
	if Float.IsApproximatelyEqual(t.X.X, -t.Y.Y) && Float.IsApproximatelyEqual(t.X.Y, t.Y.X) {
		return true // Flipped case.
	}
	return false
}

// IsApproximatelyEqual returns true if this transform and xform are approximately equal, by running
// [Vector2.IsApproximatelyEqual] on each component.
func IsApproximatelyEqual(a, b OriginXY) bool { //gd:Transform2D.is_equal_approx
	return Vector2.IsApproximatelyEqual(a.X, b.X) && Vector2.IsApproximatelyEqual(a.Y, b.Y) && Vector2.IsApproximatelyEqual(a.Origin, b.Origin)
}

// IsFinite returns true if this transform is finite, by calling [Vector2.IsFinite] on each component.
func IsFinite(t OriginXY) bool { //gd:Transform2D.is_finite
	return Vector2.IsFinite(t.X) && Vector2.IsFinite(t.Y) && Vector2.IsFinite(t.Origin)
}

// LookingAt returns a copy of the transform rotated such that the rotated X-axis points towards the target position.
//
// Operations take place in global space.
func LookingAt(target Vector2.XY, t OriginXY) OriginXY { //gd:Transform2D.looking_at
	var return_trans = RotationScaleSkewPosition(Rotation(t), Vector2.One, 0, Origin(t))
	var target_position = Vector(target, AffineInverse(t))
	return Rotated(return_trans, Rotation(return_trans)+Vector2.AngleRadians(Vector2.Mul(target_position, Scale(t))))
}

// Orthonormalized returns the transform with the basis orthogonal (90 degrees), and normalized axis vectors
// (scale of 1 or -1).
func Orthonormalized(t OriginXY) OriginXY { //gd:Transform2D.orthonormalized
	var x = t.X
	var y = t.Y
	x = Vector2.Normalized(x)
	y = Vector2.SubX(y, Vector2.Dot(Vector2.Mul(x, x), y))
	y = Vector2.Normalized(y)
	t.X = x
	t.Y = y
	return t
}

// Rotated returns a copy of the transform rotated by the given angle (in radians).
//
// This method is an optimized version of multiplying the given transform X with a corresponding rotation
// transform R from the left, i.e., R * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func Rotated(t OriginXY, angle Angle.Radians) OriginXY { //gd:Transform2D.rotated
	return Mul(RotationScaleSkewPosition(angle, Vector2.One, 0, Vector2.Zero), t)
}

// RotatedLocal returns a copy of the transform rotated by the given angle (in radians).
//
// This method is an optimized version of multiplying the given transform X with a corresponding
// rotation transform R from the right, i.e., X * R.
//
// This can be seen as transforming with respect to the local frame.
func RotatedLocal(t OriginXY, angle Angle.Radians) OriginXY { //gd:Transform2D.rotated_local
	return Mul(t, RotationScaleSkewPosition(angle, Vector2.One, 0, Vector2.Zero))
}

// Scaled returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X
// with a corresponding scaling transform S from the left, i.e., S * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func Scaled(t OriginXY, scale Vector2.XY) OriginXY { //gd:Transform2D.scaled
	t.X.X *= scale.X
	t.X.Y *= scale.Y
	t.Y.X *= scale.X
	t.Y.Y *= scale.Y
	t.Origin = Vector2.Mul(t.Origin, scale)
	return t
}

// ScaledLocal returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding scaling transform S from the right, i.e., X * S.
//
// This can be seen as transforming with respect to the local frame.
func ScaledLocal(t OriginXY, scale Vector2.XY) OriginXY { //gd:Transform2D.scaled_local
	return OriginXY{
		Vector2.MulX(t.X, scale.X),
		Vector2.MulX(t.Y, scale.Y),
		t.Origin,
	}
}

// Translated returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding translation transform T from the left, i.e., T * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func Translated(t OriginXY, offset Vector2.XY) OriginXY { //gd:Transform2D.translated
	return OriginXY{
		t.X,
		t.Y,
		Vector2.Add(t.Origin, offset),
	}
}

// TranslatedLocal returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding translation transform T from the right, i.e., X * T.
//
// This can be seen as transforming with respect to the local frame.
func TranslatedLocal(t OriginXY, offset Vector2.XY) OriginXY { //gd:Transform2D.translated_local
	return OriginXY{
		t.X,
		t.Y,
		Vector2.Add(t.Origin, BasisTransform(t, offset)),
	}
}

func Mul(a, b OriginXY) OriginXY { //gd:Transform2D*(right:Transform2D)
	a.Origin = Vector(b.Origin, a)
	var (
		x0, x1, y0, y1 = tdotx(a, b.X), tdoty(a, b.X), tdotx(a, b.Y), tdoty(a, b.Y)
	)
	a.X.X = x0
	a.X.Y = x1
	a.Y.X = y0
	a.Y.Y = y1
	return a
}

func Vector(v Vector2.XY, t OriginXY) Vector2.XY { //gd:Transform2D*(right:Vector2)
	return Vector2.Add(Vector2.New(tdotx(t, v), tdoty(t, v)), t.Origin)
}
