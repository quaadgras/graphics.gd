//go:build cgo

package gdunsafe

// #include "gd.h"
import "C"
import (
	"unsafe"

	"graphics.gd/variant"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type (
	String     uintptr
	StringName uintptr
	Array      uintptr
	Dictionary uintptr
	Pointer    uintptr

	PackedArray[T byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String] [2]uint64

	VariantType uint32

	Object              uintptr
	ObjectType          uintptr
	MethodForClass      uintptr
	ScriptInstance      uintptr
	ExtensionInstanceID uintptr
	ExtensionClassID    uintptr
	ExtensionBindingID  uintptr
	FunctionID          uintptr
	PropertyList        uintptr
	MethodList          uintptr
)

func toVariant(v C.Variant) Variant {
	return Variant{uint64(v.tag), uint64(v.payload[0]), uint64(v.payload[1])}
}
func toCallError(cerr C.CallError) error {
	if cerr.error == 0 {
		return nil
	}
	return Error{error: errorType(cerr.error), expected: int32(cerr.expected), argument: int32(cerr.argument)}
}

func (args Variants) Index(i int) Variant {
	if args.count > 0 && (i >= args.count || i < 0) {
		panic("index out of range")
	}
	slot := unsafe.Pointer(uintptr(args.first) + unsafe.Sizeof(Pointer(0))*uintptr(i))
	return *(*Variant)(*(*unsafe.Pointer)(slot))
}

func (array Array) Set(index int64, value Variant) {
	C.gd_array_set(C.Array(array), C.int64_t(index), C.uint64_t(value[0]), C.uint64_t(value[1]), C.uint64_t(value[2]))
}

func (array Array) Get(index int64) Variant {
	var r C.Variant
	C.gd_array_get(C.Array(array), C.int64_t(index), &r)
	return toVariant(r)
}

func VersionMajor() uint32     { return uint32(C.gd_version_major()) }
func VersionMinor() uint32     { return uint32(C.gd_version_minor()) }
func VersionPatch() uint32     { return uint32(C.gd_version_patch()) }
func VersionHex() uint32       { return uint32(C.gd_version_hex()) }
func VersionStatus() String    { return String(C.gd_version_status()) }
func VersionBuild() String     { return String(C.gd_version_build()) }
func VersionHash() String      { return String(C.gd_version_hash()) }
func VersionTimestamp() uint64 { return uint64(C.gd_version_timestamp()) }
func VersionString() String    { return String(C.gd_version_string()) }
func LibraryLocation() String  { return String(C.gd_library_location()) }

func Malloc(size int64) Pointer    { return Pointer(C.gd_memory_malloc(C.int64_t(size))) }
func Sizeof(name StringName) int64 { return int64(C.gd_memory_sizeof(C.uintptr_t(name))) }
func Resize(ptr Pointer, size int64) Pointer {
	return Pointer(C.gd_memory_resize(C.UnsafePointer(ptr), C.int64_t(size)))
}
func Clear(ptr Pointer, size int64) { C.gd_memory_clear(C.UnsafePointer(ptr), C.int64_t(size)) }

func (ptr Pointer) Byte() byte     { return *(*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint16() uint16 { return *(*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint32() uint32 { return *(*uint32)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint64() uint64 { return *(*uint64)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }

func (ptr Pointer) SetByte(v byte)     { *(*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }
func (ptr Pointer) SetUint16(v uint16) { *(*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }
func (ptr Pointer) SetUint32(v uint32) { *(*uint32)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }
func (ptr Pointer) SetUint64(v uint64) { *(*uint64)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }

func (ptr Pointer) SetBits128(val [2]uint64) {
	p := *(*unsafe.Pointer)(unsafe.Pointer(&ptr))
	*(*[2]uint64)(p) = val
}
func (ptr Pointer) SetBits256(val [4]uint64) {
	p := *(*unsafe.Pointer)(unsafe.Pointer(&ptr))
	*(*[4]uint64)(p) = val
}
func (ptr Pointer) SetBits512(val [8]uint64) {
	p := *(*unsafe.Pointer)(unsafe.Pointer(&ptr))
	*(*[8]uint64)(p) = val
}

func (ptr Pointer) Free() { C.gd_memory_free(C.UnsafePointer(ptr)) }

// String operations

func (s String) Access(idx int64) int32 {
	return int32(C.gd_string_access(C.uintptr_t(s), C.int64_t(idx)))
}
func (s String) Resize(size int64) String {
	return String(C.gd_string_resize(C.uintptr_t(s), C.int64_t(size)))
}
func (s String) UnsafePtr() Pointer {
	return Pointer(C.gd_string_unsafe(C.uintptr_t(s)))
}
func (s String) Append(other String) String {
	return String(C.gd_string_append(C.uintptr_t(s), C.uintptr_t(other)))
}
func (s String) AppendRune(ch int32) String {
	return String(C.gd_string_append_rune(C.uintptr_t(s), C.int32_t(ch)))
}

func (enc StringEncoding) String(s string) String {
	return String(C.gd_string_decode(C.uint8_t(enc),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

func (s String) Encode(enc StringEncoding, buf []byte) int64 {
	return int64(C.gd_string_encode(C.uint8_t(enc), C.uintptr_t(s),
		(*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func (enc StringEncoding) Intern(s string) StringName {
	return StringName(C.gd_string_intern(C.uint8_t(enc),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

func Log(level LogLevel, text, code, fn, file string, line int32, notify_editor bool) {
	C.gd_log(C.uint32_t(level),
		(*C.char)(unsafe.Pointer(unsafe.StringData(text))), C.uint32_t(len(text)),
		(*C.char)(unsafe.Pointer(unsafe.StringData(code))), C.uint32_t(len(code)),
		(*C.char)(unsafe.Pointer(unsafe.StringData(fn))), C.uint32_t(len(fn)),
		(*C.char)(unsafe.Pointer(unsafe.StringData(file))), C.uint32_t(len(file)),
		C.int32_t(line), C._Bool(notify_editor))
}

func (ptr PointerTo[T]) Get() T  { return *(*T)(unsafe.Pointer(ptr)) }
func (ptr PointerTo[T]) Set(v T) { *(*T)(unsafe.Pointer(ptr)) = v }

func (p PackedArray[T]) Access(idx int64) PointerTo[T] {
	return PointerTo[T](C.gd_packed_array_access(C.uint32_t(p.Type()), C.uintptr_t(p[0]), C.uintptr_t(p[1]), C.int64_t(idx)))
}

func (p PackedArray[T]) Modify(idx int64) PointerTo[T] {
	return PointerTo[T](C.gd_packed_array_modify(C.uint32_t(p.Type()), C.uintptr_t(p[0]), C.uintptr_t(p[1]), C.int64_t(idx)))
}

func (t VariantType) Name() String {
	return String(C.gd_variant_type_name(C.uint32_t(t)))
}

func (t VariantType) Make(args ...Variant) (Variant, error) {
	var value C.Variant
	var err C.CallError
	C.gd_variant_type_make(C.uint32_t(t), &value, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(value), toCallError(err)
}

func (t VariantType) StaticCall(method StringName, args ...Variant) (Variant, error) {
	var value C.Variant
	var err C.CallError
	C.gd_variant_type_call(C.uint32_t(t), C.uintptr_t(method), &value, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(value), toCallError(err)
}

func (t VariantType) Convertable(to VariantType, strict bool) bool {
	return bool(C.gd_variant_type_convertable(C.uint32_t(t), C.uint32_t(to), C.bool(strict)))
}

func BuiltinName(utility StringName, hash int64) FunctionID {
	return FunctionID(C.gd_builtin_name(C.uintptr_t(utility), C.int64_t(hash)))
}

func BuiltinCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_builtin_call(C.uintptr_t(fn), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}

func VariantTypeSetupArray(array Array, vtype VariantType, className StringName, script Variant) {
	C.gd_variant_type_setup_array(C.uintptr_t(array), C.uint32_t(vtype), C.uintptr_t(className), C.uint64_t(script[0]), C.uint64_t(script[1]), C.uint64_t(script[2]))
}

func VariantTypeSetupDictionary(dict Dictionary, keyType VariantType, keyClassName StringName, keyScript Variant, valType VariantType, valClassName StringName, valScript Variant) {
	C.gd_variant_type_setup_dictionary(C.uintptr_t(dict), C.uint32_t(keyType), C.uintptr_t(keyClassName), C.uint64_t(keyScript[0]), C.uint64_t(keyScript[1]), C.uint64_t(keyScript[2]), C.uint32_t(valType), C.uintptr_t(valClassName), C.uint64_t(valScript[0]), C.uint64_t(valScript[1]), C.uint64_t(valScript[2]))
}

func VariantTypeFetchConstant(vtype VariantType, constant StringName, result unsafe.Pointer) {
	C.gd_variant_type_fetch_constant(C.uint32_t(vtype), C.uintptr_t(constant), C.UnsafePointer(result))
}

func VariantTypeConstructor(vtype VariantType, n int64) FunctionID {
	return FunctionID(C.gd_variant_type_unsafe_constructor(C.uint32_t(vtype), C.int64_t(n)))
}

func VariantTypeEvaluator(op VariantOperator, a, b VariantType) FunctionID {
	return FunctionID(C.gd_variant_type_evaluator(C.uint32_t(op), C.uint32_t(a), C.uint32_t(b)))
}

func VariantTypeSetter(vtype VariantType, property StringName) FunctionID {
	return FunctionID(C.gd_variant_type_setter(C.uint32_t(vtype), C.uintptr_t(property)))
}

func VariantTypeGetter(vtype VariantType, property StringName) FunctionID {
	return FunctionID(C.gd_variant_type_getter(C.uint32_t(vtype), C.uintptr_t(property)))
}

func VariantTypeHasProperty(vtype VariantType, property StringName) bool {
	return bool(C.gd_variant_type_has_property(C.uint32_t(vtype), C.uintptr_t(property)))
}

func VariantTypeMethod(vtype VariantType, method StringName, hash int64) FunctionID {
	return FunctionID(C.gd_variant_type_builtin_method(C.uint32_t(vtype), C.uintptr_t(method), C.int64_t(hash)))
}

func VariantTypeUnsafeCall(self unsafe.Pointer, fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_variant_type_unsafe_call(C.UnsafePointer(self), C.uintptr_t(fn), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}

func VariantTypeUnsafeMake(constructor FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_variant_type_unsafe_make(C.uintptr_t(constructor), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}

func VariantTypeUnsafeFree(vtype VariantType, shape uint64, args unsafe.Pointer) {
	C.gd_variant_type_unsafe_free(C.uint32_t(vtype), C.uint64_t(shape), C.UnsafePointer(args))
}

type (
	Callable   [2]uint64
	CallableID uintptr
)

func MakeCallable(impl ExtensionCallable, obj ObjectID) Callable {
	var c C.Callable
	C.gd_callable_create(C.CallableID(callables.New(impl)), C.ObjectID(obj), &c)
	return Callable{uint64(c.opaque[0]), uint64(c.opaque[1])}
}

//export gd_on_callable_called
func gd_on_callable_called(c C.CallableID, ret *C.Variant, argc C.Int, args C.VariadicVariants, err *C.CallError) {
	r, e := callables.Get(CallableID(c)).Call(Variants{first: PointerTo[PointerTo[Variant]](unsafe.Pointer(args)), count: int(argc)})
	*ret = C.Variant{C.uint64_t(r[0]), [2]C.uint64_t{C.uint64_t(r[1]), C.uint64_t(r[2])}}
	*err = C.CallError{C.uint32_t(e.error), C.int32_t(e.argument), C.int32_t(e.expected)}
}

//export gd_on_callable_verify
func gd_on_callable_verify(c C.CallableID) C.bool {
	return C.bool(callables.Get(CallableID(c)).IsValid())
}

//export gd_on_callable_delete
func gd_on_callable_delete(c C.CallableID) { callables.Del(CallableID(c)) }

//export gd_on_callable_hashed
func gd_on_callable_hashed(c C.CallableID) C.uint32_t {
	return C.uint32_t(callables.Get(CallableID(c)).Hash())
}

//export gd_on_callable_sorted
func gd_on_callable_sorted(a, b C.CallableID) C.Int {
	return C.Int(callables.Get(CallableID(a)).Compare(callables.Get(CallableID(b))))
}

//export gd_on_callable_string
func gd_on_callable_string(c C.CallableID) C.String {
	return C.String(callables.Get(CallableID(c)).UnsafeString())
}

//export gd_on_callable_length
func gd_on_callable_length(c C.CallableID) C.Int {
	return C.Int(callables.Get(CallableID(c)).ArgumentCount())
}

// Extension binding callbacks (no-ops)

//export gd_on_extension_binding_created
func gd_on_extension_binding_created(p0 C.uintptr_t) C.uintptr_t { return 0 }

//export gd_on_extension_binding_removed
func gd_on_extension_binding_removed(p0, p1 C.uintptr_t) {}

//export gd_on_extension_binding_reference
func gd_on_extension_binding_reference(p0 C.uintptr_t, p1 C.bool) C.bool { return false }

// Extension class callbacks

//export gd_on_extension_class_create
func gd_on_extension_class_create(p0 C.uintptr_t, p1 C.bool) C.uintptr_t {
	return C.uintptr_t(classes.Get(ExtensionClassID(p0)).Create(bool(p1)))
}

//export gd_on_extension_class_method
func gd_on_extension_class_method(p0 C.uintptr_t, p1 C.uintptr_t, p2 C.uint32_t) C.uintptr_t {
	fn := classes.Get(ExtensionClassID(p0)).Method(StringName(p1), uint32(p2))
	if fn == nil {
		return 0
	}
	return C.uintptr_t(functions.New(fn))
}

//export gd_on_extension_class_caller
func gd_on_extension_class_caller(p0 C.uintptr_t, p1 C.uintptr_t, p2 C.uint32_t) C.uintptr_t {
	fn := classes.Get(ExtensionClassID(p0)).Method(StringName(p1), uint32(p2))
	if fn == nil {
		return 0
	}
	return C.uintptr_t(functions.New(fn))
}

// Extension instance callbacks

//export gd_on_extension_instance_set
func gd_on_extension_instance_set(p0 C.uintptr_t, p1 C.uintptr_t, p2, p3, p4 C.uint64_t) C.bool {
	return C.bool(instances.Get(ExtensionInstanceID(p0)).Set(
		StringName(p1), Variant{uint64(p2), uint64(p3), uint64(p4)}))
}

//export gd_on_extension_instance_get
func gd_on_extension_instance_get(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant) C.bool {
	v, ok := instances.Get(ExtensionInstanceID(p0)).Get(StringName(p1))
	if ok {
		*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
	}
	return C.bool(ok)
}

//export gd_on_extension_instance_property_list
func gd_on_extension_instance_property_list(p0 C.uintptr_t) C.uintptr_t {
	return C.uintptr_t(instances.Get(ExtensionInstanceID(p0)).PropertyList())
}

//export gd_on_extension_instance_property_has_default
func gd_on_extension_instance_property_has_default(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	return C.bool(instances.Get(ExtensionInstanceID(p0)).HasDefault(StringName(p1)))
}

//export gd_on_extension_instance_property_get_default
func gd_on_extension_instance_property_get_default(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant) C.bool {
	v, ok := instances.Get(ExtensionInstanceID(p0)).GetDefault(StringName(p1))
	if ok {
		*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
	}
	return C.bool(ok)
}

//export gd_on_extension_instance_property_validation
func gd_on_extension_instance_property_validation(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	return C.bool(instances.Get(ExtensionInstanceID(p0)).ValidateProperty(StringName(p1)))
}

//export gd_on_extension_instance_notification
func gd_on_extension_instance_notification(p0 C.uintptr_t, p1 C.int32_t, p2 C.bool) {
	instances.Get(ExtensionInstanceID(p0)).Notification(int32(p1), bool(p2))
}

//export gd_on_extension_instance_stringify
func gd_on_extension_instance_stringify(p0 C.uintptr_t) C.uintptr_t {
	return C.uintptr_t(instances.Get(ExtensionInstanceID(p0)).UnsafeString())
}

//export gd_on_extension_instance_reference
func gd_on_extension_instance_reference(p0 C.uintptr_t, p1 C.bool) C.bool {
	return C.bool(instances.Get(ExtensionInstanceID(p0)).Reference(bool(p1)))
}

//export gd_on_extension_instance_rid
func gd_on_extension_instance_rid(p0 C.uintptr_t) C.uint64_t {
	return C.uint64_t(instances.Get(ExtensionInstanceID(p0)).RID())
}

//export gd_on_extension_instance_checked_call
func gd_on_extension_instance_checked_call(p0, p1 C.uintptr_t, p2, p3 C.UnsafePointer) {
	var inst ExtensionInstance
	if ExtensionInstanceID(p0) != 0 {
		inst = instances.Get(ExtensionInstanceID(p0))
	}
	functions.Get(FunctionID(p1)).PointerCall(inst, Pointer(uintptr(p3)), Pointer(uintptr(p2)))
}

//export gd_on_extension_instance_called
func gd_on_extension_instance_called(p0, p1 C.uintptr_t, p2, p3 C.UnsafePointer) {
	inst := instances.Get(ExtensionInstanceID(p0))
	functions.Get(FunctionID(p1)).PointerCall(inst, Pointer(uintptr(p3)), Pointer(uintptr(p2)))
}

//export gd_on_extension_instance_variant_call
func gd_on_extension_instance_variant_call(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant, p3 C.VariadicVariants) {
	var inst ExtensionInstance
	if ExtensionInstanceID(p0) != 0 {
		inst = instances.Get(ExtensionInstanceID(p0))
	}
	v := functions.Get(FunctionID(p1)).CheckedCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](unsafe.Pointer(p3)),
		count: -1,
	})
	*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
}

//export gd_on_extension_instance_dynamic_call
func gd_on_extension_instance_dynamic_call(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant, p3 C.int64_t, p4 C.VariadicVariants, p5 *C.CallError) {
	var inst ExtensionInstance
	if ExtensionInstanceID(p0) != 0 {
		inst = instances.Get(ExtensionInstanceID(p0))
	}
	v, err := functions.Get(FunctionID(p1)).DynamicCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](unsafe.Pointer(p4)),
		count: int(p3),
	})
	*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
	*p5 = C.CallError{C.uint32_t(err.error), C.int32_t(err.argument), C.int32_t(err.expected)}
}

//export gd_on_extension_instance_free
func gd_on_extension_instance_free(p0 C.uintptr_t) {
	inst := instances.Get(ExtensionInstanceID(p0))
	if f, ok := inst.(interface{ Free() }); ok {
		f.Free()
	}
	instances.Del(ExtensionInstanceID(p0))
}

// Extension script callbacks

//export gd_on_extension_script_categorization
func gd_on_extension_script_categorization(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.PropertyCategory() != 0)
}

//export gd_on_extension_script_get_property_type
func gd_on_extension_script_get_property_type(p0 C.uintptr_t, name C.uintptr_t, p1 *C.CallError) C.uint32_t {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		*p1 = C.CallError{C.uint32_t(errorInvalidMethod), 0, 0}
		return 0
	}
	return C.uint32_t(script.PropertyType(StringName(name)))
}

//export gd_on_extension_script_get_owner
func gd_on_extension_script_get_owner(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.Owner())
}

//export gd_on_extension_script_get_property_state
func gd_on_extension_script_get_property_state(p0 C.uintptr_t, p1 C.uintptr_t, p2 C.uintptr_t) {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return
	}
	script.ExportedProperties(func(name StringName, value Variant) bool {
		ScriptPropertyStateAdd(FunctionID(p1), Pointer(p2), name, value)
		return true
	})
}

//export gd_on_extension_script_get_methods
func gd_on_extension_script_get_methods(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.MethodList())
}

//export gd_on_extension_script_has_method
func gd_on_extension_script_has_method(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.HasMethod(StringName(p1)))
}

//export gd_on_extension_script_get_method_argument_count
func gd_on_extension_script_get_method_argument_count(p0 C.uintptr_t, p1 C.uintptr_t) C.int64_t {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.int64_t(script.MethodArgumentCount(StringName(p1)))
}

//export gd_on_extension_script_get
func gd_on_extension_script_get(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.Script())
}

//export gd_on_extension_script_is_placeholder
func gd_on_extension_script_is_placeholder(p0 C.uintptr_t) C.bool {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.IsPlaceholder())
}

//export gd_on_extension_script_get_language
func gd_on_extension_script_get_language(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.ScriptLanguage())
}

// Non-extension callbacks

//export gd_on_engine_init
func gd_on_engine_init(p0 C.uint32_t) { onEngineInit(InitializationLevel(p0)) }

//export gd_on_engine_exit
func gd_on_engine_exit(p0 C.uint32_t) { onEngineExit(InitializationLevel(p0)) }

//export gd_on_first_frame
func gd_on_first_frame() { onFirstFrame() }

//export gd_on_every_frame
func gd_on_every_frame() { onEveryFrame() }

//export gd_on_final_frame
func gd_on_final_frame() { onFinalFrame() }

//export gd_on_worker_thread_pool_task
func gd_on_worker_thread_pool_task(p0 C.uintptr_t) { onWorkerThreadPoolTask(TaskID(p0)) }

//export gd_on_worker_thread_pool_group_task
func gd_on_worker_thread_pool_group_task(p0 C.uintptr_t, p1 C.uint32_t) {
	onWorkerThreadPoolGroupTask(TaskID(p0), int32(p1))
}

//export gd_on_editor_class_in_use_detection
func gd_on_editor_class_in_use_detection(p0, p1 C.uintptr_t, p2 *C.PackedStringArray) {
	if onEditorClassDetection != nil {
		result := onEditorClassDetection(PackedArray[String]{uint64(p0), uint64(p1)})
		p2.array = C.uint64_t(result[0])
		p2.length = C.uint64_t(result[1])
	}
}

// Object construction and identity

func MakeObject(name StringName) Object {
	return Object(C.gd_object_make(C.StringName(name)))
}
func (obj Object) Name() StringName {
	return StringName(C.gd_object_name(C.Object(obj)))
}
func ObjectTypeTag(name StringName) ObjectType {
	return ObjectType(C.gd_object_type(C.StringName(name)))
}
func (obj Object) Cast(to ObjectType) Object {
	return Object(C.gd_object_cast(C.Object(obj), C.ObjectType(to)))
}
func (id ObjectID) Lookup() Object {
	return Object(C.gd_object_lookup(C.ObjectID(id)))
}
func ObjectGlobal(name StringName) Object {
	return Object(C.gd_object_global(C.StringName(name)))
}
func (obj Object) ID() ObjectID {
	return ObjectID(C.gd_object_id(C.Object(obj)))
}
func ObjectIDInsideVariant(v Variant) ObjectID {
	return ObjectID(C.gd_object_id_inside_variant(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}
func (obj Object) Free() {
	C.gd_object_unsafe_free(C.Object(obj))
}

// Object method calls

func MethodLookup(class, method StringName, hash int64) MethodForClass {
	return MethodForClass(C.gd_object_method_lookup(C.StringName(class), C.StringName(method), C.int64_t(hash)))
}
func (obj Object) Call(method MethodForClass, args ...Variant) (Variant, error) {
	var ret C.Variant
	var err C.CallError
	C.gd_object_call(C.Object(obj), C.MethodForClass(method), &ret, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(ret), toCallError(err)
}
func (obj Object) ShapedCall(fn MethodForClass, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_object_shaped_call(C.Object(obj), C.MethodForClass(fn), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}

// Extension instance management

func (obj Object) ExtensionSetup(name StringName, inst ExtensionInstance) {
	C.gd_object_extension_setup(C.Object(obj), C.StringName(name), C.ExtensionInstanceID(instances.New(inst)))
}
func (obj Object) ExtensionFetch() ExtensionInstance {
	return instances.Get(ExtensionInstanceID(C.gd_object_extension_fetch(C.Object(obj))))
}
func (obj Object) ExtensionClose() {
	C.gd_object_extension_close(C.Object(obj))
}

// Script instance management

func ScriptMake(fn ExtensionScript) ScriptInstance {
	return ScriptInstance(C.gd_object_script_make(C.ExtensionInstanceID(instances.New(fn))))
}
func (obj Object) ScriptCall(name StringName, args ...Variant) (Variant, error) {
	var ret C.Variant
	var err C.CallError
	C.gd_object_script_call(C.Object(obj), C.StringName(name), &ret, C.int64_t(len(args)), (*C.Variant)(C.UnsafePointer(unsafe.SliceData(args))), &err)
	return toVariant(ret), toCallError(err)
}
func (obj Object) ScriptSetup(script ScriptInstance) {
	C.gd_object_script_setup(C.Object(obj), C.ScriptInstance(script))
}
func (obj Object) ScriptFetch(language Object) ScriptInstance {
	return ScriptInstance(C.gd_object_script_fetch(C.Object(obj), C.Object(language)))
}
func (obj Object) ScriptDefinesMethod(method StringName) bool {
	return bool(C.gd_object_script_defines_method(C.Object(obj), C.StringName(method)))
}
func ScriptPropertyStateAdd(fn FunctionID, arg Pointer, name StringName, state Variant) {
	C.gd_object_script_property_state_add(C.FunctionID(fn), C.uintptr_t(arg), C.StringName(name), C.uint64_t(state[0]), C.uint64_t(state[1]), C.uint64_t(state[2]))
}
func ScriptPlaceholderCreate(language, script, owner Object) ScriptInstance {
	return ScriptInstance(C.gd_object_script_placeholder_create(C.Object(language), C.Object(script), C.Object(owner)))
}
func ScriptPlaceholderUpdate(script ScriptInstance, array Array, dict Dictionary) {
	C.gd_object_script_placeholder_update(C.ScriptInstance(script), C.Array(array), C.Dictionary(dict))
}

// Variant operations

type VariantOperator = uint32

func ZeroVariant() Variant {
	var zero C.Variant
	C.gd_variant_zero(&zero)
	return toVariant(zero)
}

func (v Variant) Copy() Variant {
	var result C.Variant
	C.gd_variant_copy(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), &result)
	return toVariant(result)
}
func (v Variant) VariantCall(method StringName, args ...Variant) (Variant, error) {
	var result C.Variant
	var err C.CallError
	C.gd_variant_call(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(method), &result, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(result), toCallError(err)
}
func VariantEval(op VariantOperator, a, b Variant) (Variant, bool) {
	var result C.Variant
	ok := bool(C.gd_variant_eval(C.uint32_t(op), C.uint64_t(a[0]), C.uint64_t(a[1]), C.uint64_t(a[2]), C.uint64_t(b[0]), C.uint64_t(b[1]), C.uint64_t(b[2]), &result))
	return toVariant(result), ok
}
func (v Variant) Hash() int64 {
	return int64(C.gd_variant_hash(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}
func (v Variant) Bool() bool {
	return bool(C.gd_variant_bool(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}
func (v Variant) Text() String {
	return String(C.gd_variant_text(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}
func (v Variant) Type() variant.Type {
	return variant.Type(C.gd_variant_type(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}

// Deep variant operations

func (v Variant) DeepCopy() Variant {
	var result C.Variant
	C.gd_variant_deep_copy(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), &result)
	return toVariant(result)
}
func (v Variant) DeepHash(recursion int64) int64 {
	return int64(C.gd_variant_deep_hash(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.int64_t(recursion)))
}

// Variant get/set/has

func (v Variant) GetIndex(key Variant) (Variant, bool) {
	var result C.Variant
	ok := bool(C.gd_variant_get_index(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), &result))
	return toVariant(result), ok
}
func (v Variant) GetArray(idx int64) (Variant, bool, error) {
	var result C.Variant
	var err C.CallError
	ok := bool(C.gd_variant_get_array(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.int64_t(idx), &result, &err))
	return toVariant(result), ok, toCallError(err)
}
func (v Variant) GetField(field StringName) (Variant, bool) {
	var result C.Variant
	ok := bool(C.gd_variant_get_field(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(field), &result))
	return toVariant(result), ok
}
func (v Variant) SetIndex(key, val Variant) bool {
	return bool(C.gd_variant_set_index(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), C.uint64_t(val[0]), C.uint64_t(val[1]), C.uint64_t(val[2])))
}
func (v Variant) SetArray(idx int64, val Variant, err unsafe.Pointer) bool {
	return bool(C.gd_variant_set_array(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.int64_t(idx), C.uint64_t(val[0]), C.uint64_t(val[1]), C.uint64_t(val[2]), C.UnsafePointer(err)))
}
func (v Variant) SetField(field StringName, value Variant) bool {
	return bool(C.gd_variant_set_field(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(field), C.uint64_t(value[0]), C.uint64_t(value[1]), C.uint64_t(value[2])))
}
func (v Variant) HasIndex(index Variant) bool {
	return bool(C.gd_variant_has_index(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(index[0]), C.uint64_t(index[1]), C.uint64_t(index[2])))
}
func (v Variant) HasMethod(method StringName) bool {
	return bool(C.gd_variant_has_method(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(method)))
}

// Unsafe variant operations

func VariantUnsafeCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_call(C.FunctionID(fn), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}
func VariantUnsafeEval(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_eval(C.FunctionID(fn), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}
func (v Variant) UnsafeFree() {
	C.gd_variant_unsafe_free(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]))
}
func VariantUnsafeMakeNative(vtype VariantType, v Variant, shape uint64, result unsafe.Pointer) {
	C.gd_variant_unsafe_make_native(C.uint32_t(vtype), C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(shape), C.UnsafePointer(result))
}
func VariantUnsafeFromNative(vtype VariantType, shape uint64, args unsafe.Pointer) Variant {
	var result C.Variant
	C.gd_variant_unsafe_from_native(C.uint32_t(vtype), &result, C.uint64_t(shape), C.UnsafePointer(args))
	return toVariant(result)
}
func VariantUnsafeInternalPointer(vtype VariantType, v Variant) Pointer {
	return Pointer(C.gd_variant_unsafe_internal_pointer(C.uint32_t(vtype), C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}
func VariantUnsafeSetField(setter FunctionID, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_set_field(C.FunctionID(setter), C.uint64_t(shape), C.UnsafePointer(args))
}
func VariantUnsafeSetArray(vtype VariantType, idx int64, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_set_array(C.uint32_t(vtype), C.int64_t(idx), C.uint64_t(shape), C.UnsafePointer(args))
}
func VariantUnsafeSetIndex(vtype VariantType, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_set_index(C.uint32_t(vtype), C.uint64_t(shape), C.UnsafePointer(args))
}
func VariantUnsafeGetField(getter FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_get_field(C.FunctionID(getter), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}
func VariantUnsafeGetArray(vtype VariantType, idx int64, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_get_array(C.uint32_t(vtype), C.int64_t(idx), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}
func VariantUnsafeGetIndex(vtype VariantType, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_variant_unsafe_get_index(C.uint32_t(vtype), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}

// Dictionary operations

func (d Dictionary) Access(key Variant) Variant {
	var result C.Variant
	C.gd_packed_dictionary_access(C.uintptr_t(d), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), &result)
	return toVariant(result)
}

func (d Dictionary) Modify(key, val Variant) {
	C.gd_packed_dictionary_modify(C.uintptr_t(d), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), C.uint64_t(val[0]), C.uint64_t(val[1]), C.uint64_t(val[2]))
}

// RefCounted operations

func RefGet(ref Pointer) Object {
	return Object(C.gd_ref_get_object(C.uintptr_t(ref)))
}

func RefSet(ref Pointer, obj Object) {
	C.gd_ref_set_object(C.uintptr_t(ref), C.uintptr_t(obj))
}

// Editor operations

func EditorAddDocumentation(xml string) {
	C.gd_editor_add_documentation((*C.char)(unsafe.Pointer(unsafe.StringData(xml))), C.uint32_t(len(xml)))
}

func EditorAddPlugin(name StringName) {
	C.gd_editor_add_plugin(C.uintptr_t(name))
}

func EditorEndPlugin(name StringName) {
	C.gd_editor_end_plugin(C.uintptr_t(name))
}

// PropertyList operations

func MakePropertyList(n int64) PropertyList {
	return PropertyList(C.gd_property_list_make(C.int64_t(n)))
}

func (p PropertyList) Push(vtype VariantType, name StringName, className StringName, hint uint32, hintString String, usage uint32, meta uint32) {
	C.gd_property_list_push(C.uintptr_t(p), C.uint32_t(vtype), C.uintptr_t(name), C.uintptr_t(className), C.uint32_t(hint), C.uintptr_t(hintString), C.uint32_t(usage), C.uint32_t(meta))
}

func (p PropertyList) Free() {
	C.gd_property_list_free(C.uintptr_t(p))
}

func (p PropertyList) InfoType() VariantType {
	return VariantType(C.gd_property_info_type(C.uintptr_t(p)))
}

func (p PropertyList) InfoName() StringName {
	return StringName(C.gd_property_info_name(C.uintptr_t(p)))
}

func (p PropertyList) InfoClassName() StringName {
	return StringName(C.gd_property_info_class_name(C.uintptr_t(p)))
}

func (p PropertyList) InfoHint() uint32 {
	return uint32(C.gd_property_info_hint(C.uintptr_t(p)))
}

func (p PropertyList) InfoHintString() String {
	return String(C.gd_property_info_hint_string(C.uintptr_t(p)))
}

func (p PropertyList) InfoUsage() uint32 {
	return uint32(C.gd_property_info_usage(C.uintptr_t(p)))
}

// MethodList operations

func MakeMethodList(n int64) MethodList {
	return MethodList(C.gd_method_list_make(C.int64_t(n)))
}

func (m MethodList) Push(name StringName, call ExtensionFunction, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count int64, defaults unsafe.Pointer) {
	C.gd_method_list_push(C.uintptr_t(m), C.uintptr_t(name), C.uintptr_t(functions.New(call)), C.uint32_t(flags), C.uintptr_t(returnInfo), C.uintptr_t(argsInfo), C.int64_t(count), C.UnsafePointer(defaults))
}

func (m MethodList) Free() {
	C.gd_method_list_free(C.uintptr_t(m))
}

// ClassDB sub-API operations

func FileAccessWrite(file Object, buf []byte) {
	C.gd_classdb_FileAccess_write(C.uintptr_t(file), (*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf)))
}

func FileAccessRead(file Object, buf []byte) int {
	return int(C.gd_classdb_FileAccess_read(C.uintptr_t(file), (*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func ImageUnsafe(img Object) Pointer {
	return Pointer(C.gd_classdb_Image_unsafe(C.uintptr_t(img)))
}

func ImageAccess(img Object, offset int64) byte {
	return byte(C.gd_classdb_Image_access(C.uintptr_t(img), C.int64_t(offset)))
}

func XMLParserLoad(parser Object, buf []byte) int {
	return int(C.gd_classdb_XMLParser_load(C.uintptr_t(parser), (*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func WorkerThreadPoolAddTask(pool Object, task Pointer, priority bool, description String) {
	C.gd_classdb_WorkerThreadPool_add_task(C.uintptr_t(pool), C.uintptr_t(task), C._Bool(priority), C.uintptr_t(description))
}

func WorkerThreadPoolAddGroupTask(pool Object, task Pointer, elements, arg int32, priority bool, description String) {
	C.gd_classdb_WorkerThreadPool_add_group_task(C.uintptr_t(pool), C.uintptr_t(task), C.int32_t(elements), C.int32_t(arg), C._Bool(priority), C.uintptr_t(description))
}

// ClassDB registration

func RegisterClass(class, parent StringName, id ExtensionClass, virtual, abstract, exposed, runtime bool, icon String) {
	C.gd_classdb_register(C.uintptr_t(class), C.uintptr_t(parent), C.uintptr_t(classes.New(id)), C.bool(virtual), C.bool(abstract), C.bool(exposed), C.bool(runtime), C.uintptr_t(icon))
}

func RegisterMethods(class StringName, methods MethodList) {
	C.gd_classdb_register_methods(C.uintptr_t(class), C.uintptr_t(methods))
}

func RegisterConstant(class, enum, name StringName, value int64, bitfield bool) {
	C.gd_classdb_register_constant(C.uintptr_t(class), C.uintptr_t(enum), C.uintptr_t(name), C.int64_t(value), C.bool(bitfield))
}

func RegisterProperty(class StringName, property PropertyList, setter, getter StringName) {
	C.gd_classdb_register_property(C.uintptr_t(class), C.uintptr_t(property), C.uintptr_t(setter), C.uintptr_t(getter))
}

func RegisterPropertyIndexed(class StringName, property PropertyList, setter, getter StringName, index int) {
	C.gd_classdb_register_property_indexed(C.uintptr_t(class), C.uintptr_t(property), C.uintptr_t(setter), C.uintptr_t(getter), C.int64_t(index))
}

func RegisterPropertyGroup(class StringName, group, prefix String) {
	C.gd_classdb_register_property_group(C.uintptr_t(class), C.uintptr_t(group), C.uintptr_t(prefix))
}

func RegisterPropertySubgroup(class StringName, subgroup, prefix String) {
	C.gd_classdb_register_property_sub_group(C.uintptr_t(class), C.uintptr_t(subgroup), C.uintptr_t(prefix))
}

func RegisterSignal(class, signal StringName, args PropertyList) {
	C.gd_classdb_register_signal(C.uintptr_t(class), C.uintptr_t(signal), C.uintptr_t(args))
}

func RegisterRemoval(class StringName) {
	C.gd_classdb_register_removal(C.uintptr_t(class))
}

// Iterator operations

func (v Variant) IteratorMake(result unsafe.Pointer, err unsafe.Pointer) {
	C.gd_iterator_make(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.UnsafePointer(result), C.UnsafePointer(err))
}
func (v Variant) IteratorNext(iter unsafe.Pointer, err unsafe.Pointer) bool {
	return bool(C.gd_iterator_next(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.UnsafePointer(iter), C.UnsafePointer(err)))
}
func (v Variant) IteratorLoad(iter Variant, result unsafe.Pointer, err unsafe.Pointer) {
	C.gd_iterator_load(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(iter[0]), C.uint64_t(iter[1]), C.uint64_t(iter[2]), C.UnsafePointer(result), C.UnsafePointer(err))
}
