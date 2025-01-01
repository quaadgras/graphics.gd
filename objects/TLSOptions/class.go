package TLSOptions

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.TLSOptions
type Any interface {
	gd.IsClass
	AsTLSOptions() Instance
}

/*
Creates a TLS client configuration which validates certificates and their common names (fully qualified domain names).
You can specify a custom [param trusted_chain] of certification authorities (the default CA list will be used if [code]null[/code]), and optionally provide a [param common_name_override] if you expect the certificate to have a common name other than the server FQDN.
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
func Client() objects.TLSOptions {
	self := Instance{}
	return objects.TLSOptions(class(self).Client([1]objects.X509Certificate{}[0], gd.NewString("")))
}

/*
Creates an [b]unsafe[/b] TLS client configuration where certificate validation is optional. You can optionally provide a valid [param trusted_chain], but the common name of the certificates will never be checked. Using this configuration for purposes other than testing [b]is not recommended[/b].
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
func ClientUnsafe() objects.TLSOptions {
	self := Instance{}
	return objects.TLSOptions(class(self).ClientUnsafe([1]objects.X509Certificate{}[0]))
}

/*
Creates a TLS server configuration using the provided [param key] and [param certificate].
[b]Note:[/b] The [param certificate] should include the full certificate chain up to the signing CA (certificates file can be concatenated using a general purpose text editor).
*/
func Server(key objects.CryptoKey, certificate objects.X509Certificate) objects.TLSOptions {
	self := Instance{}
	return objects.TLSOptions(class(self).Server(key, certificate))
}

/*
Returns [code]true[/code] if created with [method TLSOptions.server], [code]false[/code] otherwise.
*/
func (self Instance) IsServer() bool {
	return bool(class(self).IsServer())
}

/*
Returns [code]true[/code] if created with [method TLSOptions.client_unsafe], [code]false[/code] otherwise.
*/
func (self Instance) IsUnsafeClient() bool {
	return bool(class(self).IsUnsafeClient())
}

/*
Returns the common name (domain name) override specified when creating with [method TLSOptions.client].
*/
func (self Instance) GetCommonNameOverride() string {
	return string(class(self).GetCommonNameOverride().String())
}

/*
Returns the CA [X509Certificate] chain specified when creating with [method TLSOptions.client] or [method TLSOptions.client_unsafe].
*/
func (self Instance) GetTrustedCaChain() objects.X509Certificate {
	return objects.X509Certificate(class(self).GetTrustedCaChain())
}

/*
Returns the [CryptoKey] specified when creating with [method TLSOptions.server].
*/
func (self Instance) GetPrivateKey() objects.CryptoKey {
	return objects.CryptoKey(class(self).GetPrivateKey())
}

/*
Returns the [X509Certificate] specified when creating with [method TLSOptions.server].
*/
func (self Instance) GetOwnCertificate() objects.X509Certificate {
	return objects.X509Certificate(class(self).GetOwnCertificate())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TLSOptions

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TLSOptions"))
	return Instance{classdb.TLSOptions(object)}
}

/*
Creates a TLS client configuration which validates certificates and their common names (fully qualified domain names).
You can specify a custom [param trusted_chain] of certification authorities (the default CA list will be used if [code]null[/code]), and optionally provide a [param common_name_override] if you expect the certificate to have a common name other than the server FQDN.
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
//go:nosplit
func (self class) Client(trusted_chain objects.X509Certificate, common_name_override gd.String) objects.TLSOptions {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(trusted_chain[0])[0])
	callframe.Arg(frame, pointers.Get(common_name_override))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.TLSOptions{classdb.TLSOptions(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Creates an [b]unsafe[/b] TLS client configuration where certificate validation is optional. You can optionally provide a valid [param trusted_chain], but the common name of the certificates will never be checked. Using this configuration for purposes other than testing [b]is not recommended[/b].
[b]Note:[/b] On the Web platform, TLS verification is always enforced against the CA list of the web browser. This is considered a security feature.
*/
//go:nosplit
func (self class) ClientUnsafe(trusted_chain objects.X509Certificate) objects.TLSOptions {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(trusted_chain[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_client_unsafe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.TLSOptions{classdb.TLSOptions(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Creates a TLS server configuration using the provided [param key] and [param certificate].
[b]Note:[/b] The [param certificate] should include the full certificate chain up to the signing CA (certificates file can be concatenated using a general purpose text editor).
*/
//go:nosplit
func (self class) Server(key objects.CryptoKey, certificate objects.X509Certificate) objects.TLSOptions {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key[0])[0])
	callframe.Arg(frame, pointers.Get(certificate[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.TLSOptions{classdb.TLSOptions(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if created with [method TLSOptions.server], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsServer() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_is_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if created with [method TLSOptions.client_unsafe], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsUnsafeClient() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_is_unsafe_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the common name (domain name) override specified when creating with [method TLSOptions.client].
*/
//go:nosplit
func (self class) GetCommonNameOverride() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_common_name_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the CA [X509Certificate] chain specified when creating with [method TLSOptions.client] or [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) GetTrustedCaChain() objects.X509Certificate {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_trusted_ca_chain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.X509Certificate{classdb.X509Certificate(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the [CryptoKey] specified when creating with [method TLSOptions.server].
*/
//go:nosplit
func (self class) GetPrivateKey() objects.CryptoKey {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_private_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.CryptoKey{classdb.CryptoKey(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the [X509Certificate] specified when creating with [method TLSOptions.server].
*/
//go:nosplit
func (self class) GetOwnCertificate() objects.X509Certificate {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TLSOptions.Bind_get_own_certificate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.X509Certificate{classdb.X509Certificate(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsTLSOptions() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTLSOptions() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("TLSOptions", func(ptr gd.Object) any { return [1]classdb.TLSOptions{classdb.TLSOptions(ptr)} })
}
