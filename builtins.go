package gd

import (
	"reflect"
	"unsafe"

	"grow.graphics/gd/gdclass"
	"grow.graphics/gd/gdclass/Node"
	"grow.graphics/gd/gdclass/ResourceLoader"
	gdResourceLoader "grow.graphics/gd/gdclass/ResourceLoader"
	gd "grow.graphics/gd/internal"
	classdb "grow.graphics/gd/internal/classdb"
	"grow.graphics/gd/internal/discreet"
)

// Class can be embedded inside of a struct to represent a new Class type.
// The extended class will be available by calling the [Class.Super] method.
type Class[T, S gd.IsClass] struct {
	gd.Class[T, S]
}

type isRefCounted interface {
	AsRefCounted() RefCounted
}

// New creates a new instance of the given class.
func New[T gd.IsClass]() *T {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName(classNameOf(reflect.TypeOf([0]T{}).Elem())))
	if native, ok := gd.Global.Instances[discreet.Get(object)[0]]; ok {
		native.SetPointer(object)
		if rc, ok := gdclass.As[RefCounted](object); ok {
			rc.Reference() // resources need to be referenced when we create them, as we will unreference them when they expire.
		}
		return any(native).(*T)
	}
	ptr := new(T)
	(*(*[1]gd.Object)(unsafe.Pointer(ptr)))[0] = object
	if rc, ok := any(ptr).(isRefCounted); ok {
		rc.AsRefCounted().Reference() // resources need to be referenced when we create them, as we will unreference them when they expire.
	}
	return ptr
}

// Const can be used to retrieve a 'constant' value from a structured type.
func Const[F func(T) T, T any](constant F) T {
	var zero T
	return constant(zero)
}

type isResource interface {
	gd.IsClass

	AsResource() gdclass.Resource
}

// Load returns a Resource from the filesystem located at the absolute path. Unless it's already
// referenced elsewhere (such as in another script or in the scene), the resource is loaded from
// disk on function call, which might cause a slight delay, especially when loading large scenes.
// To avoid unnecessary delays when loading something multiple times, either store the resource
// in a variable or use preload.
//
// Note: Resource paths can be obtained by right-clicking on a resource in the FileSystem dock
// and choosing "Copy Path", or by dragging the file from the FileSystem dock into the current
// script.
//
// Important: The path must be absolute. A relative path will always return null.
//
// This function is a simplified version of [ResourceLoader(godot).Load], which can be used for
// more advanced scenarios.
//
// Note: Files have to be imported into the engine first to load them using this function. If you
// want to load Images at run-time, you may use Image.load. If you want to import audio files,
// you can use the snippet described in AudioStreamMP3.data.
//
// Note: If ProjectSettings.editor/export/convert_text_resources_to_binary is true, load will not
// be able to read converted files in an exported project. If you rely on run-time loading of files
// present within the PCK, set ProjectSettings.editor/export/convert_text_resources_to_binary to false.
func Load[T isResource](path string) (T, bool) {
	hint := classNameOf(reflect.TypeOf([0]T{}).Elem())
	resource := ResourceLoader.GD().Load(
		gd.NewString(path), gd.NewString(hint), gdResourceLoader.CacheModeReuse)
	return gdclass.As[T](resource[0])
}

// AddChild adds a child to the parent node, returning a [NodePath] to the child
// with the specified lifetime.
func AddChild(parent, child gdclass.Node) NodePath {
	var adding gdclass.Node
	adding[0] = discreet.New[classdb.Node](discreet.Get(child[0]))
	Node.GD(parent).AddChild(adding, true, 0)
	defer discreet.End(child[0])
	return Node.GD(child).GetPath()
}
