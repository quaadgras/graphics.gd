//go:build cgo

//go:generate go run ./internal/cmd/generate
//go:generate go fmt .
package startup

/*

#include "startup_cgo.h"
*/
import "C"

import (
	"errors"
	"iter"
	"runtime"
	"runtime/cgo"
	"unsafe"

	gd "graphics.gd/internal"
	internal "graphics.gd/internal"
	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Packed"
)

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

type initialization = C.GDExtensionInitialization
type initializationLevel = C.GDExtensionInitializationLevel

func doInitialization(init *initialization) {
	C.initialization(init)
}

//export initialize
func initialize(_ unsafe.Pointer, level initializationLevel) {
	internal.Global.Init(gd.GDExtensionInitializationLevel(level))
	if level == 2 {
		for _, fn := range internal.StartupFunctions {
			fn()
		}
		close(intialized)
		resume_main, stop_main = iter.Pull(call_main_in_steps())
		resume_main()
		for _, fn := range internal.PostStartupFunctions {
			fn()
		}
	}
}

//export deinitialize
func deinitialize(_ unsafe.Pointer, level initializationLevel) {
	if level == 2 {
		for _, cleanup := range internal.Cleanups() {
			cleanup()
		}
		pointers.Cycle()
		pointers.Cycle()
		if theMainFunctionIsWaitingForTheEngineToShutDown {
			resume_main()
		}
	}
}

func get_proc_address(handle uintptr, name string) unsafe.Pointer {
	name = name + "\000"
	char := C.CString(name)
	defer C.free(unsafe.Pointer(char))
	return C.get_proc_address(C.pointer(handle), char)
}

// linkCGO implements the Godot GDExtension API via CGO.
func linkCGO(API *gd.API) {
	get_godot_version := dlsymGD("get_godot_version")
	API.GetGodotVersion = func() gd.Version {
		var version = new(C.GDExtensionGodotVersion)
		C.get_godot_version(C.uintptr_t(uintptr(get_godot_version)), version)
		return gd.Version{
			Major: uint32(version.major),
			Minor: uint32(version.minor),
			Patch: uint32(version.patch),
			Value: C.GoString(version.string),
		}
	}
	variant_new_copy := dlsymGD("variant_new_copy")
	get_variant_from_type_constructor := dlsymGD("get_variant_from_type_constructor")
	API.Variants.FromTypeConstructor = func(vt gd.VariantType) func(ret callframe.Ptr[[3]uint64], arg callframe.Addr) {
		fn := C.get_variant_from_type_constructor(
			C.uintptr_t(uintptr(get_variant_from_type_constructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret callframe.Ptr[[3]uint64], arg callframe.Addr) {
			C.call_variant_from_type_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	get_variant_to_type_constructor := dlsymGD("get_variant_to_type_constructor")
	API.Variants.ToTypeConstructor = func(vt gd.VariantType) func(ret callframe.Addr, arg callframe.Ptr[[3]uint64]) {
		fn := C.get_variant_to_type_constructor(
			C.uintptr_t(uintptr(get_variant_to_type_constructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret callframe.Addr, arg callframe.Ptr[[3]uint64]) {
			C.call_variant_to_type_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_operator_evaluator := dlsymGD("variant_get_ptr_operator_evaluator")
	API.Variants.PointerOperatorEvaluator = func(op gd.Operator, a, b gd.VariantType) func(a, b, ret callframe.Addr) {
		fn := C.variant_get_ptr_operator_evaluator(
			C.uintptr_t(uintptr(variant_get_ptr_operator_evaluator)),
			C.GDExtensionVariantOperator(op),
			C.GDExtensionVariantType(a),
			C.GDExtensionVariantType(b),
		)
		return func(a, b, ret callframe.Addr) {
			C.call_variant_ptr_operator_evaluator(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(a.Uintptr()),
				C.uintptr_t(b.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_builtin_method := dlsymGD("variant_get_ptr_builtin_method")
	API.Variants.GetPointerBuiltinMethod = func(vt gd.VariantType, sn gd.StringName, i gd.Int) func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, pointers.Get(sn))
		fn := C.variant_get_ptr_builtin_method(
			C.uintptr_t(uintptr(variant_get_ptr_builtin_method)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32) {
			C.call_variant_ptr_builtin_method(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(args.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
				C.int32_t(c),
			)
		}
	}
	variant_get_ptr_constructor := dlsymGD("variant_get_ptr_constructor")
	API.Variants.GetPointerConstructor = func(vt gd.VariantType, index int32) func(ret callframe.Addr, args callframe.Args) {
		fn := C.variant_get_ptr_constructor(
			C.uintptr_t(uintptr(variant_get_ptr_constructor)),
			C.GDExtensionVariantType(vt),
			C.int32_t(index),
		)
		return func(ret callframe.Addr, args callframe.Args) {
			C.call_variant_ptr_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(args.Uintptr()),
			)
		}
	}
	variant_get_ptr_destructor := dlsymGD("variant_get_ptr_destructor")
	API.Variants.GetPointerDestructor = func(vt gd.VariantType) func(ret callframe.Addr) {
		fn := C.variant_get_ptr_destructor(
			C.uintptr_t(uintptr(variant_get_ptr_destructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret callframe.Addr) {
			C.call_variant_ptr_destructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_construct := dlsymGD("variant_construct")
	API.Variants.Construct = func(t gd.VariantType, args ...gd.Variant) (gd.Variant, error) {
		var frame = callframe.New()
		for _, arg := range args {
			callframe.Arg(frame, pointers.Get(arg))
		}
		var r_ret = callframe.Ret[[3]uint64](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_construct(
			C.uintptr_t(uintptr(variant_construct)),
			C.GDExtensionVariantType(t),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(frame.Array(0).Uintptr()),
			C.int32_t(len(args)),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = pointers.New[gd.Variant](r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_get_ptr_setter := dlsymGD("variant_get_ptr_setter")
	API.Variants.GetPointerSetter = func(vt gd.VariantType, sn gd.StringName) func(base, arg callframe.Addr) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, pointers.Get(sn))
		fn := C.variant_get_ptr_setter(
			C.uintptr_t(uintptr(variant_get_ptr_setter)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		return func(base, arg callframe.Addr) {
			C.call_variant_ptr_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_getter := dlsymGD("variant_get_ptr_getter")
	API.Variants.GetPointerGetter = func(vt gd.VariantType, sn gd.StringName) func(base, ret callframe.Addr) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, pointers.Get(sn))
		fn := C.variant_get_ptr_getter(
			C.uintptr_t(uintptr(variant_get_ptr_getter)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		return func(base, ret callframe.Addr) {
			C.call_variant_ptr_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_indexed_setter := dlsymGD("variant_get_ptr_indexed_setter")
	API.Variants.GetPointerIndexedSetter = func(vt gd.VariantType) func(base callframe.Addr, index gd.Int, arg callframe.Addr) {
		fn := C.variant_get_ptr_indexed_setter(
			C.uintptr_t(uintptr(variant_get_ptr_indexed_setter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base callframe.Addr, index gd.Int, arg callframe.Addr) {
			C.call_variant_ptr_indexed_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.GDExtensionInt(index),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_indexed_getter := dlsymGD("variant_get_ptr_indexed_getter")
	API.Variants.GetPointerIndexedGetter = func(vt gd.VariantType) func(base callframe.Addr, index gd.Int, ret callframe.Addr) {
		fn := C.variant_get_ptr_indexed_getter(
			C.uintptr_t(uintptr(variant_get_ptr_indexed_getter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base callframe.Addr, index gd.Int, ret callframe.Addr) {
			C.call_variant_ptr_indexed_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.GDExtensionInt(index),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_keyed_setter := dlsymGD("variant_get_ptr_keyed_setter")
	API.Variants.GetPointerKeyedSetter = func(vt gd.VariantType) func(base, key, arg callframe.Addr) {
		fn := C.variant_get_ptr_keyed_setter(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_setter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base, key, arg callframe.Addr) {
			C.call_variant_ptr_keyed_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(key.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_keyed_getter := dlsymGD("variant_get_ptr_keyed_getter")
	API.Variants.GetPointerKeyedGetter = func(vt gd.VariantType) func(base, key, ret callframe.Addr) {
		fn := C.variant_get_ptr_keyed_getter(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_getter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base, key, ret callframe.Addr) {
			C.call_variant_ptr_keyed_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(key.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_keyed_checker := dlsymGD("variant_get_ptr_keyed_checker")
	API.Variants.GetPointerKeyedChecker = func(vt gd.VariantType) func(base, key callframe.Addr) uint32 {
		fn := C.variant_get_ptr_keyed_checker(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_checker)),
			C.GDExtensionVariantType(vt),
		)
		return func(base callframe.Addr, key callframe.Addr) uint32 {
			return uint32(C.call_variant_ptr_keyed_checker(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(key.Uintptr()),
			))
		}
	}
	variant_get_constant_value := dlsymGD("variant_get_constant_value")
	API.Variants.GetConstantValue = func(vt gd.VariantType, sn gd.StringName) gd.Variant {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, pointers.Get(sn))
		var r_ret = callframe.Ret[[3]uint64](frame)
		C.variant_get_constant_value(
			C.uintptr_t(uintptr(variant_get_constant_value)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = pointers.New[gd.Variant](r_ret.Get())
		frame.Free()
		return ret
	}
	variant_get_ptr_utility_function := dlsymGD("variant_get_ptr_utility_function")
	API.Variants.GetPointerUtilityFunction = func(sn gd.StringName, hash gd.Int) func(ret callframe.Addr, args callframe.Args, c int32) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, pointers.Get(sn))
		fn := C.variant_get_ptr_utility_function(
			C.uintptr_t(uintptr(variant_get_ptr_utility_function)),
			C.uintptr_t(p_method.Uintptr()),
			C.GDExtensionInt(hash),
		)
		frame.Free()
		return func(ret callframe.Addr, args callframe.Args, c int32) {
			C.call_variant_ptr_utility_function(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(args.Uintptr()),
				C.int32_t(c),
			)
		}
	}
	string_new_with_utf8_chars_and_len := dlsymGD("string_new_with_utf8_chars_and_len")
	API.Strings.New = func(s string) gd.String {
		var str = C.CString(s)
		var frame = callframe.New()
		var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
		C.string_new_with_utf8_chars_and_len(
			C.uintptr_t(uintptr(string_new_with_utf8_chars_and_len)),
			C.uintptr_t(r_ret.Uintptr()),
			str,
			C.GDExtensionInt(len(s)),
		)
		var ret = pointers.New[gd.String](r_ret.Get())
		frame.Free()
		C.free(unsafe.Pointer(str))
		return ret
	}
	string_to_utf8_chars := dlsymGD("string_to_utf8_chars")
	API.Strings.Get = func(s gd.String) string {
		var length = s.Length()
		if length == 0 {
			return ""
		}
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(s))
		var r_ret = make([]byte, length)
		C.string_to_utf8_chars(
			C.uintptr_t(uintptr(string_to_utf8_chars)),
			C.uintptr_t(p_self.Uintptr()),
			(*C.char)(unsafe.Pointer(&r_ret[0])),
			C.GDExtensionInt(length),
		)
		frame.Free()
		return string(r_ret)
	}
	string_operator_index := dlsymGD("string_operator_index")
	API.Strings.SetIndex = func(s gd.String, index gd.Int, val rune) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(s))
		*C.string_operator_index(
			C.uintptr_t(uintptr(string_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
		) = C.rune(val)
		frame.Free()
	}
	string_operator_index_const := dlsymGD("string_operator_index")
	API.Strings.Index = func(s gd.String, index gd.Int) rune {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(s))
		var ret = C.string_operator_index_const(
			C.uintptr_t(uintptr(string_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
		)
		frame.Free()
		return rune(*ret)
	}
	string_operator_plus_eq_string := dlsymGD("string_operator_plus_eq_string")
	API.Strings.Append = func(s gd.String, other gd.String) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(s))
		var p_other = callframe.Arg(frame, pointers.Get(other))
		C.string_operator_plus_eq_string(
			C.uintptr_t(uintptr(string_operator_plus_eq_string)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_other.Uintptr()),
		)
		pointers.Set(s, p_self.Get())
		frame.Free()
	}
	string_operator_plus_eq_char := dlsymGD("string_operator_plus_eq_char")
	API.Strings.AppendRune = func(s gd.String, other rune) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(s))
		C.string_operator_plus_eq_char(
			C.uintptr_t(uintptr(string_operator_plus_eq_char)),
			C.uintptr_t(p_self.Uintptr()),
			C.rune(other),
		)
		pointers.Set(s, p_self.Get())
		frame.Free()
	}
	string_resize := dlsymGD("string_resize")
	API.Strings.Resize = func(s gd.String, size gd.Int) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(s))
		var length = s.Length()
		C.string_resize(
			C.uintptr_t(uintptr(string_resize)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(size),
		)
		if size < length {
			API.Strings.SetIndex(s, size, 0)
		}
		pointers.Set(s, p_self.Get())
		frame.Free()
	}
	string_name_new_with_utf8_chars_and_len := dlsymGD("string_name_new_with_utf8_chars_and_len")
	API.StringNames.New = func(s string) gd.StringName {
		var str = C.CString(s)
		var frame = callframe.New()
		var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
		C.string_name_new_with_utf8_chars_and_len(
			C.uintptr_t(uintptr(string_name_new_with_utf8_chars_and_len)),
			C.uintptr_t(r_ret.Uintptr()),
			str,
			C.GDExtensionInt(len(s)),
		)
		var ret = pointers.New[gd.StringName](r_ret.Get())
		frame.Free()
		C.free(unsafe.Pointer(str))
		return ret
	}
	xml_parser_open_buffer := dlsymGD("xml_parser_open_buffer")
	API.XMLParser.OpenBuffer = func(x gd.Object, b []byte) error {
		var pin runtime.Pinner
		pin.Pin(&b[0])
		//mmm.Pin[pinner](mmm.Life(x.AsPointer()), &pin, [0]uintptr{})

		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(x))
		var errCode = C.xml_parser_open_buffer(
			C.uintptr_t(uintptr(xml_parser_open_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.size_t(len(b)),
		)
		frame.Free()
		if errCode != 0 {
			return errors.New("xml_parser_open_buffer failed") // TODO godot error code string?
		}
		return nil
	}
	file_access_store_buffer := dlsymGD("file_access_store_buffer")
	API.FileAccess.StoreBuffer = func(f gd.Object, b []byte) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(f))
		C.file_access_store_buffer(
			C.uintptr_t(uintptr(file_access_store_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.uint64_t(len(b)),
		)
		frame.Free()
	}
	file_access_get_buffer := dlsymGD("file_access_get_buffer")
	API.FileAccess.GetBuffer = func(f gd.Object, b []byte) int {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(f))
		var length = C.file_access_get_buffer(
			C.uintptr_t(uintptr(file_access_get_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.uint64_t(len(b)),
		)
		frame.Free()
		return int(length)
	}
	API.PackedByteArray = makePackedFunctions[gd.PackedByteArray, byte]("byte_array")
	API.PackedColorArray = makePackedFunctions[gd.PackedColorArray, gd.Color]("color_array")
	API.PackedFloat32Array = makePackedFunctions[gd.PackedFloat32Array, float32]("float32_array")
	API.PackedFloat64Array = makePackedFunctions[gd.PackedFloat64Array, float64]("float64_array")
	API.PackedInt32Array = makePackedFunctions[gd.PackedInt32Array, int32]("int32_array")
	API.PackedInt64Array = makePackedFunctions[gd.PackedInt64Array, int64]("int64_array")
	packed_string_array_operator_index_const := dlsymGD("packed_string_array_operator_index_const")
	API.PackedStringArray.Index = func(psa gd.PackedStringArray, i gd.Int) gd.String {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(psa))
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_string_array_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return pointers.Let[gd.String]([1]gd.EnginePointer{*(*gd.EnginePointer)(ret)})
	}
	packed_string_array_operator_index := dlsymGD("packed_string_array_operator_index")
	API.PackedStringArray.SetIndex = func(psa gd.PackedStringArray, i gd.Int, v gd.String) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(psa))
		var ptr = C.packed_T_operator_index(
			C.uintptr_t(uintptr(packed_string_array_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		*(*[1]gd.EnginePointer)(ptr) = pointers.Get(v)
		frame.Free()
	}
	API.PackedVector2Array = makePackedFunctions[gd.PackedVector2Array, gd.Vector2]("vector2_array")
	API.PackedVector3Array = makePackedFunctions[gd.PackedVector3Array, gd.Vector3]("vector3_array")
	API.PackedVector4Array = makePackedFunctions[gd.PackedVector4Array, gd.Vector4]("vector4_array")
	dictionary_operator_index := dlsymGD("dictionary_operator_index")
	API.Dictionary.Index = func(d gd.Dictionary, key gd.Variant) gd.Variant {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(d))
		var p_key = callframe.Arg(frame, pointers.Get(key))
		var ptr = C.dictionary_operator_index(
			C.uintptr_t(uintptr(dictionary_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
		)
		var r_ret = *(*[3]uint64)(ptr)
		var ret = pointers.Raw[gd.Variant](r_ret)
		frame.Free()
		return ret
	}
	API.Dictionary.SetIndex = func(dict gd.Dictionary, key, val gd.Variant) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get(dict))
		var p_key = callframe.Arg(frame, pointers.Get(key))
		var ptr = C.dictionary_operator_index(
			C.uintptr_t(uintptr(dictionary_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
		)
		var p_copy = callframe.Ret[[3]uint64](frame)
		var p_value = callframe.Arg(frame, pointers.Get(val))
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(p_copy.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
		)
		*(*[3]uint64)(ptr) = p_copy.Get()
		frame.Free()
	}
	object_get_instance_from_id = dlsymGD("object_get_instance_from_id")
	API.Object.GetInstanceFromID = func(id gd.ObjectID) [1]gd.Object {
		var ret = C.object_get_instance_from_id(
			C.uintptr_t(uintptr(object_get_instance_from_id)),
			C.uintptr_t(id),
		)
		if ret == 0 {
			return [1]gd.Object{}
		}
		return [1]gd.Object{gd.PointerMustAssertInstanceID[gd.Object](gd.EnginePointer(ret))}
	}
	object_method_bind_call := dlsymGD("object_method_bind_call")
	API.Object.MethodBindCall = func(method gd.MethodBind, obj [1]gd.Object, arg ...gd.Variant) (gd.Variant, error) {
		var self = pointers.Get(obj[0])
		if self[0] == 0 {
			return gd.Variant{}, errors.New("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				return gd.Variant{}, errors.New("use after free")
			}
		}
		var frame = callframe.New()
		for _, a := range arg {
			callframe.Arg(frame, pointers.Get(a))
		}
		var r_ret = callframe.Ret[[3]uint64](frame)
		var r_error = new(C.GDExtensionCallError)
		C.object_method_bind_call(
			C.uintptr_t(uintptr(object_method_bind_call)),
			C.uintptr_t(method),
			C.uintptr_t(self[0]),
			C.uintptr_t(frame.Array(0).Uintptr()),
			C.GDExtensionInt(len(arg)),
			C.uintptr_t(r_ret.Uintptr()),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = pointers.New[gd.Variant](r_ret.Get())
		frame.Free()
		return ret, nil
	}
	object_method_bind_ptrcall = dlsymGD("object_method_bind_ptrcall")
	API.Object.MethodBindPointerCall = method_bind_ptrcall
	API.Object.MethodBindPointerCallStatic = func(method gd.MethodBind, arg callframe.Args, ret callframe.Addr) {
		C.object_method_bind_ptrcall(
			C.uintptr_t(uintptr(object_method_bind_ptrcall)),
			C.uintptr_t(method),
			C.uintptr_t(0),
			C.uintptr_t(arg.Uintptr()),
			C.uintptr_t(ret.Uintptr()),
		)
	}
	object_destroy := dlsymGD("object_destroy")
	API.Object.Destroy = func(o [1]gd.Object) {
		if o == [1]gd.Object{} {
			panic("nil gd.Object dereference")
		}
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		C.object_destroy(
			C.uintptr_t(uintptr(object_destroy)),
			C.uintptr_t(self[0]),
		)
	}
	global_get_singleton := dlsymGD("global_get_singleton")
	API.Object.GetSingleton = func(name gd.StringName) [1]gd.Object {
		var frame = callframe.New()
		var p_name = callframe.Arg(frame, pointers.Get(name))
		var ret = C.global_get_singleton(
			C.uintptr_t(uintptr(global_get_singleton)),
			C.uintptr_t(p_name.Uintptr()),
		)
		frame.Free()
		return [1]gd.Object{pointers.Raw[gd.Object]([3]uint64{uint64(ret)})}
	}
	object_get_instance_binding := dlsymGD("object_get_instance_binding")
	API.Object.GetInstanceBinding = func(o [1]gd.Object, et gd.ExtensionToken, ibt gd.InstanceBindingType) any {
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var ret = C.object_get_instance_binding(
			C.uintptr_t(uintptr(object_get_instance_binding)),
			C.uintptr_t(self[0]),
			C.uintptr_t(et),
			unsafe.Pointer(ibt),
		)
		return cgo.Handle(ret).Value()
	}
	object_set_instance_binding := dlsymGD("object_set_instance_binding")
	API.Object.SetInstanceBinding = func(o [1]gd.Object, et gd.ExtensionToken, val any, ibt gd.InstanceBindingType) {
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}

		var pinner runtime.Pinner
		var bindings C.GDExtensionInstanceBindingCallbacks
		pinner.Pin(&bindings)
		defer pinner.Unpin()

		p_val := cgo.NewHandle(val)
		C.object_set_instance_binding(
			C.uintptr_t(uintptr(object_set_instance_binding)),
			C.uintptr_t(self[0]),
			C.uintptr_t(et),
			C.uintptr_t(p_val),
			unsafe.Pointer(&bindings),
		)
	}
	object_free_instance_binding := dlsymGD("object_free_instance_binding")
	API.Object.FreeInstanceBinding = func(o [1]gd.Object, et gd.ExtensionToken) {
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		C.object_free_instance_binding(
			C.uintptr_t(uintptr(object_free_instance_binding)),
			C.uintptr_t(self[0]),
			C.uintptr_t(et),
		)
	}
	object_set_instance := dlsymGD("object_set_instance")
	API.Object.SetInstance = func(o [1]gd.Object, sn gd.StringName, a gd.ObjectInterface) {
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var frame = callframe.New()
		var p_sn = callframe.Arg(frame, pointers.Get(sn))
		var p_val = cgo.NewHandle(a)
		C.object_set_instance(
			C.uintptr_t(uintptr(object_set_instance)),
			C.uintptr_t(self[0]),
			C.uintptr_t(p_sn.Uintptr()),
			C.uintptr_t(p_val),
		)
		frame.Free()
	}
	object_get_class_name := dlsymGD("object_get_class_name")
	API.Object.GetClassName = func(o [1]gd.Object, token gd.ExtensionToken) gd.String {
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var frame = callframe.New()
		var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
		C.object_get_class_name(
			C.uintptr_t(uintptr(object_get_class_name)),
			C.uintptr_t(self[0]),
			C.uintptr_t(token),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = pointers.New[gd.String](r_ret.Get())
		frame.Free()
		return ret
	}
	object_cast_to := dlsymGD("object_cast_to")
	API.Object.CastTo = func(o [1]gd.Object, ct gd.ClassTag) [1]gd.Object {
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			return [1]gd.Object{}
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var ret = C.object_cast_to(
			C.uintptr_t(uintptr(object_cast_to)),
			C.uintptr_t(self[0]),
			C.uintptr_t(ct),
		)
		if ret == 0 {
			return [1]gd.Object{}
		}
		return o // Let?
	}
	object_get_instance_id := dlsymGD("object_get_instance_id")
	API.Object.GetInstanceID = func(o [1]gd.Object) gd.ObjectID {
		var self = pointers.Get(o[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		return gd.ObjectID(C.object_get_instance_id(
			C.uintptr_t(uintptr(object_get_instance_id)),
			C.uintptr_t(self[0]),
		))
	}
	ref_get_object := dlsymGD("ref_get_object")
	API.RefCounted.GetObject = func(rc [1]gd.Object) [1]gd.Object {
		var self = pointers.Get(rc[0])
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var ret = C.ref_get_object(
			C.uintptr_t(uintptr(ref_get_object)),
			C.uintptr_t(self[0]),
		)
		return [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(ret)})}
	}
	ref_set_object := dlsymGD("ref_set_object")
	API.RefCounted.SetObject = func(rc [1]gd.Object, o [1]gd.Object) {
		C.ref_set_object(
			C.uintptr_t(uintptr(ref_set_object)),
			C.uintptr_t(pointers.Get(rc[0])[0]),
			C.uintptr_t(pointers.Get(o[0])[0]),
		)
	}
	classdb_construct_object := dlsymGD("classdb_construct_object")
	API.ClassDB.ConstructObject = func(name gd.StringName) [1]gd.Object {
		var frame = callframe.New()
		defer frame.Free()
		return [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(C.classdb_construct_object(
			C.uintptr_t(uintptr(classdb_construct_object)),
			C.uintptr_t(callframe.Arg(frame, pointers.Get(name)).Uintptr()),
		))})}
	}
	classdb_get_class_tag := dlsymGD("classdb_get_class_tag")
	API.ClassDB.GetClassTag = func(name gd.StringName) gd.ClassTag {
		var frame = callframe.New()
		defer frame.Free()
		return gd.ClassTag(C.classdb_get_class_tag(
			C.uintptr_t(uintptr(classdb_get_class_tag)),
			C.uintptr_t(callframe.Arg(frame, pointers.Get(name)).Uintptr()),
		))
	}
	classdb_get_method_bind := dlsymGD("classdb_get_method_bind")
	API.ClassDB.GetMethodBind = func(class, method gd.StringName, hash gd.Int) gd.MethodBind {
		var frame = callframe.New()
		defer frame.Free()
		return gd.MethodBind(C.classdb_get_method_bind(
			C.uintptr_t(uintptr(classdb_get_method_bind)),
			C.uintptr_t(callframe.Arg(frame, pointers.Get(class)).Uintptr()),
			C.uintptr_t(callframe.Arg(frame, pointers.Get(method)).Uintptr()),
			C.GDExtensionInt(hash),
		))
	}
	get_library_path := dlsymGD("get_library_path")
	API.GetLibraryPath = func(et gd.ExtensionToken) gd.String {
		var frame = callframe.New()
		var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
		C.get_library_path(
			C.uintptr_t(uintptr(get_library_path)),
			C.uintptr_t(et),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = pointers.New[gd.String](r_ret.Get())
		frame.Free()
		return ret
	}
	callable_custom_create := dlsymGD("callable_custom_create")
	API.Callables.Create = func(fn func(...gd.Variant) (gd.Variant, error)) gd.Callable {
		var frame = callframe.New()
		var r_callable = callframe.Ret[[2]uint64](frame)
		var info C.GDExtensionCallableCustomInfo
		*(*uintptr)(unsafe.Pointer(&info.token)) = uintptr(gd.Global.ExtensionToken)
		*(*uintptr)(unsafe.Pointer(&info.callable_userdata)) = uintptr(cgo.NewHandle(fn))
		C.callable_custom_create(
			C.uintptr_t(uintptr(callable_custom_create)),
			C.uintptr_t(r_callable.Uintptr()),
			&info,
		)
		var r_ret = pointers.New[gd.Callable](r_callable.Get())
		frame.Free()
		return r_ret
	}

	classdb_register_extension_class_signal := dlsymGD("classdb_register_extension_class_signal")
	API.ClassDB.RegisterClassSignal = func(library gd.ExtensionToken, class, signal gd.StringName, args []gd.PropertyInfo) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, pointers.Get(class))
		var p_signal = callframe.Arg(frame, pointers.Get(signal))
		p_list, free := cPropertyList(args) //FIXME
		defer free()
		C.classdb_register_extension_class_signal(
			C.uintptr_t(uintptr(classdb_register_extension_class_signal)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_signal.Uintptr()),
			p_list,
			C.GDExtensionInt(len(args)),
		)
		frame.Free()
	}

	classdb_register_extension_class2 := dlsymGD("classdb_register_extension_class2")
	API.ClassDB.RegisterClass = func(library gd.ExtensionToken, name, extends gd.StringName, info gd.ClassInterface) {
		var frame = callframe.New()
		var p_name = callframe.Arg(frame, pointers.Get(name))
		var p_extends = callframe.Arg(frame, pointers.Get(extends))
		var is_virtual C.GDExtensionBool
		if info.IsVirtual() {
			is_virtual = 1
		}
		var is_abstract C.GDExtensionBool
		if info.IsAbstract() {
			is_abstract = 1
		}
		var is_exposed C.GDExtensionBool
		if info.IsExposed() {
			is_exposed = 1
		}
		var p_info = C.GDExtensionClassCreationInfo2{
			is_virtual:  is_virtual,
			is_abstract: is_abstract,
			is_exposed:  is_exposed,
		}
		*(*uintptr)(unsafe.Pointer(&p_info.class_userdata)) = uintptr(cgo.NewHandle(info))
		C.classdb_register_extension_class2(
			C.uintptr_t(uintptr(classdb_register_extension_class2)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_name.Uintptr()),
			C.uintptr_t(p_extends.Uintptr()),
			&p_info,
		)
		frame.Free()
	}

	classdb_register_extension_class_integer_constant := dlsymGD("classdb_register_extension_class_integer_constant")
	API.ClassDB.RegisterClassIntegerConstant = func(library gd.ExtensionToken, class, enum, constant gd.StringName, value int64, bitfield bool) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, pointers.Get(class))
		var p_enum = callframe.Arg(frame, pointers.Get(enum))
		var p_constant = callframe.Arg(frame, pointers.Get(constant))
		C.classdb_register_extension_class_integer_constant(
			C.uintptr_t(uintptr(classdb_register_extension_class_integer_constant)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_enum.Uintptr()),
			C.uintptr_t(p_constant.Uintptr()),
			C.GDExtensionInt(value),
			C.GDExtensionBool(btoi(bitfield)),
		)
		frame.Free()
	}
	classdb_register_extension_class_property := dlsymGD("classdb_register_extension_class_property")
	API.ClassDB.RegisterClassProperty = func(library gd.ExtensionToken, class gd.StringName, info gd.PropertyInfo, getter, setter gd.StringName) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, pointers.Get(class))
		var p_getter = callframe.Arg(frame, pointers.Get(getter))
		var p_setter = callframe.Arg(frame, pointers.Get(setter))
		var pins runtime.Pinner
		defer pins.Unpin()
		var info_name = pointers.Get(info.Name)
		var info_className = pointers.Get(info.ClassName)
		var info_hintString = pointers.Get(info.HintString)
		pins.Pin(&info_name)
		pins.Pin(&info_className)
		pins.Pin(&info_hintString)
		var p_info = C.GDExtensionPropertyInfo{
			_type:       C.GDExtensionVariantType(info.Type),
			name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_name)),
			class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_className)),
			hint:        C.uint32_t(info.Hint),
			hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&info_hintString)),
			usage:       C.uint32_t(info.Usage),
		}
		C.classdb_register_extension_class_property(
			C.uintptr_t(uintptr(classdb_register_extension_class_property)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			&p_info,
			C.uintptr_t(p_getter.Uintptr()),
			C.uintptr_t(p_setter.Uintptr()),
		)
		frame.Free()
	}
	classdb_register_extension_class_property_indexed := dlsymGD("classdb_register_extension_class_property_indexed")
	API.ClassDB.RegisterClassPropertyIndexed = func(library gd.ExtensionToken, class gd.StringName, info gd.PropertyInfo, getter, setter gd.StringName, index int64) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, pointers.Get(class))
		var p_getter = callframe.Arg(frame, pointers.Get(getter))
		var p_setter = callframe.Arg(frame, pointers.Get(setter))
		var pins runtime.Pinner
		defer pins.Unpin()
		var info_name = pointers.Get(info.Name)
		var info_className = pointers.Get(info.ClassName)
		var info_hintString = pointers.Get(info.HintString)
		pins.Pin(&info_name)
		pins.Pin(&info_className)
		pins.Pin(&info_hintString)
		var p_info = C.GDExtensionPropertyInfo{
			_type:       C.GDExtensionVariantType(info.Type),
			name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_name)),
			class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_className)),
			hint:        C.uint32_t(info.Hint),
			hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&info_hintString)),
			usage:       C.uint32_t(info.Usage),
		}
		C.classdb_register_extension_class_property_indexed(
			C.uintptr_t(uintptr(classdb_register_extension_class_property_indexed)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			&p_info,
			C.uintptr_t(p_getter.Uintptr()),
			C.uintptr_t(p_setter.Uintptr()),
			C.GDExtensionInt(index),
		)
		frame.Free()
	}

	classdb_register_extension_class_property_group := dlsymGD("classdb_register_extension_class_property_group")
	API.ClassDB.RegisterClassPropertyGroup = func(library gd.ExtensionToken, class gd.StringName, group, prefix gd.String) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, pointers.Get(class))
		var p_group = callframe.Arg(frame, pointers.Get(group))
		var p_prefix = callframe.Arg(frame, pointers.Get(prefix))
		C.classdb_register_extension_class_property_group(
			C.uintptr_t(uintptr(classdb_register_extension_class_property_group)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_group.Uintptr()),
			C.uintptr_t(p_prefix.Uintptr()),
		)
		frame.Free()
	}
	classdb_register_extension_class_property_subgroup := dlsymGD("classdb_register_extension_class_property_subgroup")
	API.ClassDB.RegisterClassPropertySubGroup = func(library gd.ExtensionToken, class gd.StringName, subGroup, prefix gd.String) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, pointers.Get(class))
		var p_subGroup = callframe.Arg(frame, pointers.Get(subGroup))
		var p_prefix = callframe.Arg(frame, pointers.Get(prefix))
		C.classdb_register_extension_class_property_subgroup(
			C.uintptr_t(uintptr(classdb_register_extension_class_property_subgroup)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_subGroup.Uintptr()),
			C.uintptr_t(p_prefix.Uintptr()),
		)
		frame.Free()
	}
	classdb_register_extension_class_method := dlsymGD("classdb_register_extension_class_method")
	API.ClassDB.RegisterClassMethod = func(library gd.ExtensionToken, class gd.StringName, info gd.Method) {
		infoHandle := cgo.NewHandle(&info)
		gd.RegisterCleanup(func() {
			infoHandle.Delete() // FIXME link this to class registration lifetime?
		})

		var pins runtime.Pinner
		defer pins.Unpin()

		var name = pointers.Get(info.Name)
		pins.Pin(&name)

		var returnInfo *C.GDExtensionPropertyInfo

		var has_return_value C.GDExtensionBool
		if info.ReturnValueInfo != nil {
			has_return_value = 1

			var retName = pointers.Get(info.ReturnValueInfo.Name)
			pins.Pin(&retName)

			var className = pointers.Get(info.ReturnValueInfo.ClassName)
			pins.Pin(&className)

			var hintString = pointers.Get(info.ReturnValueInfo.HintString)
			pins.Pin(&hintString)

			returnInfo = &C.GDExtensionPropertyInfo{
				_type:       C.GDExtensionVariantType(info.ReturnValueInfo.Type),
				name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&retName)),
				class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&className)),
				hint:        C.uint32_t(info.ReturnValueInfo.Hint),
				hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&hintString)),
				usage:       C.uint32_t(info.ReturnValueInfo.Usage),
			}

			pins.Pin(returnInfo)
		}

		var list, free = cPropertyList(info.Arguments)
		defer free()

		var firstMetadata *C.GDExtensionClassMethodArgumentMetadata
		var metadatas = make([]C.GDExtensionClassMethodArgumentMetadata, 0, len(info.ArgumentsMetadata))
		for _, metadata := range info.ArgumentsMetadata {
			metadatas = append(metadatas, C.GDExtensionClassMethodArgumentMetadata(metadata))
		}
		if len(metadatas) > 0 {
			firstMetadata = &metadatas[0]
			pins.Pin(&metadatas[0])
		}

		var firstDefaultArgument *C.GDExtensionVariantPtr
		var defaultArguments = make([]C.GDExtensionVariantPtr, 0, len(info.DefaultArguments))
		for _, arg := range info.DefaultArguments {
			var def = pointers.Get(arg)
			pins.Pin(&def)
			defaultArguments = append(defaultArguments, C.GDExtensionVariantPtr(unsafe.Pointer(&def)))
		}
		if len(defaultArguments) > 0 {
			firstDefaultArgument = &defaultArguments[0]
			pins.Pin(&defaultArguments[0])
		}

		var frame = callframe.New()
		var p_class = callframe.Arg(frame, pointers.Get(class))
		var p_info = C.GDExtensionClassMethodInfo{
			name:                   (C.GDExtensionStringNamePtr)(unsafe.Pointer(&name)),
			method_flags:           C.uint32_t(info.MethodFlags),
			has_return_value:       has_return_value,
			return_value_info:      returnInfo,
			return_value_metadata:  C.GDExtensionClassMethodArgumentMetadata(info.ReturnValueMetadata),
			argument_count:         C.uint32_t(len(info.Arguments)),
			arguments_info:         list,
			arguments_metadata:     firstMetadata,
			default_argument_count: C.uint32_t(len(info.DefaultArguments)),
			default_arguments:      firstDefaultArgument,
		}
		*(*uintptr)(unsafe.Pointer(&p_info.method_userdata)) = uintptr(infoHandle) // FIXME leak
		C.classdb_register_extension_class_method(
			C.uintptr_t(uintptr(classdb_register_extension_class_method)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			&p_info,
		)
		frame.Free()
	}
	classdb_unregister_extension_class := dlsymGD("classdb_unregister_extension_class")
	API.ClassDB.UnregisterClass = func(library gd.ExtensionToken, name gd.StringName) {
		var frame = callframe.New()
		var p_name = callframe.Arg(frame, pointers.Get(name))
		C.classdb_unregister_extension_class(
			C.uintptr_t(uintptr(classdb_unregister_extension_class)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_name.Uintptr()),
		)
		frame.Free()
	}
}

func makePackedFunctions[T gd.Packed[T, V], V Packed.Type](prefix string) gd.PackedFunctionsFor[T, V] {
	var API gd.PackedFunctionsFor[T, V]
	packed_T_operator_index := dlsymGD("packed_" + prefix + "_operator_index")
	API.SetIndex = func(t T, i gd.Int, v V) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get[T, [2]uint64](t))
		var ptr = C.packed_T_operator_index(
			C.uintptr_t(uintptr(packed_T_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		*(*V)(unsafe.Pointer(uintptr(ptr))) = v
		frame.Free()
	}
	packed_T_operator_index_const := dlsymGD("packed_" + prefix + "_operator_index_const")
	API.Index = func(t T, i gd.Int) V {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get[T, [2]uint64](t))
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_T_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return *(*V)(unsafe.Pointer(uintptr(ret)))
	}
	API.CopyAsSlice = func(t T) []V {
		var size = t.Len()
		if size == 0 {
			return nil
		}
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get[T, [2]uint64](t))
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_T_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(0),
		)
		var slice = make([]V, size)
		copy(slice, unsafe.Slice((*V)(unsafe.Pointer(uintptr(ret))), size))
		frame.Free()
		return slice
	}
	API.CopyFromSlice = func(t T, slice []V) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, pointers.Get[T, [2]uint64](t))
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_T_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(0),
		)
		frame.Free()
		copy(unsafe.Slice((*V)(ret), t.Len()), slice)
	}
	return API
}

//export set_func
func set_func(p_instance uintptr, p_name, p_value unsafe.Pointer) bool {
	name := pointers.Let[gd.StringName](*(*[1]gd.EnginePointer)(p_name))
	value := pointers.Let[gd.Variant](*(*[3]uint64)(p_value))
	return cgo.Handle(p_instance).Value().(gd.ObjectInterface).Set(name, value)
}

//export get_func
func get_func(p_instance uintptr, p_name, p_value unsafe.Pointer) bool {
	name := pointers.Let[gd.StringName](*(*[1]gd.EnginePointer)(p_name))
	variant, ok := cgo.Handle(p_instance).Value().(gd.ObjectInterface).Get(name)
	if !ok {
		return false
	}
	*(*[3]uint64)(p_value), _ = pointers.End(variant)
	return true
}

type FIXME *struct{}

func cPropertyList(list []gd.PropertyInfo) (*C.GDExtensionPropertyInfo, func()) {
	if len(list) == 0 {
		return nil, func() {}
	}
	var pins runtime.Pinner
	var slice = make([]C.GDExtensionPropertyInfo, 0, len(list))
	for i := range list {
		property := &list[i]
		name := pointers.Get(property.Name)
		pins.Pin(&name)
		class_name := pointers.Get(property.ClassName)
		pins.Pin(&class_name)
		hint_string := pointers.Get(property.HintString)
		pins.Pin(&hint_string)
		slice = append(slice, C.GDExtensionPropertyInfo{
			_type:       C.GDExtensionVariantType(property.Type),
			name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&name)),
			class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&class_name)),
			hint:        C.uint32_t(property.Hint),
			hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&hint_string)),
			usage:       C.uint32_t(property.Usage),
		})
	}
	pins.Pin(&slice[0])
	return &slice[0], func() {
		pins.Unpin()
	}
}

var propertyLists = make(map[uintptr]func())

//export get_property_list_func
func get_property_list_func(p_instance uintptr, p_length *uint32) *C.GDExtensionPropertyInfo {
	list := cgo.Handle(p_instance).Value().(gd.ObjectInterface).GetPropertyList()
	*p_length = uint32(len(list))
	clist, free := cPropertyList(list)
	propertyLists[p_instance] = free
	return clist
}

//export free_property_list_func
func free_property_list_func(p_instance uintptr, p_properties *C.GDExtensionPropertyInfo) {
	propertyLists[p_instance]()
}

//export property_can_revert_func
func property_can_revert_func(p_instance uintptr, p_name unsafe.Pointer) bool {
	name := pointers.Let[gd.StringName](*(*[1]gd.EnginePointer)(p_name))
	return cgo.Handle(p_instance).Value().(gd.ObjectInterface).PropertyCanRevert(name)
}

//export property_get_revert_func
func property_get_revert_func(p_instance uintptr, p_name, p_value unsafe.Pointer) bool {
	name := pointers.Let[gd.StringName](*(*[1]gd.EnginePointer)(p_name))
	variant, ok := cgo.Handle(p_instance).Value().(gd.ObjectInterface).PropertyGetRevert(name)
	if ok {
		*(*[3]uint64)(p_value), _ = pointers.End(variant)
	}
	return ok
}

//export notification_func
func notification_func(p_instance uintptr, p_notification int32, p_reversed bool) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Notification(p_notification, p_reversed)
}

//export to_string_func
func to_string_func(p_instance uintptr, valid, out unsafe.Pointer) {
	s, ok := cgo.Handle(p_instance).Value().(gd.ObjectInterface).ToString()
	if !ok {
		*(*bool)(valid) = false
		return
	}
	*(*bool)(valid) = true
	*(*[1]gd.EnginePointer)(out) = pointers.Get(s)
}

//export reference_func
func reference_func(p_instance uintptr) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Reference()
}

//export unreference_func
func unreference_func(p_instance uintptr) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Unreference()
}

//export create_instance_func
func create_instance_func(p_class uintptr) uintptr {
	return uintptr(pointers.Get(cgo.Handle(p_class).Value().(gd.ClassInterface).CreateInstance()[0])[0])
}

//export free_instance_func
func free_instance_func(_, p_instance uintptr) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Free()
}

//export get_virtual_call_data_func
func get_virtual_call_data_func(p_class uintptr, p_name unsafe.Pointer) uintptr {
	var name = pointers.Let[gd.StringName](*(*[1]gd.EnginePointer)(p_name))
	virtual := cgo.Handle(p_class).Value().(gd.ClassInterface).GetVirtual(name)
	if virtual == nil {
		return 0
	}
	return uintptr(cgo.NewHandle(virtual))
}

//export call_virtual_with_data_func
func call_virtual_with_data_func(p_instance uintptr, p_name unsafe.Pointer, p_data uintptr, p_args, p_ret unsafe.Pointer) {
	var name = pointers.Let[gd.StringName](*(*[1]gd.EnginePointer)(p_name))
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).CallVirtual(name, cgo.Handle(p_data).Value(), gd.Address(p_args), gd.Address(p_ret))
}

//export get_rid_func
func get_rid_func(p_instance uintptr) C.uint64_t {
	return C.uint64_t(cgo.Handle(p_instance).Value().(gd.ObjectInterface).GetRID())
}

//export callable_call
func callable_call(p_callable uintptr, p_args unsafe.Pointer, count C.GDExtensionInt, p_ret unsafe.Pointer, issue *C.GDExtensionCallError) {
	fn := cgo.Handle(p_callable).Value().(func(...gd.Variant) (gd.Variant, error))
	var slice = unsafe.Slice((**[3]uint64)(p_args), int(count))
	var args = make([]gd.Variant, 0, len(slice))
	for _, elem := range slice {
		args = append(args, pointers.Let[gd.Variant](*elem))
	}
	ret, err := fn(args...)
	if err != nil {
		*issue = C.GDExtensionCallError{}
		issue.error = 7 // TODO no generic error>
		return
	}
	*(*[3]uint64)(p_ret), _ = pointers.End(ret)
	*issue = C.GDExtensionCallError{}
}

//export method_call
func method_call(p_method uintptr, p_instance uintptr, p_args unsafe.Pointer, count C.GDExtensionInt, p_ret unsafe.Pointer, issue *C.GDExtensionCallError) {
	method := cgo.Handle(p_method).Value().(*gd.Method)
	var variants = make([]gd.Variant, 0, int(count))
	for _, elem := range unsafe.Slice((**[3]uint64)(p_args), int(count)) {
		variants = append(variants, pointers.Let[gd.Variant](*elem))
	}
	result, err := method.Call(cgo.Handle(p_instance).Value(), variants...)
	if err != nil {
		issue.error = 7 // TODO no generic error>
		return
	}
	if result != (gd.Variant{}) {
		*(*[3]uint64)(p_ret), _ = pointers.End(result)
	}
}

//export method_ptrcall
func method_ptrcall(p_method uintptr, p_instance uintptr, p_args unsafe.Pointer, p_ret unsafe.Pointer) {
	method := cgo.Handle(p_method).Value().(*gd.Method)
	var instance any
	if p_instance != 0 {
		instance = cgo.Handle(p_instance).Value()
	}
	method.PointerCall(instance, gd.Address(p_args), gd.Address(p_ret))
}

var object_method_bind_ptrcall unsafe.Pointer
var object_get_instance_from_id unsafe.Pointer

//go:linkname method_bind_ptrcall
func method_bind_ptrcall(method gd.MethodBind, obj [1]gd.Object, arg callframe.Args, ret callframe.Addr) {
	if obj == ([1]gd.Object{}) {
		panic("nil gd.Object dereference")
	}
	var self = pointers.Get(obj[0])
	if self[0] == 0 {
		panic("nil gd.Object dereference")
	}
	if self[1] != 0 {
		var ret = C.object_get_instance_from_id(
			C.uintptr_t(uintptr(object_get_instance_from_id)),
			C.uintptr_t(self[1]),
		)
		if ret == 0 {
			panic("use after free")
		}
	}
	C.object_method_bind_ptrcall(
		C.uintptr_t(uintptr(object_method_bind_ptrcall)),
		C.uintptr_t(method),
		C.uintptr_t(self[0]),
		C.uintptr_t(arg.Uintptr()),
		C.uintptr_t(ret.Uintptr()),
	)
}
