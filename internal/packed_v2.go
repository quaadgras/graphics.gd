package gd

import (
	"graphics.gd/internal/pointers"
	ArrayType "graphics.gd/variant/Array"
	ArrayVariant "graphics.gd/variant/Array"
	PackedType "graphics.gd/variant/Packed"
	StringType "graphics.gd/variant/String"
)

func (p PackedByteArray) New() PackedByteArray       { return NewPackedByteArray() }
func (p PackedFloat32Array) New() PackedFloat32Array { return NewPackedFloat32Array() }
func (p PackedFloat64Array) New() PackedFloat64Array { return NewPackedFloat64Array() }
func (p PackedInt32Array) New() PackedInt32Array     { return NewPackedInt32Array() }
func (p PackedInt64Array) New() PackedInt64Array     { return NewPackedInt64Array() }
func (p PackedColorArray) New() PackedColorArray     { return NewPackedColorArray() }
func (p PackedVector2Array) New() PackedVector2Array { return NewPackedVector2Array() }
func (p PackedVector3Array) New() PackedVector3Array { return NewPackedVector3Array() }
func (p PackedVector4Array) New() PackedVector4Array { return NewPackedVector4Array() }

func InternalPacked[P Packed[P, V], V PackedType.Type](array PackedType.Array[V]) P {
	_, state := ArrayVariant.As(ArrayType.Contains[V](array), NewPackedProxy[P, V])
	return pointers.Load[P, PackedPointers](state)
}

func InternalPackedStrings(array PackedType.Strings) PackedStringArray {
	_, state := ArrayVariant.As(ArrayType.Contains[StringType.Readable](array), NewPackedStringProxy)
	return pointers.Load[PackedStringArray](state)
}

func NewPackedProxy[P Packed[P, V], V PackedType.Type]() (PackedProxy[P, V], complex128) {
	array := [1]P{}[0].New()
	return PackedProxy[P, V]{}, pointers.Pack[P, PackedPointers](array)
}

type PackedProxy[P Packed[P, V], V PackedType.Type] struct{}

func (PackedProxy[P, V]) Any(raw complex128) ArrayType.Any {
	panic("cannot convert packed array to any array! NOT IMPLEMENTED")
}
func (PackedProxy[P, V]) Resize(raw complex128, n int) {
	pointers.Load[P, PackedPointers](raw).Resize(Int(n))
}
func (PackedProxy[P, V]) Index(raw complex128, i int) V {
	return pointers.Load[P, PackedPointers](raw).Index(Int(i))
}
func (PackedProxy[P, V]) SetIndex(raw complex128, i int, v V) {
	pointers.Load[P, PackedPointers](raw).SetIndex(Int(i), v)
}
func (PackedProxy[P, V]) Len(raw complex128) int {
	return int(pointers.Load[P, PackedPointers](raw).Len())
}
func (PackedProxy[P, V]) IsReadOnly(complex128) bool {
	return false
}
func (PackedProxy[P, V]) MakeReadOnly(complex128) {}

type PackedStringArrayProxy struct{}

func NewPackedStringProxy() (PackedStringArrayProxy, complex128) {
	return PackedStringArrayProxy{}, pointers.Pack(NewPackedStringArray())
}

func (PackedStringArrayProxy) Any(raw complex128) ArrayType.Any {
	panic("cannot convert packed array to any array! NOT IMPLEMENTED")
}
func (PackedStringArrayProxy) Resize(raw complex128, n int) {
	pointers.Load[PackedStringArray](raw).Resize(Int(n))
}
func (PackedStringArrayProxy) Index(raw complex128, i int) StringType.Readable {
	s := pointers.Load[PackedStringArray](raw).Index(Int(i))
	return StringType.Via(StringProxy{}, pointers.Pack(s))
}
func (PackedStringArrayProxy) SetIndex(raw complex128, i int, v StringType.Readable) {
	_, s := StringType.Proxy(v, StringCacheCheck, NewStringProxy)
	pointers.Load[PackedStringArray](raw).SetIndex(Int(i), pointers.Load[String](s))
}
func (PackedStringArrayProxy) Len(raw complex128) int {
	return int(pointers.Load[PackedStringArray](raw).Size())
}
func (PackedStringArrayProxy) IsReadOnly(complex128) bool {
	return false
}
func (PackedStringArrayProxy) MakeReadOnly(complex128) {}
