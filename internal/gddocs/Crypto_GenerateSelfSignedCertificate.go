/*
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
*/

package main

import "graphics.gd/classdb/Crypto"

func ExampleCryptoSelfSigned() {
	var crypto = Crypto.New()
	// Generate 4096 bits RSA key.
	var key = crypto.GenerateRsa(4096)
	// Generate self-signed certificate using the given key.
	var cert = crypto.MoreArgs().GenerateSelfSignedCertificate(key, "CN=example.com,O=A Game Company,C=IT", "", "")
	_ = cert
}
