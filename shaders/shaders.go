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

Go for loops are unrolled. Go if statements are decided while compiling, so to branch on a
value that is only known on the GPU, use [bool.X.Evaluate], and to loop for a number of times
that is only known on the GPU, range over [graphics.gd/shaders/over.Int].

# Branches and Loops

Evaluate returns which way a GPU condition goes, so that a Go if statement can branch on it.
The pipeline function is run once for each way that its conditions can go, and the values that
differ between the runs are computed inside of an if statement on the GPU. Assigning to Go
variables inside of the branches works as usual:

	color := s.Background
	if float.Gt(density, 0.0).Evaluate() {
		color = s.lighting(density) // only computed for pixels inside of a cloud.
	}

Each condition doubles the number of runs that follow it, so to choose between values that are
cheap to compute, use [bool.Select], which computes both of them.

A Go for statement over over.Int loops on the GPU. The loop body is written like any other Go
loop body: values assigned to variables declared before the loop carry over to the next iteration,
values computed before the loop are computed once, outside of it, and continue and break (inside
of a branch on a condition evaluated with Evaluate) skip the rest of the iteration and stop the loop:

	pos, density := start, float.New(0.0)
	for range over.Int(steps) {
		pos = vec3.Add(pos, step)
		density = float.Add(density, s.density(pos))
		if float.Gt(density, 1.0).Evaluate() {
			break
		}
	}

# Varyings

The value returned by Fragment is what Material receives, blended between the vertices of
each face. Fields tagged with a built-in (such as UV or COLOR) use that built-in. To pass a
value of your own, return a struct that embeds the pipeline's Fragment and adds untagged
fields, then accept that same struct in Material. Each untagged field becomes a varying.
In the same way, untagged fields on the value returned by Material are passed to Lighting.

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

# Uniforms

Uniforms are added as fields to the shader struct. They can be written with [Set] and read with
[Get]. Uniforms wrapped inside the [PerInstance] generic type need to be accessed through the
RenderingServer/GeometryInstance3D packages.

	type MyShader struct {
		CanvasItem.Shader[MyShader]

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
	"maps"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"

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
	CompileAny(prog)
}

// CompileAny compiles the shader of a material, the code for each shader class
// is generated once and shared by all of its materials. It starts being
// generated in the background when the class is registered.
func CompileAny(val Any) {
	p := programFor(reflect.TypeOf(val))
	p.generate(val)
	if p.failure != nil {
		panic(p.failure)
	}
	// Each material binds its own uniforms, so that Set and Get act on it.
	generating.Lock()
	compileUniforms(io.Discard, val)
	generating.Unlock()
	shader := Shader.New()
	shader.SetCode(p.code)
	val.AsShaderMaterial().SetShader(shader)
}

// program is the code generated for a shader class.
type program struct {
	once    sync.Once
	code    string
	failure any // what the generation panicked with.
}

var (
	programs sync.Map // reflect.Type → *program

	// generating is held while code is generated, as the shader DSL keeps
	// its state in package variables.
	generating sync.Mutex

	precompiling     = make(chan Any, 1024)
	startPrecompiler sync.Once
)

func programFor(t reflect.Type) *program {
	p, _ := programs.LoadOrStore(t, new(program))
	return p.(*program)
}

// generate generates the code of the shader class of prog, unless it already
// has been.
func (p *program) generate(prog Any) {
	p.once.Do(func() {
		generating.Lock()
		defer generating.Unlock()
		defer func() {
			if r := recover(); r != nil {
				p.failure = r
			}
		}()
		rtype := reflect.TypeOf(prog)
		fragment, _ := rtype.MethodByName("Fragment")
		material, _ := rtype.MethodByName("Material")
		lighting, _ := rtype.MethodByName("Lighting")
		p.code = generate(prog, fragment.Type.In(1), material.Type.In(1), lighting.Type.In(1), lighting.Type.Out(0))
	})
}

func init() {
	gdclass.OnRegisterShader = precompile
}

// precompile starts generating the code of a shader class in the background,
// so that it is ready by the time the first material of the class is created.
// construct returns a new value of the class (with the defaults of its
// uniforms), it is called straight away, by classdb when the class is
// registered.
func precompile(construct func() reflect.Value) {
	prog, ok := reflect.TypeAssert[Any](construct())
	if !ok {
		return
	}
	startPrecompiler.Do(func() {
		go func() {
			for prog := range precompiling {
				programFor(reflect.TypeOf(prog)).generate(prog)
			}
		}()
	})
	select {
	case precompiling <- prog:
	default: // the queue is full, the class is compiled when it is first used.
	}
}

// generate returns the code of the shader in Godot's shading language.
func generate(prog Any, v, f, m, l reflect.Type) string {
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

	// Untagged fields returned by one stage are passed to the next stage as
	// varyings, which the GPU interpolates between the vertices of each face.
	vertexVaryings := pipeline[0] == "vertex" && pipeline[1] == "fragment"
	fragmentVaryings := pipeline[1] == "fragment" && pipeline[2] == "light"

	var vertices = reflect.New(v).Elem()
	linkup(vertices.Addr().Interface(), false)
	var fragment = reflect.New(f).Elem()
	linkup(fragment.Addr().Interface(), vertexVaryings)
	var material = reflect.New(m).Elem()
	linkup(material.Addr().Interface(), fragmentVaryings)

	compileUniforms(&writer, prog)
	declared := make(map[string]bool)
	if vertexVaryings {
		compileVaryings(&writer, rvalue.MethodByName("Fragment").Type().Out(0), declared)
	}
	if fragmentVaryings {
		compileVaryings(&writer, rvalue.MethodByName("Material").Type().Out(0), declared)
	}
	if len(declared) > 0 {
		fmt.Fprintln(&writer)
	}
	canon = newCanonicalizer()
	stages := []struct {
		method   string
		in       reflect.Value
		varyings bool
	}{
		{"Fragment", vertices, vertexVaryings},
		{"Material", fragment, fragmentVaryings},
		{"Lighting", material, false},
	}
	for i, stage := range stages {
		if pipeline[i] == "" {
			continue
		}
		if outputs := runStage(rvalue.MethodByName(stage.method), stage.in, pipeline[i], stage.varyings); outputs != nil {
			compileFunction(&writer, outputs, pipeline[i])
		}
	}
	return writer.String()
}

// linkup binds each field of the stage input in to the built-in it is tagged
// with, or when varyings is true, an untagged field to the varying it names.
func linkup(in any, varyings bool) {
	value := reflect.ValueOf(in).Elem()
	for field, rvalue := range value.Fields() {
		if !field.IsExported() {
			continue
		}
		if field.Type.Kind() == reflect.Struct && !isGPUValue(field.Type) {
			linkup(rvalue.Addr().Interface(), varyings)
			continue
		}
		if tag := field.Tag.Get("gd"); tag != "" {
			switch ptr := rvalue.Addr().Interface().(type) {
			case *vec2.XY:
				dsl.Set(&ptr.X, dsl.Identifier(tag+".x"))
				dsl.Set(&ptr.Y, dsl.Identifier(tag+".y"))
			}
			dsl.Set(rvalue.Addr().Interface().(dsl.Pointer), dsl.Identifier(tag))
		} else if varyings && isGPUValue(field.Type) {
			dsl.Set(rvalue.Addr().Interface().(dsl.Pointer), dsl.Identifier(String.ToSnakeCase(field.Name)))
		}
	}
}

// isGPUValue reports whether t is a value on the GPU (as opposed to a struct
// that groups such values together).
func isGPUValue(t reflect.Type) bool {
	return reflect.PointerTo(t).Implements(reflect.TypeFor[dsl.Pointer]())
}

// stageOutputs calls fn with the name to assign and the value of each output
// of a pipeline stage, looking inside of structs that group outputs together.
// Untagged fields are varyings, named after the field. The value of an unset
// output is nil.
func stageOutputs(value reflect.Value, fn func(name string, tagged bool, expr dsl.Evaluator)) {
	for field, rvalue := range value.Fields() {
		if !field.IsExported() {
			continue
		}
		if field.Type.Kind() == reflect.Struct && !isGPUValue(field.Type) {
			stageOutputs(rvalue, fn)
			continue
		}
		expr, ok := rvalue.Interface().(dsl.Evaluator)
		if !ok {
			continue
		}
		if rvalue.IsZero() {
			expr = nil // unset, so the built-in keeps its default.
		}
		if tag := field.Tag.Get("gd"); tag != "" {
			fn(tag, true, expr)
		} else {
			fn(String.ToSnakeCase(field.Name), false, expr)
		}
	}
}

// compileVaryings declares a varying for each untagged field of the stage output
// type t that has not already been declared.
func compileVaryings(w io.Writer, t reflect.Type, declared map[string]bool) {
	for field := range t.Fields() {
		if !field.IsExported() {
			continue
		}
		if field.Type.Kind() == reflect.Struct && !isGPUValue(field.Type) {
			compileVaryings(w, field.Type, declared)
			continue
		}
		if field.Tag.Get("gd") != "" || !isGPUValue(field.Type) {
			continue
		}
		name := String.ToSnakeCase(field.Name)
		if declared[name] {
			continue
		}
		declared[name] = true
		typ := glslTypeFor(field.Type)
		if strings.HasPrefix(typ, "i") || strings.HasPrefix(typ, "u") {
			typ = "flat " + typ // integers cannot be interpolated.
		}
		fmt.Fprintf(w, "varying %s %s;\n", typ, name)
	}
}

func compileUniforms(w io.Writer, prog Any) {
	compileUniformFields(w, prog, reflect.ValueOf(prog).Elem())
	fmt.Fprintln(w)
}

// compileUniformFields declares each field of value as a uniform. The fields of
// an embedded struct that is not itself a GPU type are declared as if they were
// fields of value, so that a set of uniforms can be shared between shaders.
func compileUniformFields(w io.Writer, prog Any, value reflect.Value) {
	for field, rvalue := range value.Fields() {
		if field.Name == "Shader" {
			continue
		}
		if field.Anonymous && field.Type.Kind() == reflect.Struct && !reflect.PointerTo(field.Type).Implements(reflect.TypeFor[dsl.Pointer]()) {
			compileUniformFields(w, prog, rvalue)
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
		// A value set on the field before the shader is compiled (by the
		// constructor passed to classdb.Register) becomes the uniform's default.
		// Only constants qualify: a copy of a compiled shader (as made when Godot
		// duplicates the material) already holds the uniform itself.
		var initial dsl.Evaluator
		if !rvalue.IsZero() {
			if expr, ok := rvalue.Interface().(dsl.Evaluator); ok && dsl.Evaluate(expr) == nil {
				initial = expr
			}
		}
		fmt.Fprintf(w, "uniform %s %s", glslTypeFor(field.Type), name)
		dsl.Set(rvalue.Addr().Interface().(dsl.Pointer), gpu.Uniform(name, prog))
		if options != "" {
			fmt.Fprintf(w, " : %s", strings.Join(splitOptions(options), ", "))
		}
		if initial != nil {
			fmt.Fprintf(w, " = ")
			compileExpression(w, initial, nil)
		}
		fmt.Fprintf(w, ";\n")
	}
}

// splitOptions splits the comma separated uniform hints of a struct tag, keeping
// the commas inside of parentheses, as in hint_range(0.0, 5.0).
func splitOptions(options string) []string {
	var (
		split []string
		depth int
		start int
	)
	for i, r := range options {
		switch r {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				split = append(split, strings.TrimSpace(options[start:i]))
				start = i + 1
			}
		}
	}
	return append(split, strings.TrimSpace(options[start:]))
}

// formatFloat formats x as the shortest GLSL float literal that represents it.
func formatFloat(x float64) string {
	s := strconv.FormatFloat(x, 'g', -1, 64)
	switch {
	case strings.ContainsAny(s, ".nN"):
		return s
	case strings.Contains(s, "e"):
		mantissa, exponent, _ := strings.Cut(s, "e")
		return mantissa + ".0e" + exponent
	default:
		return s + ".0"
	}
}

var glslTypes sync.Map // reflect.Type → string

func glslTypeFor(t reflect.Type) string {
	if name, ok := glslTypes.Load(t); ok {
		return name.(string)
	}
	name := glslTypeOf(t)
	glslTypes.Store(t, name)
	return name
}

func glslTypeOf(t reflect.Type) string {
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
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()), elem.ConvertibleTo(reflect.TypeFor[gpu.RGBA]()):
			return "sampler2D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
			return "isampler2D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
			return "usampler2D"
		}
	case t.Implements(reflect.TypeFor[gpu.IsSampler3D]()):
		elem := gpu.SamplerType(reflect.Zero(t).Interface().(gpu.IsSampler3D))
		switch {
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()), elem.ConvertibleTo(reflect.TypeFor[gpu.RGBA]()):
			return "sampler3D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
			return "isampler3D"
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
			return "usampler3D"
		}
	case t.Implements(reflect.TypeFor[gpu.IsArraySampler2D]()):
		elem := gpu.SamplerType(reflect.Zero(t).Interface().(gpu.IsArraySampler2D))
		switch {
		case elem.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()), elem.ConvertibleTo(reflect.TypeFor[gpu.RGBA]()):
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

type hasIdentity interface {
	Identity() *byte
}

type output struct {
	name string
	expr dsl.Evaluator
}

// runStage runs a pipeline function once for each way that the conditions it
// evaluates can go, and returns its outputs, or nil when it sets none.
func runStage(method reflect.Value, in reflect.Value, name string, varyings bool) []output {
	var names []string
	values := explore(func() []dsl.Evaluator {
		var values []dsl.Evaluator
		names = names[:0]
		stageOutputs(method.Call([]reflect.Value{in})[0], func(field string, tagged bool, expr dsl.Evaluator) {
			if expr != nil && !tagged && !varyings {
				panic(fmt.Sprintf("shaders: %s() cannot output %s, as only vertex() and fragment() can pass varyings to the next stage", name, field))
			}
			names = append(names, field)
			values = append(values, expr)
		})
		return values
	}, func(i int) string { return names[i] })
	var outputs []output
	for i, value := range values {
		if value != nil {
			outputs = append(outputs, output{names[i], value})
		}
	}
	return outputs
}

type varInfo struct {
	name string
	typ  string
	expr dsl.Evaluator
}

type compilationContext struct {
	available map[int]string // variables, by expression number.
}

// nested returns a context for a nested scope, whose variables are not
// available after it.
func (ctx *compilationContext) nested() *compilationContext {
	return &compilationContext{available: maps.Clone(ctx.available)}
}

// statement is either a variable declaration or a block (an if or a loop).
type statement struct {
	v     *varInfo
	block *dsl.Block
}

// emitter collects the statements of a scope, in the order that they run.
type emitter struct {
	stmts []statement
	split int // for a while loop, the statements before this one compute its condition.
}

// edge is a use of an expression (or block) by another. Body is the index of
// the block body that uses it, or -1 when it is used where its user is.
type edge struct {
	to   any // int expression number, or *dsl.Block
	body int
}

// node is an expression, or a block, in the function being compiled.
type node struct {
	expr    dsl.Evaluator // nil for a block.
	block   *dsl.Block
	edges   []edge
	split   int // number of edges computed by a while loop before its condition.
	uses    []*dsl.Scope
	homes   []*dsl.Scope // where it is computed.
	extract bool         // whether it is stored in a variable.
}

// function compiles a pipeline function. Each expression is computed in the
// innermost scope that contains all of its uses (so that work only needed by
// one branch of an if is done inside of that branch), but never inside of a
// loop that it was created outside of (so that work done outside of a loop body
// in Go, is done once, outside of the loop, on the GPU). An expression used
// more than once, or computed in a scope other than where it is used, is
// stored in a variable.
type function struct {
	nodes   map[any]*node
	post    []any // nodes in post-order, users after what they use.
	ordered map[any]bool
	scopes  map[*dsl.Scope]*emitter
	bodyOf  map[*dsl.Scope]*dsl.Block
	names   map[int]string // variables of blocks, by expression number.
	nextVar int
}

func compileFunction(w io.Writer, outputs []output, name string) {
	fn := &function{
		nodes:   make(map[any]*node),
		ordered: make(map[any]bool),
		scopes:  make(map[*dsl.Scope]*emitter),
		bodyOf:  make(map[*dsl.Scope]*dsl.Block),
		names:   make(map[int]string),
	}
	for _, out := range outputs {
		fn.visit(canon.of(out.expr))
	}
	for _, out := range outputs {
		n := fn.nodes[canon.of(out.expr)]
		n.uses = append(n.uses, nil)
	}
	fn.place()
	for _, out := range outputs {
		fn.order(canon.of(out.expr))
	}
	fmt.Fprintf(w, "void %s() {\n", name)
	ctx := &compilationContext{available: make(map[int]string)}
	for id, name := range fn.names {
		ctx.available[id] = name
	}
	fn.emit(w, fn.emitter(nil).stmts, 1, ctx)
	for _, out := range outputs {
		fmt.Fprintf(w, "\t%s = ", out.name)
		compileExpression(w, out.expr, ctx)
		fmt.Fprintf(w, ";\n")
	}
	fmt.Fprintf(w, "}\n")
}

// visit adds the node with key (and everything it uses) to the function.
func (fn *function) visit(key any) {
	if _, ok := fn.nodes[key]; ok {
		return
	}
	n := &node{}
	fn.nodes[key] = n
	use := func(e dsl.Evaluator, body int) {
		if e != nil {
			id := canon.of(e)
			fn.visit(id)
			n.edges = append(n.edges, edge{id, body})
		}
	}
	if b, ok := key.(*dsl.Block); ok {
		n.block = b
		for _, init := range b.Init {
			use(init, -1)
		}
		use(b.Count, -1)
		if b.Kind == dsl.BlockIf {
			use(b.Cond, -1)
		}
		for i, body := range b.Bodies {
			if b.Kind == dsl.BlockWhile {
				use(b.Cond, i)
				n.split = len(n.edges)
			}
			for _, next := range body.Next {
				use(next, i)
			}
			use(body.Break, i)
		}
	} else {
		n.expr = canon.reps[key.(int)]
		resolved := resolve(n.expr)
		if l, ok := resolved.(*dsl.Local); ok {
			if !l.Param {
				blk := canon.rep(l.Block)
				fn.visit(blk)
				n.edges = append(n.edges, edge{blk, -1})
			}
		} else {
			eachChild(resolved, func(e dsl.Evaluator) { use(e, -1) })
		}
	}
	fn.post = append(fn.post, key)
}

// place decides where each node is computed, visiting users before what they use.
func (fn *function) place() {
	for i := len(fn.post) - 1; i >= 0; i-- {
		n := fn.nodes[fn.post[i]]
		uses := n.uses
		var created *dsl.Scope
		clamp := true
		if n.block != nil {
			created = n.block.Created
		} else if hi, ok := n.expr.(hasIdentity); ok && hi.Identity() != nil {
			created = canon.loops[fn.post[i].(int)]
		} else {
			clamp = false // a literal value, computed wherever it is used.
		}
		if clamp {
			// A value is computed outside of the loops that it was created
			// outside of, so lift each use out of them.
			loop := created.InnermostLoop()
			lifted := make([]*dsl.Scope, len(uses))
			for j, use := range uses {
				for use.InnermostLoop() != loop {
					if use == nil {
						var code strings.Builder
						compileExpression(&code, n.expr, nil)
						panic("shaders: a value computed inside of a loop body is used outside of it, return it as part of the loop's state instead: " + code.String())
					}
					use = use.InnermostLoop().Parent
				}
				lifted[j] = use
			}
			uses = lifted
		}
		home := uses[0]
		for _, use := range uses[1:] {
			home = commonScope(home, use)
		}
		if n.block != nil {
			n.homes = []*dsl.Scope{home}
		} else {
			n.homes = fn.split(home, uses)
		}
		if n.block != nil {
			for _, body := range n.block.Bodies {
				body.Scope.Parent = home
				fn.bodyOf[body.Scope] = n.block
			}
		} else if !isSimpleExpr(resolve(n.expr)) && reflect.TypeOf(n.expr) != untyped && (clamp || hasComponents(n.expr)) {
			// Stored in a variable, unless it is used once, where it is computed.
			n.extract = len(uses) > len(n.homes) || !slices.Equal(uses, n.uses)
			for _, home := range n.homes {
				if !slices.Contains(uses, home) {
					n.extract = true
				}
			}
		}
		for _, home := range n.homes {
			for _, e := range n.edges {
				at := home
				if e.body >= 0 {
					at = n.block.Bodies[e.body].Scope
				}
				child := fn.nodes[e.to]
				child.uses = append(child.uses, at)
			}
		}
	}
}

// split returns where to compute a value used in uses, all of which are inside
// of scope. The value is computed in scope if it is used there, or if every
// branch of an if inside of it needs it. Otherwise each branch that uses it
// computes it separately, so that a branch that does not need it does not
// compute it.
func (fn *function) split(scope *dsl.Scope, uses []*dsl.Scope) []*dsl.Scope {
	var children []*dsl.Scope
	groups := make(map[*dsl.Scope][]*dsl.Scope)
	for _, use := range uses {
		if use == scope {
			return []*dsl.Scope{scope}
		}
		child := use
		for child.Parent != scope {
			child = child.Parent
		}
		if _, ok := groups[child]; !ok {
			children = append(children, child)
		}
		groups[child] = append(groups[child], use)
	}
	var homes []*dsl.Scope
	every := true // every child needs it on every path through the child.
	for _, child := range children {
		split := fn.split(child, groups[child])
		every = every && len(split) == 1 && split[0] == child
		homes = append(homes, split...)
	}
	if blk := fn.bodyOf[children[0]]; every && blk != nil && len(children) == len(blk.Bodies) {
		for _, child := range children {
			every = every && fn.bodyOf[child] == blk
		}
		if every {
			return []*dsl.Scope{scope}
		}
	}
	return homes
}

// hasComponents reports whether e is a vector or matrix literal that is made
// out of values computed on the GPU.
func hasComponents(e dsl.Evaluator) bool {
	found := false
	if dsl.Evaluate(e) == nil {
		eachComponent(reflect.ValueOf(e), func(c dsl.Evaluator) {
			if hi, ok := c.(hasIdentity); ok && hi.Identity() != nil {
				found = true
			}
		})
	}
	return found
}

// commonScope returns the innermost scope that contains both a and b.
func commonScope(a, b *dsl.Scope) *dsl.Scope {
	depth := func(s *dsl.Scope) (d int) {
		for ; s != nil; s = s.Parent {
			d++
		}
		return d
	}
	da, db := depth(a), depth(b)
	for ; da > db; da-- {
		a = a.Parent
	}
	for ; db > da; db-- {
		b = b.Parent
	}
	for a != b {
		a, b = a.Parent, b.Parent
	}
	return a
}

func (fn *function) emitter(scope *dsl.Scope) *emitter {
	e, ok := fn.scopes[scope]
	if !ok {
		e = &emitter{}
		fn.scopes[scope] = e
	}
	return e
}

// order adds the statements that compute the node with key, after the
// statements that compute what it uses.
func (fn *function) order(key any) {
	if fn.ordered[key] {
		return
	}
	fn.ordered[key] = true
	n := fn.nodes[key]
	if n.block == nil {
		for _, e := range n.edges {
			fn.order(e.to)
		}
		if n.extract {
			for _, home := range n.homes {
				name := fmt.Sprintf("_v%d", fn.nextVar)
				fn.nextVar++
				e := fn.emitter(home)
				e.stmts = append(e.stmts, statement{v: &varInfo{name: name, typ: glslTypeFor(reflect.TypeOf(n.expr)), expr: n.expr}})
			}
		}
		return
	}
	b := n.block
	prefix := map[dsl.BlockKind]string{dsl.BlockIf: "_r", dsl.BlockFor: "_s", dsl.BlockWhile: "_s"}[b.Kind]
	for i, l := range b.State {
		name := fmt.Sprintf("%s%d", prefix, fn.nextVar)
		fn.nextVar++
		fn.names[canon.of(l)] = name
		if b.Params != nil {
			fn.names[canon.of(b.Params[i])] = name
		}
	}
	if b.Index != nil {
		fn.names[canon.of(b.Index)] = fmt.Sprintf("_i%d", fn.nextVar)
		fn.nextVar++
	}
	for i, e := range n.edges {
		if e.body < 0 {
			fn.order(e.to)
			continue
		}
		body := fn.emitter(b.Bodies[e.body].Scope)
		fn.order(e.to)
		if b.Kind == dsl.BlockWhile && i == n.split-1 {
			body.split = len(body.stmts)
		}
	}
	e := fn.emitter(n.homes[0])
	e.stmts = append(e.stmts, statement{block: b})
}

func (fn *function) emit(w io.Writer, stmts []statement, depth int, ctx *compilationContext) {
	indent := strings.Repeat("\t", depth)
	for _, stmt := range stmts {
		if v := stmt.v; v != nil {
			fmt.Fprintf(w, "%s%s %s = ", indent, v.typ, v.name)
			compileExpression(w, v.expr, ctx)
			fmt.Fprintf(w, ";\n")
			ctx.available[canon.of(v.expr)] = v.name
			continue
		}
		b := stmt.block
		switch b.Kind {
		case dsl.BlockIf:
			for _, l := range b.State {
				fmt.Fprintf(w, "%s%s %s;\n", indent, glslTypeFor(l.Type), fn.names[canon.of(l)])
			}
			fmt.Fprintf(w, "%sif (", indent)
			compileExpression(w, b.Cond, ctx)
			fmt.Fprintf(w, ") {\n")
			for i, body := range b.Bodies {
				if i > 0 {
					fmt.Fprintf(w, "%s} else {\n", indent)
				}
				inner := ctx.nested()
				fn.emit(w, fn.emitter(body.Scope).stmts, depth+1, inner)
				fn.assign(w, b, body, depth+1, inner)
			}
			fmt.Fprintf(w, "%s}\n", indent)
		case dsl.BlockFor, dsl.BlockWhile:
			for i, l := range b.State {
				fmt.Fprintf(w, "%s%s %s = ", indent, glslTypeFor(l.Type), fn.names[canon.of(l)])
				compileExpression(w, b.Init[i], ctx)
				fmt.Fprintf(w, ";\n")
			}
			body := b.Bodies[0]
			e := fn.emitter(body.Scope)
			outer := ctx
			ctx := ctx.nested()
			stmts := e.stmts
			switch {
			case b.Kind == dsl.BlockFor:
				i := fn.names[canon.of(b.Index)]
				fmt.Fprintf(w, "%sfor (int %s = 0; %s < ", indent, i, i)
				compileExpression(w, b.Count, outer)
				fmt.Fprintf(w, "; %s++) {\n", i)
			case e.split == 0:
				fmt.Fprintf(w, "%swhile (", indent)
				compileExpression(w, b.Cond, ctx)
				fmt.Fprintf(w, ") {\n")
			default:
				fmt.Fprintf(w, "%swhile (true) {\n", indent)
				fn.emit(w, stmts[:e.split], depth+1, ctx)
				fmt.Fprintf(w, "%s\tif (!", indent)
				compileExpression(w, b.Cond, ctx)
				fmt.Fprintf(w, ") {\n%s\t\tbreak;\n%s\t}\n", indent, indent)
				stmts = stmts[e.split:]
			}
			fn.emit(w, stmts, depth+1, ctx)
			fn.assign(w, b, body, depth+1, ctx)
			if body.Break != nil {
				fmt.Fprintf(w, "%s\tif (", indent)
				compileExpression(w, body.Break, ctx)
				fmt.Fprintf(w, ") {\n%s\t\tbreak;\n%s\t}\n", indent, indent)
			}
			fmt.Fprintf(w, "%s}\n", indent)
		}
	}
}

// assign writes the values of a body to the variables of its block. As the
// variables are assigned one after another, a value that reads a variable
// assigned before it is first copied into a temporary variable.
func (fn *function) assign(w io.Writer, b *dsl.Block, body dsl.Body, depth int, ctx *compilationContext) {
	indent := strings.Repeat("\t", depth)
	type assignment struct {
		name  string
		typ   string
		value string
	}
	var assignments []assignment
	for i, l := range b.State {
		next := body.Next[i]
		if b.Params != nil && canon.of(next) == canon.of(b.Params[i]) {
			continue // unchanged.
		}
		var buf strings.Builder
		compileExpression(&buf, next, ctx)
		assignments = append(assignments, assignment{fn.names[canon.of(l)], glslTypeFor(l.Type), buf.String()})
	}
	for i, a := range assignments {
		for _, later := range assignments[i+1:] {
			if containsWord(later.value, a.name) {
				tmp := fmt.Sprintf("_t%d", fn.nextVar)
				fn.nextVar++
				fmt.Fprintf(w, "%s%s %s = %s;\n", indent, a.typ, tmp, a.value)
				assignments[i].value = tmp
				break
			}
		}
	}
	for _, a := range assignments {
		fmt.Fprintf(w, "%s%s = %s;\n", indent, a.name, a.value)
	}
}

// containsWord reports whether the GLSL code reads the variable name.
func containsWord(code, name string) bool {
	for i := 0; ; {
		j := strings.Index(code[i:], name)
		if j < 0 {
			return false
		}
		j += i
		end := j + len(name)
		before := j == 0 || !isIdentByte(code[j-1])
		after := end == len(code) || !isIdentByte(code[end])
		if before && after {
			return true
		}
		i = end
	}
}

func isIdentByte(c byte) bool {
	return c == '_' || c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// eachChild calls fn with each expression that resolved is computed from.
func eachChild(resolved dsl.Evaluator, fn func(dsl.Evaluator)) {
	switch v := resolved.(type) {
	case dsl.Operation:
		if v.A != nil {
			fn(v.A)
		}
		fn(v.B)
	case dsl.Select:
		fn(v.Value)
	case dsl.Index:
		fn(v.Value)
	case dsl.FunctionCall:
		for _, arg := range v.Args {
			fn(arg)
		}
	case dsl.Ternary:
		fn(v.If)
		fn(v.A)
		fn(v.B)
	case dsl.Identifier, *dsl.Local:
	default:
		eachComponent(reflect.ValueOf(resolved), fn)
	}
}

// eachComponent calls fn for each component of a vector or matrix literal,
// which are the exported fields of the value, or the entries of its columns.
func eachComponent(rvalue reflect.Value, fn func(dsl.Evaluator)) {
	for _, leaf := range leafPaths(rvalue.Type()) {
		if leaf.evaluator {
			if ev, ok := leaf.get(rvalue).Interface().(dsl.Evaluator); ok {
				fn(ev)
			}
		}
	}
}

// eachLeaf calls fn with each component of a vector or matrix literal, and
// each Go value (such as the float64 of a scalar literal) that it is made of.
func eachLeaf(rvalue reflect.Value, component func(dsl.Evaluator), basic func(reflect.Value)) {
	for _, leaf := range leafPaths(rvalue.Type()) {
		value := leaf.get(rvalue)
		if !leaf.evaluator {
			basic(value)
		} else if ev, ok := value.Interface().(dsl.Evaluator); ok {
			component(ev)
		}
	}
}

// leafPath is the path (of field or element indices) to a component of a
// value, or to a Go value that it is made of.
type leafPath struct {
	path      []int
	evaluator bool
}

func (l leafPath) get(value reflect.Value) reflect.Value {
	for _, step := range l.path {
		if value.Kind() == reflect.Array {
			value = value.Index(step)
		} else {
			value = value.Field(step)
		}
	}
	return value
}

var leafPathCache sync.Map // reflect.Type → []leafPath

var evaluatorType = reflect.TypeFor[dsl.Evaluator]()

// leafPaths returns the leaves of values of type t, worked out once for each
// type, as literals are visited for every expression.
func leafPaths(t reflect.Type) []leafPath {
	if paths, ok := leafPathCache.Load(t); ok {
		return paths.([]leafPath)
	}
	var paths []leafPath
	var walk func(t reflect.Type, path []int)
	walk = func(t reflect.Type, path []int) {
		visit := func(t reflect.Type, path []int) {
			switch {
			case t.Implements(evaluatorType) || t.Kind() == reflect.Interface:
				paths = append(paths, leafPath{slices.Clone(path), true})
			case t.Kind() == reflect.Struct || t.Kind() == reflect.Array:
				walk(t, path)
			default:
				paths = append(paths, leafPath{slices.Clone(path), false})
			}
		}
		switch t.Kind() {
		case reflect.Struct:
			for i := range t.NumField() {
				if field := t.Field(i); field.IsExported() {
					visit(field.Type, append(path, i))
				}
			}
		case reflect.Array:
			for i := range t.Len() {
				visit(t.Elem(), append(path, i))
			}
		}
	}
	walk(t, nil)
	leafPathCache.Store(t, paths)
	return paths
}

func isSimpleExpr(resolved dsl.Evaluator) bool {
	switch v := resolved.(type) {
	case dsl.Identifier, *dsl.Local:
		return true
	case dsl.Select:
		switch resolve(v.Value).(type) {
		case dsl.Identifier, *dsl.Local:
			return true
		}
	}
	return false
}

func compileExpression(w io.Writer, expression dsl.Evaluator, ctx *compilationContext) {
	if expression == nil {
		return
	}
	if ctx != nil {
		if name, avail := ctx.available[canon.of(expression)]; avail {
			fmt.Fprintf(w, "%s", name)
			return
		}
	}
	if expr := dsl.Evaluate(expression); expr != nil {
		expression = expr
	}
	switch value := expression.(type) {
	case dsl.Operation:
		fmt.Fprintf(w, "(")
		compileExpression(w, value.A, ctx)
		fmt.Fprintf(w, " %s ", value.Op)
		compileExpression(w, value.B, ctx)
		fmt.Fprintf(w, ")")
		return
	case dsl.Identifier:
		fmt.Fprintf(w, "%s", value)
		return
	case dsl.Ternary:
		fmt.Fprintf(w, "(")
		compileExpression(w, value.If, ctx)
		fmt.Fprintf(w, " ? ")
		compileExpression(w, value.A, ctx)
		fmt.Fprintf(w, " : ")
		compileExpression(w, value.B, ctx)
		fmt.Fprintf(w, ")")
		return
	case dsl.Select:
		compileExpression(w, value.Value, ctx)
		fmt.Fprintf(w, ".%s", value.Field)
		return
	case dsl.Index:
		compileExpression(w, value.Value, ctx)
		fmt.Fprintf(w, "[%d]", value.Index)
		return
	case dsl.FunctionCall:
		compileCall(w, string(value.Name), ctx, value.Args...)
		return
	case *dsl.Local:
		panic("shaders: a variable of an if or loop is read where it is not available")
	}
	rtype := reflect.TypeOf(expression)
	switch {
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.RGBA]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.RGBA]()).Interface().(gpu.RGBA)
		compileCall(w, "vec4", ctx, value.R, value.G, value.B, value.A)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4]()).Interface().(gpu.Vec4)
		compileCall(w, "vec4", ctx, value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4i]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4i]()).Interface().(gpu.Vec4i)
		compileCall(w, "ivec4", ctx, value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4u]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4u]()).Interface().(gpu.Vec4u)
		compileCall(w, "uvec4", ctx, value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec4b]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec4b]()).Interface().(gpu.Vec4b)
		compileCall(w, "bvec4", ctx, value.X, value.Y, value.Z, value.W)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3]()).Interface().(gpu.Vec3)
		compileCall(w, "vec3", ctx, value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.RGB]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.RGB]()).Interface().(gpu.RGB)
		compileCall(w, "vec3", ctx, value.R, value.G, value.B)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3i]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3i]()).Interface().(gpu.Vec3i)
		compileCall(w, "ivec3", ctx, value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3u]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3u]()).Interface().(gpu.Vec3u)
		compileCall(w, "uvec3", ctx, value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec3b]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec3b]()).Interface().(gpu.Vec3b)
		compileCall(w, "bvec3", ctx, value.X, value.Y, value.Z)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2]()).Interface().(gpu.Vec2)
		compileCall(w, "vec2", ctx, value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2i]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2i]()).Interface().(gpu.Vec2i)
		compileCall(w, "ivec2", ctx, value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2u]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2u]()).Interface().(gpu.Vec2u)
		compileCall(w, "uvec2", ctx, value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Vec2b]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Vec2b]()).Interface().(gpu.Vec2b)
		compileCall(w, "bvec2", ctx, value.X, value.Y)
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Float]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Float]()).Interface().(gpu.Float)
		fmt.Fprintf(w, "%s", formatFloat(value.X))
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
		compileMatrix(w, "mat2", "vec2", ctx, value.Columns[0][:], value.Columns[1][:])
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Mat3]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Mat3]()).Interface().(gpu.Mat3)
		compileMatrix(w, "mat3", "vec3", ctx, value.Columns[0][:], value.Columns[1][:], value.Columns[2][:])
	case rtype.ConvertibleTo(reflect.TypeFor[gpu.Mat4]()):
		value := reflect.ValueOf(expression).Convert(reflect.TypeFor[gpu.Mat4]()).Interface().(gpu.Mat4)
		compileMatrix(w, "mat4", "vec4", ctx, value.Columns[0][:], value.Columns[1][:], value.Columns[2][:], value.Columns[3][:])
	default:
		panic(fmt.Sprintf("unsupported expression type %T", expression))
	}
}

// compileMatrix writes a matrix constructor from its column vectors, as Godot's
// shading language does not accept a matrix constructor with a scalar per entry.
func compileMatrix(w io.Writer, mat, vec string, ctx *compilationContext, columns ...[]gpu.Float) {
	fmt.Fprintf(w, "%s(", mat)
	for i, column := range columns {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		var args []dsl.Evaluator
		for _, entry := range column {
			args = append(args, entry)
		}
		compileCall(w, vec, ctx, args...)
	}
	fmt.Fprintf(w, ")")
}

func compileCall(w io.Writer, name string, ctx *compilationContext, args ...dsl.Evaluator) {
	fmt.Fprintf(w, "%s(", name)
	for i, arg := range args {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		compileExpression(w, arg, ctx)
	}
	fmt.Fprintf(w, ")")
}
