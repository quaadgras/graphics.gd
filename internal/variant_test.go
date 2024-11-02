package gd_test

import (
	"testing"

	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/variant"
)

func TestVariants(t *testing.T) {
	var f = variant.New(3.14)
	if f.Interface().(gd.Float) != 3.14 {
		t.Fail()
	}
	if f.Float() != 3.14 {
		t.Fail()
	}
}
