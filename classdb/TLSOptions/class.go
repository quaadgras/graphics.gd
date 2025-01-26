// Package TLSOptions provides methods for working with TLSOptions object instances.
package TLSOptions

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

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
type Instance [1]gdclass.TLSOptions

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTLSOptions() Instance
}

/*
Creates a TLS client configuration which validates certificates and their common names (fully qualified domain names).
You can specify a custom [param trusted_chain] of certification authorities (the default CA list will be used if [code]null[/code]), and optionally provide a [param common_name_override] if you expect the certificate to have a common name other than the server FQDN.
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
func Client() [1]gdclass.TLSOptions { //gd:TLSOptions.client
	self := Instance{}
	return [1]gdclass.TLSOptions(class(self).Client([1][1]gdclass.X509Certificate{}[0], gd.NewString("")))
}

/*
Creates an [b]unsafe[/b] TLS client configuration where certificate validation is optional. You can optionally provide a valid [param trusted_chain], but the common name of the certificates will never be checked. Using this configuration for purposes other than testing [b]is not recommended[/b].
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
func ClientUnsafe() [1]gdclass.TLSOptions { //gd:TLSOptions.client_unsafe
	self := Instance{}
	return [1]gdclass.TLSOptions(class(self).ClientUnsafe([1][1]gdclass.X509Certificate{}[0]))
}

/*
Creates a TLS server configuration using the provided [param key] and [param certificate].
[b]Note:[/b] The [param certificate] should include the full certificate chain up to the signing CA (certificates file can be concatenated using a general purpose text editor).
*/
func Server(key [1]gdclass.CryptoKey, certificate [1]gdclass.X509Certificate) [1]gdclass.TLSOptions { //gd:TLSOptions.server
	self := Instance{}
	return [1]gdclass.TLSOptions(class(self).Server(key, certificate))
}

/*
Returns [code]true[/code] if created with [method TLSOptions.server], [code]false[/code] otherwise.
*/
func (self Instance) IsServer() bool { //gd:TLSOptions.is_server
	return bool(class(self).IsServer())
}

/*
Returns [code]true[/code] if created with [method TLSOptions.client_unsafe], [code]false[/code] otherwise.
*/
func (self Instance) IsUnsafeClient() bool { //gd:TLSOptions.is_unsafe_client
	return bool(class(self).IsUnsafeClient())
}

/*
Returns the common name (domain name) override specified when creating with [method TLSOptions.client].
*/
func (self Instance) GetCommonNameOverride() string { //gd:TLSOptions.get_common_name_override
	return string(class(self).GetCommonNameOverride().String())
}

/*
Returns the CA [X509Certificate] chain specified when creating with [method TLSOptions.client] or [method TLSOptions.client_unsafe].
*/
func (self Instance) GetTrustedCaChain() [1]gdclass.X509Certificate { //gd:TLSOptions.get_trusted_ca_chain
	return [1]gdclass.X509Certificate(class(self).GetTrustedCaChain())
}

/*
Returns the [CryptoKey] specified when creating with [method TLSOptions.server].
*/
func (self Instance) GetPrivateKey() [1]gdclass.CryptoKey { //gd:TLSOptions.get_private_key
	return [1]gdclass.CryptoKey(class(self).GetPrivateKey())
}

/*
Returns the [X509Certificate] specified when creating with [method TLSOptions.server].
*/
func (self Instance) GetOwnCertificate() [1]gdclass.X509Certificate { //gd:TLSOptions.get_own_certificate
	return [1]gdclass.X509Certificate(class(self).GetOwnCertificate())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TLSOptions

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TLSOptions"))
	casted := Instance{*(*gdclass.TLSOptions)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates a TLS client configuration which validates certificates and their common names (fully qualified domain names).
You can specify a custom [param trusted_chain] of certification authorities (the default CA list will be used if [code]null[/code]), and optionally provide a [param common_name_override] if you expect the certificate to have a common name other than the server FQDN.
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
//go:nosplit
func (self class) Client(trusted_chain [1]gdclass.X509Certificate, common_name_override gd.String) [1]gdclass.TLSOptions { //gd:TLSOptions.client
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(trusted_chain[0])[0])
	callframe.Arg(frame, pointers.Get(common_name_override))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_client, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TLSOptions{gd.PointerWithOwnershipTransferredToGo[gdclass.TLSOptions](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates an [b]unsafe[/b] TLS client configuration where certificate validation is optional. You can optionally provide a valid [param trusted_chain], but the common name of the certificates will never be checked. Using this configuration for purposes other than testing [b]is not recommended[/b].
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
//go:nosplit
func (self class) ClientUnsafe(trusted_chain [1]gdclass.X509Certificate) [1]gdclass.TLSOptions { //gd:TLSOptions.client_unsafe
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(trusted_chain[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_client_unsafe, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TLSOptions{gd.PointerWithOwnershipTransferredToGo[gdclass.TLSOptions](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a TLS server configuration using the provided [param key] and [param certificate].
[b]Note:[/b] The [param certificate] should include the full certificate chain up to the signing CA (certificates file can be concatenated using a general purpose text editor).
*/
//go:nosplit
func (self class) Server(key [1]gdclass.CryptoKey, certificate [1]gdclass.X509Certificate) [1]gdclass.TLSOptions { //gd:TLSOptions.server
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key[0])[0])
	callframe.Arg(frame, pointers.Get(certificate[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_server, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TLSOptions{gd.PointerWithOwnershipTransferredToGo[gdclass.TLSOptions](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if created with [method TLSOptions.server], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsServer() bool { //gd:TLSOptions.is_server
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_is_server, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if created with [method TLSOptions.client_unsafe], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsUnsafeClient() bool { //gd:TLSOptions.is_unsafe_client
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_is_unsafe_client, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the common name (domain name) override specified when creating with [method TLSOptions.client].
*/
//go:nosplit
func (self class) GetCommonNameOverride() gd.String { //gd:TLSOptions.get_common_name_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_common_name_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the CA [X509Certificate] chain specified when creating with [method TLSOptions.client] or [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) GetTrustedCaChain() [1]gdclass.X509Certificate { //gd:TLSOptions.get_trusted_ca_chain
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_trusted_ca_chain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.X509Certificate{gd.PointerWithOwnershipTransferredToGo[gdclass.X509Certificate](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [CryptoKey] specified when creating with [method TLSOptions.server].
*/
//go:nosplit
func (self class) GetPrivateKey() [1]gdclass.CryptoKey { //gd:TLSOptions.get_private_key
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_private_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CryptoKey{gd.PointerWithOwnershipTransferredToGo[gdclass.CryptoKey](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [X509Certificate] specified when creating with [method TLSOptions.server].
*/
//go:nosplit
func (self class) GetOwnCertificate() [1]gdclass.X509Certificate { //gd:TLSOptions.get_own_certificate
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_own_certificate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.X509Certificate{gd.PointerWithOwnershipTransferredToGo[gdclass.X509Certificate](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsTLSOptions() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTLSOptions() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("TLSOptions", func(ptr gd.Object) any { return [1]gdclass.TLSOptions{*(*gdclass.TLSOptions)(unsafe.Pointer(&ptr))} })
}
