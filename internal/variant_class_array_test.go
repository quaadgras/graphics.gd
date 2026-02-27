package gd_test

import (
	"testing"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/variant"
	"graphics.gd/variant/Array"

	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/Texture2DArray"
)

// TestClassInstanceVariantType verifies that a class instance (e.g. Image.Instance)
// is marshalled as a Variant of TypeObject, not TypeArray, even though
// Image.Instance is defined as [1]gdclass.Image (an array type in Go).
func TestClassInstanceVariantType(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		img := Image.Create(1, 1, false, Image.FormatRgba8)
		v := variant.New(img)
		if v.Type() != variant.TypeObject {
			t.Errorf("expected variant type Object (%d), got %s (%d)", variant.TypeObject, v.Type(), v.Type())
		}
	})
}

// TestArrayFromSliceOfClassInstances verifies that creating a Godot Array
// from a Go slice of class instances ([]Image.Instance) produces an array
// of Object variants, not nested Array variants.
func TestArrayFromSliceOfClassInstances(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		img := Image.Create(4, 4, false, Image.FormatRgba8)
		images := []Image.Instance{img}

		arr := gd.ArrayFromSlice[Array.Contains[[1]gdclass.Image]](images)
		if arr.Len() != 1 {
			t.Fatalf("expected array length 1, got %d", arr.Len())
		}

		// The element should be retrievable as an Object, not as an Array.
		elem := arr.Any().Index(0)
		if elem.Type() != variant.TypeObject {
			t.Errorf("expected element variant type Object (%d), got %s (%d)", variant.TypeObject, elem.Type(), elem.Type())
		}
	})
}

// TestTexture2DArrayCreateFromImages verifies that CreateFromImages
// succeeds when given a slice of valid Image instances.
func TestTexture2DArrayCreateFromImages(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		img := Image.Create(4, 4, false, Image.FormatRgba8)
		texArr := Texture2DArray.New()
		err := texArr.AsImageTextureLayered().CreateFromImages([]Image.Instance{img})
		if err != nil {
			t.Errorf("CreateFromImages failed: %v", err)
		}
	})
}
