package TLSOptions

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
TLSOptions abstracts the configuration options for the [StreamPeerTLS] and [PacketPeerDTLS] classes.
Objects of this class cannot be instantiated directly, and one of the static methods [method client], [method client_unsafe], or [method server] should be used instead.
[codeblocks]
[gdscript]
# Create a TLS client configuration which uses our custom trusted CA chain.
var client_trusted_cas = load("res://my_trusted_cas.crt")
var client_tls_options = TLSOptions.client(client_trusted_cas)

# Create a TLS server configuration.
var server_certs = load("res://my_server_cas.crt")
var server_key = load("res://my_server_key.key")
var server_tls_options = TLSOptions.server(server_key, server_certs)
[/gdscript]
[/codeblocks]

*/
type Simple [1]classdb.TLSOptions
func (self Simple) Client(trusted_chain [1]classdb.X509Certificate, common_name_override string) [1]classdb.TLSOptions {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TLSOptions(Expert(self).Client(gc, trusted_chain, gc.String(common_name_override)))
}
func (self Simple) ClientUnsafe(trusted_chain [1]classdb.X509Certificate) [1]classdb.TLSOptions {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TLSOptions(Expert(self).ClientUnsafe(gc, trusted_chain))
}
func (self Simple) Server(key [1]classdb.CryptoKey, certificate [1]classdb.X509Certificate) [1]classdb.TLSOptions {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TLSOptions(Expert(self).Server(gc, key, certificate))
}
func (self Simple) IsServer() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsServer())
}
func (self Simple) IsUnsafeClient() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUnsafeClient())
}
func (self Simple) GetCommonNameOverride() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCommonNameOverride(gc).String())
}
func (self Simple) GetTrustedCaChain() [1]classdb.X509Certificate {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.X509Certificate(Expert(self).GetTrustedCaChain(gc))
}
func (self Simple) GetPrivateKey() [1]classdb.CryptoKey {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.CryptoKey(Expert(self).GetPrivateKey(gc))
}
func (self Simple) GetOwnCertificate() [1]classdb.X509Certificate {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.X509Certificate(Expert(self).GetOwnCertificate(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TLSOptions
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a TLS client configuration which validates certificates and their common names (fully qualified domain names).
You can specify a custom [param trusted_chain] of certification authorities (the default CA list will be used if [code]null[/code]), and optionally provide a [param common_name_override] if you expect the certificate to have a common name other than the server FQDN.
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
//go:nosplit
func (self class) Client(ctx gd.Lifetime, trusted_chain object.X509Certificate, common_name_override gd.String) object.TLSOptions {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(trusted_chain[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(common_name_override))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.TLSOptions.Bind_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TLSOptions
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates an [b]unsafe[/b] TLS client configuration where certificate validation is optional. You can optionally provide a valid [param trusted_chain], but the common name of the certificates will never be checked. Using this configuration for purposes other than testing [b]is not recommended[/b].
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
//go:nosplit
func (self class) ClientUnsafe(ctx gd.Lifetime, trusted_chain object.X509Certificate) object.TLSOptions {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(trusted_chain[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.TLSOptions.Bind_client_unsafe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TLSOptions
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a TLS server configuration using the provided [param key] and [param certificate].
[b]Note:[/b] The [param certificate] should include the full certificate chain up to the signing CA (certificates file can be concatenated using a general purpose text editor).
*/
//go:nosplit
func (self class) Server(ctx gd.Lifetime, key object.CryptoKey, certificate object.X509Certificate) object.TLSOptions {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(key[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(certificate[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.TLSOptions.Bind_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TLSOptions
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if created with [method TLSOptions.server], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsServer() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TLSOptions.Bind_is_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if created with [method TLSOptions.client_unsafe], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsUnsafeClient() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TLSOptions.Bind_is_unsafe_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the common name (domain name) override specified when creating with [method TLSOptions.client].
*/
//go:nosplit
func (self class) GetCommonNameOverride(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TLSOptions.Bind_get_common_name_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the CA [X509Certificate] chain specified when creating with [method TLSOptions.client] or [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) GetTrustedCaChain(ctx gd.Lifetime) object.X509Certificate {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TLSOptions.Bind_get_trusted_ca_chain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.X509Certificate
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [CryptoKey] specified when creating with [method TLSOptions.server].
*/
//go:nosplit
func (self class) GetPrivateKey(ctx gd.Lifetime) object.CryptoKey {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TLSOptions.Bind_get_private_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.CryptoKey
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [X509Certificate] specified when creating with [method TLSOptions.server].
*/
//go:nosplit
func (self class) GetOwnCertificate(ctx gd.Lifetime) object.X509Certificate {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TLSOptions.Bind_get_own_certificate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.X509Certificate
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTLSOptions() Expert { return self[0].AsTLSOptions() }


//go:nosplit
func (self Simple) AsTLSOptions() Simple { return self[0].AsTLSOptions() }


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
func init() {classdb.Register("TLSOptions", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
