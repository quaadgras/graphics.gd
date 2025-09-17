/*
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
*/

package main

import (
	"crypto/sha256"

	"graphics.gd/classdb/Crypto"
	"graphics.gd/classdb/HashingContext"
)

func ExampleCrypto() {
	crypto := Crypto.New()

	// Generate new RSA key.
	key := crypto.GenerateRsa(4096)
	// Generate new self-signed certificate with the given key.
	cert := crypto.MoreArgs().GenerateSelfSignedCertificate(key, "CN=mydomain.com,O=My Game Company,C=IT", "20140101000000", "20340101000000")
	// Save key and certificate in the user folder.
	key.Save("user://generated.key")
	cert.Save("user://generated.crt")
	// Encryption
	data := "Some data"
	encrypted := crypto.Encrypt(key, []byte(data))
	// Decryption
	decrypted := crypto.Decrypt(key, encrypted)
	// Signing
	var hash = sha256.Sum256([]byte(data))
	signature := crypto.Sign(HashingContext.HashSha256, hash[:], key)
	// Verifying
	verified := crypto.Verify(HashingContext.HashSha256, hash[:], signature, key)
	// Checks
	if !verified {
		panic("verification failed")
	}
	if string(decrypted) != data {
		panic("decryption failed")
	}
}
