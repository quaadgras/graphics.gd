// Package Transform3D a 3×4 matrix representing a 3D transformation.
package Transform3D

import (
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Vector3"
)

// The BasisOrigin built-in Variant type is a 3×4 matrix representing a
// transformation in 3D space. It contains a Basis, which on its own can
// represent rotation, scale, and shear. Additionally, combined with its
// own origin, the transform can also represent a translation.
//
// For a general introduction, see the Matrices and transforms tutorial.
//
// Note: GD uses a right-handed coordinate system, which is a common
// standard. For directions, the convention for built-in types like
// Camera3D is for -Z to point forward (+X is right, +Y is up, and +Z is
// back). Other objects may use different direction conventions. For more
// information, see the 3D asset direction conventions tutorial.
type BasisOrigin = struct {

	// Basis of this transform. It is composed by 3 axes (Basis.x, Basis.y,
	// and Basis.z). Together, these represent the transform's rotation,
	// scale, and shear.
	Basis Basis.XYZ

	// The translation offset of this transform. In 3D space, this can be
	// seen as the position.
	Origin Vector3.XYZ
}

var (
	// Identity is a transform with no translation, no rotation, and its scale
	// being 1. Its basis is equal to [Basis.Identity].
	//
	// When multiplied by another Variant such as AABB or another BasisOrigin,
	// no transformation occurs.
	Identity = BasisOrigin{
		Basis: Basis.Identity,
	}

	// BasisOrigin with mirroring applied perpendicular to the YZ plane. Its basis
	// is equal to [Basis.FlipX].
	FlipX = BasisOrigin{
		Basis: Basis.FlipX,
	}

	// BasisOrigin with mirroring applied perpendicular to the XZ plane. Its basis
	// is equal to [Basis.FlipY].
	FlipY = BasisOrigin{
		Basis: Basis.FlipY,
	}

	// BasisOrigin with mirroring applied perpendicular to the XY plane. Its basis
	// is equal to [Basis.FlipZ].
	FlipZ = BasisOrigin{
		Basis: Basis.FlipZ,
	}
)

// New constructs a BasisOrigin identical to the [Identity].
func New() BasisOrigin { //gd:BasisOrigin()
	return BasisOrigin{
		Basis: Basis.New(),
	}
}

// AffineInverse returns the inverse of the transform, under the assumption that the basis is
// invertible (must have non-zero determinant).
func AffineInverse(t BasisOrigin) BasisOrigin { //gd:Transform3D.affine_inverse
	var inv = Basis.Inverse(t.Basis)
	return BasisOrigin{
		Basis:  inv,
		Origin: Basis.Transform(Vector3.Neg(t.Origin), inv),
	}
}

// Lerp returns a transform interpolated between this transform and another by a given
// weight (on the range of 0.0 to 1.0).
func Lerp[X Float.Any](t BasisOrigin, other BasisOrigin, weight X) BasisOrigin { //gd:Transform3D.interpolate_with
	var (
		src_scale = Basis.Scale(t.Basis)
		src_rot   = Basis.AsQuaternion(t.Basis)
		src_loc   = t.Origin
	)
	var (
		dst_scale = Basis.Scale(other.Basis)
		dst_rot   = Basis.AsQuaternion(other.Basis)
		dst_loc   = other.Origin
	)
	return BasisOrigin{
		Basis: Basis.RotatesScales(
			Quaternion.Normalized(Quaternion.Slerp(src_rot, dst_rot, weight)),
			Vector3.Lerp(src_scale, dst_scale, weight),
		),
		Origin: Vector3.Lerp(src_loc, dst_loc, weight),
	}
}

// Inverse returns the inverse of the transform, under the assumption that the transformation basis is
// orthonormal (i.e. rotation/reflection is fine, scaling/skew is not). Use affine_inverse for
// non-orthonormal transforms (e.g. with scaling).
func Inverse(t BasisOrigin) BasisOrigin { //gd:Transform3D.inverse
	basis := Basis.Transposed(t.Basis)
	return BasisOrigin{
		Basis:  basis,
		Origin: Basis.Transform(Vector3.Neg(t.Origin), basis),
	}
}

// IsApproximatelyEqual returns true if this transform and xform are approximately equal, by running
// [Vector3.IsApproximatelyEqual] on each component.
func IsApproximatelyEqual(a, b BasisOrigin) bool { //gd:Transform3D.is_equal_approx
	return Basis.IsApproximatelyEqual(a.Basis, b.Basis) && Vector3.IsApproximatelyEqual(a.Origin, b.Origin)
}

// IsFinite returns true if this transform is finite, by calling [IsFinite] on each component.
func IsFinite(t BasisOrigin) bool { //gd:Transform3D.is_finite
	return Basis.IsFinite(t.Basis) && Vector3.IsFinite(t.Origin)
}

// LookingAt returns a copy of the transform rotated such that the forward axis (-Z) points towards the target position.
//
// The up axis (+Y) points as close to the up vector as possible while staying perpendicular to the forward axis. The
// resulting transform is orthonormalized. The existing rotation, scale, and skew information from the original transform
// is discarded. The target and up vectors cannot be zero, cannot be parallel to each other, and are defined in global/parent space.
//
// If use_model_front is true, the +Z axis (asset front) is treated as forward (implies +X is left) and points toward the target
// position. By default, the -Z axis (camera forward) is treated as forward (implies +X is right).
func LookingAt(t BasisOrigin, target, up Vector3.XYZ) BasisOrigin { //gd:Transform3D.looking_at
	t.Basis = Basis.LookingAt(target, up)
	return t
}

// Orthonormalized returns the transform with the basis orthogonal (90 degrees), and normalized axis vectors (scale of 1 or -1).
func Orthonormalized(t BasisOrigin) BasisOrigin { //gd:Transform3D.orthonormalized
	t.Basis = Basis.Orthonormalized(t.Basis)
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
func Rotated(t BasisOrigin, axis Vector3.XYZ, phi Angle.Radians) BasisOrigin { //gd:Transform3D.rotated
	var basis = Basis.RotatesAxisAngle(axis, phi)
	return BasisOrigin{
		Basis:  Basis.Mul(basis, t.Basis),
		Origin: Basis.Transform(t.Origin, basis),
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
func RotatedLocal(t BasisOrigin, axis Vector3.XYZ, phi Angle.Radians) BasisOrigin { //gd:Transform3D.rotated_local
	var basis = Basis.RotatesAxisAngle(axis, phi)
	return BasisOrigin{
		Basis:  Basis.Mul(t.Basis, basis),
		Origin: t.Origin,
	}
}

// Scaled returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X with a corresponding
// scaling transform S from the left, i.e., S * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func Scaled(t BasisOrigin, scale Vector3.XYZ) BasisOrigin { //gd:Transform3D.scaled
	var basis = Basis.Scales(scale)
	return BasisOrigin{
		Basis:  Basis.Mul(basis, t.Basis),
		Origin: Vector3.Mul(t.Origin, scale),
	}
}

// ScaledLocal returns a copy of the transform scaled by the given scale factor.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding scaling transform S from the right, i.e., X * S.
//
// This can be seen as transforming with respect to the local frame.
func ScaledLocal(t BasisOrigin, scale Vector3.XYZ) BasisOrigin { //gd:Transform3D.scaled_local
	var basis = Basis.Scales(scale)
	return BasisOrigin{
		Basis:  Basis.Mul(t.Basis, basis),
		Origin: t.Origin,
	}
}

// Translated returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a corresponding
// translation transform T from the left, i.e., T * X.
//
// This can be seen as transforming with respect to the global/parent frame.
func Translated(t BasisOrigin, offset Vector3.XYZ) BasisOrigin { //gd:Transform3D.translated
	return BasisOrigin{
		Basis:  t.Basis,
		Origin: Vector3.Add(t.Origin, offset),
	}
}

// TranslatedLocal returns a copy of the transform translated by the given offset.
//
// This method is an optimized version of multiplying the given transform X with a
// corresponding translation transform T from the right, i.e., X * T.
//
// This can be seen as transforming with respect to the local frame.
func TranslatedLocal(t BasisOrigin, offset Vector3.XYZ) BasisOrigin { //gd:Transform3D.translated_local
	return BasisOrigin{
		Basis:  t.Basis,
		Origin: Vector3.Add(t.Origin, Basis.Transform(offset, t.Basis)),
	}
}

/*func (t BasisOrigin) Projection() Projection { // Projection(BasisOrigin)
return Projection{
	Vector4{t.Basis[0][X], t.Basis[1][X], t.Basis[2][X], 0},
	Vector4{t.Basis[0][Y], t.Basis[1][Y], t.Basis[2][Y], 0},
	Vector4{t.Basis[0][Z], t.Basis[1][Z], t.Basis[2][Z], 0},
	Vector4{t.Origin[X], t.Origin[Y], t.Origin[Z], 1},
}
}*/

func Mul(a, b BasisOrigin) BasisOrigin { //gd:Transform3D*BasisOrigin
	return BasisOrigin{
		Basis:  Basis.Mul(a.Basis, b.Basis),
		Origin: Transform(a.Origin, b),
	}
}

func Transform(v Vector3.XYZ, t BasisOrigin) Vector3.XYZ { //gd:Transform3D*Vector3
	return Vector3.XYZ{
		Vector3.Dot(t.Basis.X, v) + t.Origin.X,
		Vector3.Dot(t.Basis.Y, v) + t.Origin.Y,
		Vector3.Dot(t.Basis.Z, v) + t.Origin.Z,
	}
}
