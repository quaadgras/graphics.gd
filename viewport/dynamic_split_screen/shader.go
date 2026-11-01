package main

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/swizzle"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

// SplitScreen shows the view of the first player, or when the players are far
// apart, splits the screen between both views along the line perpendicular to
// the players. Ported from split_screen.gdshader.
type SplitScreen struct {
	CanvasItem.Shader[SplitScreen]

	ViewportSize       vec2.XY                      // Size in pixels of the viewport.
	Viewport1          texture.Sampler2D[vec4.RGBA] `gd:"viewport1,source_color"`
	Viewport2          texture.Sampler2D[vec4.RGBA] `gd:"viewport2,source_color"`
	SplitActive        bool.X                       // true: split screen, false: use view1.
	Player1Position    vec2.XY                      `gd:"player1_position"`        // Position of player 1 in UV coordinates.
	Player2Position    vec2.XY                      `gd:"player2_position"`        // Position of player 2 in UV coordinates.
	SplitLineThickness float.X                      `gd:",hint_range(0, 10, 0.1)"` // Width of the split border.
	SplitLineColor     vec3.RGB                     `gd:",source_color"`           // Color of the split border.
}

func (*SplitScreen) RenderMode() []CanvasItem.RenderMode {
	return []CanvasItem.RenderMode{CanvasItem.Unshaded}
}

// distanceToLine returns the distance from point to the line through p1 and p2.
func distanceToLine(p1, p2, point vec2.XY) float.X {
	a := float.Sub(p1.Y, p2.Y)
	b := float.Sub(p2.X, p1.X)
	cross := float.Sub(float.Add(float.Add(float.Mul(a, point.X), float.Mul(b, point.Y)), float.Mul(p1.X, p2.Y)), float.Mul(p2.X, p1.Y))
	return float.Div(float.Abs(cross), float.Sqrt(float.Add(float.Mul(a, a), float.Mul(b, b))))
}

func (s *SplitScreen) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	view1 := swizzle.RGB(s.Viewport1.Sample(frag.UV))
	view2 := swizzle.RGB(s.Viewport2.Sample(frag.UV))
	width, height := s.ViewportSize.X, s.ViewportSize.Y
	p1 := s.Player1Position

	dx := vec2.Sub(s.Player2Position, p1)
	// A high value (vertical split) when dx.y is 0.
	slope := bool.Select(float.Neq(dx.Y, 0.0), float.Div(dx.X, dx.Y), float.New(100000.0))

	// splitY is the height of the split line at x, in UV coordinates.
	splitY := func(x float.X) float.X {
		return float.Add(float.Mul(float.Sub(0.5, x), slope), 0.5)
	}
	lineStart := vec2.New(0.0, float.Mul(height, splitY(float.New(0.0))))
	lineEnd := vec2.New(width, float.Mul(height, splitY(float.New(1.0))))
	player1Y := splitY(p1.X)

	below := float.Gt(frag.UV.Y, splitY(frag.UV.X))
	color := bool.Select(below,
		bool.Select(float.Gt(p1.Y, player1Y), view1, view2),
		bool.Select(float.Lt(p1.Y, player1Y), view1, view2))

	distance := distanceToLine(lineStart, lineEnd, vec2.Mul(frag.UV, s.ViewportSize))
	onLine := float.Lt(distance, s.SplitLineThickness)
	color = bool.Select(onLine, rgb.Mix(s.SplitLineColor, color, float.Div(distance, s.SplitLineThickness)), color)

	color = bool.Select(s.SplitActive, color, view1)
	return CanvasItem.Material{
		Color: rgba.New(color.R, color.G, color.B, 1.0),
	}
}
