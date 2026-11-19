//go:build ios || nativeui_all

package nativeui

import (
	"graphics.gd/variant/Color"
)

func init() {
	register(IOS, func(o Options) spec {
		if o.Version.Major >= 26 {
			return ios26(o.Dark, !o.ReduceTransparency)
		}
		return ios(o.Dark)
	})
}

// ios26 is UIKit on iOS 26 and later (Liquid Glass), measured on an iPhone
// 16 running iOS 27: 34pt capsule buttons (.gray() / .filled()), a 63x28
// switch with a 37pt capsule knob, 32pt capsule segmented controls, 34pt
// text fields with a hairline border and 5pt corners, a 6pt slider track
// with a 38x24 capsule thumb, a 4pt progress bar.
func ios26(dark, glass bool) spec {
	s := spec{
		fonts:       []string{"SF Pro Text", "SF Pro", ".AppleSystemUIFont", "Helvetica Neue", "Inter", "Noto Sans"},
		fontSize:    17,
		boldWeight:  600,
		noFocusRing: true,
		onAccent:    white(1),
	}
	var button, disabled, field, fieldBorder, segBar, segSelected, switchOff, switchOn, track, progressTrack, popup Color.RGBA
	if dark {
		s.accent, s.window = hex("#0091FF"), black(1)
		s.text, s.dim, s.separator, s.disabled = white(1), rgba(235, 235, 245, 0.6), rgba(84, 84, 88, 0.65), hex("#5B5B5E")
		button, disabled, field, fieldBorder = hex("#272628"), hex("#1D1D1F"), black(1), hex("#343332")
		segBar, segSelected = hex("#1F1F21"), hex("#5A5A5E")
		switchOff, switchOn, track, progressTrack = hex("#474649"), hex("#30D158"), hex("#2A2A2D"), hex("#2A2A2D")
		popup = rgba(37, 37, 37, 0.96)
		s.card = box{bg: hex("#1C1C1E"), radius: 26}
	} else {
		s.accent, s.window = hex("#0088FF"), white(1)
		s.text, s.dim, s.separator, s.disabled = black(1), rgba(60, 60, 67, 0.6), rgba(60, 60, 67, 0.29), hex("#BABABC")
		button, disabled, field, fieldBorder = hex("#EAEAEB"), hex("#EFEFF0"), white(1), hex("#CCCCCC")
		segBar, segSelected = hex("#EBEBEC"), white(1)
		switchOff, switchOn, track, progressTrack = hex("#D5D5D7"), hex("#34C759"), hex("#E6E6E6"), hex("#E3E3E5")
		popup = rgba(250, 250, 250, 0.96)
		s.card = box{bg: hex("#F2F2F7"), radius: 26}
	}
	pad := [4]float64{14, 7, 14, 7} // 17pt text: 34pt capsules
	s.button = box{bg: button, radius: 17}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonHover = s.button
	s.buttonPressed = s.button.with(shade(button, 0.92))
	s.buttonDisabled = s.button.with(disabled)
	s.accentButton = box{bg: s.accent, radius: 17}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentHover = s.accentButton
	s.accentPressed = s.accentButton.with(shade(s.accent, 0.85))
	s.arrowUpDown, s.arrowColor = true, s.text
	s.entry = box{bg: field, border: fieldBorder, bw: borderAll(1), radius: 5}.padded(8, 6, 8, 6)
	s.entryFocus = s.entry
	s.popup = box{bg: popup, radius: 22, shadow: black(0.2), shadowSize: 24, shadowY: 8}.padded(8, 8, 8, 8)
	s.popupHover = box{bg: withAlpha(s.text, 0.08), radius: 12}.padded(14, 10, 14, 10)
	s.popupHoverText = s.text
	s.card = s.card.padded(16, 12, 16, 12)
	// Segmented control: the bar's colour behind each segment, the selected
	// one a raised capsule.
	s.tabSelected = box{bg: segSelected, radius: 16, shadow: black(0.08), shadowSize: 3, shadowY: 1}.padded(20, 8, 20, 8)
	s.tabUnselected = box{bg: segBar, radius: 16}.padded(20, 8, 20, 8)
	s.tabHovered = s.tabUnselected
	s.tabSelectedText, s.tabSeparation, s.tabFontSize = s.text, 0.01, 13
	s.progressTrack = box{bg: progressTrack, radius: 2}.padded(0, 2, 0, 2)
	s.progressFill = box{bg: s.accent, radius: 2}.padded(0, 2, 0, 2)
	s.sliderTrack = box{bg: track, radius: 3}.padded(0, 3, 0, 3)
	s.sliderFill = box{bg: s.accent, radius: 3}
	// iOS has no check box: the round checkmark of lists (Reminders, Mail).
	s.check = checkSpec{size: 22, radius: 11, borderW: 1.5, markW: 2, border: rgba(60, 60, 67, 0.3), radioDot: 0.2}
	if dark {
		s.check.border = rgba(235, 235, 245, 0.3)
	}
	s.toggle = switchSpec{w: 63, h: 28, off: switchOff, on: switchOn, knobOff: white(1), knobOn: white(1), knobOffR: 12, knobOnR: 12, knobW: 37, knobShadow: true}
	s.thumb = thumbSpec{size: 24, w: 38, fill: white(1), border: black(0.06), shadow: true}
	if glass {
		// Liquid Glass on the segmented control's bar.
		tint := withAlpha(segBar, 0.6)
		s.tabUnselected = box{bg: tint, radius: 16, glass: true}.padded(20, 8, 20, 8)
		s.tabHovered = s.tabUnselected
	}
	return s
}

// ios is UIKit on iOS 17+, laid out like a grouped Settings screen.
func ios(dark bool) spec {
	s := spec{
		fonts:       []string{"SF Pro Text", "SF Pro", ".AppleSystemUIFont", "Helvetica Neue", "Inter", "Noto Sans"},
		fontSize:    17,
		boldWeight:  600,
		noFocusRing: true,
	}
	var fill, segmentSelected, popup Color.RGBA
	if dark {
		s.accent = hex("#0A84FF")
		s.window = hex("#000000")
		s.text, s.dim, s.separator = white(1), rgba(235, 235, 245, 0.6), rgba(84, 84, 88, 0.65)
		s.entry = box{bg: black(1), border: hex("#333233"), bw: borderAll(1), radius: 5}
		s.disabled = rgba(235, 235, 245, 0.3)
		fill = rgba(118, 118, 128, 0.24)
		segmentSelected = hex("#636366")
		popup = rgba(37, 37, 37, 0.96)
		s.card = box{bg: hex("#1C1C1E"), radius: 10}
		s.toggle = switchSpec{off: rgba(120, 120, 128, 0.32), on: hex("#30D158")}
		s.thumb = thumbSpec{fill: white(1), border: black(0.04), shadow: true}
	} else {
		s.accent = hex("#007AFF")
		s.window = white(1)
		s.text, s.dim, s.separator = black(1), rgba(60, 60, 67, 0.6), rgba(60, 60, 67, 0.29)
		s.entry = box{bg: white(1), border: hex("#C0C0C0"), bw: borderAll(1), radius: 5}
		s.disabled = rgba(60, 60, 67, 0.3)
		fill = rgba(118, 118, 128, 0.12)
		segmentSelected = white(1)
		popup = rgba(242, 242, 242, 0.96)
		s.card = box{bg: hex("#F2F2F7"), radius: 10}
		s.toggle = switchSpec{off: rgba(120, 120, 128, 0.16), on: hex("#34C759")}
		s.thumb = thumbSpec{fill: white(1), border: black(0.04), shadow: true}
	}
	s.onAccent = white(1)

	// Gray (bordered) buttons with accent labels, filled for the accent one.
	pad := [4]float64{14, 7, 14, 7}
	s.button = box{bg: fill, radius: 8}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonHover = s.button
	s.buttonPressed = s.button.with(withAlpha(fill, fill.A*2))
	s.buttonDisabled = s.button
	s.buttonText = s.accent
	s.accentButton = box{bg: s.accent, radius: 8}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentHover = s.accentButton
	s.accentPressed = s.accentButton.with(withAlpha(s.accent, 0.7))
	s.arrowColor = s.accent

	// UITextField's default rounded rect: a hairline border, 5pt corners.
	s.entry = s.entry.padded(8, 7, 8, 7)
	s.entryFocus = s.entry

	s.popup = box{bg: popup, radius: 13, shadow: black(0.2), shadowSize: 20, shadowY: 6}.padded(0, 0, 0, 0)
	s.popupHover = box{bg: withAlpha(s.text, 0.08)}.padded(16, 11, 16, 11)
	s.popupHoverText = s.text
	s.card = s.card.padded(16, 12, 16, 12)

	// Segmented control.
	s.tabSelected = box{bg: segmentSelected, radius: 7, shadow: black(0.12), shadowSize: 3, shadowY: 1}.padded(20, 8, 20, 8)
	s.tabUnselected = box{bg: fill, radius: 7}.padded(20, 8, 20, 8)
	s.tabHovered = s.tabUnselected
	s.tabSelectedText, s.tabSeparation, s.tabFontSize = s.text, 0.01, 13
	s.arrowUpDown = true

	track := rgba(120, 120, 128, 0.2)
	s.progressTrack = box{bg: track, radius: 2}.padded(0, 2, 0, 2)
	s.progressFill = box{bg: s.accent, radius: 2}.padded(0, 2, 0, 2)
	s.sliderTrack = box{bg: track, radius: 2}.padded(0, 2, 0, 2)
	s.sliderFill = box{bg: s.accent, radius: 2}

	// iOS has no check box: use the round checkmark from lists (Reminders, Mail).
	s.check = checkSpec{size: 22, radius: 11, borderW: 1.5, markW: 2, border: rgba(60, 60, 67, 0.3), radioDot: 0.2}
	if dark {
		s.check.border = rgba(235, 235, 245, 0.3)
	}
	s.toggle.w, s.toggle.h, s.toggle.knobOffR, s.toggle.knobOnR, s.toggle.knobShadow = 51, 31, 13.5, 13.5, true
	s.toggle.knobOff, s.toggle.knobOn = white(1), white(1)
	s.thumb.size = 28
	return s
}
