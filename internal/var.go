//go:build !generate

package gd

import (
	"reflect"
	"strings"
	"unsafe"

	"runtime.link/api/call"
	"runtime.link/mmm"
)

type Bool = bool
type Float = float64
type Int = int64

type Rect2 struct {
	Position Vector2
	Size     Vector2
}

type Rect2i struct {
	Position Vector2i
	Size     Vector2i
}

type Transform2D struct {
	X Vector2
	Y Vector2
	O Vector2
}

type Vector4 struct {
	X, Y, Z, W float32
}

type Vector4i struct {
	X, Y, Z, W int32
}

type Plane struct {
	Matrix Vector4
}

type Quaternion struct {
	X, Y, Z, W float32
}

type AABB struct {
	Position Vector3
	Size     Vector3
}

type Basis struct {
	Rows [3]Vector3
}

type Transform3D struct {
	Basis  Basis
	Origin Vector3
}

const _ = unsafe.Sizeof(Transform3D{})

type Projection struct {
	X, Y, Z, W Vector4
}

type Color struct {
	R, G, B, A float32
}

type RID uint64

type Callable mmm.Pointer[API, Callable, [2]uintptr]

func (c Callable) Free() {
	var frame = call.New()
	mmm.API(c).typeset.destruct.Callable(call.Arg(frame, mmm.End(c)))
	frame.Free()
}

type Signal mmm.Pointer[API, Signal, [2]uintptr]

type IsSignal interface{ signal() }

func (s Signal) signal() {}

func (s Signal) Free() {
	var frame = call.New()
	mmm.API(s).typeset.destruct.Signal(call.Arg(frame, mmm.End(s)))
	frame.Free()
}

type Dictionary mmm.Pointer[API, Dictionary, uintptr]

func (d Dictionary) Index(ctx Context, key Variant) Variant {
	return mmm.API(d).Dictionary.Index(ctx, d, key)
}

func (d Dictionary) SetIndex(key Variant, value Variant) {
	mmm.API(d).Dictionary.SetIndex(d, key, value)
}

func (d Dictionary) Free() {
	var frame = call.New()
	mmm.API(d).typeset.destruct.Dictionary(call.Arg(frame, mmm.End(d)))
	frame.Free()
}

type Array mmm.Pointer[API, Array, uintptr]

func (a Array) Index(ctx Context, index Int) Variant {
	return mmm.API(a).Array.Index(ctx, a, index)
}

func (a Array) SetIndex(index Int, value Variant) {
	mmm.API(a).Array.SetIndex(a, index, value)
}

func (a Array) Free() {
	var frame = call.New()
	mmm.API(a).typeset.destruct.Array(call.Arg(frame, mmm.End(a)))
	frame.Free()
}

type ArrayOf[T any] Array

type PackedByteArray mmm.Pointer[API, PackedByteArray, [2]uintptr]

func (p PackedByteArray) Index(idx Int) byte {
	return mmm.API(p).PackedByteArray.Index(p, idx)
}

func (p PackedByteArray) SetIndex(idx Int, value byte) {
	mmm.API(p).PackedByteArray.SetIndex(p, idx, value)
}

func (p PackedByteArray) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedByteArray(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedInt32Array mmm.Pointer[API, PackedInt32Array, [2]uintptr]

func (p PackedInt32Array) Index(idx Int) int32 {
	return mmm.API(p).PackedInt32Array.Index(p, idx)
}

func (p PackedInt32Array) SetIndex(idx Int, value int32) {
	mmm.API(p).PackedInt32Array.SetIndex(p, idx, value)
}

func (p PackedInt32Array) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedInt32Array(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedInt64Array mmm.Pointer[API, PackedInt64Array, [2]uintptr]

func (p PackedInt64Array) Index(idx Int) int64 {
	return mmm.API(p).PackedInt64Array.Index(p, idx)
}

func (p PackedInt64Array) SetIndex(idx Int, value int64) {
	mmm.API(p).PackedInt64Array.SetIndex(p, idx, value)
}

func (p PackedInt64Array) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedInt64Array(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedFloat32Array mmm.Pointer[API, PackedFloat32Array, [2]uintptr]

func (p PackedFloat32Array) Index(idx Int) float32 {
	return mmm.API(p).PackedFloat32Array.Index(p, idx)
}

func (p PackedFloat32Array) SetIndex(idx Int, value float32) {
	mmm.API(p).PackedFloat32Array.SetIndex(p, idx, value)
}

func (p PackedFloat32Array) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedFloat32Array(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedFloat64Array mmm.Pointer[API, PackedFloat64Array, [2]uintptr]

func (p PackedFloat64Array) Index(idx Int) float64 {
	return mmm.API(p).PackedFloat64Array.Index(p, idx)
}

func (p PackedFloat64Array) SetIndex(idx Int, value float64) {
	mmm.API(p).PackedFloat64Array.SetIndex(p, idx, value)
}

func (p PackedFloat64Array) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedFloat64Array(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedStringArray mmm.Pointer[API, PackedStringArray, [2]uintptr]

func (p PackedStringArray) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	size := int(p.Size())
	ctx := NewContext(mmm.API(p))
	defer ctx.End()
	for i := 0; i < size; i++ {
		builder.WriteString(p.Index(ctx, Int(i)).String())
		if i < size-1 {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (p PackedStringArray) Index(ctx Context, idx Int) String {
	return mmm.API(p).PackedStringArray.Index(ctx, p, idx)
}

func (p PackedStringArray) SetIndex(idx Int, value String) {
	mmm.API(p).PackedStringArray.SetIndex(p, idx, value)
}

func (p PackedStringArray) AsSlice(ctx Context) []String {
	return mmm.API(p).PackedStringArray.CopyAsSlice(ctx, p)
}

func (p PackedStringArray) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedStringArray(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedVector2Array mmm.Pointer[API, PackedVector2Array, [2]uintptr]

func (p PackedVector2Array) Index(idx Int) Vector2 {
	return mmm.API(p).PackedVector2Array.Index(p, idx)
}

func (p PackedVector2Array) SetIndex(idx Int, value Vector2) {
	mmm.API(p).PackedVector2Array.SetIndex(p, idx, value)
}

func (p PackedVector2Array) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedVector2Array(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedVector3Array mmm.Pointer[API, PackedVector3Array, [2]uintptr]

func (p PackedVector3Array) Index(idx Int) Vector3 {
	return mmm.API(p).PackedVector3Array.Index(p, idx)
}

func (p PackedVector3Array) SetIndex(idx Int, value Vector3) {
	mmm.API(p).PackedVector3Array.SetIndex(p, idx, value)
}

func (p PackedVector3Array) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedVector3Array(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type PackedColorArray mmm.Pointer[API, PackedColorArray, [2]uintptr]

func (p PackedColorArray) Index(idx Int) Color {
	return mmm.API(p).PackedColorArray.Index(p, idx)
}

func (p PackedColorArray) SetIndex(idx Int, value Color) {
	mmm.API(p).PackedColorArray.SetIndex(p, idx, value)
}

func (p PackedColorArray) Free() {
	var frame = call.New()
	mmm.API(p).typeset.destruct.PackedColorArray(call.Arg(frame, mmm.End(p)))
	frame.Free()
}

type Variant mmm.Pointer[API, Variant, [3]uintptr]

func (s Variant) Free() {
	mmm.API(s).Variants.Destroy(s)
	mmm.End(s)
}

func (self Object) AsObject() Object {
	return self
}

type Iterator struct {
	self Variant
	iter Variant
}

func (iter Iterator) Next() bool {
	return mmm.API(iter.self).Variants.IteratorNext(iter.self, iter.iter)
}

func (iter Iterator) Value(ctx Context) Variant {
	val, ok := mmm.API(iter.self).Variants.IteratorGet(ctx, iter.self, iter.iter)
	if !ok {
		panic("failed to get iterator value")
	}
	return val
}

func variantTypeFromName(s string) (VariantType, reflect.Type) {
	switch s {
	case "Nil":
		return TypeNil, nil
	case "bool", "Bool":
		return TypeBool, reflect.TypeOf(false)
	case "int", "Int":
		return TypeInt, reflect.TypeOf(int64(0))
	case "float", "Float":
		return TypeFloat, reflect.TypeOf(Float(0))
	case "String":
		return TypeString, reflect.TypeOf(String{})
	case "Vector2":
		return TypeVector2, reflect.TypeOf(Vector2{})
	case "Vector2i":
		return TypeVector2i, reflect.TypeOf(Vector2i{})
	case "Rect2":
		return TypeRect2, reflect.TypeOf(Rect2{})
	case "Rect2i":
		return TypeRect2i, reflect.TypeOf(Rect2i{})
	case "Vector3":
		return TypeVector3, reflect.TypeOf(Vector3{})
	case "Vector3i":
		return TypeVector3i, reflect.TypeOf(Vector3i{})
	case "Transform2D":
		return TypeTransform2d, reflect.TypeOf(Transform2D{})
	case "Vector4":
		return TypeVector4, reflect.TypeOf(Vector4{})
	case "Vector4i":
		return TypeVector4i, reflect.TypeOf(Vector4i{})
	case "Plane":
		return TypePlane, reflect.TypeOf(Plane{})
	case "Quaternion":
		return TypeQuaternion, reflect.TypeOf(Quaternion{})
	case "AABB":
		return TypeAabb, reflect.TypeOf(AABB{})
	case "Basis":
		return TypeBasis, reflect.TypeOf(Basis{})
	case "Transform3D":
		return TypeTransform3d, reflect.TypeOf(Transform3D{})
	case "Projection":
		return TypeProjection, reflect.TypeOf(Projection{})
	case "Color":
		return TypeColor, reflect.TypeOf(Color{})
	case "StringName":
		return TypeStringName, reflect.TypeOf(StringName{})
	case "NodePath":
		return TypeNodePath, reflect.TypeOf(NodePath{})
	case "RID":
		return TypeRid, reflect.TypeOf(RID(0))
	case "Object":
		return TypeObject, reflect.TypeOf(uintptr(0))
	case "Callable":
		return TypeCallable, reflect.TypeOf(Callable{})
	case "Signal":
		return TypeSignal, reflect.TypeOf(Signal{})
	case "Dictionary":
		return TypeDictionary, reflect.TypeOf(Dictionary{})
	case "Array":
		return TypeArray, reflect.TypeOf(Array{})
	case "PackedByteArray":
		return TypePackedByteArray, reflect.TypeOf(PackedByteArray{})
	case "PackedInt32Array":
		return TypePackedInt32Array, reflect.TypeOf(PackedInt32Array{})
	case "PackedInt64Array":
		return TypePackedInt64Array, reflect.TypeOf(PackedInt64Array{})
	case "PackedFloat32Array":
		return TypePackedFloat32Array, reflect.TypeOf(PackedFloat32Array{})
	case "PackedFloat64Array":
		return TypePackedFloat64Array, reflect.TypeOf(PackedFloat64Array{})
	case "PackedStringArray":
		return TypePackedStringArray, reflect.TypeOf(PackedStringArray{})
	case "PackedVector2Array":
		return TypePackedVector2Array, reflect.TypeOf(PackedVector2Array{})
	case "PackedVector3Array":
		return TypePackedVector3Array, reflect.TypeOf(PackedVector3Array{})
	case "PackedColorArray":
		return TypePackedColorArray, reflect.TypeOf(PackedColorArray{})
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}

func operatoTypeFromName(name string) Operator {
	switch name {
	case "Equals":
		return Equal
	case "NotEqual":
		return NotEqual
	case "Less":
		return Less
	case "LessEqual":
		return LessEqual
	case "Greater":
		return Greater
	case "GreaterEqual":
		return GreaterEqual
	case "Add":
		return Add
	case "Subtract":
		return Subtract
	case "Multiply":
		return Multiply
	case "Divide":
		return Divide
	case "Negate":
		return Negate
	case "Module":
		return Module
	case "Power":
		return Power
	case "ShiftLeft":
		return ShiftLeft
	case "ShiftRight":
		return ShiftRight
	case "BitAnd":
		return BitAnd
	case "BitOr":
		return BitOr
	case "BitXor":
		return BitXor
	case "BitNegate":
		return BitNegate
	case "And":
		return LogicalAnd
	case "Or":
		return LogicalOr
	case "Xor":
		return LogicalXor
	case "Not":
		return LogicalNegate
	case "In":
		return In
	default:
		panic("gdextension.operatoTypeFromName: unknown operator " + name)
	}
}
