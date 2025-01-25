// Package ClassDB provides methods for working with ClassDB object instances.
package ClassDB

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Provides access to metadata stored for every available class.
*/
var self [1]gdclass.ClassDB
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ClassDB)
	self = *(*[1]gdclass.ClassDB)(unsafe.Pointer(&obj))
}

/*
Returns the names of all the classes available.
*/
func GetClassList() []string {
	once.Do(singleton)
	return []string(class(self).GetClassList().Strings())
}

/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
func GetInheritersFromClass(class_ string) []string {
	once.Do(singleton)
	return []string(class(self).GetInheritersFromClass(gd.NewStringName(class_)).Strings())
}

/*
Returns the parent class of [param class].
*/
func GetParentClass(class_ string) string {
	once.Do(singleton)
	return string(class(self).GetParentClass(gd.NewStringName(class_)).String())
}

/*
Returns whether the specified [param class] is available or not.
*/
func ClassExists(class_ string) bool {
	once.Do(singleton)
	return bool(class(self).ClassExists(gd.NewStringName(class_)))
}

/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
func IsParentClass(class_ string, inherits string) bool {
	once.Do(singleton)
	return bool(class(self).IsParentClass(gd.NewStringName(class_), gd.NewStringName(inherits)))
}

/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
func CanInstantiate(class_ string) bool {
	once.Do(singleton)
	return bool(class(self).CanInstantiate(gd.NewStringName(class_)))
}

/*
Creates an instance of [param class].
*/
func Instantiate(class_ string) any {
	once.Do(singleton)
	return any(class(self).Instantiate(gd.NewStringName(class_)).Interface())
}

/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
func ClassHasSignal(class_ string, signal string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasSignal(gd.NewStringName(class_), gd.NewStringName(signal)))
}

/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
func ClassGetSignal(class_ string, signal string) map[any]any {
	once.Do(singleton)
	return map[any]any(gd.DictionaryAs[any, any](class(self).ClassGetSignal(gd.NewStringName(class_), gd.NewStringName(signal))))
}

/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
func ClassGetSignalList(class_ string) []map[any]any {
	once.Do(singleton)
	return []map[any]any(gd.ArrayAs[[]map[any]any](gd.InternalArray(class(self).ClassGetSignalList(gd.NewStringName(class_), false))))
}

/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetPropertyList(class_ string) []map[any]any {
	once.Do(singleton)
	return []map[any]any(gd.ArrayAs[[]map[any]any](gd.InternalArray(class(self).ClassGetPropertyList(gd.NewStringName(class_), false))))
}

/*
Returns the value of [param property] of [param object] or its ancestry.
*/
func ClassGetProperty(obj Object.Instance, property string) any {
	once.Do(singleton)
	return any(class(self).ClassGetProperty(obj, gd.NewStringName(property)).Interface())
}

/*
Sets [param property] value of [param object] to [param value].
*/
func ClassSetProperty(obj Object.Instance, property string, value any) error {
	once.Do(singleton)
	return error(gd.ToError(class(self).ClassSetProperty(obj, gd.NewStringName(property), gd.NewVariant(value))))
}

/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
func ClassGetPropertyDefaultValue(class_ string, property string) any {
	once.Do(singleton)
	return any(class(self).ClassGetPropertyDefaultValue(gd.NewStringName(class_), gd.NewStringName(property)).Interface())
}

/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
func ClassHasMethod(class_ string, method string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasMethod(gd.NewStringName(class_), gd.NewStringName(method), false))
}

/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetMethodArgumentCount(class_ string, method string) int {
	once.Do(singleton)
	return int(int(class(self).ClassGetMethodArgumentCount(gd.NewStringName(class_), gd.NewStringName(method), false)))
}

/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
func ClassGetMethodList(class_ string) []map[any]any {
	once.Do(singleton)
	return []map[any]any(gd.ArrayAs[[]map[any]any](gd.InternalArray(class(self).ClassGetMethodList(gd.NewStringName(class_), false))))
}

/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
func ClassGetIntegerConstantList(class_ string) []string {
	once.Do(singleton)
	return []string(class(self).ClassGetIntegerConstantList(gd.NewStringName(class_), false).Strings())
}

/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
func ClassHasIntegerConstant(class_ string, name string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasIntegerConstant(gd.NewStringName(class_), gd.NewStringName(name)))
}

/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
func ClassGetIntegerConstant(class_ string, name string) int {
	once.Do(singleton)
	return int(int(class(self).ClassGetIntegerConstant(gd.NewStringName(class_), gd.NewStringName(name))))
}

/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
func ClassHasEnum(class_ string, name string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasEnum(gd.NewStringName(class_), gd.NewStringName(name), false))
}

/*
Returns an array with all the enums of [param class] or its ancestry.
*/
func ClassGetEnumList(class_ string) []string {
	once.Do(singleton)
	return []string(class(self).ClassGetEnumList(gd.NewStringName(class_), false).Strings())
}

/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
func ClassGetEnumConstants(class_ string, enum string) []string {
	once.Do(singleton)
	return []string(class(self).ClassGetEnumConstants(gd.NewStringName(class_), gd.NewStringName(enum), false).Strings())
}

/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
func ClassGetIntegerConstantEnum(class_ string, name string) string {
	once.Do(singleton)
	return string(class(self).ClassGetIntegerConstantEnum(gd.NewStringName(class_), gd.NewStringName(name), false).String())
}

/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
func IsClassEnumBitfield(class_ string, enum string) bool {
	once.Do(singleton)
	return bool(class(self).IsClassEnumBitfield(gd.NewStringName(class_), gd.NewStringName(enum), false))
}

/*
Returns whether this [param class] is enabled or not.
*/
func IsClassEnabled(class_ string) bool {
	once.Do(singleton)
	return bool(class(self).IsClassEnabled(gd.NewStringName(class_)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ClassDB

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns the names of all the classes available.
*/
//go:nosplit
func (self class) GetClassList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_class_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
//go:nosplit
func (self class) GetInheritersFromClass(class_ gd.StringName) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_inheriters_from_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the parent class of [param class].
*/
//go:nosplit
func (self class) GetParentClass(class_ gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_parent_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether the specified [param class] is available or not.
*/
//go:nosplit
func (self class) ClassExists(class_ gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_exists, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
//go:nosplit
func (self class) IsParentClass(class_ gd.StringName, inherits gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(inherits))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_parent_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) CanInstantiate(class_ gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates an instance of [param class].
*/
//go:nosplit
func (self class) Instantiate(class_ gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_instantiate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
//go:nosplit
func (self class) ClassHasSignal(class_ gd.StringName, signal gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(signal))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_signal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
//go:nosplit
func (self class) ClassGetSignal(class_ gd.StringName, signal gd.StringName) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(signal))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_signal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
//go:nosplit
func (self class) ClassGetSignalList(class_ gd.StringName, no_inheritance bool) Array.Contains[gd.Dictionary] {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_signal_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Dictionary]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetPropertyList(class_ gd.StringName, no_inheritance bool) Array.Contains[gd.Dictionary] {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Dictionary]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the value of [param property] of [param object] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetProperty(obj [1]gd.Object, property gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(property))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets [param property] value of [param object] to [param value].
*/
//go:nosplit
func (self class) ClassSetProperty(obj [1]gd.Object, property gd.StringName, value gd.Variant) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(property))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_set_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
//go:nosplit
func (self class) ClassGetPropertyDefaultValue(class_ gd.StringName, property gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(property))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_default_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
//go:nosplit
func (self class) ClassHasMethod(class_ gd.StringName, method gd.StringName, no_inheritance bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(method))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetMethodArgumentCount(class_ gd.StringName, method gd.StringName, no_inheritance bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(method))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_method_argument_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
//go:nosplit
func (self class) ClassGetMethodList(class_ gd.StringName, no_inheritance bool) Array.Contains[gd.Dictionary] {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_method_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Dictionary]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantList(class_ gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasIntegerConstant(class_ gd.StringName, name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_integer_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
//go:nosplit
func (self class) ClassGetIntegerConstant(class_ gd.StringName, name gd.StringName) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasEnum(class_ gd.StringName, name gd.StringName, no_inheritance bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_enum, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array with all the enums of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumList(class_ gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_enum_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumConstants(class_ gd.StringName, enum gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(enum))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_enum_constants, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantEnum(class_ gd.StringName, name gd.StringName, no_inheritance bool) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant_enum, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
//go:nosplit
func (self class) IsClassEnumBitfield(class_ gd.StringName, enum gd.StringName, no_inheritance bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	callframe.Arg(frame, pointers.Get(enum))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_class_enum_bitfield, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether this [param class] is enabled or not.
*/
//go:nosplit
func (self class) IsClassEnabled(class_ gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_class_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ClassDB", func(ptr gd.Object) any { return [1]gdclass.ClassDB{*(*gdclass.ClassDB)(unsafe.Pointer(&ptr))} })
}

type Error = gd.Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
