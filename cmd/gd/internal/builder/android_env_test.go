package builder

import "testing"

func TestBakeable(t *testing.T) {
	for name, want := range map[string]bool{
		"REALM_REVIEW":      true,
		"REALM_SHOT":        true,
		"MY_SETTING":        true,
		"PATH":              false,
		"HOME":              false,
		"LD_PRELOAD":        false,
		"TERMUX_VERSION":    false,
		"ANDROID_DATA":      false,
		"GOARCH":            false,
		"CGO_ENABLED":       false,
		"BOOTCLASSPATH":     false,
		"GITHUB_TOKEN":      false,
		"ANTHROPIC_API_KEY": false,
		"MY_SECRET":         false,
		"_":                 false,
	} {
		if got := bakeable(name); got != want {
			t.Errorf("bakeable(%q) = %v, want %v", name, got, want)
		}
	}
}
