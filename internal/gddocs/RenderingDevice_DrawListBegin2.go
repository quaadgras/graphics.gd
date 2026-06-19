/*
rd.draw_list_begin(fb[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS | 5)
*/

package main

import (
	"graphics.gd/classdb/Rendering"
	"graphics.gd/classdb/RenderingDevice"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
)

// Note: this doc snippet predates the 4.7 draw_flags-based draw_list_begin signature;
// translated to the current API.
func ExampleDrawListBegin(rd RenderingDevice.Instance, fb []RID.Framebuffer, clearColors []Color.RGBA) {
	rd.MoreArgs().DrawListBegin(fb[0], Rendering.DrawClearColorAll, clearColors, 1.0, 0, Rect2.PositionSize{}, 0)
}
