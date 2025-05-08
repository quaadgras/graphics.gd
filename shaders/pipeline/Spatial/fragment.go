package Spatial

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/int"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/uint"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

type Fragment struct {
	VertexPosition        vec3.XYZ         `gd:"VERTEX"`                   // Position of the vertex, in model space. In world space if world_vertex_coords is used.
	Position              vec3.XYZ         `gd:"POSITION"`                 // If written to, overrides final vertex position in clip space.
	Normal                vec3.XYZ         `gd:"NORMAL"`                   // Normal in model space. In world space if world_vertex_coords is used.
	Tangent               vec3.XYZ         `gd:"TANGENT"`                  // Tangent in model space. In world space if world_vertex_coords is used.
	Binormal              vec3.XYZ         `gd:"BINORMAL"`                 // Binormal in model space. In world space if world_vertex_coords is used.
	UV                    vec2.XY          `gd:"UV"`                       // UV main channel.
	UV2                   vec2.XY          `gd:"UV2"`                      // UV secondary channel.
	Color                 vec4.RGBA        `gd:"COLOR"`                    // Color from vertices.
	Roughness             float.X          `gd:"ROUGHNESS"`                // Roughness for vertex lighting.
	PointSize             float.X          `gd:"POINT_SIZE"`               // Point size for point rendering.
	ModelViewMatrix       mat4.ColumnMajor `gd:"MODEL_VIEW_MATRIX"`        // Model/local space to view space transform (use if possible).
	ModelViewNormalMatrix mat4.ColumnMajor `gd:"MODEL_VIEW_NORMAL_MATRIX"` //
	ProjectionMatrix      mat4.ColumnMajor `gd:"PROJECTION_MATRIX"`        // View space to clip space transform.
}

// ViewportSize returns the size of the viewport (in pixels).
func (frag Fragment) ViewportSize() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("VIEWPORT_SIZE")))
}

// Fragcoord returns the coordinate of pixel center in screen space. xy specifies position in window. Origin is lower left.
// z specifies fragment depth. It is also used as the output value for the fragment depth unless DEPTH is written to.
func (frag Fragment) Fragcoord() vec4.XYZW {
	return gpu.NewVec4Expression(gpu.New(gpu.Identifier("FRAGCOORD")))
}

// FrontFacing returns true if the current face is front facing, false otherwise.
func (frag Fragment) FrontFacing() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("FRONT_FACING")))
}

// View returns the normalized vector from fragment position to camera (in view space). This is the same for both
// perspective and orthogonal cameras.
func (frag Fragment) View() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("VIEW")))
}

// PointCoord returns the point coordinate for drawing points with POINT_SIZE.
func (frag Fragment) PointCoord() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("POINT_COORD")))
}

// ModelMatrix returns the model/local space to world space transform.
func (frag Fragment) ModelMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("MODEL_MATRIX")))
}

// ModelNormalMatrix returns Model/local space to world space transform for normals. This is the same as MODEL_MATRIX by
// default unless the object is scaled non-uniformly, in which case this is set to transpose(inverse(mat3(MODEL_MATRIX))).
func (frag Fragment) ModelNormalMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("MODEL_NORMAL_MATRIX")))
}

// ViewMatrix returns the world space to view space transform.
func (frag Fragment) ViewMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("VIEW_MATRIX")))
}

// InvViewMatrix returns the view space to world space transform.
func (frag Fragment) InvViewMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("INV_VIEW_MATRIX")))
}

// InvProjectionMatrix returns the clip space to view space transform.
func (frag Fragment) InvProjectionMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("INV_PROJECTION_MATRIX")))
}

// NodePositionWorld returns the node position in world space.
func (frag Fragment) NodePositionWorld() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("NODE_POSITION_WORLD")))
}

// NodePositionView returns the node position in view space.
func (frag Fragment) NodePositionView() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("NODE_POSITION_VIEW")))
}

// CameraPositionWorld returns the camera position in world space.
func (frag Fragment) CameraPositionWorld() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("CAMERA_POSITION_WORLD")))
}

// CameraDirectionWorld returns the camera direction in world space.
func (frag Fragment) CameraDirectionWorld() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("CAMERA_DIRECTION_WORLD")))
}

// CameraVisibleLayers returns the cull layers of the camera rendering the current pass.
func (frag Fragment) CameraVisibleLayers() uint.X {
	return gpu.NewUintExpression(gpu.New(gpu.Identifier("CAMERA_VISIBLE_LAYERS")))
}

// ViewIndex returns the view that we are rendering. Used to distinguish between views in multiview/stereo rendering.
// VIEW_MONO_LEFT (0) for Mono (not multiview) or left eye, VIEW_RIGHT (1) for right eye.
func (frag Fragment) ViewIndex() int.X {
	return gpu.NewIntExpression(gpu.New(gpu.Identifier("VIEW_INDEX")))
}

// ViewMonoLeft constant for Mono or left eye, always 0.
func (frag Fragment) ViewMonoLeft() int.X {
	return gpu.NewIntExpression(gpu.New(gpu.Identifier("VIEW_MONO_LEFT")))
}

// ViewRight constant for right eye, always 1.
func (frag Fragment) ViewRight() int.X {
	return gpu.NewIntExpression(gpu.New(gpu.Identifier("VIEW_RIGHT")))
}

// EyeOffset for the eye being rendered. Only applicable for multiview rendering.
func (frag Fragment) EyeOffset() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("EYE_OFFSET")))
}

// ScreenUV returns the screen UV coordinate for current pixel.
func (frag Fragment) ScreenUV() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("SCREEN_UV")))
}

// Backlight color (works like direct light, but it's received even if the normal is slightly facing away from the light).
// If used, backlighting will be applied to the object. Can be used as a cheaper approximation of subsurface scattering.
func (frag Fragment) Backlight() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("BACKLIGHT")))
}
