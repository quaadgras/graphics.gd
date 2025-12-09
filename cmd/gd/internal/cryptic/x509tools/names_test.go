package x509tools

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNames(t *testing.T) {
	n := pkix.Name{
		CommonName:   "foo/bar",
		Organization: []string{"ham ", "\"spam\"", " eggs"},
		Locality:     []string{"north+southville"},
		Country:      []string{""},
		ExtraNames: []pkix.AttributeTypeAndValue{
			{Type: asn1.ObjectIdentifier{1, 2, 3, 4}, Value: "5678"},
			{Type: asn1.ObjectIdentifier{2, 5, 4, 17}, Value: "888"},
		},
	}
	name, err := asn1.Marshal(n.ToRDNSequence())
	require.NoError(t, err)
	assert.Equal(t, `/C=/L=north+southville/O=ham /O= eggs/O="spam"/CN=foo\/bar/1.2.3.4=5678/postalCode=888`, FormatPkixName(name, NameStyleOpenSsl))
	assert.Equal(t, `postalCode=888, 1.2.3.4=5678, CN=foo/bar, O="ham " + O=" eggs" + O="""spam""", L="north+southville", C=""`, FormatPkixName(name, NameStyleLdap))
	assert.Equal(t, `PostalCode=888, OID.1.2.3.4=5678, CN=foo/bar, O="ham " + O=" eggs" + O="""spam""", L="north+southville", C=""`, FormatPkixName(name, NameStyleMsOsco))
}

func TestBMP(t *testing.T) {
	str := "✔"
	bmp := ToBMPString(str)
	assert.Equal(t, []byte{0x27, 0x14}, bmp.Bytes)
}

func TestBMPName(t *testing.T) {
	n := pkix.RDNSequence{
		pkix.RelativeDistinguishedNameSET{pkix.AttributeTypeAndValue{
			Type: asn1.ObjectIdentifier{2, 5, 4, 10},
			Value: asn1.RawValue{
				Tag:   asn1.TagBMPString,
				Bytes: []byte{0x27, 0x14},
			},
		}},
	}
	blob, err := asn1.Marshal(n)
	require.NoError(t, err)
	assert.Equal(t, `O=✔`, FormatPkixName(blob, NameStyleLdap))
}
