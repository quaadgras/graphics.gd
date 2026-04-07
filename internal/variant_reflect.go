package gd

// Bool returns the variant as a bool. Panics if the variant is not a bool.
func (v Variant) Bool() bool { return v.Interface().(bool) }

// Float returns the variant as a float. Panics if the variant is not a float.
func (v Variant) Float() Float { return v.Interface().(Float) }

// Int returns the variant as an int. Panics if the variant is not an int.
func (v Variant) Int() Int { return v.Interface().(Int) }

// Vector2 returns the variant as a Vector2. Panics if the variant is not a Vector2.
func (v Variant) Vector2() Vector2 { return v.Interface().(Vector2) }

// Vector2i returns the variant as a Vector2i. Panics if the variant is not a Vector2i.
func (v Variant) Vector2i() Vector2i { return v.Interface().(Vector2i) }

// Rect2 returns the variant as a Rect2. Panics if the variant is not a Rect2.
func (v Variant) Rect2() Rect2 { return v.Interface().(Rect2) }

// Rect2i returns the variant as a Rect2i. Panics if the variant is not a Rect2i.
func (v Variant) Rect2i() Rect2i { return v.Interface().(Rect2i) }

// Vector3 returns the variant as a Vector3. Panics if the variant is not a Vector3.
func (v Variant) Vector3() Vector3 { return v.Interface().(Vector3) }

// Vector3i returns the variant as a Vector3i. Panics if the variant is not a Vector3i.
func (v Variant) Vector3i() Vector3i { return v.Interface().(Vector3i) }

// Transform2D returns the variant as a Transform2D. Panics if the variant is not a Transform2D.
func (v Variant) Transform2D() Transform2D { return v.Interface().(Transform2D) }

// Vector4 returns the variant as a Vector4. Panics if the variant is not a Vector4.
func (v Variant) Vector4() Vector4 { return v.Interface().(Vector4) }

// Vector4i returns the variant as a Vector4i. Panics if the variant is not a Vector4i.
func (v Variant) Vector4i() Vector4i { return v.Interface().(Vector4i) }

// Plane returns the variant as a Plane. Panics if the variant is not a Plane.
func (v Variant) Plane() Plane { return v.Interface().(Plane) }

// Quaternion returns the variant as a Quaternion. Panics if the variant is not a Quaternion.
func (v Variant) Quaternion() Quaternion { return v.Interface().(Quaternion) }

// AABB returns the variant as an AABB. Panics if the variant is not an AABB.
func (v Variant) AABB() AABB { return v.Interface().(AABB) }

// Basis returns the variant as a Basis. Panics if the variant is not a Basis.
func (v Variant) Basis() Basis { return v.Interface().(Basis) }

// Transform3D returns the variant as a Transform3D. Panics if the variant is not a Transform3D.
func (v Variant) Transform3D() Transform3D { return v.Interface().(Transform3D) }

// Projection returns the variant as a Projection. Panics if the variant is not a Projection.
func (v Variant) Projection() Projection { return v.Interface().(Projection) }

// Color returns the variant as a Color. Panics if the variant is not a Color.
func (v Variant) Color() Color { return v.Interface().(Color) }
