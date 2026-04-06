//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gdunsafe

import (
	"math/rand"
	"reflect"
	"structs"
	"time"
	"unsafe"

	"graphics.gd/internal/threadsafe"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type Int = int64
type Variant [3]uint64
type CallError struct {
	_ structs.HostLayout

	Type     CallErrorType
	Argument int32
	Expected int32
}
type CallErrorType uint32
type ObjectID uint64

type PointerTo[T any] Pointer

type VariadicVariants struct {
	First PointerTo[PointerTo[Variant]]
	Count int
}

var callables threadsafe.Handles[ExtensionCallable, CallableID]

// ExtensionCallable can be implemented to provide an extension-implemented
// [Callable] to the engine.
type ExtensionCallable interface {

	// Call is called whenever the [Callable] is called. It's given a variadic
	// list of variants and the implementation returns a [Variant]. If the
	// arguments aren't compatible, return a non-zero error.
	Call(VariadicVariants) (Variant, CallError)

	// IsValid is called to verify that the [Callable] is valid. Return true
	// if the callable is in a valid state (and callable), otherwise false.
	IsValid() bool

	// Hash is called to hash the [Callable]. Identical underlying
	// implementations of a callable should always return the same value.
	Hash() uint32

	// UnsafeString is called when the [Callable] is being converted by the
	// engine into a string. It should string representation of the callable.
	UnsafeString() String

	// ArgumentCount is called to determine how many arguments the [Callable]
	// expects to receive. Return -1, for variadic arguments.
	ArgumentCount() int

	// Compare is called to compare two different [CustomCallable] values. Return
	// less than zero, if a < b, zero if a = b and more than zero if a > b.
	Compare(ExtensionCallable) int
}

func (ptr Pointer) SetInt32(v int32) {
	ptr.SetUint32(*(*uint32)(unsafe.Pointer(&v)))
}

// Shape is used to correctly transfer data for unsafe calls into the engine.
type Shape uint64

const (
	ShapeEmpty Shape = iota

	ShapeBytes1
	ShapeBytes2
	ShapeBytes4
	ShapeBytes8
	ShapeBytes4x2
	ShapeBytes4x3
	ShapeBytes8x2
	ShapeBytes4x4
	ShapeBytes8x3
	ShapeBytes4x6
	ShapeBytes4x9
	ShapeBytes4x12
	ShapeBytes4x16
)

const (
	SizeVariant     Shape = ShapeBytes8x3
	SizeBool        Shape = ShapeBytes1
	SizeInt         Shape = ShapeBytes8
	SizeFloat       Shape = ShapeBytes8
	SizeVector2     Shape = ShapeBytes4x2
	SizeVector3     Shape = ShapeBytes4x3
	SizeVector4     Shape = ShapeBytes4x4
	SizeColor       Shape = ShapeBytes4x4
	SizeRect2       Shape = ShapeBytes4x4
	SizeRect2i      Shape = ShapeBytes4x4
	SizeVector2i    Shape = ShapeBytes4x2
	SizeVector3i    Shape = ShapeBytes4x3
	SizeVector4i    Shape = ShapeBytes4x4
	SizeTransform2D Shape = ShapeBytes4x6
	SizeTransform3D Shape = ShapeBytes4x12
	SizePlane       Shape = ShapeBytes4x4
	SizeQuaternion  Shape = ShapeBytes4x4
	SizeAABB        Shape = ShapeBytes4x6
	SizeBasis       Shape = ShapeBytes4x9
	SizeProjection  Shape = ShapeBytes4x16
	SizeRID         Shape = ShapeBytes8
	SizeCallable    Shape = ShapeBytes8x2
	SizeSignal      Shape = ShapeBytes8x2

	SizeCallError Shape = ShapeBytes4x3
)

func (shape Shape) SizeResult() (size int) {
	switch shape & 0xF {
	case ShapeEmpty:
		return 0
	case ShapeBytes1:
		return 1
	case ShapeBytes2:
		return 2
	case ShapeBytes4:
		return 4
	case ShapeBytes8:
		return 8
	case ShapeBytes4x2:
		return 4 * 2
	case ShapeBytes4x3:
		return 4 * 3
	case ShapeBytes8x2:
		return 8 * 2
	case ShapeBytes4x4:
		return 4 * 4
	case ShapeBytes8x3:
		return 8 * 3
	case ShapeBytes4x6:
		return 4 * 6
	case ShapeBytes4x9:
		return 4 * 9
	case ShapeBytes4x12:
		return 4 * 12
	case ShapeBytes4x16:
		return 4 * 16
	default:
		panic("Shape.SizeResult: invalid shape")
	}
}

const (
	CallOK               CallErrorType = iota
	CallInvalidMethod    CallErrorType = 1
	CallInvalidArguments CallErrorType = 2
	CallTooManyArguments CallErrorType = 3
	CallTooFewArguments  CallErrorType = 4
	CallInstanceIsNull   CallErrorType = 5
	CallMethodNotConst   CallErrorType = 6
)

func (err CallError) Err() error {
	if err.Type == CallOK {
		return nil
	}
	return err
}

func (err CallError) Error() string {
	switch err.Type {
	case CallInvalidMethod:
		return "Call Invalid Method"
	case CallInvalidArguments:
		return "Call Invalid Arguments"
	case CallTooManyArguments:
		return "Call Too Many Arguments"
	case CallTooFewArguments:
		return "Call Too Few Arguments"
	case CallInstanceIsNull:
		return "Call Instance Is Null"
	case CallMethodNotConst:
		return "Call Method Not Const"
	default:
		return "Unknown Call Error"
	}
}

// LogLevel identifies the severity of a log message.
type LogLevel uint32

const (
	LogError   LogLevel = 0
	LogWarning LogLevel = 1
)

// StringEncoding identifies a string encoding for Decode, Encode, and Intern operations.
type StringEncoding uint8

const (
	Latin1  StringEncoding = iota // ISO 8859-1
	UTF8                          // UTF-8
	UTF16LE                       // UTF-16 little-endian
	UTF16BE                       // UTF-16 big-endian
	UTF32                         // UTF-32
	Wide                          // platform-native wide characters (wchar_t)
)

// ALIGN_UP aligns a value to the next multiple of align.
func alignUp(value, align uint32) uint32 {
	return (value + (align - 1)) & ^(align - 1)
}

func (shape Shape) Alignment() int {
	switch shape {
	case ShapeEmpty:
		return 0
	case ShapeBytes1:
		return 1
	case ShapeBytes2:
		return 2
	case ShapeBytes4, ShapeBytes4x2, ShapeBytes4x3, ShapeBytes4x4, ShapeBytes4x6, ShapeBytes4x9, ShapeBytes4x12, ShapeBytes4x16:
		return 4
	case ShapeBytes8, ShapeBytes8x2, ShapeBytes8x3:
		return 8
	default:
		panic("Shape.Alignment: invalid shape")
	}
}

func (shape Shape) SizeArguments() (size int) {
	for i := 1; i < 16; i++ {
		var current = (shape >> (i * 4)) & 0xF
		switch current {
		case ShapeEmpty:
			return size
		default:
			size += current.SizeResult()
			size = int(alignUp(uint32(size), uint32(current.Alignment())))
		}
	}
	return
}

const (
	TypeNil                VariantType = 0
	TypeBool               VariantType = 1
	TypeInt                VariantType = 2
	TypeFloat              VariantType = 3
	TypeString             VariantType = 4
	TypeVector2            VariantType = 5
	TypeVector2i           VariantType = 6
	TypeRect2              VariantType = 7
	TypeRect2i             VariantType = 8
	TypeVector3            VariantType = 9
	TypeVector3i           VariantType = 10
	TypeTransform2D        VariantType = 11
	TypeVector4            VariantType = 12
	TypeVector4i           VariantType = 13
	TypePlane              VariantType = 14
	TypeQuaternion         VariantType = 15
	TypeAABB               VariantType = 16
	TypeBasis              VariantType = 17
	TypeTransform3D        VariantType = 18
	TypeProjection         VariantType = 19
	TypeColor              VariantType = 20
	TypeStringName         VariantType = 21
	TypeNodePath           VariantType = 22
	TypeRID                VariantType = 23
	TypeObject             VariantType = 24
	TypeCallable           VariantType = 25
	TypeSignal             VariantType = 26
	TypeDictionary         VariantType = 27
	TypeArray              VariantType = 28
	TypePackedByteArray    VariantType = 29
	TypePackedInt32Array   VariantType = 30
	TypePackedInt64Array   VariantType = 31
	TypePackedFloat32Array VariantType = 32
	TypePackedFloat64Array VariantType = 33
	TypePackedStringArray  VariantType = 34
	TypePackedVector2Array VariantType = 35
	TypePackedVector3Array VariantType = 36
	TypePackedColorArray   VariantType = 37
	TypePackedVector4Array VariantType = 38
)

func (p PackedArray[T]) Type() VariantType {
	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[byte]():
		return TypePackedByteArray
	case reflect.TypeFor[int32]():
		return TypePackedInt32Array
	case reflect.TypeFor[int64]():
		return TypePackedInt64Array
	case reflect.TypeFor[float32]():
		return TypePackedFloat32Array
	case reflect.TypeFor[float64]():
		return TypePackedFloat64Array
	case reflect.TypeFor[String]():
		return TypePackedStringArray
	case reflect.TypeFor[Vector2.XY]():
		return TypePackedVector2Array
	case reflect.TypeFor[Vector3.XYZ]():
		return TypePackedVector3Array
	case reflect.TypeFor[Color.RGBA]():
		return TypePackedColorArray
	case reflect.TypeFor[Vector4.XYZW]():
		return TypePackedVector4Array
	default:
		return 0
	}
}

type ExtensionInstance interface {
	Set(StringName, Variant) bool
	Get(StringName) (Variant, bool)
	HasDefault(StringName) bool
	GetDefault(StringName) (Variant, bool)
	PropertyList() PropertyList
	ValidateProperty(StringName) bool
	Notification(what int32, reverse bool)
	UnsafeString() String
	Reference(bool) bool
	RID() RID.Any
}

type ExtensionFunction interface {
	PointerCall(ExtensionInstance, Pointer, Pointer)
	CheckedCall(ExtensionInstance, VariadicVariants) Variant
	DynamicCall(ExtensionInstance, VariadicVariants) (Variant, CallError)
}

type ExtensionClass interface {
	Create(notify_postinitialize bool) Object
	Method(name StringName, hash uint32) ExtensionFunction
}

type PropertyInfo PropertyList

// ExtensionScript is an interface that can be used to implement a script
// in the engine. Useful when creating new scripting languages.
type ExtensionScript interface {
	ExtensionInstance

	// PropertyCategory used in editor's inspector as a heading to
	// group script properties together.
	PropertyCategory() StringName

	// PropertyType returns the variant type for the given Property.
	PropertyType(StringName) VariantType

	// Owner should return the object that the script is attached to.
	Owner() Object

	// ExportedProperties iterator. Should call the given function
	// for each exported property. Used when serializing the script.
	ExportedProperties(func(StringName, Variant) bool)

	// MethodList should return a [MethodList] that represents each
	// of the script's defined methods.
	MethodList() MethodList

	// HasMethod returns true if the script has a method with the
	// given name.
	HasMethod(StringName) bool

	// MethodArgumentCount returns the number of arguments that the
	// given method expects. Use -1, if the method is variadic.
	MethodArgumentCount(StringName) int

	// Script returns the underlying Script object.
	Script() Object

	// IsPlaceholder returns true if the script is a placeholder script
	// ie. the script failed to load, or the language is not available.
	IsPlaceholder() bool

	// ScriptLanguage returns the underlying ScriptLanguage object.
	ScriptLanguage() Object
}

type InitializationLevel uint32

func OnEngineInit(func(InitializationLevel)) {}
func OnEngineExit(func(InitializationLevel)) {}
func OnFirstFrame(func())                    {}
func OnEveryFrame(func())                    {}
func OnFinalFrame(func())                    {}

type TaskID uintptr

func OnWorkerThreadPoolTask(func(TaskID))
func OnWorkerThreadPoolGroupTask(func(TaskID, int32))

func OnEditorClassDetection(func(PackedArray[String]) PackedArray[String]) {}

// just a placeholder for functions that don't need to be implemented
// as they are already available in the Go standard library.

func randomize() { //gd:randomize
	rand.Seed(time.Now().UnixNano())
}

func seed(s int) { //gd:seed
	rand.Seed(int64(s))
}

func rand_from_seed(seed int) *rand.Rand { //gd:rand_from_seed
	return rand.New(rand.NewSource(int64(seed)))
}

func weakref(v any) any { return v } //gd:weakref
