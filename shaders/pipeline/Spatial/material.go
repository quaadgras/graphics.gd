package Spatial

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

type Material struct {
	LightVertexPosition    vec3.XYZ  `gd:"LIGHT_VERTEX"`             // A writable version of VERTEX that can be used to alter light and shadows. Writing to this will not change the position of the fragment.
	Depth                  float.X   `gd:"DEPTH"`                    // Custom depth value (range of [0.0, 1.0]). If DEPTH is being written to in any shader branch, then you are responsible for setting the DEPTH for all other branches. Otherwise, the graphics API will leave them uninitialized.
	Normal                 vec3.XYZ  `gd:"NORMAL"`                   // Normal that comes from the vertex() function, in view space. If skip_vertex_transform is enabled, it may not be in view space.
	Tangent                vec3.XYZ  `gd:"TANGENT"`                  // Tangent that comes from the vertex() function, in view space. If skip_vertex_transform is enabled, it may not be in view space.
	Binormal               vec3.XYZ  `gd:"BINORMAL"`                 // Binormal that comes from the vertex() function, in view space. If skip_vertex_transform is enabled, it may not be in view space.
	NormalMap              vec3.XYZ  `gd:"NORMAL_MAP"`               // Set normal here if reading normal from a texture instead of NORMAL.
	NormalMapDepth         float.X   `gd:"NORMAL_MAP_DEPTH"`         // Depth from NORMAL_MAP. Defaults to 1.0.
	Albedo                 vec4.RGBA `gd:"ALBEDO"`                   // Albedo (default white). Base color.
	Alpha                  float.X   `gd:"ALPHA"`                    // Alpha (range of [0.0, 1.0]). If read from or written to, the material will go to the transparent pipeline.
	AlphaScizzorThreshold  float.X   `gd:"ALPHA_SCISSOR_THRESHOLD"`  // If written to, values below a certain amount of alpha are discarded.
	AlphaHashScale         float.X   `gd:"ALPHA_HASH_SCALE"`         // Alpha hash scale when using the alpha hash transparency mode. Defaults to 1.0. Higher values result in more visible pixels in the dithering pattern.
	AlphaAntialiasingEdge  float.X   `gd:"ALPHA_ANTIALIASING_EDGE"`  // The threshold below which alpha to coverage antialiasing should be used. Defaults to 0.0. Requires the alpha_to_coverage render mode. Should be set to a value lower than ALPHA_SCISSOR_THRESHOLD to be effective.
	AlphaTextureCoordinate vec2.XY   `gd:"ALPHA_TEXTURE_COORDINATE"` // The texture coordinate to use for alpha-to-coverge antialiasing. Requires the alpha_to_coverage render mode. Typically set to UV * vec2(albedo_texture_size) where albedo_texture_size is the size of the albedo texture in pixels.
	PremultiplyAlphaFactor float.X   `gd:"PREMUL_ALPHA_FACTOR"`      // Premultiplied alpha factor. Only effective if render_mode blend_premul_alpha; is used. This should be written to when using a shaded material with premultiplied alpha blending for interaction with lighting. This is not required for unshaded materials.
	Metallic               float.X   `gd:"METALLIC"`                 // Metallic (range of [0.0, 1.0]).
	Specular               float.X   `gd:"SPECULAR"`                 // Specular (not physically accurate to change). Defaults to 0.5. 0.0 disables reflections.
	Roughness              float.X   `gd:"ROUGHNESS"`                // Roughness (range of [0.0, 1.0]).
	Rim                    float.X   `gd:"RIM"`                      // Rim (range of [0.0, 1.0]). If used, Godot calculates rim lighting. Rim size depends on ROUGHNESS.
	RimTint                float.X   `gd:"RIM_TINT"`                 // Rim Tint, range of 0.0 (white) to 1.0 (albedo). If used, Godot calculates rim lighting.
	ClearCoat              float.X   `gd:"CLEARCOAT"`                // Small specular blob added on top of the existing one. If used, Godot calculates clearcoat.
	ClearCoatGloss         float.X   `gd:"CLEARCOAT_GLOSS"`          // Gloss of clearcoat. If used, Godot calculates clearcoat.
	Anisotropy             float.X   `gd:"ANISOTROPY"`               // For distorting the specular blob according to tangent space.
	AnisotropyFlow         vec2.XY   `gd:"ANISOTROPY_FLOW"`          // Distortion direction, use with flowmaps.
	SubsurfaceScattering   SubsurfaceScattering
	Backlight              float.X `gd:"BACKLIGHT"` // Color of backlighting (works like direct light, but it's received even if the normal is slightly facing away from the light). If used, backlighting will be applied to the object. Can be used as a cheaper approximation of subsurface scattering.
	AmbientOcclusion       AmbientOcclusion
	Emission               vec3.RGB  `gd:"EMISSION"`   // Emission color (can go over (1.0, 1.0, 1.0) for HDR).
	Fog                    vec4.RGBA `gd:"FOG"`        // If written to, blends final pixel color with FOG.rgb based on FOG.a.
	Radiance               vec4.RGBA `gd:"RADIANCE"`   // If written to, blends environment map radiance with RADIANCE.rgb based on RADIANCE.a.
	Irradiance             vec4.RGBA `gd:"IRRADIANCE"` // If written to, blends environment map irradiance with IRRADIANCE.rgb based on IRRADIANCE.a.
}

type AmbientOcclusion struct {
	Strength    float.X `gd:"AO"`       // Strength of ambient occlusion. For use with pre-baked AO.
	LightEffect float.X `gd:"AO_LIGHT"` // How much ambient occlusion affects direct light (range of [0.0, 1.0], default 0.0).
}

type SubsurfaceScattering struct {
	Strength           float.X   `gd:"SUBSURFACE_SCATTERING"` // Strength of subsurface scattering. If used, subsurface scattering will be applied to the object.
	TransmittanceColor vec4.RGBA `gd:"TRANSMITTANCE_COLOR"`   // Color of subsurface scattering transmittance. If used, subsurface scattering transmittance will be applied to the object.
	TransmittanceDepth float.X   `gd:"TRANSMITTANCE_DEPTH"`   // Depth of subsurface scattering transmittance. Higher values allow the effect to reach deeper into the object.
	TransmittanceBoost float.X   `gd:"TRANSMITTANCE_BOOST"`   // Boosts the subsurface scattering transmittance if set above 0.0. This makes the effect show up even on directly lit surfaces
}

// ViewportSize returns the size of the viewport (in pixels).
func (mat Material) ViewportSize() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("VIEWPORT_SIZE")))
}

// Fragcoord returns the coordinate of pixel center in screen space. xy specifies position in window. Origin is lower left.
// z specifies fragment depth. It is also used as the output value for the fragment depth unless DEPTH is written to.
func (mat Material) Fragcoord() vec4.XYZW {
	return gpu.NewVec4Expression(gpu.New(gpu.Identifier("FRAGCOORD")))
}

// ModelMatrix returns the model/local space to world space transform.
func (mat Material) ModelMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("MODEL_MATRIX")))
}

// InverseModelMatrix returns the inverse of the model/local space to world space transform.
func (mat Material) InverseModelMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("INV_MODEL_MATRIX")))
}

// ViewMatrix returns the world space to view space transform.
func (mat Material) ViewMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("VIEW_MATRIX")))
}

// ProjectionMatrix returns the view space to clip space transform.
func (mat Material) ProjectionMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("PROJECTION_MATRIX")))
}

// InverseProjectionMatrix returns the clip space to view space transform.
func (mat Material) InverseProjectionMatrix() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("INV_PROJECTION_MATRIX")))
}

// ScreenUV returns the screen UV coordinate for current pixel.
func (mat Material) ScreenUV() vec2.XY {
	return gpu.NewVec2Expression(gpu.New(gpu.Identifier("SCREEN_UV")))
}

// View returns the normalized vector from fragment position to camera (in view space). This is the same for both
// perspective and orthogonal cameras.
func (mat Material) View() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("VIEW")))
}

// Light vector, in view space.
func (mat Material) Light() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("LIGHT")))
}

// LightColor multiplied by light energy multiplied by PI. The PI multiplication is present because
// physically-based lighting models include a division by PI.
func (mat Material) LightColor() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("LIGHT_COLOR")))
}

// SpecularAmount for OmniLight3D and SpotLight3D, 2.0 multiplied by light_specular. For DirectionalLight3D, 1.0.
func (mat Material) SpecularAmount() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("SPECULAR_AMOUNT")))
}

// LightIsDirectional returns true if this pass is a DirectionalLight3D.
func (mat Material) LightIsDirectional() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("LIGHT_IS_DIRECTIONAL")))
}

// Attenuation returns the attenuation based on distance or shadow.
func (mat Material) Attenuation() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("ATTENUATION")))
}
