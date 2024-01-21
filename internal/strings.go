//go:build !generate

package gd

import (
	"runtime.link/api/call"
	"runtime.link/mmm"
)

type String mmm.Pointer[API, String, uintptr]

func (s String) StringName(ctx Context) StringName {
	var frame = call.New()
	call.Arg(frame, mmm.Get(s))
	var r_ret = call.Ret[uintptr](frame)
	mmm.API(s).typeset.creation.StringName[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[StringName](ctx.Lifetime, ctx.API, raw)
}

// Copy returns a copy of the string that is owned by the provided context.
func (s String) Copy(ctx Context) String {
	var frame = call.New()
	call.Arg(frame, mmm.Get(s))
	var ret = call.Ret[uintptr](frame)
	ctx.API.typeset.creation.String[1](ret.Uintptr(), frame.Array(0))
	var raw = ret.Get()
	frame.Free()
	return mmm.New[String](ctx.Lifetime, ctx.API, raw)
}

func (s String) Free() {
	var frame = call.New()
	mmm.API(s).typeset.destruct.String(call.Arg(frame, mmm.End(s)).Uintptr())
	frame.Free()
}

func (s *String) Append(ctx Context, other String) {
	*s = mmm.API(*s).Strings.Append(ctx, *s, other)
}

func (s String) String() string {
	if mmm.Get(s) == 0 {
		return ""
	}
	return mmm.API(s).Strings.Get(s)
}

func (Godot *API) StringFromStringName(ctx Context, s StringName) String {
	var frame = call.New()
	call.Arg(frame, mmm.Get(s))
	var r_ret = call.Ret[uintptr](frame)
	Godot.typeset.creation.String[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[String](ctx.Lifetime, ctx.API, raw)
}

type StringName mmm.Pointer[API, StringName, uintptr]

func (Godot *API) StringNameFromString(ctx Context, s String) StringName {
	var frame = call.New()
	call.Arg(frame, mmm.Get(s))
	var r_ret = call.Ret[uintptr](frame)
	Godot.typeset.creation.StringName[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[StringName](ctx.Lifetime, ctx.API, raw)
}

func (s StringName) Free() {
	var frame = call.New()
	mmm.API(s).typeset.destruct.StringName(call.Arg(frame, mmm.End(s)).Uintptr())
	frame.Free()
}

func (s StringName) String() string {
	if mmm.Get(s) == 0 {
		return ""
	}
	ctx := NewContext(mmm.API(s))
	var tmp = mmm.API(s).StringFromStringName(ctx, s)
	defer ctx.End()
	return tmp.String()
}

type NodePath mmm.Pointer[API, NodePath, uintptr]

func (n NodePath) Free() {
	var frame = call.New()
	mmm.API(n).typeset.destruct.NodePath(call.Arg(frame, mmm.End(n)).Uintptr())
	frame.Free()
}
