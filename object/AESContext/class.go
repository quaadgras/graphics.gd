package AESContext

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
This class holds the context information required for encryption and decryption operations with AES (Advanced Encryption Standard). Both AES-ECB and AES-CBC modes are supported.
[codeblocks]
[gdscript]
extends Node

var aes = AESContext.new()

func _ready():
    var key = "My secret key!!!" # Key must be either 16 or 32 bytes.
    var data = "My secret text!!" # Data size must be multiple of 16 bytes, apply padding if needed.
    # Encrypt ECB
    aes.start(AESContext.MODE_ECB_ENCRYPT, key.to_utf8_buffer())
    var encrypted = aes.update(data.to_utf8_buffer())
    aes.finish()
    # Decrypt ECB
    aes.start(AESContext.MODE_ECB_DECRYPT, key.to_utf8_buffer())
    var decrypted = aes.update(encrypted)
    aes.finish()
    # Check ECB
    assert(decrypted == data.to_utf8_buffer())

    var iv = "My secret iv!!!!" # IV must be of exactly 16 bytes.
    # Encrypt CBC
    aes.start(AESContext.MODE_CBC_ENCRYPT, key.to_utf8_buffer(), iv.to_utf8_buffer())
    encrypted = aes.update(data.to_utf8_buffer())
    aes.finish()
    # Decrypt CBC
    aes.start(AESContext.MODE_CBC_DECRYPT, key.to_utf8_buffer(), iv.to_utf8_buffer())
    decrypted = aes.update(encrypted)
    aes.finish()
    # Check CBC
    assert(decrypted == data.to_utf8_buffer())
[/gdscript]
[csharp]
using Godot;
using System.Diagnostics;

public partial class MyNode : Node
{
    private AesContext _aes = new AesContext();

    public override void _Ready()
    {
        string key = "My secret key!!!"; // Key must be either 16 or 32 bytes.
        string data = "My secret text!!"; // Data size must be multiple of 16 bytes, apply padding if needed.
        // Encrypt ECB
        _aes.Start(AesContext.Mode.EcbEncrypt, key.ToUtf8Buffer());
        byte[] encrypted = _aes.Update(data.ToUtf8Buffer());
        _aes.Finish();
        // Decrypt ECB
        _aes.Start(AesContext.Mode.EcbDecrypt, key.ToUtf8Buffer());
        byte[] decrypted = _aes.Update(encrypted);
        _aes.Finish();
        // Check ECB
        Debug.Assert(decrypted == data.ToUtf8Buffer());

        string iv = "My secret iv!!!!"; // IV must be of exactly 16 bytes.
        // Encrypt CBC
        _aes.Start(AesContext.Mode.EcbEncrypt, key.ToUtf8Buffer(), iv.ToUtf8Buffer());
        encrypted = _aes.Update(data.ToUtf8Buffer());
        _aes.Finish();
        // Decrypt CBC
        _aes.Start(AesContext.Mode.EcbDecrypt, key.ToUtf8Buffer(), iv.ToUtf8Buffer());
        decrypted = _aes.Update(encrypted);
        _aes.Finish();
        // Check CBC
        Debug.Assert(decrypted == data.ToUtf8Buffer());
    }
}
[/csharp]
[/codeblocks]

*/
type Simple [1]classdb.AESContext
func (self Simple) Start(mode classdb.AESContextMode, key []byte, iv []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Start(mode, gc.PackedByteSlice(key), gc.PackedByteSlice(iv)))
}
func (self Simple) Update(src []byte) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).Update(gc, gc.PackedByteSlice(src)).Bytes())
}
func (self Simple) GetIvState() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetIvState(gc).Bytes())
}
func (self Simple) Finish() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Finish()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AESContext
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Start the AES context in the given [param mode]. A [param key] of either 16 or 32 bytes must always be provided, while an [param iv] (initialization vector) of exactly 16 bytes, is only needed when [param mode] is either [constant MODE_CBC_ENCRYPT] or [constant MODE_CBC_DECRYPT].
*/
//go:nosplit
func (self class) Start(mode classdb.AESContextMode, key gd.PackedByteArray, iv gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, mmm.Get(iv))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AESContext.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Run the desired operation for this AES context. Will return a [PackedByteArray] containing the result of encrypting (or decrypting) the given [param src]. See [method start] for mode of operation.
[b]Note:[/b] The size of [param src] must be a multiple of 16. Apply some padding if needed.
*/
//go:nosplit
func (self class) Update(ctx gd.Lifetime, src gd.PackedByteArray) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AESContext.Bind_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Get the current IV state for this context (IV gets updated when calling [method update]). You normally don't need this function.
[b]Note:[/b] This function only makes sense when the context is started with [constant MODE_CBC_ENCRYPT] or [constant MODE_CBC_DECRYPT].
*/
//go:nosplit
func (self class) GetIvState(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AESContext.Bind_get_iv_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Close this AES context so it can be started again. See [method start].
*/
//go:nosplit
func (self class) Finish()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AESContext.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsAESContext() Expert { return self[0].AsAESContext() }


//go:nosplit
func (self Simple) AsAESContext() Simple { return self[0].AsAESContext() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AESContext", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Mode = classdb.AESContextMode

const (
/*AES electronic codebook encryption mode.*/
	ModeEcbEncrypt Mode = 0
/*AES electronic codebook decryption mode.*/
	ModeEcbDecrypt Mode = 1
/*AES cipher blocker chaining encryption mode.*/
	ModeCbcEncrypt Mode = 2
/*AES cipher blocker chaining decryption mode.*/
	ModeCbcDecrypt Mode = 3
/*Maximum value for the mode enum.*/
	ModeMax Mode = 4
)
