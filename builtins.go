package gd

import (
	"reflect"

	internal "grow.graphics/gd/internal"
)

// Class can be embedded inside of a struct to represent a new Class type.
// The extended class will be available by calling the [Class.Super] method.
type Class[T, S internal.IsClass] struct {
	internal.Class[T, S]
}

// Context for ownership and a reference to the Godot API, apart from
// its use as an ordinary [context.Context] to signal cancellation, this
// value is not safe to use concurrently. Each goroutine should create
// its own [Context] and use that instead.
//
//	newctx := gd.NewContext(oldctx.API())
//
// When a [Context] is freed, it will free all of the objects that were
// created using it. A [Context] should not be used after free, as it
// will be recycled and will cause values to be unexpectedly freed.
//
// When the context has been passed in as a function argument, always
// assume that the [Context] will be freed when the function returns.
// Classes can be moved between contexts using their [KeepAlive] method.
type Context = internal.Context

// Create a new instance of the given class, which should be an uninitialised
// pointer to a value of that class. T must be a class from this package.
func Create[T internal.PointerToClass](ctx Context, ptr T) T {
	return internal.Create[T](ctx, ptr)
}

// Const can be used to retrieve a 'constant' value from a structured type.
func Const[F func(T) T, T any](constant F) T {
	var zero T
	return constant(zero)
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T internal.IsClass](godot Context, class internal.IsClass) (T, bool) {
	return internal.As[T](godot, class)
}

type isResource interface {
	internal.IsClass

	AsResource() Resource
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
func Load[T isResource](godot Context, path string) (T, bool) {
	tmp := internal.NewContext(godot.API)
	defer tmp.End()
	hint := classNameOf(reflect.TypeOf([0]T{}).Elem())
	resource := ResourceLoader(godot).Load(godot,
		tmp.String(path), tmp.String(hint), ResourceLoaderCacheModeReuse)
	return As[T](godot, resource)
}
