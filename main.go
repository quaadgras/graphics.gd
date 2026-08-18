// The graphics.gd harness: a coding agent in a terminal rendered by the
// engine itself. It edits, builds and ships graphics.gd projects — the
// long game is doing all of that on-device, iPhone included.
//
// Run it with gd (the harness is itself a graphics.gd project):
//
//	ANTHROPIC_API_KEY=... gd run
//
// It operates on the project named by GD_HARNESS_PROJECT (default: the
// working directory; on iOS, the app sandbox). With /remote configured
// it instead operates on another machine's checkout over SSH — which is
// how an iPhone gets a full build loop today: the remote runs
// GOOS=ios gd run, and the harness opens the served SideStore install
// link on-device.
package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"

	"graphics.gd/harness/internal/agent"
	"graphics.gd/harness/internal/term"
	"graphics.gd/harness/internal/toolchain"
)

func main() {
	startup.LoadingScene()
	project, err := filepath.Abs(defaultProject())
	if err != nil {
		panic(err)
	}
	stateDir := filepath.Join(project, ".harness")

	console := term.New()
	SceneTree.Add(console.Root())
	console.Focus()

	local := toolchain.New(project, console)
	remote, _ := toolchain.LoadRemote(stateDir, console)
	if remote != nil {
		remote.OnURL = installURL(console)
	}
	kit := func() toolchain.Kit {
		if remote != nil {
			return remote
		}
		return local
	}
	ai := agent.New(stateDir, kit, console)

	console.Print("[color=#50fa7b]graphics.gd harness[/color] — project: " + term.Escape(kit().Describe()))
	if !ai.Ready() {
		console.System("no API key: /provider to pick an AI (anthropic, grok, qwen, openai), /key to set its key.")
	}
	console.System("/help for commands")

	// One worker goroutine owns the agent and every blocking operation;
	// the main thread only renders and forwards input.
	inbox := make(chan string, 16)
	go func() {
		for line := range inbox {
			switch {
			case strings.HasPrefix(line, "!"):
				out, err := kit().Shell(strings.TrimPrefix(line, "!"), 5*time.Minute)
				if out != "" {
					console.Printf("%s", out)
				}
				if err != nil {
					console.Error(err.Error())
				}
			case strings.HasPrefix(line, "/"):
				remote = command(line, console, kit, remote, stateDir, ai)
			default:
				ai.Turn(line)
			}
		}
	}()

	console.OnSubmit(func(line string) {
		select {
		case inbox <- line:
		default:
			console.Error("busy — previous command is still running")
		}
	})

	for range startup.Rendering() {
		console.Flush()
	}
}

// command handles a /verb line and returns the (possibly changed)
// remote kit.
func command(line string, console *term.Console, kit func() toolchain.Kit, remote *toolchain.Remote, stateDir string, ai *agent.Agent) *toolchain.Remote {
	verb, rest, _ := strings.Cut(strings.TrimPrefix(line, "/"), " ")
	rest = strings.TrimSpace(rest)
	var err error
	switch verb {
	case "help":
		console.System("/provider [name] — list AI providers, or switch (anthropic, grok, qwen, openai, custom)")
		console.System("/model <name> — override the model for the current provider")
		console.System("/key <key> — set and remember the API key for the current provider")
		console.System("/remote user@host[:port] /path/to/project — build via SSH on another machine")
		console.System("/remote off — back to the local project")
		console.System("/build — gd build for the host")
		console.System("/test — gd test")
		console.System("/deploy — GOOS=ios gd run: build, export and serve a SideStore install")
		console.System("/android — GOOS=android gd run: build, install and launch on that device")
		console.System("/clear — forget the conversation")
		console.System("/quit — exit")
		console.System("!command — run a shell command in the project")
		console.System("anything else — talk to the agent")
	case "provider":
		switch {
		case rest == "":
			for _, line := range ai.Status() {
				console.System(line)
			}
		case strings.HasPrefix(rest, "custom"):
			fields := strings.Fields(rest)
			if len(fields) != 3 {
				console.Error("usage: /provider custom <base-url> <model>")
				break
			}
			if err = ai.SetCustomProvider(fields[1], fields[2]); err == nil {
				console.System("custom provider set: " + fields[1] + " / " + fields[2] + " — set its key with /key")
			}
		default:
			if err = ai.SetProvider(rest); err == nil {
				console.System("provider: " + rest)
				if !ai.Ready() {
					console.System("no key yet — set one with /key <key>")
				}
			}
		}
	case "model":
		if rest == "" {
			console.Error("usage: /model <name>")
			break
		}
		if err = ai.SetModel(rest); err == nil {
			console.System("model: " + rest)
		}
	case "key":
		if rest == "" {
			console.Error("usage: /key <key>")
			break
		}
		if err = ai.SetKey(rest); err == nil {
			console.System("API key saved — agent chat enabled")
		}
	case "remote":
		if rest == "off" || rest == "" {
			toolchain.ForgetRemote(stateDir)
			if rest == "off" {
				console.System("remote cleared — using the local project")
			} else {
				console.Error("usage: /remote user@host[:port] /path/to/project (or /remote off)")
			}
			return nil
		}
		next, err := toolchain.NewRemote(rest, stateDir, console)
		if err != nil {
			console.Error(err.Error())
			return remote
		}
		next.OnURL = installURL(console)
		console.System("remote set: " + next.Describe())
		console.System("authorize it by adding this key to the remote's ~/.ssh/authorized_keys:")
		console.Printf("%s", next.PublicKey())
		return next
	case "build":
		err = kit().GD("build", "")
	case "test":
		err = kit().GD("test", "")
	case "deploy":
		err = kit().GD("run", "ios")
	case "android":
		err = kit().GD("run", "android")
	case "clear":
		ai.Reset()
		console.System("conversation cleared")
	case "quit", "exit":
		os.Exit(0)
	default:
		console.Error("unknown command " + line + " — /help")
	}
	if err != nil {
		console.Error(err.Error())
	}
	return remote
}

// installURL reacts to a sidestore:// link appearing in remote build
// output: on iOS it opens SideStore to install onto this device.
func installURL(console *term.Console) func(string) {
	return func(url string) {
		openURL(console, url)
	}
}
