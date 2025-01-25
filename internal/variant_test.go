package gd_test

import (
	"testing"

	"graphics.gd/variant"
)

func TestVariants(t *testing.T) {
	var f = variant.New(3.14)
	if f.Float64() != 3.14 {
		t.Fail()
	}
}
