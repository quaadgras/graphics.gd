//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gdunsafe

import (
	"reflect"
	"strconv"
	"structs"
	"unsafe"

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
	// String is a UTF-32 encoded string.
	String struct {
		_   [0]*String
		raw gdString
	}
	// Object is a reference to a Godot object.
	Object struct {
		_   [0]*Object
		raw gdObject
	}
	// Array is a sequentially-indexed list of variant values.
	Array struct {
		_   [0]*Array
		raw gdArray
	}
	// Dictionary is a key-value map of variant values.
	Dictionary struct {
		_   [0]*Dictionary
		raw gdDictionary
	}
	// StringName is a string that is used as a unique identifier for a Godot object.
	StringName struct {
		_   [0]*StringName
		raw gdStringName
	}
	// NodePath is a path to a Node.
	NodePath struct {
		_   [0]*NodePath
		raw gdNodePath
	}
	// Signal is a broadcast queue.
	Signal struct {
		_   [0]*Signal
		raw gdSignal
	}
	// Callable is a closure or function.
	Callable struct {
		_   [0]*Callable
		raw gdCallable
	}

	// RefCounted is a reference-counted [Object].
	RefCounted struct {
		_   [0]*RefCounted
		raw gdObject
	}
	// Script is a reference to a [Script] object.
	Script struct {
		_   [0]*Script
		raw gdObject
	}
	// ScriptLanguage is a reference to a [ScriptLanguage] object.
	ScriptLanguage struct {
		_   [0]*ScriptLanguage
		raw gdObject
	}

	PropertyList  Pointer // PropertyList is a container for a list of [Property] values.
	MethodList    Pointer // MethodList is a container for a list of [Method] values.
	MethodPointer Pointer // MethodPointer is a pointer to a method on an [Class].

	Class StringName // Class is a [StringName] that is used as a unique identifier for a class.

	// ClassTag is an opaque identifier used for a class.
	ClassTag struct {
		_   [0]*ClassTag
		raw gdClassTag
	}
)

type Type struct {
	vtype     variant.Type
	shape     Shape
	class     Class
	class_tag ClassTag
	script    Variant
}

func TypeFrom(vtype variant.Type) Type {
	return Type{
		vtype: vtype,
	}
}

func TypeByName(name string) Type {
	return Type{
		vtype: variant.TypeObject,
		shape: ShapeObject,
	}
}

func TypeFor[T Any | Class]() Type {
	return Type{}
}

// ClassWithScript returns a [Type] with the given [Class] and [Script].
func ClassWithScript(class Class, script Script) Type {
	return Type{
		vtype:     variant.TypeObject,
		shape:     ShapeObject,
		class:     class,
		class_tag: class.Tag(),
		script:    VariantFrom[Object](Object{raw: script.raw}),
	}
}

func (t Type) Shape() Shape { return t.shape }

func (t Type) Free() {
	if t.class != (Class{}) {
		Free(StringName(t.class))
	}
}

// Packable types that can be packed into a [PackedArray].
type Packable interface {
	byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String
}

// Addressable types that can be used as a [Pointer] or [MutablePointer].
type Addressable interface {
	Any | Packable | Variant | PointerTo[Variant] | Pointer | uint64 | uint32 | uint16 | [2]uint64 | struct{} | Error
}

type Returnable interface {
	Any | Variant | struct{}
}

type ObjectID uint64

// PointerTo is a [Pointer] that points to a value of type T, it should
// be treated as if it were an [unsafe.Pointer] from a different process.
type PointerTo[T Addressable] Pointer

// MutablePointerTo is a [MutablePointer] that points to a value of type T, it should
// be treated as if it were an [unsafe.Pointer] from a different process.
type MutablePointerTo[T Addressable] MutablePointer

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

var callables threadsafe.Handles[ExtensionCallable, uintptr]

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

// Encoding interface for supported string encodings, [Latin1], [UTF8], [UTF16], [UTF32] and [Wide].
type Encoding interface {
	Decode(s String, buf []byte) int // Decode the string into the buffer. Returns bytes written.
	String(s string) String          // String returns a new [String] from the given string.
	Intern(s string) StringName      // Intern returns a [StringName] for the given string.
}

// Latin1 implements [Encoding].
var Latin1 latin1

type latin1 struct{}

// UTF8 implements [Encoding] via UTF-8.
var UTF8 utf8

type utf8 struct{}

// UTF16 implements [Encoding], can be set to true for
// little-endian UTF-16 encoding, or false for big-endian.
type UTF16 = utf16

type utf16 struct{}

// UTF32 implements [Encoding] via UTF-32.
var UTF32 utf32

type utf32 struct{}

// Wide implements [Encoding] via platform-native wide characters (wchar_t).
var Wide wide

type wide struct{}

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

// Any is a constraint that matches any Variant type.
type Any interface {
	bool |
		int64 |
		float64 |
		String |
		Vector2.XY |
		Vector2i.XY |
		Rect2.PositionSize |
		Rect2i.PositionSize |
		Vector3.XYZ |
		Vector3i.XYZ |
		Transform2D.OriginXY |
		Vector4.XYZW |
		Vector4i.XYZW |
		Plane.NormalD |
		Quaternion.IJKX |
		AABB.PositionSize |
		Basis.XYZ |
		Transform3D.BasisOrigin |
		Projection.XYZW |
		Color.RGBA |
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

	// Set is called when the engine attempts to set a value for
	// the given field/property in the extension instance.
	Set(field StringName, value Variant) bool

	// Get is called when the engine attempts to get the value of
	// the given field/property in the extension instance.
	Get(field StringName) (Variant, bool)

	// HasDefault is called when the engine attempts to check if
	// the given field/property has a default value.
	HasDefault(field StringName) bool

	// GetDefault is called when the engine attempts to get the
	// default value of the given field/property.
	GetDefault(field StringName) (Variant, bool)

	// PropertyList is called when the engine attempts to get the
	// list of properties for the extension instance.
	PropertyList() PropertyList

	// ValidateProperty is called when the engine attempts to validate
	// a property for the extension instance.
	ValidateProperty(Property) bool

	// Notification is called when the engine sends an Object
	// notification to the extension instance.
	Notification(what int32, reverse bool)

	// UnsafeString is called when the engine attempts to get a
	// string representation of the extension instance.
	UnsafeString() String

	// RID is called when the engine attempts to get the RID of
	// the extension instance.
	RID() RID.Any

	// Reference is called when the engine attempts to increment
	// or decrement the reference count of the extension instance.
	Reference(increment bool) bool

	//ReferenceIncremented()
	//ReferenceDecremented() bool
}

// Property description.
type Property struct {
	Type       variant.Type // Type of the property.
	Name       StringName   // Name of the property.
	ClassName  StringName   // Class name of the property (if it's an [Object] property).
	Hint       uint32       // Hint for the property.
	HintString String       // Hint-specific string for the property.
	Usage      uint32       // Usage flags.
}

// ExtensionFunction can be implemented by an [Object] to provide a mechanism for the
// engine to call methods on an [ExtensionInstance].
type ExtensionFunction interface {
	// PointerCall expects the return value to be written into the result [Pointer]
	// and arguments are passed as an array of [Pointer] values.
	PointerCall(instance ExtensionInstance, result MutablePointer, args PointerTo[Pointer])

	// CheckedCall is called in the engine in rare cases where it knows the number of
	// arguments in advance and can avoid the overhead of dynamic calling.
	CheckedCall(instance ExtensionInstance, args Variants) Variant

	// DynamicCall is called when the engine needs to call a method on an
	// [ExtensionInstance] from a [Script], or a [Signal] attached in the editor.
	DynamicCall(instance ExtensionInstance, args Variants) (Variant, Error)
}

// ExtensionClass is an interface that can be implemented to create a new
// class in the engine. Pass an implementation of this interface to the
// [RegisterClass] function to register it with the engine.
type ExtensionClass interface {

	// Create should instantiate the underlying object with [New] and then
	// call [SetupExtension] on it.
	Create(notify_postinitialize bool) Object

	// Parent returns the parent class for this extension class.
	Parent() Class

	// Method should return the [ExtensionFunction] for the given method.
	Method(name StringName, hash uint32) ExtensionFunction

	Virtual() bool  // Virtual class.
	Abstract() bool // Abstract class.
	Exposed() bool  // Exposed to scripts and the editor.
	Runtime() bool  // Runtime classes only run outside the editor.

	Icon() String // Icon to display in the editor for this class.
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

const (
	InitializationLevelCore    InitializationLevel = 0
	InitializationLevelServers InitializationLevel = 1
	InitializationLevelScene   InitializationLevel = 2
	InitializationLevelEditor  InitializationLevel = 3
)

var (
	onEngineInit []func(InitializationLevel)
	onEngineExit []func(InitializationLevel)
	onFirstFrame []func()
	onEveryFrame []func()
	onFinalFrame []func()
)

// OnEngineInit registers a function to be called when the engine initializes,
// this should be called early inside an init function.
func OnEngineInit(fn func(level InitializationLevel)) { onEngineInit = append(onEngineInit, fn) }

// OnEngineExit registers a function to be called when the engine exits,
// this should be called early inside an init function.
func OnEngineExit(fn func(level InitializationLevel)) { onEngineExit = append(onEngineExit, fn) }

// OnFirstFrame registers a function to be called on the first frame,
// this should be called early inside an init function.
func OnFirstFrame(fn func()) { onFirstFrame = append(onFirstFrame, fn) }

// OnEveryFrame registers a function to be called on every frame. Can
// be called at any time but cannot be removed after being added.
func OnEveryFrame(fn func()) { onEveryFrame = append(onEveryFrame, fn) }

// OnFinalFrame registers a function to be called on the final frame,
// this should be called before the engine shuts down.
func OnFinalFrame(fn func()) { onFinalFrame = append(onFinalFrame, fn) }

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

// shapeFor maps a Go type in the [Any] constraint to its [Shape].
func shapeFor[T Returnable]() Shape {
	return shapeOf(reflect.TypeFor[T]())
}

func shapeOf(rtype reflect.Type) Shape {
	switch rtype {
	case reflect.TypeFor[struct{}]():
		return empty
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
		panic("shapeOf: invalid type " + rtype.String())
	}
}

// VariantOperator is an enumeration of the supported variant operators.
type VariantOperator uint32

const (
	OpEqual         VariantOperator = iota // ==
	OpNotEqual                             // !=
	OpLess                                 // <
	OpLessEqual                            // <=
	OpGreater                              // >
	OpGreaterEqual                         // >=
	OpAdd                                  // +
	OpSubtract                             // -
	OpMultiply                             // *
	OpDivide                               // /
	OpNegate                               // -
	OpPositive                             // +
	OpModule                               // %
	OpPower                                // ^
	OpShiftLeft                            // <<
	OpShiftRight                           // >>
	OpBitwiseAnd                           // &
	OpBitwiseOr                            // |
	OpBitwiseXor                           // ^
	OpBitwiseNegate                        // ~
	OpLogicalAnd                           // &&
	OpLogicalOr                            // ||
	OpLogicalXor                           // !=
	OpLogicalNegate                        // !
	OpIn                                   // in
)

func builtinMethodShapeFor[T Any, Args any, Result Returnable]() (Shape, bool) {
	var shape = shapeFor[Result]()
	var shift = 4
	shape |= shapeFor[T]() << shift
	shift += 4
	rtype := reflect.TypeFor[Args]()
	if rtype == reflect.TypeFor[[]Variant]() {
		return shape, true
	}
	for field := range rtype.Fields() {
		shape |= shapeOf(field.Type) << shift
		shift += 4
	}
	return shape, false
}

func shapeVariants(receiver bool, length int) Shape {
	var start = 1
	if receiver {
		start = 2
	}
	var shape Shape
	for i := range length {
		shape |= ShapeVariant << ((i + start) * 4) // +2 to leave room for T and Result
	}
	return shape
}

type BuiltinMethod[T Any, Args any, Result Returnable] struct {
	_ [0]*T
	_ [0]*Args
	_ [0]*Result

	entry Pointer
	shape Shape
	vargs bool
}

type BuiltinMethodMutable[T Any, Args any, Result Returnable] struct {
	_   [0]*T
	_   [0]*Args
	_   [0]*Result
	mut struct{}

	entry Pointer
	shape Shape
	vargs bool
}

func Import[T any]() *T {
	var spec = new(T)
	OnEngineInit(func(level InitializationLevel) {
		if level == InitializationLevelCore {
			link(reflect.ValueOf(spec).Elem())
		}
	})
	return spec
}

func link(spec reflect.Value) {
	for method := range spec.Fields() {
		value := reflect.NewAt(method.Type, unsafe.Add(spec.Addr().UnsafePointer(), method.Offset))
		mb, ok := reflect.TypeAssert[linkable](value)
		if ok {
			hash, err := strconv.ParseInt(method.Tag.Get("hash"), 10, 64)
			if err == nil {
				mb.link(method.Name, hash)
			}
		} else if method.Type.Kind() == reflect.Struct {
			link(value.Elem())
		}
	}
}

type linkable interface {
	link(string, int64)
}

func (m *BuiltinMethodMutable[T, Args, Result]) link(method string, hash int64) {
	var builtin BuiltinMethod[T, Args, Result]
	builtin.link(method, hash)
	m.entry = builtin.entry
	m.shape = builtin.shape
	m.vargs = builtin.vargs
}

// BuiltinMethodByName returns a BuiltinMethodPointer that can be used to call the given builtin method on a value of type T.
func (builtin *BuiltinMethod[T, Args, Result]) link(method string, hash int64) {
	method_name := UTF8.Intern(method)
	defer Free(method_name)
	shape, vargs := builtinMethodShapeFor[T, Args, Result]()
	*builtin = BuiltinMethod[T, Args, Result]{
		entry: Pointer(gd_builtin_method(uint32(variantTypeOf[T]()), method_name.raw, hash)),
		shape: shape,
		vargs: vargs,
	}
}
