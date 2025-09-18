/*
func _ready():
	for _i in 2:
		await get_tree().process_frame

	print(
			RenderingServer.viewport_get_render_info(get_viewport().get_viewport_rid(),
			RenderingServer.VIEWPORT_RENDER_INFO_TYPE_VISIBLE,
			RenderingServer.VIEWPORT_RENDER_INFO_DRAW_CALLS_IN_FRAME)
	)
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/Viewport"
)

func RenderingServer_ViewportGetRenderInfo() {
	if Engine.GetFramesDrawn() == 2 {
		fmt.Println(RenderingServer.ViewportGetRenderInfo(Viewport.Get(node).GetViewportRid(),
			RenderingServer.ViewportRenderInfoTypeVisible,
			RenderingServer.ViewportRenderInfoDrawCallsInFrame))
	}
}
