package spatial

import "unsafe"

/*
Projection is a 4x4 matrix used for 3D projective transformations. It can represent transformations such as translation,
rotation, scaling, shearing, and perspective division. It consists of four Vector4 columns.

For purely linear transformations (translation, rotation, and scale), it is recommended to use Transform3D, as it is
more performant and requires less memory.

Generally used for a camera's projection matrix.
*/
type Projection [4]Vector4

type ProjectionPlane int

const (
	PlaneNear ProjectionPlane = iota
	PlaneFar
	PlaneLeft
	PlaneTop
	PlaneRight
	PlaneBottom
)

// "Fields"

func (p Projection) X() Vector4 { return p[0] }
func (p Projection) Y() Vector4 { return p[1] }
func (p Projection) Z() Vector4 { return p[2] }
func (p Projection) W() Vector4 { return p[3] }

// "Constants"

func (Projection) IDENTITY() Projection {
	return Projection{
		Vector4{1, 0, 0, 0},
		Vector4{0, 1, 0, 0},
		Vector4{0, 0, 1, 0},
		Vector4{0, 0, 0, 1},
	}
}

// "Methods"

// NewProjectionWithDepthCorrection creates a new Projection that projects positions from a depth range of -1 to 1
// to one that ranges from 0 to 1, and flips the projected positions vertically, according to fliy.
func NewProjectionWithDepthCorrection(fliy bool) Projection {
	return Projection{
		Vector4{1, 0, 0, 0},
		Vector4{0, ʕ[float](fliy, -1, 1), 0, 0},
		Vector4{0, 0, 0.5, 0},
		Vector4{0, 0, 0.5, 1},
	}
}

// NewProjectionForHeadMountedDisplay creates a new Projection for projecting positions onto a head-mounted display
// with the given X:Y aspect ratio, distance between eyes, display width, distance to lens, oversampling factor, and
// depth clipping planes.
//
// eye creates the projection for the left eye when set to 1, or the right eye when set to 2.
func NewProjectionForHeadMountedDisplay(eye int, aspect, intraocular_dist, display_width, display_to_lens, oversample, z_near, z_far Float) Projection {
	// we first calculate our base frustum on our values without taking our lens magnification into account.
	var f1 = (intraocular_dist * 0.5) / display_to_lens
	var f2 = ((display_width - intraocular_dist) * 0.5) / display_to_lens
	var f3 = (display_width / 4.0) / display_to_lens

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
		return NewProjectionWithFrustum(-f2*z_near, f1*z_near, -f3*z_near, f3*z_near, z_near, z_far)
	case 2: // right eye
		return NewProjectionWithFrustum(-f1*z_near, f2*z_near, -f3*z_near, f3*z_near, z_near, z_far)
	}
	panic("NewProjectionForHeadMountedDisplay(): invalid eye")
}

// NewProjectionWithFrustum creates a new Projection that projects positions in a frustum with the given clipping planes.
func NewProjectionWithFrustum(left, right, bottom, top, z_near, z_far Float) Projection {
	return Projection{
		Vector4{float(2 * z_near / (right - left)), 0, float((right + left) / (right - left)), 0},
		Vector4{0, float(2 * z_near / (top - bottom)), float((top + bottom) / (top - bottom)), 0},
		Vector4{0, 0, float(-(z_far + z_near) / (z_far - z_near)), float(-2 * z_far * z_near / (z_far - z_near))},
		Vector4{0, 0, -1, 0},
	}
}

// NewProjectionWithFrustumAspectRatio creates a new Projection that projects positions in a frustum with the given size,
// X:Y aspect ratio, offset, and clipping planes.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal.
func NewProjectionWithFrustumAspectRatio(size, aspect Float, offset Vector2, z_near, z_far Float, flip_fov bool) Projection {
	if !flip_fov {
		size *= aspect
	}
	return NewProjectionWithFrustum(-size/2+Float(offset[X]), size/2+Float(offset[X]), -size/aspect/2+Float(offset[Y]), size/aspect/2+Float(offset[Y]), z_near, z_far)
}

// NewProjectionThatIsOrthogonal creates a new Projection that projects positions using an orthogonal projection with the given
// clipping planes.
func NewProjectionThatIsOrthogonal(left, right, bottom, top, z_near, z_far Float) Projection {
	return Projection{
		Vector4{float(2 / (right - left)), 0, 0, 0},
		Vector4{0, float(2 / (top - bottom)), 0, 0},
		Vector4{0, 0, float(-2 / (z_far - z_near)), 0},
		Vector4{float(-(right + left) / (right - left)), float(-(top + bottom) / (top - bottom)), float(-(z_far + z_near) / (z_far - z_near)), 1},
	}
}

// NewProjectionThatIsOrthogonalWithAspectRatio creates a new Projection that projects positions using an orthogonal projection with
// the given size, X:Y aspect ratio, and clipping planes.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal.
func NewProjectionThatIsOrthogonalWithAspectRatio(size, aspect Float, z_near, z_far Float, flip_fov bool) Projection {
	if !flip_fov {
		size *= aspect
	}
	return NewProjectionThatIsOrthogonal(-size/2, size/2, -size/aspect/2, size/aspect/2, z_near, z_far)
}

// NewProjectionWithPerspective creates a new Projection that projects positions using a perspective projection with the given Y-axis
// field of view (in degrees), X:Y aspect ratio, and clipping planes.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal.
func NewProjectionWithPerspective(fovy Degrees, aspect, z_near, z_far Float, flip_fov bool) Projection {
	if flip_fov {
		fovy = ProjectionFovy(fovy, 1/aspect)
	}
	var sine, cotangent Radians
	var deltaZ float
	var radians = (fovy / 2).Radians()

	deltaZ = float(z_far - z_near)
	sine = Sin(radians)

	if (deltaZ == 0) || (sine == 0) || (aspect == 0) {
		return Projection{}
	}
	cotangent = Cos(radians) / sine

	return Projection{
		Vector4{float(cotangent) / float(aspect), 0, 0, 0},
		Vector4{0, float(cotangent), 0, 0},
		Vector4{0, 0, float(-float(z_far+z_near) / deltaZ), float(-float(2*z_near*z_far) / deltaZ)},
		Vector4{0, 0, -1, 1},
	}
}

// NewProjectionWithPerspectiveForHeadMountedDisplay creates a new Projection that projects positions using a perspective projection
// with the given Y-axis field of view (in degrees), X:Y aspect ratio, and clipping distances. The projection is adjusted for a
// head-mounted display with the given distance between eyes and distance to a point that can be focused on.
//
// eye creates the projection for the left eye when set to 1, or the right eye when set to 2.
//
// flip_fov determines whether the projection's field of view is flipped over its diagonal
func NewProjectionWithPerspectiveForHeadMountedDisplay(fovy Degrees, aspect, z_near, z_far Float, flip_fov bool, eye int, intraocular_dist, convergence_dist Float) Projection {
	if flip_fov {
		fovy = ProjectionFovy(fovy, 1.0/aspect)
	}

	var left, right, ymax, xmax, frustumshift Float
	var modeltranslation float

	ymax = z_near * Float(Tan((fovy / 2.0).Radians()))
	xmax = ymax * aspect
	frustumshift = (intraocular_dist / 2.0) * z_near / convergence_dist

	switch eye {
	case 1: // left eye
		left = -xmax + frustumshift
		right = xmax + frustumshift
		modeltranslation = float(intraocular_dist / 2.0)
	case 2: // right eye
		left = -xmax - frustumshift
		right = xmax - frustumshift
		modeltranslation = -float(intraocular_dist / 2.0)
	default: // mono, should give the same result as set_perspective(p_fovy_degrees,p_aspect,p_z_near,p_z_far,p_flip_fov)
		left = -xmax
		right = xmax
		modeltranslation = 0.0
	}

	var proj = NewProjectionWithFrustum(left, right, -ymax, ymax, z_near, z_far)

	// translate matrix by (modeltranslation, 0.0, 0.0)
	var cm Projection
	cm = cm.IDENTITY()
	cm[3][0] = modeltranslation
	return proj.Mul(cm)
}

// ProjectionFovy returns the vertical field of view of the projection (in degrees) associated with the given
// horizontal field of view (in degrees) and aspect ratio.
func ProjectionFovy(hfov Degrees, aspect Float) Degrees {
	return Degrees(2 * Atan(Radians(aspect)*Tan(hfov.Radians()/2)).Degrees())
}

// Determinant returns a scalar value that is the signed factor by which areas are scaled by this matrix. If the
// sign is negative, the matrix flips the orientation of the area.
//
// The determinant can be used to calculate the invertibility of a matrix or solve linear systems of equations
// involving the matrix, among other applications.
func (p Projection) Determinant() Float {
	return Float(p[0][3]*p[1][2]*p[2][1]*p[3][0] - p[0][2]*p[1][3]*p[2][1]*p[3][0] -
		p[0][3]*p[1][1]*p[2][2]*p[3][0] + p[0][1]*p[1][3]*p[2][2]*p[3][0] +
		p[0][2]*p[1][1]*p[2][3]*p[3][0] - p[0][1]*p[1][2]*p[2][3]*p[3][0] -
		p[0][3]*p[1][2]*p[2][0]*p[3][1] + p[0][2]*p[1][3]*p[2][0]*p[3][1] +
		p[0][3]*p[1][0]*p[2][2]*p[3][1] - p[0][0]*p[1][3]*p[2][2]*p[3][1] -
		p[0][2]*p[1][0]*p[2][3]*p[3][1] + p[0][0]*p[1][2]*p[2][3]*p[3][1] +
		p[0][3]*p[1][1]*p[2][0]*p[3][2] - p[0][1]*p[1][3]*p[2][0]*p[3][2] -
		p[0][3]*p[1][0]*p[2][1]*p[3][2] + p[0][0]*p[1][3]*p[2][1]*p[3][2] +
		p[0][1]*p[1][0]*p[2][3]*p[3][2] - p[0][0]*p[1][1]*p[2][3]*p[3][2] -
		p[0][2]*p[1][1]*p[2][0]*p[3][3] + p[0][1]*p[1][2]*p[2][0]*p[3][3] +
		p[0][2]*p[1][0]*p[2][1]*p[3][3] - p[0][0]*p[1][2]*p[2][1]*p[3][3] -
		p[0][1]*p[1][0]*p[2][2]*p[3][3] + p[0][0]*p[1][1]*p[2][2]*p[3][3])
}

// FlippedY returns a copy of this Projection with the signs of the values of the Y column flipped.
func (p Projection) FlippedY() Projection {
	for i := range p[0] {
		p[1][i] = -p[1][i]
	}
	return p
}

// AspectRatio returns the X:Y aspect ratio of this Projection's viewport.
func (p Projection) AspectRatio() Float {
	vp_he := p.ViewportHalfExtents()
	return Float(vp_he[X] / vp_he[Y])
}

// FarPlaneHalfExtents returns the dimensions of the far clipping plane of the projection, divided by two.
func (p Projection) FarPlaneHalfExtents() Vector2 {
	m := (*[16]float)(unsafe.Pointer(&p[0][0]))
	var far_plane = Plane{
		Normal: Vector3{
			m[3] - m[2],
			m[7] - m[6],
			m[11] - m[10],
		},
		D: -m[15] + m[14],
	}.Normalized()
	var right_plane = Plane{
		Normal: Vector3{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: -m[15] + m[12],
	}.Normalized()
	var top_plane = Plane{
		Normal: Vector3{
			m[3] - m[1],
			m[7] - m[5],
			m[11] - m[9],
		},
		D: -m[15] + m[13],
	}.Normalized()
	res, _ := far_plane.Intersect3(right_plane, top_plane)
	return Vector2{res[X], res[Y]}
}

// FieldOfView returns the horizontal field of view of the projection (in degrees).
func (p Projection) FieldOfView() Degrees {
	m := (*[16]float)(unsafe.Pointer(&p[0][0]))
	var right_plane = Plane{
		Normal: Vector3{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: -m[15] + m[12],
	}.Normalized()
	if (m[8] == 0) && (m[9] == 0) {
		return Acos(Absf(right_plane.Normal[X])).Degrees() * 2.0
	}
	// our frustum is asymmetrical need to calculate the left planes angle separately..
	var left_plane = Plane{
		Normal: Vector3{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: m[15] + m[12],
	}.Normalized()
	return Acos(Absf(left_plane.Normal[X])).Degrees() + Acos(Absf(right_plane.Normal[X])).Degrees()
}

// LevelOfDetailMultiplier returns the factor by which the visible level of detail is scaled by
// this Projection.
func (p Projection) LevelOfDetailMultiplier() Float {
	if p.IsOrthogonal() {
		return Float(p.ViewportHalfExtents()[X])
	}
	var zn = p.NearZ()
	var width = Float(p.ViewportHalfExtents()[X]) * 2.0
	return 1.0 / (zn / width)
}

// PixelsPerMeter returns the number of pixels with the given pixel width displayed per meter,
// after this Projection is applied.
func (p Projection) PixelsPerMeter(pixel_width Int) Int {
	var result = Vector4{1, 0, -1, 1}.Transform(p)
	return Int((result[X]*0.5 + 0.5) * float(pixel_width))
}

// ProjectionPlane returns the clipping plane of this Projection whose index is given by plane.
//
// plane should be equal to one of PLANE_NEAR, PLANE_FAR, PLANE_LEFT, PLANE_TOP, PLANE_RIGHT, or
// PLANE_BOTTOM.
func (p Projection) ProjectionPlane(plane ProjectionPlane) (new_plane Plane) {
	m := (*[16]float)(unsafe.Pointer(&p[0][0]))
	switch plane {
	case PlaneNear:
		new_plane = Plane{
			Normal: Vector3{
				m[3] + m[2],
				m[7] + m[6],
				m[11] + m[10],
			},
			D: m[15] + m[14],
		}
	case PlaneFar:
		new_plane = Plane{
			Normal: Vector3{
				m[3] - m[2],
				m[7] - m[6],
				m[11] - m[10],
			},
			D: m[15] + m[14],
		}
	case PlaneLeft:
		new_plane = Plane{
			Normal: Vector3{
				m[3] + m[0],
				m[7] + m[4],
				m[11] + m[8],
			},
			D: m[15] + m[12],
		}
	case PlaneTop:
		new_plane = Plane{
			Normal: Vector3{
				m[3] - m[1],
				m[7] - m[5],
				m[11] - m[9],
			},
			D: m[15] + m[13],
		}
	case PlaneRight:
		new_plane = Plane{
			Normal: Vector3{
				m[3] - m[0],
				m[7] - m[4],
				m[11] - m[8],
			},
			D: m[15] + m[12],
		}
	case PlaneBottom:
		new_plane = Plane{
			Normal: Vector3{
				m[3] + m[1],
				m[7] + m[5],
				m[11] + m[9],
			},
			D: m[15] + m[13],
		}
	}
	new_plane.Normal = new_plane.Normal.Neg()
	return new_plane.Normalized()
}

// ViewportHalfExtents returns the dimensions of the viewport plane that this Projection projects positions onto,
// divided by two.
func (p Projection) ViewportHalfExtents() Vector2 {
	m := (*[16]float)(unsafe.Pointer(&p[0][0]))
	near_plane := Plane{
		Normal: Vector3{
			m[3] + m[2],
			m[7] + m[6],
			m[11] + m[10],
		},
		D: -m[15] + m[14],
	}.Normalized()
	right_plane := Plane{
		Normal: Vector3{
			m[3] - m[0],
			m[7] - m[4],
			m[11] - m[8],
		},
		D: m[15] + m[12],
	}.Normalized()
	top_plane := Plane{
		Normal: Vector3{
			m[3] - m[1],
			m[7] - m[5],
			m[11] - m[9],
		},
		D: -m[15] + m[13],
	}.Normalized()
	res, _ := near_plane.Intersect3(right_plane, top_plane)
	return Vector2{res[X], res[Y]}
}

// FarZ returns the distance for this Projection beyond which positions are clipped.
func (p Projection) FarZ() Float {
	m := (*[16]float)(unsafe.Pointer(&p[0][0]))
	var new_plane = Plane{
		Normal: Vector3{
			m[3] - m[2],
			m[7] - m[6],
			m[11] - m[10],
		},
		D: m[15] - m[14],
	}.Normalized()
	return Float(new_plane.D)
}

// NearZ returns the distance for this Projection before which positions are clipped.
func (p Projection) NearZ() Float {
	m := (*[16]float)(unsafe.Pointer(&p[0][0]))
	var new_plane = Plane{
		Normal: Vector3{
			m[3] + m[2],
			m[7] + m[6],
			m[11] + m[10],
		},
		D: -m[15] - m[14],
	}.Normalized()
	return Float(new_plane.D)
}

// Inverse returns a Projection that performs the inverse of this Projection's projective transformation.
func (p Projection) Inverse() Projection {
	var i, j int
	var pvt_i, pvt_j [4]int /* Locations of pivot matrix */
	var pvt_val float       /* Value of current pivot element */
	var hold float          /* Temporary storage */
	var determinant float = 1.0
	for k := 0; k < 4; k++ {
		/** Locate k'th pivot element **/
		pvt_val = p[k][k] /** Initialize for search **/
		pvt_i[k] = k
		pvt_j[k] = k
		for i := k; i < 4; i++ {
			for j := k; j < 4; j++ {
				if Absf(p[i][j]) > Absf(pvt_val) {
					pvt_i[k] = i
					pvt_j[k] = j
					pvt_val = p[i][j]
				}
			}
		}
		/** Product of pivots, gives determinant when finished **/
		determinant *= pvt_val
		if IsApproximatelyZero(determinant) {
			return Projection{} /** Matrix is singular (zero determinant). **/
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
	return p
}

// IsOrthogonal returns true if this Projection performs an orthogonal projection.
func (p Projection) IsOrthogonal() bool { return p[3][3] == 1.0 }

// JitterOffseted returns a Projection with the X and Y values from the given Vector2
// added to the first and second values of the final column respectively.
func (p Projection) JitterOffseted(jitter Vector2) Projection {
	p[3][0] += jitter[X]
	p[3][1] += jitter[Y]
	return p
}

// PerspectiveAdjustedNearZ returns a Projection with the near clipping distance adjusted to be new_znear.
//
// Note: The original Projection must be a perspective projection.
func (p Projection) PerspectiveAdjustedNearZ(new_near_z Float) Projection {
	var zfar = p.FarZ()
	var znear = new_near_z
	var deltaZ = zfar - znear
	p[2][2] = float(-(zfar + znear) / deltaZ)
	p[3][2] = float(-2 * znear * zfar / deltaZ)
	return p
}

func (p Projection) Transform3D() Transform3D {
	return Transform3D{
		Basis: Basis{
			Vector3{p[0][0], p[0][1], p[0][2]},
			Vector3{p[1][0], p[1][1], p[1][2]},
			Vector3{p[2][0], p[2][1], p[2][2]},
		},
		Origin: Vector3{p[3][0], p[3][1], p[3][2]},
	}
}

func (p Projection) Mul(other Projection) Projection {
	var new_matrix Projection
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			var ab float = 0
			for k := 0; k < 4; k++ {
				ab += p[k][i] * other[j][k]
			}
			new_matrix[j][i] = ab
		}
	}
	return new_matrix
}
