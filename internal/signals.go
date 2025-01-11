//go:build !generate

package gd

import (
	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
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
	var r_ret = callframe.Ret[SignalPointers](frame)
	Global.typeset.creation.Signal[2](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[Signal](raw)
}
