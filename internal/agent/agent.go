// Package agent is a minimal coding-agent loop over the Anthropic
// Messages API: send the conversation, print text, execute tool calls,
// repeat until the model stops asking for tools. Pure net/http, no SDK.
package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"graphics.gd/harness/internal/term"
	"graphics.gd/harness/internal/toolchain"
)

const (
	apiURL       = "https://api.anthropic.com/v1/messages"
	apiVersion   = "2023-06-01"
	defaultModel = "claude-sonnet-5"
	maxTokens    = 8192
	maxToolTurns = 50
)

type Agent struct {
	stateDir string
	kit      func() toolchain.Kit
	console  *term.Console

	key   string
	model string
	httpc *http.Client
	msgs  []message
}

type message struct {
	Role    string  `json:"role"`
	Content []block `json:"content"`
}

type block struct {
	Type string `json:"type"`
	// text
	Text string `json:"text,omitempty"`
	// tool_use (from the model)
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name,omitempty"`
	Input json.RawMessage `json:"input,omitempty"`
	// tool_result (from us)
	ToolUseID string `json:"tool_use_id,omitempty"`
	Content   string `json:"content,omitempty"`
	IsError   bool   `json:"is_error,omitempty"`
}

type request struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	System    string    `json:"system"`
	Messages  []message `json:"messages"`
	Tools     []tool    `json:"tools"`
}

type response struct {
	Content    []block `json:"content"`
	StopReason string  `json:"stop_reason"`
	Error      *struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

// New builds an agent whose tools operate through kit — a function so
// the caller can swap between local and remote kits at runtime.
// stateDir holds harness state such as the saved API key.
func New(stateDir string, kit func() toolchain.Kit, console *term.Console) *Agent {
	model := os.Getenv("GD_HARNESS_MODEL")
	if model == "" {
		model = defaultModel
	}
	a := &Agent{
		stateDir: stateDir,
		kit:      kit,
		console:  console,
		key:      os.Getenv("ANTHROPIC_API_KEY"),
		model:    model,
		httpc:    &http.Client{Timeout: 5 * time.Minute},
	}
	if a.key == "" {
		if saved, err := os.ReadFile(a.keyFile()); err == nil {
			a.key = strings.TrimSpace(string(saved))
		}
	}
	return a
}

func (a *Agent) Ready() bool { return a.key != "" }

func (a *Agent) keyFile() string {
	return filepath.Join(a.stateDir, "key")
}

// SetKey stores the API key for this and future sessions. On platforms
// without environment variables (iOS) this is the only way in.
func (a *Agent) SetKey(key string) error {
	a.key = strings.TrimSpace(key)
	if err := os.MkdirAll(filepath.Dir(a.keyFile()), 0o700); err != nil {
		return err
	}
	return os.WriteFile(a.keyFile(), []byte(a.key+"\n"), 0o600)
}

// Reset clears the conversation.
func (a *Agent) Reset() { a.msgs = nil }

// Turn runs one user turn to completion, executing tool calls as the
// model requests them. Call from a single worker goroutine.
func (a *Agent) Turn(input string) {
	if !a.Ready() {
		a.console.Error("ANTHROPIC_API_KEY is not set — export it and restart, or use ! shell commands.")
		return
	}
	a.msgs = append(a.msgs, message{Role: "user", Content: []block{{Type: "text", Text: input}}})
	for range maxToolTurns {
		res, err := a.call()
		if err != nil {
			a.console.Error("api: " + err.Error())
			return
		}
		a.msgs = append(a.msgs, message{Role: "assistant", Content: res.Content})
		var results []block
		for _, b := range res.Content {
			switch b.Type {
			case "text":
				a.console.Print(term.Escape(b.Text))
			case "tool_use":
				a.console.System("⚙ " + b.Name + " " + summarize(b.Input))
				out, isErr := a.dispatch(b.Name, b.Input)
				results = append(results, block{
					Type: "tool_result", ToolUseID: b.ID, Content: out, IsError: isErr,
				})
			}
		}
		if res.StopReason != "tool_use" || len(results) == 0 {
			return
		}
		a.msgs = append(a.msgs, message{Role: "user", Content: results})
	}
	a.console.Error("stopping: too many tool calls in one turn")
}

func (a *Agent) call() (*response, error) {
	body, err := json.Marshal(request{
		Model:     a.model,
		MaxTokens: maxTokens,
		System:    a.systemPrompt(),
		Messages:  a.msgs,
		Tools:     toolDefinitions,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-api-key", a.key)
	req.Header.Set("anthropic-version", apiVersion)
	httpRes, err := a.httpc.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()
	data, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}
	var res response
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("bad response (%s): %.200s", httpRes.Status, data)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("%s: %s", res.Error.Type, res.Error.Message)
	}
	return &res, nil
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
