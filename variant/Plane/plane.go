// Package Plane provides a plane in Hessian normal form.
package Plane

import (
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
	"graphics.gd/variant/Vector3"
)

// NormalD represents a normalized plane equation. normal is the normal
// of the plane (a, b, c normalized), and d is the distance from the
// origin to the plane (in the direction of "normal"). "Over" or "Above"
// the plane is considered the side of the plane towards where the normal
// is pointing.
type NormalD = struct {
	// Normal of the plane, typically a unit vector. Shouldn't be a zero vector as
	// Plane with such normal does not represent a valid plane.
	//
	// In the scalar equation of the plane ax + by + cz = d, this is the
	// vector (a, b, c), where d is the d property.
	Normal Vector3.XYZ

	// D is the distance from the origin to the plane, expressed in terms of normal
	// (according to its direction and magnitude). Actual absolute distance from the
	// origin to the plane can be calculated as Abs(d) / normal.Length() (if normal
	// has zero length then this Plane does not represent a valid plane).
	//
	// In the scalar equation of the plane ax + by + cz = d, this is d, while the (a, b, c) coordinates are represented by the
	// normal property.
	D Float.X
}

var (
	YZ = NormalD{Normal: Vector3.XYZ{1, 0, 0}} // A plane that extends in the Y and Z axes (normal vector points +X).
	XZ = NormalD{Normal: Vector3.XYZ{0, 1, 0}} // A plane that extends in the X and Z axes (normal vector points +Y).
	XY = NormalD{Normal: Vector3.XYZ{0, 0, 1}} // A plane that extends in the X and Y axes (normal vector points +Z).
)

// New creates a plane from the four parameters. The three components of the resulting
// plane's normal are a, b and c, and the plane has a distance of d from the origin.
func New[X Float.Any | Int.Any](a, b, c, d X) NormalD { //gd:Plane(a:float,b:float,c:float,d:float)
	return NormalD{Normal: Vector3.New(a, b, c), D: Float.X(d)}
}

// Points creates a plane from the three points, given in clockwise order.
func Points(a, b, c Vector3.XYZ) NormalD { //gd:Plane(point1:Vector3,point2:Vector3,point3:Vector3)
	normal := Vector3.Normalized(Vector3.Cross(Vector3.Sub(a, c), Vector3.Sub(a, b)))
	return NormalD{
		Normal: normal,
		D:      Vector3.Dot(normal, a),
	}
}

// NormalPoint creates a plane from the normal vector and a point on the plane.
//
// The normal of the plane must be a unit vector.
func NormalPoint(normal, point Vector3.XYZ) NormalD { //gd:Plane(normal:Vector3,point:Vector3)
	return NormalD{
		Normal: normal,
		D:      Vector3.Dot(normal, point),
	}
}

// DistanceToPoint returns the shortest distance from the plane to the position point.
// If the point is above the plane, the distance will be positive. If below,
// the distance will be negative.
func DistanceToPoint(point Vector3.XYZ, p NormalD) Float.X { //gd:Plane.distance_to
	return Vector3.Dot(p.Normal, point) - p.D
}

// GetCenter returns the center of the plane.
func Center(p NormalD) Vector3.XYZ { //gd:Plane.get_center
	return Vector3.MulX(p.Normal, p.D)
}

// HasPoint returns true if point is inside the plane. Comparison uses a custom minimum
// tolerance threshold.
func HasPoint(point Vector3.XYZ, p NormalD) bool { //gd:Plane.has_point
	return Float.Abs(Vector3.Dot(p.Normal, point)-p.D) <= Float.Epsilon
}

// Intersect3 returns the intersection point of the three planes b, c and this plane. If
// no intersection is found, false is returned.
func Intersect3(a, b, c NormalD) (Vector3.XYZ, bool) { //gd:Plane.intersect_3
	var normal0 = a.Normal
	var normal1 = b.Normal
	var normal2 = c.Normal
	var (
		denom = Vector3.Dot(Vector3.Cross(normal0, normal1), normal2)
	)
	if Float.IsApproximatelyZero(denom) {
		return Vector3.Zero, false
	}
	return Vector3.Add(
		Vector3.MulX(Vector3.Cross(normal1, normal2), a.D),
		Vector3.DivX(Vector3.Add(
			Vector3.MulX(Vector3.Cross(normal2, normal0), b.D),
			Vector3.MulX(Vector3.Cross(normal0, normal1), c.D)), denom)), true
}

// IntersectsRay returns the intersection point of a ray consisting of the position from and the
// direction normal dir with this plane. If no intersection is found, false is returned.
func IntersectsRay(p NormalD, from, dir Vector3.XYZ) (Vector3.XYZ, bool) { //gd:Plane.intersects_ray
	var segment = dir
	var den = Vector3.Dot(p.Normal, segment)
	if Float.IsApproximatelyZero(den) {
		return Vector3.Zero, false
	}
	var dist = (Vector3.Dot(p.Normal, from) - p.D) / den
	if dist > Float.Epsilon { //this is a ray, before the emitting pos (p_from) doesn't exist
		return Vector3.Zero, false
	}
	dist = -dist
	return Vector3.Add(from, Vector3.MulX(segment, dist)), true
}

// IntersectsSegment returns the intersection point of a segment from position from to position to
// with this plane. If no intersection is found, false is returned.
func IntersectsSegment(p NormalD, from, to Vector3.XYZ) (Vector3.XYZ, bool) { //gd:Plane.intersects_segment
	var segment = Vector3.Sub(to, from)
	var den = Vector3.Dot(p.Normal, segment)
	if Float.IsApproximatelyZero(den) {
		return Vector3.Zero, false
	}
	var dist = (Vector3.Dot(p.Normal, from) - p.D) / den
	if dist < -Float.Epsilon || dist > 1+Float.Epsilon {
		return Vector3.Zero, false
	}
	return Vector3.Add(from, Vector3.MulX(segment, dist)), true
}

// IsApproximatelyEqual returns true if this plane and other are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func IsApproximatelyEqual(p, other NormalD) bool { //gd:Plane.is_equal_approx
	return Vector3.IsApproximatelyEqual(p.Normal, other.Normal) && Float.IsApproximatelyEqual(p.D, other.D)
}

// IsFinite returns true if this plane is finite, by calling [IsFinite] on each component.
func IsFinite(p NormalD) bool { return Vector3.IsFinite(p.Normal) && Float.IsFinite(p.D) } //gd:Plane.is_finite

// IsPointOver returns true if point is located above the plane.
func IsPointOver(p NormalD, point Vector3.XYZ) bool { //gd:Plane.is_point_over
	return Vector3.Dot(p.Normal, point) > p.D
}

// Normalized returns a copy of the plane, with normalized normal (so it's a unit vector).
// Returns a zero value if normal can't be normalized (it has zero length).
func Normalized(p NormalD) NormalD { //gd:Plane.normalized
	var l = Vector3.Length(p.Normal)
	if l == 0 {
		return NormalD{}
	}
	return NormalD{
		Normal: Vector3.DivX(p.Normal, l),
		D:      p.D / l,
	}
}

// Project returns the orthogonal projection of point into a point in the plane.
func Project(point Vector3.XYZ, p NormalD) Vector3.XYZ { //gd:Plane.project
	return Vector3.Sub(point, Vector3.MulX(p.Normal, DistanceToPoint(point, p)))
}
