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
	"graphics.gd/variant"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

// Variant is the raw representation for a variant value in the engine.
// It should be destroyed with [Variant.Free] when no longer in use.
type Variant [3]uint64

type ObjectID uint64

// PointerTo is a [Pointer] that points to a value of type T, it should
// be treated as if it were an [unsafe.Pointer] from a different process.
type PointerTo[T any] Pointer

// Variants accessor, used to represent zero or more arguments.
type Variants struct {
	first PointerTo[PointerTo[Variant]]
	count int
}

// Len returns the number of variants in the [Variants] list.
func (args Variants) Len() int { return max(0, args.count) }

// ExpectedLen returns an error if the number of variants does not match
// the expected length.
func (args Variants) ExpectedLen(n int) Error {
	if args.count == n {
		return Error{}
	}
	var kind = errorTooFewArguments
	if args.count > n {
		kind = errorTooManyArguments
	}
	return Error{
		error:    kind,
		argument: -1,
		expected: int32(n),
	}
}

// ExpectedArg returns the i-th variant and an error if it does not match
// the expected type.
func (args Variants) ExpectedArg(i int, vtype variant.Type) (Variant, Error) {
	arg := args.Index(i)
	if arg.Type() != vtype {
		return arg, Error{
			error:    errorInvalidArgumentType,
			argument: int32(i),
			expected: int32(vtype),
		}
	}
	return arg, Error{}
}

var callables threadsafe.Handles[ExtensionCallable, CallableID]

var (
	classes   threadsafe.Handles[ExtensionClass, ExtensionClassID]
	instances threadsafe.Handles[ExtensionInstance, ExtensionInstanceID]
	functions threadsafe.Handles[ExtensionFunction, FunctionID]
)

// ExtensionInstances is an iterator over each active [ExtensionInstance].
func ExtensionInstances(fn func(ExtensionInstance) bool) { instances.All(fn) }

// ExtensionCallable can be implemented to provide an extension-implemented
// [Callable] that can be passed to the engine.
type ExtensionCallable interface {

	// Call is called whenever the [Callable] is called. It's given a variadic
	// list of variants and the implementation returns a [Variant]. If the
	// arguments aren't compatible, return a non-nil [error].
	Call(Variants) (Variant, Error)

	// IsValid is called to verify that the [Callable] is valid. Return true
	// if the callable is in a valid state (and callable), otherwise false.
	IsValid() bool

	// Hash is called to hash the [Callable]. Identical underlying
	// implementations of a callable should always return the same value.
	Hash() uint32

	// UnsafeString is called when the [Callable] is being converted by the
	// engine into a [String]. This should be a human-readable representation.
	UnsafeString() String

	// ArgumentCount is called to determine how many arguments the [Callable]
	// expects to receive. Return -1, for a variadic number of arguments.
	ArgumentCount() int

	// Compare is called to compare two different [CustomCallable] values. Return
	// less than zero, if a < b, zero if a = b and more than zero if a > b.
	Compare(ExtensionCallable) int
}

// SetInt16 writes a uint16 value to the underlying memory address of the [Pointer].
func (ptr Pointer) SetInt16(v int16) { ptr.SetUint16(*(*uint16)(unsafe.Pointer(&v))) }

// SetInt32 writes a uint32 value to the underlying memory address of the [Pointer].
func (ptr Pointer) SetInt32(v int32) {
	ptr.SetUint32(*(*uint32)(unsafe.Pointer(&v)))
}

// Shape is used to represent the shape (structure, size and alignment) for a value
// being passed to the engine. It's a four bit value that can represent up to 16
// components. Typically the return value is represented in the lowest four bits,
// the arguments (including the receiver) are represented in subsequent sets of four
// bits each, increasing in significance.
//
// The [Shape] is used to correctly copy arguments between the engine and Go on
// platforms where the engine exists in a different address space (ie. on WASM).
type Shape uint64

const (
	empty Shape = iota // means no data.

	bytes1   // shape for a single byte value
	bytes2   // shape for a two byte value
	bytes4   // shape for a four byte value
	bytes8   // shape for an eight byte value
	bytes4x2 // shape for two 4-byte values
	bytes4x3 // shape for three 4-byte values
	bytes8x2 // shape for two 8-byte values
	bytes4x4 // shape for four 4-byte values

)

const (
	ShapeBool     Shape = bytes1   // shape of a [bool]
	ShapeInt      Shape = bytes8   // shape of an [int64]
	ShapeFloat    Shape = bytes8   // shape of a [float64]
	ShapeColor    Shape = bytes4x4 // shape of a [Color.RGBA]
	ShapeRect2i   Shape = bytes4x4 // shape of a [Rect2i.PositionSize]
	ShapeVector2i Shape = bytes4x2 // shape of a [Vector2i.XY]
	ShapeVector3i Shape = bytes4x3 // shape of a [Vector3i.XYZ]
	ShapeVector4i Shape = bytes4x4 // shape of a [Vector4i.XYZW]
	ShapeRID      Shape = bytes8   // shape of an [RID.Any]
	ShapeCallable Shape = bytes8x2 // shape of a [Callable]
	ShapeSignal   Shape = bytes8x2 // shape of a [Signal]

	ShapeError Shape = bytes4x3 // shape of an [Error]
)

// Error is the underlying structure for an engine [error].
type Error struct {
	_ structs.HostLayout

	error    errorType
	argument int32
	expected int32
}

type errorType uint32

const (
	errorInvalidMethod       errorType = 1
	errorInvalidArgumentType errorType = 2
	errorTooManyArguments    errorType = 3
	errorTooFewArguments     errorType = 4
	errorInstanceIsNull      errorType = 5
	errorMethodNotConst      errorType = 6
)

func (err Error) Error() string {
	switch err.error {
	case errorInvalidMethod:
		return "Call Invalid Method"
	case errorInvalidArgumentType:
		return "Call Invalid Arguments"
	case errorTooManyArguments:
		return "Call Too Many Arguments"
	case errorTooFewArguments:
		return "Call Too Few Arguments"
	case errorInstanceIsNull:
		return "Call Instance Is Null"
	case errorMethodNotConst:
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

// SizeArguments returns the total size of the argument structure for a [Shape].
func (shape Shape) SizeArguments() (size int) {
	for i := 1; i < 16; i++ {
		var current = (shape >> (i * 4)) & 0xF
		switch current {
		case empty:
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
	CheckedCall(ExtensionInstance, Variants) Variant
	DynamicCall(ExtensionInstance, Variants) (Variant, Error)
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

var (
	onEngineInit                func(InitializationLevel)
	onEngineExit                func(InitializationLevel)
	onFirstFrame                func()
	onEveryFrame                func()
	onFinalFrame                func()
	onWorkerThreadPoolTask      func(TaskID)
	onWorkerThreadPoolGroupTask func(TaskID, int32)
	onEditorClassDetection      func(PackedArray[String]) PackedArray[String]
)

func OnEngineInit(fn func(InitializationLevel)) { onEngineInit = fn }
func OnEngineExit(fn func(InitializationLevel)) { onEngineExit = fn }
func OnFirstFrame(fn func())                    { onFirstFrame = fn }
func OnEveryFrame(fn func())                    { onEveryFrame = fn }
func OnFinalFrame(fn func())                    { onFinalFrame = fn }

type TaskID uintptr

func OnWorkerThreadPoolTask(fn func(TaskID))             { onWorkerThreadPoolTask = fn }
func OnWorkerThreadPoolGroupTask(fn func(TaskID, int32)) { onWorkerThreadPoolGroupTask = fn }

func OnEditorClassDetection(fn func(PackedArray[String]) PackedArray[String]) {
	onEditorClassDetection = fn
}

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
