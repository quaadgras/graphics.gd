package gd_test

import (
	"testing"

	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
)

func Something[T any]() T {
	var zero T
	return zero
}

func TestStatic(t *testing.T) {
	var godot = internal.NewContext(API)
	defer godot.End()

	image := gd.Image.Create(gd.Image{}, godot, 1, 1, false, gd.ImageFormatRgb8)
	if image.GetWidth() != 1 {
		t.Fail()
	}
}
