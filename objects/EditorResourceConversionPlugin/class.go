package EditorResourceConversionPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
		Handles(resource objects.Resource) bool
		//Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
		Convert(resource objects.Resource) objects.Resource
	}
*/
type Instance [1]classdb.EditorResourceConversionPlugin

/*
Returns the class name of the target type of [Resource] that this plugin converts source resources to.
*/
func (Instance) _converts_to(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
Called to determine whether a particular [Resource] can be converted to the target resource type by this plugin.
*/
func (Instance) _handles(impl func(ptr unsafe.Pointer, resource objects.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
*/
func (Instance) _convert(impl func(ptr unsafe.Pointer, resource objects.Resource) objects.Resource) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorResourceConversionPlugin

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourceConversionPlugin"))
	return Instance{classdb.EditorResourceConversionPlugin(object)}
}

/*
Returns the class name of the target type of [Resource] that this plugin converts source resources to.
*/
func (class) _converts_to(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
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
Called to determine whether a particular [Resource] can be converted to the target resource type by this plugin.
*/
func (class) _handles(impl func(ptr unsafe.Pointer, resource objects.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
*/
func (class) _convert(impl func(ptr unsafe.Pointer, resource objects.Resource) objects.Resource) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorResourceConversionPlugin() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorResourceConversionPlugin() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_converts_to":
		return reflect.ValueOf(self._converts_to)
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_convert":
		return reflect.ValueOf(self._convert)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_converts_to":
		return reflect.ValueOf(self._converts_to)
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_convert":
		return reflect.ValueOf(self._convert)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorResourceConversionPlugin", func(ptr gd.Object) any { return classdb.EditorResourceConversionPlugin(ptr) })
}
