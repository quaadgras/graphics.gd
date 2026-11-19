//go:build android || nativeui_all

package nativeui

import (
	"graphics.gd/variant/Color"
)

func init() {
	register(lookAndroid, func(o options) spec {
		if o.Accent.A > 0 {
			return materialDynamic(o.Accent, o.Dark)
		}
		return material(o.Dark)
	})
}

// materialDynamic recolours the scheme from the system's accent colour, as
// Android 12+ does from the wallpaper ("tonal spot"): the primary palette
// keeps the accent's hue and chroma (at least 36), secondary is the same
// hue at chroma 16, the neutrals at chroma 4 and 8. Each role takes its
// tone from Material's scheme.
func materialDynamic(accent Color.RGBA, dark bool) spec {
	p := tonalOf(accent)
	primary := p.withChroma(max(p.chroma, 36))
	secondary, neutral, neutralVariant := p.withChroma(16), p.withChroma(4), p.withChroma(8)
	t := func(lightTone, darkTone float64) float64 {
		if dark {
			return darkTone
		}
		return lightTone
	}
	accentC, onAccent := primary.tone(t(40, 80)), primary.tone(t(100, 20))
	surface, highest := neutral.tone(t(98, 6)), neutral.tone(t(90, 22))
	text, onVariant := neutral.tone(t(10, 90)), neutralVariant.tone(t(30, 80))
	outline, outlineVariant := neutralVariant.tone(t(50, 60)), neutralVariant.tone(t(80, 30))
	secondaryContainer := secondary.tone(t(90, 30))
	m := material(dark)
	m.accent, m.onAccent, m.window, m.text, m.dim, m.separator = accentC, onAccent, surface, text, onVariant, outlineVariant
	m.disabled = withAlpha(text, 0.38)
	repaint := func(b *box, from, to Color.RGBA) {
		if b.bg == from {
			b.bg = to
		}
		if b.border == from {
			b.border = to
		}
	}
	base := material(dark)
	for _, pair := range [][2]Color.RGBA{{base.accent, accentC}, {base.onAccent, onAccent}, {base.window, surface}, {base.text, text},
		{base.dim, onVariant}, {base.separator, outlineVariant}, {base.check.border, onVariant}, {base.toggle.offBorder, outline},
		{base.toggle.off, highest}, {base.entry.bg, highest}, {base.progressTrack.bg, secondaryContainer}, {base.sliderTrack.bg, secondaryContainer}} {
		for _, b := range []*box{&m.button, &m.buttonHover, &m.buttonPressed, &m.buttonDisabled, &m.accentButton, &m.accentHover, &m.accentPressed,
			&m.entry, &m.entryFocus, &m.popup, &m.popupHover, &m.card, &m.tabSelected, &m.tabUnselected, &m.tabHovered,
			&m.progressTrack, &m.progressFill, &m.sliderTrack, &m.sliderFill} {
			repaint(b, pair[0], pair[1])
		}
	}
	if m.tabPanel != nil {
		repaint(m.tabPanel, base.separator, outlineVariant)
	}
	m.buttonText, m.tabSelectedText = accentC, accentC
	m.buttonHover.bg, m.buttonPressed.bg = withAlpha(accentC, 0.08), withAlpha(accentC, 0.12)
	m.check.border = text
	m.toggle.off, m.toggle.offBorder, m.toggle.knobOff, m.toggle.knobOn = highest, outline, outline, onAccent
	m.thumb.fill, m.thumb.border, m.thumb.gap = accentC, accentC, surface
	return m
}

// material is Android's Material 3 (baseline color scheme), as Material
// Components 1.12 draws it (measured on Android 16).
func material(dark bool) spec {
	s := spec{
		fonts:       []string{"Roboto", "Google Sans Text", "Google Sans", "Noto Sans"},
		fontSize:    14,
		boldWeight:  500,
		noFocusRing: true,
	}
	var surface, container, containerHigh, containerHighest, onVariant, outline, outlineVariant, secondaryContainer Color.RGBA
	if dark {
		s.accent, s.onAccent = hex("#D0BCFF"), hex("#381E72")
		surface, container, containerHigh, containerHighest = hex("#141218"), hex("#211F26"), hex("#2B2930"), hex("#36343B")
		s.text, onVariant, outline, outlineVariant = hex("#E6E0E9"), hex("#CAC4D0"), hex("#938F99"), hex("#49454F")
		secondaryContainer = hex("#4A4458")
		s.card = box{bg: hex("#1D1B20"), radius: 12}
	} else {
		s.accent, s.onAccent = hex("#6750A4"), white(1)
		surface, container, containerHigh, containerHighest = hex("#FEF7FF"), hex("#F3EDF7"), hex("#ECE6F0"), hex("#E6E0E9")
		s.text, onVariant, outline, outlineVariant = hex("#1D1B20"), hex("#49454F"), hex("#79747E"), hex("#CAC4D0")
		secondaryContainer = hex("#E8DEF8")
		s.card = box{bg: hex("#F7F2FA"), radius: 12}
	}
	_ = containerHigh
	s.window = surface
	s.dim, s.separator = onVariant, outlineVariant
	s.disabled = withAlpha(s.text, 0.38)

	// Outlined buttons for the default style, filled for the accent one.
	pad := [4]float64{24, 10, 24, 10}
	s.button = box{border: outline, bw: borderAll(1), radius: 20}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonHover = s.button.with(withAlpha(s.accent, 0.08))
	s.buttonPressed = s.button.with(withAlpha(s.accent, 0.12))
	s.buttonDisabled = box{border: withAlpha(s.text, 0.12), bw: borderAll(1), radius: 20}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonText = s.accent
	s.accentButton = box{bg: s.accent, radius: 20}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentHover = s.accentButton.with(shade(s.accent, 0.95))
	s.accentPressed = s.accentButton.with(shade(s.accent, 0.9))

	// Filled text fields.
	s.entry = box{bg: containerHighest, border: onVariant, bw: [4]int{0, 0, 0, 1}, radius: 4, squareBottom: true}.padded(16, 18, 16, 18) // 56dp
	s.entryFocus = s.entry
	s.entryFocus.border, s.entryFocus.bw = s.accent, [4]int{0, 0, 0, 2}

	s.popup = box{bg: container, radius: 4, shadow: black(0.25), shadowSize: 8, shadowY: 2}.padded(0, 8, 0, 8)
	s.popupHover = box{bg: withAlpha(s.text, 0.08)}.padded(12, 8, 12, 8)
	s.popupHoverText = s.text
	s.card = s.card.padded(16, 16, 16, 16)

	// Primary tabs: accent label with an underline.
	s.tabSelected = box{border: s.accent, bw: [4]int{0, 0, 0, 3}, noCenter: true}.padded(16, 12, 16, 11)
	s.tabUnselected = box{}.padded(16, 12, 16, 11)
	s.tabHovered = box{bg: withAlpha(s.text, 0.08)}.padded(16, 12, 16, 11)
	s.tabSelectedText = s.accent

	// Tabs sit on the surface over a hairline, without a card.
	s.tabPanel = &box{border: outlineVariant, bw: [4]int{0, 1, 0, 0}, noCenter: true}
	s.tabPanel.pad, s.tabPanel.setPad = [4]float64{0, 16, 0, 16}, true

	s.progressTrack = box{bg: secondaryContainer, radius: 2}.padded(0, 2, 0, 2)
	s.progressFill = box{bg: s.accent, radius: 2}.padded(0, 2, 0, 2)
	// The 2024 slider: a 16dp track, a 4dp bar for the handle with a gap
	// either side.
	s.sliderTrack = box{bg: secondaryContainer, radius: 8}.padded(0, 8, 0, 8)
	s.sliderFill = box{bg: s.accent, radius: 8}

	s.check = checkSpec{size: 18, radius: 2, borderW: 2, markW: 2, border: s.text, radioDot: 0.25, radioRing: true}
	s.toggle = switchSpec{w: 52, h: 32, off: containerHighest, offBorder: outline, offBorderW: 2,
		knobOff: outline, knobOn: s.onAccent, knobOffR: 8, knobOnR: 12}
	s.thumb = thumbSpec{size: 44, w: 4, gapW: 6, gap: surface, fill: s.accent, border: s.accent, radius: 2}
	return s
}
