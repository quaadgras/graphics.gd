---
title: Underlying Data Structures
slug: guide/variants
sidebar:
  order: 20
---

The gdextension interface to the engine supports a fixed set of underlying data types, all of
them have an equivalent convenience type in Go and a high-performance type (this type is
allocation efficient and suitable for use in hot functions).

All of the data types have a package underneath [graphics.gd/variant](https://pkg.go.dev/graphics.gd/variant).
With the exception of `Object` and `RefCounted`, all these packages are implemented in pure Go to avoid any
coupling or overhead when making calls to the engine. You can import these packages in any Go project.

### Core Types

| Engine Type        | Convenience Type          | High Performance Type           |
| ------------------ | ------------------------- | ------------------------------- |
| Variant            | `any`                     | `variant.Any`                   |
| bool               | `bool`                    | `bool`                          |
| int                | `int`                     | `int64`                         |
| float              | `Float.X`                 | `float64`                       |
| String             | `string`                  | `String.Unicode`                |
| Vector2            | `Vector2.XY`              | `Vector2.XY`                    |
| Vector2i           | `Vector2i.XY`             | `Vector2i.XY`                   |
| Rect2              | `Rect2.PositionSize`      | `Rect2.PositionSize`            |
| Rect2i             | `Rect2i.PositionSize`     | `Rect2i.PositionSize`           |
| Vector3            | `Vector3.XYZ`             | `Vector3.XYZ`                   |
| Vector3i           | `Vector3i.XYZ`            | `Vector3i.XYZ`                  |
| Transform2D        | `Transform2D.OriginXY`    | `Transform2D.OriginXY`          |
| Vector4            | `Vector4.XYZW`            | `Vector4.XYZW`                  |
| Vector4i           | `Vector4i.XYZW`           | `Vector4i.XYZW`                 |
| Plane              | `Plane.NormalD`           | `Plane.NormalD`                 |
| Quaternion         | `Quaternion.IJKL`         | `Quaternion.IJKL`               |
| AABB               | `AABB.PositionSize`       | `AABB.PositionSize`             |
| Basis              | `Basis.XYZ`               | `Basis.XYZ`                     |
| Transform3D        | `Transform3D.BasisOrigin` | `Transform3D.BasisOrigin`       |
| Projection         | `Projection.XYZW`         | `Projection.XYZW`               |
| Color              | `Color.RGBA`              | `Color.RGBA`                    |
| StringName         | `string`                  | `String.Name`                   |
| NodePath           | `string`                  | `Path.ToNode`                   |
| Signal             | `chan T`                  | `Signal.Any`                    |
| RID                | `RID.T`                   | `RID.Any`                       |
| Object             | `*T \| T.Instance`        | `T.Advanced`                    |
| Callable           | `func(...T) (...T)`       | `Callable.Function`             |
| Dictionary         | `struct \| map[T]T`       | `Dictionary.Any`                |
| Array              | `[]T`                     | `Array.Any`                     |
| PackedByteArray    | `[]byte`                  | `Packed.Bytes`                  |
| PackedInt32Array   | `[]int32`                 | `Packed.Array[int32]`           |
| PackedInt64Array   | `[]int64`                 | `Packed.Array[int64]`           |
| PackedFloat32Array | `[]float32`               | `Packed.Array[float32]`         |
| PackedFloat64Array | `[]float64`               | `Packed.Array[float64]`         |
| PackedStringArray  | `[]string`                | `Packed.Strings`                |
| PackedVector2Array | `[]Vector2.XY`            | `Packed.Array[Vector2.XY]`      |
| PackedVector3Array | `[]Vector3.XYZ`           | `Packed.Array[Vector3.XYZ]`     |
| PackedColorArray   | `[]Color.RGBA`            | `Packed.Array[Color.RGBA]`      |
| PackedVector4Array | `[]Vector4.XYZW`          | `Packed.Array[Vector4.XYZW]`    |


### Additional Types
`graphics.gd` defines some additional variant types to improve type-safety and readability.

| Additional Type    | Underlying Engine Type  |
|--------------------|-------------------------|
| Enum.Int[T]        | `int`                   |
| Angle.Radians      | `float`                 |
| Angle.Degrees      | `float`                 |
| Error.Code         | `int`                   |
| Euler.Radians      | `Vector3`               |
| Euler.Degrees      | `Vector3`               |
| Path.ToFile        | `String`                |
| Path.ToResource    | `String`                |
| Path.ToDirectory   | `String`                |

### Bring Your Own Vectors
The functions available within the `Vector2`, `Vector2i`, `Vector3`, `Vector3i`, `Vector4`, and `Vector4i` packages
can operate with any matching vector types (as long as they share the same underlying `struct`), they are not requred by
`graphics.gd` and you are more than welcome to define your own vector types with their own set of methods .

```go
package myvectors

type MyVector2 struct {
	X float32
	Y float32
}

func (v MyVector2) Add(other MyVector2) MyVector2 {
	return Vector2.Add(v, other) // this works without any special type conversions.
}
```

In this example, `MyVector2` can be passed to engine methods as if it were a `Vector2.XY`.
