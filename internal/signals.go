//go:build !generate

package gd

import (
	"iter"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	CallableType "graphics.gd/variant/Callable"
	DictionaryType "graphics.gd/variant/Dictionary"
	ErrorType "graphics.gd/variant/Error"
	SignalType "graphics.gd/variant/Signal"
	StringType "graphics.gd/variant/String"
)

type IsSignal interface{ setSignal(Signal) }

func SetSignal(signal IsSignal, val Signal) {
	signal.setSignal(val)
}

func (s Signal) signal() {}

func (s *Signal) setSignal(val Signal) {
	*s = val
}

func (s Signal) Free() {
	if ptr, ok := pointers.End(s); ok {
		var frame = callframe.New()
		Global.typeset.destruct.Signal(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func NewSignalOf(object [1]Object, signal StringName) Signal {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(object[0]))
	callframe.Arg(frame, pointers.Get(signal))
	var r_ret = callframe.Ret[[2]uint64](frame)
	Global.typeset.creation.Signal[2](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[Signal](raw)
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
