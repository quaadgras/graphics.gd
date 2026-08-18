package agent

import (
	"encoding/json"
	"fmt"
	"regexp"
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
		Description: "Replace a string in a file. old_string must match exactly (include enough surrounding lines to be unique), and by default must occur exactly once. Set replace_all to replace every occurrence.",
		InputSchema: schema(`"path":{"type":"string"},"old_string":{"type":"string"},"new_string":{"type":"string"},"replace_all":{"type":"boolean"}`, "path", "old_string", "new_string"),
	},
	{
		Name:        "ls",
		Description: "List a directory in the project (default: the project root).",
		InputSchema: schema(`"path":{"type":"string"}`),
	},
	{
		Name:        "grep",
		Description: "Search text files under path (default the whole project) for a literal substring. Returns file:line:text. Use this to locate code before editing.",
		InputSchema: schema(`"pattern":{"type":"string"},"path":{"type":"string"}`, "pattern"),
	},
	{
		Name:        "replace",
		Description: "Regex find-and-replace in one file (sed-like). pattern is an RE2 regular expression; replacement may use $1 or ${name} capture-group backrefs. Replaces every match. Use edit for a literal one-off change.",
		InputSchema: schema(`"path":{"type":"string"},"pattern":{"type":"string"},"replacement":{"type":"string"}`, "path", "pattern", "replacement"),
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
		Path       string `json:"path"`
		Content     string `json:"content"`
		OldString   string `json:"old_string"`
		NewString   string `json:"new_string"`
		ReplaceAll  bool   `json:"replace_all"`
		Pattern     string `json:"pattern"`
		Replacement string `json:"replacement"`
		Command     string `json:"command"`
		Verb       string `json:"verb"`
		Goos       string `json:"goos"`
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
		if args.OldString == "" {
			return "old_string is empty — use write to create a file", true
		}
		data, err := kit.ReadFile(args.Path)
		if err != nil {
			return fail(err)
		}
		text := string(data)
		n := strings.Count(text, args.OldString)
		switch {
		case n == 0:
			return editNotFound(text, args.OldString, args.Path), true
		case n > 1 && !args.ReplaceAll:
			return fmt.Sprintf("old_string matches %d times in %s — add surrounding lines to make it unique, or pass replace_all=true", n, args.Path), true
		}
		if args.ReplaceAll {
			text = strings.ReplaceAll(text, args.OldString, args.NewString)
		} else {
			text = strings.Replace(text, args.OldString, args.NewString, 1)
		}
		if err := kit.WriteFile(args.Path, []byte(text)); err != nil {
			return fail(err)
		}
		return fmt.Sprintf("edited %s (%d replacement(s))", args.Path, n), false
	case "grep":
		out, err := kit.Grep(args.Pattern, args.Path)
		if err != nil {
			return fail(err)
		}
		if strings.TrimSpace(out) == "" {
			return "no matches for " + args.Pattern, false
		}
		return out, false
	case "replace":
		re, err := regexp.Compile(args.Pattern)
		if err != nil {
			return "bad regex: " + err.Error(), true
		}
		data, err := kit.ReadFile(args.Path)
		if err != nil {
			return fail(err)
		}
		text := string(data)
		n := len(re.FindAllStringIndex(text, -1))
		if n == 0 {
			return "no matches for /" + args.Pattern + "/ in " + args.Path, true
		}
		if err := kit.WriteFile(args.Path, []byte(re.ReplaceAllString(text, args.Replacement))); err != nil {
			return fail(err)
		}
		return fmt.Sprintf("replaced %d match(es) in %s", n, args.Path), false
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

// editNotFound explains a failed edit match, catching the common case
// where only whitespace differs — the failure mode weaker models hit
// most — so the model can correct rather than flail.
func editNotFound(text, old, path string) string {
	if strings.Contains(squeeze(text), squeeze(old)) {
		return "old_string not found in " + path + " — the text matches except for whitespace/indentation. Read the file and copy the exact bytes, including leading tabs/spaces."
	}
	// If the first non-blank line of old_string exists, point there.
	for _, line := range strings.Split(old, "\n") {
		if t := strings.TrimSpace(line); t != "" {
			if strings.Contains(text, t) {
				return fmt.Sprintf("old_string not found in %s — a line like %q exists but the surrounding text differs. Read the file and copy the current contents exactly.", path, truncate(t, 60))
			}
			break
		}
	}
	return "old_string not found in " + path + " — read the file to see its current contents."
}

// squeeze removes all whitespace, for a whitespace-insensitive
// containment check.
func squeeze(s string) string { return strings.Join(strings.Fields(s), "") }

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n] + "…"
	}
	return s
}
