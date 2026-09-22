package builder

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"
)

// androidAPI is the API level the bundled link stubs were generated for
// (see bundled/gen-android), which is also Godot's minimum.
const androidAPI = "24"

// androidTriple returns the target triple for an android GOARCH.
func androidTriple(GOARCH string) (string, error) {
	switch GOARCH {
	case "arm64":
		return "aarch64-linux-android", nil
	case "amd64":
		return "x86_64-linux-android", nil
	}
	return "", fmt.Errorf("gd: unsupported android architecture %q", GOARCH)
}

// androidSDK describes the bundled Android SDK, materialized for a triple.
type androidSDK struct {
	Include     string // architecture independent headers.
	IncludeArch string // headers specific to the triple (asm).
	Lib         string // link stubs and C runtime objects for the triple.
	LibC        string // zig libc installation file for the triple, see ZIG_LIBC.
}

// setupAndroidSDK writes the bundled Android SDK to dir and compiles the
// link stubs plus C runtime objects for the given triple. The stubs only
// satisfy the linker, on-device the dynamic linker resolves their symbols
// to the real system libraries.
func setupAndroidSDK(dir, triple string) (androidSDK, error) {
	sdk := androidSDK{
		Include:     filepath.Join(dir, "usr", "include"),
		IncludeArch: filepath.Join(dir, "usr", "include", triple),
		Lib:         filepath.Join(dir, "usr", "lib", triple),
	}
	sdk.LibC = filepath.Join(sdk.Lib, "libc.txt")
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return sdk, err
	}
	// The stubs take a moment to compile, so they are only rebuilt when the
	// bundled SDK (or where it lives, libc.txt is absolute) changes.
	hash := sha256.New()
	fmt.Fprintln(hash, dir, zig)
	err = fs.WalkDir(android_sdk, "bundled/android", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := android_sdk.ReadFile(path)
		fmt.Fprintln(hash, path, len(data))
		hash.Write(data)
		return err
	})
	if err != nil {
		return sdk, err
	}
	stamp, sum := filepath.Join(sdk.Lib, "stamp"), hex.EncodeToString(hash.Sum(nil))
	if existing, err := os.ReadFile(stamp); err == nil && string(existing) == sum {
		return sdk, nil
	}
	if err := project.SetupFiles(android_sdk, "bundled/android", dir); err != nil {
		return sdk, err
	}
	sources, err := filepath.Glob(filepath.Join(sdk.Lib, "*.c"))
	if err != nil {
		return sdk, err
	}
	for _, source := range sources {
		name := strings.TrimSuffix(filepath.Base(source), ".c")
		// No ZIG_LIBC and no API level here: with either, zig links libc
		// into the stub, which is what is being built.
		cmd := exec.Command(zig, append([]string{"cc", "-target", triple, "-shared", "-nostdlib", "-fno-builtin", "-w"}, append(tooling.CGOCFlags(),
			"-Wl,-soname,"+name+".so", "-o", filepath.Join(sdk.Lib, name+".so"), source)...)...)
		if output, err := cmd.CombinedOutput(); err != nil {
			return sdk, fmt.Errorf("build %s stub for %s: %w\n%s", name, triple, err, output)
		}
	}
	libc := fmt.Sprintf("include_dir=%s\nsys_include_dir=%s\ncrt_dir=%s\nmsvc_lib_dir=\nkernel32_lib_dir=\ngcc_dir=\n",
		sdk.Include, sdk.IncludeArch, sdk.Lib)
	if err := os.WriteFile(sdk.LibC, []byte(libc), 0644); err != nil {
		return sdk, err
	}
	for _, crt := range []string{"crtbegin_so", "crtend_so"} {
		cmd := exec.Command(zig, "cc", "-target", triple+"."+androidAPI, "-fPIC", "-O2", "-fvisibility=hidden",
			"-c", filepath.Join(dir, "usr", "lib", crt+".c"), "-o", filepath.Join(sdk.Lib, crt+".o"))
		cmd.Env = append(os.Environ(), "ZIG_LIBC="+sdk.LibC)
		if output, err := cmd.CombinedOutput(); err != nil {
			return sdk, fmt.Errorf("build %s for %s: %w\n%s", crt, triple, err, output)
		}
	}
	return sdk, os.WriteFile(stamp, []byte(sum), 0644)
}
