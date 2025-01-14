//go:build !generate

package gd

import (
	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
)

func (s String) StringName() StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(s))
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Global.typeset.creation.StringName[2](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[StringName](raw)
}

// Copy returns a copy of the string that is owned by the provided context.
func (s String) Copy() String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(s))
	var ret = callframe.Ret[[1]EnginePointer](frame)
	Global.typeset.creation.String[1](ret.Addr(), frame.Array(0))
	var raw = ret.Get()
	frame.Free()
	return pointers.New[String](raw)
}

func (s String) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	var frame = callframe.New()
	Global.typeset.destruct.String(callframe.Arg(frame, ptr).Addr())
	frame.Free()
}

func (s *String) Append(other String) {
	modified := Global.Strings.Append(*s, other)
	pointers.End(*s)
	*s = modified
}

func (s String) Len() int { return int(s.Length()) }
func (s String) Cap() int { return int(s.Length()) }

func (s String) String() string {
	if pointers.Get(s) == ([1]EnginePointer{}) {
		return ""
	}
	return Global.Strings.Get(s)
}

func (Godot *API) StringFromStringName(s StringName) String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(s))
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Godot.typeset.creation.String[2](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[String](raw)
}

func (Godot *API) StringFromNodePath(s NodePath) String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(s))
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Godot.typeset.creation.String[3](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[String](raw)
}

func NewStringNameFromString(s String) StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(s))
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Global.typeset.creation.StringName[2](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[StringName](raw)
}

func (s StringName) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	var frame = callframe.New()
	Global.typeset.destruct.StringName(callframe.Arg(frame, ptr).Addr())
	frame.Free()
}

func (s StringName) String() string {
	if pointers.Get(s) == ([1]EnginePointer{}) {
		return ""
	}
	var tmp = Global.StringFromStringName(s)
	return tmp.String()
}

func (s String) NodePath() NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(s))
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Global.typeset.creation.NodePath[2](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[NodePath](raw)
}

func (n NodePath) String() string {
	return Global.StringFromNodePath(n).String()
}

func (n NodePath) Free() {
	ptr, ok := pointers.End(n)
	if !ok {
		return
	}
	var frame = callframe.New()
	Global.typeset.destruct.NodePath(callframe.Arg(frame, ptr).Addr())
	frame.Free()
}
