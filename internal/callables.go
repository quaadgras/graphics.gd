//go:build !generate

package gd

import (
	"reflect"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	CallableType "graphics.gd/variant/Callable"
)

func init() {
	gdextension.On.Callables = gdextension.CallbacksForCallables{
		Call: func(fn gdextension.FunctionID, result gdextension.Returns[gdextension.Variant], arg_count int, args gdextension.Accepts[gdextension.Variant], call_error gdextension.Returns[gdextension.CallError]) {
			value := cgoHandle(fn).Value()
			switch cb := value.(type) {
			case func():
				cb()
				gdmemory.Set(gdextension.Pointer(result), gdextension.Variant{})
				gdmemory.Set(gdextension.Pointer(call_error), CallError{})
				return
			}
			vargs := make([]reflect.Value, min(arg_count, 16))
			rtype := reflect.TypeOf(value)
			for i := range arg_count {
				var to_type reflect.Type
				if rtype.IsVariadic() && i >= rtype.NumIn()-1 {
					to_type = rtype.In(rtype.NumIn() - 1).Elem()
				} else {
					if i >= rtype.NumIn() {
						gdmemory.Set(gdextension.Pointer(call_error), gdextension.CallError{
							Type:     gdextension.CallTooManyArguments,
							Expected: int32(rtype.NumIn()),
						})
						return
					}
					to_type = rtype.In(i)
				}
				var err error
				vargs[i], err = ConvertToDesiredGoType(pointers.Let[Variant](gdmemory.IndexVariants(args, arg_count, i)), to_type)
				if err != nil {
					vtype, _ := VariantTypeOf(rtype.In(i))
					gdmemory.Set(gdextension.Pointer(call_error), gdextension.CallError{
						Type:     gdextension.CallInvalidArguments,
						Argument: int32(i),
						Expected: int32(vtype),
					})
					return
				}
			}
			if len(vargs) < rtype.NumIn() && (!rtype.IsVariadic() && len(vargs) == rtype.NumIn()-1) {
				gdmemory.Set(gdextension.Pointer(call_error), gdextension.CallError{
					Type:     gdextension.CallTooFewArguments,
					Expected: int32(rtype.NumIn()),
				})
				return
			}
			results := reflect.ValueOf(value).Call(vargs)
			if len(results) > 0 {
				raw, _ := pointers.End(CutVariant(results[0].Interface(), true))
				gdmemory.Set(gdextension.Pointer(result), raw)
			}
			gdmemory.Set(gdextension.Pointer(call_error), CallError{})
		},
		Hash: func(fn gdextension.FunctionID) uint32 {
			return uint32(reflect.ValueOf(cgoHandle(fn).Value()).Pointer())
		},
		Stringify: func(fn gdextension.FunctionID, err gdextension.Returns[gdextension.CallError]) gdextension.String {
			s := NewString(reflect.ValueOf(cgoHandle(fn).Value()).String())
			raw, _ := pointers.End(s)
			return raw
		},
		ArgumentCount: func(fn gdextension.FunctionID, err gdextension.Returns[gdextension.CallError]) int {
			value := reflect.ValueOf(cgoHandle(fn).Value())
			return value.Type().NumIn()
		},
		Validation: func(fn gdextension.FunctionID) bool {
			return cgoHandle(fn).Value() != nil
		},
		Compare: func(fn, other gdextension.FunctionID) bool {
			return reflect.ValueOf(cgoHandle(fn).Value()).Pointer() == reflect.ValueOf(cgoHandle(other).Value()).Pointer()
		},
		LessThan: func(fn, other gdextension.FunctionID) bool {
			return reflect.ValueOf(cgoHandle(fn).Value()).Pointer() < reflect.ValueOf(cgoHandle(other).Value()).Pointer()
		},
		Free: func(fn gdextension.FunctionID) {
			cgoHandle(fn).Delete()
		},
	}
}

// Callable creates a new callable out of the given function which must only accept
// godot-compatible types and return up to one godot-compatible type.
func NewCallable(fn any) Callable {
	var result gdextension.Callable
	gdextension.Host.Callables.Create(gdextension.CallableID(cgoNewHandle(fn)), 0, gdextension.CallReturns[gdextension.Callable](&result))
	return pointers.New[Callable](result)
}

func InternalCallable(fn CallableType.Function) Callable {
	return NewCallable(fn.Call)
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
