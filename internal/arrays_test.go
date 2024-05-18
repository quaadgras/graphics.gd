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
}
