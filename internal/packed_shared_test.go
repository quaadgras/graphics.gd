package gd_test

import (
	"testing"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Packed"
	"graphics.gd/variant/String"
	"graphics.gd/variant/Vector3"
)

// A packed array read out of a variant shares the variant's buffer, so
// the bulk read must take the engine's CONST index. The writable one
// copies the array on write into a temporary the caller never sees: the
// handle is left pointing at the variant's own buffer with its share
// already given up, and freeing the handle — which the read does, as
// soon as it has the bytes — takes the variant's array away with it.
// Nothing shows at the time; the fault comes a frame or two later, when
// [pointers.Cycle] frees the variant and the engine destroys an array
// that is already gone. Mesh.SurfaceGetArrays, which reads a dozen
// packed arrays out of one engine array, brought a whole realm down
// this way.
//
// Reading the same variant twice is enough to show it without the
// crash: with the writable index the second read comes back empty.
func TestBulkReadLeavesASharedPackedArrayAlone(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		want := []Vector3.XYZ{{X: 1}, {Y: 2}, {Z: 3}}
		var array Array.Any
		array.Append(variant.New(Packed.New(want...)))
		held := gd.InternalArray(array).Index(0)
		for read := 1; read <= 3; read++ {
			points, ok := held.ConvenientInterface().([]gd.Vector3)
			if !ok {
				t.Fatalf("read %d: the variant came back as %T", read, held.ConvenientInterface())
			}
			if len(points) != len(want) {
				t.Fatalf("read %d: %d points, want %d — the read took the variant's array with it",
					read, len(points), len(want))
			}
			for i, p := range points {
				if p != (gd.Vector3{X: want[i].X, Y: want[i].Y, Z: want[i].Z}) {
					t.Fatalf("read %d: point %d is %v, want %v", read, i, p, want[i])
				}
			}
		}
		// And every variant the reads made is freed without a fault.
		for range 4 {
			pointers.Cycle()
		}
	})
}

// Writing to a packed array that shares its buffer must unshare it first,
// the way the engine's own copy-on-write does, and must leave the handle
// holding the buffer it was moved to: the variant keeps what it had, the
// view sees the write, and both are freed without a fault.
func TestSetIndexUnsharesFirst(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		held := gd.NewVariant(gd.NewPackedInt32Slice([]int32{1, 2, 3}))
		view := held.Interface().(Packed.Array[int32])
		before := gd.InternalPacked[gd.PackedInt32Array](view) // a handle made before the write
		view.SetIndex(0, 99)
		if got := view.Index(0); got != 99 {
			t.Fatalf("the view reads %d after writing 99", got)
		}
		if got := held.ConvenientInterface().([]int32); len(got) != 3 || got[0] != 1 {
			t.Fatalf("the write went through to the variant's own array: %v", got)
		}
		// Freeing through the older handle frees the buffer the array was
		// moved to, not the one it shared with the variant back then.
		before.Free()
		if got := held.ConvenientInterface().([]int32); len(got) != 3 || got[0] != 1 {
			t.Fatalf("freeing the view took the variant's array with it: %v", got)
		}

		names := gd.NewVariant(gd.NewPackedStringSlice([]string{"a", "b"}))
		renamed := names.Interface().(Packed.Strings)
		renamed.SetIndex(1, String.New("z"))
		if got := renamed.Index(1).String(); got != "z" {
			t.Fatalf("the view reads %q after writing z", got)
		}
		if got := names.ConvenientInterface().([]string); len(got) != 2 || got[1] != "b" {
			t.Fatalf("the write went through to the variant's own array: %v", got)
		}

		// The bulk write, through a handle that was born empty and is
		// freed by hand: what is freed must be the buffer the handle
		// holds now, not the one it was made with.
		points := gd.NewPackedVector3Slice([]gd.Vector3{{X: 1}, {Y: 2}})
		shared := gd.NewVariant(points)
		points.CopyFromSlice([]gd.Vector3{{X: 7}, {Y: 8}})
		if got := points.Index(1); got != (gd.Vector3{Y: 8}) {
			t.Fatalf("the array reads %v after the bulk write", got)
		}
		if got := shared.ConvenientInterface().([]gd.Vector3); len(got) != 2 || got[1] != (gd.Vector3{Y: 2}) {
			t.Fatalf("the bulk write went through to the variant's own array: %v", got)
		}
		points.Free()
		if got := shared.ConvenientInterface().([]gd.Vector3); len(got) != 2 || got[0] != (gd.Vector3{X: 1}) {
			t.Fatalf("freeing the array took the variant's with it: %v", got)
		}
		for range 4 {
			pointers.Cycle()
		}
	})
}
