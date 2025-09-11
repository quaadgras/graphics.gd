package Transform3D_test

import (
	"math"
	"testing"

	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector3"
)

// Helper function to check if two float values are approximately equal
func approxEqual(a, b Float.X, tolerance Float.X) bool {
	return math.Abs(float64(a-b)) < float64(tolerance)
}

// Helper function to check if two Vector3 values are approximately equal
func vector3ApproxEqual(a, b Vector3.XYZ, tolerance Float.X) bool {
	return approxEqual(a.X, b.X, tolerance) &&
		approxEqual(a.Y, b.Y, tolerance) &&
		approxEqual(a.Z, b.Z, tolerance)
}

// Helper function to check if two AABB values are approximately equal
func aabbApproxEqual(a, b AABB.PositionSize, tolerance Float.X) bool {
	return vector3ApproxEqual(a.Position, b.Position, tolerance) &&
		vector3ApproxEqual(a.Size, b.Size, tolerance)
}

func TestTransformAABB_Identity(t *testing.T) {
	// Test that identity transform doesn't change the AABB
	aabb := AABB.PositionSize{
		Position: Vector3.XYZ{X: 1, Y: 2, Z: 3},
		Size:     Vector3.XYZ{X: 4, Y: 5, Z: 6},
	}

	result := Transform3D.TransformAABB(aabb, Transform3D.Identity)

	if !aabbApproxEqual(result, aabb, 1e-6) {
		t.Errorf("Identity transform should not change AABB. Expected %+v, got %+v", aabb, result)
	}
}

func TestTransformAABB_Translation(t *testing.T) {
	// Test pure translation
	aabb := AABB.PositionSize{
		Position: Vector3.XYZ{X: 0, Y: 0, Z: 0},
		Size:     Vector3.XYZ{X: 2, Y: 2, Z: 2},
	}

	translation := Vector3.XYZ{X: 10, Y: 20, Z: 30}
	transform := Transform3D.BasisOrigin{
		Basis:  Basis.Identity,
		Origin: translation,
	}

	result := Transform3D.TransformAABB(aabb, transform)
	expected := AABB.PositionSize{
		Position: translation,
		Size:     aabb.Size,
	}

	if !aabbApproxEqual(result, expected, 1e-6) {
		t.Errorf("Translation transform failed. Expected %+v, got %+v", expected, result)
	}
}

func TestTransformAABB_UniformScale(t *testing.T) {
	// Test uniform scaling
	aabb := AABB.PositionSize{
		Position: Vector3.XYZ{X: 1, Y: 1, Z: 1},
		Size:     Vector3.XYZ{X: 2, Y: 2, Z: 2},
	}

	scale := Float.X(2.0)
	transform := Transform3D.BasisOrigin{
		Basis: Basis.XYZ{
			X: Vector3.XYZ{X: scale, Y: 0, Z: 0},
			Y: Vector3.XYZ{X: 0, Y: scale, Z: 0},
			Z: Vector3.XYZ{X: 0, Y: 0, Z: scale},
		},
		Origin: Vector3.Zero,
	}

	result := Transform3D.TransformAABB(aabb, transform)
	expected := AABB.PositionSize{
		Position: Vector3.XYZ{X: 2, Y: 2, Z: 2}, // Position scaled
		Size:     Vector3.XYZ{X: 4, Y: 4, Z: 4}, // Size scaled
	}

	if !aabbApproxEqual(result, expected, 1e-6) {
		t.Errorf("Uniform scale transform failed. Expected %+v, got %+v", expected, result)
	}
}

func TestTransformAABB_Rotation90Z(t *testing.T) {
	// Test 90-degree rotation around Z-axis
	aabb := AABB.PositionSize{
		Position: Vector3.XYZ{X: 1, Y: 0, Z: 0},
		Size:     Vector3.XYZ{X: 1, Y: 1, Z: 1},
	}

	// 90-degree rotation around Z-axis: X becomes Y, Y becomes -X
	transform := Transform3D.BasisOrigin{
		Basis: Basis.XYZ{
			X: Vector3.XYZ{X: 0, Y: 1, Z: 0},  // Original X axis becomes Y
			Y: Vector3.XYZ{X: -1, Y: 0, Z: 0}, // Original Y axis becomes -X
			Z: Vector3.XYZ{X: 0, Y: 0, Z: 1},  // Z unchanged
		},
		Origin: Vector3.Zero,
	}

	result := Transform3D.TransformAABB(aabb, transform)

	// After rotation, the AABB should encompass the rotated box
	// Original box spans from (1,0,0) to (2,1,1)
	// After rotation: (1,0,0) -> (0,1,0), (2,1,1) -> (-1,2,1)
	// So the AABB should span from (-1,1,0) to (0,2,1)
	expected := AABB.PositionSize{
		Position: Vector3.XYZ{X: -1, Y: 1, Z: 0},
		Size:     Vector3.XYZ{X: 1, Y: 1, Z: 1},
	}

	if !aabbApproxEqual(result, expected, 1e-6) {
		t.Errorf("90-degree Z rotation failed. Expected %+v, got %+v", expected, result)
	}
}

func TestTransformAABB_NegativeScale(t *testing.T) {
	// Test negative scaling (mirroring)
	aabb := AABB.PositionSize{
		Position: Vector3.XYZ{X: 1, Y: 1, Z: 1},
		Size:     Vector3.XYZ{X: 2, Y: 2, Z: 2},
	}

	// Mirror across X-axis (flip X)
	transform := Transform3D.BasisOrigin{
		Basis: Basis.XYZ{
			X: Vector3.XYZ{X: -1, Y: 0, Z: 0}, // Flip X
			Y: Vector3.XYZ{X: 0, Y: 1, Z: 0},  // Y unchanged
			Z: Vector3.XYZ{X: 0, Y: 0, Z: 1},  // Z unchanged
		},
		Origin: Vector3.Zero,
	}

	result := Transform3D.TransformAABB(aabb, transform)

	// Original box spans from (1,1,1) to (3,3,3)
	// After X flip: (1,1,1) -> (-1,1,1), (3,3,3) -> (-3,3,3)
	// AABB should span from (-3,1,1) to (-1,3,3)
	expected := AABB.PositionSize{
		Position: Vector3.XYZ{X: -3, Y: 1, Z: 1},
		Size:     Vector3.XYZ{X: 2, Y: 2, Z: 2},
	}

	if !aabbApproxEqual(result, expected, 1e-6) {
		t.Errorf("Negative scale (mirror) failed. Expected %+v, got %+v", expected, result)
	}
}
