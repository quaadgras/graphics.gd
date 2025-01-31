package shaders

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/int"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec4"
)

type CanvasItemVertexAttributes = Vertex2D

type Vertex2D struct {
	Global

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

type Fragment2D struct {
	Position vec2.XY                      `gd:"VERTEX"`  // Vertex, in local space.
	UV       vec2.XY                      `gd:"UV"`      // Normalized texture coordinates. Range from 0 to 1.
	Texture  texture.Sampler2D[vec4.RGBA] `gd:"TEXTURE"` // Texture sampler.
}

type Material2D struct {
	Color vec4.RGBA `gd:"COLOR"` // Color from vertex primitive.
}
