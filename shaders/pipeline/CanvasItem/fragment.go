package CanvasItem

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

type Fragment struct {
	SDF SDF

	Position  vec2.XY   `gd:"VERTEX"`     // Vertex, in local space.
	UV        vec2.XY   `gd:"UV"`         // Normalized texture coordinates. Range from 0 to 1.
	Color     vec4.RGBA `gd:"COLOR"`      // Color from vertex primitive.
	PointSize float.X   `gd:"POINT_SIZE"` // Point size for point drawing.
}

// Fragcoord returns the coordinate of pixel center. In screen space. xy specifies position in viewport.
// Upper-left of the viewport is the origin, (0.0, 0.0).
func (frag Fragment) Fragcoord() vec4.XYZW {
	return gpu.NewVec4Expression(gpu.New(gpu.Identifier("FRAGCOORD")))
}

// ScreenPixelSize of individual pixels. Equal to inverse of resolution.
func (frag Fragment) ScreenPixelSize() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("SCREEN_PIXEL_SIZE")))
}

// PointCoord for drawing points.
func (frag Fragment) PointCoord() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("POINT_COORD")))
}

// Texture returns the default 2D texture.
func (frag Fragment) Texture() texture.Sampler2D[vec4.RGBA] {
	var sampler texture.Sampler2D[vec4.RGBA]
	gpu.Set(&sampler, gpu.New(gpu.Identifier("TEXTURE")))
	return sampler
}

// TexturePixelSize returns the normalized pixel size of default 2D texture.
// For a Sprite2D with a texture of size 64x32px, TEXTURE_PIXEL_SIZE = vec2(1/64, 1/32)
func (frag Fragment) TexturePixelSize() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("TEXTURE_PIXEL_SIZE")))
}

// AtLightPass returns true if the fragment is being processed at light pass.
func (frag Fragment) AtLightPass() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("AT_LIGHT_PASS")))
}

// SpecularShininessTexture of this object.
func (frag Fragment) SpecularShininessTexture() texture.Sampler2D[vec4.RGBA] {
	var sampler texture.Sampler2D[vec4.RGBA]
	gpu.Set(&sampler, gpu.New(gpu.Identifier("SPECULAR_SHININESS_TEXTURE")))
	return sampler
}

// SpecularShininess returns the specular shininess color, as sampled from the texture.
func (frag Fragment) SpecularShininess() vec4.RGBA {
	return gpu.NewRGBAExpression(gpu.New(gpu.Identifier("SPECULAR_SHININESS")))
}

// ScreenUV returns the screen UV coordinate for current pixel.
func (frag Fragment) ScreenUV() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("SCREEN_UV")))
}

// Normal read from NormalTexture.
func (frag Fragment) Normal() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("NORMAL")))
}

// NormalTexture returns the default 2D normal texture.
func (frag Fragment) NormalTexture() texture.Sampler2D[vec4.XYZW] {
	var sampler texture.Sampler2D[vec4.XYZW]
	gpu.Set(&sampler, gpu.New(gpu.Identifier("NORMAL_TEXTURE")))
	return sampler
}
