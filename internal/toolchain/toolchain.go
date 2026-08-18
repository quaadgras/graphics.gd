// Package toolchain abstracts how the harness reads, edits, compiles,
// runs and ships the project it is working on.
//
// A Kit is one implementation of that: the local kit uses the
// filesystem and (where the OS allows processes) the gd command; the
// remote kit in remote.go drives another machine's checkout over SSH,
// which is how an iPhone gets a full build loop before the in-process
// toolchain tier exists. On iOS the local kit's Shell is a small
// built-in userland and GD reports that building needs the roadmap's
// in-process tier (compiler.gd + embedded lld).
package toolchain

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Kit reads, edits, builds and runs a project.
type Kit interface {
	// Describe says where this kit's project lives, for the banner.
	Describe() string
	// Shell runs a command in the project root, returning combined output.
	Shell(command string, timeout time.Duration) (string, error)
	// GD runs a gd verb (build, run, test) against the project,
	// streaming progress. An empty goos targets the default platform.
	GD(verb string, goos string) error
	// ReadFile, WriteFile, List and paths in general are relative to
	// the project root; implementations must refuse escapes.
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
	List(path string) (string, error)
	// Grep searches text files under path (default the project root)
	// for a literal substring, returning file:line:text lines.
	Grep(pattern, path string) (string, error)
}

// grepSkipDirs are directories Grep never descends into.
var grepSkipDirs = map[string]bool{
	".git": true, ".godot": true, "node_modules": true, "releases": true,
}

const grepMaxMatches = 200

// Toolchain is the local Kit, rooted at Project.
type Toolchain struct {
	Project string
	Out     io.Writer
}

func New(project string, out io.Writer) *Toolchain {
	return &Toolchain{Project: project, Out: out}
}

func (t *Toolchain) Describe() string { return t.Project }

func (t *Toolchain) resolve(path string) (string, error) {
	if path == "" {
		path = "."
	}
	joined := filepath.Clean(filepath.Join(t.Project, path))
	if joined != t.Project && !strings.HasPrefix(joined, t.Project+string(filepath.Separator)) {
		return "", fmt.Errorf("path escapes the project: %s", path)
	}
	return joined, nil
}

func (t *Toolchain) ReadFile(path string) ([]byte, error) {
	resolved, err := t.resolve(path)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(resolved)
}

func (t *Toolchain) WriteFile(path string, data []byte) error {
	resolved, err := t.resolve(path)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(resolved), 0o755); err != nil {
		return err
	}
	return os.WriteFile(resolved, data, 0o644)
}

func (t *Toolchain) List(path string) (string, error) {
	resolved, err := t.resolve(path)
	if err != nil {
		return "", err
	}
	entries, err := os.ReadDir(resolved)
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

func (t *Toolchain) Grep(pattern, path string) (string, error) {
	if pattern == "" {
		return "", fmt.Errorf("grep needs a pattern")
	}
	root, err := t.resolve(path)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	matches := 0
	filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if p != root && grepSkipDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if matches >= grepMaxMatches {
			return filepath.SkipAll
		}
		data, err := os.ReadFile(p)
		if err != nil || bytes.IndexByte(data, 0) >= 0 { // unreadable or binary
			return nil
		}
		rel, _ := filepath.Rel(t.Project, p)
		for i, line := range strings.Split(string(data), "\n") {
			if strings.Contains(line, pattern) {
				fmt.Fprintf(&b, "%s:%d:%s\n", rel, i+1, strings.TrimSpace(line))
				if matches++; matches >= grepMaxMatches {
					break
				}
			}
		}
		return nil
	})
	if matches >= grepMaxMatches {
		b.WriteString("(truncated at " + fmt.Sprint(grepMaxMatches) + " matches)\n")
	}
	return b.String(), nil
}
