package shaders

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/int"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

type Vertex2D struct {
	Globals

	// Local space to world space transform. World space is the coordinates you normally use in the editor.
	ModelMatrix mat4.ColumnMajor `gd:"MODEL_MATRIX"`
	// World space to canvas space transform. In canvas space the origin is the upper-left corner of the
	// screen and coordinates ranging from (0, 0) to viewport size.
	CanvasMatrix mat4.ColumnMajor `gd:"CANVAS_MATRIX"`
	// Canvas space to clip space. In clip space coordinates ranging from (-1, -1) to (1, 1).
	ScreenMatrix mat4.ColumnMajor `gd:"SCREEN_MATRIX"`

	InstanceID     int.X     `gd:"INSTANCE_ID"`     // InstanceID for instancing.
	InstanceCustom vec4.XYZW `gd:"INSTANCE_CUSTOM"` // InstanceCustom data.
	AtLightPass    bool.X    `gd:"AT_LIGHT_PASS"`   // Always false.

	// Normalized pixel size of default 2D texture. For a Sprite2D with a texture of size 64x32px,
	// TEXTURE_PIXEL_SIZE = vec2(1/64, 1/32)
	TexturePixelSize vec2.XY `gd:"TEXTURE_PIXEL_SIZE"`

	Position vec2.XY `gd:"VERTEX"`    // Vertex, in local space.
	ID       float.X `gd:"VERTEX_ID"` // The index of the current vertex in the vertex buffer.

	UV        vec2.XY   `gd:"UV"`         // Normalized texture coordinates. Range from 0 to 1.
	Color     vec4.RGBA `gd:"COLOR"`      // Color from vertex primitive.
	PointSize float.X   `gd:"POINT_SIZE"` // Point size for point drawing.
	Custom0   vec4.XYZW `gd:"CUSTOM0"`    // Custom value from vertex primitive.
	Custom1   vec4.XYZW `gd:"CUSTOM1"`    // Custom value from vertex primitive.
}

type FragmentReadOnly2D struct {
	Globals

	Pixel                    vec4.XYZW                    `gd:"FRAGCOORD"`                  // Coordinate of pixel center. In screen space. xy specifies position in window. Origin is upper-left.
	PixelSize                vec2.XY                      `gd:"SCREEN_PIXEL_SIZE"`          // Size of individual pixels. Equal to inverse of resolution.
	Point                    vec2.XY                      `gd:"POINT_COORD"`                // Coordinate for drawing points.
	Texture                  texture.Sampler2D[vec4.RGBA] `gd:"TEXTURE"`                    // Default 2D texture.
	TexturePixelSize         vec2.XY                      `gd:"TEXTURE_PIXEL_SIZE"`         // Normalized pixel size of default 2D texture. For a Sprite2D with a texture of size 64x32px, TexturePixelSize = vec2(1/64, 1/32)
	AtLightPass              bool.X                       `gd:"AT_LIGHT_PASS"`              // Always false.
	SpecularShininessTexture texture.Sampler2D[vec4.RGBA] `gd:"SPECULAR_SHININESS_TEXTURE"` // Specular shininess texture of this object.
	SpecularShininess        vec4.XYZW                    `gd:"SPECULAR_SHININESS"`         // Specular shininess color, as sampled from the texture.
	UV                       vec2.XY                      `gd:"UV"`                         // UV from vertex function.
	ScreenUV                 vec2.XY                      `gd:"SCREEN_UV"`                  // Screen UV coordinate for current pixel.
	Normal                   vec3.XYZ                     `gd:"NORMAL"`                     // Normal from vertex function.
	NormalTexture            texture.Sampler2D[vec4.RGBA] `gd:"NORMAL_TEXTURE"`             // Default 2D normal texture.
}

type Fragment2D struct {
	FragmentReadOnly2D

	Position  vec2.XY   `gd:"VERTEX"`     // Vertex, in local space.
	UV        vec2.XY   `gd:"UV"`         // Normalized texture coordinates. Range from 0 to 1.
	Color     vec4.RGBA `gd:"COLOR"`      // Color from vertex primitive.
	PointSize float.X   `gd:"POINT_SIZE"` // Point size for point drawing.
}

type MaterialReadOnly2D struct {
	Globals
	Light2D

	Pixel             vec4.XYZW                    `gd:"FRAGCOORD"`          // Coordinate of pixel center. In screen space. xy specifies position in window. Origin is upper-left.
	UV                vec2.XY                      `gd:"UV"`                 // UV from vertex function.
	Texture           texture.Sampler2D[vec4.RGBA] `gd:"TEXTURE"`            // Current texture in use for CanvasItem.
	TexturePixelSize  vec2.XY                      `gd:"TEXTURE_PIXEL_SIZE"` // Normalized pixel size of default 2D texture. For a Sprite2D with a texture of size 64x32px, TexturePixelSize = vec2(1/64, 1/32)
	ScreenUV          vec2.XY                      `gd:"SCREEN_UV"`          // Screen UV coordinate for current pixel.
	Point             vec2.XY                      `gd:"POINT_COORD"`        // Coordinate for drawing points.
	SpecularShininess vec4.XYZW                    `gd:"SPECULAR_SHININESS"` // Specular shininess, as set in the object's texture.
}

type Light2D struct {
	Color         vec4.RGBA `gd:"LIGHT_COLOR"`          // Color of Light multiplied by Light's texture.
	Energy        float.X   `gd:"LIGHT_ENERGY"`         // Energy multiplier of Light.
	Position      vec3.XYZ  `gd:"LIGHT_POSITION"`       // Position of Light in screen space. If using a DirectionalLight2D this is always vec3(0,0,0).
	Direction     vec3.XYZ  `gd:"LIGHT_DIRECTION"`      // Direction of Light in screen space.
	IsDirectional bool.X    `gd:"LIGHT_IS_DIRECTIONAL"` // true if this pass is a DirectionalLight2D.
}

type Material2D struct {
	MaterialReadOnly2D

	SDF MaterialSDF

	Position            vec2.XY   `gd:"VERTEX"`        // Pixel position in screen space.
	PositionForShadows  vec2.XY   `gd:"SHADOW_VERTEX"` // Same as Position but can be written to alter shadows.
	PositionForLighting vec3.XYZ  `gd:"LIGHT_VERTEX"`  // Same as Position but can be written to alter lighting. Z component represents height.
	Normal              vec3.XYZ  `gd:"NORMAL"`        // Normal from vertex function.
	Color               vec4.RGBA `gd:"COLOR"`         // Color from vertex primitive.
}

type MaterialSDF struct{}

func (mat MaterialSDF) Texture(pos vec2.XY) float.X {
	return gpu.NewFloatExpression(gpu.Fn("texture_sdf", pos))
}
func (mat MaterialSDF) TextureNormal(pos vec2.XY) vec3.XYZ {
	return gpu.NewVec3Expression(gpu.Fn("texture_sdf_normal", pos))
}
func (mat Material2D) ToScreenUV(pos vec2.XY) vec2.XY {
	return gpu.NewVec2Expression(gpu.Fn("sdf_to_screen_uv", pos))
}
func (mat Material2D) FromScreenUV(uv vec2.XY) vec2.XY {
	return gpu.NewVec2Expression(gpu.Fn("screen_uv_to_sdf", uv))
}
