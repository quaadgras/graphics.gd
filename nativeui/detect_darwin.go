//go:build darwin && !ios

package nativeui

func platformLook(osVersion) lookID { return lookMacOS }
