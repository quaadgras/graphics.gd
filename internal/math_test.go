//go:build !generate

package gd_test

import (
	"testing"

	gd "grow.graphics/gd/internal"
)

func TestABS(t *testing.T) {
	if gd.Abs(gd.Int(-1)) != 1 {
		t.Error("expected 1")
	}
	if gd.Abs(gd.Float(-1.2)) != 1.2 {
		t.Error("expected 1.2")
	}
}
