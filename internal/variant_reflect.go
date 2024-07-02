package gd

// Bool returns the variant as a bool. Panics if the variant is not a bool.
func (variant Variant) Bool() Bool { return variant.Interface(Lifetime{}).(Bool) }

// Float returns the variant as a float. Panics if the variant is not a float.
func (variant Variant) Float() Float { return variant.Interface(Lifetime{}).(Float) }

// Int returns the variant as an int. Panics if the variant is not an int.
func (variant Variant) Int() Int { return variant.Interface(Lifetime{}).(Int) }

// Vector2 returns the variant as a Vector2. Panics if the variant is not a Vector2.
func (variant Variant) Vector2() Vector2 { return variant.Interface(Lifetime{}).(Vector2) }

// Vector2i returns the variant as a Vector2i. Panics if the variant is not a Vector2i.
func (variant Variant) Vector2i() Vector2i { return variant.Interface(Lifetime{}).(Vector2i) }

// Rect2 returns the variant as a Rect2. Panics if the variant is not a Rect2.
func (variant Variant) Rect2() Rect2 { return variant.Interface(Lifetime{}).(Rect2) }

// Rect2i returns the variant as a Rect2i. Panics if the variant is not a Rect2i.
func (variant Variant) Rect2i() Rect2i { return variant.Interface(Lifetime{}).(Rect2i) }

// Vector3 returns the variant as a Vector3. Panics if the variant is not a Vector3.
func (variant Variant) Vector3() Vector3 { return variant.Interface(Lifetime{}).(Vector3) }

// Vector3i returns the variant as a Vector3i. Panics if the variant is not a Vector3i.
func (variant Variant) Vector3i() Vector3i { return variant.Interface(Lifetime{}).(Vector3i) }

// Transform2D returns the variant as a Transform2D. Panics if the variant is not a Transform2D.
func (variant Variant) Transform2D() Transform2D { return variant.Interface(Lifetime{}).(Transform2D) }

// Vector4 returns the variant as a Vector4. Panics if the variant is not a Vector4.
func (variant Variant) Vector4() Vector4 { return variant.Interface(Lifetime{}).(Vector4) }

// Vector4i returns the variant as a Vector4i. Panics if the variant is not a Vector4i.
func (variant Variant) Vector4i() Vector4i { return variant.Interface(Lifetime{}).(Vector4i) }

// Plane returns the variant as a Plane. Panics if the variant is not a Plane.
func (variant Variant) Plane() Plane { return variant.Interface(Lifetime{}).(Plane) }

// Quaternion returns the variant as a Quaternion. Panics if the variant is not a Quaternion.
func (variant Variant) Quaternion() Quaternion { return variant.Interface(Lifetime{}).(Quaternion) }

// AABB returns the variant as an AABB. Panics if the variant is not an AABB.
func (variant Variant) AABB() AABB { return variant.Interface(Lifetime{}).(AABB) }

// Basis returns the variant as a Basis. Panics if the variant is not a Basis.
func (variant Variant) Basis() Basis { return variant.Interface(Lifetime{}).(Basis) }

// Transform3D returns the variant as a Transform3D. Panics if the variant is not a Transform3D.
func (variant Variant) Transform3D() Transform3D { return variant.Interface(Lifetime{}).(Transform3D) }

// Projection returns the variant as a Projection. Panics if the variant is not a Projection.
func (variant Variant) Projection() Projection { return variant.Interface(Lifetime{}).(Projection) }

// Color returns the variant as a Color. Panics if the variant is not a Color.
func (variant Variant) Color() Color { return variant.Interface(Lifetime{}).(Color) }
