package builder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const samplePlist = `<dict>
	<key>UISupportedInterfaceOrientations</key>
	<array>
		<string>UIInterfaceOrientationLandscapeLeft</string>
	</array>
	<key>UISupportedInterfaceOrientations~ipad</key>
	<array>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
	<key>CADisableMinimumFrameDurationOnPhone</key><true/>
</dict>`

func TestEnforceOrientationPortrait(t *testing.T) {
	out := enforceOrientation(samplePlist, 1) // SCREEN_PORTRAIT
	iphone := section(out, "UISupportedInterfaceOrientations</key>")
	ipad := section(out, "UISupportedInterfaceOrientations~ipad</key>")
	if !strings.Contains(iphone, "Portrait<") || strings.Contains(iphone, "Landscape") {
		t.Fatalf("iphone not portrait:\n%s", iphone)
	}
	if !strings.Contains(ipad, "Portrait<") || strings.Contains(ipad, "Landscape") {
		t.Fatalf("ipad not portrait:\n%s", ipad)
	}
	if !strings.Contains(out, "CADisableMinimumFrameDurationOnPhone") {
		t.Fatal("rest of plist mangled")
	}
}

func TestEnforceOrientationLandscapeDefault(t *testing.T) {
	out := enforceOrientation(samplePlist, 0) // SCREEN_LANDSCAPE
	if !strings.Contains(section(out, "UISupportedInterfaceOrientations</key>"), "LandscapeLeft") {
		t.Fatal("iphone landscape-left expected")
	}
	if !strings.Contains(section(out, "UISupportedInterfaceOrientations~ipad</key>"), "LandscapeRight") {
		t.Fatal("ipad landscape-right expected")
	}
}

func TestEnforceOrientationSensor(t *testing.T) {
	out := enforceOrientation(samplePlist, 6) // SCREEN_SENSOR
	iphone := section(out, "UISupportedInterfaceOrientations</key>")
	for _, want := range []string{"LandscapeLeft", "LandscapeRight", "Portrait<", "PortraitUpsideDown"} {
		if !strings.Contains(iphone, want) {
			t.Fatalf("sensor missing %s:\n%s", want, iphone)
		}
	}
}

func TestReadOrientation(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "project.godot"),
		[]byte("[display]\n\nwindow/handheld/orientation=1\n"), 0o644)
	if got := readOrientation(dir); got != 1 {
		t.Fatalf("readOrientation = %d, want 1", got)
	}
	os.WriteFile(filepath.Join(dir, "project.godot"), []byte("[application]\n"), 0o644)
	if got := readOrientation(dir); got != 0 {
		t.Fatalf("readOrientation default = %d, want 0", got)
	}
}

// section returns the text from a key to the next </array>.
func section(plist, key string) string {
	i := strings.Index(plist, key)
	if i < 0 {
		return ""
	}
	rest := plist[i:]
	if j := strings.Index(rest, "</array>"); j >= 0 {
		return rest[:j]
	}
	return rest
}
