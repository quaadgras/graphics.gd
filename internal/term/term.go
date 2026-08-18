// Package term renders a terminal-like console using engine nodes.
//
// The console is safe to write to from any goroutine: writes land in a
// channel and are appended to the scrollback by Flush, which the main
// loop calls once per frame. Engine nodes are only ever touched on the
// main thread.
package term

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/RichTextLabel"
	"graphics.gd/classdb/SystemFont"
	"graphics.gd/classdb/VBoxContainer"
)

// Console is a scrollback log plus an input line.
type Console struct {
	root  VBoxContainer.Instance
	log   RichTextLabel.Instance
	input LineEdit.Instance

	pending chan string // bbcode fragments awaiting Flush

	mu     sync.Mutex
	submit func(string)
}

// New builds the console's node tree; add Root to the scene afterwards.
func New() *Console {
	c := &Console{pending: make(chan string, 4096)}

	mono := SystemFont.New()
	mono.SetFontNames([]string{"JetBrains Mono", "Cascadia Mono", "Menlo", "monospace"})

	// Phone screens are dense; desktop keeps the theme default.
	fontSize := 0
	if runtime.GOOS == "ios" || runtime.GOOS == "android" {
		fontSize = 28
	}

	c.root = VBoxContainer.New()
	c.root.AsControl().SetAnchorsPreset(Control.PresetFullRect)

	c.log = RichTextLabel.New()
	c.log.SetBbcodeEnabled(true)
	c.log.SetScrollFollowing(true)
	c.log.SetSelectionEnabled(true)
	c.log.AsControl().SetSizeFlagsVertical(Control.SizeExpandFill)
	c.log.AsControl().AddThemeFontOverride("normal_font", mono.AsFont())
	if fontSize > 0 {
		c.log.AsControl().AddThemeFontSizeOverride("normal_font_size", fontSize)
	}
	c.root.AsNode().AddChild(c.log.AsNode())

	c.input = LineEdit.New()
	c.input.SetPlaceholderText("talk to the agent, !command for shell, /help for commands")
	c.input.AsControl().AddThemeFontOverride("font", mono.AsFont())
	if fontSize > 0 {
		c.input.AsControl().AddThemeFontSizeOverride("font_size", fontSize)
	}
	c.input.OnTextSubmitted(func(text string) {
		c.input.SetText("")
		line := strings.TrimSpace(text)
		if line == "" {
			return
		}
		echo := line
		if strings.HasPrefix(echo, "/key ") { // never render secrets into the scrollback
			echo = "/key ••••"
		}
		c.Print("[color=#8be9fd]> " + Escape(echo) + "[/color]")
		c.mu.Lock()
		submit := c.submit
		c.mu.Unlock()
		if submit != nil {
			submit(line)
		}
	})
	c.root.AsNode().AddChild(c.input.AsNode())

	return c
}

// Focus moves keyboard focus to the input line; call after the console
// has been added to the scene tree.
func (c *Console) Focus() { c.input.AsControl().GrabFocus() }

// Root is the node to add to the scene tree.
func (c *Console) Root() VBoxContainer.Instance { return c.root }

// OnSubmit sets the handler for submitted input lines. The handler runs
// on the main thread and must not block; hand the line to a goroutine.
func (c *Console) OnSubmit(fn func(line string)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.submit = fn
}

// Flush appends pending output to the scrollback. Main thread only;
// call once per frame.
func (c *Console) Flush() {
	for {
		select {
		case s := <-c.pending:
			c.log.AppendText(s)
		default:
			return
		}
	}
}

// Print queues a line of bbcode, unescaped, for the scrollback.
func (c *Console) Print(bbcode string) { c.pending <- bbcode + "\n" }

// Printf queues formatted plain text, escaped, for the scrollback.
func (c *Console) Printf(format string, args ...any) {
	c.Print(Escape(fmt.Sprintf(format, args...)))
}

// System prints a dim status line.
func (c *Console) System(msg string) {
	c.Print("[color=#6272a4]" + Escape(msg) + "[/color]")
}

// Error prints an error line.
func (c *Console) Error(msg string) {
	c.Print("[color=#ff5555]" + Escape(msg) + "[/color]")
}

// Write implements io.Writer with plain (escaped) text, so subprocess
// output and agent text can stream straight into the scrollback.
func (c *Console) Write(p []byte) (int, error) {
	c.pending <- Escape(string(p))
	return len(p), nil
}

// Escape makes arbitrary text safe to embed in bbcode.
func Escape(s string) string {
	return strings.ReplaceAll(s, "[", "[lb]")
}
