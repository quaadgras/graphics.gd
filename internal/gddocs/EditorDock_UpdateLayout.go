/*
func _update_layout(layout):
	box_container.vertical = (layout == DOCK_LAYOUT_VERTICAL)
*/

package main

import (
	"graphics.gd/classdb/BoxContainer"
	"graphics.gd/classdb/EditorDock"
)

type editorDockUpdateLayout struct {
	EditorDock.Extension[editorDockUpdateLayout]

	BoxContainer BoxContainer.Instance
}

func (n editorDockUpdateLayout) UpdateLayout(layout EditorDock.DockLayout) {
	n.BoxContainer.SetVertical(layout == EditorDock.DockLayoutVertical)
}
