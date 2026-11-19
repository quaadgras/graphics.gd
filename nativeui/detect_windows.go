//go:build windows

package nativeui

import (
	"unsafe"

	win "golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"

	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

func platformLook(osVersion) lookID { return lookWindows }

func init() { platformDetect = detectWindows }

var (
	user32                = win.NewLazySystemDLL("user32.dll")
	procGetSysColor       = user32.NewProc("GetSysColor")
	procSystemParametersW = user32.NewProc("SystemParametersInfoW")
)

// detectWindows reads what Godot's DisplayServer doesn't: the accent shade
// Fluent controls actually use, the text size, transparency effects and the
// high contrast colours.
func detectWindows(opts *options) {
	const personalize = `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`
	if v, ok := regInt(personalize, "AppsUseLightTheme"); ok {
		opts.Dark = v == 0
	}
	if v, ok := regInt(personalize, "EnableTransparency"); ok && v == 0 {
		opts.ReduceTransparency = true
	}
	// Windows derives seven shades from the accent colour (AccentPalette:
	// Light3, Light2, Light1, Base, Dark1, Dark2, Dark3 as RGBA bytes).
	// WinUI fills accent controls with Dark1 in light mode, Light2 in dark.
	if k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Accent`, registry.QUERY_VALUE); err == nil {
		if b, _, err := k.GetBinaryValue("AccentPalette"); err == nil && len(b) >= 28 {
			i := 4
			if opts.Dark {
				i = 1
			}
			opts.Accent = Color.RGBA{R: Float.X(b[i*4]) / 255, G: Float.X(b[i*4+1]) / 255, B: Float.X(b[i*4+2]) / 255, A: 1}
		}
		k.Close()
	}
	if v, ok := regInt(`Software\Microsoft\Accessibility`, "TextScaleFactor"); ok && v >= 100 && v <= 225 {
		opts.textScale = Float.X(v) / 100
	}
	// Windows sizes are for 96 dpi; Godot's screen scale is always 1 there.
	if dpi := DisplayServer.ScreenGetDpi(); dpi > 0 {
		opts.Scale = Float.X(dpi) / 96
	}
	if highContrastOn() {
		opts.HighContrast = true
		opts.Palette = &palette{
			Window: sysColor(5), WindowText: sysColor(8), // COLOR_WINDOW, COLOR_WINDOWTEXT
			Base: sysColor(5), Text: sysColor(8),
			Button: sysColor(15), ButtonText: sysColor(18), // COLOR_BTNFACE, COLOR_BTNTEXT
			Highlight: sysColor(13), HighlightText: sysColor(14), // COLOR_HIGHLIGHT, COLOR_HIGHLIGHTTEXT
			Disabled: sysColor(17), // COLOR_GRAYTEXT
		}
	}
}

func regInt(path, name string) (uint64, bool) {
	k, err := registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE)
	if err != nil {
		return 0, false
	}
	defer k.Close()
	v, _, err := k.GetIntegerValue(name)
	return v, err == nil
}

func sysColor(index int) Color.RGBA {
	c, _, _ := procGetSysColor.Call(uintptr(index)) // COLORREF: 0x00BBGGRR
	return Color.RGBA{R: Float.X(c&0xff) / 255, G: Float.X(c>>8&0xff) / 255, B: Float.X(c>>16&0xff) / 255, A: 1}
}

func highContrastOn() bool {
	type highContrast struct {
		size, flags   uint32
		defaultScheme *uint16
	}
	const spiGetHighContrast, hcfHighContrastOn = 0x42, 1
	hc := highContrast{size: uint32(unsafe.Sizeof(highContrast{}))}
	ok, _, _ := procSystemParametersW.Call(spiGetHighContrast, uintptr(hc.size), uintptr(unsafe.Pointer(&hc)), 0)
	return ok != 0 && hc.flags&hcfHighContrastOn != 0
}
