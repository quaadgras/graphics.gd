// Package Callable provides generic methods for working with callable functions.
package Callable

import (
	"math"
	"reflect"
	"runtime"

	gd "graphics.gd/internal"
	"graphics.gd/variant/Array"
)

// Func representation.
type Func interface {
	Name() string
	Args() (args int, bind Array.Contains[any])
	Call(arg ...any) any
	Bind(args ...any) Func
}

type Any = gd.Callable

// New returns a new [Func] from the given value, if the value is not a Go func
// then it will be wrapped as if it were a function without any arguments that
// returns the specified value.
func New(value any) Any {
	return gd.NewCallable(value)
}

type local struct {
	value any
	cache []reflect.Value
	binds Array.Contains[any]
	trims int
	proxy Func
}

func (l *local) Name() string {
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(l.value).Pointer()).Name()
	}
	return "<unknown callable>"
}

func (l *local) Args() (int, Array.Contains[any]) {
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		return reflect.TypeOf(l.value).NumIn(), l.binds
	}
	return 0, l.binds
}

func (l *local) Call(args ...any) any {
	argc, binds := l.Args()
	if len(args)-binds.Len() != argc {
		panic("invalid number of arguments")
	}
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		for i := range args {
			l.cache[i] = reflect.ValueOf(args[i])
		}
		for i, v := range binds.Iter() {
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
func Bind(fn Func, args ...any) Func { //gd:Callable.bind Callable.bindv
	if fn == nil {
		return nil
	}
	return fn.Bind(args...)
}

// Call calls the method represented by this Callable. Arguments can be passed and should
// match the method's signature.
func Call(fn Func, args ...any) any { //gd:Callable.call Callable.call_deferred Callable.callv Callable.rpc Callable.rpc_id
	if fn == nil {
		return nil
	}
	return fn.Call(args...)
}

// Create creates a new Callable for the method named method in the specified value.
func Create(value any, method string) Any { //gd:Callable.create
	if value == nil {
		return Any{}
	}
	return New(reflect.ValueOf(value).MethodByName(method).Interface())
}

// ArgumentCount returns the total number of arguments this Callable should take.
func ArgumentCount(fn Func) int { //gd:Callable.get_argument_count
	if fn == nil {
		return 0
	}
	argc, _ := fn.Args()
	return argc
}

// BoundArguments returns the arguments that have been bound to this Callable.
func BoundArguments(fn Func) Array.Contains[any] { //gd:Callable.get_bound_arguments
	_, binds := fn.Args()
	return binds
}

// BoundArgumentsCount returns the total amount of arguments bound (or unbound)
// via successive bind or unbind calls. If the amount of arguments unbound is
// greater than the ones bound, this function returns a value less than zero.
func BoundArgumentsCount(fn Func) int { //gd:Callable.get_bound_arguments_count
	if fn == nil {
		return 0
	}
	bound := BoundArguments(fn)
	return bound.Len()
}

// Method returns the name of the function represented by this Callable or
// an empty string if a name is not available.
func Method(fn Func) string { //gd:Callable.get_method
	if fn == nil {
		return ""
	}
	return fn.Name()
}

// IsProxy returns true if the given value is not backed by a Go function.
func IsProxy(fn Func) bool { //gd:Callable.is_custom
	if fn == nil {
		return true
	}
	return reflect.TypeOf(fn).Kind() != reflect.Func
}

// Hash returns the 32-bit hash value of this Callable's object.
//
// Note: Callables with equal content will always produce identical hash values. However,
// the reverse is not true. Returning identical hash values does not imply the
// callables are equal, because different callables can have identical hash values due
// to hash collisions. The engine uses a 32-bit hash algorithm for hash.
func Hash(fn Func) uint32 { //gd:Callable.hash
	return uint32(reflect.ValueOf(reflect.TypeOf(fn)).Pointer() % math.MaxUint32)
}

// Receiver returns the receiver of the method represented by this Callable.
// If no receiver is available, this function returns nil.
func Receiver(fn Func) any { //gd:Callable.get_object Callable.get_object_id
	if fn == nil {
		return nil
	}
	l, ok := fn.(*local)
	if !ok {
		return nil
	}
	return l.value
}

// IsNull returns true if the given value is nil.
func IsNull(fn Func) bool { //gd:Callable.is_null
	return fn == nil
}

// IsStandard returns true if the given value is backed by a Go function.
func IsStandard(fn Func) bool { //gd:Callable.is_standard
	if fn == nil {
		return false
	}
	return reflect.TypeOf(fn).Kind() == reflect.Func
}

// IsValid returns true if the given value is not nil.
func IsValid(fn Func) bool { //gd:Callable.is_valid
	return fn != nil
}

// Unbind returns a copy of this Callable with a number of arguments unbound. In other
// words, when the new callable is called the last few arguments supplied by the user
// are ignored, according to argcount. The remaining arguments are passed to the callable.
// This allows to use the original callable in a context that attempts to pass more
// arguments than this callable can handle, e.g. a signal with a fixed number of arguments.
// See also [Bind].
//
// Note: When this method is chained with other similar methods, the order in which the
// argument list is modified is read from right to left.
func Unbind(fn Func, argcount int) Func { //gd:Callable.unbind
	if fn == nil {
		return nil
	}
	return &unbind{Func: fn, count: argcount}
}

type unbind struct {
	Func
	count int
}

func (u *unbind) Call(args ...any) any {
	if u.Func == nil {
		return nil
	}
	if len(args) >= u.count {
		args = args[:len(args)-u.count]
	}
	return u.Func.Call(args...)
}
