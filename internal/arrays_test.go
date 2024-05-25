package gd_test

import (
	"testing"

	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
)

func TestArrays(t *testing.T) {
	godot := internal.NewContext(API)
	defer godot.End()

	var numbers = gd.NewArrayOf[gd.Int](godot)
	numbers.Append(1)

	if numbers.Index(godot, 0) != 1 {
		t.Error("expected 1")
	}

	var packed = godot.PackedVector3Array()
	packed.Append(gd.Vector3{1, 2, 3})

	if packed.Index(0) != (gd.Vector3{1, 2, 3}) {
		t.Error("expected 1, 2, 3")
	}
}
