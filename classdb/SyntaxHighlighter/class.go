package SyntaxHighlighter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Base class for syntax highlighters. Provides syntax highlighting data to a [TextEdit]. The associated [TextEdit] will call into the [SyntaxHighlighter] on an as-needed basis.
[b]Note:[/b] A [SyntaxHighlighter] instance should not be used across multiple [TextEdit] nodes.

	// SyntaxHighlighter methods that can be overridden by a [Class] that extends it.
	type SyntaxHighlighter interface {
		//Virtual method which can be overridden to return syntax highlighting data.
		//See [method get_line_syntax_highlighting] for more details.
		GetLineSyntaxHighlighting(line int) Dictionary.Any
		//Virtual method which can be overridden to clear any local caches.
		ClearHighlightingCache()
		//Virtual method which can be overridden to update any local caches.
		UpdateCache()
	}
*/
type Instance [1]gdclass.SyntaxHighlighter
type Any interface {
	gd.IsClass
	AsSyntaxHighlighter() Instance
}

/*
Virtual method which can be overridden to return syntax highlighting data.
See [method get_line_syntax_highlighting] for more details.
*/
func (Instance) _get_line_syntax_highlighting(impl func(ptr unsafe.Pointer, line int) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var line = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(line))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method which can be overridden to clear any local caches.
*/
func (Instance) _clear_highlighting_cache(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Virtual method which can be overridden to update any local caches.
*/
func (Instance) _update_cache(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns syntax highlighting data for a single line. If the line is not cached, calls [method _get_line_syntax_highlighting] to calculate the data.
The return [Dictionary] is column number to [Dictionary]. The column number notes the start of a region, the region will end if another region is found, or at the end of the line. The nested [Dictionary] contains the data for that region, currently only the key "color" is supported.
[b]Example return:[/b]
[codeblock]

	var color_map = {
	    0: {
	        "color": Color(1, 0, 0)
	    },
	    5: {
	        "color": Color(0, 1, 0)
	    }
	}

[/codeblock]
This will color columns 0-4 red, and columns 5-eol in green.
*/
func (self Instance) GetLineSyntaxHighlighting(line int) Dictionary.Any {
	return Dictionary.Any(class(self).GetLineSyntaxHighlighting(gd.Int(line)))
}

/*
Clears then updates the [SyntaxHighlighter] caches. Override [method _update_cache] for a callback.
[b]Note:[/b] This is called automatically when the associated [TextEdit] node, updates its own cache.
*/
func (self Instance) UpdateCache() {
	class(self).UpdateCache()
}

/*
Clears all cached syntax highlighting data.
Then calls overridable method [method _clear_highlighting_cache].
*/
func (self Instance) ClearHighlightingCache() {
	class(self).ClearHighlightingCache()
}

/*
Returns the associated [TextEdit] node.
*/
func (self Instance) GetTextEdit() [1]gdclass.TextEdit {
	return [1]gdclass.TextEdit(class(self).GetTextEdit())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SyntaxHighlighter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SyntaxHighlighter"))
	return Instance{*(*gdclass.SyntaxHighlighter)(unsafe.Pointer(&object))}
}

/*
Virtual method which can be overridden to return syntax highlighting data.
See [method get_line_syntax_highlighting] for more details.
*/
func (class) _get_line_syntax_highlighting(impl func(ptr unsafe.Pointer, line gd.Int) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var line = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, line)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method which can be overridden to clear any local caches.
*/
func (class) _clear_highlighting_cache(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Virtual method which can be overridden to update any local caches.
*/
func (class) _update_cache(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns syntax highlighting data for a single line. If the line is not cached, calls [method _get_line_syntax_highlighting] to calculate the data.
The return [Dictionary] is column number to [Dictionary]. The column number notes the start of a region, the region will end if another region is found, or at the end of the line. The nested [Dictionary] contains the data for that region, currently only the key "color" is supported.
[b]Example return:[/b]
[codeblock]
var color_map = {
    0: {
        "color": Color(1, 0, 0)
    },
    5: {
        "color": Color(0, 1, 0)
    }
}
[/codeblock]
This will color columns 0-4 red, and columns 5-eol in green.
*/
//go:nosplit
func (self class) GetLineSyntaxHighlighting(line gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SyntaxHighlighter.Bind_get_line_syntax_highlighting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears then updates the [SyntaxHighlighter] caches. Override [method _update_cache] for a callback.
[b]Note:[/b] This is called automatically when the associated [TextEdit] node, updates its own cache.
*/
//go:nosplit
func (self class) UpdateCache() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SyntaxHighlighter.Bind_update_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears all cached syntax highlighting data.
Then calls overridable method [method _clear_highlighting_cache].
*/
//go:nosplit
func (self class) ClearHighlightingCache() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SyntaxHighlighter.Bind_clear_highlighting_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the associated [TextEdit] node.
*/
//go:nosplit
func (self class) GetTextEdit() [1]gdclass.TextEdit {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SyntaxHighlighter.Bind_get_text_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.TextEdit{gd.PointerMustAssertInstanceID[gdclass.TextEdit](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsSyntaxHighlighter() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSyntaxHighlighter() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_get_line_syntax_highlighting":
		return reflect.ValueOf(self._get_line_syntax_highlighting)
	case "_clear_highlighting_cache":
		return reflect.ValueOf(self._clear_highlighting_cache)
	case "_update_cache":
		return reflect.ValueOf(self._update_cache)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_line_syntax_highlighting":
		return reflect.ValueOf(self._get_line_syntax_highlighting)
	case "_clear_highlighting_cache":
		return reflect.ValueOf(self._clear_highlighting_cache)
	case "_update_cache":
		return reflect.ValueOf(self._update_cache)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	gdclass.Register("SyntaxHighlighter", func(ptr gd.Object) any {
		return [1]gdclass.SyntaxHighlighter{*(*gdclass.SyntaxHighlighter)(unsafe.Pointer(&ptr))}
	})
}
