//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/variant/String"
)

func TestStrings(t *testing.T) {
	var str = String.New("Hello, World!")
	if str.String() != "Hello, World!" {
		t.Fail()
	}
	str.Append(String.New(" from Go!"))
	if str.String() != "Hello, World! from Go!" {
		t.Fail()
	}
}
