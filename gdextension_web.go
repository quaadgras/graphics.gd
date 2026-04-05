//go:build wasm

package gdunsafe

import (
	"sync"
	"unsafe"
)

type (
	String     uint32
	StringName uint32
	Array      uint32
	Dictionary uint32
	Pointer    uint32

	VariantType uint32

	Object              uint32
	ObjectType          uint32
	MethodForClass      uint32
	ScriptInstance      uint32
	ExtensionInstanceID uint32
	ExtensionClassID    uint32
	ExtensionBindingID  uint32
	FunctionID          uint32
	PropertyList        uint32
	MethodList          uint32
)

//go:wasmimport gd array_set
func gd_array_set(array Array, index Int, v1, v2, v3 uint64)

func (array Array) Set(index Int, value Variant) {
	gd_array_set(array, index, value[0], value[1], value[2])
}

//go:wasmimport gd array_get
func gd_array_get(array Array, index Int, result uint32)

func (array Array) Get(index Int) Variant {
	var value Variant
	result := makeResult(SizeVariant)
	gd_array_get(array, index, uint32(result))
	loadResult(SizeVariant, &value, result)
	return value
}

//go:wasmimport gd version_major
func VersionMajor() uint32

//go:wasmimport gd version_minor
func VersionMinor() uint32

//go:wasmimport gd version_patch
func VersionPatch() uint32

//go:wasmimport gd version_hex
func VersionHex() uint32

//go:wasmimport gd version_status
func VersionStatus() String

//go:wasmimport gd version_build
func VersionBuild() String

//go:wasmimport gd version_hash
func VersionHash() String

//go:wasmimport gd gd_version_timestamp
func VersionTimestamp() uint64

//go:wasmimport gd version_string
func VersionString() String

//go:wasmimport gd library_location
func LibraryLocation() String

//go:wasmimport gd memory_malloc
func Malloc(Int) Pointer

//go:wasmimport gd memory_sizeof
func Sizeof(StringName) Int

//go:wasmimport gd memory_resize
func Resize(Pointer, Int) Pointer

//go:wasmimport gd memory_clear
func Clear(Pointer, Int)

//go:wasmimport gd memory_load_byte
func gd_memory_load_byte(Pointer) uint32

//go:wasmimport gd memory_load_u16
func gd_memory_load_u16(Pointer) uint32

func (ptr Pointer) Byte() byte     { return byte(gd_memory_load_byte(ptr)) }
func (ptr Pointer) Uint16() uint16 { return uint16(gd_memory_load_u16(ptr)) }

//go:wasmimport gd memory_load_u32
func (ptr Pointer) Uint32() uint32

//go:wasmimport gd memory_load_u64
func (ptr Pointer) Uint64() uint64

//go:wasmimport gd memory_edit_byte
func gd_memory_edit_byte(Pointer, uint32)

//go:wasmimport gd memory_edit_u16
func gd_memory_edit_u16(Pointer, uint32)

func (ptr Pointer) SetByte(val byte)     { gd_memory_edit_byte(ptr, uint32(val)) }
func (ptr Pointer) SetUint16(val uint16) { gd_memory_edit_u16(ptr, uint32(val)) }

//go:wasmimport gd memory_edit_u32
func (ptr Pointer) SetUint32(val uint32)

//go:wasmimport gd memory_edit_u64
func (ptr Pointer) SetUint64(val uint64)

//go:wasmimport gd memory_edit_128
func gd_memory_edit_128(Pointer, uint64, uint64)

//go:wasmimport gd memory_edit_256
func gd_memory_edit_256(Pointer, uint64, uint64, uint64, uint64)

//go:wasmimport gd memory_edit_512
func gd_memory_edit_512(Pointer, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64)

func (ptr Pointer) SetBits128(val [2]uint64) {
	gd_memory_edit_128(ptr, val[0], val[1])
}
func (ptr Pointer) SetBits256(val [4]uint64) {
	gd_memory_edit_256(ptr, val[0], val[1], val[2], val[3])
}
func (ptr Pointer) SetBits512(val [8]uint64) {
	gd_memory_edit_512(ptr, val[0], val[1], val[2], val[3], val[4], val[5], val[6], val[7])
}

//go:wasmimport gd memory_free
func (ptr Pointer) Free()

//go:wasmimport gd log_error
func LogError(text, code, fn, file string, line int32, notify_editor bool)

//go:wasmimport gd log_warning
func LogWarning(text, code, fn, file string, line int32, notify_editor bool)

//go:wasmimport gd variant_type_name
func (t VariantType) Name() String

//go:wasmimport gd variant_type_make
func gd_variant_type_make(t VariantType, result Pointer, arg_count Int, args, err Pointer)

func (t VariantType) Make(args ...Variant) (value Variant, err CallError) {
	var param = copyVariants(unsafe.SliceData(args), len(args))
	result := makeResult(SizeVariant)
	result_err := makeResult(SizeCallError)
	gd_variant_type_make(t, Pointer(result), Int(len(args)), Pointer(param), Pointer(result_err))
	loadResult(SizeVariant, &value, result)
	loadResult(SizeCallError, &value, result_err)
	return value, err
}

//go:wasmimport gd variant_type_call
func gd_variant_type_call_wasm(t VariantType, method StringName, result Pointer, argc Int, args Pointer, err Pointer)

func (t VariantType) StaticCall(method StringName, args ...Variant) (value Variant, callErr CallError) {
	param := copyVariants(unsafe.SliceData(args), len(args))
	result := makeResult(SizeVariant)
	result_err := makeResult(SizeCallError)
	gd_variant_type_call_wasm(t, method, Pointer(result), Int(len(args)), Pointer(param), Pointer(result_err))
	loadResult(SizeVariant, &value, result)
	loadResult(SizeCallError, &callErr, result_err)
	return value, callErr
}

//go:wasmimport gd variant_type_convertable
func gd_variant_type_convertable_wasm(t VariantType, to VariantType, strict uint32) uint32

func (t VariantType) Convertable(to VariantType, strict bool) bool {
	var s uint32
	if strict {
		s = 1
	}
	return gd_variant_type_convertable_wasm(t, to, s) != 0
}

//go:wasmimport gd builtin_name
func gd_builtin_name_wasm(utility StringName, hash int64) FunctionID

func BuiltinName(utility StringName, hash int64) FunctionID {
	return gd_builtin_name_wasm(utility, hash)
}

//go:wasmimport gd builtin_call
func gd_builtin_call_wasm(fn FunctionID, result Pointer, shape uint64, args Pointer)

func BuiltinCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_builtin_call_wasm(fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_type_setup_array
func gd_variant_type_setup_array_wasm(array Array, vtype VariantType, className StringName, v1, v2, v3 uint64)

func VariantTypeSetupArray(array Array, vtype VariantType, className StringName, script Variant) {
	gd_variant_type_setup_array_wasm(array, vtype, className, script[0], script[1], script[2])
}

//go:wasmimport gd variant_type_setup_dictionary
func gd_variant_type_setup_dictionary_wasm(dict Dictionary, keyType VariantType, keyClassName StringName, ks1, ks2, ks3 uint64, valType VariantType, valClassName StringName, vs1, vs2, vs3 uint64)

func VariantTypeSetupDictionary(dict Dictionary, keyType VariantType, keyClassName StringName, keyScript Variant, valType VariantType, valClassName StringName, valScript Variant) {
	gd_variant_type_setup_dictionary_wasm(dict, keyType, keyClassName, keyScript[0], keyScript[1], keyScript[2], valType, valClassName, valScript[0], valScript[1], valScript[2])
}

//go:wasmimport gd variant_type_fetch_constant
func gd_variant_type_fetch_constant_wasm(vtype VariantType, constant StringName, result Pointer)

func VariantTypeFetchConstant(vtype VariantType, constant StringName, result unsafe.Pointer) {
	mem := makeResult(SizeVariant)
	gd_variant_type_fetch_constant_wasm(vtype, constant, Pointer(mem))
	loadResult(SizeVariant, result, mem)
}

//go:wasmimport gd variant_type_unsafe_constructor
func gd_variant_type_unsafe_constructor_wasm(vtype VariantType, n Int) FunctionID

func VariantTypeConstructor(vtype VariantType, n Int) FunctionID {
	return gd_variant_type_unsafe_constructor_wasm(vtype, n)
}

//go:wasmimport gd variant_type_evaluator
func gd_variant_type_evaluator_wasm(op VariantOperator, a, b VariantType) FunctionID

func VariantTypeEvaluator(op VariantOperator, a, b VariantType) FunctionID {
	return gd_variant_type_evaluator_wasm(op, a, b)
}

//go:wasmimport gd variant_type_setter
func gd_variant_type_setter_wasm(vtype VariantType, property StringName) FunctionID

func VariantTypeSetter(vtype VariantType, property StringName) FunctionID {
	return gd_variant_type_setter_wasm(vtype, property)
}

//go:wasmimport gd variant_type_getter
func gd_variant_type_getter_wasm(vtype VariantType, property StringName) FunctionID

func VariantTypeGetter(vtype VariantType, property StringName) FunctionID {
	return gd_variant_type_getter_wasm(vtype, property)
}

//go:wasmimport gd variant_type_has_property
func gd_variant_type_has_property_wasm(vtype VariantType, property StringName) uint32

func VariantTypeHasProperty(vtype VariantType, property StringName) bool {
	return gd_variant_type_has_property_wasm(vtype, property) != 0
}

//go:wasmimport gd variant_type_builtin_method
func gd_variant_type_builtin_method_wasm(vtype VariantType, method StringName, hash int64) FunctionID

func VariantTypeMethod(vtype VariantType, method StringName, hash int64) FunctionID {
	return gd_variant_type_builtin_method_wasm(vtype, method, hash)
}

//go:wasmimport gd variant_type_unsafe_call
func gd_variant_type_unsafe_call_wasm(self Pointer, fn FunctionID, result Pointer, shape uint64, args Pointer)

func VariantTypeUnsafeCall(self unsafe.Pointer, fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_self := copyArguments(Shape(shape)>>4, self)
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_type_unsafe_call_wasm(Pointer(mem_self), fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape)>>4, self, mem_self)
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_type_unsafe_make
func gd_variant_type_unsafe_make_wasm(constructor FunctionID, result Pointer, shape uint64, args Pointer)

func VariantTypeUnsafeMake(constructor FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_type_unsafe_make_wasm(constructor, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_type_unsafe_free
func gd_variant_type_unsafe_free_wasm(vtype VariantType, shape uint64, args Pointer)

func VariantTypeUnsafeFree(vtype VariantType, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_type_unsafe_free_wasm(vtype, shape, Pointer(mem_args))
}

func (args VariadicVariants) Index(i int) Variant {
	if i >= args.Count || i < 0 {
		panic("index out of range")
	}
	// args.First points to an engine-side array of pointers-to-Variant.
	// Read the i-th pointer, then read the Variant it points to.
	ptr := Pointer(Pointer(args.First) + Pointer(i)*Pointer(4)).Uint32()
	if ptr == 0 {
		return Variant{}
	}
	return readVariant(Pointer(ptr))
}

type (
	Callable   [2]uint64
	CallableID uint32
)

//go:wasmimport gd callable_create
func gd_callable_create(id CallableID, object ObjectID, result Pointer)

func MakeCallable(impl CallableImplementation, obj ObjectID) Callable {
	result := makeResult(SizeCallable)
	gd_callable_create(callables.New(impl), obj, Pointer(result))
	var c Callable
	loadResult(SizeCallable, unsafe.Pointer(&c), result)
	return c
}

//go:wasmexport gd_on_callable_called
func gd_on_callable_called(c CallableID, ret Pointer, argc Int, args Pointer, err Pointer) {
	r, e := callables.Get(c).Call(VariadicVariants{First: PointerTo[PointerTo[Variant]](args), Count: int(argc)})
	ret.SetBits128([2]uint64{r[0], r[1]})
	(ret + 16).SetUint64(r[2])
	(err + 0).SetUint32(uint32(e.Type))
	(err + 4).SetInt32(e.Argument)
	(err + 8).SetInt32(e.Expected)
}

//go:wasmexport gd_on_callable_verify
func gd_on_callable_verify(c CallableID) bool {
	return callables.Get(c).IsValid()
}

//go:wasmexport gd_on_callable_delete
func gd_on_callable_delete(c CallableID) { callables.Del(c) }

//go:wasmexport gd_on_callable_hashed
func gd_on_callable_hashed(c CallableID) uint32 {
	return callables.Get(c).Hash()
}

//go:wasmexport gd_on_callable_sorted
func gd_on_callable_sorted(a, b CallableID) int64 {
	return int64(callables.Get(a).Compare(callables.Get(b)))
}

//go:wasmexport gd_on_callable_string
func gd_on_callable_string(c CallableID) String {
	return callables.Get(c).UnsafeString()
}

//go:wasmexport gd_on_callable_length
func gd_on_callable_length(c CallableID) int64 {
	return int64(callables.Get(c).NumIn())
}

// Cross-memory helpers for transferring data between Go and engine address spaces.

var wasmResultBufs [2]Pointer
var wasmResultIdx int
var wasmArgBuf Pointer

var wasmSetup = sync.OnceFunc(func() {
	wasmArgBuf = Malloc(64 * 64)
	for i := range wasmResultBufs {
		wasmResultBufs[i] = Malloc(64 * 64)
		Clear(wasmResultBufs[i], 64*64)
	}
})

func makeResult(shape Shape) Pointer {
	wasmSetup()
	wasmResultIdx ^= 1
	return wasmResultBufs[wasmResultIdx]
}

func loadResult[T ~unsafe.Pointer | *Variant | *CallError](shape Shape, result T, from Pointer) {
	wasmSetup()
	if from == 0 {
		panic("nil pointer dereference")
	}
	data := unsafe.Pointer(result)
	done := 0
	size := shape.SizeResult()
	if size == 0 {
		return
	}
	defer Clear(Pointer(from), Int(size))
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(data, done)) = Pointer(from + Pointer(done)).Uint32()
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = Pointer(from + Pointer(done)).Uint16()
			done += 2
			size -= 2
		default:
			*(*uint8)(unsafe.Add(data, done)) = Pointer(from + Pointer(done)).Byte()
			done += 1
			size -= 1
		}
	}
}

func copyVariants[T ~unsafe.Pointer | *Variant](args T, n int) Pointer {
	wasmSetup()
	var offset int
	var data = unsafe.Pointer(args)
	for i := range n {
		Pointer(wasmArgBuf + Pointer(offset)).SetBits128(*(*[2]uint64)(unsafe.Add(data, uintptr(i*24))))
		Pointer(wasmArgBuf + Pointer(offset+16)).SetUint64(*(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
		offset += 24
	}
	return wasmArgBuf
}

func copyArguments(shape Shape, args unsafe.Pointer) Pointer {
	wasmSetup()
	if args == nil {
		return 0
	}
	bytes := shape.SizeArguments()
	buf := unsafe.Slice((*byte)(args), bytes)
	ptr := wasmArgBuf
	off := Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			Pointer(ptr + off).SetUint64(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			Pointer(ptr + off).SetUint32(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			Pointer(ptr + off).SetUint16(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		default:
			Pointer(ptr + off).SetByte(*(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func readVariant(addr Pointer) Variant {
	if addr == 0 {
		panic("nil pointer dereference")
	}
	var v Variant
	v[0] = Pointer(addr).Uint64()
	v[1] = Pointer(addr + 8).Uint64()
	v[2] = Pointer(addr + 16).Uint64()
	return v
}

// Object construction and identity

//go:wasmimport gd object_make
func gd_object_make(name StringName) Object

func MakeObject(name StringName) Object { return gd_object_make(name) }

//go:wasmimport gd object_name
func gd_object_name(obj Object) StringName

func (obj Object) Name() StringName { return gd_object_name(obj) }

//go:wasmimport gd object_type
func gd_object_type(name StringName) ObjectType

func ObjectTypeTag(name StringName) ObjectType { return gd_object_type(name) }

//go:wasmimport gd object_cast
func gd_object_cast(obj Object, to ObjectType) Object

func (obj Object) Cast(to ObjectType) Object { return gd_object_cast(obj, to) }

//go:wasmimport gd object_lookup
func gd_object_lookup(id ObjectID) Object

func (id ObjectID) Lookup() Object { return gd_object_lookup(id) }

//go:wasmimport gd object_global
func gd_object_global(name StringName) Object

func ObjectGlobal(name StringName) Object { return gd_object_global(name) }

//go:wasmimport gd object_id
func gd_object_id(obj Object) ObjectID

func (obj Object) ID() ObjectID { return gd_object_id(obj) }

//go:wasmimport gd object_id_inside_variant
func gd_object_id_inside_variant(v1, v2, v3 uint64) ObjectID

func ObjectIDInsideVariant(v Variant) ObjectID {
	return gd_object_id_inside_variant(v[0], v[1], v[2])
}

//go:wasmimport gd object_unsafe_free
func gd_object_unsafe_free(obj Object)

func (obj Object) Free() { gd_object_unsafe_free(obj) }

// Object method calls

//go:wasmimport gd object_method_lookup
func gd_object_method_lookup(class, method StringName, hash int64) MethodForClass

func MethodLookup(class, method StringName, hash int64) MethodForClass {
	return gd_object_method_lookup(class, method, hash)
}

//go:wasmimport gd object_call
func gd_object_call(obj Object, method MethodForClass, result Pointer, argc Int, args Pointer, err Pointer)

func (obj Object) Call(method MethodForClass, args ...Variant) (Variant, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(SizeCallError)
	gd_object_call(obj, method, Pointer(mem_result), Int(len(args)), Pointer(mem_args), Pointer(mem_err))
	var result Variant
	var errResult CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &errResult, mem_err)
	return result, errResult
}

//go:wasmimport gd object_unsafe_call
func gd_object_unsafe_call(obj Object, fn MethodForClass, result Pointer, shape uint64, args Pointer)

func (obj Object) UnsafeCall(fn MethodForClass, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_object_unsafe_call(obj, fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

// Extension instance management

//go:wasmimport gd object_extension_setup
func gd_object_extension_setup(obj Object, name StringName, inst ExtensionInstanceID)

func (obj Object) ExtensionSetup(name StringName, inst ExtensionInstanceID) {
	gd_object_extension_setup(obj, name, inst)
}

//go:wasmimport gd object_extension_fetch
func gd_object_extension_fetch(obj Object) ExtensionInstanceID

func (obj Object) ExtensionFetch() ExtensionInstanceID { return gd_object_extension_fetch(obj) }

//go:wasmimport gd object_extension_close
func gd_object_extension_close(obj Object)

func (obj Object) ExtensionClose() { gd_object_extension_close(obj) }

// Script instance management

//go:wasmimport gd object_script_make
func gd_object_script_make(fn ExtensionInstanceID) ScriptInstance

func ScriptMake(fn ExtensionInstanceID) ScriptInstance { return gd_object_script_make(fn) }

//go:wasmimport gd object_script_call
func gd_object_script_call(obj Object, name StringName, result Pointer, argc Int, args Pointer, err Pointer)

func (obj Object) ScriptCall(name StringName, args ...Variant) (Variant, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(SizeCallError)
	gd_object_script_call(obj, name, Pointer(mem_result), Int(len(args)), Pointer(mem_args), Pointer(mem_err))
	var result Variant
	var err CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &err, mem_err)
	return result, err
}

//go:wasmimport gd object_script_setup
func gd_object_script_setup(obj Object, script ScriptInstance)

func (obj Object) ScriptSetup(script ScriptInstance) { gd_object_script_setup(obj, script) }

//go:wasmimport gd object_script_fetch
func gd_object_script_fetch(obj Object, language Object) ScriptInstance

func (obj Object) ScriptFetch(language Object) ScriptInstance {
	return gd_object_script_fetch(obj, language)
}

//go:wasmimport gd object_script_defines_method
func gd_object_script_defines_method(obj Object, method StringName) uint32

func (obj Object) ScriptDefinesMethod(method StringName) bool {
	return gd_object_script_defines_method(obj, method) != 0
}

//go:wasmimport gd object_script_property_state_add
func gd_object_script_property_state_add(fn FunctionID, arg uint32, name StringName, s1, s2, s3 uint64)

func ScriptPropertyStateAdd(fn FunctionID, arg Pointer, name StringName, state Variant) {
	gd_object_script_property_state_add(fn, uint32(arg), name, state[0], state[1], state[2])
}

//go:wasmimport gd object_script_placeholder_create
func gd_object_script_placeholder_create(language, script, owner Object) ScriptInstance

func ScriptPlaceholderCreate(language, script, owner Object) ScriptInstance {
	return gd_object_script_placeholder_create(language, script, owner)
}

//go:wasmimport gd object_script_placeholder_update
func gd_object_script_placeholder_update(script ScriptInstance, array Array, dict Dictionary)

func ScriptPlaceholderUpdate(script ScriptInstance, array Array, dict Dictionary) {
	gd_object_script_placeholder_update(script, array, dict)
}

// Variant operations

type VariantOperator = uint32

//go:wasmimport gd variant_zero
func gd_variant_zero(result Pointer)

func ZeroVariant() Variant {
	mem := makeResult(SizeVariant)
	gd_variant_zero(Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_copy
func gd_variant_copy(v1, v2, v3 uint64, result Pointer)

func (v Variant) Copy() Variant {
	mem := makeResult(SizeVariant)
	gd_variant_copy(v[0], v[1], v[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_call
func gd_variant_call(v1, v2, v3 uint64, method StringName, result Pointer, argc Int, args Pointer, err Pointer)

func (v Variant) VariantCall(method StringName, args ...Variant) (Variant, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(SizeCallError)
	gd_variant_call(v[0], v[1], v[2], method, Pointer(mem_result), Int(len(args)), Pointer(mem_args), Pointer(mem_err))
	var result Variant
	var err CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &err, mem_err)
	return result, err
}

//go:wasmimport gd variant_eval
func gd_variant_eval(op uint32, a1, a2, a3, b1, b2, b3 uint64, result Pointer) uint32

func VariantEval(op VariantOperator, a, b Variant) (Variant, bool) {
	mem := makeResult(SizeVariant)
	r := gd_variant_eval(op, a[0], a[1], a[2], b[0], b[1], b[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_hash
func gd_variant_hash(v1, v2, v3 uint64) int64

func (v Variant) Hash() Int { return gd_variant_hash(v[0], v[1], v[2]) }

//go:wasmimport gd variant_bool
func gd_variant_bool(v1, v2, v3 uint64) uint32

func (v Variant) Bool() bool {
	return gd_variant_bool(v[0], v[1], v[2]) != 0
}

//go:wasmimport gd variant_text
func gd_variant_text(v1, v2, v3 uint64) String

func (v Variant) Text() String {
	return gd_variant_text(v[0], v[1], v[2])
}

//go:wasmimport gd variant_type
func gd_variant_type(v1, v2, v3 uint64) VariantType

func (v Variant) Type() VariantType {
	return gd_variant_type(v[0], v[1], v[2])
}

// Deep variant operations

//go:wasmimport gd variant_deep_copy
func gd_variant_deep_copy(v1, v2, v3 uint64, result Pointer)

func (v Variant) DeepCopy() Variant {
	mem := makeResult(SizeVariant)
	gd_variant_deep_copy(v[0], v[1], v[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_deep_hash
func gd_variant_deep_hash(v1, v2, v3 uint64, recursion Int) Int

func (v Variant) DeepHash(recursion Int) Int {
	return gd_variant_deep_hash(v[0], v[1], v[2], recursion)
}

// Variant get/set/has

//go:wasmimport gd variant_get_index
func gd_variant_get_index(v1, v2, v3, k1, k2, k3 uint64, result Pointer) uint32

func (v Variant) GetIndex(key Variant) (Variant, bool) {
	mem := makeResult(SizeVariant)
	r := gd_variant_get_index(v[0], v[1], v[2], key[0], key[1], key[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_get_array
func gd_variant_get_array(v1, v2, v3 uint64, idx Int, result Pointer, err Pointer) uint32

func (v Variant) GetArray(idx Int) (Variant, bool, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_err := makeResult(SizeCallError)
	r := gd_variant_get_array(v[0], v[1], v[2], idx, Pointer(mem_result), Pointer(mem_err))
	var result Variant
	var callErr CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &callErr, mem_err)
	return result, r != 0, callErr
}

//go:wasmimport gd variant_get_field
func gd_variant_get_field(v1, v2, v3 uint64, field StringName, result Pointer) uint32

func (v Variant) GetField(field StringName) (Variant, bool) {
	mem := makeResult(SizeVariant)
	r := gd_variant_get_field(v[0], v[1], v[2], field, Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_set_index
func gd_variant_set_index(v1, v2, v3, k1, k2, k3, val1, val2, val3 uint64) uint32

func (v Variant) SetIndex(key, val Variant) bool {
	return gd_variant_set_index(v[0], v[1], v[2], key[0], key[1], key[2], val[0], val[1], val[2]) != 0
}

//go:wasmimport gd variant_set_array
func gd_variant_set_array(v1, v2, v3 uint64, idx Int, val1, val2, val3 uint64, err Pointer) uint32

func (v Variant) SetArray(idx Int, val Variant) (bool, CallError) {
	mem_err := makeResult(SizeCallError)
	r := gd_variant_set_array(v[0], v[1], v[2], idx, val[0], val[1], val[2], Pointer(mem_err))
	var callErr CallError
	loadResult(SizeCallError, &callErr, mem_err)
	return r != 0, callErr
}

//go:wasmimport gd variant_set_field
func gd_variant_set_field(v1, v2, v3 uint64, field StringName, val1, val2, val3 uint64) uint32

func (v Variant) SetField(field StringName, value Variant) bool {
	return gd_variant_set_field(v[0], v[1], v[2], field, value[0], value[1], value[2]) != 0
}

//go:wasmimport gd variant_has_index
func gd_variant_has_index(v1, v2, v3, i1, i2, i3 uint64) uint32

func (v Variant) HasIndex(index Variant) bool {
	return gd_variant_has_index(v[0], v[1], v[2], index[0], index[1], index[2]) != 0
}

//go:wasmimport gd variant_has_method
func gd_variant_has_method(v1, v2, v3 uint64, method StringName) uint32

func (v Variant) HasMethod(method StringName) bool {
	return gd_variant_has_method(v[0], v[1], v[2], method) != 0
}

// Unsafe variant operations

//go:wasmimport gd variant_unsafe_call
func gd_variant_unsafe_call(fn FunctionID, result Pointer, shape uint64, args Pointer)

func VariantUnsafeCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_call(fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_eval
func gd_variant_unsafe_eval(fn FunctionID, result Pointer, shape uint64, args Pointer)

func VariantUnsafeEval(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_eval(fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_free
func gd_variant_unsafe_free(v1, v2, v3 uint64)

func (v Variant) UnsafeFree() {
	gd_variant_unsafe_free(v[0], v[1], v[2])
}

//go:wasmimport gd variant_unsafe_make_native
func gd_variant_unsafe_make_native(vtype VariantType, v1, v2, v3 uint64, shape uint64, result Pointer)

func VariantUnsafeMakeNative(vtype VariantType, v Variant, shape uint64, result unsafe.Pointer) {
	mem := makeResult(Shape(shape))
	gd_variant_unsafe_make_native(vtype, v[0], v[1], v[2], shape, Pointer(mem))
	loadResult(Shape(shape), result, mem)
}

//go:wasmimport gd variant_unsafe_from_native
func gd_variant_unsafe_from_native(vtype VariantType, result Pointer, shape uint64, args Pointer)

func VariantUnsafeFromNative(vtype VariantType, shape uint64, args unsafe.Pointer) Variant {
	mem_result := makeResult(SizeVariant)
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_from_native(vtype, Pointer(mem_result), shape, Pointer(mem_args))
	var result Variant
	loadResult(SizeVariant, &result, mem_result)
	return result
}

//go:wasmimport gd variant_unsafe_internal_pointer
func gd_variant_unsafe_internal_pointer(vtype VariantType, v1, v2, v3 uint64) Pointer

func VariantUnsafeInternalPointer(vtype VariantType, v Variant) Pointer {
	return gd_variant_unsafe_internal_pointer(vtype, v[0], v[1], v[2])
}

//go:wasmimport gd variant_unsafe_set_field
func gd_variant_unsafe_set_field(setter FunctionID, shape uint64, args Pointer)

func VariantUnsafeSetField(setter FunctionID, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_set_field(setter, shape, Pointer(mem_args))
}

//go:wasmimport gd variant_unsafe_set_array
func gd_variant_unsafe_set_array(vtype VariantType, idx Int, shape uint64, args Pointer)

func VariantUnsafeSetArray(vtype VariantType, idx Int, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_set_array(vtype, idx, shape, Pointer(mem_args))
}

//go:wasmimport gd variant_unsafe_set_index
func gd_variant_unsafe_set_index(vtype VariantType, shape uint64, args Pointer)

func VariantUnsafeSetIndex(vtype VariantType, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_set_index(vtype, shape, Pointer(mem_args))
}

//go:wasmimport gd variant_unsafe_get_field
func gd_variant_unsafe_get_field(getter FunctionID, result Pointer, shape uint64, args Pointer)

func VariantUnsafeGetField(getter FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_get_field(getter, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_get_array
func gd_variant_unsafe_get_array(vtype VariantType, idx Int, result Pointer, shape uint64, args Pointer)

func VariantUnsafeGetArray(vtype VariantType, idx Int, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_get_array(vtype, idx, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_get_index
func gd_variant_unsafe_get_index(vtype VariantType, result Pointer, shape uint64, args Pointer)

func VariantUnsafeGetIndex(vtype VariantType, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_get_index(vtype, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

// Iterator operations

//go:wasmimport gd iterator_make
func gd_iterator_make(v1, v2, v3 uint64, result Pointer, err Pointer)

func (v Variant) IteratorMake(result unsafe.Pointer, err unsafe.Pointer) {
	mem_result := makeResult(SizeVariant)
	mem_err := makeResult(SizeCallError)
	gd_iterator_make(v[0], v[1], v[2], Pointer(mem_result), Pointer(mem_err))
	loadResult(SizeVariant, result, mem_result)
	loadResult(SizeCallError, err, mem_err)
}

//go:wasmimport gd iterator_next
func gd_iterator_next(v1, v2, v3 uint64, iter Pointer, err Pointer) uint32

func (v Variant) IteratorNext(iter unsafe.Pointer, err unsafe.Pointer) bool {
	mem_iter := makeResult(SizeVariant)
	// Copy iter into engine memory, call, then copy back.
	mem_args := copyArguments(SizeVariant, iter)
	mem_err := makeResult(SizeCallError)
	r := gd_iterator_next(v[0], v[1], v[2], Pointer(mem_args), Pointer(mem_err))
	loadResult(SizeVariant, iter, mem_args)
	_ = mem_iter
	loadResult(SizeCallError, err, mem_err)
	return r != 0
}

//go:wasmimport gd iterator_load
func gd_iterator_load(v1, v2, v3, i1, i2, i3 uint64, result Pointer, err Pointer)

func (v Variant) IteratorLoad(iter Variant, result unsafe.Pointer, err unsafe.Pointer) {
	mem_result := makeResult(SizeVariant)
	mem_err := makeResult(SizeCallError)
	gd_iterator_load(v[0], v[1], v[2], iter[0], iter[1], iter[2], Pointer(mem_result), Pointer(mem_err))
	loadResult(SizeVariant, result, mem_result)
	loadResult(SizeCallError, err, mem_err)
}
