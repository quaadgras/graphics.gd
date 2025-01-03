package pointers_test

import (
	"testing"

	"graphics.gd/internal/mmm"
	"graphics.gd/internal/pointers"
)

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

var simulated_pointers = make(map[[1]uintptr]bool)

type MyPointer pointers.Solo[MyPointer]

func (ptr MyPointer) Free() {
	raw, ok := pointers.End(ptr)
	if ok {
		delete(simulated_pointers, raw)
	}
}

func BenchmarkDiscrete(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%512 == 0 {
			pointers.Cycle()
		}
		var p = pointers.New[MyPointer]([1]uintptr{1})
		if pointers.Get(p) != [1]uintptr{1} {
			b.Fatal("bad")
		}
	}
}

type MyString pointers.Pair[MyString]

func (ptr MyString) Free() {}

type MySlice pointers.Trio[MySlice]

func (ptr MySlice) Free() {}

func TestPointersV2(t *testing.T) {
	pointers.New[MyPointer]([1]uintptr{1})
	simulated_pointers[[1]uintptr{1}] = true

	pointers.Cycle()

	if len(simulated_pointers) != 0 {
		t.Fatal("simulated pointers not freed")
	}
}
