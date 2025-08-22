//go:build !generate

package gd

import (
	"reflect"
	"structs"
	"unsafe"

	"graphics.gd/internal/gdextension"

	"runtime.link/api"
)

type VariantType = gdextension.VariantType

type Address uintptr

// API specification for Godot's GDExtension.
type API struct {
	api.Specification

	GetGodotVersion func() Version

	Object struct {
		GetInstanceBinding  func([1]Object, ExtensionToken, InstanceBindingType) any
		SetInstanceBinding  func([1]Object, ExtensionToken, any, InstanceBindingType)
		FreeInstanceBinding func([1]Object, ExtensionToken)
		SetInstance         func([1]Object, StringName, ObjectInterface)
		CastTo              func([1]Object, ClassTag) [1]Object
	}
	Callables struct {
		Create func(fn func(...Variant) (Variant, error)) Callable
		Get    func(Callable) (func(...Variant) (Variant, error), bool)
	}
	ClassDB struct {
		GetClassTag func(StringName) ClassTag

		RegisterClass                func(library ExtensionToken, name, extends StringName, info ClassInterface)
		RegisterClassMethod          func(library ExtensionToken, class StringName, info Method)
		RegisterClassProperty        func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName)
		RegisterClassPropertyIndexed func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName, index int64)
		RegisterClassSignal          func(library ExtensionToken, class, signal StringName, args []PropertyInfo)
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
	_ structs.HostLayout

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

type Operator = gdextension.VariantOperator

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

	CreateInstance() [1]Object
	GetVirtual(StringName) any
}

type ObjectInterface interface {
	OnCreate(reflect.Value)
	Set(StringName, Variant) bool
	Get(StringName) (Variant, bool)
	GetPropertyList() []PropertyInfo
	PropertyCanRevert(StringName) bool
	PropertyGetRevert(StringName) (Variant, bool)
	ValidateProperty(*PropertyInfo) bool
	Notification(int32, bool)
	ToString() (String, bool)
	Reference()
	Unreference()
	CallVirtual(StringName, any, Address, Address)
	GetRID() RID
	Free()
}

type ExtensionClassCallVirtualFunc func(any, Address, Address)

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
	PointerCall         func(any, Address, Address)
	MethodFlags         MethodFlags
	ReturnValueInfo     *PropertyInfo
	ReturnValueMetadata ClassMethodArgumentMetadata

	Arguments         []PropertyInfo
	ArgumentsMetadata []ClassMethodArgumentMetadata

	DefaultArguments []Variant
}
