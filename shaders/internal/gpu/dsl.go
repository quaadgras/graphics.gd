package gpu

import (
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"graphics.gd/classdb/ShaderMaterial"
)

type Expression struct {
	identity *byte
	indirect Evaluator
	uniform  string
	shader   ShaderMaterial.Any

	// literal marks a constant written in Go, so that a zero constant, such as
	// the 0.0 in an output of ROUGHNESS = 0.0, is not mistaken for an unset value.
	literal bool

	// scope is the block that the expression was created in, nil for the top
	// level of a pipeline function.
	scope *Scope

	// stamp numbers the expressions created during a run of a pipeline function
	// (outside of loop bodies) in the order they were created, so that runs can
	// be compared.
	stamp int
}

func Uniform(name string, shader ShaderMaterial.Any) Evaluator {
	return Expression{identity: new(byte), uniform: name, shader: shader}
}

func New(e Evaluator) Expression {
	if t := tracing; t != nil {
		target := Expression{identity: new(byte), indirect: e, scope: current}
		c := &Cell{Target: target, Trace: t, Iteration: t.Iteration, Seq: len(t.Cells[t.Iteration-1])}
		t.Cells[t.Iteration-1] = append(t.Cells[t.Iteration-1], c)
		return Expression{identity: new(byte), indirect: c, scope: current}
	}
	stamps++
	return Expression{identity: new(byte), indirect: e, scope: current, stamp: stamps}
}

// NewUnstamped is like New, for an expression that is not created by the
// pipeline function itself.
func NewUnstamped(e Evaluator) Expression {
	return Expression{identity: new(byte), indirect: e, scope: current}
}

func Op(a Evaluator, op string, b Evaluator) Expression {
	return New(Operation{A: a, B: b, Op: op})
}

func (e Expression) evaluate() Evaluator {
	if e.uniform != "" {
		return Identifier(e.uniform)
	}
	if e.indirect != nil {
		return e.indirect.evaluate()
	}
	return nil
}

func (e *Expression) set(ptr Pointer, val Evaluator) {
	if expr, ok := val.(Expression); ok {
		*e = expr
	} else {
		*e = New(val)
	}
	plan := componentsOf(reflect.TypeOf(ptr).Elem())
	if len(plan) == 0 {
		return
	}
	ifc := reflect.ValueOf(ptr).Elem()
	// Components select from a typed copy of the parent, so that a parent that is
	// only ever used through its components can still be extracted into a variable.
	parent, ok := reflect.TypeAssert[Evaluator](ifc)
	if !ok {
		parent = val
	}
	base := ifc.Addr().UnsafePointer()
	for _, c := range plan {
		if c.scalar {
			// The expression of a scalar is the first field of its struct.
			*(*Expression)(unsafe.Add(base, c.offset)) = New(Select{Value: parent, Field: c.name})
			continue
		}
		if ptr, ok := reflect.TypeAssert[Pointer](ifc.Field(c.index).Addr()); ok {
			ptr.set(ptr, New(Select{Value: parent, Field: c.name}))
		}
	}
}

// component is a field of a GPU value that selects one of its components.
type component struct {
	index  int
	offset uintptr
	name   string
	scalar bool // a scalar, which has no components of its own.
}

var components sync.Map // reflect.Type → []component

// componentsOf returns the components of values of type t, working them out
// once for each type, as values are created for every operation.
func componentsOf(t reflect.Type) []component {
	if plan, ok := components.Load(t); ok {
		return plan.([]component)
	}
	var plan []component
	if t.Kind() == reflect.Struct {
		for i := range t.NumField() {
			field := t.Field(i)
			if !field.IsExported() || field.Anonymous || !reflect.PointerTo(field.Type).Implements(reflect.TypeFor[Pointer]()) {
				continue
			}
			plan = append(plan, component{
				index:  i,
				offset: field.Offset,
				name:   strings.ToLower(field.Name),
				scalar: isScalar(field.Type),
			})
		}
	}
	components.Store(t, plan)
	return plan
}

// isScalar reports whether t is a struct that starts with an embedded
// Expression and has no components of its own.
func isScalar(t reflect.Type) bool {
	if t.Kind() != reflect.Struct || t.NumField() == 0 {
		return false
	}
	first := t.Field(0)
	if !first.Anonymous || first.Type != reflect.TypeFor[Expression]() || first.Offset != 0 {
		return false
	}
	return len(componentsOf(t)) == 0
}

func (e Expression) Identity() *byte {
	return e.identity
}

func (e Expression) getScope() *Scope {
	if c, ok := e.indirect.(*Cell); ok {
		return ScopeOf(c.Target)
	}
	return e.scope
}

func (e Expression) info() Info {
	c, _ := e.indirect.(*Cell)
	return Info{Cell: c, Stamp: e.stamp, Identity: e.identity, Literal: e.literal}
}

func (e Expression) getShader() ShaderMaterial.Any {
	return e.shader
}

func Shader(ptr HasShader) ShaderMaterial.Any {
	return ptr.getShader()
}

type Evaluator interface {
	evaluate() Evaluator
}

type EquivalentTo[T any] interface {
	Evaluator
	equivalentTo(T)
	getShader() ShaderMaterial.Any
}

type HasShader interface {
	getShader() ShaderMaterial.Any
}

type Pointer interface {
	set(Pointer, Evaluator)
}

type Operation struct {
	A  Evaluator
	B  Evaluator
	Op string
}

type Identifier string

func (n Identifier) evaluate() Evaluator { return n }

func (e Operation) evaluate() Evaluator {
	return e
}

func Evaluate(e Evaluator) Evaluator {
	return e.evaluate()
}

func Set(ptr Pointer, val Evaluator) {
	ptr.set(ptr, val)
}

type Select struct {
	Value Evaluator
	Field string
}

func (i Select) evaluate() Evaluator {
	return i
}

// Index is element Index of an array, or column Index of a matrix.
type Index struct {
	Value Evaluator
	Index int
}

func (i Index) evaluate() Evaluator {
	return i
}

type Ternary struct {
	If Evaluator
	A  Evaluator
	B  Evaluator
}

func (i Ternary) evaluate() Evaluator {
	return i
}

type FunctionCall struct {
	Name string
	Args []Evaluator
}

func (f FunctionCall) evaluate() Evaluator {
	return f
}

type Output struct {
	Index *int
	Type  string
}

func Out(t string) Expression {
	return New(Output{Index: new(int), Type: t})
}

func (o Output) evaluate() Evaluator { return o }

func Fn(name string, args ...Evaluator) Expression {
	return New(FunctionCall{Name: name, Args: args})
}

// If returns a when cond is true, otherwise b.
func If[T Evaluator](cond Evaluator, a, b T) T {
	var result T
	ptr, ok := any(&result).(Pointer)
	if !ok {
		panic("gpu.If: unsupported type " + reflect.TypeFor[T]().String())
	}
	Set(ptr, New(Ternary{If: cond, A: a, B: b}))
	return result
}
