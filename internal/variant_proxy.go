package gd

import (
	"iter"
	"reflect"
	"unsafe"

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
	return variantAsValueType[bool](pointers.Load[Variant](raw), TypeBool)
}
func (VariantProxy) Int(raw complex128) int64 {
	return variantAsValueType[int64](pointers.Load[Variant](raw), TypeInt)
}
func (VariantProxy) Float(raw complex128) float64 {
	return variantAsValueType[float64](pointers.Load[Variant](raw), TypeFloat)
}
func (VariantProxy) Vector2(raw complex128) Vector2Type.XY {
	return variantAsValueType[Vector2Type.XY](pointers.Load[Variant](raw), TypeVector2)
}
func (VariantProxy) Vector2i(raw complex128) Vector2iType.XY {
	return variantAsValueType[Vector2iType.XY](pointers.Load[Variant](raw), TypeVector2i)
}
func (VariantProxy) Rect2(raw complex128) Rect2Type.PositionSize {
	return variantAsValueType[Rect2Type.PositionSize](pointers.Load[Variant](raw), TypeRect2)
}
func (VariantProxy) Rect2i(raw complex128) Rect2iType.PositionSize {
	return variantAsValueType[Rect2iType.PositionSize](pointers.Load[Variant](raw), TypeRect2i)
}
func (VariantProxy) Vector3(raw complex128) Vector3Type.XYZ {
	return variantAsValueType[Vector3Type.XYZ](pointers.Load[Variant](raw), TypeVector3)
}
func (VariantProxy) Vector3i(raw complex128) Vector3iType.XYZ {
	return variantAsValueType[Vector3iType.XYZ](pointers.Load[Variant](raw), TypeVector3i)
}
func (VariantProxy) Transform2D(raw complex128) Transform2DType.OriginXY {
	return variantAsValueType[Transform2DType.OriginXY](pointers.Load[Variant](raw), TypeTransform2D)
}
func (VariantProxy) Vector4(raw complex128) Vector4Type.XYZW {
	return variantAsValueType[Vector4Type.XYZW](pointers.Load[Variant](raw), TypeVector4)
}
func (VariantProxy) Vector4i(raw complex128) Vector4iType.XYZW {
	return variantAsValueType[Vector4iType.XYZW](pointers.Load[Variant](raw), TypeVector4i)
}
func (VariantProxy) Plane(raw complex128) PlaneType.NormalD {
	return variantAsValueType[PlaneType.NormalD](pointers.Load[Variant](raw), TypePlane)
}
func (VariantProxy) Quaternion(raw complex128) QuaternionType.IJKX {
	return variantAsValueType[QuaternionType.IJKX](pointers.Load[Variant](raw), TypeQuaternion)
}
func (VariantProxy) AABB(raw complex128) AABBType.PositionSize {
	return variantAsValueType[AABBType.PositionSize](pointers.Load[Variant](raw), TypeAABB)
}
func (VariantProxy) Basis(raw complex128) BasisType.XYZ {
	return variantAsValueType[BasisType.XYZ](pointers.Load[Variant](raw), TypeBasis)
}
func (VariantProxy) Transform3D(raw complex128) Transform3DType.BasisOrigin {
	return variantAsValueType[Transform3DType.BasisOrigin](pointers.Load[Variant](raw), TypeTransform3D)
}
func (VariantProxy) Projection(raw complex128) ProjectionType.XYZW {
	return variantAsValueType[ProjectionType.XYZW](pointers.Load[Variant](raw), TypeProjection)
}
func (VariantProxy) Color(raw complex128) ColorType.RGBA {
	return variantAsValueType[ColorType.RGBA](pointers.Load[Variant](raw), TypeColor)
}
func (VariantProxy) Interface(raw complex128) interface{} {
	return pointers.Load[Variant](raw).Interface()
}
func (VariantProxy) RID(raw complex128) RIDType.Any {
	return variantAsValueType[RIDType.Any](pointers.Load[Variant](raw), TypeRID)
}
func (VariantProxy) Bytes(raw complex128) []byte {
	packed := variantAsPointerType[PackedByteArray](pointers.Load[Variant](raw), TypePackedByteArray)
	return packed.Bytes()
}

func (VariantProxy) New(val any) complex128 {
	return pointers.Pack(NewVariant(val))
}
func (VariantProxy) NewBool(val bool) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeBool, gdextension.SizeBool, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewInt(val int64) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeInt, gdextension.SizeInt, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewFloat(val float64) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeFloat, gdextension.SizeFloat, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewVector2(val Vector2Type.XY) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeVector2, gdextension.SizeVector2, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewVector2i(val Vector2iType.XY) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeVector2i, gdextension.SizeVector2i, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewRect2(val Rect2Type.PositionSize) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeRect2, gdextension.SizeRect2, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewRect2i(val Rect2iType.PositionSize) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeRect2i, gdextension.SizeRect2i, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewVector3(val Vector3Type.XYZ) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeVector3, gdextension.SizeVector3, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewVector3i(val Vector3iType.XYZ) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeVector3i, gdextension.SizeVector3i, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewTransform2D(val Transform2DType.OriginXY) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeTransform2D, gdextension.SizeTransform2D, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewVector4(val Vector4Type.XYZW) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeVector4, gdextension.SizeVector4, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewVector4i(val Vector4iType.XYZW) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeVector4i, gdextension.SizeVector4i, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewPlane(val PlaneType.NormalD) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypePlane, gdextension.SizePlane, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewQuaternion(val QuaternionType.IJKX) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeQuaternion, gdextension.SizeQuaternion, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewAABB(val AABBType.PositionSize) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeAABB, gdextension.SizeAABB, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewBasis(val BasisType.XYZ) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeBasis, gdextension.SizeBasis, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewTransform3D(val Transform3DType.BasisOrigin) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeTransform3D, gdextension.SizeTransform3D, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewProjection(val ProjectionType.XYZW) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeProjection, gdextension.SizeProjection, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewColor(val ColorType.RGBA) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeColor, gdextension.SizeColor, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}
func (VariantProxy) NewRID(val RIDType.Any) complex128 {
	var ret gdextension.Variant
	ret.LoadNative(TypeRID, gdextension.SizeRID, unsafe.Pointer(&val))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
}

func (VariantProxy) NewBytes(val []byte) complex128 {
	var ret gdextension.Variant
	var arr = gdextension.ToPackedArray(pointers.Get(NewPackedByteSlice(val)))
	ret.LoadNative(TypePackedByteArray, gdextension.SizePackedArray, unsafe.Pointer(&arr))
	return pointers.Pack(pointers.New[Variant]([3]uint64(ret)))
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
func (VariantProxy) Call(complex128, StringType.Readable, ...VariantPkg.Any) VariantPkg.Any {
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
	return pointers.New[String]([1]EnginePointer{EnginePointer(gdextension.Host.Variants.Text(pointers.Get(pointers.Load[Variant](raw))))}).String()
}

func (VariantProxy) KeepAlive(val complex128) bool {
	return !pointers.Bad(pointers.Load[Variant](val))
}
func (VariantProxy) Free(val complex128) {
	pointers.Load[Variant](val).Free()
}
