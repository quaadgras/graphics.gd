# Write Shaders In Go

Shaders can be written by calling shader-specific functions defined by the
packages in this directory. The shader is then run once in Go which essentially
records itself as an AST which is compiled down into Godot's GLSL variant.

This sort of approach is often referred to as a language-hosted DSL. It's well
suited for branchless shaders as any Go branches or side effects will only be
evaluated at shader 'compile time'.

The key benefits are the type safety, composition, language familiarity and IDE
integration that result from this.

## Example:

```go
package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec4"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/pipeline/CanvasItem"
)

type MyShader struct {
	CanvasItem.Shader[MyShader]
}

// The pipeline functions are named after what they return, not what they accept as
// input.

// Fragment runs for each point/vertex of the shape/mesh being rendered, should return
// fragment parameters for each point (also known as a vertex shader).
func (MyShader) Fragment(vertex CanvasItem.Vertex) CanvasItem.Fragment {
	return CanvasItem.Fragment{
		Position: vertex.Position,
	}
}

// Material runs for each pixel on each face of the shape being rendered, should return
// the surface parameters for each pixel (also known as a fragment shader). The input
// fragment is a blend of each contributing vertex point.
func (MyShader) Material(fragment CanvasItem.Fragment) CanvasItem.Material {
	return CanvasItem.Material{
		Color: rgba.New(1, 0, 0, 1),
	}
}

// Lighting runs for each light, per pixel for each face of the shape being rendered, should
// return the final color for each pixel (also known as a lighting pass).
func (MyShader) Lighting(material CanvasItem.Material) CanvasItem.Lighting {
	return CanvasItem.Lighting{
		Color: material.Color,
	}
}
```

## Branches and Loops

Call evaluate on a gpu bool to branch, avoid branching where possible, as it
can be expensive.

```go
color := s.Background
if float.Gt(density, 0.0).Evaluate() {
	color = s.lighting(density) // only computed for pixels inside of a cloud.
}
```

A `for` statement over `over.Int` loops on the GPU. `continue` and `break` work
as expected.

```go
pos, density := start, float.New(0.0)
for range over.Int(steps) {
	pos = vec3.Add(pos, step)
	density = float.Add(density, s.density(pos))
	if float.Gt(density, 1.0).Evaluate() {
		break
	}
}
```

## Uniforms

Uniforms are added as fields to the shader struct. They can be written with [Set] and read with
[Get]. Uniforms wrapped inside the [PerInstance] generic type need to be accessed through the
RenderingServer/GeometryInstance3D packages.

```go
	type MyShader struct {
		CanvasItem.Shader[MyShader]

		MyUniform vec2.XY `gd:"my_uniform"`

		Color shaders.PerInstance[vec4.XYZW] `gd:"color"`
	}

	ms.Color.Value()

	var shader = new(MyShader)
	shaders.Set(&shader.MyUniform, Vector2.New(1, 2))
```

## Varyings

The value returned by `Fragment` is what `Material` receives, blended between
the vertices of each face. Fields tagged with a built-in (such as `UV` or
`COLOR`) use that built-in. To pass a value of your own, return a struct that
embeds the pipeline's `Fragment` and adds untagged fields, then accept that same
struct in `Material`. Each untagged field becomes a varying. In the same way,
untagged fields on the value returned by `Material` are passed to `Lighting`.

```go
type SmokeFragment struct {
	Spatial.Fragment

	Voronoi float.X // becomes: varying float voronoi;
}

func (s *SmokeBall) Fragment(vertex Spatial.Vertex) SmokeFragment {
	voronoi := float.X(s.VoronoiSampler.Sample(vertex.UV).X)
	return SmokeFragment{
		Fragment: Spatial.Fragment{
			VertexPosition: vec3.Add(vertex.Position, vec3.Mul(vertex.Normal, voronoi)),
		},
		Voronoi: voronoi,
	}
}

func (s *SmokeBall) Material(frag SmokeFragment) Spatial.Material {
	return Spatial.Material{
		Alpha: float.SmoothStep(0.4, 0.6, frag.Voronoi),
	}
}
```

Integer varyings are declared `flat`, as integers cannot be blended.
