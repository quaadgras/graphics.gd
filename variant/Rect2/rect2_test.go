package Rect2_test

import (
	"testing"

	"grow.graphics/gd/internal/gdtests"
	"grow.graphics/gd/variant/Rect2"
	"grow.graphics/gd/variant/Vector2"
)

func TestAbs(t *testing.T) {
	var rect = Rect2.New(25, 25, -100, -50)
	var absolute = Rect2.Abs(rect) // absolute is Rect2(-75, -25, 100, 50)
	gdtests.That(t, absolute, Rect2.New(-75, -25, 100, 50))
}

func TestExpand(t *testing.T) {
	var rect = Rect2.New(0, 0, 5, 2)

	rect = Rect2.ExpandTo(Vector2.New(10, 0), rect) // rect is Rect2(0, 0, 10, 2)
	gdtests.That(t, rect, Rect2.New(0, 0, 10, 2))
	rect = Rect2.ExpandTo(Vector2.New(-5, 5), rect) // rect is Rect2(-5, 0, 15, 5)
	gdtests.That(t, rect, Rect2.New(-5, 0, 15, 5))
}

func TestGrow(t *testing.T) {
	var a = Rect2.Expand(Rect2.New(4, 4, 8, 8), 4) // a is Rect2(0, 0, 16, 16)
	var b = Rect2.Expand(Rect2.New(0, 0, 8, 4), 2) // b is Rect2(-2, -2, 12, 8)
	gdtests.That(t, a, Rect2.New(0, 0, 16, 16))
	gdtests.That(t, b, Rect2.New(-2, -2, 12, 8))
}
