package gdextension

import (
	"unsafe"
)

func (method MethodForClass) Call(self Object, args ...Variant) (Variant, error) {
	var result Variant
	var err CallError
	object_method_call_noescape(self, method, &result, args, &err)
	return result, err.Err()
}

//go:noescape
func object_method_call_noescape(object Object, method MethodForClass, result *Variant, args []Variant, err *CallError)

//go:linkname object_method_call graphics.gd/internal/gdextension.object_method_call_noescape
func object_method_call(object Object, method MethodForClass, result *Variant, args []Variant, err *CallError) {
	Host.Objects.Call(object, method, CallReturns[Variant](result), len(args), CallAccepts[Variant](unsafe.SliceData(args)), CallReturns[CallError](err))
}

func Call[T any](object Object, method MethodForClass, shape Shape, args unsafe.Pointer) T {
	var result T
	call_noescape(object, method, unsafe.Pointer(&result), shape, args)
	return result
}

func CallStatic[T any](method MethodForClass, shape Shape, args unsafe.Pointer) T {
	var result T
	call_noescape(0, method, unsafe.Pointer(&result), shape, args)
	return result
}

//go:noescape
func call_noescape(object Object, method MethodForClass, result unsafe.Pointer, shape Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/gdextension.call_noescape
func call(object Object, method MethodForClass, result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	Host.Objects.Unsafe.Call(object, method, CallReturns[any](result), shape, CallAccepts[any](args))
}

func (v *Variant) LoadNative(vtype VariantType, size Shape, ptr unsafe.Pointer) {
	variant_from_native_noescape(vtype, v, size, ptr)
}

//go:noescape
func variant_from_native_noescape(vtype VariantType, result *Variant, size Shape, ptr unsafe.Pointer)

//go:linkname variant_from_native graphics.gd/internal/gdextension.variant_from_native_noescape
func variant_from_native(vtype VariantType, result *Variant, size Shape, ptr unsafe.Pointer) {
	Host.Variants.Unsafe.FromNative(vtype, CallReturns[Variant](result), SizeVariant|size<<4, CallAccepts[any](ptr))
}

func LoadNative[T any](vtype VariantType, variant Variant) T {
	var result T
	variant_into_native_noescape(vtype, variant, unsafe.Pointer(&result), SizeOf[T]())
	return result
}

//go:noescape
func variant_into_native_noescape(vtype VariantType, variant Variant, ptr unsafe.Pointer, size Shape)

//go:linkname variant_into_native graphics.gd/internal/gdextension.variant_into_native_noescape
func variant_into_native(vtype VariantType, variant Variant, ptr unsafe.Pointer, size Shape) {
	Host.Variants.Unsafe.MakeNative(vtype, variant, SizeVariant|size<<4, CallReturns[any](ptr))
}
