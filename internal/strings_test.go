//go:build !generate

package gd_test

import (
	"testing"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/String"
)

func TestStrings(t *testing.T) {
	var str = gd.NewString("Hello, World!")
	if str.String() != "Hello, World!" {
		t.Fail()
	}
	str = pointers.New[gd.String]([1]gd.EnginePointer{gd.EnginePointer(gdextension.Host.Strings.Append.String(gdextension.String(pointers.Get(str)[0]), gdextension.String(pointers.Get(gd.NewString(" from Go!"))[0])))})
	if str.String() != "Hello, World! from Go!" {
		t.Fail()
	}
}

func TestVariantStrings(t *testing.T) {
	var str = gd.NewVariant(gd.NewString("Hello, Variant!"))
	if gd.VariantAs[string](str) != "Hello, Variant!" {
		t.Fail()
	}
}

func TestStringNames(t *testing.T) {
	var str = gd.NewStringName("Hello, World!")
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
