// Package CanvasItem provides a canvas item shader pipeline used for shading 2D objects.
package CanvasItem

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/int"
	"graphics.gd/shaders/internal"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

/*
Shader used to draw all 2D elements in Godot. These include all nodes that inherit from CanvasItems, and all GUI elements.

CanvasItem shaders contain fewer built-in variables and functionality than Spatial shaders, but they maintain the same basic
structure with vertex, fragment, and light processor functions.
*/
type Shader struct {
	classdb.Extension[internal.Shader, ShaderMaterial.Instance]
}

func (Shader) ShaderType() string       { return "canvas_item" }
func (Shader) RenderMode() []RenderMode { return nil }

func (Shader) Fragment(vertex Vertex) Fragment     { return Fragment{} }
func (Shader) Material(fragment Fragment) Material { return Material{} }
func (Shader) Lighting(material Material) Lighting { return Lighting{} }

type RenderMode string

const (
	BlendMix            RenderMode = "blend_mix"             // Mix blend mode (alpha is transparency), default.
	BlendAdd            RenderMode = "blend_add"             // Additive blend mode.
	BlendSub            RenderMode = "blend_sub"             // Subtractive blend mode.
	BlendMul            RenderMode = "blend_mul"             // Multiplicative blend mode.
	BlendPremulAlpha    RenderMode = "blend_premul_alpha"    // Premultiplied alpha blend mode (fully transparent RenderMode = add, fully opaque RenderMode = mix).
	BlendDisabled       RenderMode = "blend_disabled"        // Disable blending, values (including alpha) are written as-is.
	Unshaded            RenderMode = "unshaded"              // Result is just albedo. No lighting/shading happens in material, making it faster to render.
	LightOnly           RenderMode = "light_only"            // Only draw on light pass.
	SkipVertexTransform RenderMode = "skip_vertex_transform" // VERTEX needs to be transformed manually in the vertex() function.
	WorldVertexCoords   RenderMode = "world_vertex_coords"   // VERTEX is modified in world coordinates instead of local.
)

// Globals are available everywhere, including custom functions.
var (
	Time = gpu.NewFloatExpression(gpu.New(gpu.Identifier("TIME")))
	PI   = gpu.NewFloatExpression(gpu.New(gpu.Identifier("PI")))
	TAU  = gpu.NewFloatExpression(gpu.New(gpu.Identifier("TAU")))
	E    = gpu.NewFloatExpression(gpu.New(gpu.Identifier("E")))
)

type Vertex struct {
	ModelMatrix mat4.ColumnMajor `gd:"MODEL_MATRIX"` // Model/local space to world space transform. World space is the coordinates you normally use in the editor.

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

type Lighting struct {
	Color            vec4.RGBA `gd:"LIGHT"`           // Output color for this Light2D.
	ShadowModulation vec4.RGBA `gd:"SHADOW_MODULATE"` // Multiply shadows cast at this point by this color.
}

/*
SDF provides a few additional functions implemented to sample an automatically generated Signed Distance Field texture.
These functions available for the fragment() and light() functions of CanvasItem shaders. Custom functions may also use
them as long as they called from supported functions.

The signed distance field is generated from LightOccluder2D nodes present in the scene with the SDF Collision property
enabled (which is the default). See the 2D lights and shadows documentation for more information.
*/
type SDF struct{}

// Texture performs an SDF texture lookup.
func (mat SDF) Texture(pos vec2.XY) float.X {
	return gpu.NewFloatExpression(gpu.Fn("texture_sdf", pos))
}

// TextureNormal calculates a normal from the SDF texture.
func (mat SDF) TextureNormal(pos vec2.XY) vec3.XYZ {
	return gpu.NewVec3Expression(gpu.Fn("texture_sdf_normal", pos))
}

// ToScreenUV converts an SDF to screen UV.
func (mat SDF) ToScreenUV(pos vec2.XY) vec2.XY {
	return gpu.NewVec2Expression(gpu.Fn("sdf_to_screen_uv", pos))
}

// FromScreenUV converts screen UV to an SDF.
func (mat SDF) FromScreenUV(uv vec2.XY) vec2.XY {
	return gpu.NewVec2Expression(gpu.Fn("screen_uv_to_sdf", uv))
}
