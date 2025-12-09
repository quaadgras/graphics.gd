package pkcs7

import (
	"encoding/asn1"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// marshal and unmarshal so FullBytes is set
func roundTrip(t *testing.T, l AttributeList) AttributeList {
	t.Helper()
	raw, err := marshalUnsortedSet(l)
	require.NoError(t, err)
	var l2 AttributeList
	_, err = asn1.UnmarshalWithParams(raw, &l2, "set")
	require.NoError(t, err)
	return l2
}

func TestAttributeList(t *testing.T) {
	var l AttributeList
	assert.False(t, l.Exists(OidAttributeSigningTime))
	a := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.NoError(t, l.Add(OidAttributeSigningTime, a))
	ll := roundTrip(t, l)
	assert.True(t, ll.Exists(OidAttributeSigningTime))
	var x time.Time
	if assert.NoError(t, ll.GetOne(OidAttributeSigningTime, &x)) {
		assert.Equal(t, a, x)
	}

	b := a.AddDate(0, 0, 1)
	assert.NoError(t, l.Add(OidAttributeSigningTime, b))
	ll = roundTrip(t, l)
	assert.Error(t, ll.GetOne(OidAttributeSigningTime, &x))
	var times []time.Time
	if assert.NoError(t, ll.GetAll(OidAttributeSigningTime, &times)) {
		assert.Equal(t, []time.Time{a, b}, times)
	}
}
