/*
shader_type canvas_item;
render_mode unshaded;

uniform sampler2D screen_texture : hint_screen_texture, repeat_disable, filter_nearest;

void fragment() {
    vec4 c = textureLod(screen_texture, SCREEN_UV, 0.0);

    if (c.a > 0.0001) {
        c.rgb /= c.a;
    }

    COLOR *= c;
}
*/

package main

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec4"
)

type CanvasGroupShader struct {
	CanvasItem.Shader[CanvasGroupShader]

	ScreenTexture texture.Sampler2D[vec4.RGBA]
}

func (sl *CanvasGroupShader) RenderMode() []CanvasItem.RenderMode {
	return []CanvasItem.RenderMode{CanvasItem.Unshaded}
}

func (sl *CanvasGroupShader) Fragment(vertex CanvasItem.Vertex) CanvasItem.Fragment {
	var c = sl.ScreenTexture.SampleLOD(vertex.ScreenUV, float.New(0))
	var threshold = float.Gt(c.A, float.New(0.0001))
	c.R = bool.Mix(c.R, float.Div(c.R, c.A), threshold)
	c.G = bool.Mix(c.G, float.Div(c.G, c.A), threshold)
	c.B = bool.Mix(c.B, float.Div(c.B, c.A), threshold)
	return CanvasItem.Fragment{
		Color: rgba.Mul(vertex.Color, c),
	}
}
