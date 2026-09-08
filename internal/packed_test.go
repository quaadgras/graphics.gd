package gd_test

import (
	"testing"

	gd "graphics.gd/internal"
	"graphics.gd/variant/Packed"
	"graphics.gd/variant/Vector2"
)

func TestPacked(t *testing.T) {
	var array = gd.NewPackedInt32Array()
	array.Resize(2)
	if array.Size() != 2 {
		t.Fatal("packed array failed to resize")
	}
	array.SetIndex(0, 1)
	if array.Index(0) != 1 {
		t.Fatal("bad")
	}

	var sliced = gd.NewPackedInt64Slice([]int64{1, 2, 3})
	if sliced.Index(0) != 1 {
		t.Fatal("bad")
	}
}

// An empty packed array has no engine buffer, so its unsafe pointer is
// null: converting an empty Go slice into one (every omitted packed
// argument, like DrawPolygon's uvs) and reading an empty one back must
// both be no-ops rather than nil dereferences. This only bit the wasm
// builds (the web export and the hot-reload guest), where the copy goes
// through host memory accessors instead of a plain copy.
func TestPackedEmpty(t *testing.T) {
	var points = gd.InternalPacked[gd.PackedVector2Array, gd.Vector2](Packed.New[Vector2.XY]())
	if points.Size() != 0 {
		t.Fatal("empty packed array should stay empty")
	}
	if got := points.AsSlice(); len(got) != 0 {
		t.Fatal("empty packed array should read back empty")
	}
	var empty = gd.NewPackedVector2Slice(nil)
	if empty.Size() != 0 || len(empty.AsSlice()) != 0 {
		t.Fatal("empty packed slice should stay empty")
	}
	var bytes = gd.NewPackedByteSlice(nil)
	if bytes.Size() != 0 || len(bytes.Bytes()) != 0 {
		t.Fatal("empty packed byte slice should stay empty")
	}
	// A refilled array still copies in bulk after the empty round trip.
	var filled = gd.InternalPacked[gd.PackedVector2Array, gd.Vector2](Packed.New(Vector2.XY{X: 1, Y: 2}, Vector2.XY{X: 3, Y: 4}))
	if got := filled.AsSlice(); len(got) != 2 || got[1] != (gd.Vector2{X: 3, Y: 4}) {
		t.Fatal("packed array bulk copy lost data", got)
	}
}
