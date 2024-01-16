//go:build !generate

package gd

import (
	"context"
	"reflect"
	"unsafe"

	"runtime.link/mmm"
)

type Float = float64
type Int = int64

type String mmm.Pointer[API, String, uintptr]

func (s String) Free() {
	var frame = s.API.newFrame()
	frameSet[uintptr](0, frame, s.Pointer())
	s.API.variant.destruct.String(frame.Args())
	frame.free()
}

func (s String) String() string {
	if s.Pointer() == 0 {
		return ""
	}
	var buf = make([]byte, s.Length())
	var frame = s.API.newFrame()
	frameSet[uintptr](0, frame, s.Pointer())
	s.API.Strings.Get(frame.Get(0), buf)
	frame.free()
	return string(buf)
}

func (Godot *API) StringFromStringName(ctx context.Context, s StringName) String {
	var frame = Godot.newFrame()
	frameSet[uintptr](0, frame, s.Pointer())
	Godot.variant.creation.String[2](frame.Back(), frame.Args())
	var raw = frameGet[uintptr](frame)
	frame.free()
	return mmm.Make[API, String](ctx, Godot, raw)
}

type Vector2 struct {
	X, Y float32
}

type Vector2i struct {
	X, Y int32
}

type Rect2 struct {
	Position Vector2
	Size     Vector2
}

type Rect2i struct {
	Position Vector2i
	Size     Vector2i
}

type Vector3 struct {
	X, Y, Z float32
}

type Vector3i struct {
	X, Y, Z int32
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

type StringName mmm.Pointer[API, StringName, uintptr]

func (Godot *API) StringName(ctx context.Context, s string) StringName {
	var frame = Godot.newFrame()
	Godot.StringNames.New(frame.Back(), s)
	var raw = frameGet[uintptr](frame)
	frame.free()
	return mmm.Make[API, StringName](ctx, Godot, raw)
}

func (Godot *API) StringNameFromString(ctx context.Context, s String) StringName {
	var frame = Godot.newFrame()
	frameSet[uintptr](0, frame, s.Pointer())
	Godot.variant.creation.StringName[2](frame.Back(), frame.Args())
	var raw = frameGet[uintptr](frame)
	frame.free()
	return mmm.Make[API, StringName](ctx, Godot, raw)
}

func (s StringName) Free() {
	var frame = s.API.newFrame()
	frameSet[uintptr](0, frame, s.Pointer())
	s.API.variant.destruct.StringName(frame.Get(0))
	frame.free()
	mmm.MarkFree(s)
}

func (s StringName) String() string {
	if s.Pointer() == 0 {
		return ""
	}
	var tmp = s.API.StringFromStringName(s.Context(), s)
	defer tmp.Free()
	return tmp.String()
}

type NodePath mmm.Pointer[API, NodePath, uintptr]

func (n NodePath) Free() {
	var frame = n.API.newFrame()
	frameSet[uintptr](0, frame, n.Pointer())
	n.API.variant.destruct.NodePath(frame.Get(0))
	frame.free()
	mmm.MarkFree(n)
}

type RID uint64

type Callable mmm.Pointer[API, Callable, [2]uintptr]

func (c Callable) Free() {
	var frame = c.API.newFrame()
	frameSet[[2]uintptr](0, frame, c.Pointer())
	c.API.variant.destruct.Callable(frame.Get(0))
	frame.free()
	mmm.MarkFree(c)
}

type Signal mmm.Pointer[API, Signal, [2]uintptr]

func (s Signal) Free() {
	var frame = s.API.newFrame()
	frameSet[[2]uintptr](0, frame, s.Pointer())
	s.API.variant.destruct.Signal(frame.Get(0))
	frame.free()
	mmm.MarkFree(s)
}

type Dictionary mmm.Pointer[API, Dictionary, uintptr]

func (d Dictionary) Free() {
	var frame = d.API.newFrame()
	frameSet[uintptr](0, frame, d.Pointer())
	d.API.variant.destruct.Dictionary(frame.Get(0))
	frame.free()
	mmm.MarkFree(d)
}

type Array mmm.Pointer[API, Array, uintptr]

func (a Array) Free() {
	var frame = a.API.newFrame()
	frameSet[uintptr](0, frame, a.Pointer())
	a.API.variant.destruct.Array(frame.Get(0))
	frame.free()
	mmm.MarkFree(a)
}

type ArrayOf[T any] Array

type PackedByteArray mmm.Pointer[API, PackedByteArray, [2]uintptr]

func (p PackedByteArray) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedByteArray(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedInt32Array mmm.Pointer[API, PackedInt32Array, [2]uintptr]

func (p PackedInt32Array) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedInt32Array(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedInt64Array mmm.Pointer[API, PackedInt64Array, [2]uintptr]

func (p PackedInt64Array) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedInt64Array(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedFloat32Array mmm.Pointer[API, PackedFloat32Array, [2]uintptr]

func (p PackedFloat32Array) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedFloat32Array(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedFloat64Array mmm.Pointer[API, PackedFloat64Array, [2]uintptr]

func (p PackedFloat64Array) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedFloat64Array(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedStringArray mmm.Pointer[API, PackedStringArray, [2]uintptr]

func (p PackedStringArray) AsSlice() []string {
	var slice = make([]string, p.Size())
	for i := 0; i < len(slice); i++ {
		tmp := p.API.PackedStringArray.Index(&p, int64(i))
		slice[i] = (mmm.Make[API, String, uintptr](nil, p.API, *tmp)).String()
	}
	return slice
}

func (p PackedStringArray) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedStringArray(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedVector2Array mmm.Pointer[API, PackedVector2Array, [2]uintptr]

func (p PackedVector2Array) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedVector2Array(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedVector3Array mmm.Pointer[API, PackedVector3Array, [2]uintptr]

func (p PackedVector3Array) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedVector3Array(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type PackedColorArray mmm.Pointer[API, PackedColorArray, [2]uintptr]

func (p PackedColorArray) Free() {
	var frame = p.API.newFrame()
	frameSet[[2]uintptr](0, frame, p.Pointer())
	p.API.variant.destruct.PackedColorArray(frame.Get(0))
	frame.free()
	mmm.MarkFree(p)
}

type Variant mmm.Pointer[API, Variant, [3]uintptr]

func (s Variant) Free() {
	var frame = s.API.newFrame()
	frameSet[[3]uintptr](0, frame, s.Pointer())
	s.API.Variants.Free(frame.Get(0))
	frame.free()
	mmm.MarkFree(s)
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
