package variant

import (
	"iter"
	"reflect"

	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/String"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
	"graphics.gd/variant/Vector4i"
)

func Proxy[T API](current Any, alloc func() (T, complex128)) (T, complex128) {
	already, ok := current.value.(T)
	if ok && already.KeepAlive(current.local) {
		return already, current.local
	}
	impl, state := alloc()
	switch current.Type() {
	case TypeNil:
		return impl, state
	case TypeBool:
		return impl, impl.NewBool(current.Bool())
	case TypeInt:
		return impl, impl.NewInt(current.Int64())
	case TypeFloat:
		return impl, impl.NewFloat(current.Float64())
	case TypeVector2:
		return impl, impl.NewVector2(current.Vector2())
	case TypeVector2i:
		return impl, impl.NewVector2i(current.Vector2i())
	case TypeRect2:
		return impl, impl.NewRect2(current.Rect2())
	case TypeRect2i:
		return impl, impl.NewRect2i(current.Rect2i())
	case TypeVector3:
		return impl, impl.NewVector3(current.Vector3())
	case TypeVector3i:
		return impl, impl.NewVector3i(current.Vector3i())
	case TypeTransform2D:
		return impl, impl.NewTransform2D(current.Transform2D())
	case TypeVector4:
		return impl, impl.NewVector4(current.Vector4())
	case TypeVector4i:
		return impl, impl.NewVector4i(current.Vector4i())
	case TypePlane:
		return impl, impl.NewPlane(current.Plane())
	case TypeQuaternion:
		return impl, impl.NewQuaternion(current.Quaternion())
	case TypeAABB:
		return impl, impl.NewAABB(current.AABB())
	case TypeBasis:
		return impl, impl.NewBasis(current.Basis())
	case TypeTransform3D:
		return impl, impl.NewTransform3D(current.Transform3D())
	case TypeProjection:
		return impl, impl.NewProjection(current.Projection())
	case TypeColor:
		return impl, impl.NewColor(current.Color())
	default:
		return impl, impl.New(current.Interface())
	}
}

func Implementation(impl API, state complex128) Any { return Any{state, impl} }

type API interface {
	Bool(complex128) bool
	Int(complex128) int64
	Float(complex128) float64
	Vector2(complex128) Vector2.XY
	Vector2i(complex128) Vector2i.XY
	Rect2(complex128) Rect2.PositionSize
	Rect2i(complex128) Rect2i.PositionSize
	Vector3(complex128) Vector3.XYZ
	Vector3i(complex128) Vector3i.XYZ
	Transform2D(complex128) Transform2D.OriginXY
	Vector4(complex128) Vector4.XYZW
	Vector4i(complex128) Vector4i.XYZW
	Plane(complex128) Plane.NormalD
	Quaternion(complex128) Quaternion.IJKX
	AABB(complex128) AABB.PositionSize
	Basis(complex128) Basis.XYZ
	Transform3D(complex128) Transform3D.BasisOrigin
	Projection(complex128) Projection.XYZW
	Color(complex128) Color.RGBA
	Interface(complex128) interface{}

	New(any) complex128
	NewBool(bool) complex128
	NewInt(int64) complex128
	NewFloat(float64) complex128
	NewVector2(Vector2.XY) complex128
	NewVector2i(Vector2i.XY) complex128
	NewRect2(Rect2.PositionSize) complex128
	NewRect2i(Rect2i.PositionSize) complex128
	NewVector3(Vector3.XYZ) complex128
	NewVector3i(Vector3i.XYZ) complex128
	NewTransform2D(Transform2D.OriginXY) complex128
	NewVector4(Vector4.XYZW) complex128
	NewVector4i(Vector4i.XYZW) complex128
	NewPlane(Plane.NormalD) complex128
	NewQuaternion(Quaternion.IJKX) complex128
	NewAABB(AABB.PositionSize) complex128
	NewBasis(Basis.XYZ) complex128
	NewTransform3D(Transform3D.BasisOrigin) complex128
	NewProjection(Projection.XYZW) complex128
	NewColor(Color.RGBA) complex128

	Convert(complex128, reflect.Type) reflect.Value
	AssignableTo(complex128, reflect.Type) bool
	ConvertibleTo(complex128, reflect.Type) bool
	Calculate(complex128, Operator, complex128) Any
	Call(complex128, String.Readable, ...Any) Any
	Has(complex128, Any) bool
	Set(complex128, Any, Any) bool
	Get(complex128, Any) (Any, bool)
	Iter(complex128) iter.Seq2[Any, Any]
	Hash(complex128, int) uint32
	Duplicate(complex128) Any
	Type(complex128) Type
	String(complex128) string

	KeepAlive(complex128) bool
	Free(complex128)
}
