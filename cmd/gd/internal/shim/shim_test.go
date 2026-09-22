package shim

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestCompiler(t *testing.T) {
	dir := t.TempDir()
	object := filepath.Join(dir, "os_android.os")
	if err := os.WriteFile(object, []byte("ELF"), 0644); err != nil {
		t.Fatal(err)
	}
	config := Config{Targets: map[string]Target{
		"aarch64-linux-android": {LibC: "arm64.txt", Lib: "arm64"},
		"arm-linux-androideabi": {LibC: "arm.txt"},
		"aarch64-ios":           {System: []string{"/sdk/usr/include"}, Frameworks: []string{"/sdk/Frameworks"}},
	}}
	for _, test := range []struct {
		name string
		args []string
		want []string
		libc string
	}{
		{"api level", []string{"-target", "aarch64-linux-android24", "-c", "a.c"},
			[]string{"-target", "aarch64-linux-android.24", "-c", "a.c", "-g0", "-Larm64"}, "arm64.txt"},
		{"no api level", []string{"--target=aarch64-linux-android", "-g"},
			[]string{"--target=aarch64-linux-android", "-g", "-Larm64"}, "arm64.txt"},
		{"clang architecture", []string{"-target", "armv7a-linux-androideabi24"},
			[]string{"-target", "arm-linux-androideabi.24", "-g0"}, "arm.txt"},
		{"other targets", []string{"-target", "x86_64-linux-musl"},
			[]string{"-target", "x86_64-linux-musl", "-g0"}, ""},
		{"shared objects", []string{"-shared", object, "missing.os"},
			[]string{"-shared", object + ".o", "missing.os", "-g0"}, ""},
		{"xcode", []string{"-c", "-miphoneos-version-min=14.0", "-arch", "arm64", "-I/sdk/usr/include/", "-Icore", "-isysroot", "/sdk"},
			[]string{"-target", "aarch64-ios.14.0", "-c", "-Icore", "-isysroot", "/sdk", "-g0", "-isystem", "/sdk/usr/include", "-iframework", "/sdk/Frameworks"}, ""},
	} {
		got, target := compiler(nil, test.args, config)
		if test.name == "other targets" {
			if defaulted, _ := compiler(nil, []string{"-c"}, Config{Default: "x86_64-linux-musl"}); !slices.Equal(defaulted, []string{"-target", "x86_64-linux-musl", "-c", "-g0"}) {
				t.Errorf("default target: %q", defaulted)
			}
		}
		if !slices.Equal(got, test.want) {
			t.Errorf("%s:\n got %q\nwant %q", test.name, got, test.want)
		}
		if target.LibC != test.libc {
			t.Errorf("%s: got libc %q, want %q", test.name, target.LibC, test.libc)
		}
	}
	if data, err := os.ReadFile(object + ".o"); err != nil || string(data) != "ELF" {
		t.Errorf("alias of shared object: %q %v", data, err)
	}
}

func TestSwift(t *testing.T) {
	config := Config{Swift: "app.m", Targets: map[string]Target{
		"aarch64-ios": {Frameworks: []string{"/sdk/Frameworks"}},
	}}
	got, _ := swift(nil, []string{"-frontend", "-c", "-emit-object", "-target", "arm64-apple-ios14.0", "-sdk", "/sdk", "-o", "app.o", "app.swift"}, config)
	want := []string{"-target", "aarch64-ios.14.0", "-fobjc-arc", "-fblocks", "-fvisibility=hidden", "-O2", "-g0", "-c", "app.m", "-o", "app.o", "-iframework", "/sdk/Frameworks"}
	if !slices.Equal(got, want) {
		t.Errorf("got %q\nwant %q", got, want)
	}
	if zig, key := retarget("x86_64-apple-ios14.0-simulator"); zig != "x86_64-ios.14.0-simulator" || key != "x86_64-ios-simulator" {
		t.Errorf("simulator: %q %q", zig, key)
	}
}
