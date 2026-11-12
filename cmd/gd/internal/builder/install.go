package builder

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"
)

// CanInstall reports why `gd install` cannot install a goos/goarch release
// on this system, or nil when it can: only a release that runs here is
// installed (GOOS=musl included, on linux, and android on-device, under
// Termux, where Android.InstallOnDevice does the installing).
func CanInstall(goos, goarch string) error {
	host := runtime.GOOS
	if goos == "musl" {
		goos = "linux"
	}
	if goos != host || (goarch != runtime.GOARCH && !(goos == "darwin" && os.Getenv("GOARCH") == "")) {
		return fmt.Errorf("gd install: cannot install a %v/%v release on %v/%v, use gd build to export it instead", goos, goarch, runtime.GOOS, runtime.GOARCH)
	}
	switch goos {
	case "linux", "darwin", "windows", "android":
		return nil
	}
	return fmt.Errorf("gd install: installing is not supported on %v, use gd build to export a release instead", host)
}

// Install installs the release that BuildMain exported for goos/goarch to
// release (see ReleasePath) on this system: on linux into GOBIN (GOPATH/bin), on macOS into
// /Applications (~/Applications when that is not writable) and on windows
// into %LOCALAPPDATA%\Programs, with a Start menu shortcut. Either way GOBIN
// gets a command that runs it, so it is on the PATH like go install's.
func Install(goos, release string) error {
	if _, err := os.Stat(release); err != nil {
		return fmt.Errorf("gd install: the release build did not produce %s", release)
	}
	switch goos {
	case "linux", "musl":
		return installLinux(release)
	case "darwin":
		return installMacOS(release)
	case "windows":
		return installWindows(release)
	}
	return fmt.Errorf("gd install: cannot install a %v release", goos)
}

// ReleasePath returns where BuildMain exports goos/goarch to. It is to be
// asked before the build, which sets GOARCH for the architectures it builds.
func ReleasePath(goos, goarch string) (string, error) {
	var arch string
	switch goarch {
	case "amd64":
		arch = "x86_64"
	case "arm64":
		arch = "arm64"
	default:
		return "", fmt.Errorf("gd install: cannot install %v/%v", goos, goarch)
	}
	switch goos {
	case "linux", "musl":
		if goos == "musl" || BuildMode() == "libgodot" {
			return filepath.Join(project.ReleasesDirectory, "musl", goarch, project.Name), nil // see Musl.BuildMain
		}
		return presetExportPath("Linux " + arch)
	case "windows":
		return presetExportPath("Windows " + arch)
	case "darwin":
		if BuildMode() == "libgodot" {
			variant := "universal" // see MacOS.buildMainLibgodot
			if GOARCH := os.Getenv("GOARCH"); GOARCH != "" {
				variant = GOARCH
			}
			return filepath.Join(project.ReleasesDirectory, "darwin", variant, project.AppleSafePackageName(project.Name)+".app"), nil
		}
		return presetExportPath("macOS")
	}
	return "", fmt.Errorf("gd install: cannot install a %v release", goos)
}

// installLinux copies an executable that is all of its release into GOBIN.
// A release with files beside the executable (the extension's library, in
// the default build mode) is copied into GOPATH/pkg/gd/<name> as a whole,
// and GOBIN gets a symlink to its executable: the engine finds those files
// beside the executable it resolves /proc/self/exe to.
func installLinux(release string) error {
	gobin, gopath, err := goPaths()
	if err != nil {
		return err
	}
	name := filepath.Base(release)
	dst := filepath.Join(gobin, name)
	entries, err := os.ReadDir(filepath.Dir(release))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(gobin, 0755); err != nil {
		return err
	}
	if len(entries) == 1 {
		os.Remove(dst) // a symlink from an earlier install of the other kind.
		if err := copyFile(release, dst); err != nil {
			return err
		}
		fmt.Println("gd: installed", dst)
		return nil
	}
	dir := filepath.Join(gopath, "pkg", "gd", name)
	if err := replaceDir(filepath.Dir(release), dir); err != nil {
		return err
	}
	os.Remove(dst)
	if err := os.Symlink(filepath.Join(dir, name), dst); err != nil {
		return err
	}
	fmt.Println("gd: installed", dir)
	fmt.Println("gd: installed", dst, "->", filepath.Join(dir, name))
	return nil
}

// goPaths returns where go install puts executables, and the first entry
// of GOPATH.
func goPaths() (gobin, gopath string, err error) {
	out, err := tooling.Go.Output("env", "GOBIN", "GOPATH")
	if err != nil {
		return "", "", err
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		return "", "", fmt.Errorf("gd install: unexpected output from go env: %q", out)
	}
	gobin, gopath = strings.TrimSpace(lines[0]), filepath.SplitList(strings.TrimSpace(lines[1]))[0]
	if gopath == "" {
		return "", "", errors.New("gd install: GOPATH is not set")
	}
	if gobin == "" {
		gobin = filepath.Join(gopath, "bin")
	}
	return gobin, gopath, nil
}

// installMacOS copies the app bundle into /Applications, or into the
// user's own Applications folder when /Applications is not writable, and
// puts a GOBIN/<name> script that execs its executable. Not a symbolic
// link: the engine looks for the bundle's data beside the path it was run
// as (_NSGetExecutablePath does not resolve links).
func installMacOS(app string) error {
	gobin, _, err := goPaths()
	if err != nil {
		return err
	}
	dir := "/Applications"
	if probe, err := os.CreateTemp(dir, ".gd-install-"); err == nil {
		probe.Close()
		os.Remove(probe.Name())
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		dir = filepath.Join(home, "Applications")
	}
	dst := filepath.Join(dir, filepath.Base(app))
	if err := replaceDir(app, dst); err != nil {
		return err
	}
	fmt.Println("gd: installed", dst)
	exe, err := bundleExecutable(dst)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(gobin, 0755); err != nil {
		return err
	}
	script := filepath.Join(gobin, strings.TrimSuffix(filepath.Base(app), ".app"))
	os.Remove(script)
	if err := os.WriteFile(script, []byte("#!/bin/sh\nexec '"+strings.ReplaceAll(exe, "'", `'\''`)+"' \"$@\"\n"), 0755); err != nil {
		return err
	}
	fmt.Println("gd: installed", script)
	return nil
}

// bundleExecutable returns the executable of the app bundle, the one that
// its Info.plist names (gd and Godot both write the XML form).
func bundleExecutable(app string) (string, error) {
	plist, err := os.ReadFile(filepath.Join(app, "Contents", "Info.plist"))
	if err == nil {
		if m := regexp.MustCompile(`<key>CFBundleExecutable</key>\s*<string>([^<]+)</string>`).FindSubmatch(plist); m != nil {
			exe := filepath.Join(app, "Contents", "MacOS", string(m[1]))
			if _, err := os.Stat(exe); err == nil {
				return exe, nil
			}
		}
	}
	entries, err := os.ReadDir(filepath.Join(app, "Contents", "MacOS"))
	if err != nil {
		return "", err
	}
	if len(entries) != 1 {
		return "", fmt.Errorf("gd install: cannot tell which executable of %s to put on the PATH", app)
	}
	return filepath.Join(app, "Contents", "MacOS", entries[0].Name()), nil
}

// installWindows copies the release (the executable and the libraries
// beside it) into %LOCALAPPDATA%\Programs\<name>, puts a shortcut to it in
// the Start menu and a GOBIN\<name>.cmd that runs it. Not a symbolic link:
// those need developer mode, and the engine would look for the libraries
// beside the link.
func installWindows(exe string) error {
	gobin, _, err := goPaths()
	if err != nil {
		return err
	}
	local := os.Getenv("LOCALAPPDATA")
	if local == "" {
		return errors.New("gd install: LOCALAPPDATA is not set")
	}
	name := strings.TrimSuffix(filepath.Base(exe), filepath.Ext(exe))
	dir := filepath.Join(local, "Programs", name)
	if err := replaceDir(filepath.Dir(exe), dir); err != nil {
		return err
	}
	fmt.Println("gd: installed", dir)
	if err := os.MkdirAll(gobin, 0755); err != nil {
		return err
	}
	cmd := filepath.Join(gobin, name+".cmd")
	if err := os.WriteFile(cmd, []byte("@\""+filepath.Join(dir, filepath.Base(exe))+"\" %*\r\n"), 0755); err != nil {
		return err
	}
	fmt.Println("gd: installed", cmd)
	if appdata := os.Getenv("APPDATA"); appdata != "" {
		lnk := filepath.Join(appdata, "Microsoft", "Windows", "Start Menu", "Programs", name+".lnk")
		quote := func(s string) string { return "'" + strings.ReplaceAll(s, "'", "''") + "'" }
		script := fmt.Sprintf("$s = (New-Object -ComObject WScript.Shell).CreateShortcut(%s); $s.TargetPath = %s; $s.WorkingDirectory = %s; $s.Save()",
			quote(lnk), quote(filepath.Join(dir, filepath.Base(exe))), quote(dir))
		if out, err := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", script).CombinedOutput(); err != nil {
			fmt.Fprintf(os.Stderr, "gd: could not add a Start menu shortcut: %v\n%s", err, out)
		} else {
			fmt.Println("gd: installed", lnk)
		}
	}
	return nil
}

// replaceDir replaces dst with a copy of the directory src.
func replaceDir(src, dst string) error {
	tmp := dst + ".gd-install"
	if err := os.RemoveAll(tmp); err != nil {
		return err
	}
	if err := copyTree(src, tmp); err != nil {
		os.RemoveAll(tmp)
		return err
	}
	if err := os.RemoveAll(dst); err != nil {
		return err
	}
	return os.Rename(tmp, dst)
}

// copyTree copies the directory src to dst, keeping file modes and
// symbolic links (app bundles have both).
func copyTree(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		switch {
		case info.Mode()&os.ModeSymlink != 0:
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			return os.Symlink(link, target)
		case info.IsDir():
			return os.MkdirAll(target, info.Mode().Perm()|0700)
		default:
			in, err := os.Open(path)
			if err != nil {
				return err
			}
			defer in.Close()
			out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode().Perm())
			if err != nil {
				return err
			}
			if _, err := io.Copy(out, in); err != nil {
				out.Close()
				return err
			}
			return out.Close()
		}
	})
}
