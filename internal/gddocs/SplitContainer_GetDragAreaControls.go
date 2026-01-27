/*
$BarnacleButton.reparent($SplitContainer.get_drag_area_controls()[0])
*/

package main

func SplitContainer_GetDragAreaControls() {
	barnacleButton.AsNode().Reparent(splitContainer.GetDragAreaControls()[0].AsNode())
}
