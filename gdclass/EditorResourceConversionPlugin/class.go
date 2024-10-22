package EditorResourceConversionPlugin

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
[EditorResourceConversionPlugin] is invoked when the context menu is brought up for a resource in the editor inspector. Relevant conversion plugins will appear as menu options to convert the given resource to a target type.
Below shows an example of a basic plugin that will convert an [ImageTexture] to a [PortableCompressedTexture2D].
[codeblocks]
[gdscript]
extends EditorResourceConversionPlugin

func _handles(resource: Resource):
    return resource is ImageTexture

func _converts_to():
    return "PortableCompressedTexture2D"

func _convert(itex: Resource):
    var ptex = PortableCompressedTexture2D.new()
    ptex.create_from_image(itex.get_image(), PortableCompressedTexture2D.COMPRESSION_MODE_LOSSLESS)
    return ptex
[/gdscript]
[/codeblocks]
To use an [EditorResourceConversionPlugin], register it using the [method EditorPlugin.add_resource_conversion_plugin] method first.
	// EditorResourceConversionPlugin methods that can be overridden by a [Class] that extends it.
	type EditorResourceConversionPlugin interface {
		//Returns the class name of the target type of [Resource] that this plugin converts source resources to.
		ConvertsTo() string
		//Called to determine whether a particular [Resource] can be converted to the target resource type by this plugin.
		Handles(resource gdclass.Resource) bool
		//Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
		Convert(resource gdclass.Resource) gdclass.Resource
	}

*/
type Go [1]classdb.EditorResourceConversionPlugin

/*
Returns the class name of the target type of [Resource] that this plugin converts source resources to.
*/
func (Go) _converts_to(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Called to determine whether a particular [Resource] can be converted to the target resource type by this plugin.
*/
func (Go) _handles(impl func(ptr unsafe.Pointer, resource gdclass.Resource) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var resource gdclass.Resource
		resource[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
*/
func (Go) _convert(impl func(ptr unsafe.Pointer, resource gdclass.Resource) gdclass.Resource, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var resource gdclass.Resource
		resource[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorResourceConversionPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorResourceConversionPlugin"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the class name of the target type of [Resource] that this plugin converts source resources to.
*/
func (class) _converts_to(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Called to determine whether a particular [Resource] can be converted to the target resource type by this plugin.
*/
func (class) _handles(impl func(ptr unsafe.Pointer, resource gdclass.Resource) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var resource gdclass.Resource
		resource[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
*/
func (class) _convert(impl func(ptr unsafe.Pointer, resource gdclass.Resource) gdclass.Resource, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var resource gdclass.Resource
		resource[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

func (self class) AsEditorResourceConversionPlugin() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorResourceConversionPlugin() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_converts_to": return reflect.ValueOf(self._converts_to);
	case "_handles": return reflect.ValueOf(self._handles);
	case "_convert": return reflect.ValueOf(self._convert);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_converts_to": return reflect.ValueOf(self._converts_to);
	case "_handles": return reflect.ValueOf(self._handles);
	case "_convert": return reflect.ValueOf(self._convert);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorResourceConversionPlugin", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
