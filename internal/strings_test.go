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
	pointers.Set(str, gdextension.Host.Strings.Append.String(pointers.Get(str), pointers.Get(gd.NewString(" from Go!"))))
	if str.String() != "Hello, World! from Go!" {
		t.Fail()
	}
}

func TestVariantStrings(t *testing.T) {
	var str = gd.NewVariant(gd.NewString("Hello, Variant!"))
	if gd.VariantAs[string](str) != "Hello, Variant!" {
		t.Fail()
	}
	var str_name = gd.NewVariant(gd.NewStringName("Hello, StringName!"))
	if gd.VariantAs[string](str_name) != "Hello, StringName!" {
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

// TestUTF8StringRoundTrip guards the Godot->Go conversion: String.String()
// used to size its UTF-8 buffer by character count (String.Length()), which
// truncates multibyte sequences mid-character for non-ASCII text and yields
// invalid UTF-8 (issue: "Unicode parsing error" in the engine).
func TestUTF8StringRoundTrip(t *testing.T) {
	for _, text := range []string{
		// CJK: 3-byte characters.
		"中文测试",
		"你好，世界！",
		"こんにちは",
		"日本語のテスト",
		"한국어 테스트",
		// Latin-1 / Latin Extended-A: 2-byte characters.
		"café",
		"Grüße aus München",
		"naïve — résumé",
		"Español ñ",
		// Supplementary planes: 4-byte characters.
		"🚀",
		// Mixed widths in one string.
		"mixed 中文 + 日本語 + émigré 123",
	} {
		t.Run(text, func(t *testing.T) {
			if got := gd.NewString(text).String(); got != text {
				t.Errorf("String round-trip: got %q, want %q", got, text)
			}
			if got := gd.NewStringName(text).String(); got != text {
				t.Errorf("StringName round-trip: got %q, want %q", got, text)
			}
			if got := gd.VariantAs[string](gd.NewVariant(gd.NewString(text))); got != text {
				t.Errorf("Variant round-trip: got %q, want %q", got, text)
			}
		})
	}
}
