//go:build !cgo && !wasm

// Package gdunsafe provides unsafe access to the gdextension API (no protections for double-free, pointer-aliasing or use-after-free).
package gdunsafe

import (
	"unsafe"

	"graphics.gd/variant"
)

const unavailable = "gdunsafe: unavailable without cgo or wasm"

// LibraryLocation returns a string representing the location of the current extension.
func LibraryLocation() String { panic(unavailable) }

// Index returns the i'th variant, for checked calls to instance methods,
// you may not pass an index here beyond the number of arguments declared.
func (args Variants) Index(i int) Variant { panic(unavailable) }

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

func Malloc(size uintptr) MutablePointer                     { panic(unavailable) } // Malloc allocates optionally 8-byte aligned memory.
func Resize(ptr MutablePointer, size uintptr) MutablePointer { panic(unavailable) } // Resize resizes optionally 8-byte aligned memory.
func Memset(ptr MutablePointer, size uintptr, value byte)    { panic(unavailable) } // Memset sets a block of memory to a given value.

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

func (enc latin1) Decode(s String, buf []byte) int { panic(unavailable) } // Decode the string into the buffer. Returns bytes written.
func (enc latin1) String(s string) String          { panic(unavailable) } // String returns a new [String] from the given string.
func (enc latin1) Intern(s string) StringName      { panic(unavailable) } // Intern returns a [StringName] for the given string.

func (enc utf8) Decode(s String, buf []byte) int { panic(unavailable) } // Decode the string into the buffer. Returns bytes written.
func (enc utf8) String(s string) String          { panic(unavailable) } // String returns a new [String] from the given string.
func (enc utf8) Intern(s string) StringName      { panic(unavailable) } // Intern returns a [StringName] for the given string.

func (enc utf16) Decode(s String, buf []byte) int { panic(unavailable) } // Decode the string into the buffer. Returns bytes written.
func (enc utf16) String(s string) String          { panic(unavailable) } // String returns a new [String] from the given string.
func (enc utf16) Intern(s string) StringName      { panic(unavailable) } // Intern returns a [StringName] for the given string.

func (enc utf32) Decode(s String, buf []byte) int { panic(unavailable) } // Decode the string into the buffer. Returns bytes written.
func (enc utf32) String(s string) String          { panic(unavailable) } // String returns a new [String] from the given string.
func (enc utf32) Intern(s string) StringName      { panic(unavailable) } // Intern returns a [StringName] for the given string.

func (enc wide) Decode(s String, buf []byte) int { panic(unavailable) } // Decode the string into the buffer. Returns bytes written.
func (enc wide) String(s string) String          { panic(unavailable) } // String returns a new [String] from the given string.
func (enc wide) Intern(s string) StringName      { panic(unavailable) } // Intern returns a [StringName] for the given string.

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
func Utility(utility StringName, hash int64) func(result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	panic(unavailable)
}

// Constant returns T's 'name' constant of type E.
func Constant[T, E Any](name StringName) E {
	panic(unavailable)
}

// Constructor returns a function that can be used to construct a value of type T with the given
// shape and arguments.
func Constructor[T Any](n int) func(shape Shape, args unsafe.Pointer) T {
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

func (class Class) Tag() ClassTag { panic(unavailable) }

func New(name Class) Object                          { panic(unavailable) } // New [Object] of the given type.
func (obj Object) Class() Class                      { panic(unavailable) } // Type of the object.
func (obj Object) Cast(to ClassTag) Object           { panic(unavailable) } // Cast the object to the given [Type].
func (obj Object) Script(lang ScriptLanguage) Script { panic(unavailable) } // Script attached to the object.
func (obj Object) AttachScript(script Script)        { panic(unavailable) } // AttachScript to the object.
func (obj Object) ID() ObjectID                      { panic(unavailable) } // ID of the object.
func (obj Object) Free()                             { panic(unavailable) } // Free the object.

func (id ObjectID) Object() Object { panic(unavailable) } // Object associated with the ID.

// Singleton returns the singleton object with the given name.
func Singleton(name StringName) Object { panic(unavailable) }

// Method returns the [MethodPointer] for the given class and method.
func Method(class, method StringName, hash int64) MethodPointer { panic(unavailable) }

// Call the method with the given arguments.
func (obj Object) Call(method MethodPointer, args ...Variant) (Variant, Error) {
	panic(unavailable)
}

// ShapedCall calls the method with unsafely shaped arguments and result.
func (obj Object) ShapedCall(method MethodPointer, result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	panic(unavailable)
}

func (obj Object) SetupExtension(name StringName, inst ExtensionInstance) { panic(unavailable) } // SetupExtension in [ExtensionClass.Create].
func (obj Object) ExtensionInstance() ExtensionInstance                   { panic(unavailable) } // ExtensionInstance is non-nil if the object has one.

// MakeScript creates a new [Script] using the given [ExtensionScript] implementation.
func MakeScript(fn ExtensionScript) Script { panic(unavailable) }

// Call the script's method with the given arguments.
func (obj Script) Call(name StringName, args ...Variant) (Variant, Error) {
	panic(unavailable)
}

// HasMethod returns true if the script has a method with the given name.
func (obj Script) HasMethod(method StringName) bool { panic(unavailable) }

// MakePlaceholderScript creates a new [Script] that acts as a placeholder for the given script.
func MakePlaceholderScript(language ScriptLanguage, script Script, owner Object) Script {
	panic(unavailable)
}

// UpdatePlaceholder updates the placeholder script with the given array and dictionary.
func (s Script) UpdatePlaceholder(array Array, dict Dictionary) {
	panic(unavailable)
}

func Nil() Variant                                                         { panic(unavailable) } // Nil returns a nil [Variant].
func (v Variant) Copy(deep bool) Variant                                   { panic(unavailable) } // Copy returns a copy of the [Variant]. Deep copy is performed if `deep` is true.
func (v Variant) Call(method StringName, args ...Variant) (Variant, Error) { panic(unavailable) } // Call calls the method with the given arguments on the [Variant] and returns the result.
func (v Variant) Hash(depth int64) int64                                   { panic(unavailable) } // Hash returns the hash of the [Variant]. A depth > 0 returns a deep hash.
func (v Variant) Bool() bool                                               { panic(unavailable) } // Bool returns the [Variant] as a truthy [bool].
func (v Variant) UnsafeString() String                                     { panic(unavailable) } // Text returns the [Variant] as a [String].
func (v Variant) Type() variant.Type                                       { panic(unavailable) } // Type returns the type of the [Variant].
func (v Variant) ObjectID() ObjectID                                       { panic(unavailable) } // ObjectID returns the [ObjectID] inside the [Variant].
func (v Variant) Lookup(key Variant) (Variant, bool)                       { panic(unavailable) } // Lookup returns the value for the given key in the [Variant].
func (v Variant) Index(idx int) (Variant, bool, Error)                     { panic(unavailable) } // Index returns the value at the given index in the [Variant].
func (v Variant) Field(field StringName) (Variant, bool)                   { panic(unavailable) } // Field returns the value for the given field in the [Variant].
func (v Variant) Insert(key, val Variant) bool                             { panic(unavailable) } // Insert inserts the given key-value pair into the [Variant].
func (v Variant) SetIndex(idx int, val Variant) (bool, Error)              { panic(unavailable) } // SetIndex sets the value at the given index in the [Variant].
func (v Variant) SetField(field StringName, value Variant) bool            { panic(unavailable) } // SetField sets the value for the given field in the [Variant].
func (v Variant) Has(index Variant) bool                                   { panic(unavailable) } // Has returns true if the [Variant] has the given key.
func (v Variant) HasMethod(method StringName) bool                         { panic(unavailable) } // HasMethod returns true if the [Variant] has the given method.
func (v Variant) Free()                                                    { panic(unavailable) } // Free releases any resources associated by the [Variant].

// Evaluate evaluates the [VariantOperator] on the given [Variant] operands and returns the result.
func (op VariantOperator) Evaluate(a, b Variant) (Variant, bool) {
	panic(unavailable)
}

// VariantInto converts a [Variant] to the given type T.
func VariantInto[T Any](v Variant) T {
	panic(unavailable)
}

// VariantFrom converts a native value of type T to a [Variant].
func VariantFrom[T Any](native T) Variant {
	panic(unavailable)
}

// PointerIntoVariant returns a pointer to the underlying T value inside the [Variant].
func PointerIntoVariant[T Any](v Variant) PointerTo[T] { panic(unavailable) }

func (d Dictionary) Lookup(key Variant) Variant { panic(unavailable) } // Lookup returns the value associated with the given key in the dictionary.
func (d Dictionary) Insert(key, val Variant)    { panic(unavailable) } // Insert the value associated with the given key into the dictionary.
func (dict Dictionary) SetType(key, val Type)   { panic(unavailable) } // SetType sets the type of the key and value in the dictionary.

func (ref RefCounted) Get() Object    { panic(unavailable) } // Get returns the object held by the reference.
func (ref RefCounted) Set(obj Object) { panic(unavailable) } // Set sets the object held by the reference.

func EditorDocumentation(xml string) { panic(unavailable) } // EditorDocumentation adds documentation to the editor.
func EnableEditorPlugin(name Class)  { panic(unavailable) } // EnableEditorPlugin with the given name.
func RemoveEditorPlugin(name Class)  { panic(unavailable) } // RemoveEditorPlugin with the given name.

func MakePropertyList(n int64) PropertyList { panic(unavailable) } // MakePropertyList creates a new property list with the given size.
func (p PropertyList) Push(Property)        { panic(unavailable) } // Push adds a property to the list.
func (p PropertyList) Free()                { panic(unavailable) } // Free frees the property list.

func MakeMethodList(n int64) MethodList { panic(unavailable) } // MakeMethodList creates a new method list with the given size.
// Push adds a method to the list.
func (m MethodList) Push(name StringName, call ExtensionFunction, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count int64, defaults unsafe.Pointer) {
	panic(unavailable)
}
func (m MethodList) Free() { panic(unavailable) } // Free frees the method list.

// RegisterClass registers a new extension class with the given name, parent, and ID.
func RegisterClass(class string, id ExtensionClass) Class {
	panic(unavailable)
}

// RegisterMethods registers new methods for the given class.
func (class Class) RegisterMethods(methods MethodList) {
	panic(unavailable)
}

// RegisterConstant registers a new constant for the given class and enum.
func (class Class) RegisterConstant(enum, name StringName, value int64, bitfield bool) {
	panic(unavailable)
}

// RegisterProperty registers a new property for the given class.
func (class Class) RegisterProperty(property Property, setter, getter StringName) {
	panic(unavailable)
}

// RegisterPropertyIndexed registers a new indexed property for the given class.
func (class Class) RegisterPropertyIndexed(property Property, setter, getter StringName, index int) {
	panic(unavailable)
}

// RegisterPropertyGroup registers a new property group for the given class.
func (class Class) RegisterPropertyGroup(group, prefix String) {
	panic(unavailable)
}

// RegisterPropertySubgroup registers a new property subgroup for the given class.
func (class Class) RegisterPropertySubgroup(subgroup, prefix String) {
	panic(unavailable)
}

// RegisterSignal registers a new signal for the given class.
func (class Class) RegisterSignal(signal StringName, args PropertyList) {
	panic(unavailable)
}

// Free deregisters the class and releases any resources associated with it.
func (class Class) Free() { panic(unavailable) }

// Iterator returns an [Iterator] for the given variant.
func (v Variant) Iterator() (Iterator, Error) { panic(unavailable) }

// Next returns the next value from the iterator, or false if there are no more values.
func (iter *Iterator) Next() (bool, Error) {
	panic(unavailable)
}

// Value returns the current value from the iterator.
func (iter Iterator) Value() (Variant, Error) {
	panic(unavailable)
}
