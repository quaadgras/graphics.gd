//go:build !generate

package gd

import (
	"grow.graphics/gd/internal/callframe"
	"grow.graphics/gd/internal/discreet"
)

type IsSignal interface{ signal() }

func (s Signal) signal() {}

func (s Signal) Free() {
	if ptr, ok := discreet.End(s); ok {
		var frame = callframe.New()
		Global.typeset.destruct.Signal(callframe.Arg(frame, ptr).Uintptr())
		frame.Free()
	}
}

func NewSignalOf(object Object, signal StringName) Signal {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(object))
	callframe.Arg(frame, discreet.Get(signal))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.Signal[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return discreet.New[Signal](raw)
}
