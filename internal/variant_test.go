package gd_test

import (
	"testing"

	"grow.graphics/gd/variant"
)

func TestVariants(t *testing.T) {
	var f = variant.New(3.14)
	if variant.Get[float64](f) != 3.14 {
		t.Fail()
	}
	if f.Float() != 3.14 {
		t.Fail()
	}
}
