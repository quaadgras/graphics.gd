//go:build !generate

package gd

import (
	"unsafe"

	"grow.graphics/gd/internal/callframe"

	"runtime.link/mmm"
)

type String mmm.Pointer[API, String, uintptr]

func (s String) StringName(ctx Lifetime) StringName {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(s).typeset.creation.StringName[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[StringName](ctx.Lifetime, ctx.API, raw)
}

// Copy returns a copy of the string that is owned by the provided context.
func (s String) Copy(ctx Lifetime) String {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var ret = callframe.Ret[uintptr](frame)
	ctx.API.typeset.creation.String[1](ret.Uintptr(), frame.Array(0))
	var raw = ret.Get()
	frame.Free()
	return mmm.New[String](ctx.Lifetime, ctx.API, raw)
}

func (s String) Free() {
	var frame = callframe.New()
	mmm.API(s).typeset.destruct.String(callframe.Arg(frame, mmm.End(s)).Uintptr())
	frame.Free()
}

func (s *String) Append(ctx Lifetime, other String) {
	*s = mmm.API(*s).Strings.Append(ctx, *s, other)
}

func (s String) UnsafePointer() unsafe.Pointer {
	return mmm.API(s).Strings.UnsafePointer(s)
}

func (s String) Len() int { return int(s.Length()) }
func (s String) Cap() int { return int(s.Length()) }

func (s String) String() string {
	if mmm.Get(s) == 0 {
		return ""
	}
	return mmm.API(s).Strings.Get(s)
}

func (Godot *API) StringFromStringName(ctx Lifetime, s StringName) String {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[uintptr](frame)
	Godot.typeset.creation.String[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[String](ctx.Lifetime, ctx.API, raw)
}

type StringName mmm.Pointer[API, StringName, uintptr]

func (Godot *API) StringNameFromString(ctx Lifetime, s String) StringName {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[uintptr](frame)
	Godot.typeset.creation.StringName[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[StringName](ctx.Lifetime, ctx.API, raw)
}

func (s StringName) Free() {
	var frame = callframe.New()
	mmm.API(s).typeset.destruct.StringName(callframe.Arg(frame, mmm.End(s)).Uintptr())
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

func (s String) NodePath(ctx Lifetime) NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(s).typeset.creation.NodePath[2](r_ret.Uintptr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[NodePath](ctx.Lifetime, ctx.API, raw)
}

func (n NodePath) Free() {
	var frame = callframe.New()
	mmm.API(n).typeset.destruct.NodePath(callframe.Arg(frame, mmm.End(n)).Uintptr())
	frame.Free()
}
