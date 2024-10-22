package CryptoKey

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The CryptoKey class represents a cryptographic key. Keys can be loaded and saved like any other [Resource].
They can be used to generate a self-signed [X509Certificate] via [method Crypto.generate_self_signed_certificate] and as private key in [method StreamPeerTLS.accept_stream] along with the appropriate certificate.

*/
type Go [1]classdb.CryptoKey

/*
Saves a key to the given [param path]. If [param public_only] is [code]true[/code], only the public key will be saved.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
func (self Go) Save(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Save(gc.String(path), false))
}

/*
Loads a key from [param path]. If [param public_only] is [code]true[/code], only the public key will be loaded.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
func (self Go) Load(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Load(gc.String(path), false))
}

/*
Returns [code]true[/code] if this CryptoKey only has the public part, and not the private one.
*/
func (self Go) IsPublicOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPublicOnly())
}

/*
Returns a string containing the key in PEM format. If [param public_only] is [code]true[/code], only the public key will be included.
*/
func (self Go) SaveToString() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).SaveToString(gc, false).String())
}

/*
Loads a key from the given [param string_key]. If [param public_only] is [code]true[/code], only the public key will be loaded.
*/
func (self Go) LoadFromString(string_key string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).LoadFromString(gc.String(string_key), false))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CryptoKey
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CryptoKey"))
	return *(*Go)(unsafe.Pointer(&object))
}

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
func (self class) AsCryptoKey() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCryptoKey() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("CryptoKey", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
