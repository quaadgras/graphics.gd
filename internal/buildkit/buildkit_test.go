package buildkit

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestZeroExecBuild proves the whole point: compile and link a Go
// program with the toolchain running as wasm inside this process.
// Opt-in (it cross-compiles cmd/compile and cmd/link to wasip1 with
// the host go, ~60MB of wasm): GD_HARNESS_WASMTC_TEST=1.
//
// The host go command is used only to PREPARE the fixture (the wasm
// tools, and export data for the hello program's dependencies); the
// build under test spawns no processes.
func TestZeroExecBuild(t *testing.T) {
	if os.Getenv("GD_HARNESS_WASMTC_TEST") != "1" {
		t.Skip("set GD_HARNESS_WASMTC_TEST=1 to run (builds the toolchain as wasm with the host go)")
	}
	work := t.TempDir()
	host := func(dir string, env []string, args ...string) string {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(), env...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%v: %v\n%s", args, err, out)
		}
		return string(out)
	}

	// Fixture: the toolchain as wasm, a hello module, and importcfgs.
	host(work, []string{"GOOS=wasip1", "GOARCH=wasm"}, "go", "build", "-o", work+"/compile.wasm", "cmd/compile")
	host(work, []string{"GOOS=wasip1", "GOARCH=wasm"}, "go", "build", "-o", work+"/link.wasm", "cmd/link")
	hello := filepath.Join(work, "hello")
	os.MkdirAll(hello, 0o755)
	os.WriteFile(filepath.Join(hello, "go.mod"), []byte("module hello\n\ngo 1.24\n"), 0o644)
	os.WriteFile(filepath.Join(hello, "main.go"), []byte("package main\n\nfunc main() { println(\"zero-exec build\") }\n"), 0o644)
	importcfg := host(hello, nil,
		"go", "list", "-deps", "-export",
		"-f", "{{if .Export}}packagefile {{.ImportPath}}={{.Export}}{{end}}", ".")
	os.WriteFile(filepath.Join(work, "importcfg"), []byte(importcfg), 0o644)
	os.WriteFile(filepath.Join(work, "importcfg.link"), []byte(importcfg), 0o644)

	compileWasm, err := os.ReadFile(work + "/compile.wasm")
	if err != nil {
		t.Fatal(err)
	}
	linkWasm, err := os.ReadFile(work + "/link.wasm")
	if err != nil {
		t.Fatal(err)
	}

	// The build under test: zero process spawns from here on.
	ctx := context.Background()
	runner := NewRunner(ctx)
	defer runner.Close(ctx)
	env := map[string]string{"GOOS": "linux", "GOARCH": "amd64", "TMPDIR": work}
	run := func(tool []byte, args ...string) {
		t.Helper()
		var out bytes.Buffer
		if err := runner.Run(ctx, Invocation{
			Tool: tool, Args: args, Env: env,
			Stdout: &out, Stderr: &out,
		}); err != nil {
			t.Fatalf("%v: %v\n%s", args, err, out.String())
		}
	}
	// Twice, same Runner: fresh instances make the compiler re-entrant.
	for round := range 2 {
		object := fmt.Sprintf("%s/main%d.a", work, round)
		run(compileWasm, "compile", "-p", "main", "-complete",
			"-importcfg", work+"/importcfg", "-pack", "-o", object,
			hello+"/main.go")
		run(linkWasm, "link", "-importcfg", work+"/importcfg.link",
			"-buildmode=exe", "-linkmode=internal",
			"-o", fmt.Sprintf("%s/hello%d", work, round), object)
	}

	built, err := os.ReadFile(work + "/hello1")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(built, []byte("\x7fELF")) {
		t.Fatalf("output is not an ELF executable (starts %q)", built[:4])
	}
	// Host exec only to VERIFY the artifact, not to build it.
	os.Chmod(work+"/hello1", 0o755)
	if out := host(work, nil, work+"/hello1"); !strings.Contains(out, "zero-exec build") {
		t.Fatalf("built binary printed %q", out)
	}
}
