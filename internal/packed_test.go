package gd_test

import (
	"testing"

	gd "graphics.gd/internal"
)

func TestPacked(t *testing.T) {
	var array = gd.NewPackedInt32Array()
	array.Resize(2)
	if array.Size() != 2 {
		t.Fatal("packed array failed to resize")
	}
	array.Set(0, 1)
	if array.Index(0) != 1 {
		t.Fatal("bad")
	}

	var sliced = gd.NewPackedInt64Slice([]int64{1, 2, 3})
	if sliced.Index(0) != 1 {
		t.Fatal("bad")
	}
}
