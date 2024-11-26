package EditorSyntaxHighlighter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/SyntaxHighlighter"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Base class that all [SyntaxHighlighter]s used by the [ScriptEditor] extend from.
Add a syntax highlighter to an individual script by calling [method ScriptEditorBase.add_syntax_highlighter]. To apply to all scripts on open, call [method ScriptEditor.register_syntax_highlighter].

	// EditorSyntaxHighlighter methods that can be overridden by a [Class] that extends it.
	type EditorSyntaxHighlighter interface {
		//Virtual method which can be overridden to return the syntax highlighter name.
		GetName() string
		//Virtual method which can be overridden to return the supported language names.
		GetSupportedLanguages() []string
	}
*/
type Instance [1]classdb.EditorSyntaxHighlighter

/*
Virtual method which can be overridden to return the syntax highlighter name.
*/
func (Instance) _get_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method which can be overridden to return the supported language names.
*/
func (Instance) _get_supported_languages(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorSyntaxHighlighter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSyntaxHighlighter"))
	return Instance{classdb.EditorSyntaxHighlighter(object)}
}

/*
Virtual method which can be overridden to return the syntax highlighter name.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method which can be overridden to return the supported language names.
*/
func (class) _get_supported_languages(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorSyntaxHighlighter() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorSyntaxHighlighter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsSyntaxHighlighter() SyntaxHighlighter.Advanced {
	return *((*SyntaxHighlighter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSyntaxHighlighter() SyntaxHighlighter.Instance {
	return *((*SyntaxHighlighter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_supported_languages":
		return reflect.ValueOf(self._get_supported_languages)
	default:
		return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_supported_languages":
		return reflect.ValueOf(self._get_supported_languages)
	default:
		return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}
func init() {
	classdb.Register("EditorSyntaxHighlighter", func(ptr gd.Object) any { return classdb.EditorSyntaxHighlighter(ptr) })
}
