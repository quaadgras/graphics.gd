//go:build !generate

package gd

/*
Transform2D is a 2×3 matrix (2 rows, 3 columns) used for 2D linear transformations. It can represent transformations such as translation,
rotation, and scaling. It consists of three Vector2 values: x, y, and the origin.

For more information, read the "Matrices and transforms" documentation article.
*/
type Transform2D [3]Vector2

// NewTransform2D constructs a new Transform2D from the given rotation and position.
func NewTransform2D(rotation Radians, scale Vector2, skew Radians, position Vector2) Transform2D {
	r := float32(rotation)
	s := scale
	k := float32(skew)
	p := position
	return Transform2D{
		Vector2{s[X] * Cos(r), s[Y] * Sin(r)},
		Vector2{s[X] * -Sin(r+k), s[Y] * Cos(r+k)},
		p,
	}
}

// "Fields"

func (t Transform2D) X() Vector2      { return t[0] }
func (t Transform2D) Y() Vector2      { return t[1] }
func (t Transform2D) Origin() Vector2 { return t[2] }
func (t *Transform2D) SetX(x Vector2) { t[0] = x }
func (t *Transform2D) SetY(y Vector2) { t[1] = y }
func (t *Transform2D) SetOrigin(o Vector2) {
	t[2] = o
}

// "Constants"

func (t Transform2D) IDENTITY() Transform2D {
	return Transform2D{
		Vector2{1, 0},
		Vector2{0, 1},
		Vector2{0, 0},
	}
}
func (t Transform2D) FLIP_X() Transform2D {
	return Transform2D{
		Vector2{-1, 0},
		Vector2{0, 1},
		Vector2{0, 0},
	}
}
func (t Transform2D) FLIP_Y() Transform2D {
	return Transform2D{
		Vector2{1, 0},
		Vector2{0, -1},
		Vector2{0, 0},
	}
}

func (t Transform2D) tdotx(v Vector2) float32 { return t[0][x]*v[x] + t[1][x]*v[y] }
func (t Transform2D) tdoty(v Vector2) float32 { return t[0][y]*v[x] + t[1][y]*v[y] }

// AffineInverse returns the inverse of the transform, under the assumption that the basis is
// invertible (must have non-zero determinant).
func (t Transform2D) AffineInverse() Transform2D {
	det := t.Determinant()
	idet := float32(1 / det)
	t[0][0], t[1][1] = t[1][1], t[0][0]
	t[0] = t[0].Mul(Vector2{idet, -idet})
	t[1] = t[1].Mul(Vector2{-idet, idet})
	t[2] = t.BasisTransform(t[2].Neg())
	return t
}

// BasisTransform returns a vector transformed (multiplied) by the basis matrix.
//
// This method does not account for translation (the origin vector).
func (t Transform2D) BasisTransform(v Vector2) Vector2 { return Vector2{t.tdotx(v), t.tdoty(v)} }

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
func (t Transform2D) InverseBasisTransform(v Vector2) Vector2 {
	return t.Inverse().BasisTransform(v)
}

// Determinant returns the determinant of the basis matrix. If the basis is uniformly scaled, then its
// determinant equals the square of the scale factor.
//
// A negative determinant means the basis was flipped, so one part of the scale is negative. A zero
// determinant means the basis isn't invertible, and is usually considered invalid.
func (t Transform2D) Determinant() Float {
	return Float(t[0][X]*t[1][Y] - t[0][Y]*t[1][X])
}

// GetOrigin returns the transform's origin (translation).
func (t Transform2D) GetOrigin() Vector2 { return t[2] }

// GetRotation returns the transform's rotation.
func (t Transform2D) GetRotation() Radians { return Atan2(t[0][Y], t[0][X]) }

// GetScale returns the transform's scale.
func (t Transform2D) GetScale() Vector2 {
	var det_sign = Signf(t.Determinant())
	return Vector2{float32(t[0].Length()), float32(det_sign * t[1].Length())}
}

// GetSkew returns the transform's skew (in radians)
func (t Transform2D) GetSkew() Radians {
	return Radians(Acos(t[0].Normalized().Dot(t[1].Normalized().Mulf(Sign(t.Determinant()))))) - Pi*0.5
}

// InterpolateWith returns a transform interpolated between this transform and
// another by a given weight (on the range of 0.0 to 1.0).
func (t Transform2D) InterpolateWith(b Transform2D, weight Float) Transform2D {
	return NewTransform2D(
		LerpAngle(t.GetRotation(), b.GetRotation(), Radians(weight)),
		t.GetScale().Lerp(b.GetScale(), weight),
		LerpAngle(t.GetSkew(), b.GetSkew(), Radians(weight)),
		t.GetOrigin().Lerp(b.GetOrigin(), weight),
	)
}

// Inverse returns the inverse of the transform, under the assumption that the transformation basis is
// orthonormal (i.e. rotation/reflection is fine, scaling/skew is not). Use affine_inverse for
// non-orthonormal transforms (e.g. with scaling).
func (t Transform2D) Inverse() Transform2D {
	t[0][1], t[1][0] = t[1][0], t[0][1]
	t[2] = t.InverseBasisTransform(t[2].Neg())
	return t
}

// IsConformal returns true if the transform's basis is conformal, meaning it preserves angles and distance
// ratios, and may only be composed of rotation and uniform scale. Returns false if the transform's basis
// has non-uniform scale or shear/skew. This can be used to validate if the transform is non-distorted,
// which is important for physics and other use cases.
func (t Transform2D) IsConformal() bool {
	// Non-flipped case.
	if IsApproximatelyEqual(t[0][0], t[1][1]) && IsApproximatelyEqual(t[0][1], -t[1][0]) {
		return true
	}
	// Flipped case.
	if IsApproximatelyEqual(t[0][0], -t[1][1]) && IsApproximatelyEqual(t[0][1], t[1][0]) {
		return true
	}
	return false
}

// IsApproximatelyEqual returns true if this transform and xform are approximately equal, by running
// [Vector2.IsApproximatelyEqual] on each component.
func (t Transform2D) IsApproximatelyEqual(xform Transform2D) bool {
	return t[0].IsApproximatelyEqual(xform[0]) && t[1].IsApproximatelyEqual(xform[1]) && t[2].IsApproximatelyEqual(xform[2])
}

// IsFinite returns true if this transform is finite, by calling [Vector2.IsFinite] on each component.
func (t Transform2D) IsFinite() bool { return t[0].IsFinite() && t[1].IsFinite() && t[2].IsFinite() }

// LookingAt returns a copy of the transform rotated such that the rotated X-axis points towards the target position.
//
// Operations take place in global space.
func (t Transform2D) LookingAt(target Vector2) Transform2D {
	var return_trans = NewTransform2D(t.GetRotation(), Vector2{1, 1}, 0, t.GetOrigin())
	var target_position = target.Transform(t.AffineInverse())
	return return_trans.Rotated(return_trans.GetRotation() + (target_position.Mul(t.GetScale())).Angle())
}

// Orthonormalized returns the transform with the basis orthogonal (90 degrees), and normalized axis vectors
// (scale of 1 or -1).
func (t Transform2D) Orthonormalized() Transform2D {
	var x = t[0]
	var y = t[1]
	x = x.Normalized()
	y = y.Subf(x.Mul(x).Dot(y))
	y = y.Normalized()
	t[0] = x
	t[1] = y
	return t
}

// Rotated returns a copy of the transform rotated by the given angle (in radians).
//
// This method is an optimized version of multiplying the given transform X with a corresponding rotation
// transform R from the left, i.e., R * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func (t Transform2D) Rotated(angle Radians) Transform2D {
	return NewTransform2D(angle, Vector2{1, 1}, 0, Vector2{}).Mul(t)
}

// RotatedLocal returns a copy of the transform rotated by the given angle (in radians).
//
// This method is an optimized version of multiplying the given transform X with a corresponding
// rotation transform R from the right, i.e., X * R.
//
// This can be seen as transforming with respect to the local frame.
func (t Transform2D) RotatedLocal(angle Radians) Transform2D {
	return t.Mul(NewTransform2D(angle, Vector2{1, 1}, 0, Vector2{}))
}

// Scaled returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X
// with a corresponding scaling transform S from the left, i.e., S * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func (t Transform2D) Scaled(scale Vector2) Transform2D {
	t[0][0] *= scale[X]
	t[0][1] *= scale[Y]
	t[1][0] *= scale[X]
	t[1][1] *= scale[Y]
	t[2] = t[2].Mul(scale)
	return t
}

// ScaledLocal returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding scaling transform S from the right, i.e., X * S.
//
// This can be seen as transforming with respect to the local frame.
func (t Transform2D) ScaledLocal(scale Vector2) Transform2D {
	return Transform2D{
		t[0].Mulf(Float(scale[X])),
		t[1].Mulf(Float(scale[Y])),
		t[2],
	}
}

// Translated returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding translation transform T from the left, i.e., T * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func (t Transform2D) Translated(offset Vector2) Transform2D {
	return Transform2D{
		t[0],
		t[1],
		t[2].Add(offset),
	}
}

// TranslatedLocal returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding translation transform T from the right, i.e., X * T.
//
// This can be seen as transforming with respect to the local frame.
func (t Transform2D) TranslatedLocal(offset Vector2) Transform2D {
	return Transform2D{
		t[0],
		t[1],
		t[2].Add(t.BasisTransform(offset)),
	}
}

func (t Transform2D) Mul(other Transform2D) Transform2D {
	t[2] = other[2].Transform(t)

	var x0, x1, y0, y1 float32

	x0 = t.tdotx(other[0])
	x1 = t.tdoty(other[0])
	y0 = t.tdotx(other[1])
	y1 = t.tdoty(other[1])

	t[0][0] = x0
	t[0][1] = x1
	t[1][0] = y0
	t[1][1] = y1
	return t
}
