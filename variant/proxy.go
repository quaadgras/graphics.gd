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

func Through(proxy Proxy, state complex128) Any {
	return Any{value: proxy, local: state}
}

type Proxy interface {
	Bool(complex128) bool
	Int(complex128) int64
	Float(complex128) float64
	String(complex128) String.Readable
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

	Convert(complex128, reflect.Type) reflect.Value
	AssignableTo(complex128, reflect.Type) bool
	ConvertibleTo(complex128, reflect.Type) bool
	Calculate(complex128, Operator, complex128) Any
	Call(complex128, String.Readable, ...Any) Any
	Set(complex128, Any, Any) bool
	Get(complex128, Any) (Any, bool)
	Iter(complex128) iter.Seq[Any]
	Hash(complex128, int) uint32
	Duplicate(complex128) Any
	Type(complex128) Type
	Interface(complex128) interface{}
	Has(complex128, Any) bool
}
