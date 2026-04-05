package noescape

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

type MethodForClass gdextension.MethodForClass

type Variant gdextension.Variant

func (v *Variant) LoadNative(vtype gdextension.VariantType, size gdextension.Shape, ptr unsafe.Pointer) {
	variant_from_native_noescape(vtype, (*gdextension.Variant)(v), size, ptr)
}

//go:noescape
func variant_from_native_noescape(vtype gdextension.VariantType, result *gdextension.Variant, size gdextension.Shape, ptr unsafe.Pointer)

//go:linkname variant_from_native graphics.gd/internal/noescape.variant_from_native_noescape
func variant_from_native(vtype gdextension.VariantType, result *gdextension.Variant, size gdextension.Shape, ptr unsafe.Pointer) {
	*result = gdunsafe.VariantUnsafeFromNative(gdunsafe.VariantType(vtype), uint64(gdextension.SizeVariant|size<<4), ptr)
}

func LoadNative[T gdextension.AnyVariant](vtype gdextension.VariantType, variant gdextension.Variant) T {
	var result T
	variant_into_native_noescape(vtype, variant, unsafe.Pointer(&result), gdextension.SizeOf[T]())
	return result
}

//go:noescape
func variant_into_native_noescape(vtype gdextension.VariantType, variant gdextension.Variant, ptr unsafe.Pointer, size gdextension.Shape)

//go:linkname variant_into_native graphics.gd/internal/noescape.variant_into_native_noescape
func variant_into_native(vtype gdextension.VariantType, variant gdextension.Variant, ptr unsafe.Pointer, size gdextension.Shape) {
	gdunsafe.VariantUnsafeMakeNative(gdunsafe.VariantType(vtype), gdunsafe.Variant(variant), uint64(size|gdextension.SizeVariant<<4), ptr)
}

func Free[T gdextension.AnyVariant](vtype gdextension.VariantType, val *T) {
	free_noescape(vtype, gdextension.SizeOf[T](), unsafe.Pointer(val))
}

//go:noescape
func free_noescape(vtype gdextension.VariantType, size gdextension.Shape, ptr unsafe.Pointer)

//go:linkname free graphics.gd/internal/noescape.free_noescape
func free(vtype gdextension.VariantType, size gdextension.Shape, ptr unsafe.Pointer) {
	gdextension.Host.Builtin.Types.Unsafe.Free(vtype, size<<4, gdextension.CallAccepts[any](ptr))
}

func Make[T gdextension.AnyVariant](constructor gdextension.FunctionID, size gdextension.Shape, ptr unsafe.Pointer) T {
	var result T
	make_native_noescape(constructor, unsafe.Pointer(&result), gdextension.SizeOf[T]()|size, ptr)
	return result
}

//go:noescape
func make_native_noescape(constructor gdextension.FunctionID, result unsafe.Pointer, shape gdextension.Shape, ptr unsafe.Pointer)

//go:linkname make_native graphics.gd/internal/noescape.make_native_noescape
func make_native(constructor gdextension.FunctionID, result unsafe.Pointer, shape gdextension.Shape, ptr unsafe.Pointer) {
	gdextension.Host.Builtin.Types.Unsafe.Make(constructor, gdextension.CallReturns[any](unsafe.Pointer(result)), shape, gdextension.CallAccepts[any](ptr))
}

func IndexPacked[T gdextension.Packable](access func(p gdextension.PackedArray[T], idx int, result gdextension.CallReturns[T]), arr gdextension.PackedArray[T], index int) T {
	var simple = *(*func(p gdextension.PackedArray[byte], idx int, result unsafe.Pointer))(unsafe.Pointer(&access))
	var result T
	index_packed_noescape(simple, gdextension.PackedArray[byte](arr), index, unsafe.Pointer(&result))
	return result
}

//go:noescape
func index_packed_noescape(access func(p gdextension.PackedArray[byte], idx int, result unsafe.Pointer), p gdextension.PackedArray[byte], index int, result unsafe.Pointer)

//go:linkname index_packed graphics.gd/internal/noescape.index_packed_noescape
func index_packed(access func(p gdextension.PackedArray[byte], idx int, result unsafe.Pointer), p gdextension.PackedArray[byte], index int, result unsafe.Pointer) {
	access(p, index, result)
}

func CallStatic[T any](method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return Call[T](0, method, shape, args)
}
