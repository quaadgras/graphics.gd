// Package Fog provides a fog shader pipeline used for shading 3D objects.
package Fog

import (
	"reflect"

	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/internal/gdclass"
	"graphics.gd/shaders"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/vec3"
)

type Shader[T gdclass.Interface] struct {
	ShaderMaterial.Extension[T]
}

func (s *Shader[T]) OnCreate(value reflect.Value) {
	shaders.CompileAny(value.Interface().(shaders.Any))
}

func (*Shader[T]) ShaderType() string       { return "fog" }
func (*Shader[T]) RenderMode() []RenderMode { return nil }
func (*Shader[T]) Pipeline() [3]string {
	return [3]string{"", "", "fog"}
}

func (*Shader[T]) Fragment(state struct{}) Fragment { return Fragment{} }
func (*Shader[T]) Material(state Fragment) Material { return Material{} }
func (*Shader[T]) Lighting(Material) struct{}       { return struct{}{} }

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
