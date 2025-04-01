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
)

type MyShader struct {
	shaders.Type2D

	MyUniform vec2.XY `gd:"my_uniform"`
}

// The pipeline functions are named after what they return, not what they accept as
// input.

// Fragment returns a fragment for the given vertex (also known as a vertex shader).
func (MyShader) Fragment(vertex shaders.Vertex2D) shaders.Fragment2D {
	return shaders.Fragment2D{
		Position: vertex.Position,
	}
}

// Material returns a material for the given fragment (also known as a fragment shader).
func (MyShader) Material(fragment shaders.Fragment2D) shaders.Material2D {
	return shaders.Material2D{
		Color: rgba.New(1, 0, 0, 1),
	}
}

// Lighting calculates the lighting for the given material (also known as a lighting pass).
func (MyShader) Lighting(material shaders.Material2D) vec4.RGBA {
	return material.Color
}
```
