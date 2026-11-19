//go:build !(darwin || windows || android) || nativeui_all

package nativeui

import "graphics.gd/variant/Color"

func init() { register(Fusion, fusion) }

// fusionPalette is Qt's built-in default palette (what LXQt starts from).
var fusionPalette = Palette{Window: hex("#efefef"), WindowText: hex("#000000"), Button: hex("#efefef"), ButtonText: hex("#000000"),
	Base: hex("#ffffff"), Text: hex("#000000"), Highlight: hex("#308cc6"), HighlightText: hex("#ffffff")}

func lighter(c Color.RGBA, f float32) Color.RGBA {
	clamp := func(v float32) float32 { return min(1, v) }
	return Color.RGBA{R: clamp(c.R * f), G: clamp(c.G * f), B: clamp(c.B * f), A: c.A}
}

func darker(c Color.RGBA, f float32) Color.RGBA {
	return Color.RGBA{R: c.R / f, G: c.G / f, B: c.B / f, A: c.A}
}

// fusion is Qt's Fusion style, the default of LXQt and of Qt programs
// outside KDE. Shapes and colour derivations were measured from renders of
// the real style.
func fusion(o Options) spec {
	p := fusionPalette
	if o.Palette != nil {
		p = *o.Palette
	}
	outline := darker(p.Window, 1.4)
	page := lighter(p.Window, 1.05)
	top, bottom := lighter(p.Button, 1.065), darker(p.Button, 1.005)
	s := spec{
		fonts:      []string{"sans-serif"},
		fontSize:   13,
		boldWeight: 700,
		accent:     p.Highlight,
		onAccent:   p.HighlightText,
		window:     p.Window,
		text:       p.WindowText,
		dim:        blend(p.WindowText, p.Window, 0.4),
		disabled:   blend(p.WindowText, p.Window, 0.6),
		separator:  darker(p.Window, 1.5),
		buttonText: p.ButtonText,
		accentText: p.ButtonText,
	}
	btn := box{border: outline, bw: borderAll(1), radius: 2}.padded(12, 4, 12, 4).gradient(top, bottom)
	s.button = btn
	s.buttonHover = btn.gradient(lighter(top, 1.02), lighter(bottom, 1.02))
	s.buttonPressed = btn.gradient(darker(bottom, 1.07), darker(bottom, 1.03))
	s.buttonDisabled = btn
	s.buttonDisabled.border = blend(outline, p.Window, 0.37)
	s.accentButton = btn.gradient(blend(top, p.Highlight, 0.03), blend(bottom, p.Highlight, 0.09))
	s.accentButton.border = darker(p.Highlight, 1.25)
	s.accentHover = s.accentButton
	s.accentPressed = s.buttonPressed
	s.accentPressed.border = s.accentButton.border
	s.entry = box{bg: p.Base, border: outline, bw: borderAll(1), radius: 2}.padded(4, 3, 4, 3)
	s.entryFocus = s.entry
	s.entryFocus.border = darker(p.Highlight, 1.25)
	s.focus = box{border: darker(p.Highlight, 1.25), bw: borderAll(1), radius: 2, noCenter: true}
	s.noFocusRing = true // Fusion shows focus on the field's own border
	s.popup = box{bg: p.Base, border: darker(p.Window, 1.6), bw: borderAll(1)}.padded(1, 1, 1, 1)
	s.popupHover = box{bg: p.Highlight}.padded(10, 4, 10, 4)
	s.popupHoverText = p.HighlightText
	s.card = box{bg: page, border: outline, bw: borderAll(1), radius: 2}.padded(16, 14, 16, 14)
	s.tabSelected = box{border: outline, bw: [4]int{1, 1, 1, 0}, radius: 2}.padded(12, 5, 12, 5).gradient(p.Base, page)
	s.tabUnselected = box{bg: darker(p.Window, 1.03), border: outline, bw: [4]int{1, 1, 1, 1}, radius: 2}.padded(12, 5, 12, 5)
	s.tabHovered = s.tabUnselected.with(p.Window)
	s.tabSelectedText = p.WindowText
	s.sliderTrack = box{border: outline, bw: borderAll(1), radius: 2}.padded(0, 2, 0, 3).gradient(darker(p.Window, 1.14), darker(p.Window, 1.07))
	s.sliderFill = box{border: darker(p.Highlight, 1.5), bw: borderAll(1), radius: 2}.gradient(darker(p.Highlight, 1.1), lighter(p.Highlight, 1.15))
	s.progressTrack = box{bg: p.Base, border: outline, bw: borderAll(1), radius: 2}.padded(0, 12, 0, 12)
	s.progressFill = box{border: darker(p.Highlight, 1.4), bw: borderAll(1), radius: 2}.padded(0, 12, 0, 12).gradient(lighter(p.Highlight, 1.18), p.Highlight)
	s.check = checkSpec{size: 14, radius: 1, borderW: 1, markW: 1.8, bg: p.Base, border: darker(p.Window, 1.27), radioDot: 0.2,
		onFill: p.Base, onBorder: darker(p.Window, 1.27), mark: p.Text}
	s.toggle = switchSpec{w: 32, h: 16, off: p.Base, offBorder: outline, offBorderW: 1, knobOff: p.Button, knobOn: p.Button, knobOffR: 6, knobOnR: 6}
	s.thumb = thumbSpec{size: 14, fill: darker(p.Button, 1.03), top: lighter(p.Button, 1.07), border: darker(p.Window, 1.24), radius: 3}
	return s
}
