package builder

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMoltenVK(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	library := filepath.Join(home, "sdk", "MoltenVK", "MoltenVK.xcframework",
		"macos-arm64", "libMoltenVK.a")
	if err := os.MkdirAll(filepath.Dir(library), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(library, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		options []string
		want    string
	}{
		{nil, ""},
		{[]string{"vulkan=no", "vulkan_sdk_path=~/sdk"}, ""},
		{[]string{"vulkan=yes", "vulkan=no", "vulkan_sdk_path=~/sdk"}, ""},
		{[]string{"vulkan=True", "use_volk=yes", "vulkan_sdk_path=~/sdk"}, ""},
		{[]string{"vulkan=no", "vulkan=True", "vulkan_sdk_path=~/sdk"},
			library},
	} {
		got, err := moltenVK(test.options, "arm64")
		if err != nil || got != test.want {
			t.Errorf("moltenVK(%q) = %q, %v; want %q",
				test.options, got, err, test.want)
		}
	}
}
