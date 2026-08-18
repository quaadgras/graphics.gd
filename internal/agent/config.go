package agent

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Provider describes one model backend. Anthropic speaks its Messages
// API; everyone else here (xAI Grok, Qwen/DashScope, OpenAI, local
// servers) speaks the OpenAI chat-completions API, so one "openai"
// format covers them all — only the base URL and default model change.
type Provider struct {
	Name    string `json:"name"`
	BaseURL string `json:"base_url"`
	Model   string `json:"model"`
	Format  string `json:"format"` // "anthropic" | "openai"
	KeyHint string `json:"key_hint,omitempty"`
	KeyEnv  string `json:"-"` // env var consulted for an initial key
}

// builtinProviders are the presets. "custom" is added at runtime from
// Config.Custom.
var builtinProviders = map[string]Provider{
	"anthropic": {
		Name: "anthropic", BaseURL: "https://api.anthropic.com",
		Model: "claude-sonnet-5", Format: "anthropic",
		KeyHint: "console.anthropic.com/settings/keys", KeyEnv: "ANTHROPIC_API_KEY",
	},
	"grok": {
		Name: "grok", BaseURL: "https://api.x.ai",
		Model: "grok-4", Format: "openai",
		KeyHint: "console.x.ai", KeyEnv: "XAI_API_KEY",
	},
	"qwen": {
		Name: "qwen", BaseURL: "https://dashscope-intl.aliyuncs.com/compatible-mode",
		Model: "qwen-max", Format: "openai",
		KeyHint: "dashscope.console.aliyun.com/apiKey", KeyEnv: "DASHSCOPE_API_KEY",
	},
	"openai": {
		Name: "openai", BaseURL: "https://api.openai.com",
		Model: "gpt-4o", Format: "openai",
		KeyHint: "platform.openai.com/api-keys", KeyEnv: "OPENAI_API_KEY",
	},
}

// Config is the persisted agent settings.
type Config struct {
	Provider string            `json:"provider"`         // preset name or "custom"
	Model    string            `json:"model,omitempty"`  // overrides the provider default
	Keys     map[string]string `json:"keys,omitempty"`   // api key per provider name
	Custom   *Provider         `json:"custom,omitempty"` // user-defined endpoint
}

func configFile(stateDir string) string { return filepath.Join(stateDir, "config.json") }

// loadConfig reads settings, seeding from the environment and migrating
// a legacy .harness/key (an Anthropic key) on first run.
func loadConfig(stateDir string) *Config {
	cfg := &Config{Keys: map[string]string{}}
	if data, err := os.ReadFile(configFile(stateDir)); err == nil {
		_ = json.Unmarshal(data, cfg)
	}
	if cfg.Keys == nil {
		cfg.Keys = map[string]string{}
	}
	// Seed keys from the environment for any provider missing one.
	for name, p := range builtinProviders {
		if cfg.Keys[name] == "" && p.KeyEnv != "" {
			if v := strings.TrimSpace(os.Getenv(p.KeyEnv)); v != "" {
				cfg.Keys[name] = v
			}
		}
	}
	// Migrate the legacy single-key file (Anthropic only).
	if cfg.Keys["anthropic"] == "" {
		if legacy, err := os.ReadFile(filepath.Join(stateDir, "key")); err == nil {
			cfg.Keys["anthropic"] = strings.TrimSpace(string(legacy))
		}
	}
	// Pick a default provider: explicit env, else the first preset that
	// has a key, else anthropic.
	if cfg.Provider == "" {
		cfg.Provider = strings.TrimSpace(os.Getenv("GD_HARNESS_PROVIDER"))
	}
	if cfg.Provider == "" {
		cfg.Provider = "anthropic"
		for _, name := range providerOrder() {
			if cfg.Keys[name] != "" {
				cfg.Provider = name
				break
			}
		}
	}
	if env := strings.TrimSpace(os.Getenv("GD_HARNESS_MODEL")); env != "" && cfg.Model == "" {
		cfg.Model = env
	}
	return cfg
}

func (c *Config) save(stateDir string) error {
	if err := os.MkdirAll(stateDir, 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile(stateDir), append(data, '\n'), 0o600)
}

// provider returns the resolved provider (preset or custom) with the
// configured model applied.
func (c *Config) provider() (Provider, error) {
	var p Provider
	if c.Provider == "custom" {
		if c.Custom == nil {
			return Provider{}, fmt.Errorf("provider is 'custom' but no custom endpoint is set (see /provider custom <url> <model>)")
		}
		p = *c.Custom
		p.Name = "custom"
	} else {
		preset, ok := builtinProviders[c.Provider]
		if !ok {
			return Provider{}, fmt.Errorf("unknown provider %q (see /provider)", c.Provider)
		}
		p = preset
	}
	if c.Model != "" {
		p.Model = c.Model
	}
	return p, nil
}

func (c *Config) key() string { return c.Keys[c.Provider] }

// providerOrder lists preset names in a stable, friendly order.
func providerOrder() []string {
	names := make([]string, 0, len(builtinProviders))
	for name := range builtinProviders {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		rank := map[string]int{"anthropic": 0, "grok": 1, "qwen": 2, "openai": 3}
		return rank[names[i]] < rank[names[j]]
	})
	return names
}
