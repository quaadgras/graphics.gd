package tooling

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/schollz/progressbar/v3"
	"runtime.link/api/xray"
)

var debug = os.Getenv("DEBUG_CMD") != ""

// GOTOOLCHAIN=local will disable automatic toolchain downloads.

type toolchain struct {
	Name          string                       // as found in $PATH
	Version       string                       // expected version
	VersionFlag   string                       // to extract version
	VersionPrefix string                       // version prefix
	Downloads     map[string]map[string]string // specific GOOS/GOARCH download URLs
	DownloadURL   string                       // base Download URL with $(VERSION), $(OS), $(ARCH) variables
	DownloadARCH  map[string]string            // to map GOARCH to $(ARCH)
	DownloadOS    map[string]string            // to map GOOS to $(OS)
	DownloadEXT   map[string]string            // to map GOOS to download file extension.
	DownloadHint  string                       // where to get it
	Unzip         string                       // rename the binary named this inside the zip to Name
	IsApp         bool
	Installations map[string]string // expected installations for specific GOOS (with $(HOME) variable)

	RequiredFor string // description of why gd needs this toolchain dependency

	ConvertArguments map[string]string

	DarwinUniversal bool

	IsLibrary bool

	Path string // cached by [toolchain.Lookup]
}

func (exe toolchain) PathToCommand() string {
	if exe.Path == "" {
		panic("toolchain.PathToCommand: toolchain not yet looked up")
	}
	if exe.IsApp && runtime.GOOS == "darwin" {
		return filepath.Join(exe.Path, "Contents", "MacOS", exe.Name)
	}
	return exe.Path
}

func (exe toolchain) Exec(args ...string) error {
	for i, arg := range args {
		if strings.Contains(arg, "=") {
			arg, _, _ = strings.Cut(arg, "=")
		}
		if newarg, ok := exe.ConvertArguments[arg]; ok {
			if newarg == "" {
				args = append(args[:i], args[i+1:]...)
				continue
			}
			args[i] = newarg
		}
	}
	path, err := exe.Lookup()
	if err != nil {
		return xray.New(err)
	}
	cmd := exec.Command(path, args...)
	if debug {
		fmt.Println(path, strings.Join(args, " "))
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func (exe toolchain) Action(name string, suffix_args []string, args ...string) error {
	var suffix = make([]string, 0, len(suffix_args))
	for _, arg := range suffix_args {
		suffix = append(suffix, arg)
	}
	for i, arg := range args {
		if newarg, ok := exe.ConvertArguments[arg]; ok {
			args[i] = newarg
		}
	}
	for i, arg := range suffix {
		if newarg, ok := exe.ConvertArguments[arg]; ok {
			suffix[i] = newarg
		}
	}
	path, err := exe.Lookup()
	if err != nil {
		return xray.New(err)
	}
	args = append(append([]string{name}, args...), suffix...)
	cmd := exec.Command(path, args...)
	if debug {
		fmt.Println(path, strings.Join(args, " "))
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func (exe toolchain) Output(args ...string) (string, error) {
	path, err := exe.Lookup()
	if err != nil {
		return "", err
	}
	if debug {
		fmt.Println(path, strings.Join(args, " "))
	}
	out, err := exec.Command(path, args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func (exe toolchain) CombinedOutput(args ...string) (string, error) {
	path, err := exe.Lookup()
	if err != nil {
		return "", err
	}
	if debug {
		fmt.Println(path, strings.Join(args, " "))
	}
	out, err := exec.Command(path, args...).CombinedOutput()
	if debug {
		fmt.Println(string(out))
	}
	return strings.TrimSpace(string(out)), err
}

func (exe *toolchain) Lookup() (string, error) {
	return exe.LookupPlatform(runtime.GOOS, runtime.GOARCH)
}

func (exe *toolchain) LookupPlatform(GOOS, GOARCH string) (string, error) {
	if exe.Path != "" {
		return exe.Path, nil
	}
	my, err := user.Current()
	if err != nil {
		return "", xray.New(err)
	}
	HOME := my.HomeDir
	GDPATH := os.Getenv("GDPATH")
	if GDPATH == "" && HOME != "" {
		GDPATH = filepath.Join(HOME, "gd")
	}
	ARCH := exe.DownloadARCH[GOARCH]
	if ARCH == "" {
		ARCH = "$(MISSING)"
	}
	OS := strings.ReplaceAll(exe.DownloadOS[GOOS], "$(ARCH)", ARCH)
	if OS == "" {
		OS = "$(MISSING)"
	}
	EXT, ok := exe.DownloadEXT[GOOS]
	if !ok {
		EXT = "$(MISSING)"
	}
	var MaybeUniversal = GOARCH
	if GOOS == "darwin" && exe.DarwinUniversal {
		MaybeUniversal = "universal"
	}
	var variables = strings.NewReplacer(
		"$(VERSION)", exe.Version, "$(ARCH)", ARCH, "$(OS)", OS, "$(GOARCH)", MaybeUniversal, "$(GOOS)", GOOS, "$(HOME)", HOME, "$(GDPATH)", GDPATH, "$(EXT)", EXT,
	)
	var install_dir = filepath.Join(GDPATH, "bin")
	if exe.IsLibrary {
		install_dir = filepath.Join(GDPATH, "lib")
	}
	if dir, ok := exe.Installations[GOOS]; ok {
		install_dir = variables.Replace(dir)
	}
	var name = variables.Replace(exe.Name)
	var install_path = filepath.Join(install_dir, name)
	if runtime.GOOS == "windows" {
		install_path += ".exe"
	}
	if exe.IsApp && runtime.GOOS == "darwin" {
		install_path += ".app"
	}
	// always prefer the GDPATH-installed version if it matches the expected version.
	if _, err := os.Stat(install_path); err == nil {
		if exe.IsLibrary {
			return install_path, nil
		}
		var exe_path = install_path
		if exe.IsApp && runtime.GOOS == "darwin" {
			exe_path = filepath.Join(install_path, "Contents", "MacOS", name)
		}
		version, err := exec.Command(exe_path, exe.VersionFlag).CombinedOutput()
		version = bytes.TrimSpace(version)
		if err == nil {
			if (exe.Version != "" && string(version) == exe.Version) || (exe.VersionPrefix != "" && strings.HasPrefix(string(version), exe.VersionPrefix)) {
				exe.Path = install_path
				return exe.PathToCommand(), nil
			}
		}
	}
	// some users (ie. NixOS) don't want things to be automatically installed, they
	// can set their toolchain to local and download/install everything themselves.
	if os.Getenv("GOTOOLCHAIN") == "local" || os.Getenv("GDTOOLCHAIN") == "local" || GDPATH == "" {
		path, err := exec.LookPath(name)
		if err != nil {
			return "", fmt.Errorf(
				"'%v' %s not found in $PATH (required for %v) and automatic-downloads are disabled, please install it, ie. %v",
				name, exe.Version, exe.RequiredFor, exe.DownloadHint,
			)
		}
		exe.Path = path
		return exe.PathToCommand(), nil
	}
	if !exe.IsLibrary {
		// if the expected version of the tool is already installed in $PATH, then we can
		// just use it.
		if path, err := exec.LookPath(name); err == nil {
			version, _ := exec.Command(path, exe.VersionFlag).CombinedOutput()
			if (exe.Version != "" && string(version) == exe.Version) || (exe.VersionPrefix != "" && strings.HasPrefix(string(version), exe.VersionPrefix)) || (exe.Version == "" && exe.VersionPrefix == "") {
				exe.Path = path
				if exe.IsApp {
					exe.IsApp = false
				}
				return exe.PathToCommand(), nil
			}
		}
	}
	// attempt to automatically download and install the toolchain.
	url, ok := exe.Downloads[GOOS][GOARCH]
	if !ok {
		url = variables.Replace(exe.DownloadURL)
	}
	if url == "" || strings.Contains(url, "$(MISSING)") {
		return "", fmt.Errorf(
			"'%v' %s not found in $PATH (required for %v) and no automatic-download is available, please install it, ie. %v",
			name, exe.Version, exe.RequiredFor, exe.DownloadHint,
		)
	}
	if err := os.MkdirAll(install_dir, 0755); err != nil {
		return "", xray.New(err)
	}
	var dest = install_path
	dest += "." + exe.Version + ".download"
	if err := func() error {
		out, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return xray.New(err)
		}
		defer out.Close()
		stat, err := out.Stat()
		if err != nil {
			return xray.New(err)
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return xray.New(err)
		}
		if stat.Size() > 0 {
			req.Header.Set("Range", fmt.Sprintf("bytes=%d-", stat.Size()))
		}
		req.Header.Set("User-Agent", "graphics.gd/cmd/gd")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return xray.New(err)
		}
		defer resp.Body.Close()
		switch resp.StatusCode {
		case 200:
		case 206:
			if _, err := out.Seek(stat.Size(), io.SeekStart); err != nil {
				return xray.New(err)
			}
		case 416:
			contentRange := resp.Header.Get("Content-Range")
			if contentRange != fmt.Sprintf("bytes */%d", stat.Size()) {
				return fmt.Errorf("unable to resume download of '%v' (required for %v), please delete %v and try again\nGET %s HTTP status: %v", name, exe.RequiredFor, dest, url, resp.StatusCode)
			}
		default:
			return fmt.Errorf(
				"unable to download '%v' (required for %v) and not found in $PATH, please install it, ie. %v\nGET %s HTTP status: %v",
				name, exe.RequiredFor, exe.DownloadHint, url, resp.StatusCode,
			)
		}
		if resp.StatusCode != 416 {
			bar := progressbar.DefaultBytes(
				resp.ContentLength,
				fmt.Sprintf("gd: downloading %s v%s", name, exe.Version),
			)
			if _, err := io.Copy(io.MultiWriter(out, bar), resp.Body); err != nil {
				return xray.New(err)
			}
		}
		return nil
	}(); err != nil {
		return "", xray.New(err)
	}
	var unzip = variables.Replace(exe.Unzip)
	if exe.IsApp && runtime.GOOS == "darwin" {
		unzip = ""
	}
	switch {
	case strings.HasSuffix(url, ".zip"):
		if err := ExtractArchive(dest, install_dir, "zip", unzip, runtime.GOOS != "darwin" || !exe.IsApp); err != nil {
			return "", xray.New(err)
		}
		if err := os.Remove(dest); err != nil {
			return "", xray.New(err)
		}
	case strings.HasSuffix(url, ".tar.gz"):
		if err := ExtractArchive(dest, install_dir, "tar.gz", unzip, true); err != nil {
			return "", xray.New(err)
		}
		if err := os.Remove(dest); err != nil {
			return "", xray.New(err)
		}
	case strings.HasSuffix(url, ".tar.xz"):
		if err := ExtractArchive(dest, install_dir, "tar.xz", unzip, true); err != nil {
			return "", xray.New(err)
		}
		if err := os.Remove(dest); err != nil {
			return "", xray.New(err)
		}
	default:
		if err := os.Rename(dest, install_path); err != nil {
			return "", xray.New(err)
		}
	}
	if unzip != "" {
		if err := os.Rename(filepath.Join(install_dir, unzip), install_path); err != nil {
			return "", xray.New(err)
		}
	}
	exe.Path = install_path
	return exe.PathToCommand(), nil
}
