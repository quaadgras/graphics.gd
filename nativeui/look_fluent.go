//go:build windows || nativeui_all

package nativeui

func init() {
	register(Windows, func(o Options) spec {
		win10 := o.Version.Major == 10 && o.Version.Patch > 0 && o.Version.Patch < 22000
		switch {
		case o.HighContrast && o.Palette != nil:
			return windowsHighContrast(*o.Palette, win10)
		case win10:
			return windows10(o.Dark)
		}
		return windows(o.Dark)
	})
}

func windows(dark bool) spec {
	s := spec{
		fonts:      []string{"Segoe UI Variable Text", "Segoe UI Variable", "Segoe UI", "Selawik", "Noto Sans"},
		fontSize:   14,
		boldWeight: 600,
	}
	if dark {
		s.accent, s.onAccent = hex("#60CDFF"), black(1)
		s.window = hex("#202020")
		s.text, s.dim, s.disabled, s.separator = white(1), white(0.79), white(0.36), white(0.08)
		s.button = box{bg: hex("#2D2D2D"), border: hex("#353535"), borderBottom: hex("#303030"), bw: borderAll(1), radius: 4}
		s.buttonHover = s.button.with(hex("#323232"))
		s.buttonPressed = s.button.with(hex("#272727"))
		s.buttonDisabled = s.button.with(hex("#2A2A2A"))
		s.entry = box{bg: hex("#2D2D2D"), border: hex("#303030"), borderBottom: hex("#9A9A9A"), bw: borderAll(1), radius: 4}
		s.entryFocus = box{bg: hex("#1F1F1F"), border: hex("#303030"), borderBottom: s.accent, bw: [4]int{1, 1, 1, 2}, radius: 4}
		s.focus = box{border: white(1), bw: borderAll(2), radius: 7, expand: 3, noCenter: true}
		s.popup = box{bg: hex("#2C2C2C"), border: black(0.2), bw: borderAll(1), radius: 8, shadow: black(0.26), shadowSize: 16, shadowY: 8}
		s.popupHover = box{bg: white(0.06), radius: 4}
		s.popupHoverText = s.text
		s.card = box{bg: white(0.05), border: black(0.1), bw: borderAll(1), radius: 8}
		s.tabSelected = box{bg: hex("#282828"), border: hex("#1D1D1D"), bw: [4]int{1, 1, 1, 0}, radius: 4, squareBottom: true}
		s.tabHovered = box{bg: white(0.04), radius: 4, squareBottom: true}
		s.tabPanel = &box{border: hex("#1D1D1D"), bw: [4]int{0, 1, 0, 0}, noCenter: true}
		s.progressTrack = box{bg: hex("#9A9A9A"), radius: 1}
		s.sliderTrack = box{bg: white(0.54), radius: 2}
		s.check = checkSpec{bg: black(0.1), border: white(0.54)}
		s.toggle = switchSpec{off: black(0.1), offBorder: white(0.6), offBorderW: 1, knobOff: hex("#CECECE"), knobOn: black(1)}
		s.thumb = thumbSpec{fill: hex("#454545"), border: hex("#4A4A4A")}
	} else {
		s.accent, s.onAccent = hex("#005FB8"), white(1)
		s.window = hex("#F3F3F3")
		s.text, s.dim, s.disabled, s.separator = black(0.894), black(0.6), black(0.36), black(0.08)
		s.button = box{bg: hex("#FBFBFB"), border: hex("#E5E5E5"), borderBottom: hex("#CCCCCC"), bw: borderAll(1), radius: 4}
		s.buttonHover = s.button.with(hex("#F6F6F6"))
		s.buttonPressed = s.button.with(hex("#F5F5F5"))
		s.buttonDisabled = s.button.with(hex("#F5F5F5"))
		s.entry = box{bg: hex("#FBFBFB"), border: hex("#E5E5E5"), borderBottom: hex("#868686"), bw: borderAll(1), radius: 4}
		s.entryFocus = box{bg: white(1), border: hex("#E5E5E5"), borderBottom: s.accent, bw: [4]int{1, 1, 1, 2}, radius: 4}
		s.focus = box{border: black(0.9), bw: borderAll(2), radius: 7, expand: 3, noCenter: true}
		s.popup = box{bg: hex("#F9F9F9"), border: black(0.06), bw: borderAll(1), radius: 8, shadow: black(0.14), shadowSize: 16, shadowY: 8}
		s.popupHover = box{bg: black(0.04), radius: 4}
		s.popupHoverText = s.text
		s.card = box{bg: white(0.7), border: black(0.06), bw: borderAll(1), radius: 8}
		s.tabSelected = box{bg: hex("#F9F9F9"), border: hex("#E5E5E5"), bw: [4]int{1, 1, 1, 0}, radius: 4, squareBottom: true}
		s.tabHovered = box{bg: black(0.03), radius: 4, squareBottom: true}
		s.tabPanel = &box{border: hex("#E5E5E5"), bw: [4]int{0, 1, 0, 0}, noCenter: true}
		s.progressTrack = box{bg: hex("#868686"), radius: 1}
		s.sliderTrack = box{bg: black(0.45), radius: 2}
		s.check = checkSpec{bg: black(0.02), border: black(0.45)}
		s.toggle = switchSpec{off: black(0.02), offBorder: black(0.6), offBorderW: 1, knobOff: black(0.6), knobOn: white(1)}
		s.thumb = thumbSpec{fill: white(1), border: black(0.08), shadow: true}
	}
	pad := [4]float64{11, 6, 11, 7} // 32px tall with Segoe UI's 19px line
	for _, b := range []*box{&s.button, &s.buttonHover, &s.buttonPressed, &s.buttonDisabled} {
		*b = b.padded(pad[0], pad[1], pad[2], pad[3])
	}
	s.accentButton = box{bg: s.accent, radius: 4}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentHover = s.accentButton.with(shade(s.accent, 1.08))
	s.accentPressed = s.accentButton.with(shade(s.accent, 0.88))
	s.entry = s.entry.padded(10, 6, 10, 7)
	s.entryFocus = s.entryFocus.padded(10, 6, 10, 7)
	s.popup = s.popup.padded(4, 4, 4, 4)
	s.popupHover = s.popupHover.padded(10, 4, 10, 4)
	s.card = s.card.padded(16, 14, 16, 14)
	// WinUI-style tabs: the selected one a raised sheet joined to the
	// page, the others flat (WPF's Fluent TabControl).
	s.tabSelected = s.tabSelected.padded(6, 8, 6, 9)
	s.tabUnselected = box{}.padded(6, 8, 6, 9)
	s.tabHovered = s.tabHovered.padded(6, 8, 6, 9)
	s.tabSelectedText, s.tabSeparation = s.text, 0.01
	s.tabPanel.pad, s.tabPanel.setPad = [4]float64{16, 16, 16, 16}, true
	s.progressTrack = s.progressTrack.padded(0, 1, 0, 1)
	s.progressFill = box{bg: s.accent, radius: 1}.padded(0, 1, 0, 1)
	s.checkRow = 32
	s.sliderTrack = s.sliderTrack.padded(0, 2, 0, 2)
	s.sliderFill = box{bg: s.accent, radius: 2}
	s.check.size, s.check.radius, s.check.borderW, s.check.markW, s.check.radioDot = 20, 4, 1, 1.5, 0.25
	s.toggle.w, s.toggle.h, s.toggle.knobOffR, s.toggle.knobOnR = 40, 20, 6, 7
	s.thumb.size, s.thumb.dot = 20, 5
	return s
}

// windows10 is Windows 10: in light mode the Win32 common controls that
// most desktop programs show, in dark mode the UWP controls of Settings
// (Win32 controls have no dark mode there).
func windows10(dark bool) spec {
	s := spec{fonts: []string{"Segoe UI", "Selawik", "Noto Sans"}, fontSize: 12, boldWeight: 600}
	pad := [4]float64{10, 3, 10, 4}
	if dark {
		s.fontSize = 14
		pad = [4]float64{11, 4, 11, 6}
		s.accent, s.onAccent = hex("#0078D7"), white(1)
		s.window = hex("#1F1F1F")
		s.text, s.dim, s.disabled, s.separator = white(1), white(0.6), white(0.4), white(0.2)
		s.button = box{bg: hex("#333333"), border: hex("#333333"), bw: borderAll(2)}
		s.buttonHover = box{bg: hex("#333333"), border: hex("#858585"), bw: borderAll(2)}
		s.buttonPressed = box{bg: hex("#666666"), border: hex("#666666"), bw: borderAll(2)}
		s.buttonDisabled = s.button
		s.accentButton = box{bg: s.accent, border: s.accent, bw: borderAll(2)}
		s.accentHover = box{bg: s.accent, border: hex("#858585"), bw: borderAll(2)}
		s.accentPressed = box{bg: hex("#666666"), border: hex("#666666"), bw: borderAll(2)}
		s.entry = box{bg: hex("#131313"), border: white(0.4), bw: borderAll(2)}.padded(10, 4, 10, 5)
		s.entryFocus = box{bg: black(1), border: s.accent, bw: borderAll(2)}.padded(10, 4, 10, 5)
		s.focus = box{border: white(1), bw: borderAll(2), expand: 2, noCenter: true}
		s.popup = box{bg: hex("#2B2B2B"), border: hex("#767676"), bw: borderAll(1), shadow: black(0.3), shadowSize: 8, shadowY: 4}.padded(0, 4, 0, 4)
		s.popupHover = box{bg: white(0.1)}.padded(12, 6, 12, 6)
		s.popupHoverText = s.text
		s.card = box{bg: white(0.05)}.padded(16, 14, 16, 14)
		s.tabSelected = box{border: s.accent, bw: [4]int{0, 0, 0, 2}, noCenter: true}.padded(12, 6, 12, 6)
		s.tabUnselected = box{}.padded(12, 6, 12, 6)
		s.tabHovered = box{}.padded(12, 6, 12, 6)
		s.progressTrack = box{bg: white(0.2)}.padded(0, 2, 0, 2)
		s.progressFill = box{bg: s.accent}.padded(0, 2, 0, 2)
		s.sliderTrack = box{bg: white(0.4)}.padded(0, 1, 0, 1)
		s.sliderFill = box{bg: s.accent}
		s.check = checkSpec{size: 20, borderW: 2, markW: 2, border: white(0.8), mark: white(1), radioDot: 0.25}
		s.toggle = switchSpec{w: 44, h: 20, offBorder: white(0.8), offBorderW: 2, knobOff: white(0.8), knobOn: white(1), knobOffR: 5, knobOnR: 5}
		s.thumb = thumbSpec{size: 16, fill: s.accent, border: s.accent, radius: 4, top: s.accent}
	} else {
		s.accent, s.onAccent = hex("#0078D7"), black(1)
		s.window = hex("#F0F0F0")
		s.text, s.dim, s.disabled, s.separator = black(1), hex("#6D6D6D"), hex("#838383"), hex("#A0A0A0")
		s.button = box{bg: hex("#E1E1E1"), border: hex("#ADADAD"), bw: borderAll(1)}
		s.buttonHover = box{bg: hex("#E5F1FB"), border: s.accent, bw: borderAll(1)}
		s.buttonPressed = box{bg: hex("#CCE4F7"), border: hex("#005499"), bw: borderAll(1)}
		s.buttonDisabled = box{bg: hex("#CCCCCC"), border: hex("#BFBFBF"), bw: borderAll(1)}
		// The default push button: a plain button with a blue frame.
		s.accentButton = box{bg: hex("#E1E1E1"), border: s.accent, bw: borderAll(2)}
		s.accentHover = box{bg: hex("#E5F1FB"), border: s.accent, bw: borderAll(2)}
		s.accentPressed = box{bg: hex("#CCE4F7"), border: hex("#005499"), bw: borderAll(2)}
		s.accentText = black(1)
		s.entry = box{bg: white(1), border: hex("#7A7A7A"), bw: borderAll(1)}.padded(4, 3, 4, 3)
		s.entryFocus = box{bg: white(1), border: s.accent, bw: borderAll(1)}.padded(4, 3, 4, 3)
		s.focus = box{border: black(0.55), bw: borderAll(1), expand: -3, noCenter: true}
		s.popup = box{bg: hex("#F2F2F2"), border: hex("#CCCCCC"), bw: borderAll(1), shadow: black(0.3), shadowSize: 4, shadowY: 2}.padded(2, 2, 2, 2)
		s.popupHover = box{bg: hex("#91C9F7")}.padded(10, 3, 10, 3)
		s.popupHoverText = s.text
		s.card = box{bg: white(1), border: hex("#D9D9D9"), bw: borderAll(1)}.padded(12, 10, 12, 10)
		s.tabSelected = box{bg: white(1), border: hex("#D9D9D9"), bw: [4]int{1, 1, 1, 0}}.padded(8, 3, 8, 3)
		s.tabUnselected = box{bg: hex("#F0F0F0"), border: hex("#D9D9D9"), bw: borderAll(1)}.padded(6, 2, 6, 2)
		s.tabHovered = box{bg: hex("#D8EAF9"), border: hex("#D9D9D9"), bw: borderAll(1)}.padded(6, 2, 6, 2)
		s.progressTrack = box{bg: hex("#E6E6E6"), border: hex("#BCBCBC"), bw: borderAll(1)}.padded(0, 7, 0, 7)
		s.progressFill = box{bg: hex("#06B025")}.padded(0, 7, 0, 7)
		s.sliderTrack = box{bg: hex("#E7EAEA"), border: hex("#D6D6D6"), bw: borderAll(1)}.padded(0, 2, 0, 2)
		s.sliderFill = s.sliderTrack
		s.check = checkSpec{size: 13, borderW: 1, markW: 1.5, bg: white(1), border: hex("#333333"),
			onFill: white(1), onBorder: hex("#333333"), mark: hex("#333333"), radioDot: 0.2}
		s.toggle = switchSpec{w: 44, h: 20, offBorder: hex("#333333"), offBorderW: 2, knobOff: hex("#333333"), knobOn: white(1), knobOffR: 5, knobOnR: 5}
		s.thumb = thumbSpec{size: 11, fill: hex("#007AD9"), border: hex("#007AD9"), radius: 1, top: hex("#007AD9")}
	}
	for _, b := range []*box{&s.button, &s.buttonHover, &s.buttonPressed, &s.buttonDisabled, &s.accentButton, &s.accentHover, &s.accentPressed} {
		*b = b.padded(pad[0], pad[1], pad[2], pad[3])
	}
	return s
}

// windowsHighContrast draws every control in the user's contrast theme
// colours (Settings, Accessibility, Contrast themes), the way WinUI does:
// outlines in the text colour, the highlight colour for selection and
// accented controls.
func windowsHighContrast(p Palette, win10 bool) spec {
	r := 4
	fonts, size := []string{"Segoe UI Variable Text", "Segoe UI Variable", "Segoe UI", "Selawik", "Noto Sans"}, 14
	if win10 {
		r, fonts, size = 0, []string{"Segoe UI", "Selawik", "Noto Sans"}, 12
	}
	disabled := p.Disabled
	if disabled.A == 0 {
		disabled = blend(p.Window, p.WindowText, 0.5)
	}
	s := spec{fonts: fonts, fontSize: size, boldWeight: 600,
		window: p.Window, text: p.WindowText, dim: p.WindowText, disabled: disabled, separator: p.WindowText,
		accent: p.Highlight, onAccent: p.HighlightText, buttonText: p.ButtonText}
	pad := [4]float64{11, 5, 11, 6}
	s.button = box{bg: p.Button, border: p.ButtonText, bw: borderAll(1), radius: r}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonHover = box{bg: p.Button, border: p.Highlight, bw: borderAll(2), radius: r}.padded(pad[0], pad[1], pad[2], pad[3])
	s.buttonPressed = s.buttonHover
	s.buttonDisabled = box{bg: p.Button, border: disabled, bw: borderAll(1), radius: r}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentButton = box{bg: p.Highlight, border: p.Highlight, bw: borderAll(1), radius: r}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentHover = box{bg: p.Highlight, border: p.WindowText, bw: borderAll(2), radius: r}.padded(pad[0], pad[1], pad[2], pad[3])
	s.accentPressed = s.accentHover
	s.entry = box{bg: p.Base, border: p.Text, bw: borderAll(1), radius: r}.padded(10, 5, 10, 6)
	s.entryFocus = box{bg: p.Base, border: p.Highlight, bw: borderAll(2), radius: r}.padded(10, 5, 10, 6)
	s.focus = box{border: p.WindowText, bw: borderAll(2), radius: r + 3, expand: 3, noCenter: true}
	s.popup = box{bg: p.Window, border: p.WindowText, bw: borderAll(1), radius: 2 * r}.padded(4, 4, 4, 4)
	s.popupHover = box{bg: p.Highlight, radius: r}.padded(10, 4, 10, 4)
	s.popupHoverText = p.HighlightText
	s.card = box{bg: p.Window, border: p.WindowText, bw: borderAll(1), radius: 2 * r}.padded(16, 14, 16, 14)
	s.tabSelected = box{border: p.Highlight, bw: [4]int{0, 0, 0, 3}, noCenter: true}.padded(12, 6, 12, 6)
	s.tabUnselected = box{}.padded(12, 6, 12, 6)
	s.tabHovered = box{border: p.WindowText, bw: [4]int{0, 0, 0, 1}, noCenter: true}.padded(12, 6, 12, 6)
	s.progressTrack = box{bg: p.Window, border: p.WindowText, bw: borderAll(1), radius: 2}.padded(0, 2, 0, 2)
	s.progressFill = box{bg: p.Highlight, radius: 2}.padded(0, 2, 0, 2)
	s.sliderTrack = box{bg: p.WindowText, radius: 2}.padded(0, 2, 0, 2)
	s.sliderFill = box{bg: p.Highlight, radius: 2}
	s.check = checkSpec{size: 20, radius: float64(r), borderW: 1, markW: 1.5, bg: p.Window, border: p.WindowText,
		onBorder: p.Highlight, radioDot: 0.22}
	s.toggle = switchSpec{w: 40, h: 20, off: p.Window, offBorder: p.WindowText, offBorderW: 1, knobOff: p.WindowText, knobOn: p.HighlightText, knobOffR: 6, knobOnR: 7}
	s.thumb = thumbSpec{size: 20, fill: p.Window, border: p.WindowText, dot: 6}
	return s
}
