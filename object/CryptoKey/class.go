package CryptoKey

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The CryptoKey class represents a cryptographic key. Keys can be loaded and saved like any other [Resource].
They can be used to generate a self-signed [X509Certificate] via [method Crypto.generate_self_signed_certificate] and as private key in [method StreamPeerTLS.accept_stream] along with the appropriate certificate.

*/
type Simple [1]classdb.CryptoKey
func (self Simple) Save(path string, public_only bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Save(gc.String(path), public_only))
}
func (self Simple) Load(path string, public_only bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Load(gc.String(path), public_only))
}
func (self Simple) IsPublicOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPublicOnly())
}
func (self Simple) SaveToString(public_only bool) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).SaveToString(gc, public_only).String())
}
func (self Simple) LoadFromString(string_key string, public_only bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadFromString(gc.String(string_key), public_only))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CryptoKey
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Saves a key to the given [param path]. If [param public_only] is [code]true[/code], only the public key will be saved.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
//go:nosplit
func (self class) Save(path gd.String, public_only bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CryptoKey.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads a key from [param path]. If [param public_only] is [code]true[/code], only the public key will be loaded.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
//go:nosplit
func (self class) Load(path gd.String, public_only bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CryptoKey.Bind_load, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this CryptoKey only has the public part, and not the private one.
*/
//go:nosplit
func (self class) IsPublicOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CryptoKey.Bind_is_public_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a string containing the key in PEM format. If [param public_only] is [code]true[/code], only the public key will be included.
*/
//go:nosplit
func (self class) SaveToString(ctx gd.Lifetime, public_only bool) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CryptoKey.Bind_save_to_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Loads a key from the given [param string_key]. If [param public_only] is [code]true[/code], only the public key will be loaded.
*/
//go:nosplit
func (self class) LoadFromString(string_key gd.String, public_only bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(string_key))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CryptoKey.Bind_load_from_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCryptoKey() Expert { return self[0].AsCryptoKey() }


//go:nosplit
func (self Simple) AsCryptoKey() Simple { return self[0].AsCryptoKey() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
func init() {classdb.Register("CryptoKey", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
