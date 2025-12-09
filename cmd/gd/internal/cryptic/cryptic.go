// Package cryptic provides functions for generating deterministic RSA keys and self-signed certificates.
//
// Sub-packages for signing .jar files have been primarily sourced from the Go relic signing server.
package cryptic

import (
	stdrsa "crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/binary"
	"math/rand"
	"time"

	"graphics.gd/cmd/gd/internal/cryptic/rsa" // deterministic rsa
	"graphics.gd/cmd/gd/internal/project"

	"golang.org/x/crypto/pbkdf2"
)

type reader struct {
	r *rand.Rand
}

func (rr *reader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(rr.r.Intn(256))
	}
	return len(p), nil
}

func DeterministicCertificate(passphrase, salt string) (*stdrsa.PrivateKey, *x509.Certificate, error) {
	dk := pbkdf2.Key([]byte(passphrase), []byte(salt), 100000, 64, sha256.New)
	seed := int64(binary.LittleEndian.Uint64(dk))
	r := rand.New(rand.NewSource(seed))
	key, err := rsa.GenerateKey(&reader{r: r}, 4096)
	if err != nil {
		return nil, nil, err
	}
	// Compute SubjectKeyId as SHA-1 of the DER-encoded public key
	pubDER, err := x509.MarshalPKIXPublicKey(key.Public())
	if err != nil {
		return nil, nil, err
	}
	ski := sha1.Sum(pubDER)
	template := x509.Certificate{
		Subject: pkix.Name{
			CommonName: project.Name,
		},
		Issuer: pkix.Name{
			CommonName: project.Name,
		},
		NotBefore:          time.Date(2025, time.December, 9, 0, 0, 0, 0, time.UTC),
		NotAfter:           time.Date(2055, time.January, 1, 0, 0, 0, 0, time.UTC),
		SubjectKeyId:       ski[:],
		SignatureAlgorithm: x509.SHA256WithRSA,
	}
	derBytes, err := x509.CreateCertificate(&reader{r: r}, &template, &template, key.Public(), key)
	if err != nil {
		return nil, nil, err
	}
	cert, err := x509.ParseCertificate(derBytes)
	if err != nil {
		return nil, nil, err
	}
	return (*stdrsa.PrivateKey)(key), cert, nil
}
