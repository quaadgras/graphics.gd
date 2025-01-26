// Package Crypto provides methods for working with Crypto object instances.
package Crypto

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
The Crypto class provides access to advanced cryptographic functionalities.
Currently, this includes asymmetric key encryption/decryption, signing/verification, and generating cryptographically secure random bytes, RSA keys, HMAC digests, and self-signed [X509Certificate]s.
[codeblocks]
[gdscript]
var crypto = Crypto.new()

# Generate new RSA key.
var key = crypto.generate_rsa(4096)

# Generate new self-signed certificate with the given key.
var cert = crypto.generate_self_signed_certificate(key, "CN=mydomain.com,O=My Game Company,C=IT")

# Save key and certificate in the user folder.
key.save("user://generated.key")
cert.save("user://generated.crt")

# Encryption
var data = "Some data"
var encrypted = crypto.encrypt(key, data.to_utf8_buffer())

# Decryption
var decrypted = crypto.decrypt(key, encrypted)

# Signing
var signature = crypto.sign(HashingContext.HASH_SHA256, data.sha256_buffer(), key)

# Verifying
var verified = crypto.verify(HashingContext.HASH_SHA256, data.sha256_buffer(), signature, key)

# Checks
assert(verified)
assert(data.to_utf8_buffer() == decrypted)
[/gdscript]
[csharp]
using Godot;
using System.Diagnostics;

Crypto crypto = new Crypto();

// Generate new RSA key.
CryptoKey key = crypto.GenerateRsa(4096);

// Generate new self-signed certificate with the given key.
X509Certificate cert = crypto.GenerateSelfSignedCertificate(key, "CN=mydomain.com,O=My Game Company,C=IT");

// Save key and certificate in the user folder.
key.Save("user://generated.key");
cert.Save("user://generated.crt");

// Encryption
string data = "Some data";
byte[] encrypted = crypto.Encrypt(key, data.ToUtf8Buffer());

// Decryption
byte[] decrypted = crypto.Decrypt(key, encrypted);

// Signing
byte[] signature = crypto.Sign(HashingContext.HashType.Sha256, Data.Sha256Buffer(), key);

// Verifying
bool verified = crypto.Verify(HashingContext.HashType.Sha256, Data.Sha256Buffer(), signature, key);

// Checks
Debug.Assert(verified);
Debug.Assert(data.ToUtf8Buffer() == decrypted);
[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.Crypto

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCrypto() Instance
}

/*
Generates a [PackedByteArray] of cryptographically secure random bytes with given [param size].
*/
func (self Instance) GenerateRandomBytes(size int) []byte { //gd:Crypto.generate_random_bytes
	return []byte(class(self).GenerateRandomBytes(gd.Int(size)).Bytes())
}

/*
Generates an RSA [CryptoKey] that can be used for creating self-signed certificates and passed to [method StreamPeerTLS.accept_stream].
*/
func (self Instance) GenerateRsa(size int) [1]gdclass.CryptoKey { //gd:Crypto.generate_rsa
	return [1]gdclass.CryptoKey(class(self).GenerateRsa(gd.Int(size)))
}

/*
Generates a self-signed [X509Certificate] from the given [CryptoKey] and [param issuer_name]. The certificate validity will be defined by [param not_before] and [param not_after] (first valid date and last valid date). The [param issuer_name] must contain at least "CN=" (common name, i.e. the domain name), "O=" (organization, i.e. your company name), "C=" (country, i.e. 2 lettered ISO-3166 code of the country the organization is based in).
A small example to generate an RSA key and an X509 self-signed certificate.
[codeblocks]
[gdscript]
var crypto = Crypto.new()
# Generate 4096 bits RSA key.
var key = crypto.generate_rsa(4096)
# Generate self-signed certificate using the given key.
var cert = crypto.generate_self_signed_certificate(key, "CN=example.com,O=A Game Company,C=IT")
[/gdscript]
[csharp]
var crypto = new Crypto();
// Generate 4096 bits RSA key.
CryptoKey key = crypto.GenerateRsa(4096);
// Generate self-signed certificate using the given key.
X509Certificate cert = crypto.GenerateSelfSignedCertificate(key, "CN=mydomain.com,O=My Game Company,C=IT");
[/csharp]
[/codeblocks]
*/
func (self Instance) GenerateSelfSignedCertificate(key [1]gdclass.CryptoKey) [1]gdclass.X509Certificate { //gd:Crypto.generate_self_signed_certificate
	return [1]gdclass.X509Certificate(class(self).GenerateSelfSignedCertificate(key, gd.NewString("CN=myserver,O=myorganisation,C=IT"), gd.NewString("20140101000000"), gd.NewString("20340101000000")))
}

/*
Sign a given [param hash] of type [param hash_type] with the provided private [param key].
*/
func (self Instance) Sign(hash_type gdclass.HashingContextHashType, hash []byte, key [1]gdclass.CryptoKey) []byte { //gd:Crypto.sign
	return []byte(class(self).Sign(hash_type, gd.NewPackedByteSlice(hash), key).Bytes())
}

/*
Verify that a given [param signature] for [param hash] of type [param hash_type] against the provided public [param key].
*/
func (self Instance) Verify(hash_type gdclass.HashingContextHashType, hash []byte, signature []byte, key [1]gdclass.CryptoKey) bool { //gd:Crypto.verify
	return bool(class(self).Verify(hash_type, gd.NewPackedByteSlice(hash), gd.NewPackedByteSlice(signature), key))
}

/*
Encrypt the given [param plaintext] with the provided public [param key].
[b]Note:[/b] The maximum size of accepted plaintext is limited by the key size.
*/
func (self Instance) Encrypt(key [1]gdclass.CryptoKey, plaintext []byte) []byte { //gd:Crypto.encrypt
	return []byte(class(self).Encrypt(key, gd.NewPackedByteSlice(plaintext)).Bytes())
}

/*
Decrypt the given [param ciphertext] with the provided private [param key].
[b]Note:[/b] The maximum size of accepted ciphertext is limited by the key size.
*/
func (self Instance) Decrypt(key [1]gdclass.CryptoKey, ciphertext []byte) []byte { //gd:Crypto.decrypt
	return []byte(class(self).Decrypt(key, gd.NewPackedByteSlice(ciphertext)).Bytes())
}

/*
Generates an [url=https://en.wikipedia.org/wiki/HMAC]HMAC[/url] digest of [param msg] using [param key]. The [param hash_type] parameter is the hashing algorithm that is used for the inner and outer hashes.
Currently, only [constant HashingContext.HASH_SHA256] and [constant HashingContext.HASH_SHA1] are supported.
*/
func (self Instance) HmacDigest(hash_type gdclass.HashingContextHashType, key []byte, msg []byte) []byte { //gd:Crypto.hmac_digest
	return []byte(class(self).HmacDigest(hash_type, gd.NewPackedByteSlice(key), gd.NewPackedByteSlice(msg)).Bytes())
}

/*
Compares two [PackedByteArray]s for equality without leaking timing information in order to prevent timing attacks.
See [url=https://paragonie.com/blog/2015/11/preventing-timing-attacks-on-string-comparison-with-double-hmac-strategy]this blog post[/url] for more information.
*/
func (self Instance) ConstantTimeCompare(trusted []byte, received []byte) bool { //gd:Crypto.constant_time_compare
	return bool(class(self).ConstantTimeCompare(gd.NewPackedByteSlice(trusted), gd.NewPackedByteSlice(received)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Crypto

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Crypto"))
	casted := Instance{*(*gdclass.Crypto)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Generates a [PackedByteArray] of cryptographically secure random bytes with given [param size].
*/
//go:nosplit
func (self class) GenerateRandomBytes(size gd.Int) gd.PackedByteArray { //gd:Crypto.generate_random_bytes
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_generate_random_bytes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Generates an RSA [CryptoKey] that can be used for creating self-signed certificates and passed to [method StreamPeerTLS.accept_stream].
*/
//go:nosplit
func (self class) GenerateRsa(size gd.Int) [1]gdclass.CryptoKey { //gd:Crypto.generate_rsa
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_generate_rsa, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CryptoKey{gd.PointerWithOwnershipTransferredToGo[gdclass.CryptoKey](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Generates a self-signed [X509Certificate] from the given [CryptoKey] and [param issuer_name]. The certificate validity will be defined by [param not_before] and [param not_after] (first valid date and last valid date). The [param issuer_name] must contain at least "CN=" (common name, i.e. the domain name), "O=" (organization, i.e. your company name), "C=" (country, i.e. 2 lettered ISO-3166 code of the country the organization is based in).
A small example to generate an RSA key and an X509 self-signed certificate.
[codeblocks]
[gdscript]
var crypto = Crypto.new()
# Generate 4096 bits RSA key.
var key = crypto.generate_rsa(4096)
# Generate self-signed certificate using the given key.
var cert = crypto.generate_self_signed_certificate(key, "CN=example.com,O=A Game Company,C=IT")
[/gdscript]
[csharp]
var crypto = new Crypto();
// Generate 4096 bits RSA key.
CryptoKey key = crypto.GenerateRsa(4096);
// Generate self-signed certificate using the given key.
X509Certificate cert = crypto.GenerateSelfSignedCertificate(key, "CN=mydomain.com,O=My Game Company,C=IT");
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GenerateSelfSignedCertificate(key [1]gdclass.CryptoKey, issuer_name gd.String, not_before gd.String, not_after gd.String) [1]gdclass.X509Certificate { //gd:Crypto.generate_self_signed_certificate
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key[0])[0])
	callframe.Arg(frame, pointers.Get(issuer_name))
	callframe.Arg(frame, pointers.Get(not_before))
	callframe.Arg(frame, pointers.Get(not_after))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_generate_self_signed_certificate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.X509Certificate{gd.PointerWithOwnershipTransferredToGo[gdclass.X509Certificate](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sign a given [param hash] of type [param hash_type] with the provided private [param key].
*/
//go:nosplit
func (self class) Sign(hash_type gdclass.HashingContextHashType, hash gd.PackedByteArray, key [1]gdclass.CryptoKey) gd.PackedByteArray { //gd:Crypto.sign
	var frame = callframe.New()
	callframe.Arg(frame, hash_type)
	callframe.Arg(frame, pointers.Get(hash))
	callframe.Arg(frame, pointers.Get(key[0])[0])
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_sign, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Verify that a given [param signature] for [param hash] of type [param hash_type] against the provided public [param key].
*/
//go:nosplit
func (self class) Verify(hash_type gdclass.HashingContextHashType, hash gd.PackedByteArray, signature gd.PackedByteArray, key [1]gdclass.CryptoKey) bool { //gd:Crypto.verify
	var frame = callframe.New()
	callframe.Arg(frame, hash_type)
	callframe.Arg(frame, pointers.Get(hash))
	callframe.Arg(frame, pointers.Get(signature))
	callframe.Arg(frame, pointers.Get(key[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_verify, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Encrypt the given [param plaintext] with the provided public [param key].
[b]Note:[/b] The maximum size of accepted plaintext is limited by the key size.
*/
//go:nosplit
func (self class) Encrypt(key [1]gdclass.CryptoKey, plaintext gd.PackedByteArray) gd.PackedByteArray { //gd:Crypto.encrypt
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key[0])[0])
	callframe.Arg(frame, pointers.Get(plaintext))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_encrypt, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Decrypt the given [param ciphertext] with the provided private [param key].
[b]Note:[/b] The maximum size of accepted ciphertext is limited by the key size.
*/
//go:nosplit
func (self class) Decrypt(key [1]gdclass.CryptoKey, ciphertext gd.PackedByteArray) gd.PackedByteArray { //gd:Crypto.decrypt
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key[0])[0])
	callframe.Arg(frame, pointers.Get(ciphertext))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_decrypt, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Generates an [url=https://en.wikipedia.org/wiki/HMAC]HMAC[/url] digest of [param msg] using [param key]. The [param hash_type] parameter is the hashing algorithm that is used for the inner and outer hashes.
Currently, only [constant HashingContext.HASH_SHA256] and [constant HashingContext.HASH_SHA1] are supported.
*/
//go:nosplit
func (self class) HmacDigest(hash_type gdclass.HashingContextHashType, key gd.PackedByteArray, msg gd.PackedByteArray) gd.PackedByteArray { //gd:Crypto.hmac_digest
	var frame = callframe.New()
	callframe.Arg(frame, hash_type)
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, pointers.Get(msg))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_hmac_digest, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Compares two [PackedByteArray]s for equality without leaking timing information in order to prevent timing attacks.
See [url=https://paragonie.com/blog/2015/11/preventing-timing-attacks-on-string-comparison-with-double-hmac-strategy]this blog post[/url] for more information.
*/
//go:nosplit
func (self class) ConstantTimeCompare(trusted gd.PackedByteArray, received gd.PackedByteArray) bool { //gd:Crypto.constant_time_compare
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(trusted))
	callframe.Arg(frame, pointers.Get(received))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_constant_time_compare, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCrypto() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCrypto() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Crypto", func(ptr gd.Object) any { return [1]gdclass.Crypto{*(*gdclass.Crypto)(unsafe.Pointer(&ptr))} })
}
