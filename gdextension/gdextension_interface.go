package gdextension

/*
#include <stdlib.h>
#include "gdextension_interface.h"

typedef uintptr_t pointer;
typedef const char* string;

void get_godot_version(pointer fn, GDExtensionGodotVersion *r_version) {
	((GDExtensionInterfaceGetGodotVersion)fn)(r_version);
}
void *mem_alloc(pointer fn, size_t size) {
	return ((GDExtensionInterfaceMemAlloc)fn)(size);
}
void *mem_realloc(pointer fn, void *ptr, size_t size) {
	return ((GDExtensionInterfaceMemRealloc)fn)(ptr, size);
}
void mem_free(pointer fn, void *ptr) {
	((GDExtensionInterfaceMemFree)fn)(ptr);
}
void print_error(pointer fn, string p_description, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintError)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
void print_error_with_message(pointer fn, string p_description, string p_message, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintErrorWithMessage)fn)(p_description, p_message, p_function, p_file, p_line, p_notify_editor);
}
void print_warning(pointer fn, string p_description, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintWarning)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
uint64_t get_native_struct_size(pointer fn, pointer p_classname) {
	return ((GDExtensionInterfaceGetNativeStructSize)fn)((GDExtensionConstStringNamePtr)p_classname);
}
void variant_new_copy(pointer fn, pointer r_dest, pointer p_self) {
	((GDExtensionInterfaceVariantNewCopy)fn)((GDExtensionUninitializedVariantPtr)r_dest, (GDExtensionConstVariantPtr)p_self);
}
void variant_new_nil(pointer fn, pointer r_dest) {
	((GDExtensionInterfaceVariantNewNil)fn)((GDExtensionUninitializedVariantPtr)r_dest);
}
void variant_destroy(pointer fn, pointer p_self) {
	((GDExtensionInterfaceVariantDestroy)fn)((GDExtensionVariantPtr)p_self);
}
void variant_call(pointer fn, pointer p_self, pointer p_method, pointer p_args, GDExtensionInt p_argument_count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCall)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
void variant_call_static(pointer fn, GDExtensionVariantType p_type, pointer p_method, pointer p_args, GDExtensionInt p_argument_count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCallStatic)fn)((GDExtensionVariantType)p_type, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
void variant_evaluate(pointer fn, GDExtensionVariantOperator p_operator, pointer p_a, pointer p_b, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantEvaluate)fn)(p_operator, (GDExtensionConstVariantPtr)p_a, (GDExtensionConstVariantPtr)p_b, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_set(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSet)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
void variant_set_named(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSetNamed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
void variant_set_keyed(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSetKeyed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
void variant_set_indexed(pointer fn, pointer p_self, GDExtensionInt p_index, pointer p_value, pointer r_valid, pointer r_oob) {
	((GDExtensionInterfaceVariantSetIndexed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionInt)p_index, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid, (GDExtensionBool*)r_oob);
}
void variant_get(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGet)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_get_named(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGetNamed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_get_keyed(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGetKeyed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_get_indexed(pointer fn, pointer p_self, GDExtensionInt p_index, pointer r_ret, pointer r_valid, pointer r_oob) {
	((GDExtensionInterfaceVariantGetIndexed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionInt)p_index, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid, (GDExtensionBool*)r_oob);
}
void variant_iter_init(pointer fn, pointer p_self, pointer r_iter, pointer r_valid) {
	((GDExtensionInterfaceVariantIterInit)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedVariantPtr)r_iter, (GDExtensionBool*)r_valid);
}
void variant_iter_next(pointer fn, pointer p_self, pointer p_iter, pointer r_valid) {
	((GDExtensionInterfaceVariantIterNext)fn)((GDExtensionVariantPtr)p_self, (GDExtensionVariantPtr)p_iter, (GDExtensionBool*)r_valid);
}
void variant_iter_get(pointer fn, pointer p_self, pointer p_iter, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantIterGet)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionVariantPtr)p_iter, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
GDExtensionInt variant_hash(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantHash)fn)((GDExtensionConstVariantPtr)p_self);
}
GDExtensionInt variant_recursive_hash(pointer fn, pointer p_self, GDExtensionInt p_count) {
	return ((GDExtensionInterfaceVariantRecursiveHash)fn)((GDExtensionConstVariantPtr)p_self, p_count);
}
GDExtensionBool variant_hash_compare(pointer fn, pointer p_self, pointer p_other) {
	return ((GDExtensionInterfaceVariantHashCompare)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_other);
}
GDExtensionBool variant_booleanize(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantBooleanize)fn)((GDExtensionConstVariantPtr)p_self);
}
void variant_duplicate(pointer fn, pointer p_self, pointer r_ret, GDExtensionBool p_deep) {
	((GDExtensionInterfaceVariantDuplicate)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedVariantPtr)r_ret, p_deep);
}
void variant_stringify(pointer fn, pointer p_self, pointer r_ret) {
	((GDExtensionInterfaceVariantStringify)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedStringPtr)r_ret);
}
GDExtensionVariantType variant_get_type(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantGetType)fn)((GDExtensionConstVariantPtr)p_self);
}
GDExtensionBool variant_has_method(pointer fn, pointer p_self, pointer p_method) {
	return ((GDExtensionInterfaceVariantHasMethod)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_method);
}
GDExtensionBool variant_has_member(pointer fn, pointer p_self, pointer p_member) {
	return ((GDExtensionInterfaceVariantHasMember)fn)((GDExtensionVariantType)p_self, (GDExtensionConstStringNamePtr)p_member);
}
GDExtensionBool variant_has_key(pointer fn, pointer p_self, pointer p_key, pointer r_valid) {
	return ((GDExtensionInterfaceVariantHasKey)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionBool*)r_valid);
}
void variant_get_type_name(pointer fn, GDExtensionVariantType p_type, pointer r_ret) {
	((GDExtensionInterfaceVariantGetTypeName)fn)(p_type, (GDExtensionUninitializedStringPtr)r_ret);
}
GDExtensionBool variant_can_convert(pointer fn, pointer p_self, GDExtensionVariantType p_type) {
	return ((GDExtensionInterfaceVariantCanConvert)fn)((GDExtensionVariantType)p_self, p_type);
}
GDExtensionBool variant_can_convert_strict(pointer fn, pointer p_self, GDExtensionVariantType p_type) {
	return ((GDExtensionInterfaceVariantCanConvertStrict)fn)((GDExtensionVariantType)p_self, p_type);
}

pointer classdb_construct_object(pointer fn, pointer p_classname) {
	return (pointer)((GDExtensionInterfaceClassdbConstructObject)fn)((GDExtensionConstStringNamePtr)p_classname);
}
*/
import "C"

import (
	"unsafe"

	gd "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

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
	mem_alloc := dlsymGD("mem_alloc")
	API.Memory.Allocate = func(size uintptr) unsafe.Pointer {
		return unsafe.Pointer(C.mem_alloc(C.uintptr_t(uintptr(mem_alloc)), C.size_t(size)))
	}
	mem_realloc := dlsymGD("mem_realloc")
	API.Memory.Reallocate = func(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
		return unsafe.Pointer(C.mem_realloc(C.uintptr_t(uintptr(mem_realloc)), ptr, C.size_t(size)))
	}
	mem_free := dlsymGD("mem_free")
	API.Memory.Free = func(ptr unsafe.Pointer) {
		C.mem_free(C.uintptr_t(uintptr(mem_free)), ptr)
	}
	print_error := dlsymGD("print_error")
	API.PrintError = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error(
			C.uintptr_t(uintptr(print_error)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_error_with_message := dlsymGD("print_error_with_message")
	API.PrintErrorMessage = func(code, message, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_message := C.CString(message)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error_with_message(
			C.uintptr_t(uintptr(print_error_with_message)),
			p_description,
			p_message,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_message))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_warning := dlsymGD("print_warning")
	API.PrintWarning = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_warning(
			C.uintptr_t(uintptr(print_warning)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_script_error := dlsymGD("print_script_error")
	API.PrintScriptError = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error(
			C.uintptr_t(uintptr(print_script_error)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_script_error_with_message := dlsymGD("print_script_error_with_message")
	API.PrintScriptErrorMessage = func(code, message, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_message := C.CString(message)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error_with_message(
			C.uintptr_t(uintptr(print_script_error_with_message)),
			p_description,
			p_message,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_message))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	get_native_struct_size := dlsymGD("get_native_struct_size")
	API.GetNativeStructSize = func(name gd.StringName) uint64 {
		var frame = call.New()
		defer frame.Free()
		return uint64(C.get_native_struct_size(
			C.uintptr_t(uintptr(get_native_struct_size)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))
	}
	variant_new_copy := dlsymGD("variant_new_copy")
	API.Variants.NewCopy = func(ctx gd.Context, self gd.Variant) gd.Variant {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_dest = call.Ret[[3]uintptr](frame)
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(r_dest.Uintptr()),
			C.uintptr_t(p_self.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_dest.Get())
		frame.Free()
		return ret
	}
	variant_new_nil := dlsymGD("variant_new_nil")
	API.Variants.NewNil = func(ctx gd.Context) gd.Variant {
		var frame = call.New()
		var r_dest = call.Ret[[3]uintptr](frame)
		C.variant_new_nil(
			C.uintptr_t(uintptr(variant_new_nil)),
			C.uintptr_t(r_dest.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_dest.Get())
		frame.Free()
		return ret
	}
	variant_destroy := dlsymGD("variant_destroy")
	API.Variants.Destroy = func(self gd.Variant) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		C.variant_destroy(
			C.uintptr_t(uintptr(variant_destroy)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
	}
	variant_call := dlsymGD("variant_call")
	API.Variants.Call = func(ctx gd.Context, self gd.Variant, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_method = call.Arg(frame, method.Pointer())
		for _, arg := range args {
			call.Arg(frame, arg.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_call(
			C.uintptr_t(uintptr(variant_call)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(frame.Array(2).Uintptr()),
			C.GDExtensionInt(len(args)),
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
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_call_static := dlsymGD("variant_call_static")
	API.Variants.CallStatic = func(ctx gd.Context, vtype gd.VariantType, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		var p_method = call.Arg(frame, method.Pointer())
		for _, arg := range args {
			call.Arg(frame, arg.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_call_static(
			C.uintptr_t(uintptr(variant_call_static)),
			C.GDExtensionVariantType(vtype),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(frame.Array(1).Uintptr()),
			C.GDExtensionInt(len(args)),
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
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_evaluate := dlsymGD("variant_evaluate")
	API.Variants.Evaluate = func(ctx gd.Context, operator gd.Operator, a, b gd.Variant) (ret gd.Variant, ok bool) {
		var frame = call.New()
		var p_a = call.Arg(frame, a.Pointer())
		var p_b = call.Arg(frame, b.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_evaluate(
			C.uintptr_t(uintptr(variant_evaluate)),
			C.GDExtensionVariantOperator(operator),
			C.uintptr_t(p_a.Uintptr()),
			C.uintptr_t(p_b.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, r_valid.Get()
	}
	variant_set := dlsymGD("variant_set")
	API.Variants.Set = func(self, key, val gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
		C.variant_set(
			C.uintptr_t(uintptr(variant_set)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		frame.Free()
		return r_valid.Get()
	}
	variant_set_named := dlsymGD("variant_set_named")
	API.Variants.SetNamed = func(self gd.Variant, key gd.StringName, val gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
		C.variant_set_named(
			C.uintptr_t(uintptr(variant_set_named)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		frame.Free()
		return r_valid.Get()
	}
	variant_set_keyed := dlsymGD("variant_set_keyed")
	API.Variants.SetKeyed = func(self, key, val gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
		C.variant_set_keyed(
			C.uintptr_t(uintptr(variant_set_keyed)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		frame.Free()
		return r_valid.Get()
	}
	variant_set_indexed := dlsymGD("variant_set_indexed")
	API.Variants.SetIndexed = func(self gd.Variant, index gd.Int, val gd.Variant) (ok, oob bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
		var r_oob = call.Ret[bool](frame)
		C.variant_set_indexed(
			C.uintptr_t(uintptr(variant_set_indexed)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
			C.uintptr_t(r_oob.Uintptr()),
		)
		frame.Free()
		return r_valid.Get(), r_oob.Get()
	}
	variant_get := dlsymGD("variant_get")
	API.Variants.Get = func(ctx gd.Context, self, key gd.Variant) (val gd.Variant, ok bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_get(
			C.uintptr_t(uintptr(variant_get)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_named := dlsymGD("variant_get_named")
	API.Variants.GetNamed = func(ctx gd.Context, self gd.Variant, key gd.StringName) (val gd.Variant, ok bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_get_named(
			C.uintptr_t(uintptr(variant_get_named)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_keyed := dlsymGD("variant_get_keyed")
	API.Variants.GetKeyed = func(ctx gd.Context, self, key gd.Variant) (val gd.Variant, ok bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_get_keyed(
			C.uintptr_t(uintptr(variant_get_keyed)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_indexed := dlsymGD("variant_get_indexed")
	API.Variants.GetIndexed = func(ctx gd.Context, self gd.Variant, index gd.Int) (gd.Variant, bool, bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		var r_oob = call.Ret[bool](frame)
		C.variant_get_indexed(
			C.uintptr_t(uintptr(variant_get_indexed)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
			C.uintptr_t(r_oob.Uintptr()),
		)
		frame.Free()
		return mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get()), r_valid.Get(), r_oob.Get()
	}
	variant_iter_init := dlsymGD("variant_iter_init")
	API.Variants.IteratorInitialize = func(ctx gd.Context, self gd.Variant) (gd.Variant, bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_valid = call.Ret[bool](frame)
		var r_iter = call.Ret[[3]uintptr](frame)
		C.variant_iter_init(
			C.uintptr_t(uintptr(variant_iter_init)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_iter.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_iter.Get())
		var ok = r_valid.Get()
		frame.Free()
		return ret, ok
	}
	variant_iter_next := dlsymGD("variant_iter_next")
	API.Variants.IteratorNext = func(self gd.Variant, iter gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_iter = call.Arg(frame, iter.Pointer())
		var r_valid = call.Ret[bool](frame)
		C.variant_iter_next(
			C.uintptr_t(uintptr(variant_iter_next)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_iter.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ok = r_valid.Get()
		frame.Free()
		return ok
	}
	variant_iter_get := dlsymGD("variant_iter_get")
	API.Variants.IteratorGet = func(ctx gd.Context, self, iter gd.Variant) (gd.Variant, bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_iter = call.Arg(frame, iter.Pointer())
		var r_valid = call.Ret[bool](frame)
		var r_ret = call.Ret[[3]uintptr](frame)
		C.variant_iter_get(
			C.uintptr_t(uintptr(variant_iter_get)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_iter.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		var ok = r_valid.Get()
		frame.Free()
		return ret, ok
	}
	variant_hash := dlsymGD("variant_hash")
	API.Variants.Hash = func(self gd.Variant) gd.Int {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_hash(
			C.uintptr_t(uintptr(variant_hash)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		return gd.Int(ret)
	}
	variant_recursive_hash := dlsymGD("variant_recursive_hash")
	API.Variants.RecursiveHash = func(self gd.Variant, count gd.Int) gd.Int {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_recursive_hash(
			C.uintptr_t(uintptr(variant_recursive_hash)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(count),
		)
		frame.Free()
		return gd.Int(ret)
	}
	variant_hash_compare := dlsymGD("variant_hash_compare")
	API.Variants.HashCompare = func(self, other gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_other = call.Arg(frame, other.Pointer())
		var ret = C.variant_hash_compare(
			C.uintptr_t(uintptr(variant_hash_compare)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_other.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_booleanize := dlsymGD("variant_booleanize")
	API.Variants.Booleanize = func(self gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_booleanize(
			C.uintptr_t(uintptr(variant_booleanize)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_duplicate := dlsymGD("variant_duplicate")
	API.Variants.Duplicate = func(self gd.Variant, deep bool) gd.Variant {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var p_deep = 0
		if deep {
			p_deep = 1
		}
		C.variant_duplicate(
			C.uintptr_t(uintptr(variant_duplicate)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.GDExtensionBool(p_deep),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](self.Context(), self.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_stringify := dlsymGD("variant_stringify")
	API.Variants.Stringify = func(self gd.Variant) gd.String {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_ret = call.Ret[uintptr](frame)
		C.variant_stringify(
			C.uintptr_t(uintptr(variant_stringify)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.String, uintptr](self.Context(), self.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_get_type := dlsymGD("variant_get_type")
	API.Variants.GetType = func(self gd.Variant) gd.VariantType {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_get_type(
			C.uintptr_t(uintptr(variant_get_type)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		return gd.VariantType(ret)
	}
	variant_has_method := dlsymGD("variant_has_method")
	API.Variants.HasMethod = func(self gd.Variant, method gd.StringName) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_method = call.Arg(frame, method.Pointer())
		var ret = C.variant_has_method(
			C.uintptr_t(uintptr(variant_has_method)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_has_member := dlsymGD("variant_has_member")
	API.Variants.HasMember = func(self gd.Variant, member gd.StringName) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_member = call.Arg(frame, member.Pointer())
		var ret = C.variant_has_member(
			C.uintptr_t(uintptr(variant_has_member)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_member.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_has_key := dlsymGD("variant_has_key")
	API.Variants.HasKey = func(self gd.Variant, key gd.Variant) (hasKey, valid bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_valid = call.Ret[bool](frame)
		var ret = C.variant_has_key(
			C.uintptr_t(uintptr(variant_has_key)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		valid = r_valid.Get()
		frame.Free()
		if ret != 0 {
			return true, valid
		}
		return false, valid
	}
	variant_get_type_name := dlsymGD("variant_get_type_name")
	API.Variants.GetTypeName = func(ctx gd.Context, vtype gd.VariantType) gd.String {
		var frame = call.New()
		var r_ret = call.Ret[uintptr](frame)
		C.variant_get_type_name(
			C.uintptr_t(uintptr(variant_get_type_name)),
			C.GDExtensionVariantType(vtype),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.String, uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret
	}
	variant_can_convert := dlsymGD("variant_can_convert")
	API.Variants.CanConvert = func(self gd.Variant, toType gd.VariantType) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_can_convert(
			C.uintptr_t(uintptr(variant_can_convert)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionVariantType(toType),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_can_convert_strict := dlsymGD("variant_can_convert_strict")
	API.Variants.CanConvertStrict = func(self gd.Variant, toType gd.VariantType) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_can_convert_strict(
			C.uintptr_t(uintptr(variant_can_convert_strict)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionVariantType(toType),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}

	classdb_construct_object := dlsymGD("classdb_construct_object")
	API.ClassDB.ConstructObject = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = call.New()
		var obj gd.Object
		obj.SetPointer(mmm.Make[gd.API, gd.Pointer, uintptr](ctx, ctx.API(), uintptr(C.classdb_construct_object(
			C.uintptr_t(uintptr(classdb_construct_object)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))))
		frame.Free()
		return obj
	}

}
