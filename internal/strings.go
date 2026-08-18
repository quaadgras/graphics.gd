//go:build !generate

package gd

import (
	"unsafe"

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
	// Godot's String.Length() counts characters, but the UTF-8 encoding of a
	// non-ASCII string can need up to 4 bytes per character, and
	// string_to_utf8_chars writes at most len(buf) bytes. A character-counted
	// buffer therefore truncates multibyte sequences mid-character and yields
	// invalid UTF-8 (Godot then reports "Unicode parsing error" when the
	// string is handed back). Size the buffer for the worst-case byte length
	// and slice by the number of bytes actually written (which includes the
	// NUL terminator when the string fits, so drop a trailing NUL).
	var buf = make([]byte, s.Length()*4+1)
	n := gdextension.Host.Strings.Encode.UTF8(pointers.Get(s), buf)
	if n > 0 && buf[n-1] == 0 {
		n--
	}
	// string(buf[:n]) copies; unsafe.String would alias buf without keeping it
	// alive for the GC, so the returned value could dangle after collection.
	return string(buf[:n])
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
	// Build the temporary conversion String as an UNTRACKED (Raw) handle
	// and free it explicitly. A tracked pointers.New here would be
	// reclaimable by the main thread's per-frame Cycle, which — when this
	// runs off the loader thread — races our read and crashes (this was
	// the original SIGSEGV). A Raw handle is invisible to Cycle, so it
	// can't be freed out from under us; we own its lifetime and free it
	// once we've copied the bytes out.
	arg := pointers.Get(s)
	tmp := pointers.Raw[String](noescape.Make[gdextension.String](builtin.creation.String[2], gdextension.SizeStringName<<4, unsafe.Pointer(&arg)))
	out := tmp.String()
	raw := pointers.Get(tmp)
	noescape.Free(gdextension.TypeString, &raw)
	return out
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
	// Untracked temporary conversion String, freed explicitly — same
	// rationale as StringName.String: a tracked temp here races the
	// main-thread Cycle when called off-thread.
	arg := pointers.Get(n)
	tmp := pointers.Raw[String](noescape.Make[gdextension.String](builtin.creation.String[3], gdextension.SizeNodePath<<4, unsafe.Pointer(&arg)))
	out := tmp.String()
	raw := pointers.Get(tmp)
	noescape.Free(gdextension.TypeString, &raw)
	return out
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
		name := gdextension.Host.Strings.Decode.UTF8(str)
		static_strings[str] = name
		return pointers.Raw[String](name)
	}
	_, ptr := StringType.Proxy(s, StringCacheCheck, NewStringProxy)
	return pointers.Load[String](ptr)
}

func StringCacheCheck(_ StringProxy, raw complex128) bool { return true }

func NewStringProxy() (StringProxy, complex128) {
	return WrapString(NewString(""))
}

// StringProxy is engine-backed proxy state for a unicode string. The anchor
// is nil for main-thread (frame-temporary) strings and carries the GC
// lifetime of goroutine-created strings — see anchors.go. Methods always
// operate on the packed state; the anchor only needs to travel with the
// proxy value.
type StringProxy struct {
	anchor *String
}

func (proxy StringProxy) Len(raw complex128) int {
	return pointers.Load[String](raw).Len()
}
func (proxy StringProxy) Slice(raw complex128, index int, close int) StringType.Unicode {
	s := pointers.Load[String](raw)
	s = s.Substr(Int(index), Int(close))
	return StringType.Via(WrapString(s))
}
func (proxy StringProxy) String(raw complex128) string {
	return pointers.Load[String](raw).String()
}
func (proxy StringProxy) Index(raw complex128, n int) byte {
	return byte(gdextension.Host.Strings.Access(pointers.Get(pointers.Load[String](raw)), n))
}
func (proxy StringProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Unicode) {
	s := pointers.Load[String](raw)
	next := s.Substr(0, s.Length())
	return StringType.Rune(gdextension.Host.Strings.Access(pointers.Get(pointers.Load[String](raw)), 0)), 0, StringType.Via(WrapString(next))
}
func (proxy StringProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Unicode {
	s := pointers.Load[String](raw)
	str := s.Substr(0, s.Length())
	pointers.Set(str, gdextension.Host.Strings.Append.Rune(pointers.Get(str), rune(r)))
	return StringType.Via(WrapString(str))
}
func (proxy StringProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Unicode {
	s := pointers.Load[String](raw)
	s2 := pointers.Load[String](raw2)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(s2)))
	return StringType.Via(WrapString(sub))
}
func (proxy StringProxy) AppendString(raw complex128, str string) StringType.Unicode {
	s := pointers.Load[String](raw)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(str))))
	return StringType.Via(WrapString(sub))
}
func (proxy StringProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	return int(pointers.Load[String](raw).CasecmpTo(pointers.Load[String](raw2)))
}

func InternalNodePath(s Path.ToNode) NodePath {
	if str := s.String(); rodatacheck.String(str) && threadcheck.Main() {
		if name, ok := static_nodepaths[str]; ok {
			return pointers.Raw[NodePath](name)
		}
		name := gdextension.Host.Strings.Decode.UTF8(str)
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
	return WrapNodePath(NewString("").NodePath())
}

// NodePathProxy is engine-backed proxy state for a node path, the anchor
// carries the GC lifetime of goroutine-created paths — see anchors.go.
type NodePathProxy struct {
	anchor *NodePath
}

func (NodePathProxy) Len(raw complex128) int {
	return pointers.Load[NodePath](raw).InternalString().Len()
}
func (NodePathProxy) Slice(raw complex128, index int, close int) StringType.Unicode {
	// The slice of a path is returned as a plain string wrapper (previously
	// it round-tripped through NodePath but was packed as a StringProxy
	// anyway, relying on the two engine types sharing a pointer shape —
	// the typed wrap is required now that the anchor's cleanup runs the
	// wrapper's own destructor).
	s := pointers.Load[NodePath](raw)
	return StringType.Via(WrapString(s.InternalString().Substr(Int(index), Int(close))))
}
func (NodePathProxy) String(raw complex128) string {
	return pointers.Load[NodePath](raw).String()
}
func (NodePathProxy) Index(raw complex128, n int) byte {
	return byte(gdextension.Host.Strings.Access(pointers.Get(pointers.Load[NodePath](raw).InternalString()), n))
}
func (NodePathProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Unicode) {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	next := str.Substr(0, 1)
	return StringType.Rune(gdextension.Host.Strings.Access(pointers.Get(str), 0)), 0, StringType.Via(WrapString(next))
}
func (NodePathProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Unicode {
	s := pointers.Load[NodePath](raw)
	str := s.InternalString()
	pointers.Set(str, gdextension.Host.Strings.Append.Rune(pointers.Get(str), rune(r)))
	return StringType.Via(WrapNodePath(str.NodePath()))
}
func (NodePathProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Unicode {
	s := pointers.Load[NodePath](raw)
	s2 := pointers.Load[NodePath](raw2)
	sub := s.InternalString()
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(s2.InternalString())))
	return StringType.Via(WrapNodePath(sub.NodePath()))
}
func (NodePathProxy) AppendString(raw complex128, str string) StringType.Unicode {
	s := pointers.Load[NodePath](raw)
	sub := s.InternalString()
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(str))))
	return StringType.Via(WrapNodePath(sub.NodePath()))
}
func (NodePathProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	return int(pointers.Load[NodePath](raw).InternalString().CasecmpTo(pointers.Load[NodePath](raw2).InternalString()))
}

func InternalStringName(s StringType.Name) StringName {
	if str := s.String(); rodatacheck.String(str) && threadcheck.Main() {
		if name, ok := static_string_names[str]; ok {
			return pointers.Raw[StringName](name)
		}
		name := gdextension.Host.Strings.Intern.UTF8(str)
		static_string_names[str] = name
		return pointers.Raw[StringName](name)
	}
	_, ptr := StringType.Proxy(s, StringNameCheck, NewStringNameProxy)
	return pointers.Load[StringName](ptr)
}

func StringNameCheck(_ StringNameProxy, raw complex128) bool { return true }

func NewStringNameProxy() (StringNameProxy, complex128) {
	return WrapStringName(NewStringName(""))
}

// StringNameProxy is engine-backed proxy state for a string name, the anchor
// carries the GC lifetime of goroutine-created names — see anchors.go.
type StringNameProxy struct {
	anchor *StringName
}

func (StringNameProxy) Len(raw complex128) int {
	return int(pointers.Load[StringName](raw).Length())
}
func (StringNameProxy) Slice(raw complex128, index int, close int) StringType.Unicode {
	s := pointers.Load[StringName](raw)
	return StringType.Via(WrapStringName(s.Substr(Int(index), Int(close)).StringName()))
}
func (StringNameProxy) String(raw complex128) string {
	return pointers.Load[StringName](raw).String()
}
func (StringNameProxy) Index(raw complex128, n int) byte {
	name := pointers.Load[StringName](raw)
	s := name.Substr(0, name.Length())
	return byte(gdextension.Host.Strings.Access(pointers.Get(s), n))
}
func (StringNameProxy) DecodeRune(raw complex128) (StringType.Rune, int, StringType.Unicode) {
	s := pointers.Load[StringName](raw)
	next := s.Substr(0, 1).StringName()
	return StringType.Rune(gdextension.Host.Strings.Access(pointers.Get(s.Substr(0, s.Length())), 0)), 0, StringType.Via(WrapStringName(next))
}
func (StringNameProxy) AppendRune(raw complex128, r StringType.Rune) StringType.Unicode {
	s := pointers.Load[StringName](raw)
	str := s.Substr(0, s.Length())
	pointers.Set(str, gdextension.Host.Strings.Append.Rune(pointers.Get(str), rune(r)))
	return StringType.Via(WrapStringName(str.StringName()))
}
func (StringNameProxy) AppendOther(raw complex128, api StringType.API, raw2 complex128) StringType.Unicode {
	s := pointers.Load[StringName](raw)
	s2 := pointers.Load[StringName](raw2).String()
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(s2))))
	return StringType.Via(WrapStringName(sub.StringName()))
}
func (StringNameProxy) AppendString(raw complex128, str string) StringType.Unicode {
	s := pointers.Load[StringName](raw)
	sub := s.Substr(0, s.Length())
	pointers.Set(sub, gdextension.Host.Strings.Append.String(pointers.Get(sub), pointers.Get(NewString(str))))
	return StringType.Via(WrapStringName(sub.StringName()))
}
func (StringNameProxy) CompareOther(raw complex128, other_api StringType.API, raw2 complex128) int {
	other := pointers.Load[StringName](raw2)
	return int(pointers.Load[StringName](raw).CasecmpTo(other.Substr(0, other.Length())))
}
