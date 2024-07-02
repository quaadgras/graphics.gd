//go:build !generate

package gd

import (
	"grow.graphics/gd/internal/callframe"

	"runtime.link/mmm"
)

type Signal mmm.Pointer[API, Signal, [2]uintptr]

type IsSignal interface{ signal() }

func (s Signal) signal() {}

func (s Signal) Free() {
	var frame = callframe.New()
	mmm.API(s).typeset.destruct.Signal(callframe.Arg(frame, mmm.End(s)).Uintptr())
	frame.Free()
}

func (Godot *Lifetime) SignalOf(ctx Lifetime, object Object, signal StringName) Signal {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(object.AsPointer()))
	callframe.Arg(frame, mmm.Get(signal))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Godot.API.typeset.creation.Signal[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[Signal](ctx.Lifetime, ctx.API, raw)
}
