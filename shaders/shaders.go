/*
Package shaders provides a ShaderMaterial.Instance with the shader pipeline written within Go.

Shaders are multi-stage programs executed on the GPU. They can be used to gain precise control
over rendering calculations, such as lighting, shadows, and post-processing effects.

To create a new 2D shader in Go, define a struct that embeds shaders.Type2D and implements the
pipeline methods you would like to overide. For example:

	type MyShader struct {
		CanvasItem.Shader
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

Each sub-package provides GPU-specific shader types that can be used within a shader pipeline.
Keep in mind that the Go code is compiled to run on the GPU, so non-GPU values, function
calls or branches will only take affect during compilation and not when rendering.

All for loops will be unrolled. The shaders package does not currently support non-constant
loops.

# Uniforms

Uniforms are added as fields to the shader struct. They can be written with [Set] and read with
[Get]. Uniforms wrapped inside the [PerInstance] generic type need to be accessed through the
RenderingServer/GeometryInstance3D packages.

	type MyShader struct {
		CanvasItem.Shader

		MyUniform vec2.XY `gd:"my_uniform"`

		Color shaders.PerInstance[vec4.XYZW] `gd:"color"`
	}

	var shader = new(MyShader)
	shaders.Compile(&shader)
	shaders.Set(&shader.MyUniform, Vector2.New(1, 2))
*/
package shaders

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Shader"
	"graphics.gd/classdb/ShaderMaterial"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/shaders/internal/gpu"
	dsl "graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/vec2"
	"graphics.gd/variant/String"
)

// Globals are available everywhere, including custom functions.
var (
	Time = gpu.NewFloatExpression(gpu.New(gpu.Identifier("TIME")))
	PI   = gpu.NewFloatExpression(gpu.New(gpu.Identifier("PI")))
	TAU  = gpu.NewFloatExpression(gpu.New(gpu.Identifier("TAU")))
	E    = gpu.NewFloatExpression(gpu.New(gpu.Identifier("E")))
)

// PerInstance uniforms can be added inside a shader struct.
type PerInstance[T gpu.Evaluator] struct {
	value T
}

// Value returns the value of the uniform for the current instance.
func (p PerInstance[T]) Value() T { return p.value }

// Set sets the value of a uniform.
func Set[T gpu.EquivalentTo[G], G any](uniform *T, value G) {
	gpu.Shader(gpu.EquivalentTo[G](*uniform)).AsShaderMaterial().SetShaderParameter(string(gpu.Evaluate(gpu.EquivalentTo[G](*uniform)).(gpu.Identifier)), value)
}

// Get gets the value of a uniform.
func Get[T gpu.EquivalentTo[G], G any](uniform *T) G {
	rvalue, err := gd.ConvertToDesiredGoType(gpu.Shader(gpu.EquivalentTo[G](*uniform)).AsShaderMaterial().GetShaderParameter(string(gpu.Evaluate(gpu.EquivalentTo[G](*uniform)).(gpu.Identifier))), reflect.TypeFor[G]())
	if err != nil {
		Engine.Raise(err)
	}
	return rvalue.Interface().(G)
}

type Program[Vertex, Fragment, Material, Lighting any, RenderMode ~string] interface {
	gdclass.Pointer

	ShaderMaterial.Any

	ShaderType() string
	RenderMode() []RenderMode

	Pipeline() [3]string

	Fragment(Vertex) Fragment
	Material(Fragment) Material
	Lighting(Material) Lighting
}

type Any interface {
	ShaderMaterial.Any

	ShaderType() string
	Pipeline() [3]string
}

func Compile[V, F, M, L comparable, RM ~string](prog Program[V, F, M, L, RM]) {
	compile(prog, reflect.TypeFor[V](), reflect.TypeFor[F](), reflect.TypeFor[M](), reflect.TypeFor[L]())
}

func CompileAny(val Any) {
	rtype := reflect.TypeOf(val)
	fragment, _ := rtype.MethodByName("Fragment")
	material, _ := rtype.MethodByName("Material")
	lighting, _ := rtype.MethodByName("Lighting")
	compile(val, fragment.Type.In(1), material.Type.In(1), lighting.Type.In(1), lighting.Type.Out(0))
}

func compile(prog Any, v, f, m, l reflect.Type) {
	super := prog.AsShaderMaterial()
	shader := Shader.New()
	writer := strings.Builder{}
	fmt.Fprintf(&writer, "// Code generated by graphics.gd/shaders DO NOT EDIT!\n")
	fmt.Fprintf(&writer, "shader_type %s;\n\n", prog.ShaderType())

	rvalue := reflect.ValueOf(prog)

	render_modes := rvalue.MethodByName("RenderMode").Call(nil)[0]
	if render_modes.Len() > 0 {
		fmt.Fprintf(&writer, "render_mode ")
		for i := range render_modes.Len() {
			mode := render_modes.Index(i).String()
			if i > 0 {
				fmt.Fprintf(&writer, ", ")
			}
			fmt.Fprintf(&writer, "%s", mode)
		}
		fmt.Fprintf(&writer, ";\n")
	}

	pipeline := prog.Pipeline()

	var vertices = reflect.New(v).Elem()
	linkup(vertices.Addr().Interface())
	var fragment = reflect.New(f).Elem()
	linkup(fragment.Addr().Interface())
	var material = reflect.New(m).Elem()
	linkup(material.Addr().Interface())

	compileUniforms(&writer, prog)
	if frag := rvalue.MethodByName("Fragment").Call([]reflect.Value{vertices}); !frag[0].IsZero() && pipeline[0] != "" {
		compileFunction(&writer, frag[0].Interface(), pipeline[0])
	}
	if matl := rvalue.MethodByName("Material").Call([]reflect.Value{fragment}); !matl[0].IsZero() && pipeline[1] != "" {
		compileFunction(&writer, matl[0].Interface(), pipeline[1])
	}
	if lght := rvalue.MethodByName("Lighting").Call([]reflect.Value{material}); !lght[0].IsZero() && pipeline[2] != "" {
		compileFunction(&writer, lght[0].Interface(), pipeline[2])
	}
	shader.SetCode(writer.String())
	super.SetShader(shader)
}

func linkup(in any) {
	value := reflect.ValueOf(in).Elem()
	rtype := value.Type()
	for i := range rtype.NumField() {
		if value.Field(i).Kind() == reflect.Struct && rtype.Field(i).IsExported() {
			linkup(value.Field(i).Addr().Interface())
		}
		if tag := rtype.Field(i).Tag.Get("gd"); tag != "" {
			field := value.Field(i)
			switch ptr := field.Addr().Interface().(type) {
			case *vec2.XY:
				dsl.Set(&ptr.X, dsl.Identifier(tag+".x"))
				dsl.Set(&ptr.Y, dsl.Identifier(tag+".y"))
			}
			dsl.Set(value.Field(i).Addr().Interface().(dsl.Pointer), dsl.Identifier(tag))
		}
	}
}

func compileUniforms(w io.Writer, prog Any) {
	value := reflect.ValueOf(prog).Elem()
	rtype := value.Type()
	for i := range rtype.NumField() {
		field := rtype.Field(i)
		if field.Name == "Shader" {
			continue
		}
		name := String.ToSnakeCase(field.Name)
		options := ""
		if tag := field.Tag.Get("gd"); tag != "" {
			tag, options, _ = strings.Cut(tag, ",")
			if tag != "" {
				name = tag
			}
		}
		fmt.Fprintf(w, "uniform %s %s", glslTypeFor(field.Type), name)
		dsl.Set(value.Field(i).Addr().Interface().(dsl.Pointer), gpu.Uniform(name, prog))
		if options != "" {
			fmt.Fprintf(w, ": ")
			var first bool = true
			for option := range strings.SplitSeq(options, ",") {
				if first {
					first = false
				} else {
					fmt.Fprintf(w, ",")
				}
				fmt.Fprintf(w, " %s", option)
			}
		}
		fmt.Fprintf(w, ";\n")
	}
	fmt.Fprintln(w)
}

func glslTypeFor(t reflect.Type) string {
	switch {
	case t.ConvertibleTo(reflect.TypeFor[gpu.Bool]()):
		return "bool"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec2b]()):
		return "bvec2"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec3b]()):
		return "bvec3"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec4b]()):
		return "bvec4"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Float]()):
		return "float"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Int]()):
		return "int"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec2i]()):
		return "ivec2"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec3i]()):
		return "ivec3"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
		return "ivec4"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Mat2]()):
		return "mat2"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Mat3]()):
		return "mat3"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Mat4]()):
		return "mat4"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec2]()):
		return "vec2"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec3]()):
		return "vec3"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()):
		return "vec4"
	case t.ConvertibleTo(reflect.TypeFor[gpu.RGB]()):
		return "vec3"
	case t.ConvertibleTo(reflect.TypeFor[gpu.RGBA]()):
		return "vec4"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Uint]()):
		return "uint"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec2u]()):
		return "uvec2"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec3u]()):
		return "uvec3"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
		return "uvec4"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec2]()):
		return "vec2"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec3]()):
		return "vec3"
	case t.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()):
		return "vec4"
	case t.Implements(reflect.TypeFor[gpu.IsSampler2D]()):
		elem := gpu.SamplerType(reflect.Zero(t).Interface().(gpu.IsSampler2D))
		switch {
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()):
			return "sampler2D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
			return "isampler2D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
			return "usampler2D"
		}
	case t.Implements(reflect.TypeFor[gpu.IsSampler3D]()):
		elem := gpu.SamplerType(reflect.Zero(t).Interface().(gpu.IsSampler3D))
		switch {
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()):
			return "sampler3D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
			return "isampler3D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
			return "usampler3D"
		}
	case t.Implements(reflect.TypeFor[gpu.IsArraySampler2D]()):
		elem := gpu.SamplerType(reflect.Zero(t).Interface().(gpu.IsArraySampler2D))
		switch {
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()):
			return "sampler2DArray"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
			return "isampler2DArray"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
			return "usampler2DArray"
		}
	case t.Implements(reflect.TypeFor[gpu.IsCubeSampler]()):
		return "samplerCube"
	}
	panic(fmt.Sprintf("unsupported GPU type %s", t))
}

func compileFunction(w io.Writer, data any, name string) {
	fmt.Fprintf(w, "void %s() {\n", name)
	value := reflect.ValueOf(data)
	rtype := value.Type()
	for i := range rtype.NumField() {
		field := rtype.Field(i)
		expr, ok := value.Field(i).Interface().(dsl.Evaluator)
		if ok && !value.Field(i).IsZero() {
			fmt.Fprintf(w, "\t%s = ", field.Tag.Get("gd"))
			compileExpression(w, expr)
			fmt.Fprintf(w, ";\n")
		}
	}
	fmt.Fprintf(w, "}\n")
}

func compileExpression(w io.Writer, expression dsl.Evaluator) {
	if expr := dsl.Evaluate(expression); expr != nil {
		expression = expr
	}
	rtype := reflect.TypeOf(expression)
	switch {
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.RGBA]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.RGBA]()).Interface().(gpu.RGBA)
		compileCall(w, "vec4", value.R, value.G, value.B, value.A)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4]()).Interface().(gpu.Vec4)
		compileCall(w, "vec4", value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4i]()).Interface().(gpu.Vec4i)
		compileCall(w, "ivec4", value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4u]()).Interface().(gpu.Vec4u)
		compileCall(w, "uvec4", value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4b]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4b]()).Interface().(gpu.Vec4b)
		compileCall(w, "bvec4", value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3]()).Interface().(gpu.Vec3)
		compileCall(w, "vec3", value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.RGB]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.RGB]()).Interface().(gpu.RGB)
		compileCall(w, "vec3", value.R, value.G, value.B)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3i]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3i]()).Interface().(gpu.Vec3i)
		compileCall(w, "ivec3", value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3u]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3u]()).Interface().(gpu.Vec3u)
		compileCall(w, "uvec3", value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3b]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3b]()).Interface().(gpu.Vec3b)
		compileCall(w, "bvec3", value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2]()).Interface().(gpu.Vec2)
		compileCall(w, "vec2", value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2i]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2i]()).Interface().(gpu.Vec2i)
		compileCall(w, "ivec2", value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2u]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2u]()).Interface().(gpu.Vec2u)
		compileCall(w, "uvec2", value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2b]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2b]()).Interface().(gpu.Vec2b)
		compileCall(w, "bvec2", value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Float]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Float]()).Interface().(gpu.Float)
		fmt.Fprintf(w, "%f", value.X)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Int]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Int]()).Interface().(gpu.Int)
		fmt.Fprintf(w, "%d", value.X)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Uint]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Uint]()).Interface().(gpu.Uint)
		fmt.Fprintf(w, "%d", value.X)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Bool]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Bool]()).Interface().(gpu.Bool)
		fmt.Fprintf(w, "%t", value.X)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Mat2]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Mat2]()).Interface().(gpu.Mat2)
		compileCall(w, "mat2", value.Columns[0][0], value.Columns[0][1], value.Columns[1][0], value.Columns[1][1])
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Mat3]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Mat3]()).Interface().(gpu.Mat3)
		compileCall(w, "mat3", value.Columns[0][0], value.Columns[0][1], value.Columns[0][2], value.Columns[1][0], value.Columns[1][1], value.Columns[1][2], value.Columns[2][0], value.Columns[2][1], value.Columns[2][2])
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Mat4]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Mat4]()).Interface().(gpu.Mat4)
		compileCall(w, "mat4", value.Columns[0][0], value.Columns[0][1], value.Columns[0][2], value.Columns[0][3], value.Columns[1][0], value.Columns[1][1], value.Columns[1][2], value.Columns[1][3], value.Columns[2][0], value.Columns[2][1], value.Columns[2][2], value.Columns[2][3], value.Columns[3][0], value.Columns[3][1], value.Columns[3][2], value.Columns[3][3])
	default:
		switch value := expression.(type) {
		case dsl.Operation:
			fmt.Fprintf(w, "(")
			compileExpression(w, value.A)
			fmt.Fprintf(w, " %s ", value.Op)
			compileExpression(w, value.B)
			fmt.Fprintf(w, ")")
		case dsl.Identifier:
			fmt.Fprintf(w, "%s", value)
		case dsl.Ternary:
			fmt.Fprintf(w, "(")
			compileExpression(w, value.If)
			fmt.Fprintf(w, " ? ")
			compileExpression(w, value.A)
			fmt.Fprintf(w, " : ")
			compileExpression(w, value.B)
			fmt.Fprintf(w, ")")
		case dsl.FunctionCall:
			compileCall(w, string(value.Name), value.Args...)
		default:
			panic(fmt.Sprintf("unsupported expression type %T", expression))
		}
	}
}

func compileCall(w io.Writer, name string, args ...dsl.Evaluator) {
	fmt.Fprintf(w, "%s(", name)
	for i, arg := range args {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		compileExpression(w, arg)
	}
	fmt.Fprintf(w, ")")
}
