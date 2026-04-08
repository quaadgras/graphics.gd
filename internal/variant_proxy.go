package gd

import (
	"iter"
	"reflect"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	AABBType "graphics.gd/variant/AABB"
	BasisType "graphics.gd/variant/Basis"
	ColorType "graphics.gd/variant/Color"
	PlaneType "graphics.gd/variant/Plane"
	ProjectionType "graphics.gd/variant/Projection"
	QuaternionType "graphics.gd/variant/Quaternion"
	RIDType "graphics.gd/variant/RID"
	Rect2Type "graphics.gd/variant/Rect2"
	Rect2iType "graphics.gd/variant/Rect2i"
	StringType "graphics.gd/variant/String"
	Transform2DType "graphics.gd/variant/Transform2D"
	Transform3DType "graphics.gd/variant/Transform3D"
	Vector2Type "graphics.gd/variant/Vector2"
	Vector2iType "graphics.gd/variant/Vector2i"
	Vector3Type "graphics.gd/variant/Vector3"
	Vector3iType "graphics.gd/variant/Vector3i"
	Vector4Type "graphics.gd/variant/Vector4"
	Vector4iType "graphics.gd/variant/Vector4i"
)

func InternalVariant(extract VariantPkg.Any) Variant {
	return NewVariant(extract)
}

type VariantProxy struct{}

func (VariantProxy) Bool(raw complex128) bool {
	return gdunsafe.VariantInto[bool](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Int(raw complex128) int64 {
	return gdunsafe.VariantInto[int64](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Float(raw complex128) float64 {
	return gdunsafe.VariantInto[float64](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Vector2(raw complex128) Vector2Type.XY {
	return gdunsafe.VariantInto[Vector2Type.XY](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Vector2i(raw complex128) Vector2iType.XY {
	return gdunsafe.VariantInto[Vector2iType.XY](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Rect2(raw complex128) Rect2Type.PositionSize {
	return gdunsafe.VariantInto[Rect2Type.PositionSize](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Rect2i(raw complex128) Rect2iType.PositionSize {
	return gdunsafe.VariantInto[Rect2iType.PositionSize](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Vector3(raw complex128) Vector3Type.XYZ {
	return gdunsafe.VariantInto[Vector3Type.XYZ](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Vector3i(raw complex128) Vector3iType.XYZ {
	return gdunsafe.VariantInto[Vector3iType.XYZ](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Transform2D(raw complex128) Transform2DType.OriginXY {
	return gdunsafe.VariantInto[Transform2DType.OriginXY](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Vector4(raw complex128) Vector4Type.XYZW {
	return gdunsafe.VariantInto[Vector4Type.XYZW](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Vector4i(raw complex128) Vector4iType.XYZW {
	return gdunsafe.VariantInto[Vector4iType.XYZW](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Plane(raw complex128) PlaneType.NormalD {
	return gdunsafe.VariantInto[PlaneType.NormalD](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Quaternion(raw complex128) QuaternionType.IJKX {
	return gdunsafe.VariantInto[QuaternionType.IJKX](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) AABB(raw complex128) AABBType.PositionSize {
	return gdunsafe.VariantInto[AABBType.PositionSize](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Basis(raw complex128) BasisType.XYZ {
	return gdunsafe.VariantInto[BasisType.XYZ](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Transform3D(raw complex128) Transform3DType.BasisOrigin {
	return gdunsafe.VariantInto[Transform3DType.BasisOrigin](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Projection(raw complex128) ProjectionType.XYZW {
	return gdunsafe.VariantInto[ProjectionType.XYZW](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Color(raw complex128) ColorType.RGBA {
	return gdunsafe.VariantInto[ColorType.RGBA](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Interface(raw complex128) any {
	return pointers.Load[Variant](raw).Interface()
}
func (VariantProxy) RID(raw complex128) RIDType.Any {
	return gdunsafe.VariantInto[RIDType.Any](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
}
func (VariantProxy) Bytes(raw complex128) []byte {
	packed := gdunsafe.VariantInto[gdunsafe.PackedArray[byte]](gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))))
	return pointers.New[PackedByteArray](gdextension.PackedArray[byte](packed)).Bytes()
}

func (VariantProxy) New(val any) complex128 {
	return pointers.Pack(NewVariant(val))
}
func (VariantProxy) NewBool(val bool) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewInt(val int64) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewFloat(val float64) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewVector2(val Vector2Type.XY) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewVector2i(val Vector2iType.XY) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewRect2(val Rect2Type.PositionSize) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewRect2i(val Rect2iType.PositionSize) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewVector3(val Vector3Type.XYZ) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewVector3i(val Vector3iType.XYZ) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewTransform2D(val Transform2DType.OriginXY) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewVector4(val Vector4Type.XYZW) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewVector4i(val Vector4iType.XYZW) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewPlane(val PlaneType.NormalD) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewQuaternion(val QuaternionType.IJKX) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewAABB(val AABBType.PositionSize) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewBasis(val BasisType.XYZ) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewTransform3D(val Transform3DType.BasisOrigin) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewProjection(val ProjectionType.XYZW) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewColor(val ColorType.RGBA) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}
func (VariantProxy) NewRID(val RIDType.Any) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(val)))
}

func (VariantProxy) NewBytes(val []byte) complex128 {
	return pointers.Pack(pointers.New[Variant](gdunsafe.VariantFrom(gdunsafe.PackedArray[byte](pointers.Get(NewPackedByteSlice(val))))))
}

func (VariantProxy) Convert(raw complex128, rtype reflect.Type) reflect.Value {
	rvalue, err := convertVariantToDesiredGoType(pointers.Load[Variant](raw), rtype)
	if err != nil {
		panic(err)
	}
	return rvalue
}
func (VariantProxy) AssignableTo(raw complex128, rtype reflect.Type) bool { return true }
func (VariantProxy) ConvertibleTo(complex128, reflect.Type) bool          { return true }
func (VariantProxy) Calculate(complex128, VariantPkg.Operator, complex128) VariantPkg.Any {
	panic("not implemented")
}
func (VariantProxy) Call(complex128, StringType.Unicode, ...VariantPkg.Any) VariantPkg.Any {
	panic("not implemented")
}
func (VariantProxy) Has(complex128, VariantPkg.Any) bool {
	panic("not implemented")
}
func (VariantProxy) Set(complex128, VariantPkg.Any, VariantPkg.Any) bool {
	panic("not implemented")
}
func (VariantProxy) Get(complex128, VariantPkg.Any) (VariantPkg.Any, bool) {
	panic("not implemented")
}
func (VariantProxy) Iter(complex128) iter.Seq2[VariantPkg.Any, VariantPkg.Any] {
	panic("not implemented")
}
func (VariantProxy) Hash(complex128, int) uint32 {
	panic("not implemented")
}
func (VariantProxy) Duplicate(raw complex128) VariantPkg.Any {
	panic("not implemented")
}
func (VariantProxy) Type(raw complex128) VariantPkg.Type {
	return VariantPkg.Type(pointers.Load[Variant](raw).Type())
}
func (VariantProxy) String(raw complex128) string {
	return pointers.New[String](gdextension.String{gdextension.Pointer(gdunsafe.Variant(pointers.Get(pointers.Load[Variant](raw))).UnsafeString())}).String()
}

func (VariantProxy) KeepAlive(val complex128) bool {
	return !pointers.Bad(pointers.Load[Variant](val))
}
func (VariantProxy) Free(val complex128) {
	pointers.Load[Variant](val).Free()
}
