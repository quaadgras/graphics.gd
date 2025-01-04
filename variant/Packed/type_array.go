package Packed

import (
	"graphics.gd/variant/Color"
	"graphics.gd/variant/String"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"

	gd "graphics.gd/internal"
)

// ColorArray is an array specifically designed to hold Color. Packs data tightly, so it saves
// memory for large array sizes.
type ColorArray struct {
	Array[ColorArray, Color.RGBA, gd.Color, gd.PackedColorArray, *gd.PackedColorArray]
}

func (array ColorArray) less(a, b Color.RGBA) bool  { return Color.Less(a, b) }
func (array ColorArray) conv(c Color.RGBA) gd.Color { return c }
func (array ColorArray) wrap(c gd.Color) Color.RGBA { return c }
func (array ColorArray) alloc() gd.PackedColorArray { return gd.NewPackedColorArray() }

func (array ColorArray) make(local []Color.RGBA, proxy gd.PackedColorArray) ColorArray {
	array.local = local
	array.proxy = proxy
	return array
}

// Float32Array is an array specifically designed to hold float32. Packs data tightly, so it saves
// memory for large array sizes.
type Float32Array struct {
	Array[Float32Array, float32, gd.Float, gd.PackedFloat32Array, *gd.PackedFloat32Array]
}

func (array Float32Array) less(a, b float32) bool       { return a < b }
func (array Float32Array) conv(c float32) gd.Float      { return gd.Float(c) }
func (array Float32Array) wrap(c gd.Float) float32      { return float32(c) }
func (array Float32Array) alloc() gd.PackedFloat32Array { return gd.NewPackedFloat32Array() }

func (array Float32Array) make(local []float32, proxy gd.PackedFloat32Array) Float32Array {
	array.local = local
	array.proxy = proxy
	return array
}

// Float64Array is an array specifically designed to hold float64. Packs data tightly, so it saves
// memory for large array sizes.
type Float64Array struct {
	Array[Float64Array, float64, gd.Float, gd.PackedFloat64Array, *gd.PackedFloat64Array]
}

func (array Float64Array) less(a, b float64) bool       { return a < b }
func (array Float64Array) conv(c float64) gd.Float      { return gd.Float(c) }
func (array Float64Array) wrap(c gd.Float) float64      { return float64(c) }
func (array Float64Array) alloc() gd.PackedFloat64Array { return gd.NewPackedFloat64Array() }

func (array Float64Array) make(local []float64, proxy gd.PackedFloat64Array) Float64Array {
	array.local = local
	array.proxy = proxy
	return array
}

// Int32Array is an array specifically designed to hold int32. Packs data tightly, so it saves
// memory for large array sizes.
type Int32Array struct {
	Array[Int32Array, int32, gd.Int, gd.PackedInt32Array, *gd.PackedInt32Array]
}

func (array Int32Array) less(a, b int32) bool       { return a < b }
func (array Int32Array) conv(c int32) gd.Int        { return gd.Int(c) }
func (array Int32Array) wrap(c gd.Int) int32        { return int32(c) }
func (array Int32Array) alloc() gd.PackedInt32Array { return gd.NewPackedInt32Array() }

func (array Int32Array) make(local []int32, proxy gd.PackedInt32Array) Int32Array {
	array.local = local
	array.proxy = proxy
	return array
}

// Int64Array is an array specifically designed to hold int64. Packs data tightly, so it saves
// memory for large array sizes.
type Int64Array struct {
	Array[Int64Array, int64, gd.Int, gd.PackedInt64Array, *gd.PackedInt64Array]
}

func (array Int64Array) less(a, b int64) bool       { return a < b }
func (array Int64Array) conv(c int64) gd.Int        { return gd.Int(c) }
func (array Int64Array) wrap(c gd.Int) int64        { return int64(c) }
func (array Int64Array) alloc() gd.PackedInt64Array { return gd.NewPackedInt64Array() }

func (array Int64Array) make(local []int64, proxy gd.PackedInt64Array) Int64Array {
	array.local = local
	array.proxy = proxy
	return array
}

// StringArray is an array specifically designed to hold string. Packs data tightly, so it saves
// memory for large array sizes.
type StringArray struct {
	Array[StringArray, string, gd.String, gd.PackedStringArray, *gd.PackedStringArray]
}

func (array StringArray) less(a, b string) bool       { return a < b }
func (array StringArray) conv(c string) gd.String     { return String.New(c) }
func (array StringArray) wrap(c gd.String) string     { return c.String() }
func (array StringArray) alloc() gd.PackedStringArray { return gd.NewPackedStringArray() }

func (array StringArray) make(local []string, proxy gd.PackedStringArray) StringArray {
	array.local = local
	array.proxy = proxy
	return array
}

// Vector2Array is an array specifically designed to hold Vector2. Packs data tightly, so it saves
// memory for large array sizes.
type Vector2Array struct {
	Array[Vector2Array, Vector2.XY, gd.Vector2, gd.PackedVector2Array, *gd.PackedVector2Array]
}

func (array Vector2Array) less(a, b Vector2.XY) bool    { return Vector2.Less(a, b) }
func (array Vector2Array) conv(c Vector2.XY) gd.Vector2 { return gd.Vector2(c) }
func (array Vector2Array) wrap(c gd.Vector2) Vector2.XY { return Vector2.XY(c) }
func (array Vector2Array) alloc() gd.PackedVector2Array { return gd.NewPackedVector2Array() }

func (array Vector2Array) make(local []Vector2.XY, proxy gd.PackedVector2Array) Vector2Array {
	array.local = local
	array.proxy = proxy
	return array
}

// Vector3Array is an array specifically designed to hold Vector3. Packs data tightly, so it saves
// memory for large array sizes.
type Vector3Array struct {
	Array[Vector3Array, Vector3.XYZ, gd.Vector3, gd.PackedVector3Array, *gd.PackedVector3Array]
}

func (array Vector3Array) less(a, b Vector3.XYZ) bool    { return Vector3.Less(a, b) }
func (array Vector3Array) conv(c Vector3.XYZ) gd.Vector3 { return gd.Vector3(c) }
func (array Vector3Array) wrap(c gd.Vector3) Vector3.XYZ { return Vector3.XYZ(c) }
func (array Vector3Array) alloc() gd.PackedVector3Array  { return gd.NewPackedVector3Array() }

func (array Vector3Array) make(local []Vector3.XYZ, proxy gd.PackedVector3Array) Vector3Array {
	array.local = local
	array.proxy = proxy
	return array
}

// Vector4Array is an array specifically designed to hold Vector4. Packs data tightly, so it saves
// memory for large array sizes.
type Vector4Array struct {
	Array[Vector4Array, Vector4.XYZW, gd.Vector4, gd.PackedVector4Array, *gd.PackedVector4Array]
}

func (array Vector4Array) less(a, b Vector4.XYZW) bool    { return Vector4.Less(a, b) }
func (array Vector4Array) conv(c Vector4.XYZW) gd.Vector4 { return gd.Vector4(c) }
func (array Vector4Array) wrap(c gd.Vector4) Vector4.XYZW { return Vector4.XYZW(c) }
func (array Vector4Array) alloc() gd.PackedVector4Array   { return gd.NewPackedVector4Array() }

func (array Vector4Array) make(local []Vector4.XYZW, proxy gd.PackedVector4Array) Vector4Array {
	array.local = local
	array.proxy = proxy
	return array
}
