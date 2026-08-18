package agent

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"graphics.gd/harness/internal/toolchain"
)

// memKit is an in-memory Kit for exercising the editing tools without
// touching disk or the engine.
type memKit struct{ files map[string]string }

func (m *memKit) Describe() string { return "mem" }
func (m *memKit) Shell(string, time.Duration) (string, error) {
	return "", nil
}
func (m *memKit) GD(string, string) error { return nil }
func (m *memKit) ReadFile(p string) ([]byte, error) {
	s, ok := m.files[p]
	if !ok {
		return nil, &notExist{p}
	}
	return []byte(s), nil
}
func (m *memKit) WriteFile(p string, data []byte) error {
	m.files[p] = string(data)
	return nil
}
func (m *memKit) List(string) (string, error) { return "", nil }
func (m *memKit) Grep(pattern, _ string) (string, error) {
	var b strings.Builder
	for name, body := range m.files {
		for i, line := range strings.Split(body, "\n") {
			if strings.Contains(line, pattern) {
				b.WriteString(name + ":" + itoa(i+1) + ":" + line + "\n")
			}
		}
	}
	return b.String(), nil
}

type notExist struct{ p string }

func (e *notExist) Error() string { return "no such file: " + e.p }

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}

var _ toolchain.Kit = (*memKit)(nil)

func newTestAgent(files map[string]string) *Agent {
	kit := &memKit{files: files}
	return &Agent{kit: func() toolchain.Kit { return kit }}
}

func call(a *Agent, name string, input string) (string, bool) {
	return a.dispatch(name, json.RawMessage(input))
}

func TestEditUniqueAndReplaceAll(t *testing.T) {
	a := newTestAgent(map[string]string{
		"a.go": "x := 1\ny := 1\nz := 1\n",
	})

	// Ambiguous single edit is refused with a helpful message.
	out, isErr := call(a, "edit", `{"path":"a.go","old_string":"1","new_string":"2"}`)
	if !isErr || !strings.Contains(out, "matches 3 times") {
		t.Fatalf("expected ambiguity error, got (%v) %q", isErr, out)
	}

	// replace_all succeeds and reports the count.
	out, isErr = call(a, "edit", `{"path":"a.go","old_string":"1","new_string":"2","replace_all":true}`)
	if isErr || !strings.Contains(out, "3 replacement") {
		t.Fatalf("replace_all failed: (%v) %q", isErr, out)
	}
	if got := a.kit().(*memKit).files["a.go"]; !strings.Contains(got, "x := 2") || strings.Contains(got, "1") {
		t.Fatalf("replace_all did not replace all: %q", got)
	}

	// A unique edit with surrounding context works.
	call(a, "write", `{"path":"b.go","content":"func A() {}\nfunc B() {}\n"}`)
	out, isErr = call(a, "edit", `{"path":"b.go","old_string":"func B() {}","new_string":"func B() int { return 0 }"}`)
	if isErr || !strings.Contains(out, "1 replacement") {
		t.Fatalf("unique edit failed: (%v) %q", isErr, out)
	}
}

func TestEditWhitespaceHint(t *testing.T) {
	a := newTestAgent(map[string]string{
		"c.go": "func main() {\n\treturn\n}\n", // real indent is a tab
	})
	// Model sends spaces where the file has a tab.
	out, isErr := call(a, "edit", `{"path":"c.go","old_string":"    return","new_string":"    return 1"}`)
	if !isErr || !strings.Contains(out, "whitespace") {
		t.Fatalf("expected whitespace hint, got (%v) %q", isErr, out)
	}
}

func TestReplaceTool(t *testing.T) {
	a := newTestAgent(map[string]string{
		"a.go": "oldName(1)\noldName(2)\nkeep\n",
	})
	// Capture-group backref rename.
	out, isErr := call(a, "replace", `{"path":"a.go","pattern":"oldName\\((\\d)\\)","replacement":"newName($1)"}`)
	if isErr || !strings.Contains(out, "2 match") {
		t.Fatalf("replace failed: (%v) %q", isErr, out)
	}
	got := a.kit().(*memKit).files["a.go"]
	if !strings.Contains(got, "newName(1)") || !strings.Contains(got, "newName(2)") || strings.Contains(got, "oldName") {
		t.Fatalf("replace wrong result: %q", got)
	}
	// Bad regex is reported, file untouched.
	out, isErr = call(a, "replace", `{"path":"a.go","pattern":"(unclosed","replacement":"x"}`)
	if !isErr || !strings.Contains(out, "bad regex") {
		t.Fatalf("expected bad-regex error, got (%v) %q", isErr, out)
	}
	// No matches is an error the model can act on.
	out, isErr = call(a, "replace", `{"path":"a.go","pattern":"zzz","replacement":"y"}`)
	if !isErr || !strings.Contains(out, "no matches") {
		t.Fatalf("expected no-matches, got (%v) %q", isErr, out)
	}
}

func TestGrepTool(t *testing.T) {
	a := newTestAgent(map[string]string{
		"x.go": "package main\nfunc Hello() {}\n",
		"y.go": "package main\n// Hello world\n",
	})
	out, isErr := call(a, "grep", `{"pattern":"Hello"}`)
	if isErr {
		t.Fatalf("grep errored: %q", out)
	}
	if !strings.Contains(out, "x.go:2:") || !strings.Contains(out, "y.go:2:") {
		t.Fatalf("grep missed matches: %q", out)
	}
	out, _ = call(a, "grep", `{"pattern":"Nonexistent"}`)
	if !strings.Contains(out, "no matches") {
		t.Fatalf("expected no-matches message, got %q", out)
	}
}
