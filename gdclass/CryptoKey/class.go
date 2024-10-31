package CryptoKey

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
The CryptoKey class represents a cryptographic key. Keys can be loaded and saved like any other [Resource].
They can be used to generate a self-signed [X509Certificate] via [method Crypto.generate_self_signed_certificate] and as private key in [method StreamPeerTLS.accept_stream] along with the appropriate certificate.
*/
type Instance [1]classdb.CryptoKey

/*
Saves a key to the given [param path]. If [param public_only] is [code]true[/code], only the public key will be saved.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
func (self Instance) Save(path string) gd.Error {
	return gd.Error(class(self).Save(gd.NewString(path), false))
}

/*
Loads a key from [param path]. If [param public_only] is [code]true[/code], only the public key will be loaded.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
func (self Instance) Load(path string) gd.Error {
	return gd.Error(class(self).Load(gd.NewString(path), false))
}

/*
Returns [code]true[/code] if this CryptoKey only has the public part, and not the private one.
*/
func (self Instance) IsPublicOnly() bool {
	return bool(class(self).IsPublicOnly())
}

/*
Returns a string containing the key in PEM format. If [param public_only] is [code]true[/code], only the public key will be included.
*/
func (self Instance) SaveToString() string {
	return string(class(self).SaveToString(false).String())
}

/*
Loads a key from the given [param string_key]. If [param public_only] is [code]true[/code], only the public key will be loaded.
*/
func (self Instance) LoadFromString(string_key string) gd.Error {
	return gd.Error(class(self).LoadFromString(gd.NewString(string_key), false))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CryptoKey

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CryptoKey"))
	return Instance{classdb.CryptoKey(object)}
}

/*
Saves a key to the given [param path]. If [param public_only] is [code]true[/code], only the public key will be saved.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
//go:nosplit
func (self class) Save(path gd.String, public_only bool) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_load, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this CryptoKey only has the public part, and not the private one.
*/
//go:nosplit
func (self class) IsPublicOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_is_public_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a string containing the key in PEM format. If [param public_only] is [code]true[/code], only the public key will be included.
*/
//go:nosplit
func (self class) SaveToString(public_only bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_save_to_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads a key from the given [param string_key]. If [param public_only] is [code]true[/code], only the public key will be loaded.
*/
//go:nosplit
func (self class) LoadFromString(string_key gd.String, public_only bool) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(string_key))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_load_from_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCryptoKey() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCryptoKey() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("CryptoKey", func(ptr gd.Object) any { return classdb.CryptoKey(ptr) }) }
