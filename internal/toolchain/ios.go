//go:build ios

package toolchain

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// iOS forbids process creation, so Shell here is a small built-in
// userland (pure Go, runs in-process) rather than /bin/sh: enough for
// the agent to look around and manage files. No pipes, no quoting —
// one command per line, whitespace-separated. Compilation is the next
// tier: compiler.gd for Go and lld's library entry point for the
// Mach-O link, per the Readme roadmap.

func (t *Toolchain) GD(verb string, goos string) error {
	return errors.New("building on iOS needs the in-process toolchain, which is not implemented yet (see Readme roadmap)")
}

func (t *Toolchain) Shell(command string, timeout time.Duration) (string, error) {
	args := strings.Fields(command)
	if len(args) == 0 {
		return "", nil
	}
	name, args := args[0], args[1:]
	builtin, ok := builtins[name]
	if !ok {
		known := make([]string, 0, len(builtins))
		for k := range builtins {
			known = append(known, k)
		}
		return "", fmt.Errorf("no such builtin %q (no exec on iOS; have: %s)", name, strings.Join(known, " "))
	}
	return builtin(t.Project, args)
}

var builtins = map[string]func(root string, args []string) (string, error){
	"pwd":   func(root string, _ []string) (string, error) { return root, nil },
	"echo":  func(_ string, args []string) (string, error) { return strings.Join(args, " "), nil },
	"ls":    builtinLs,
	"cat":   builtinCat,
	"grep":  builtinGrep,
	"find":  builtinFind,
	"mkdir": builtinMkdir,
	"rm":    builtinRm,
	"mv":    builtinMv,
	"cp":    builtinCp,
}

func within(root string, path string) (string, error) {
	if path == "" {
		path = "."
	}
	joined := filepath.Clean(filepath.Join(root, path))
	if joined != root && !strings.HasPrefix(joined, root+string(filepath.Separator)) {
		return "", fmt.Errorf("path escapes the project: %s", path)
	}
	return joined, nil
}

func builtinLs(root string, args []string) (string, error) {
	path := "."
	if len(args) > 0 {
		path = args[len(args)-1]
	}
	dir, err := within(root, path)
	if err != nil {
		return "", err
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	for _, e := range entries {
		if e.IsDir() {
			b.WriteString(e.Name() + "/\n")
		} else {
			b.WriteString(e.Name() + "\n")
		}
	}
	return b.String(), nil
}

func builtinCat(root string, args []string) (string, error) {
	var b strings.Builder
	for _, arg := range args {
		path, err := within(root, arg)
		if err != nil {
			return "", err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		b.Write(data)
	}
	return b.String(), nil
}

func builtinGrep(root string, args []string) (string, error) {
	recursive := len(args) > 0 && (args[0] == "-r" || args[0] == "-R")
	if recursive {
		args = args[1:]
	}
	if len(args) < 1 {
		return "", errors.New("usage: grep [-r] substring [path]")
	}
	pattern := args[0]
	start := "."
	if len(args) > 1 {
		start = args[1]
	}
	path, err := within(root, start)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	match := func(file string) {
		data, err := os.ReadFile(file)
		if err != nil {
			return
		}
		rel, _ := filepath.Rel(root, file)
		for i, line := range strings.Split(string(data), "\n") {
			if strings.Contains(line, pattern) {
				fmt.Fprintf(&b, "%s:%d:%s\n", rel, i+1, line)
			}
		}
	}
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		match(path)
	} else if recursive {
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err == nil && !d.IsDir() {
				match(p)
			}
			return nil
		})
	} else {
		return "", errors.New("grep on a directory needs -r")
	}
	return b.String(), nil
}

func builtinFind(root string, args []string) (string, error) {
	start := "."
	pattern := ""
	for i := 0; i < len(args); i++ {
		if args[i] == "-name" && i+1 < len(args) {
			pattern = args[i+1]
			i++
		} else {
			start = args[i]
		}
	}
	path, err := within(root, start)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if pattern != "" {
			if ok, _ := filepath.Match(pattern, d.Name()); !ok {
				return nil
			}
		}
		rel, _ := filepath.Rel(root, p)
		b.WriteString(rel + "\n")
		return nil
	})
	return b.String(), nil
}

func builtinMkdir(root string, args []string) (string, error) {
	for _, arg := range args {
		if arg == "-p" {
			continue
		}
		path, err := within(root, arg)
		if err != nil {
			return "", err
		}
		if err := os.MkdirAll(path, 0o755); err != nil {
			return "", err
		}
	}
	return "", nil
}

func builtinRm(root string, args []string) (string, error) {
	recursive := len(args) > 0 && (args[0] == "-r" || args[0] == "-rf")
	if recursive {
		args = args[1:]
	}
	for _, arg := range args {
		path, err := within(root, arg)
		if err != nil {
			return "", err
		}
		if recursive {
			err = os.RemoveAll(path)
		} else {
			err = os.Remove(path)
		}
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func builtinMv(root string, args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("usage: mv src dst")
	}
	src, err := within(root, args[0])
	if err != nil {
		return "", err
	}
	dst, err := within(root, args[1])
	if err != nil {
		return "", err
	}
	return "", os.Rename(src, dst)
}

func builtinCp(root string, args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("usage: cp src dst")
	}
	src, err := within(root, args[0])
	if err != nil {
		return "", err
	}
	dst, err := within(root, args[1])
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return "", err
	}
	return "", os.WriteFile(dst, data, 0o644)
}
