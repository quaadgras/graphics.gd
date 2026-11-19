package nativeui

import "testing"

func TestBrowserFromUserAgent(t *testing.T) {
	for ua, want := range map[string]browserID{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/155.0.0.0 Safari/537.36":                         browserChromium,
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/155.0.0.0 Safari/537.36 Edg/155.0.0.0": browserChromium,
		"Mozilla/5.0 (X11; Linux x86_64; rv:160.0) Gecko/20100101 Firefox/160.0":                                                        browserFirefox,
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15":         browserSafari,
		"": browserChromium,
	} {
		if got := browserFromUserAgent(ua); got != want {
			t.Errorf("%q: got %v, want %v", ua, got, want)
		}
	}
}

func TestIsOneUI(t *testing.T) {
	if !isOneUI("samsung", "80500") || !isOneUI("samsung", "") || isOneUI("Google", "") {
		t.Fatal("isOneUI")
	}
}
