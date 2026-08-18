package agent

import (
	"encoding/json"
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

// dispatch executes one tool call through the active kit, returning
// output and whether it errored.
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
	kit := a.kit()
	fail := func(err error) (string, bool) { return err.Error(), true }
	switch name {
	case "read":
		data, err := kit.ReadFile(args.Path)
		if err != nil {
			return fail(err)
		}
		const cap = 200 << 10
		if len(data) > cap {
			return string(data[:cap]) + "\n(truncated)", false
		}
		return string(data), false
	case "write":
		if err := kit.WriteFile(args.Path, []byte(args.Content)); err != nil {
			return fail(err)
		}
		return "wrote " + args.Path, false
	case "edit":
		data, err := kit.ReadFile(args.Path)
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
		if err := kit.WriteFile(args.Path, []byte(text)); err != nil {
			return fail(err)
		}
		return "edited " + args.Path, false
	case "ls":
		listing, err := kit.List(args.Path)
		if err != nil {
			return fail(err)
		}
		return listing, false
	case "run":
		out, err := kit.Shell(args.Command, 5*time.Minute)
		if err != nil {
			return out + "\nerror: " + err.Error(), true
		}
		return out, false
	case "gd":
		if err := kit.GD(args.Verb, args.Goos); err != nil {
			return fail(err)
		}
		return "gd " + args.Verb + " succeeded", false
	}
	return "unknown tool " + name, true
}
