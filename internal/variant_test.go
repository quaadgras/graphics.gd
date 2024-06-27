package gd_test

import (
	"testing"

	gd "grow.graphics/gd/internal"
	internal "grow.graphics/gd/internal"
)

func TestVariants(t *testing.T) {
	var godot = internal.NewContext(API)
	defer godot.End()

	var f = godot.Variant(gd.Float(3.14))

	if f.Interface(godot).(gd.Float) != 3.14 {
		t.Fail()
	}

	if f.Float() != 3.14 {
		t.Fail()
	}
}
