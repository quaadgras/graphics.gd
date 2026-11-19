//go:build js

package nativeui

import "syscall/js"

func platformLook(osVersion) lookID { return lookWeb }

func init() { platformDetect = detectBrowser }

// detectBrowser picks the browser engine from the user agent and dark mode
// from the page's prefers-color-scheme, straight from the JS host.
func detectBrowser(opts *options) {
	defer func() { recover() }() // a JS runtime without navigator or matchMedia
	global := js.Global()
	if nav := global.Get("navigator"); nav.Truthy() {
		opts.Browser = browserFromUserAgent(nav.Get("userAgent").String())
	}
	if global.Get("matchMedia").Truthy() {
		opts.Dark = global.Call("matchMedia", "(prefers-color-scheme: dark)").Get("matches").Bool()
	}
}
