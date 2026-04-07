//go:build !cgo && !wasm

// Package gdunsafe provides unsafe access to the gdextension API (no protections for double-free, pointer-aliasing or use-after-free).
package gdunsafe

import (
	"unsafe"

	"graphics.gd/variant"
)

type (
	PackedArray[T Packable] [2]uint64

	MethodForClass uintptr

	PropertyList uintptr
	MethodList   uintptr
)

type (
	Callable        [2]uint64
	CallableID      uintptr
	VariantOperator = uint32
)

const unavailable = "gdunsafe: unavailable without cgo or wasm"

// LibraryLocation returns a string representing the location of the current extension.
func LibraryLocation() String { panic(unavailable) }

// Index returns the i'th variant, for checked calls to instance methods,
// you may not pass an index here beyond the number of arguments declared.
func (args Variants) Index(i int) Variant { panic(unavailable) }

// Array is a sequentially-indexed list of variant values.
type Array struct{ _ [0]*Array }

func (array Array) Index(index int) Variant           { panic(unavailable) } // Index returns the variant value at the given index.
func (array Array) SetIndex(index int, value Variant) { panic(unavailable) } // SetIndex sets the variant value at the given index.
func (array Array) SetType(t Type)                    { panic(unavailable) } // SetType configures the element type for the array.

// Size returns the size of the type in bytes.
func (t Type) Size() uintptr { panic(unavailable) }

func Version() String          { panic(unavailable) } // Version of the engine.
func VersionMajor() uint32     { panic(unavailable) } // Major version of the engine.
func VersionMinor() uint32     { panic(unavailable) } // Minor version of the engine.
func VersionPatch() uint32     { panic(unavailable) } // Patch version of the engine.
func VersionHexed() uint32     { panic(unavailable) } // Hexed version of the engine.
func VersionState() String     { panic(unavailable) } // State of the engine (e.g. "stable", "beta", "alpha").
func VersionBuild() String     { panic(unavailable) } // Build type.
func VersionCommit() String    { panic(unavailable) } // Commit hash of the build.
func VersionTimestamp() uint64 { panic(unavailable) } // Timestamp of the engine build.

func Malloc(size uintptr, align bool) MutablePointer                     { panic(unavailable) } // Malloc allocates optionally 8-byte aligned memory.
func Resize(ptr MutablePointer, size uintptr, align bool) MutablePointer { panic(unavailable) } // Resize resizes optionally 8-byte aligned memory.

type Pointer struct{ _ [0]*Pointer }

type MutablePointer struct{ _ [0]*MutablePointer }

func (ptr MutablePointer) Free() { panic(unavailable) } // Free releases memory allocated by [Malloc] or [Resize].

func (ptr PointerTo[T]) Get() T         { panic(unavailable) }
func (ptr MutablePointerTo[T]) Set(v T) { panic(unavailable) }

// String operations

func (s String) Index(idx int) rune                     { panic(unavailable) } // Index returns the rune at the given index.
func (s String) SetIndex(idx int, char rune)            { panic(unavailable) } // SetIndex sets the rune at the given index.
func (s String) Resize(size int) String                 { panic(unavailable) } // Resize resizes the string to the given size.
func (s String) Pointer() PointerTo[rune]               { panic(unavailable) } // Pointer to the first rune of the string.
func (s String) MutablePointer() MutablePointerTo[rune] { panic(unavailable) } // MutablePointer to the first rune in the string.

func (s *String) Append(other String) { panic(unavailable) } // Append appends another string to this string.
func (s *String) AppendRune(ch rune)  { panic(unavailable) } // AppendRune appends the given rune to this string.

func (enc StringEncoding) Decode(s String, buf []byte) int { panic(unavailable) } // Decode the string into the buffer. Returns bytes written.
func (enc StringEncoding) String(s string) String          { panic(unavailable) } // String returns a new [String] from the given string.
func (enc StringEncoding) Intern(s string) StringName      { panic(unavailable) } // Intern returns a [StringName] for the given string.

// Log a warning or error message.
func Log(level LogLevel, text, code, fn, file string, line int32, notify_editor bool) {
	panic(unavailable)
}

func (p PackedArray[T]) Index(idx int64) T         { panic(unavailable) } // Index returns the element at the given index.
func (p PackedArray[T]) SetIndex(idx int64, val T) { panic(unavailable) } // SetIndex sets the element at the given index.

func (p PackedArray[T]) Pointer() PointerTo[T]               { panic(unavailable) } // Pointer to the first element in the array.
func (p PackedArray[T]) MutablePointer() MutablePointerTo[T] { panic(unavailable) } // MutablePointer to the first element in the array.

// MakeVariant creates a new [Variant] of the given [variant.Type] using the given arguments
// as constructor arguments. Returns an error if the argument combination is invalid.
func MakeVariant(vtype variant.Type, args ...Variant) (Variant, Error) { panic(unavailable) }

// Call invokes a static method on T with the given arguments. Returns an error if the method
// does not exist or the arguments are invalid.
func Call[T Any](method StringName, args ...Variant) (Variant, Error) {
	panic(unavailable)
}

// Convertable returns true if values of type A can be converted to type B. If strict is true,
// the conversion must be exact (no data loss).
func Convertable[A, B Any](strict bool) bool { panic(unavailable) }

// Utility returns the utility function with the given name and hash as a function that can be
// shape called.
func Utility(utility StringName, hash int64) func(result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}

// Constant returns T's 'name' constant of type E.
func Constant[T, E Any](name StringName) E {
	panic(unavailable)
}

// Constructor returns a function that can be used to construct a value of type T with the given
// shape and arguments.
func Constructor[T Any](n int64) func(shape uint64, args unsafe.Pointer) T {
	panic(unavailable)
}

// Evaluator returns a function that can be used to evaluate an operation on two values of types A
// and B, returning a value of type R.
func Evaluator[A, B, R Any](op VariantOperator) func(a A, b B) R {
	panic(unavailable)
}

// Setter returns a function that can be used to set the value of the given field on a value of type T.
func Setter[T Any, E Any](field StringName) func(v T, field E) {
	panic(unavailable)
}

// Getter returns a function that can be used to get the value of the given field from a value of type T.
func Getter[T Any, E Any](field StringName) func(v T) E {
	panic(unavailable)
}

// PropertyExists returns true if the given property exists on the given type.
func PropertyExists[T Any](property StringName) bool {
	panic(unavailable)
}

// BuiltinMethod returns a function that can be used to call the given builtin method on a value of type T.
func BuiltinMethod[T Any](method StringName, hash int64) func(self T, ret unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	panic(unavailable)
}

// SetIndex sets the value of the given index on the given value of type T.
func SetIndex[T, V Any](self T, index int64, value V) {
	panic(unavailable)
}

// Index returns the value of the given index on the given value of type T.
func Index[T, V Any](self T, index int64) V {
	panic(unavailable)
}

// Insert inserts the given value at the given index on the given value of type T.
func Insert[T Any](self T, index, value Variant) {
	panic(unavailable)
}

// Lookup returns the value of the given key on the given value of type T.
func Lookup[T Any](self T, key Variant) Variant {
	panic(unavailable)
}

// Free releases any resources associated with the given value of type T.
func Free[T Any](val T) { panic(unavailable) }

// MakeCallable returns a [Callable] backed by the given [ExtensionCallable] implementation.
// It can optionally be bound to an object, when the object is freed, the callable is freed.
func MakeCallable(impl ExtensionCallable, obj ObjectID) Callable { panic(unavailable) }

func New(name StringName) Object                     { panic(unavailable) } // New [Object] of the given type.
func (obj Object) Type() Type                        { panic(unavailable) } // Type of the object.
func (obj Object) Cast(to Type) Object               { panic(unavailable) } // Cast the object to the given [Type].
func (obj Object) Script(lang ScriptLanguage) Script { panic(unavailable) } // Script attached to the object.
func (obj Object) AttachScript(script Script)        { panic(unavailable) } // AttachScript to the object.
func (obj Object) ID() ObjectID                      { panic(unavailable) } // ID of the object.
func (obj Object) Free()                             { panic(unavailable) } // Free the object.

func (id ObjectID) Object() Object { panic(unavailable) } // Object associated with the ID.

// Singleton returns the singleton object with the given name.
func Singleton(name StringName) Object { panic(unavailable) }

// Object method calls
func MethodLookup(class, method StringName, hash int64) MethodForClass { panic(unavailable) }
func (obj Object) Call(method MethodForClass, args ...Variant) (Variant, Error) {
	panic(unavailable)
}
func (obj Object) ShapedCall(fn MethodForClass, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}

// Extension instance management

func (obj Object) SetupExtension(name StringName, inst ExtensionInstance) { panic(unavailable) }
func (obj Object) FetchExtension() ExtensionInstance                      { panic(unavailable) }
func (obj Object) CloseExtension()                                        { panic(unavailable) }

// Script instance management

func MakeScript(fn ExtensionScript) Script { panic(unavailable) }
func (obj Script) Call(name StringName, args ...Variant) (Variant, Error) {
	panic(unavailable)
}

func (obj Script) HasMethod(method StringName) bool { panic(unavailable) }

func MakePlaceholderScript(language ScriptLanguage, script Script, owner Object) Script {
	panic(unavailable)
}
func (s Script) UpdatePlaceholder(array Array, dict Dictionary) {
	panic(unavailable)
}

// Variant operations

func ZeroVariant() Variant               { panic(unavailable) }
func (v Variant) Copy(deep bool) Variant { panic(unavailable) }
func (v Variant) VariantCall(method StringName, args ...Variant) (Variant, Error) {
	panic(unavailable)
}
func (v Variant) Hash(depth int64) int64 { panic(unavailable) }
func (v Variant) Bool() bool             { panic(unavailable) }
func (v Variant) Text() String           { panic(unavailable) }
func (v Variant) Type() variant.Type     { panic(unavailable) }

// Variant get/set/has

func (v Variant) ObjectID() ObjectID                                       { panic(unavailable) }
func (v Variant) GetIndex(key Variant) (Variant, bool)                     { panic(unavailable) }
func (v Variant) GetArray(idx int64) (Variant, bool, Error)                { panic(unavailable) }
func (v Variant) GetField(field StringName) (Variant, bool)                { panic(unavailable) }
func (v Variant) SetIndex(key, val Variant) bool                           { panic(unavailable) }
func (v Variant) SetArray(idx int64, val Variant, err unsafe.Pointer) bool { panic(unavailable) }
func (v Variant) SetField(field StringName, value Variant) bool            { panic(unavailable) }
func (v Variant) HasIndex(index Variant) bool                              { panic(unavailable) }
func (v Variant) HasMethod(method StringName) bool                         { panic(unavailable) }
func (v Variant) Free()                                                    { panic(unavailable) }

func (op VariantOperator) Evaluate(a, b Variant) (Variant, bool) {
	panic(unavailable)
}

func VariantInto[T Any](v Variant) T {
	panic(unavailable)
}

func VariantFrom[T Any](native T) Variant {
	panic(unavailable)
}

func PointerIntoVariant[T Any](v Variant) PointerTo[T] { panic(unavailable) }

// Dictionary operations

func (d Dictionary) Index(key Variant) Variant { panic(unavailable) }
func (d Dictionary) SetIndex(key, val Variant) { panic(unavailable) }

func (dict Dictionary) SetType(key, val Type) {
	panic(unavailable)
}

// RefCounted operations

func RefGet(ref Pointer) Object      { panic(unavailable) }
func RefSet(ref Pointer, obj Object) { panic(unavailable) }

// Editor operations

func EditorAddDocumentation(xml string) { panic(unavailable) }
func EditorAddPlugin(name StringName)   { panic(unavailable) }
func EditorEndPlugin(name StringName)   { panic(unavailable) }

// PropertyList operations

func MakePropertyList(n int64) PropertyList { panic(unavailable) }
func (p PropertyList) Push(Property) {
	panic(unavailable)
}
func (p PropertyList) Free() { panic(unavailable) }

// MethodList operations

func MakeMethodList(n int64) MethodList { panic(unavailable) }
func (m MethodList) Push(name StringName, call ExtensionFunction, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count int64, defaults unsafe.Pointer) {
	panic(unavailable)
}
func (m MethodList) Free() { panic(unavailable) }

// ClassDB registration

func RegisterClass(class, parent StringName, id ExtensionClass, virtual, abstract, exposed, runtime bool, icon String) {
	panic(unavailable)
}
func RegisterMethods(class StringName, methods MethodList)                      { panic(unavailable) }
func RegisterConstant(class, enum, name StringName, value int64, bitfield bool) { panic(unavailable) }
func RegisterProperty(class StringName, property PropertyList, setter, getter StringName) {
	panic(unavailable)
}
func RegisterPropertyIndexed(class StringName, property PropertyList, setter, getter StringName, index int) {
	panic(unavailable)
}
func RegisterPropertyGroup(class StringName, group, prefix String)       { panic(unavailable) }
func RegisterPropertySubgroup(class StringName, subgroup, prefix String) { panic(unavailable) }
func RegisterSignal(class, signal StringName, args PropertyList)         { panic(unavailable) }
func RegisterRemoval(class StringName)                                   { panic(unavailable) }

// ClassDB sub-API operations

/*func FileAccessWrite(file Object, buf []byte)     { panic(unavailable) }
func FileAccessRead(file Object, buf []byte) int  { panic(unavailable) }
func ImageUnsafe(img Object) Pointer              { panic(unavailable) }
func ImageAccess(img Object, offset int64) byte   { panic(unavailable) }
func XMLParserLoad(parser Object, buf []byte) int { panic(unavailable) }
func WorkerThreadPoolAddTask(pool Object, task Pointer, priority bool, description String) {
	panic(unavailable)
}
func WorkerThreadPoolAddGroupTask(pool Object, task Pointer, elements, arg int32, priority bool, description String) {
	panic(unavailable)
	}*/

// Iterator operations

func (v Variant) Iterator() (Iterator, Error) { panic(unavailable) }

func (iter *Iterator) Next() (bool, Error) {
	panic(unavailable)
}

func (iter *Iterator) Value() (Variant, Error) {
	panic(unavailable)
}
