package agent

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type tool struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	InputSchema json.RawMessage `json:"input_schema"`
}

func schema(properties string, required ...string) json.RawMessage {
	req, _ := json.Marshal(required)
	return json.RawMessage(`{"type":"object","properties":{` + properties + `},"required":` + string(req) + `}`)
}

var toolDefinitions = []tool{
	{
		Name:        "read",
		Description: "Read a file in the project.",
		InputSchema: schema(`"path":{"type":"string"}`, "path"),
	},
	{
		Name:        "write",
		Description: "Create or overwrite a file in the project.",
		InputSchema: schema(`"path":{"type":"string"},"content":{"type":"string"}`, "path", "content"),
	},
	{
		Name:        "edit",
		Description: "Replace an exact string in a file. old_string must match exactly once.",
		InputSchema: schema(`"path":{"type":"string"},"old_string":{"type":"string"},"new_string":{"type":"string"}`, "path", "old_string", "new_string"),
	},
	{
		Name:        "ls",
		Description: "List a directory in the project (default: the project root).",
		InputSchema: schema(`"path":{"type":"string"}`),
	},
	{
		Name:        "run",
		Description: "Run a shell command in the project root and return its combined output.",
		InputSchema: schema(`"command":{"type":"string"}`, "command"),
	},
	{
		Name:        "gd",
		Description: "Run the gd build tool on the project. verb is build, run or test. goos optionally targets another platform: ios exports and serves a SideStore install, android builds and installs an APK.",
		InputSchema: schema(`"verb":{"type":"string","enum":["build","run","test"]},"goos":{"type":"string"}`, "verb"),
	},
}

// dispatch executes one tool call, returning output and whether it errored.
func (a *Agent) dispatch(name string, input json.RawMessage) (string, bool) {
	var args struct {
		Path      string `json:"path"`
		Content   string `json:"content"`
		OldString string `json:"old_string"`
		NewString string `json:"new_string"`
		Command   string `json:"command"`
		Verb      string `json:"verb"`
		Goos      string `json:"goos"`
	}
	if err := json.Unmarshal(input, &args); err != nil {
		return "bad tool input: " + err.Error(), true
	}
	fail := func(err error) (string, bool) { return err.Error(), true }
	switch name {
	case "read":
		path, err := a.resolve(args.Path)
		if err != nil {
			return fail(err)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fail(err)
		}
		const cap = 200 << 10
		if len(data) > cap {
			return string(data[:cap]) + "\n(truncated)", false
		}
		return string(data), false
	case "write":
		path, err := a.resolve(args.Path)
		if err != nil {
			return fail(err)
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return fail(err)
		}
		if err := os.WriteFile(path, []byte(args.Content), 0o644); err != nil {
			return fail(err)
		}
		return "wrote " + args.Path, false
	case "edit":
		path, err := a.resolve(args.Path)
		if err != nil {
			return fail(err)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fail(err)
		}
		text := string(data)
		switch strings.Count(text, args.OldString) {
		case 0:
			return "old_string not found in " + args.Path, true
		case 1:
		default:
			return "old_string matches more than once in " + args.Path, true
		}
		text = strings.Replace(text, args.OldString, args.NewString, 1)
		if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
			return fail(err)
		}
		return "edited " + args.Path, false
	case "ls":
		path, err := a.resolve(args.Path)
		if err != nil {
			return fail(err)
		}
		entries, err := os.ReadDir(path)
		if err != nil {
			return fail(err)
		}
		var b strings.Builder
		for _, e := range entries {
			if e.IsDir() {
				fmt.Fprintf(&b, "%s/\n", e.Name())
			} else {
				fmt.Fprintf(&b, "%s\n", e.Name())
			}
		}
		return b.String(), false
	case "run":
		out, err := a.tc.Shell(args.Command, 5*time.Minute)
		if err != nil {
			return out + "\nerror: " + err.Error(), true
		}
		return out, false
	case "gd":
		if err := a.tc.GD(args.Verb, args.Goos); err != nil {
			return fail(err)
		}
		return "gd " + args.Verb + " succeeded", false
	}
	return "unknown tool " + name, true
}

// resolve joins a tool-supplied path onto the project root and refuses
// to escape it.
func (a *Agent) resolve(path string) (string, error) {
	if path == "" {
		path = "."
	}
	joined := filepath.Clean(filepath.Join(a.project, path))
	if joined != a.project && !strings.HasPrefix(joined, a.project+string(filepath.Separator)) {
		return "", fmt.Errorf("path escapes the project: %s", path)
	}
	return joined, nil
}
