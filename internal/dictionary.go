package gd

import (
	"grow.graphics/gd/internal/callframe"
	"runtime.link/mmm"
)

type Dictionary mmm.Pointer[API, Dictionary, uintptr]

func (d Dictionary) Index(ctx Context, key Variant) Variant {
	return mmm.API(d).Dictionary.Index(ctx, d, key)
}

func (d Dictionary) SetIndex(key Variant, value Variant) {
	mmm.API(d).Dictionary.SetIndex(d, key, value)
}

func (d Dictionary) Free() {
	var frame = callframe.New()
	mmm.API(d).typeset.destruct.Dictionary(callframe.Arg(frame, mmm.End(d)).Uintptr())
	frame.Free()
}

func (godot Context) Dictionary() Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	godot.API.typeset.creation.Dictionary[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[Dictionary](godot.Lifetime, godot.API, raw)
}
