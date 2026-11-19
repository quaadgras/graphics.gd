package nativeui

import "strings"

// browserFromUserAgent picks the engine from a user agent string.
// Chromium's agent also says "Safari", and iOS browsers all use WebKit.
func browserFromUserAgent(ua string) browserID {
	switch {
	case strings.Contains(ua, "Firefox/"):
		return browserFirefox
	case strings.Contains(ua, "Chrome/") || strings.Contains(ua, "Chromium/") || strings.Contains(ua, "Edg/"):
		return browserChromium
	case strings.Contains(ua, "AppleWebKit/"):
		return browserSafari
	}
	return browserChromium
}
