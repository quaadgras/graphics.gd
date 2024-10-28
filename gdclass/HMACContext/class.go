package HMACContext

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The HMACContext class is useful for advanced HMAC use cases, such as streaming the message as it supports creating the message over time rather than providing it all at once.
[codeblocks]
[gdscript]
extends Node
var ctx = HMACContext.new()

func _ready():
    var key = "supersecret".to_utf8_buffer()
    var err = ctx.start(HashingContext.HASH_SHA256, key)
    assert(err == OK)
    var msg1 = "this is ".to_utf8_buffer()
    var msg2 = "super duper secret".to_utf8_buffer()
    err = ctx.update(msg1)
    assert(err == OK)
    err = ctx.update(msg2)
    assert(err == OK)
    var hmac = ctx.finish()
    print(hmac.hex_encode())

[/gdscript]
[csharp]
using Godot;
using System.Diagnostics;

public partial class MyNode : Node
{
    private HmacContext _ctx = new HmacContext();

    public override void _Ready()
    {
        byte[] key = "supersecret".ToUtf8Buffer();
        Error err = _ctx.Start(HashingContext.HashType.Sha256, key);
        Debug.Assert(err == Error.Ok);
        byte[] msg1 = "this is ".ToUtf8Buffer();
        byte[] msg2 = "super duper secret".ToUtf8Buffer();
        err = _ctx.Update(msg1);
        Debug.Assert(err == Error.Ok);
        err = _ctx.Update(msg2);
        Debug.Assert(err == Error.Ok);
        byte[] hmac = _ctx.Finish();
        GD.Print(hmac.HexEncode());
    }
}
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.HMACContext

/*
Initializes the HMACContext. This method cannot be called again on the same HMACContext until [method finish] has been called.
*/
func (self Go) Start(hash_type classdb.HashingContextHashType, key []byte) gd.Error {
	return gd.Error(class(self).Start(hash_type, gd.NewPackedByteSlice(key)))
}

/*
Updates the message to be HMACed. This can be called multiple times before [method finish] is called to append [param data] to the message, but cannot be called until [method start] has been called.
*/
func (self Go) Update(data []byte) gd.Error {
	return gd.Error(class(self).Update(gd.NewPackedByteSlice(data)))
}

/*
Returns the resulting HMAC. If the HMAC failed, an empty [PackedByteArray] is returned.
*/
func (self Go) Finish() []byte {
	return []byte(class(self).Finish().Bytes())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.HMACContext
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HMACContext"))
	return Go{classdb.HMACContext(object)}
}

/*
Initializes the HMACContext. This method cannot be called again on the same HMACContext until [method finish] has been called.
*/
//go:nosplit
func (self class) Start(hash_type classdb.HashingContextHashType, key gd.PackedByteArray) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, hash_type)
	callframe.Arg(frame, discreet.Get(key))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HMACContext.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the message to be HMACed. This can be called multiple times before [method finish] is called to append [param data] to the message, but cannot be called until [method start] has been called.
*/
//go:nosplit
func (self class) Update(data gd.PackedByteArray) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HMACContext.Bind_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the resulting HMAC. If the HMAC failed, an empty [PackedByteArray] is returned.
*/
//go:nosplit
func (self class) Finish() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HMACContext.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsHMACContext() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsHMACContext() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("HMACContext", func(ptr gd.Object) any { return classdb.HMACContext(ptr) })}
