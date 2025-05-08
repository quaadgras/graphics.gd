package internal

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/ShaderMaterial"
	gd "graphics.gd/internal"
)

type Shader struct {
	classdb.Extension[Shader, ShaderMaterial.Instance] `gd:"GoShader"`
}

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		classdb.Register[Shader]()
	})
}
