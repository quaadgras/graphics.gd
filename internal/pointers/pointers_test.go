package pointers_test

import (
	"testing"

	"graphics.gd/internal/pointers"
)

var simulated_pointers = make(map[[1]uint64]bool)

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
		var p = pointers.New[MyPointer]([1]uint64{1})
		if pointers.Get(p) != [1]uint64{1} {
			b.Fatal("bad")
		}
	}
}

type MyString pointers.Pair[MyString]

func (ptr MyString) Free() {}

type MySlice pointers.Trio[MySlice]

func (ptr MySlice) Free() {}

func TestPointersV2(t *testing.T) {
	pointers.New[MyPointer]([1]uint64{1})
	simulated_pointers[[1]uint64{1}] = true

	pointers.Cycle()
	pointers.Cycle()

	if len(simulated_pointers) != 0 {
		t.Fatal("simulated pointers not freed")
	}
}
