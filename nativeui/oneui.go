package nativeui

import "strings"

// isOneUI reports whether an Android device runs Samsung's One UI, from its
// ro.product.manufacturer and ro.build.version.oneui properties (the latter
// is the One UI version, 80500 for 8.5, and only set on One UI).
func isOneUI(manufacturer, oneUIVersion string) bool {
	return oneUIVersion != "" || strings.EqualFold(manufacturer, "samsung")
}
