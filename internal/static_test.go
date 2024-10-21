package gd_test

import (
	"testing"

	internal "grow.graphics/gd/internal"
	"grow.graphics/gd/object/Image"
)

func TestStatic(t *testing.T) {
	var godot = internal.NewContext(API)
	defer godot.End()

	var image Image.Expert = Image.Expert.Create(Image.Expert{}, godot, 1, 1, false, Image.FormatRgb8)
	if image.GetWidth() != 1 {
		t.Fail()
	}
}
