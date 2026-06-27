package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

type Linux struct{}

func (Linux) Build(args ...string) error {
	if !project.IncludesGo {
		return nil
	}
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	var glibc bool
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
		zig, err := tooling.Zig.Lookup()
		if err != nil {
			return xray.New(err)
		}
		switch GOARCH {
		case "amd64":
			if err := os.Setenv("CC", zig+" cc -target x86_64-linux-gnu"); err != nil {
				return xray.New(err)
			}
		case "arm64":
			if err := os.Setenv("CC", zig+" cc -target aarch64-linux-gnu"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
		}
		glibc = true // cross-compiles always target *-linux-gnu above
	} else {
		// Native build: the target libc is the host libc. ldd reports "musl ..." on
		// musl systems and "ldd (GNU libc) ..." on glibc ones.
		version, _ := tooling.ListDynamicDependencies.CombinedOutput("--version")
		glibc = !strings.HasPrefix(strings.TrimSpace(version), "musl")
	}
	if glibc {
		// The c-shared library references the libgcc unwinder (_Unwind_GetIP and the rest of
		// the _Unwind_* family, pulled in by cgo / Go's C objects), but the default external
		// link can record no DT_NEEDED for libgcc_s.so.1 — it then relies on those symbols
		// already being in the global scope when the host engine dlopen()s the extension. That
		// holds where libgcc_s.so.1 is loaded eagerly, but on glibc >= 2.34 (Fedora/Nobara,
		// etc.) it is only loaded lazily on demand, so the extension fails to relocate with
		// "undefined symbol: _Unwind_GetIP". Force a recorded dependency so the loader pulls
		// libgcc_s.so.1 in for us: --no-as-needed defeats the linker's as-needed pruning (which
		// drops the library when no pending undefined reference is visible as it is processed),
		// then --as-needed restores the default for the libraries that follow. (musl builds
		// link the unwinder in statically and must not get -lgcc_s; this is gated to glibc.
		// The zig toolchain links its own unwinder in too, so there the flags are a harmless
		// no-op — but it rejects -Wl,--push-state, so we use the as-needed pair, not push/pop.)
		const forceUnwinder = "-Wl,--no-as-needed -lgcc_s -Wl,--as-needed"
		ldflags := strings.TrimSpace(os.Getenv("CGO_LDFLAGS") + " " + forceUnwinder)
		if err := os.Setenv("CGO_LDFLAGS", ldflags); err != nil {
			return xray.New(err)
		}
	}
	return tooling.Go.Action("build", args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("linux_%v.so", GOARCH)))
}

func (linux Linux) BuildMain(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if err := linux.Build(args...); err != nil {
		return xray.New(err)
	}
	var export []string
	switch GOARCH {
	case "amd64":
		export = []string{"--headless", "--export-release", "Linux x86_64"}
	case "arm64":
		export = []string{"--headless", "--export-release", "Linux arm64"}
	default:
		return fmt.Errorf("gd export: cannot export linux %v", GOARCH)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec(export...); err != nil {
		return xray.New(err)
	}
	return nil
}

func (linux Linux) Run(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd run: cannot run linux/%v executable on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := linux.Build(args...); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return tooling.Godot.Exec(args...)
}

func (Linux) Test(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd test: cannot run linux/%v tests on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := tooling.Go.Action("test", args, "-c", "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("linux_%v.so", GOARCH))); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	args = append(args, "--headless")
	return tooling.Godot.Exec(args...)
}
