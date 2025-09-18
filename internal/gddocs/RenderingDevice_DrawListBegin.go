/*
var rd = RenderingDevice.new()
var clear_colors = PackedColorArray([Color(0, 0, 0, 0), Color(0, 0, 0, 0), Color(0, 0, 0, 0)])
var draw_list = rd.draw_list_begin(framebuffers[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS)

# Draw opaque.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)
# Draw wire.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline_wire)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)

rd.draw_list_end()
*/

package main

import (
	"graphics.gd/classdb/Rendering"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
)

var framebuffers []RID.Framebuffer
var raster_pipeline RID.RenderPipeline
var raster_base_uniform RID.UniformSet
var raster_push_constant []byte
var slice_triangle_count []int
var i int

func RenderingDevice_DrawListBegin() {
	var clear_colors = [3]Color.RGBA{}
	var draw_list = rd.MoreArgs().DrawListBegin(framebuffers[i], Rendering.DrawClearColorAll, clear_colors[:], 0, 0, Rect2.PositionSize{}, 0)

	rd.DrawListBindRenderPipeline(draw_list, raster_pipeline)
	rd.DrawListBindUniformSet(draw_list, raster_base_uniform, 0)
	rd.DrawListSetPushConstant(draw_list, raster_push_constant, len(raster_push_constant))
	rd.MoreArgs().DrawListDraw(draw_list, false, 1, slice_triangle_count[i]*3)

	rd.DrawListEnd()

}
