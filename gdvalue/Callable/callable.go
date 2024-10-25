// Package Callable provides generic methods for working with callable functions.
package Callable

import (
	"reflect"
	"runtime"

	"grow.graphics/gd/gdvalue/Array"
)

// Func representation.
type Func interface {
	Name() string
	Args() (args int, bind Array.Of[any])
	Call(arg ...any) any
	Bind(args ...any) Func
}

// New returns a new [Func] from the given value, if the value is not a Go func
// then it will be wrapped as if it were a function without any arguments that
// returns the specified value.
func New(value any) Func {
	return &local{value: value}
}

type local struct {
	value any
	cache []reflect.Value
	binds Array.Of[any]
	trims int
	proxy Func
}

func (l *local) Name() string {
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(l.value).Pointer()).Name()
	}
	return "<unknown callable>"
}

func (l *local) Args() (int, Array.Of[any]) {
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		return reflect.TypeOf(l.value).NumIn(), l.binds
	}
	return 0, l.binds
}

func (l *local) Call(args ...any) any {
	argc, binds := l.Args()
	if len(args)-Array.Size(binds) != argc {
		panic("invalid number of arguments")
	}
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		for i := range args {
			l.cache[i] = reflect.ValueOf(args[i])
		}
		for i, v := range Array.Iter(binds) {
			l.cache[i+len(args)] = reflect.ValueOf(v)
		}
		result := reflect.ValueOf(l.value).Call(l.cache)
		if len(result) == 0 {
			return nil
		}
		return result[0].Interface()
	}
	return l.value
}

func (l local) Bind(args ...any) Func {
	for _, arg := range args {
		l.cache = append(l.cache, reflect.ValueOf(arg))
	}
	return &l
}

// Bind returns a copy of this Callable with one or more arguments bound. When called,
// the bound arguments are passed after the arguments supplied by call. See also [Unbind].
//
// Note: When this method is chained with other similar methods, the order in which the
// argument list is modified is read from right to left.
func Bind(fn Func, args ...any) Func {
	if fn == nil {
		return nil
	}
	return fn.Bind(args...)
}

// Call calls the method represented by this Callable. Arguments can be passed and should
// match the method's signature.
func Call(fn Func, args ...any) any {
	if fn == nil {
		return nil
	}
	return fn.Call(args...)
}

// Create creates a new Callable for the method named method in the specified value.
func Create(value any, method string) Func {
	if value == nil {
		return nil
	}
	return New(reflect.ValueOf(value).MethodByName(method).Interface())
}

// ArgumentCount returns the total number of arguments this Callable should take.
func ArgumentCount(fn Func) int {
	if fn == nil {
		return 0
	}
	argc, _ := fn.Args()
	return argc
}

// BoundArguments returns the arguments that have been bound to this Callable.
func BoundArguments(fn Func) Array.Of[any] {
	if fn == nil {
		return nil
	}
	_, binds := fn.Args()
	return binds
}

// BoundArgumentsCount returns the total amount of arguments bound (or unbound)
// via successive bind or unbind calls. If the amount of arguments unbound is
// greater than the ones bound, this function returns a value less than zero.
func BoundArgumentsCount(fn Func) int {
	if fn == nil {
		return 0
	}
	return Array.Size(BoundArguments(fn))
}

// Method returns the name of the function represented by this Callable or
// an empty string if a name is not available.
func Method(fn Func) string {
	if fn == nil {
		return ""
	}
	return fn.Name()
}

// IsProxy returns true if the given value is not backed by a Go function.
func IsProxy(fn Func) bool {
	if fn == nil {
		return true
	}
	return reflect.TypeOf(fn).Kind() != reflect.Func
}

// IsNull returns true if the given value is nil.
func IsNull(fn Func) bool {
	return fn == nil
}

// IsStandard returns true if the given value is backed by a Go function.
func IsStandard(fn Func) bool {
	if fn == nil {
		return false
	}
	return reflect.TypeOf(fn).Kind() == reflect.Func
}

// IsValid returns true if the given value is not nil.
func IsValid(fn Func) bool {
	return fn != nil
}
