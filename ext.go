//go:build !generate

package gd

import (
	"runtime/cgo"
	"unsafe"

	"runtime.link/api"
	"runtime.link/api/call"
)

type ExtensionAPI struct {
	api.Specification

	GetGodotVersion func(*ExtensionGodotVersion) `call:"get_godot_version func(&void)"`

	Allocate   func(uintptr) unsafe.Pointer                 `call:"mem_alloc func(size_t)$void"`
	Reallocate func(unsafe.Pointer, uintptr) unsafe.Pointer `call:"mem_realloc func($void,size_t)$void"`
	Release    func(unsafe.Pointer)                         `call:"mem_free func($void)"`

	PrintError              func(code, function, file string, line int32, notifyEditor bool)                             `call:"print_error func(&char,&char,&char,int32_t,bool)"`
	PrintErrorMessage       func(code, message, function, file string, line int32, notifyEditor bool)                    `call:"print_error func(&char,&char,&char,&char,int32_t,bool)"`
	PrintWarning            func(code, function, file string, line int32, notifyEditor bool)                             `call:"print_warning func(&char,&char,&char,int32_t,bool)"`
	PrintWarningMessage     func(code, message, function, file string, line int32, notifyEditor bool)                    `call:"print_warning func(&char,&char,&char,&char,int32_t,bool)"`
	PrintScriptError        func(code, function, file string, line int32, stackTrace string, notifyEditor bool)          `call:"print_script_error func(&char,&char,&char,int32_t,&char,bool)"`
	PrintScriptErrorMessage func(code, message, function, file string, line int32, stackTrace string, notifyEditor bool) `call:"print_script_error func(&char,&char,&char,&char,int32_t,&char,bool)"`

	GetNativeStructSize func(name *StringName) uint64 `call:"get_native_struct_size func(&void)size_t"`

	Free struct {
		StringName func() func(*StringName) `call:"variant_get_ptr_destructor func(-int=21)func($void)"`
	}

	Variants struct {
		// Copy dst must be unititialized
		Copy func(dst, src *Variant) `call:"variant_new_copy func(+void,$void)"`
		// Zero dst must be unititialized
		Zero          func(dst *Variant)                                                                        `call:"variant_new_nil func(+void)"`
		Free          func(*Variant)                                                                            `call:"variant_destroy func($void)"`
		Call          func(self *Variant, method *StringName, args []*Variant, ret *Variant, err CallError)     `call:"variant_call func($void,&void,&void,-int64_t=@3,+void,+void)"`
		CallStatic    func(vtype VariantType, method *StringName, args []*Variant, ret *Variant, err CallError) `call:"variant_call_static func(int,&void,&void,&void,-int64_t=@3,+void,+void)"`
		Evaluate      func(operator Operator, a, b *Variant, ret *Variant, ok *bool)                            `call:"variant_evaluate func(int,&void,&void,+void,+bool)"`
		Set           func(self, key, val *Variant, ok *bool)                                                   `call:"variant_set func(&void,&void,+bool)"`
		SetNamed      func(self *Variant, key *StringName, val *Variant, ok *bool)                              `call:"variant_set_named func(&void,&void,+bool)"`
		SetKeyed      func(self, key, val *Variant, ok *bool)                                                   `call:"variant_set_keyed func(&void,&void,+bool)"`
		SetIndex      func(self *Variant, index int64, val *Variant, ok *bool)                                  `call:"variant_set_indexed func(&void,int64_t,+bool)"`
		Get           func(self, key *Variant, ret *Variant, ok *bool)                                          `call:"variant_get func(&void,&void,+void,+bool)"`
		GetNamed      func(self *Variant, key *StringName, ret *Variant, ok *bool)                              `call:"variant_get_named func(&void,&void,+void,+bool)"`
		GetKeyed      func(self *Variant, key *Variant, ret *Variant, ok *bool)                                 `call:"variant_get_keyed func(&void,&void,+void,+bool)"`
		GetIndex      func(self *Variant, index int64, ret *Variant, ok *bool)                                  `call:"variant_get_indexed func(&void,int64_t,+void,+bool)"`
		Iterator      func(self *Variant, iter *Iterator) bool                                                  `call:"variant_iter_init func(&void,+void,+bool)"`
		Next          func(self *Variant, iterator Iterator, ret *Variant) bool                                 `call:"variant_iter_next func(&void,&void,+void,+bool)"`
		IteratorValue func(self *Variant, iterator Iterator, ret *Variant) bool                                 `call:"variant_iter_get func(&void,&void,+void,+bool)"`
		Hash          func(self *Variant) int64                                                                 `call:"variant_hash func(&void)int64_t"`
		RecursiveHash func(self *Variant) int64                                                                 `call:"variant_recursive_hash func(&void)int64_t"`
		Compare       func(self, variant *Variant) bool                                                         `call:"variant_hash_compare func(&void,&void)bool"`
		ToBool        func(self *Variant) bool                                                                  `call:"variant_booleanize func(&void)bool"`
		Duplicate     func(self *Variant, ret *Variant, deep bool)                                              `call:"variant_duplicate func(&void,+void,bool)"`
		ToString      func(self *Variant, out *String)                                                          `call:"variant_stringify func(&void,+void)"`
		Type          func(self *Variant) VariantType                                                           `call:"variant_get_type func(&void)int"`
		HasMethod     func(self *Variant, method *StringName) bool                                              `call:"variant_has_method func(&void,&void)bool"`
		HasMember     func(self *Variant, member *StringName) bool                                              `call:"variant_has_member func(&void,&void)bool"`
		HasKey        func(self *Variant, key *Variant) (bool, bool)                                            `call:"variant_has_key func(&void,&void,+bool)bool"`
		// TypeName dst must be unititialized
		TypeName         func(self *Variant, dst *String)        `call:"variant_get_type_name func(&void,+void)"`
		CanConvert       func(fromType, toType VariantType) bool `call:"variant_can_convert func(int,int)bool"`
		CanConvertStrict func(fromType, toType VariantType) bool `call:"variant_can_convert_strict func(int,int)bool"`

		Encoder func(VariantType) func(out *Variant, in unsafe.Pointer) `call:"get_variant_from_type_constructor func(int)func(+void,&void)"`
		Decoder func(VariantType) func(out unsafe.Pointer, in *Variant) `call:"get_variant_to_type_constructor func(int)func(+void,&void)"`

		Evaulator func(Operator) func(a *Variant, b *Variant, ret *Variant) `call:"variant_get_ptr_operator_evaluator func(int)func(&void,&void,+void)"`

		BuiltinMethodCaller func(VariantType, *StringName, int64) func(base unsafe.Pointer, args []unsafe.Pointer, ret unsafe.Pointer) `call:"variant_get_ptr_builtin_method func(int,&void,int64_t)func(&void,&void,+void,-int=@2)"`
		Constructor         func(VariantType, int32) func(base unsafe.Pointer, args *unsafe.Pointer)                                   `call:"variant_get_ptr_constructor func(int,int32_t)func(+void,&void)"`
		Destructor          func(VariantType) func(unsafe.Pointer)                                                                     `call:"variant_get_ptr_destructor func(int)func($void)"`

		Construct func(t VariantType, out *Variant, args []*Variant, err CallError) `call:"variant_construct func(int,+void,&void,-int64_t=@3,+void)"`

		Setter      func(VariantType, *StringName) func(*Variant, unsafe.Pointer)             `call:"variant_get_ptr_setter func(int,&void)func(&void,&void)"`
		Getter      func(VariantType, *StringName) func(base *Variant, ret unsafe.Pointer)    `call:"variant_get_ptr_getter func(int,&void)func(&void,+void)"`
		IndexSetter func(VariantType) func(*Variant, int64, unsafe.Pointer)                   `call:"variant_get_ptr_indexed_setter func(int)func(&void,int64_t,&void)"`
		IndexGetter func(VariantType) func(val *Variant, idx int64, ret unsafe.Pointer)       `call:"variant_get_ptr_indexed_getter func(int)func(&void,int64_t)&void"`
		KeySetter   func(VariantType) func(*Variant, *Variant, unsafe.Pointer)                `call:"variant_get_ptr_keyed_setter func(int)func(&void,&void,&void)"`
		KeyGetter   func(VariantType) func(val *Variant, key *Variant, ret unsafe.Pointer)    `call:"variant_get_ptr_keyed_getter func(int)func(&void,&void)&void"`
		KeyChecker  func(VariantType) func(*Variant, *Variant) uint32                         `call:"variant_get_ptr_keyed_checker func(int)func(&void,&void)uint32_t"`
		Constant    func(t VariantType, name *StringName, ret *Variant)                       `call:"variant_get_constant_value func(int,&void,+void)"`
		Utility     func(VariantType, *StringName, int64) func(val *Variant, args []*Variant) `call:"variant_get_ptr_utility_function func(int,&void,int64_t)func(+void,&void,-int=@2)"`
	}

	Strings struct {
		New        func(dst *String, s string) `call:"string_new_with_utf8_chars_and_len func(+void,&char,int64_t=@2)"`
		Get        func(*String, []byte)       `call:"string_to_utf8_chars func(&void,&char,-int64_t=@2)"`
		Index      func(*String, int64) rune   `call:"string_operator_index func(&void,int64_t)&char32_t"`
		Append     func(*String, *String)      `call:"string_operator_plus_eq_string func(&void,&void)"`
		AppendRune func(*String, rune)         `call:"string_operator_plus_eq_char func(&void,&char32_t)"`
		Resize     func(*String, int64)        `call:"string_resize func(&void,int64_t)"`
	}

	StringNames struct {
		New func(dst *StringName, s string) `call:"string_name_new_with_utf8_chars_and_len func(+void,&char,-int64_t=@2)"`
	}

	XMLParser struct {
		OpenBuffer func(*XMLParser, []byte) Error `call:"xml_parser_open_buffer func(&void,&char,int64_t=@2)int64_t"`
	}

	FileAccess struct {
		StoreBuffer func(*FileAccess, []byte) Error `call:"file_access_store_buffer func(&void,&char,int64_t=@2)int64_t"`
		ReadBuffer  func(*FileAccess, []byte) Error `call:"file_access_get_buffer func(&void,&char,int64_t=@2)int64_t"`
	}

	PackedByteArray struct {
		Index func(*PackedByteArray, int64) *byte `call:"packed_byte_array_operator_index func(&void,int64_t)&char"`
	}
	PackedColorArray struct {
		Index func(*PackedColorArray, int64) *Color `call:"packed_color_array_operator_index func(&void,int64_t)&Color"`
	}
	PackedFloat32Array struct {
		Index func(*PackedFloat32Array, int64) *float32 `call:"packed_float32_array_operator_index func(&void,int64_t)&float"`
	}
	PackedFloat64Array struct {
		Index func(*PackedFloat64Array, int64) *float64 `call:"packed_float64_array_operator_index func(&void,int64_t)&double"`
	}
	PackedInt32Array struct {
		Index func(*PackedInt32Array, int64) *int32 `call:"packed_int32_array_operator_index func(&void,int64_t)&int32_t"`
	}
	PackedInt64Array struct {
		Index func(*PackedInt64Array, int64) *int64 `call:"packed_int64_array_operator_index func(&void,int64_t)&int64_t"`
	}
	PackedStringArray struct {
		Index func(*PackedStringArray, int64) *String `call:"packed_string_array_operator_index func(&void,int64_t)&char"`
	}
	PackedVector2Array struct {
		Index func(*PackedVector2Array, int64) *Vector2 `call:"packed_vector2_array_operator_index func(&void,int64_t)&Vector2"`
	}
	PackedVector3Array struct {
		Index func(*PackedVector3Array, int64) *Vector3 `call:"packed_vector3_array_operator_index func(&void,int64_t)&Vector3"`
	}
	Array struct {
		Index    func(*Array, int64) *Variant                                         `call:"array_operator_index func(&void,int64_t)&void"`
		Set      func(*Array, *Array)                                                 `call:"array_ref func(&void,&void)"`
		SetTyped func(a *Array, t VariantType, className *StringName, script *Script) `call:"array_set_typed func(&void,int,&void,&void)"`
	}
	Dictionary struct {
		Index func(d *Dictionary, key *Variant, ret *Variant) `call:"dictionary_operator_index func(&void,&void,&void)"`
	}

	Object struct {
		Call                func(method MethodBind, obj uintptr, arg []*Variant, ret *Variant, err CallError)     `call:"object_method_bind_call func(&void,&void,&void,-int64_t=@3,+void,+void)"`
		UnsafeCall          func(method MethodBind, obj uintptr, args []unsafe.Pointer, ret unsafe.Pointer)       `call:"object_method_bind_ptrcall func(&void,&void,&void,&void)"`
		Destroy             func(obj uintptr)                                                                     `call:"object_destroy func(&void)"`
		GetSingleton        func(name *StringName) uintptr                                                        `call:"global_get_singleton func(&void)&void"`
		GetInstanceBinding  func(obj uintptr, token unsafe.Pointer, callbacks *InstanceBindingCallbacks)          `call:"object_get_instance_binding func(&void,$void,&void)"`
		SetInstanceBinding  func(obj uintptr, token, binding unsafe.Pointer, callbacks *InstanceBindingCallbacks) `call:"object_set_instance_binding func(&void,$void,&void)"`
		FreeInstanceBinding func(obj uintptr, token unsafe.Pointer)                                               `call:"object_free_instance_binding func(&void,$void)"`
		SetInstance         func(obj uintptr, name *StringName, value cgo.Handle)                                 `call:"object_set_instance func(&void,&void,&void)"`
		GetClassName        func(obj uintptr, library ExtensionToken, out *StringName)                            `call:"object_get_class_name func(&void,+void)"`
		CastTo              func(obj uintptr, class ClassTag) uintptr                                             `call:"object_cast_to func(&void,&void)&void"`
		GetFromID           func(id InstanceID) uintptr                                                           `call:"object_get_instance_from_id func(uint64_t)&void"`
		GetID               func(obj uintptr) InstanceID                                                          `call:"object_get_instance_id func(&void)uint64_t"`
		GetFromReference    func(ref unsafe.Pointer) uintptr                                                      `call:"ref_get_object func(&void)&void"`
		Reference           func(ref unsafe.Pointer, obj uintptr)                                                 `call:"ref_set_object func(&void,&void)"`
	}

	Scripts struct {
		Create            func(info *ScriptInstanceInfo, script unsafe.Pointer) Script             `call:"script_instance_create2 func(&void,&void)&void"`
		CreatePlaceholder func(lang *ScriptLanguage, script *Script, owner *uintptr) Script        `call:"placeholder_script_instance_create func(&void,&void,&void)&void"`
		UpdatePlaceholder func(script Script, properties *ArrayOf[Dictionary], values *Dictionary) `call:"placeholder_script_instance_update func(&void)"`
		Get               func(obj *uintptr, lang *ScriptLanguage) Script                          `call:"object_get_script_instance func(&void)&void"`
	}

	Callables struct {
		Create      func(callable *Callable, info *CallableCustomInfo)            `call:"callable_custom_create func(+void,&void)"`
		GetUserData func(callable *Callable, token unsafe.Pointer) unsafe.Pointer `call:"callable_custom_get_userdata func(&void,&void)$void"`
	}

	ClassDB struct {
		CreateObject  func(class *StringName) uintptr                                    `call:"classdb_construct_object func(&void)$void"`
		GetMethodBind func(class *StringName, method *StringName, hash int64) MethodBind `call:"classdb_get_method_bind func(&void,&void,int64_t)void"`
		GetClassTag   func(class *StringName) ClassTag                                   `call:"classdb_get_class_tag func(&void)void"`

		RegisterClass                 func(library ExtensionToken, name, extends *StringName, info *ClassCreationInfo)                             `call:"classdb_register_extension_class2 func(&void,&void,&void,&void)"`
		RegisterClassMethod           func(library ExtensionToken, class *StringName, info *ClassMethodInfo)                                       `call:"classdb_register_extension_class_method func(&void,&void,&void)"`
		RegisterClassIntegerConstant  func(library ExtensionToken, class, enum, name *StringName, value int64, bitfield bool)                      `call:"classdb_register_extension_class_integer_constant func(&void,&void,&void,&void,&void,int64_t)"`
		RegisterClassProperty         func(library ExtensionToken, class *StringName, info *PropertyInfo, getter, setter *StringName)              `call:"classdb_register_extension_class_property func(&void,&void,&void,&void,&void)"`
		RegisterClassPropertyIndexed  func(library ExtensionToken, class *StringName, info *PropertyInfo, getter, setter *StringName, index int64) `call:"classdb_register_extension_class_property_indexed func(&void,&void,&void,&void,&void,int64_t)"`
		RegisterClassPropertyGroup    func(library ExtensionToken, class *StringName, group, prefix *String)                                       `call:"classdb_register_extension_class_property_group func(&void,&void,&void,&void)"`
		RegisterClassPropertySubGroup func(library ExtensionToken, class *StringName, subGroup, prefix *String)                                    `call:"classdb_register_extension_class_property_subgroup func(&void,&void,&void,&void,&void)"`
		RegisterClassSignal           func(library ExtensionToken, class, signal *StringName, args []PropertyInfo)                                 `call:"classdb_register_extension_class_signal func(&void,&void,&void,&void,-int64_t=@4)"`
		UnregisterClass               func(library ExtensionToken, class *StringName)                                                              `call:"classdb_unregister_extension_class func(&void,&void)"`
		GetLibraryPath                func(library ExtensionToken, out *String)                                                                    `call:"get_library_path func(&void,&void)"`
	}

	EditorPlugins struct {
		Add    func(plugin *StringName) `call:"editor_add_plugin func(&void)"`
		Remove func(plugin *StringName) `call:"editor_remove_plugin func(&void)"`
	}
}

type CallErrorType int32

const (
	OK CallErrorType = iota
	ErrInvalidMethod
	ErrInvalidArgument
	ErrTooManyArguments
	ErrTooFewArguments
	ErrInstanceIsNil
	ErrMethodNotConst
)

type CallError struct {
	ErrorType CallErrorType
	Argument  int32
	Expected  int32
}

type ExtensionGodotVersion struct {
	Major  uint32
	Minor  uint32
	Patch  uint32
	String *byte
}

type Iterator uintptr

type MethodBind uintptr

type ClassTag uintptr
type ExtensionToken uintptr

type InstanceID uint64

type Operator int32

const (
	Equal Operator = iota
	NotEqual
	Less
	LessEqual
	Greater
	GreaterEqual
	Add
	Subtract
	Multiply
	Divide
	Negate
	Positive
	Module
	Power
	ShiftLeft
	ShiftRight
	BitAnd
	BitOr
	BitXor
	BitNegate
	LogicalAnd
	LogicalOr
	LogicalXor
	LogicalNegate
	In
)

type ExtensionInitialization[T any] struct {
	MinimumInitializationLevel ExtensionInitializationLevel
	Userdata                   T
	Initialize                 call.Back[func(userdata T, level ExtensionInitializationLevel)]
	Deinitialize               call.Back[func(userdata T, level ExtensionInitializationLevel)]
}

type InstanceBindingCallbacks struct {
	Create    call.Back[func(token, instance unsafe.Pointer) unsafe.Pointer]
	Free      call.Back[func(token, instance, binding unsafe.Pointer)]
	Reference call.Back[func(token, binding unsafe.Pointer, reference bool) bool]
}

type PropertyInfo struct {
	Type       VariantType
	Name       StringName
	ClassName  StringName
	Hint       uint32 // PropertyHint
	HintString String
	Usage      uint32 // PropertyUsageFlags
}

type MethodInfo struct {
	Name                 StringName
	ReturnValue          PropertyInfo
	Flags                uint32 // MethodFlags
	ID                   int32
	ArgumentCount        uint32
	Arguments            *PropertyInfo
	DefaultArgumentCount uint32
	DefaultArguments     *Variant
}

type ClassCreationInfo struct {
	IsVirtual  bool
	IsAbstract bool
	IsExposed  bool

	Set call.Back[func(instance cgo.Handle, name *StringName, value *Variant)]
	Get call.Back[func(instance cgo.Handle, name *StringName, ret *Variant)]

	GetPropertyList   call.Back[func(instance cgo.Handle, len *uint32) *PropertyInfo]
	FreePropertyList  call.Back[func(instance cgo.Handle, list *PropertyInfo)]
	PropertyCanRevert call.Back[func(instance cgo.Handle, name *StringName) bool]
	PropertyGetRevert call.Back[func(instance cgo.Handle, name *StringName, ret *Variant)]
	ValidateProperty  call.Back[func(instance cgo.Handle, name *StringName, property *PropertyInfo) bool]
	Notification      call.Back[func(instance cgo.Handle, notification int32, reversed bool)]
	ToString          call.Back[func(instance cgo.Handle, valid *bool, out *String)]
	Reference         call.Back[func(instance cgo.Handle)]
	Unreference       call.Back[func(instance cgo.Handle)]

	// Mandatory
	CreateInstance         call.Back[func(userdata unsafe.Pointer) uintptr]
	FreeInstance           call.Back[func(userdata unsafe.Pointer, instance cgo.Handle)]
	RecreateInstance       call.Back[func(userdata unsafe.Pointer, instance cgo.Handle) uintptr]
	GetVirtual             call.Back[func(userdata unsafe.Pointer, name *StringName) call.Back[func(instance cgo.Handle, args unsafe.Pointer, ret unsafe.Pointer)]]
	GetVirtualCallWithData call.Back[func(userdata unsafe.Pointer, name *StringName) unsafe.Pointer]
	CallVirtualWithData    call.Back[func(instance cgo.Handle, name *StringName, userdata unsafe.Pointer, args unsafe.Pointer, ret unsafe.Pointer)]
	GetRID                 call.Back[func(instance cgo.Handle) uint64]

	UserData unsafe.Pointer
}

type ClassMethodArgumentMetadata uint32

const (
	ArgumentMetadataNone ClassMethodArgumentMetadata = iota
	ArgumentMetadataIntIsInt8
	ArgumentMetadataIntIsInt16
	ArgumentMetadataIntIsInt32
	ArgumentMetadataIntIsInt64
	ArgumentMetadataIntIsUint8
	ArgumentMetadataIntIsUint16
	ArgumentMetadataIntIsUint32
	ArgumentMetadataIntIsUint64
	ArgumentMetadataRealIsFloat32
	ArgumentMetadataRealIsFloat64
)

type ClassMethodInfo struct {
	Name                *StringName
	MethodUserdata      unsafe.Pointer
	Call                call.Back[func(userdata, instance unsafe.Pointer, args *Variant, arg_count int64, ret *Variant, err *CallError)]
	PointerCall         call.Back[func(userdata, instance unsafe.Pointer, args unsafe.Pointer, ret unsafe.Pointer, err *CallError)]
	MethodFlags         uint32 // MethodFlags
	HasReturnValue      bool
	ReturnValueInfo     *PropertyInfo
	ReturnValueMetadata ClassMethodArgumentMetadata

	ArgumentCount     uint32
	ArgumentsInfo     *PropertyInfo
	ArgumentsMetadata *ClassMethodArgumentMetadata

	DefaultArgumentCount uint32
	DefaultArguments     *Variant
}

type CallableCustomInfo struct {
	CallableUserdata unsafe.Pointer
	Token            unsafe.Pointer
	ObjectID         InstanceID
	Call             call.Back[func(userdata, instance unsafe.Pointer, args *Variant, arg_count int64, ret *Variant, err *CallError)]
	IsValid          call.Back[func(userdata unsafe.Pointer) bool]
	Free             call.Back[func(userdata unsafe.Pointer)]
	Hash             call.Back[func(userdata unsafe.Pointer) uint32]
	Equal            call.Back[func(a unsafe.Pointer, b unsafe.Pointer) bool]
	LessThan         call.Back[func(a unsafe.Pointer, b unsafe.Pointer) bool]
	ToString         call.Back[func(userdata unsafe.Pointer, valid *bool, out *String)]
}

type ScriptInstanceInfo struct {
	Set               call.Back[func(instance unsafe.Pointer, name *StringName, value *Variant)]
	Get               call.Back[func(instance unsafe.Pointer, name *StringName, ret *Variant)]
	GetPropertyList   call.Back[func(instance unsafe.Pointer, len *uint32) *PropertyInfo]
	FreePropertyList  call.Back[func(instance unsafe.Pointer, list *PropertyInfo)]
	PropertyCanRevert call.Back[func(instance unsafe.Pointer, name *StringName) bool]
	PropertyGetRevert call.Back[func(instance unsafe.Pointer, name *StringName, ret *Variant)]

	GetOwner         call.Back[func(instance unsafe.Pointer) *Object]
	GetPropertyState call.Back[func(instance unsafe.Pointer, add func(name *StringName, value *Variant, userdata unsafe.Pointer), userdata unsafe.Pointer)]

	GetMethodList    call.Back[func(instance unsafe.Pointer, len *uint32) *MethodInfo]
	FreeMethodList   call.Back[func(instance unsafe.Pointer, list *MethodInfo)]
	GetPropertyType  call.Back[func(instance unsafe.Pointer, name *StringName, valid *bool) VariantType]
	ValidateProperty call.Back[func(instance unsafe.Pointer, name *StringName, property *PropertyInfo) bool]

	HasMethod call.Back[func(instance unsafe.Pointer, name *StringName) bool]

	Call         call.Back[func(instance unsafe.Pointer, name *StringName, args *Variant, arg_count int64, ret *Variant, err *CallError)]
	Notification call.Back[func(instance unsafe.Pointer, notification int32, reversed bool)]

	ToString call.Back[func(instance unsafe.Pointer, valid *bool, out *String)]

	RefcountIncremented call.Back[func(instance unsafe.Pointer)]
	RefcountDecremented call.Back[func(instance unsafe.Pointer)]

	GetScript call.Back[func(instance unsafe.Pointer) *Script]

	IsPlaceholder call.Back[func(instance unsafe.Pointer) bool]

	SetFallback call.Back[func(instance unsafe.Pointer, name *StringName, value *Variant)]
	GetFallback call.Back[func(instance unsafe.Pointer, name *StringName, ret *Variant)]

	GetLanguage call.Back[func(instance unsafe.Pointer) *ScriptLanguage]

	Free call.Back[func(instance unsafe.Pointer)]
}
