//go:build !(darwin || windows || android) || nativeui_all

package nativeui

import (
	"graphics.gd/nativeui/internal/cosmic"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

func init() { register(lookCOSMIC, cosmicLook) }

// cosmicDefaults are COSMIC's stock dark theme colours, for builds that
// show the look without the theme files (screenshots, other platforms).
var cosmicDefaults = map[string][4]float64{
	"background.base": {0.106, 0.106, 0.106, 1}, "background.on": {0.91, 0.91, 0.91, 1},
	"background.divider": {0.27, 0.27, 0.27, 1}, "background.component.base": {0.18, 0.18, 0.18, 1},
	"background.component.hover": {0.26, 0.26, 0.26, 1}, "background.component.divider": {0.75, 0.75, 0.75, 0.2},
	"primary.base": {0.153, 0.153, 0.153, 1}, "primary.on": {0.91, 0.91, 0.91, 1},
	"primary.divider":        {0.29, 0.29, 0.29, 1},
	"primary.component.base": {0.212, 0.212, 0.212, 1}, "primary.component.hover": {0.291, 0.291, 0.291, 1},
	"primary.component.pressed": {0.37, 0.37, 0.37, 1}, "primary.component.selected": {0.291, 0.291, 0.291, 1},
	"primary.component.border": {0.74, 0.74, 0.74, 1}, "primary.component.divider": {0.75, 0.75, 0.75, 0.2},
	"primary.component.on": {0.75, 0.75, 0.75, 1},
	"accent.base":          {0.39, 0.82, 0.87, 1}, "accent.on": {0, 0, 0, 1}, "accent.focus": {0.39, 0.82, 0.87, 1},
	"accent.selected_text": {0.39, 0.82, 0.87, 1},
	"button.base":          {0.62, 0.62, 0.62, 0.25}, "button.hover": {0.39, 0.39, 0.39, 0.4},
	"button.pressed": {0.17, 0.17, 0.17, 0.625}, "button.on": {0.91, 0.91, 0.91, 1},
	"button.disabled": {0.62, 0.62, 0.62, 0.125}, "button.on_disabled": {0.91, 0.91, 0.91, 0.5},
	"accent_button.base": {0.39, 0.82, 0.87, 1}, "accent_button.hover": {0.47, 0.86, 0.9, 1},
	"accent_button.pressed": {0.3, 0.62, 0.66, 1}, "accent_button.on": {0, 0, 0, 1},
}

// cosmicLook is libcosmic's widget style, coloured from the user's theme.
func cosmicLook(o options) spec {
	th, _ := o.cosmicTheme.(*cosmic.Theme)
	col := func(file string, path ...string) Color.RGBA {
		c, ok := cosmic.RGBA{}, false
		if th != nil {
			c, ok = th.Color(file, path...)
		}
		if !ok {
			key := file
			for _, p := range path {
				key += "." + p
			}
			d := cosmicDefaults[key]
			c = cosmic.RGBA{R: d[0], G: d[1], B: d[2], A: d[3]}
		}
		return Color.RGBA{R: Float.X(c.R), G: Float.X(c.G), B: Float.X(c.B), A: Float.X(c.A)}
	}
	radius := func(name string, def float64) int {
		if th != nil {
			return int(th.Radius(name, def))
		}
		return int(def)
	}
	over := func(top, under Color.RGBA) Color.RGBA { // flatten a translucent colour
		a := top.A
		return Color.RGBA{R: under.R*(1-a) + top.R*a, G: under.G*(1-a) + top.G*a, B: under.B*(1-a) + top.B*a, A: 1}
	}
	window, page := col("background", "base"), col("primary", "base")
	accent := col("accent", "base")
	rs, rm, rxs := radius("radius_s", 8), radius("radius_m", 16), radius("radius_xs", 4)
	s := spec{
		fonts:      []string{"Open Sans", "Fira Sans", "sans-serif"},
		fontSize:   14,
		boldWeight: 700,
		accent:     accent,
		onAccent:   col("accent", "on"),
		window:     window,
		text:       col("background", "on"),
		dim:        blend(col("background", "on"), window, 0.35),
		disabled:   col("button", "on_disabled"),
		separator:  col("primary", "divider"),
		buttonText: col("button", "on"),
		accentText: col("accent_button", "on"),
	}
	// Buttons: the button colour over the window, with half the small radius
	// (measured from COSMIC Settings, "slightly round" style).
	btnPad := [4]float64{14, 6, 14, 6}
	btnR := max(0, rs/2)
	btn := func(c Color.RGBA) box {
		return box{bg: over(c, window), radius: btnR}.padded(btnPad[0], btnPad[1], btnPad[2], btnPad[3])
	}
	s.button = btn(col("button", "base"))
	s.buttonHover = btn(col("button", "hover"))
	s.buttonPressed = btn(col("button", "pressed"))
	s.buttonDisabled = btn(col("button", "disabled"))
	s.accentButton = btn(col("accent_button", "base"))
	s.accentHover = btn(col("accent_button", "hover"))
	s.accentPressed = btn(col("accent_button", "pressed"))
	field := over(col("primary", "component", "base"), page)
	s.entry = box{bg: field, border: over(col("primary", "component", "divider"), field), bw: borderAll(1), radius: rs}.padded(12, 7, 12, 7)
	s.entryFocus = s.entry
	s.entryFocus.border, s.entryFocus.bw = col("accent", "focus"), borderAll(2)
	s.focus = box{border: col("accent", "focus"), bw: borderAll(2), radius: btnR + 2, expand: 2, noCenter: true}
	s.popup = box{bg: over(col("background", "component", "base"), window), border: over(col("background", "component", "divider"), window), bw: borderAll(1), radius: rm, shadow: Color.RGBA{A: 0.3}, shadowSize: 8, shadowY: 3}.padded(8, 8, 8, 8)
	s.popupHover = box{bg: over(col("background", "component", "hover"), window), radius: rs}.padded(12, 6, 12, 6)
	s.popupHoverText = s.text
	s.card = box{bg: page, radius: rm}.padded(16, 14, 16, 14)
	s.tabSelected = box{bg: over(col("primary", "component", "selected"), window), radius: rs}.padded(16, 6, 16, 6)
	s.tabUnselected = box{}.padded(16, 6, 16, 6)
	s.tabHovered = box{bg: over(col("primary", "component", "hover"), window), radius: rs}.padded(16, 6, 16, 6)
	s.tabSelectedText = col("accent", "selected_text")
	rail := over(col("primary", "component", "divider"), page)
	s.sliderTrack = box{bg: rail, radius: 2}.padded(0, 2, 0, 2)
	s.sliderFill = box{bg: accent, radius: 2}
	s.progressTrack = box{bg: rail, radius: 2}.padded(0, 2, 0, 2)
	s.progressFill = box{bg: accent, radius: 2}.padded(0, 2, 0, 2)
	// Checks and radios: a light ring around a dark centre; checked ones
	// are accent-filled with a dark mark. Toggles: a light grey track and
	// a dark knob, the track turning accent when on.
	knob := over(col("background", "component", "base"), window)
	s.check = checkSpec{size: 18, radius: float64(rxs), borderW: 1.5, markW: 2, bg: knob, border: col("primary", "component", "border"), radioDot: 0.15,
		mark: col("accent", "on")}
	s.toggle = switchSpec{w: 48, h: 24, off: blend(col("background", "on"), window, 0.42), knobOff: knob, knobOn: knob, knobOffR: 9, knobOnR: 9}
	s.thumb = thumbSpec{size: 20, fill: accent, border: accent}
	return s
}
