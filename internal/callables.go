//go:build !generate

package gd

import (
	"reflect"

	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	CallableType "graphics.gd/variant/Callable"
)

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
