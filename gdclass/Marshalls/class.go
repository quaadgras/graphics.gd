package Marshalls

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Provides data transformation and encoding utility functions.
*/
var self gdclass.Marshalls
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Marshalls)
	self = *(*gdclass.Marshalls)(unsafe.Pointer(&obj))
}

/*
Returns a Base64-encoded string of the [Variant] [param variant]. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
func VariantToBase64(variant gd.Variant) string {
	once.Do(singleton)
	return string(class(self).VariantToBase64(variant, false).String())
}

/*
Returns a decoded [Variant] corresponding to the Base64-encoded string [param base64_str]. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
func Base64ToVariant(base64_str string) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).Base64ToVariant(gd.NewString(base64_str), false))
}

/*
Returns a Base64-encoded string of a given [PackedByteArray].
*/
func RawToBase64(array []byte) string {
	once.Do(singleton)
	return string(class(self).RawToBase64(gd.NewPackedByteSlice(array)).String())
}

/*
Returns a decoded [PackedByteArray] corresponding to the Base64-encoded string [param base64_str].
*/
func Base64ToRaw(base64_str string) []byte {
	once.Do(singleton)
	return []byte(class(self).Base64ToRaw(gd.NewString(base64_str)).Bytes())
}

/*
Returns a Base64-encoded string of the UTF-8 string [param utf8_str].
*/
func Utf8ToBase64(utf8_str string) string {
	once.Do(singleton)
	return string(class(self).Utf8ToBase64(gd.NewString(utf8_str)).String())
}

/*
Returns a decoded string corresponding to the Base64-encoded string [param base64_str].
*/
func Base64ToUtf8(base64_str string) string {
	once.Do(singleton)
	return string(class(self).Base64ToUtf8(gd.NewString(base64_str)).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.Marshalls

func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Returns a Base64-encoded string of the [Variant] [param variant]. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) VariantToBase64(variant gd.Variant, full_objects bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(variant))
	callframe.Arg(frame, full_objects)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_variant_to_base64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a decoded [Variant] corresponding to the Base64-encoded string [param base64_str]. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) Base64ToVariant(base64_str gd.String, allow_objects bool) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base64_str))
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_base64_to_variant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a Base64-encoded string of a given [PackedByteArray].
*/
//go:nosplit
func (self class) RawToBase64(array gd.PackedByteArray) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_raw_to_base64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a decoded [PackedByteArray] corresponding to the Base64-encoded string [param base64_str].
*/
//go:nosplit
func (self class) Base64ToRaw(base64_str gd.String) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base64_str))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_base64_to_raw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a Base64-encoded string of the UTF-8 string [param utf8_str].
*/
//go:nosplit
func (self class) Utf8ToBase64(utf8_str gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(utf8_str))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_utf8_to_base64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a decoded string corresponding to the Base64-encoded string [param base64_str].
*/
//go:nosplit
func (self class) Base64ToUtf8(base64_str gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base64_str))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_base64_to_utf8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() { classdb.Register("Marshalls", func(ptr gd.Object) any { return classdb.Marshalls(ptr) }) }
