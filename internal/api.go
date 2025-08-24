//go:build !generate

package gd

import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/gdextension"

	"runtime.link/api"
)

type VariantType = gdextension.VariantType

type Address uintptr

// API specification for Godot's GDExtension.
type API struct {
	api.Specification

	// The following fields are primarily reserved for internal use within the gd module,
	// no backwards compatibility is guaranteed for these fields.

	ExtensionToken
	cache

	refCountedClassTag gdextension.ObjectType
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

type CallError = gdextension.CallError

type InstanceBindingType unsafe.Pointer

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
