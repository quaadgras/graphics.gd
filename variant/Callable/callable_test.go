package Callable_test

import (
	"testing"

	"graphics.gd/variant/Callable"
)

func TestCallable(t *testing.T) {
	var fn = Callable.New(func() string {
		return "hello"
	})
	if fn.Call().String() != "hello" {
		t.Fatal("invalid call")
	}
}

func BenchmarkAllocations(b *testing.B) {
	b.ReportAllocs()
	var fn = Callable.New(func() string {
		return "hello"
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn.Call()
	}
}
