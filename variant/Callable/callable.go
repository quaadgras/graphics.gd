// Package Callable provides generic methods for working with callable functions.
package Callable

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sync"

	"graphics.gd/variant"
	"graphics.gd/variant/Array"
)

// Function represents a function. It can either be a method on a named type, or a custom callable used
// for different purposes
type Function struct {
	state complex128
	proxy Proxy
}

// Nil represents a nil function.
var Nil Function

type functionCall struct {
	function  Function
	arguments []variant.Any
}

// queue of functions to call later. A plain Go slice, NOT an engine Array:
// Defer runs on arbitrary goroutines while holding the mutex, and an engine
// Array append from a user goroutine is a cross-thread engine call serviced
// by the main thread's ring flush — which runs AFTER Callable.Cycle in the
// frame, and Cycle starts by taking this same mutex. Holding the mutex across
// that append therefore deadlocks the main thread against the deferring
// goroutine whenever the Defer misses the flush window (reliably on slower
// machines, e.g. the whole test suite hung at its second test on windows).
var queue []functionCall
var mutex sync.Mutex

// Cycle calls all functions in the defer queue.
//
// The pending batch is taken ownership of under the mutex and the mutex is
// released before any function runs: a Defer landing mid-cycle appends to the
// fresh queue and runs next frame (the previous iterate-then-Clear shape
// silently discarded such entries without calling them), and the functions
// themselves are free to Defer without re-entrancy trouble.
func Cycle() {
	mutex.Lock()
	batch := queue
	queue = nil
	mutex.Unlock()
	for _, queued := range batch {
		queued.function.Call(queued.arguments...)
	}
}

// New returns a new [Func] from the given value, if the value is not a Go func
// then it will be wrapped as if it were a function without any arguments that
// returns the specified value.
func New(value any) Function {
	if value == nil {
		return Function{}
	}
	if already, ok := value.(Function); ok {
		return already
	}
	rvalue := reflect.ValueOf(value)
	if rvalue.Kind() == reflect.Func && rvalue.IsNil() {
		return Function{}
	}
	return Function{
		proxy: &local{
			value: value,
		},
	}
}

// wrapObject fits an engine object reference to a class instance parameter.
// The engine hands an object to a Go function as a bare reference, while
// the function asks for an instance type: [Object.Instance] is a one-element
// array of that reference, and every class instance offers SetObject to take
// one. Without this, a func(Object.Instance) connected from Go panics in
// reflect.Call the first time the engine invokes it.
func wrapObject(value reflect.Value, param reflect.Type) (reflect.Value, bool) {
	if param.Kind() == reflect.Array && param.Len() == 1 && param.Elem() == value.Type() {
		wrapped := reflect.New(param).Elem()
		wrapped.Index(0).Set(value)
		return wrapped, true
	}
	setter, ok := reflect.PointerTo(param).MethodByName("SetObject")
	if !ok || setter.Type.NumIn() != 2 {
		return reflect.Value{}, false
	}
	arg := setter.Type.In(1)
	if arg.Kind() != reflect.Array || arg.Len() != 1 || !value.Type().ConvertibleTo(arg.Elem()) {
		return reflect.Value{}, false
	}
	packed := reflect.New(arg).Elem()
	packed.Index(0).Set(value.Convert(arg.Elem()))
	instance := reflect.New(param)
	setter.Func.Call([]reflect.Value{instance, packed})
	return instance.Elem(), true
}

type local struct {
	value any
	binds Array.Any
	cache sync.Pool
	trims int
	proxy Function
}

func (l *local) Name(complex128) string {
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(l.value).Pointer()).Name()
	}
	return "<unknown callable>"
}

func (l *local) Args(complex128) (int, Array.Any) {
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		return reflect.TypeOf(l.value).NumIn(), l.binds
	}
	return 0, l.binds
}

func (l *local) Call(_ complex128, args ...variant.Any) variant.Any {
	argc, binds := l.Args(0)
	if len(args)-binds.Len() != argc {
		panic(fmt.Sprintf("%s: invalid number of arguments, expected %d, found %d",
			runtime.FuncForPC(reflect.ValueOf(l.value).Pointer()).Name(), argc, len(args)-binds.Len()))
	}
	if reflect.TypeOf(l.value).Kind() == reflect.Func {
		ftype := reflect.TypeOf(l.value)
		values, ok := l.cache.Get().(Array.Contains[reflect.Value])
		if !ok {
			values = Array.New[reflect.Value]()
		}
		defer l.cache.Put(values)
		values.Resize(len(args) + binds.Len())
		for i := range args {
			values.SetIndex(i, reflect.ValueOf(args[i].Interface()))
		}
		for i, v := range binds.Iter() {
			values.SetIndex(i+len(args), reflect.ValueOf(v.Interface()))
		}
		for i := range ftype.NumIn() {
			value := values.Index(i)
			if value.Type().Implements(reflect.TypeFor[interface {
				ConvertTo(reflect.Type) (reflect.Value, error)
			}]()) {
				converted, err := value.Interface().(interface {
					ConvertTo(reflect.Type) (reflect.Value, error)
				}).ConvertTo(ftype.In(i))
				if err != nil {
					panic(err)
				}
				values.SetIndex(i, converted)
			} else if value.Type() != ftype.In(i) {
				if value.Type().ConvertibleTo(ftype.In(i)) {
					values.SetIndex(i, value.Convert(ftype.In(i)))
				} else if wrapped, ok := wrapObject(value, ftype.In(i)); ok {
					values.SetIndex(i, wrapped)
				}
			}
		}
		result := reflect.ValueOf(l.value).Call(values.Slice())
		if len(result) == 0 {
			return variant.Nil
		}
		return variant.New(result[0].Interface())
	}
	return variant.New(l.value)
}

func (l *local) Bind(_ complex128, args ...variant.Any) (Proxy, complex128) {
	result := local{}
	result.binds = Array.Duplicate(l.binds)
	for _, arg := range args {
		result.binds.Append(arg)
	}
	return &result, 0
}

// Bind returns a copy of this Callable with one or more arguments bound. When called,
// the bound arguments are passed after the arguments supplied by call. See also [Unbind].
//
// Note: When this method is chained with other similar methods, the order in which the
// argument list is modified is read from right to left.
func Bind(fn Function, args ...variant.Any) Function { //gd:Callable.bind Callable.bindv
	if fn == Nil {
		return Nil
	}
	return Through(fn.proxy.Bind(fn.state, args...))
}

// Call calls the method represented by this Callable. Arguments can be passed and should
// match the method's signature.
func (fn Function) Call(args ...variant.Any) variant.Any { //gd:Callable.call Callable.callv Callable.rpc Callable.rpc_id
	if fn == Nil {
		return variant.Nil
	}
	return fn.proxy.Call(fn.state, args...)
}

// Defer calls the function represented by this callable at the end of the current frame.
// Arguments can be passed and should match the method's signature. This is safe to call
// from multiple goroutines.
func Defer(fn Function, args ...variant.Any) { //gd:Callable.call_deferred
	if fn == Nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	queue = append(queue, functionCall{
		function:  fn,
		arguments: args,
	})
}

// Create creates a new Callable for the method named method in the specified value.
func Create(value any, method string) Function { //gd:Callable.create
	if value == nil {
		return Nil
	}
	return New(reflect.ValueOf(value).MethodByName(method).Interface())
}

// ArgumentCount returns the total number of arguments this Callable should take.
func ArgumentCount(fn Function) int { //gd:Callable.get_argument_count
	if fn == Nil {
		return 0
	}
	argc, _ := fn.proxy.Args(fn.state)
	return argc
}

// BoundArguments returns the arguments that have been bound to this Callable.
func BoundArguments(fn Function) Array.Any { //gd:Callable.get_bound_arguments
	_, binds := fn.proxy.Args(fn.state)
	return binds
}

// BoundArgumentsCount returns the total amount of arguments bound (or unbound)
// via successive bind or unbind calls. If the amount of arguments unbound is
// greater than the ones bound, this function returns a value less than zero.
func BoundArgumentsCount(fn Function) int { //gd:Callable.get_bound_arguments_count
	if fn == Nil {
		return 0
	}
	bound := BoundArguments(fn)
	return bound.Len()
}

// UnboundArgumentsCount returns returns the total amount of arguments unbound via
// successive bind or unbind calls.
func UnboundArgumentsCount(fn Function) int { //gd:Callable.get_unbound_arguments_count
	if fn == Nil {
		return 0
	}
	argc, binds := fn.proxy.Args(fn.state)
	return argc - binds.Len()
}

// Method returns the name of the function represented by this Callable or
// an empty string if a name is not available.
func Method(fn Function) string { //gd:Callable.get_method
	if fn == Nil {
		return ""
	}
	return fn.proxy.Name(fn.state)
}

// IsProxy returns true if the given value is not backed by a Go function.
func IsProxy(fn Function) bool { //gd:Callable.is_custom
	if fn == Nil {
		return true
	}
	return reflect.TypeFor[Function]().Kind() != reflect.Func
}

// Hash returns the 32-bit hash value of this Callable's object.
//
// Note: Callables with equal content will always produce identical hash values. However,
// the reverse is not true. Returning identical hash values does not imply the
// callables are equal, because different callables can have identical hash values due
// to hash collisions. The engine uses a 32-bit hash algorithm for hash.
func Hash(fn Function) uint32 { //gd:Callable.hash
	return uint32(reflect.ValueOf(reflect.TypeFor[Function]()).Pointer() % math.MaxUint32)
}

// Receiver returns the receiver of the method represented by this Callable.
// If no receiver is available, this function returns nil.
func Receiver(fn Function) any { //gd:Callable.get_object Callable.get_object_id
	if fn == Nil {
		return nil
	}
	l, ok := fn.proxy.(*local)
	if !ok {
		return nil
	}
	return l.value
}

// IsNil returns true if the given value is nil.
func IsNil(fn Function) bool { //gd:Callable.is_null
	return fn == Nil
}

// IsStandard returns true if the given value is backed by a Go function.
func IsStandard(fn Function) bool { //gd:Callable.is_standard
	if fn == Nil {
		return false
	}
	return reflect.TypeFor[Function]().Kind() == reflect.Func
}

// IsValid returns true if the given value is not nil.
func IsValid(fn Function) bool { //gd:Callable.is_valid
	return fn != Nil
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
func Unbind(fn Function, argcount int) Function { //gd:Callable.unbind
	if fn == Nil {
		return Nil
	}
	return Function{
		proxy: &unbind{Proxy: fn.proxy, count: argcount},
	}
}

type unbind struct {
	Proxy
	count int
}

func (u *unbind) Call(state complex128, args ...variant.Any) variant.Any {
	if u.Proxy == nil {
		return variant.Nil
	}
	if len(args) >= u.count {
		args = args[:len(args)-u.count]
	}
	return u.Proxy.Call(state, args...)
}
