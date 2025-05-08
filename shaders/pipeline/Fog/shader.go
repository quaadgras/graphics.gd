// Package Fog provides a fog shader pipeline used for shading 3D objects.
package Fog

import (
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/internal"
	"graphics.gd/shaders/vec3"
)

type Shader struct {
	shader
}

type shader = internal.Shader

func (Shader) ShaderType() string       { return "fog" }
func (Shader) RenderMode() []RenderMode { return nil }
func (Shader) Pipeline() [3]string {
	return [3]string{"", "", "fog"}
}

func (Shader) Fragment(state struct{}) Fragment { return Fragment{} }
func (Shader) Material(state Fragment) Material { return Material{} }
func (Shader) Lighting(Material) struct{}       { return struct{}{} }

type RenderMode string

type Fragment struct {
	WorldPosition  vec3.XYZ `gd:"WORLD_POSITION"`  // Position of current froxel cell in world space.
	ObjectPosition vec3.XYZ `gd:"OBJECT_POSITION"` // Position of the center of the current FogVolume in world space.
	UVW            vec3.XYZ `gd:"UVW"`             // 3-dimensional UV, used to map a 3D texture to the current FogVolume.
	Size           vec3.XYZ `gd:"SIZE"`            // Size of the current FogVolume when its shape has a size.
	SDF            vec3.XYZ `gd:"SDF"`             // Signed distance field to the surface of the FogVolume. Negative if inside volume, positive otherwise.
}

type Material struct {
	Albedo   vec3.RGB `gd:"ALBEDO"`   // Output base color value, interacts with light to produce final color. Only written to fog volume if used.
	Density  float.X  `gd:"DENSITY"`  // Output density value. Can be negative to allow subtracting one volume from another. Density must be used for fog shader to write anything at all.
	Emission vec3.RGB `gd:"EMISSION"` // Output emission color value, added to color during light pass to produce final color. Only written to fog volume if used.
}
