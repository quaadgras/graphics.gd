package gdextension

import (
	"reflect"
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

func Call[T any](object Object, method MethodForClass, shape Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
	return result
}

func CallStatic[T any](method MethodForClass, shape Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	call_noescape(0, method, unsafe.Pointer(&result), shape, argptr)
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

func LoadNative[T AnyVariant](vtype VariantType, variant Variant) T {
	var result T
	variant_into_native_noescape(vtype, variant, unsafe.Pointer(&result), SizeOf[T]())
	return result
}

//go:noescape
func variant_into_native_noescape(vtype VariantType, variant Variant, ptr unsafe.Pointer, size Shape)

//go:linkname variant_into_native graphics.gd/internal/gdextension.variant_into_native_noescape
func variant_into_native(vtype VariantType, variant Variant, ptr unsafe.Pointer, size Shape) {
	Host.Variants.Unsafe.MakeNative(vtype, variant, size|SizeVariant<<4, CallReturns[any](ptr))
}

func Free[T AnyVariant](vtype VariantType, val *T) {
	free_noescape(vtype, SizeOf[T](), unsafe.Pointer(val))
}

//go:noescape
func free_noescape(vtype VariantType, size Shape, ptr unsafe.Pointer)

//go:linkname free graphics.gd/internal/gdextension.free_noescape
func free(vtype VariantType, size Shape, ptr unsafe.Pointer) {
	Host.Builtin.Types.Unsafe.Free(vtype, size<<4, CallAccepts[any](ptr))
}

func Make[T AnyVariant](constructor FunctionID, size Shape, ptr unsafe.Pointer) T {
	var result T
	make_native_noescape(constructor, unsafe.Pointer(&result), SizeOf[T]()|size, ptr)
	return result
}

//go:noescape
func make_native_noescape(constructor FunctionID, result unsafe.Pointer, shape Shape, ptr unsafe.Pointer)

//go:linkname make_native graphics.gd/internal/gdextension.make_native_noescape
func make_native(constructor FunctionID, result unsafe.Pointer, shape Shape, ptr unsafe.Pointer) {
	Host.Builtin.Types.Unsafe.Make(constructor, CallReturns[any](unsafe.Pointer(result)), shape, CallAccepts[any](ptr))
}

func IndexPacked[T Packable](access func(p PackedArray[T], idx int, result CallReturns[T]), arr PackedArray[T], index int) T {
	var simple = *(*func(p PackedArray[byte], idx int, result unsafe.Pointer))(unsafe.Pointer(&access))
	var result T
	index_packed_noescape(simple, PackedArray[byte](arr), index, unsafe.Pointer(&result))
	return result
}

//go:noescape
func index_packed_noescape(access func(p PackedArray[byte], idx int, result unsafe.Pointer), p PackedArray[byte], index int, result unsafe.Pointer)

//go:linkname index_packed graphics.gd/internal/gdextension.index_packed_noescape
func index_packed(access func(p PackedArray[byte], idx int, result unsafe.Pointer), p PackedArray[byte], index int, result unsafe.Pointer) {
	access(p, index, result)
}
