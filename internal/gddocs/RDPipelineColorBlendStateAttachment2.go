/*
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.color_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_SRC_ALPHA
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_SRC_ALPHA
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
*/

package main

import (
	"graphics.gd/classdb/RDPipelineColorBlendStateAttachment"
	"graphics.gd/classdb/Rendering"
)

func ExampleBlendAdd() {
	var attachment = RDPipelineColorBlendStateAttachment.New()
	attachment.SetEnableBlend(true)
	attachment.SetColorBlendOp(Rendering.BlendOpAdd)
	attachment.SetSrcColorBlendFactor(Rendering.BlendFactorOne)
	attachment.SetDstColorBlendFactor(Rendering.BlendFactorOne)
	attachment.SetAlphaBlendOp(Rendering.BlendOpAdd)
	attachment.SetSrcAlphaBlendFactor(Rendering.BlendFactorOne)
	attachment.SetDstAlphaBlendFactor(Rendering.BlendFactorOne)
}
