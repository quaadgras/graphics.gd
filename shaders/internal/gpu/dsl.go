package gpu

import (
	"weak"

	"graphics.gd/shaders/internal"
)

type Expression struct {
	indirect Evaluator
	uniform  string
	shader   weak.Pointer[internal.Shader]
}

func Uniform(name string, shader *internal.Shader) Evaluator {
	return Expression{uniform: name, shader: weak.Make(shader)}
}

func New(e Evaluator) Expression {
	return Expression{indirect: e}
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

func (e *Expression) set(val Evaluator) {
	if expr, ok := val.(Expression); ok {
		*e = expr
		return
	}
	*e = New(val)
}

func (e Expression) getShader() *internal.Shader {
	return e.shader.Value()
}

func Shader(ptr HasShader) *internal.Shader {
	return ptr.getShader()
}

type Evaluator interface {
	evaluate() Evaluator
}

type EquivalentTo[T any] interface {
	Evaluator
	equivalentTo(T)
	getShader() *internal.Shader
}

type HasShader interface {
	getShader() *internal.Shader
}

type Pointer interface {
	set(Evaluator)
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
	ptr.set(val)
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
