//go:build !generate

package gd

import (
	"runtime/cgo"
	"unsafe"

	"runtime.link/api"
	"runtime.link/api/call"
)

type API struct {
	api.Specification

	GetGodotVersion     func() Version
	GetNativeStructSize func(StringName) uint64

	Memory struct {
		Allocate   func(uintptr) unsafe.Pointer
		Reallocate func(unsafe.Pointer, uintptr) unsafe.Pointer
		Free       func(unsafe.Pointer)
	}

	PrintError              func(code, function, file string, line int32, notifyEditor bool)
	PrintErrorMessage       func(code, message, function, file string, line int32, notifyEditor bool)
	PrintWarning            func(code, function, file string, line int32, notifyEditor bool)
	PrintWarningMessage     func(code, message, function, file string, line int32, notifyEditor bool)
	PrintScriptError        func(code, function, file string, line int32, notifyEditor bool)
	PrintScriptErrorMessage func(code, message, function, file string, line int32, notifyEditor bool)

	Variants struct {
		NewCopy                  func(ctx Context, src Variant) Variant
		NewNil                   func(ctx Context) Variant
		Destroy                  func(self Variant)
		Call                     func(ctx Context, self Variant, method StringName, args ...Variant) (Variant, error)
		CallStatic               func(ctx Context, vtype VariantType, method StringName, args ...Variant) (Variant, error)
		Evaluate                 func(ctx Context, operator Operator, a, b Variant) (ret Variant, ok bool)
		Set                      func(self, key, val Variant) bool
		SetNamed                 func(self Variant, key StringName, val Variant) bool
		SetKeyed                 func(self, key, val Variant) bool
		SetIndexed               func(self Variant, index int64, val Variant) (ok, oob bool)
		Get                      func(ctx Context, self, key Variant) (Variant, bool)
		GetNamed                 func(ctx Context, self Variant, key StringName) (Variant, bool)
		GetKeyed                 func(ctx Context, self, key Variant) (Variant, bool)
		GetIndexed               func(ctx Context, self Variant, index int64) (val Variant, ok, oob bool)
		IteratorInitialize       func(ctx Context, self Variant) (Variant, bool)
		IteratorNext             func(self Variant, iterator Variant) bool
		IteratorGet              func(ctx Context, self Variant, iterator Variant) (Variant, bool)
		Hash                     func(self Variant) int64
		RecursiveHash            func(self Variant, count Int) int64
		HashCompare              func(self, variant Variant) bool
		Booleanize               func(self Variant) bool
		Duplicate                func(self Variant, deep bool) Variant
		Stringify                func(self Variant) String
		GetType                  func(self Variant) VariantType
		HasMethod                func(self Variant, method StringName) bool
		HasMember                func(self Variant, member StringName) bool
		HasKey                   func(self Variant, key Variant) (hasKey, valid bool)
		GetTypeName              func(ctx Context, self VariantType) String
		CanConvert               func(self Variant, to VariantType) bool
		CanConvertStrict         func(self Variant, to VariantType) bool
		FromTypeConstructor      func(VariantType) func(ret call.Ptr[[3]uintptr], arg call.Any)
		ToTypeConstructor        func(VariantType) func(ret call.Any, arg call.Ptr[[3]uintptr])
		PointerOperatorEvaluator func(op Operator, a, b VariantType) func(a, b, ret call.Any)
		GetPointerBuiltinMethod  func(VariantType, StringName, Int) func(base call.Any, args call.Args, ret call.Any, c int32)

		BuiltinMethodCaller func(VariantType, StringNamePtr, int64) func(base, args CallFrameArgs, ret CallFrameBack, c int32) `call:"variant_get_ptr_builtin_method func(int,&void,int64_t)func(&void,&void,+void,int)"`

		Constructor func(VariantType, int32) func(base CallFrameBack, args CallFrameArgs) `call:"variant_get_ptr_constructor func(int,int32_t)func(+void,&void)"`
		Destructor  func(VariantType) func(CallFrameArgs)                                 `call:"variant_get_ptr_destructor func(int)func($void)"`

		Construct func(t VariantType, out *Variant, args []*Variant, err CallError) `call:"variant_construct func(int,+void,&void,-int64_t=@3,+void)"`

		Setter      func(VariantType, StringNamePtr) func(*Variant, unsafe.Pointer)                 `call:"variant_get_ptr_setter func(int,&void)func(&void,&void)"`
		Getter      func(VariantType, StringNamePtr) func(base *Variant, ret unsafe.Pointer)        `call:"variant_get_ptr_getter func(int,&void)func(&void,+void)"`
		IndexSetter func(VariantType) func(*Variant, int64, unsafe.Pointer)                         `call:"variant_get_ptr_indexed_setter func(int)func(&void,int64_t,&void)"`
		IndexGetter func(VariantType) func(val *Variant, idx int64, ret unsafe.Pointer)             `call:"variant_get_ptr_indexed_getter func(int)func(&void,int64_t)&void"`
		KeySetter   func(VariantType) func(*Variant, *Variant, unsafe.Pointer)                      `call:"variant_get_ptr_keyed_setter func(int)func(&void,&void,&void)"`
		KeyGetter   func(VariantType) func(val *Variant, key *Variant, ret unsafe.Pointer)          `call:"variant_get_ptr_keyed_getter func(int)func(&void,&void)&void"`
		KeyChecker  func(VariantType) func(*Variant, *Variant) uint32                               `call:"variant_get_ptr_keyed_checker func(int)func(&void,&void)uint32_t"`
		Constant    func(t VariantType, name StringNamePtr, ret *Variant)                           `call:"variant_get_constant_value func(int,&void,+void)"`
		Utility     func(StringNamePtr, int64) func(ret CallFrameBack, args CallFrameArgs, c int32) `call:"variant_get_ptr_utility_function func(&void,int64_t)func(+void,&void,int)"`
	}

	Strings struct {
		New        func(dst CallFrameBack, s string) `call:"string_new_with_utf8_chars_and_len func(+void,&char,-int64_t=@2)"`
		Get        func(CallFrameArgs, []byte)       `call:"string_to_utf8_chars func(&void,&char,-int64_t=@2)"`
		Index      func(StringPtr, int64) rune       `call:"string_operator_index func(&void,int64_t)&char32_t"`
		Append     func(StringPtr, *String)          `call:"string_operator_plus_eq_string func(&void,&void)"`
		AppendRune func(StringPtr, rune)             `call:"string_operator_plus_eq_char func(&void,&char32_t)"`
		Resize     func(StringPtr, int64)            `call:"string_resize func(&void,int64_t)"`
	}

	StringNames struct {
		New func(dst CallFrameBack, s string) `call:"string_name_new_with_utf8_chars_and_len func(+void,&char,-int64_t=@2)"`
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
		Index func(*PackedStringArray, int64) *uintptr `call:"packed_string_array_operator_index func(&void,int64_t)$void"`
	}
	PackedVector2Array struct {
		Index func(*PackedVector2Array, int64) *Vector2 `call:"packed_vector2_array_operator_index func(&void,int64_t)&Vector2"`
	}
	PackedVector3Array struct {
		Index func(*PackedVector3Array, int64) *Vector3 `call:"packed_vector3_array_operator_index func(&void,int64_t)&Vector3"`
	}
	Array struct {
		Index    func(*Array, int64) *Variant                                           `call:"array_operator_index func(&void,int64_t)&void"`
		Set      func(*Array, *Array)                                                   `call:"array_ref func(&void,&void)"`
		SetTyped func(a *Array, t VariantType, className StringNamePtr, script *Script) `call:"array_set_typed func(&void,int,&void,&void)"`
	}
	Dictionary struct {
		Index func(d *Dictionary, key *Variant, ret *Variant) `call:"dictionary_operator_index func(&void,&void,&void)"`
	}

	Object struct {
		MethodBindPointerCall func(method MethodBind, obj Object, arg call.Args, ret call.Any)

		Call                func(method MethodBind, obj uintptr, arg []*Variant, ret *Variant, err CallError)     `call:"object_method_bind_call func(&void,&void,&void,-int64_t=@3,+void,+void)"`
		UnsafeCall          func(method MethodBind, obj uintptr, args CallFrameArgs, ret CallFrameBack)           `call:"object_method_bind_ptrcall func(&void,&void,&void,&void)"`
		Destroy             func(obj uintptr)                                                                     `call:"object_destroy func(&void)"`
		GetSingleton        func(name StringNamePtr) uintptr                                                      `call:"global_get_singleton func(&void)&void"`
		GetInstanceBinding  func(obj uintptr, token unsafe.Pointer, callbacks *InstanceBindingCallbacks)          `call:"object_get_instance_binding func(&void,$void,&void)"`
		SetInstanceBinding  func(obj uintptr, token, binding unsafe.Pointer, callbacks *InstanceBindingCallbacks) `call:"object_set_instance_binding func(&void,$void,&void)"`
		FreeInstanceBinding func(obj uintptr, token unsafe.Pointer)                                               `call:"object_free_instance_binding func(&void,$void)"`
		SetInstance         func(obj uintptr, name StringNamePtr, value cgo.Handle)                               `call:"object_set_instance func(&void,&void,&void)"`
		GetClassName        func(obj uintptr, library ExtensionToken, out StringNamePtr)                          `call:"object_get_class_name func(&void,+void)"`
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
		ConstructObject func(Context, StringName) Object
		GetMethodBind   func(class, method StringNamePtr, hash int64) MethodBind `call:"classdb_get_method_bind func(&void,&void,int64_t)void"`
		GetClassTag     func(class StringNamePtr) ClassTag                       `call:"classdb_get_class_tag func(&void)void"`

		RegisterClass                 func(library ExtensionToken, name, extends StringNamePtr, info *ClassCreationInfo)                               `call:"classdb_register_extension_class2 func(&void,&void,&void,&void)"`
		RegisterClassMethod           func(library ExtensionToken, class StringNamePtr, info *ClassMethodInfo)                                         `call:"classdb_register_extension_class_method func(&void,&void,&void)"`
		RegisterClassIntegerConstant  func(library ExtensionToken, class, enum, name StringNamePtr, value int64, bitfield bool)                        `call:"classdb_register_extension_class_integer_constant func(&void,&void,&void,&void,&void,int64_t)"`
		RegisterClassProperty         func(library ExtensionToken, class StringNamePtr, info *PropertyInfo, getter, setter StringNamePtr)              `call:"classdb_register_extension_class_property func(&void,&void,&void,&void,&void)"`
		RegisterClassPropertyIndexed  func(library ExtensionToken, class StringNamePtr, info *PropertyInfo, getter, setter StringNamePtr, index int64) `call:"classdb_register_extension_class_property_indexed func(&void,&void,&void,&void,&void,int64_t)"`
		RegisterClassPropertyGroup    func(library ExtensionToken, class CallFrameArgs, group, prefix CallFrameArgs)                                   `call:"classdb_register_extension_class_property_group func(&void,&void,&void,&void)"`
		RegisterClassPropertySubGroup func(library ExtensionToken, class CallFrameArgs, subGroup, prefix CallFrameArgs)                                `call:"classdb_register_extension_class_property_subgroup func(&void,&void,&void,&void,&void)"`
		RegisterClassSignal           func(library ExtensionToken, class, signal StringNamePtr, args []PropertyInfo)                                   `call:"classdb_register_extension_class_signal func(&void,&void,&void,&void,-int64_t=@4)"`
		UnregisterClass               func(library ExtensionToken, class StringNamePtr)                                                                `call:"classdb_unregister_extension_class func(&void,&void)"`
		GetLibraryPath                func(library ExtensionToken, out CallFrameBack)                                                                  `call:"get_library_path func(&void,&void)"`
	}

	EditorPlugins struct {
		Add    func(plugin StringNamePtr) `call:"editor_add_plugin func(&void)"`
		Remove func(plugin StringNamePtr) `call:"editor_remove_plugin func(&void)"`
	}

	ExtensionToken

	cache
}

type (
	StringPtr     *uintptr
	StringNamePtr *uintptr
	VariantPtr    *[3]uintptr
)

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

func (err CallError) Error() string {
	switch err.ErrorType {
	case ErrInvalidMethod:
		return "invalid method"
	case ErrInvalidArgument:
		return "invalid argument"
	case ErrTooManyArguments:
		return "too many arguments"
	case ErrTooFewArguments:
		return "too few arguments"
	case ErrInstanceIsNil:
		return "instance is nil"
	case ErrMethodNotConst:
		return "method not const"
	default:
		return "unknown error"
	}
}

type Version struct {
	Major uint32
	Minor uint32
	Patch uint32
	Value string
}

func (v Version) String() string {
	return v.Value
}

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
	MinimumInitializationLevel GDExtensionInitializationLevel
	Userdata                   T
	Initialize                 call.Back[func(userdata T, level GDExtensionInitializationLevel)]
	Deinitialize               call.Back[func(userdata T, level GDExtensionInitializationLevel)]
}

type InstanceBindingCallbacks struct {
	Create    call.Back[func(token, instance unsafe.Pointer) unsafe.Pointer]
	Free      call.Back[func(token, instance, binding unsafe.Pointer)]
	Reference call.Back[func(token, binding unsafe.Pointer, reference bool) bool]
}

type PropertyInfo struct {
	Type       uint32 // VariantType
	Name       StringNamePtr
	ClassName  StringNamePtr
	Hint       uint32 // PropertyHint
	HintString StringPtr
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

	Set call.Back[func(instance cgo.Handle, name StringNamePtr, value VariantPtr) bool]
	Get call.Back[func(instance cgo.Handle, name StringNamePtr, ret VariantPtr) bool]

	GetPropertyList   call.Back[func(instance cgo.Handle, len *uint32) *PropertyInfo]
	FreePropertyList  call.Back[func(instance cgo.Handle, list *PropertyInfo)]
	PropertyCanRevert call.Back[func(instance cgo.Handle, name StringNamePtr) bool]
	PropertyGetRevert call.Back[func(instance cgo.Handle, name StringNamePtr, ret VariantPtr)]
	ValidateProperty  call.Back[func(instance cgo.Handle, name StringNamePtr, property *PropertyInfo) bool]
	Notification      call.Back[func(instance cgo.Handle, notification int32, reversed bool)]
	ToString          call.Back[func(instance cgo.Handle, valid *bool, out *String)]
	Reference         call.Back[func(instance cgo.Handle)]
	Unreference       call.Back[func(instance cgo.Handle)]

	// Mandatory
	CreateInstance         call.Back[func(userdata unsafe.Pointer) uintptr]
	FreeInstance           call.Back[func(userdata unsafe.Pointer, instance cgo.Handle)]
	RecreateInstance       call.Back[func(userdata unsafe.Pointer, instance cgo.Handle) uintptr]
	GetVirtual             call.Back[func(userdata unsafe.Pointer, name StringNamePtr) ExtensionClassCallVirtualFunc]
	GetVirtualCallWithData call.Back[func(userdata unsafe.Pointer, name StringNamePtr) unsafe.Pointer]
	CallVirtualWithData    call.Back[func(instance cgo.Handle, name StringNamePtr, userdata, args, ret unsafe.Pointer)]
	GetRID                 call.Back[func(instance cgo.Handle) uint64]

	UserData unsafe.Pointer
}

type ExtensionClassCallVirtualFunc = call.Back[func(cgo.Handle, UnsafeArgs, UnsafeBack)]

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
	Name                StringNamePtr
	MethodUserdata      unsafe.Pointer
	Call                call.Back[func(userdata unsafe.Pointer, instance cgo.Handle, args *Variant, arg_count int64, ret *Variant, err *CallError)]
	PointerCall         call.Back[func(userdata unsafe.Pointer, instance cgo.Handle, args unsafe.Pointer, ret unsafe.Pointer, err *CallError)]
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
