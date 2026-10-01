package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
)

type RichTextBBCode struct {
	Control.Extension[RichTextBBCode] `gd:"RichTextBBCode"`
}

func (r *RichTextBBCode) OnRichTextLabelMetaClicked(meta any) {
	link := fmt.Sprint(meta)
	if err := OS.ShellOpen(link); err == nil {
		fmt.Printf("Opened link '%s' successfully!\n", link)
	} else {
		fmt.Printf("Failed opening the link '%s'!\n", link)
	}
}

func (r *RichTextBBCode) OnPauseToggled(buttonPressed bool) {
	SceneTree.Get(r.AsNode()).SetPaused(buttonPressed)
}

func main() {
	classdb.Register[RichTextBBCode]()
	startup.Scene()
}
