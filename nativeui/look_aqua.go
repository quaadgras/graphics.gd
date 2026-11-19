//go:build (darwin && !ios) || nativeui_all

package nativeui

import "graphics.gd/variant/Color"

func init() {
	register(lookMacOS, func(o options) spec {
		if o.Version.Major >= 26 {
			return macOS26(o.Dark, !o.ReduceTransparency)
		}
		return macOS(o.Dark)
	})
}

// macOS26 is macOS 26 Tahoe and later. Sizes and colours are measured from
// AppKit's own controls (regular size, 1x) drawn on macOS 27: 24pt push
// buttons, pop-ups and fields; flat grey fills without borders on buttons,
// checkboxes and radios; capsule knobs on switches and sliders; a tab view
// whose centered segmented tabs straddle the top edge of its box.
func macOS26(dark, glass bool) spec {
	s := spec{
		fonts:      []string{"SF Pro Text", "SF Pro", ".AppleSystemUIFont", "Helvetica Neue", "Inter", "Noto Sans"},
		fontSize:   13,
		boldWeight: 700,
		accent:     hex("#007AFF"),
		onAccent:   white(1),
	}
	var button, buttonDisabled, field, fieldBorder, fill, track, card, cardBorder, tabBar, tabSelected, knobOff, thumb Color.RGBA
	if dark {
		s.accent = hex("#0A84FF")
		s.window = hex("#1E1E1E")
		s.text, s.dim, s.disabled, s.separator = white(0.9), white(0.55), white(0.25), hex("#3A3A3A")
		button, buttonDisabled = hex("#363636"), hex("#2D2D2D")
		field, fieldBorder = hex("#171717"), hex("#393939")
		fill, track = hex("#3A3A3A"), hex("#363636")
		card, cardBorder = hex("#252525"), hex("#404040")
		tabBar, tabSelected = hex("#363636"), hex("#5C5C5C")
		knobOff, thumb = hex("#E2E2E2"), hex("#DFDFDF")
		s.popup = box{bg: hex("#2A2A2A"), border: white(0.12), bw: borderAll(1), radius: 12, shadow: black(0.45), shadowSize: 16, shadowY: 6}
	} else {
		s.window = white(1)
		s.text, s.dim, s.disabled, s.separator = black(0.88), black(0.5), black(0.3), hex("#DFDFDF")
		button, buttonDisabled = hex("#E4E4E4"), hex("#EDEDED")
		field, fieldBorder = white(1), hex("#D2D2D2")
		fill, track = hex("#DFDFDF"), hex("#E8E8E8")
		card, cardBorder = hex("#F7F7F7"), hex("#ECECEC")
		tabBar, tabSelected = hex("#E6E6E6"), white(1)
		knobOff, thumb = white(1), white(1)
		s.popup = box{bg: rgba(250, 250, 250, 0.98), border: black(0.08), bw: borderAll(1), radius: 12, shadow: black(0.2), shadowSize: 16, shadowY: 6}
	}
	const r = 6
	pad := [4]float64{12, 3, 12, 3}
	s.button = box{bg: button, radius: r}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonHover = s.button
	s.buttonPressed = s.button.with(shade(button, 0.92))
	s.buttonDisabled = s.button.with(buttonDisabled)
	s.accentButton = box{bg: s.accent, radius: r}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentHover = s.accentButton
	s.accentPressed = s.accentButton.with(shade(s.accent, 0.85))
	s.entry = box{bg: field, border: fieldBorder, bw: borderAll(1), radius: r}.padded(7, 3, 7, 3)
	s.entryFocus = s.entry
	s.focus = box{border: withAlpha(s.accent, 0.5), bw: borderAll(3), radius: r + 3, expand: 3, noCenter: true}
	s.popup = s.popup.padded(5, 5, 5, 5)
	s.popupHover = box{bg: s.accent, radius: 6}.padded(9, 3, 9, 3)
	s.popupHoverText = white(1)
	s.card = box{bg: card, border: cardBorder, bw: borderAll(1), radius: 5}.padded(16, 14, 16, 14)
	// Segmented tabs: the bar's grey behind every tab, the selected one a
	// raised pill; the box reaches up to the tabs' middle.
	s.tabSelected = box{bg: tabSelected, radius: r, shadow: black(0.12), shadowSize: 1, shadowY: 0.5}.padded(16, 3, 16, 3)
	s.tabUnselected = box{bg: tabBar, radius: r}.padded(16, 3, 16, 3)
	s.tabHovered = box{bg: shade(tabBar, 0.96), radius: r}.padded(16, 3, 16, 3)
	s.tabSeparation = 0.01 // touching (zero would mean the default)
	s.card.expandTop = 12
	s.progressTrack = box{bg: track, radius: 4}.padded(0, 4, 0, 4)
	s.progressFill = box{bg: s.accent, radius: 4}.padded(0, 4, 0, 4)
	s.sliderTrack = box{bg: fill, radius: 3}.padded(0, 3, 0, 3)
	s.sliderFill = box{bg: s.accent, radius: 3}
	s.check = checkSpec{size: 16, radius: 4.5, markW: 2, bg: fill, radioDot: 0.16}
	s.toggle = switchSpec{w: 54, h: 24, off: fill, knobOff: knobOff, knobOn: white(1), knobOffR: 10, knobOnR: 10, knobW: 32, knobShadow: !dark}
	s.thumb = thumbSpec{size: 16, w: 20, fill: thumb, border: black(0.1), shadow: true}
	s.arrowUpDown, s.arrowColor = true, s.text
	if glass {
		// Liquid Glass on the tab bar. Menus stay opaque: Godot draws popups in
		// their own viewport, where the shader can't see the window behind.
		tint := withAlpha(white(1), 0.35)
		if dark {
			tint = Color.RGBA{R: 0.16, G: 0.16, B: 0.16, A: 0.45}
		}
		s.tabUnselected = box{bg: tint, radius: r, glass: true, shadow: black(0.08), shadowSize: 2}.padded(16, 3, 16, 3)
		s.tabHovered = s.tabUnselected
	}
	return s
}

func macOS(dark bool) spec {
	s := spec{
		fonts:      []string{"SF Pro Text", "SF Pro", ".AppleSystemUIFont", "Helvetica Neue", "Inter", "Noto Sans"},
		fontSize:   13,
		boldWeight: 600,
		accent:     hex("#007AFF"),
		onAccent:   white(1),
	}
	if dark {
		s.accent = hex("#0A84FF")
		s.window = hex("#2B2B2B")
		s.text, s.dim, s.disabled, s.separator = white(0.85), white(0.55), white(0.25), white(0.1)
		s.button = box{bg: white(0.2), radius: 5, shadow: black(0.3), shadowSize: 1, shadowY: 0.5}
		s.buttonPressed = s.button.with(white(0.32))
		s.buttonDisabled = s.button.with(white(0.08))
		s.entry = box{bg: white(0.05), border: white(0.15), bw: borderAll(1), radius: 3}
		s.popup = box{bg: hex("#323232"), border: white(0.12), bw: borderAll(1), radius: 6, shadow: black(0.4), shadowSize: 14, shadowY: 5}
		s.card = box{bg: white(0.04), border: white(0.07), bw: borderAll(1), radius: 8}
		s.tabSelected = box{bg: hex("#636366"), radius: 5, shadow: black(0.3), shadowSize: 1, shadowY: 0.5}
		s.tabHovered = box{bg: white(0.05), radius: 5}
		s.progressTrack = box{bg: white(0.1), radius: 3}
		s.sliderTrack = box{bg: white(0.12), radius: 2}
		s.check = checkSpec{bg: white(0.08), border: white(0.25)}
		s.toggle = switchSpec{off: white(0.15), knobOff: hex("#DCDCDC"), knobOn: white(1)}
		s.thumb = thumbSpec{fill: hex("#DCDCDC"), border: black(0.2), shadow: true}
	} else {
		s.window = hex("#ECECEC")
		s.text, s.dim, s.disabled, s.separator = black(0.85), black(0.5), black(0.25), black(0.1)
		s.button = box{bg: white(1), border: black(0.06), bw: borderAll(1), radius: 5, shadow: black(0.18), shadowSize: 1, shadowY: 0.5}
		s.buttonPressed = s.button.with(hex("#E0E0E0"))
		s.buttonDisabled = s.button.with(white(0.5))
		s.entry = box{bg: white(1), border: black(0.15), bw: borderAll(1), radius: 3}
		s.popup = box{bg: rgba(246, 246, 246, 0.98), border: black(0.1), bw: borderAll(1), radius: 6, shadow: black(0.22), shadowSize: 14, shadowY: 5}
		s.card = box{bg: black(0.035), border: black(0.06), bw: borderAll(1), radius: 8}
		s.tabSelected = box{bg: white(1), radius: 5, shadow: black(0.2), shadowSize: 1, shadowY: 0.5}
		s.tabHovered = box{bg: black(0.04), radius: 5}
		s.progressTrack = box{bg: black(0.08), radius: 3}
		s.sliderTrack = box{bg: black(0.1), radius: 2}
		s.check = checkSpec{bg: white(1), border: black(0.25)}
		s.toggle = switchSpec{off: black(0.09), knobOff: white(1), knobOn: white(1)}
		s.thumb = thumbSpec{fill: white(1), border: black(0.12), shadow: true}
	}
	s.button = s.button.padded(12, 3, 12, 4)
	s.buttonHover = s.button
	s.buttonPressed.pad, s.buttonDisabled.pad = s.button.pad, s.button.pad
	s.buttonPressed.setPad, s.buttonDisabled.setPad = true, true
	s.accentButton = box{bg: s.accent, radius: 5, shadow: black(0.2), shadowSize: 1, shadowY: 0.5}.padded(12, 3, 12, 4)
	s.accentHover = s.accentButton
	s.accentPressed = s.accentButton.with(shade(s.accent, 0.85))
	s.entry = s.entry.padded(6, 3, 6, 3)
	s.entryFocus = s.entry
	s.focus = box{border: s.accent, bw: borderAll(3), radius: 7, expand: 3, noCenter: true}
	s.focus.border.A = 0.5
	s.popup = s.popup.padded(5, 5, 5, 5)
	s.popupHover = box{bg: s.accent, radius: 4}.padded(8, 2, 8, 2)
	s.popupHoverText = white(1)
	s.card = s.card.padded(16, 14, 16, 14)
	s.tabSelected = s.tabSelected.padded(14, 2, 14, 3)
	s.tabUnselected = box{}.padded(14, 2, 14, 3)
	s.tabHovered = s.tabHovered.padded(14, 2, 14, 3)
	s.progressTrack = s.progressTrack.padded(0, 3, 0, 3)
	s.progressFill = box{bg: s.accent, radius: 3}.padded(0, 3, 0, 3)
	s.sliderTrack = s.sliderTrack.padded(0, 2, 0, 2)
	s.sliderFill = box{bg: s.accent, radius: 2}
	s.check.size, s.check.radius, s.check.borderW, s.check.markW, s.check.radioDot = 14, 3.5, 1, 2, 0.2
	s.toggle.w, s.toggle.h, s.toggle.knobOffR, s.toggle.knobOnR, s.toggle.knobShadow = 38, 22, 10, 10, true
	s.thumb.size = 20
	s.arrowBadge = true
	return s
}
