//go:build !generate

package gd

import (
	"runtime.link/api/call"
	"runtime.link/mmm"
)

type Signal mmm.Pointer[API, Signal, [2]uintptr]

type IsSignal interface{ signal() }

func (s Signal) signal() {}

func (s Signal) Free() {
	var frame = call.New()
	mmm.API(s).typeset.destruct.Signal(call.Arg(frame, mmm.End(s)).Uintptr())
	frame.Free()
}

func (Godot *Context) SignalOf(ctx Context, object Object, signal StringName) Signal {
	var frame = call.New()
	call.Arg(frame, mmm.Get(object.AsPointer()))
	call.Arg(frame, mmm.Get(signal))
	var r_ret = call.Ret[[2]uintptr](frame)
	Godot.API.typeset.creation.Signal[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[Signal](ctx.Lifetime, ctx.API, raw)
}
