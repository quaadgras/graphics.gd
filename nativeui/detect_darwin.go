//go:build darwin && !ios

package nativeui

func platformLook(Version) Look { return MacOS }
