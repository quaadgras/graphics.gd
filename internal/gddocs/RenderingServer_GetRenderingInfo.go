/*
func _ready():
	for _i in 2:
		await get_tree().process_frame

	print(RenderingServer.get_rendering_info(RENDERING_INFO_TOTAL_DRAW_CALLS_IN_FRAME))
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/RenderingServer"
)

func RenderingServer_GetRenderingInfo() {
	if Engine.GetFramesDrawn() == 2 {
		fmt.Println(RenderingServer.GetRenderingInfo(RenderingServer.RenderingInfoTotalDrawCallsInFrame))
	}
}
