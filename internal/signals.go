//go:build !generate

package gd

import (
	"iter"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/threadcheck"
	VariantPkg "graphics.gd/variant"
	CallableType "graphics.gd/variant/Callable"
	DictionaryType "graphics.gd/variant/Dictionary"
	ErrorType "graphics.gd/variant/Error"
	SignalType "graphics.gd/variant/Signal"
	StringType "graphics.gd/variant/String"
)

func (s Signal) Free() {
	if ptr, ok := pointers.End(s); ok {
		gdunsafe.Free(gdunsafe.Signal(ptr))
	}
}

func NewSignalOf(object [1]gdreference.Object, signal gdunsafe.StringName) gdunsafe.Signal {
	return builtin.creation.Signal[2](gdunsafe.ShapeObject<<4|gdunsafe.ShapeStringName<<8, unsafe.Pointer(&struct {
		object gdunsafe.Object
		signal gdunsafe.StringName
	}{
		object: gdreference.GetObject(gdreference.Object(object[0])),
		signal: signal,
	}))
}

func InternalSignal(signal SignalType.Any) gdunsafe.Signal {
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
	return ToError(ErrorType.Code(sig.Connect(InternalCallable(fn), Int(flags))))
}
func (SignalProxy) Remove(raw complex128, fn CallableType.Function) {
	sig := pointers.Load[Signal](raw)
	sig.Disconnect(InternalCallable(fn))
}
func (SignalProxy) Name(raw complex128) StringType.Unicode {
	sig := pointers.Load[Signal](raw)
	return StringType.New(sig.GetName().String())
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
	if threadcheck.Main() {
		sig := pointers.Load[Signal](raw)
		vargs := make([]Variant, len(values))
		for i, v := range values {
			vargs[i] = NewVariant(v)
		}
		sig.Emit(vargs...)
		return
	}
	CallableType.Defer(CallableType.New(func() {
		sig := pointers.Load[Signal](raw)
		vargs := make([]Variant, len(values))
		for i, v := range values {
			vargs[i] = NewVariant(v)
		}
		sig.Emit(vargs...)
	}))
}
func (SignalProxy) Emitter(raw complex128) VariantPkg.Any {
	return VariantPkg.New(gdreference.GetObject(gdreference.Object(pointers.Load[Signal](raw).GetObject())))
}
