package nativeui

import "testing"

// TestEveryLook builds the spec of every look compiled in (all of them with
// -tags nativeui_all), light and dark, with and without an accent, and
// checks it can make a usable theme.
func TestEveryLook(t *testing.T) {
	if len(looks) == 0 {
		t.Skip("no looks compiled in")
	}
	for look, fn := range looks {
		if fn == nil {
			continue
		}
		for _, opts := range []options{
			{Look: look},
			{Look: look, Dark: true},
			{Look: look, Accent: hex("#387AFF")},
			{Look: look, Dark: true, HighContrast: true, Palette: &palette{Window: black(1), WindowText: white(1), Button: black(1),
				ButtonText: white(1), Base: black(1), Text: white(1), Highlight: hex("#1AEBFF"), HighlightText: black(1)}},
			{Look: look, Version: osVersion{Major: 26}},
			{Look: look, Version: osVersion{Major: 10, Patch: 19045}},
		} {
			s := lookSpec(opts)
			name := look.String()
			if s.fontSize <= 0 || len(s.fonts) == 0 {
				t.Errorf("%s %+v: no font (%d %v)", name, opts, s.fontSize, s.fonts)
			}
			if s.window.A == 0 && look != lookGTK && look != lookGTK2 {
				t.Errorf("%s: transparent window", name)
			}
			if s.text.A == 0 || s.accent.A == 0 {
				t.Errorf("%s: no text or accent colour", name)
			}
			if s.check.size <= 0 || s.toggle.w <= 0 && !s.switchAsCheck || s.thumb.size <= 0 {
				t.Errorf("%s: check %v toggle %v thumb %v", name, s.check.size, s.toggle.w, s.thumb.size)
			}
			if !s.button.setPad || !s.entry.setPad {
				t.Errorf("%s: buttons or fields without padding", name)
			}
		}
	}
}
