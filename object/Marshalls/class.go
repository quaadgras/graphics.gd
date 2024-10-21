package Marshalls

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Provides data transformation and encoding utility functions.

*/
type Simple [1]classdb.Marshalls
func (self Simple) VariantToBase64(variant gd.Variant, full_objects bool) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).VariantToBase64(gc, variant, full_objects).String())
}
func (self Simple) Base64ToVariant(base64_str string, allow_objects bool) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).Base64ToVariant(gc, gc.String(base64_str), allow_objects))
}
func (self Simple) RawToBase64(array []byte) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).RawToBase64(gc, gc.PackedByteSlice(array)).String())
}
func (self Simple) Base64ToRaw(base64_str string) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).Base64ToRaw(gc, gc.String(base64_str)).Bytes())
}
func (self Simple) Utf8ToBase64(utf8_str string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).Utf8ToBase64(gc, gc.String(utf8_str)).String())
}
func (self Simple) Base64ToUtf8(base64_str string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).Base64ToUtf8(gc, gc.String(base64_str)).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Marshalls
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns a Base64-encoded string of the [Variant] [param variant]. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) VariantToBase64(ctx gd.Lifetime, variant gd.Variant, full_objects bool) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(variant))
	callframe.Arg(frame, full_objects)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marshalls.Bind_variant_to_base64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a decoded [Variant] corresponding to the Base64-encoded string [param base64_str]. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) Base64ToVariant(ctx gd.Lifetime, base64_str gd.String, allow_objects bool) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base64_str))
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marshalls.Bind_base64_to_variant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a Base64-encoded string of a given [PackedByteArray].
*/
//go:nosplit
func (self class) RawToBase64(ctx gd.Lifetime, array gd.PackedByteArray) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(array))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marshalls.Bind_raw_to_base64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a decoded [PackedByteArray] corresponding to the Base64-encoded string [param base64_str].
*/
//go:nosplit
func (self class) Base64ToRaw(ctx gd.Lifetime, base64_str gd.String) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base64_str))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marshalls.Bind_base64_to_raw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a Base64-encoded string of the UTF-8 string [param utf8_str].
*/
//go:nosplit
func (self class) Utf8ToBase64(ctx gd.Lifetime, utf8_str gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(utf8_str))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marshalls.Bind_utf8_to_base64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a decoded string corresponding to the Base64-encoded string [param base64_str].
*/
//go:nosplit
func (self class) Base64ToUtf8(ctx gd.Lifetime, base64_str gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base64_str))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marshalls.Bind_base64_to_utf8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsMarshalls() Expert { return self[0].AsMarshalls() }


//go:nosplit
func (self Simple) AsMarshalls() Simple { return self[0].AsMarshalls() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Marshalls", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
