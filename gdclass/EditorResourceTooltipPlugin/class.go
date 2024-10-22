package EditorResourceTooltipPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Resource tooltip plugins are used by [FileSystemDock] to generate customized tooltips for specific resources. E.g. tooltip for a [Texture2D] displays a bigger preview and the texture's dimensions.
A plugin must be first registered with [method FileSystemDock.add_resource_tooltip_plugin]. When the user hovers a resource in filesystem dock which is handled by the plugin, [method _make_tooltip_for_path] is called to create the tooltip. It works similarly to [method Control._make_custom_tooltip].
	// EditorResourceTooltipPlugin methods that can be overridden by a [Class] that extends it.
	type EditorResourceTooltipPlugin interface {
		//Return [code]true[/code] if the plugin is going to handle the given [Resource] [param type].
		Handles(atype string) bool
		//Create and return a tooltip that will be displayed when the user hovers a resource under the given [param path] in filesystem dock.
		//The [param metadata] dictionary is provided by preview generator (see [method EditorResourcePreviewGenerator._generate]).
		//[param base] is the base default tooltip, which is a [VBoxContainer] with a file name, type and size labels. If another plugin handled the same file type, [param base] will be output from the previous plugin. For best result, make sure the base tooltip is part of the returned [Control].
		//[b]Note:[/b] It's unadvised to use [method ResourceLoader.load], especially with heavy resources like models or textures, because it will make the editor unresponsive when creating the tooltip. You can use [method request_thumbnail] if you want to display a preview in your tooltip.
		//[b]Note:[/b] If you decide to discard the [param base], make sure to call [method Node.queue_free], because it's not freed automatically.
		//[codeblock]
		//func _make_tooltip_for_path(path, metadata, base):
		//    var t_rect = TextureRect.new()
		//    request_thumbnail(path, t_rect)
		//    base.add_child(t_rect) # The TextureRect will appear at the bottom of the tooltip.
		//    return base
		//[/codeblock]
		MakeTooltipForPath(path string, metadata gd.Dictionary, base gdclass.Control) gdclass.Control
	}

*/
type Go [1]classdb.EditorResourceTooltipPlugin

/*
Return [code]true[/code] if the plugin is going to handle the given [Resource] [param type].
*/
func (Go) _handles(impl func(ptr unsafe.Pointer, atype string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var atype = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Create and return a tooltip that will be displayed when the user hovers a resource under the given [param path] in filesystem dock.
The [param metadata] dictionary is provided by preview generator (see [method EditorResourcePreviewGenerator._generate]).
[param base] is the base default tooltip, which is a [VBoxContainer] with a file name, type and size labels. If another plugin handled the same file type, [param base] will be output from the previous plugin. For best result, make sure the base tooltip is part of the returned [Control].
[b]Note:[/b] It's unadvised to use [method ResourceLoader.load], especially with heavy resources like models or textures, because it will make the editor unresponsive when creating the tooltip. You can use [method request_thumbnail] if you want to display a preview in your tooltip.
[b]Note:[/b] If you decide to discard the [param base], make sure to call [method Node.queue_free], because it's not freed automatically.
[codeblock]
func _make_tooltip_for_path(path, metadata, base):
    var t_rect = TextureRect.new()
    request_thumbnail(path, t_rect)
    base.add_child(t_rect) # The TextureRect will appear at the bottom of the tooltip.
    return base
[/codeblock]
*/
func (Go) _make_tooltip_for_path(impl func(ptr unsafe.Pointer, path string, metadata gd.Dictionary, base gdclass.Control) gdclass.Control, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var metadata = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var base gdclass.Control
		base[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), metadata, base)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}

/*
Requests a thumbnail for the given [TextureRect]. The thumbnail is created asynchronously by [EditorResourcePreview] and automatically set when available.
*/
func (self Go) RequestThumbnail(path string, control gdclass.TextureRect) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RequestThumbnail(gc.String(path), control)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorResourceTooltipPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorResourceTooltipPlugin"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Return [code]true[/code] if the plugin is going to handle the given [Resource] [param type].
*/
func (class) _handles(impl func(ptr unsafe.Pointer, atype gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var atype = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Create and return a tooltip that will be displayed when the user hovers a resource under the given [param path] in filesystem dock.
The [param metadata] dictionary is provided by preview generator (see [method EditorResourcePreviewGenerator._generate]).
[param base] is the base default tooltip, which is a [VBoxContainer] with a file name, type and size labels. If another plugin handled the same file type, [param base] will be output from the previous plugin. For best result, make sure the base tooltip is part of the returned [Control].
[b]Note:[/b] It's unadvised to use [method ResourceLoader.load], especially with heavy resources like models or textures, because it will make the editor unresponsive when creating the tooltip. You can use [method request_thumbnail] if you want to display a preview in your tooltip.
[b]Note:[/b] If you decide to discard the [param base], make sure to call [method Node.queue_free], because it's not freed automatically.
[codeblock]
func _make_tooltip_for_path(path, metadata, base):
    var t_rect = TextureRect.new()
    request_thumbnail(path, t_rect)
    base.add_child(t_rect) # The TextureRect will appear at the bottom of the tooltip.
    return base
[/codeblock]
*/
func (class) _make_tooltip_for_path(impl func(ptr unsafe.Pointer, path gd.String, metadata gd.Dictionary, base gdclass.Control) gdclass.Control, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var metadata = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var base gdclass.Control
		base[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, metadata, base)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Requests a thumbnail for the given [TextureRect]. The thumbnail is created asynchronously by [EditorResourcePreview] and automatically set when available.
*/
//go:nosplit
func (self class) RequestThumbnail(path gd.String, control gdclass.TextureRect)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(control[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourceTooltipPlugin.Bind_request_thumbnail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsEditorResourceTooltipPlugin() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorResourceTooltipPlugin() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_handles": return reflect.ValueOf(self._handles);
	case "_make_tooltip_for_path": return reflect.ValueOf(self._make_tooltip_for_path);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_handles": return reflect.ValueOf(self._handles);
	case "_make_tooltip_for_path": return reflect.ValueOf(self._make_tooltip_for_path);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorResourceTooltipPlugin", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
