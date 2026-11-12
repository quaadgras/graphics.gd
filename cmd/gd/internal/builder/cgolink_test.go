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

func TestLd64Flags(t *testing.T) {
	in := []string{
		"-L/x", "-lfoo", "/abs/libbar.a", "-framework", "Metal", "-weak_framework", "ARKit",
		"-Wl,-rpath,@executable_path", "-Xlinker", "-dead_strip", "-force_load", "/abs/libz.a",
		"-pthread", "-O2", "-g", "-m64", "-fPIC", "-flto=thin", "-fuse-ld=lld", "-mios-version-min=15.0",
		"-target", "arm64-apple-ios", "-arch", "arm64", "-isysroot", "/sdk", "-map", "out.map", "-filelist", "files",
	}
	want := []string{
		"-L/x", "-lfoo", "/abs/libbar.a", "-framework", "Metal", "-weak_framework", "ARKit",
		"-rpath", "@executable_path", "-dead_strip", "-force_load", "/abs/libz.a",
		"-map", "out.map", "-filelist", "files",
	}
	if got := ld64Flags(in); !slices.Equal(got, want) {
		t.Errorf("ld64Flags =\n%q\nwant\n%q", got, want)
	}
}
