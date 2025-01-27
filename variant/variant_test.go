package variant_test

import (
	"fmt"
	"testing"

	"graphics.gd/variant"
)

func TestAny(t *testing.T) {
	var i16 = variant.New(int16(42))
	if i16.Interface() != int16(42) {
		t.Error("i16.Interface() != 42")
	}
	var hello = variant.New("hello")
	if hello.Interface() != "hello" {
		fmt.Println(hello.Interface())
		t.Error("hello.Interface() != 'hello'")
	}
	var bytes = variant.New([]byte{0x01, 0x02, 0x03})
	if fmt.Sprintf("%v", bytes.Interface()) != "[1 2 3]" {
		t.Error("bytes.Interface() != [1 2 3]")
	}

	var i8 = variant.New(int8(22))
	if i8.Int8() != int8(22) {
		t.Error("i8.Int8() != 22")
	}
}

func BenchmarkNewAllocs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		variant.New(42)
	}
}
