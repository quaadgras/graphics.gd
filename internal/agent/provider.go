package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// entry is one turn in the neutral conversation the agent owns,
// independent of any provider's wire format. Providers translate a
// slice of these into their own request shape.
type entry struct {
	role    string       // "user" | "assistant"
	text    string       // model or user prose
	calls   []toolCall   // assistant-requested tool calls
	results []toolResult // user-supplied tool results
}

type toolCall struct {
	id    string
	name  string
	input json.RawMessage
}

type toolResult struct {
	id      string
	output  string
	isError bool
}

// reply is a provider's response for one round.
type reply struct {
	text  string
	calls []toolCall
}

// complete sends the conversation to the provider and returns its
// reply. It dispatches on p.Format.
func complete(ctx context.Context, client *http.Client, p Provider, key, system string, history []entry, tools []tool) (reply, error) {
	switch p.Format {
	case "anthropic":
		return completeAnthropic(ctx, client, p, key, system, history, tools)
	case "openai":
		return completeOpenAI(ctx, client, p, key, system, history, tools)
	default:
		return reply{}, fmt.Errorf("unknown provider format %q", p.Format)
	}
}

func do(ctx context.Context, client *http.Client, url string, headers map[string]string, body any) ([]byte, string, error) {
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, "", err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(raw))
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("content-type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	return data, res.Status, err
}

// ---- Anthropic Messages ----

func completeAnthropic(ctx context.Context, client *http.Client, p Provider, key, system string, history []entry, tools []tool) (reply, error) {
	type block struct {
		Type      string          `json:"type"`
		Text      string          `json:"text,omitempty"`
		ID        string          `json:"id,omitempty"`
		Name      string          `json:"name,omitempty"`
		Input     json.RawMessage `json:"input,omitempty"`
		ToolUseID string          `json:"tool_use_id,omitempty"`
		Content   string          `json:"content,omitempty"`
		IsError   bool            `json:"is_error,omitempty"`
	}
	type message struct {
		Role    string  `json:"role"`
		Content []block `json:"content"`
	}
	var messages []message
	for _, e := range history {
		var blocks []block
		if e.text != "" {
			blocks = append(blocks, block{Type: "text", Text: e.text})
		}
		for _, c := range e.calls {
			blocks = append(blocks, block{Type: "tool_use", ID: c.id, Name: c.name, Input: c.input})
		}
		for _, r := range e.results {
			blocks = append(blocks, block{Type: "tool_result", ToolUseID: r.id, Content: r.output, IsError: r.isError})
		}
		if len(blocks) > 0 {
			messages = append(messages, message{Role: e.role, Content: blocks})
		}
	}
	atools := make([]map[string]any, len(tools))
	for i, t := range tools {
		atools[i] = map[string]any{"name": t.Name, "description": t.Description, "input_schema": t.InputSchema}
	}
	body := map[string]any{
		"model": p.Model, "max_tokens": maxTokens, "system": system,
		"messages": messages, "tools": atools,
	}
	data, status, err := do(ctx, client, p.BaseURL+"/v1/messages", map[string]string{
		"x-api-key": key, "anthropic-version": "2023-06-01",
	}, body)
	if err != nil {
		return reply{}, err
	}
	var res struct {
		Content []block `json:"content"`
		Error   *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return reply{}, fmt.Errorf("bad response (%s): %.200s", status, data)
	}
	if res.Error != nil {
		return reply{}, fmt.Errorf("%s", res.Error.Message)
	}
	var out reply
	for _, b := range res.Content {
		switch b.Type {
		case "text":
			out.text += b.Text
		case "tool_use":
			out.calls = append(out.calls, toolCall{id: b.ID, name: b.Name, input: b.Input})
		}
	}
	return out, nil
}

// ---- OpenAI chat completions (Grok, Qwen, OpenAI, local) ----

func completeOpenAI(ctx context.Context, client *http.Client, p Provider, key, system string, history []entry, tools []tool) (reply, error) {
	type function struct {
		Name      string          `json:"name"`
		Arguments string          `json:"arguments,omitempty"`
		Params    json.RawMessage `json:"parameters,omitempty"`
	}
	type toolCallJSON struct {
		ID       string   `json:"id"`
		Type     string   `json:"type"`
		Function function `json:"function"`
	}
	type message struct {
		Role       string         `json:"role"`
		Content    string         `json:"content,omitempty"`
		ToolCalls  []toolCallJSON `json:"tool_calls,omitempty"`
		ToolCallID string         `json:"tool_call_id,omitempty"`
	}
	messages := []message{{Role: "system", Content: system}}
	for _, e := range history {
		switch {
		case e.role == "assistant":
			m := message{Role: "assistant", Content: e.text}
			for _, c := range e.calls {
				m.ToolCalls = append(m.ToolCalls, toolCallJSON{
					ID: c.id, Type: "function",
					Function: function{Name: c.name, Arguments: string(c.input)},
				})
			}
			messages = append(messages, m)
		case len(e.results) > 0:
			// Tool results are their own role:tool messages.
			for _, r := range e.results {
				content := r.output
				if r.isError {
					content = "ERROR: " + content
				}
				messages = append(messages, message{Role: "tool", ToolCallID: r.id, Content: content})
			}
		default:
			messages = append(messages, message{Role: "user", Content: e.text})
		}
	}
	otools := make([]map[string]any, len(tools))
	for i, t := range tools {
		otools[i] = map[string]any{
			"type": "function",
			"function": map[string]any{
				"name": t.Name, "description": t.Description, "parameters": t.InputSchema,
			},
		}
	}
	body := map[string]any{
		"model": p.Model, "max_tokens": maxTokens, "messages": messages, "tools": otools,
	}
	data, status, err := do(ctx, client, p.BaseURL+"/v1/chat/completions", map[string]string{
		"authorization": "Bearer " + key,
	}, body)
	if err != nil {
		return reply{}, err
	}
	var res struct {
		Choices []struct {
			Message struct {
				Content   string         `json:"content"`
				ToolCalls []toolCallJSON `json:"tool_calls"`
			} `json:"message"`
		} `json:"choices"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return reply{}, fmt.Errorf("bad response (%s): %.200s", status, data)
	}
	if res.Error != nil {
		return reply{}, fmt.Errorf("%s", res.Error.Message)
	}
	if len(res.Choices) == 0 {
		return reply{}, fmt.Errorf("no choices (%s): %.200s", status, data)
	}
	msg := res.Choices[0].Message
	out := reply{text: msg.Content}
	for _, c := range msg.ToolCalls {
		input := json.RawMessage(c.Function.Arguments)
		if len(input) == 0 {
			input = json.RawMessage("{}")
		}
		out.calls = append(out.calls, toolCall{id: c.ID, name: c.Function.Name, input: input})
	}
	return out, nil
}
