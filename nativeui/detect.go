package nativeui

import (
	"strconv"
	"strings"

	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/OS"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

// options are what a theme is built from: the look and the system settings it
// takes, detected by detect.
type options struct {
	Look    lookID
	Version osVersion // of the operating system, selects between a look's generations
	Dark    bool
	Accent  Color.RGBA // zero value: use the look's default accent
	Scale   Float.X    // zero value: 1

	// Palette is the desktop's colour scheme, where it has one (KDE's
	// kdeglobals); looks that support it take their colours from it.
	Palette *palette

	// Browser is the web browser's engine, for the Web look.
	Browser browserID

	// Font is the interface font the desktop asks for, if it says.
	Font typeface

	// textScale multiplies the look's font size (Windows' "Text size"
	// accessibility setting), zero means 1.
	textScale Float.X

	// cosmicTheme is the COSMIC theme files, read by detect on COSMIC.
	cosmicTheme any

	// Desktop is the Linux/BSD desktop environment ("XFCE", "GNOME"...), if any.
	Desktop string

	// GTKTheme is the directory of a GTK 3 theme (containing gtk-3.0/gtk.css)
	// for the GTK look, on Linux and the BSDs.
	GTKTheme string

	// Accessibility preferences reported by the system.
	HighContrast        bool
	ReduceMotion        bool
	ReduceTransparency  bool
	ScreenReaderEnabled bool
}

// Palette is a desktop colour scheme, in Qt's terms.
type palette struct {
	Window, WindowText, Button, ButtonText, Base, Text, Highlight, HighlightText Color.RGBA

	// Disabled is the colour of disabled text (Windows' GrayText), zero
	// where the scheme doesn't say.
	Disabled Color.RGBA
}

// Browser is a web browser engine. Each draws form controls its own way,
// the same on every desktop operating system.
type browserID int

const (
	browserChromium browserID = iota // Chrome, Edge, Opera, Brave...
	browserFirefox
	browserSafari // WebKit on macOS
)

func (b browserID) String() string {
	switch b {
	case browserFirefox:
		return "Firefox"
	case browserSafari:
		return "Safari"
	}
	return "Chromium"
}

// Typeface describes an interface font.
type typeface struct {
	Family string
	Weight int     // CSS weight, 400 regular, 700 bold
	Size   Float.X // pixels, before options.Scale
}

// platformDetect fills in what only some platforms can tell (desktop settings).
var platformDetect func(*options)

// Version of the operating system, as reported by OS.GetVersion. On Windows,
// Patch is the build number (Windows 11 is 10.0.22000 and later).
type osVersion struct {
	Major, Minor, Patch int
}

func parseVersion(s string) osVersion {
	var v osVersion
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '.' || r == ' ' || r == '-' })
	for i, dst := range []*int{&v.Major, &v.Minor, &v.Patch} {
		if i < len(parts) {
			*dst, _ = strconv.Atoi(parts[i])
		}
	}
	return v
}

// detect returns options matching the current system: its look and version,
// dark mode, accent color, display scale and accessibility preferences (where
// the platform reports them).
func detect() options {
	var opts options
	opts.Version = parseVersion(OS.GetVersion())
	opts.Look = platformLook(opts.Version)
	if DisplayServer.IsDarkModeSupported() {
		opts.Dark = DisplayServer.IsDarkMode()
	}
	if accent := DisplayServer.GetAccentColor(); accent.A > 0 {
		opts.Accent = accent
	}
	opts.Scale = Float.X(DisplayServer.ScreenGetScale())
	// These return -1 where the platform can't tell.
	opts.HighContrast = DisplayServer.AccessibilityShouldIncreaseContrast() == 1
	opts.ReduceMotion = DisplayServer.AccessibilityShouldReduceAnimation() == 1
	opts.ReduceTransparency = DisplayServer.AccessibilityShouldReduceTransparency() == 1
	opts.ScreenReaderEnabled = DisplayServer.AccessibilityScreenReaderActive() == 1
	if platformDetect != nil {
		platformDetect(&opts)
	}
	return opts
}
