//go:build !generate

package gd

import (
	"context"
	"reflect"
	"strings"

	"runtime.link/mmm"
)

type godotContext = mmm.ContextWith[API]

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
	godotContext
}

// NewContext returns a new [Context] with the given API, the [Context]
// will need to be freed manually by calling [Context.Free].
func NewContext(api *API) Context {
	return newContext(api)
}

func newContext(api *API) Context {
	var ctx Context
	ctx.godotContext = mmm.NewContextWith[API](context.Background(), api)
	return ctx
}

func Create[T PointerToClass](ctx Context, ptr T) T {
	object := ctx.API().ClassDB.ConstructObject(ctx, ctx.StringName(strings.TrimPrefix(reflect.TypeOf(ptr).Elem().Name(), "class")))
	ptr.SetPointer(mmm.Make[API, Pointer](ctx, ctx.API(), object.Pointer()))
	mmm.MarkFree(object.AsPointer())
	return ptr
}

// Version returns the version of the Godot API that we are linked in to.
func (ctx Context) Version() Version {
	return ctx.API().GetGodotVersion()
}

// GetLibraryPath returns the path to the library that was loaded.
func (ctx Context) GetLibraryPath() string {
	var godot = ctx.API()
	var frame = godot.NewFrame()
	ctx.API().ClassDB.GetLibraryPath(godot.ExtensionToken, frame.Back())
	var path = FrameGet[uintptr](frame)
	frame.Free()
	return mmm.Make[API, String](ctx, godot, path).String()
}

// String returns a [String] from a standard UTF8 Go string.
func (ctx Context) String(s string) String {
	return ctx.API().Strings.New(ctx, s)
}

// StringName returns a [StringName] from a standard UTF8 Go string.
func (ctx Context) StringName(s string) StringName {
	return ctx.API().StringNames.New(ctx, s)
}
