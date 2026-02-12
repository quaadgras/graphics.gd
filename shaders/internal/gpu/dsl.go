package gpu

import (
	"reflect"
	"strings"

	"graphics.gd/classdb/ShaderMaterial"
)

type Expression struct {
	identity *byte
	indirect Evaluator
	uniform  string
	shader   ShaderMaterial.Any
}

func Uniform(name string, shader ShaderMaterial.Any) Evaluator {
	return Expression{identity: new(byte), uniform: name, shader: shader}
}

func New(e Evaluator) Expression {
	return Expression{identity: new(byte), indirect: e}
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
	ifc := reflect.ValueOf(ptr).Elem()
	for field, value := range ifc.Fields() {
		if !field.IsExported() || field.Anonymous {
			continue
		}
		if ptr, ok := reflect.TypeAssert[Pointer](value.Addr()); ok {
			ptr.set(ptr, New(Select{
				Value: val,
				Field: strings.ToLower(field.Name),
			}))
		}
	}
}

func (e Expression) Identity() *byte {
	return e.identity
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
