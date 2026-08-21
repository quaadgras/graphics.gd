package tooling

import "testing"

func TestDottedVersionPrefix(t *testing.T) {
	for reported, expect := range map[string]string{
		"4.7.2.stable.official.ed1daf0bf":   "4.7.2",
		"4.7.stable.official.abcdef123\n":   "4.7",
		"4.7.stable.custom_build.4923e4443": "4.7",
		"4.8.dev.custom_build":              "4.8",
		"":                                  "",
		"unknown":                           "",
	} {
		if got := dottedVersionPrefix(reported); got != expect {
			t.Errorf("dottedVersionPrefix(%q) = %q, want %q", reported, got, expect)
		}
	}
}
