// Command gen-android regenerates ../android, the Android SDK that gd
// bundles, from pinned commits of the Android Open Source Project.
//
//	go run ./cmd/gd/internal/builder/bundled/gen-android
//
// The SDK is enough to compile and link both the Go extension and the
// Godot engine itself for Android with zig, without the NDK:
//
//   - usr/include holds bionic's libc headers, the NDK platform headers
//     (android/, camera/, media/, jni.h), the Khronos EGL/GLES/OpenSL ES
//     headers and the kernel's linux/*.h headers.
//   - usr/include/<triple>/asm holds the per-architecture kernel headers.
//   - usr/lib/<triple>/*.c are link stubs generated from the platform's
//     *.map.txt symbol lists, the same lists the NDK generates its own
//     stub libraries from. gd compiles them with zig at build time; the
//     device's real libraries resolve the symbols at runtime.
//   - usr/lib/crtbegin_so.c and crtend_so.c are the C runtime objects
//     every shared library links.
//
// Everything is copied verbatim from AOSP (BSD, Apache-2.0 and MIT, see
// the NOTICE files written alongside), nothing is taken from the NDK.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// API is the Android API level the link stubs are generated for, symbols
// introduced later are left out. It matches Godot's minimum (android-24).
const API = 24

const aosp = "https://android.googlesource.com/platform/"

type repository struct {
	name   string
	commit string
	sparse []string // sparse-checkout patterns.
}

var repositories = []repository{
	{"bionic", "731631f300090436d7f5df80d50b6275c8c60a93", []string{
		"/libc/include/", "/libc/kernel/uapi/", "/libc/kernel/android/uapi/", "/libc/arch-common/bionic/",
		"/libc/NOTICE", "/libc/libc.map.txt", "/libm/libm.map.txt", "/libdl/libdl.map.txt",
	}},
	{"frameworks/native", "4f463a6b1de9198963dc6aff74154a504ba3f8f6", []string{
		"/include/android/", "/libs/nativewindow/include/android/", "/libs/arect/include/android/",
		"/libs/nativewindow/libnativewindow.map.txt", "/opengl/include/", "/opengl/libs/*.map.txt",
		"/vulkan/libvulkan/libvulkan.map.txt", "/NOTICE",
	}},
	{"frameworks/base", "1cdfff555f4a21f71ccc978290e2e212e2f8b168", []string{"/native/android/libandroid.map.txt"}},
	{"frameworks/av", "e2f098935447ca4945946de5cb69db843fe3f003", []string{
		"/camera/ndk/include/camera/", "/camera/ndk/libcamera2ndk.map.txt",
		"/media/ndk/include/media/", "/media/ndk/libmediandk.map.txt",
	}},
	{"frameworks/wilhelm", "5674f27e4c8333495518d07a676d0664d92fd74d", []string{"/include/SLES/", "/src/libOpenSLES.map.txt"}},
	{"libnativehelper", "aef2939781fc0b57b4477df7160935cdf5697919", []string{"/include_jni/jni.h"}},
	{"system/logging", "bcac7c30d88a3773a7c0bc9f5617a23a886331fd", []string{"/liblog/include/android/", "/liblog/liblog.map.txt"}},
	{"external/zlib", "46c6da99965067627e6c078e197106988d57d4ff", []string{"/libz.map.txt"}},
}

// headers maps source directories (or files) onto usr/include.
var headers = [][2]string{
	{"bionic/libc/include", ""},
	{"bionic/libc/kernel/android/uapi/linux", "linux"},
	{"frameworks/native/include/android", "android"},
	{"frameworks/native/libs/nativewindow/include/android", "android"},
	{"frameworks/native/libs/arect/include/android", "android"},
	{"system/logging/liblog/include/android", "android"},
	{"frameworks/av/camera/ndk/include/camera", "camera"},
	{"frameworks/av/media/ndk/include/media", "media"},
	{"frameworks/native/opengl/include/EGL", "EGL"},
	{"frameworks/native/opengl/include/GLES2", "GLES2"},
	{"frameworks/native/opengl/include/GLES3", "GLES3"},
	{"frameworks/native/opengl/include/KHR", "KHR"},
	{"frameworks/wilhelm/include/SLES", "SLES"},
	{"libnativehelper/include_jni/jni.h", "jni.h"},
}

// architectures gd can build for, with their kernel asm directory and the
// tag *.map.txt files use for them.
var architectures = []struct{ triple, asm, tag string }{
	{"aarch64-linux-android", "asm-arm64", "arm64"},
	{"x86_64-linux-android", "asm-x86", "x86_64"},
}

// stubs are the libraries a link may name, each generated from a map.
var stubs = map[string]string{
	"libc":            "bionic/libc/libc.map.txt",
	"libm":            "bionic/libm/libm.map.txt",
	"libdl":           "bionic/libdl/libdl.map.txt",
	"liblog":          "system/logging/liblog/liblog.map.txt",
	"libandroid":      "frameworks/base/native/android/libandroid.map.txt",
	"libnativewindow": "frameworks/native/libs/nativewindow/libnativewindow.map.txt",
	"libEGL":          "frameworks/native/opengl/libs/libEGL.map.txt",
	"libGLESv2":       "frameworks/native/opengl/libs/libGLESv2.map.txt",
	"libGLESv3":       "frameworks/native/opengl/libs/libGLESv3.map.txt",
	"libvulkan":       "frameworks/native/vulkan/libvulkan/libvulkan.map.txt",
	"libOpenSLES":     "frameworks/wilhelm/src/libOpenSLES.map.txt",
	"libcamera2ndk":   "frameworks/av/camera/ndk/libcamera2ndk.map.txt",
	"libmediandk":     "frameworks/av/media/ndk/libmediandk.map.txt",
	"libz":            "external/zlib/libz.map.txt",
}

// libpthread is named by `-lpthread`, which cgo passes for every target.
const libpthread = `// Stub so the linker can satisfy -lpthread. Android has no libpthread.so,
// pthreads live in libc.so, so this stub must stay empty: defining symbols
// here would record a DT_NEEDED on a library that does not exist on-device.
`

func main() {
	if err := generate(); err != nil {
		fmt.Fprintln(os.Stderr, "gen-android:", err)
		os.Exit(1)
	}
}

func generate() error {
	out, err := outputDirectory()
	if err != nil {
		return err
	}
	src, err := os.MkdirTemp("", "gen-android")
	if err != nil {
		return err
	}
	defer os.RemoveAll(src)
	for _, repo := range repositories {
		if err := fetch(src, repo); err != nil {
			return fmt.Errorf("%s: %w", repo.name, err)
		}
	}
	// Everything is staged first, so that only the headers reachable from
	// the public ones need to be kept.
	stage := filepath.Join(src, "include")
	for _, mapping := range headers {
		if err := copyAll(filepath.Join(src, mapping[0]), filepath.Join(stage, mapping[1])); err != nil {
			return err
		}
	}
	// linux/ and asm/ are kept whole, as they are what gets included
	// directly (the engine reads linux/input.h, linux/rtnetlink.h...), the
	// kernel's device-class directories (drm, sound...) are 'optional', as
	// are the subdirectories of linux/ (netfilter has names that only differ
	// by case, which neither go:embed nor every filesystem can represent).
	optional := map[string]bool{}
	uapi := filepath.Join(src, "bionic/libc/kernel/uapi")
	entries, err := os.ReadDir(uapi)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, "asm-") && name != "asm-generic" {
			continue // per-architecture, see below.
		}
		optional[name] = name != "linux" && name != "asm-generic"
		if err := copyAll(filepath.Join(uapi, name), filepath.Join(stage, name)); err != nil {
			return err
		}
	}
	for _, arch := range architectures {
		if err := copyAll(filepath.Join(uapi, arch.asm, "asm"), filepath.Join(stage, arch.triple, "asm")); err != nil {
			return err
		}
	}
	keep, err := closure(stage, optional)
	if err != nil {
		return err
	}
	for _, dir := range []string{"usr", "include", "lib"} { // include & lib are the pre-sysroot layout.
		if err := os.RemoveAll(filepath.Join(out, dir)); err != nil {
			return err
		}
	}
	for _, name := range keep {
		if err := copyFile(filepath.Join(stage, name), filepath.Join(out, "usr", "include", name)); err != nil {
			return err
		}
	}
	for _, arch := range architectures {
		for lib, path := range stubs {
			stub, err := stub(filepath.Join(src, path), arch.tag)
			if err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			dst := filepath.Join(out, "usr", "lib", arch.triple, lib+".c")
			if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
				return err
			}
			if err := os.WriteFile(dst, []byte(stub), 0644); err != nil {
				return err
			}
		}
	}
	for _, arch := range architectures {
		if err := os.WriteFile(filepath.Join(out, "usr", "lib", arch.triple, "libpthread.c"), []byte(libpthread), 0644); err != nil {
			return err
		}
	}
	crt := filepath.Join(src, "bionic/libc/arch-common/bionic")
	for _, name := range []string{"crtbegin_so.c", "__dso_handle_so.h", "atexit.h", "pthread_atfork.h"} {
		if err := copyFile(filepath.Join(crt, name), filepath.Join(out, "usr", "lib", name)); err != nil {
			return err
		}
	}
	crtend := "// crtend_so closes the sections crtbegin_so opens, of which zig's lld needs none.\n"
	if err := os.WriteFile(filepath.Join(out, "usr", "lib", "crtend_so.c"), []byte(crtend), 0644); err != nil {
		return err
	}
	var notice strings.Builder
	notice.WriteString("The files in this directory are copied from the Android Open Source Project:\n\n")
	for _, repo := range repositories {
		fmt.Fprintf(&notice, "\t%s%s @ %s\n", aosp, repo.name, repo.commit)
	}
	notice.WriteString("\nregenerate them with: go run ./cmd/gd/internal/builder/bundled/gen-android\n")
	if err := os.WriteFile(filepath.Join(out, "usr", "README.txt"), []byte(notice.String()), 0644); err != nil {
		return err
	}
	for from, to := range map[string]string{
		"bionic/libc/NOTICE":       "NOTICE.bionic.txt",
		"frameworks/native/NOTICE": "NOTICE.frameworks.txt",
	} {
		if err := copyFile(filepath.Join(src, from), filepath.Join(out, "usr", to)); err != nil {
			return err
		}
	}
	fmt.Printf("gen-android: wrote %d headers to %s\n", len(keep), out)
	return nil
}

// outputDirectory locates ../android relative to this source file's package,
// wherever in the module the generator happens to be run from.
func outputDirectory() (string, error) {
	root, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}").Output()
	if err != nil {
		return "", err
	}
	return filepath.Join(strings.TrimSpace(string(root)), "cmd", "gd", "internal", "builder", "bundled", "android"), nil
}

func fetch(dir string, repo repository) error {
	dst := filepath.Join(dir, repo.name)
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	git := func(args ...string) error {
		cmd := exec.Command("git", append([]string{"-C", dst}, args...)...)
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	fmt.Println("gen-android: fetching", repo.name)
	for _, args := range [][]string{
		{"init", "-q"},
		{"remote", "add", "origin", aosp + repo.name},
		{"sparse-checkout", "set", "--no-cone"},
		append([]string{"sparse-checkout", "add"}, repo.sparse...),
		{"fetch", "-q", "--depth", "1", "--filter=blob:none", "origin", repo.commit},
		{"checkout", "-q", "FETCH_HEAD"},
	} {
		if err := git(args...); err != nil {
			return err
		}
	}
	return nil
}

var include = regexp.MustCompile(`(?m)^\s*#\s*include(?:_next)?\s*[<"]([^>"]+)[>"]`)

// closure returns every header under dir, except for those within an
// optional top-level directory that no other header transitively includes.
func closure(dir string, optional map[string]bool) ([]string, error) {
	files := map[string]bool{}
	if err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		if base := filepath.Base(rel); strings.HasSuffix(base, ".h") || !strings.Contains(base, ".") {
			files[filepath.ToSlash(rel)] = true
		}
		return nil
	}); err != nil {
		return nil, err
	}
	var seen = map[string]bool{}
	var work []string
	for name := range files {
		top, rest, _ := strings.Cut(name, "/")
		if !optional[top] && !(top == "linux" && strings.Contains(rest, "/")) {
			seen[name] = true
			work = append(work, name)
		}
	}
	for len(work) > 0 {
		name := work[len(work)-1]
		work = work[:len(work)-1]
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return nil, err
		}
		for _, match := range include.FindAllSubmatch(data, -1) {
			included := string(match[1])
			candidates := []string{filepath.ToSlash(filepath.Join(filepath.Dir(name), included)), included}
			for _, arch := range architectures {
				candidates = append(candidates, arch.triple+"/"+included)
			}
			for _, candidate := range candidates {
				if files[candidate] && !seen[candidate] {
					seen[candidate] = true
					work = append(work, candidate)
				}
			}
		}
	}
	keep := make([]string, 0, len(seen))
	folded := map[string]string{}
	for name := range seen {
		if other, ok := folded[strings.ToLower(name)]; ok {
			return nil, fmt.Errorf("%s and %s only differ by case", name, other)
		}
		folded[strings.ToLower(name)] = name
		keep = append(keep, name)
	}
	slices.Sort(keep)
	return keep, nil
}

var symbol = regexp.MustCompile(`^\s*([A-Za-z_][A-Za-z0-9_]*);`)

// stub returns C source defining every symbol of the *.map.txt at path
// that is public and available to arch at [API].
func stub(path, arch string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var out strings.Builder
	fmt.Fprintf(&out, "// Code generated by gen-android from %s; DO NOT EDIT.\n", filepath.Base(path))
	out.WriteString("//\n// Link stub: the dynamic linker resolves these symbols to the device's real\n")
	out.WriteString("// library at runtime, the definitions here only exist to satisfy the linker.\n\n")
	var (
		seen    = map[string]bool{}
		private bool
		lines   []string
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code, comment, _ := strings.Cut(scanner.Text(), "#")
		tags := strings.Fields(comment)
		switch {
		case strings.Contains(code, "{"):
			private = strings.Contains(code, "PRIVATE") || strings.Contains(code, "PLATFORM") || slices.Contains(tags, "platform-only")
			continue
		case strings.Contains(code, "}"):
			private = false
			continue
		}
		match := symbol.FindStringSubmatch(code)
		if match == nil || private || seen[match[1]] {
			continue
		}
		introduced, restricted, supported, platform := 0, false, false, false
		for _, tag := range tags {
			key, value, _ := strings.Cut(tag, "=")
			switch key {
			case "introduced", "introduced-" + arch:
				if n, err := strconv.Atoi(value); err == nil && (introduced == 0 || n < introduced) {
					introduced = n
				}
			case "arm", "arm64", "x86", "x86_64", "riscv64":
				restricted = true
				supported = supported || key == arch
			case "platform-only", "apex", "llndk", "systemapi":
				platform = true
			}
		}
		if (restricted && !supported) || introduced > API || (platform && introduced == 0) {
			continue
		}
		seen[match[1]] = true
		if slices.Contains(tags, "var") {
			lines = append(lines, "int "+match[1]+" = 0;")
		} else {
			lines = append(lines, "void "+match[1]+"(void) {}")
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	out.WriteString(strings.Join(lines, "\n"))
	out.WriteString("\n")
	return out.String(), nil
}

func copyAll(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return copyFile(src, dst)
	}
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		return copyFile(path, filepath.Join(dst, rel))
	})
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}
