/*
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.color_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
*/

package main

import (
	"graphics.gd/classdb/RDPipelineColorBlendStateAttachment"
	"graphics.gd/classdb/Rendering"
)

func ExampleBlendPremultipliedAlpha() {
	var attachment = RDPipelineColorBlendStateAttachment.New()
	attachment.SetEnableBlend(true)
	attachment.SetColorBlendOp(Rendering.BlendOpAdd)
	attachment.SetSrcColorBlendFactor(Rendering.BlendFactorOne)
	attachment.SetDstColorBlendFactor(Rendering.BlendFactorOneMinusSrcAlpha)
	attachment.SetAlphaBlendOp(Rendering.BlendOpAdd)
	attachment.SetSrcAlphaBlendFactor(Rendering.BlendFactorOne)
	attachment.SetDstAlphaBlendFactor(Rendering.BlendFactorOneMinusSrcAlpha)
}
