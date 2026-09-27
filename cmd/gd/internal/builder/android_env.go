package builder

import (
	"encoding/base64"
	"os"
	"slices"
	"sort"
	"strings"
)

// bakedEnvFlags returns the linker flag that bakes the shell's own
// environment variables into the library, for internal/bakedenv to set as
// the app starts. On-device (Termux) the app is launched by the system and
// inherits nothing from the shell that ran gd, and Godot's exported
// launcher drops any command line handed to it, so `NAME=value gd run` would
// otherwise have no way to reach the program. What the system, Termux, the
// shell and the Go toolchain set is left out (it would be wrong inside the
// app), as is anything named like a credential, which should not end up in
// an APK.
func bakedEnvFlags() string {
	var kept []string
	for _, kv := range os.Environ() {
		name, _, _ := strings.Cut(kv, "=")
		if name == "" || !bakeable(name) {
			continue
		}
		kept = append(kept, kv)
	}
	if len(kept) == 0 {
		return ""
	}
	sort.Strings(kept)
	return " -X graphics.gd/internal/bakedenv.env=" + base64.StdEncoding.EncodeToString([]byte(strings.Join(kept, "\x00")))
}

// bakeable reports whether a variable is the user's own, to bake into the
// library, rather than one the system, Termux, the shell or the toolchain
// set.
func bakeable(name string) bool {
	if slices.Contains([]string{
		"PATH", "HOME", "PREFIX", "TMPDIR", "SHELL", "TERM", "COLORTERM", "LANG",
		"PWD", "OLDPWD", "SHLVL", "_", "USER", "LOGNAME", "HOSTNAME", "MAIL",
		"EDITOR", "VISUAL", "PAGER", "BROWSER", "TZ", "HISTFILE", "HISTSIZE",
		"EXTERNAL_STORAGE", "SECONDARY_STORAGE", "ASEC_MOUNTPOINT", "BOOTCLASSPATH",
		"DEX2OATBOOTCLASSPATH", "SYSTEMSERVERCLASSPATH", "STANDALONE_SYSTEMSERVER_JARS",
		"JAVA_HOME", "CC", "CXX", "AR", "CFLAGS", "CXXFLAGS", "LDFLAGS", "CPPFLAGS",
		"PKG_CONFIG_PATH", "ZIG_GLOBAL_CACHE_DIR", "ZIG_LOCAL_CACHE_DIR",
	}, name) {
		return false
	}
	for _, prefix := range []string{
		"ANDROID_", "TERMUX", "LD_", "LC_", "GO", "CGO_", "GD", "GODOT", "SSH_",
		"XDG_", "TMUX", "STY", "DBUS_", "DISPLAY", "WAYLAND_", "NPM_", "NODE_",
		"PYTHON", "VIRTUAL_ENV", "CONDA",
	} {
		if strings.HasPrefix(name, prefix) {
			return false
		}
	}
	upper := strings.ToUpper(name)
	for _, secret := range []string{"TOKEN", "SECRET", "PASSWORD", "PASSWD", "API_KEY", "APIKEY", "PRIVATE", "CREDENTIAL", "AUTH"} {
		if strings.Contains(upper, secret) {
			return false
		}
	}
	return true
}
