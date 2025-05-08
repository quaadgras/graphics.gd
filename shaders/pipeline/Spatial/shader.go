// Package Spatial provides the spatial shader pipeline used for shading 3D objects.
package Spatial

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/int"
	"graphics.gd/shaders/internal"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/uint"
	"graphics.gd/shaders/uvec4"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

/*
Shader used for shading 3D objects. They are the most complex type of shader that the engine offers. Spatial
shaders are highly configurable with different render modes and different rendering options (e.g. Subsurface
Scattering, Transmission, Ambient Occlusion, Rim lighting etc). Users can optionally write vertex, fragment,
and light processor functions to affect how objects are drawn.
*/
type Shader struct {
	classdb.Extension[internal.Shader, ShaderMaterial.Instance]
}

func (Shader) ShaderType() string       { return "spatial" }
func (Shader) RenderMode() []RenderMode { return nil }

func (Shader) Fragment(vertex Vertex) Fragment     { return Fragment{} }
func (Shader) Material(fragment Fragment) Material { return Material{} }
func (Shader) Lighting(material Material) Lighting { return Lighting{} }

type RenderMode string

const (
	BlendMix                     RenderMode = "blend_mix"                 // Mix blend mode (alpha is transparency), default.
	BlendAdd                     RenderMode = "blend_add"                 // Additive blend mode.
	BlendSub                     RenderMode = "blend_sub"                 // Subtractive blend mode.
	BlendMul                     RenderMode = "blend_mul"                 // Multiplicative blend mode.
	BlendPremulAlpha             RenderMode = "blend_premul_alpha"        // Premultiplied alpha blend mode (fully transparent RenderMode = add, fully opaque RenderMode = mix).
	DepthDrawOpaque              RenderMode = "depth_draw_opaque"         // Only draw depth for opaque geometry (not transparent).
	DepthDrawAlways              RenderMode = "depth_draw_always"         // Always draw depth (opaque and transparent).
	DepthDrawNever               RenderMode = "depth_draw_never"          // Never draw depth.
	DepthPrepassAlpha            RenderMode = "depth_prepass_alpha"       // Do opaque depth pre-pass for transparent geometry.
	DepthTestDisabled            RenderMode = "depth_test_disabled"       // Disable depth testing.
	SubSurfaceScatteringModeSkin RenderMode = "sss_mode_skin"             // Subsurface Scattering mode for skin (optimizes visuals for human skin, e.g. boosted red channel).
	CullBack                     RenderMode = "cull_back"                 // Cull back-faces (default).
	CullFront                    RenderMode = "cull_front"                // Cull front-faces.
	CullDisabled                 RenderMode = "cull_disabled"             // Culling disabled (double sided).
	Unshaded                     RenderMode = "unshaded"                  // Result is just albedo. No lighting/shading happens in material, making it faster to render.
	Wireframe                    RenderMode = "wireframe"                 // Geometry draws using lines (useful for troubleshooting).
	DebugShadowSplits            RenderMode = "debug_shadow_splits"       // Directional shadows are drawn using different colors for each split (useful for troubleshooting).
	DiffuseBurley                RenderMode = "diffuse_burley"            // Burley (Disney PBS) for diffuse (default).
	DiffuseLambert               RenderMode = "diffuse_lambert"           // Lambert shading for diffuse.
	DiffuseLambertWrap           RenderMode = "diffuse_lambert_wrap"      // Lambert-wrap shading (roughness-dependent) for diffuse.
	DiffuseToon                  RenderMode = "diffuse_toon"              // Toon shading for diffuse.
	SpecularSchlickGGX           RenderMode = "specular_schlick_ggx"      // Schlick-GGX for direct light specular lobes (default).
	SpecularToon                 RenderMode = "specular_toon"             // Toon for direct light specular lobes.
	SpecularDisabled             RenderMode = "specular_disabled"         // Disable direct light specular lobes. Doesn't affect reflected light (use SPECULAR RenderMode = 0.0 instead).
	SkipVertexTransform          RenderMode = "skip_vertex_transform"     // VERTEX, NORMAL, TANGENT, and BITANGENT need to be transformed manually in the vertex() function.
	WorldVertexCoords            RenderMode = "world_vertex_coords"       // VERTEX, NORMAL, TANGENT, and BITANGENT are modified in world space instead of model space.
	EnsureCorrectNormals         RenderMode = "ensure_correct_normals"    // Use when non-uniform scale is applied to mesh (note: currently unimplemented).
	ShadowsDisabled              RenderMode = "shadows_disabled"          // Disable computing shadows in shader. The shader will not cast shadows, but can still receive them.
	AmbientLightDisabled         RenderMode = "ambient_light_disabled"    // Disable contribution from ambient light and radiance map.
	ShadowToOpacity              RenderMode = "shadow_to_opacity"         // Lighting modifies the alpha so shadowed areas are opaque and non-shadowed areas are transparent. Useful for overlaying shadows onto a camera feed in AR.
	VertexLighting               RenderMode = "vertex_lighting"           // Use vertex-based lighting instead of per-pixel lighting.
	ParticleTrails               RenderMode = "particle_trails"           // Enables the trails when used on particles geometry.
	AlphaToCoverage              RenderMode = "alpha_to_coverage"         // Alpha antialiasing mode, see here for more.
	AlphaToCoverageAndOne        RenderMode = "alpha_to_coverage_and_one" // Alpha antialiasing mode, see here for more.
	FogDisabled                  RenderMode = "fog_disabled"              // Disable receiving depth-based or volumetric fog. Useful for blend_add materials like particles.
)

// Globals are available everywhere, including custom functions.
var (
	Time         = gpu.NewFloatExpression(gpu.New(gpu.Identifier("TIME")))
	PI           = gpu.NewFloatExpression(gpu.New(gpu.Identifier("PI")))
	TAU          = gpu.NewFloatExpression(gpu.New(gpu.Identifier("TAU")))
	E            = gpu.NewFloatExpression(gpu.New(gpu.Identifier("E")))
	OutputIsSRGB = gpu.NewBoolExpression(gpu.New(gpu.Identifier("OUTPUT_IS_SRGB")))
	ClipSpaceFar = gpu.NewFloatExpression(gpu.New(gpu.Identifier("CLIP_SPACE_FAR")))
)

type Vertex struct {
	ViewportSize          vec2.XY          `gd:"VIEWPORT_SIZE"`            // Size of viewport (in pixels).
	ViewMatrix            mat4.ColumnMajor `gd:"VIEW_MATRIX"`              // View matrix. Transforms from world space to camera space.
	InvViewMatrix         mat4.ColumnMajor `gd:"INV_VIEW_MATRIX"`          // View space to world space transform.
	MainCamInvViewMatrix  mat4.ColumnMajor `gd:"MAIN_CAM_INV_VIEW_MATRIX"` // View space to world space transform of camera used to draw the current viewport.
	InvProjectionMatrix   mat4.ColumnMajor `gd:"INV_PROJECTION_MATRIX"`    // Clip space to view space transform.
	NodePositionWorld     vec3.XYZ         `gd:"NODE_POSITION_WORLD"`      // Node position in world space.
	NodePositionView      vec3.XYZ         `gd:"NODE_POSITION_VIEW"`       // Node position in view space.
	CameraPositionWorld   vec3.XYZ         `gd:"CAMERA_POSITION_WORLD"`    // Camera position in world space.
	CameraPositionView    vec3.XYZ         `gd:"CAMERA_POSITION_VIEW"`     // Camera position in view space.
	CameraDirectionWorld  vec3.XYZ         `gd:"CAMERA_DIRECTION_WORLD"`   // Camera direction in world space.
	CameraVisibleLayers   uint.X           `gd:"CAMERA_VISIBLE_LAYERS"`    // Camera visible layers.
	InstanceID            int.X            `gd:"INSTANCE_ID"`              // InstanceID for instancing.
	InstanceCustom        vec4.XYZW        `gd:"INSTANCE_CUSTOM"`          // Instance custom data (for particles, mostly).
	ViewIndex             int.X            `gd:"VIEW_INDEX"`               // The view that we are rendering. VIEW_MONO_LEFT (0) for Mono (not multiview) or left eye, VIEW_RIGHT (1) for right eye.
	ViewMonoLeft          int.X            `gd:"VIEW_MONO_LEFT"`           // Constant for Mono or left eye, always 0.
	ViewRight             int.X            `gd:"VIEW_RIGHT"`               // Constant for right eye, always 1.
	EyeOffset             vec3.XYZ         `gd:"EYE_OFFSET"`               // Position offset for the eye being rendered. Only applicable for multiview rendering.
	Position              vec3.XYZ         `gd:"VERTEX"`                   // Position of the vertex, in model space. In world space if world_vertex_coords is used.
	VertexID              int.X            `gd:"VERTEX_ID"`                // The index of the current vertex in the vertex buffer.
	Normal                vec3.XYZ         `gd:"NORMAL"`                   // Normal in model space. In world space if world_vertex_coords is used.
	Tangent               vec3.XYZ         `gd:"TANGENT"`                  // Tangent in model space. In world space if world_vertex_coords is used.
	Binormal              vec3.XYZ         `gd:"BINORMAL"`                 // Binormal in model space. In world space if world_vertex_coords is used.
	UV                    vec2.XY          `gd:"UV"`                       // UV main channel.
	UV2                   vec2.XY          `gd:"UV2"`                      // UV secondary channel.
	Color                 vec4.RGBA        `gd:"COLOR"`                    // Color from vertices.
	PointSize             float.X          `gd:"POINT_SIZE"`               // Point size for point rendering.
	ModelViewMatrix       mat4.ColumnMajor `gd:"MODEL_VIEW_MATRIX"`        // Model/local space to view space transform (use if possible).
	ModelViewNormalMatrix mat4.ColumnMajor `gd:"MODEL_VIEW_NORMAL_MATRIX"` //
	ModelMatrix           mat4.ColumnMajor `gd:"MODEL_MATRIX"`             // Model/local space to world space transform.
	ModelNormalMatrix     mat4.ColumnMajor `gd:"MODEL_NORMAL_MATRIX"`      //
	ProjectionMatrix      mat4.ColumnMajor `gd:"PROJECTION_MATRIX"`        // View space to clip space transform.
	BoneIndices           uvec4.XYZW       `gd:"BONE_INDICES"`             //
	BoneWeights           vec4.XYZW        `gd:"BONE_WEIGHTS"`             //
	Custom0               vec4.XYZW        `gd:"CUSTOM0"`                  //
	Custom1               vec4.XYZW        `gd:"CUSTOM1"`                  //
	Custom2               vec4.XYZW        `gd:"CUSTOM2"`                  //
	Custom3               vec4.XYZW        `gd:"CUSTOM3"`                  //
}

type Lighting struct {
	Diffuse  vec3.RGB `gd:"DIFFUSE"`  // Diffuse light result.
	Specular vec3.RGB `gd:"SPECULAR"` // Specular light result.
	Alpha    float.X  `gd:"ALPHA"`    // Alpha (range of [0.0, 1.0]). If written to, the material will go to the transparent pipeline.
}
