package agent

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestProviderFormats checks that both wire formats serialize a
// conversation with a tool-call round trip correctly, and that each
// parses its own response shape back into the neutral reply.
func TestProviderFormats(t *testing.T) {
	history := []entry{
		{role: "user", text: "list the files"},
		{role: "assistant", text: "", calls: []toolCall{{id: "c1", name: "ls", input: json.RawMessage(`{"path":"."}`)}}},
		{role: "user", results: []toolResult{{id: "c1", output: "main.go\n"}}},
	}
	tools := toolDefinitions

	t.Run("anthropic", func(t *testing.T) {
		var got map[string]any
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("x-api-key") != "k" || r.Header.Get("anthropic-version") == "" {
				t.Errorf("missing anthropic auth headers")
			}
			body, _ := io.ReadAll(r.Body)
			json.Unmarshal(body, &got)
			io.WriteString(w, `{"content":[{"type":"text","text":"done"},{"type":"tool_use","id":"c2","name":"read","input":{"path":"main.go"}}]}`)
		}))
		defer srv.Close()
		p := Provider{Name: "a", BaseURL: srv.URL, Model: "m", Format: "anthropic"}
		rep, err := complete(context.Background(), srv.Client(), p, "k", "sys", history, tools)
		if err != nil {
			t.Fatal(err)
		}
		if got["system"] != "sys" || got["model"] != "m" {
			t.Fatalf("system/model not sent: %v", got["system"])
		}
		msgs, _ := got["messages"].([]any)
		if len(msgs) != 3 {
			t.Fatalf("want 3 messages, got %d", len(msgs))
		}
		if rep.text != "done" || len(rep.calls) != 1 || rep.calls[0].name != "read" {
			t.Fatalf("bad reply: %+v", rep)
		}
	})

	t.Run("openai", func(t *testing.T) {
		var got map[string]any
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("authorization") != "Bearer k" {
				t.Errorf("missing bearer auth")
			}
			body, _ := io.ReadAll(r.Body)
			json.Unmarshal(body, &got)
			io.WriteString(w, `{"choices":[{"message":{"content":"done","tool_calls":[{"id":"c2","type":"function","function":{"name":"read","arguments":"{\"path\":\"main.go\"}"}}]}}]}`)
		}))
		defer srv.Close()
		p := Provider{Name: "o", BaseURL: srv.URL, Model: "m", Format: "openai"}
		rep, err := complete(context.Background(), srv.Client(), p, "k", "sys", history, tools)
		if err != nil {
			t.Fatal(err)
		}
		msgs, _ := got["messages"].([]any)
		// system + user + assistant(with tool_calls) + tool result = 4
		if len(msgs) != 4 {
			t.Fatalf("want 4 messages, got %d", len(msgs))
		}
		first, _ := msgs[0].(map[string]any)
		if first["role"] != "system" {
			t.Fatalf("first message should be system, got %v", first["role"])
		}
		last, _ := msgs[3].(map[string]any)
		if last["role"] != "tool" || last["tool_call_id"] != "c1" {
			t.Fatalf("tool result not mapped to role:tool: %v", last)
		}
		if rep.text != "done" || len(rep.calls) != 1 || rep.calls[0].name != "read" {
			t.Fatalf("bad reply: %+v", rep)
		}
		if !strings.Contains(string(rep.calls[0].input), "main.go") {
			t.Fatalf("tool args not parsed: %s", rep.calls[0].input)
		}
	})
}

// TestConfigProviderResolution checks preset lookup, model override, and
// key seeding from the environment.
func TestConfigProviderResolution(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XAI_API_KEY", "grok-key")
	t.Setenv("GD_HARNESS_PROVIDER", "grok")
	cfg := loadConfig(dir)
	if cfg.Provider != "grok" {
		t.Fatalf("provider = %q, want grok", cfg.Provider)
	}
	if cfg.key() != "grok-key" {
		t.Fatalf("grok key not seeded from env: %q", cfg.key())
	}
	p, err := cfg.provider()
	if err != nil {
		t.Fatal(err)
	}
	if p.Format != "openai" || p.BaseURL != "https://api.x.ai" {
		t.Fatalf("grok resolved wrong: %+v", p)
	}
	cfg.Model = "grok-4-fast"
	p, _ = cfg.provider()
	if p.Model != "grok-4-fast" {
		t.Fatalf("model override ignored: %s", p.Model)
	}
	if err := cfg.save(dir); err != nil {
		t.Fatal(err)
	}
	if reloaded := loadConfig(dir); reloaded.Provider != "grok" || reloaded.Model != "grok-4-fast" {
		t.Fatalf("config did not round-trip: %+v", reloaded)
	}
}
