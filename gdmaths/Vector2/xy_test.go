package Vector2_test

import (
	"testing"

	"grow.graphics/gd/variant/Vector2"
)

func TestVector2(t *testing.T) {
	if Vector2.Abs(Vector2.XY{-2, -2}) != (Vector2.XY{2, 2}) {
		t.Fatal("Abs failed")
	}
}
