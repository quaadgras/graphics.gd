package spatial

/*
Transform3D is a 3×4 matrix (3 rows, 4 columns) used for 3D linear transformations. It can represent transformations such as
translation, rotation, and scaling. It consists of a basis (first 3 columns) and a Vector3 for the origin (last column).
*/
type Transform3D struct {
	Basis  Basis
	Origin Vector3
}

// "Constants"

func (Transform3D) IDENTITY() Transform3D {
	return Transform3D{
		Basis:  Basis{Vector3{1, 0, 0}, Vector3{0, 1, 0}, Vector3{0, 0, 1}},
		Origin: Vector3{0, 0, 0},
	}
}
func (Transform3D) FLIP_X() Transform3D {
	return Transform3D{
		Basis:  Basis{Vector3{-1, 0, 0}, Vector3{0, 1, 0}, Vector3{0, 0, 1}},
		Origin: Vector3{0, 0, 0},
	}
}
func (Transform3D) FLIP_Y() Transform3D {
	return Transform3D{
		Basis:  Basis{Vector3{1, 0, 0}, Vector3{0, -1, 0}, Vector3{0, 0, 1}},
		Origin: Vector3{0, 0, 0},
	}
}
func (Transform3D) FLIP_Z() Transform3D {
	return Transform3D{
		Basis:  Basis{Vector3{1, 0, 0}, Vector3{0, 1, 0}, Vector3{0, 0, -1}},
		Origin: Vector3{0, 0, 0},
	}
}

// "Methods"

// AffineInverse returns the inverse of the transform, under the assumption that the basis is
// invertible (must have non-zero determinant).
func (t Transform3D) AffineInverse() Transform3D {
	var inv = t.Basis.Inverse()
	return Transform3D{
		Basis:  inv,
		Origin: inv.Transform(t.Origin.Neg()),
	}
}

// InterpolateWith returns a transform interpolated between this transform and another by a given
// weight (on the range of 0.0 to 1.0).
func (t Transform3D) InterpolateWith(other Transform3D, weight Float) Transform3D {
	var (
		src_scale = t.Basis.GetScale()
		src_rot   = t.Basis.Quaternion()
		src_loc   = t.Origin
	)
	var (
		dst_scale = other.Basis.GetScale()
		dst_rot   = other.Basis.Quaternion()
		dst_loc   = other.Origin
	)
	return Transform3D{
		Basis:  newBasisWithQuaternionScale(src_rot.Slerp(dst_rot, weight).Normalized(), src_scale.lerp(dst_scale, weight)),
		Origin: src_loc.Lerp(dst_loc, weight),
	}
}

// Inverse returns the inverse of the transform, under the assumption that the transformation basis is
// orthonormal (i.e. rotation/reflection is fine, scaling/skew is not). Use affine_inverse for
// non-orthonormal transforms (e.g. with scaling).
func (t Transform3D) Inverse() Transform3D {
	basis := t.Basis.Transposed()
	return Transform3D{
		Basis:  basis,
		Origin: basis.Transform(t.Origin.Neg()),
	}
}

// IsApproximatelyEqual returns true if this transform and xform are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func (t Transform3D) IsApproximatelyEqual(other Transform3D) bool {
	return t.Basis.IsApproximatelyEqual(other.Basis) && t.Origin.IsApproximatelyEqual(other.Origin)
}

// IsFinite returns true if this transform is finite, by calling [IsFinite] on each component.
func (t Transform3D) IsFinite() bool {
	return t.Basis.IsFinite() && t.Origin.IsFinite()
}

// LookingAt returns a copy of the transform rotated such that the forward axis (-Z) points towards the target position.
//
// The up axis (+Y) points as close to the up vector as possible while staying perpendicular to the forward axis. The
// resulting transform is orthonormalized. The existing rotation, scale, and skew information from the original transform
// is discarded. The target and up vectors cannot be zero, cannot be parallel to each other, and are defined in global/parent space.
//
// If use_model_front is true, the +Z axis (asset front) is treated as forward (implies +X is left) and points toward the target
// position. By default, the -Z axis (camera forward) is treated as forward (implies +X is right).
func (t Transform3D) LookingAt(target, up Vector3, use_model_front bool) Transform3D {
	t.Basis = t.Basis.LookingAt(target, up, use_model_front)
	return t
}

// Orthonormalized returns the transform with the basis orthogonal (90 degrees), and normalized axis vectors (scale of 1 or -1).
func (t Transform3D) Orthonormalized() Transform3D {
	t.Basis = t.Basis.Orthonormalized()
	return t
}

// Rotated returns a copy of the transform rotated around the given axis by the given angle (in radians).
//
// The axis must be a normalized vector.
//
// This method is an optimized version of multiplying the given transform X with a corresponding rotation
// transform R from the left, i.e., R * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func (t Transform3D) Rotated(axis Vector3, phi Radians) Transform3D {
	var basis = NewBasisRotatedAround(axis, phi)
	return Transform3D{
		Basis:  basis.Mul(t.Basis),
		Origin: basis.Transform(t.Origin),
	}
}

// RotatedLocal returns a copy of the transform rotated around the given axis by the given angle
// (in radians).
//
// The axis must be a normalized vector.
//
// This method is an optimized version of multiplying the given transform X with a corresponding
// rotation transform R from the right, i.e., X * R.
//
// This can be seen as transforming with respect to the local frame.
func (t Transform3D) RotatedLocal(axis Vector3, phi Radians) Transform3D {
	var basis = NewBasisRotatedAround(axis, phi)
	return Transform3D{
		Basis:  t.Basis.Mul(basis),
		Origin: t.Origin,
	}
}

// Scaled returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X with a corresponding
// scaling transform S from the left, i.e., S * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func (t Transform3D) Scaled(scale Vector3) Transform3D {
	var basis = NewBasisScaledBy(scale)
	return Transform3D{
		Basis:  basis.Mul(t.Basis),
		Origin: t.Origin.Mul(scale),
	}
}

// ScaledLocal returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding scaling transform S from the right, i.e., X * S.
//
// This can be seen as transforming with respect to the local frame.
func (t Transform3D) ScaledLocal(scale Vector3) Transform3D {
	var basis = NewBasisScaledBy(scale)
	return Transform3D{
		Basis:  t.Basis.Mul(basis),
		Origin: t.Origin,
	}
}

// Translated returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a corresponding
// translation transform T from the left, i.e., T * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func (t Transform3D) Translated(offset Vector3) Transform3D {
	return Transform3D{
		Basis:  t.Basis,
		Origin: t.Origin.Add(offset),
	}
}

// TranslatedLocal returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding translation transform T from the right, i.e., X * T.
//
// This can be seen as transforming with respect to the local frame.
func (t Transform3D) TranslatedLocal(offset Vector3) Transform3D {
	return Transform3D{
		Basis:  t.Basis,
		Origin: t.Origin.Add(t.Basis.Transform(offset)),
	}
}

func (t Transform3D) Mul(other Transform3D) Transform3D {
	return Transform3D{
		Basis:  t.Basis.Mul(other.Basis),
		Origin: t.Origin.Transform(other),
	}
}
