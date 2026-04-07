//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gdunsafe

import (
	"reflect"
	"structs"

	"graphics.gd/internal/threadsafe"
	"graphics.gd/variant"
	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
	"graphics.gd/variant/Vector4i"
)

type Iterator Variant

type (
	String struct{}
	Object struct{}

	Dictionary struct{}
	StringName struct{}
	NodePath   struct{}

	Signal [2]uint64

	Script         Object
	ScriptLanguage Object
)

type Type struct {
	vtype      variant.Type
	shape      Shape
	class_tag  uintptr
	class_name StringName
	script     Object
}

func (t Type) Shape() Shape { return t.shape }

func (t Type) Free() {
	Free(t.class_name)
}

type Packable interface {
	byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String
}

type ObjectID uint64

// PointerTo is a [Pointer] that points to a value of type T, it should
// be treated as if it were an [unsafe.Pointer] from a different process.
type PointerTo[T Any | Variant | PointerTo[Variant] | Packable] Pointer

// MutablePointerTo is a [MutablePointer] that points to a value of type T, it should
// be treated as if it were an [unsafe.Pointer] from a different process.
type MutablePointerTo[T Any | Variant | Packable] MutablePointer

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
	classes   threadsafe.Handles[ExtensionClass, uintptr]
	instances threadsafe.Handles[ExtensionInstance, uintptr]
	functions threadsafe.Handles[ExtensionFunction, uintptr]
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

// Error implements the [error] interface.
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

type Any interface {
	bool |
		int64 |
		float64 |
		String |
		~Vector2.XY |
		~Vector2i.XY |
		~Rect2.PositionSize |
		~Rect2i.PositionSize |
		~Vector3.XYZ |
		~Vector3i.XYZ |
		~Transform2D.OriginXY |
		~Vector4.XYZW |
		~Vector4i.XYZW |
		~Plane.NormalD |
		~Quaternion.IJKX |
		~AABB.PositionSize |
		~Basis.XYZ |
		~Transform3D.BasisOrigin |
		~Projection.XYZW |
		~Color.RGBA |
		StringName |
		NodePath |
		RID.Any |
		Object |
		Callable |
		Signal |
		Dictionary |
		Array |
		PackedArray[byte] |
		PackedArray[int32] |
		PackedArray[int64] |
		PackedArray[float32] |
		PackedArray[float64] |
		PackedArray[Color.RGBA] |
		PackedArray[Vector2.XY] |
		PackedArray[Vector3.XYZ] |
		PackedArray[Vector4.XYZW] |
		PackedArray[String]
}

// Type returns the [variant.Type] of the packed array.
func (p PackedArray[T]) Type() variant.Type {
	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[byte]():
		return variant.TypePackedByteArray
	case reflect.TypeFor[int32]():
		return variant.TypePackedInt32Array
	case reflect.TypeFor[int64]():
		return variant.TypePackedInt64Array
	case reflect.TypeFor[float32]():
		return variant.TypePackedFloat32Array
	case reflect.TypeFor[float64]():
		return variant.TypePackedFloat64Array
	case reflect.TypeFor[String]():
		return variant.TypePackedStringArray
	case reflect.TypeFor[Vector2.XY]():
		return variant.TypePackedVector2Array
	case reflect.TypeFor[Vector3.XYZ]():
		return variant.TypePackedVector3Array
	case reflect.TypeFor[Color.RGBA]():
		return variant.TypePackedColorArray
	case reflect.TypeFor[Vector4.XYZW]():
		return variant.TypePackedVector4Array
	default:
		return 0
	}
}

// ExtensionInstance is an interface that can be used to implement an instance
// of an [ExtensionClass] in the engine.
type ExtensionInstance interface {
	Set(field StringName, value Variant) bool
	Get(field StringName) (Variant, bool)
	HasDefault(field StringName) bool
	GetDefault(field StringName) (Variant, bool)
	PropertyList() PropertyList
	ValidateProperty(Property) bool
	Notification(what int32, reverse bool)
	UnsafeString() String

	RID() RID.Any

	Reference(increment bool) bool

	//ReferenceIncremented()
	//ReferenceDecremented() bool
}

type Property struct {
	Type       variant.Type
	Name       StringName
	ClassName  StringName
	Hint       uint32
	HintString String
	Usage      uint32
}

type ExtensionFunction interface {
	PointerCall(instance ExtensionInstance, result Pointer, args Pointer)
	CheckedCall(instance ExtensionInstance, args Variants) Variant
	DynamicCall(instance ExtensionInstance, args Variants) (Variant, Error)
}

// ExtensionClass is an interface that can be implemented to create a new
// class in the engine. Pass an implementation of this interface to the
// [RegisterClass] function to register it with the engine.
type ExtensionClass interface {

	// Create should instantiate the underlying object with [New] and then
	// call [SetupExtension] on it.
	Create(notify_postinitialize bool) Object

	// Method should return the [ExtensionFunction] for the given method.
	Method(name StringName, hash uint32) ExtensionFunction
}

// ExtensionScript is an interface that can be used to implement a script
// in the engine. Useful when creating new scripting languages.
type ExtensionScript interface {
	ExtensionInstance

	// PropertyCategory used in editor's inspector as a heading to
	// group script properties together.
	PropertyCategory() StringName

	// PropertyType returns the variant type for the given Property.
	PropertyType(field StringName) variant.Type

	// Owner should return the object that the script is attached to.
	Owner() Object

	// ExportedProperties iterator. Should call the given function
	// for each exported property. Used when serializing the script.
	ExportedProperties(func(field StringName, value Variant) bool)

	// MethodList should return a [MethodList] that represents each
	// of the script's defined methods.
	MethodList() MethodList

	// HasMethod returns true if the script has a method with the
	// given name.
	HasMethod(name StringName) bool

	// MethodArgumentCount returns the number of arguments that the
	// given method expects. Use -1, if the method is variadic.
	MethodArgumentCount(method_name StringName) int

	// Script returns the underlying Script object.
	Script() Object

	// IsPlaceholder returns true if the script is a placeholder script
	// ie. the script failed to load, or the language is not available.
	IsPlaceholder() bool

	// ScriptLanguage returns the underlying ScriptLanguage object.
	ScriptLanguage() Object
}

// InitializationLevel values are documented in the [GDExtension] package.
//
// [GDExtension]: https://pkg.go.dev/graphics.gd/classdb/GDExtension#InitializationLevel
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

// OnEngineInit registers a function to be called when the engine initializes,
// this should be called early inside an init function.
func OnEngineInit(fn func(level InitializationLevel)) { onEngineInit = fn }

// OnEngineExit registers a function to be called when the engine exits,
// this should be called early inside an init function.
func OnEngineExit(fn func(level InitializationLevel)) { onEngineExit = fn }

// OnFirstFrame registers a function to be called on the first frame,
// this should be called early inside an init function.
func OnFirstFrame(fn func()) { onFirstFrame = fn }

// OnEveryFrame registers a function to be called on every frame. Can
// be called at any time but cannot be removed after being added.
func OnEveryFrame(fn func()) { onEveryFrame = fn }

// OnFinalFrame registers a function to be called on the final frame,
// this should be called before the engine shuts down.
func OnFinalFrame(fn func()) { onFinalFrame = fn }

type TaskID uintptr

// variantTypeOf maps a Go type in the [Any] constraint to its [variant.Type].
func variantTypeOf[T Any]() variant.Type {
	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[bool]():
		return variant.TypeBool
	case reflect.TypeFor[int64]():
		return variant.TypeInt
	case reflect.TypeFor[float64]():
		return variant.TypeFloat
	case reflect.TypeFor[String]():
		return variant.TypeString
	case reflect.TypeFor[Vector2.XY]():
		return variant.TypeVector2
	case reflect.TypeFor[Vector2i.XY]():
		return variant.TypeVector2i
	case reflect.TypeFor[Rect2.PositionSize]():
		return variant.TypeRect2
	case reflect.TypeFor[Rect2i.PositionSize]():
		return variant.TypeRect2i
	case reflect.TypeFor[Vector3.XYZ]():
		return variant.TypeVector3
	case reflect.TypeFor[Vector3i.XYZ]():
		return variant.TypeVector3i
	case reflect.TypeFor[Transform2D.OriginXY]():
		return variant.TypeTransform2D
	case reflect.TypeFor[Vector4.XYZW]():
		return variant.TypeVector4
	case reflect.TypeFor[Vector4i.XYZW]():
		return variant.TypeVector4i
	case reflect.TypeFor[Plane.NormalD]():
		return variant.TypePlane
	case reflect.TypeFor[Quaternion.IJKX]():
		return variant.TypeQuaternion
	case reflect.TypeFor[AABB.PositionSize]():
		return variant.TypeAABB
	case reflect.TypeFor[Basis.XYZ]():
		return variant.TypeBasis
	case reflect.TypeFor[Transform3D.BasisOrigin]():
		return variant.TypeTransform3D
	case reflect.TypeFor[Projection.XYZW]():
		return variant.TypeProjection
	case reflect.TypeFor[Color.RGBA]():
		return variant.TypeColor
	case reflect.TypeFor[StringName]():
		return variant.TypeStringName
	case reflect.TypeFor[NodePath]():
		return variant.TypeNodePath
	case reflect.TypeFor[RID.Any]():
		return variant.TypeRID
	case reflect.TypeFor[Object]():
		return variant.TypeObject
	case reflect.TypeFor[Callable]():
		return variant.TypeCallable
	case reflect.TypeFor[Signal]():
		return variant.TypeSignal
	case reflect.TypeFor[Dictionary]():
		return variant.TypeDictionary
	case reflect.TypeFor[Array]():
		return variant.TypeArray
	case reflect.TypeFor[PackedArray[byte]]():
		return variant.TypePackedByteArray
	case reflect.TypeFor[PackedArray[int32]]():
		return variant.TypePackedInt32Array
	case reflect.TypeFor[PackedArray[int64]]():
		return variant.TypePackedInt64Array
	case reflect.TypeFor[PackedArray[float32]]():
		return variant.TypePackedFloat32Array
	case reflect.TypeFor[PackedArray[float64]]():
		return variant.TypePackedFloat64Array
	case reflect.TypeFor[PackedArray[String]]():
		return variant.TypePackedStringArray
	case reflect.TypeFor[PackedArray[Vector2.XY]]():
		return variant.TypePackedVector2Array
	case reflect.TypeFor[PackedArray[Vector3.XYZ]]():
		return variant.TypePackedVector3Array
	case reflect.TypeFor[PackedArray[Color.RGBA]]():
		return variant.TypePackedColorArray
	case reflect.TypeFor[PackedArray[Vector4.XYZW]]():
		return variant.TypePackedVector4Array
	default:
		return variant.TypeNil
	}
}

// shapeOf maps a Go type in the [Any] constraint to its [Shape].
func shapeOf[T Any]() Shape {
	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[bool]():
		return ShapeBool
	case reflect.TypeFor[int64]():
		return ShapeInt
	case reflect.TypeFor[float64]():
		return ShapeFloat
	case reflect.TypeFor[String]():
		return ShapeString
	case reflect.TypeFor[Vector2.XY]():
		return ShapeVector2
	case reflect.TypeFor[Vector2i.XY]():
		return ShapeVector2i
	case reflect.TypeFor[Rect2.PositionSize]():
		return ShapeRect2
	case reflect.TypeFor[Rect2i.PositionSize]():
		return ShapeRect2i
	case reflect.TypeFor[Vector3.XYZ]():
		return ShapeVector3
	case reflect.TypeFor[Vector3i.XYZ]():
		return ShapeVector3i
	case reflect.TypeFor[Transform2D.OriginXY]():
		return ShapeTransform2D
	case reflect.TypeFor[Vector4.XYZW]():
		return ShapeVector4
	case reflect.TypeFor[Vector4i.XYZW]():
		return ShapeVector4i
	case reflect.TypeFor[Plane.NormalD]():
		return ShapePlane
	case reflect.TypeFor[Quaternion.IJKX]():
		return ShapeQuaternion
	case reflect.TypeFor[AABB.PositionSize]():
		return ShapeAABB
	case reflect.TypeFor[Basis.XYZ]():
		return ShapeBasis
	case reflect.TypeFor[Transform3D.BasisOrigin]():
		return ShapeTransform3D
	case reflect.TypeFor[Projection.XYZW]():
		return ShapeProjection
	case reflect.TypeFor[Color.RGBA]():
		return ShapeColor
	case reflect.TypeFor[StringName]():
		return ShapeStringName
	case reflect.TypeFor[NodePath]():
		return ShapeNodePath
	case reflect.TypeFor[RID.Any]():
		return ShapeRID
	case reflect.TypeFor[Object]():
		return ShapeObject
	case reflect.TypeFor[Callable]():
		return ShapeCallable
	case reflect.TypeFor[Signal]():
		return ShapeSignal
	case reflect.TypeFor[Dictionary]():
		return ShapeDictionary
	case reflect.TypeFor[Array]():
		return ShapeArray
	case reflect.TypeFor[PackedArray[byte]](),
		reflect.TypeFor[PackedArray[int32]](),
		reflect.TypeFor[PackedArray[int64]](),
		reflect.TypeFor[PackedArray[float32]](),
		reflect.TypeFor[PackedArray[float64]](),
		reflect.TypeFor[PackedArray[String]](),
		reflect.TypeFor[PackedArray[Vector2.XY]](),
		reflect.TypeFor[PackedArray[Vector3.XYZ]](),
		reflect.TypeFor[PackedArray[Color.RGBA]](),
		reflect.TypeFor[PackedArray[Vector4.XYZW]]():
		return ShapePackedArray
	case reflect.TypeFor[Variant]():
		return ShapeVariant
	default:
		return 0
	}
}

/*
func OnWorkerThreadPoolTask(fn func(TaskID))             { onWorkerThreadPoolTask = fn }
func OnWorkerThreadPoolGroupTask(fn func(TaskID, int32)) { onWorkerThreadPoolGroupTask = fn }

func OnEditorClassDetection(fn func(PackedArray[String]) PackedArray[String]) {
	onEditorClassDetection = fn
}
*/
