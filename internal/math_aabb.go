//go:build !generate

package gd

/*
AABB consists of a position, a size, and several utility functions. It is typically used for fast overlap tests.

It uses floating-point coordinates. The 2D counterpart to AABB is Rect2.

Negative values for size are not supported and will not work for most methods. Use abs to get an AABB with a positive size.

Note: Unlike Rect2, AABB does not have a variant that uses integer coordinates.
*/
type AABB struct {
	Position Vector3
	Size     Vector3
}
