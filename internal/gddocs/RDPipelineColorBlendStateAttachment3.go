/*
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_REVERSE_SUBTRACT
attachment.color_blend_op = RenderingDevice.BLEND_OP_REVERSE_SUBTRACT
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

func ExampleBlendSubtract() {
	var attachment = RDPipelineColorBlendStateAttachment.New()
	attachment.SetEnableBlend(true)
	attachment.SetColorBlendOp(Rendering.BlendOpReverseSubtract)
	attachment.SetSrcColorBlendFactor(Rendering.BlendFactorOne)
	attachment.SetDstColorBlendFactor(Rendering.BlendFactorOne)
	attachment.SetAlphaBlendOp(Rendering.BlendOpReverseSubtract)
	attachment.SetSrcAlphaBlendFactor(Rendering.BlendFactorOne)
	attachment.SetDstAlphaBlendFactor(Rendering.BlendFactorOne)
}
