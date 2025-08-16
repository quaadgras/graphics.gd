//go:build !generate

package gd

import (
	"reflect"

	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	CallableType "graphics.gd/variant/Callable"
)

/*func init() {
type Callframe struct {
	_ structs.HostLayout

	r_return *gdextension.Variant
	r_error  *gdextension.CallError
	p_args   **gdextension.Variant
	p_argc   int64
}

gdextension.On.Callables.Call = func(fn gdextension.FunctionID, arg_count int64, args gdextension.CallAccepts[gdextension.Variant]) (gdextension.CallReturns[gdextension.Variant], gdextension.MaybeError) {
	frame := (*Callframe)(args)
	result, _ := cgo.Handle(fn).Value().(func(args []*Variant) (Variant, error))(nil)
	frame.r_error.Type = 0
	*frame.r_return, _ = pointers.End(result)
	return gdextension.CallReturns[gdextension.Variant]{}, gdextension.MaybeError{}
}
gdextension.On.Callables.Hash = func(fn gdextension.FunctionID) int64 {
	return int64(fn)
}
gdextension.On.Callables.Validation = func(fn gdextension.FunctionID) bool {
	return fn != 0
}
gdextension.On.Callables.Stringify = func(fn gdextension.FunctionID, call gdextension.Call) (gdextension.String, gdextension.MaybeError) {
	frame := (*Callframe)(call)
	s := NewString(reflect.ValueOf(cgo.Handle(fn).Value()).String())
	raw, _ := pointers.End(s)
	frame.r_error.Type = 0
	return gdextension.String(raw[0]), gdextension.MaybeError{}
}
}*/

// Callable creates a new callable out of the given function which must only accept
// godot-compatible types and return up to one godot-compatible type.
func NewCallable(fn any) Callable {
	rvalue := reflect.ValueOf(fn)
	ftype := rvalue.Type()
	return Global.Callables.Create(func(args ...Variant) (_ Variant, err error) {
		var vargs = make([]reflect.Value, len(args))
		for i, arg := range args {
			vargs[i], err = ConvertToDesiredGoType(arg, ftype.In(i))
			if err != nil {
				panic(err)
			}
		}
		results := rvalue.Call(vargs)
		if len(results) == 0 {
			return Global.Variants.NewNil(), nil
		}
		result := results[0].Interface()
		return CutVariant(result, true), nil
	})
}

func InternalCallable(fn CallableType.Function) Callable {
	return Global.Callables.Create(func(args ...Variant) (_ Variant, err error) {
		converted := make([]VariantPkg.Any, len(args))
		for i, arg := range args {
			converted[i] = VariantPkg.New(arg.Interface())
		}
		result := fn.Call(converted...).Interface()
		return CutVariant(result, true), nil
	})
}

type CallableProxy struct{}

func (CallableProxy) Name(state complex128) string {
	return pointers.Load[Callable](state).GetMethod().String()
}
func (CallableProxy) Args(state complex128) (args int, bind ArrayType.Any) {
	c := pointers.Load[Callable](state)
	b := c.GetBoundArguments()
	return int(c.GetArgumentCount()), ArrayType.Through(ArrayProxy[VariantPkg.Any]{}, pointers.Pack(b))
}
func (CallableProxy) Call(state complex128, args ...VariantPkg.Any) VariantPkg.Any {
	c := pointers.Load[Callable](state)
	vargs := make([]Variant, len(args))
	for i, arg := range args {
		vargs[i] = NewVariant(arg.Interface())
	}
	return VariantPkg.New(c.Call(vargs...).Interface())
}
func (CallableProxy) Bind(state complex128, args ...VariantPkg.Any) (CallableType.Proxy, complex128) {
	c := pointers.Load[Callable](state)
	vargs := make([]Variant, len(args))
	for i, arg := range args {
		vargs[i] = NewVariant(arg.Interface())
	}
	return CallableProxy{}, pointers.Pack(c.Bind(vargs...))
}
