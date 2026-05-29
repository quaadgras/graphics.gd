// Package Sky provides a sky shader pipeline used for shading 3D objects.
package Sky

import (
	"fmt"
	"reflect"

	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/internal/gdclass"
	"graphics.gd/shaders"
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

type Shader[T gdclass.Interface] struct {
	ShaderMaterial.Extension[T]
}

func (s *Shader[T]) OnCreate(value reflect.Value) {
	shaders.CompileAny(value.Interface().(shaders.Any))
}

func (*Shader[T]) ShaderType() string       { return "sky" }
func (*Shader[T]) RenderMode() []RenderMode { return nil }
func (*Shader[T]) Pipeline() [3]string {
	return [3]string{"", "", "sky"}
}

func (*Shader[T]) Fragment(state struct{}) struct{} { return struct{}{} }
func (*Shader[T]) Material(state struct{}) Snapshot { return Snapshot{} }
func (*Shader[T]) Lighting(Snapshot) Lighting       { return Lighting{} }

type RenderMode string

const (
	UseHalfResPass    RenderMode = "use_half_res_pass"    // Allows the shader to write to and access the half resolution pass.
	UseQuarterResPass RenderMode = "use_quarter_res_pass" // Allows the shader to write to and access the quarter resolution pass.
	DisableFog        RenderMode = "disable_fog"          // If used, fog will not affect the sky.
)

type Snapshot struct {
	CameraPosition   vec3.XYZ                       `gd:"POSITION"`            // Camera position in world space.
	Radiance         texture.CubeSampler[vec4.RGBA] `gd:"RADIANCE"`            // Radiance cubemap. Can only be read from during background pass. Check !AT_CUBEMAP_PASS before using.
	AtHalfResPass    bool.X                         `gd:"AT_HALF_RES_PASS"`    // True if the shader is being processed at half resolution pass.
	AtQuarterResPass bool.X                         `gd:"AT_QUARTER_RES_PASS"` // True if the shader is being processed at quarter resolution pass.
	AtCubemapPass    bool.X                         `gd:"AT_CUBEMAP_PASS"`     // True if the shader is being processed at cubemap pass.
	EyeDirection     vec3.XYZ                       `gd:"EYE_DIRECTION"`       // Normalized direction of current pixel. Use this as your basic direction for procedural effects.
	ScreenUV         vec2.XY                        `gd:"SCREEN_UV"`           // Screen UV coordinate for current pixel. Used to map a texture to the full screen.
	SkyCoords        vec2.XY                        `gd:"SKY_COORDS"`          // Sphere UV. Used to map a panorama texture to the sky.
	HalfResColor     vec4.RGBA                      `gd:"HALF_RES_COLOR"`      // Color value of corresponding pixel from half resolution pass. Uses linear filter.
	QuarterResColor  vec4.RGBA                      `gd:"QUARTER_RES_COLOR"`   // Color value of corresponding pixel from quarter resolution pass. Uses linear filter.
}

func (Snapshot) Lights() [4]Light {
	var lights [4]Light
	for i := range 4 {
		lights[i] = Light{
			Enabled:   gpu.NewBoolExpression(gpu.New(gpu.Identifier(fmt.Sprintf("LIGHT%d_ENABLED", i+1)))),
			Energy:    gpu.NewFloatExpression(gpu.New(gpu.Identifier(fmt.Sprintf("LIGHT%d_ENERGY", i+1)))),
			Direction: gpu.NewVec3Expression(gpu.New(gpu.Identifier(fmt.Sprintf("LIGHT%d_DIRECTION", i+1)))),
			Color:     gpu.NewRGBExpression(gpu.New(gpu.Identifier(fmt.Sprintf("LIGHT%d_COLOR", i+1)))),
			Size:      gpu.NewFloatExpression(gpu.New(gpu.Identifier(fmt.Sprintf("LIGHT%d_SIZE", i+1)))),
		}
	}
	return lights
}

type Light struct {
	Enabled   bool.X   // true if LIGHTX is visible and in the scene. If false, other light properties may be garbage.
	Energy    float.X  // Energy multiplier for LIGHTX.
	Direction vec3.XYZ // Direction that LIGHTX is facing.
	Color     vec3.RGB // Color of LIGHTX.
	Size      float.X  // Angular diameter of LIGHTX in the sky. Expressed in radians. For reference, the sun from earth is about .0087 radians (0.5 degrees).
}

type Lighting struct {
	Color vec3.RGB  `gd:"COLOR"` // Color of the light.
	Alpha float.X   `gd:"ALPHA"` // Alpha of the light.
	Fog   vec4.RGBA `gd:"FOG"`   // Fog color of the light.
}
