package gd_test

import (
	"testing"

	"graphics.gd/classdb/Image"
)

func TestStatic(t *testing.T) {
	var image Image.Instance = Image.Create(1, 1, false, Image.FormatRgb8)
	if image.GetWidth() != 1 {
		t.Fail()
	}
}
