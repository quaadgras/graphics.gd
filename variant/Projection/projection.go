// Package Projection provides a 4×4 matrix for 3D projective transformations.
package Projection

import (
	"unsafe"

	"grow.graphics/gd/variant/AABB"
	"grow.graphics/gd/variant/Angle"
	"grow.graphics/gd/variant/Float"
	"grow.graphics/gd/variant/Int"
	"grow.graphics/gd/variant/Plane"
	"grow.graphics/gd/variant/Rect2"
	"grow.graphics/gd/variant/Vector2"
	"grow.graphics/gd/variant/Vector3"
	"grow.graphics/gd/variant/Vector4"
)

// ʕ is a little ternary operator for porting C code.
func ʕ[T any](q bool, a T, b T) T {
	if q {
		return a
	}
	return b
}

// XYZW is a 4×4 matrix used for 3D projective transformations. It can represent
// transformations such as translation, rotation, scaling, shearing, and
// perspective division. It consists of four Vector4 columns.
//
// For purely linear transformations (translation, rotation, and scale), it is
// recommended to use Transform3D, as it is more performant and requires less memory.
type XYZW = struct {
	X Vector4.XYZW // The projection matrix's X vector (column 0). Equivalent to array index 0.
	Y Vector4.XYZW // The projection matrix's Y vector (column 1). Equivalent to array index 1.
	Z Vector4.XYZW // The projection matrix's Z vector (column 2). Equivalent to array index 2.
	W Vector4.XYZW // The projection matrix's W vector (column 3). Equivalent to array index 3.
}

type ClippingPlane int

const (
	PlaneNear ClippingPlane = iota
	PlaneFar
	PlaneLeft
	PlaneTop
	PlaneRight
	PlaneBottom
)

var (
	// Identity projection with no transformation defined. When applied to other data
	// structures, no transformation is performed.
	Identity = XYZW{
		X: Vector4.XYZW{X: 1, Y: 0, Z: 0, W: 0},
		Y: Vector4.XYZW{X: 0, Y: 1, Z: 0, W: 0},
		Z: Vector4.XYZW{X: 0, Y: 0, Z: 1, W: 0},
		W: Vector4.XYZW{X: 0, Y: 0, Z: 0, W: 1},
	}

	// Zero projection with all values initialized to 0. When applied to other data
	// structures, they will be zeroed.
	Zero = XYZW{}
)

// New creates a new projection set to [Identity].
func New() XYZW { return Identity }

// DepthCorrection creates a new Projection that projects positions from a depth range of -1 to 1
// to one that ranges from 0 to 1, and flips the projected positions vertically, according to fliy.
func DepthCorrection(fliy bool) XYZW { //gd:Projection.create_depth_correction
	return XYZW{
		Vector4.XYZW{1, 0, 0, 0},
		Vector4.XYZW{0, ʕ[Float.X](fliy, -1, 1), 0, 0},
		Vector4.XYZW{0, 0, 0.5, 0},
		Vector4.XYZW{0, 0, 0.5, 1},
	}
}

// IntoRect2 creates a new Projection that projects positions into the given Rect2.
func IntoRect2(rect Rect2.PositionSize) XYZW { //gd:Projection.create_light_atlas_rect
	return XYZW{
		Vector4.XYZW{rect.Size.X, 0, 0, 0},
		Vector4.XYZW{0, rect.Size.Y, 0, 0},
		Vector4.XYZW{0, 0, 1, 0},
		Vector4.XYZW{0, rect.Position.X, rect.Position.Y, 1},
	}
}

// FitAABB creates a new Projection that scales a given projection to fit around a given
// AABB in projection space.
func FitAABB(aabb AABB.PositionSize) XYZW { //gd:Projection.create_fit_aabb
	var columns [4][4]Float.X
	var min = aabb.Position
	var max = Vector3.Add(aabb.Position, aabb.Size)
	columns[0][0] = 2 / (max.X - min.X)
	columns[1][0] = 0
	columns[2][0] = 0
	columns[3][0] = -(max.X + min.X) / (max.X - min.X)
	columns[0][1] = 0
	columns[1][1] = 2 / (max.Y - min.Y)
	columns[2][1] = 0
	columns[3][1] = -(max.Y + min.Y) / (max.Y - min.Y)
	columns[0][2] = 0
	columns[1][2] = 0
	columns[2][2] = 2 / (max.Z - min.Z)
	columns[3][2] = -(max.Z + min.Z) / (max.Z - min.Z)
	columns[0][3] = 0
	columns[1][3] = 0
	columns[2][3] = 0
	columns[3][3] = 1
	return *(*XYZW)(unsafe.Pointer(&columns))
}

// HeadMountedDisplay creates a new Projection for projecting positions onto a head-mounted display
// with the given X:Y aspect ratio, distance between eyes, display width, distance to lens, oversampling
// factor, and depth clipping planes.
//
// eye creates the projection for the left eye when set to 1, or the right eye when set to 2.
func HeadMountedDisplay[X Float.Any, I Int.Any](eye I, aspect, intraocular_dist, display_width, display_to_lens, oversample, z_near, z_far Float.X) XYZW { //gd:Projection.create_for_hmd
	// we first calculate our base frustum on our values without taking our lens magnification into account.
	var (
		f1 = (intraocular_dist * 0.5) / display_to_lens
		f2 = ((display_width - intraocular_dist) * 0.5) / display_to_lens
		f3 = (display_width / 4.0) / display_to_lens
	)
	// now we apply our oversample factor to increase our FOV. how much we oversample is always a balance we strike between performance and how much
	// we're willing to sacrifice in FOV.
	var add = ((f1 + f2) * (oversample - 1.0)) / 2.0
	f1 += add
	f2 += add
	f3 *= oversample
	// always apply KEEWIDTH aspect ratio
	f3 /= aspect
	switch eye {
	case 1: // left eye
		return Frustum(-f2*z_near, f1*z_near, -f3*z_near, f3*z_near, z_near, z_far)
	case 2: // right eye
		return Frustum(-f1*z_near, f2*z_near, -f3*z_near, f3*z_near, z_near, z_far)
	}
	panic("NewProjectionForHeadMountedDisplay(): invalid eye")
}

// Frustum creates a new Projection that projects positions in a frustum with the given clipping planes.
func Frustum[X Float.Any](left, right, bottom, top, z_near, z_far X) XYZW { //gd:Projection.create_frustum
	return XYZW{
		Vector4.XYZW{Float.X(2 * z_near / (right - left)), 0, Float.X((right + left) / (right - left)), 0},
		Vector4.XYZW{0, Float.X(2 * z_near / (top - bottom)), Float.X((top + bottom) / (top - bottom)), 0},
		Vector4.XYZW{0, 0, Float.X(-(z_far + z_near) / (z_far - z_near)), Float.X(-2 * z_far * z_near / (z_far - z_near))},
		Vector4.XYZW{0, 0, -1, 0},
	}
}

// FrustumAspectRatio creates a new Projection that projects positions in a frustum with the given size,
// X:Y aspect ratio, offset, and clipping planes.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal.
func FrustumAspectRatio[X Float.Any](size, aspect X, offset Vector2.XY, z_near, z_far X, flip_fov bool) XYZW { //gd:Projection.create_frustum_aspect
	if !flip_fov {
		size *= aspect
	}
	return Frustum(-size/2+X(offset.X), size/2+X(offset.X), -size/aspect/2+X(offset.Y), size/aspect/2+X(offset.Y), z_near, z_far)
}

// Orthogonal creates a new Projection that projects positions using an orthogonal projection with the given
// clipping planes.
func Orthogonal[X Float.Any](left, right, bottom, top, z_near, z_far X) XYZW { //gd:Projection.create_orthogonal
	return XYZW{
		Vector4.XYZW{Float.X(2 / (right - left)), 0, 0, 0},
		Vector4.XYZW{0, Float.X(2 / (top - bottom)), 0, 0},
		Vector4.XYZW{0, 0, Float.X(-2 / (z_far - z_near)), 0},
		Vector4.XYZW{Float.X(-(right + left) / (right - left)), Float.X(-(top + bottom) / (top - bottom)), Float.X(-(z_far + z_near) / (z_far - z_near)), 1},
	}
}

// OrthogonalAspectRatio creates a new Projection that projects positions using an orthogonal projection with
// the given size, X:Y aspect ratio, and clipping planes.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal.
func OrthogonalAspectRatio[X Float.Any](size, aspect, z_near, z_far X, flip_fov bool) XYZW { //gd:Projection.create_orthogonal_aspect
	if !flip_fov {
		size *= aspect
	}
	return Orthogonal(-size/2, size/2, -size/aspect/2, size/aspect/2, z_near, z_far)
}

// Perspective creates a new Projection that projects positions using a perspective projection with the given Y-axis
// field of view (in degrees), X:Y aspect ratio, and clipping planes.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal.
func Perspective[X Float.Any](fovy Angle.Degrees, aspect, z_near, z_far X, flip_fov bool) XYZW { //gd:Projection.create_perspective
	if flip_fov {
		fovy = Fovy(fovy, 1/aspect)
	}
	var (
		sine, cotangent Float.X
		deltaZ          Float.X
		radians         = Angle.InRadians(fovy / 2)
	)
	deltaZ = Float.X(z_far - z_near)
	sine = Angle.Sin(radians)
	if (deltaZ == 0) || (sine == 0) || (aspect == 0) {
		return Zero
	}
	cotangent = Angle.Cos(radians) / sine
	return XYZW{
		Vector4.XYZW{Float.X(cotangent) / Float.X(aspect), 0, 0, 0},
		Vector4.XYZW{0, Float.X(cotangent), 0, 0},
		Vector4.XYZW{0, 0, Float.X(-Float.X(z_far+z_near) / deltaZ), Float.X(-Float.X(2*z_near*z_far) / deltaZ)},
		Vector4.XYZW{0, 0, -1, 1},
	}
}

// HeadMountedDisplayPerspective creates a new Projection that projects positions using a perspective
// projection with the given Y-axis field of view (in degrees), X:Y aspect ratio, and clipping distances.
// The projection is adjusted for a head-mounted display with the given distance between eyes and
// distance to a point that can be focused on.
//
// eye creates the projection for the left eye when set to 1, or the right eye when set to 2.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal
func HeadMountedDisplayPerspective(fovy Angle.Degrees, aspect, z_near, z_far Float.X, flip_fov bool, eye int, intraocular_dist, convergence_dist Float.X) XYZW { //gd:Projection.create_perspective_hmd
	if flip_fov {
		fovy = Fovy(fovy, 1.0/aspect)
	}
	var (
		left, right, ymax, xmax, frustumshift Float.X
		modeltranslation                      Float.X
	)
	ymax = z_near * Angle.Tan(Angle.InRadians(fovy/2.0))
	xmax = ymax * Float.X(aspect)
	frustumshift = Float.X((intraocular_dist / 2.0) * z_near / convergence_dist)
	switch eye {
	case 1: // left eye
		left = -xmax + frustumshift
		right = xmax + frustumshift
		modeltranslation = Float.X(intraocular_dist / 2.0)
	case 2: // right eye
		left = -xmax - frustumshift
		right = xmax - frustumshift
		modeltranslation = -Float.X(intraocular_dist / 2.0)
	default: // mono, should give the same result as set_perspective(p_fovy_degrees,p_aspect,p_z_near,p_z_far,p_flip_fov)
		left = -xmax
		right = xmax
		modeltranslation = 0.0
	}
	var (
		proj = Frustum(left, right, -ymax, ymax, z_near, z_far)
	)
	// translate matrix by (modeltranslation, 0.0, 0.0)
	var cm = New()
	cm.Z.X = modeltranslation
	return Mul(proj, cm)
}

// Fovy returns the vertical field of view of the projection (in degrees) associated with the given
// horizontal field of view (in degrees) and aspect ratio.
func Fovy[X Float.Any](hfov Angle.Degrees, aspect X) Angle.Degrees { //gd:Projection.get_fovy
	return Angle.InDegrees(2 * Angle.Atan(Float.X(aspect)*Angle.Tan(Angle.InRadians(hfov)/2)))
}

// Determinant returns a scalar value that is the signed factor by which areas are scaled by this matrix. If the
// sign is negative, the matrix flips the orientation of the area.
//
// The determinant can be used to calculate the invertibility of a matrix or solve linear systems of equations
// involving the matrix, among other applications.
func Determinant(p XYZW) Float.X { //gd:Projection.determinant
	return p.X.W*p.Y.Z*p.Z.Y*p.W.X - p.X.Z*p.Y.W*p.Z.Y*p.W.X -
		p.X.W*p.Y.Y*p.Z.Z*p.W.X + p.X.Y*p.Y.W*p.Z.Z*p.W.X +
		p.X.Z*p.Y.Y*p.Z.W*p.W.X - p.X.Y*p.Y.Z*p.Z.W*p.W.X -
		p.X.W*p.Y.Z*p.Z.X*p.W.Y + p.X.Z*p.Y.W*p.Z.X*p.W.Y +
		p.X.W*p.Y.X*p.Z.Z*p.W.Y - p.X.X*p.Y.W*p.Z.Z*p.W.Y -
		p.X.Z*p.Y.X*p.Z.W*p.W.Y + p.X.X*p.Y.Z*p.Z.W*p.W.Y +
		p.X.W*p.Y.Y*p.Z.X*p.W.Z - p.X.Y*p.Y.W*p.Z.X*p.W.Z -
		p.X.W*p.Y.X*p.Z.Y*p.W.Z + p.X.X*p.Y.W*p.Z.Y*p.W.Z +
		p.X.Y*p.Y.X*p.Z.W*p.W.Z - p.X.X*p.Y.Y*p.Z.W*p.W.Z -
		p.X.Z*p.Y.Y*p.Z.X*p.W.W + p.X.Y*p.Y.Z*p.Z.X*p.W.W +
		p.X.Z*p.Y.X*p.Z.Y*p.W.W - p.X.X*p.Y.Z*p.Z.Y*p.W.W -
		p.X.Y*p.Y.X*p.Z.Z*p.W.W + p.X.X*p.Y.Y*p.Z.Z*p.W.W
}

// FlippedY returns a copy of this Projection with the signs of the values of the Y column flipped.
func FlippedY(p XYZW) XYZW { //gd:Projection.flipped_y
	p.Y = Vector4.Neg(p.Y)
	return p
}

// AspectRatio returns the X:Y aspect ratio of this Projection's viewport.
func AspectRatio(p XYZW) Float.X { //gd:Projection.get_aspect
	vp_he := ViewportHalfExtents(p)
	return vp_he.X / vp_he.Y
}

// FarPlaneHalfExtents returns the dimensions of the far clipping plane of the projection, divided by two.
func FarPlaneHalfExtents(p XYZW) Vector2.XY { //gd:Projection.get_far_plane_half_extents
	m := (*[16]Float.X)(unsafe.Pointer(&p))
	var far_plane = Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[2],
			m[7] - m[6],
			m[11] - m[10],
		},
		D: -m[15] + m[14],
	})
	var right_plane = Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: -m[15] + m[12],
	})
	var top_plane = Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[1],
			m[7] - m[5],
			m[11] - m[9],
		},
		D: -m[15] + m[13],
	})
	res, _ := Plane.Intersect3(far_plane, right_plane, top_plane)
	return Vector2.XY{res.X, res.Y}
}

// FieldOfView returns the horizontal field of view of the projection (in degrees).
func FieldOfView(p XYZW) Angle.Degrees { //gd:Projection.get_fov
	m := (*[16]Float.X)(unsafe.Pointer(&p))
	var right_plane = Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: -m[15] + m[12],
	})
	if (m[8] == 0) && (m[9] == 0) {
		return Angle.InDegrees(Angle.Acos(Float.Abs(right_plane.Normal.X))) * 2.0
	}
	// our frustum is asymmetrical need to calculate the left planes angle separately..
	var left_plane = Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: m[15] + m[12],
	})
	return Angle.InDegrees(Angle.Acos(Float.Abs(left_plane.Normal.X))) + Angle.InDegrees(Angle.Acos(Float.Abs(right_plane.Normal.X)))
}

// LevelOfDetailMultiplier returns the factor by which the visible level of detail is scaled by
// this Projection.
func LevelOfDetailMultiplier(p XYZW) Float.X { //gd:Projection.get_lod_multiplier
	if IsOrthogonal(p) {
		return ViewportHalfExtents(p).X
	}
	var zn = NearZ(p)
	var width = ViewportHalfExtents(p).X * 2.0
	return 1.0 / (zn / width)
}

// PixelsPerMeter returns the number of pixels with the given pixel width displayed per meter,
// after this Projection is applied.
func PixelsPerMeter(p XYZW, pixel_width int) int { //gd:Projection.get_pixels_per_meter
	var result = Transform(Vector4.XYZW{1, 0, -1, 1}, p)
	return int((result.X*0.5 + 0.5) * Float.X(pixel_width))
}

// For returns the clipping plane of this Projection whose index is given by plane.
//
// plane should be equal to one of PLANE_NEAR, PLANE_FAR, PLANE_LEFT, PLANE_TOP, PLANE_RIGHT, or
// PLANE_BOTTOM.
func (index ClippingPlane) For(p XYZW) (new_plane Plane.NormalD) { //gd:Projection.get_projection_plane
	m := (*[16]Float.X)(unsafe.Pointer(&p))
	switch index {
	case PlaneNear:
		new_plane = Plane.NormalD{
			Normal: Vector3.XYZ{
				m[3] + m[2],
				m[7] + m[6],
				m[11] + m[10],
			},
			D: m[15] + m[14],
		}
	case PlaneFar:
		new_plane = Plane.NormalD{
			Normal: Vector3.XYZ{
				m[3] - m[2],
				m[7] - m[6],
				m[11] - m[10],
			},
			D: m[15] + m[14],
		}
	case PlaneLeft:
		new_plane = Plane.NormalD{
			Normal: Vector3.XYZ{
				m[3] + m[0],
				m[7] + m[4],
				m[11] + m[8],
			},
			D: m[15] + m[12],
		}
	case PlaneTop:
		new_plane = Plane.NormalD{
			Normal: Vector3.XYZ{
				m[3] - m[1],
				m[7] - m[5],
				m[11] - m[9],
			},
			D: m[15] + m[13],
		}
	case PlaneRight:
		new_plane = Plane.NormalD{
			Normal: Vector3.XYZ{
				m[3] - m[0],
				m[7] - m[4],
				m[11] - m[8],
			},
			D: m[15] + m[12],
		}
	case PlaneBottom:
		new_plane = Plane.NormalD{
			Normal: Vector3.XYZ{
				m[3] + m[1],
				m[7] + m[5],
				m[11] + m[9],
			},
			D: m[15] + m[13],
		}
	}
	new_plane.Normal = Vector3.Neg(new_plane.Normal)
	return Plane.Normalized(new_plane)
}

// ViewportHalfExtents returns the dimensions of the viewport plane that this Projection projects positions onto,
// divided by two.
func ViewportHalfExtents(p XYZW) Vector2.XY { //gd:Projection.get_viewport_half_extents
	m := (*[16]Float.X)(unsafe.Pointer(&p))
	near_plane := Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] + m[2],
			m[7] + m[6],
			m[11] + m[10],
		},
		D: -m[15] + m[14],
	})
	right_plane := Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: m[15] + m[12],
	})
	top_plane := Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[1],
			m[7] - m[5],
			m[11] - m[9],
		},
		D: -m[15] + m[13],
	})
	res, _ := Plane.Intersect3(near_plane, right_plane, top_plane)
	return Vector2.XY{res.X, res.Y}
}

// FarZ returns the distance for this Projection beyond which positions are clipped.
func FarZ(p XYZW) Float.X { //gd:Projection.get_z_far
	m := (*[16]Float.X)(unsafe.Pointer(&p))
	var new_plane = Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] - m[2],
			m[7] - m[6],
			m[11] - m[10],
		},
		D: m[15] - m[14],
	})
	return new_plane.D
}

// NearZ returns the distance for this Projection before which positions are clipped.
func NearZ(p XYZW) Float.X { //gd:Projection.get_z_near
	m := (*[16]Float.X)(unsafe.Pointer(&p))
	var new_plane = Plane.Normalized(Plane.NormalD{
		Normal: Vector3.XYZ{
			m[3] + m[2],
			m[7] + m[6],
			m[11] + m[10],
		},
		D: -m[15] - m[14],
	})
	return new_plane.D
}

// Inverse returns a Projection that performs the inverse of this Projection's projective transformation.
func Inverse(proj XYZW) XYZW { //gd:Projection.inverse
	p := (*[4][4]Float.X)(unsafe.Pointer(&proj))
	var i, j int
	var pvt_i, pvt_j [4]int /* Locations of pivot matrix */
	var pvt_val Float.X     /* Value of current pivot element */
	var hold Float.X        /* Temporary storage */
	var determinant Float.X = 1.0
	for k := 0; k < 4; k++ {
		/** Locate k'th pivot element **/
		pvt_val = p[k][k] /** Initialize for search **/
		pvt_i[k] = k
		pvt_j[k] = k
		for i := k; i < 4; i++ {
			for j := k; j < 4; j++ {
				if Float.Abs(p[i][j]) > Float.Abs(pvt_val) {
					pvt_i[k] = i
					pvt_j[k] = j
					pvt_val = p[i][j]
				}
			}
		}
		/** Product of pivots, gives determinant when finished **/
		determinant *= pvt_val
		if Float.IsApproximatelyZero(determinant) {
			return Zero /** Matrix is singular (zero determinant). **/
		}
		/** "Interchange" rows (with sign change stuff) **/
		i = pvt_i[k]
		if i != k { /** If rows are different **/
			for j := 0; j < 4; j++ {
				hold = -p[k][j]
				p[k][j] = p[i][j]
				p[i][j] = hold
			}
		}
		/** "Interchange" columns **/
		j = pvt_j[k]
		if j != k { /** If columns are different **/
			for i := 0; i < 4; i++ {
				hold = -p[i][k]
				p[i][k] = p[i][j]
				p[i][j] = hold
			}
		}
		/** Divide column by minus pivot value **/
		for i := 0; i < 4; i++ {
			if i != k {
				p[i][k] /= (-pvt_val)
			}
		}
		/** Reduce the matrix **/
		for i := 0; i < 4; i++ {
			hold = p[i][k]
			for j := 0; j < 4; j++ {
				if i != k && j != k {
					p[i][j] += hold * p[k][j]
				}
			}
		}
		/** Divide row by pivot **/
		for j := 0; j < 4; j++ {
			if j != k {
				p[k][j] /= pvt_val
			}
		}
		/** Replace pivot by reciprocal (at last we can touch it). **/
		p[k][k] = 1.0 / pvt_val
	}
	/* That was most of the work, one final pass of row/column interchange */
	/* to finish */
	for k := 4 - 2; k >= 0; k-- { /* Don't need to work with 1 by 1 corner*/
		i = pvt_j[k] /* Rows to swap correspond to pivot COLUMN */
		if i != k {  /* If rows are different */
			for j := 0; j < 4; j++ {
				hold = p[k][j]
				p[k][j] = -p[i][j]
				p[i][j] = hold
			}
		}
		j = pvt_i[k] /* Columns to swap correspond to pivot ROW */
		if j != k {  /* If columns are different */
			for i := 0; i < 4; i++ {
				hold = p[i][k]
				p[i][k] = -p[i][j]
				p[i][j] = hold
			}
		}
	}
	return proj
}

// IsOrthogonal returns true if this Projection performs an orthogonal projection.
func IsOrthogonal(p XYZW) bool { return p.Z.Z == 1.0 } //gd:Projection.is_orthogonal

// JitterOffseted returns a Projection with the X and Y values from the given Vector2
// added to the first and second values of the final column respectively.
func JitterOffseted(p XYZW, jitter Vector2.XY) XYZW { //gd:Projection.jitter_offseted
	p.Z.X += jitter.X
	p.Z.Y += jitter.Y
	return p
}

// PerspectiveAdjustedNearZ returns a Projection with the near clipping distance adjusted to be new_znear.
//
// Note: The original Projection must be a perspective projection.
func PerspectiveAdjustedNearZ(p XYZW, new_near_z Float.X) XYZW { //gd:Projection.perspective_znear_adjusted
	var zfar = FarZ(p)
	var znear = new_near_z
	var deltaZ = zfar - znear
	p.Z.Z = -(zfar + znear) / deltaZ
	p.W.Y = -2 * znear * zfar / deltaZ
	return p
}

func Mul(a, b XYZW) XYZW { //Projection*Projection
	var result XYZW
	new_matrix := (*[4][4]Float.X)(unsafe.Pointer(&a))
	p := (*[4][4]Float.X)(unsafe.Pointer(&a))
	other := (*[4][4]Float.X)(unsafe.Pointer(&b))
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			var ab Float.X = 0
			for k := 0; k < 4; k++ {
				ab += p[k][i] * other[j][k]
			}
			new_matrix[j][i] = ab
		}
	}
	return result
}

// Transform transforms (multiplies) the [Vector4.XYZW] by the given projection's transformation matrix.
func Transform(v Vector4.XYZW, p XYZW) Vector4.XYZW { //Projection*Vector4
	return Vector4.XYZW{
		p.X.X*v.X + p.Y.X*v.Y + p.Z.X*v.Z + p.W.X*v.W,
		p.X.Y*v.X + p.Y.Y*v.Y + p.Z.Y*v.Z + p.W.Y*v.W,
		p.X.Z*v.X + p.Y.Z*v.Y + p.Z.Z*v.Z + p.W.Z*v.W,
		p.X.W*v.X + p.Y.W*v.Y + p.Z.W*v.Z + p.W.W*v.W,
	}
}
