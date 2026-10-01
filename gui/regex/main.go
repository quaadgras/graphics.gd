package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/HBoxContainer"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/RegEx"
	"graphics.gd/classdb/TextEdit"
	"graphics.gd/classdb/VBoxContainer"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
)

// Regex is a playground for testing regular expressions: it lists every
// match of the expression (and its capture groups) found in the text.
type Regex struct {
	VBoxContainer.Extension[Regex] `gd:"Regex"`

	Text          TextEdit.Instance      `gd:"%Text"`
	Expression    LineEdit.Instance      `gd:"%Expression"`
	List          VBoxContainer.Instance `gd:"%List"`
	HBoxContainer HBoxContainer.Instance

	regex RegEx.Instance
}

func (r *Regex) Ready() {
	r.regex = RegEx.New()
	r.Text.SetText(`They asked me "What's going on \"in the manor\"?"`)
	r.UpdateExpression(r.Expression.Text())
}

func (r *Regex) UpdateExpression(text string) {
	r.regex.Compile(text)
	r.UpdateText()
}

func (r *Regex) addLabel(text string, modulate Color.RGBA) {
	label := Label.New()
	label.SetText(text)
	label.AsCanvasItem().SetModulate(modulate)
	r.List.AsNode().AddChild(label.AsNode())
}

func (r *Regex) UpdateText() {
	for _, child := range r.List.AsNode().GetChildren() {
		child.QueueFree()
	}
	if r.regex.IsValid() {
		r.HBoxContainer.AsCanvasItem().SetModulate(Color.RGBA{R: 1, G: 1, B: 1, A: 1})
		for matchNumber, regexMatch := range r.regex.SearchAll(r.Text.Text()) {
			// List all matches and their respective captures.
			r.addLabel(fmt.Sprintf("RegEx match #%d:", matchNumber+1), Color.RGBA{R: 0.6, G: 0.9, B: 1, A: 1})
			for captureNumber, result := range regexMatch.Strings() {
				r.addLabel(fmt.Sprintf("    Capture group #%d: %s", captureNumber+1, result),
					Color.RGBA{R: 1, G: 1, B: 1, A: 1})
			}
		}
	} else {
		r.HBoxContainer.AsCanvasItem().SetModulate(Color.RGBA{R: 1, G: 0.2, B: 0.1, A: 1})
		r.addLabel("Error: Invalid regular expression. Check if the expression is correctly escaped and terminated.",
			Color.RGBA{R: 1, G: 1, B: 1, A: 1})
	}
}

func (r *Regex) OnHelpMetaClicked(_ any) {
	// Workaround for clickable link doing nothing when clicked.
	OS.ShellOpen("https://regexr.com")
}

func main() {
	classdb.Register[Regex]()
	startup.Scene()
}
