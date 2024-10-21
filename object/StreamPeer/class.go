package StreamPeer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
StreamPeer is an abstract base class mostly used for stream-based protocols (such as TCP). It provides an API for sending and receiving data through streams as raw data or strings.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.StreamPeer
func (self Simple) PutData(data []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).PutData(gc.PackedByteSlice(data)))
}
func (self Simple) PutPartialData(data []byte) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).PutPartialData(gc, gc.PackedByteSlice(data)))
}
func (self Simple) GetData(bytes int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetData(gc, gd.Int(bytes)))
}
func (self Simple) GetPartialData(bytes int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetPartialData(gc, gd.Int(bytes)))
}
func (self Simple) GetAvailableBytes() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAvailableBytes()))
}
func (self Simple) SetBigEndian(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBigEndian(enable)
}
func (self Simple) IsBigEndianEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBigEndianEnabled())
}
func (self Simple) Put8(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Put8(gd.Int(value))
}
func (self Simple) PutU8(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutU8(gd.Int(value))
}
func (self Simple) Put16(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Put16(gd.Int(value))
}
func (self Simple) PutU16(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutU16(gd.Int(value))
}
func (self Simple) Put32(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Put32(gd.Int(value))
}
func (self Simple) PutU32(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutU32(gd.Int(value))
}
func (self Simple) Put64(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Put64(gd.Int(value))
}
func (self Simple) PutU64(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutU64(gd.Int(value))
}
func (self Simple) PutFloat(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutFloat(gd.Float(value))
}
func (self Simple) PutDouble(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutDouble(gd.Float(value))
}
func (self Simple) PutString(value string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutString(gc.String(value))
}
func (self Simple) PutUtf8String(value string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutUtf8String(gc.String(value))
}
func (self Simple) PutVar(value gd.Variant, full_objects bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PutVar(value, full_objects)
}
func (self Simple) Get8() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get8()))
}
func (self Simple) GetU8() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetU8()))
}
func (self Simple) Get16() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get16()))
}
func (self Simple) GetU16() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetU16()))
}
func (self Simple) Get32() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get32()))
}
func (self Simple) GetU32() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetU32()))
}
func (self Simple) Get64() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get64()))
}
func (self Simple) GetU64() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetU64()))
}
func (self Simple) GetFloat() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFloat()))
}
func (self Simple) GetDouble() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDouble()))
}
func (self Simple) GetString(bytes int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetString(gc, gd.Int(bytes)).String())
}
func (self Simple) GetUtf8String(bytes int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetUtf8String(gc, gd.Int(bytes)).String())
}
func (self Simple) GetVar(allow_objects bool) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetVar(gc, allow_objects))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StreamPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sends a chunk of data through the connection, blocking if necessary until the data is done sending. This function returns an [enum Error] code.
*/
//go:nosplit
func (self class) PutData(data gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sends a chunk of data through the connection. If all the data could not be sent at once, only part of it will. This function returns two values, an [enum Error] code and an integer, describing how much data was actually sent.
*/
//go:nosplit
func (self class) PutPartialData(ctx gd.Lifetime, data gd.PackedByteArray) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_partial_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the [param bytes] argument. If not enough bytes are available, the function will block until the desired amount is received. This function returns two values, an [enum Error] code and a data array.
*/
//go:nosplit
func (self class) GetData(ctx gd.Lifetime, bytes gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the "bytes" argument. If not enough bytes are available, the function will return how many were actually received. This function returns two values, an [enum Error] code, and a data array.
*/
//go:nosplit
func (self class) GetPartialData(ctx gd.Lifetime, bytes gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_partial_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the number of bytes this [StreamPeer] has available.
*/
//go:nosplit
func (self class) GetAvailableBytes() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_available_bytes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBigEndian(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_set_big_endian, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsBigEndianEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_is_big_endian_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Puts a signed byte into the stream.
*/
//go:nosplit
func (self class) Put8(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts an unsigned byte into the stream.
*/
//go:nosplit
func (self class) PutU8(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_u8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a signed 16-bit value into the stream.
*/
//go:nosplit
func (self class) Put16(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts an unsigned 16-bit value into the stream.
*/
//go:nosplit
func (self class) PutU16(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_u16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a signed 32-bit value into the stream.
*/
//go:nosplit
func (self class) Put32(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts an unsigned 32-bit value into the stream.
*/
//go:nosplit
func (self class) PutU32(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_u32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a signed 64-bit value into the stream.
*/
//go:nosplit
func (self class) Put64(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts an unsigned 64-bit value into the stream.
*/
//go:nosplit
func (self class) PutU64(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_u64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a single-precision float into the stream.
*/
//go:nosplit
func (self class) PutFloat(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a double-precision float into the stream.
*/
//go:nosplit
func (self class) PutDouble(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_double, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a zero-terminated ASCII string into the stream prepended by a 32-bit unsigned integer representing its size.
[b]Note:[/b] To put an ASCII string without prepending its size, you can use [method put_data]:
[codeblocks]
[gdscript]
put_data("Hello world".to_ascii_buffer())
[/gdscript]
[csharp]
PutData("Hello World".ToAsciiBuffer());
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) PutString(value gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a zero-terminated UTF-8 string into the stream prepended by a 32 bits unsigned integer representing its size.
[b]Note:[/b] To put a UTF-8 string without prepending its size, you can use [method put_data]:
[codeblocks]
[gdscript]
put_data("Hello world".to_utf8_buffer())
[/gdscript]
[csharp]
PutData("Hello World".ToUtf8Buffer());
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) PutUtf8String(value gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_utf8_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a Variant into the stream. If [param full_objects] is [code]true[/code] encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) PutVar(value gd.Variant, full_objects bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value))
	callframe.Arg(frame, full_objects)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_put_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a signed byte from the stream.
*/
//go:nosplit
func (self class) Get8() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets an unsigned byte from the stream.
*/
//go:nosplit
func (self class) GetU8() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_u8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets a signed 16-bit value from the stream.
*/
//go:nosplit
func (self class) Get16() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets an unsigned 16-bit value from the stream.
*/
//go:nosplit
func (self class) GetU16() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_u16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets a signed 32-bit value from the stream.
*/
//go:nosplit
func (self class) Get32() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets an unsigned 32-bit value from the stream.
*/
//go:nosplit
func (self class) GetU32() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_u32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets a signed 64-bit value from the stream.
*/
//go:nosplit
func (self class) Get64() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets an unsigned 64-bit value from the stream.
*/
//go:nosplit
func (self class) GetU64() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_u64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets a single-precision float from the stream.
*/
//go:nosplit
func (self class) GetFloat() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets a double-precision float from the stream.
*/
//go:nosplit
func (self class) GetDouble() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_double, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets an ASCII string with byte-length [param bytes] from the stream. If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_string].
*/
//go:nosplit
func (self class) GetString(ctx gd.Lifetime, bytes gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Gets a UTF-8 string with byte-length [param bytes] from the stream (this decodes the string sent as UTF-8). If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_utf8_string].
*/
//go:nosplit
func (self class) GetUtf8String(ctx gd.Lifetime, bytes gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_utf8_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Gets a Variant from the stream. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) GetVar(ctx gd.Lifetime, allow_objects bool) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeer.Bind_get_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsStreamPeer() Expert { return self[0].AsStreamPeer() }


//go:nosplit
func (self Simple) AsStreamPeer() Simple { return self[0].AsStreamPeer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("StreamPeer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
