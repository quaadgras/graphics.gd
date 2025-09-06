//go:build !generate

package gd

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Path"
	StringType "graphics.gd/variant/String"
)

func (s String) StringName() StringName {
	var arg = pointers.Get(s)
	return pointers.New[StringName](gdextension.Make[gdextension.StringName](builtin.creation.StringName[2], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

// Copy returns a copy of the string that is owned by the provided context.
func (s String) Copy() String {
	var arg = pointers.Get(s)
	return pointers.New[String](gdextension.Make[gdextension.String](builtin.creation.String[1], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

func (s String) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	gdextension.Free(gdextension.TypeString, &ptr)
}

func (s String) Len() int { return int(s.Length()) }
func (s String) Cap() int { return int(s.Length()) }

func (s String) String() string {
	if pointers.Get(s) == (gdextension.String{}) {
		return ""
	}
	if s.Length() == 0 {
		return ""
	}
	var buf = make([]byte, s.Length())
	gdextension.Host.Strings.Encode.UTF8(pointers.Get(s), buf)
	return unsafe.String(&buf[0], len(buf))
}

func StringFromStringName(s StringName) String {
	var arg = pointers.Get(s)
	return pointers.New[String](gdextension.Make[gdextension.String](builtin.creation.String[2], gdextension.SizeStringName<<4, unsafe.Pointer(&arg)))
}

func StringFromNodePath(s NodePath) String {
	var arg = pointers.Get(s)
	return pointers.New[String](gdextension.Make[gdextension.String](builtin.creation.String[3], gdextension.SizeNodePath<<4, unsafe.Pointer(&arg)))
}

func NewStringNameFromString(s String) StringName {
	var arg = pointers.Get(s)
	return pointers.New[StringName](gdextension.Make[gdextension.StringName](builtin.creation.StringName[2], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

func (s StringName) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	if ptr == (gdextension.StringName{}) {
		return
	}
	gdextension.Free(gdextension.TypeStringName, &ptr)
}

func (s StringName) String() string {
	if pointers.Get(s) == (gdextension.StringName{}) {
		return ""
	}
	var tmp = StringFromStringName(s)
	return tmp.String()
}

func (s String) NodePath() NodePath {
	var arg = pointers.Get(s)
	return pointers.New[NodePath](gdextension.Make[gdextension.NodePath](builtin.creation.NodePath[2], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

func (n NodePath) InternalString() String {
	var ptr = pointers.Get(n)
	return pointers.New[String](gdextension.Make[gdextension.String](builtin.creation.String[2], gdextension.SizeNodePath<<4, unsafe.Pointer(&ptr)))
}

func (n NodePath) String() string {
	return StringFromNodePath(n).String()
}

func (n NodePath) Free() {
	ptr, ok := pointers.End(n)
	if !ok {
		return
	}
	gdextension.Free(gdextension.TypeNodePath, &ptr)
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
	return byte(gdextension.Host.Strings.Access(pointers.Get(pointers.Load[String](raw)), n))
}
func (StringProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Readable) {
	s := pointers.Load[String](raw)
	next := s.Substr(0, s.Length())
	return StringType.Rune(gdextension.Host.Strings.Access(pointers.Get(pointers.Load[String](raw)), 0)), 0, StringType.Via(StringProxy{}, pointers.Pack(next))
}
func (StringProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Readable {
	s := pointers.Load[String](raw)
	str := s.Substr(0, s.Length())
	pointers.Set(str, gdextension.Host.Strings.Append.Rune(pointers.Get(str), rune(r)))
	return StringType.Via(StringProxy{}, pointers.Pack(str))
}
func (StringProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Readable {
	s := pointers.Load[String](raw)
	s2 := pointers.Load[String](raw2)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(s2)))
	return StringType.Via(StringProxy{}, pointers.Pack(sub))
}
func (StringProxy) AppendString(raw complex128, str string) StringType.Readable {
	s := pointers.Load[String](raw)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(str))))
	return StringType.Via(StringProxy{}, pointers.Pack(sub))
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
	return byte(gdextension.Host.Strings.Access(pointers.Get(pointers.Load[NodePath](raw).InternalString()), n))
}
func (NodePathProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Readable) {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	next := str.Substr(0, 1).NodePath()
	return StringType.Rune(gdextension.Host.Strings.Access(pointers.Get(str), 0)), 0, StringType.Via(StringProxy{}, pointers.Pack(next))
}
func (NodePathProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Readable {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	pointers.Set(str, gdextension.Host.Strings.Append.Rune(pointers.Get(str), rune(r)))
	return StringType.Via(NodePathProxy{}, pointers.Pack(str.NodePath()))
}
func (NodePathProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Readable {
	s := pointers.Load[NodePath](raw)
	s2 := pointers.Load[NodePath](raw2)
	sub := s.InternalString()
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(s2.InternalString())))
	return StringType.Via(StringNameProxy{}, pointers.Pack(sub.NodePath()))
}
func (NodePathProxy) AppendString(raw complex128, str string) StringType.Readable {
	s := pointers.Load[NodePath](raw)
	sub := s.InternalString()
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(str))))
	return StringType.Via(NodePathProxy{}, pointers.Pack(sub.NodePath()))
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
	return byte(gdextension.Host.Strings.Access(pointers.Get(s), n))
}
func (StringNameProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Readable) {
	s := pointers.Load[StringName](raw)
	next := s.Substr(0, 1).StringName()
	return StringType.Rune(gdextension.Host.Strings.Access(pointers.Get(s.Substr(0, s.Length())), 0)), 0, StringType.Via(StringNameProxy{}, pointers.Pack(next))
}
func (StringNameProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Readable {
	s := pointers.Load[StringName](raw)
	str := s.Substr(0, s.Length())
	pointers.Set(str, gdextension.Host.Strings.Append.Rune(pointers.Get(str), rune(r)))
	return StringType.Via(StringNameProxy{}, pointers.Pack(str.StringName()))
}
func (StringNameProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Readable {
	s := pointers.Load[StringName](raw)
	s2 := pointers.Load[StringName](raw2).String()
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(s2))))
	return StringType.Via(StringNameProxy{}, pointers.Pack(sub.StringName()))
}
func (StringNameProxy) AppendString(raw complex128, str string) StringType.Readable {
	s := pointers.Load[StringName](raw)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(str))))
	return StringType.Via(StringNameProxy{}, pointers.Pack(sub.StringName()))
}
func (StringNameProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	other := pointers.Load[StringName](raw2)
	return int(pointers.Load[StringName](raw).CasecmpTo(other.Substr(0, other.Length())))
}
