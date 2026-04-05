//go:build !generate

package gd

import (
	"runtime"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/noescape"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/rodatacheck"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/variant/Path"
	StringType "graphics.gd/variant/String"
)

var (
	static_string_names = make(map[string]gdextension.StringName)
	static_strings      = make(map[string]gdextension.String)
	static_nodepaths    = make(map[string]gdextension.NodePath)
)

func init() {
	RegisterCleanup(func() {
		for _, name := range static_string_names {
			noescape.Free(gdextension.TypeStringName, &name)
		}
		for _, name := range static_strings {
			noescape.Free(gdextension.TypeString, &name)
		}
		for _, name := range static_nodepaths {
			noescape.Free(gdextension.TypeNodePath, &name)
		}
	})
}

func (s String) StringName() StringName {
	var arg = pointers.Get(s)
	return pointers.New[StringName](noescape.Make[gdextension.StringName](builtin.creation.StringName[2], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

// Copy returns a copy of the string that is owned by the provided context.
func (s String) Copy() String {
	var arg = pointers.Get(s)
	return pointers.New[String](noescape.Make[gdextension.String](builtin.creation.String[1], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

func (s String) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	noescape.Free(gdextension.TypeString, &ptr)
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
	gdunsafe.String(pointers.Get(s)[0]).Encode(gdunsafe.UTF8, buf)
	return unsafe.String(&buf[0], len(buf))
}

func StringFromStringName(s StringName) String {
	var arg = pointers.Get(s)
	return pointers.New[String](noescape.Make[gdextension.String](builtin.creation.String[2], gdextension.SizeStringName<<4, unsafe.Pointer(&arg)))
}

func StringFromNodePath(s NodePath) String {
	var arg = pointers.Get(s)
	return pointers.New[String](noescape.Make[gdextension.String](builtin.creation.String[3], gdextension.SizeNodePath<<4, unsafe.Pointer(&arg)))
}

func NewStringNameFromString(s String) StringName {
	var arg = pointers.Get(s)
	return pointers.New[StringName](noescape.Make[gdextension.StringName](builtin.creation.StringName[2], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

func (s StringName) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	if ptr == (gdextension.StringName{}) {
		return
	}
	noescape.Free(gdextension.TypeStringName, &ptr)
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
	return pointers.New[NodePath](noescape.Make[gdextension.NodePath](builtin.creation.NodePath[2], gdextension.SizeString<<4, unsafe.Pointer(&arg)))
}

func (n NodePath) InternalString() String {
	var ptr = pointers.Get(n)
	return pointers.New[String](noescape.Make[gdextension.String](builtin.creation.String[2], gdextension.SizeNodePath<<4, unsafe.Pointer(&ptr)))
}

func (n NodePath) String() string {
	return StringFromNodePath(n).String()
}

func (n NodePath) Free() {
	ptr, ok := pointers.End(n)
	if !ok {
		return
	}
	noescape.Free(gdextension.TypeNodePath, &ptr)
}

func InternalString(s StringType.Unicode) String {
	if str := s.String(); rodatacheck.String(str) && threadcheck.Main() {
		if name, ok := static_strings[str]; ok {
			return pointers.Raw[String](name)
		}
		name := gdextension.String{gdextension.Pointer(gdunsafe.UTF8.String(str))}
		static_strings[str] = name
		return pointers.Raw[String](name)
	}
	_, ptr := StringType.Proxy(s, StringCacheCheck, NewStringProxy)
	return pointers.Load[String](ptr)
}

func StringCacheCheck(_ StringProxy, raw complex128) bool { return true }

func NewStringProxy() (StringProxy, complex128) {
	if !threadcheck.Main() {
		ptr := pointers.Pin(NewString(""))
		runtime.AddCleanup(&ptr, func(raw String) {
			raw.Free()
		}, ptr)
		return StringProxy{indirect: &ptr}, 0
	}
	return StringProxy{}, pointers.Pack(NewString(""))
}

type StringProxy struct {
	indirect *String
}

func (proxy StringProxy) Len(raw complex128) int {
	if proxy.indirect != nil {
		return StringProxy{}.Len(pointers.Pack(*proxy.indirect))
	}
	return pointers.Load[String](raw).Len()
}
func (proxy StringProxy) Slice(raw complex128, index int, close int) StringType.Unicode {
	if proxy.indirect != nil {
		return StringProxy{}.Slice(pointers.Pack(*proxy.indirect), index, close)
	}
	s := pointers.Load[String](raw)
	s = s.Substr(Int(index), Int(close))
	return StringType.Via(StringProxy{}, pointers.Pack(s))
}
func (proxy StringProxy) String(raw complex128) string {
	if proxy.indirect != nil {
		return StringProxy{}.String(pointers.Pack(*proxy.indirect))
	}
	return pointers.Load[String](raw).String()
}
func (proxy StringProxy) Index(raw complex128, n int) byte {
	if proxy.indirect != nil {
		return StringProxy{}.Index(pointers.Pack(*proxy.indirect), n)
	}
	return byte(gdunsafe.String(pointers.Get(pointers.Load[String](raw))[0]).Access(gdunsafe.Int(n)))
}
func (proxy StringProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Unicode) {
	if proxy.indirect != nil {
		return StringProxy{}.DecodeRune(pointers.Pack(*proxy.indirect))
	}
	s := pointers.Load[String](raw)
	next := s.Substr(0, s.Length())
	return StringType.Rune(gdunsafe.String(pointers.Get(pointers.Load[String](raw))[0]).Access(0)), 0, StringType.Via(StringProxy{}, pointers.Pack(next))
}
func (proxy StringProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Unicode {
	if proxy.indirect != nil {
		return StringProxy{}.AppendRune(pointers.Pack(*proxy.indirect), r)
	}
	s := pointers.Load[String](raw)
	str := s.Substr(0, s.Length())
	pointers.Set(str, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(str)[0]).AppendRune(int32(r)))})
	return StringType.Via(StringProxy{}, pointers.Pack(str))
}
func (proxy StringProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Unicode {
	if proxy.indirect != nil {
		return StringProxy{}.AppendOther(pointers.Pack(*proxy.indirect), api, raw2)
	}
	s := pointers.Load[String](raw)
	s2 := pointers.Load[String](raw2)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(sub)[0]).Append(gdunsafe.String(pointers.Get(s2)[0])))})
	return StringType.Via(StringProxy{}, pointers.Pack(sub))
}
func (proxy StringProxy) AppendString(raw complex128, str string) StringType.Unicode {
	if proxy.indirect != nil {
		return StringProxy{}.AppendString(pointers.Pack(*proxy.indirect), str)
	}
	s := pointers.Load[String](raw)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(sub)[0]).Append(gdunsafe.String(pointers.Get(NewString(str))[0])))})
	return StringType.Via(StringProxy{}, pointers.Pack(sub))
}
func (proxy StringProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	if proxy.indirect != nil {
		return StringProxy{}.CompareOther(pointers.Pack(*proxy.indirect), other_api, raw2)
	}
	return int(pointers.Load[String](raw).CasecmpTo(pointers.Load[String](raw2)))
}

func InternalNodePath(s Path.ToNode) NodePath {
	if str := s.String(); rodatacheck.String(str) && threadcheck.Main() {
		if name, ok := static_nodepaths[str]; ok {
			return pointers.Raw[NodePath](name)
		}
		name := gdextension.String{gdextension.Pointer(gdunsafe.UTF8.String(str))}
		path := noescape.Make[gdextension.NodePath](builtin.creation.NodePath[2], gdextension.SizeString<<4, unsafe.Pointer(&name))
		static_nodepaths[str] = path
		noescape.Free(gdextension.TypeString, &name)
		return pointers.Raw[NodePath](path)
	}
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
func (NodePathProxy) Slice(raw complex128, index int, close int) StringType.Unicode {
	s := pointers.Load[NodePath](raw)
	s = s.InternalString().Substr(Int(index), Int(close)).NodePath()
	return StringType.Via(StringProxy{}, pointers.Pack(s))
}
func (NodePathProxy) String(raw complex128) string {
	return pointers.Load[NodePath](raw).String()
}
func (NodePathProxy) Index(raw complex128, n int) byte {
	return byte(gdunsafe.String(pointers.Get(pointers.Load[NodePath](raw).InternalString())[0]).Access(gdunsafe.Int(n)))
}
func (NodePathProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Unicode) {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	next := str.Substr(0, 1).NodePath()
	return StringType.Rune(gdunsafe.String(pointers.Get(str)[0]).Access(0)), 0, StringType.Via(StringProxy{}, pointers.Pack(next))
}
func (NodePathProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Unicode {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	pointers.Set(str, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(str)[0]).AppendRune(int32(r)))})
	return StringType.Via(NodePathProxy{}, pointers.Pack(str.NodePath()))
}
func (NodePathProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Unicode {
	s := pointers.Load[NodePath](raw)
	s2 := pointers.Load[NodePath](raw2)
	sub := s.InternalString()
	pointers.Set(sub, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(sub)[0]).Append(gdunsafe.String(pointers.Get(s2.InternalString())[0])))})
	return StringType.Via(StringNameProxy{}, pointers.Pack(sub.NodePath()))
}
func (NodePathProxy) AppendString(raw complex128, str string) StringType.Unicode {
	s := pointers.Load[NodePath](raw)
	sub := s.InternalString()
	pointers.Set(sub, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(sub)[0]).Append(gdunsafe.String(pointers.Get(NewString(str))[0])))})
	return StringType.Via(NodePathProxy{}, pointers.Pack(sub.NodePath()))
}
func (NodePathProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	return int(pointers.Load[NodePath](raw).InternalString().CasecmpTo(pointers.Load[NodePath](raw2).InternalString()))
}

func InternalStringName(s StringType.Name) StringName {
	if str := s.String(); rodatacheck.String(str) && threadcheck.Main() {
		if name, ok := static_string_names[str]; ok {
			return pointers.Raw[StringName](name)
		}
		name := gdextension.StringName{gdextension.Pointer(gdunsafe.UTF8.Intern(str))}
		static_string_names[str] = name
		return pointers.Raw[StringName](name)
	}
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
func (StringNameProxy) Slice(raw complex128, index int, close int) StringType.Unicode {
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
	return byte(gdunsafe.String(pointers.Get(s)[0]).Access(gdunsafe.Int(n)))
}
func (StringNameProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Unicode) {
	s := pointers.Load[StringName](raw)
	next := s.Substr(0, 1).StringName()
	return StringType.Rune(gdunsafe.String(pointers.Get(s.Substr(0, s.Length()))[0]).Access(0)), 0, StringType.Via(StringNameProxy{}, pointers.Pack(next))
}
func (StringNameProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Unicode {
	s := pointers.Load[StringName](raw)
	str := s.Substr(0, s.Length())
	pointers.Set(str, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(str)[0]).AppendRune(int32(r)))})
	return StringType.Via(StringNameProxy{}, pointers.Pack(str.StringName()))
}
func (StringNameProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Unicode {
	s := pointers.Load[StringName](raw)
	s2 := pointers.Load[StringName](raw2).String()
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(sub)[0]).Append(gdunsafe.String(pointers.Get(NewString(s2))[0])))})
	return StringType.Via(StringNameProxy{}, pointers.Pack(sub.StringName()))
}
func (StringNameProxy) AppendString(raw complex128, str string) StringType.Unicode {
	s := pointers.Load[StringName](raw)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.String{gdextension.Pointer(gdunsafe.String(pointers.Get(sub)[0]).Append(gdunsafe.String(pointers.Get(NewString(str))[0])))})
	return StringType.Via(StringNameProxy{}, pointers.Pack(sub.StringName()))
}
func (StringNameProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	other := pointers.Load[StringName](raw2)
	return int(pointers.Load[StringName](raw).CasecmpTo(other.Substr(0, other.Length())))
}
