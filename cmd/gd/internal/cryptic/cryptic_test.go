package cryptic_test

import (
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"graphics.gd/cmd/gd/internal/cryptic"
)

func TestDeterminisim(t *testing.T) {
	key1, cert1, err := cryptic.DeterministicCertificate("testpass", "testsalt")
	if err != nil {
		t.Error(err)
		return
	}
	key2, cert2, err := cryptic.DeterministicCertificate("testpass", "testsalt")
	if err != nil {
		t.Error(err)
		return
	}
	hk1 := sha256.Sum256(x509.MarshalPKCS1PrivateKey(key1))
	hk2 := sha256.Sum256(x509.MarshalPKCS1PrivateKey(key2))
	if hk1 != hk2 {
		t.Fatal("Private keys differ")
	}
	h1 := sha256.Sum256(cert1.Raw)
	h2 := sha256.Sum256(cert2.Raw)
	fmt.Println("Hash1:", fmt.Sprintf("%x", h1))
	fmt.Println("Hash2:", fmt.Sprintf("%x", h2))
	if h1 != h2 {
		fmt.Println(CompareCertificates(cert1, cert2))
		t.Fatal("Certificates differ")
	}

}

// CompareCertificates compares two X.509 certificates and returns a string describing their differences.
func CompareCertificates(cert1, cert2 *x509.Certificate) (string, error) {
	if cert1 == nil || cert2 == nil {
		return "", fmt.Errorf("one or both certificates are nil")
	}

	var differences strings.Builder

	// Compare basic fields
	if cert1.Version != cert2.Version {
		differences.WriteString(fmt.Sprintf("Version differs: %d vs %d\n", cert1.Version, cert2.Version))
	}
	if cert1.SerialNumber.Cmp(cert2.SerialNumber) != 0 {
		differences.WriteString(fmt.Sprintf("SerialNumber differs: %s vs %s\n", cert1.SerialNumber, cert2.SerialNumber))
	}
	if cert1.SignatureAlgorithm != cert2.SignatureAlgorithm {
		differences.WriteString(fmt.Sprintf("SignatureAlgorithm differs: %s vs %s\n", cert1.SignatureAlgorithm.String(), cert2.SignatureAlgorithm.String()))
	}
	if !cert1.NotBefore.Equal(cert2.NotBefore) {
		differences.WriteString(fmt.Sprintf("NotBefore differs: %s vs %s\n", cert1.NotBefore, cert2.NotBefore))
	}
	if !cert1.NotAfter.Equal(cert2.NotAfter) {
		differences.WriteString(fmt.Sprintf("NotAfter differs: %s vs %s\n", cert1.NotAfter, cert2.NotAfter))
	}
	if cert1.IsCA != cert2.IsCA {
		differences.WriteString(fmt.Sprintf("IsCA differs: %v vs %v\n", cert1.IsCA, cert2.IsCA))
	}
	if cert1.KeyUsage != cert2.KeyUsage {
		differences.WriteString(fmt.Sprintf("KeyUsage differs: %d vs %d\n", cert1.KeyUsage, cert2.KeyUsage))
	}

	// Compare Subject and Issuer
	if !reflect.DeepEqual(cert1.Subject, cert2.Subject) {
		differences.WriteString(fmt.Sprintf("Subject differs: %v vs %v\n", cert1.Subject, cert2.Subject))
	}
	if !reflect.DeepEqual(cert1.Issuer, cert2.Issuer) {
		differences.WriteString(fmt.Sprintf("Issuer differs: %v vs %v\n", cert1.Issuer, cert2.Issuer))
	}

	// Compare DNSNames, EmailAddresses, IPAddresses, URIs
	if !reflect.DeepEqual(cert1.DNSNames, cert2.DNSNames) {
		differences.WriteString(fmt.Sprintf("DNSNames differ: %v vs %v\n", cert1.DNSNames, cert2.DNSNames))
	}
	if !reflect.DeepEqual(cert1.EmailAddresses, cert2.EmailAddresses) {
		differences.WriteString(fmt.Sprintf("EmailAddresses differ: %v vs %v\n", cert1.EmailAddresses, cert2.EmailAddresses))
	}
	if !reflect.DeepEqual(cert1.IPAddresses, cert2.IPAddresses) {
		differences.WriteString(fmt.Sprintf("IPAddresses differ: %v vs %v\n", cert1.IPAddresses, cert2.IPAddresses))
	}
	if !reflect.DeepEqual(cert1.URIs, cert2.URIs) {
		differences.WriteString(fmt.Sprintf("URIs differ: %v vs %v\n", cert1.URIs, cert2.URIs))
	}

	// Compare Extensions
	if len(cert1.Extensions) != len(cert2.Extensions) {
		differences.WriteString(fmt.Sprintf("Number of extensions differs: %d vs %d\n", len(cert1.Extensions), len(cert2.Extensions)))
	} else {
		for i := range cert1.Extensions {
			if !reflect.DeepEqual(cert1.Extensions[i], cert2.Extensions[i]) {
				differences.WriteString(fmt.Sprintf("Extension #%d differs: %v vs %v\n", i, cert1.Extensions[i], cert2.Extensions[i]))
			}
		}
	}

	// Compare Signature
	if !reflect.DeepEqual(cert1.Signature, cert2.Signature) {
		differences.WriteString("Signature differs\n")
	}

	// Compare Public Key (basic check)
	if !reflect.DeepEqual(cert1.PublicKeyAlgorithm, cert2.PublicKeyAlgorithm) {
		differences.WriteString(fmt.Sprintf("PublicKeyAlgorithm differs: %s vs %s\n", cert1.PublicKeyAlgorithm, cert2.PublicKeyAlgorithm))
	}
	if !reflect.DeepEqual(cert1.PublicKey, cert2.PublicKey) {
		differences.WriteString("PublicKey differs\n")
	}

	if differences.Len() == 0 {
		return "No differences found", nil
	}
	return differences.String(), nil
}
