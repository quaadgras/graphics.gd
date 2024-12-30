package gd_test

import (
	"testing"

	"graphics.gd"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Packed"
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
	var float32s Packed.Float32Array
	float32s.Resize(2)
	float32s.Set(0, 1)
	float32s.Set(1, 2)
	if float32s.Index(0) != 1 {
		t.Error("expected 1")
	}
	if float32s.Index(1) != 2 {
		t.Error("expected 2")
	}
}
