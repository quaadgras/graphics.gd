package ClassDB

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Provides access to metadata stored for every available class.

*/
var self gdclass.ClassDB
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.ClassDB)
	self = *(*gdclass.ClassDB)(unsafe.Pointer(&obj))
}

/*
Returns the names of all the classes available.
*/
func GetClassList() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetClassList(gc).Strings(gc))
}

/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
func GetInheritersFromClass(class_ string) []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetInheritersFromClass(gc, gc.StringName(class_)).Strings(gc))
}

/*
Returns the parent class of [param class].
*/
func GetParentClass(class_ string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetParentClass(gc, gc.StringName(class_)).String())
}

/*
Returns whether the specified [param class] is available or not.
*/
func ClassExists(class_ string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).ClassExists(gc.StringName(class_)))
}

/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
func IsParentClass(class_ string, inherits string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsParentClass(gc.StringName(class_), gc.StringName(inherits)))
}

/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
func CanInstantiate(class_ string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).CanInstantiate(gc.StringName(class_)))
}

/*
Creates an instance of [param class].
*/
func Instantiate(class_ string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Variant(class(self).Instantiate(gc, gc.StringName(class_)))
}

/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
func ClassHasSignal(class_ string, signal string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).ClassHasSignal(gc.StringName(class_), gc.StringName(signal)))
}

/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
func ClassGetSignal(class_ string, signal string) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Dictionary(class(self).ClassGetSignal(gc, gc.StringName(class_), gc.StringName(signal)))
}

/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
func ClassGetSignalList(class_ string) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.ArrayOf[gd.Dictionary](class(self).ClassGetSignalList(gc, gc.StringName(class_), false))
}

/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetPropertyList(class_ string) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.ArrayOf[gd.Dictionary](class(self).ClassGetPropertyList(gc, gc.StringName(class_), false))
}

/*
Returns the value of [param property] of [param object] or its ancestry.
*/
func ClassGetProperty(obj gd.Object, property string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Variant(class(self).ClassGetProperty(gc, obj, gc.StringName(property)))
}

/*
Sets [param property] value of [param object] to [param value].
*/
func ClassSetProperty(obj gd.Object, property string, value gd.Variant) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Error(class(self).ClassSetProperty(obj, gc.StringName(property), value))
}

/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
func ClassGetPropertyDefaultValue(class_ string, property string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Variant(class(self).ClassGetPropertyDefaultValue(gc, gc.StringName(class_), gc.StringName(property)))
}

/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
func ClassHasMethod(class_ string, method string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).ClassHasMethod(gc.StringName(class_), gc.StringName(method), false))
}

/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetMethodArgumentCount(class_ string, method string) int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).ClassGetMethodArgumentCount(gc.StringName(class_), gc.StringName(method), false)))
}

/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
func ClassGetMethodList(class_ string) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.ArrayOf[gd.Dictionary](class(self).ClassGetMethodList(gc, gc.StringName(class_), false))
}

/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
func ClassGetIntegerConstantList(class_ string) []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).ClassGetIntegerConstantList(gc, gc.StringName(class_), false).Strings(gc))
}

/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
func ClassHasIntegerConstant(class_ string, name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).ClassHasIntegerConstant(gc.StringName(class_), gc.StringName(name)))
}

/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
func ClassGetIntegerConstant(class_ string, name string) int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).ClassGetIntegerConstant(gc.StringName(class_), gc.StringName(name))))
}

/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
func ClassHasEnum(class_ string, name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).ClassHasEnum(gc.StringName(class_), gc.StringName(name), false))
}

/*
Returns an array with all the enums of [param class] or its ancestry.
*/
func ClassGetEnumList(class_ string) []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).ClassGetEnumList(gc, gc.StringName(class_), false).Strings(gc))
}

/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
func ClassGetEnumConstants(class_ string, enum string) []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).ClassGetEnumConstants(gc, gc.StringName(class_), gc.StringName(enum), false).Strings(gc))
}

/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
func ClassGetIntegerConstantEnum(class_ string, name string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).ClassGetIntegerConstantEnum(gc, gc.StringName(class_), gc.StringName(name), false).String())
}

/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
func IsClassEnumBitfield(class_ string, enum string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsClassEnumBitfield(gc.StringName(class_), gc.StringName(enum), false))
}

/*
Returns whether this [param class] is enabled or not.
*/
func IsClassEnabled(class_ string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsClassEnabled(gc.StringName(class_)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.ClassDB
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the names of all the classes available.
*/
//go:nosplit
func (self class) GetClassList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_get_class_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
//go:nosplit
func (self class) GetInheritersFromClass(ctx gd.Lifetime, class_ gd.StringName) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_get_inheriters_from_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the parent class of [param class].
*/
//go:nosplit
func (self class) GetParentClass(ctx gd.Lifetime, class_ gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_get_parent_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether the specified [param class] is available or not.
*/
//go:nosplit
func (self class) ClassExists(class_ gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
//go:nosplit
func (self class) IsParentClass(class_ gd.StringName, inherits gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(inherits))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_is_parent_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) CanInstantiate(class_ gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an instance of [param class].
*/
//go:nosplit
func (self class) Instantiate(ctx gd.Lifetime, class_ gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
//go:nosplit
func (self class) ClassHasSignal(class_ gd.StringName, signal gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(signal))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_has_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
//go:nosplit
func (self class) ClassGetSignal(ctx gd.Lifetime, class_ gd.StringName, signal gd.StringName) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(signal))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
//go:nosplit
func (self class) ClassGetSignalList(ctx gd.Lifetime, class_ gd.StringName, no_inheritance bool) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_signal_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetPropertyList(ctx gd.Lifetime, class_ gd.StringName, no_inheritance bool) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_property_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns the value of [param property] of [param object] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetProperty(ctx gd.Lifetime, obj gd.Object, property gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(property))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets [param property] value of [param object] to [param value].
*/
//go:nosplit
func (self class) ClassSetProperty(obj gd.Object, property gd.StringName, value gd.Variant) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_set_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
//go:nosplit
func (self class) ClassGetPropertyDefaultValue(ctx gd.Lifetime, class_ gd.StringName, property gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(property))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_property_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
//go:nosplit
func (self class) ClassHasMethod(class_ gd.StringName, method gd.StringName, no_inheritance bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(method))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_has_method, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetMethodArgumentCount(class_ gd.StringName, method gd.StringName, no_inheritance bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(method))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_method_argument_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
//go:nosplit
func (self class) ClassGetMethodList(ctx gd.Lifetime, class_ gd.StringName, no_inheritance bool) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_method_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantList(ctx gd.Lifetime, class_ gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_integer_constant_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasIntegerConstant(class_ gd.StringName, name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_has_integer_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
//go:nosplit
func (self class) ClassGetIntegerConstant(class_ gd.StringName, name gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_integer_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasEnum(class_ gd.StringName, name gd.StringName, no_inheritance bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_has_enum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with all the enums of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumList(ctx gd.Lifetime, class_ gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_enum_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumConstants(ctx gd.Lifetime, class_ gd.StringName, enum gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(enum))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_enum_constants, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantEnum(ctx gd.Lifetime, class_ gd.StringName, name gd.StringName, no_inheritance bool) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_class_get_integer_constant_enum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
//go:nosplit
func (self class) IsClassEnumBitfield(class_ gd.StringName, enum gd.StringName, no_inheritance bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	callframe.Arg(frame, mmm.Get(enum))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_is_class_enum_bitfield, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether this [param class] is enabled or not.
*/
//go:nosplit
func (self class) IsClassEnabled(class_ gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ClassDB.Bind_is_class_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("ClassDB", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
