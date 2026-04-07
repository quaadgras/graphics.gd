//go:build !generate

package gd

import (
	"hash/maphash"
	"reflect"

	gdunsafe "graphics.gd"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/threadsafe"
	"graphics.gd/variant"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	CallableType "graphics.gd/variant/Callable"
)

var callables threadsafe.Handles[comparableCallable, gdextension.FunctionID]
var callables_hash = maphash.MakeSeed()

// comparableCallable wraps a Go function together with a stable identity,
// so that two Godot callables created from the same [CallableType.Function]
// can be recognized as equal even though they have different CGO handles.
type comparableCallable struct {
	fn any
	id any // comparable
}

func (c comparableCallable) Call(args gdunsafe.Variants) (gdunsafe.Variant, gdunsafe.Error) {
	defer Recover()
	switch cb := c.fn.(type) {
	case func():
		cb()
		return gdunsafe.Variant{}, gdunsafe.Error{}
	case func() int:
		raw, _ := pointers.End(CutVariant(cb(), true))
		return raw, gdunsafe.Error{}
	}
	vargs := make([]reflect.Value, min(args.Len(), 16))
	rtype := reflect.TypeOf(c.fn)
	for i := range args.Len() {
		var to_type reflect.Type
		if rtype.IsVariadic() && i >= rtype.NumIn()-1 {
			to_type = rtype.In(rtype.NumIn() - 1).Elem()
		} else {
			if i >= rtype.NumIn() {
				return gdunsafe.Variant{}, args.ExpectedLen(rtype.NumIn())
			}
			to_type = rtype.In(i)
		}
		var err error
		vargs[i], err = ConvertToDesiredGoType(pointers.Let[Variant](args.Index(i)), to_type)
		if err != nil {
			vtype, _ := VariantTypeOf(rtype.In(i))
			return args.ExpectedArg(i, variant.Type(vtype))
		}
	}
	if len(vargs) < rtype.NumIn() && (!rtype.IsVariadic() && len(vargs) == rtype.NumIn()-1) {
		return gdunsafe.Variant{}, args.ExpectedLen(rtype.NumIn())
	}
	results := reflect.ValueOf(c.fn).Call(vargs)
	if len(results) > 0 {
		raw, _ := pointers.End(CutVariant(results[0].Interface(), true))
		return raw, gdunsafe.Error{}
	}
	return gdunsafe.Variant{}, gdunsafe.Error{}
}

func (c comparableCallable) IsValid() bool {
	return c.fn != nil
}

func (c comparableCallable) Hash() uint32 {
	return uint32(maphash.Comparable(callables_hash, c.id))
}

func (c comparableCallable) UnsafeString() gdunsafe.String {
	s := NewString(reflect.ValueOf(c.fn).String())
	raw, _ := pointers.End(s)
	return gdunsafe.String(raw[0])
}

func (c comparableCallable) ArgumentCount() int {
	return reflect.TypeOf(c.fn).NumIn()
}

func (c comparableCallable) Compare(other gdunsafe.ExtensionCallable) int {
	if o, ok := other.(comparableCallable); ok {
		if c.id != nil && o.id != nil && c.id == o.id {
			return 0
		}
	}
	if c.Hash() < other.Hash() {
		return -1
	}
	return 1
}

// NewCallable creates a new callable out of the given function which must only accept
// godot-compatible types and return up to one godot-compatible type.
func NewCallable(fn any) Callable {
	var result = gdextension.Callable(gdunsafe.MakeCallable(comparableCallable{fn: fn}, 0))
	return pointers.New[Callable](result)
}

// NewComparableCallable creates a new comparable callable from the given function which must only accept
// godot-compatible types and return up to one godot-compatible type.
func NewComparableCallable[T comparable](fn any, id T) Callable {
	var result = gdextension.Callable(gdunsafe.MakeCallable(comparableCallable{fn: fn, id: id}, 0))
	return pointers.New[Callable](result)
}

func InternalCallable(fn CallableType.Function) Callable {
	if fn == (CallableType.Function{}) {
		return Callable{}
	}
	return NewComparableCallable(fn.Call, fn)
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

func CallableAs[T any](callable Callable) T {
	fn, _ := reflect.TypeAssert[T](reflect.MakeFunc(reflect.TypeFor[T](), func(args []reflect.Value) (results []reflect.Value) {
		vargs := make([]Variant, len(args))
		for i, arg := range args {
			vargs[i] = NewVariant(arg.Interface())
		}
		result := callable.Call(vargs...)
		if reflect.TypeFor[T]().NumOut() == 0 {
			return nil
		}
		converted, err := ConvertToDesiredGoType(result, reflect.TypeFor[T]().Out(0))
		if err != nil {
			panic(err)
		}
		return []reflect.Value{converted}
	}))
	return fn
}
