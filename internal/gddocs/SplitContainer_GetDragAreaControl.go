/*
$BarnacleButton.reparent($SplitContainer.get_drag_area_control())
*/

package main

import (
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/SplitContainer"
)

var barnacleButton Button.Instance
var splitContainer SplitContainer.Instance

func SplitContainer_GetDragAreaControl() {
	barnacleButton.AsNode().Reparent(splitContainer.GetDragAreaControl().AsNode())
}
