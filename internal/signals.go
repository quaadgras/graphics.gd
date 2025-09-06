//go:build !generate

package gd

import (
	"iter"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	CallableType "graphics.gd/variant/Callable"
	DictionaryType "graphics.gd/variant/Dictionary"
	ErrorType "graphics.gd/variant/Error"
	SignalType "graphics.gd/variant/Signal"
	StringType "graphics.gd/variant/String"
)

func (s Signal) Free() {
	if ptr, ok := pointers.End(s); ok {
		gdextension.Free(gdextension.TypeSignal, &ptr)
	}
}

func NewSignalOf(object [1]Object, signal StringName) Signal {
	return pointers.New[Signal](gdextension.Make[gdextension.Signal](builtin.creation.Signal[2], gdextension.SizeObject<<4|gdextension.SizeStringName<<8, unsafe.Pointer(&struct {
		object gdextension.Object
		signal gdextension.StringName
	}{
		object: gdextension.Object(pointers.Get(object[0])[0]),
		signal: pointers.Get(signal),
	})))
}

func InternalSignal(signal SignalType.Any) Signal {
	_, state := SignalType.Proxy(signal, NewSignalCheck, NewSignalProxy)
	return pointers.Load[Signal](state)
}

type SignalProxy struct{}

func NewSignalProxy() (SignalProxy, complex128) {
	panic("NewSignalProxy: not implemented")
}

func NewSignalCheck(SignalProxy, complex128) bool {
	return true
}

func (SignalProxy) Attach(raw complex128, fn CallableType.Function, flags SignalType.Flags) error {
	sig := pointers.Load[Signal](raw)
	return ErrorType.Code(sig.Connect(InternalCallable(fn), Int(flags)))
}
func (SignalProxy) Remove(raw complex128, fn CallableType.Function) {
	sig := pointers.Load[Signal](raw)
	sig.Disconnect(InternalCallable(fn))
}
func (SignalProxy) Name(raw complex128) StringType.Readable {
	sig := pointers.Load[Signal](raw)
	return StringType.New(sig.GetName())
}
func (SignalProxy) Consumers(raw complex128) iter.Seq[SignalType.Consumer] {
	sig := pointers.Load[Signal](raw)
	return func(fn func(SignalType.Consumer) bool) {
		for _, conn := range sig.GetConnections().Iter() {
			if !fn(DictionaryAs[SignalType.Consumer](conn.Interface().(DictionaryType.Any))) {
				break
			}
		}
	}
}
func (SignalProxy) Emit(raw complex128, values ...VariantPkg.Any) {
	sig := pointers.Load[Signal](raw)
	vargs := make([]Variant, len(values))
	for i, v := range values {
		vargs[i] = NewVariant(v)
	}
	sig.Emit(vargs...)
}
func (SignalProxy) Emitter(raw complex128) VariantPkg.Any {
	return VariantPkg.New(pointers.Get(pointers.Load[Signal](raw).GetObject()))
}
