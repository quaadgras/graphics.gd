package gd

import (
	"reflect"
	"unsafe"

	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/object"
	"grow.graphics/gd/object/Node"
	gdResourceLoader "grow.graphics/gd/object/ResourceLoader"
	"grow.graphics/gd/internal/mmm"
)

// Class can be embedded inside of a struct to represent a new Class type.
// The extended class will be available by calling the [Class.Super] method.
type Class[T, S gd.IsClass] struct {
	gd.Class[T, S]
}

// Lifetime for ownership and a reference to the Godot API, this
// value is not safe to use concurrently. Each goroutine should create
// its own [Lifetime] and use that instead.
//
//	newctx := gd.NewLifetime(other_lifetime)
//	defer newctx.End()
//
// When a [Lifetime] is freed, it will free all of the objects that were
// created using it. A [Lifetime] should not be used after free, as it
// will be recycled and will cause values to be unexpectedly freed.
//
// When the lifetime has been passed in as a function argument, always
// assume that the [Lifetime] will be freed when the function returns.
type Lifetime = gd.Lifetime

// Context is deprecated, use [Lifetime] instead.
type Context = Lifetime

// NewLifetime returns a new [Lifetime] with the given anchor lifetime
// always defer [Lifetime.Free] unless you know what you are doing.
func NewLifetime(anchor Lifetime) Lifetime {
	return gd.NewLifetime(anchor.API)
}

type isRefCounted interface {
	AsRefCounted() RefCounted
}

// New creates a new instance of the given class.
func New[T gd.IsClass](ctx Lifetime) *T {
	object := ctx.API.ClassDB.ConstructObject(ctx, ctx.StringName(classNameOf(reflect.TypeOf([0]T{}).Elem())))
	if native, ok := ctx.API.Instances[mmm.Get(object.AsPointer())[0]]; ok {
		native.SetPointer(object.AsPointer())
		native.SetTemporary(ctx)
		if rc, ok := As[RefCounted](ctx, object); ok {
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

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T gd.IsClass](godot Lifetime, class gd.IsClass) (T, bool) {
	return gd.As[T](godot, class)
}

type isResource interface {
	gd.IsClass

	AsResource() object.Resource
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
func Load[T isResource](godot Lifetime, path string) (T, bool) {
	tmp := NewLifetime(godot)
	defer tmp.End()
	hint := classNameOf(reflect.TypeOf([0]T{}).Elem())
	resource := gdResourceLoader.Expert(ResourceLoader(godot)).Load(godot,
		tmp.String(path), tmp.String(hint), gdResourceLoader.CacheModeReuse)
	return As[T](godot, resource[0])
}

// AddChild adds a child to the parent node, returning a [NodePath] to the child
// with the specified lifetime.
func AddChild(godot Lifetime, parent, child object.Node) NodePath {
	tmp := NewLifetime(godot)
	defer tmp.End()
	var adding object.Node
	adding[0].SetPointer(mmm.New[gd.Pointer](tmp.Lifetime, godot.API, mmm.Get(child[0].AsPointer())))
	Node.Expert(parent).AddChild(adding, true, 0)
	defer mmm.End(child[0].AsPointer())
	return Node.Expert(child).GetPath(godot)
}
