//go:build (linux && !android) || freebsd || openbsd || netbsd || dragonfly

package nativeui

import (
	"graphics.gd/nativeui/internal/cosmic"
	"strings"

	"graphics.gd/nativeui/internal/desktop"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

// kdePalette maps kdeglobals colours onto a palette, Breeze's defaults
// filling in what the scheme leaves out.
func kdePalette(kde map[string][3]uint8, dark bool) *Palette {
	p := breezePalettes[0]
	if dark {
		p = breezePalettes[1]
	}
	return qtPalette(p, kde)
}

// qtPalette overrides a palette with the colours a desktop configures.
func qtPalette(p Palette, kde map[string][3]uint8) *Palette {
	set := func(dst *Color.RGBA, key string) {
		if c, ok := kde[key]; ok {
			*dst = Color.RGBA{R: Float.X(c[0]) / 255, G: Float.X(c[1]) / 255, B: Float.X(c[2]) / 255, A: 1}
		}
	}
	set(&p.Window, "Window.BackgroundNormal")
	set(&p.WindowText, "Window.ForegroundNormal")
	set(&p.Button, "Button.BackgroundNormal")
	set(&p.ButtonText, "Button.ForegroundNormal")
	set(&p.Base, "View.BackgroundNormal")
	set(&p.Text, "View.ForegroundNormal")
	set(&p.Highlight, "Selection.BackgroundNormal")
	set(&p.HighlightText, "Selection.ForegroundNormal")
	return &p
}

// platformLook is Adwaita until a desktop's settings say otherwise.
func platformLook(Version) Look { return Adwaita }

func init() { platformDetect = detectDesktop }

// detectDesktop reads the desktop environment's settings: a GTK theme with
// CSS files on disk is loaded as is (built-in themes like Adwaita keep the
// Adwaita look), along with the desktop's font.
func detectDesktop(opts *Options) {
	s := desktop.Env{}.Read()
	opts.Desktop = string(s.Desktop)
	opts.Font = Typeface{Family: s.Font.Family, Weight: s.Font.Weight, Size: Float.X(s.Pixels())}
	if s.Dark {
		opts.Dark = true
	}
	switch {
	case s.Desktop == desktop.COSMIC:
		if th, err := cosmic.Load(desktop.Env{}.ConfigHome()); err == nil {
			opts.Look, opts.cosmicTheme, opts.Dark = COSMIC, th, th.Dark
		}
	case s.Desktop == desktop.KDE:
		opts.Look = Breeze
		opts.Palette = kdePalette(s.KDE, opts.Dark)
	case s.Desktop == desktop.LXQt:
		opts.Look = Fusion
		base := fusionPalette
		if strings.EqualFold(s.QtStyle, "Breeze") {
			opts.Look = Breeze
			base = breezePalettes[0]
		}
		opts.Palette = qtPalette(base, s.KDE)
	case s.Desktop == desktop.LXDE && s.ThemeDir2 != "": // LXDE's own apps are GTK 2
		opts.Look, opts.GTKTheme = GTK2, s.ThemeDir2
	case s.ThemeDir != "":
		opts.Look, opts.GTKTheme = GTK, s.ThemeDir
	case s.ThemeDir2 != "":
		opts.Look, opts.GTKTheme = GTK2, s.ThemeDir2
	}
}
