//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/variant/String"
	"graphics.gd/variant/StringName"
)

func TestStrings(t *testing.T) {
	var str = String.New("Hello, World!")
	if str.String() != "Hello, World!" {
		t.Fail()
	}
	str = String.Append(str, String.New(" from Go!"))
	if str.String() != "Hello, World! from Go!" {
		t.Fail()
	}
}

func TestStringNames(t *testing.T) {
	var str = StringName.New("Hello, World!")
	if str.String() != "Hello, World!" {
		t.Fail()
	}
}

var HelloWorld = String.New("Hello, World!")

func TestStaticStrings(t *testing.T) {
	if HelloWorld.String() != "Hello, World!" {
		t.Fail()
	}
}
