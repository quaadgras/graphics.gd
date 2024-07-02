//go:build !generate

package gd

import (
	"reflect"
)

// Callable creates a new callable out of the given function which must only accept
// godot-compatible types and return up to one godot-compatible type.
func (godot Lifetime) Callable(fn any) Callable {
	rvalue := reflect.ValueOf(fn)

	var hasContext bool
	if rvalue.Type().NumIn() > 0 {
		if rvalue.Type().In(0) == reflect.TypeOf(godot) {
			hasContext = true
		}
	}
	var offset = 0
	if hasContext {
		offset = 1
	}

	return godot.API.Callables.Create(godot, func(args ...Variant) (Variant, error) {
		godot := NewContext(godot.API)
		defer godot.End()

		var vargs = make([]reflect.Value, 0, len(args)+offset)
		if hasContext {
			vargs = append(vargs, reflect.ValueOf(godot))
		}
		for i, arg := range args {
			vargs = append(vargs, reflect.ValueOf(arg.Interface(godot)).Convert(rvalue.Type().In(i+offset)))
		}
		results := rvalue.Call(vargs)
		if len(results) == 0 {
			return godot.API.Variants.NewNil(godot), nil
		}
		return godot.Variant(results[0].Interface()), nil
	})
}
