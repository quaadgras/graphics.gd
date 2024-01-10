package gdextension

import (
	"unsafe"

	"grow.graphics/gd"
	"runtime.link/api"
	"runtime.link/api/call"
)

type API struct {
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

	GetNativeStructSize func(name *gd.StringName) uint64 `call:"get_native_struct_size func(&char)size_t"`

	Free struct {
		StringName func() func(*gd.StringName) `call:"variant_get_ptr_destructor func(-int=21)func($void)"`
	}

	Variants struct {
		// Copy dst must be unititialized
		Copy func(dst, src *gd.Variant) `call:"variant_new_copy func(+void,$void)"`
		// Zero dst must be unititialized
		Zero          func(dst *gd.Variant)                                                                                 `call:"variant_new_nil func(+void)"`
		Free          func(*gd.Variant)                                                                                     `call:"variant_destroy func($void)"`
		Call          func(self *gd.Variant, method *gd.StringName, args []*gd.Variant, ret *gd.Variant, err CallError)     `call:"variant_call func($void,&char,&void,-int64_t=@3,+void,+void)"`
		CallStatic    func(vtype gd.VariantType, method *gd.StringName, args []*gd.Variant, ret *gd.Variant, err CallError) `call:"variant_call_static func(int,&void,&char,&void,-int64_t=@3,+void,+void)"`
		Evaluate      func(operator Operator, a, b *gd.Variant, ret *gd.Variant, ok *bool)                                  `call:"variant_evaluate func(int,&void,&void,+void,+bool)"`
		Set           func(self, key, val *gd.Variant, ok *bool)                                                            `call:"variant_set func(&void,&void,+bool)"`
		SetNamed      func(self *gd.Variant, key *gd.StringName, val *gd.Variant, ok *bool)                                 `call:"variant_set_named func(&void,&char,+bool)"`
		SetKeyed      func(self, key, val *gd.Variant, ok *bool)                                                            `call:"variant_set_keyed func(&void,&void,+bool)"`
		SetIndex      func(self *gd.Variant, index int64, val *gd.Variant, ok *bool)                                        `call:"variant_set_indexed func(&void,int64_t,+bool)"`
		Get           func(self, key *gd.Variant, ret *gd.Variant, ok *bool)                                                `call:"variant_get func(&void,&void,+void,+bool)"`
		GetNamed      func(self *gd.Variant, key *gd.StringName, ret *gd.Variant, ok *bool)                                 `call:"variant_get_named func(&void,&char,+void,+bool)"`
		GetKeyed      func(self *gd.Variant, key *gd.Variant, ret *gd.Variant, ok *bool)                                    `call:"variant_get_keyed func(&void,&void,+void,+bool)"`
		GetIndex      func(self *gd.Variant, index int64, ret *gd.Variant, ok *bool)                                        `call:"variant_get_indexed func(&void,int64_t,+void,+bool)"`
		Iterator      func(self *gd.Variant, iter *Iterator) bool                                                           `call:"variant_iter_init func(&void,+void,+bool)"`
		Next          func(self *gd.Variant, iterator Iterator, ret *gd.Variant) bool                                       `call:"variant_iter_next func(&void,&void,+void,+bool)"`
		IteratorValue func(self *gd.Variant, iterator Iterator, ret *gd.Variant) bool                                       `call:"variant_iter_get func(&void,&void,+void,+bool)"`
		Hash          func(self *gd.Variant) int64                                                                          `call:"variant_hash func(&void)int64_t"`
		RecursiveHash func(self *gd.Variant) int64                                                                          `call:"variant_recursive_hash func(&void)int64_t"`
		Compare       func(self, variant *gd.Variant) bool                                                                  `call:"variant_hash_compare func(&void,&void)bool"`
		ToBool        func(self *gd.Variant) bool                                                                           `call:"variant_booleanize func(&void)bool"`
		Duplicate     func(self *gd.Variant, ret *gd.Variant, deep bool)                                                    `call:"variant_duplicate func(&void,+void,bool)"`
		ToString      func(self *gd.Variant, out *gd.String)                                                                `call:"variant_stringify func(&void,+char)"`
		Type          func(self *gd.Variant) gd.VariantType                                                                 `call:"variant_get_type func(&void)int"`
		HasMethod     func(self *gd.Variant, method *gd.StringName) bool                                                    `call:"variant_has_method func(&void,&char)bool"`
		HasMember     func(self *gd.Variant, member *gd.StringName) bool                                                    `call:"variant_has_member func(&void,&char)bool"`
		HasKey        func(self *gd.Variant, key *gd.Variant) (bool, bool)                                                  `call:"variant_has_key func(&void,&void,+bool)bool"`
		// TypeName dst must be unititialized
		TypeName         func(self *gd.Variant, dst *gd.String)     `call:"variant_get_type_name func(&void,+char)"`
		CanConvert       func(fromType, toType gd.VariantType) bool `call:"variant_can_convert func(int,int)bool"`
		CanConvertStrict func(fromType, toType gd.VariantType) bool `call:"variant_can_convert_strict func(int,int)bool"`

		Encoder func(gd.VariantType) func(out *gd.Variant, in unsafe.Pointer) `call:"get_variant_from_type_constructor func(int)func(+void,&void)"`
		Decoder func(gd.VariantType) func(out unsafe.Pointer, in *gd.Variant) `call:"get_variant_to_type_constructor func(int)func(+void,&void)"`

		Evaulator func(Operator) func(a *gd.Variant, b *gd.Variant, ret *gd.Variant) `call:"variant_get_ptr_operator_evaluator func(int)func(&void,&void,+void)"`

		BuiltinMethodCaller func(gd.VariantType, *gd.StringName, int64) func(base *gd.Variant, args []*gd.Variant, ret *gd.Variant) `call:"variant_get_ptr_builtin_method func(int,&char,int64_t)func(&void,&void,+void,-int=@2)"`
		Constructor         func(gd.VariantType, int32) func(base unsafe.Pointer, args *unsafe.Pointer)                             `call:"variant_get_ptr_constructor func(int,int64_t)func(+void,-int=@1)"`
		Destructor          func(gd.VariantType) func(unsafe.Pointer)                                                               `call:"variant_get_ptr_destructor func(int)func($void)"`

		Construct func(t gd.VariantType, out *gd.Variant, args []*gd.Variant, err CallError) `call:"variant_construct func(int,+void,&void,-int64_t=@3,+void)"`

		Setter      func(gd.VariantType, *gd.StringName) func(*gd.Variant, unsafe.Pointer)                `call:"variant_get_ptr_setter func(int,&char)func(&void,&void)"`
		Getter      func(gd.VariantType, *gd.StringName) func(base *gd.Variant, ret unsafe.Pointer)       `call:"variant_get_ptr_getter func(int,&char)func(&void,+void)"`
		IndexSetter func(gd.VariantType) func(*gd.Variant, int64, unsafe.Pointer)                         `call:"variant_get_ptr_indexed_setter func(int)func(&void,int64_t,&void)"`
		IndexGetter func(gd.VariantType) func(val *gd.Variant, idx int64, ret unsafe.Pointer)             `call:"variant_get_ptr_indexed_getter func(int)func(&void,int64_t)&void"`
		KeySetter   func(gd.VariantType) func(*gd.Variant, *gd.Variant, unsafe.Pointer)                   `call:"variant_get_ptr_keyed_setter func(int)func(&void,&void,&void)"`
		KeyGetter   func(gd.VariantType) func(val *gd.Variant, key *gd.Variant, ret unsafe.Pointer)       `call:"variant_get_ptr_keyed_getter func(int)func(&void,&void)&void"`
		KeyChecker  func(gd.VariantType) func(*gd.Variant, *gd.Variant) uint32                            `call:"variant_get_ptr_keyed_checker func(int)func(&void,&void)uint32_t"`
		Constant    func(t gd.VariantType, name *gd.StringName, ret *gd.Variant)                          `call:"variant_get_constant_value func(int,&char,+void)"`
		Utility     func(gd.VariantType, *gd.StringName, int64) func(val *gd.Variant, args []*gd.Variant) `call:"variant_get_ptr_utility_function func(int,&char,int64_t)func(+void,&void,-int=@2)"`
	}

	Strings struct {
		New        func(dst *gd.String, s string) `call:"string_new_with_utf8_chars_and_len func(+void,&char,int64_t=@2)"`
		Get        func(*gd.String, []byte)       `call:"string_to_utf8_chars func(&void,&char,-int64_t=@2)"`
		Index      func(*gd.String, int64) rune   `call:"string_operator_index func(&void,int64_t)&char32_t"`
		Append     func(*gd.String, *gd.String)   `call:"string_operator_plus_eq_string func(&void,&void)"`
		AppendRune func(*gd.String, rune)         `call:"string_operator_plus_eq_char func(&void,&char32_t)"`
		Resize     func(*gd.String, int64)        `call:"string_resize func(&void,int64_t)"`
	}

	StringNames struct {
		New func(dst *gd.StringName, s string) `call:"string_name_new_with_utf8_chars_and_len func(+void,&char,-int64_t=@2)"`
	}

	XMLParser struct {
		OpenBuffer func(*gd.XMLParser, []byte) gd.Error `call:"xml_parser_open_buffer func(&void,&char,int64_t=@2)int64_t"`
	}

	FileAccess struct {
		StoreBuffer func(*gd.FileAccess, []byte) gd.Error `call:"file_access_store_buffer func(&void,&char,int64_t=@2)int64_t"`
		ReadBuffer  func(*gd.FileAccess, []byte) gd.Error `call:"file_access_get_buffer func(&void,&char,int64_t=@2)int64_t"`
	}

	PackedByteArray struct {
		Index func(*gd.PackedByteArray, int64) *byte `call:"packed_byte_array_operator_index func(&void,int64_t)&char"`
	}
	PackedColorArray struct {
		Index func(*gd.PackedColorArray, int64) *gd.Color `call:"packed_color_array_operator_index func(&void,int64_t)&Color"`
	}
	PackedFloat32Array struct {
		Index func(*gd.PackedFloat32Array, int64) *float32 `call:"packed_float32_array_operator_index func(&void,int64_t)&float"`
	}
	PackedFloat64Array struct {
		Index func(*gd.PackedFloat64Array, int64) *float64 `call:"packed_float64_array_operator_index func(&void,int64_t)&double"`
	}
	PackedInt32Array struct {
		Index func(*gd.PackedInt32Array, int64) *int32 `call:"packed_int32_array_operator_index func(&void,int64_t)&int32_t"`
	}
	PackedInt64Array struct {
		Index func(*gd.PackedInt64Array, int64) *int64 `call:"packed_int64_array_operator_index func(&void,int64_t)&int64_t"`
	}
	PackedStringArray struct {
		Index func(*gd.PackedStringArray, int64) *gd.String `call:"packed_string_array_operator_index func(&void,int64_t)&char"`
	}
	PackedVector2Array struct {
		Index func(*gd.PackedVector2Array, int64) *gd.Vector2 `call:"packed_vector2_array_operator_index func(&void,int64_t)&Vector2"`
	}
	PackedVector3Array struct {
		Index func(*gd.PackedVector3Array, int64) *gd.Vector3 `call:"packed_vector3_array_operator_index func(&void,int64_t)&Vector3"`
	}
	Array struct {
		Index    func(*gd.Array, int64) *gd.Variant                                               `call:"array_operator_index func(&void,int64_t)&void"`
		Set      func(*gd.Array, *gd.Array)                                                       `call:"array_ref func(&void,&void)"`
		SetTyped func(a *gd.Array, t gd.VariantType, className *gd.StringName, script *gd.Script) `call:"array_set_typed func(&void,int,&void,&void)"`
	}
	Dictionary struct {
		Index func(d *gd.Dictionary, key *gd.Variant, ret *gd.Variant) `call:"dictionary_operator_index func(&void,&void,&void)"`
	}

	Object struct {
		Call                func(method MethodBind, obj *Object, arg []*gd.Variant, ret *gd.Variant, err CallError) `call:"object_method_bind_call func(&void,&void,&void,-int64_t=@3,+void,+void)"`
		UnsafeCall          func(method MethodBind, obj *Object, args []unsafe.Pointer, ret unsafe.Pointer)         `call:"object_method_bind_ptrcall func(&void,&void,&void,&void)"`
		Destroy             func(obj *Object)                                                                       `call:"object_destroy func(&void)"`
		GetSingleton        func(name *gd.StringName) Object                                                        `call:"global_get_singleton func(&void)&void"`
		GetInstanceBinding  func(obj *Object, token unsafe.Pointer, callbacks *InstanceBindingCallbacks)            `call:"object_get_instance_binding func(&void,$void,&void)"`
		SetInstanceBinding  func(obj *Object, token, binding unsafe.Pointer, callbacks *InstanceBindingCallbacks)   `call:"object_set_instance_binding func(&void,$void,&void)"`
		FreeInstanceBinding func(obj *Object, token unsafe.Pointer)                                                 `call:"object_free_instance_binding func(&void,$void)"`
		SetInstance         func(obj *Object, name *gd.StringName, value unsafe.Pointer)                            `call:"object_set_instance func(&void,&char,&void)"`
		GetClassName        func(obj *Object, library unsafe.Pointer, out *gd.StringName)                           `call:"object_get_class_name func(&void,+void)"`
		CastTo              func(obj *Object, class ClassTag) *Object                                               `call:"object_cast_to func(&void,&char)&void"`
		GetFromID           func(id InstanceID) *Object                                                             `call:"object_get_instance_from_id func(uint64_t)&void"`
		GetID               func(obj *Object) InstanceID                                                            `call:"object_get_instance_id func(&void)uint64_t"`
		GetFromReference    func(ref unsafe.Pointer) *Object                                                        `call:"ref_get_object func(&void)&void"`
		Reference           func(ref unsafe.Pointer, obj *Object)                                                   `call:"ref_set_object func(&void,&void)"`
	}

	Scripts struct {
		Create            func(info *ScriptInstanceInfo, script unsafe.Pointer) Script                      `call:"script_instance_create2 func(&void,&char,&char)&void"`
		CreatePlaceholder func(lang *gd.ScriptLanguage, script *gd.Script, owner *Object) Script            `call:"placeholder_script_instance_create func(&void,&char,&char)&void"`
		UpdatePlaceholder func(script Script, properties *gd.ArrayOf[gd.Dictionary], values *gd.Dictionary) `call:"placeholder_script_instance_update func(&void)"`
		Get               func(obj *Object, lang *gd.ScriptLanguage) Script                                 `call:"object_get_script_instance func(&void)&void"`
	}

	Callables struct {
		Create      func(callable *gd.Callable, info *CallableCustomInfo)            `call:"callable_custom_create func(+void,&void)"`
		GetUserData func(callable *gd.Callable, token unsafe.Pointer) unsafe.Pointer `call:"callable_custom_get_userdata func(&void,&void)$void"`
	}

	ClassDB struct {
		CreateObject  func(class *gd.StringName) Object                                        `call:"classdb_construct_object func(&void)$void"`
		GetMethodBind func(class *gd.StringName, method *gd.StringName, hash int64) MethodBind `call:"classdb_get_method_bind func(&void,&void,int)void"`
		GetClassTag   func(class *gd.StringName) ClassTag                                      `call:"classdb_get_class_tag func(&void)void"`

		RegisterClass                 func(library unsafe.Pointer, name, extends *gd.StringName, info *ClassCreationInfo)                                `call:"classdb_register_extension_class2 func(&void,&char,&void)"`
		RegisterClassMethod           func(library unsafe.Pointer, class *gd.StringName, info *ClassMethodInfo)                                          `call:"classdb_register_extension_class_method func(&void,&char,&void)"`
		RegisterClassIntegerConstant  func(library unsafe.Pointer, class, enum, name *gd.StringName, value int64, bitfield bool)                         `call:"classdb_register_extension_class_integer_constant func(&void,&void,&void,&void,&void,int64_t)"`
		RegisterClassProperty         func(library unsafe.Pointer, class *gd.StringName, info *PropertyInfo, getter, setter *gd.StringName)              `call:"classdb_register_extension_class_property func(&void,&void,&void,&void,&void)"`
		RegisterClassPropertyIndexed  func(library unsafe.Pointer, class *gd.StringName, info *PropertyInfo, getter, setter *gd.StringName, index int64) `call:"classdb_register_extension_class_property_indexed func(&void,&void,&void,&void,&void,int64_t)"`
		RegisterClassPropertyGroup    func(library unsafe.Pointer, class *gd.StringName, group, prefix *gd.String)                                       `call:"classdb_register_extension_class_property_group func(&void,&void,&void,&void)"`
		RegisterClassPropertySubGroup func(library unsafe.Pointer, class *gd.StringName, subGroup, prefix *gd.String)                                    `call:"classdb_register_extension_class_property_subgroup func(&void,&void,&void,&void,&void)"`
		RegisterClassSignal           func(library unsafe.Pointer, class, signal *gd.StringName, args []PropertyInfo)                                    `call:"classdb_register_extension_class_signal func(&void,&void,&void,&void,-int64_t=@4)"`
		UnregisterClass               func(library unsafe.Pointer, class *gd.StringName)                                                                 `call:"classdb_unregister_extension_class func(&void,&void)"`
		GetLibraryPath                func(library unsafe.Pointer, out *gd.String)                                                                       `call:"get_library_path func(&void,&void)"`
	}

	EditorPlugins struct {
		Add    func(plugin *gd.StringName) `call:"editor_add_plugin func(&void)"`
		Remove func(plugin *gd.StringName) `call:"editor_remove_plugin func(&void)"`
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

type Script uintptr

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
	MinimumInitializationLevel gd.ExtensionInitializationLevel
	Userdata                   T
	Initialize                 call.Back[func(userdata T, level gd.ExtensionInitializationLevel)]
	Deinitialize               call.Back[func(userdata T, level gd.ExtensionInitializationLevel)]
}

type InstanceBindingCallbacks struct {
	Create    call.Back[func(token, instance unsafe.Pointer) unsafe.Pointer]
	Free      call.Back[func(token, instance, binding unsafe.Pointer)]
	Reference call.Back[func(token, binding unsafe.Pointer, reference bool) bool]
}

type PropertyInfo struct {
	Type       gd.VariantType
	Name       gd.StringName
	ClassName  gd.StringName
	Hint       uint32 // gd.PropertyHint
	HintString gd.String
	Usage      uint32 // gd.PropertyUsageFlags
}

type MethodInfo struct {
	Name                 gd.StringName
	ReturnValue          PropertyInfo
	Flags                ClassMethodFlags
	ID                   int32
	ArgumentCount        uint32
	Arguments            *PropertyInfo
	DefaultArgumentCount uint32
	DefaultArguments     *gd.Variant
}

type ClassCreationInfo struct {
	IsVirtual  bool
	IsAbstract bool
	IsExposed  bool

	Set call.Back[func(instance unsafe.Pointer, name *gd.StringName, value *gd.Variant)]
	Get call.Back[func(instance unsafe.Pointer, name *gd.StringName, ret *gd.Variant)]

	GetPropertyList   call.Back[func(instance unsafe.Pointer, len *uint32) *PropertyInfo]
	FreePropertyList  call.Back[func(instance unsafe.Pointer, list *PropertyInfo)]
	PropertyCanRevert call.Back[func(instance unsafe.Pointer, name *gd.StringName) bool]
	PropertyGetRevert call.Back[func(instance unsafe.Pointer, name *gd.StringName, ret *gd.Variant)]
	ValidateProperty  call.Back[func(instance unsafe.Pointer, name *gd.StringName, property *PropertyInfo) bool]
	Notification      call.Back[func(instance unsafe.Pointer, notification int32, reversed bool)]
	ToString          call.Back[func(instance unsafe.Pointer, valid *bool, out *gd.String)]
	Reference         call.Back[func(instance unsafe.Pointer)]
	Unreference       call.Back[func(instance unsafe.Pointer)]

	// Mandatory
	CreateInstance   call.Back[func(userdata unsafe.Pointer) *gd.Object]
	FreeInstance     call.Back[func(userdata unsafe.Pointer, instance *gd.Object)]
	RecreateInstance call.Back[func(userdata unsafe.Pointer, instance *gd.Object) *gd.Object]
	GetVirtual       call.Back[func(userdata unsafe.Pointer, name *gd.StringName) call.Back[func(instance unsafe.Pointer, name *gd.StringName, userdata unsafe.Pointer, args []unsafe.Pointer, ret unsafe.Pointer)]]
	GetRID           call.Back[func(instance unsafe.Pointer) uint64]

	UserData unsafe.Pointer
}

type ClassMethodFlags uint32

const (
	MethodFlagNormal ClassMethodFlags = 1 << iota
	MethodFlagEditor
	MethodFlagConst
	MethodFlagVirtual
	MethodFlagVararg
	MethodFlagStatic
	MethodFlagsDefault ClassMethodFlags = MethodFlagNormal
)

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
	Name                *gd.StringName
	MethodUserdata      unsafe.Pointer
	Call                call.Back[func(userdata, instance unsafe.Pointer, args *gd.Variant, arg_count int64, ret *gd.Variant, err *CallError)]
	PointerCall         call.Back[func(userdata, instance unsafe.Pointer, args unsafe.Pointer, ret unsafe.Pointer, err *CallError)]
	MethodFlags         ClassMethodFlags
	HasReturnValue      bool
	ReturnValueInfo     *PropertyInfo
	ReturnValueMetadata ClassMethodArgumentMetadata

	ArgumentCount     uint32
	ArgumentsInfo     *PropertyInfo
	ArgumentsMetadata *ClassMethodArgumentMetadata

	DefaultArgumentCount uint32
	DefaultArguments     *gd.Variant
}

type CallableCustomInfo struct {
	CallableUserdata unsafe.Pointer
	Token            unsafe.Pointer
	ObjectID         InstanceID
	Call             call.Back[func(userdata, instance unsafe.Pointer, args *gd.Variant, arg_count int64, ret *gd.Variant, err *CallError)]
	IsValid          call.Back[func(userdata unsafe.Pointer) bool]
	Free             call.Back[func(userdata unsafe.Pointer)]
	Hash             call.Back[func(userdata unsafe.Pointer) uint32]
	Equal            call.Back[func(a unsafe.Pointer, b unsafe.Pointer) bool]
	LessThan         call.Back[func(a unsafe.Pointer, b unsafe.Pointer) bool]
	ToString         call.Back[func(userdata unsafe.Pointer, valid *bool, out *gd.String)]
}

type ScriptInstanceInfo struct {
	Set               call.Back[func(instance unsafe.Pointer, name *gd.StringName, value *gd.Variant)]
	Get               call.Back[func(instance unsafe.Pointer, name *gd.StringName, ret *gd.Variant)]
	GetPropertyList   call.Back[func(instance unsafe.Pointer, len *uint32) *PropertyInfo]
	FreePropertyList  call.Back[func(instance unsafe.Pointer, list *PropertyInfo)]
	PropertyCanRevert call.Back[func(instance unsafe.Pointer, name *gd.StringName) bool]
	PropertyGetRevert call.Back[func(instance unsafe.Pointer, name *gd.StringName, ret *gd.Variant)]

	GetOwner         call.Back[func(instance unsafe.Pointer) *gd.Object]
	GetPropertyState call.Back[func(instance unsafe.Pointer, add func(name *gd.StringName, value *gd.Variant, userdata unsafe.Pointer), userdata unsafe.Pointer)]

	GetMethodList    call.Back[func(instance unsafe.Pointer, len *uint32) *MethodInfo]
	FreeMethodList   call.Back[func(instance unsafe.Pointer, list *MethodInfo)]
	GetPropertyType  call.Back[func(instance unsafe.Pointer, name *gd.StringName, valid *bool) gd.VariantType]
	ValidateProperty call.Back[func(instance unsafe.Pointer, name *gd.StringName, property *PropertyInfo) bool]

	HasMethod call.Back[func(instance unsafe.Pointer, name *gd.StringName) bool]

	Call         call.Back[func(instance unsafe.Pointer, name *gd.StringName, args *gd.Variant, arg_count int64, ret *gd.Variant, err *CallError)]
	Notification call.Back[func(instance unsafe.Pointer, notification int32, reversed bool)]

	ToString call.Back[func(instance unsafe.Pointer, valid *bool, out *gd.String)]

	RefcountIncremented call.Back[func(instance unsafe.Pointer)]
	RefcountDecremented call.Back[func(instance unsafe.Pointer)]

	GetScript call.Back[func(instance unsafe.Pointer) *gd.Script]

	IsPlaceholder call.Back[func(instance unsafe.Pointer) bool]

	SetFallback call.Back[func(instance unsafe.Pointer, name *gd.StringName, value *gd.Variant)]
	GetFallback call.Back[func(instance unsafe.Pointer, name *gd.StringName, ret *gd.Variant)]

	GetLanguage call.Back[func(instance unsafe.Pointer) *gd.ScriptLanguage]

	Free call.Back[func(instance unsafe.Pointer)]
}

type Object uintptr
