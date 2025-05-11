package internal

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/ShaderMaterial"
	gd "graphics.gd/internal"
)

type Shader struct {
	ShaderMaterial.Extension[Shader] `gd:"GoShader"`
}

func (shader *Shader) shader() *Shader {
	return shader
}

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		classdb.Register[Shader]()
	})
}

type ShaderPointer interface {
	shader() *Shader
}

func Pointer(shader ShaderPointer) *Shader {
	return shader.shader()
}
