package gd_test

import (
	"testing"

	"grow.graphics/gd"
	"grow.graphics/gd/variant/Array"
	"grow.graphics/gd/variant/Packed"
)

func TestArrays(t *testing.T) {
	var numbers = Array.New[int]()
	numbers.Append(1)
	if numbers.Index(0) != 1 {
		t.Error("expected 1")
	}
	var packed Packed.Vector3Array
	packed.Append(gd.Vector3{1, 2, 3})
	if packed.Index(0) != (gd.Vector3{1, 2, 3}) {
		t.Error("expected 1, 2, 3")
	}
}
