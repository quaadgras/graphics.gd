//go:build !generate

package gd

import (
	"unsafe"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"

	"runtime.link/api"
)

type VariantType int

func (t VariantType) String() string {
	return Global.Variants.GetTypeName(t).String()
}

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
		NewCopy                   func(src Variant) Variant
		NewNil                    func() Variant
		Destroy                   func(self Variant)
		Call                      func(self Variant, method StringName, args ...Variant) (Variant, error)
		CallStatic                func(vtype VariantType, method StringName, args ...Variant) (Variant, error)
		Evaluate                  func(operator Operator, a, b Variant) (ret Variant, ok bool)
		Set                       func(self, key, val Variant) bool
		SetNamed                  func(self Variant, key StringName, val Variant) bool
		SetKeyed                  func(self, key, val Variant) bool
		SetIndexed                func(self Variant, index Int, val Variant) (ok, oob bool)
		Get                       func(self, key Variant) (Variant, bool)
		GetNamed                  func(self Variant, key StringName) (Variant, bool)
		GetKeyed                  func(self, key Variant) (Variant, bool)
		GetIndexed                func(self Variant, index Int) (val Variant, ok, oob bool)
		IteratorInitialize        func(self Variant) (Variant, bool)
		IteratorNext              func(self Variant, iterator Variant) bool
		IteratorGet               func(self Variant, iterator Variant) (Variant, bool)
		Hash                      func(self Variant) Int
		RecursiveHash             func(self Variant, count Int) Int
		HashCompare               func(self, variant Variant) bool
		Booleanize                func(self Variant) bool
		Duplicate                 func(self Variant, deep bool) Variant
		Stringify                 func(self Variant) String
		GetType                   func(self Variant) VariantType
		HasMethod                 func(self Variant, method StringName) bool
		HasMember                 func(self Variant, member StringName) bool
		HasKey                    func(self Variant, key Variant) (hasKey, valid bool)
		GetTypeName               func(self VariantType) String
		CanConvert                func(self Variant, to VariantType) bool
		CanConvertStrict          func(self Variant, to VariantType) bool
		FromTypeConstructor       func(VariantType) func(ret callframe.Ptr[[3]uintptr], arg uintptr)
		ToTypeConstructor         func(VariantType) func(ret uintptr, arg callframe.Ptr[[3]uintptr])
		PointerOperatorEvaluator  func(op Operator, a, b VariantType) func(a, b, ret uintptr)
		GetPointerBuiltinMethod   func(VariantType, StringName, Int) func(base uintptr, args callframe.Args, ret uintptr, c int32)
		GetPointerConstructor     func(vtype VariantType, index int32) func(base uintptr, args callframe.Args)
		GetPointerDestructor      func(VariantType) func(base uintptr)
		Construct                 func(t VariantType, args ...Variant) (Variant, error)
		GetPointerSetter          func(VariantType, StringName) func(base, arg uintptr)
		GetPointerGetter          func(VariantType, StringName) func(base, ret uintptr)
		GetPointerIndexedSetter   func(VariantType) func(base uintptr, index Int, arg uintptr)
		GetPointerIndexedGetter   func(VariantType) func(base uintptr, index Int, ret uintptr)
		GetPointerKeyedSetter     func(VariantType) func(base uintptr, key uintptr, arg uintptr)
		GetPointerKeyedGetter     func(VariantType) func(base uintptr, key uintptr, ret uintptr)
		GetPointerKeyedChecker    func(VariantType) func(base uintptr, key uintptr) uint32
		GetConstantValue          func(t VariantType, name StringName) Variant
		GetPointerUtilityFunction func(name StringName, hash Int) func(ret uintptr, args callframe.Args, c int32)
	}
	Strings struct {
		New        func(string) String
		Get        func(String) string
		SetIndex   func(String, Int, rune)
		Index      func(String, Int) rune
		Append     func(String, String) String
		AppendRune func(String, rune)
		Resize     func(String, Int)

		UnsafePointer func(String) unsafe.Pointer
	}
	StringNames struct {
		New func(string) StringName
	}
	XMLParser struct {
		OpenBuffer func(Object, []byte) error
	}
	FileAccess struct {
		StoreBuffer func(Object, []byte)
		GetBuffer   func(Object, []byte) int
	}
	PackedByteArray    PackedFunctionsFor[PackedByteArray, byte]
	PackedColorArray   PackedFunctionsFor[PackedColorArray, Color]
	PackedFloat32Array PackedFunctionsFor[PackedFloat32Array, float32]
	PackedFloat64Array PackedFunctionsFor[PackedFloat64Array, float64]
	PackedInt32Array   PackedFunctionsFor[PackedInt32Array, int32]
	PackedInt64Array   PackedFunctionsFor[PackedInt64Array, int64]
	PackedStringArray  struct {
		Index       func(PackedStringArray, Int) String
		SetIndex    func(PackedStringArray, Int, String)
		CopyAsSlice func(PackedStringArray) []String
	}
	PackedVector2Array PackedFunctionsFor[PackedVector2Array, Vector2]
	PackedVector3Array PackedFunctionsFor[PackedVector3Array, Vector3]
	PackedVector4Array PackedFunctionsFor[PackedVector4Array, Vector4]
	Array              struct {
		Index    func(Array, Int) Variant
		Set      func(self, from Array)
		SetIndex func(Array, Int, Variant)
		SetTyped func(self Array, t VariantType, className StringName, script Object)
	}
	Dictionary struct {
		Index    func(dict Dictionary, key Variant) Variant
		SetIndex func(dict Dictionary, key, val Variant)
	}
	Object struct {
		MethodBindCall        func(method MethodBind, obj Object, arg ...Variant) (Variant, error)
		MethodBindPointerCall func(method MethodBind, obj Object, arg callframe.Args, ret uintptr)
		Destroy               func(Object)
		GetSingleton          func(name StringName) Object
		GetInstanceBinding    func(Object, ExtensionToken, InstanceBindingType) any
		SetInstanceBinding    func(Object, ExtensionToken, any, InstanceBindingType)
		FreeInstanceBinding   func(Object, ExtensionToken)
		SetInstance           func(Object, StringName, ObjectInterface)
		GetClassName          func(Object, ExtensionToken) String
		CastTo                func(Object, ClassTag) Object
		GetInstanceID         func(Object) ObjectID
		GetInstanceFromID     func(ObjectID) Object
	}
	RefCounted struct {
		GetObject func(Object) Object
		SetObject func(Object, Object)
	}
	Callables struct {
		Create func(fn func(...Variant) (Variant, error)) Callable
		Get    func(Callable) (func(...Variant) (Variant, error), bool)
	}
	ClassDB struct {
		ConstructObject func(StringName) Object
		GetClassTag     func(StringName) ClassTag
		GetMethodBind   func(class, method StringName, hash Int) MethodBind

		RegisterClass                 func(library ExtensionToken, name, extends StringName, info ClassInterface)
		RegisterClassMethod           func(library ExtensionToken, class StringName, info Method)
		RegisterClassIntegerConstant  func(library ExtensionToken, class, enum, name StringName, value int64, bitfield bool)
		RegisterClassProperty         func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName)
		RegisterClassPropertyIndexed  func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName, index int64)
		RegisterClassPropertyGroup    func(library ExtensionToken, class StringName, group, prefix String)
		RegisterClassPropertySubGroup func(library ExtensionToken, class StringName, subGroup, prefix String)
		RegisterClassSignal           func(library ExtensionToken, class, signal StringName, args []PropertyInfo)
		UnregisterClass               func(library ExtensionToken, class StringName)
	}
	EditorPlugins struct {
		Add    func(plugin StringName)
		Remove func(plugin StringName)
	}

	GetLibraryPath func(ExtensionToken) String

	// The following fields are primarily reserved for internal use within the gd module,
	// no backwards compatibility is guaranteed for these fields.

	ExtensionToken
	cache

	refCountedClassTag ClassTag

	// extensions instances are mapped here.
	Singletons singletons
}

type Packed[T any] interface {
	PackedByteArray | PackedInt32Array | PackedInt64Array | PackedFloat32Array |
		PackedFloat64Array | PackedStringArray |
		PackedVector2Array | PackedVector3Array | PackedVector4Array |
		PackedColorArray

	pointers.Generic[T, [2]uintptr]
	Len() int
}

type PackedFunctionsFor[T Packed[T], V any] struct {
	Index         func(T, Int) V
	SetIndex      func(T, Int, V)
	CopyAsSlice   func(T) []V
	UnsafePointer func(T) unsafe.Pointer
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

type GDExtensionInitializationLevel int64

const (
	GDExtensionInitializationLevelCore    GDExtensionInitializationLevel = 0
	GDExtensionInitializationLevelServers GDExtensionInitializationLevel = 1
	GDExtensionInitializationLevelScene   GDExtensionInitializationLevel = 2
	GDExtensionInitializationLevelEditor  GDExtensionInitializationLevel = 3
)

type PropertyInfo struct {
	Type       VariantType
	Name       StringName
	ClassName  StringName
	Hint       int64
	HintString String
	Usage      int64
}

type MethodFlags int64

type MethodInfo struct {
	Name             StringName
	ReturnValue      PropertyInfo
	Flags            MethodFlags
	ID               int32
	Arguments        []PropertyInfo
	DefaultArguments []Variant
}

type ClassInterface interface {
	IsVirtual() bool
	IsAbstract() bool
	IsExposed() bool

	CreateInstance() Object

	// ReloadInstance is used to reload an existing object instance
	// when the Go shared library is reloaded from the editor.
	ReloadInstance(Object) ObjectInterface

	GetVirtual(StringName) any
}

type ObjectInterface interface {
	OnCreate()
	Set(StringName, Variant) bool
	Get(StringName) (Variant, bool)
	GetPropertyList() []PropertyInfo
	PropertyCanRevert(StringName) bool
	PropertyGetRevert(StringName) (Variant, bool)
	ValidateProperty(StringName, *PropertyInfo) bool
	Notification(int32, bool)
	ToString() (String, bool)
	Reference()
	Unreference()
	CallVirtual(StringName, any, UnsafeArgs, UnsafeBack)
	GetRID() RID
	Free()
}

type ExtensionClassCallVirtualFunc func(any, UnsafeArgs, UnsafeBack)

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

type Method struct {
	Name                StringName
	Call                func(any, ...Variant) (Variant, error)
	PointerCall         func(any, UnsafeArgs, UnsafeBack)
	MethodFlags         MethodFlags
	ReturnValueInfo     *PropertyInfo
	ReturnValueMetadata ClassMethodArgumentMetadata

	Arguments         []PropertyInfo
	ArgumentsMetadata []ClassMethodArgumentMetadata

	DefaultArguments []Variant
}
