// Package EditorResourceConversionPlugin provides methods for working with EditorResourceConversionPlugin object instances.
package EditorResourceConversionPlugin

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorResourceConversionPlugin)
*/
type Instance [1]gdclass.EditorResourceConversionPlugin
type Any interface {
	gd.IsClass
	AsEditorResourceConversionPlugin() Instance
}
type Interface interface {
	//Returns the class name of the target type of [Resource] that this plugin converts source resources to.
	ConvertsTo() string
	//Called to determine whether a particular [Resource] can be converted to the target resource type by this plugin.
	Handles(resource [1]gdclass.Resource) bool
	//Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
	Convert(resource [1]gdclass.Resource) [1]gdclass.Resource
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) ConvertsTo() (_ string)                                       { return }
func (self Implementation) Handles(resource [1]gdclass.Resource) (_ bool)                { return }
func (self Implementation) Convert(resource [1]gdclass.Resource) (_ [1]gdclass.Resource) { return }

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
func (Instance) _handles(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
*/
func (Instance) _convert(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) [1]gdclass.Resource) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
type class [1]gdclass.EditorResourceConversionPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourceConversionPlugin"))
	casted := Instance{*(*gdclass.EditorResourceConversionPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
func (class) _handles(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Takes an input [Resource] and converts it to the type given in [method _converts_to]. The returned [Resource] is the result of the conversion, and the input [Resource] remains unchanged.
*/
func (class) _convert(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) [1]gdclass.Resource) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_converts_to":
		return reflect.ValueOf(self._converts_to)
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_convert":
		return reflect.ValueOf(self._convert)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
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
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorResourceConversionPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorResourceConversionPlugin{*(*gdclass.EditorResourceConversionPlugin)(unsafe.Pointer(&ptr))}
	})
}
