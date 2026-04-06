//go:build !cgo && !wasm

// Package gdunsafe provides direct 'unsafe' access to the graphics/game development engine's extension API.
package gdunsafe

import (
	"unsafe"

	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type (
	String     uintptr
	StringName uintptr
	Array      uintptr
	Dictionary uintptr
	Pointer    uintptr

	PackedArray[T byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String] [2]uint64

	VariantType uint32

	Object              uintptr
	ObjectType          uintptr
	MethodForClass      uintptr
	ScriptInstance      uintptr
	ExtensionInstanceID uintptr
	ExtensionClassID    uintptr
	ExtensionBindingID  uintptr
	FunctionID          uintptr
	PropertyList        uintptr
	MethodList          uintptr
)

type (
	Callable        [2]uint64
	CallableID      uintptr
	VariantOperator = uint32
)

const unavailable = "gdunsafe: unavailable without cgo or wasm build tags"

// VariadicVariants

func (args VariadicVariants) Index(i int) Variant { panic(unavailable) }

// Array operations

func (array Array) Set(index Int, value Variant) { panic(unavailable) }
func (array Array) Get(index Int) Variant        { panic(unavailable) }

// Version information

func VersionMajor() uint32     { panic(unavailable) }
func VersionMinor() uint32     { panic(unavailable) }
func VersionPatch() uint32     { panic(unavailable) }
func VersionHex() uint32       { panic(unavailable) }
func VersionStatus() String    { panic(unavailable) }
func VersionBuild() String     { panic(unavailable) }
func VersionHash() String      { panic(unavailable) }
func VersionTimestamp() uint64 { panic(unavailable) }
func VersionString() String    { panic(unavailable) }
func LibraryLocation() String  { panic(unavailable) }

// Memory operations

func Malloc(size Int) Pointer              { panic(unavailable) }
func Sizeof(name StringName) Int           { panic(unavailable) }
func Resize(ptr Pointer, size Int) Pointer { panic(unavailable) }
func Clear(ptr Pointer, size Int)          { panic(unavailable) }

// Pointer operations

func (ptr Pointer) Byte() byte               { panic(unavailable) }
func (ptr Pointer) Uint16() uint16           { panic(unavailable) }
func (ptr Pointer) Uint32() uint32           { panic(unavailable) }
func (ptr Pointer) Uint64() uint64           { panic(unavailable) }
func (ptr Pointer) SetByte(v byte)           { panic(unavailable) }
func (ptr Pointer) SetUint16(v uint16)       { panic(unavailable) }
func (ptr Pointer) SetUint32(v uint32)       { panic(unavailable) }
func (ptr Pointer) SetUint64(v uint64)       { panic(unavailable) }
func (ptr Pointer) SetBits128(val [2]uint64) { panic(unavailable) }
func (ptr Pointer) SetBits256(val [4]uint64) { panic(unavailable) }
func (ptr Pointer) SetBits512(val [8]uint64) { panic(unavailable) }
func (ptr Pointer) Free()                    { panic(unavailable) }

// String operations

func (s String) Access(idx Int) int32                      { panic(unavailable) }
func (s String) Resize(size Int) String                    { panic(unavailable) }
func (s String) UnsafePtr() Pointer                        { panic(unavailable) }
func (s String) Append(other String) String                { panic(unavailable) }
func (s String) AppendRune(ch int32) String                { panic(unavailable) }
func (s String) Encode(enc StringEncoding, buf []byte) Int { panic(unavailable) }
func (enc StringEncoding) String(s string) String          { panic(unavailable) }
func (enc StringEncoding) Intern(s string) StringName      { panic(unavailable) }

// Log

func Log(level LogLevel, text, code, fn, file string, line int32, notify_editor bool) {
	panic(unavailable)
}

// PackedArray operations

func (ptr PointerTo[T]) Get() T                      { panic(unavailable) }
func (ptr PointerTo[T]) Set(v T)                     { panic(unavailable) }
func (p PackedArray[T]) Access(idx Int) PointerTo[T] { panic(unavailable) }
func (p PackedArray[T]) Modify(idx Int) PointerTo[T] { panic(unavailable) }

// VariantType operations

func (t VariantType) Name() String                              { panic(unavailable) }
func (t VariantType) Make(args ...Variant) (Variant, CallError) { panic(unavailable) }
func (t VariantType) StaticCall(method StringName, args ...Variant) (Variant, CallError) {
	panic(unavailable)
}
func (t VariantType) Convertable(to VariantType, strict bool) bool { panic(unavailable) }

// Builtin and utility functions

func BuiltinName(utility StringName, hash int64) FunctionID { panic(unavailable) }
func BuiltinCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}

// Variant type setup

func VariantTypeSetupArray(array Array, vtype VariantType, className StringName, script Variant) {
	panic(unavailable)
}
func VariantTypeSetupDictionary(dict Dictionary, keyType VariantType, keyClassName StringName, keyScript Variant, valType VariantType, valClassName StringName, valScript Variant) {
	panic(unavailable)
}
func VariantTypeFetchConstant(vtype VariantType, constant StringName, result unsafe.Pointer) {
	panic(unavailable)
}
func VariantTypeConstructor(vtype VariantType, n Int) FunctionID           { panic(unavailable) }
func VariantTypeEvaluator(op VariantOperator, a, b VariantType) FunctionID { panic(unavailable) }
func VariantTypeSetter(vtype VariantType, property StringName) FunctionID  { panic(unavailable) }
func VariantTypeGetter(vtype VariantType, property StringName) FunctionID  { panic(unavailable) }
func VariantTypeHasProperty(vtype VariantType, property StringName) bool   { panic(unavailable) }
func VariantTypeMethod(vtype VariantType, method StringName, hash int64) FunctionID {
	panic(unavailable)
}

// Variant type unsafe operations

func VariantTypeUnsafeCall(self unsafe.Pointer, fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantTypeUnsafeMake(constructor FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantTypeUnsafeFree(vtype VariantType, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}

// Callable operations

func MakeCallable(impl CallableImplementation, obj ObjectID) Callable { panic(unavailable) }

// Object operations

func MakeObject(name StringName) Object        { panic(unavailable) }
func (obj Object) Name() StringName            { panic(unavailable) }
func ObjectTypeTag(name StringName) ObjectType { panic(unavailable) }
func (obj Object) Cast(to ObjectType) Object   { panic(unavailable) }
func (id ObjectID) Lookup() Object             { panic(unavailable) }
func ObjectGlobal(name StringName) Object      { panic(unavailable) }
func (obj Object) ID() ObjectID                { panic(unavailable) }
func ObjectIDInsideVariant(v Variant) ObjectID { panic(unavailable) }
func (obj Object) Free()                       { panic(unavailable) }

// Object method calls

func MethodLookup(class, method StringName, hash int64) MethodForClass { panic(unavailable) }
func (obj Object) Call(method MethodForClass, args ...Variant) (Variant, CallError) {
	panic(unavailable)
}
func (obj Object) ShapedCall(fn MethodForClass, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}

// Extension instance management

func (obj Object) ExtensionSetup(name StringName, inst ExtensionInstanceID) { panic(unavailable) }
func (obj Object) ExtensionFetch() ExtensionInstanceID                      { panic(unavailable) }
func (obj Object) ExtensionClose()                                          { panic(unavailable) }

// Script instance management

func ScriptMake(fn ExtensionInstanceID) ScriptInstance { panic(unavailable) }
func (obj Object) ScriptCall(name StringName, args ...Variant) (Variant, CallError) {
	panic(unavailable)
}
func (obj Object) ScriptSetup(script ScriptInstance)          { panic(unavailable) }
func (obj Object) ScriptFetch(language Object) ScriptInstance { panic(unavailable) }
func (obj Object) ScriptDefinesMethod(method StringName) bool { panic(unavailable) }
func ScriptPropertyStateAdd(fn FunctionID, arg Pointer, name StringName, state Variant) {
	panic(unavailable)
}
func ScriptPlaceholderCreate(language, script, owner Object) ScriptInstance { panic(unavailable) }
func ScriptPlaceholderUpdate(script ScriptInstance, array Array, dict Dictionary) {
	panic(unavailable)
}

// Variant operations

func ZeroVariant() Variant      { panic(unavailable) }
func (v Variant) Copy() Variant { panic(unavailable) }
func (v Variant) VariantCall(method StringName, args ...Variant) (Variant, CallError) {
	panic(unavailable)
}
func VariantEval(op VariantOperator, a, b Variant) (Variant, bool) { panic(unavailable) }
func (v Variant) Hash() Int                                        { panic(unavailable) }
func (v Variant) Bool() bool                                       { panic(unavailable) }
func (v Variant) Text() String                                     { panic(unavailable) }
func (v Variant) Type() VariantType                                { panic(unavailable) }
func (v Variant) DeepCopy() Variant                                { panic(unavailable) }
func (v Variant) DeepHash(recursion Int) Int                       { panic(unavailable) }

// Variant get/set/has

func (v Variant) GetIndex(key Variant) (Variant, bool)                   { panic(unavailable) }
func (v Variant) GetArray(idx Int) (Variant, bool, CallError)            { panic(unavailable) }
func (v Variant) GetField(field StringName) (Variant, bool)              { panic(unavailable) }
func (v Variant) SetIndex(key, val Variant) bool                         { panic(unavailable) }
func (v Variant) SetArray(idx Int, val Variant, err unsafe.Pointer) bool { panic(unavailable) }
func (v Variant) SetField(field StringName, value Variant) bool          { panic(unavailable) }
func (v Variant) HasIndex(index Variant) bool                            { panic(unavailable) }
func (v Variant) HasMethod(method StringName) bool                       { panic(unavailable) }

// Unsafe variant operations

func VariantUnsafeCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantUnsafeEval(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func (v Variant) UnsafeFree() { panic(unavailable) }
func VariantUnsafeMakeNative(vtype VariantType, v Variant, shape uint64, result unsafe.Pointer) {
	panic(unavailable)
}
func VariantUnsafeFromNative(vtype VariantType, shape uint64, args unsafe.Pointer) Variant {
	panic(unavailable)
}
func VariantUnsafeInternalPointer(vtype VariantType, v Variant) Pointer { panic(unavailable) }
func VariantUnsafeSetField(setter FunctionID, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantUnsafeSetArray(vtype VariantType, idx Int, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantUnsafeSetIndex(vtype VariantType, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantUnsafeGetField(getter FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantUnsafeGetArray(vtype VariantType, idx Int, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}
func VariantUnsafeGetIndex(vtype VariantType, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	panic(unavailable)
}

// Dictionary operations

func (d Dictionary) Access(key Variant) Variant { panic(unavailable) }
func (d Dictionary) Modify(key, val Variant)    { panic(unavailable) }

// RefCounted operations

func RefGet(ref Pointer) Object      { panic(unavailable) }
func RefSet(ref Pointer, obj Object) { panic(unavailable) }

// Editor operations

func EditorAddDocumentation(xml string) { panic(unavailable) }
func EditorAddPlugin(name StringName)   { panic(unavailable) }
func EditorEndPlugin(name StringName)   { panic(unavailable) }

// PropertyList operations

func MakePropertyList(n Int) PropertyList { panic(unavailable) }
func (p PropertyList) Push(vtype VariantType, name StringName, className StringName, hint uint32, hintString String, usage uint32, meta uint32) {
	panic(unavailable)
}
func (p PropertyList) Free()                     { panic(unavailable) }
func (p PropertyList) InfoType() VariantType     { panic(unavailable) }
func (p PropertyList) InfoName() StringName      { panic(unavailable) }
func (p PropertyList) InfoClassName() StringName { panic(unavailable) }
func (p PropertyList) InfoHint() uint32          { panic(unavailable) }
func (p PropertyList) InfoHintString() String    { panic(unavailable) }
func (p PropertyList) InfoUsage() uint32         { panic(unavailable) }

// MethodList operations

func MakeMethodList(n Int) MethodList { panic(unavailable) }
func (m MethodList) Push(name StringName, call FunctionID, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count Int, defaults unsafe.Pointer) {
	panic(unavailable)
}
func (m MethodList) Free() { panic(unavailable) }

// ClassDB sub-API operations

func FileAccessWrite(file Object, buf []byte)     { panic(unavailable) }
func FileAccessRead(file Object, buf []byte) int  { panic(unavailable) }
func ImageUnsafe(img Object) Pointer              { panic(unavailable) }
func ImageAccess(img Object, offset Int) byte     { panic(unavailable) }
func XMLParserLoad(parser Object, buf []byte) int { panic(unavailable) }
func WorkerThreadPoolAddTask(pool Object, task Pointer, priority bool, description String) {
	panic(unavailable)
}
func WorkerThreadPoolAddGroupTask(pool Object, task Pointer, elements, arg int32, priority bool, description String) {
	panic(unavailable)
}

// Iterator operations

func (v Variant) IteratorMake(result unsafe.Pointer, err unsafe.Pointer) { panic(unavailable) }
func (v Variant) IteratorNext(iter unsafe.Pointer, err unsafe.Pointer) bool {
	panic(unavailable)
}
func (v Variant) IteratorLoad(iter Variant, result unsafe.Pointer, err unsafe.Pointer) {
	panic(unavailable)
}
