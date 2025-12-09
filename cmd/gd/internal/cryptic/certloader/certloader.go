//
// Copyright (c) SAS Institute Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package certloader

import (
	"bytes"
	"crypto"
	"crypto/tls"
	"crypto/x509"

	"graphics.gd/cmd/gd/internal/cryptic/pkcs9"
)

const asn1Magic = 0x30 // weak but good enough?
var pkcs7SignedData = []byte{0x06, 0x09, 0x2A, 0x86, 0x48, 0x86, 0xF7, 0x0D, 0x01, 0x07, 0x02}

// A bundle of X509 certificate chain and/or PGP certificate, with optional private key
type Certificate struct {
	Leaf         *x509.Certificate
	Certificates []*x509.Certificate
	PrivateKey   crypto.PrivateKey
	Timestamper  pkcs9.Timestamper
	KeyName      string
}

// Return the X509 certificates in the chain up to, but not including, the root CA certificate
func (s *Certificate) Chain() []*x509.Certificate {
	var chain []*x509.Certificate
	if s.Leaf != nil {
		// ensure leaf comes first
		chain = append(chain, s.Leaf)
	}
	for i, cert := range s.Certificates {
		if i > 0 && bytes.Equal(cert.RawIssuer, cert.RawSubject) {
			// omit root CA
			continue
		} else if cert == s.Leaf {
			// already in list
			continue
		}
		chain = append(chain, cert)
	}
	return chain
}

// Return the certificate that issued the leaf certificate
func (s *Certificate) Issuer() *x509.Certificate {
	if s.Leaf == nil {
		return nil
	}
	for _, cert := range s.Certificates {
		if bytes.Equal(cert.RawSubject, s.Leaf.RawIssuer) {
			return cert
		}
	}
	return nil
}

// Return the private key in the form of a crypto.Signer
func (s *Certificate) Signer() crypto.Signer {
	if s.PrivateKey == nil {
		return nil
	}
	return s.PrivateKey.(crypto.Signer)
}

// Return a tls.Certificate structure containing the X509 certificate chain and
// private key
func (s *Certificate) TLS() tls.Certificate {
	var raw [][]byte
	for _, cert := range s.Certificates {
		raw = append(raw, cert.Raw)
	}
	return tls.Certificate{Leaf: s.Leaf, Certificate: raw, PrivateKey: s.PrivateKey}
}
