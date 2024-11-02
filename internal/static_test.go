package gd_test

import (
	"testing"

	"grow.graphics/gd/objects/Image"
)

func TestStatic(t *testing.T) {
	var image Image.Instance = Image.Instance.Create(Image.Instance{}, 1, 1, false, Image.FormatRgb8)
	if image.GetWidth() != 1 {
		t.Fail()
	}
}
