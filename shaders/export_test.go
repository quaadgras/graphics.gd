package shaders

import "reflect"

// Generate returns the code that CompileAny would compile.
func Generate(val Any) string {
	rtype := reflect.TypeOf(val)
	fragment, _ := rtype.MethodByName("Fragment")
	material, _ := rtype.MethodByName("Material")
	lighting, _ := rtype.MethodByName("Lighting")
	return generate(val, fragment.Type.In(1), material.Type.In(1), lighting.Type.In(1), lighting.Type.Out(0))
}
