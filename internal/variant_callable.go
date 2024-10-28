//go:build !generate

package gd

import (
	"reflect"
)

// Callable creates a new callable out of the given function which must only accept
// godot-compatible types and return up to one godot-compatible type.
func NewCallable(fn any) Callable {
	rvalue := reflect.ValueOf(fn)
	var offset = 0
	return Global.Callables.Create(func(args ...Variant) (Variant, error) {
		var vargs = make([]reflect.Value, 0, len(args)+offset)
		for i, arg := range args {
			vargs = append(vargs, reflect.ValueOf(arg.Interface()).Convert(rvalue.Type().In(i+offset)))
		}
		results := rvalue.Call(vargs)
		if len(results) == 0 {
			return Global.Variants.NewNil(), nil
		}
		return NewVariant(results[0].Interface()), nil
	})
}
