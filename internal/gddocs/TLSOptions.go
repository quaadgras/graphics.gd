/*
[gdscript]
# Create a TLS client configuration which uses our custom trusted CA chain.
var client_trusted_cas = load("res://my_trusted_cas.crt")
var client_tls_options = TLSOptions.client(client_trusted_cas)

# Create a TLS server configuration.
var server_certs = load("res://my_server_cas.crt")
var server_key = load("res://my_server_key.key")
var server_tls_options = TLSOptions.server(server_key, server_certs)
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/TLSOptions"
	"graphics.gd/classdb/X509Certificate"
)

func ExampleTLSOptions() {
	var client_trusted_cas = Resource.Load[X509Certificate.Instance]("res://my_trusted_cas.crt")
	var client_tls_options = TLSOptions.Client(client_trusted_cas, "")
	_ = client_tls_options
}
