//go:build !generate

package gd

import (
	"runtime/cgo"
	"unsafe"

	"runtime.link/api"
	"runtime.link/api/call"
)

// API specification for Godot's GDExtension.
type API struct {
	api.Specification

	GetGodotVersion     func() Version
	GetNativeStructSize func(StringName) uintptr

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
		NewCopy                   func(ctx Context, src Variant) Variant
		NewNil                    func(ctx Context) Variant
		Destroy                   func(self Variant)
		Call                      func(ctx Context, self Variant, method StringName, args ...Variant) (Variant, error)
		CallStatic                func(ctx Context, vtype VariantType, method StringName, args ...Variant) (Variant, error)
		Evaluate                  func(ctx Context, operator Operator, a, b Variant) (ret Variant, ok bool)
		Set                       func(self, key, val Variant) bool
		SetNamed                  func(self Variant, key StringName, val Variant) bool
		SetKeyed                  func(self, key, val Variant) bool
		SetIndexed                func(self Variant, index int64, val Variant) (ok, oob bool)
		Get                       func(ctx Context, self, key Variant) (Variant, bool)
		GetNamed                  func(ctx Context, self Variant, key StringName) (Variant, bool)
		GetKeyed                  func(ctx Context, self, key Variant) (Variant, bool)
		GetIndexed                func(ctx Context, self Variant, index int64) (val Variant, ok, oob bool)
		IteratorInitialize        func(ctx Context, self Variant) (Variant, bool)
		IteratorNext              func(self Variant, iterator Variant) bool
		IteratorGet               func(ctx Context, self Variant, iterator Variant) (Variant, bool)
		Hash                      func(self Variant) int64
		RecursiveHash             func(self Variant, count Int) int64
		HashCompare               func(self, variant Variant) bool
		Booleanize                func(self Variant) bool
		Duplicate                 func(self Variant, deep bool) Variant
		Stringify                 func(self Variant) String
		GetType                   func(self Variant) VariantType
		HasMethod                 func(self Variant, method StringName) bool
		HasMember                 func(self Variant, member StringName) bool
		HasKey                    func(self Variant, key Variant) (hasKey, valid bool)
		GetTypeName               func(ctx Context, self VariantType) String
		CanConvert                func(self Variant, to VariantType) bool
		CanConvertStrict          func(self Variant, to VariantType) bool
		FromTypeConstructor       func(VariantType) func(ret call.Ptr[[3]uintptr], arg call.Any)
		ToTypeConstructor         func(VariantType) func(ret call.Any, arg call.Ptr[[3]uintptr])
		PointerOperatorEvaluator  func(op Operator, a, b VariantType) func(a, b, ret call.Any)
		GetPointerBuiltinMethod   func(VariantType, StringName, Int) func(base call.Any, args call.Args, ret call.Any, c int32)
		GetPointerConstructor     func(vtype VariantType, index int32) func(base call.Any, args call.Args)
		GetPointerDestructor      func(VariantType) func(base call.Any)
		Construct                 func(ctx Context, t VariantType, args ...Variant) (Variant, error)
		GetPointerSetter          func(VariantType, StringName) func(base, arg call.Any)
		GetPointerGetter          func(VariantType, StringName) func(base, ret call.Any)
		GetPointerIndexedSetter   func(VariantType) func(base call.Any, index Int, arg call.Any)
		GetPointerIndexedGetter   func(VariantType) func(base call.Any, index Int, ret call.Any)
		GetPointerKeyedSetter     func(VariantType) func(base call.Any, key call.Any, arg call.Any)
		GetPointerKeyedGetter     func(VariantType) func(base call.Any, key call.Any, ret call.Any)
		GetPointerKeyedChecker    func(VariantType) func(base call.Any, key call.Any) uint32
		GetConstantValue          func(ctx Context, t VariantType, name StringName) Variant
		GetPointerUtilityFunction func(name StringName, hash Int) func(ret call.Any, args call.Args, c int32)
	}
	Strings struct {
		New        func(Context, string) String
		Get        func(String) string
		SetIndex   func(String, Int, rune)
		Index      func(String, Int) rune
		Append     func(String, String)
		AppendRune func(String, rune)
		Resize     func(String, Int)
	}
	StringNames struct {
		New func(Context, string) StringName
	}
	XMLParser struct {
		OpenBuffer func(XMLParser, []byte) error
	}
	FileAccess struct {
		StoreBuffer func(FileAccess, []byte)
		GetBuffer   func(FileAccess, []byte) int
	}
	PackedByteArray    PackedFunctionsFor[PackedByteArray, byte]
	PackedColorArray   PackedFunctionsFor[PackedColorArray, Color]
	PackedFloat32Array PackedFunctionsFor[PackedFloat32Array, float32]
	PackedFloat64Array PackedFunctionsFor[PackedFloat64Array, float64]
	PackedInt32Array   PackedFunctionsFor[PackedInt32Array, int32]
	PackedInt64Array   PackedFunctionsFor[PackedInt64Array, int64]
	PackedStringArray  struct {
		Index       func(Context, PackedStringArray, Int) String
		SetIndex    func(PackedStringArray, Int, String)
		CopyAsSlice func(Context, PackedStringArray) []String
	}
	PackedVector2Array PackedFunctionsFor[PackedVector2Array, Vector2]
	PackedVector3Array PackedFunctionsFor[PackedVector3Array, Vector3]
	Array              struct {
		Index    func(Context, Array, Int) Variant
		Set      func(self, from Array)
		SetIndex func(Array, Int, Variant)
		SetTyped func(self Array, t VariantType, className StringName, script Script)
	}
	Dictionary struct {
		Index    func(ctx Context, dict Dictionary, key Variant) Variant
		SetIndex func(dict Dictionary, key, val Variant)
	}
	Object struct {
		MethodBindCall        func(ctx Context, method MethodBind, obj Object, arg ...Variant) (Variant, error)
		MethodBindPointerCall func(method MethodBind, obj Object, arg call.Args, ret call.Any)
		Destroy               func(Object)
		GetSingleton          func(ctx Context, name StringName) Object
		GetInstanceBinding    func(Object, ExtensionToken, InstanceBindingType) any
		SetInstanceBinding    func(Object, ExtensionToken, any, InstanceBindingType)
		FreeInstanceBinding   func(Object, ExtensionToken)
		SetInstance           func(Object, StringName, any)
		GetClassName          func(Context, Object, ExtensionToken) String
		CastTo                func(Context, Object, ClassTag) Object
		GetInstanceID         func(Object) ObjectID
	}
	RefCounted struct {
		GetObject func(Context, RefCounted) Object
		SetObject func(RefCounted, Object)
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

type Packed interface {
	PackedByteArray | PackedInt32Array | PackedInt64Array | PackedFloat32Array |
		PackedFloat64Array | PackedStringArray | PackedVector2Array | PackedVector3Array |
		PackedColorArray

	Pointer() [2]uintptr
	Size() Int
}

type PackedFunctionsFor[T Packed, V any] struct {
	Index       func(T, Int) V
	SetIndex    func(T, Int, V)
	CopyAsSlice func(T) []V
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

type InstanceBindingType unsafe.Pointer

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
