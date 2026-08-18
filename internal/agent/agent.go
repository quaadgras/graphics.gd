// Package agent is a minimal coding-agent loop: send the conversation,
// print text, execute tool calls, repeat until the model stops asking
// for tools. Pure net/http, no SDK. The model backend is pluggable —
// Anthropic's Messages API or any OpenAI-compatible endpoint (xAI Grok,
// Qwen/DashScope, OpenAI, a local server) — selected in settings.
package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"graphics.gd/harness/internal/toolchain"
)

const (
	maxTokens    = 8192
	maxToolTurns = 50
)

// Console is the sink the agent writes to. term.Console satisfies it;
// keeping it an interface lets the agent's tests run without linking the
// engine.
type Console interface {
	// Printf prints escaped plain text (model prose, tool output).
	Printf(format string, args ...any)
	// System prints a dim status line.
	System(msg string)
	// Error prints an error line.
	Error(msg string)
}

type Agent struct {
	stateDir string
	kit      func() toolchain.Kit
	console  Console

	cfg     *Config
	httpc   *http.Client
	history []entry
}

// New builds an agent whose tools operate through kit — a function so
// the caller can swap between local and remote kits at runtime.
// stateDir holds harness state (settings, keys).
func New(stateDir string, kit func() toolchain.Kit, console Console) *Agent {
	return &Agent{
		stateDir: stateDir,
		kit:      kit,
		console:  console,
		cfg:      loadConfig(stateDir),
		httpc:    &http.Client{Timeout: 5 * time.Minute},
	}
}

// Ready reports whether the current provider has an API key.
func (a *Agent) Ready() bool { return a.cfg.key() != "" }

// SetKey stores the API key for the current provider.
func (a *Agent) SetKey(key string) error {
	a.cfg.Keys[a.cfg.Provider] = strings.TrimSpace(key)
	return a.cfg.save(a.stateDir)
}

// SetProvider switches to a named preset provider.
func (a *Agent) SetProvider(name string) error {
	name = strings.TrimSpace(name)
	if name != "custom" {
		if _, ok := builtinProviders[name]; !ok {
			return fmt.Errorf("unknown provider %q — try: %s, custom", name, strings.Join(providerOrder(), ", "))
		}
	}
	a.cfg.Provider = name
	a.cfg.Model = "" // clear any per-provider model override
	return a.cfg.save(a.stateDir)
}

// SetCustomProvider configures and selects a custom OpenAI-compatible
// endpoint.
func (a *Agent) SetCustomProvider(baseURL, model string) error {
	a.cfg.Custom = &Provider{
		Name: "custom", BaseURL: strings.TrimRight(strings.TrimSpace(baseURL), "/"),
		Model: strings.TrimSpace(model), Format: "openai",
	}
	a.cfg.Provider = "custom"
	a.cfg.Model = ""
	return a.cfg.save(a.stateDir)
}

// SetModel overrides the model for the current provider.
func (a *Agent) SetModel(model string) error {
	a.cfg.Model = strings.TrimSpace(model)
	return a.cfg.save(a.stateDir)
}

// Status returns provider/model/key lines for the /provider command.
func (a *Agent) Status() []string {
	var lines []string
	for _, name := range providerOrder() {
		mark := "  "
		if name == a.cfg.Provider {
			mark = "▸ "
		}
		have := "no key"
		if a.cfg.Keys[name] != "" {
			have = "key set"
		}
		p := builtinProviders[name]
		lines = append(lines, fmt.Sprintf("%s%-9s %-10s (%s) — %s", mark, name, p.Model, have, p.KeyHint))
	}
	if a.cfg.Custom != nil {
		mark := "  "
		if a.cfg.Provider == "custom" {
			mark = "▸ "
		}
		lines = append(lines, fmt.Sprintf("%scustom    %-10s %s", mark, a.cfg.Custom.Model, a.cfg.Custom.BaseURL))
	}
	if p, err := a.cfg.provider(); err == nil {
		lines = append(lines, fmt.Sprintf("using %s / %s", p.Name, p.Model))
	}
	return lines
}

// Reset clears the conversation.
func (a *Agent) Reset() { a.history = nil }

// Turn runs one user turn to completion, executing tool calls as the
// model requests them. Call from a single worker goroutine.
func (a *Agent) Turn(input string) {
	p, err := a.cfg.provider()
	if err != nil {
		a.console.Error(err.Error())
		return
	}
	if !a.Ready() {
		a.console.Error(fmt.Sprintf("no API key for %s — set one with /key <key> (from %s), or /provider to switch.", p.Name, p.KeyHint))
		return
	}
	ctx := context.Background()
	a.history = append(a.history, entry{role: "user", text: input})
	for range maxToolTurns {
		res, err := complete(ctx, a.httpc, p, a.cfg.key(), a.systemPrompt(), a.history, toolDefinitions)
		if err != nil {
			a.console.Error(p.Name + ": " + err.Error())
			return
		}
		if res.text != "" {
			a.console.Printf("%s", res.text)
		}
		a.history = append(a.history, entry{role: "assistant", text: res.text, calls: res.calls})
		if len(res.calls) == 0 {
			return
		}
		var results []toolResult
		for _, c := range res.calls {
			a.console.System("⚙ " + c.name + " " + summarize(c.input))
			out, isErr := a.dispatch(c.name, c.input)
			results = append(results, toolResult{id: c.id, output: out, isError: isErr})
		}
		a.history = append(a.history, entry{role: "user", results: results})
	}
	a.console.Error("stopping: too many tool calls in one turn")
}

func (a *Agent) systemPrompt() string {
	return "You are the graphics.gd harness, a coding agent embedded in a " +
		"terminal rendered by the very engine your user builds games with. " +
		"You edit and build the graphics.gd (Go bindings for Godot 4) project " +
		"at " + a.kit().Describe() + ". Paths in tool calls are relative to " +
		"that root. Use the gd tool to build, test or run the project (goos " +
		"ios exports and serves a SideStore install for an iPhone). Be direct " +
		"and keep output brief: it renders in a small scrollback."
}

// summarize renders tool input compactly for the scrollback.
func summarize(input json.RawMessage) string {
	var m map[string]any
	if json.Unmarshal(input, &m) != nil {
		return ""
	}
	for _, key := range []string{"path", "command", "verb"} {
		if v, ok := m[key].(string); ok {
			if len(v) > 80 {
				v = v[:80] + "…"
			}
			return v
		}
	}
	return ""
}
