//go:build !(darwin || windows || android || linux || freebsd || openbsd || netbsd || dragonfly || js)

package nativeui

// platformLook on the web, until browser detection picks its own looks.
func platformLook(Version) Look { return Adwaita }
