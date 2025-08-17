//go:build !generate

package gd

import (
	"unsafe"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Path"
	StringType "graphics.gd/variant/String"
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

func (s String) Len() int { return int(s.Length()) }
func (s String) Cap() int { return int(s.Length()) }

func (s String) String() string {
	if pointers.Get(s) == ([1]EnginePointer{}) {
		return ""
	}
	var buf = make([]byte, s.Length())
	gdextension.Host.Strings.Encode.UTF8(gdextension.String(pointers.Get(s)[0]), buf)
	return unsafe.String(&buf[0], len(buf))
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

func (n NodePath) InternalString() String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(n))
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Global.typeset.creation.String[2](r_ret.Addr(), frame.Array(0))
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[String](raw)
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

func InternalString(s StringType.Readable) String {
	_, ptr := StringType.Proxy(s, StringCacheCheck, NewStringProxy)
	return pointers.Load[String](ptr)
}

func StringCacheCheck(_ StringProxy, raw complex128) bool { return true }

func NewStringProxy() (StringProxy, complex128) {
	return StringProxy{}, pointers.Pack(NewString(""))
}

type StringProxy struct{}

func (StringProxy) Len(raw complex128) int {
	return pointers.Load[String](raw).Len()
}
func (StringProxy) Slice(raw complex128, index int, close int) StringType.Readable {
	s := pointers.Load[String](raw)
	s = s.Substr(Int(index), Int(close))
	return StringType.Via(StringProxy{}, pointers.Pack(s))
}
func (StringProxy) String(raw complex128) string {
	return pointers.Load[String](raw).String()
}
func (StringProxy) Index(raw complex128, n int) byte {
	return byte(gdextension.Host.Strings.Access(gdextension.String(pointers.Get(pointers.Load[String](raw))[0]), n))
}
func (StringProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Readable) {
	s := pointers.Load[String](raw)
	next := s.Substr(0, s.Length())
	return StringType.Rune(gdextension.Host.Strings.Access(gdextension.String(pointers.Get(pointers.Load[String](raw))[0]), 0)), 0, StringType.Via(StringProxy{}, pointers.Pack(next))
}
func (StringProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Readable {
	s := gdextension.String(pointers.Get(pointers.Load[String](raw))[0])
	s = gdextension.Host.Strings.Append.Rune(s, rune(r))
	return StringType.Via(StringProxy{}, pointers.Pack(pointers.New[String]([1]EnginePointer{EnginePointer(s)})))
}
func (StringProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Readable {
	s := gdextension.String(pointers.Get(pointers.Load[String](raw))[0])
	s2 := gdextension.String(pointers.Get(pointers.Load[String](raw2))[0])
	s = gdextension.Host.Strings.Append.String(s, s2)
	return StringType.Via(StringProxy{}, pointers.Pack(pointers.New[String]([1]EnginePointer{EnginePointer(s)})))
}
func (StringProxy) AppendString(raw complex128, str string) StringType.Readable {
	s := gdextension.String(pointers.Get(pointers.Load[String](raw))[0])
	s2 := gdextension.String(pointers.Get(NewString(str))[0])
	s = gdextension.Host.Strings.Append.String(s, s2)
	return StringType.Via(StringProxy{}, pointers.Pack(pointers.New[String]([1]EnginePointer{EnginePointer(s)})))
}
func (StringProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	return int(pointers.Load[String](raw).CasecmpTo(pointers.Load[String](raw2)))
}

func InternalNodePath(s Path.ToNode) NodePath {
	_, ptr := StringType.Proxy(s, NodePathCheck, NewNodePathProxy)
	return pointers.Load[NodePath](ptr)
}

func NodePathCheck(_ NodePathProxy, raw complex128) bool { return true }

func NewNodePathProxy() (NodePathProxy, complex128) {
	return NodePathProxy{}, pointers.Pack(NewString("").NodePath())
}

type NodePathProxy struct{}

func (NodePathProxy) Len(raw complex128) int {
	return pointers.Load[NodePath](raw).InternalString().Len()
}
func (NodePathProxy) Slice(raw complex128, index int, close int) StringType.Readable {
	s := pointers.Load[NodePath](raw)
	s = s.InternalString().Substr(Int(index), Int(close)).NodePath()
	return StringType.Via(StringProxy{}, pointers.Pack(s))
}
func (NodePathProxy) String(raw complex128) string {
	return pointers.Load[NodePath](raw).String()
}
func (NodePathProxy) Index(raw complex128, n int) byte {
	return byte(gdextension.Host.Strings.Access(gdextension.String(pointers.Get(pointers.Load[NodePath](raw).InternalString())[0]), n))
}
func (NodePathProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Readable) {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	next := str.Substr(0, 1).NodePath()
	return StringType.Rune(gdextension.Host.Strings.Access(gdextension.String(pointers.Get(str)[0]), 0)), 0, StringType.Via(StringProxy{}, pointers.Pack(next))
}
func (NodePathProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Readable {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	s = pointers.New[String]([1]EnginePointer{EnginePointer(gdextension.Host.Strings.Append.Rune(gdextension.String(pointers.Get(str)[0]), rune(r)))}).NodePath()
	return StringType.Via(NodePathProxy{}, pointers.Pack(s))
}
func (NodePathProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Readable {
	s := pointers.Load[NodePath](raw).InternalString()
	s2 := pointers.Load[NodePath](raw2).InternalString()
	gdextension.Host.Strings.Append.String(gdextension.String(pointers.Get(s)[0]), gdextension.String(pointers.Get(s2)[0]))
	r := pointers.New[String]([1]EnginePointer{EnginePointer(gdextension.Host.Strings.Append.String(gdextension.String(pointers.Get(s)[0]), gdextension.String(pointers.Get(s2)[0])))}).NodePath()
	return StringType.Via(NodePathProxy{}, pointers.Pack(r))
}
func (NodePathProxy) AppendString(raw complex128, str string) StringType.Readable {
	s := pointers.Load[NodePath](raw).InternalString()
	s2 := NewString(str)
	gdextension.Host.Strings.Append.String(gdextension.String(pointers.Get(s)[0]), gdextension.String(pointers.Get(s2)[0]))
	r := pointers.New[String]([1]EnginePointer{EnginePointer(gdextension.Host.Strings.Append.String(gdextension.String(pointers.Get(s)[0]), gdextension.String(pointers.Get(s2)[0])))}).NodePath()
	return StringType.Via(NodePathProxy{}, pointers.Pack(r))
}
func (NodePathProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	return int(pointers.Load[NodePath](raw).InternalString().CasecmpTo(pointers.Load[NodePath](raw2).InternalString()))
}

func InternalStringName(s StringType.Name) StringName {
	_, ptr := StringType.Proxy(s, StringNameCheck, NewStringNameProxy)
	return pointers.Load[StringName](ptr)
}

func StringNameCheck(_ StringNameProxy, raw complex128) bool { return true }

func NewStringNameProxy() (StringNameProxy, complex128) {
	return StringNameProxy{}, pointers.Pack(NewStringName(""))
}

type StringNameProxy struct{}

func (StringNameProxy) Len(raw complex128) int {
	return int(pointers.Load[StringName](raw).Length())
}
func (StringNameProxy) Slice(raw complex128, index int, close int) StringType.Readable {
	s := pointers.Load[StringName](raw)
	s = s.Substr(Int(index), Int(close)).StringName()
	return StringType.Via(StringNameProxy{}, pointers.Pack(s))
}
func (StringNameProxy) String(raw complex128) string {
	return pointers.Load[StringName](raw).String()
}
func (StringNameProxy) Index(raw complex128, n int) byte {
	name := pointers.Load[StringName](raw)
	s := name.Substr(0, name.Length())
	return byte(gdextension.Host.Strings.Access(gdextension.String(pointers.Get(s)[0]), n))
}
func (StringNameProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Readable) {
	s := pointers.Load[StringName](raw)
	next := s.Substr(0, 1).StringName()
	return StringType.Rune(gdextension.Host.Strings.Access(gdextension.String(pointers.Get(s)[0]), 0)), 0, StringType.Via(StringNameProxy{}, pointers.Pack(next))
}
func (StringNameProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Readable {
	s := pointers.Load[StringName](raw)
	str := s.Substr(0, s.Length())
	result := pointers.New[String]([1]EnginePointer{EnginePointer(gdextension.Host.Strings.Append.Rune(gdextension.String(pointers.Get(str)[0]), rune(r)))})
	return StringType.Via(StringNameProxy{}, pointers.Pack(result.StringName()))
}
func (StringNameProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Readable {
	s := pointers.Load[StringName](raw)
	s2 := pointers.Load[StringName](raw2)
	sub := s.Substr(0, s.Length())
	result := pointers.New[String]([1]EnginePointer{EnginePointer(gdextension.Host.Strings.Append.String(gdextension.String(pointers.Get(sub)[0]), gdextension.String(pointers.Get(s2)[0])))})
	return StringType.Via(StringNameProxy{}, pointers.Pack(result.StringName()))
}
func (StringNameProxy) AppendString(raw complex128, str string) StringType.Readable {
	s := pointers.Load[StringName](raw)
	sub := s.Substr(0, s.Length())
	result := pointers.New[String]([1]EnginePointer{EnginePointer(gdextension.Host.Strings.Append.String(gdextension.String(pointers.Get(sub)[0]), gdextension.String(pointers.Get(NewString(str))[0])))})
	return StringType.Via(StringNameProxy{}, pointers.Pack(result.StringName()))
}
func (StringNameProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	other := pointers.Load[StringName](raw2)
	return int(pointers.Load[StringName](raw).CasecmpTo(other.Substr(0, other.Length())))
}
