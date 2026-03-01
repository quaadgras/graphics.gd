package rodatacheck_test

import (
	"testing"

	"graphics.gd/internal/rodatacheck"
)

func TestLiteral(t *testing.T) {
	if !rodatacheck.String("hello") {
		t.Skip("rodatacheck not supported on this platform")
	}
}

func TestHeap(t *testing.T) {
	s := string([]byte("hello"))
	if rodatacheck.String(s) {
		t.Fatal("expected String() == false for heap-allocated string")
	}
}

func TestEmpty(t *testing.T) {
	if rodatacheck.String("") {
		t.Fatal("expected String() == false for empty string")
	}
}

func TestConcat(t *testing.T) {
	a := "hello"
	b := " world"
	s := a + b
	if rodatacheck.String(s) {
		t.Fatal("expected String() == false for runtime-concatenated string")
	}
}
