/*
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.color_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_DST_COLOR
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ZERO
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_DST_ALPHA
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ZERO
*/

package main

import (
	"graphics.gd/classdb/RDPipelineColorBlendStateAttachment"
	"graphics.gd/classdb/Rendering"
)

func ExampleBlendMultiply() {
	var attachment = RDPipelineColorBlendStateAttachment.New()
	attachment.SetEnableBlend(true)
	attachment.SetColorBlendOp(Rendering.BlendOpAdd)
	attachment.SetSrcColorBlendFactor(Rendering.BlendFactorDstColor)
	attachment.SetDstColorBlendFactor(Rendering.BlendFactorZero)
	attachment.SetAlphaBlendOp(Rendering.BlendOpAdd)
	attachment.SetSrcAlphaBlendFactor(Rendering.BlendFactorDstAlpha)
	attachment.SetDstAlphaBlendFactor(Rendering.BlendFactorZero)
}
