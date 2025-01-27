//go:build !generate

package gd_test

import (
	"testing"

	gd "graphics.gd/internal"
)

func BenchmarkMethodBindPointerCall(B *testing.B) {
	B.ReportAllocs()
	s := gd.NewString("Hello, World!")
	for i := 0; i < B.N; i++ {
		_ = s.Length()
	}
}
