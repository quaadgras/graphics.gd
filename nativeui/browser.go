package nativeui

import "strings"

// browserFromUserAgent picks the engine from a user agent string.
// Chromium's agent also says "Safari", and iOS browsers all use WebKit.
func browserFromUserAgent(ua string) Browser {
	switch {
	case strings.Contains(ua, "Firefox/"):
		return Firefox
	case strings.Contains(ua, "Chrome/") || strings.Contains(ua, "Chromium/") || strings.Contains(ua, "Edg/"):
		return Chromium
	case strings.Contains(ua, "AppleWebKit/"):
		return Safari
	}
	return Chromium
}
