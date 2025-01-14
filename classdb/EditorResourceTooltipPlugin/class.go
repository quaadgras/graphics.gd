// Package EditorResourceTooltipPlugin provides methods for working with EditorResourceTooltipPlugin object instances.
package EditorResourceTooltipPlugin

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Resource tooltip plugins are used by [FileSystemDock] to generate customized tooltips for specific resources. E.g. tooltip for a [Texture2D] displays a bigger preview and the texture's dimensions.
A plugin must be first registered with [method FileSystemDock.add_resource_tooltip_plugin]. When the user hovers a resource in filesystem dock which is handled by the plugin, [method _make_tooltip_for_path] is called to create the tooltip. It works similarly to [method Control._make_custom_tooltip].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorResourceTooltipPlugin)
*/
type Instance [1]gdclass.EditorResourceTooltipPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorResourceTooltipPlugin() Instance
}
type Interface interface {
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
	MakeTooltipForPath(path string, metadata Dictionary.Any, base [1]gdclass.Control) [1]gdclass.Control
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Handles(atype string) (_ bool) { return }
func (self implementation) MakeTooltipForPath(path string, metadata Dictionary.Any, base [1]gdclass.Control) (_ [1]gdclass.Control) {
	return
}

/*
Return [code]true[/code] if the plugin is going to handle the given [Resource] [param type].
*/
func (Instance) _handles(impl func(ptr unsafe.Pointer, atype string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(atype)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype.String())
		gd.UnsafeSet(p_back, ret)
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
func (Instance) _make_tooltip_for_path(impl func(ptr unsafe.Pointer, path string, metadata Dictionary.Any, base [1]gdclass.Control) [1]gdclass.Control) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var metadata = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(metadata)
		var base = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 2))})}
		defer pointers.End(base[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), metadata, base)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Requests a thumbnail for the given [TextureRect]. The thumbnail is created asynchronously by [EditorResourcePreview] and automatically set when available.
*/
func (self Instance) RequestThumbnail(path string, control [1]gdclass.TextureRect) {
	class(self).RequestThumbnail(gd.NewString(path), control)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorResourceTooltipPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourceTooltipPlugin"))
	casted := Instance{*(*gdclass.EditorResourceTooltipPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Return [code]true[/code] if the plugin is going to handle the given [Resource] [param type].
*/
func (class) _handles(impl func(ptr unsafe.Pointer, atype gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
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
func (class) _make_tooltip_for_path(impl func(ptr unsafe.Pointer, path gd.String, metadata gd.Dictionary, base [1]gdclass.Control) [1]gdclass.Control) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var metadata = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		var base = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}
		defer pointers.End(base[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, metadata, base)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Requests a thumbnail for the given [TextureRect]. The thumbnail is created asynchronously by [EditorResourcePreview] and automatically set when available.
*/
//go:nosplit
func (self class) RequestThumbnail(path gd.String, control [1]gdclass.TextureRect) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourceTooltipPlugin.Bind_request_thumbnail, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsEditorResourceTooltipPlugin() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorResourceTooltipPlugin() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_make_tooltip_for_path":
		return reflect.ValueOf(self._make_tooltip_for_path)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_make_tooltip_for_path":
		return reflect.ValueOf(self._make_tooltip_for_path)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorResourceTooltipPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorResourceTooltipPlugin{*(*gdclass.EditorResourceTooltipPlugin)(unsafe.Pointer(&ptr))}
	})
}
