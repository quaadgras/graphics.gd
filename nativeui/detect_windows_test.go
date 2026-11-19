//go:build windows

package nativeui

import "testing"

// TestDetectWindows logs what this machine's settings say.
func TestDetectWindows(t *testing.T) {
	var opts Options
	detectWindows(&opts)
	t.Logf("dark %v accent %v textScale %v reduceTransparency %v highContrast %v", opts.Dark, opts.Accent, opts.textScale, opts.ReduceTransparency, opts.HighContrast)
	if opts.Palette != nil {
		t.Logf("palette %+v", *opts.Palette)
	}
}
