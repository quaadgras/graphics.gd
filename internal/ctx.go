//go:build !generate

package gd

import (
	"reflect"
	"strings"

	"runtime.link/mmm"
)

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
type Context struct {
	mmm.Lifetime

	API *API
}

// NewContext returns a new [Context] with the given API, the [Context]
// will need to be freed manually by calling [Context.Free].
func NewContext(api *API) Context {
	return newContext(api)
}

func newContext(api *API) Context {
	var ctx Context
	ctx.Lifetime = mmm.NewLifetime(api)
	ctx.API = api
	return ctx
}

func Create[T PointerToClass](ctx Context, ptr T) T {
	object := ctx.API.ClassDB.ConstructObject(ctx, ctx.StringName(strings.TrimPrefix(reflect.TypeOf(ptr).Elem().Name(), "class")))
	if native, ok := ctx.API.Instances[mmm.Get(object.AsPointer())]; ok {
		return native.(T)
	}
	ptr.SetPointer(mmm.Let[Pointer](ctx.Lifetime, ctx.API, mmm.Get(object.AsPointer())))
	return ptr
}

// Version returns the version of the Godot API that we are linked in to.
func (godot Context) Version() Version {
	return godot.API.GetGodotVersion()
}

// GetLibraryPath returns the path to the library that was loaded.
func (godot Context) GetLibraryPath() string {
	return godot.API.GetLibraryPath(godot, godot.API.ExtensionToken).String()
}

// String returns a [String] from a standard UTF8 Go string.
func (godot Context) String(s string) String {
	return godot.API.Strings.New(godot, s)
}

// StringName returns a [StringName] from a standard UTF8 Go string.
func (godot Context) StringName(s string) StringName {
	return godot.API.StringNames.New(godot, s)
}
