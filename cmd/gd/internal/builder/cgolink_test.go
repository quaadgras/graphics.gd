package builder

import (
	"slices"
	"testing"
)

func TestSplitPkgConfigOutput(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want []string
	}{
		{"-L/usr/lib -lfoo\n", []string{"-L/usr/lib", "-lfoo"}},
		{"  -lfoo   -lbar  ", []string{"-lfoo", "-lbar"}},
		{`-L/my\ libs -lfoo`, []string{"-L/my libs", "-lfoo"}},
		{`-L"/my libs" '-Wl,-rpath,/a b' -lfoo`, []string{"-L/my libs", "-Wl,-rpath,/a b", "-lfoo"}},
		{`'a\b' "c\"d"`, []string{`a\b`, `c"d`}},
		{`-framework Metal`, []string{"-framework", "Metal"}},
		{"\n", nil},
	} {
		got, err := splitPkgConfigOutput(tc.in)
		if err != nil || !slices.Equal(got, tc.want) {
			t.Errorf("splitPkgConfigOutput(%q) = %q, %v; want %q", tc.in, got, err, tc.want)
		}
	}
	for _, bad := range []string{`-L"/unterminated`, `trailing\`} {
		if _, err := splitPkgConfigOutput(bad); err == nil {
			t.Errorf("splitPkgConfigOutput(%q) did not fail", bad)
		}
	}
}
