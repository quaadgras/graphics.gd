//go:build !generate

package gd

import (
	"grow.graphics/gd/internal/callframe"
	"grow.graphics/gd/internal/pointers"
)

type IsSignal interface{ signal() }

func (s Signal) signal() {}

func (s Signal) Free() {
	if ptr, ok := pointers.End(s); ok {
		var frame = callframe.New()
		Global.typeset.destruct.Signal(callframe.Arg(frame, ptr).Uintptr())
		frame.Free()
	}
}

func NewSignalOf(object Object, signal StringName) Signal {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(object))
	callframe.Arg(frame, pointers.Get(signal))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.Signal[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[Signal](raw)
}
