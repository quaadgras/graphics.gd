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
	runOnMain(t, func(t testing.TB) {
		var str = gd.NewString("Hello, World!")
		if str.String() != "Hello, World!" {
			t.Fail()
		}
		pointers.Set(str, gdextension.Host.Strings.Append.String(pointers.Get(str), pointers.Get(gd.NewString(" from Go!"))))
		if str.String() != "Hello, World! from Go!" {
			t.Fail()
		}
	})
}

func TestVariantStrings(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var str = gd.NewVariant(gd.NewString("Hello, Variant!"))
		if gd.VariantAs[string](str) != "Hello, Variant!" {
			t.Fail()
		}
		var str_name = gd.NewVariant(gd.NewStringName("Hello, StringName!"))
		if gd.VariantAs[string](str_name) != "Hello, StringName!" {
			t.Fail()
		}
	})
}

func TestStringNames(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var str = gd.NewStringName("Hello, World!")
		if str.String() != "Hello, World!" {
			t.Fail()
		}
	})
}

var HelloWorld = String.New("Hello, World!")

func TestStaticStrings(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		if HelloWorld.String() != "Hello, World!" {
			t.Fail()
		}
	})
}
