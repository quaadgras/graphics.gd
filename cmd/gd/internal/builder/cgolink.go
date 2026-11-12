package builder

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"graphics.gd/cmd/gd/internal/tooling"
)

// cgoLinkFlags returns the linker flags that go build would give the C code
// of the packages linked into the program in the current directory, for a
// link that gd does itself: `go build -buildmode=c-archive` records them in
// the archive but nothing applies them when another linker (zig, ld64)
// links it, so a package that links a C/C++ library would otherwise fail
// with undefined symbols. These are each package's #cgo LDFLAGS (go list
// expands their ${SRCDIR}) followed, as go build orders them, by the
// `pkg-config --libs` of its #cgo pkg-config packages, which go list leaves
// unresolved.
func cgoLinkFlags(tags string) ([]string, error) {
	out, err := tooling.Go.Output("list", "-tags", tags, "-deps", "-f",
		`{{range .CgoLDFLAGS}}{{println "L" .}}{{end}}{{if .CgoPkgConfig}}{{println "P" .Dir}}{{range .CgoPkgConfig}}{{println "A" .}}{{end}}{{end}}`, ".")
	if err != nil {
		return nil, err
	}
	var (
		flags []string
		dir   string
		args  []string
	)
	flush := func() error {
		if len(args) == 0 {
			return nil
		}
		libs, err := pkgConfigLibs(dir, args)
		args = nil
		flags = append(flags, libs...)
		return err
	}
	for _, line := range strings.Split(out, "\n") {
		kind, value, _ := strings.Cut(strings.TrimRight(line, "\r"), " ")
		if kind != "A" {
			if err := flush(); err != nil {
				return nil, err
			}
		}
		switch kind {
		case "L":
			if value = strings.TrimSpace(value); value != "" {
				flags = append(flags, value)
			}
		case "P":
			dir = value
		case "A":
			args = append(args, value)
		}
	}
	if err := flush(); err != nil {
		return nil, err
	}
	return flags, nil
}

// pkgConfigLibs runs pkg-config (or $PKG_CONFIG) for the arguments of a
// package's #cgo pkg-config directives, in its directory, the way go build
// does: those that start with -- are options, the rest packages.
func pkgConfigLibs(dir string, directive []string) ([]string, error) {
	pkgconfig := os.Getenv("PKG_CONFIG")
	if pkgconfig == "" {
		pkgconfig = "pkg-config"
	}
	var options, packages []string
	for _, arg := range directive {
		if strings.HasPrefix(arg, "--") {
			options = append(options, arg)
		} else {
			packages = append(packages, arg)
		}
	}
	if len(packages) == 0 {
		return nil, nil
	}
	cmd := exec.Command(pkgconfig, append(append(append([]string{"--libs"}, options...), "--"), packages...)...)
	cmd.Dir = dir
	var stderr strings.Builder
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return nil, fmt.Errorf("gd: %s needs %s to link %s: %w", dir, pkgconfig, strings.Join(packages, " "), err)
		}
		return nil, fmt.Errorf("gd: %s --libs %s: %w\n%s", pkgconfig, strings.Join(directive, " "), err, stderr.String())
	}
	return splitPkgConfigOutput(string(out))
}

// splitPkgConfigOutput splits pkg-config's output into its arguments, with
// pkg-config's quoting: backslash escapes, and single or double quotes.
func splitPkgConfigOutput(out string) ([]string, error) {
	var (
		args    []string
		arg     strings.Builder
		inArg   bool
		quote   rune
		escaped bool
	)
	for _, r := range out {
		switch {
		case escaped:
			escaped = false
			arg.WriteRune(r)
		case r == '\\' && quote != '\'':
			escaped, inArg = true, true
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
				arg.WriteRune(r)
			}
		case r == '\'' || r == '"':
			quote, inArg = r, true
		case strings.ContainsRune(" \t\r\n", r):
			if inArg {
				args = append(args, arg.String())
				arg.Reset()
				inArg = false
			}
		default:
			arg.WriteRune(r)
			inArg = true
		}
	}
	if quote != 0 || escaped {
		return nil, fmt.Errorf("gd: cannot parse the output of pkg-config: %q", out)
	}
	if inArg {
		args = append(args, arg.String())
	}
	return args, nil
}

// ld64Flags turns cgo linker flags, which are written for a C compiler that
// links (clang, zig cc), into the arguments of ld64.lld, which gd runs
// itself for Apple platforms: -Wl, and -Xlinker are unwrapped and flags
// that only mean something to the compiler (-pthread, -O2, -g, -m64,
// -fPIC...) are left out. The rest (-l, -L, -F, -framework and the linker's own
// options, files) are the linker's already.
func ld64Flags(flags []string) []string {
	var out []string
	for i := 0; i < len(flags); i++ {
		flag := flags[i]
		switch {
		case strings.HasPrefix(flag, "-Wl,"):
			out = append(out, strings.Split(strings.TrimPrefix(flag, "-Wl,"), ",")...)
		case flag == "-Xlinker":
			if i+1 < len(flags) {
				i++
				out = append(out, flags[i])
			}
		case flag == "-target" || flag == "-isysroot" || flag == "--sysroot" || flag == "-arch":
			i++ // a compiler flag with its value (the link has its own -arch).
		case compilerOnlyFlag(flag):
		default:
			out = append(out, flag)
		}
	}
	return out
}

// compilerOnlyFlag reports whether a cgo linker flag is the compiler's and
// not the linker's. Many of ld64's own options start with -f and -m too
// (-framework, -force_load, -fixup_chains, -filelist, -map), so only the
// compiler's spellings of those are matched.
func compilerOnlyFlag(flag string) bool {
	switch {
	case flag == "-pthread", flag == "-rdynamic", flag == "-static-libgcc", flag == "-static-libstdc++",
		flag == "-m32", flag == "-m64",
		strings.HasPrefix(flag, "-O"), strings.HasPrefix(flag, "-g"), strings.HasPrefix(flag, "-std="),
		strings.HasPrefix(flag, "--sysroot="), strings.HasPrefix(flag, "-isysroot"), strings.HasPrefix(flag, "--target="),
		strings.HasPrefix(flag, "-march="), strings.HasPrefix(flag, "-mcpu="), strings.HasPrefix(flag, "-mtune="),
		strings.HasPrefix(flag, "-m") && strings.Contains(flag, "-version-min="),
		strings.HasPrefix(flag, "-fuse-ld="), strings.HasPrefix(flag, "-flto"), strings.HasPrefix(flag, "-fno-"),
		strings.HasPrefix(flag, "-fsanitize"), strings.HasPrefix(flag, "-fprofile"), strings.HasPrefix(flag, "-fstack-"),
		strings.HasPrefix(flag, "-fobjc-"), strings.HasPrefix(flag, "-fopenmp"),
		flag == "-fPIC", flag == "-fpic", flag == "-fPIE", flag == "-fpie":
		return true
	}
	return false
}
