package x509tools

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	packed := []byte{1, 0, 2, 0}
	sig, err := UnpackEcdsaSignature(packed)
	require.NoError(t, err)
	assert.Equal(t, int64(256), sig.R.Int64())
	assert.Equal(t, int64(512), sig.S.Int64())
	invalid := []byte{1, 2, 3}
	_, err = UnpackEcdsaSignature(invalid)
	require.Error(t, err)
}

func TestUnmarshal(t *testing.T) {
	packed := []byte{0x30, 0x8, 0x2, 0x2, 0x2, 0x0, 0x2, 0x2, 0x0, 0xff}
	sig, err := UnmarshalEcdsaSignature(packed)
	require.NoError(t, err)
	assert.Equal(t, int64(512), sig.R.Int64())
	assert.Equal(t, int64(255), sig.S.Int64())
	invalid := []byte{1, 2, 3}
	_, err = UnmarshalEcdsaSignature(invalid)
	require.Error(t, err)
}

func TestMarshal(t *testing.T) {
	sig := EcdsaSignature{
		R: big.NewInt(512),
		S: big.NewInt(255),
	}
	assert.Equal(t, []byte{0x30, 0x8, 0x2, 0x2, 0x2, 0x0, 0x2, 0x2, 0x0, 0xff}, sig.Marshal())

	// test both ways around to ensure it's padded correctly
	assert.Equal(t, []byte{2, 0, 0, 255}, sig.Pack())
	sig.R, sig.S = sig.S, sig.R
	assert.Equal(t, []byte{0, 255, 2, 0}, sig.Pack())
}
