//go:build !generate

package gd

import (
	"reflect"
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
		return NewVariant(results[0].Interface()), nil
	})
}
