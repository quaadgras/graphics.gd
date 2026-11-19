//go:build !(darwin || windows || android) || nativeui_all

package nativeui

import (
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

func init() { register(Adwaita, func(o Options) spec { return adwaita(o.Dark) }) }

func adwaita(dark bool) spec {
	s := spec{
		fonts:      []string{"Adwaita Sans", "Cantarell", "Inter", "Noto Sans"},
		fontSize:   15,
		boldWeight: 700,
		buttonBold: true,
		accent:     hex("#3584E4"),
		onAccent:   white(1),
	}
	var fg func(Float.X) Color.RGBA
	if dark {
		fg = white
		s.window = hex("#242424")
		s.text, s.dim, s.disabled, s.separator = white(1), white(0.55), white(0.3), white(0.1)
		s.popup = box{bg: hex("#383838"), radius: 12, border: white(0.07), bw: borderAll(1), shadow: black(0.35), shadowSize: 12, shadowY: 3}
		s.card = box{bg: white(0.08), radius: 12, shadow: black(0.25), shadowSize: 2, shadowY: 1}
		s.toggle = switchSpec{off: white(0.2), knobOff: hex("#E8E8E8"), knobOn: white(1), knobShadow: true}
		s.thumb = thumbSpec{fill: hex("#E8E8E8"), border: black(0.2), shadow: true}
	} else {
		fg = black
		s.window = hex("#FAFAFA")
		s.text, s.dim, s.disabled, s.separator = black(0.8), black(0.55), black(0.3), black(0.1)
		s.popup = box{bg: white(1), radius: 12, border: black(0.08), bw: borderAll(1), shadow: black(0.18), shadowSize: 12, shadowY: 3}
		s.card = box{bg: white(1), radius: 12, border: black(0.03), bw: borderAll(1), shadow: black(0.08), shadowSize: 3, shadowY: 1}
		s.toggle = switchSpec{off: black(0.15), knobOff: white(1), knobOn: white(1), knobShadow: true}
		s.thumb = thumbSpec{fill: white(1), border: black(0.14), shadow: true}
	}
	s.button = box{bg: fg(0.1), radius: 6}.padded(16, 7, 16, 7)
	s.buttonHover = s.button.with(fg(0.15))
	s.buttonPressed = s.button.with(fg(0.28))
	s.buttonDisabled = s.button.with(fg(0.05))
	s.accentButton = box{bg: s.accent, radius: 6}.padded(16, 7, 16, 7)
	s.accentHover = s.accentButton.with(shade(s.accent, 1.08))
	s.accentPressed = s.accentButton.with(shade(s.accent, 0.85))
	s.entry = box{bg: fg(0.1), radius: 6}.padded(9, 7, 9, 7)
	s.entryFocus = s.entry
	s.focus = box{border: s.accent, bw: borderAll(2), radius: 8, expand: 1, noCenter: true}
	s.focus.border.A = 0.5
	s.popup = s.popup.padded(6, 6, 6, 6)
	s.popupHover = box{bg: fg(0.08), radius: 6}.padded(10, 6, 10, 6)
	s.popupHoverText = s.text
	s.card = s.card.padded(16, 14, 16, 14)
	s.tabSelected = box{bg: fg(0.1), radius: 6}.padded(14, 6, 14, 6)
	s.tabUnselected = box{}.padded(14, 6, 14, 6)
	s.tabHovered = box{bg: fg(0.05), radius: 6}.padded(14, 6, 14, 6)
	s.progressTrack = box{bg: fg(0.15), radius: 3}.padded(0, 3, 0, 3)
	s.progressFill = box{bg: s.accent, radius: 3}.padded(0, 3, 0, 3)
	s.sliderTrack = box{bg: fg(0.15), radius: 3}.padded(0, 3, 0, 3)
	s.sliderFill = box{bg: s.accent, radius: 3}
	s.check = checkSpec{size: 18, radius: 5, borderW: 2, markW: 2.5, border: fg(0.15), radioDot: 0.17}
	s.toggle.w, s.toggle.h, s.toggle.knobOffR, s.toggle.knobOnR = 46, 26, 10, 10
	s.thumb.size = 20
	return s
}
