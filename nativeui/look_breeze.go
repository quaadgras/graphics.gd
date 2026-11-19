//go:build !(darwin || windows || android) || nativeui_all

package nativeui

import "graphics.gd/variant/Color"

func init() { register(Breeze, breeze) }

// breezePalettes are Plasma 6's Breeze Light and Breeze Dark colour schemes.
var breezePalettes = [2]Palette{
	{Window: hex("#eff0f1"), WindowText: hex("#232629"), Button: hex("#fcfcfc"), ButtonText: hex("#232629"),
		Base: hex("#ffffff"), Text: hex("#232629"), Highlight: hex("#3daee9"), HighlightText: hex("#ffffff")},
	{Window: hex("#202326"), WindowText: hex("#fcfcfc"), Button: hex("#292c30"), ButtonText: hex("#fcfcfc"),
		Base: hex("#141618"), Text: hex("#fcfcfc"), Highlight: hex("#3daee9"), HighlightText: hex("#fcfcfc")},
}

// breeze is KDE's Breeze widget style. Its shapes were measured from
// renders of the real style; every colour is derived from the palette, so
// any KDE colour scheme applies.
func breeze(o Options) spec {
	p := breezePalettes[0]
	if o.Dark {
		p = breezePalettes[1]
	}
	if o.Palette != nil {
		p = *o.Palette
	}
	outline := blend(p.Window, p.WindowText, 0.2)
	page := blend(p.Window, p.Base, 0.3)
	s := spec{
		fonts:      []string{"Noto Sans", "sans-serif"},
		fontSize:   13,
		boldWeight: 700,
		accent:     p.Highlight,
		onAccent:   p.HighlightText,
		window:     p.Window,
		text:       p.WindowText,
		dim:        blend(p.WindowText, p.Window, 0.35),
		disabled:   blend(p.WindowText, p.Window, 0.55),
		separator:  outline,
	}
	pad := [4]float64{12, 5, 12, 5}
	btn := box{bg: p.Button, border: blend(p.Button, p.ButtonText, 0.2), bw: borderAll(1), radius: 3}.padded(pad[0], pad[1], pad[2], pad[3])
	s.button = btn
	s.buttonHover = btn
	s.buttonHover.border = p.Highlight
	s.buttonPressed = btn.with(blend(p.Button, p.Highlight, 0.3))
	s.buttonPressed.border = p.Highlight
	s.buttonDisabled = btn
	s.buttonDisabled.border = blend(p.Button, p.ButtonText, 0.08)
	s.buttonText = p.ButtonText
	s.accentText = p.ButtonText
	s.accentButton = btn.with(blend(p.Button, p.Highlight, 0.2))
	s.accentButton.border = blend(p.Button, p.Highlight, 0.5)
	s.accentHover = s.accentButton
	s.accentHover.border = p.Highlight
	s.accentPressed = s.buttonPressed
	s.entry = box{bg: p.Base, border: outline, bw: borderAll(1), radius: 3}.padded(7, 5, 7, 5)
	s.entryFocus = s.entry
	s.entryFocus.border = p.Highlight
	s.focus = box{border: p.Highlight, bw: borderAll(1), radius: 3, noCenter: true}
	s.noFocusRing = true // Breeze highlights the border of the focused widget instead
	s.popup = box{bg: page, border: outline, bw: borderAll(1), radius: 3, shadow: Color.RGBA{A: 0.2}, shadowSize: 6, shadowY: 2}.padded(4, 4, 4, 4)
	s.popupHover = box{bg: blend(page, p.Highlight, 0.3), border: p.Highlight, bw: borderAll(1), radius: 3}.padded(8, 4, 8, 4)
	s.popupHoverText = p.WindowText
	s.card = box{bg: page, border: outline, bw: borderAll(1), radius: 3}.padded(16, 14, 16, 14)
	inactive := Color.RGBA{R: p.Window.R * 0.833, G: p.Window.G * 0.833, B: p.Window.B * 0.833, A: 1}
	s.tabSelected = box{bg: page, border: p.Highlight, bw: [4]int{0, 3, 0, 0}, radius: 3}.padded(14, 6, 14, 6)
	s.tabUnselected = box{bg: inactive}.padded(14, 6, 14, 6)
	s.tabHovered = box{bg: blend(inactive, p.Highlight, 0.2)}.padded(14, 6, 14, 6)
	s.tabSelectedText = p.WindowText
	groove := box{bg: blend(p.Window, p.WindowText, 0.14), border: blend(p.Window, p.WindowText, 0.245), bw: borderAll(1), radius: 3}.padded(0, 3, 0, 3)
	fillCol := blend(p.Highlight, blend(p.Window, p.WindowText, 0.25), 0.4)
	s.sliderTrack = groove
	s.sliderFill = box{bg: fillCol, border: blend(fillCol, p.WindowText, 0.15), bw: borderAll(1), radius: 3}
	s.progressTrack = groove
	s.progressFill = s.sliderFill.padded(0, 3, 0, 3)
	s.check = checkSpec{size: 16, radius: 3, borderW: 1, markW: 2, bg: p.Button, border: outline, radioDot: 0.22,
		onFill: blend(p.Button, p.Highlight, 0.3), onBorder: p.Highlight, mark: p.ButtonText}
	s.toggle = switchSpec{w: 36, h: 18, off: p.Window, offBorder: outline, offBorderW: 1,
		knobOff: p.Button, knobOn: p.Button, knobOffR: 7, knobOnR: 7}
	s.thumb = thumbSpec{size: 18, fill: p.Button, border: outline}
	return s
}
