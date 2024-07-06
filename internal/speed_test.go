//go:build !generate

package gd_test

import (
	"testing"

	gd "grow.graphics/gd/internal"
)

func BenchmarkMethodBindPointerCall(B *testing.B) {
	godot := gd.NewLifetime(API)
	defer godot.End()

	s := godot.String("Hello, World!")

	for i := 0; i < B.N; i++ {
		_ = s.Length()
	}
}
