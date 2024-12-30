package Packed_test

import (
	"fmt"
	"testing"

	"graphics.gd/variant/Color"
	"graphics.gd/variant/Packed"
)

func BenchmarkPacked(b *testing.B) {
	var packed Packed.ColorArray
	packed.Append(Color.RGBA{R: 0, G: 0, B: 0, A: 0})
	var sum Color.RGBA
	for i := 0; i < b.N; i++ {
		sum = Color.Add(sum, packed.Index(0))
	}
	fmt.Println(sum)
}
