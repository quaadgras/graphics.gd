//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/variant/String"
)

func BenchmarkMethodBindPointerCall(B *testing.B) {
	s := String.New("Hello, World!")
	for i := 0; i < B.N; i++ {
		_ = s.Length()
	}
}
