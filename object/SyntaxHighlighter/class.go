package SyntaxHighlighter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for syntax highlighters. Provides syntax highlighting data to a [TextEdit]. The associated [TextEdit] will call into the [SyntaxHighlighter] on an as-needed basis.
[b]Note:[/b] A [SyntaxHighlighter] instance should not be used across multiple [TextEdit] nodes.
	// SyntaxHighlighter methods that can be overridden by a [Class] that extends it.
	type SyntaxHighlighter interface {
		//Virtual method which can be overridden to return syntax highlighting data.
		//See [method get_line_syntax_highlighting] for more details.
		GetLineSyntaxHighlighting(line gd.Int) gd.Dictionary
		//Virtual method which can be overridden to clear any local caches.
		ClearHighlightingCache() 
		//Virtual method which can be overridden to update any local caches.
		UpdateCache() 
	}

*/
type Simple [1]classdb.SyntaxHighlighter
func (Simple) _get_line_syntax_highlighting(impl func(ptr unsafe.Pointer, line int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var line = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(line))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _clear_highlighting_cache(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _update_cache(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (self Simple) GetLineSyntaxHighlighting(line int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetLineSyntaxHighlighting(gc, gd.Int(line)))
}
func (self Simple) UpdateCache() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateCache()
}
func (self Simple) ClearHighlightingCache() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearHighlightingCache()
}
func (self Simple) GetTextEdit() [1]classdb.TextEdit {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TextEdit(Expert(self).GetTextEdit(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SyntaxHighlighter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Virtual method which can be overridden to return syntax highlighting data.
See [method get_line_syntax_highlighting] for more details.
*/
func (class) _get_line_syntax_highlighting(impl func(ptr unsafe.Pointer, line gd.Int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var line = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, line)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Virtual method which can be overridden to clear any local caches.
*/
func (class) _clear_highlighting_cache(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Virtual method which can be overridden to update any local caches.
*/
func (class) _update_cache(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
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
func (self class) GetLineSyntaxHighlighting(ctx gd.Lifetime, line gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SyntaxHighlighter.Bind_get_line_syntax_highlighting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears then updates the [SyntaxHighlighter] caches. Override [method _update_cache] for a callback.
[b]Note:[/b] This is called automatically when the associated [TextEdit] node, updates its own cache.
*/
//go:nosplit
func (self class) UpdateCache()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SyntaxHighlighter.Bind_update_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all cached syntax highlighting data.
Then calls overridable method [method _clear_highlighting_cache].
*/
//go:nosplit
func (self class) ClearHighlightingCache()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SyntaxHighlighter.Bind_clear_highlighting_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the associated [TextEdit] node.
*/
//go:nosplit
func (self class) GetTextEdit(ctx gd.Lifetime) object.TextEdit {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SyntaxHighlighter.Bind_get_text_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TextEdit
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSyntaxHighlighter() Expert { return self[0].AsSyntaxHighlighter() }


//go:nosplit
func (self Simple) AsSyntaxHighlighter() Simple { return self[0].AsSyntaxHighlighter() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_line_syntax_highlighting": return reflect.ValueOf(self._get_line_syntax_highlighting);
	case "_clear_highlighting_cache": return reflect.ValueOf(self._clear_highlighting_cache);
	case "_update_cache": return reflect.ValueOf(self._update_cache);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_get_line_syntax_highlighting": return reflect.ValueOf(self._get_line_syntax_highlighting);
	case "_clear_highlighting_cache": return reflect.ValueOf(self._clear_highlighting_cache);
	case "_update_cache": return reflect.ValueOf(self._update_cache);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SyntaxHighlighter", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
