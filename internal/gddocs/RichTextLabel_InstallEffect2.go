/*
# rich_text_label.gd
extends RichTextLabel

func _ready():
	install_effect(MyCustomEffect.new())

	# Alternatively, if not using `class_name` in the script that extends RichTextEffect:
	install_effect(preload("res://effect.gd").new())
*/

package main

import (
	"graphics.gd/classdb/RichTextEffect"
	"graphics.gd/classdb/RichTextLabel"
)

type richTextLabelInstallEffect struct {
	RichTextLabel.Extension[richTextLabelInstallEffect]

	Effect RichTextEffect.Instance
}

func (n richTextLabelInstallEffect) Ready() {
	n.AsRichTextLabel().InstallEffect(n.Effect)
}
