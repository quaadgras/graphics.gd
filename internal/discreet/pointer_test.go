package discreet_test

import (
	"testing"

	"grow.graphics/gd/internal/discreet"
	"grow.graphics/gd/internal/mmm"
)

type MyPointer discreet.PointerNamed[MyPointer, [1]uintptr, [1]discreet.Type]

func (p MyPointer) Free() {
	if ptr, ok := discreet.End(p); ok {
		_ = ptr // free the pointer
	}
}

func TestPointer(t *testing.T) {
	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			discreet.Cycle()
			discreet.MarkAndSweep()
		}
		var p = discreet.New[MyPointer]([1]uintptr{1})
		if discreet.Get(p) != [1]uintptr{1} {
			t.Fatal("bad")
		}
	}
}

func BenchmarkDiscrete(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%512 == 0 {
			discreet.Cycle()
			discreet.MarkAndSweep()
		}
		var p = discreet.New[MyPointer]([1]uintptr{1})
		if discreet.Get(p) != [1]uintptr{1} {
			b.Fatal("bad")
		}

	}

}

type MmmPointer mmm.Pointer[struct{}, MmmPointer, [1]uintptr]

func (p MmmPointer) Free() {
	mmm.End(p)
}

func BenchmarkMMM(b *testing.B) {
	lt := mmm.NewLifetime()
	api := new(struct{})
	for i := 0; i < b.N; i++ {
		if i%512 == 0 {
			lt.End()
			lt = mmm.NewLifetime()
		}
		var p = mmm.New[MmmPointer](lt, api, [1]uintptr{1})
		if mmm.Get(p) != [1]uintptr{1} {
			b.Fatal("bad")
		}
	}
}
