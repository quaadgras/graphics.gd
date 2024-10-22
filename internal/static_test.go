package gd_test

import (
	"testing"

	"grow.graphics/gd/gdclass/Image"
	internal "grow.graphics/gd/internal"
)

func TestStatic(t *testing.T) {
	var godot = internal.NewContext(API)
	defer godot.End()

	var image Image.Expert = Image.Expert.Create(Image.Expert{}, godot, 1, 1, false, Image.FormatRgb8)
	if image.GetWidth() != 1 {
		t.Fail()
	}
}
