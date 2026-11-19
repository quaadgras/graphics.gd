//go:build js || nativeui_all

package nativeui

import "graphics.gd/variant/Color"

func init() { register(lookWeb, func(o options) spec { return web(o.Browser, o.Dark) }) }

// web is the browser's own form controls: what an unstyled <button>,
// <input> or <select> looks like. Each engine draws them its own way, the
// same on every desktop operating system; measured from Chromium 155,
// Firefox 160 and Safari (macOS 27) rendering the same page. Page text is
// the document font (serif, 16px); controls use a smaller sans.
func web(browser browserID, dark bool) spec {
	switch browser {
	case browserFirefox:
		return webFirefox(dark)
	case browserSafari:
		return webSafari(dark)
	}
	return webChromium(dark)
}

var webSerif = []string{"Times New Roman", "Times", "Liberation Serif", "DejaVu Serif", "Noto Serif", "serif"}

// webChromium: Chrome's "refreshed" form controls (2020 onwards).
func webChromium(dark bool) spec {
	s := spec{
		fonts:      []string{"Arial", "Liberation Sans", "Helvetica", "Arimo", "sans-serif"},
		fontSize:   13,
		boldWeight: 700,
		labelFonts: webSerif, labelSize: 16,
		switchAsCheck: true,
	}
	var button, buttonBorder, hover, hoverBorder, pressed, pressedBorder, field, fieldBorder, ring Color.RGBA
	if dark {
		s.window, s.text, s.dim = hex("#121212"), white(1), hex("#757575")
		s.disabled, s.separator = white(0.3), hex("#858585")
		s.accent, s.onAccent = hex("#99C8FF"), hex("#3B3B3B")
		button, buttonBorder = hex("#6B6B6B"), hex("#6B6B6B")
		hover, hoverBorder = hex("#7B7B7B"), hex("#7B7B7B")
		pressed, pressedBorder = hex("#5B5B5B"), hex("#5B5B5B")
		field, fieldBorder, ring = hex("#3B3B3B"), hex("#858585"), hex("#EEEEEE")
	} else {
		s.window, s.text, s.dim = white(1), black(1), hex("#757575")
		s.disabled, s.separator = rgba(16, 16, 16, 0.3), hex("#767676")
		s.accent, s.onAccent = hex("#0075FF"), white(1)
		button, buttonBorder = hex("#EFEFEF"), hex("#767676")
		hover, hoverBorder = hex("#E5E5E5"), hex("#4F4F4F")
		pressed, pressedBorder = hex("#F5F5F5"), hex("#8D8D8D")
		field, fieldBorder, ring = white(1), hex("#767676"), hex("#101010")
	}
	pad := [4]float64{8, 3, 8, 3} // 2px border + 1px 6px padding: 21px tall
	btn := func(bg, border Color.RGBA) box {
		return box{bg: bg, border: border, bw: borderAll(1), radius: 2}.padded(pad[0], pad[1], pad[2], pad[3])
	}
	s.button, s.buttonHover, s.buttonPressed = btn(button, buttonBorder), btn(hover, hoverBorder), btn(pressed, pressedBorder)
	s.buttonDisabled = btn(withAlpha(button, 0.3), withAlpha(buttonBorder, 0.3))
	// The web has no default button: Save looks like Cancel.
	s.accentButton, s.accentHover, s.accentPressed, s.accentText = s.button, s.buttonHover, s.buttonPressed, s.text
	s.entry = box{bg: field, border: fieldBorder, bw: borderAll(1), radius: 2}.padded(4, 3, 4, 3)
	s.entryFocus = box{bg: field, border: ring, bw: borderAll(2), radius: 2}.padded(4, 3, 4, 3)
	s.focus = box{border: ring, bw: borderAll(2), radius: 2, noCenter: true}
	s.popup = box{bg: field, border: fieldBorder, bw: borderAll(1)}.padded(1, 1, 1, 1)
	s.popupHover = box{bg: hex("#1967D2")}.padded(4, 1, 4, 1)
	s.popupHoverText = white(1)
	s.card = box{border: hex("#C0C0C0"), bw: borderAll(2), noCenter: true}.padded(12, 10, 12, 12) // <fieldset>
	s.tabSelected = btn(field, fieldBorder)
	s.tabUnselected = s.button
	s.tabHovered = s.buttonHover
	s.tabSeparation = 4
	track := box{bg: field, border: hex("#B2B2B2"), bw: borderAll(1), radius: 4}
	if dark {
		track = box{bg: field, border: fieldBorder, bw: borderAll(1), radius: 4}
	}
	s.progressTrack, s.progressFill = track.padded(0, 4, 0, 4), box{bg: s.accent, radius: 4}.padded(0, 4, 0, 4)
	s.sliderTrack, s.sliderFill = track.padded(0, 4, 0, 4), box{bg: s.accent, radius: 4}
	s.check = checkSpec{size: 13, radius: 2, borderW: 1, markW: 2, bg: field, border: fieldBorder,
		radioRing: true, radioGap: field, radioDot: 0.3, mark: s.onAccent}
	s.thumb = thumbSpec{size: 16, fill: s.accent, border: s.accent}
	return s
}

// webFirefox: Firefox's own ("non-native") form controls, used on Windows
// and Linux since Firefox 89.
func webFirefox(dark bool) spec {
	s := spec{
		fonts:      []string{"Segoe UI", "sans-serif"}, // system-ui: Segoe UI on Windows, fontconfig's default sans on Linux
		fontSize:   13,
		boldWeight: 700,
		labelFonts: []string{"Times New Roman", "serif"}, labelSize: 16, // Firefox asks fontconfig for "serif"
		switchAsCheck: true,
		accent:        hex("#0060DF"),
		onAccent:      white(1),
	}
	var button, hover, pressed, field Color.RGBA
	border := hex("#8F8F9D")
	if dark {
		s.window, s.text, s.dim = hex("#1C1B22"), hex("#FBFBFE"), hex("#9B9BA0")
		s.disabled, s.separator = withAlpha(hex("#FBFBFE"), 0.4), border
		button, hover, pressed, field = hex("#2B2A33"), hex("#52525E"), hex("#5B5B66"), hex("#2B2A33")
	} else {
		s.window, s.text, s.dim = white(1), black(1), hex("#757575")
		s.disabled, s.separator = black(0.4), border
		button, hover, pressed, field = hex("#E9E9ED"), hex("#D0D0D7"), hex("#B1B1B9"), white(1)
	}
	pad := [4]float64{9, 4, 9, 4}
	btn := func(bg Color.RGBA) box {
		return box{bg: bg, border: border, bw: borderAll(1), radius: 4}.padded(pad[0], pad[1], pad[2], pad[3])
	}
	s.button, s.buttonHover, s.buttonPressed = btn(button), btn(hover), btn(pressed)
	s.buttonDisabled = box{bg: withAlpha(button, 0.4), border: withAlpha(border, 0.4), bw: borderAll(1), radius: 4}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentButton, s.accentHover, s.accentPressed, s.accentText = s.button, s.buttonHover, s.buttonPressed, s.text
	s.entry = box{bg: field, border: border, bw: borderAll(1), radius: 4}.padded(5, 4, 5, 4)
	s.entryFocus = box{bg: field, border: s.accent, bw: borderAll(2), radius: 4}.padded(5, 4, 5, 4)
	s.focus = box{border: s.accent, bw: borderAll(2), radius: 5, expand: 1, noCenter: true}
	s.popup = box{bg: field, border: border, bw: borderAll(1), radius: 4}.padded(2, 2, 2, 2)
	s.popupHover = box{bg: s.accent, radius: 2}.padded(6, 2, 6, 2)
	s.popupHoverText = white(1)
	s.card = box{border: hex("#C0C0C0"), bw: borderAll(2), noCenter: true}.padded(12, 10, 12, 12)
	s.tabSelected, s.tabUnselected, s.tabHovered, s.tabSeparation = btn(field), s.button, s.buttonHover, 4
	track := box{bg: hex("#E9E9ED"), border: border, bw: borderAll(1), radius: 3}
	s.progressTrack, s.progressFill = track.padded(0, 3, 0, 3), box{bg: s.accent, radius: 3}.padded(0, 3, 0, 3)
	s.sliderTrack, s.sliderFill = track.padded(0, 3, 0, 3), box{bg: s.accent, radius: 3}
	s.check = checkSpec{size: 14, radius: 2.5, borderW: 1, markW: 2, bg: field, border: border,
		radioRing: true, radioGap: white(1), radioDot: 0.22}
	s.thumb = thumbSpec{size: 16, fill: hex("#676774"), border: white(1)}
	return s
}

// webSafari: WebKit on macOS draws AppKit's small controls with the
// system font at 11px.
func webSafari(dark bool) spec {
	s := spec{
		fonts:      []string{"SF Pro Text", "SF Pro", ".AppleSystemUIFont", "Helvetica Neue", "Inter", "Noto Sans"},
		fontSize:   11,
		boldWeight: 700,
		labelFonts: []string{"Times", "Times New Roman", "Liberation Serif", "serif"}, labelSize: 16,
		accent:      hex("#007AFF"),
		onAccent:    white(1),
		arrowUpDown: true,
	}
	var button, disabled, field, border, track, knob Color.RGBA
	if dark {
		s.window, s.text, s.dim, s.disabled, s.separator = hex("#1E1E1E"), white(0.9), white(0.5), white(0.25), white(0.6)
		button, disabled, field, border, track, knob = hex("#2F2F2F"), hex("#252525"), hex("#1E1E1E"), hex("#393939"), hex("#343434"), hex("#D9EBFF")
	} else {
		s.window, s.text, s.dim, s.disabled, s.separator = white(1), black(0.88), black(0.5), black(0.3), black(0.6)
		button, disabled, field, border, track, knob = hex("#EBEBEB"), hex("#F7F7F7"), white(1), hex("#E0E0E0"), hex("#E6E6E6"), white(1)
	}
	pad := [4]float64{9, 2, 9, 2}
	s.button = box{bg: button, radius: 5}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonHover, s.buttonPressed = s.button, s.button.with(shade(button, 0.9))
	s.buttonDisabled = s.button.with(disabled)
	s.accentButton, s.accentHover, s.accentPressed, s.accentText = s.button, s.buttonHover, s.buttonPressed, s.text
	s.arrowColor = s.text
	s.entry = box{bg: field, border: border, bw: borderAll(1), radius: 5}.padded(6, 3, 6, 3)
	s.entryFocus = s.entry
	s.focus = box{border: withAlpha(s.accent, 0.5), bw: borderAll(3), radius: 7, expand: 2, noCenter: true}
	s.popup = box{bg: field, border: black(0.1), bw: borderAll(1), radius: 8, shadow: black(0.2), shadowSize: 12, shadowY: 4}.padded(5, 5, 5, 5)
	s.popupHover = box{bg: s.accent, radius: 4}.padded(8, 1, 8, 1)
	s.popupHoverText = white(1)
	s.card = box{border: hex("#C0C0C0"), bw: borderAll(2), noCenter: true}.padded(12, 10, 12, 12)
	s.tabSelected = s.button.with(shade(button, 0.9))
	s.tabUnselected, s.tabHovered, s.tabSeparation = s.button, s.button, 4
	s.progressTrack, s.progressFill = box{bg: track, radius: 4}.padded(0, 4, 0, 4), box{bg: s.accent, radius: 4}.padded(0, 4, 0, 4)
	s.sliderTrack, s.sliderFill = box{bg: track, radius: 4}.padded(0, 4, 0, 4), box{bg: s.accent, radius: 4}
	s.check = checkSpec{size: 12, radius: 3, borderW: 1, markW: 1.5, bg: field, border: hex("#808080"), radioDot: 0.17}
	if dark {
		s.check.border = hex("#9A9A9A")
	}
	s.toggle = switchSpec{w: 41, h: 18, off: track, knobOff: white(1), knobOn: knob, knobOffR: 7, knobOnR: 7, knobW: 24, knobShadow: !dark}
	s.thumb = thumbSpec{size: 16, w: 24, fill: white(1), border: black(0.08), shadow: true}
	return s
}
