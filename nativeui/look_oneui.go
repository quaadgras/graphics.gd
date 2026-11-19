//go:build android || nativeui_all

package nativeui

import "graphics.gd/variant/Color"

func init() { register(lookOneUI, func(o options) spec { return oneUI(o.Accent, o.Dark) }) }

// oneUI is Samsung's One UI as its own apps draw it (Samsung's SESL
// widgets), measured on a Galaxy Z Fold5 with One UI 8.5 (Android 16): the
// framework DeviceDefault widgets for 36dp pill buttons, round 20dp check
// boxes, underlined text fields and a 4dp slider track with a hollow 16dp
// thumb; the Settings app for the colours (a black page with #171719 cards
// in dark mode) and the 32x19dp switch with its knob inside the track. The
// accent is the system's (wallpaper) colour.
func oneUI(accent Color.RGBA, dark bool) spec {
	s := spec{
		fonts:       []string{"One UI Sans", "SamsungOneUI", "Roboto", "Noto Sans"},
		fontSize:    16,
		boldWeight:  600,
		buttonBold:  true,
		noFocusRing: true,
		onAccent:    white(1),
	}
	if accent.A == 0 {
		accent = hex("#387AFF")
	}
	s.accent = accent
	var button, disabled, line, track, ring, switchOff, knobOff, thumbInside, popup Color.RGBA
	if dark {
		s.window, s.text, s.dim, s.separator, s.disabled = black(1), hex("#FDFDFF"), hex("#A4A3A8"), hex("#3A3A3C"), white(0.4)
		button, disabled, line, track, ring = hex("#3A3A3C"), hex("#252527"), hex("#8C8C8C"), hex("#3A3A3C"), hex("#707071")
		switchOff, knobOff, thumbInside, popup = hex("#5A5A5E"), hex("#FCFCFE"), black(1), hex("#171719")
		s.buttonText = hex("#FAFAFF")
	} else {
		s.window, s.text, s.dim, s.separator, s.disabled = hex("#FAFAFF"), hex("#323232"), hex("#848487"), hex("#DCDCDF"), black(0.35)
		button, disabled, line, track, ring = hex("#E4E4E7"), hex("#F1F1F5"), hex("#8C8C8C"), hex("#E7E7EA"), hex("#B5B5B7")
		switchOff, knobOff, thumbInside, popup = hex("#C0C0C5"), white(1), hex("#FCFCFF"), white(1)
		s.buttonText = hex("#313132")
	}
	pad := [4]float64{24, 7, 24, 7} // 36dp pills
	s.button = box{bg: button, radius: 18}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonHover = s.button
	s.buttonPressed = s.button.with(shade(button, 0.92))
	s.buttonDisabled = s.button.with(disabled)
	s.accentButton = box{bg: accent, radius: 18}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentHover = s.accentButton
	s.accentPressed = s.accentButton.with(shade(accent, 0.85))
	s.entry = box{border: line, bw: [4]int{0, 0, 0, 1}, noCenter: true}.padded(4, 8, 4, 8)
	s.entryFocus = box{border: accent, bw: [4]int{0, 0, 0, 2}, noCenter: true}.padded(4, 8, 4, 7)
	s.focus = box{}
	s.popup = box{bg: popup, radius: 26, shadow: black(0.2), shadowSize: 20, shadowY: 6}.padded(0, 12, 0, 12)
	if dark { // menus are raised above the #171719 cards
		s.popup.bg = hex("#2A2A2D")
	}
	s.popupHover = box{bg: withAlpha(s.text, 0.08)}.padded(24, 12, 24, 12)
	s.popupHoverText = s.text
	s.card = box{bg: popup, radius: 26}.padded(20, 16, 20, 16)
	// Sub tabs: the selected one is a filled pill.
	s.tabSelected = box{bg: button, radius: 18}.padded(16, 7, 16, 7)
	s.tabUnselected = box{}.padded(16, 7, 16, 7)
	s.tabHovered = s.tabUnselected
	s.tabSelectedText, s.tabSeparation = s.text, 4
	s.tabPanel = &box{bg: popup, radius: 26}
	s.tabPanel.pad, s.tabPanel.setPad = [4]float64{20, 16, 20, 16}, true // the page is a card: inset like One UI lists
	s.progressTrack = box{bg: track, radius: 2}.padded(0, 2, 0, 2)
	s.progressFill = box{bg: accent, radius: 2}.padded(0, 2, 0, 2)
	s.sliderTrack = box{bg: track, radius: 2}.padded(0, 2, 0, 2)
	s.sliderFill = box{bg: accent, radius: 2}
	s.check = checkSpec{size: 20, radius: 10, borderW: 2, markW: 2, bg: s.window, border: ring,
		radioRing: true, radioGap: s.window, radioDot: 0.28}
	s.checkRow = 32
	// Sub-headers are small and grey; text fields use a bigger font.
	s.headerColor, s.headerFontSize, s.entryFontSize = s.dim, 14, 18
	// Settings' switch: the knob inside an accent track.
	s.toggle = switchSpec{w: 32, h: 19, off: switchOff, on: accent, knobOff: knobOff, knobOn: hex("#FCFCFE"), knobOffR: 7, knobOnR: 7}
	s.thumb = thumbSpec{size: 16, fill: thumbInside, border: accent, borderW: 2}
	return s
}
