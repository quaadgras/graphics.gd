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

type Material struct {
	SDF SDF

	Normal             vec3.XYZ  `gd:"NORMAL"`           // Normal vector, in local space.
	NormalMap          vec3.XYZ  `gd:"NORMAL_MAP"`       // Configures normal maps meant for 3D for use in 2D. If used, overrides NORMAL.
	NormalMapDepth     float.X   `gd:"NORMAL_MAP_DEPTH"` // Normal map depth for scaling.
	PixelPosition      vec2.XY   `gd:"VERTEX"`           // Pixel position in screen space.
	ShadowPosition     vec2.XY   `gd:"SHADOW_VERTEX"`    // Same as VERTEX but can be written to alter shadows.
	LightPixelPosition vec3.XYZ  `gd:"LIGHT_VERTEX"`     // Same as VERTEX but can be written to alter lighting. Z component represents height.
	Color              vec4.RGBA `gd:"COLOR"`            // COLOR from the vertex() function multiplied by the TEXTURE color. Also output color value.
}

// Fragcoord returns the coordinate of pixel center in screen space. xy specifies position in window. Origin is lower left.
// z specifies fragment depth. It is also used as the output value for the fragment depth unless DEPTH is written to.
func (mat Material) Fragcoord() vec4.XYZW {
	return gpu.NewVec4Expression(gpu.New(gpu.Identifier("FRAGCOORD")))
}

// UV from the vertex() function, equivalent to the UV in the fragment() function.
func (mat Material) UV() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("UV")))
}

// Texture returns the default 2D texture.
func (mat Material) Texture() texture.Sampler2D[vec4.RGBA] {
	var sampler texture.Sampler2D[vec4.RGBA]
	gpu.Set(&sampler, gpu.New(gpu.Identifier("TEXTURE")))
	return sampler
}

// TexturePixelSize returns the normalized pixel size of default 2D texture.
// For a Sprite2D with a texture of size 64x32px, TEXTURE_PIXEL_SIZE = vec2(1/64, 1/32)
func (mat Material) TexturePixelSize() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("TEXTURE_PIXEL_SIZE")))
}

// ScreenUV returns the screen UV coordinate for current pixel.
func (mat Material) ScreenUV() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("SCREEN_UV")))
}

// PointCoord returns the point coordinate for drawing points with POINT_SIZE.
func (mat Material) PointCoord() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("POINT_COORD")))
}

// LightColor returns the color of the light, multiplied by the light's texture.
func (mat Material) LightColor() vec4.RGBA {
	return gpu.NewRGBAExpression(gpu.New(gpu.Identifier("LIGHT_COLOR")))
}

// LightEnergy energy multiplier of the Light2D.
func (mat Material) LightEnergy() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("LIGHT_ENERGY")))
}

// LightPosition in screen space. If using a DirectionalLight2D this is always (0.0, 0.0, 0.0).
func (mat Material) LightPosition() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("LIGHT_POSITION")))
}

// LightDirection in screen space.
func (mat Material) LightDirection() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("LIGHT_DIRECTION")))
}

// LightIsDirectional returns true if this pass is a DirectionalLight2D.
func (mat Material) LightIsDirectional() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("LIGHT_IS_DIRECTIONAL")))
}

// Light output color for this Light2D.
func (mat Material) Light() vec4.RGBA {
	return gpu.NewRGBAExpression(gpu.New(gpu.Identifier("LIGHT")))
}

// SpecularShininess as set in the object's texture.
func (mat Material) SpecularShininess() vec4.XYZW {
	return gpu.NewVec4Expression(gpu.New(gpu.Identifier("SPECULAR_SHININESS")))
}
