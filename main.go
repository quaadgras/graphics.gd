// The graphics.gd harness: a coding agent in a terminal rendered by the
// engine itself. It edits, builds and ships graphics.gd projects — the
// long game is doing all of that on-device, iPhone included.
//
// Run it with gd (the harness is itself a graphics.gd project):
//
//	ANTHROPIC_API_KEY=... gd run
//
// It operates on the project named by GD_HARNESS_PROJECT, defaulting to
// the harness's own working directory (it can work on itself).
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
	project := os.Getenv("GD_HARNESS_PROJECT")
	if project == "" {
		project, _ = os.Getwd()
	}
	project, err := filepath.Abs(project)
	if err != nil {
		panic(err)
	}

	startup.LoadingScene()
	console := term.New()
	SceneTree.Add(console.Root())
	console.Focus()

	tools := toolchain.New(project, console)
	ai := agent.New(project, tools, console)

	console.Print("[color=#50fa7b]graphics.gd harness[/color] — project: " + term.Escape(project))
	if !ai.Ready() {
		console.System("ANTHROPIC_API_KEY is not set: agent chat is disabled, ! and / still work.")
	}
	console.System("/help for commands")

	// One worker goroutine owns the agent and every blocking operation;
	// the main thread only renders and forwards input.
	inbox := make(chan string, 16)
	go worker(inbox, console, tools, ai)

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

func worker(inbox <-chan string, console *term.Console, tools *toolchain.Toolchain, ai *agent.Agent) {
	for line := range inbox {
		switch {
		case strings.HasPrefix(line, "!"):
			out, err := tools.Shell(strings.TrimPrefix(line, "!"), 5*time.Minute)
			if out != "" {
				console.Printf("%s", out)
			}
			if err != nil {
				console.Error(err.Error())
			}
		case strings.HasPrefix(line, "/"):
			command(line, console, tools, ai)
		default:
			ai.Turn(line)
		}
	}
}

func command(line string, console *term.Console, tools *toolchain.Toolchain, ai *agent.Agent) {
	verb, _, _ := strings.Cut(strings.TrimPrefix(line, "/"), " ")
	var err error
	switch verb {
	case "help":
		console.System("/build — gd build for the host")
		console.System("/test — gd test")
		console.System("/deploy — GOOS=ios gd run: build, export and serve a SideStore install")
		console.System("/android — GOOS=android gd run: build, install and launch on this device")
		console.System("/clear — forget the conversation")
		console.System("/quit — exit")
		console.System("!command — run a shell command in the project")
		console.System("anything else — talk to the agent")
	case "build":
		err = tools.GD("build", "")
	case "test":
		err = tools.GD("test", "")
	case "deploy":
		err = tools.GD("run", "ios")
	case "android":
		err = tools.GD("run", "android")
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
}
